package resources

import (
	"net/http"
)

// ConfiguredResource ...
type ConfiguredResource interface {
	GetResourceConfig() *ResourceConfig
}

// CreateableResource ...
type CreateableResource interface {
	Create() http.Handler
}

// ReadableResource ...
type ReadableResource interface {
	Read() http.Handler
}

// VersionReadableResource ...
type VersionReadableResource interface {
	VersionRead() http.Handler
}

// UpdateableResource ...
type UpdateableResource interface {
	Update() http.Handler
}

// DeleteableResource ...
type DeleteableResource interface {
	Delete() http.Handler
}

// SearchableResource ...
type SearchableResource interface {
	Search() http.Handler
}

// ValidateableResource ...
type ValidateableResource interface {
	Validate() http.Handler
}

// PatchableResource ...
type PatchableResource interface {
	Patch() http.Handler
}
