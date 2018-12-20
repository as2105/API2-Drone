package resources

import (
	"github.com/SynapticHealthAlliance/fhir-api/internal/metadata"
	"github.com/SynapticHealthAlliance/fhir-api/models"
	"github.com/SynapticHealthAlliance/fhir-api/utils"
)

// CapabilityConfig ...
type CapabilityConfig struct {
	*models.CapabilityStatement
}

func (c *CapabilityConfig) AddResource(i interface{}) {
	newR := &models.CapabilityStatement_Resource{
		ConditionalCreate: true,
		ConditionalDelete: string(models.ConditionalDeleteStatusSingle),
		ConditionalUpdate: true,
		Type:              utils.GetBaseTypeName(i),
		Interaction:       c.getInteractions(i),
		Versioning:        string(models.ResourceVersionPolicyVersionedUpdate),
	}

	if t, ok := i.(SearchableResource); ok {
		newR.SearchInclude = t.SearchIncludes()
		newR.SearchParam = c.getSearchParam(t)
	}

	rArr := append(c.CapabilityStatement.Rest[0].Resource, newR)
	c.CapabilityStatement.Rest[0].Resource = rArr
}

func (c *CapabilityConfig) getSearchParam(t SearchableResource) []*models.CapabilityStatement_SearchParam {
	params := t.SearchParams()
	modelParams := []*models.CapabilityStatement_SearchParam{}
	for _, p := range params {
		modelParams = append(modelParams, &models.CapabilityStatement_SearchParam{
			Name: p.Name,
			Type: string(p.Type),
		})
	}
	return modelParams
}

func (c *CapabilityConfig) getInteractions(i interface{}) []*models.CapabilityStatement_Interaction {
	ints := c.getRestfulInteractionTypes(i)
	mods := []*models.CapabilityStatement_Interaction{}
	for _, i := range ints {
		mods = append(mods, &models.CapabilityStatement_Interaction{Code: string(i)})
	}
	return mods
}

func (c *CapabilityConfig) getRestfulInteractionTypes(i interface{}) []models.TypeRestfulInteraction {
	ints := []models.TypeRestfulInteraction{}
	if _, ok := i.(CreateableResource); ok {
		ints = append(ints, models.TypeRestfulInteractionCreate)
	}
	if _, ok := i.(SearchableResource); ok {
		ints = append(ints, models.TypeRestfulInteractionSearchType)
	}
	if _, ok := i.(UpdateableResource); ok {
		ints = append(ints, models.TypeRestfulInteractionUpdate)
	}
	if _, ok := i.(DeleteableResource); ok {
		ints = append(ints, models.TypeRestfulInteractionDelete)
	}
	if _, ok := i.(ReadableResource); ok {
		ints = append(ints, models.TypeRestfulInteractionRead)
	}
	if _, ok := i.(PatchableResource); ok {
		ints = append(ints, models.TypeRestfulInteractionPatch)
	}
	if _, ok := i.(VersionReadableResource); ok {
		ints = append(ints, models.TypeRestfulInteractionVRead)
	}
	return ints
}

func NewCapabilityConfig() *CapabilityConfig {
	newCS := models.NewCapabilityStatement()
	newCS.AcceptUnknown = string(models.UnknownContentCodeExtensions)
	newCS.Date = metadata.Metadata.BuildTime
	newCS.Description = metadata.Metadata.AppName
	newCS.FhirVersion = models.FHIRVersion
	newCS.Format = []string{string(models.MimeTypeJSON)}
	newCS.Kind = string(models.CapabilityStatementKindInstance)
	newCS.Name = metadata.Metadata.AppName
	newCS.Status = string(models.PublicationStatusDraft)
	newCS.Version = metadata.Metadata.Version
	newCS.Rest = []*models.CapabilityStatement_Rest{
		{Mode: string(models.RestfulCapabilityModeServer)},
	}

	newConfig := &CapabilityConfig{CapabilityStatement: newCS}

	for _, i := range GetRegisteredResources() {
		newConfig.AddResource(i)
	}

	return newConfig
}
