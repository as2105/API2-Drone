package resources

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/config"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/SynapticHealthAlliance/fhir-api/storage/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gobuffalo/packr/v2"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
)

// Practitioner ...
// TODO: This entire struct can be refactored to share code with other resource handlers
type Practitioner struct {
	jsonValidator *models.JSONValidator
	adapter       *ethereum.Adapter
	config        *ResourceConfig
}

// Create ...
func (r *Practitioner) Create(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		p := &models.Practitioner{}

		err := loadResourceFromBody(p, req, r.jsonValidator)
		if err != nil {
			log.WithError(err).Panic("failed to load resource")
		}

		newUUID := uuid.NewUUID()
		now := time.Now().UTC()

		p.ID = newUUID.String()
		if p.Meta == nil {
			p.Meta = &models.Meta{}
		}
		p.Meta.VersionID = initialResourceVersionID
		p.Meta.LastUpdated = now.Format(time.RFC3339)

		jsonBytes, err := json.Marshal(p)
		if err != nil {
			log.WithError(err).Panic("failed to marshal object as JSON")
		}

		elemData := ethereum.NewObjectCollectionElementFHIRJSONData(jsonBytes)
		if err := r.adapter.Create(req.Context(), newUUID, elemData); err != nil {
			log.WithError(err).Panic("failed to save object to smart contract")
		}

		resourceCreated(rndr, rw, newUUID, p.Meta.VersionID, now, p)
	})
}

// Read ...
func (r *Practitioner) Read(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resourceID, err := getResourceID(req)
		if err != nil {
			log.WithError(err).Error("invalid resource ID provided")
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}
		p := &models.Practitioner{}
		if err := r.adapter.ReadJSONResource(req.Context(), resourceID, p); err != nil {
			log.WithError(err).Panic("failed to read record")
		}
		// TODO: support deleted records, versioning
		resourceRead(rndr, rw, req, http.StatusOK, p.Meta.VersionID, p.Meta.LastUpdated, p, true)
	})
}

// Update ...
func (r *Practitioner) Update(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resourceID, err := getResourceID(req)
		if err != nil {
			log.WithError(err).Error("invalid resource ID provided")
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}

		oldP := &models.Practitioner{}
		if err := r.adapter.ReadJSONResource(req.Context(), resourceID, oldP); err != nil {
			log.WithError(err).Panic("failed to read record")
		}
		oldTime, err := time.Parse(time.RFC3339, oldP.Meta.LastUpdated)
		if err != nil {
			log.WithError(err).Panic("failed to parse last updated timestamp on original resource")
		}

		// optimistic locking
		expected := req.Header.Get("If-Match")
		if expected == "" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		if expected != generateETag(oldP.Meta.VersionID) {
			rw.WriteHeader(http.StatusPreconditionFailed)
			return
		}

		newP := &models.Practitioner{}
		if err := loadResourceFromBody(newP, req, r.jsonValidator); err != nil {
			log.WithError(err).Panic("failed to load resource")
		}

		now := time.Now().UTC()
		uCount, err := getUpdateCountFromVersionID(oldP.Meta.VersionID)
		if err != nil {
			log.WithError(err).Panic("failed to get current update count")
		}
		curBlockNumber, err := r.adapter.CurrentBlock(req.Context())
		if err != nil {
			log.WithError(err).Panic("failed to get current block number")
		}
		newVersionID := generateResourceVersionID(uCount+1, curBlockNumber)

		if newP.Meta == nil {
			newP.Meta = &models.Meta{}
		}
		newP.Meta.VersionID = newVersionID
		newP.Meta.LastUpdated = now.Format(time.RFC3339)

		jsonBytes, err := json.Marshal(newP)
		if err != nil {
			log.WithError(err).Panic("failed to marshal object as JSON")
		}

		changeScore := uint8(0) // TODO

		elemData := ethereum.NewObjectCollectionElementFHIRJSONData(jsonBytes)
		if err := r.adapter.Update(req.Context(), resourceID, oldTime, elemData, changeScore); err != nil {
			log.WithError(err).Panic("failed to save object to smart contract")
		}

		// TODO: support upsert

		// TODO: Return 409 if contract failed due to timestamp change

		resourceRead(rndr, rw, req, http.StatusOK, newP.Meta.VersionID, now, newP, false)
	})
}

// Delete ...
func (r *Practitioner) Delete(log *logging.Logger, rndr *render.Render) http.Handler {
	return ethereumResourceDestroy(log, rndr, r)
}

// Search ... (TODO)
func (r *Practitioner) Search(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
}

// Validate ...
func (r *Practitioner) Validate(log *logging.Logger, rndr *render.Render) http.Handler {
	return validateJSONResource(log, rndr, r.jsonValidator)
}

// GetResourceConfig ...
func (r *Practitioner) GetResourceConfig() *ResourceConfig {
	return r.config
}

// GetAdapter ...
func (r *Practitioner) getAdapter() *ethereum.Adapter {
	return r.adapter
}

func (r *Practitioner) getModel() interface{} {
	return &models.Practitioner{}
}

// NewPractitioner ...
func NewPractitioner(
	box *packr.Box,
	connection *ethclient.Client,
	transactOpts *bind.TransactOpts,
	appConfig *config.Config,
	log *logging.Logger,
	txnsChan ethereum.TransactionsChannel,
) (*Practitioner, error) {
	v, err := models.NewJSONValidator(box, "Practitioner")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}

	rConfig := NewResourceConfig()
	rConfig.Versioning = models.CapabilityStatementResourceVersioningVersionedUpdate
	rConfig.SearchIncludes = searchIncludes{
		"Practitioner:location",
		"*",
	}
	rConfig.SearchParams = []searchParam{
		{Name: "_id", Type: models.SearchParameterTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParameterTypeDate},
		{Name: "active", Type: models.SearchParameterTypeToken},
		{Name: "identifier", Type: models.SearchParameterTypeToken},
		{Name: "location", Type: models.SearchParameterTypeReference},
		{Name: "name", Type: models.SearchParameterTypeString},
		{Name: "telecom", Type: models.SearchParameterTypeToken},
	}

	collContract := appConfig.ObjectCollectionContracts["Practitioner"]
	if collContract == nil {
		return nil, errors.New("no collection contract found")
	}

	adapter, err := ethereum.NewAdapter(
		connection,
		transactOpts,
		appConfig.OrganizationContract,
		collContract,
		txnsChan,
		log,
	)

	return &Practitioner{
		jsonValidator: v,
		config:        rConfig,
		adapter:       adapter,
	}, nil
}
