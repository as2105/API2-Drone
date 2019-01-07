package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/config"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/storage/database"
	"github.com/SynapticHealthAlliance/fhir-api/storage/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gobuffalo/packr/v2"
)

// ResourceRegistry ...
type ResourceRegistry []interface{}

// NewResourceRegistry ...
func NewResourceRegistry(
	box *packr.Box,
	db *database.DB,
	connection *ethclient.Client,
	transactOpts *bind.TransactOpts,
	appConfig *config.Config,
	log *logging.Logger,
	txnsChan ethereum.TransactionsChannel,
) (ResourceRegistry, error) {
	registry := ResourceRegistry{}

	// Practitioner
	practitioner, err := NewPractitioner(box, connection, transactOpts, appConfig, log, txnsChan)
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
