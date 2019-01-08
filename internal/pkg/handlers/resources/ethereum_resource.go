package resources

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/logging"
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/storage/ethereum"
	"github.com/SynapticHealthAlliance/fhir-api/pkg/models"
	"github.com/pborman/uuid"
	"github.com/unrolled/render"
)

// EthereumResource provides a standard set of handlers for Ethereum-backed resources
type EthereumResource struct {
	adapter       *ethereum.Adapter
	config        *ResourceConfig
	jsonValidator *models.JSONValidator
	log           *logging.Logger
	newModelFunc  func() models.Resource
	renderer      *render.Render
}

// Create ...
func (h *EthereumResource) Create() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resource := h.newModelFunc()

		err := loadResourceFromBody(resource, req, h.jsonValidator)
		if err != nil {
			h.log.WithError(err).Panic("failed to load resource")
		}

		newUUID := uuid.NewUUID()
		now := time.Now().UTC()

		resource.SetID(newUUID.String())
		meta := resource.GetMeta()
		if meta == nil {
			meta = &models.Meta{}
		}
		meta.VersionID = initialResourceVersionID
		meta.LastUpdated = now.Format(time.RFC3339)
		resource.SetMeta(meta)

		jsonBytes, err := json.Marshal(resource)
		if err != nil {
			h.log.WithError(err).Panic("failed to marshal object as JSON")
		}

		elemData := ethereum.NewObjectCollectionElementFHIRJSONData(jsonBytes)
		if err := h.adapter.Create(req.Context(), newUUID, elemData); err != nil {
			h.log.WithError(err).Panic("failed to save object to smart contract")
		}

		resourceCreated(h.renderer, rw, newUUID, meta.VersionID, now, resource)
	})
}

// Read ...
func (h *EthereumResource) Read() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resourceID, err := getResourceID(req)
		if err != nil {
			h.log.WithError(err).Error("invalid resource ID provided")
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}
		resource := h.newModelFunc()
		if err := h.adapter.ReadJSONResource(req.Context(), resourceID, resource); err != nil {
			h.log.WithError(err).Panic("failed to read record")
		}
		// TODO: support deleted records, versioning
		meta := resource.GetMeta()
		resourceRead(h.renderer, rw, req, http.StatusOK, meta.VersionID, meta.LastUpdated, resource, true)
	})
}

// Update ...
func (h *EthereumResource) Update() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resourceID, err := getResourceID(req)
		if err != nil {
			h.log.WithError(err).Error("invalid resource ID provided")
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}

		oldResource := h.newModelFunc()
		if err := h.adapter.ReadJSONResource(req.Context(), resourceID, oldResource); err != nil {
			h.log.WithError(err).Panic("failed to read record")
		}
		oldMeta := oldResource.GetMeta()
		oldTime, err := time.Parse(time.RFC3339, oldMeta.LastUpdated)
		if err != nil {
			h.log.WithError(err).Panic("failed to parse last updated timestamp on original resource")
		}

		// optimistic locking
		expected := req.Header.Get("If-Match")
		if expected == "" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		if expected != generateETag(oldMeta.VersionID) {
			rw.WriteHeader(http.StatusPreconditionFailed)
			return
		}

		newResource := h.newModelFunc()
		if err := loadResourceFromBody(newResource, req, h.jsonValidator); err != nil {
			h.log.WithError(err).Panic("failed to load resource")
		}

		now := time.Now().UTC()
		uCount, err := getUpdateCountFromVersionID(oldMeta.VersionID)
		if err != nil {
			h.log.WithError(err).Panic("failed to get current update count")
		}
		curBlockNumber, err := h.adapter.CurrentBlock(req.Context())
		if err != nil {
			h.log.WithError(err).Panic("failed to get current block number")
		}
		newVersionID := generateResourceVersionID(uCount+1, curBlockNumber)

		newMeta := newResource.GetMeta()
		if newMeta == nil {
			newMeta = &models.Meta{}
		}
		newMeta.VersionID = newVersionID
		newMeta.LastUpdated = now.Format(time.RFC3339)
		newResource.SetMeta(newMeta)

		jsonBytes, err := json.Marshal(newResource)
		if err != nil {
			h.log.WithError(err).Panic("failed to marshal object as JSON")
		}

		changeScore := uint8(0) // TODO

		elemData := ethereum.NewObjectCollectionElementFHIRJSONData(jsonBytes)
		if err := h.adapter.Update(req.Context(), resourceID, oldTime, elemData, changeScore); err != nil {
			h.log.WithError(err).Panic("failed to save object to smart contract")
		}

		// TODO: support upsert

		// TODO: Return 409 if contract failed due to timestamp change

		resourceRead(h.renderer, rw, req, http.StatusOK, newMeta.VersionID, now, newResource, false)
	})
}

// Delete ...
func (h *EthereumResource) Delete() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resourceID, err := getResourceID(req)
		if err != nil {
			h.log.WithError(err).Error("invalid resource ID provided")
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}
		if err := h.adapter.Destroy(req.Context(), resourceID); err != nil {
			h.log.WithError(err).Panic("failed to destroy object")
		}
		rw.WriteHeader(http.StatusNoContent)
	})
}

// Validate ...
func (h *EthereumResource) Validate() http.Handler {
	return validateJSONResource(h.log, h.renderer, h.jsonValidator)
}

// GetResourceConfig ...
func (h *EthereumResource) GetResourceConfig() *ResourceConfig {
	return h.config
}
