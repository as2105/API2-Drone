package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/storage/ethereum"
	"github.com/SynapticHealthAlliance/fhir-api/pkg/models"
	"github.com/pkg/errors"
)

// PractitionerRole ...
type PractitionerRole struct {
	EthereumResource
}

// NewPractitionerRole ...
func NewPractitionerRole(registry *Registry) (*PractitionerRole, error) {
	validator, err := models.NewJSONValidator(registry.box, "PractitionerRole")
	if err != nil {
		return nil, errors.Wrap(err, "could not create JSON validator")
	}

	newConfig := NewResourceConfig()
	newConfig.Versioning = models.CapabilityStatementResourceVersioningVersionedUpdate
	newConfig.SearchIncludes = searchIncludes{
		"PractitionerRole:location",
		"PractitionerRole:practitioner",
		"*",
	}
	newConfig.SearchParams = []searchParam{
		{Name: "_id", Type: models.SearchParameterTypeToken},
		{Name: "_lastUpdated", Type: models.SearchParameterTypeDate},
		{Name: "identifier", Type: models.SearchParameterTypeToken},
		{Name: "location", Type: models.SearchParameterTypeReference},
		{Name: "practitioner", Type: models.SearchParameterTypeReference},
		{Name: "telecom", Type: models.SearchParameterTypeToken},
	}

	collContract := registry.appConfig.ObjectCollectionContracts["PractitionerRole"]
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

	return &PractitionerRole{
		EthereumResource{
			adapter:       newAdapter,
			config:        newConfig,
			jsonValidator: validator,
			log:           registry.log,
			newModelFunc:  func() models.Resource { return &models.PractitionerRole{} },
			renderer:      registry.renderer,
		},
	}, nil
}
