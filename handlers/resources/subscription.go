package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
)

// Subscription ...
type Subscription struct {
	jsonValidator *models.JSONValidator
}

// Create ...
func (s *Subscription) Create(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
}

// Read ...
func (s *Subscription) Read(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
}

// Update ...
func (s *Subscription) Update(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
}

// Delete ...
func (s *Subscription) Delete(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
}

// Validate ...
func (s *Subscription) Validate(log *logging.Logger, rndr *render.Render) http.Handler {
	return validateJSONResource(log, rndr, s.jsonValidator)
}

func NewSubscription(box *packr.Box) (*Subscription, error) {
	v, err := models.NewJSONValidator(box, "Subscription")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}
	return &Subscription{jsonValidator: v}, nil
}
