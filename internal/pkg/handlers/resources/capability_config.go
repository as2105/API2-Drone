package resources

import (
	"encoding/json"
	"time"

	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/metadata"
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/utils"
	"github.com/SynapticHealthAlliance/fhir-api/pkg/models"
	"github.com/pkg/errors"
)

// CapabilityConfig ...
type CapabilityConfig struct {
	*models.CapabilityStatement
}

// AddResource ...
func (c *CapabilityConfig) AddResource(typeName string, i interface{}) error {
	newR := &models.CapabilityStatementResource{
		Type:        typeName,
		Interaction: c.getInteractions(i),
	}

	if t, ok := i.(ConfiguredResource); ok {
		config := t.GetResourceConfig()
		if config == nil {
			return errors.New("no resource configuration present")
		}
		newR.ConditionalCreate = config.ConditionalCreate
		newR.ConditionalDelete = config.ConditionalDelete
		newR.ConditionalUpdate = config.ConditionalUpdate
		newR.ConditionalRead = config.ConditionalRead
		newR.SearchInclude = config.SearchIncludes
		newR.SearchParam = c.getSearchParams(config.SearchParams)
		newR.UpdateCreate = config.UpdateCreate
		newR.Versioning = config.Versioning
	} else {
		return errors.New("resource is not configurable")
	}

	rArr := append(c.CapabilityStatement.Rest[0].Resource, newR)
	c.CapabilityStatement.Rest[0].Resource = rArr

	return nil
}

func (c *CapabilityConfig) getSearchParams(params []searchParam) []*models.CapabilityStatementSearchParam {
	modelParams := []*models.CapabilityStatementSearchParam{}
	for _, p := range params {
		modelParams = append(modelParams, &models.CapabilityStatementSearchParam{
			Name: p.Name,
			Type: models.CapabilityStatementSearchParamType(p.Type),
		})
	}
	return modelParams
}

func (c *CapabilityConfig) getInteractions(i interface{}) []*models.CapabilityStatementInteraction {
	ints := c.getRestfulInteractionTypes(i)
	mods := []*models.CapabilityStatementInteraction{}
	for _, i := range ints {
		mods = append(mods, &models.CapabilityStatementInteraction{Code: i})
	}
	return mods
}

func (c *CapabilityConfig) getRestfulInteractionTypes(i interface{}) []models.CapabilityStatementInteractionCode {
	ints := []models.CapabilityStatementInteractionCode{}
	if _, ok := i.(CreateableResource); ok {
		ints = append(ints, models.CapabilityStatementInteractionCodeCreate)
	}
	if _, ok := i.(SearchableResource); ok {
		ints = append(ints, models.CapabilityStatementInteractionCodeSearchType)
	}
	if _, ok := i.(UpdateableResource); ok {
		ints = append(ints, models.CapabilityStatementInteractionCodeUpdate)
	}
	if _, ok := i.(DeleteableResource); ok {
		ints = append(ints, models.CapabilityStatementInteractionCodeDelete)
	}
	if _, ok := i.(ReadableResource); ok {
		ints = append(ints, models.CapabilityStatementInteractionCodeRead)
	}
	if _, ok := i.(PatchableResource); ok {
		ints = append(ints, models.CapabilityStatementInteractionCodePatch)
	}
	if _, ok := i.(VersionReadableResource); ok {
		ints = append(ints, models.CapabilityStatementInteractionCodeVread)
	}
	return ints
}

// NewCapabilityConfig ...
func NewCapabilityConfig(registry *Registry) *CapabilityConfig {
	log := registry.log

	log.Debug("entering NewCapabilityConfig")

	newCS := &models.CapabilityStatement{}
	newCS.Date = metadata.Data.BuildTime.UTC().Format(time.RFC3339)
	newCS.Description = metadata.Data.AppName
	newCS.FhirVersion = models.FHIRVersion
	newCS.Format = []string{"json"}
	newCS.Kind = models.CapabilityStatementKindInstance
	newCS.Name = metadata.Data.AppName
	newCS.PatchFormat = []string{"application/json-patch+json"}
	newCS.Status = models.CapabilityStatementStatusDraft
	newCS.Version = metadata.Data.Version
	// newCS.Copyright = "Synaptic Health Alliance 2019"
	// newCS.Publisher = "Synaptic Health Alliance"
	// newCS.Purpose = "Synaptic Health Alliance"
	// newCS.Title = "Synaptic Health Alliance PDX"
	// newCS.UseContext = []*models.UsageContext{}
	// newCS.Jurisdiction = []*models.CodeableConcept{}
	// newCS.Format = []string{}
	// newCS.PatchFormat = []string{}
	// newCS.ImplementationGuide = []string{}

	newCS.Rest = []*models.CapabilityStatementRest{
		{Mode: models.CapabilityStatementRestModeServer},
	}

	newConfig := &CapabilityConfig{CapabilityStatement: newCS}

	for _, i := range registry.Resources {
		tName := utils.GetBaseTypeName(i)
		log.Debugf("adding registered resource %q", tName)
		if err := newConfig.AddResource(tName, i); err != nil {
			log.WithError(err).Panicf("failed to add registered resource %q to capability config", tName)
		}
	}

	// validate JSON
	log.Debug("validating generated capability statement")
	v, err := models.NewJSONValidator(registry.box, "CapabilityStatement")
	if err != nil {
		log.WithError(err).Panic("could not create JSON validator")
	}
	jsonBytes, err := json.Marshal(newCS)
	if err != nil {
		log.WithError(err).Panic("could not marshal object as JSON bytes")
	}
	if valid, vErrs, err := v.Validate(jsonBytes); !valid {
		if err != nil {
			log.Error(err)
		}
		for _, e := range vErrs {
			log.Error(e.String())
		}
		log.Error("generated capability statement is not valid")
	}
	log.Debug("capability statement is valid")

	return newConfig
}
