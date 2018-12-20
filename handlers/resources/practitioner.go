package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/unrolled/render"
)

// Practitioner ...
type Practitioner struct{}

// Create ...
func (r *Practitioner) Create(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Read ...
func (r *Practitioner) Read(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Update ...
func (r *Practitioner) Update(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Delete ...
func (r *Practitioner) Delete(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}

// Search ...
func (r *Practitioner) Search(log *logging.Logger, rndr *render.Render) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})
}
