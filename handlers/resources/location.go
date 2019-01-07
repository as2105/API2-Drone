package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
)

// Location ...
type Location struct {
	jsonValidator *models.JSONValidator
	config        *ResourceConfig
}

// Create ...
func (l *Location) Create(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Read ...
func (l *Location) Read(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Update ...
func (l *Location) Update(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Delete ...
func (l *Location) Delete(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Search ...
func (l *Location) Search(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Validate ...
func (l *Location) Validate(log *logging.Logger, rndr *render.Render) http.Handler {
	return validateJSONResource(log, rndr, l.jsonValidator)
}

// GetResourceConfig ...
func (l *Location) GetResourceConfig() *ResourceConfig {
	return l.config
}

// NewLocation ...
func NewLocation(box *packr.Box) (*Location, error) {
	v, err := models.NewJSONValidator(box, "Location")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}

	config := NewResourceConfig()
	config.SearchIncludes = searchIncludes{"*"}
	config.SearchParams = []searchParam{
		{Name: "_id", Type: models.SearchParameterTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParameterTypeDate},
		{Name: "identifier", Type: models.SearchParameterTypeToken},
		{Name: "status", Type: models.SearchParameterTypeToken},
		{Name: "type", Type: models.SearchParameterTypeToken},
	}

	return &Location{
		jsonValidator: v,
		config:        config,
	}, nil
}
