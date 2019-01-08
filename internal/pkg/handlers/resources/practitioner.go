package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/storage/ethereum"
	"github.com/SynapticHealthAlliance/fhir-api/pkg/models"
	"github.com/pkg/errors"
)

// Practitioner ...
type Practitioner struct {
	EthereumResource
}

// NewPractitioner ...
func NewPractitioner(registry *Registry) (*Practitioner, error) {
	validator, err := models.NewJSONValidator(registry.box, "Practitioner")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}

	newConfig := NewResourceConfig()
	newConfig.Versioning = models.CapabilityStatementResourceVersioningVersionedUpdate
	newConfig.SearchIncludes = searchIncludes{
		"Practitioner:location",
		"*",
	}
	newConfig.SearchParams = []searchParam{
		{Name: "_id", Type: models.SearchParameterTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParameterTypeDate},
		{Name: "active", Type: models.SearchParameterTypeToken},
		{Name: "identifier", Type: models.SearchParameterTypeToken},
		{Name: "location", Type: models.SearchParameterTypeReference},
		{Name: "name", Type: models.SearchParameterTypeString},
		{Name: "telecom", Type: models.SearchParameterTypeToken},
	}

	collContract := registry.appConfig.ObjectCollectionContracts["Practitioner"]
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

	return &Practitioner{
		EthereumResource{
			adapter:       newAdapter,
			config:        newConfig,
			jsonValidator: validator,
			log:           registry.log,
			newModelFunc:  func() models.Resource { return &models.Practitioner{} },
			renderer:      registry.renderer,
		},
	}, nil
}
