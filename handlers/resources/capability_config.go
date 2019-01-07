package resources

import (
	"encoding/json"

	"github.com/SynapticHealthAlliance/fhir-api/internal/metadata"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/SynapticHealthAlliance/fhir-api/utils"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
)

// CapabilityConfig ...
type CapabilityConfig struct {
	*models.CapabilityStatement
}

// AddResource ...
func (c *CapabilityConfig) AddResource(typeName string, i interface{}) error {
	newR := &models.CapabilityStatement_Resource{
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
		newR.Versioning = config.Versioning
		newR.SearchInclude = config.SearchIncludes
		newR.SearchParam = c.getSearchParams(config.SearchParams)
	} else {
		return errors.New("resource is not configurable")
	}

	rArr := append(c.CapabilityStatement.Rest[0].Resource, newR)
	c.CapabilityStatement.Rest[0].Resource = rArr

	return nil
}

func (c *CapabilityConfig) getSearchParams(params []searchParam) []*models.CapabilityStatement_SearchParam {
	modelParams := []*models.CapabilityStatement_SearchParam{}
	for _, p := range params {
		modelParams = append(modelParams, &models.CapabilityStatement_SearchParam{
			Name: p.Name,
			Type: models.CapabilityStatement_SearchParamType(p.Type),
		})
	}
	return modelParams
}

func (c *CapabilityConfig) getInteractions(i interface{}) []*models.CapabilityStatement_Interaction {
	ints := c.getRestfulInteractionTypes(i)
	mods := []*models.CapabilityStatement_Interaction{}
	for _, i := range ints {
		mods = append(mods, &models.CapabilityStatement_Interaction{Code: i})
	}
	return mods
}

func (c *CapabilityConfig) getRestfulInteractionTypes(i interface{}) []models.CapabilityStatement_InteractionCode {
	ints := []models.CapabilityStatement_InteractionCode{}
	if _, ok := i.(CreateableResource); ok {
		ints = append(ints, models.CapabilityStatement_InteractionCodeCreate)
	}
	if _, ok := i.(SearchableResource); ok {
		ints = append(ints, models.CapabilityStatement_InteractionCodeSearchType)
	}
	if _, ok := i.(UpdateableResource); ok {
		ints = append(ints, models.CapabilityStatement_InteractionCodeUpdate)
	}
	if _, ok := i.(DeleteableResource); ok {
		ints = append(ints, models.CapabilityStatement_InteractionCodeDelete)
	}
	if _, ok := i.(ReadableResource); ok {
		ints = append(ints, models.CapabilityStatement_InteractionCodeRead)
	}
	if _, ok := i.(PatchableResource); ok {
		ints = append(ints, models.CapabilityStatement_InteractionCodePatch)
	}
	if _, ok := i.(VersionReadableResource); ok {
		ints = append(ints, models.CapabilityStatement_InteractionCodeVread)
	}
	return ints
}

// NewCapabilityConfig ...
func NewCapabilityConfig(registry ResourceRegistry, log *logging.Logger, box *packr.Box) *CapabilityConfig {
	log.Debug("entering NewCapabilityConfig")

	newCS := &models.CapabilityStatement{}
	newCS.Date = metadata.Metadata.BuildTime
	newCS.Description = metadata.Metadata.AppName
	newCS.FhirVersion = models.FHIRVersion
	newCS.Format = []string{"json"}
	newCS.PatchFormat = []string{"application/json-patch+json"}
	newCS.Kind = models.CapabilityStatementKindInstance
	newCS.Name = metadata.Metadata.AppName
	newCS.Status = models.CapabilityStatementStatusDraft
	newCS.Version = metadata.Metadata.Version
	// newCS.Copyright = "Synaptic Health Alliance 2019"
	// newCS.Publisher = "Synaptic Health Alliance"
	// newCS.Purpose = "Synaptic Health Alliance"
	// newCS.Title = "Synaptic Health Alliance PDX"
	// newCS.UseContext = []*models.UsageContext{}
	// newCS.Jurisdiction = []*models.CodeableConcept{}
	// newCS.Format = []string{}
	// newCS.PatchFormat = []string{}
	// newCS.ImplementationGuide = []string{}

	newCS.Rest = []*models.CapabilityStatement_Rest{
		{Mode: models.CapabilityStatement_RestModeServer},
	}

	newConfig := &CapabilityConfig{CapabilityStatement: newCS}

	for _, i := range registry {
		tName := utils.GetBaseTypeName(i)
		log.Debugf("adding registered resource %q", tName)
		if err := newConfig.AddResource(tName, i); err != nil {
			log.WithError(err).Panicf("failed to add registered resource %q to capability config", tName)
		}
	}

	// validate JSON
	log.Debug("validating generated capability statement")
	v, err := models.NewJSONValidator(box, "CapabilityStatement")
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
