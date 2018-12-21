package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/gobuffalo/packr/v2"
)

type SearchIncludes []string

type SearchParam struct {
	Name                       string
	ObjectIndexContractAddress string
	Type                       models.SearchParamType
}

type ResourceRegistry []interface{}

func NewResourceRegistry(box *packr.Box) (ResourceRegistry, error) {
	registry := ResourceRegistry{}

	// Practitioner
	practitioner, err := NewPractitioner(box)
	if err != nil {
		return registry, err
	}
	registry = append(registry, practitioner)

	// PractitionerRole
	practitionerRole, err := NewPractitionerRole(box)
	if err != nil {
		return registry, err
	}
	registry = append(registry, practitionerRole)

	// Location
	location, err := NewLocation(box)
	if err != nil {
		return registry, err
	}
	registry = append(registry, location)

	// Subscription
	subscription, err := NewSubscription(box)
	if err != nil {
		return registry, err
	}
	registry = append(registry, subscription)

	return registry, nil
}
