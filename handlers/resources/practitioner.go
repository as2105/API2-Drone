package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
)

// Practitioner ...
type Practitioner struct {
	jsonValidator *models.JSONValidator
	config        *ResourceConfig
}

// Create ...
func (r *Practitioner) Create(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
}

// Read ...
func (r *Practitioner) Read(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
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
func NewPractitioner(box *packr.Box) (*Practitioner, error) {
	v, err := models.NewJSONValidator(box, "Practitioner")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}

	config := NewResourceConfig()
	config.SearchIncludes = searchIncludes{
		"Practitioner:location",
		"*",
	}
	config.SearchParams = []searchParam{
		{Name: "_id", Type: models.SearchParameterTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParameterTypeDate},
		{Name: "active", Type: models.SearchParameterTypeToken},
		{Name: "identifier", Type: models.SearchParameterTypeToken},
		{Name: "location", Type: models.SearchParameterTypeReference},
		{Name: "name", Type: models.SearchParameterTypeString},
		{Name: "telecom", Type: models.SearchParameterTypeToken},
	}

	return &Practitioner{jsonValidator: v, config: config}, nil
}
