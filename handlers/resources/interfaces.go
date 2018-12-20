package resources

import (
	"net/http"

	"github.com/SynapticHealthAlliance/fhir-api/logging"
)

// CreateableResource ...
type CreateableResource interface {
	Create(log *logging.Logger) http.Handler
}

// ReadableResource ...
type ReadableResource interface {
	Read(log *logging.Logger) http.Handler
}

// VersionReadableResource ...
type VersionReadableResource interface {
	VersionRead(log *logging.Logger) http.Handler
}

// UpdateableResource ...
type UpdateableResource interface {
	Update(log *logging.Logger) http.Handler
}

// DeleteableResource ...
type DeleteableResource interface {
	Delete(log *logging.Logger) http.Handler
}

// SearchableResource ...
type SearchableResource interface {
	Search(log *logging.Logger) http.Handler
	SearchIncludes() SearchIncludes
	SearchParams() []SearchParam
}

// ValidateableResource ...
type ValidateableResource interface {
	Validate(log *logging.Logger) http.Handler
}

// PatchableResource ...
type PatchableResource interface {
	Patch(log *logging.Logger) http.Handler
}
