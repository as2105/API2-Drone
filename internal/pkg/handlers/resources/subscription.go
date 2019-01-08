package resources

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/logging"
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/storage/database"
	"github.com/SynapticHealthAlliance/fhir-api/pkg/models"
	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
)

type subscriptionDB struct {
	UUID      string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Data      []byte
	Active    bool
}

// Subscription ...
type Subscription struct {
	config        *ResourceConfig
	db            *database.DB
	jsonValidator *models.JSONValidator
	log           *logging.Logger
	renderer      *render.Render
}

// Create ...
func (h *Subscription) Create() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		newSub := &models.Subscription{}
		err := loadResourceFromBody(newSub, req, h.jsonValidator)
		if err != nil {
			h.log.WithError(err).Panic("failed to load resource")
		}
		if newSub.ID == "" {
			newSub.ID = uuid.NewUUID().String()
		}
		now := time.Now().UTC()
		if newSub.Meta == nil {
			newSub.Meta = &models.Meta{}
		}
		newSub.Meta.VersionID = ""
		newSub.Meta.LastUpdated = now.Format(time.RFC3339)
		if newSub.Status == "" || newSub.Status == "active" {
			newSub.Status = "requested" // activate subscription later
		}
		jsonBytes, err := json.Marshal(newSub)
		if err != nil {
			h.log.WithError(err).Panic("failed to marshal object as JSON")
		}
		newDBRec := &subscriptionDB{
			UUID:      newSub.ID,
			Data:      jsonBytes,
			CreatedAt: now,
			UpdatedAt: now,
		}
		if err := h.db.Create(newDBRec).Error; err != nil {
			h.log.WithError(err).Panic("failed to save object to database")
		}
		resourceCreated(h.renderer, rw, uuid.Parse(newSub.ID), "", now, newSub)
	})
}

// Read ...
func (h *Subscription) Read() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		uuid := vars["resourceID"]
		if uuid == "" {
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}
		dbRec := &subscriptionDB{}
		query := h.db.Unscoped().Where(&subscriptionDB{UUID: uuid}).First(dbRec)
		if query.RecordNotFound() {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else if err := query.Error; err != nil {
			h.log.WithError(err).Panic("failed to query database")
		}
		if dbRec.DeletedAt != nil {
			rw.WriteHeader(http.StatusGone)
			return
		}
		newSub := &models.Subscription{}
		if err := json.Unmarshal(dbRec.Data, newSub); err != nil {
			h.log.WithError(err).Panic("failed to unmarshal data")
		}
		resourceRead(h.renderer, rw, req, http.StatusOK, "", dbRec.UpdatedAt, newSub, true)
	})
}

// Update ...
func (h *Subscription) Update() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resourceID, err := getResourceID(req)
		if err != nil {
			h.log.WithError(err).Error("invalid resource ID provided")
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}

		newSub := &models.Subscription{}
		if err := loadResourceFromBody(newSub, req, h.jsonValidator); err != nil {
			h.log.WithError(err).Panic("failed to load resource")
		}
		if newSub.ID != resourceID.String() {
			h.log.Errorf("resourceID provided (%q) does not match ID inside of document (%q)", resourceID.String(), newSub.ID)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		scope := h.db.Unscoped()
		dbRec := &subscriptionDB{}
		query := scope.Where(&subscriptionDB{UUID: resourceID.String()}).First(dbRec)
		if query.RecordNotFound() {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else if err := query.Error; err != nil {
			h.log.WithError(err).Panic("failed to query database")
		}
		oldSub := &models.Subscription{}
		if err := json.Unmarshal(dbRec.Data, oldSub); err != nil {
			h.log.WithError(err).Panic("failed to unmarshal data")
		}

		now := time.Now().UTC()

		status := http.StatusOK
		if dbRec.DeletedAt != nil {
			status = http.StatusCreated
		}

		if newSub.Meta == nil {
			newSub.Meta = &models.Meta{}
		}
		newSub.Meta.VersionID = ""
		newSub.Meta.LastUpdated = now.Format(time.RFC3339)
		if newSub.Status == "" || newSub.Status == "active" {
			newSub.Status = "requested" // activate subscription later
		}
		jsonBytes, err := json.Marshal(newSub)
		if err != nil {
			h.log.WithError(err).Panic("failed to marshal object as JSON")
		}

		dbRec.UpdatedAt = now
		dbRec.Data = jsonBytes
		dbRec.DeletedAt = nil
		dbRec.Active = false

		if err := scope.Save(dbRec).Error; err != nil {
			h.log.WithError(err).Panic("failed to update record in database")
		}

		resourceRead(h.renderer, rw, req, status, "", dbRec.UpdatedAt, newSub, false)
	})
}

// Delete ...
func (h *Subscription) Delete() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		uuid := vars["resourceID"]
		if uuid == "" {
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}
		if err := h.db.Where(&subscriptionDB{UUID: uuid}).Delete(subscriptionDB{}).Error; err != nil {
			h.log.WithError(err).Panic("failed to delete record")
		}
		rw.WriteHeader(http.StatusNoContent)
	})
}

// Validate ...
func (h *Subscription) Validate() http.Handler {
	return validateJSONResource(h.log, h.renderer, h.jsonValidator)
}

// GetResourceConfig ...
func (h *Subscription) GetResourceConfig() *ResourceConfig {
	return h.config
}

// NewSubscription ...
func NewSubscription(registry *Registry) (*Subscription, error) {
	v, err := models.NewJSONValidator(registry.box, "Subscription")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}
	registry.db.AutoMigrate(&subscriptionDB{})
	config := NewResourceConfig()
	return &Subscription{
		config:        config,
		db:            registry.db,
		jsonValidator: v,
		log:           registry.log,
		renderer:      registry.renderer,
	}, nil
}
