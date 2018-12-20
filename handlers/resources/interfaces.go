package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/unrolled/render"
)

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
	SearchIncludes() SearchIncludes
	SearchParams() []SearchParam
}

// ValidateableResource ...
type ValidateableResource interface {
	Validate(log *logging.Logger, rndr *render.Render) http.Handler
}

// PatchableResource ...
type PatchableResource interface {
	Patch(log *logging.Logger, rndr *render.Render) http.Handler
}
