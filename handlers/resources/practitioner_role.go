package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
)

// PractitionerRole ...
type PractitionerRole struct {
	jsonValidator *models.JSONValidator
	config        *ResourceConfig
}

// Create ...
func (r *PractitionerRole) Create(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Read ...
func (r *PractitionerRole) Read(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Update ...
func (r *PractitionerRole) Update(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Delete ...
func (r *PractitionerRole) Delete(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Search ...
func (r *PractitionerRole) Search(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Validate ...
func (r *PractitionerRole) Validate(log *logging.Logger, rndr *render.Render) http.Handler {
	return validateJSONResource(log, rndr, r.jsonValidator)
}

// GetResourceConfig ...
func (r *PractitionerRole) GetResourceConfig() *ResourceConfig {
	return r.config
}

// NewPractitionerRole ...
func NewPractitionerRole(box *packr.Box) (*PractitionerRole, error) {
	v, err := models.NewJSONValidator(box, "PractitionerRole")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}

	config := NewResourceConfig()
	config.SearchIncludes = searchIncludes{
		"PractitionerRole:location",
		"PractitionerRole:practitioner",
		"*",
	}
	config.SearchParams = []searchParam{
		{Name: "_id", Type: models.SearchParameterTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParameterTypeDate},
		{Name: "identifier", Type: models.SearchParameterTypeToken},
		{Name: "location", Type: models.SearchParameterTypeReference},
		{Name: "practitioner", Type: models.SearchParameterTypeReference},
		{Name: "telecom", Type: models.SearchParameterTypeToken},
	}

	return &PractitionerRole{jsonValidator: v, config: config}, nil
}
