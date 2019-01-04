package resources

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/SynapticHealthAlliance/fhir-api/storage/database"
	"github.com/gobuffalo/packr/v2"
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
	jsonValidator *models.JSONValidator
	db            *database.DB
	config        *ResourceConfig
}

// Create ...
func (s *Subscription) Create(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		newSub := models.NewSubscription()
		err := loadResourceFromBody(newSub, req, s.jsonValidator)
		if err != nil {
			log.WithError(err).Panic("failed to load resource")
		}

		newUUID := uuid.NewUUID()
		now := time.Now().UTC()

		newSub.ID = newUUID.String()
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
			log.WithError(err).Panic("failed to marshal object as JSON")
		}

		newDBRec := &subscriptionDB{
			UUID:      newSub.ID,
			Data:      jsonBytes,
			CreatedAt: now,
			UpdatedAt: now,
		}
		if err := s.db.Create(newDBRec).Error; err != nil {
			log.WithError(err).Panic("failed to save object to database")
		}

		resourceCreated(rndr, rw, newUUID, "", now, newSub)
	})
}

// Read ...
func (s *Subscription) Read(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		uuid := vars["resourceID"]
		if uuid == "" {
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}
		dbRec := &subscriptionDB{}
		query := s.db.Unscoped().Where(&subscriptionDB{UUID: uuid}).First(dbRec)
		if query.RecordNotFound() {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else if err := query.Error; err != nil {
			log.WithError(err).Panic("failed to query database")
		}
		if dbRec.DeletedAt != nil {
			rw.WriteHeader(http.StatusGone)
			return
		}
		newSub := models.NewSubscription()
		if err := json.Unmarshal(dbRec.Data, newSub); err != nil {
			log.WithError(err).Panic("failed to unmarshal data")
		}
		resourceRead(rndr, rw, http.StatusOK, "", dbRec.UpdatedAt, newSub)
	})
}

// Update ...
func (s *Subscription) Update(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resourceID, err := getResourceID(req)
		if err != nil {
			log.WithError(err).Error("invalid resource ID provided")
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}

		newSub := models.NewSubscription()
		if err := loadResourceFromBody(newSub, req, s.jsonValidator); err != nil {
			log.WithError(err).Panic("failed to load resource")
		}
		if newSub.ID != resourceID.String() {
			log.Errorf("resourceID provided (%q) does not match ID inside of document (%q)", resourceID.String(), newSub.ID)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		scope := s.db.Unscoped()
		dbRec := &subscriptionDB{}
		query := scope.Where(&subscriptionDB{UUID: resourceID.String()}).First(dbRec)
		if query.RecordNotFound() {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else if err := query.Error; err != nil {
			log.WithError(err).Panic("failed to query database")
		}
		oldSub := models.NewSubscription()
		if err := json.Unmarshal(dbRec.Data, oldSub); err != nil {
			log.WithError(err).Panic("failed to unmarshal data")
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
			log.WithError(err).Panic("failed to marshal object as JSON")
		}

		dbRec.UpdatedAt = now
		dbRec.Data = jsonBytes
		dbRec.DeletedAt = nil
		dbRec.Active = false

		if err := scope.Save(dbRec).Error; err != nil {
			log.WithError(err).Panic("failed to update record in database")
		}

		resourceRead(rndr, rw, status, "", dbRec.UpdatedAt, newSub)
	})
}

// Delete ...
func (s *Subscription) Delete(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		uuid := vars["resourceID"]
		if uuid == "" {
			rw.WriteHeader(http.StatusBadRequest) // TODO: More verbose errors?
			return
		}
		if err := s.db.Where(&subscriptionDB{UUID: uuid}).Delete(subscriptionDB{}).Error; err != nil {
			log.WithError(err).Panic("failed to delete record")
		}
		rw.WriteHeader(http.StatusNoContent)
	})
}

// Validate ...
func (s *Subscription) Validate(log *logging.Logger, rndr *render.Render) http.Handler {
	return validateJSONResource(log, rndr, s.jsonValidator)
}

// GetResourceConfig ...
func (s *Subscription) GetResourceConfig() *ResourceConfig {
	return s.config
}

// NewSubscription ...
func NewSubscription(box *packr.Box, db *database.DB) (*Subscription, error) {
	v, err := models.NewJSONValidator(box, "Subscription")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}
	db.AutoMigrate(&subscriptionDB{})
	config := NewResourceConfig()
	return &Subscription{jsonValidator: v, db: db, config: config}, nil
}
