package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
)

// Practitioner ...
type Practitioner struct{}

// Create ...
func (r *Practitioner) Create(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Read ...
func (r *Practitioner) Read(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Update ...
func (r *Practitioner) Update(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Delete ...
func (r *Practitioner) Delete(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Search ...
func (r *Practitioner) Search(log *logging.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}
