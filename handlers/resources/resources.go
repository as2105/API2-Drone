package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/SynapticHealthAlliance/fhir-api/storage/database"
	"github.com/gobuffalo/packr/v2"
)

type SearchIncludes []string

type SearchParam struct {
	Name                       string
	ObjectIndexContractAddress string
	Type                       models.SearchParamType
}

type ResourceRegistry []interface{}

func NewResourceRegistry(box *packr.Box, db *database.DB) (ResourceRegistry, error) {
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
	subscription, err := NewSubscription(box, db)
	if err != nil {
		return registry, err
	}
	registry = append(registry, subscription)

	return registry, nil
}
