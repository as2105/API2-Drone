package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/storage/ethereum"
	"github.com/unrolled/render"
)

// ConfiguredResource ...
type ConfiguredResource interface {
	GetResourceConfig() *ResourceConfig
}

// CreateableResource ...
type CreateableResource interface {
	Create(log *logging.Logger, rndr *render.Render) http.Handler
}

// ReadableResource ...
type ReadableResource interface {
	Read(log *logging.Logger, rndr *render.Render) http.Handler
}

// VersionReadableResource ...
type VersionReadableResource interface {
	VersionRead(log *logging.Logger, rndr *render.Render) http.Handler
}

// UpdateableResource ...
type UpdateableResource interface {
	Update(log *logging.Logger, rndr *render.Render) http.Handler
}

// DeleteableResource ...
type DeleteableResource interface {
	Delete(log *logging.Logger, rndr *render.Render) http.Handler
}

// SearchableResource ...
type SearchableResource interface {
	Search(log *logging.Logger, rndr *render.Render) http.Handler
}

// ValidateableResource ...
type ValidateableResource interface {
	Validate(log *logging.Logger, rndr *render.Render) http.Handler
}

// PatchableResource ...
type PatchableResource interface {
	Patch(log *logging.Logger, rndr *render.Render) http.Handler
}

type ethereumResource interface {
	getAdapter() *ethereum.Adapter
	getModel() interface{}
}
