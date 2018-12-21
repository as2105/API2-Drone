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

func (r *Practitioner) SearchIncludes() SearchIncludes {
	return SearchIncludes{
		"Practitioner:location",
		"*",
	}
}

func (r *Practitioner) SearchParams() []SearchParam {
	return []SearchParam{
		{Name: "_id", Type: models.SearchParamTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParamTypeDate},
		{Name: "active", Type: models.SearchParamTypeToken},
		{Name: "identifier", Type: models.SearchParamTypeToken},
		{Name: "location", Type: models.SearchParamTypeReference},
		{Name: "name", Type: models.SearchParamTypeString},
		{Name: "telecom", Type: models.SearchParamTypeToken},
	}
}

// Validate ...
func (r *Practitioner) Validate(log *logging.Logger, rndr *render.Render) http.Handler {
	return validateJSONResource(log, rndr, r.jsonValidator)
}

func NewPractitioner(box *packr.Box) (*Practitioner, error) {
	v, err := models.NewJSONValidator(box, "Practitioner")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}
	return &Practitioner{jsonValidator: v}, nil
}
