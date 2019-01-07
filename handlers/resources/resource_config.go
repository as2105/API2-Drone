package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/models"
)

type searchIncludes []string

type searchParam struct {
	Name                       string
	ObjectIndexContractAddress string
	Type                       models.SearchParamType
}

// ResourceConfig ...
type ResourceConfig struct {
	ConditionalCreate bool
	ConditionalDelete models.ConditionalDeleteStatus
	ConditionalUpdate bool
	ConditionalRead   models.ConditionalReadStatus
	SearchIncludes    searchIncludes
	SearchParams      []searchParam
	UpdateCreate      bool
	Versioning        models.ResourceVersionPolicy
}

// NewResourceConfig ...
func NewResourceConfig() *ResourceConfig {
	return &ResourceConfig{
		ConditionalDelete: models.ConditionalDeleteStatusNotSupported,
		ConditionalRead:   models.ConditionalReadStatusFullSupport,
		Versioning:        models.ResourceVersionPolicyNoVersion,
	}
}
