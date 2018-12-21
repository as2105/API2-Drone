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

func (r *PractitionerRole) SearchIncludes() SearchIncludes {
	return SearchIncludes{
		"PractitionerRole:location",
		"PractitionerRole:practitioner",
		"*",
	}
}

func (r *PractitionerRole) SearchParams() []SearchParam {
	return []SearchParam{
		{Name: "_id", Type: models.SearchParamTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParamTypeDate},
		{Name: "identifier", Type: models.SearchParamTypeToken},
		{Name: "location", Type: models.SearchParamTypeReference},
		{Name: "practitioner", Type: models.SearchParamTypeReference},
		{Name: "telecom", Type: models.SearchParamTypeToken},
	}
}

// Validate ...
func (r *PractitionerRole) Validate(log *logging.Logger, rndr *render.Render) http.Handler {
	return validateJSONResource(log, rndr, r.jsonValidator)
}

func NewPractitionerRole(box *packr.Box) (*PractitionerRole, error) {
	v, err := models.NewJSONValidator(box, "PractitionerRole")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}
	return &PractitionerRole{jsonValidator: v}, nil
}
