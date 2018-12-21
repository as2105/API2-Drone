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

func (l *Location) SearchIncludes() SearchIncludes {
	return SearchIncludes{"*"}
}

func (l *Location) SearchParams() []SearchParam {
	return []SearchParam{
		{Name: "_id", Type: models.SearchParamTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParamTypeDate},
		{Name: "identifier", Type: models.SearchParamTypeToken},
		{Name: "status", Type: models.SearchParamTypeToken},
		{Name: "type", Type: models.SearchParamTypeToken},
	}
}

// Validate ...
func (l *Location) Validate(log *logging.Logger, rndr *render.Render) http.Handler {
	return validateJSONResource(log, rndr, l.jsonValidator)
}

func NewLocation(box *packr.Box) (*Location, error) {
	v, err := models.NewJSONValidator(box, "Location")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}
	return &Location{jsonValidator: v}, nil
}
