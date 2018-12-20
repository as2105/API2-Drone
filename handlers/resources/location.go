package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
)

// Location ...
type Location struct{}

// Create ...
func (l *Location) Create(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Read ...
func (l *Location) Read(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Update ...
func (l *Location) Update(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Delete ...
func (l *Location) Delete(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Search ...
func (l *Location) Search(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

func (r *Practitioner) SearchIncludes() SearchIncludes {
	return SearchIncludes{"*"}
}

func (r *Practitioner) SearchParams() []SearchParam {
	return []SearchParam{
		{Name: "_id", Type: models.SearchParamTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParamTypeDate},
		{Name: "identifier", Type: models.SearchParamTypeToken},
		{Name: "status", Type: models.SearchParamTypeToken},
		{Name: "type", Type: models.SearchParamTypeToken},
	}
}
