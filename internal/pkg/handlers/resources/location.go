package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/storage/ethereum"
	"github.com/SynapticHealthAlliance/fhir-api/pkg/models"
	"github.com/pkg/errors"
)

// Location ...
type Location struct {
	EthereumResource
}

// NewLocation ...
func NewLocation(registry *Registry) (*Location, error) {
	validator, err := models.NewJSONValidator(registry.box, "Location")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}

	newConfig := NewResourceConfig()
	newConfig.Versioning = models.CapabilityStatementResourceVersioningVersionedUpdate
	newConfig.SearchParams = []searchParam{
		{Name: "_id", Type: models.SearchParameterTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParameterTypeDate},
		{Name: "identifier", Type: models.SearchParameterTypeToken},
		{Name: "status", Type: models.SearchParameterTypeToken},
		{Name: "type", Type: models.SearchParameterTypeToken},
	}

	collContract := registry.appConfig.ObjectCollectionContracts["Location"]
	if collContract == nil {
		return nil, errors.New("no collection contract found")
	}
	newAdapter, err := ethereum.NewAdapter(
		registry.connection,
		registry.transactOpts,
		registry.appConfig.OrganizationContract,
		collContract,
		registry.txnsChan,
		registry.log,
	)
	if err != nil {
		return nil, errors.New("failed to create ethereum adapter")
	}

	return &Location{
		EthereumResource{
			adapter:       newAdapter,
			config:        newConfig,
			jsonValidator: validator,
			log:           registry.log,
			newModelFunc:  func() models.Resource { return &models.Location{} },
			renderer:      registry.renderer,
		},
	}, nil
}
