package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
)

// PractitionerRole ...
type PractitionerRole struct{}

// Create ...
func (r *PractitionerRole) Create(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Read ...
func (r *PractitionerRole) Read(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Update ...
func (r *PractitionerRole) Update(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Delete ...
func (r *PractitionerRole) Delete(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Search ...
func (r *PractitionerRole) Search(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}
