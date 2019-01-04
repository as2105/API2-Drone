package resources

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"net/http"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/config"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/SynapticHealthAlliance/fhir-api/storage/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gobuffalo/packr/v2"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
)

// Practitioner ...
type Practitioner struct {
	jsonValidator *models.JSONValidator
	adapter       *ethereum.Adapter
	config        *ResourceConfig
}

// Create ...
func (r *Practitioner) Create(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		p := models.NewPractitioner()

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
		p := models.NewPractitioner()
		if err := r.adapter.ReadJSONResource(req.Context(), resourceID, p); err != nil {
			log.WithError(err).Panic("failed to read record")
		}
		// TODO: support deleted records, versioning
		resourceRead(rndr, rw, http.StatusOK, p.Meta.VersionID, p.Meta.LastUpdated, p)
	})
}

// Update ...
func (r *Practitioner) Update(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
}

// Delete ...
func (r *Practitioner) Delete(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
}

// Search ...
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
	rConfig.SearchIncludes = searchIncludes{
		"Practitioner:location",
		"*",
	}
	rConfig.SearchParams = []searchParam{
		{Name: "_id", Type: models.SearchParamTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParamTypeDate},
		{Name: "active", Type: models.SearchParamTypeToken},
		{Name: "identifier", Type: models.SearchParamTypeToken},
		{Name: "location", Type: models.SearchParamTypeReference},
		{Name: "name", Type: models.SearchParamTypeString},
		{Name: "telecom", Type: models.SearchParamTypeToken},
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
