package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/config"
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/logging"
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/storage/database"
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/storage/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gobuffalo/packr/v2"
	"github.com/unrolled/render"
)

// Registry ...
type Registry struct {
	Resources []interface{}

	box          *packr.Box
	db           *database.DB
	connection   *ethclient.Client
	transactOpts *bind.TransactOpts
	appConfig    *config.Config
	log          *logging.Logger
	renderer     *render.Render
	txnsChan     ethereum.TransactionsChannel
}

func (r *Registry) add(resource interface{}) {
	r.Resources = append(r.Resources, resource)
}

// NewRegistry ...
func NewRegistry(
	box *packr.Box,
	db *database.DB,
	connection *ethclient.Client,
	transactOpts *bind.TransactOpts,
	appConfig *config.Config,
	log *logging.Logger,
	renderer *render.Render,
	txnsChan ethereum.TransactionsChannel,
) (*Registry, error) {
	registry := &Registry{
		box:          box,
		db:           db,
		connection:   connection,
		transactOpts: transactOpts,
		appConfig:    appConfig,
		log:          log,
		renderer:     renderer,
		txnsChan:     txnsChan,
	}

	// Practitioner
	practitioner, err := NewPractitioner(registry)
	if err != nil {
		return registry, err
	}
	registry.add(practitioner)

	// PractitionerRole
	practitionerRole, err := NewPractitionerRole(registry)
	if err != nil {
		return registry, err
	}
	registry.add(practitionerRole)

	// Location
	location, err := NewLocation(registry)
	if err != nil {
		return registry, err
	}
	registry.add(location)

	// Subscription
	subscription, err := NewSubscription(registry)
	if err != nil {
		return registry, err
	}
	registry.add(subscription)

	return registry, nil
}
