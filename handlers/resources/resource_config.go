package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/models"
)

type searchIncludes []string

type searchParam struct {
	Name                       string
	ObjectIndexContractAddress string
	Type                       models.SearchParameterType
}

// ResourceConfig ...
type ResourceConfig struct {
	ConditionalCreate bool
	ConditionalDelete models.CapabilityStatementResourceConditionalDelete
	ConditionalUpdate bool
	ConditionalRead   models.CapabilityStatementResourceConditionalRead
	SearchIncludes    searchIncludes
	SearchParams      []searchParam
	UpdateCreate      bool
	Versioning        models.CapabilityStatementResourceVersioning
}

// NewResourceConfig ...
func NewResourceConfig() *ResourceConfig {
	return &ResourceConfig{
		ConditionalDelete: models.CapabilityStatementResourceConditionalDeleteNotSupported,
		ConditionalRead:   models.CapabilityStatementResourceConditionalReadFullSupport,
		Versioning:        models.CapabilityStatementResourceVersioningNoVersion,
	}
}
