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
	ConditionalDelete models.CapabilityStatement_ResourceConditionalDelete
	ConditionalUpdate bool
	SearchIncludes    searchIncludes
	SearchParams      []searchParam
	Versioning        models.CapabilityStatement_ResourceVersioning
}

// NewResourceConfig ...
func NewResourceConfig() *ResourceConfig {
	return &ResourceConfig{
		ConditionalDelete: models.CapabilityStatement_ResourceConditionalDeleteNotSupported,
		Versioning:        models.CapabilityStatement_ResourceVersioningNoVersion,
	}
}
