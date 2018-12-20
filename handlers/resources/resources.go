package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/models"
)

type SearchIncludes []string

type SearchParam struct {
	Name                       string
	ObjectIndexContractAddress string
	Type                       models.SearchParamType
}

// Add models to this registry to expose them
var registry = []interface{}{
	&Practitioner{},
	&PractitionerRole{},
	&Location{},
}

// GetRegisteredResources ...
func GetRegisteredResources() []interface{} {
	return registry
}
