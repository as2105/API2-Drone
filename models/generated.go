package models

// see http://hl7.org/fhir/json.html#schema for information about the FHIR Json Schemas
// Schema: http://json-schema.org/draft-04/schema#
import (
	"time"
)

// CarePlan is Describes the intention of how one or more practitioners intend to deliver care for a particular patient, group or community for a period of time, possibly limited to care for a specific condition or set of conditions.
type CarePlan struct {
	_ *DomainResource
	// A care plan that is fulfilled in whole or in part by this care plan.
	BasedOn []*Reference `json:"basedOn"`
	// A larger care plan of which this particular care plan is a component or step.
	PartOf []*Reference `json:"partOf"`
	// Indicates whether the plan is currently being acted upon, represents future intentions or is now a historical record.
	Status string `json:"status"`
	// Identifies the conditions/problems/concerns/diagnoses/etc. whose management and/or mitigation are handled by this plan.
	Addresses []*Reference `json:"addresses"`
	// General notes about the care plan not covered elsewhere.
	Note []*Annotation `json:"note"`
	// Completed or terminated care plan whose function is taken by this new care plan.
	Replaces []*Reference `json:"replaces"`
	// A description of the scope and nature of the plan.
	Description string `json:"description"`
	// Identifies the individual(s) or ogranization who is responsible for the content of the care plan.
	Author []*Reference `json:"author"`
	// Identifies a planned action to occur as part of the plan.  For example, a medication to be used, lab tests to perform, self-monitoring, education, etc.
	Activity []*CarePlan_Activity `json:"activity"`
	// Identifies portions of the patient's record that specifically influenced the formation of the plan.  These might include co-morbidities, recent procedures, limitations, recent assessments, etc.
	SupportingInfo []*Reference `json:"supportingInfo"`
	// Extensions for intent
	Intent_ext *Element `json:"_intent"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Identifies the original context in which this particular CarePlan was created.
	Context *Reference `json:"context"`
	// Indicates when the plan did (or is intended to) come into effect and end.
	Period *Period `json:"period"`
	// Indicates the level of authority/intentionality associated with the care plan and where the care plan fits into the workflow chain.
	Intent string `json:"intent"`
	// Identifies what "kind" of plan this is to support differentiation between multiple co-existing plans; e.g. "Home health", "psychiatric", "asthma", "disease management", "wellness plan", etc.
	Category []*CodeableConcept `json:"category"`
	// Human-friendly name for the CarePlan.
	Title string `json:"title"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// This is a CarePlan resource
	ResourceType string `json:"resourceType,omitempty"`
	// This records identifiers associated with this care plan that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []*Identifier `json:"identifier"`
	// Identifies the protocol, questionnaire, guideline or other specification the care plan should be conducted in accordance with.
	Definition []*Reference `json:"definition"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Identifies the patient or group whose intended care is described by the plan.
	Subject *Reference `json:"subject,omitempty"`
	// Identifies all people and organizations who are expected to be involved in the care envisioned by this plan.
	CareTeam []*Reference `json:"careTeam"`
	// Describes the intended objective(s) of carrying out the care plan.
	Goal []*Reference `json:"goal"`
}

func NewCarePlan() *CarePlan { return &CarePlan{ResourceType: "CarePlan"} }

// ExplanationOfBenefit is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit struct {
	_ *DomainResource
	// Other claims which are related to this claim such as prior claim versions or for related services.
	Related []*ExplanationOfBenefit_Related `json:"related"`
	// Balance by Benefit Category.
	BenefitBalance []*ExplanationOfBenefit_BenefitBalance `json:"benefitBalance"`
	// The referral resource which lists the date, practitioner, reason and other supporting information.
	Referral *Reference `json:"referral"`
	// Processing outcome errror, partial or complete processing.
	Outcome *CodeableConcept `json:"outcome"`
	// The business identifier for the instance: invoice number, claim number, pre-determination or pre-authorization number.
	ClaimResponse *Reference `json:"claimResponse"`
	// Prescription to support the dispensing of Pharmacy or Vision products.
	Prescription *Reference `json:"prescription"`
	// Financial instrument by which payment information for health care.
	Insurance *ExplanationOfBenefit_Insurance `json:"insurance"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The provider which is responsible for the claim.
	Provider *Reference `json:"provider"`
	// Precedence (primary, secondary, etc.).
	// pattern [1-9][0-9]*
	Precedence uint64 `json:"precedence"`
	// Extensions for precedence
	Precedence_ext *Element `json:"_precedence"`
	// The form to be used for printing the content.
	Form *CodeableConcept `json:"form"`
	// The status of the resource instance.
	Status string `json:"status"`
	// The members of the team who provided the overall service as well as their role and whether responsible and qualifications.
	CareTeam []*ExplanationOfBenefit_CareTeam `json:"careTeam"`
	// The person who created the explanation of benefit.
	Enterer *Reference `json:"enterer"`
	// The provider which is responsible for the claim.
	Organization *Reference `json:"organization"`
	// The business identifier for the instance: invoice number, claim number, pre-determination or pre-authorization number.
	Claim *Reference `json:"claim"`
	// An accident which resulted in the need for healthcare services.
	Accident *ExplanationOfBenefit_Accident `json:"accident"`
	// The start and optional end dates of when the patient was precluded from working due to the treatable condition(s).
	EmploymentImpacted *Period `json:"employmentImpacted"`
	// The first tier service adjudications for payor added services.
	AddItem []*ExplanationOfBenefit_AddItem `json:"addItem"`
	// A finer grained suite of claim subtype codes which may convey Inpatient vs Outpatient and/or a specialty service. In the US the BillType.
	SubType []*CodeableConcept `json:"subType"`
	// The date when the EOB was created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// Original prescription which has been superceded by this prescription to support the dispensing of pharmacy services, medications or products. For example, a physician may prescribe a medication which the pharmacy determines is contraindicated, or for which the patient has an intolerance, and therefor issues a new precription for an alternate medication which has the same theraputic intent. The prescription from the pharmacy becomes the 'prescription' and that from the physician becomes the 'original prescription'.
	OriginalPrescription *Reference `json:"originalPrescription"`
	// The total cost of the services reported.
	TotalCost *Money `json:"totalCost"`
	// The amount of deductable applied which was not allocated to any particular service line.
	UnallocDeductable *Money `json:"unallocDeductable"`
	// Note text.
	ProcessNote []*ExplanationOfBenefit_ProcessNote `json:"processNote"`
	// The EOB Business Identifier.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for disposition
	Disposition_ext *Element `json:"_disposition"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// A description of the status of the adjudication.
	Disposition string `json:"disposition"`
	// Ordered list of patient diagnosis for which care is sought.
	Diagnosis []*ExplanationOfBenefit_Diagnosis `json:"diagnosis"`
	// The category of claim, eg, oral, pharmacy, vision, insitutional, professional.
	Type *CodeableConcept `json:"type"`
	// Patient Resource.
	Patient *Reference `json:"patient"`
	// Facility where the services were provided.
	Facility *Reference `json:"facility"`
	// Payment details for the claim if the claim has been paid.
	Payment *ExplanationOfBenefit_Payment `json:"payment"`
	// This is a ExplanationOfBenefit resource
	ResourceType string `json:"resourceType,omitempty"`
	// The insurer which is responsible for the explanation of benefit.
	Insurer *Reference `json:"insurer"`
	// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues. Often there are mutiple jurisdiction specific valuesets which are required.
	Information []*ExplanationOfBenefit_Information `json:"information"`
	// Ordered list of patient procedures performed to support the adjudication.
	Procedure []*ExplanationOfBenefit_Procedure `json:"procedure"`
	// The start and optional end dates of when the patient was confined to a treatment center.
	Hospitalization *Period `json:"hospitalization"`
	// First tier of goods and services.
	Item []*ExplanationOfBenefit_Item `json:"item"`
	// Total amount of benefit payable (Equal to sum of the Benefit amounts from all detail lines and additions less the Unallocated Deductable).
	TotalBenefit *Money `json:"totalBenefit"`
	// The billable period for which charges are being submitted.
	BillablePeriod *Period `json:"billablePeriod"`
	// The party to be reimbursed for the services.
	Payee *ExplanationOfBenefit_Payee `json:"payee"`
}

func NewExplanationOfBenefit() *ExplanationOfBenefit {
	return &ExplanationOfBenefit{ResourceType: "ExplanationOfBenefit"}
}

// ExplanationOfBenefit_Accident is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Accident struct {
	_ *BackboneElement
	// Where the accident occurred.
	LocationAddress *Address `json:"locationAddress"`
	// Where the accident occurred.
	LocationReference *Reference `json:"locationReference"`
	// Date of an accident which these services are addressing.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Type of accident: work, auto, etc.
	Type *CodeableConcept `json:"type"`
}

// MessageHeader_Response is The header for a message exchange that is either requesting or responding to an action.  The reference(s) that are the subject of the action as well as other information related to the action are typically transmitted in a bundle in which the MessageHeader resource instance is the first resource in the bundle.
type MessageHeader_Response struct {
	_ *BackboneElement
	// The MessageHeader.id of the message to which this message is a response.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Identifier string `json:"identifier"`
	// Extensions for identifier
	Identifier_ext *Element `json:"_identifier"`
	// Code that identifies the type of response to the message - whether it was successful or not, and whether it should be resent or not.
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// Full details of any issues found in the message.
	Details *Reference `json:"details"`
}

// PlanDefinition_DynamicValue is This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinition_DynamicValue struct {
	_ *BackboneElement
	// The path to the element to be customized. This is the path on the resource that will hold the result of the calculation defined by the expression.
	Path string `json:"path"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// The media type of the language for the expression.
	Language string `json:"language"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
	// An expression specifying the value of the customized element.
	Expression string `json:"expression"`
	// Extensions for expression
	Expression_ext *Element `json:"_expression"`
	// A brief, natural language description of the intended semantics of the dynamic value.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
}

// ClinicalImpression is A record of a clinical assessment performed to determine what problem(s) may affect the patient and before planning the treatments or management strategies that are best to manage a patient's condition. Assessments are often 1:1 with a clinical consultation / encounter,  but this varies greatly depending on the clinical workflow. This resource is called "ClinicalImpression" rather than "ClinicalAssessment" to avoid confusion with the recording of assessment tools such as Apgar score.
type ClinicalImpression struct {
	_ *DomainResource
	// Action taken as part of assessment procedure.
	Action []*Reference `json:"action"`
	// The encounter or episode of care this impression was created as part of.
	Context *Reference `json:"context"`
	// Extensions for protocol
	Protocol_ext []*Element `json:"_protocol"`
	// Extensions for summary
	Summary_ext *Element `json:"_summary"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// This is a ClinicalImpression resource
	ResourceType string `json:"resourceType,omitempty"`
	// Indicates when the documentation of the assessment was complete.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// The clinician performing the assessment.
	Assessor *Reference `json:"assessor"`
	// Reference to a specific published clinical protocol that was followed during this assessment, and/or that provides evidence in support of the diagnosis.
	Protocol []string `json:"protocol"`
	// A summary of the context and/or cause of the assessment - why / where was it performed, and what patient events/status prompted it.
	Description string `json:"description"`
	// Commentary about the impression, typically recorded after the impression itself was made, though supplemental notes by the original author could also appear.
	Note []*Annotation `json:"note"`
	// One or more sets of investigations (signs, symptions, etc.). The actual grouping of investigations vary greatly depending on the type and context of the assessment. These investigations may include data generated during the assessment process, or data previously generated and recorded that is pertinent to the outcomes.
	Investigation []*ClinicalImpression_Investigation `json:"investigation"`
	// Estimate of likely outcome.
	PrognosisCodeableConcept []*CodeableConcept `json:"prognosisCodeableConcept"`
	// A text summary of the investigations and the diagnosis.
	Summary string `json:"summary"`
	// A unique identifier assigned to the clinical impression that remains consistent regardless of what server the impression is stored on.
	Identifier []*Identifier `json:"identifier"`
	// The patient or group of individuals assessed as part of this record.
	Subject *Reference `json:"subject,omitempty"`
	// The point in time or period over which the subject was assessed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	EffectiveDateTime time.Time `json:"effectiveDateTime"`
	// Categorizes the type of clinical assessment performed.
	Code *CodeableConcept `json:"code"`
	// The point in time or period over which the subject was assessed.
	EffectivePeriod *Period `json:"effectivePeriod"`
	// This a list of the relevant problems/conditions for a patient.
	Problem []*Reference `json:"problem"`
	// RiskAssessment expressing likely outcome.
	PrognosisReference []*Reference `json:"prognosisReference"`
	// A reference to the last assesment that was conducted bon this patient. Assessments are often/usually ongoing in nature; a care provider (practitioner or team) will make new assessments on an ongoing basis as new data arises or the patient's conditions changes.
	Previous *Reference `json:"previous"`
	// Specific findings or diagnoses that was considered likely or relevant to ongoing treatment.
	Finding []*ClinicalImpression_Finding `json:"finding"`
	// Identifies the workflow status of the assessment.
	Status string `json:"status"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for effectiveDateTime
	EffectiveDateTime_ext *Element `json:"_effectiveDateTime"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
}

func NewClinicalImpression() *ClinicalImpression {
	return &ClinicalImpression{ResourceType: "ClinicalImpression"}
}

// TestReport_Operation is A summary of information based on the results of executing a TestScript.
type TestReport_Operation struct {
	_ *BackboneElement
	// The result of this operation.
	Result string `json:"result"`
	// Extensions for result
	Result_ext *Element `json:"_result"`
	// An explanatory message associated with the result.
	Message string `json:"message"`
	// Extensions for message
	Message_ext *Element `json:"_message"`
	// A link to further details on the result.
	Detail string `json:"detail"`
	// Extensions for detail
	Detail_ext *Element `json:"_detail"`
}

// Bundle_Request is A container for a collection of resources.
type Bundle_Request struct {
	_ *BackboneElement
	// If the ETag values match, return a 304 Not Modified status. See the API documentation for ["Conditional Read"](http.html#cread).
	IfNoneMatch string `json:"ifNoneMatch"`
	// Extensions for ifModifiedSince
	IfModifiedSince_ext *Element `json:"_ifModifiedSince"`
	// Extensions for ifNoneExist
	IfNoneExist_ext *Element `json:"_ifNoneExist"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for method
	Method_ext *Element `json:"_method"`
	// The URL for this entry, relative to the root (the address to which the request is posted).
	Url string `json:"url"`
	// Extensions for ifNoneMatch
	IfNoneMatch_ext *Element `json:"_ifNoneMatch"`
	// Only perform the operation if the last updated date matches. See the API documentation for ["Conditional Read"](http.html#cread).
	IfModifiedSince string `json:"ifModifiedSince"`
	// Only perform the operation if the Etag value matches. For more information, see the API section ["Managing Resource Contention"](http.html#concurrency).
	IfMatch string `json:"ifMatch"`
	// Extensions for ifMatch
	IfMatch_ext *Element `json:"_ifMatch"`
	// Instruct the server not to perform the create if a specified resource already exists. For further information, see the API documentation for ["Conditional Create"](http.html#ccreate). This is just the query portion of the URL - what follows the "?" (not including the "?").
	IfNoneExist string `json:"ifNoneExist"`
	// The HTTP verb for this entry in either a change history, or a transaction/ transaction response.
	Method string `json:"method"`
}

// CodeSystem_Concept is A code system resource specifies a set of codes drawn from one or more code systems.
type CodeSystem_Concept struct {
	_ *BackboneElement
	// A code - a text symbol - that uniquely identifies the concept within the code system.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// A human readable string that is the recommended default way to present this concept to a user.
	Display string `json:"display"`
	// A property value for this concept.
	Property []*CodeSystem_Property1 `json:"property"`
	// Defines children of a concept to produce a hierarchy of concepts. The nature of the relationships is variable (is-a/contains/categorizes) - see hierarchyMeaning.
	Concept []*CodeSystem_Concept `json:"concept"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// Extensions for display
	Display_ext *Element `json:"_display"`
	// The formal definition of the concept. The code system resource does not make formal definitions required, because of the prevalence of legacy systems. However, they are highly recommended, as without them there is no formal meaning associated with the concept.
	Definition string `json:"definition"`
	// Extensions for definition
	Definition_ext *Element `json:"_definition"`
	// Additional representations for the concept - other languages, aliases, specialized purposes, used for particular purposes, etc.
	Designation []*CodeSystem_Designation `json:"designation"`
}

// ExplanationOfBenefit_Insurance is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Insurance struct {
	_ *BackboneElement
	// Reference to the program or plan identification, underwriter or payor.
	Coverage *Reference `json:"coverage"`
	// A list of references from the Insurer to which these services pertain.
	PreAuthRef []string `json:"preAuthRef"`
	// Extensions for preAuthRef
	PreAuthRef_ext []*Element `json:"_preAuthRef"`
}

// ExplanationOfBenefit_BenefitBalance is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_BenefitBalance struct {
	_ *BackboneElement
	// A short name or tag for the benefit, for example MED01, or DENT2.
	Name string `json:"name"`
	// A richer description of the benefit, for example 'DENT2 covers 100% of basic, 50% of major but exclused Ortho, Implants and Costmetic services'.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Unit designation: individual or family.
	Unit *CodeableConcept `json:"unit"`
	// The term or period of the values such as 'maximum lifetime benefit' or 'maximum annual vistis'.
	Term *CodeableConcept `json:"term"`
	// Benefits Used to date.
	Financial []*ExplanationOfBenefit_Financial `json:"financial"`
	// Dental, Vision, Medical, Pharmacy, Rehab etc.
	Category *CodeableConcept `json:"category,omitempty"`
	// Dental: basic, major, ortho; Vision exam, glasses, contacts; etc.
	SubCategory *CodeableConcept `json:"subCategory"`
	// True if the indicated class of service is excluded from the plan, missing or False indicated the service is included in the coverage.
	Excluded bool `json:"excluded"`
	// Extensions for excluded
	Excluded_ext *Element `json:"_excluded"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Network designation.
	Network *CodeableConcept `json:"network"`
}

// CodeSystem_Property is A code system resource specifies a set of codes drawn from one or more code systems.
type CodeSystem_Property struct {
	_ *BackboneElement
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The type of the property value. Properties of type "code" contain a code defined by the code system (e.g. a reference to anotherr defined concept).
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// A code that is used to identify the property. The code is used internally (in CodeSystem.concept.property.code) and also externally, such as in property filters.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// Reference to the formal meaning of the property. One possible source of meaning is the [Concept Properties](codesystem-concept-properties.html) code system.
	Uri string `json:"uri"`
	// Extensions for uri
	Uri_ext *Element `json:"_uri"`
	// A description of the property- why it is defined, and how its value might be used.
	Description string `json:"description"`
}

// ImmunizationRecommendation is A patient's point-in-time immunization and recommendation (i.e. forecasting a patient's immunization eligibility according to a published schedule) with optional supporting justification.
type ImmunizationRecommendation struct {
	_ *DomainResource
	// This is a ImmunizationRecommendation resource
	ResourceType string `json:"resourceType,omitempty"`
	// A unique identifier assigned to this particular recommendation record.
	Identifier []*Identifier `json:"identifier"`
	// The patient the recommendations are for.
	Patient *Reference `json:"patient,omitempty"`
	// Vaccine administration recommendations.
	Recommendation []*ImmunizationRecommendation_Recommendation `json:"recommendation,omitempty"`
}

func NewImmunizationRecommendation() *ImmunizationRecommendation {
	return &ImmunizationRecommendation{ResourceType: "ImmunizationRecommendation"}
}

// MedicationRequest is An order or request for both supply of the medication and the instructions for administration of the medication to a patient. The resource is called "MedicationRequest" rather than "MedicationPrescription" or "MedicationOrder" to generalize the use across inpatient and outpatient settings, including care plans, etc., and to harmonize with workflow patterns.
type MedicationRequest struct {
	_ *DomainResource
	// This is a MedicationRequest resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The individual, organization or device that initiated the request and has responsibility for its activation.
	Requester *MedicationRequest_Requester `json:"requester"`
	// Indicates how the medication is to be used by the patient.
	DosageInstruction []*Dosage `json:"dosageInstruction"`
	// A shared identifier common to all requests that were authorized more or less simultaneously by a single author, representing the identifier of the requisition or prescription.
	GroupIdentifier *Identifier `json:"groupIdentifier"`
	// Extensions for priority
	Priority_ext *Element `json:"_priority"`
	// Protocol or definition followed by this request.
	Definition []*Reference `json:"definition"`
	// Indicates the type of medication order and where the medication is expected to be consumed or administered.
	Category *CodeableConcept `json:"category"`
	// The date (and perhaps time) when the prescription was initially written or authored on.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AuthoredOn time.Time `json:"authoredOn"`
	// Condition or observation that supports why the medication was ordered.
	ReasonReference []*Reference `json:"reasonReference"`
	// Extra information about the prescription that could not be conveyed by the other attributes.
	Note []*Annotation `json:"note"`
	// Indicates whether or not substitution can or should be part of the dispense. In some cases substitution must happen, in other cases substitution must not happen. This block explains the prescriber's intent. If nothing is specified substitution may be done.
	Substitution *MedicationRequest_Substitution `json:"substitution"`
	// This records identifiers associated with this medication request that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate. For example a re-imbursement system might issue its own id for each prescription that is created.  This is particularly important where FHIR only provides part of an entire workflow process where records must be tracked through an entire system.
	Identifier []*Identifier `json:"identifier"`
	// A plan or request that is fulfilled in whole or in part by this medication request.
	BasedOn []*Reference `json:"basedOn"`
	// Whether the request is a proposal, plan, or an original order.
	Intent string `json:"intent"`
	// Identifies the medication being requested. This is a link to a resource that represents the medication which may be the details of the medication or simply an attribute carrying a code that identifies the medication from a known list of medications.
	MedicationCodeableConcept *CodeableConcept `json:"medicationCodeableConcept"`
	// A link to an encounter, or episode of care, that identifies the particular occurrence or set occurrences of contact between patient and health care provider.
	Context *Reference `json:"context"`
	// Links to Provenance records for past versions of this resource or fulfilling request or event resources that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the resource.
	EventHistory []*Reference `json:"eventHistory"`
	// Include additional information (for example, patient height and weight) that supports the ordering of the medication.
	SupportingInformation []*Reference `json:"supportingInformation"`
	// The person who entered the order on behalf of another individual for example in the case of a verbal or a telephone order.
	Recorder *Reference `json:"recorder"`
	// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. Drug-drug interaction, duplicate therapy, dosage alert etc.
	DetectedIssue []*Reference `json:"detectedIssue"`
	// The reason or the indication for ordering the medication.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// A code specifying the current state of the order.  Generally this will be active or completed state.
	Status string `json:"status"`
	// Extensions for intent
	Intent_ext *Element `json:"_intent"`
	// Indicates how quickly the Medication Request should be addressed with respect to other requests.
	Priority string `json:"priority"`
	// Identifies the medication being requested. This is a link to a resource that represents the medication which may be the details of the medication or simply an attribute carrying a code that identifies the medication from a known list of medications.
	MedicationReference *Reference `json:"medicationReference"`
	// A link to a resource representing the person or set of individuals to whom the medication will be given.
	Subject *Reference `json:"subject,omitempty"`
	// Indicates the specific details for the dispense or medication supply part of a medication request (also known as a Medication Prescription or Medication Order).  Note that this information is not always sent with the order.  There may be in some settings (e.g. hospitals) institutional or system support for completing the dispense details in the pharmacy department.
	DispenseRequest *MedicationRequest_DispenseRequest `json:"dispenseRequest"`
	// A link to a resource representing an earlier order related order or prescription.
	PriorPrescription *Reference `json:"priorPrescription"`
	// Extensions for authoredOn
	AuthoredOn_ext *Element `json:"_authoredOn"`
}

func NewMedicationRequest() *MedicationRequest {
	return &MedicationRequest{ResourceType: "MedicationRequest"}
}

// Subscription_Channel is The subscription resource is used to define a push based subscription from a server to another system. Once a subscription is registered with the server, the server checks every resource that is created or updated, and if the resource matches the given criteria, it sends a message on the defined "channel" so that another system is able to take an appropriate action.
type Subscription_Channel struct {
	_ *BackboneElement
	// Extensions for payload
	Payload_ext *Element `json:"_payload"`
	// Additional headers / information to send as part of the notification.
	Header []string `json:"header"`
	// Extensions for header
	Header_ext []*Element `json:"_header"`
	// The type of channel to send notifications on.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The uri that describes the actual end-point to send messages to.
	Endpoint string `json:"endpoint"`
	// Extensions for endpoint
	Endpoint_ext *Element `json:"_endpoint"`
	// The mime type to send the payload in - either application/fhir+xml, or application/fhir+json. If the payload is not present, then there is no payload in the notification, just a notification.
	Payload string `json:"payload"`
}

// Substance_Instance is A homogeneous material with a definite composition.
type Substance_Instance struct {
	_ *BackboneElement
	// Identifier associated with the package/container (usually a label affixed directly).
	Identifier *Identifier `json:"identifier"`
	// When the substance is no longer valid to use. For some substances, a single arbitrary date is used for expiry.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Expiry time.Time `json:"expiry"`
	// Extensions for expiry
	Expiry_ext *Element `json:"_expiry"`
	// The amount of the substance.
	Quantity *Quantity `json:"quantity"`
}

// TestScript_Param is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Param struct {
	_ *BackboneElement
	// Descriptive name for this parameter that matches the external assert rule parameter name.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The explicit or dynamic value for the parameter that will be passed on to the external rule template.
	Value string `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
}

// TestScript_Rule1 is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Rule1 struct {
	_ *BackboneElement
	// Id of the referenced rule within the external ruleset template.
	// pattern [A-Za-z0-9\-\.]{1,64}
	RuleId string `json:"ruleId"`
	// Extensions for ruleId
	RuleId_ext *Element `json:"_ruleId"`
	// Each rule template can take one or more parameters for rule evaluation.
	Param []*TestScript_Param1 `json:"param"`
}

// EligibilityResponse_Error is This resource provides eligibility and plan details from the processing of an Eligibility resource.
type EligibilityResponse_Error struct {
	_ *BackboneElement
	// An error code,from a specified code system, which details why the eligibility check could not be performed.
	Code *CodeableConcept `json:"code,omitempty"`
}

// OperationOutcome_Issue is A collection of error, warning or information messages that result from a system action.
type OperationOutcome_Issue struct {
	_ *BackboneElement
	// Indicates whether the issue indicates a variation from successful processing.
	Severity string `json:"severity"`
	// Extensions for severity
	Severity_ext *Element `json:"_severity"`
	// Extensions for diagnostics
	Diagnostics_ext *Element `json:"_diagnostics"`
	// For resource issues, this will be a simple XPath limited to element names, repetition indicators and the default child access that identifies one of the elements in the resource that caused this issue to be raised.  For HTTP errors, will be "http." + the parameter name.
	Location []string `json:"location"`
	// Extensions for location
	Location_ext []*Element `json:"_location"`
	// Extensions for expression
	Expression_ext []*Element `json:"_expression"`
	// Describes the type of the issue. The system that creates an OperationOutcome SHALL choose the most applicable code from the IssueType value set, and may additional provide its own code for the error in the details element.
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// Additional details about the error. This may be a text description of the error, or a system code that identifies the error.
	Details *CodeableConcept `json:"details"`
	// Additional diagnostic information about the issue.  Typically, this may be a description of how a value is erroneous, or a stack dump to help trace the issue.
	Diagnostics string `json:"diagnostics"`
	// A simple FHIRPath limited to element names, repetition indicators and the default child access that identifies one of the elements in the resource that caused this issue to be raised.
	Expression []string `json:"expression"`
}

// StructureMap_Dependent is A Map of relationships between 2 structures that can be used to transform data.
type StructureMap_Dependent struct {
	_ *BackboneElement
	// Variable to pass to the rule or group.
	Variable []string `json:"variable"`
	// Extensions for variable
	Variable_ext []*Element `json:"_variable"`
	// Name of a rule or group to apply.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
}

// Consent_Data1 is A record of a healthcare consumer's policy choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent_Data1 struct {
	_ *BackboneElement
	// How the resource reference is interpreted when testing consent restrictions.
	Meaning string `json:"meaning"`
	// Extensions for meaning
	Meaning_ext *Element `json:"_meaning"`
	// A reference to a specific resource that defines which resources are covered by this consent.
	Reference *Reference `json:"reference,omitempty"`
}

// EpisodeOfCare_StatusHistory is An association between a patient and an organization / healthcare provider(s) during which time encounters may occur. The managing organization assumes a level of responsibility for the patient during this time.
type EpisodeOfCare_StatusHistory struct {
	_ *BackboneElement
	// The period during this EpisodeOfCare that the specific status applied.
	Period *Period `json:"period,omitempty"`
	// planned | waitlist | active | onhold | finished | cancelled.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
}

// Group_Characteristic is Represents a defined collection of entities that may be discussed or acted upon collectively but which are not expected to act collectively and are not formally or legally recognized; i.e. a collection of entities that isn't an Organization.
type Group_Characteristic struct {
	_ *BackboneElement
	// The period over which the characteristic is tested; e.g. the patient had an operation during the month of June.
	Period *Period `json:"period"`
	// A code that identifies the kind of trait being asserted.
	Code *CodeableConcept `json:"code,omitempty"`
	// The value of the trait that holds (or does not hold - see 'exclude') for members of the group.
	ValueBoolean bool `json:"valueBoolean"`
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// If true, indicates the characteristic is one that is NOT held by members of the group.
	Exclude bool `json:"exclude"`
	// Extensions for exclude
	Exclude_ext *Element `json:"_exclude"`
	// The value of the trait that holds (or does not hold - see 'exclude') for members of the group.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	// The value of the trait that holds (or does not hold - see 'exclude') for members of the group.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// The value of the trait that holds (or does not hold - see 'exclude') for members of the group.
	ValueRange *Range `json:"valueRange"`
}

// MedicationRequest_DispenseRequest is An order or request for both supply of the medication and the instructions for administration of the medication to a patient. The resource is called "MedicationRequest" rather than "MedicationPrescription" or "MedicationOrder" to generalize the use across inpatient and outpatient settings, including care plans, etc., and to harmonize with workflow patterns.
type MedicationRequest_DispenseRequest struct {
	_ *BackboneElement
	// An integer indicating the number of times, in addition to the original dispense, (aka refills or repeats) that the patient can receive the prescribed medication. Usage Notes: This integer does not include the original order dispense. This means that if an order indicates dispense 30 tablets plus "3 repeats", then the order can be dispensed a total of 4 times and the patient can receive a total of 120 tablets.
	// pattern [1-9][0-9]*
	NumberOfRepeatsAllowed uint64 `json:"numberOfRepeatsAllowed"`
	// Extensions for numberOfRepeatsAllowed
	NumberOfRepeatsAllowed_ext *Element `json:"_numberOfRepeatsAllowed"`
	// The amount that is to be dispensed for one fill.
	Quantity *Quantity `json:"quantity"`
	// Identifies the period time over which the supplied product is expected to be used, or the length of time the dispense is expected to last.
	ExpectedSupplyDuration *Duration `json:"expectedSupplyDuration"`
	// Indicates the intended dispensing Organization specified by the prescriber.
	Performer *Reference `json:"performer"`
	// This indicates the validity period of a prescription (stale dating the Prescription).
	ValidityPeriod *Period `json:"validityPeriod"`
}

// NutritionOrder_Nutrient is A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrder_Nutrient struct {
	_ *BackboneElement
	// The nutrient that is being modified such as carbohydrate or sodium.
	Modifier *CodeableConcept `json:"modifier"`
	// The quantity of the specified nutrient to include in diet.
	Amount *Quantity `json:"amount"`
}

// Provenance is Provenance of a resource is a record that describes entities and processes involved in producing and delivering or otherwise influencing that resource. Provenance provides a critical foundation for assessing authenticity, enabling trust, and allowing reproducibility. Provenance assertions are a form of contextual metadata and can themselves become important records with their own provenance. Provenance statement indicates clinical significance in terms of confidence in authenticity, reliability, and trustworthiness, integrity, and stage in lifecycle (e.g. Document Completion - has the artifact been legally authenticated), all of which may impact security, privacy, and trust policies.
type Provenance struct {
	_ *DomainResource
	// An actor taking a role in an activity  for which it can be assigned some degree of responsibility for the activity taking place.
	Agent []*Provenance_Agent `json:"agent,omitempty"`
	// A digital signature on the target Reference(s). The signer should match a Provenance.agent. The purpose of the signature is indicated.
	Signature []*Signature `json:"signature"`
	// The period during which the activity occurred.
	Period *Period `json:"period"`
	// The instant of time at which the activity was recorded.
	Recorded string `json:"recorded"`
	// Extensions for recorded
	Recorded_ext *Element `json:"_recorded"`
	// Extensions for policy
	Policy_ext []*Element `json:"_policy"`
	// An activity is something that occurs over a period of time and acts upon or with entities; it may include consuming, processing, transforming, modifying, relocating, using, or generating entities.
	Activity *Coding `json:"activity"`
	// An entity used in this activity.
	Entity []*Provenance_Entity `json:"entity"`
	// This is a Provenance resource
	ResourceType string `json:"resourceType,omitempty"`
	// The Reference(s) that were generated or updated by  the activity described in this resource. A provenance can point to more than one target if multiple resources were created/updated by the same activity.
	Target []*Reference `json:"target,omitempty"`
	// Policy or plan the activity was defined by. Typically, a single activity may have multiple applicable policy documents, such as patient consent, guarantor funding, etc.
	Policy []string `json:"policy"`
	// Where the activity occurred, if relevant.
	Location *Reference `json:"location"`
	// The reason that the activity was taking place.
	Reason []*Coding `json:"reason"`
}

func NewProvenance() *Provenance { return &Provenance{ResourceType: "Provenance"} }

// Encounter_ClassHistory is An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type Encounter_ClassHistory struct {
	_ *BackboneElement
	// The time that the episode was in the specified class.
	Period *Period `json:"period,omitempty"`
	// inpatient | outpatient | ambulatory | emergency +.
	Class *Coding `json:"class,omitempty"`
}

// MeasureReport_Population1 is The MeasureReport resource contains the results of evaluating a measure.
type MeasureReport_Population1 struct {
	_ *BackboneElement
	// The type of the population.
	Code *CodeableConcept `json:"code"`
	// The number of members of the population in this stratum.
	// pattern -?([0]|([1-9][0-9]*))
	Count int64 `json:"count"`
	// Extensions for count
	Count_ext *Element `json:"_count"`
	// This element refers to a List of patient level MeasureReport resources, one for each patient in this population in this stratum.
	Patients *Reference `json:"patients"`
	// The identifier of the population being reported, as defined by the population element of the measure.
	Identifier *Identifier `json:"identifier"`
}

// Medication_Content is This resource is primarily used for the identification and definition of a medication. It covers the ingredients and the packaging for a medication.
type Medication_Content struct {
	_ *BackboneElement
	// The amount of the product that is in the package.
	Amount *Quantity `json:"amount"`
	// Identifies one of the items in the package.
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept"`
	// Identifies one of the items in the package.
	ItemReference *Reference `json:"itemReference"`
}

// PaymentReconciliation_ProcessNote is This resource provides payment details and claim references supporting a bulk payment.
type PaymentReconciliation_ProcessNote struct {
	_ *BackboneElement
	// The note purpose: Print/Display.
	Type *CodeableConcept `json:"type"`
	// The note text.
	Text string `json:"text"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
}

// TestReport_Teardown is A summary of information based on the results of executing a TestScript.
type TestReport_Teardown struct {
	_ *BackboneElement
	// The teardown action will only contain an operation.
	Action []*TestReport_Action2 `json:"action,omitempty"`
}

// Claim_Information is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_Information struct {
	_ *BackboneElement
	// Sequence of the information element which serves to provide a link.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// The general class of the information supplied: information; exception; accident, employment; onset, etc.
	Category *CodeableConcept `json:"category,omitempty"`
	// The date when or period to which this information refers.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	TimingDate time.Time `json:"timingDate"`
	// The date when or period to which this information refers.
	TimingPeriod *Period `json:"timingPeriod"`
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	ValueAttachment *Attachment `json:"valueAttachment"`
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	ValueReference *Reference `json:"valueReference"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// System and code pertaining to the specific information regarding special conditions relating to the setting, treatment or patient  for which care is sought which may influence the adjudication.
	Code *CodeableConcept `json:"code"`
	// Extensions for timingDate
	TimingDate_ext *Element `json:"_timingDate"`
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	ValueString string `json:"valueString"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// For example, provides the reason for: the additional stay, or missing tooth or any other situation where a reason code is required in addition to the content.
	Reason *CodeableConcept `json:"reason"`
}

// SupplyRequest_Requester is A record of a request for a medication, substance or device used in the healthcare setting.
type SupplyRequest_Requester struct {
	_ *BackboneElement
	// The device, practitioner, etc. who initiated the request.
	Agent *Reference `json:"agent,omitempty"`
	// The organization the device or practitioner was acting on behalf of.
	OnBehalfOf *Reference `json:"onBehalfOf"`
}

// TestScript_Metadata is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Metadata struct {
	_ *BackboneElement
	// A link to the FHIR specification that this test is covering.
	Link []*TestScript_Link `json:"link"`
	// Capabilities that must exist and are assumed to function correctly on the FHIR server being tested.
	Capability []*TestScript_Capability `json:"capability,omitempty"`
}

// ValueSet_Parameter is A value set specifies a set of codes drawn from one or more code systems.
type ValueSet_Parameter struct {
	_ *BackboneElement
	// Extensions for valueInteger
	ValueInteger_ext *Element `json:"_valueInteger"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// The value of the parameter.
	// pattern -?([0]|([1-9][0-9]*))
	ValueInteger int64 `json:"valueInteger"`
	// Extensions for valueCode
	ValueCode_ext *Element `json:"_valueCode"`
	// The value of the parameter.
	ValueString string `json:"valueString"`
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// The name of the parameter.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for valueDecimal
	ValueDecimal_ext *Element `json:"_valueDecimal"`
	// The value of the parameter.
	ValueUri string `json:"valueUri"`
	// Extensions for valueUri
	ValueUri_ext *Element `json:"_valueUri"`
	// The value of the parameter.
	// pattern [^\s]+([\s]?[^\s]+)*
	ValueCode string `json:"valueCode"`
	// The value of the parameter.
	ValueBoolean bool `json:"valueBoolean"`
	// The value of the parameter.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	ValueDecimal float64 `json:"valueDecimal"`
}

// Timing is Specifies an event that may occur multiple times. Timing schedules are used to record when things are planned, expected or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
type Timing struct {
	_ *Element
	// A set of rules that describe when the event is scheduled.
	Repeat *Timing_Repeat `json:"repeat"`
	// A code for the timing schedule. Some codes such as BID are ubiquitous, but many institutions define their own additional codes. If a code is provided, the code is understood to be a complete statement of whatever is specified in the structured timing data, and either the code or the data may be used to interpret the Timing, with the exception that .repeat.bounds still applies over the code (and is not contained in the code).
	Code *CodeableConcept `json:"code"`
	// Identifies specific times when the event occurs.
	Event []time.Time `json:"event"`
	// Extensions for event
	Event_ext []*Element `json:"_event"`
}

// ElementDefinition_Slicing is Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition_Slicing struct {
	_ *BackboneElement
	// If the matching elements have to occur in the same order as defined in the profile.
	Ordered bool `json:"ordered"`
	// Extensions for ordered
	Ordered_ext *Element `json:"_ordered"`
	// Whether additional slices are allowed or not. When the slices are ordered, profile authors can also say that additional slices are only allowed at the end.
	Rules string `json:"rules"`
	// Extensions for rules
	Rules_ext *Element `json:"_rules"`
	// Designates which child elements are used to discriminate between the slices when processing an instance. If one or more discriminators are provided, the value of the child elements in the instance data SHALL completely distinguish which slice the element in the resource matches based on the allowed values for those elements in each of the slices.
	Discriminator []*ElementDefinition_Discriminator `json:"discriminator"`
	// A human-readable text description of how the slicing works. If there is no discriminator, this is required to be present to provide whatever information is possible about how the slices can be differentiated.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
}

// CommunicationRequest is A request to convey information; e.g. the CDS system proposes that an alert be sent to a responsible provider, the CDS system proposes that the public health agency be notified about a reportable condition.
type CommunicationRequest struct {
	_ *DomainResource
	// Characterizes how quickly the proposed act must be initiated. Includes concepts such as stat, urgent, routine.
	// pattern [^\s]+([\s]?[^\s]+)*
	Priority string `json:"priority"`
	// The entity (e.g. person, organization, clinical information system, device, group, or care team) which is the intended target of the communication.
	Recipient []*Reference `json:"recipient"`
	// The resources which were related to producing this communication request.
	Topic []*Reference `json:"topic"`
	// Extensions for occurrenceDateTime
	OccurrenceDateTime_ext *Element `json:"_occurrenceDateTime"`
	// Describes why the request is being made in coded or textual form.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Comments made about the request by the requester, sender, recipient, subject or other participants.
	Note []*Annotation `json:"note"`
	// This is a CommunicationRequest resource
	ResourceType string `json:"resourceType,omitempty"`
	// Completed or terminated request(s) whose function is taken by this new request.
	Replaces []*Reference `json:"replaces"`
	// A shared identifier common to all requests that were authorized more or less simultaneously by a single author, representing the identifier of the requisition, prescription or similar form.
	GroupIdentifier *Identifier `json:"groupIdentifier"`
	// The status of the proposal or order.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Extensions for priority
	Priority_ext *Element `json:"_priority"`
	// A channel that was used for this communication (e.g. email, fax).
	Medium []*CodeableConcept `json:"medium"`
	// The patient or group that is the focus of this communication request.
	Subject *Reference `json:"subject"`
	// The encounter or episode of care within which the communication request was created.
	Context *Reference `json:"context"`
	// A unique ID of this request for reference purposes. It must be provided if user wants it returned as part of any output, otherwise it will be autogenerated, if needed, by CDS system. Does not need to be the actual ID of the source system.
	Identifier []*Identifier `json:"identifier"`
	// The individual who initiated the request and has responsibility for its activation.
	Requester *CommunicationRequest_Requester `json:"requester"`
	// The entity (e.g. person, organization, clinical information system, or device) which is to be the source of the communication.
	Sender *Reference `json:"sender"`
	// The type of message to be sent such as alert, notification, reminder, instruction, etc.
	Category []*CodeableConcept `json:"category"`
	// Text, attachment(s), or resource(s) to be communicated to the recipient.
	Payload []*CommunicationRequest_Payload `json:"payload"`
	// The time when this communication is to occur.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OccurrenceDateTime time.Time `json:"occurrenceDateTime"`
	// The time when this communication is to occur.
	OccurrencePeriod *Period `json:"occurrencePeriod"`
	// For draft requests, indicates the date of initial creation.  For requests with other statuses, indicates the date of activation.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AuthoredOn time.Time `json:"authoredOn"`
	// Extensions for authoredOn
	AuthoredOn_ext *Element `json:"_authoredOn"`
	// Indicates another resource whose existence justifies this request.
	ReasonReference []*Reference `json:"reasonReference"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A plan or proposal that is fulfilled in whole or in part by this request.
	BasedOn []*Reference `json:"basedOn"`
}

func NewCommunicationRequest() *CommunicationRequest {
	return &CommunicationRequest{ResourceType: "CommunicationRequest"}
}

// ExpansionProfile_Designation2 is Resource to define constraints on the Expansion of a FHIR ValueSet.
type ExpansionProfile_Designation2 struct {
	_ *BackboneElement
	// The language this designation is defined for.
	// pattern [^\s]+([\s]?[^\s]+)*
	Language string `json:"language"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
	// Which kinds of designation to exclude from the expansion.
	Use *Coding `json:"use"`
}

// AllergyIntolerance is Risk of harmful or undesirable, physiological response which is unique to an individual and associated with exposure to a substance.
type AllergyIntolerance struct {
	_ *DomainResource
	// Code for an allergy or intolerance statement (either a positive or a negated/excluded statement).  This may be a code for a substance or pharmaceutical product that is considered to be responsible for the adverse reaction risk (e.g., "Latex"), an allergy or intolerance condition (e.g., "Latex allergy"), or a negated/excluded code for a specific substance or class (e.g., "No latex allergy") or a general or categorical negated statement (e.g.,  "No known allergy", "No known drug allergies").
	Code *CodeableConcept `json:"code"`
	// The date on which the existance of the AllergyIntolerance was first asserted or acknowledged.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AssertedDate time.Time `json:"assertedDate"`
	// Extensions for assertedDate
	AssertedDate_ext *Element `json:"_assertedDate"`
	// The source of the information about the allergy that is recorded.
	Asserter *Reference `json:"asserter"`
	// The clinical status of the allergy or intolerance.
	ClinicalStatus string `json:"clinicalStatus"`
	// Identification of the underlying physiological mechanism for the reaction risk.
	Type string `json:"type"`
	// Extensions for category
	Category_ext []*Element `json:"_category"`
	// Category of the identified substance.
	Category []string `json:"category"`
	// Estimated or actual date,  date-time, or age when allergy or intolerance was identified.
	OnsetAge *Age `json:"onsetAge"`
	// Estimated or actual date,  date-time, or age when allergy or intolerance was identified.
	OnsetPeriod *Period `json:"onsetPeriod"`
	// Estimated or actual date,  date-time, or age when allergy or intolerance was identified.
	OnsetRange *Range `json:"onsetRange"`
	// Extensions for lastOccurrence
	LastOccurrence_ext *Element `json:"_lastOccurrence"`
	// This is a AllergyIntolerance resource
	ResourceType string `json:"resourceType,omitempty"`
	// Individual who recorded the record and takes responsibility for its content.
	Recorder *Reference `json:"recorder"`
	// Extensions for clinicalStatus
	ClinicalStatus_ext *Element `json:"_clinicalStatus"`
	// Assertion about certainty associated with the propensity, or potential risk, of a reaction to the identified substance (including pharmaceutical product).
	VerificationStatus string `json:"verificationStatus"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The patient who has the allergy or intolerance.
	Patient *Reference `json:"patient,omitempty"`
	// Estimated or actual date,  date-time, or age when allergy or intolerance was identified.
	OnsetString string `json:"onsetString"`
	// Additional narrative about the propensity for the Adverse Reaction, not captured in other fields.
	Note []*Annotation `json:"note"`
	// Extensions for verificationStatus
	VerificationStatus_ext *Element `json:"_verificationStatus"`
	// Estimated or actual date,  date-time, or age when allergy or intolerance was identified.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OnsetDateTime time.Time `json:"onsetDateTime"`
	// Extensions for onsetString
	OnsetString_ext *Element `json:"_onsetString"`
	// Represents the date and/or time of the last known occurrence of a reaction event.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	LastOccurrence time.Time `json:"lastOccurrence"`
	// This records identifiers associated with this allergy/intolerance concern that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []*Identifier `json:"identifier"`
	// Estimate of the potential clinical harm, or seriousness, of the reaction to the identified substance.
	Criticality string `json:"criticality"`
	// Extensions for criticality
	Criticality_ext *Element `json:"_criticality"`
	// Extensions for onsetDateTime
	OnsetDateTime_ext *Element `json:"_onsetDateTime"`
	// Details about each adverse reaction event linked to exposure to the identified substance.
	Reaction []*AllergyIntolerance_Reaction `json:"reaction"`
}

func NewAllergyIntolerance() *AllergyIntolerance {
	return &AllergyIntolerance{ResourceType: "AllergyIntolerance"}
}

// CapabilityStatement_Software is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Software struct {
	_ *BackboneElement
	// Name software is known by.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The version identifier for the software covered by this statement.
	Version string `json:"version"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Date this version of the software was released.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ReleaseDate time.Time `json:"releaseDate"`
	// Extensions for releaseDate
	ReleaseDate_ext *Element `json:"_releaseDate"`
}

// Condition_Evidence is A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
type Condition_Evidence struct {
	_ *BackboneElement
	// Links to other relevant information, including pathology reports.
	Detail []*Reference `json:"detail"`
	// A manifestation or symptom that led to the recording of this condition.
	Code []*CodeableConcept `json:"code"`
}

// Encounter_Participant is An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type Encounter_Participant struct {
	_ *BackboneElement
	// Role of participant in encounter.
	Type []*CodeableConcept `json:"type"`
	// The period of time that the specified participant participated in the encounter. These can overlap or be sub-sets of the overall encounter's period.
	Period *Period `json:"period"`
	// Persons involved in the encounter other than the patient.
	Individual *Reference `json:"individual"`
}

// ExplanationOfBenefit_Diagnosis is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Diagnosis struct {
	_ *BackboneElement
	// Sequence of diagnosis which serves to provide a link.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// The diagnosis.
	DiagnosisCodeableConcept *CodeableConcept `json:"diagnosisCodeableConcept"`
	// The diagnosis.
	DiagnosisReference *Reference `json:"diagnosisReference"`
	// The type of the Diagnosis, for example: admitting, primary, secondary, discharge.
	Type []*CodeableConcept `json:"type"`
	// The package billing code, for example DRG, based on the assigned grouping code system.
	PackageCode *CodeableConcept `json:"packageCode"`
}

// Patient_Link is Demographics and other administrative information about an individual or animal receiving care or other health-related services.
type Patient_Link struct {
	_ *BackboneElement
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The other patient resource that the link refers to.
	Other *Reference `json:"other,omitempty"`
	// The type of link between this patient resource and another patient resource.
	Type string `json:"type"`
}

// PaymentNotice is This resource provides the status of the payment for goods and services rendered, and the request and response resource references.
type PaymentNotice struct {
	_ *DomainResource
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// The practitioner who is responsible for the services rendered to the patient.
	Provider *Reference `json:"provider"`
	// The payment status, typically paid: payment sent, cleared: payment received.
	PaymentStatus *CodeableConcept `json:"paymentStatus"`
	// The notice business identifier.
	Identifier []*Identifier `json:"identifier"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Reference of response to resource for which payment is being made.
	Response *Reference `json:"response"`
	// The date when this resource was created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Reference of resource for which payment is being made.
	Request *Reference `json:"request"`
	// The date when the above payment action occurrred.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	StatusDate time.Time `json:"statusDate"`
	// The Insurer who is target  of the request.
	Target *Reference `json:"target"`
	// This is a PaymentNotice resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for statusDate
	StatusDate_ext *Element `json:"_statusDate"`
	// The organization which is responsible for the services rendered to the patient.
	Organization *Reference `json:"organization"`
}

func NewPaymentNotice() *PaymentNotice { return &PaymentNotice{ResourceType: "PaymentNotice"} }

// ValueSet is A value set specifies a set of codes drawn from one or more code systems.
type ValueSet struct {
	_ *DomainResource
	// The identifier that is used to identify this version of the value set when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the value set author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// A natural language name identifying the value set. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// The name of the individual or organization that published the value set.
	Publisher string `json:"publisher"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// If this is set to 'true', then no new versions of the content logical definition can be created.  Note: Other metadata might still change.
	Immutable bool `json:"immutable"`
	// A value set can also be "expanded", where the value set is turned into a simple collection of enumerated codes. This element holds the expansion, if it has been performed.
	Expansion *ValueSet_Expansion `json:"expansion"`
	// This is a ValueSet resource
	ResourceType string `json:"resourceType,omitempty"`
	// An absolute URI that is used to identify this value set when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this value set is (or will be) published. The URL SHOULD include the major version of the value set. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// Extensions for extensible
	Extensible_ext *Element `json:"_extensible"`
	// A set of criteria that define the content logical definition of the value set by including or excluding codes from outside this value set. This I also known as the "Content Logical Definition" (CLD).
	Compose *ValueSet_Compose `json:"compose"`
	// A formal identifier that is used to identify this value set when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []*Identifier `json:"identifier"`
	// A copyright statement relating to the value set and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the value set.
	Copyright string `json:"copyright"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A boolean value to indicate that this value set is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// The date  (and optionally time) when the value set was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the value set changes. (e.g. the 'content logical definition').
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// A short, descriptive, user-friendly title for the value set.
	Title string `json:"title"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// A free text natural language description of the value set from a consumer's perspective.
	Description string `json:"description"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Whether this is intended to be used with an extensible binding or not.
	Extensible bool `json:"extensible"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// A legal or geographic region in which the value set is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Extensions for immutable
	Immutable_ext *Element `json:"_immutable"`
	// Explaination of why this value set is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// The status of this value set. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate value set instances.
	UseContext []*UsageContext `json:"useContext"`
}

func NewValueSet() *ValueSet { return &ValueSet{ResourceType: "ValueSet"} }

// AdverseEvent is Actual or  potential/avoided event causing unintended physical injury resulting from or contributed to by medical care, a research study or other healthcare setting factors that requires additional monitoring, treatment, or hospitalization, or that results in death.
type AdverseEvent struct {
	_ *DomainResource
	// Information on who recorded the adverse event.  May be the patient or a practitioner.
	Recorder *Reference `json:"recorder"`
	// This element defines the specific type of event that occurred or that was prevented from occurring.
	Type *CodeableConcept `json:"type"`
	// This subject or group impacted by the event.  With a prospective adverse event, there will be no subject as the adverse event was prevented.
	Subject *Reference `json:"subject"`
	// Includes information about the reaction that occurred as a result of exposure to a substance (for example, a drug or a chemical).
	Reaction []*Reference `json:"reaction"`
	// Describes the seriousness or severity of the adverse event.
	Seriousness *CodeableConcept `json:"seriousness"`
	// The information about where the adverse event occurred.
	Location *Reference `json:"location"`
	// Describes the type of outcome from the adverse event.
	Outcome *CodeableConcept `json:"outcome"`
	// Parties that may or should contribute or have contributed information to the Act. Such information includes information leading to the decision to perform the Act and how to perform the Act (e.g. consultant), information that the Act itself seeks to reveal (e.g. informant of clinical history), or information about what Act was performed (e.g. informant witness).
	EventParticipant *Reference `json:"eventParticipant"`
	// AdverseEvent.study.
	Study []*Reference `json:"study"`
	// This is a AdverseEvent resource
	ResourceType string `json:"resourceType,omitempty"`
	// The type of event which is important to characterize what occurred and caused harm to the subject, or had the potential to cause harm to the subject.
	Category string `json:"category"`
	// The date (and perhaps time) when the adverse event occurred.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Extensions for category
	Category_ext *Element `json:"_category"`
	// Describes the adverse event in text.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Describes the entity that is suspected to have caused the adverse event.
	SuspectEntity []*AdverseEvent_SuspectEntity `json:"suspectEntity"`
	// The identifier(s) of this adverse event that are assigned by business processes and/or used to refer to it when a direct URL reference to the resource itsefl is not appropriate.
	Identifier *Identifier `json:"identifier"`
	// AdverseEvent.subjectMedicalHistory.
	SubjectMedicalHistory []*Reference `json:"subjectMedicalHistory"`
	// AdverseEvent.referenceDocument.
	ReferenceDocument []*Reference `json:"referenceDocument"`
}

func NewAdverseEvent() *AdverseEvent { return &AdverseEvent{ResourceType: "AdverseEvent"} }

// Binary is A binary resource can contain any content, whether text, image, pdf, zip archive, etc.
type Binary struct {
	_ *Resource
	// This is a Binary resource
	ResourceType string `json:"resourceType,omitempty"`
	// MimeType of the binary content represented as a standard MimeType (BCP 13).
	// pattern [^\s]+([\s]?[^\s]+)*
	ContentType string `json:"contentType"`
	// Extensions for contentType
	ContentType_ext *Element `json:"_contentType"`
	// Treat this binary as if it was this other resource for access control purposes.
	SecurityContext *Reference `json:"securityContext"`
	// The actual content, base64 encoded.
	Content string `json:"content"`
	// Extensions for content
	Content_ext *Element `json:"_content"`
}

func NewBinary() *Binary { return &Binary{ResourceType: "Binary"} }

// ClinicalImpression_Investigation is A record of a clinical assessment performed to determine what problem(s) may affect the patient and before planning the treatments or management strategies that are best to manage a patient's condition. Assessments are often 1:1 with a clinical consultation / encounter,  but this varies greatly depending on the clinical workflow. This resource is called "ClinicalImpression" rather than "ClinicalAssessment" to avoid confusion with the recording of assessment tools such as Apgar score.
type ClinicalImpression_Investigation struct {
	_ *BackboneElement
	// A name/code for the group ("set") of investigations. Typically, this will be something like "signs", "symptoms", "clinical", "diagnostic", but the list is not constrained, and others such groups such as (exposure|family|travel|nutitirional) history may be used.
	Code *CodeableConcept `json:"code,omitempty"`
	// A record of a specific investigation that was undertaken.
	Item []*Reference `json:"item"`
}

// ExplanationOfBenefit_Adjudication is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Adjudication struct {
	_ *BackboneElement
	// Monitory amount associated with the code.
	Amount *Money `json:"amount"`
	// A non-monetary value for example a percentage. Mutually exclusive to the amount element above.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Value float64 `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// Code indicating: Co-Pay, deductable, elegible, benefit, tax, etc.
	Category *CodeableConcept `json:"category,omitempty"`
	// Adjudication reason such as limit reached.
	Reason *CodeableConcept `json:"reason"`
}

// NutritionOrder_OralDiet is A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrder_OralDiet struct {
	_ *BackboneElement
	// Class that defines the quantity and type of nutrient modifications (for example carbohydrate, fiber or sodium) required for the oral diet.
	Nutrient []*NutritionOrder_Nutrient `json:"nutrient"`
	// Class that describes any texture modifications required for the patient to safely consume various types of solid foods.
	Texture []*NutritionOrder_Texture `json:"texture"`
	// The required consistency (e.g. honey-thick, nectar-thick, thin, thickened.) of liquids or fluids served to the patient.
	FluidConsistencyType []*CodeableConcept `json:"fluidConsistencyType"`
	// Free text or additional instructions or information pertaining to the oral diet.
	Instruction string `json:"instruction"`
	// Extensions for instruction
	Instruction_ext *Element `json:"_instruction"`
	// The kind of diet or dietary restriction such as fiber restricted diet or diabetic diet.
	Type []*CodeableConcept `json:"type"`
	// The time period and frequency at which the diet should be given.  The diet should be given for the combination of all schedules if more than one schedule is present.
	Schedule []*Timing `json:"schedule"`
}

// ValueSet_Include is A value set specifies a set of codes drawn from one or more code systems.
type ValueSet_Include struct {
	_ *BackboneElement
	// Extensions for valueSet
	ValueSet_ext []*Element `json:"_valueSet"`
	// An absolute URI which is the code system from which the selected codes come from.
	System string `json:"system"`
	// Extensions for system
	System_ext *Element `json:"_system"`
	// The version of the code system that the codes are selected from.
	Version string `json:"version"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Specifies a concept to be included or excluded.
	Concept []*ValueSet_Concept `json:"concept"`
	// Select concepts by specify a matching criteria based on the properties (including relationships) defined by the system. If multiple filters are specified, they SHALL all be true.
	Filter []*ValueSet_Filter `json:"filter"`
	// Selects concepts found in this value set. This is an absolute URI that is a reference to ValueSet.url.
	ValueSet []string `json:"valueSet"`
}

// Claim_Diagnosis is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_Diagnosis struct {
	_ *BackboneElement
	// The type of the Diagnosis, for example: admitting, primary, secondary, discharge.
	Type []*CodeableConcept `json:"type"`
	// The package billing code, for example DRG, based on the assigned grouping code system.
	PackageCode *CodeableConcept `json:"packageCode"`
	// Sequence of diagnosis which serves to provide a link.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// The diagnosis.
	DiagnosisCodeableConcept *CodeableConcept `json:"diagnosisCodeableConcept"`
	// The diagnosis.
	DiagnosisReference *Reference `json:"diagnosisReference"`
}

// Claim_Item is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_Item struct {
	_ *BackboneElement
	// If this is an actual service or product line, ie. not a Group, then use code to indicate the Professional Service or Product supplied (eg. CTP, HCPCS,USCLS,ICD10, NCPDP,DIN,RXNorm,ACHI,CCI). If a grouping item then use a group code to indicate the type of thing being grouped eg. 'glasses' or 'compound'.
	Service *CodeableConcept `json:"service"`
	// Where the service was provided.
	LocationCodeableConcept *CodeableConcept `json:"locationCodeableConcept"`
	// The quantity times the unit price for an addittional service or product or charge. For example, the formula: unit Quantity * unit Price (Cost per Point) * factor Number  * points = net Amount. Quantity, factor and points are assumed to be 1 if not supplied.
	Net *Money `json:"net"`
	// Extensions for procedureLinkId
	ProcedureLinkId_ext []*Element `json:"_procedureLinkId"`
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Factor float64 `json:"factor"`
	// Extensions for informationLinkId
	InformationLinkId_ext []*Element `json:"_informationLinkId"`
	// Second tier of goods and services.
	Detail []*Claim_Detail `json:"detail"`
	// Procedures applicable for this service or product line.
	ProcedureLinkId []uint64 `json:"procedureLinkId"`
	// The type of reveneu or cost center providing the product and/or service.
	Revenue *CodeableConcept `json:"revenue"`
	// The number of repetitions of a service or product.
	Quantity *Quantity `json:"quantity"`
	// Extensions for factor
	Factor_ext *Element `json:"_factor"`
	// A service line number.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// The date or dates when the enclosed suite of services were performed or completed.
	ServicedPeriod *Period `json:"servicedPeriod"`
	// Exceptions, special conditions and supporting information pplicable for this service or product line.
	InformationLinkId []uint64 `json:"informationLinkId"`
	// For programs which require reason codes for the inclusion or covering of this billed item under the program or sub-program.
	ProgramCode []*CodeableConcept `json:"programCode"`
	// List of Unique Device Identifiers associated with this line item.
	Udi []*Reference `json:"udi"`
	// Physical service site on the patient (limb, tooth, etc).
	BodySite *CodeableConcept `json:"bodySite"`
	// Diagnosis applicable for this service or product line.
	DiagnosisLinkId []uint64 `json:"diagnosisLinkId"`
	// Extensions for diagnosisLinkId
	DiagnosisLinkId_ext []*Element `json:"_diagnosisLinkId"`
	// Item typification or modifiers codes, eg for Oral whether the treatment is cosmetic or associated with TMJ, or for medical whether the treatment was outside the clinic or out of office hours.
	Modifier []*CodeableConcept `json:"modifier"`
	// Extensions for servicedDate
	ServicedDate_ext *Element `json:"_servicedDate"`
	// Where the service was provided.
	LocationAddress *Address `json:"locationAddress"`
	// Where the service was provided.
	LocationReference *Reference `json:"locationReference"`
	// If the item is a node then this is the fee for the product or service, otherwise this is the total of the fees for the children of the group.
	UnitPrice *Money `json:"unitPrice"`
	// A region or surface of the site, eg. limb region or tooth surface(s).
	SubSite []*CodeableConcept `json:"subSite"`
	// CareTeam applicable for this service or product line.
	CareTeamLinkId []uint64 `json:"careTeamLinkId"`
	// Extensions for careTeamLinkId
	CareTeamLinkId_ext []*Element `json:"_careTeamLinkId"`
	// The date or dates when the enclosed suite of services were performed or completed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ServicedDate time.Time `json:"servicedDate"`
	// A billed item may include goods or services provided in multiple encounters.
	Encounter []*Reference `json:"encounter"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// Health Care Service Type Codes  to identify the classification of service or benefits.
	Category *CodeableConcept `json:"category"`
}

// Claim_Detail is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_Detail struct {
	_ *BackboneElement
	// Extensions for factor
	Factor_ext *Element `json:"_factor"`
	// Third tier of goods and services.
	SubDetail []*Claim_SubDetail `json:"subDetail"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Factor float64 `json:"factor"`
	// List of Unique Device Identifiers associated with this line item.
	Udi []*Reference `json:"udi"`
	// A service line number.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// The type of reveneu or cost center providing the product and/or service.
	Revenue *CodeableConcept `json:"revenue"`
	// Health Care Service Type Codes  to identify the classification of service or benefits.
	Category *CodeableConcept `json:"category"`
	// If this is an actual service or product line, ie. not a Group, then use code to indicate the Professional Service or Product supplied (eg. CTP, HCPCS,USCLS,ICD10, NCPDP,DIN,ACHI,CCI). If a grouping item then use a group code to indicate the type of thing being grouped eg. 'glasses' or 'compound'.
	Service *CodeableConcept `json:"service"`
	// For programs which require reson codes for the inclusion, covering, of this billed item under the program or sub-program.
	ProgramCode []*CodeableConcept `json:"programCode"`
	// The number of repetitions of a service or product.
	Quantity *Quantity `json:"quantity"`
	// Item typification or modifiers codes, eg for Oral whether the treatment is cosmetic or associated with TMJ, or for medical whether the treatment was outside the clinic or out of office hours.
	Modifier []*CodeableConcept `json:"modifier"`
	// If the item is a node then this is the fee for the product or service, otherwise this is the total of the fees for the children of the group.
	UnitPrice *Money `json:"unitPrice"`
	// The quantity times the unit price for an addittional service or product or charge. For example, the formula: unit Quantity * unit Price (Cost per Point) * factor Number  * points = net Amount. Quantity, factor and points are assumed to be 1 if not supplied.
	Net *Money `json:"net"`
}

// Measure_Stratifier is The Measure resource provides the definition of a quality measure.
type Measure_Stratifier struct {
	_ *BackboneElement
	// The path to an element that defines the stratifier, specified as a valid FHIR resource path.
	Path string `json:"path"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// The identifier for the stratifier used to coordinate the reported data back to this stratifier.
	Identifier *Identifier `json:"identifier"`
	// The criteria for the stratifier. This must be the name of an expression defined within a referenced library.
	Criteria string `json:"criteria"`
	// Extensions for criteria
	Criteria_ext *Element `json:"_criteria"`
}

// Sequence is Raw data describing a biological sequence.
type Sequence struct {
	_ *DomainResource
	// The definition of variant here originates from Sequence ontology ([variant_of](http://www.sequenceontology.org/browser/current_svn/term/variant_of)). This element can represent amino acid or nucleic sequence change(including insertion,deletion,SNP,etc.)  It can represent some complex mutation or segment variation with the assist of CIGAR string.
	Variant []*Sequence_Variant `json:"variant"`
	// Extensions for readCoverage
	ReadCoverage_ext *Element `json:"_readCoverage"`
	// Whether the sequence is numbered starting at 0 (0-based numbering or coordinates, inclusive start, exclusive end) or starting at 1 (1-based numbering, inclusive start and inclusive end).
	// pattern -?([0]|([1-9][0-9]*))
	CoordinateSystem int64 `json:"coordinateSystem"`
	// Extensions for coordinateSystem
	CoordinateSystem_ext *Element `json:"_coordinateSystem"`
	// A sequence that is used as a reference to describe variants that are present in a sequence analyzed.
	ReferenceSeq *Sequence_ReferenceSeq `json:"referenceSeq"`
	// The method for sequencing, for example, chip information.
	Device *Reference `json:"device"`
	// The number of copies of the seqeunce of interest. (RNASeq).
	Quantity *Quantity `json:"quantity"`
	// Extensions for observedSeq
	ObservedSeq_ext *Element `json:"_observedSeq"`
	// Configurations of the external repository. The repository shall store target's observedSeq or records related with target's observedSeq.
	Repository []*Sequence_Repository `json:"repository"`
	// Pointer to next atomic sequence which at most contains one variant.
	Pointer []*Reference `json:"pointer"`
	// This is a Sequence resource
	ResourceType string `json:"resourceType,omitempty"`
	// Amino Acid Sequence/ DNA Sequence / RNA Sequence.
	Type string `json:"type"`
	// Specimen used for sequencing.
	Specimen *Reference `json:"specimen"`
	// The patient whose sequencing results are described by this resource.
	Patient *Reference `json:"patient"`
	// An experimental feature attribute that defines the quality of the feature in a quantitative way, such as a phred quality score ([SO:0001686](http://www.sequenceontology.org/browser/current_svn/term/SO:0001686)).
	Quality []*Sequence_Quality `json:"quality"`
	// Coverage (read depth or depth) is the average number of reads representing a given nucleotide in the reconstructed sequence.
	// pattern -?([0]|([1-9][0-9]*))
	ReadCoverage int64 `json:"readCoverage"`
	// Sequence that was observed. It is the result marked by referenceSeq along with variant records on referenceSeq. This shall starts from referenceSeq.windowStart and end by referenceSeq.windowEnd.
	ObservedSeq string `json:"observedSeq"`
	// A unique identifier for this particular sequence instance. This is a FHIR-defined id.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The organization or lab that should be responsible for this result.
	Performer *Reference `json:"performer"`
}

func NewSequence() *Sequence { return &Sequence{ResourceType: "Sequence"} }

// Claim_Related is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_Related struct {
	_ *BackboneElement
	// Other claims which are related to this claim such as prior claim versions or for related services.
	Claim *Reference `json:"claim"`
	// For example prior or umbrella.
	Relationship *CodeableConcept `json:"relationship"`
	// An alternate organizational reference to the case or file to which this particular claim pertains - eg Property/Casualy insurer claim # or Workers Compensation case # .
	Reference *Identifier `json:"reference"`
}

// StructureDefinition_Mapping is A definition of a FHIR structure. This resource is used to describe the underlying resources, data types defined in FHIR, and also for describing extensions and constraints on resources and data types.
type StructureDefinition_Mapping struct {
	_ *BackboneElement
	// Extensions for uri
	Uri_ext *Element `json:"_uri"`
	// A name for the specification that is being mapped to.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Comments about this mapping, including version notes, issues, scope limitations, and other important notes for usage.
	Comment string `json:"comment"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// An Internal id that is used to identify this mapping set when specific mappings are made.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Identity string `json:"identity"`
	// Extensions for identity
	Identity_ext *Element `json:"_identity"`
	// An absolute URI that identifies the specification that this mapping is expressed to.
	Uri string `json:"uri"`
}

// ElementDefinition_Base is Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition_Base struct {
	_ *BackboneElement
	// Extensions for min
	Min_ext *Element `json:"_min"`
	// Maximum cardinality of the base element identified by the path.
	Max string `json:"max"`
	// Extensions for max
	Max_ext *Element `json:"_max"`
	// The Path that identifies the base element - this matches the ElementDefinition.path for that element. Across FHIR, there is only one base definition of any element - that is, an element definition on a [[[StructureDefinition]]] without a StructureDefinition.base.
	Path string `json:"path"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// Minimum cardinality of the base element identified by the path.
	// pattern [0]|([1-9][0-9]*)
	Min uint64 `json:"min"`
}

// DataRequirement is Describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type DataRequirement struct {
	_ *Element
	// Extensions for mustSupport
	MustSupport_ext []*Element `json:"_mustSupport"`
	// Code filters specify additional constraints on the data, specifying the value set of interest for a particular element of the data.
	CodeFilter []*DataRequirement_CodeFilter `json:"codeFilter"`
	// Date filters specify additional constraints on the data in terms of the applicable date range for specific elements.
	DateFilter []*DataRequirement_DateFilter `json:"dateFilter"`
	// The type of the required data, specified as the type name of a resource. For profiles, this value is set to the type of the base resource of the profile.
	// pattern [^\s]+([\s]?[^\s]+)*
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The profile of the required data, specified as the uri of the profile definition.
	Profile []string `json:"profile"`
	// Extensions for profile
	Profile_ext []*Element `json:"_profile"`
	// Indicates that specific elements of the type are referenced by the knowledge module and must be supported by the consumer in order to obtain an effective evaluation. This does not mean that a value is required for this element, only that the consuming system must understand the element and be able to provide values for it if they are available. Note that the value for this element can be a path to allow references to nested elements. In that case, all the elements along the path must be supported.
	MustSupport []string `json:"mustSupport"`
}

// ActivityDefinition_Participant is This resource allows for the definition of some activity to be performed, independent of a particular patient, practitioner, or other performance context.
type ActivityDefinition_Participant struct {
	_ *BackboneElement
	// The type of participant in the action.
	// pattern [^\s]+([\s]?[^\s]+)*
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The role the participant should play in performing the described action.
	Role *CodeableConcept `json:"role"`
}

// DocumentManifest_Content is A collection of documents compiled for a purpose together with metadata that applies to the collection.
type DocumentManifest_Content struct {
	_ *BackboneElement
	// The list of references to document content, or Attachment that consist of the parts of this document manifest. Usually, these would be document references, but direct references to Media or Attachments are also allowed.
	PAttachment *Attachment `json:"pAttachment"`
	// The list of references to document content, or Attachment that consist of the parts of this document manifest. Usually, these would be document references, but direct references to Media or Attachments are also allowed.
	PReference *Reference `json:"pReference"`
}

// ProcedureRequest_Requester is A record of a request for diagnostic investigations, treatments, or operations to be performed.
type ProcedureRequest_Requester struct {
	_ *BackboneElement
	// The device, practitioner or organization who initiated the request.
	Agent *Reference `json:"agent,omitempty"`
	// The organization the device or practitioner was acting on behalf of.
	OnBehalfOf *Reference `json:"onBehalfOf"`
}

// TestReport is A summary of information based on the results of executing a TestScript.
type TestReport struct {
	_ *DomainResource
	// The current state of this test report.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for score
	Score_ext *Element `json:"_score"`
	// Name of the tester producing this report (Organization or individual).
	Tester string `json:"tester"`
	// When the TestScript was executed and this TestReport was generated.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Issued time.Time `json:"issued"`
	// This is a TestReport resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for result
	Result_ext *Element `json:"_result"`
	// The final score (percentage of tests passed) resulting from the execution of the TestScript.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Score float64 `json:"score"`
	// Extensions for issued
	Issued_ext *Element `json:"_issued"`
	// The results of the series of required setup operations before the tests were executed.
	Setup *TestReport_Setup `json:"setup"`
	// Identifier for the TestScript assigned for external purposes outside the context of FHIR.
	Identifier *Identifier `json:"identifier"`
	// Ideally this is an absolute URL that is used to identify the version-specific TestScript that was executed, matching the `TestScript.url`.
	TestScript *Reference `json:"testScript,omitempty"`
	// The overall result from the execution of the TestScript.
	Result string `json:"result"`
	// A participant in the test execution, either the execution engine, a client, or a server.
	Participant []*TestReport_Participant `json:"participant"`
	// The results of the series of operations required to clean up after the all the tests were executed (successfully or otherwise).
	Teardown *TestReport_Teardown `json:"teardown"`
	// A free text natural language name identifying the executed TestScript.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for tester
	Tester_ext *Element `json:"_tester"`
	// A test executed from the test script.
	Test []*TestReport_Test `json:"test"`
}

func NewTestReport() *TestReport { return &TestReport{ResourceType: "TestReport"} }

// Ratio is A relationship of two Quantity values - expressed as a numerator and a denominator.
type Ratio struct {
	_ *Element
	// The value of the numerator.
	Numerator *Quantity `json:"numerator"`
	// The value of the denominator.
	Denominator *Quantity `json:"denominator"`
}

// BodySite is Record details about the anatomical location of a specimen or body part.  This resource may be used when a coded concept does not provide the necessary detail needed for the use case.
type BodySite struct {
	_ *DomainResource
	// Whether this body site is in active use.
	Active bool `json:"active"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
	// Qualifier to refine the anatomical location.  These include qualifiers for laterality, relative location, directionality, number, and plane.
	Qualifier []*CodeableConcept `json:"qualifier"`
	// Image or images used to identify a location.
	Image []*Attachment `json:"image"`
	// This is a BodySite resource
	ResourceType string `json:"resourceType,omitempty"`
	// Identifier for this instance of the anatomical location.
	Identifier []*Identifier `json:"identifier"`
	// Named anatomical location - ideally coded where possible.
	Code *CodeableConcept `json:"code"`
	// A summary, charactarization or explanation of the anatomic location.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The person to which the body site belongs.
	Patient *Reference `json:"patient,omitempty"`
}

func NewBodySite() *BodySite { return &BodySite{ResourceType: "BodySite"} }

// ConceptMap is A statement of relationships from one set of concepts to one or more other concepts - either code systems or data elements, or classes in class models.
type ConceptMap struct {
	_ *DomainResource
	// A legal or geographic region in which the concept map is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// The source value set that specifies the concepts that are being mapped.
	SourceReference *Reference `json:"sourceReference"`
	// Extensions for targetUri
	TargetUri_ext *Element `json:"_targetUri"`
	// A short, descriptive, user-friendly title for the concept map.
	Title string `json:"title"`
	// The status of this concept map. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// A free text natural language description of the concept map from a consumer's perspective.
	Description string `json:"description"`
	// The source value set that specifies the concepts that are being mapped.
	SourceUri string `json:"sourceUri"`
	// A boolean value to indicate that this concept map is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Explaination of why this concept map is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// A formal identifier that is used to identify this concept map when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier *Identifier `json:"identifier"`
	// The date  (and optionally time) when the concept map was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the concept map changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// The name of the individual or organization that published the concept map.
	Publisher string `json:"publisher"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate concept map instances.
	UseContext []*UsageContext `json:"useContext"`
	// Extensions for sourceUri
	SourceUri_ext *Element `json:"_sourceUri"`
	// An absolute URI that is used to identify this concept map when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this concept map is (or will be) published. The URL SHOULD include the major version of the concept map. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// The identifier that is used to identify this version of the concept map when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the concept map author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// The target value set provides context to the mappings. Note that the mapping is made between concepts, not between value sets, but the value set provides important context about how the concept mapping choices are made.
	TargetUri string `json:"targetUri"`
	// A natural language name identifying the concept map. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// A group of mappings that all have the same source and target system.
	Group []*ConceptMap_Group `json:"group"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// A copyright statement relating to the concept map and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the concept map.
	Copyright string `json:"copyright"`
	// The target value set provides context to the mappings. Note that the mapping is made between concepts, not between value sets, but the value set provides important context about how the concept mapping choices are made.
	TargetReference *Reference `json:"targetReference"`
	// This is a ConceptMap resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
}

func NewConceptMap() *ConceptMap { return &ConceptMap{ResourceType: "ConceptMap"} }

// EligibilityResponse is This resource provides eligibility and plan details from the processing of an Eligibility resource.
type EligibilityResponse struct {
	_ *DomainResource
	// The date when the enclosed suite of services were performed or completed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// The practitioner who is responsible for the services rendered to the patient.
	RequestProvider *Reference `json:"requestProvider"`
	// A description of the status of the adjudication.
	Disposition string `json:"disposition"`
	// Extensions for disposition
	Disposition_ext *Element `json:"_disposition"`
	// The form to be used for printing the content.
	Form *CodeableConcept `json:"form"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The Response business identifier.
	Identifier []*Identifier `json:"identifier"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Original request resource reference.
	Request *Reference `json:"request"`
	// The Insurer who produced this adjudicated response.
	Insurer *Reference `json:"insurer"`
	// Extensions for inforce
	Inforce_ext *Element `json:"_inforce"`
	// The insurer may provide both the details for the requested coverage as well as details for additional coverages known to the insurer.
	Insurance []*EligibilityResponse_Insurance `json:"insurance"`
	// Mutually exclusive with Services Provided (Item).
	Error []*EligibilityResponse_Error `json:"error"`
	// This is a EligibilityResponse resource
	ResourceType string `json:"resourceType,omitempty"`
	// Flag indicating if the coverage provided is inforce currently  if no service date(s) specified or for the whole duration of the service dates.
	Inforce bool `json:"inforce"`
	// The organization which is responsible for the services rendered to the patient.
	RequestOrganization *Reference `json:"requestOrganization"`
	// Transaction status: error, complete.
	Outcome *CodeableConcept `json:"outcome"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
}

func NewEligibilityResponse() *EligibilityResponse {
	return &EligibilityResponse{ResourceType: "EligibilityResponse"}
}

// FamilyMemberHistory is Significant health events and conditions for a person related to the patient relevant in the context of care for the patient.
type FamilyMemberHistory struct {
	_ *DomainResource
	// If true, indicates the taking of an individual family member's history did not occur. The notDone element should not be used to document negated conditions, such as a family member that did not have a condition.
	NotDone bool `json:"notDone"`
	// The actual or approximate date of birth of the relative.
	BornString string `json:"bornString"`
	// A protocol or questionnaire that was adhered to in whole or in part by this event.
	Definition []*Reference `json:"definition"`
	// A code specifying the status of the record of the family history of a specific family member.
	Status string `json:"status"`
	// Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
	DeceasedAge *Age `json:"deceasedAge"`
	// Extensions for bornString
	BornString_ext *Element `json:"_bornString"`
	// The age of the relative at the time the family member history is recorded.
	AgeAge *Age `json:"ageAge"`
	// The age of the relative at the time the family member history is recorded.
	AgeString string `json:"ageString"`
	// Extensions for deceasedDate
	DeceasedDate_ext *Element `json:"_deceasedDate"`
	// Extensions for notDone
	NotDone_ext *Element `json:"_notDone"`
	// The actual or approximate date of birth of the relative.
	BornPeriod *Period `json:"bornPeriod"`
	// Extensions for estimatedAge
	EstimatedAge_ext *Element `json:"_estimatedAge"`
	// Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
	DeceasedRange *Range `json:"deceasedRange"`
	// Extensions for deceasedString
	DeceasedString_ext *Element `json:"_deceasedString"`
	// Describes why the family member history occurred in coded or textual form.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Extensions for ageString
	AgeString_ext *Element `json:"_ageString"`
	// The person who this history concerns.
	Patient *Reference `json:"patient,omitempty"`
	// This will either be a name or a description; e.g. "Aunt Susan", "my cousin with the red hair".
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// If true, indicates that the age value specified is an estimated value.
	EstimatedAge bool `json:"estimatedAge"`
	// Extensions for deceasedBoolean
	DeceasedBoolean_ext *Element `json:"_deceasedBoolean"`
	// This property allows a non condition-specific note to the made about the related person. Ideally, the note would be in the condition property, but this is not always possible.
	Note []*Annotation `json:"note"`
	// This is a FamilyMemberHistory resource
	ResourceType string `json:"resourceType,omitempty"`
	// Describes why the family member's history is absent.
	NotDoneReason *CodeableConcept `json:"notDoneReason"`
	// The date (and possibly time) when the family member history was taken.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Administrative Gender - the gender that the relative is considered to have for administration and record keeping purposes.
	Gender string `json:"gender"`
	// Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
	DeceasedString string `json:"deceasedString"`
	// Indicates a Condition, Observation, AllergyIntolerance, or QuestionnaireResponse that justifies this family member history event.
	ReasonReference []*Reference `json:"reasonReference"`
	// The significant Conditions (or condition) that the family member had. This is a repeating section to allow a system to represent more than one condition per resource, though there is nothing stopping multiple resources - one per condition.
	Condition []*FamilyMemberHistory_Condition `json:"condition"`
	// This records identifiers associated with this family member history record that are defined by business processes and/ or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []*Identifier `json:"identifier"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The age of the relative at the time the family member history is recorded.
	AgeRange *Range `json:"ageRange"`
	// Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	DeceasedDate time.Time `json:"deceasedDate"`
	// The type of relationship this person has to the patient (father, mother, brother etc.).
	Relationship *CodeableConcept `json:"relationship,omitempty"`
	// The actual or approximate date of birth of the relative.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	BornDate time.Time `json:"bornDate"`
	// Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
	DeceasedBoolean bool `json:"deceasedBoolean"`
	// Extensions for gender
	Gender_ext *Element `json:"_gender"`
	// Extensions for bornDate
	BornDate_ext *Element `json:"_bornDate"`
}

func NewFamilyMemberHistory() *FamilyMemberHistory {
	return &FamilyMemberHistory{ResourceType: "FamilyMemberHistory"}
}

// GraphDefinition_Target is A formal computable definition of a graph of resources - that is, a coherent set of resources that form a graph by following references. The Graph Definition resource defines a set and makes rules about the set.
type GraphDefinition_Target struct {
	_ *BackboneElement
	// Extensions for profile
	Profile_ext *Element `json:"_profile"`
	// Compartment Consistency Rules.
	Compartment []*GraphDefinition_Compartment `json:"compartment"`
	// Additional links from target resource.
	Link []*GraphDefinition_Link `json:"link"`
	// Type of resource this link refers to.
	// pattern [^\s]+([\s]?[^\s]+)*
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Profile for the target resource.
	Profile string `json:"profile"`
}

// Immunization is Describes the event of a patient being administered a vaccination or a record of a vaccination as reported by a patient, a clinician or another party and may include vaccine reaction information and what vaccination protocol was followed.
type Immunization struct {
	_ *DomainResource
	// Indicates who or what performed the event.
	Practitioner []*Immunization_Practitioner `json:"practitioner"`
	// Categorical data indicating that an adverse event is associated in time to an immunization.
	Reaction []*Immunization_Reaction `json:"reaction"`
	// The source of the data when the report of the immunization event is not based on information from the person who administered the vaccine.
	ReportOrigin *CodeableConcept `json:"reportOrigin"`
	// Name of vaccine manufacturer.
	Manufacturer *Reference `json:"manufacturer"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// An indication that the content of the record is based on information from the person who administered the vaccine. This reflects the context under which the data was originally recorded.
	PrimarySource bool `json:"primarySource"`
	// The service delivery location where the vaccine administration occurred.
	Location *Reference `json:"location"`
	// A unique identifier assigned to this immunization record.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for notGiven
	NotGiven_ext *Element `json:"_notGiven"`
	// The patient who either received or did not receive the immunization.
	Patient *Reference `json:"patient,omitempty"`
	// Date vaccine administered or was to be administered.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for expirationDate
	ExpirationDate_ext *Element `json:"_expirationDate"`
	// Reasons why a vaccine was or was not administered.
	Explanation *Immunization_Explanation `json:"explanation"`
	// Vaccine that was administered or was to be administered.
	VaccineCode *CodeableConcept `json:"vaccineCode,omitempty"`
	// The visit or admission or other contact between patient and health care provider the immunization was performed as part of.
	Encounter *Reference `json:"encounter"`
	// Lot number of the  vaccine product.
	LotNumber string `json:"lotNumber"`
	// The path by which the vaccine product is taken into the body.
	Route *CodeableConcept `json:"route"`
	// The quantity of vaccine product that was administered.
	DoseQuantity *Quantity `json:"doseQuantity"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Indicates if the vaccination was or was not given.
	NotGiven bool `json:"notGiven"`
	// Extra information about the immunization that is not conveyed by the other attributes.
	Note []*Annotation `json:"note"`
	// Contains information about the protocol(s) under which the vaccine was administered.
	VaccinationProtocol []*Immunization_VaccinationProtocol `json:"vaccinationProtocol"`
	// Extensions for primarySource
	PrimarySource_ext *Element `json:"_primarySource"`
	// Extensions for lotNumber
	LotNumber_ext *Element `json:"_lotNumber"`
	// Body site where vaccine was administered.
	Site *CodeableConcept `json:"site"`
	// This is a Immunization resource
	ResourceType string `json:"resourceType,omitempty"`
	// Indicates the current status of the vaccination event.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Date vaccine batch expires.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ExpirationDate time.Time `json:"expirationDate"`
}

func NewImmunization() *Immunization { return &Immunization{ResourceType: "Immunization"} }

// Medication_Ingredient is This resource is primarily used for the identification and definition of a medication. It covers the ingredients and the packaging for a medication.
type Medication_Ingredient struct {
	_ *BackboneElement
	// The actual ingredient - either a substance (simple ingredient) or another medication.
	ItemReference *Reference `json:"itemReference"`
	// Indication of whether this ingredient affects the therapeutic action of the drug.
	IsActive bool `json:"isActive"`
	// Extensions for isActive
	IsActive_ext *Element `json:"_isActive"`
	// Specifies how many (or how much) of the items there are in this Medication.  For example, 250 mg per tablet.  This is expressed as a ratio where the numerator is 250mg and the denominator is 1 tablet.
	Amount *Ratio `json:"amount"`
	// The actual ingredient - either a substance (simple ingredient) or another medication.
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept"`
}

// MedicationAdministration is Describes the event of a patient consuming or otherwise being administered a medication.  This may be as simple as swallowing a tablet or it may be a long running infusion.  Related resources tie this event to the authorizing prescription, and the specific encounter between patient and health care practitioner.
type MedicationAdministration struct {
	_ *DomainResource
	// This is a MedicationAdministration resource
	ResourceType string `json:"resourceType,omitempty"`
	// Indicates the type of medication administration and where the medication is expected to be consumed or administered.
	Category *CodeableConcept `json:"category"`
	// Identifies the medication that was administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	MedicationReference *Reference `json:"medicationReference"`
	// The person or animal or group receiving the medication.
	Subject *Reference `json:"subject,omitempty"`
	// Extensions for notGiven
	NotGiven_ext *Element `json:"_notGiven"`
	// A code indicating why the administration was not performed.
	ReasonNotGiven []*CodeableConcept `json:"reasonNotGiven"`
	// The original request, instruction or authority to perform the administration.
	Prescription *Reference `json:"prescription"`
	// Will generally be set to show that the administration has been completed.  For some long running administrations such as infusions it is possible for an administration to be started but not completed or it may be paused while some other process is under way.
	Status string `json:"status"`
	// Identifies the medication that was administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	MedicationCodeableConcept *CodeableConcept `json:"medicationCodeableConcept"`
	// Additional information (for example, patient height and weight) that supports the administration of the medication.
	SupportingInformation []*Reference `json:"supportingInformation"`
	// A specific date/time or interval of time during which the administration took place (or did not take place, when the 'notGiven' attribute is true). For many administrations, such as swallowing a tablet the use of dateTime is more appropriate.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	EffectiveDateTime time.Time `json:"effectiveDateTime"`
	// A specific date/time or interval of time during which the administration took place (or did not take place, when the 'notGiven' attribute is true). For many administrations, such as swallowing a tablet the use of dateTime is more appropriate.
	EffectivePeriod *Period `json:"effectivePeriod"`
	// The device used in administering the medication to the patient.  For example, a particular infusion pump.
	Device []*Reference `json:"device"`
	// Extra information about the medication administration that is not conveyed by the other attributes.
	Note []*Annotation `json:"note"`
	// Describes the medication dosage information details e.g. dose, rate, site, route, etc.
	Dosage *MedicationAdministration_Dosage `json:"dosage"`
	// External identifier - FHIR will generate its own internal identifiers (probably URLs) which do not need to be explicitly managed by the resource.  The identifier here is one that would be used by another non-FHIR system - for example an automated medication pump would provide a record each time it operated; an administration while the patient was off the ward might be made with a different system and entered after the event.  Particularly important if these records have to be updated.
	Identifier []*Identifier `json:"identifier"`
	// The visit, admission or other contact between patient and health care provider the medication administration was performed as part of.
	Context *Reference `json:"context"`
	// Set this to true if the record is saying that the medication was NOT administered.
	NotGiven bool `json:"notGiven"`
	// A code indicating why the medication was given.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Condition or observation that supports why the medication was administered.
	ReasonReference []*Reference `json:"reasonReference"`
	// A summary of the events of interest that have occurred, such as when the administration was verified.
	EventHistory []*Reference `json:"eventHistory"`
	// A protocol, guideline, orderset or other definition that was adhered to in whole or in part by this event.
	Definition []*Reference `json:"definition"`
	// A larger event of which this particular event is a component or step.
	PartOf []*Reference `json:"partOf"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for effectiveDateTime
	EffectiveDateTime_ext *Element `json:"_effectiveDateTime"`
	// The individual who was responsible for giving the medication to the patient.
	Performer []*MedicationAdministration_Performer `json:"performer"`
}

func NewMedicationAdministration() *MedicationAdministration {
	return &MedicationAdministration{ResourceType: "MedicationAdministration"}
}

// NamingSystem_UniqueId is A curated namespace that issues unique symbols within that namespace for the identification of concepts, people, devices, etc.  Represents a "System" used within the Identifier and Coding data types.
type NamingSystem_UniqueId struct {
	_ *BackboneElement
	// Identifies the unique identifier scheme used for this particular identifier.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The string that should be sent over the wire to identify the code system or identifier system.
	Value string `json:"value"`
	// Identifies the period of time over which this identifier is considered appropriate to refer to the naming system.  Outside of this window, the identifier might be non-deterministic.
	Period *Period `json:"period"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// Indicates whether this identifier is the "preferred" identifier of this type.
	Preferred bool `json:"preferred"`
	// Extensions for preferred
	Preferred_ext *Element `json:"_preferred"`
	// Notes about the past or intended usage of this identifier.
	Comment string `json:"comment"`
}

// ElementDefinition_Type is Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition_Type struct {
	_ *BackboneElement
	// Identifies a profile structure or implementation Guide that SHALL hold for the datatype this element refers to. Can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the resource SHALL conform to at least one profile defined in the implementation guide.
	Profile string `json:"profile"`
	// Identifies a profile structure or implementation Guide that SHALL hold for the target of the reference this element refers to. Can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the resource SHALL conform to at least one profile defined in the implementation guide.
	TargetProfile string `json:"targetProfile"`
	// Extensions for targetProfile
	TargetProfile_ext *Element `json:"_targetProfile"`
	// Extensions for aggregation
	Aggregation_ext []*Element `json:"_aggregation"`
	// Whether this reference needs to be version specific or version independent, or whether either can be used.
	Versioning string `json:"versioning"`
	// Extensions for versioning
	Versioning_ext *Element `json:"_versioning"`
	// URL of Data type or Resource that is a(or the) type used for this element. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition e.g. "string" is a reference to http://hl7.org/fhir/StructureDefinition/string. Absolute URLs are only allowed in logical models.
	Code string `json:"code"`
	// Extensions for profile
	Profile_ext *Element `json:"_profile"`
	// If the type is a reference to another resource, how the resource is or can be aggregated - is it a contained resource, or a reference, and if the context is a bundle, is it included in the bundle.
	Aggregation []string `json:"aggregation"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
}

// UsageContext is Specifies clinical/business/etc metadata that can be used to retrieve, index and/or categorize an artifact. This metadata can either be specific to the applicable population (e.g., age category, DRG) or the specific context of care (e.g., venue, care setting, provider of care).
type UsageContext struct {
	_ *Element
	// A code that identifies the type of context being specified by this usage context.
	Code *Coding `json:"code,omitempty"`
	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	ValueRange *Range `json:"valueRange"`
}

// ClaimResponse_Error is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse_Error struct {
	_ *BackboneElement
	// Extensions for subdetailSequenceLinkId
	SubdetailSequenceLinkId_ext *Element `json:"_subdetailSequenceLinkId"`
	// An error code,from a specified code system, which details why the claim could not be adjudicated.
	Code *CodeableConcept `json:"code,omitempty"`
	// The sequence number of the line item submitted which contains the error. This value is omitted when the error is elsewhere.
	// pattern [1-9][0-9]*
	SequenceLinkId uint64 `json:"sequenceLinkId"`
	// Extensions for sequenceLinkId
	SequenceLinkId_ext *Element `json:"_sequenceLinkId"`
	// The sequence number of the addition within the line item submitted which contains the error. This value is omitted when the error is not related to an Addition.
	// pattern [1-9][0-9]*
	DetailSequenceLinkId uint64 `json:"detailSequenceLinkId"`
	// Extensions for detailSequenceLinkId
	DetailSequenceLinkId_ext *Element `json:"_detailSequenceLinkId"`
	// The sequence number of the addition within the line item submitted which contains the error. This value is omitted when the error is not related to an Addition.
	// pattern [1-9][0-9]*
	SubdetailSequenceLinkId uint64 `json:"subdetailSequenceLinkId"`
}

// Consent_Policy is A record of a healthcare consumer's policy choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent_Policy struct {
	_ *BackboneElement
	// The references to the policies that are included in this consent scope. Policies may be organizational, but are often defined jurisdictionally, or in law.
	Uri string `json:"uri"`
	// Extensions for uri
	Uri_ext *Element `json:"_uri"`
	// Entity or Organization having regulatory jurisdiction or accountability for  enforcing policies pertaining to Consent Directives.
	Authority string `json:"authority"`
	// Extensions for authority
	Authority_ext *Element `json:"_authority"`
}

// Timing_Repeat is Specifies an event that may occur multiple times. Timing schedules are used to record when things are planned, expected or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
type Timing_Repeat struct {
	_ *BackboneElement
	// Extensions for periodUnit
	PeriodUnit_ext *Element `json:"_periodUnit"`
	// Extensions for timeOfDay
	TimeOfDay_ext []*Element `json:"_timeOfDay"`
	// The number of minutes from the event. If the event code does not indicate whether the minutes is before or after the event, then the offset is assumed to be after the event.
	// pattern [0]|([1-9][0-9]*)
	Offset uint64 `json:"offset"`
	// Indicates the duration of time over which repetitions are to occur; e.g. to express "3 times per day", 3 would be the frequency and "1 day" would be the period.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Period float64 `json:"period"`
	// A maximum value for the count of the desired repetitions (e.g. do something 6-8 times).
	// pattern -?([0]|([1-9][0-9]*))
	CountMax int64 `json:"countMax"`
	// The upper limit of how long this thing happens for when it happens.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	DurationMax float64 `json:"durationMax"`
	// The units of time for the duration, in UCUM units.
	DurationUnit string `json:"durationUnit"`
	// Extensions for frequency
	Frequency_ext *Element `json:"_frequency"`
	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsDuration *Duration `json:"boundsDuration"`
	// Extensions for count
	Count_ext *Element `json:"_count"`
	// Extensions for duration
	Duration_ext *Element `json:"_duration"`
	// Extensions for dayOfWeek
	DayOfWeek_ext []*Element `json:"_dayOfWeek"`
	// Extensions for period
	Period_ext *Element `json:"_period"`
	// If present, indicates that the period is a range from [period] to [periodMax], allowing expressing concepts such as "do this once every 3-5 days.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	PeriodMax float64 `json:"periodMax"`
	// Extensions for offset
	Offset_ext *Element `json:"_offset"`
	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsPeriod *Period `json:"boundsPeriod"`
	// A total count of the desired number of repetitions.
	// pattern -?([0]|([1-9][0-9]*))
	Count int64 `json:"count"`
	// Extensions for countMax
	CountMax_ext *Element `json:"_countMax"`
	// How long this thing happens for when it happens.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Duration float64 `json:"duration"`
	// The number of times to repeat the action within the specified period / period range (i.e. both period and periodMax provided).
	// pattern -?([0]|([1-9][0-9]*))
	Frequency int64 `json:"frequency"`
	// Extensions for frequencyMax
	FrequencyMax_ext *Element `json:"_frequencyMax"`
	// If one or more days of week is provided, then the action happens only on the specified day(s).
	DayOfWeek []string `json:"dayOfWeek"`
	// The units of time for the period in UCUM units.
	PeriodUnit string `json:"periodUnit"`
	// Specified time of day for action to take place.
	TimeOfDay []time.Time `json:"timeOfDay"`
	// Real world events that the occurrence of the event should be tied to.
	When []string `json:"when"`
	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	BoundsRange *Range `json:"boundsRange"`
	// Extensions for durationUnit
	DurationUnit_ext *Element `json:"_durationUnit"`
	// If present, indicates that the frequency is a range - so to repeat between [frequency] and [frequencyMax] times within the period or period range.
	// pattern -?([0]|([1-9][0-9]*))
	FrequencyMax int64 `json:"frequencyMax"`
	// Extensions for periodMax
	PeriodMax_ext *Element `json:"_periodMax"`
	// Extensions for durationMax
	DurationMax_ext *Element `json:"_durationMax"`
	// Extensions for when
	When_ext []*Element `json:"_when"`
}

// ActivityDefinition is This resource allows for the definition of some activity to be performed, independent of a particular patient, practitioner, or other performance context.
type ActivityDefinition struct {
	_ *DomainResource
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate activity definition instances.
	UseContext []*UsageContext `json:"useContext"`
	// The status of this activity definition. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// The period, timing or frequency upon which the described activity is to occur.
	TimingTiming *Timing `json:"timingTiming"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// A reference to a Library resource containing any formal logic used by the asset.
	Library []*Reference `json:"library"`
	// The period, timing or frequency upon which the described activity is to occur.
	TimingPeriod *Period `json:"timingPeriod"`
	// Extensions for lastReviewDate
	LastReviewDate_ext *Element `json:"_lastReviewDate"`
	// The date  (and optionally time) when the activity definition was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the activity definition changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Detailed description of the type of activity; e.g. What lab test, what procedure, what kind of encounter.
	Code *CodeableConcept `json:"code"`
	// Dynamic values that will be evaluated to produce values for elements of the resulting resource. For example, if the dosage of a medication must be computed based on the patient's weight, a dynamic value would be used to specify an expression that calculated the weight, and the path on the intent resource that would contain the result.
	DynamicValue []*ActivityDefinition_DynamicValue `json:"dynamicValue"`
	// A natural language name identifying the activity definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// A description of the kind of resource the activity definition is representing. For example, a MedicationRequest, a ProcedureRequest, or a CommunicationRequest. Typically, but not always, this is a Request resource.
	// pattern [^\s]+([\s]?[^\s]+)*
	Kind string `json:"kind"`
	// Extensions for timingDateTime
	TimingDateTime_ext *Element `json:"_timingDateTime"`
	// A legal or geographic region in which the activity definition is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Extensions for usage
	Usage_ext *Element `json:"_usage"`
	// A copyright statement relating to the activity definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the activity definition.
	Copyright string `json:"copyright"`
	// The period, timing or frequency upon which the described activity is to occur.
	TimingRange *Range `json:"timingRange"`
	// Identifies the food, drug or other product being consumed or supplied in the activity.
	ProductReference *Reference `json:"productReference"`
	// Provides detailed dosage instructions in the same way that they are described for MedicationRequest resources.
	Dosage []*Dosage `json:"dosage"`
	// Indicates the sites on the subject's body where the procedure should be performed (I.e. the target sites).
	BodySite []*CodeableConcept `json:"bodySite"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// A free text natural language description of the activity definition from a consumer's perspective.
	Description string `json:"description"`
	// Explaination of why this activity definition is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// The identifier that is used to identify this version of the activity definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the activity definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active assets.
	Version string `json:"version"`
	// Descriptive topics related to the content of the activity. Topics provide a high-level categorization of the activity that can be useful for filtering and searching.
	Topic []*CodeableConcept `json:"topic"`
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []*RelatedArtifact `json:"relatedArtifact"`
	// The period, timing or frequency upon which the described activity is to occur.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	TimingDateTime time.Time `json:"timingDateTime"`
	// The period during which the activity definition content was or is planned to be in active use.
	EffectivePeriod *Period `json:"effectivePeriod"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ApprovalDate time.Time `json:"approvalDate"`
	// A contributor to the content of the asset, including authors, editors, reviewers, and endorsers.
	Contributor []*Contributor `json:"contributor"`
	// Identifies the food, drug or other product being consumed or supplied in the activity.
	ProductCodeableConcept *CodeableConcept `json:"productCodeableConcept"`
	// A short, descriptive, user-friendly title for the activity definition.
	Title string `json:"title"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A detailed description of how the asset is used from a clinical perspective.
	Usage string `json:"usage"`
	// Extensions for approvalDate
	ApprovalDate_ext *Element `json:"_approvalDate"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// The date on which the resource content was last reviewed. Review happens periodically after approval, but doesn't change the original approval date.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	LastReviewDate time.Time `json:"lastReviewDate"`
	// Identifies the facility where the activity will occur; e.g. home, hospital, specific clinic, etc.
	Location *Reference `json:"location"`
	// Identifies the quantity expected to be consumed at once (per dose, per meal, etc.).
	Quantity *Quantity `json:"quantity"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// An absolute URI that is used to identify this activity definition when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this activity definition is (or will be) published. The URL SHOULD include the major version of the activity definition. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// Indicates who should participate in performing the action described.
	Participant []*ActivityDefinition_Participant `json:"participant"`
	// This is a ActivityDefinition resource
	ResourceType string `json:"resourceType,omitempty"`
	// A boolean value to indicate that this activity definition is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// A reference to a StructureMap resource that defines a transform that can be executed to produce the intent resource using the ActivityDefinition instance as the input.
	Transform *Reference `json:"transform"`
	// A formal identifier that is used to identify this activity definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for kind
	Kind_ext *Element `json:"_kind"`
	// The name of the individual or organization that published the activity definition.
	Publisher string `json:"publisher"`
}

func NewActivityDefinition() *ActivityDefinition {
	return &ActivityDefinition{ResourceType: "ActivityDefinition"}
}

// ExplanationOfBenefit_ProcessNote is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_ProcessNote struct {
	_ *BackboneElement
	// The note text.
	Text string `json:"text"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// The ISO-639-1 alpha 2 code in lower case for the language, optionally followed by a hyphen and the ISO-3166-1 alpha 2 code for the region in upper case; e.g. "en" for English, or "en-US" for American English versus "en-EN" for England English.
	Language *CodeableConcept `json:"language"`
	// An integer associated with each note which may be referred to from each service line item.
	// pattern [1-9][0-9]*
	Number uint64 `json:"number"`
	// Extensions for number
	Number_ext *Element `json:"_number"`
	// The note purpose: Print/Display.
	Type *CodeableConcept `json:"type"`
}

// MeasureReport_Population is The MeasureReport resource contains the results of evaluating a measure.
type MeasureReport_Population struct {
	_ *BackboneElement
	// The type of the population.
	Code *CodeableConcept `json:"code"`
	// The number of members of the population.
	// pattern -?([0]|([1-9][0-9]*))
	Count int64 `json:"count"`
	// Extensions for count
	Count_ext *Element `json:"_count"`
	// This element refers to a List of patient level MeasureReport resources, one for each patient in this population.
	Patients *Reference `json:"patients"`
	// The identifier of the population being reported, as defined by the population element of the measure.
	Identifier *Identifier `json:"identifier"`
}

// Patient_Communication is Demographics and other administrative information about an individual or animal receiving care or other health-related services.
type Patient_Communication struct {
	_ *BackboneElement
	// The ISO-639-1 alpha 2 code in lower case for the language, optionally followed by a hyphen and the ISO-3166-1 alpha 2 code for the region in upper case; e.g. "en" for English, or "en-US" for American English versus "en-EN" for England English.
	Language *CodeableConcept `json:"language,omitempty"`
	// Indicates whether or not the patient prefers this language (over other languages he masters up a certain level).
	Preferred bool `json:"preferred"`
	// Extensions for preferred
	Preferred_ext *Element `json:"_preferred"`
}

// Meta is The metadata about a resource. This is content in the resource that is maintained by the infrastructure. Changes to the content may not always be associated with version changes to the resource.
type Meta struct {
	_ *Element
	// When the resource last changed - e.g. when the version changed.
	LastUpdated string `json:"lastUpdated"`
	// Extensions for lastUpdated
	LastUpdated_ext *Element `json:"_lastUpdated"`
	// A list of profiles (references to [[[StructureDefinition]]] resources) that this resource claims to conform to. The URL is a reference to [[[StructureDefinition.url]]].
	Profile []string `json:"profile"`
	// Extensions for profile
	Profile_ext []*Element `json:"_profile"`
	// Security labels applied to this resource. These tags connect specific resources to the overall security policy and infrastructure.
	Security []*Coding `json:"security"`
	// Tags applied to this resource. Tags are intended to be used to identify and relate resources to process and workflow, and applications are not required to consider the tags when interpreting the meaning of a resource.
	Tag []*Coding `json:"tag"`
	// The version specific identifier, as it appears in the version portion of the URL. This values changes when the resource is created, updated, or deleted.
	// pattern [A-Za-z0-9\-\.]{1,64}
	VersionId string `json:"versionId"`
	// Extensions for versionId
	VersionId_ext *Element `json:"_versionId"`
}

// AuditEvent_Source is A record of an event made for purposes of maintaining a security log. Typical uses include detection of intrusion attempts and monitoring for inappropriate usage.
type AuditEvent_Source struct {
	_ *BackboneElement
	// Logical source location within the healthcare enterprise network.  For example, a hospital or other provider location within a multi-entity provider group.
	Site string `json:"site"`
	// Extensions for site
	Site_ext *Element `json:"_site"`
	// Identifier of the source where the event was detected.
	Identifier *Identifier `json:"identifier,omitempty"`
	// Code specifying the type of source where event originated.
	Type []*Coding `json:"type"`
}

// Practitioner is A person who is directly or indirectly involved in the provisioning of healthcare.
type Practitioner struct {
	_ *DomainResource
	// An identifier that applies to this person in this role.
	Identifier []*Identifier `json:"identifier"`
	// Whether this practitioner's record is in active use.
	Active bool `json:"active"`
	// A language the practitioner is able to use in patient communication.
	Communication []*CodeableConcept `json:"communication"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
	// The name(s) associated with the practitioner.
	Name []*HumanName `json:"name"`
	// A contact detail for the practitioner, e.g. a telephone number or an email address.
	Telecom []*ContactPoint `json:"telecom"`
	// Address(es) of the practitioner that are not role specific (typically home address).
	// Work addresses are not typically entered in this property as they are usually role dependent.
	Address []*Address `json:"address"`
	// The date of birth for the practitioner.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	BirthDate time.Time `json:"birthDate"`
	// Image of the person.
	Photo []*Attachment `json:"photo"`
	// This is a Practitioner resource
	ResourceType string `json:"resourceType,omitempty"`
	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes.
	Gender string `json:"gender"`
	// Extensions for gender
	Gender_ext *Element `json:"_gender"`
	// Extensions for birthDate
	BirthDate_ext *Element `json:"_birthDate"`
	// Qualifications obtained by training and certification.
	Qualification []*Practitioner_Qualification `json:"qualification"`
}

func NewPractitioner() *Practitioner { return &Practitioner{ResourceType: "Practitioner"} }

// Range is A set of ordered Quantities defined by a low and high limit.
type Range struct {
	_ *Element
	// The low limit. The boundary is inclusive.
	Low *Quantity `json:"low"`
	// The high limit. The boundary is inclusive.
	High *Quantity `json:"high"`
}

// OperationDefinition is A formal computable definition of an operation (on the RESTful interface) or a named query (using the search interaction).
type OperationDefinition struct {
	_ *DomainResource
	// An absolute URI that is used to identify this operation definition when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this operation definition is (or will be) published. The URL SHOULD include the major version of the operation definition. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// Whether this is an operation or a named query.
	Kind string `json:"kind"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The identifier that is used to identify this version of the operation definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the operation definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Indicates that this operation definition is a constraining profile on the base.
	Base *Reference `json:"base"`
	// Indicates whether this operation or named query can be invoked at the system level (e.g. without needing to choose a resource type for the context).
	System bool `json:"system"`
	// This is a OperationDefinition resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// Additional information about how to use this operation or named query.
	Comment string `json:"comment"`
	// Indicates whether this operation can be invoked on a particular instance of one of the given types.
	Instance bool `json:"instance"`
	// Extensions for instance
	Instance_ext *Element `json:"_instance"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// The types on which this operation can be executed.
	Resource []string `json:"resource"`
	// Extensions for resource
	Resource_ext []*Element `json:"_resource"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A legal or geographic region in which the operation definition is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// The name used to invoke the operation.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The status of this operation definition. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for kind
	Kind_ext *Element `json:"_kind"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Extensions for system
	System_ext *Element `json:"_system"`
	// Operations that are idempotent (see [HTTP specification definition of idempotent](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html)) may be invoked by performing an HTTP GET operation instead of a POST.
	Idempotent bool `json:"idempotent"`
	// Extensions for idempotent
	Idempotent_ext *Element `json:"_idempotent"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// Defines an appropriate combination of parameters to use when invoking this operation, to help code generators when generating overloaded parameter sets for this operation.
	Overload []*OperationDefinition_Overload `json:"overload"`
	// A natural language name identifying the operation definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// The name of the individual or organization that published the operation definition.
	Publisher string `json:"publisher"`
	// A free text natural language description of the operation definition from a consumer's perspective.
	Description string `json:"description"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// Indicates whether this operation or named query can be invoked at the resource type level for any given resource type level (e.g. without needing to choose a specific resource id for the context).
	Type bool `json:"type"`
	// The parameters for the operation/query.
	Parameter []*OperationDefinition_Parameter `json:"parameter"`
	// A boolean value to indicate that this operation definition is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// The date  (and optionally time) when the operation definition was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the operation definition changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate operation definition instances.
	UseContext []*UsageContext `json:"useContext"`
	// Explaination of why this operation definition is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
}

func NewOperationDefinition() *OperationDefinition {
	return &OperationDefinition{ResourceType: "OperationDefinition"}
}

// Parameters is This special resource type is used to represent an operation request and response (operations.html). It has no other use, and there is no RESTful endpoint associated with it.
type Parameters struct {
	_ *Resource
	// A parameter passed to or received from the operation.
	Parameter []*Parameters_Parameter `json:"parameter"`
}

// Questionnaire_Option is A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
type Questionnaire_Option struct {
	_ *BackboneElement
	// Extensions for valueInteger
	ValueInteger_ext *Element `json:"_valueInteger"`
	// Extensions for valueTime
	ValueTime_ext *Element `json:"_valueTime"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// A potential answer that's allowed as the answer to this question.
	ValueCoding *Coding `json:"valueCoding"`
	// A potential answer that's allowed as the answer to this question.
	// pattern -?([0]|([1-9][0-9]*))
	ValueInteger int64 `json:"valueInteger"`
	// A potential answer that's allowed as the answer to this question.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ValueDate time.Time `json:"valueDate"`
	// Extensions for valueDate
	ValueDate_ext *Element `json:"_valueDate"`
	// A potential answer that's allowed as the answer to this question.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	ValueTime time.Time `json:"valueTime"`
	// A potential answer that's allowed as the answer to this question.
	ValueString string `json:"valueString"`
}

// TestScript_Destination is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Destination struct {
	_ *BackboneElement
	// Abstract name given to a destination server in this test script.  The name is provided as a number starting at 1.
	// pattern -?([0]|([1-9][0-9]*))
	Index int64 `json:"index"`
	// Extensions for index
	Index_ext *Element `json:"_index"`
	// The type of destination profile the test system supports.
	Profile *Coding `json:"profile,omitempty"`
}

// ValueSet_Expansion is A value set specifies a set of codes drawn from one or more code systems.
type ValueSet_Expansion struct {
	_ *BackboneElement
	// If paging is being used, the offset at which this resource starts.  I.e. this resource is a partial view into the expansion. If paging is not being used, this element SHALL not be present.
	// pattern -?([0]|([1-9][0-9]*))
	Offset int64 `json:"offset"`
	// A parameter that controlled the expansion process. These parameters may be used by users of expanded value sets to check whether the expansion is suitable for a particular purpose, or to pick the correct expansion.
	Parameter []*ValueSet_Parameter `json:"parameter"`
	// The codes that are contained in the value set expansion.
	Contains []*ValueSet_Contains `json:"contains"`
	// An identifier that uniquely identifies this expansion of the valueset. Systems may re-use the same identifier as long as the expansion and the definition remain the same, but are not required to do so.
	Identifier string `json:"identifier"`
	// Extensions for identifier
	Identifier_ext *Element `json:"_identifier"`
	// The time at which the expansion was produced by the expanding system.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Timestamp time.Time `json:"timestamp"`
	// Extensions for timestamp
	Timestamp_ext *Element `json:"_timestamp"`
	// The total number of concepts in the expansion. If the number of concept nodes in this resource is less than the stated number, then the server can return more using the offset parameter.
	// pattern -?([0]|([1-9][0-9]*))
	Total int64 `json:"total"`
	// Extensions for total
	Total_ext *Element `json:"_total"`
	// Extensions for offset
	Offset_ext *Element `json:"_offset"`
}

// Contract is A formal agreement between parties regarding the conduct of business, exchange of information or other matters.
type Contract struct {
	_ *DomainResource
	// The "patient friendly language" versionof the Contract in whole or in parts. "Patient friendly language" means the representation of the Contract and Contract Provisions in a manner that is readily accessible and understandable by a layperson in accordance with best practices for communication styles that ensure that those agreeing to or signing the Contract understand the roles, actions, obligations, responsibilities, and implication of the agreement.
	Friendly []*Contract_Friendly `json:"friendly"`
	// The target entity impacted by or of interest to parties to the agreement.
	Subject []*Reference `json:"subject"`
	// Action stipulated by this Contract.
	Action []*CodeableConcept `json:"action"`
	// The type of decision made by a grantor with respect to an offer made by a grantee.
	DecisionType *CodeableConcept `json:"decisionType"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Extensions for issued
	Issued_ext *Element `json:"_issued"`
	// A set of security labels that define which resources are controlled by this consent. If more than one label is specified, all resources must have all the specified labels.
	SecurityLabel []*Coding `json:"securityLabel"`
	// Legally binding Contract: This is the signed and legally recognized representation of the Contract, which is considered the "source of truth" and which would be the basis for legal action related to enforcement of this Contract.
	BindingAttachment *Attachment `json:"bindingAttachment"`
	// Unique identifier for this Contract.
	Identifier *Identifier `json:"identifier"`
	// When this  Contract was issued.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Issued time.Time `json:"issued"`
	// The matter of concern in the context of this agreement.
	Topic []*Reference `json:"topic"`
	// Type of Contract such as an insurance policy, real estate contract, a will, power of attorny, Privacy or Security policy , trust framework agreement, etc.
	Type *CodeableConcept `json:"type"`
	// Relevant time or time-period when this Contract is applicable.
	Applies *Period `json:"applies"`
	// Reason for action stipulated by this Contract.
	ActionReason []*CodeableConcept `json:"actionReason"`
	// One or more Contract Provisions, which may be related and conveyed as a group, and may contain nested groups.
	Term []*Contract_Term `json:"term"`
	// List of Legal expressions or representations of this Contract.
	Legal []*Contract_Legal `json:"legal"`
	// List of Computable Policy Rule Language Representations of this Contract.
	Rule []*Contract_Rule `json:"rule"`
	// Legally binding Contract: This is the signed and legally recognized representation of the Contract, which is considered the "source of truth" and which would be the basis for legal action related to enforcement of this Contract.
	BindingReference *Reference `json:"bindingReference"`
	// This is a Contract resource
	ResourceType string `json:"resourceType,omitempty"`
	// Recognized governance framework or system operating with a circumscribed scope in accordance with specified principles, policies, processes or procedures for managing rights, actions, or behaviors of parties or principals relative to resources.
	Domain []*Reference `json:"domain"`
	// An actor taking a role in an activity for which it can be assigned some degree of responsibility for the activity taking place.
	Agent []*Contract_Agent `json:"agent"`
	// Parties with legal standing in the Contract, including the principal parties, the grantor(s) and grantee(s), which are any person or organization bound by the contract, and any ancillary parties, which facilitate the execution of the contract such as a notary or witness.
	Signer []*Contract_Signer `json:"signer"`
	// Contract Valued Item List.
	ValuedItem []*Contract_ValuedItem `json:"valuedItem"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A formally or informally recognized grouping of people, principals, organizations, or jurisdictions formed for the purpose of achieving some form of collective action such as the promulgation, administration and enforcement of contracts and policies.
	Authority []*Reference `json:"authority"`
	// More specific type or specialization of an overarching or more general contract such as auto insurance, home owner  insurance, prenupial agreement, Advanced-Directive, or privacy consent.
	SubType []*CodeableConcept `json:"subType"`
	// The minimal content derived from the basal information source at a specific stage in its lifecycle.
	ContentDerivative *CodeableConcept `json:"contentDerivative"`
}

func NewContract() *Contract { return &Contract{ResourceType: "Contract"} }

// ElementDefinition_Example is Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition_Example struct {
	_ *BackboneElement
	// Extensions for valueInteger
	ValueInteger_ext *Element `json:"_valueInteger"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	ValueDecimal float64 `json:"valueDecimal"`
	// Extensions for valueUnsignedInt
	ValueUnsignedInt_ext *Element `json:"_valueUnsignedInt"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueContactPoint *ContactPoint `json:"valueContactPoint"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueBoolean bool `json:"valueBoolean"`
	// Extensions for valueUuid
	ValueUuid_ext *Element `json:"_valueUuid"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueCoding *Coding `json:"valueCoding"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueSampledData *SampledData `json:"valueSampledData"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueContactDetail *ContactDetail `json:"valueContactDetail"`
	// Extensions for valueDecimal
	ValueDecimal_ext *Element `json:"_valueDecimal"`
	// Extensions for valueOid
	ValueOid_ext *Element `json:"_valueOid"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueMarkdown string `json:"valueMarkdown"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueAddress *Address `json:"valueAddress"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern -?([0]|([1-9][0-9]*))
	ValueInteger int64 `json:"valueInteger"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueString string `json:"valueString"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern [0]|([1-9][0-9]*)
	ValueUnsignedInt uint64 `json:"valueUnsignedInt"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueDosage *Dosage `json:"valueDosage"`
	// Extensions for valueCode
	ValueCode_ext *Element `json:"_valueCode"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern [1-9][0-9]*
	ValuePositiveInt uint64 `json:"valuePositiveInt"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueBackboneElement *BackboneElement `json:"valueBackboneElement"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueRange *Range `json:"valueRange"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueRatio *Ratio `json:"valueRatio"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueDuration *Duration `json:"valueDuration"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueTiming *Timing `json:"valueTiming"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueUsageContext *UsageContext `json:"valueUsageContext"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	ValueTime time.Time `json:"valueTime"`
	// Extensions for valueMarkdown
	ValueMarkdown_ext *Element `json:"_valueMarkdown"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueAnnotation *Annotation `json:"valueAnnotation"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueIdentifier *Identifier `json:"valueIdentifier"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition"`
	// Describes the purpose of this example amoung the set of examples.
	Label string `json:"label"`
	// Extensions for label
	Label_ext *Element `json:"_label"`
	// Extensions for valueBase64Binary
	ValueBase64Binary_ext *Element `json:"_valueBase64Binary"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	ValueUuid string `json:"valueUuid"`
	// Extensions for valuePositiveInt
	ValuePositiveInt_ext *Element `json:"_valuePositiveInt"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueElement *Element `json:"valueElement"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueDistance *Distance `json:"valueDistance"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueCount *Count `json:"valueCount"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueUri string `json:"valueUri"`
	// Extensions for valueTime
	ValueTime_ext *Element `json:"_valueTime"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueNarrative *Narrative `json:"valueNarrative"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueSignature *Signature `json:"valueSignature"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// Extensions for valueDate
	ValueDate_ext *Element `json:"_valueDate"`
	// Extensions for valueDateTime
	ValueDateTime_ext *Element `json:"_valueDateTime"`
	// Extensions for valueId
	ValueId_ext *Element `json:"_valueId"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueExtension *Extension `json:"valueExtension"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueMoney *Money `json:"valueMoney"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ValueDateTime time.Time `json:"valueDateTime"`
	// Extensions for valueInstant
	ValueInstant_ext *Element `json:"_valueInstant"`
	// Extensions for valueUri
	ValueUri_ext *Element `json:"_valueUri"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueElementDefinition *ElementDefinition `json:"valueElementDefinition"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueBase64Binary string `json:"valueBase64Binary"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern [^\s]+([\s]?[^\s]+)*
	ValueCode string `json:"valueCode"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	ValueOid string `json:"valueOid"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueSimpleQuantity *Quantity `json:"valueSimpleQuantity"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueAge *Age `json:"valueAge"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValuePeriod *Period `json:"valuePeriod"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueReference *Reference `json:"valueReference"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueContributor *Contributor `json:"valueContributor"`
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueInstant string `json:"valueInstant"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ValueDate time.Time `json:"valueDate"`
	// The actual value for the element, which must be one of the types allowed for this element.
	// pattern [A-Za-z0-9\-\.]{1,64}
	ValueId string `json:"valueId"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueAttachment *Attachment `json:"valueAttachment"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueHumanName *HumanName `json:"valueHumanName"`
	// The actual value for the element, which must be one of the types allowed for this element.
	ValueMeta *Meta `json:"valueMeta"`
}

// Bundle is A container for a collection of resources.
type Bundle struct {
	_ *Resource
	// Indicates the purpose of this bundle - how it was intended to be used.
	Type string `json:"type"`
	// Extensions for total
	Total_ext *Element `json:"_total"`
	// This is a Bundle resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// If a set of search matches, this is the total number of matches for the search (as opposed to the number of results in this bundle).
	// pattern [0]|([1-9][0-9]*)
	Total uint64 `json:"total"`
	// A series of links that provide context to this bundle.
	Link []*Bundle_Link `json:"link"`
	// An entry in a bundle resource - will either contain a resource, or information about a resource (transactions and history only).
	Entry []*Bundle_Entry `json:"entry"`
	// Digital Signature - base64 encoded. XML-DSIg or a JWT.
	Signature *Signature `json:"signature"`
	// A persistent identifier for the batch that won't change as a batch is copied from server to server.
	Identifier *Identifier `json:"identifier"`
}

func NewBundle() *Bundle { return &Bundle{ResourceType: "Bundle"} }

// HealthcareService_NotAvailable is The details of a healthcare service available at a location.
type HealthcareService_NotAvailable struct {
	_ *BackboneElement
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Service is not available (seasonally or for a public holiday) from this date.
	During *Period `json:"during"`
	// The reason that can be presented to the user as to why this time is not available.
	Description string `json:"description"`
}

// MessageDefinition_AllowedResponse is Defines the characteristics of a message that can be shared between systems, including the type of event that initiates the message, the content to be transmitted and what response(s), if any, are permitted.
type MessageDefinition_AllowedResponse struct {
	_ *BackboneElement
	// Extensions for situation
	Situation_ext *Element `json:"_situation"`
	// A reference to the message definition that must be adhered to by this supported response.
	Message *Reference `json:"message,omitempty"`
	// Provides a description of the circumstances in which this response should be used (as opposed to one of the alternative responses).
	Situation string `json:"situation"`
}

// Provenance_Agent is Provenance of a resource is a record that describes entities and processes involved in producing and delivering or otherwise influencing that resource. Provenance provides a critical foundation for assessing authenticity, enabling trust, and allowing reproducibility. Provenance assertions are a form of contextual metadata and can themselves become important records with their own provenance. Provenance statement indicates clinical significance in terms of confidence in authenticity, reliability, and trustworthiness, integrity, and stage in lifecycle (e.g. Document Completion - has the artifact been legally authenticated), all of which may impact security, privacy, and trust policies.
type Provenance_Agent struct {
	_ *BackboneElement
	// Extensions for whoUri
	WhoUri_ext *Element `json:"_whoUri"`
	// The individual, device or organization that participated in the event.
	WhoReference *Reference `json:"whoReference"`
	// The individual, device, or organization for whom the change was made.
	OnBehalfOfUri string `json:"onBehalfOfUri"`
	// Extensions for onBehalfOfUri
	OnBehalfOfUri_ext *Element `json:"_onBehalfOfUri"`
	// The individual, device, or organization for whom the change was made.
	OnBehalfOfReference *Reference `json:"onBehalfOfReference"`
	// The type of relationship between agents.
	RelatedAgentType *CodeableConcept `json:"relatedAgentType"`
	// The function of the agent with respect to the activity. The security role enabling the agent with respect to the activity.
	Role []*CodeableConcept `json:"role"`
	// The individual, device or organization that participated in the event.
	WhoUri string `json:"whoUri"`
}

// RelatedPerson is Information about a person that is involved in the care for a patient, but who is not the target of healthcare, nor has a formal responsibility in the care process.
type RelatedPerson struct {
	_ *DomainResource
	// Extensions for gender
	Gender_ext *Element `json:"_gender"`
	// Extensions for birthDate
	BirthDate_ext *Element `json:"_birthDate"`
	// The period of time that this relationship is considered to be valid. If there are no dates defined, then the interval is unknown.
	Period *Period `json:"period"`
	// Whether this related person record is in active use.
	Active bool `json:"active"`
	// The nature of the relationship between a patient and the related person.
	Relationship *CodeableConcept `json:"relationship"`
	// A name associated with the person.
	Name []*HumanName `json:"name"`
	// The patient this person is related to.
	Patient *Reference `json:"patient,omitempty"`
	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes.
	Gender string `json:"gender"`
	// Image of the person.
	Photo []*Attachment `json:"photo"`
	// A contact detail for the person, e.g. a telephone number or an email address.
	Telecom []*ContactPoint `json:"telecom"`
	// This is a RelatedPerson resource
	ResourceType string `json:"resourceType,omitempty"`
	// Identifier for a person within a particular scope.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
	// The date on which the related person was born.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	BirthDate time.Time `json:"birthDate"`
	// Address where the related person can be contacted or visited.
	Address []*Address `json:"address"`
}

func NewRelatedPerson() *RelatedPerson { return &RelatedPerson{ResourceType: "RelatedPerson"} }

// TestScript_Rule2 is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Rule2 struct {
	_ *BackboneElement
	// The TestScript.rule id value this assert will evaluate.
	// pattern [A-Za-z0-9\-\.]{1,64}
	RuleId string `json:"ruleId"`
	// Extensions for ruleId
	RuleId_ext *Element `json:"_ruleId"`
	// Each rule template can take one or more parameters for rule evaluation.
	Param []*TestScript_Param2 `json:"param"`
}

// Composition_Event is A set of healthcare-related information that is assembled together into a single logical document that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. While a Composition defines the structure, it does not actually contain the content: rather the full content of a document is contained in a Bundle, of which the Composition is the first resource contained.
type Composition_Event struct {
	_ *BackboneElement
	// The period of time covered by the documentation. There is no assertion that the documentation is a complete representation for this period, only that it documents events during this time.
	Period *Period `json:"period"`
	// The description and/or reference of the event(s) being documented. For example, this could be used to document such a colonoscopy or an appendectomy.
	Detail []*Reference `json:"detail"`
	// This list of codes represents the main clinical acts, such as a colonoscopy or an appendectomy, being documented. In some cases, the event is inherent in the typeCode, such as a "History and Physical Report" in which the procedure being documented is necessarily a "History and Physical" act.
	Code []*CodeableConcept `json:"code"`
}

// Measure_SupplementalData is The Measure resource provides the definition of a quality measure.
type Measure_SupplementalData struct {
	_ *BackboneElement
	// Extensions for criteria
	Criteria_ext *Element `json:"_criteria"`
	// The supplemental data to be supplied as part of the measure response, specified as a valid FHIR Resource Path.
	Path string `json:"path"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// An identifier for the supplemental data.
	Identifier *Identifier `json:"identifier"`
	// An indicator of the intended usage for the supplemental data element. Supplemental data indicates the data is additional information requested to augment the measure information. Risk adjustment factor indicates the data is additional information used to calculate risk adjustment factors when applying a risk model to the measure calculation.
	Usage []*CodeableConcept `json:"usage"`
	// The criteria for the supplemental data. This must be the name of a valid expression defined within a referenced library, and defines the data to be returned for this element.
	Criteria string `json:"criteria"`
}

// MedicationDispense_Substitution is Indicates that a medication product is to be or has been dispensed for a named person/patient.  This includes a description of the medication product (supply) provided and the instructions for administering the medication.  The medication dispense is the result of a pharmacy system responding to a medication order.
type MedicationDispense_Substitution struct {
	_ *BackboneElement
	// The person or organization that has primary responsibility for the substitution.
	ResponsibleParty []*Reference `json:"responsibleParty"`
	// True if the dispenser dispensed a different drug or product from what was prescribed.
	WasSubstituted bool `json:"wasSubstituted"`
	// Extensions for wasSubstituted
	WasSubstituted_ext *Element `json:"_wasSubstituted"`
	// A code signifying whether a different drug was dispensed from what was prescribed.
	Type *CodeableConcept `json:"type"`
	// Indicates the reason for the substitution of (or lack of substitution) from what was prescribed.
	Reason []*CodeableConcept `json:"reason"`
}

// OperationOutcome is A collection of error, warning or information messages that result from a system action.
type OperationOutcome struct {
	_ *DomainResource
	// This is a OperationOutcome resource
	ResourceType string `json:"resourceType,omitempty"`
	// An error, warning or information message that results from a system action.
	Issue []*OperationOutcome_Issue `json:"issue,omitempty"`
}

func NewOperationOutcome() *OperationOutcome {
	return &OperationOutcome{ResourceType: "OperationOutcome"}
}

// Practitioner_Qualification is A person who is directly or indirectly involved in the provisioning of healthcare.
type Practitioner_Qualification struct {
	_ *BackboneElement
	// Period during which the qualification is valid.
	Period *Period `json:"period"`
	// Organization that regulates and issues the qualification.
	Issuer *Reference `json:"issuer"`
	// An identifier that applies to this person's qualification in this role.
	Identifier []*Identifier `json:"identifier"`
	// Coded representation of the qualification.
	Code *CodeableConcept `json:"code,omitempty"`
}

// PractitionerRole_AvailableTime is A specific set of Roles/Locations/specialties/services that a practitioner may perform at an organization for a period of time.
type PractitionerRole_AvailableTime struct {
	_ *BackboneElement
	// Is this always available? (hence times are irrelevant) e.g. 24 hour service.
	AllDay bool `json:"allDay"`
	// Extensions for allDay
	AllDay_ext *Element `json:"_allDay"`
	// The opening time of day. Note: If the AllDay flag is set, then this time is ignored.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	AvailableStartTime time.Time `json:"availableStartTime"`
	// Extensions for availableStartTime
	AvailableStartTime_ext *Element `json:"_availableStartTime"`
	// The closing time of day. Note: If the AllDay flag is set, then this time is ignored.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	AvailableEndTime time.Time `json:"availableEndTime"`
	// Extensions for availableEndTime
	AvailableEndTime_ext *Element `json:"_availableEndTime"`
	// Indicates which days of the week are available between the start and end Times.
	DaysOfWeek []string `json:"daysOfWeek"`
	// Extensions for daysOfWeek
	DaysOfWeek_ext []*Element `json:"_daysOfWeek"`
}

// Consent_Except is A record of a healthcare consumer's policy choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent_Except struct {
	_ *BackboneElement
	// The class of information covered by this exception. The type can be a FHIR resource type, a profile on a type, or a CDA document, or some other type that indicates what sort of information the consent relates to.
	Class []*Coding `json:"class"`
	// If this code is found in an instance, then the exception applies.
	Code []*Coding `json:"code"`
	// Clinical or Operational Relevant period of time that bounds the data controlled by this exception.
	DataPeriod *Period `json:"dataPeriod"`
	// Who or what is controlled by this Exception. Use group to identify a set of actors by some property they share (e.g. 'admitting officers').
	Actor []*Consent_Actor1 `json:"actor"`
	// The context of the activities a user is taking - why the user is accessing the data - that are controlled by this exception.
	Purpose []*Coding `json:"purpose"`
	// The timeframe in this exception is valid.
	Period *Period `json:"period"`
	// Actions controlled by this Exception.
	Action []*CodeableConcept `json:"action"`
	// A set of security labels that define which resources are controlled by this exception. If more than one label is specified, all resources must have all the specified labels.
	SecurityLabel []*Coding `json:"securityLabel"`
	// The resources controlled by this exception, if specific resources are referenced.
	Data []*Consent_Data1 `json:"data"`
	// Action  to take - permit or deny - when the exception conditions are met.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
}

// ConceptMap_Unmapped is A statement of relationships from one set of concepts to one or more other concepts - either code systems or data elements, or classes in class models.
type ConceptMap_Unmapped struct {
	_ *BackboneElement
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// The fixed code to use when the mode = 'fixed'  - all unmapped codes are mapped to a single fixed code.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The display for the code. The display is only provided to help editors when editing the concept map.
	Display string `json:"display"`
	// Extensions for display
	Display_ext *Element `json:"_display"`
	// The canonical URL of the map to use if this map contains no mapping.
	Url string `json:"url"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Defines which action to take if there is no match in the group. One of 3 actions is possible: use the unmapped code (this is useful when doing a mapping between versions, and only a few codes have changed), use a fixed code (a default code), or alternatively, a reference to a different concept map can be provided (by canonical URL).
	Mode string `json:"mode"`
}

// Group_Member is Represents a defined collection of entities that may be discussed or acted upon collectively but which are not expected to act collectively and are not formally or legally recognized; i.e. a collection of entities that isn't an Organization.
type Group_Member struct {
	_ *BackboneElement
	// A reference to the entity that is a member of the group. Must be consistent with Group.type.
	Entity *Reference `json:"entity,omitempty"`
	// The period that the member was in the group, if known.
	Period *Period `json:"period"`
	// A flag to indicate that the member is no longer in the group, but previously may have been a member.
	Inactive bool `json:"inactive"`
	// Extensions for inactive
	Inactive_ext *Element `json:"_inactive"`
}

// Medication_Batch is This resource is primarily used for the identification and definition of a medication. It covers the ingredients and the packaging for a medication.
type Medication_Batch struct {
	_ *BackboneElement
	// When this specific batch of product will expire.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ExpirationDate time.Time `json:"expirationDate"`
	// Extensions for expirationDate
	ExpirationDate_ext *Element `json:"_expirationDate"`
	// The assigned lot number of a batch of the specified product.
	LotNumber string `json:"lotNumber"`
	// Extensions for lotNumber
	LotNumber_ext *Element `json:"_lotNumber"`
}

// Observation_ReferenceRange is Measurements and simple assertions made about a patient, device or other subject.
type Observation_ReferenceRange struct {
	_ *BackboneElement
	// Text based reference range in an observation which may be used when a quantitative range is not appropriate for an observation.  An example would be a reference value of "Negative" or a list or table of 'normals'.
	Text string `json:"text"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// The value of the low bound of the reference range.  The low bound of the reference range endpoint is inclusive of the value (e.g.  reference range is >=5 - <=9).   If the low bound is omitted,  it is assumed to be meaningless (e.g. reference range is <=2.3).
	Low *Quantity `json:"low"`
	// The value of the high bound of the reference range.  The high bound of the reference range endpoint is inclusive of the value (e.g.  reference range is >=5 - <=9).   If the high bound is omitted,  it is assumed to be meaningless (e.g. reference range is >= 2.3).
	High *Quantity `json:"high"`
	// Codes to indicate the what part of the targeted reference population it applies to. For example, the normal or therapeutic range.
	Type *CodeableConcept `json:"type"`
	// Codes to indicate the target population this reference range applies to.  For example, a reference range may be based on the normal population or a particular sex or race.
	AppliesTo []*CodeableConcept `json:"appliesTo"`
	// The age at which this reference range is applicable. This is a neonatal age (e.g. number of weeks at term) if the meaning says so.
	Age *Range `json:"age"`
}

// Age is A duration of time during which an organism (or a process) has existed.
type Age struct {
	_ *Quantity
}

// HealthcareService_AvailableTime is The details of a healthcare service available at a location.
type HealthcareService_AvailableTime struct {
	_ *BackboneElement
	// Extensions for allDay
	AllDay_ext *Element `json:"_allDay"`
	// The opening time of day. Note: If the AllDay flag is set, then this time is ignored.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	AvailableStartTime time.Time `json:"availableStartTime"`
	// Extensions for availableStartTime
	AvailableStartTime_ext *Element `json:"_availableStartTime"`
	// The closing time of day. Note: If the AllDay flag is set, then this time is ignored.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	AvailableEndTime time.Time `json:"availableEndTime"`
	// Extensions for availableEndTime
	AvailableEndTime_ext *Element `json:"_availableEndTime"`
	// Indicates which days of the week are available between the start and end Times.
	DaysOfWeek []string `json:"daysOfWeek"`
	// Extensions for daysOfWeek
	DaysOfWeek_ext []*Element `json:"_daysOfWeek"`
	// Is this always available? (hence times are irrelevant) e.g. 24 hour service.
	AllDay bool `json:"allDay"`
}

// Medication_Package is This resource is primarily used for the identification and definition of a medication. It covers the ingredients and the packaging for a medication.
type Medication_Package struct {
	_ *BackboneElement
	// A set of components that go to make up the described item.
	Content []*Medication_Content `json:"content"`
	// Information about a group of medication produced or packaged from one production run.
	Batch []*Medication_Batch `json:"batch"`
	// The kind of container that this package comes as.
	Container *CodeableConcept `json:"container"`
}

// MessageHeader_Source is The header for a message exchange that is either requesting or responding to an action.  The reference(s) that are the subject of the action as well as other information related to the action are typically transmitted in a bundle in which the MessageHeader resource instance is the first resource in the bundle.
type MessageHeader_Source struct {
	_ *BackboneElement
	// Can convey versions of multiple systems in situations where a message passes through multiple hands.
	Version string `json:"version"`
	// Extensions for endpoint
	Endpoint_ext *Element `json:"_endpoint"`
	// May include configuration or other information useful in debugging.
	Software string `json:"software"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for software
	Software_ext *Element `json:"_software"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// An e-mail, phone, website or other contact point to use to resolve issues with message communications.
	Contact *ContactPoint `json:"contact"`
	// Identifies the routing target to send acknowledgements to.
	Endpoint string `json:"endpoint"`
	// Human-readable name for the source system.
	Name string `json:"name"`
}

// CapabilityStatement_Security is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Security struct {
	_ *BackboneElement
	// Extensions for cors
	Cors_ext *Element `json:"_cors"`
	// Types of security services that are supported/required by the system.
	Service []*CodeableConcept `json:"service"`
	// General description of how security works.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Certificates associated with security profiles.
	Certificate []*CapabilityStatement_Certificate `json:"certificate"`
	// Server adds CORS headers when responding to requests - this enables javascript applications to use the server.
	Cors bool `json:"cors"`
}

// Device is This resource identifies an instance or a type of a manufactured item that is used in the provision of healthcare without being substantially changed through that activity. The device may be a medical or non-medical device.  Medical devices include durable (reusable) medical equipment, implantable devices, as well as disposable equipment used for diagnostic, treatment, and research for healthcare and public health.  Non-medical devices may include items such as a machine, cellphone, computer, application, etc.
type Device struct {
	_ *DomainResource
	// A network address on which the device may be contacted directly.
	Url string `json:"url"`
	// Unique instance identifiers assigned to a device by manufacturers other organizations or owners.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Code or identifier to identify a kind of device.
	Type *CodeableConcept `json:"type"`
	// Lot number assigned by the manufacturer.
	LotNumber string `json:"lotNumber"`
	// A name of the manufacturer.
	Manufacturer string `json:"manufacturer"`
	// Extensions for manufacturer
	Manufacturer_ext *Element `json:"_manufacturer"`
	// Extensions for model
	Model_ext *Element `json:"_model"`
	// The version of the device, if the device has multiple releases under the same model, or if the device is software or carries firmware.
	Version string `json:"version"`
	// This is a Device resource
	ResourceType string `json:"resourceType,omitempty"`
	// Patient information, If the device is affixed to a person.
	Patient *Reference `json:"patient"`
	// Extensions for lotNumber
	LotNumber_ext *Element `json:"_lotNumber"`
	// The date and time when the device was manufactured.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ManufactureDate time.Time `json:"manufactureDate"`
	// Extensions for manufactureDate
	ManufactureDate_ext *Element `json:"_manufactureDate"`
	// Extensions for expirationDate
	ExpirationDate_ext *Element `json:"_expirationDate"`
	// The place where the device can be found.
	Location *Reference `json:"location"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Descriptive information, usage information or implantation information that is not captured in an existing element.
	Note []*Annotation `json:"note"`
	// [Unique device identifier (UDI)](device.html#5.11.3.2.2) assigned to device label or package.
	Udi *Device_Udi `json:"udi"`
	// The date and time beyond which this device is no longer valid or should not be used (if applicable).
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ExpirationDate time.Time `json:"expirationDate"`
	// The "model" is an identifier assigned by the manufacturer to identify the product by its type. This number is shared by the all devices sold as the same type.
	Model string `json:"model"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// An organization that is responsible for the provision and ongoing maintenance of the device.
	Owner *Reference `json:"owner"`
	// Contact details for an organization or a particular human that is responsible for the device.
	Contact []*ContactPoint `json:"contact"`
	// Provides additional safety characteristics about a medical device.  For example devices containing latex.
	Safety []*CodeableConcept `json:"safety"`
	// Status of the Device availability.
	Status string `json:"status"`
}

func NewDevice() *Device { return &Device{ResourceType: "Device"} }

// ImplementationGuide_Resource is A set of rules of how FHIR is used to solve a particular problem. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuide_Resource struct {
	_ *BackboneElement
	// Whether a resource is included in the guide as part of the rules defined by the guide, or just as an example of a resource that conforms to the rules and/or help implementers understand the intent of the guide.
	Example bool `json:"example"`
	// Extensions for example
	Example_ext *Element `json:"_example"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for acronym
	Acronym_ext *Element `json:"_acronym"`
	// Where this resource is found.
	SourceUri string `json:"sourceUri"`
	// Where this resource is found.
	SourceReference *Reference `json:"sourceReference"`
	// A human assigned name for the resource. All resources SHOULD have a name, but the name may be extracted from the resource (e.g. ValueSet.name).
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// A description of the reason that a resource has been included in the implementation guide.
	Description string `json:"description"`
	// A short code that may be used to identify the resource throughout the implementation guide.
	Acronym string `json:"acronym"`
	// Extensions for sourceUri
	SourceUri_ext *Element `json:"_sourceUri"`
	// Another resource that this resource is an example for. This is mostly used for resources that are included as examples of StructureDefinitions.
	ExampleFor *Reference `json:"exampleFor"`
}

// Patient_Animal is Demographics and other administrative information about an individual or animal receiving care or other health-related services.
type Patient_Animal struct {
	_ *BackboneElement
	// Identifies the high level taxonomic categorization of the kind of animal.
	Species *CodeableConcept `json:"species,omitempty"`
	// Identifies the detailed categorization of the kind of animal.
	Breed *CodeableConcept `json:"breed"`
	// Indicates the current state of the animal's reproductive organs.
	GenderStatus *CodeableConcept `json:"genderStatus"`
}

// Substance_Ingredient is A homogeneous material with a definite composition.
type Substance_Ingredient struct {
	_ *BackboneElement
	// The amount of the ingredient in the substance - a concentration ratio.
	Quantity *Ratio `json:"quantity"`
	// Another substance that is a component of this substance.
	SubstanceCodeableConcept *CodeableConcept `json:"substanceCodeableConcept"`
	// Another substance that is a component of this substance.
	SubstanceReference *Reference `json:"substanceReference"`
}

// DeviceComponent is The characteristics, operational status and capabilities of a medical-related component of a medical device.
type DeviceComponent struct {
	_ *DomainResource
	// The parameter group supported by the current device component that is based on some nomenclature, e.g. cardiovascular.
	ParameterGroup *CodeableConcept `json:"parameterGroup"`
	// The physical principle of the measurement. For example: thermal, chemical, acoustical, etc.
	MeasurementPrinciple string `json:"measurementPrinciple"`
	// The production specification such as component revision, serial number, etc.
	ProductionSpecification []*DeviceComponent_ProductionSpecification `json:"productionSpecification"`
	// The language code for the human-readable text string produced by the device. This language code will follow the IETF language tag. Example: en-US.
	LanguageCode *CodeableConcept `json:"languageCode"`
	// Extensions for lastSystemChange
	LastSystemChange_ext *Element `json:"_lastSystemChange"`
	// The link to the source Device that contains administrative device information such as manufacture, serial number, etc.
	Source *Reference `json:"source"`
	// The current operational status of the device. For example: On, Off, Standby, etc.
	OperationalStatus []*CodeableConcept `json:"operationalStatus"`
	// The timestamp for the most recent system change which includes device configuration or setting change.
	LastSystemChange string `json:"lastSystemChange"`
	// The link to the parent resource. For example: Channel is linked to its VMD parent.
	Parent *Reference `json:"parent"`
	// Extensions for measurementPrinciple
	MeasurementPrinciple_ext *Element `json:"_measurementPrinciple"`
	// This is a DeviceComponent resource
	ResourceType string `json:"resourceType,omitempty"`
	// The locally assigned unique identification by the software. For example: handle ID.
	Identifier *Identifier `json:"identifier,omitempty"`
	// The component type as defined in the object-oriented or metric nomenclature partition.
	Type *CodeableConcept `json:"type,omitempty"`
}

func NewDeviceComponent() *DeviceComponent { return &DeviceComponent{ResourceType: "DeviceComponent"} }

// ProcessResponse is This resource provides processing status, errors and notes from the processing of a resource.
type ProcessResponse struct {
	_ *DomainResource
	// Suite of processing notes or additional requirements if the processing has been held.
	ProcessNote []*ProcessResponse_ProcessNote `json:"processNote"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// Original request resource reference.
	Request *Reference `json:"request"`
	// Transaction status: error, complete, held.
	Outcome *CodeableConcept `json:"outcome"`
	// The form to be used for printing the content.
	Form *CodeableConcept `json:"form"`
	// Processing errors.
	Error []*CodeableConcept `json:"error"`
	// Request for additional supporting or authorizing information, such as: documents, images or resources.
	CommunicationRequest []*Reference `json:"communicationRequest"`
	// This is a ProcessResponse resource
	ResourceType string `json:"resourceType,omitempty"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// The date when the enclosed suite of services were performed or completed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// The organization which is responsible for the services rendered to the patient.
	RequestOrganization *Reference `json:"requestOrganization"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The organization who produced this adjudicated response.
	Organization *Reference `json:"organization"`
	// A description of the status of the adjudication or processing.
	Disposition string `json:"disposition"`
	// Extensions for disposition
	Disposition_ext *Element `json:"_disposition"`
	// The Response business identifier.
	Identifier []*Identifier `json:"identifier"`
	// The practitioner who is responsible for the services rendered to the patient.
	RequestProvider *Reference `json:"requestProvider"`
}

func NewProcessResponse() *ProcessResponse { return &ProcessResponse{ResourceType: "ProcessResponse"} }

// TestReport_Action is A summary of information based on the results of executing a TestScript.
type TestReport_Action struct {
	_ *BackboneElement
	// The operation performed.
	Operation *TestReport_Operation `json:"operation"`
	// The results of the assertion performed on the previous operations.
	Assert *TestReport_Assert `json:"assert"`
}

// HumanName is A human's name with the ability to identify parts and usage.
type HumanName struct {
	_ *Element
	// Extensions for given
	Given_ext []*Element `json:"_given"`
	// Extensions for prefix
	Prefix_ext []*Element `json:"_prefix"`
	// Extensions for suffix
	Suffix_ext []*Element `json:"_suffix"`
	// Indicates the period of time when this name was valid for the named person.
	Period *Period `json:"period"`
	// Identifies the purpose for this name.
	Use string `json:"use"`
	// Extensions for use
	Use_ext *Element `json:"_use"`
	// A full text representation of the name.
	Text string `json:"text"`
	// Given name.
	Given []string `json:"given"`
	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.
	Suffix []string `json:"suffix"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// The part of a name that links to the genealogy. In some cultures (e.g. Eritrea) the family name of a son is the first name of his father.
	Family string `json:"family"`
	// Extensions for family
	Family_ext *Element `json:"_family"`
	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.
	Prefix []string `json:"prefix"`
}

// CareTeam_Participant is The Care Team includes all the people and organizations who plan to participate in the coordination and delivery of care for a patient.
type CareTeam_Participant struct {
	_ *BackboneElement
	// Indicates specific responsibility of an individual within the care team, such as "Primary care physician", "Trained social worker counselor", "Caregiver", etc.
	Role *CodeableConcept `json:"role"`
	// The specific person or organization who is participating/expected to participate in the care team.
	Member *Reference `json:"member"`
	// The organization of the practitioner.
	OnBehalfOf *Reference `json:"onBehalfOf"`
	// Indicates when the specific member or organization did (or is intended to) come into effect and end.
	Period *Period `json:"period"`
}

// ClaimResponse_Adjudication is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse_Adjudication struct {
	_ *BackboneElement
	// A non-monetary value for example a percentage. Mutually exclusive to the amount element above.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Value float64 `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// Code indicating: Co-Pay, deductible, eligible, benefit, tax, etc.
	Category *CodeableConcept `json:"category,omitempty"`
	// Adjudication reason such as limit reached.
	Reason *CodeableConcept `json:"reason"`
	// Monetary amount associated with the code.
	Amount *Money `json:"amount"`
}

// Consent is A record of a healthcare consumer's policy choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent struct {
	_ *DomainResource
	// The source on which this consent statement is based. The source might be a scanned original paper form, or a reference to a consent that links back to such a source, a reference to a document repository (e.g. XDS) that stores the original consent document.
	SourceIdentifier *Identifier `json:"sourceIdentifier"`
	// The source on which this consent statement is based. The source might be a scanned original paper form, or a reference to a consent that links back to such a source, a reference to a document repository (e.g. XDS) that stores the original consent document.
	SourceReference *Reference `json:"sourceReference"`
	// The context of the activities a user is taking - why the user is accessing the data - that are controlled by this consent.
	Purpose []*Coding `json:"purpose"`
	// Indicates the current state of this consent.
	Status string `json:"status"`
	// The patient/healthcare consumer to whom this consent applies.
	Patient *Reference `json:"patient,omitempty"`
	// Extensions for dateTime
	DateTime_ext *Element `json:"_dateTime"`
	// An exception to the base policy of this consent. An exception can be an addition or removal of access permissions.
	Except []*Consent_Except `json:"except"`
	// Either the Grantor, which is the entity responsible for granting the rights listed in a Consent Directive or the Grantee, which is the entity responsible for complying with the Consent Directive, including any obligations or limitations on authorizations and enforcement of prohibitions.
	ConsentingParty []*Reference `json:"consentingParty"`
	// Actions controlled by this consent.
	Action []*CodeableConcept `json:"action"`
	// A set of security labels that define which resources are controlled by this consent. If more than one label is specified, all resources must have all the specified labels.
	SecurityLabel []*Coding `json:"securityLabel"`
	// Relevant time or time-period when this Consent is applicable.
	Period *Period `json:"period"`
	// Who or what is controlled by this consent. Use group to identify a set of actors by some property they share (e.g. 'admitting officers').
	Actor []*Consent_Actor `json:"actor"`
	// The organization that manages the consent, and the framework within which it is executed.
	Organization []*Reference `json:"organization"`
	// The source on which this consent statement is based. The source might be a scanned original paper form, or a reference to a consent that links back to such a source, a reference to a document repository (e.g. XDS) that stores the original consent document.
	SourceAttachment *Attachment `json:"sourceAttachment"`
	// The references to the policies that are included in this consent scope. Policies may be organizational, but are often defined jurisdictionally, or in law.
	Policy []*Consent_Policy `json:"policy"`
	// This is a Consent resource
	ResourceType string `json:"resourceType,omitempty"`
	// Unique identifier for this copy of the Consent Statement.
	Identifier *Identifier `json:"identifier"`
	// A classification of the type of consents found in the statement. This element supports indexing and retrieval of consent statements.
	Category []*CodeableConcept `json:"category"`
	// Extensions for policyRule
	PolicyRule_ext *Element `json:"_policyRule"`
	// Clinical or Operational Relevant period of time that bounds the data controlled by this consent.
	DataPeriod *Period `json:"dataPeriod"`
	// The resources controlled by this consent, if specific resources are referenced.
	Data []*Consent_Data `json:"data"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// When this  Consent was issued / created / indexed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	DateTime time.Time `json:"dateTime"`
	// A referece to the specific computable policy.
	PolicyRule string `json:"policyRule"`
}

func NewConsent() *Consent { return &Consent{ResourceType: "Consent"} }

// List is A set of information summarized from a list of other resources.
type List struct {
	_ *DomainResource
	// What order applies to the items in the list.
	OrderedBy *CodeableConcept `json:"orderedBy"`
	// Identifier for the List assigned for business purposes outside the context of FHIR.
	Identifier []*Identifier `json:"identifier"`
	// Indicates the current state of this list.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// The common subject (or patient) of the resources that are in the list, if there is one.
	Subject *Reference `json:"subject"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// This is a List resource
	ResourceType string `json:"resourceType,omitempty"`
	// A label for the list assigned by the author.
	Title string `json:"title"`
	// The encounter that is the context in which this list was created.
	Encounter *Reference `json:"encounter"`
	// The date that the list was prepared.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// How this list was prepared - whether it is a working list that is suitable for being maintained on an ongoing basis, or if it represents a snapshot of a list of items from another source, or whether it is a prepared list where items may be marked as added, modified or deleted.
	Mode string `json:"mode"`
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// This code defines the purpose of the list - why it was created.
	Code *CodeableConcept `json:"code"`
	// The entity responsible for deciding what the contents of the list were. Where the list was created by a human, this is the same as the author of the list.
	Source *Reference `json:"source"`
	// Comments that apply to the overall list.
	Note []*Annotation `json:"note"`
	// If the list is empty, why the list is empty.
	EmptyReason *CodeableConcept `json:"emptyReason"`
	// Entries in this list.
	Entry []*List_Entry `json:"entry"`
}

func NewList() *List { return &List{ResourceType: "List"} }

// ConceptMap_Element is A statement of relationships from one set of concepts to one or more other concepts - either code systems or data elements, or classes in class models.
type ConceptMap_Element struct {
	_ *BackboneElement
	// Extensions for display
	Display_ext *Element `json:"_display"`
	// A concept from the target value set that this concept maps to.
	Target []*ConceptMap_Target `json:"target"`
	// Identity (code or path) or the element/item being mapped.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The display for the code. The display is only provided to help editors when editing the concept map.
	Display string `json:"display"`
}

// DetectedIssue_Mitigation is Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. Drug-drug interaction, Ineffective treatment frequency, Procedure-condition conflict, etc.
type DetectedIssue_Mitigation struct {
	_ *BackboneElement
	// Describes the action that was taken or the observation that was made that reduces/eliminates the risk associated with the identified issue.
	Action *CodeableConcept `json:"action,omitempty"`
	// Indicates when the mitigating action was documented.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Identifies the practitioner who determined the mitigation and takes responsibility for the mitigation step occurring.
	Author *Reference `json:"author"`
}

// GraphDefinition is A formal computable definition of a graph of resources - that is, a coherent set of resources that form a graph by following references. The Graph Definition resource defines a set and makes rules about the set.
type GraphDefinition struct {
	_ *DomainResource
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate graph definition instances.
	UseContext []*UsageContext `json:"useContext"`
	// A legal or geographic region in which the graph definition is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Explaination of why this graph definition is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// Links this graph makes rules about.
	Link []*GraphDefinition_Link `json:"link"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// The date  (and optionally time) when the graph definition was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the graph definition changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// A free text natural language description of the graph definition from a consumer's perspective.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for start
	Start_ext *Element `json:"_start"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// A natural language name identifying the graph definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// The status of this graph definition. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// An absolute URI that is used to identify this graph definition when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this graph definition is (or will be) published. The URL SHOULD include the major version of the graph definition. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// The profile that describes the use of the base resource.
	Profile string `json:"profile"`
	// The type of FHIR resource at which instances of this graph start.
	// pattern [^\s]+([\s]?[^\s]+)*
	Start string `json:"start"`
	// Extensions for profile
	Profile_ext *Element `json:"_profile"`
	// This is a GraphDefinition resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The name of the individual or organization that published the graph definition.
	Publisher string `json:"publisher"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// The identifier that is used to identify this version of the graph definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the graph definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// A boolean value to indicate that this graph definition is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
}

func NewGraphDefinition() *GraphDefinition { return &GraphDefinition{ResourceType: "GraphDefinition"} }

// NutritionOrder is A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrder struct {
	_ *DomainResource
	// This modifier is used to convey order-specific modifiers about the type of food that should NOT be given. These can be derived from patient allergies, intolerances, or preferences such as No Red Meat, No Soy or No Wheat or  Gluten-Free.  While it should not be necessary to repeat allergy or intolerance information captured in the referenced AllergyIntolerance resource in the excludeFoodModifier, this element may be used to convey additional specificity related to foods that should be eliminated from the patient's diet for any reason.  This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional supplements and enteral formula feedings.
	ExcludeFoodModifier []*CodeableConcept `json:"excludeFoodModifier"`
	// Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma that delivers nutrition distal to the oral cavity.
	EnteralFormula *NutritionOrder_EnteralFormula `json:"enteralFormula"`
	// This is a NutritionOrder resource
	ResourceType string `json:"resourceType,omitempty"`
	// An encounter that provides additional information about the healthcare context in which this request is made.
	Encounter *Reference `json:"encounter"`
	// The date and time that this nutrition order was requested.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	DateTime time.Time `json:"dateTime"`
	// Oral nutritional products given in order to add further nutritional value to the patient's diet.
	Supplement []*NutritionOrder_Supplement `json:"supplement"`
	// Identifiers assigned to this order by the order sender or by the order receiver.
	Identifier []*Identifier `json:"identifier"`
	// The workflow status of the nutrition order/request.
	Status string `json:"status"`
	// Diet given orally in contrast to enteral (tube) feeding.
	OralDiet *NutritionOrder_OralDiet `json:"oralDiet"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for dateTime
	DateTime_ext *Element `json:"_dateTime"`
	// This modifier is used to convey order-specific modifiers about the type of food that should be given. These can be derived from patient allergies, intolerances, or preferences such as Halal, Vegan or Kosher. This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional supplements and enteral formula feedings.
	FoodPreferenceModifier []*CodeableConcept `json:"foodPreferenceModifier"`
	// The person (patient) who needs the nutrition order for an oral diet, nutritional supplement and/or enteral or formula feeding.
	Patient *Reference `json:"patient,omitempty"`
	// The practitioner that holds legal responsibility for ordering the diet, nutritional supplement, or formula feedings.
	Orderer *Reference `json:"orderer"`
	// A link to a record of allergies or intolerances  which should be included in the nutrition order.
	AllergyIntolerance []*Reference `json:"allergyIntolerance"`
}

func NewNutritionOrder() *NutritionOrder { return &NutritionOrder{ResourceType: "NutritionOrder"} }

// Count is A measured amount (or an amount that can potentially be measured). Note that measured amounts include amounts that are not precisely quantified, including amounts involving arbitrary units and floating currencies.
type Count struct {
	_ *Quantity
}

// CompartmentDefinition is A compartment definition that defines how resources are accessed on a server.
type CompartmentDefinition struct {
	_ *DomainResource
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate compartment definition instances.
	UseContext []*UsageContext `json:"useContext"`
	// This is a CompartmentDefinition resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// A boolean value to indicate that this compartment definition is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// A free text natural language description of the compartment definition from a consumer's perspective.
	Description string `json:"description"`
	// Explaination of why this compartment definition is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// Whether the search syntax is supported,.
	Search bool `json:"search"`
	// Which compartment this definition describes.
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The date  (and optionally time) when the compartment definition was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the compartment definition changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// An absolute URI that is used to identify this compartment definition when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this compartment definition is (or will be) published. The URL SHOULD include the major version of the compartment definition. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// The status of this compartment definition. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Information about how a resource is related to the compartment.
	Resource []*CompartmentDefinition_Resource `json:"resource"`
	// Extensions for search
	Search_ext *Element `json:"_search"`
	// A short, descriptive, user-friendly title for the compartment definition.
	Title string `json:"title"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// The name of the individual or organization that published the compartment definition.
	Publisher string `json:"publisher"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// A natural language name identifying the compartment definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// A legal or geographic region in which the compartment definition is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
}

func NewCompartmentDefinition() *CompartmentDefinition {
	return &CompartmentDefinition{ResourceType: "CompartmentDefinition"}
}

// ExplanationOfBenefit_Item is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Item struct {
	_ *BackboneElement
	// A service line number.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// Health Care Service Type Codes  to identify the classification of service or benefits.
	Category *CodeableConcept `json:"category"`
	// For programs which require reson codes for the inclusion, covering, of this billed item under the program or sub-program.
	ProgramCode []*CodeableConcept `json:"programCode"`
	// Where the service was provided.
	LocationReference *Reference `json:"locationReference"`
	// Extensions for noteNumber
	NoteNumber_ext []*Element `json:"_noteNumber"`
	// Extensions for procedureLinkId
	ProcedureLinkId_ext []*Element `json:"_procedureLinkId"`
	// The type of reveneu or cost center providing the product and/or service.
	Revenue *CodeableConcept `json:"revenue"`
	// Where the service was provided.
	LocationCodeableConcept *CodeableConcept `json:"locationCodeableConcept"`
	// Where the service was provided.
	LocationAddress *Address `json:"locationAddress"`
	// Careteam applicable for this service or product line.
	CareTeamLinkId []uint64 `json:"careTeamLinkId"`
	// Procedures applicable for this service or product line.
	ProcedureLinkId []uint64 `json:"procedureLinkId"`
	// Extensions for informationLinkId
	InformationLinkId_ext []*Element `json:"_informationLinkId"`
	// The date or dates when the enclosed suite of services were performed or completed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ServicedDate time.Time `json:"servicedDate"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// Extensions for diagnosisLinkId
	DiagnosisLinkId_ext []*Element `json:"_diagnosisLinkId"`
	// A region or surface of the site, eg. limb region or tooth surface(s).
	SubSite []*CodeableConcept `json:"subSite"`
	// If this is an actual service or product line, ie. not a Group, then use code to indicate the Professional Service or Product supplied (eg. CTP, HCPCS,USCLS,ICD10, NCPDP,DIN,ACHI,CCI). If a grouping item then use a group code to indicate the type of thing being grouped eg. 'glasses' or 'compound'.
	Service *CodeableConcept `json:"service"`
	// Item typification or modifiers codes, eg for Oral whether the treatment is cosmetic or associated with TMJ, or for medical whether the treatment was outside the clinic or out of office hours.
	Modifier []*CodeableConcept `json:"modifier"`
	// The number of repetitions of a service or product.
	Quantity *Quantity `json:"quantity"`
	// List of Unique Device Identifiers associated with this line item.
	Udi []*Reference `json:"udi"`
	// Physical service site on the patient (limb, tooth, etc).
	BodySite *CodeableConcept `json:"bodySite"`
	// Diagnosis applicable for this service or product line.
	DiagnosisLinkId []uint64 `json:"diagnosisLinkId"`
	// Exceptions, special conditions and supporting information pplicable for this service or product line.
	InformationLinkId []uint64 `json:"informationLinkId"`
	// Extensions for servicedDate
	ServicedDate_ext *Element `json:"_servicedDate"`
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Factor float64 `json:"factor"`
	// Second tier of goods and services.
	Detail []*ExplanationOfBenefit_Detail `json:"detail"`
	// If the item is a node then this is the fee for the product or service, otherwise this is the total of the fees for the children of the group.
	UnitPrice *Money `json:"unitPrice"`
	// Extensions for factor
	Factor_ext *Element `json:"_factor"`
	// A billed item may include goods or services provided in multiple encounters.
	Encounter []*Reference `json:"encounter"`
	// A list of note references to the notes provided below.
	NoteNumber []uint64 `json:"noteNumber"`
	// The adjudications results.
	Adjudication []*ExplanationOfBenefit_Adjudication `json:"adjudication"`
	// Extensions for careTeamLinkId
	CareTeamLinkId_ext []*Element `json:"_careTeamLinkId"`
	// The date or dates when the enclosed suite of services were performed or completed.
	ServicedPeriod *Period `json:"servicedPeriod"`
	// The quantity times the unit price for an addittional service or product or charge. For example, the formula: unit Quantity * unit Price (Cost per Point) * factor Number  * points = net Amount. Quantity, factor and points are assumed to be 1 if not supplied.
	Net *Money `json:"net"`
}

// RequestGroup_RelatedAction is A group of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
type RequestGroup_RelatedAction struct {
	_ *BackboneElement
	// The relationship of this action to the related action.
	// pattern [^\s]+([\s]?[^\s]+)*
	Relationship string `json:"relationship"`
	// Extensions for relationship
	Relationship_ext *Element `json:"_relationship"`
	// A duration or range of durations to apply to the relationship. For example, 30-60 minutes before.
	OffsetDuration *Duration `json:"offsetDuration"`
	// A duration or range of durations to apply to the relationship. For example, 30-60 minutes before.
	OffsetRange *Range `json:"offsetRange"`
	// The element id of the action this is related to.
	// pattern [A-Za-z0-9\-\.]{1,64}
	ActionId string `json:"actionId"`
	// Extensions for actionId
	ActionId_ext *Element `json:"_actionId"`
}

// Identifier is A technical identifier - identifies some entity uniquely and unambiguously.
type Identifier struct {
	_ *Element
	// Organization that issued/manages the identifier.
	Assigner *Reference `json:"assigner"`
	// The purpose of this identifier.
	Use string `json:"use"`
	// A coded type for the identifier that can be used to determine which identifier to use for a specific purpose.
	Type *CodeableConcept `json:"type"`
	// Establishes the namespace for the value - that is, a URL that describes a set values that are unique.
	System string `json:"system"`
	// Extensions for system
	System_ext *Element `json:"_system"`
	// The portion of the identifier typically relevant to the user and which is unique within the context of the system.
	Value string `json:"value"`
	// Extensions for use
	Use_ext *Element `json:"_use"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// Time period during which identifier is/was valid for use.
	Period *Period `json:"period"`
}

// DocumentReference_Context is A reference to a document.
type DocumentReference_Context struct {
	_ *BackboneElement
	// The time period over which the service that is described by the document was provided.
	Period *Period `json:"period"`
	// The kind of facility where the patient was seen.
	FacilityType *CodeableConcept `json:"facilityType"`
	// This property may convey specifics about the practice setting where the content was created, often reflecting the clinical specialty.
	PracticeSetting *CodeableConcept `json:"practiceSetting"`
	// The Patient Information as known when the document was published. May be a reference to a version specific, or contained.
	SourcePatientInfo *Reference `json:"sourcePatientInfo"`
	// Related identifiers or resources associated with the DocumentReference.
	Related []*DocumentReference_Related `json:"related"`
	// Describes the clinical encounter or type of care that the document content is associated with.
	Encounter *Reference `json:"encounter"`
	// This list of codes represents the main clinical acts, such as a colonoscopy or an appendectomy, being documented. In some cases, the event is inherent in the typeCode, such as a "History and Physical Report" in which the procedure being documented is necessarily a "History and Physical" act.
	Event []*CodeableConcept `json:"event"`
}

// Task_Requester is A task to be performed.
type Task_Requester struct {
	_ *BackboneElement
	// The device, practitioner, etc. who initiated the task.
	Agent *Reference `json:"agent,omitempty"`
	// The organization the device or practitioner was acting on behalf of when they initiated the task.
	OnBehalfOf *Reference `json:"onBehalfOf"`
}

// CapabilityStatement is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement struct {
	_ *DomainResource
	// A list of the patch formats supported by this implementation using their content types.
	PatchFormat []string `json:"patchFormat"`
	// A definition of the restful capabilities of the solution, if any.
	Rest []*CapabilityStatement_Rest `json:"rest"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// A copyright statement relating to the capability statement and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the capability statement.
	Copyright string `json:"copyright"`
	// Extensions for instantiates
	Instantiates_ext []*Element `json:"_instantiates"`
	// Software that is covered by this capability statement.  It is used when the capability statement describes the capabilities of a particular software version, independent of an installation.
	Software *CapabilityStatement_Software `json:"software"`
	// Extensions for patchFormat
	PatchFormat_ext []*Element `json:"_patchFormat"`
	// The name of the individual or organization that published the capability statement.
	Publisher string `json:"publisher"`
	// Explaination of why this capability statement is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// Reference to a canonical URL of another CapabilityStatement that this software implements or uses. This capability statement is a published API description that corresponds to a business service. The rest of the capability statement does not need to repeat the details of the referenced resource, but can do so.
	Instantiates []string `json:"instantiates"`
	// Identifies a specific implementation instance that is described by the capability statement - i.e. a particular installation, rather than the capabilities of a software program.
	Implementation *CapabilityStatement_Implementation `json:"implementation"`
	// Extensions for format
	Format_ext []*Element `json:"_format"`
	// The status of this capability statement. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// Extensions for fhirVersion
	FhirVersion_ext *Element `json:"_fhirVersion"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// A list of the formats supported by this implementation using their content types.
	Format []string `json:"format"`
	// An absolute URI that is used to identify this capability statement when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this capability statement is (or will be) published. The URL SHOULD include the major version of the capability statement. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// A short, descriptive, user-friendly title for the capability statement.
	Title string `json:"title"`
	// A boolean value to indicate that this capability statement is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// The date  (and optionally time) when the capability statement was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the capability statement changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// A list of profiles that represent different use cases supported by the system. For a server, "supported by the system" means the system hosts/produces a set of resources that are conformant to a particular profile, and allows clients that use its services to search using this profile and to find appropriate data. For a client, it means the system will search by this profile and process data according to the guidance implicit in the profile. See further discussion in [Using Profiles](profiling.html#profile-uses).
	Profile []*Reference `json:"profile"`
	// A description of the messaging capabilities of the solution.
	Messaging []*CapabilityStatement_Messaging `json:"messaging"`
	// A free text natural language description of the capability statement from a consumer's perspective. Typically, this is used when the capability statement describes a desired rather than an actual solution, for example as a formal expression of requirements as part of an RFP.
	Description string `json:"description"`
	// Extensions for kind
	Kind_ext *Element `json:"_kind"`
	// A natural language name identifying the capability statement. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// The way that this statement is intended to be used, to describe an actual running instance of software, a particular product (kind not instance of software) or a class of implementation (e.g. a desired purchase).
	Kind string `json:"kind"`
	// Extensions for acceptUnknown
	AcceptUnknown_ext *Element `json:"_acceptUnknown"`
	// Extensions for implementationGuide
	ImplementationGuide_ext []*Element `json:"_implementationGuide"`
	// A list of implementation guides that the server does (or should) support in their entirety.
	ImplementationGuide []string `json:"implementationGuide"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// A legal or geographic region in which the capability statement is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// A code that indicates whether the application accepts unknown elements or extensions when reading resources.
	AcceptUnknown string `json:"acceptUnknown"`
	// A document definition.
	Document []*CapabilityStatement_Document `json:"document"`
	// This is a CapabilityStatement resource
	ResourceType string `json:"resourceType,omitempty"`
	// The identifier that is used to identify this version of the capability statement when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the capability statement author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate capability statement instances.
	UseContext []*UsageContext `json:"useContext"`
	// The version of the FHIR specification on which this capability statement is based.
	// pattern [A-Za-z0-9\-\.]{1,64}
	FhirVersion string `json:"fhirVersion"`
}

func NewCapabilityStatement() *CapabilityStatement {
	return &CapabilityStatement{ResourceType: "CapabilityStatement"}
}

// Measure_Group is The Measure resource provides the definition of a quality measure.
type Measure_Group struct {
	_ *BackboneElement
	// The human readable description of this population group.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A population criteria for the measure.
	Population []*Measure_Population `json:"population"`
	// The stratifier criteria for the measure report, specified as either the name of a valid CQL expression defined within a referenced library, or a valid FHIR Resource Path.
	Stratifier []*Measure_Stratifier `json:"stratifier"`
	// A unique identifier for the group. This identifier will used to report data for the group in the measure report.
	Identifier *Identifier `json:"identifier,omitempty"`
	// Optional name or short description of this group.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
}

// ExplanationOfBenefit_Financial is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Financial struct {
	_ *BackboneElement
	// Extensions for allowedString
	AllowedString_ext *Element `json:"_allowedString"`
	// Benefits allowed.
	AllowedMoney *Money `json:"allowedMoney"`
	// Extensions for usedUnsignedInt
	UsedUnsignedInt_ext *Element `json:"_usedUnsignedInt"`
	// Benefits used.
	UsedMoney *Money `json:"usedMoney"`
	// Deductable, visits, benefit amount.
	Type *CodeableConcept `json:"type,omitempty"`
	// Benefits allowed.
	// pattern [0]|([1-9][0-9]*)
	AllowedUnsignedInt uint64 `json:"allowedUnsignedInt"`
	// Extensions for allowedUnsignedInt
	AllowedUnsignedInt_ext *Element `json:"_allowedUnsignedInt"`
	// Benefits allowed.
	AllowedString string `json:"allowedString"`
	// Benefits used.
	// pattern [0]|([1-9][0-9]*)
	UsedUnsignedInt uint64 `json:"usedUnsignedInt"`
}

// ImagingManifest is A text description of the DICOM SOP instances selected in the ImagingManifest; or the reason for, or significance of, the selection.
type ImagingManifest struct {
	_ *DomainResource
	// Study identity and locating information of the DICOM SOP instances in the selection.
	Study []*ImagingManifest_Study `json:"study,omitempty"`
	// Unique identifier of the DICOM Key Object Selection (KOS) that this resource represents.
	Identifier *Identifier `json:"identifier"`
	// A patient resource reference which is the patient subject of all DICOM SOP Instances in this ImagingManifest.
	Patient *Reference `json:"patient,omitempty"`
	// Date and time when the selection of the referenced instances were made. It is (typically) different from the creation date of the selection resource, and from dates associated with the referenced instances (e.g. capture time of the referenced image).
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AuthoringTime time.Time `json:"authoringTime"`
	// Extensions for authoringTime
	AuthoringTime_ext *Element `json:"_authoringTime"`
	// Author of ImagingManifest. It can be a human author or a device which made the decision of the SOP instances selected. For example, a radiologist selected a set of imaging SOP instances to attach in a diagnostic report, and a CAD application may author a selection to describe SOP instances it used to generate a detection conclusion.
	Author *Reference `json:"author"`
	// Free text narrative description of the ImagingManifest.
	// The value may be derived from the DICOM Standard Part 16, CID-7010 descriptions (e.g. Best in Set, Complete Study Content). Note that those values cover the wide range of uses of the DICOM Key Object Selection object, several of which are not supported by ImagingManifest. Specifically, there is no expected behavior associated with descriptions that suggest referenced images be removed or not used.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// This is a ImagingManifest resource
	ResourceType string `json:"resourceType,omitempty"`
}

func NewImagingManifest() *ImagingManifest { return &ImagingManifest{ResourceType: "ImagingManifest"} }

// Patient_Contact is Demographics and other administrative information about an individual or animal receiving care or other health-related services.
type Patient_Contact struct {
	_ *BackboneElement
	// Address for the contact person.
	Address *Address `json:"address"`
	// Administrative Gender - the gender that the contact person is considered to have for administration and record keeping purposes.
	Gender string `json:"gender"`
	// Extensions for gender
	Gender_ext *Element `json:"_gender"`
	// Organization on behalf of which the contact is acting or for which the contact is working.
	Organization *Reference `json:"organization"`
	// The period during which this contact person or organization is valid to be contacted relating to this patient.
	Period *Period `json:"period"`
	// The nature of the relationship between the patient and the contact person.
	Relationship []*CodeableConcept `json:"relationship"`
	// A name associated with the contact person.
	Name *HumanName `json:"name"`
	// A contact detail for the person, e.g. a telephone number or an email address.
	Telecom []*ContactPoint `json:"telecom"`
}

// Basic is Basic is used for handling concepts not yet defined in FHIR, narrative-only resources that don't map to an existing resource, and custom resources not appropriate for inclusion in the FHIR specification.
type Basic struct {
	_ *DomainResource
	// Identifies the 'type' of resource - equivalent to the resource name for other resources.
	Code *CodeableConcept `json:"code,omitempty"`
	// Identifies the patient, practitioner, device or any other resource that is the "focus" of this resource.
	Subject *Reference `json:"subject"`
	// Identifies when the resource was first created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	Created time.Time `json:"created"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// Indicates who was responsible for creating the resource instance.
	Author *Reference `json:"author"`
	// This is a Basic resource
	ResourceType string `json:"resourceType,omitempty"`
	// Identifier assigned to the resource for business purposes, outside the context of FHIR.
	Identifier []*Identifier `json:"identifier"`
}

func NewBasic() *Basic { return &Basic{ResourceType: "Basic"} }

// Device_Udi is This resource identifies an instance or a type of a manufactured item that is used in the provision of healthcare without being substantially changed through that activity. The device may be a medical or non-medical device.  Medical devices include durable (reusable) medical equipment, implantable devices, as well as disposable equipment used for diagnostic, treatment, and research for healthcare and public health.  Non-medical devices may include items such as a machine, cellphone, computer, application, etc.
type Device_Udi struct {
	_ *BackboneElement
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for carrierHRF
	CarrierHRF_ext *Element `json:"_carrierHRF"`
	// Organization that is charged with issuing UDIs for devices.  For example, the US FDA issuers include :
	// 1) GS1:
	// http://hl7.org/fhir/NamingSystem/gs1-di,
	// 2) HIBCC:
	// http://hl7.org/fhir/NamingSystem/hibcc-dI,
	// 3) ICCBBA for blood containers:
	// http://hl7.org/fhir/NamingSystem/iccbba-blood-di,
	// 4) ICCBA for other devices:
	// http://hl7.org/fhir/NamingSystem/iccbba-other-di.
	Issuer string `json:"issuer"`
	// Extensions for jurisdiction
	Jurisdiction_ext *Element `json:"_jurisdiction"`
	// The full UDI carrier as the human readable form (HRF) representation of the barcode string as printed on the packaging of the device.
	CarrierHRF string `json:"carrierHRF"`
	// Extensions for carrierAIDC
	CarrierAIDC_ext *Element `json:"_carrierAIDC"`
	// Extensions for entryType
	EntryType_ext *Element `json:"_entryType"`
	// The device identifier (DI) is a mandatory, fixed portion of a UDI that identifies the labeler and the specific version or model of a device.
	DeviceIdentifier string `json:"deviceIdentifier"`
	// Extensions for deviceIdentifier
	DeviceIdentifier_ext *Element `json:"_deviceIdentifier"`
	// Name of device as used in labeling or catalog.
	Name string `json:"name"`
	// The full UDI carrier of the Automatic Identification and Data Capture (AIDC) technology representation of the barcode string as printed on the packaging of the device - E.g a barcode or RFID.   Because of limitations on character sets in XML and the need to round-trip JSON data through XML, AIDC Formats *SHALL* be base64 encoded.
	CarrierAIDC string `json:"carrierAIDC"`
	// Extensions for issuer
	Issuer_ext *Element `json:"_issuer"`
	// The identity of the authoritative source for UDI generation within a  jurisdiction.  All UDIs are globally unique within a single namespace. with the appropriate repository uri as the system.  For example,  UDIs of devices managed in the U.S. by the FDA, the value is  http://hl7.org/fhir/NamingSystem/fda-udi.
	Jurisdiction string `json:"jurisdiction"`
	// A coded entry to indicate how the data was entered.
	EntryType string `json:"entryType"`
}

// ImplementationGuide_Dependency is A set of rules of how FHIR is used to solve a particular problem. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuide_Dependency struct {
	_ *BackboneElement
	// How the dependency is represented when the guide is published.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Where the dependency is located.
	Uri string `json:"uri"`
	// Extensions for uri
	Uri_ext *Element `json:"_uri"`
}

// QuestionnaireResponse_Answer is A structured set of questions and their answers. The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to.
type QuestionnaireResponse_Answer struct {
	_ *BackboneElement
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// Extensions for valueDecimal
	ValueDecimal_ext *Element `json:"_valueDecimal"`
	// Extensions for valueDate
	ValueDate_ext *Element `json:"_valueDate"`
	// Extensions for valueTime
	ValueTime_ext *Element `json:"_valueTime"`
	// The answer (or one of the answers) provided by the respondent to the question.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// The answer (or one of the answers) provided by the respondent to the question.
	ValueReference *Reference `json:"valueReference"`
	// Extensions for valueInteger
	ValueInteger_ext *Element `json:"_valueInteger"`
	// The answer (or one of the answers) provided by the respondent to the question.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ValueDateTime time.Time `json:"valueDateTime"`
	// Extensions for valueDateTime
	ValueDateTime_ext *Element `json:"_valueDateTime"`
	// The answer (or one of the answers) provided by the respondent to the question.
	ValueAttachment *Attachment `json:"valueAttachment"`
	// The answer (or one of the answers) provided by the respondent to the question.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	ValueDecimal float64 `json:"valueDecimal"`
	// The answer (or one of the answers) provided by the respondent to the question.
	ValueString string `json:"valueString"`
	// The answer (or one of the answers) provided by the respondent to the question.
	ValueUri string `json:"valueUri"`
	// The answer (or one of the answers) provided by the respondent to the question.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	ValueTime time.Time `json:"valueTime"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// Extensions for valueUri
	ValueUri_ext *Element `json:"_valueUri"`
	// The answer (or one of the answers) provided by the respondent to the question.
	ValueCoding *Coding `json:"valueCoding"`
	// Nested groups and/or questions found within this particular answer.
	Item []*QuestionnaireResponse_Item `json:"item"`
	// The answer (or one of the answers) provided by the respondent to the question.
	ValueBoolean bool `json:"valueBoolean"`
	// The answer (or one of the answers) provided by the respondent to the question.
	// pattern -?([0]|([1-9][0-9]*))
	ValueInteger int64 `json:"valueInteger"`
	// The answer (or one of the answers) provided by the respondent to the question.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ValueDate time.Time `json:"valueDate"`
}

// Sequence_Quality is Raw data describing a biological sequence.
type Sequence_Quality struct {
	_ *BackboneElement
	// Extensions for queryTP
	QueryTP_ext *Element `json:"_queryTP"`
	// Extensions for gtFP
	GtFP_ext *Element `json:"_gtFP"`
	// Extensions for recall
	Recall_ext *Element `json:"_recall"`
	// Harmonic mean of Recall and Precision, computed as: 2 * precision * recall / (precision + recall).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	FScore float64 `json:"fScore"`
	// Extensions for start
	Start_ext *Element `json:"_start"`
	// End position of the sequence.If the coordinate system is 0-based then end is is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
	// pattern -?([0]|([1-9][0-9]*))
	End int64 `json:"end"`
	// The score of an experimentally derived feature such as a p-value ([SO:0001685](http://www.sequenceontology.org/browser/current_svn/term/SO:0001685)).
	Score *Quantity `json:"score"`
	// Which method is used to get sequence quality.
	Method *CodeableConcept `json:"method"`
	// Extensions for precision
	Precision_ext *Element `json:"_precision"`
	// INDEL / SNP / Undefined variant.
	Type string `json:"type"`
	// False negatives, i.e. the number of sites in the Truth Call Set for which there is no path through the Query Call Set that is consistent with all of the alleles at this site, or sites for which there is an inaccurate genotype call for the event. Sites with correct variant but incorrect genotype are counted here.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	TruthFN float64 `json:"truthFN"`
	// False positives, i.e. the number of sites in the Query Call Set for which there is no path through the Truth Call Set that is consistent with this site. Sites with correct variant but incorrect genotype are counted here.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	QueryFP float64 `json:"queryFP"`
	// QUERY.TP / (QUERY.TP + QUERY.FP).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Precision float64 `json:"precision"`
	// Extensions for truthFN
	TruthFN_ext *Element `json:"_truthFN"`
	// Extensions for fScore
	FScore_ext *Element `json:"_fScore"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Extensions for end
	End_ext *Element `json:"_end"`
	// True positives, from the perspective of the truth data, i.e. the number of sites in the Truth Call Set for which there are paths through the Query Call Set that are consistent with all of the alleles at this site, and for which there is an accurate genotype call for the event.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	TruthTP float64 `json:"truthTP"`
	// Extensions for truthTP
	TruthTP_ext *Element `json:"_truthTP"`
	// The number of false positives where the non-REF alleles in the Truth and Query Call Sets match (i.e. cases where the truth is 1/1 and the query is 0/1 or similar).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	GtFP float64 `json:"gtFP"`
	// TRUTH.TP / (TRUTH.TP + TRUTH.FN).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Recall float64 `json:"recall"`
	// Gold standard sequence used for comparing against.
	StandardSequence *CodeableConcept `json:"standardSequence"`
	// Start position of the sequence. If the coordinate system is either 0-based or 1-based, then start position is inclusive.
	// pattern -?([0]|([1-9][0-9]*))
	Start int64 `json:"start"`
	// True positives, from the perspective of the query data, i.e. the number of sites in the Query Call Set for which there are paths through the Truth Call Set that are consistent with all of the alleles at this site, and for which there is an accurate genotype call for the event.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	QueryTP float64 `json:"queryTP"`
	// Extensions for queryFP
	QueryFP_ext *Element `json:"_queryFP"`
}

// StructureMap_Target is A Map of relationships between 2 structures that can be used to transform data.
type StructureMap_Target struct {
	_ *BackboneElement
	// Extensions for contextType
	ContextType_ext *Element `json:"_contextType"`
	// Field to create in the context.
	Element string `json:"element"`
	// Named context for field, if desired, and a field is specified.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Variable string `json:"variable"`
	// How the data is copied / created.
	Transform string `json:"transform"`
	// Extensions for listRuleId
	ListRuleId_ext *Element `json:"_listRuleId"`
	// Extensions for transform
	Transform_ext *Element `json:"_transform"`
	// How to interpret the context.
	ContextType string `json:"contextType"`
	// Extensions for variable
	Variable_ext *Element `json:"_variable"`
	// If field is a list, how to manage the list.
	ListMode []string `json:"listMode"`
	// Extensions for listMode
	ListMode_ext []*Element `json:"_listMode"`
	// Type or variable this rule applies to.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Context string `json:"context"`
	// Extensions for context
	Context_ext *Element `json:"_context"`
	// Extensions for element
	Element_ext *Element `json:"_element"`
	// Internal rule reference for shared list items.
	// pattern [A-Za-z0-9\-\.]{1,64}
	ListRuleId string `json:"listRuleId"`
	// Parameters to the transform.
	Parameter []*StructureMap_Parameter `json:"parameter"`
}

// Claim_Payee is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_Payee struct {
	_ *BackboneElement
	// organization | patient | practitioner | relatedperson.
	ResourceType *Coding `json:"resourceType"`
	// Party to be reimbursed: Subscriber, provider, other.
	Party *Reference `json:"party"`
	// Type of Party to be reimbursed: Subscriber, provider, other.
	Type *CodeableConcept `json:"type,omitempty"`
}

// Claim_Procedure is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_Procedure struct {
	_ *BackboneElement
	// The procedure code.
	ProcedureCodeableConcept *CodeableConcept `json:"procedureCodeableConcept"`
	// The procedure code.
	ProcedureReference *Reference `json:"procedureReference"`
	// Sequence of procedures which serves to order and provide a link.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// Date and optionally time the procedure was performed .
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
}

// CodeSystem_Property1 is A code system resource specifies a set of codes drawn from one or more code systems.
type CodeSystem_Property1 struct {
	_ *BackboneElement
	// The value of this property.
	ValueCoding *Coding `json:"valueCoding"`
	// The value of this property.
	// pattern -?([0]|([1-9][0-9]*))
	ValueInteger int64 `json:"valueInteger"`
	// Extensions for valueInteger
	ValueInteger_ext *Element `json:"_valueInteger"`
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The value of this property.
	// pattern [^\s]+([\s]?[^\s]+)*
	ValueCode string `json:"valueCode"`
	// Extensions for valueCode
	ValueCode_ext *Element `json:"_valueCode"`
	// The value of this property.
	ValueString string `json:"valueString"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// The value of this property.
	ValueBoolean bool `json:"valueBoolean"`
	// The value of this property.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ValueDateTime time.Time `json:"valueDateTime"`
	// Extensions for valueDateTime
	ValueDateTime_ext *Element `json:"_valueDateTime"`
	// A code that is a reference to CodeSystem.property.code.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
}

// NutritionOrder_Supplement is A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrder_Supplement struct {
	_ *BackboneElement
	// The time period and frequency at which the supplement(s) should be given.  The supplement should be given for the combination of all schedules if more than one schedule is present.
	Schedule []*Timing `json:"schedule"`
	// The amount of the nutritional supplement to be given.
	Quantity *Quantity `json:"quantity"`
	// Free text or additional instructions or information pertaining to the oral supplement.
	Instruction string `json:"instruction"`
	// Extensions for instruction
	Instruction_ext *Element `json:"_instruction"`
	// The kind of nutritional supplement product required such as a high protein or pediatric clear liquid supplement.
	Type *CodeableConcept `json:"type"`
	// The product or brand name of the nutritional supplement such as "Acme Protein Shake".
	ProductName string `json:"productName"`
	// Extensions for productName
	ProductName_ext *Element `json:"_productName"`
}

// Patient is Demographics and other administrative information about an individual or animal receiving care or other health-related services.
type Patient struct {
	_ *DomainResource
	// Extensions for multipleBirthBoolean
	MultipleBirthBoolean_ext *Element `json:"_multipleBirthBoolean"`
	// Extensions for deceasedBoolean
	DeceasedBoolean_ext *Element `json:"_deceasedBoolean"`
	// Addresses for the individual.
	Address []*Address `json:"address"`
	// Patient's nominated care provider.
	GeneralPractitioner []*Reference `json:"generalPractitioner"`
	// Organization that is the custodian of the patient record.
	ManagingOrganization *Reference `json:"managingOrganization"`
	// Whether this patient record is in active use.
	Active bool `json:"active"`
	// A contact party (e.g. guardian, partner, friend) for the patient.
	Contact []*Patient_Contact `json:"contact"`
	// An identifier for this patient.
	Identifier []*Identifier `json:"identifier"`
	// A contact detail (e.g. a telephone number or an email address) by which the individual may be contacted.
	Telecom []*ContactPoint `json:"telecom"`
	// Extensions for gender
	Gender_ext *Element `json:"_gender"`
	// Indicates whether the patient is part of a multiple (bool) or indicates the actual birth order (integer).
	// pattern -?([0]|([1-9][0-9]*))
	MultipleBirthInteger int64 `json:"multipleBirthInteger"`
	// Link to another patient resource that concerns the same actual patient.
	Link []*Patient_Link `json:"link"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
	// Extensions for deceasedDateTime
	DeceasedDateTime_ext *Element `json:"_deceasedDateTime"`
	// Extensions for multipleBirthInteger
	MultipleBirthInteger_ext *Element `json:"_multipleBirthInteger"`
	// This patient is known to be an animal.
	Animal *Patient_Animal `json:"animal"`
	// This is a Patient resource
	ResourceType string `json:"resourceType,omitempty"`
	// The date of birth for the individual.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	BirthDate time.Time `json:"birthDate"`
	// Indicates if the individual is deceased or not.
	DeceasedBoolean bool `json:"deceasedBoolean"`
	// Indicates whether the patient is part of a multiple (bool) or indicates the actual birth order (integer).
	MultipleBirthBoolean bool `json:"multipleBirthBoolean"`
	// Languages which may be used to communicate with the patient about his or her health.
	Communication []*Patient_Communication `json:"communication"`
	// A name associated with the individual.
	Name []*HumanName `json:"name"`
	// Administrative Gender - the gender that the patient is considered to have for administration and record keeping purposes.
	Gender string `json:"gender"`
	// This field contains a patient's most recent marital (civil) status.
	MaritalStatus *CodeableConcept `json:"maritalStatus"`
	// Image of the patient.
	Photo []*Attachment `json:"photo"`
	// Extensions for birthDate
	BirthDate_ext *Element `json:"_birthDate"`
	// Indicates if the individual is deceased or not.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	DeceasedDateTime time.Time `json:"deceasedDateTime"`
}

func NewPatient() *Patient { return &Patient{ResourceType: "Patient"} }

// PaymentReconciliation_Detail is This resource provides payment details and claim references supporting a bulk payment.
type PaymentReconciliation_Detail struct {
	_ *BackboneElement
	// The Organization which submitted the claim or financial transaction.
	Submitter *Reference `json:"submitter"`
	// The organization which is receiving the payment.
	Payee *Reference `json:"payee"`
	// The date of the invoice or financial resource.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Amount paid for this detail.
	Amount *Money `json:"amount"`
	// Code to indicate the nature of the payment, adjustment, funds advance, etc.
	Type *CodeableConcept `json:"type,omitempty"`
	// The claim or financial resource.
	Request *Reference `json:"request"`
	// The claim response resource.
	Response *Reference `json:"response"`
}

type ResourceList struct {
}

// Contract_Rule is A formal agreement between parties regarding the conduct of business, exchange of information or other matters.
type Contract_Rule struct {
	_ *BackboneElement
	// Computable Contract conveyed using a policy rule language (e.g. XACML, DKAL, SecPal).
	ContentAttachment *Attachment `json:"contentAttachment"`
	// Computable Contract conveyed using a policy rule language (e.g. XACML, DKAL, SecPal).
	ContentReference *Reference `json:"contentReference"`
}

// DetectedIssue is Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. Drug-drug interaction, Ineffective treatment frequency, Procedure-condition conflict, etc.
type DetectedIssue struct {
	_ *DomainResource
	// Business identifier associated with the detected issue record.
	Identifier *Identifier `json:"identifier"`
	// Indicates the status of the detected issue.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Indicates the resource representing the current activity or proposed activity that is potentially problematic.
	Implicated []*Reference `json:"implicated"`
	// Extensions for severity
	Severity_ext *Element `json:"_severity"`
	// Individual or device responsible for the issue being raised.  For example, a decision support application or a pharmacist conducting a medication review.
	Author *Reference `json:"author"`
	// Extensions for detail
	Detail_ext *Element `json:"_detail"`
	// The literature, knowledge-base or similar reference that describes the propensity for the detected issue identified.
	Reference string `json:"reference"`
	// Extensions for reference
	Reference_ext *Element `json:"_reference"`
	// Indicates the degree of importance associated with the identified issue based on the potential impact on the patient.
	Severity string `json:"severity"`
	// Indicates the patient whose record the detected issue is associated with.
	Patient *Reference `json:"patient"`
	// The date or date-time when the detected issue was initially identified.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// A textual explanation of the detected issue.
	Detail string `json:"detail"`
	// Indicates an action that has been taken or is committed to to reduce or eliminate the likelihood of the risk identified by the detected issue from manifesting.  Can also reflect an observation of known mitigating factors that may reduce/eliminate the need for any action.
	Mitigation []*DetectedIssue_Mitigation `json:"mitigation"`
	// This is a DetectedIssue resource
	ResourceType string `json:"resourceType,omitempty"`
	// Identifies the general type of issue identified.
	Category *CodeableConcept `json:"category"`
}

func NewDetectedIssue() *DetectedIssue { return &DetectedIssue{ResourceType: "DetectedIssue"} }

// Sequence_ReferenceSeq is Raw data describing a biological sequence.
type Sequence_ReferenceSeq struct {
	_ *BackboneElement
	// A string like "ACGT".
	ReferenceSeqString string `json:"referenceSeqString"`
	// Extensions for strand
	Strand_ext *Element `json:"_strand"`
	// Start position of the window on the reference sequence. If the coordinate system is either 0-based or 1-based, then start position is inclusive.
	// pattern -?([0]|([1-9][0-9]*))
	WindowStart int64 `json:"windowStart"`
	// End position of the window on the reference sequence. If the coordinate system is 0-based then end is is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
	// pattern -?([0]|([1-9][0-9]*))
	WindowEnd int64 `json:"windowEnd"`
	// A Pointer to another Sequence entity as reference sequence.
	ReferenceSeqPointer *Reference `json:"referenceSeqPointer"`
	// The Genome Build used for reference, following GRCh build versions e.g. 'GRCh 37'.  Version number must be included if a versioned release of a primary build was used.
	GenomeBuild string `json:"genomeBuild"`
	// Extensions for genomeBuild
	GenomeBuild_ext *Element `json:"_genomeBuild"`
	// Reference identifier of reference sequence submitted to NCBI. It must match the type in the Sequence.type field. For example, the prefix, "NG_" identifies reference sequence for genes, "NM_" for messenger RNA transcripts, and "NP_" for amino acid sequences.
	ReferenceSeqId *CodeableConcept `json:"referenceSeqId"`
	// Extensions for referenceSeqString
	ReferenceSeqString_ext *Element `json:"_referenceSeqString"`
	// Directionality of DNA sequence. Available values are "1" for the plus strand (5' to 3')/Watson/Sense/positive  and "-1" for the minus strand(3' to 5')/Crick/Antisense/negative.
	// pattern -?([0]|([1-9][0-9]*))
	Strand int64 `json:"strand"`
	// Extensions for windowStart
	WindowStart_ext *Element `json:"_windowStart"`
	// Extensions for windowEnd
	WindowEnd_ext *Element `json:"_windowEnd"`
	// Structural unit composed of a nucleic acid molecule which controls its own replication through the interaction of specific proteins at one or more origins of replication ([SO:0000340](http://www.sequenceontology.org/browser/current_svn/term/SO:0000340)).
	Chromosome *CodeableConcept `json:"chromosome"`
}

// CarePlan_Activity is Describes the intention of how one or more practitioners intend to deliver care for a particular patient, group or community for a period of time, possibly limited to care for a specific condition or set of conditions.
type CarePlan_Activity struct {
	_ *BackboneElement
	// Identifies the outcome at the point when the status of the activity is assessed.  For example, the outcome of an education activity could be patient understands (or not).
	OutcomeCodeableConcept []*CodeableConcept `json:"outcomeCodeableConcept"`
	// Details of the outcome or action resulting from the activity.  The reference to an "event" resource, such as Procedure or Encounter or Observation, is the result/outcome of the activity itself.  The activity can be conveyed using CarePlan.activity.detail OR using the CarePlan.activity.reference (a reference to a "request" resource).
	OutcomeReference []*Reference `json:"outcomeReference"`
	// Notes about the adherence/status/progress of the activity.
	Progress []*Annotation `json:"progress"`
	// The details of the proposed activity represented in a specific resource.
	Reference *Reference `json:"reference"`
	// A simple summary of a planned activity suitable for a general care plan system (e.g. form driven) that doesn't know about specific resources such as procedure etc.
	Detail *CarePlan_Detail `json:"detail"`
}

// ClinicalImpression_Finding is A record of a clinical assessment performed to determine what problem(s) may affect the patient and before planning the treatments or management strategies that are best to manage a patient's condition. Assessments are often 1:1 with a clinical consultation / encounter,  but this varies greatly depending on the clinical workflow. This resource is called "ClinicalImpression" rather than "ClinicalAssessment" to avoid confusion with the recording of assessment tools such as Apgar score.
type ClinicalImpression_Finding struct {
	_ *BackboneElement
	// Specific text, code or reference for finding or diagnosis, which may include ruled-out or resolved conditions.
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept"`
	// Specific text, code or reference for finding or diagnosis, which may include ruled-out or resolved conditions.
	ItemReference *Reference `json:"itemReference"`
	// Which investigations support finding or diagnosis.
	Basis string `json:"basis"`
	// Extensions for basis
	Basis_ext *Element `json:"_basis"`
}

// Encounter_Diagnosis is An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type Encounter_Diagnosis struct {
	_ *BackboneElement
	// Role that this diagnosis has within the encounter (e.g. admission, billing, discharge ...).
	Role *CodeableConcept `json:"role"`
	// Ranking of the diagnosis (for each role type).
	// pattern [1-9][0-9]*
	Rank uint64 `json:"rank"`
	// Extensions for rank
	Rank_ext *Element `json:"_rank"`
	// Reason the encounter takes place, as specified using information from another resource. For admissions, this is the admission diagnosis. The indication will typically be a Condition (with other resources referenced in the evidence.detail), or a Procedure.
	Condition *Reference `json:"condition,omitempty"`
}

// ExpansionProfile_Designation1 is Resource to define constraints on the Expansion of a FHIR ValueSet.
type ExpansionProfile_Designation1 struct {
	_ *BackboneElement
	// The language this designation is defined for.
	// pattern [^\s]+([\s]?[^\s]+)*
	Language string `json:"language"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
	// Which kinds of designation to include in the expansion.
	Use *Coding `json:"use"`
}

// PlanDefinition_Goal is This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinition_Goal struct {
	_ *BackboneElement
	// Indicates what should be done and within what timeframe.
	Target []*PlanDefinition_Target `json:"target"`
	// Indicates a category the goal falls within.
	Category *CodeableConcept `json:"category"`
	// Human-readable and/or coded description of a specific desired objective of care, such as "control blood pressure" or "negotiate an obstacle course" or "dance with child at wedding".
	Description *CodeableConcept `json:"description,omitempty"`
	// Identifies the expected level of importance associated with reaching/sustaining the defined goal.
	Priority *CodeableConcept `json:"priority"`
	// The event after which the goal should begin being pursued.
	Start *CodeableConcept `json:"start"`
	// Identifies problems, conditions, issues, or concerns the goal is intended to address.
	Addresses []*CodeableConcept `json:"addresses"`
	// Didactic or other informational resources associated with the goal that provide further supporting information about the goal. Information resources can include inline text commentary and links to web resources.
	Documentation []*RelatedArtifact `json:"documentation"`
}

// Specimen_Container is A sample to be used for analysis.
type Specimen_Container struct {
	_ *BackboneElement
	// The capacity (volume or other measure) the container may contain.
	Capacity *Quantity `json:"capacity"`
	// The quantity of specimen in the container; may be volume, dimensions, or other appropriate measurements, depending on the specimen type.
	SpecimenQuantity *Quantity `json:"specimenQuantity"`
	// Introduced substance to preserve, maintain or enhance the specimen. Examples: Formalin, Citrate, EDTA.
	AdditiveCodeableConcept *CodeableConcept `json:"additiveCodeableConcept"`
	// Introduced substance to preserve, maintain or enhance the specimen. Examples: Formalin, Citrate, EDTA.
	AdditiveReference *Reference `json:"additiveReference"`
	// Id for container. There may be multiple; a manufacturer's bar code, lab assigned identifier, etc. The container ID may differ from the specimen id in some circumstances.
	Identifier []*Identifier `json:"identifier"`
	// Textual description of the container.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The type of container associated with the specimen (e.g. slide, aliquot, etc.).
	Type *CodeableConcept `json:"type"`
}

// Period is A time period defined by a start and end date and optionally time.
type Period struct {
	_ *Element
	// The start of the period. The boundary is inclusive.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Start time.Time `json:"start"`
	// Extensions for start
	Start_ext *Element `json:"_start"`
	// The end of the period. If the end of the period is missing, it means that the period is ongoing. The start may be in the past, and the end date in the future, which means that period is expected/planned to end at that time.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	End time.Time `json:"end"`
	// Extensions for end
	End_ext *Element `json:"_end"`
}

// Claim_SubDetail is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_SubDetail struct {
	_ *BackboneElement
	// A service line number.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// Health Care Service Type Codes  to identify the classification of service or benefits.
	Category *CodeableConcept `json:"category"`
	// Item typification or modifiers codes, eg for Oral whether the treatment is cosmetic or associated with TMJ, or for medical whether the treatment was outside the clinic or out of office hours.
	Modifier []*CodeableConcept `json:"modifier"`
	// For programs which require reson codes for the inclusion, covering, of this billed item under the program or sub-program.
	ProgramCode []*CodeableConcept `json:"programCode"`
	// The fee for an addittional service or product or charge.
	UnitPrice *Money `json:"unitPrice"`
	// The quantity times the unit price for an addittional service or product or charge. For example, the formula: unit Quantity * unit Price (Cost per Point) * factor Number  * points = net Amount. Quantity, factor and points are assumed to be 1 if not supplied.
	Net *Money `json:"net"`
	// List of Unique Device Identifiers associated with this line item.
	Udi []*Reference `json:"udi"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// The type of reveneu or cost center providing the product and/or service.
	Revenue *CodeableConcept `json:"revenue"`
	// A code to indicate the Professional Service or Product supplied (eg. CTP, HCPCS,USCLS,ICD10, NCPDP,DIN,ACHI,CCI).
	Service *CodeableConcept `json:"service"`
	// The number of repetitions of a service or product.
	Quantity *Quantity `json:"quantity"`
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Factor float64 `json:"factor"`
	// Extensions for factor
	Factor_ext *Element `json:"_factor"`
}

// CodeSystem_Designation is A code system resource specifies a set of codes drawn from one or more code systems.
type CodeSystem_Designation struct {
	_ *BackboneElement
	// A code that details how this designation would be used.
	Use *Coding `json:"use"`
	// The text value for this designation.
	Value string `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// The language this designation is defined for.
	// pattern [^\s]+([\s]?[^\s]+)*
	Language string `json:"language"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
}

// ImagingManifest_Instance is A text description of the DICOM SOP instances selected in the ImagingManifest; or the reason for, or significance of, the selection.
type ImagingManifest_Instance struct {
	_ *BackboneElement
	// SOP class UID of the selected instance.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	SopClass string `json:"sopClass"`
	// Extensions for sopClass
	SopClass_ext *Element `json:"_sopClass"`
	// SOP Instance UID of the selected instance.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	Uid string `json:"uid"`
	// Extensions for uid
	Uid_ext *Element `json:"_uid"`
}

// ImagingStudy is Representation of the content produced in a DICOM imaging study. A study comprises a set of series, each of which includes a set of Service-Object Pair Instances (SOP Instances - images or other data) acquired or produced in a common context.  A series is of only one modality (e.g. X-ray, CT, MR, ultrasound), but a study may have multiple series of different modalities.
type ImagingStudy struct {
	_ *DomainResource
	// Formal identifier for the study.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	Uid string `json:"uid"`
	// The encounter or episode at which the request is initiated.
	Context *Reference `json:"context"`
	// A reference to the performed Procedure.
	ProcedureReference []*Reference `json:"procedureReference"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A list of the diagnostic requests that resulted in this imaging study being performed.
	BasedOn []*Reference `json:"basedOn"`
	// Extensions for numberOfSeries
	NumberOfSeries_ext *Element `json:"_numberOfSeries"`
	// Number of SOP Instances in Study. This value given may be larger than the number of instance elements this resource contains due to resource availability, security, or other factors. This element should be present if any instance elements are present.
	// pattern [0]|([1-9][0-9]*)
	NumberOfInstances uint64 `json:"numberOfInstances"`
	// Each study has one or more series of images or other content.
	Series []*ImagingStudy_Series `json:"series"`
	// This is a ImagingStudy resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for availability
	Availability_ext *Element `json:"_availability"`
	// The patient imaged in the study.
	Patient *Reference `json:"patient,omitempty"`
	// Extensions for started
	Started_ext *Element `json:"_started"`
	// Who read the study and interpreted the images or other content.
	Interpreter []*Reference `json:"interpreter"`
	// The network service providing access (e.g., query, view, or retrieval) for the study. See implementation notes for information about using DICOM endpoints. A study-level endpoint applies to each series in the study, unless overridden by a series-level endpoint with the same Endpoint.type.
	Endpoint []*Reference `json:"endpoint"`
	// Number of Series in the Study. This value given may be larger than the number of series elements this Resource contains due to resource availability, security, or other factors. This element should be present if any series elements are present.
	// pattern [0]|([1-9][0-9]*)
	NumberOfSeries uint64 `json:"numberOfSeries"`
	// Extensions for numberOfInstances
	NumberOfInstances_ext *Element `json:"_numberOfInstances"`
	// Extensions for uid
	Uid_ext *Element `json:"_uid"`
	// Availability of study (online, offline, or nearline).
	Availability string `json:"availability"`
	// A list of all the Series.ImageModality values that are actual acquisition modalities, i.e. those in the DICOM Context Group 29 (value set OID 1.2.840.10008.6.1.19).
	ModalityList []*Coding `json:"modalityList"`
	// The requesting/referring physician.
	Referrer *Reference `json:"referrer"`
	// Description of clinical condition indicating why the ImagingStudy was requested.
	Reason *CodeableConcept `json:"reason"`
	// Institution-generated description or classification of the Study performed.
	Description string `json:"description"`
	// Accession Number is an identifier related to some aspect of imaging workflow and data management. Usage may vary across different institutions.  See for instance [IHE Radiology Technical Framework Volume 1 Appendix A](http://www.ihe.net/uploadedFiles/Documents/Radiology/IHE_RAD_TF_Rev13.0_Vol1_FT_2014-07-30.pdf).
	Accession *Identifier `json:"accession"`
	// Other identifiers for the study.
	Identifier []*Identifier `json:"identifier"`
	// Date and time the study started.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Started time.Time `json:"started"`
	// The code for the performed procedure type.
	ProcedureCode []*CodeableConcept `json:"procedureCode"`
}

func NewImagingStudy() *ImagingStudy { return &ImagingStudy{ResourceType: "ImagingStudy"} }

// ResearchStudy_Arm is A process where a researcher or organization plans and then executes a series of steps intended to increase the field of healthcare-related knowledge.  This includes studies of safety, efficacy, comparative effectiveness and other information about medications, devices, therapies and other interventional and investigative techniques.  A ResearchStudy involves the gathering of information about human or animal subjects.
type ResearchStudy_Arm struct {
	_ *BackboneElement
	// Unique, human-readable label for this arm of the study.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Categorization of study arm, e.g. experimental, active comparator, placebo comparater.
	Code *CodeableConcept `json:"code"`
	// A succinct description of the path through the study that would be followed by a subject adhering to this arm.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
}

// Task_Output is A task to be performed.
type Task_Output struct {
	_ *BackboneElement
	// The value of the Output parameter as a basic type.
	ValueCoding *Coding `json:"valueCoding"`
	// Extensions for valueBase64Binary
	ValueBase64Binary_ext *Element `json:"_valueBase64Binary"`
	// The value of the Output parameter as a basic type.
	ValueSimpleQuantity *Quantity `json:"valueSimpleQuantity"`
	// Extensions for valueUuid
	ValueUuid_ext *Element `json:"_valueUuid"`
	// The value of the Output parameter as a basic type.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ValueDateTime time.Time `json:"valueDateTime"`
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// The value of the Output parameter as a basic type.
	// pattern [A-Za-z0-9\-\.]{1,64}
	ValueId string `json:"valueId"`
	// The value of the Output parameter as a basic type.
	ValueMarkdown string `json:"valueMarkdown"`
	// The value of the Output parameter as a basic type.
	ValueBackboneElement *BackboneElement `json:"valueBackboneElement"`
	// The value of the Output parameter as a basic type.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// The value of the Output parameter as a basic type.
	ValuePeriod *Period `json:"valuePeriod"`
	// The value of the Output parameter as a basic type.
	ValueRatio *Ratio `json:"valueRatio"`
	// Extensions for valueDecimal
	ValueDecimal_ext *Element `json:"_valueDecimal"`
	// The value of the Output parameter as a basic type.
	ValueCount *Count `json:"valueCount"`
	// The value of the Output parameter as a basic type.
	ValueAge *Age `json:"valueAge"`
	// The value of the Output parameter as a basic type.
	// pattern urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	ValueUuid string `json:"valueUuid"`
	// The value of the Output parameter as a basic type.
	ValueUri string `json:"valueUri"`
	// Extensions for valueTime
	ValueTime_ext *Element `json:"_valueTime"`
	// The value of the Output parameter as a basic type.
	ValueAnnotation *Annotation `json:"valueAnnotation"`
	// The value of the Output parameter as a basic type.
	ValueContactPoint *ContactPoint `json:"valueContactPoint"`
	// The value of the Output parameter as a basic type.
	ValueContactDetail *ContactDetail `json:"valueContactDetail"`
	// The value of the Output parameter as a basic type.
	ValueBase64Binary string `json:"valueBase64Binary"`
	// Extensions for valueOid
	ValueOid_ext *Element `json:"_valueOid"`
	// The value of the Output parameter as a basic type.
	ValueAttachment *Attachment `json:"valueAttachment"`
	// The value of the Output parameter as a basic type.
	ValueRange *Range `json:"valueRange"`
	// The value of the Output parameter as a basic type.
	ValueMeta *Meta `json:"valueMeta"`
	// The value of the Output parameter as a basic type.
	ValueDosage *Dosage `json:"valueDosage"`
	// Extensions for valueCode
	ValueCode_ext *Element `json:"_valueCode"`
	// The value of the Output parameter as a basic type.
	ValueString string `json:"valueString"`
	// Extensions for valueDate
	ValueDate_ext *Element `json:"_valueDate"`
	// Extensions for valueMarkdown
	ValueMarkdown_ext *Element `json:"_valueMarkdown"`
	// The value of the Output parameter as a basic type.
	ValueSampledData *SampledData `json:"valueSampledData"`
	// The value of the Output parameter as a basic type.
	ValueContributor *Contributor `json:"valueContributor"`
	// The value of the Output parameter as a basic type.
	ValueInstant string `json:"valueInstant"`
	// Extensions for valueInteger
	ValueInteger_ext *Element `json:"_valueInteger"`
	// The value of the Output parameter as a basic type.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	ValueDecimal float64 `json:"valueDecimal"`
	// Extensions for valueInstant
	ValueInstant_ext *Element `json:"_valueInstant"`
	// Extensions for valueDateTime
	ValueDateTime_ext *Element `json:"_valueDateTime"`
	// The value of the Output parameter as a basic type.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	ValueTime time.Time `json:"valueTime"`
	// The value of the Output parameter as a basic type.
	// pattern [^\s]+([\s]?[^\s]+)*
	ValueCode string `json:"valueCode"`
	// Extensions for valuePositiveInt
	ValuePositiveInt_ext *Element `json:"_valuePositiveInt"`
	// The value of the Output parameter as a basic type.
	// pattern -?([0]|([1-9][0-9]*))
	ValueInteger int64 `json:"valueInteger"`
	// The value of the Output parameter as a basic type.
	ValueMoney *Money `json:"valueMoney"`
	// The value of the Output parameter as a basic type.
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement"`
	// The value of the Output parameter as a basic type.
	ValueDuration *Duration `json:"valueDuration"`
	// Extensions for valueUnsignedInt
	ValueUnsignedInt_ext *Element `json:"_valueUnsignedInt"`
	// The value of the Output parameter as a basic type.
	ValueDistance *Distance `json:"valueDistance"`
	// The value of the Output parameter as a basic type.
	ValueBoolean bool `json:"valueBoolean"`
	// The value of the Output parameter as a basic type.
	// pattern [1-9][0-9]*
	ValuePositiveInt uint64 `json:"valuePositiveInt"`
	// The value of the Output parameter as a basic type.
	ValueElement *Element `json:"valueElement"`
	// The value of the Output parameter as a basic type.
	ValueIdentifier *Identifier `json:"valueIdentifier"`
	// The value of the Output parameter as a basic type.
	ValueAddress *Address `json:"valueAddress"`
	// The value of the Output parameter as a basic type.
	ValueTiming *Timing `json:"valueTiming"`
	// The value of the Output parameter as a basic type.
	ValueElementDefinition *ElementDefinition `json:"valueElementDefinition"`
	// The value of the Output parameter as a basic type.
	// pattern [0]|([1-9][0-9]*)
	ValueUnsignedInt uint64 `json:"valueUnsignedInt"`
	// Extensions for valueUri
	ValueUri_ext *Element `json:"_valueUri"`
	// The value of the Output parameter as a basic type.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	ValueOid string `json:"valueOid"`
	// The value of the Output parameter as a basic type.
	ValueNarrative *Narrative `json:"valueNarrative"`
	// The value of the Output parameter as a basic type.
	ValueReference *Reference `json:"valueReference"`
	// The value of the Output parameter as a basic type.
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact"`
	// The value of the Output parameter as a basic type.
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// The value of the Output parameter as a basic type.
	ValueHumanName *HumanName `json:"valueHumanName"`
	// The value of the Output parameter as a basic type.
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition"`
	// The value of the Output parameter as a basic type.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	// The value of the Output parameter as a basic type.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ValueDate time.Time `json:"valueDate"`
	// Extensions for valueId
	ValueId_ext *Element `json:"_valueId"`
	// The value of the Output parameter as a basic type.
	ValueExtension *Extension `json:"valueExtension"`
	// The value of the Output parameter as a basic type.
	ValueSignature *Signature `json:"valueSignature"`
	// The value of the Output parameter as a basic type.
	ValueUsageContext *UsageContext `json:"valueUsageContext"`
	// The name of the Output parameter.
	Type *CodeableConcept `json:"type,omitempty"`
}

// Annotation is A  text note which also  contains information about who made the statement and when.
type Annotation struct {
	_ *Element
	// Extensions for time
	Time_ext *Element `json:"_time"`
	// The text of the annotation.
	Text string `json:"text"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// The individual responsible for making the annotation.
	AuthorReference *Reference `json:"authorReference"`
	// The individual responsible for making the annotation.
	AuthorString string `json:"authorString"`
	// Extensions for authorString
	AuthorString_ext *Element `json:"_authorString"`
	// Indicates when this particular annotation was made.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Time time.Time `json:"time"`
}

// Bundle_Search is A container for a collection of resources.
type Bundle_Search struct {
	_ *BackboneElement
	// Why this entry is in the result set - whether it's included as a match or because of an _include requirement.
	Mode string `json:"mode"`
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// When searching, the server's search ranking score for the entry.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Score float64 `json:"score"`
	// Extensions for score
	Score_ext *Element `json:"_score"`
}

// ClaimResponse_Detail is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse_Detail struct {
	_ *BackboneElement
	// The third tier service adjudications for submitted services.
	SubDetail []*ClaimResponse_SubDetail `json:"subDetail"`
	// A service line number.
	// pattern [1-9][0-9]*
	SequenceLinkId uint64 `json:"sequenceLinkId"`
	// Extensions for sequenceLinkId
	SequenceLinkId_ext *Element `json:"_sequenceLinkId"`
	// A list of note references to the notes provided below.
	NoteNumber []uint64 `json:"noteNumber"`
	// Extensions for noteNumber
	NoteNumber_ext []*Element `json:"_noteNumber"`
	// The adjudications results.
	Adjudication []*ClaimResponse_Adjudication `json:"adjudication"`
}

// ClaimResponse_SubDetail is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse_SubDetail struct {
	_ *BackboneElement
	// Extensions for noteNumber
	NoteNumber_ext []*Element `json:"_noteNumber"`
	// The adjudications results.
	Adjudication []*ClaimResponse_Adjudication `json:"adjudication"`
	// A service line number.
	// pattern [1-9][0-9]*
	SequenceLinkId uint64 `json:"sequenceLinkId"`
	// Extensions for sequenceLinkId
	SequenceLinkId_ext *Element `json:"_sequenceLinkId"`
	// A list of note references to the notes provided below.
	NoteNumber []uint64 `json:"noteNumber"`
}

// EnrollmentRequest is This resource provides the insurance enrollment details to the insurer regarding a specified coverage.
type EnrollmentRequest struct {
	_ *DomainResource
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// The organization which is responsible for the services rendered to the patient.
	Organization *Reference `json:"organization"`
	// Patient Resource.
	Subject *Reference `json:"subject"`
	// Reference to the program or plan identification, underwriter or payor.
	Coverage *Reference `json:"coverage"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The date when this resource was created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// The Insurer who is target  of the request.
	Insurer *Reference `json:"insurer"`
	// The practitioner who is responsible for the services rendered to the patient.
	Provider *Reference `json:"provider"`
	// This is a EnrollmentRequest resource
	ResourceType string `json:"resourceType,omitempty"`
	// The Response business identifier.
	Identifier []*Identifier `json:"identifier"`
}

func NewEnrollmentRequest() *EnrollmentRequest {
	return &EnrollmentRequest{ResourceType: "EnrollmentRequest"}
}

// Immunization_VaccinationProtocol is Describes the event of a patient being administered a vaccination or a record of a vaccination as reported by a patient, a clinician or another party and may include vaccine reaction information and what vaccination protocol was followed.
type Immunization_VaccinationProtocol struct {
	_ *BackboneElement
	// Extensions for doseSequence
	DoseSequence_ext *Element `json:"_doseSequence"`
	// Indicates the authority who published the protocol.  E.g. ACIP.
	Authority *Reference `json:"authority"`
	// Extensions for series
	Series_ext *Element `json:"_series"`
	// The recommended number of doses to achieve immunity.
	// pattern [1-9][0-9]*
	SeriesDoses uint64 `json:"seriesDoses"`
	// Nominal position in a series.
	// pattern [1-9][0-9]*
	DoseSequence uint64 `json:"doseSequence"`
	// Contains the description about the protocol under which the vaccine was administered.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series string `json:"series"`
	// Extensions for seriesDoses
	SeriesDoses_ext *Element `json:"_seriesDoses"`
	// The targeted disease.
	TargetDisease []*CodeableConcept `json:"targetDisease,omitempty"`
	// Indicates if the immunization event should "count" against  the protocol.
	DoseStatus *CodeableConcept `json:"doseStatus,omitempty"`
	// Provides an explanation as to why an immunization event should or should not count against the protocol.
	DoseStatusReason *CodeableConcept `json:"doseStatusReason"`
}

// Questionnaire_Item is A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
type Questionnaire_Item struct {
	_ *BackboneElement
	// A reference to a value set containing a list of codes representing permitted answers for a "choice" or "open-choice" question.
	Options *Reference `json:"options"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	InitialDecimal float64 `json:"initialDecimal"`
	// Extensions for definition
	Definition_ext *Element `json:"_definition"`
	// A terminology code that corresponds to this group or question (e.g. a code from LOINC, which defines many questions and answers).
	Code []*Coding `json:"code"`
	// The name of a section, the text of a question or text content for a display item.
	Text string `json:"text"`
	// An indication, when true, that the value cannot be changed by a human respondent to the Questionnaire.
	ReadOnly bool `json:"readOnly"`
	// Extensions for maxLength
	MaxLength_ext *Element `json:"_maxLength"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	InitialBoolean bool `json:"initialBoolean"`
	// Extensions for initialInteger
	InitialInteger_ext *Element `json:"_initialInteger"`
	// Extensions for initialDateTime
	InitialDateTime_ext *Element `json:"_initialDateTime"`
	// Extensions for initialString
	InitialString_ext *Element `json:"_initialString"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	InitialUri string `json:"initialUri"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	InitialCoding *Coding `json:"initialCoding"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	InitialQuantity *Quantity `json:"initialQuantity"`
	// An identifier that is unique within the Questionnaire allowing linkage to the equivalent item in a QuestionnaireResponse resource.
	LinkId string `json:"linkId"`
	// An indication, if true, that the item must be present in a "completed" QuestionnaireResponse.  If false, the item may be skipped when answering the questionnaire.
	Required bool `json:"required"`
	// Extensions for required
	Required_ext *Element `json:"_required"`
	// Extensions for repeats
	Repeats_ext *Element `json:"_repeats"`
	// A short label for a particular group, question or set of display text within the questionnaire used for reference by the individual completing the questionnaire.
	Prefix string `json:"prefix"`
	// The maximum number of characters that are permitted in the answer to be considered a "valid" QuestionnaireResponse.
	// pattern -?([0]|([1-9][0-9]*))
	MaxLength int64 `json:"maxLength"`
	// Extensions for initialBoolean
	InitialBoolean_ext *Element `json:"_initialBoolean"`
	// Extensions for initialDate
	InitialDate_ext *Element `json:"_initialDate"`
	// The type of questionnaire item this is - whether text for display, a grouping of other items or a particular type of data to be captured (string, integer, coded choice, etc.).
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	InitialString string `json:"initialString"`
	// Extensions for readOnly
	ReadOnly_ext *Element `json:"_readOnly"`
	// One of the permitted answers for a "choice" or "open-choice" question.
	Option []*Questionnaire_Option `json:"option"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	// pattern -?([0]|([1-9][0-9]*))
	InitialInteger int64 `json:"initialInteger"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	InitialDate time.Time `json:"initialDate"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	InitialDateTime time.Time `json:"initialDateTime"`
	// Extensions for initialTime
	InitialTime_ext *Element `json:"_initialTime"`
	// A reference to an [[[ElementDefinition]]] that provides the details for the item. If a definition is provided, then the following element values can be inferred from the definition:
	//
	// * code (ElementDefinition.code)
	// * type (ElementDefinition.type)
	// * required (ElementDefinition.min)
	// * repeats (ElementDefinition.max)
	// * maxLength (ElementDefinition.maxLength)
	// * options (ElementDefinition.binding)
	//
	// Any information provided in these elements on a Questionnaire Item overrides the information from the definition.
	Definition string `json:"definition"`
	// An indication, if true, that the item may occur multiple times in the response, collecting multiple answers answers for questions or multiple sets of answers for groups.
	Repeats bool `json:"repeats"`
	// Extensions for initialDecimal
	InitialDecimal_ext *Element `json:"_initialDecimal"`
	// Extensions for linkId
	LinkId_ext *Element `json:"_linkId"`
	// Extensions for prefix
	Prefix_ext *Element `json:"_prefix"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// A constraint indicating that this item should only be enabled (displayed/allow answers to be captured) when the specified condition is true.
	EnableWhen []*Questionnaire_EnableWhen `json:"enableWhen"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	InitialTime time.Time `json:"initialTime"`
	// Extensions for initialUri
	InitialUri_ext *Element `json:"_initialUri"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	InitialAttachment *Attachment `json:"initialAttachment"`
	// The value that should be defaulted when initially rendering the questionnaire for user input.
	InitialReference *Reference `json:"initialReference"`
	// Text, questions and other groups to be nested beneath a question or group.
	Item []*Questionnaire_Item `json:"item"`
}

// Specimen_Collection is A sample to be used for analysis.
type Specimen_Collection struct {
	_ *BackboneElement
	// Time when specimen was collected from subject - the physiologically relevant time.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	CollectedDateTime time.Time `json:"collectedDateTime"`
	// Extensions for collectedDateTime
	CollectedDateTime_ext *Element `json:"_collectedDateTime"`
	// Time when specimen was collected from subject - the physiologically relevant time.
	CollectedPeriod *Period `json:"collectedPeriod"`
	// The quantity of specimen collected; for instance the volume of a blood sample, or the physical measurement of an anatomic pathology sample.
	Quantity *Quantity `json:"quantity"`
	// A coded value specifying the technique that is used to perform the procedure.
	Method *CodeableConcept `json:"method"`
	// Anatomical location from which the specimen was collected (if subject is a patient). This is the target site.  This element is not used for environmental specimens.
	BodySite *CodeableConcept `json:"bodySite"`
	// Person who collected the specimen.
	Collector *Reference `json:"collector"`
}

// BackboneElement is Base definition for all elements that are defined inside a resource - but not those in a data type.
type BackboneElement struct {
	_ *Element
	// May be used to represent additional information that is not part of the basic definition of the element, and that modifies the understanding of the element that contains it. Usually modifier elements provide negation or qualification. In order to make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	ModifierExtension []*Extension `json:"modifierExtension"`
}

// Questionnaire_EnableWhen is A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
type Questionnaire_EnableWhen struct {
	_ *BackboneElement
	// An answer that the referenced question must match in order for the item to be enabled.
	AnswerReference *Reference `json:"answerReference"`
	// Extensions for question
	Question_ext *Element `json:"_question"`
	// An answer that the referenced question must match in order for the item to be enabled.
	AnswerBoolean bool `json:"answerBoolean"`
	// An answer that the referenced question must match in order for the item to be enabled.
	// pattern -?([0]|([1-9][0-9]*))
	AnswerInteger int64 `json:"answerInteger"`
	// Extensions for answerInteger
	AnswerInteger_ext *Element `json:"_answerInteger"`
	// An answer that the referenced question must match in order for the item to be enabled.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	AnswerDate time.Time `json:"answerDate"`
	// An answer that the referenced question must match in order for the item to be enabled.
	AnswerCoding *Coding `json:"answerCoding"`
	// The linkId for the question whose answer (or lack of answer) governs whether this item is enabled.
	Question string `json:"question"`
	// Extensions for hasAnswer
	HasAnswer_ext *Element `json:"_hasAnswer"`
	// Extensions for answerBoolean
	AnswerBoolean_ext *Element `json:"_answerBoolean"`
	// An answer that the referenced question must match in order for the item to be enabled.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	AnswerTime time.Time `json:"answerTime"`
	// An answer that the referenced question must match in order for the item to be enabled.
	AnswerString string `json:"answerString"`
	// An answer that the referenced question must match in order for the item to be enabled.
	AnswerAttachment *Attachment `json:"answerAttachment"`
	// An indication that this item should be enabled only if the specified question is answered (hasAnswer=true) or not answered (hasAnswer=false).
	HasAnswer bool `json:"hasAnswer"`
	// An answer that the referenced question must match in order for the item to be enabled.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AnswerDateTime time.Time `json:"answerDateTime"`
	// Extensions for answerTime
	AnswerTime_ext *Element `json:"_answerTime"`
	// An answer that the referenced question must match in order for the item to be enabled.
	AnswerUri string `json:"answerUri"`
	// An answer that the referenced question must match in order for the item to be enabled.
	AnswerQuantity *Quantity `json:"answerQuantity"`
	// An answer that the referenced question must match in order for the item to be enabled.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	AnswerDecimal float64 `json:"answerDecimal"`
	// Extensions for answerDecimal
	AnswerDecimal_ext *Element `json:"_answerDecimal"`
	// Extensions for answerDate
	AnswerDate_ext *Element `json:"_answerDate"`
	// Extensions for answerDateTime
	AnswerDateTime_ext *Element `json:"_answerDateTime"`
	// Extensions for answerString
	AnswerString_ext *Element `json:"_answerString"`
	// Extensions for answerUri
	AnswerUri_ext *Element `json:"_answerUri"`
}

// ActivityDefinition_DynamicValue is This resource allows for the definition of some activity to be performed, independent of a particular patient, practitioner, or other performance context.
type ActivityDefinition_DynamicValue struct {
	_ *BackboneElement
	// Extensions for expression
	Expression_ext *Element `json:"_expression"`
	// A brief, natural language description of the intended semantics of the dynamic value.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The path to the element to be customized. This is the path on the resource that will hold the result of the calculation defined by the expression.
	Path string `json:"path"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// The media type of the language for the expression.
	Language string `json:"language"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
	// An expression specifying the value of the customized element.
	Expression string `json:"expression"`
}

// ExpansionProfile_FixedVersion is Resource to define constraints on the Expansion of a FHIR ValueSet.
type ExpansionProfile_FixedVersion struct {
	_ *BackboneElement
	// Extensions for system
	System_ext *Element `json:"_system"`
	// The version of the code system from which codes in the expansion should be included.
	Version string `json:"version"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// How to manage the intersection between a fixed version in a value set, and this fixed version of the system in the expansion profile.
	Mode string `json:"mode"`
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// The specific system for which to fix the version.
	System string `json:"system"`
}

// ContactDetail is Specifies contact information for a person or organization.
type ContactDetail struct {
	_ *Element
	// The name of an individual to contact.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The contact details for the individual (if a name was provided) or the organization.
	Telecom []*ContactPoint `json:"telecom"`
}

// CodeSystem is A code system resource specifies a set of codes drawn from one or more code systems.
type CodeSystem struct {
	_ *DomainResource
	// The total number of concepts defined by the code system. Where the code system has a compositional grammar, the count refers to the number of base (primitive) concepts.
	// pattern [0]|([1-9][0-9]*)
	Count uint64 `json:"count"`
	// The name of the individual or organization that published the code system.
	Publisher string `json:"publisher"`
	// Extensions for valueSet
	ValueSet_ext *Element `json:"_valueSet"`
	// True If code system defines a post-composition grammar.
	Compositional bool `json:"compositional"`
	// The status of this code system. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A copyright statement relating to the code system and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the code system.
	Copyright string `json:"copyright"`
	// This is a CodeSystem resource
	ResourceType string `json:"resourceType,omitempty"`
	// An absolute URI that is used to identify this code system when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this code system is (or will be) published. The URL SHOULD include the major version of the code system. For more information see [Technical and Business Versions](resource.html#versions). This is used in [Coding]{datatypes.html#Coding}.system.
	Url string `json:"url"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for caseSensitive
	CaseSensitive_ext *Element `json:"_caseSensitive"`
	// The meaning of the hierarchy of concepts.
	HierarchyMeaning string `json:"hierarchyMeaning"`
	// Extensions for content
	Content_ext *Element `json:"_content"`
	// The date  (and optionally time) when the code system was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the code system changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Explaination of why this code system is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// Extensions for versionNeeded
	VersionNeeded_ext *Element `json:"_versionNeeded"`
	// How much of the content of the code system - the concepts and codes it defines - are represented in this resource.
	Content string `json:"content"`
	// A filter that can be used in a value set compose statement when selecting concepts using a filter.
	Filter []*CodeSystem_Filter `json:"filter"`
	// A short, descriptive, user-friendly title for the code system.
	Title string `json:"title"`
	// A legal or geographic region in which the code system is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Extensions for hierarchyMeaning
	HierarchyMeaning_ext *Element `json:"_hierarchyMeaning"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// If code comparison is case sensitive when codes within this system are compared to each other.
	CaseSensitive bool `json:"caseSensitive"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate code system instances.
	UseContext []*UsageContext `json:"useContext"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// Canonical URL of value set that contains the entire code system.
	ValueSet string `json:"valueSet"`
	// Extensions for compositional
	Compositional_ext *Element `json:"_compositional"`
	// Extensions for count
	Count_ext *Element `json:"_count"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// A natural language name identifying the code system. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// A free text natural language description of the code system from a consumer's perspective.
	Description string `json:"description"`
	// Concepts that are in the code system. The concept definitions are inherently hierarchical, but the definitions must be consulted to determine what the meaning of the hierarchical relationships are.
	Concept []*CodeSystem_Concept `json:"concept"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// This flag is used to signify that the code system has not (or does not) maintain the definitions, and a version must be specified when referencing this code system.
	VersionNeeded bool `json:"versionNeeded"`
	// A property defines an additional slot through which additional information can be provided about a concept.
	Property []*CodeSystem_Property `json:"property"`
	// A formal identifier that is used to identify this code system when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier *Identifier `json:"identifier"`
	// The identifier that is used to identify this version of the code system when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the code system author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. This is used in [Coding]{datatypes.html#Coding}.version.
	Version string `json:"version"`
	// A boolean value to indicate that this code system is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
}

func NewCodeSystem() *CodeSystem { return &CodeSystem{ResourceType: "CodeSystem"} }

// EligibilityRequest is The EligibilityRequest provides patient and insurance coverage information to an insurer for them to respond, in the form of an EligibilityResponse, with information regarding whether the stated coverage is valid and in-force and optionally to provide the insurance details of the policy.
type EligibilityRequest struct {
	_ *DomainResource
	// Immediate (STAT), best effort (NORMAL), deferred (DEFER).
	Priority *CodeableConcept `json:"priority"`
	// The date when this resource was created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// Dental, Vision, Medical, Pharmacy, Rehab etc.
	BenefitCategory *CodeableConcept `json:"benefitCategory"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for servicedDate
	ServicedDate_ext *Element `json:"_servicedDate"`
	// The contract number of a business agreement which describes the terms and conditions.
	BusinessArrangement string `json:"businessArrangement"`
	// This is a EligibilityRequest resource
	ResourceType string `json:"resourceType,omitempty"`
	// Patient Resource.
	Patient *Reference `json:"patient"`
	// The date or dates when the enclosed suite of services were performed or completed.
	ServicedPeriod *Period `json:"servicedPeriod"`
	// Person who created the invoice/claim/pre-determination or pre-authorization.
	Enterer *Reference `json:"enterer"`
	// The practitioner who is responsible for the services rendered to the patient.
	Provider *Reference `json:"provider"`
	// Financial instrument by which payment information for health care.
	Coverage *Reference `json:"coverage"`
	// Extensions for businessArrangement
	BusinessArrangement_ext *Element `json:"_businessArrangement"`
	// The Response business identifier.
	Identifier []*Identifier `json:"identifier"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// The date or dates when the enclosed suite of services were performed or completed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ServicedDate time.Time `json:"servicedDate"`
	// The organization which is responsible for the services rendered to the patient.
	Organization *Reference `json:"organization"`
	// The Insurer who is target  of the request.
	Insurer *Reference `json:"insurer"`
	// Facility where the services were provided.
	Facility *Reference `json:"facility"`
	// Dental: basic, major, ortho; Vision exam, glasses, contacts; etc.
	BenefitSubCategory *CodeableConcept `json:"benefitSubCategory"`
}

func NewEligibilityRequest() *EligibilityRequest {
	return &EligibilityRequest{ResourceType: "EligibilityRequest"}
}

// Linkage_Item is Identifies two or more records (resource instances) that are referring to the same real-world "occurrence".
type Linkage_Item struct {
	_ *BackboneElement
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The resource instance being linked as part of the group.
	Resource *Reference `json:"resource,omitempty"`
	// Distinguishes which item is "source of truth" (if any) and which items are no longer considered to be current representations.
	Type string `json:"type"`
}

// MeasureReport_Group is The MeasureReport resource contains the results of evaluating a measure.
type MeasureReport_Group struct {
	_ *BackboneElement
	// The identifier of the population group as defined in the measure definition.
	Identifier *Identifier `json:"identifier,omitempty"`
	// The populations that make up the population group, one for each type of population appropriate for the measure.
	Population []*MeasureReport_Population `json:"population"`
	// The measure score for this population group, calculated as appropriate for the measure type and scoring method, and based on the contents of the populations defined in the group.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	MeasureScore float64 `json:"measureScore"`
	// Extensions for measureScore
	MeasureScore_ext *Element `json:"_measureScore"`
	// When a measure includes multiple stratifiers, there will be a stratifier group for each stratifier defined by the measure.
	Stratifier []*MeasureReport_Stratifier `json:"stratifier"`
}

// Schedule is A container for slots of time that may be available for booking appointments.
type Schedule struct {
	_ *DomainResource
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
	// A broad categorisation of the service that is to be performed during this appointment.
	ServiceCategory *CodeableConcept `json:"serviceCategory"`
	// The specific service that is to be performed during this appointment.
	ServiceType []*CodeableConcept `json:"serviceType"`
	// The specialty of a practitioner that would be required to perform the service requested in this appointment.
	Specialty []*CodeableConcept `json:"specialty"`
	// The period of time that the slots that are attached to this Schedule resource cover (even if none exist). These  cover the amount of time that an organization's planning horizon; the interval for which they are currently accepting appointments. This does not define a "template" for planning outside these dates.
	PlanningHorizon *Period `json:"planningHorizon"`
	// Comments on the availability to describe any extended information. Such as custom constraints on the slots that may be associated.
	Comment string `json:"comment"`
	// This is a Schedule resource
	ResourceType string `json:"resourceType,omitempty"`
	// External Ids for this item.
	Identifier []*Identifier `json:"identifier"`
	// Whether this schedule record is in active use, or should not be used (such as was entered in error).
	Active bool `json:"active"`
	// The resource this Schedule resource is providing availability information for. These are expected to usually be one of HealthcareService, Location, Practitioner, PractitionerRole, Device, Patient or RelatedPerson.
	Actor []*Reference `json:"actor,omitempty"`
}

func NewSchedule() *Schedule { return &Schedule{ResourceType: "Schedule"} }

// Sequence_Repository is Raw data describing a biological sequence.
type Sequence_Repository struct {
	_ *BackboneElement
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// URI of an external repository which contains further details about the genetics data.
	Name string `json:"name"`
	// Id of the variant in this external repository. The server will understand how to use this id to call for more info about datasets in external repository.
	DatasetId string `json:"datasetId"`
	// Id of the variantset in this external repository. The server will understand how to use this id to call for more info about variantsets in external repository.
	VariantsetId string `json:"variantsetId"`
	// Extensions for readsetId
	ReadsetId_ext *Element `json:"_readsetId"`
	// Click and see / RESTful API / Need login to see / RESTful API with authentication / Other ways to see resource.
	Type string `json:"type"`
	// URI of an external repository which contains further details about the genetics data.
	Url string `json:"url"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for datasetId
	DatasetId_ext *Element `json:"_datasetId"`
	// Extensions for variantsetId
	VariantsetId_ext *Element `json:"_variantsetId"`
	// Id of the read in this external repository.
	ReadsetId string `json:"readsetId"`
}

// TestScript_Teardown is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Teardown struct {
	_ *BackboneElement
	// The teardown action will only contain an operation.
	Action []*TestScript_Action2 `json:"action,omitempty"`
}

// AuditEvent_Network is A record of an event made for purposes of maintaining a security log. Typical uses include detection of intrusion attempts and monitoring for inappropriate usage.
type AuditEvent_Network struct {
	_ *BackboneElement
	// An identifier for the network access point of the user device for the audit event.
	Address string `json:"address"`
	// Extensions for address
	Address_ext *Element `json:"_address"`
	// An identifier for the type of network access point that originated the audit event.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
}

// PractitionerRole is A specific set of Roles/Locations/specialties/services that a practitioner may perform at an organization for a period of time.
type PractitionerRole struct {
	_ *DomainResource
	// A collection of times that the Service Site is available.
	AvailableTime []*PractitionerRole_AvailableTime `json:"availableTime"`
	// Technical endpoints providing access to services operated for the practitioner with this role.
	Endpoint []*Reference `json:"endpoint"`
	// Business Identifiers that are specific to a role/location.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
	// The period during which the person is authorized to act as a practitioner in these role(s) for the organization.
	Period *Period `json:"period"`
	// The organization where the Practitioner performs the roles associated.
	Organization *Reference `json:"organization"`
	// Roles which this practitioner is authorized to perform for the organization.
	Code []*CodeableConcept `json:"code"`
	// The list of healthcare services that this worker provides for this role's Organization/Location(s).
	HealthcareService []*Reference `json:"healthcareService"`
	// Practitioner that is able to provide the defined services for the organation.
	Practitioner *Reference `json:"practitioner"`
	// Contact details that are specific to the role/location/service.
	Telecom []*ContactPoint `json:"telecom"`
	// The HealthcareService is not available during this period of time due to the provided reason.
	NotAvailable []*PractitionerRole_NotAvailable `json:"notAvailable"`
	// Whether this practitioner's record is in active use.
	Active bool `json:"active"`
	// The location(s) at which this practitioner provides care.
	Location []*Reference `json:"location"`
	// Extensions for availabilityExceptions
	AvailabilityExceptions_ext *Element `json:"_availabilityExceptions"`
	// This is a PractitionerRole resource
	ResourceType string `json:"resourceType,omitempty"`
	// Specific specialty of the practitioner.
	Specialty []*CodeableConcept `json:"specialty"`
	// A description of site availability exceptions, e.g. public holiday availability. Succinctly describing all possible exceptions to normal site availability as details in the available Times and not available Times.
	AvailabilityExceptions string `json:"availabilityExceptions"`
}

func NewPractitionerRole() *PractitionerRole {
	return &PractitionerRole{ResourceType: "PractitionerRole"}
}

// Specimen_Processing is A sample to be used for analysis.
type Specimen_Processing struct {
	_ *BackboneElement
	// Extensions for timeDateTime
	TimeDateTime_ext *Element `json:"_timeDateTime"`
	// A record of the time or period when the specimen processing occurred.  For example the time of sample fixation or the period of time the sample was in formalin.
	TimePeriod *Period `json:"timePeriod"`
	// Textual description of procedure.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A coded value specifying the procedure used to process the specimen.
	Procedure *CodeableConcept `json:"procedure"`
	// Material used in the processing step.
	Additive []*Reference `json:"additive"`
	// A record of the time or period when the specimen processing occurred.  For example the time of sample fixation or the period of time the sample was in formalin.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	TimeDateTime time.Time `json:"timeDateTime"`
}

// Encounter_StatusHistory is An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type Encounter_StatusHistory struct {
	_ *BackboneElement
	// planned | arrived | triaged | in-progress | onleave | finished | cancelled +.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The time that the episode was in the specified status.
	Period *Period `json:"period,omitempty"`
}

// Substance is A homogeneous material with a definite composition.
type Substance struct {
	_ *DomainResource
	// A description of the substance - its appearance, handling requirements, and other usage notes.
	Description string `json:"description"`
	// This is a Substance resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A code that classifies the general type of substance.  This is used  for searching, sorting and display purposes.
	Category []*CodeableConcept `json:"category"`
	// A code (or set of codes) that identify this substance.
	Code *CodeableConcept `json:"code,omitempty"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Substance may be used to describe a kind of substance, or a specific package/container of the substance: an instance.
	Instance []*Substance_Instance `json:"instance"`
	// A substance can be composed of other substances.
	Ingredient []*Substance_Ingredient `json:"ingredient"`
	// Unique identifier for the substance.
	Identifier []*Identifier `json:"identifier"`
	// A code to indicate if the substance is actively used.
	Status string `json:"status"`
}

func NewSubstance() *Substance { return &Substance{ResourceType: "Substance"} }

// Narrative is A human-readable formatted text, including images.
type Narrative struct {
	_ *Element
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The actual narrative content, a stripped down version of XHTML.
	Div string `json:"div,omitempty"`
	// The status of the narrative - whether it's entirely generated (from just the defined data or the extensions too), or whether a human authored it and it may contain additional data.
	Status string `json:"status"`
}

// TriggerDefinition is A description of a triggering event.
type TriggerDefinition struct {
	_ *Element
	// Extensions for eventName
	EventName_ext *Element `json:"_eventName"`
	// The timing of the event (if this is a period trigger).
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	EventTimingDate time.Time `json:"eventTimingDate"`
	// Extensions for eventTimingDateTime
	EventTimingDateTime_ext *Element `json:"_eventTimingDateTime"`
	// The triggering data of the event (if this is a data trigger).
	EventData *DataRequirement `json:"eventData"`
	// The type of triggering event.
	Type string `json:"type"`
	// The name of the event (if this is a named-event trigger).
	EventName string `json:"eventName"`
	// The timing of the event (if this is a period trigger).
	EventTimingTiming *Timing `json:"eventTimingTiming"`
	// The timing of the event (if this is a period trigger).
	EventTimingReference *Reference `json:"eventTimingReference"`
	// Extensions for eventTimingDate
	EventTimingDate_ext *Element `json:"_eventTimingDate"`
	// The timing of the event (if this is a period trigger).
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	EventTimingDateTime time.Time `json:"eventTimingDateTime"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
}

// CommunicationRequest_Requester is A request to convey information; e.g. the CDS system proposes that an alert be sent to a responsible provider, the CDS system proposes that the public health agency be notified about a reportable condition.
type CommunicationRequest_Requester struct {
	_ *BackboneElement
	// The organization the device or practitioner was acting on behalf of.
	OnBehalfOf *Reference `json:"onBehalfOf"`
	// The device, practitioner, etc. who initiated the request.
	Agent *Reference `json:"agent,omitempty"`
}

// TestScript_Ruleset is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Ruleset struct {
	_ *BackboneElement
	// The referenced rule within the external ruleset template.
	Rule []*TestScript_Rule1 `json:"rule,omitempty"`
	// Reference to the resource (containing the contents of the ruleset needed for assertions).
	Resource *Reference `json:"resource,omitempty"`
}

// ParameterDefinition is The parameters to the module. This collection specifies both the input and output parameters. Input parameters are provided by the caller as part of the $evaluate operation. Output parameters are included in the GuidanceResponse.
type ParameterDefinition struct {
	_ *Element
	// The name of the parameter used to allow access to the value of the parameter in evaluation contexts.
	// pattern [^\s]+([\s]?[^\s]+)*
	Name string `json:"name"`
	// Whether the parameter is input or output for the module.
	// pattern [^\s]+([\s]?[^\s]+)*
	Use string `json:"use"`
	// Extensions for use
	Use_ext *Element `json:"_use"`
	// Extensions for max
	Max_ext *Element `json:"_max"`
	// The type of the parameter.
	// pattern [^\s]+([\s]?[^\s]+)*
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// If specified, this indicates a profile that the input data must conform to, or that the output data will conform to.
	Profile *Reference `json:"profile"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The minimum number of times this parameter SHALL appear in the request or response.
	// pattern -?([0]|([1-9][0-9]*))
	Min int64 `json:"min"`
	// Extensions for min
	Min_ext *Element `json:"_min"`
	// The maximum number of times this element is permitted to appear in the request or response.
	Max string `json:"max"`
	// A brief discussion of what the parameter is for and how it is used by the module.
	Documentation string `json:"documentation"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
}

// HealthcareService is The details of a healthcare service available at a location.
type HealthcareService struct {
	_ *DomainResource
	// Program Names that can be used to categorize the service.
	ProgramName []string `json:"programName"`
	// Identifies the broad category of service being performed or delivered.
	Category *CodeableConcept `json:"category"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for extraDetails
	ExtraDetails_ext *Element `json:"_extraDetails"`
	// Extensions for programName
	ProgramName_ext []*Element `json:"_programName"`
	// Extensions for appointmentRequired
	AppointmentRequired_ext *Element `json:"_appointmentRequired"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
	// The organization that provides this healthcare service.
	ProvidedBy *Reference `json:"providedBy"`
	// Does this service have specific eligibility requirements that need to be met in order to use the service?
	Eligibility *CodeableConcept `json:"eligibility"`
	// Collection of characteristics (attributes).
	Characteristic []*CodeableConcept `json:"characteristic"`
	// External identifiers for this item.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// Extra details about the service that can't be placed in the other fields.
	ExtraDetails string `json:"extraDetails"`
	// If there is a photo/symbol associated with this HealthcareService, it may be included here to facilitate quick identification of the service in a list.
	Photo *Attachment `json:"photo"`
	// The code(s) that detail the conditions under which the healthcare service is available/offered.
	ServiceProvisionCode []*CodeableConcept `json:"serviceProvisionCode"`
	// A description of site availability exceptions, e.g. public holiday availability. Succinctly describing all possible exceptions to normal site availability as details in the available Times and not available Times.
	AvailabilityExceptions string `json:"availabilityExceptions"`
	// Technical endpoints providing access to services operated for the specific healthcare services defined at this resource.
	Endpoint []*Reference `json:"endpoint"`
	// This is a HealthcareService resource
	ResourceType string `json:"resourceType,omitempty"`
	// Whether this healthcareservice record is in active use.
	Active bool `json:"active"`
	// The HealthcareService is not available during this period of time due to the provided reason.
	NotAvailable []*HealthcareService_NotAvailable `json:"notAvailable"`
	// List of contacts related to this specific healthcare service.
	Telecom []*ContactPoint `json:"telecom"`
	// Describes the eligibility conditions for the service.
	EligibilityNote string `json:"eligibilityNote"`
	// Indicates whether or not a prospective consumer will require an appointment for a particular service at a site to be provided by the Organization. Indicates if an appointment is required for access to this service.
	AppointmentRequired bool `json:"appointmentRequired"`
	// Extensions for availabilityExceptions
	AvailabilityExceptions_ext *Element `json:"_availabilityExceptions"`
	// Collection of specialties handled by the service site. This is more of a medical term.
	Specialty []*CodeableConcept `json:"specialty"`
	// Extensions for eligibilityNote
	EligibilityNote_ext *Element `json:"_eligibilityNote"`
	// The location(s) that this service is available to (not where the service is provided).
	CoverageArea []*Reference `json:"coverageArea"`
	// The specific type of service that may be delivered or performed.
	Type []*CodeableConcept `json:"type"`
	// Any additional description of the service and/or any specific issues not covered by the other attributes, which can be displayed as further detail under the serviceName.
	Comment string `json:"comment"`
	// Ways that the service accepts referrals, if this is not provided then it is implied that no referral is required.
	ReferralMethod []*CodeableConcept `json:"referralMethod"`
	// A collection of times that the Service Site is available.
	AvailableTime []*HealthcareService_AvailableTime `json:"availableTime"`
	// The location(s) where this healthcare service may be provided.
	Location []*Reference `json:"location"`
	// Further description of the service as it would be presented to a consumer while searching.
	Name string `json:"name"`
}

func NewHealthcareService() *HealthcareService {
	return &HealthcareService{ResourceType: "HealthcareService"}
}

// Contract_ValuedItem is A formal agreement between parties regarding the conduct of business, exchange of information or other matters.
type Contract_ValuedItem struct {
	_ *BackboneElement
	// An amount that expresses the weighting (based on difficulty, cost and/or resource intensiveness) associated with the Contract Valued Item delivered. The concept of Points allows for assignment of point values for a Contract Valued Item, such that a monetary amount can be assigned to each point.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Points float64 `json:"points"`
	// Specific type of Contract Valued Item that may be priced.
	EntityCodeableConcept *CodeableConcept `json:"entityCodeableConcept"`
	// Specific type of Contract Valued Item that may be priced.
	EntityReference *Reference `json:"entityReference"`
	// Indicates the time during which this Contract ValuedItem information is effective.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	EffectiveTime time.Time `json:"effectiveTime"`
	// A real number that represents a multiplier used in determining the overall value of the Contract Valued Item delivered. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Factor float64 `json:"factor"`
	// Extensions for factor
	Factor_ext *Element `json:"_factor"`
	// Extensions for points
	Points_ext *Element `json:"_points"`
	// Expresses the product of the Contract Valued Item unitQuantity and the unitPriceAmt. For example, the formula: unit Quantity * unit Price (Cost per Point) * factor Number  * points = net Amount. Quantity, factor and points are assumed to be 1 if not supplied.
	Net *Money `json:"net"`
	// Identifies a Contract Valued Item instance.
	Identifier *Identifier `json:"identifier"`
	// Extensions for effectiveTime
	EffectiveTime_ext *Element `json:"_effectiveTime"`
	// Specifies the units by which the Contract Valued Item is measured or counted, and quantifies the countable or measurable Contract Valued Item instances.
	Quantity *Quantity `json:"quantity"`
	// A Contract Valued Item unit valuation measure.
	UnitPrice *Money `json:"unitPrice"`
}

// TestScript_Test is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Test struct {
	_ *BackboneElement
	// The name of this test used for tracking/logging purposes by test engines.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// A short description of the test used by test engines for tracking and reporting purposes.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Action would contain either an operation or an assertion.
	Action []*TestScript_Action1 `json:"action,omitempty"`
}

// Communication is An occurrence of information being transmitted; e.g. an alert that was sent to a responsible provider, a public health agency was notified about a reportable condition.
type Communication struct {
	_ *DomainResource
	// The patient or group that was the focus of this communication.
	Subject *Reference `json:"subject"`
	// The entity (e.g. person, organization, clinical information system, or device) which was the target of the communication. If receipts need to be tracked by individual, a separate resource instance will need to be created for each recipient.  Multiple recipient communications are intended where either a receipt(s) is not tracked (e.g. a mass mail-out) or is captured in aggregate (all emails confirmed received by a particular time).
	Recipient []*Reference `json:"recipient"`
	// The time when this communication arrived at the destination.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Received time.Time `json:"received"`
	// Extensions for received
	Received_ext *Element `json:"_received"`
	// Extensions for sent
	Sent_ext *Element `json:"_sent"`
	// This is a Communication resource
	ResourceType string `json:"resourceType,omitempty"`
	// Identifiers associated with this Communication that are defined by business processes and/ or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []*Identifier `json:"identifier"`
	// A protocol, guideline, or other definition that was adhered to in whole or in part by this communication event.
	Definition []*Reference `json:"definition"`
	// Part of this action.
	PartOf []*Reference `json:"partOf"`
	// The status of the transmission.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// The type of message conveyed such as alert, notification, reminder, instruction, etc.
	Category []*CodeableConcept `json:"category"`
	// The encounter within which the communication was sent.
	Context *Reference `json:"context"`
	// Extensions for notDone
	NotDone_ext *Element `json:"_notDone"`
	// A channel that was used for this communication (e.g. email, fax).
	Medium []*CodeableConcept `json:"medium"`
	// The resources which were responsible for or related to producing this communication.
	Topic []*Reference `json:"topic"`
	// The time when this communication was sent.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Sent time.Time `json:"sent"`
	// The entity (e.g. person, organization, clinical information system, or device) which was the source of the communication.
	Sender *Reference `json:"sender"`
	// Additional notes or commentary about the communication by the sender, receiver or other interested parties.
	Note []*Annotation `json:"note"`
	// An order, proposal or plan fulfilled in whole or in part by this Communication.
	BasedOn []*Reference `json:"basedOn"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// If true, indicates that the described communication event did not actually occur.
	NotDone bool `json:"notDone"`
	// Describes why the communication event did not occur in coded and/or textual form.
	NotDoneReason *CodeableConcept `json:"notDoneReason"`
	// The reason or justification for the communication.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Indicates another resource whose existence justifies this communication.
	ReasonReference []*Reference `json:"reasonReference"`
	// Text, attachment(s), or resource(s) that was communicated to the recipient.
	Payload []*Communication_Payload `json:"payload"`
}

func NewCommunication() *Communication { return &Communication{ResourceType: "Communication"} }

// MessageHeader is The header for a message exchange that is either requesting or responding to an action.  The reference(s) that are the subject of the action as well as other information related to the action are typically transmitted in a bundle in which the MessageHeader resource instance is the first resource in the bundle.
type MessageHeader struct {
	_ *DomainResource
	// Code that identifies the event this message represents and connects it with its definition. Events defined as part of the FHIR specification have the system value "http://hl7.org/fhir/message-events".
	Event *Coding `json:"event,omitempty"`
	// Identifies the sending system to allow the use of a trust relationship.
	Sender *Reference `json:"sender"`
	// Extensions for timestamp
	Timestamp_ext *Element `json:"_timestamp"`
	// The source application from which this message originated.
	Source *MessageHeader_Source `json:"source,omitempty"`
	// The person or organization that accepts overall responsibility for the contents of the message. The implication is that the message event happened under the policies of the responsible party.
	Responsible *Reference `json:"responsible"`
	// This is a MessageHeader resource
	ResourceType string `json:"resourceType,omitempty"`
	// The logical author of the message - the person or device that decided the described event should happen. When there is more than one candidate, pick the most proximal to the MessageHeader. Can provide other authors in extensions.
	Author *Reference `json:"author"`
	// Allows data conveyed by a message to be addressed to a particular person or department when routing to a specific application isn't sufficient.
	Receiver *Reference `json:"receiver"`
	// The person or device that performed the data entry leading to this message. When there is more than one candidate, pick the most proximal to the message. Can provide other enterers in extensions.
	Enterer *Reference `json:"enterer"`
	// Coded indication of the cause for the event - indicates  a reason for the occurrence of the event that is a focus of this message.
	Reason *CodeableConcept `json:"reason"`
	// The actual data of the message - a reference to the root/focus class of the event.
	Focus []*Reference `json:"focus"`
	// The time that the message was sent.
	Timestamp string `json:"timestamp"`
	// Information about the message that this message is a response to.  Only present if this message is a response.
	Response *MessageHeader_Response `json:"response"`
	// The destination application which the message is intended for.
	Destination []*MessageHeader_Destination `json:"destination"`
}

func NewMessageHeader() *MessageHeader { return &MessageHeader{ResourceType: "MessageHeader"} }

// StructureDefinition_Snapshot is A definition of a FHIR structure. This resource is used to describe the underlying resources, data types defined in FHIR, and also for describing extensions and constraints on resources and data types.
type StructureDefinition_Snapshot struct {
	_ *BackboneElement
	// Captures constraints on each element within the resource.
	Element []*ElementDefinition `json:"element,omitempty"`
}

// Bundle_Entry is A container for a collection of resources.
type Bundle_Entry struct {
	_ *BackboneElement
	// Information about the search process that lead to the creation of this entry.
	Search *Bundle_Search `json:"search"`
	// Additional information about how this entry should be processed as part of a transaction.
	Request *Bundle_Request `json:"request"`
	// Additional information about how this entry should be processed as part of a transaction.
	Response *Bundle_Response `json:"response"`
	// A series of links that provide context to this entry.
	Link []*Bundle_Link `json:"link"`
	// The Absolute URL for the resource.  The fullUrl SHALL not disagree with the id in the resource. The fullUrl is a version independent reference to the resource. The fullUrl element SHALL have a value except that:
	// * fullUrl can be empty on a POST (although it does not need to when specifying a temporary id for reference in the bundle)
	// * Results from operations might involve resources that are not identified.
	FullUrl string `json:"fullUrl"`
	// Extensions for fullUrl
	FullUrl_ext *Element `json:"_fullUrl"`
	// The Resources for the entry.
	Resource interface{} `json:"resource"`
}

// Contract_Agent is A formal agreement between parties regarding the conduct of business, exchange of information or other matters.
type Contract_Agent struct {
	_ *BackboneElement
	// Who or what parties are assigned roles in this Contract.
	Actor *Reference `json:"actor,omitempty"`
	// Role type of agent assigned roles in this Contract.
	Role []*CodeableConcept `json:"role"`
}

// MedicationDispense_Performer is Indicates that a medication product is to be or has been dispensed for a named person/patient.  This includes a description of the medication product (supply) provided and the instructions for administering the medication.  The medication dispense is the result of a pharmacy system responding to a medication order.
type MedicationDispense_Performer struct {
	_ *BackboneElement
	// The device, practitioner, etc. who performed the action.  It should be assumed that the actor is the dispenser of the medication.
	Actor *Reference `json:"actor,omitempty"`
	// The organization the device or practitioner was acting on behalf of.
	OnBehalfOf *Reference `json:"onBehalfOf"`
}

// RequestGroup_Action is A group of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
type RequestGroup_Action struct {
	_ *BackboneElement
	// Extensions for label
	Label_ext *Element `json:"_label"`
	// A text equivalent of the action to be performed. This provides a human-interpretable description of the action when the definition is consumed by a system that may not be capable of interpreting it dynamically.
	TextEquivalent string `json:"textEquivalent"`
	// The type of action to perform (create, update, remove).
	Type *Coding `json:"type"`
	// Extensions for selectionBehavior
	SelectionBehavior_ext *Element `json:"_selectionBehavior"`
	// An optional value describing when the action should be performed.
	TimingRange *Range `json:"timingRange"`
	// Extensions for precheckBehavior
	PrecheckBehavior_ext *Element `json:"_precheckBehavior"`
	// The resource that is the target of the action (e.g. CommunicationRequest).
	Resource *Reference `json:"resource"`
	// Sub actions.
	Action []*RequestGroup_Action `json:"action"`
	// The title of the action displayed to a user.
	Title string `json:"title"`
	// A short description of the action used to provide a summary to display to the user.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for timingDateTime
	TimingDateTime_ext *Element `json:"_timingDateTime"`
	// Extensions for cardinalityBehavior
	CardinalityBehavior_ext *Element `json:"_cardinalityBehavior"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Extensions for textEquivalent
	TextEquivalent_ext *Element `json:"_textEquivalent"`
	// An expression that describes applicability criteria, or start/stop conditions for the action.
	Condition []*RequestGroup_Condition `json:"condition"`
	// An optional value describing when the action should be performed.
	TimingPeriod *Period `json:"timingPeriod"`
	// Defines whether the action can be selected multiple times.
	// pattern [^\s]+([\s]?[^\s]+)*
	CardinalityBehavior string `json:"cardinalityBehavior"`
	// Extensions for groupingBehavior
	GroupingBehavior_ext *Element `json:"_groupingBehavior"`
	// A user-visible label for the action.
	Label string `json:"label"`
	// A code that provides meaning for the action or action group. For example, a section may have a LOINC code for a the section of a documentation template.
	Code []*CodeableConcept `json:"code"`
	// An optional value describing when the action should be performed.
	TimingDuration *Duration `json:"timingDuration"`
	// An optional value describing when the action should be performed.
	TimingTiming *Timing `json:"timingTiming"`
	// Defines the grouping behavior for the action and its children.
	// pattern [^\s]+([\s]?[^\s]+)*
	GroupingBehavior string `json:"groupingBehavior"`
	// The participant that should perform or be responsible for this action.
	Participant []*Reference `json:"participant"`
	// Extensions for requiredBehavior
	RequiredBehavior_ext *Element `json:"_requiredBehavior"`
	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient. Information resources can include inline text commentary and links to web resources.
	Documentation []*RelatedArtifact `json:"documentation"`
	// Defines the requiredness behavior for the action.
	// pattern [^\s]+([\s]?[^\s]+)*
	RequiredBehavior string `json:"requiredBehavior"`
	// A relationship to another action such as "before" or "30-60 minutes after start of".
	RelatedAction []*RequestGroup_RelatedAction `json:"relatedAction"`
	// An optional value describing when the action should be performed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	TimingDateTime time.Time `json:"timingDateTime"`
	// Defines the selection behavior for the action and its children.
	// pattern [^\s]+([\s]?[^\s]+)*
	SelectionBehavior string `json:"selectionBehavior"`
	// Defines whether the action should usually be preselected.
	// pattern [^\s]+([\s]?[^\s]+)*
	PrecheckBehavior string `json:"precheckBehavior"`
}

// SampledData is A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.
type SampledData struct {
	_ *Element
	// Extensions for factor
	Factor_ext *Element `json:"_factor"`
	// The lower limit of detection of the measured points. This is needed if any of the data points have the value "L" (lower than detection limit).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	LowerLimit float64 `json:"lowerLimit"`
	// Extensions for upperLimit
	UpperLimit_ext *Element `json:"_upperLimit"`
	// The number of sample points at each time point. If this value is greater than one, then the dimensions will be interlaced - all the sample points for a point in time will be recorded at once.
	// pattern [1-9][0-9]*
	Dimensions uint64 `json:"dimensions"`
	// The base quantity that a measured value of zero represents. In addition, this provides the units of the entire measurement series.
	Origin *Quantity `json:"origin,omitempty"`
	// The length of time between sampling times, measured in milliseconds.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Period float64 `json:"period"`
	// Extensions for period
	Period_ext *Element `json:"_period"`
	// A correction factor that is applied to the sampled data points before they are added to the origin.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Factor float64 `json:"factor"`
	// A series of data points which are decimal values separated by a single space (character u20). The special values "E" (error), "L" (below detection limit) and "U" (above detection limit) can also be used in place of a decimal value.
	Data string `json:"data"`
	// Extensions for lowerLimit
	LowerLimit_ext *Element `json:"_lowerLimit"`
	// The upper limit of detection of the measured points. This is needed if any of the data points have the value "U" (higher than detection limit).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	UpperLimit float64 `json:"upperLimit"`
	// Extensions for dimensions
	Dimensions_ext *Element `json:"_dimensions"`
	// Extensions for data
	Data_ext *Element `json:"_data"`
}

// CapabilityStatement_Event is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Event struct {
	_ *BackboneElement
	// The impact of the content of the message.
	Category string `json:"category"`
	// Extensions for category
	Category_ext *Element `json:"_category"`
	// The mode of this event declaration - whether an application is a sender or receiver.
	Mode string `json:"mode"`
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// Extensions for focus
	Focus_ext *Element `json:"_focus"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// A coded identifier of a supported messaging event.
	Code *Coding `json:"code,omitempty"`
	// A resource associated with the event.  This is the resource that defines the event.
	// pattern [^\s]+([\s]?[^\s]+)*
	Focus string `json:"focus"`
	// Information about the request for this event.
	Request *Reference `json:"request,omitempty"`
	// Information about the response for this event.
	Response *Reference `json:"response,omitempty"`
	// Guidance on how this event is handled, such as internal system trigger points, business rules, etc.
	Documentation string `json:"documentation"`
}

// TestScript_Rule is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Rule struct {
	_ *BackboneElement
	// Reference to the resource (containing the contents of the rule needed for assertions).
	Resource *Reference `json:"resource,omitempty"`
	// Each rule template can take one or more parameters for rule evaluation.
	Param []*TestScript_Param `json:"param"`
}

// Encounter_Location is An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type Encounter_Location struct {
	_ *BackboneElement
	// The location where the encounter takes place.
	Location *Reference `json:"location,omitempty"`
	// The status of the participants' presence at the specified location during the period specified. If the participant is is no longer at the location, then the period will have an end date/time.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Time period during which the patient was present at the location.
	Period *Period `json:"period"`
}

// Resource is This is the base resource type for everything.
type Resource struct {
	_ *Element
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Id string `json:"id"`
	// Extensions for id
	Id_ext *Element `json:"_id"`
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content may not always be associated with version changes to the resource.
	Meta *Meta `json:"meta"`
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content.
	ImplicitRules string `json:"implicitRules"`
	// Extensions for implicitRules
	ImplicitRules_ext *Element `json:"_implicitRules"`
	// The base language in which the resource is written.
	// pattern [^\s]+([\s]?[^\s]+)*
	Language string `json:"language"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
}

// TestScript_Variable is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Variable struct {
	_ *BackboneElement
	// A default, hard-coded, or user-defined value for this variable.
	DefaultValue string `json:"defaultValue"`
	// Extensions for defaultValue
	DefaultValue_ext *Element `json:"_defaultValue"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The fluentpath expression to evaluate against the fixture body. When variables are defined, only one of either expression, headerField or path must be specified.
	Expression string `json:"expression"`
	// Extensions for sourceId
	SourceId_ext *Element `json:"_sourceId"`
	// Extensions for expression
	Expression_ext *Element `json:"_expression"`
	// Will be used to grab the HTTP header field value from the headers that sourceId is pointing to.
	HeaderField string `json:"headerField"`
	// Displayable text string with hint help information to the user when entering a default value.
	Hint string `json:"hint"`
	// Extensions for hint
	Hint_ext *Element `json:"_hint"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// Descriptive name for this variable.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// A free text natural language description of the variable and its purpose.
	Description string `json:"description"`
	// Extensions for headerField
	HeaderField_ext *Element `json:"_headerField"`
	// Fixture to evaluate the XPath/JSONPath expression or the headerField  against within this variable.
	// pattern [A-Za-z0-9\-\.]{1,64}
	SourceId string `json:"sourceId"`
	// XPath or JSONPath to evaluate against the fixture body.  When variables are defined, only one of either expression, headerField or path must be specified.
	Path string `json:"path"`
}

// ValueSet_Filter is A value set specifies a set of codes drawn from one or more code systems.
type ValueSet_Filter struct {
	_ *BackboneElement
	// A code that identifies a property defined in the code system.
	// pattern [^\s]+([\s]?[^\s]+)*
	Property string `json:"property"`
	// Extensions for property
	Property_ext *Element `json:"_property"`
	// The kind of operation to perform as a part of the filter criteria.
	Op string `json:"op"`
	// Extensions for op
	Op_ext *Element `json:"_op"`
	// The match value may be either a code defined by the system, or a string value, which is a regex match on the literal string of the property value when the operation is 'regex', or one of the values (true and false), when the operation is 'exists'.
	// pattern [^\s]+([\s]?[^\s]+)*
	Value string `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
}

// VisionPrescription_Dispense is An authorization for the supply of glasses and/or contact lenses to a patient.
type VisionPrescription_Dispense struct {
	_ *BackboneElement
	// Power adjustment for astigmatism measured in diopters (0.25 units).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Cylinder float64 `json:"cylinder"`
	// Adjustment for astigmatism measured in integer degrees.
	// pattern -?([0]|([1-9][0-9]*))
	Axis int64 `json:"axis"`
	// Extensions for axis
	Axis_ext *Element `json:"_axis"`
	// The recommended maximum wear period for the lens.
	Duration *Quantity `json:"duration"`
	// Brand recommendations or restrictions.
	Brand string `json:"brand"`
	// Notes for special requirements such as coatings and lens materials.
	Note []*Annotation `json:"note"`
	// Power adjustment for multifocal lenses measured in diopters (0.25 units).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Add float64 `json:"add"`
	// Extensions for power
	Power_ext *Element `json:"_power"`
	// The eye for which the lens applies.
	Eye string `json:"eye"`
	// Extensions for eye
	Eye_ext *Element `json:"_eye"`
	// Amount of prism to compensate for eye alignment in fractional units.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Prism float64 `json:"prism"`
	// Extensions for base
	Base_ext *Element `json:"_base"`
	// Extensions for cylinder
	Cylinder_ext *Element `json:"_cylinder"`
	// Extensions for diameter
	Diameter_ext *Element `json:"_diameter"`
	// Extensions for prism
	Prism_ext *Element `json:"_prism"`
	// Extensions for add
	Add_ext *Element `json:"_add"`
	// Contact lens power measured in diopters (0.25 units).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Power float64 `json:"power"`
	// Lens power measured in diopters (0.25 units).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Sphere float64 `json:"sphere"`
	// The relative base, or reference lens edge, for the prism.
	Base string `json:"base"`
	// Back curvature measured in millimeters.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	BackCurve float64 `json:"backCurve"`
	// Extensions for backCurve
	BackCurve_ext *Element `json:"_backCurve"`
	// Contact lens diameter measured in millimeters.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Diameter float64 `json:"diameter"`
	// Extensions for sphere
	Sphere_ext *Element `json:"_sphere"`
	// Extensions for color
	Color_ext *Element `json:"_color"`
	// Extensions for brand
	Brand_ext *Element `json:"_brand"`
	// Identifies the type of vision correction product which is required for the patient.
	Product *CodeableConcept `json:"product"`
	// Special color or pattern.
	Color string `json:"color"`
}

// ExpansionProfile_Designation is Resource to define constraints on the Expansion of a FHIR ValueSet.
type ExpansionProfile_Designation struct {
	_ *BackboneElement
	// Designations to be included.
	Include *ExpansionProfile_Include `json:"include"`
	// Designations to be excluded.
	Exclude *ExpansionProfile_Exclude `json:"exclude"`
}

// MedicationAdministration_Dosage is Describes the event of a patient consuming or otherwise being administered a medication.  This may be as simple as swallowing a tablet or it may be a long running infusion.  Related resources tie this event to the authorizing prescription, and the specific encounter between patient and health care practitioner.
type MedicationAdministration_Dosage struct {
	_ *BackboneElement
	// Identifies the speed with which the medication was or will be introduced into the patient.  Typically the rate for an infusion e.g. 100 ml per 1 hour or 100 ml/hr.  May also be expressed as a rate per unit of time e.g. 500 ml per 2 hours.  Other examples:  200 mcg/min or 200 mcg/1 minute; 1 liter/8 hours.
	RateSimpleQuantity *Quantity `json:"rateSimpleQuantity"`
	// Free text dosage can be used for cases where the dosage administered is too complex to code. When coded dosage is present, the free text dosage may still be present for display to humans.
	//
	// The dosage instructions should reflect the dosage of the medication that was administered.
	Text string `json:"text"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// A coded specification of the anatomic site where the medication first entered the body.  For example, "left arm".
	Site *CodeableConcept `json:"site"`
	// A code specifying the route or physiological path of administration of a therapeutic agent into or onto the patient.  For example, topical, intravenous, etc.
	Route *CodeableConcept `json:"route"`
	// A coded value indicating the method by which the medication is intended to be or was introduced into or on the body.  This attribute will most often NOT be populated.  It is most commonly used for injections.  For example, Slow Push, Deep IV.
	Method *CodeableConcept `json:"method"`
	// The amount of the medication given at one administration event.   Use this value when the administration is essentially an instantaneous event such as a swallowing a tablet or giving an injection.
	Dose *Quantity `json:"dose"`
	// Identifies the speed with which the medication was or will be introduced into the patient.  Typically the rate for an infusion e.g. 100 ml per 1 hour or 100 ml/hr.  May also be expressed as a rate per unit of time e.g. 500 ml per 2 hours.  Other examples:  200 mcg/min or 200 mcg/1 minute; 1 liter/8 hours.
	RateRatio *Ratio `json:"rateRatio"`
}

// SupplyRequest is A record of a request for a medication, substance or device used in the healthcare setting.
type SupplyRequest struct {
	_ *DomainResource
	// The item being requested.
	OrderedItem *SupplyRequest_OrderedItem `json:"orderedItem"`
	// Status of the supply request.
	Status string `json:"status"`
	// Extensions for priority
	Priority_ext *Element `json:"_priority"`
	// Extensions for authoredOn
	AuthoredOn_ext *Element `json:"_authoredOn"`
	// Who is intended to fulfill the request.
	Supplier []*Reference `json:"supplier"`
	// Where the supply is expected to come from.
	DeliverFrom *Reference `json:"deliverFrom"`
	// Unique identifier for this supply request.
	Identifier *Identifier `json:"identifier"`
	// When the request should be fulfilled.
	OccurrenceTiming *Timing `json:"occurrenceTiming"`
	// Indicates how quickly this SupplyRequest should be addressed with respect to other requests.
	// pattern [^\s]+([\s]?[^\s]+)*
	Priority string `json:"priority"`
	// When the request should be fulfilled.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OccurrenceDateTime time.Time `json:"occurrenceDateTime"`
	// When the request should be fulfilled.
	OccurrencePeriod *Period `json:"occurrencePeriod"`
	// The individual who initiated the request and has responsibility for its activation.
	Requester *SupplyRequest_Requester `json:"requester"`
	// This is a SupplyRequest resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// When the request was made.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AuthoredOn time.Time `json:"authoredOn"`
	// Why the supply item was requested.
	ReasonCodeableConcept *CodeableConcept `json:"reasonCodeableConcept"`
	// Why the supply item was requested.
	ReasonReference *Reference `json:"reasonReference"`
	// Where the supply is destined to go.
	DeliverTo *Reference `json:"deliverTo"`
	// Category of supply, e.g.  central, non-stock, etc. This is used to support work flows associated with the supply process.
	Category *CodeableConcept `json:"category"`
	// Extensions for occurrenceDateTime
	OccurrenceDateTime_ext *Element `json:"_occurrenceDateTime"`
}

func NewSupplyRequest() *SupplyRequest { return &SupplyRequest{ResourceType: "SupplyRequest"} }

// Contract_Friendly is A formal agreement between parties regarding the conduct of business, exchange of information or other matters.
type Contract_Friendly struct {
	_ *BackboneElement
	// Human readable rendering of this Contract in a format and representation intended to enhance comprehension and ensure understandability.
	ContentAttachment *Attachment `json:"contentAttachment"`
	// Human readable rendering of this Contract in a format and representation intended to enhance comprehension and ensure understandability.
	ContentReference *Reference `json:"contentReference"`
}

// ImagingManifest_Study is A text description of the DICOM SOP instances selected in the ImagingManifest; or the reason for, or significance of, the selection.
type ImagingManifest_Study struct {
	_ *BackboneElement
	// Extensions for uid
	Uid_ext *Element `json:"_uid"`
	// Reference to the Imaging Study in FHIR form.
	ImagingStudy *Reference `json:"imagingStudy"`
	// The network service providing access (e.g., query, view, or retrieval) for the study. See implementation notes for information about using DICOM endpoints. A study-level endpoint applies to each series in the study, unless overridden by a series-level endpoint with the same Endpoint.type.
	Endpoint []*Reference `json:"endpoint"`
	// Series identity and locating information of the DICOM SOP instances in the selection.
	Series []*ImagingManifest_Series `json:"series,omitempty"`
	// Study instance UID of the SOP instances in the selection.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	Uid string `json:"uid"`
}

// MeasureReport_Stratum is The MeasureReport resource contains the results of evaluating a measure.
type MeasureReport_Stratum struct {
	_ *BackboneElement
	// The value for this stratum, expressed as a string. When defining stratifiers on complex values, the value must be rendered such that the value for each stratum within the stratifier is unique.
	Value string `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// The populations that make up the stratum, one for each type of population appropriate to the measure.
	Population []*MeasureReport_Population1 `json:"population"`
	// The measure score for this stratum, calculated as appropriate for the measure type and scoring method, and based on only the members of this stratum.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	MeasureScore float64 `json:"measureScore"`
	// Extensions for measureScore
	MeasureScore_ext *Element `json:"_measureScore"`
}

// ProcessResponse_ProcessNote is This resource provides processing status, errors and notes from the processing of a resource.
type ProcessResponse_ProcessNote struct {
	_ *BackboneElement
	// The note purpose: Print/Display.
	Type *CodeableConcept `json:"type"`
	// The note text.
	Text string `json:"text"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
}

// AdverseEvent_SuspectEntity is Actual or  potential/avoided event causing unintended physical injury resulting from or contributed to by medical care, a research study or other healthcare setting factors that requires additional monitoring, treatment, or hospitalization, or that results in death.
type AdverseEvent_SuspectEntity struct {
	_ *BackboneElement
	// method1 | method2.
	CausalityMethod *CodeableConcept `json:"causalityMethod"`
	// result1 | result2.
	CausalityResult *CodeableConcept `json:"causalityResult"`
	// Extensions for causality
	Causality_ext *Element `json:"_causality"`
	// assess1 | assess2.
	CausalityAssessment *CodeableConcept `json:"causalityAssessment"`
	// AdverseEvent.suspectEntity.causalityProductRelatedness.
	CausalityProductRelatedness string `json:"causalityProductRelatedness"`
	// Extensions for causalityProductRelatedness
	CausalityProductRelatedness_ext *Element `json:"_causalityProductRelatedness"`
	// AdverseEvent.suspectEntity.causalityAuthor.
	CausalityAuthor *Reference `json:"causalityAuthor"`
	// Identifies the actual instance of what caused the adverse event.  May be a substance, medication, medication administration, medication statement or a device.
	Instance *Reference `json:"instance,omitempty"`
	// causality1 | causality2.
	Causality string `json:"causality"`
}

// Medication is This resource is primarily used for the identification and definition of a medication. It covers the ingredients and the packaging for a medication.
type Medication struct {
	_ *DomainResource
	// A code (or set of codes) that specify this medication, or a textual description if no code is available. Usage note: This could be a standard medication code such as a code from RxNorm, SNOMED CT, IDMP etc. It could also be a national or local formulary code, optionally with translations to other code systems.
	Code *CodeableConcept `json:"code"`
	// A code to indicate if the medication is in active use.
	Status string `json:"status"`
	// Set to true if the item is attributable to a specific manufacturer.
	IsBrand bool `json:"isBrand"`
	// Extensions for isBrand
	IsBrand_ext *Element `json:"_isBrand"`
	// Identifies a particular constituent of interest in the product.
	Ingredient []*Medication_Ingredient `json:"ingredient"`
	// Photo(s) or graphic representation(s) of the medication.
	Image []*Attachment `json:"image"`
	// Information that only applies to packages (not products).
	Package *Medication_Package `json:"package"`
	// This is a Medication resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Set to true if the medication can be obtained without an order from a prescriber.
	IsOverTheCounter bool `json:"isOverTheCounter"`
	// Extensions for isOverTheCounter
	IsOverTheCounter_ext *Element `json:"_isOverTheCounter"`
	// Describes the details of the manufacturer of the medication product.  This is not intended to represent the distributor of a medication product.
	Manufacturer *Reference `json:"manufacturer"`
	// Describes the form of the item.  Powder; tablets; capsule.
	Form *CodeableConcept `json:"form"`
}

func NewMedication() *Medication { return &Medication{ResourceType: "Medication"} }

// Procedure_FocalDevice is An action that is or was performed on a patient. This can be a physical intervention like an operation, or less invasive like counseling or hypnotherapy.
type Procedure_FocalDevice struct {
	_ *BackboneElement
	// The kind of change that happened to the device during the procedure.
	Action *CodeableConcept `json:"action"`
	// The device that was manipulated (changed) during the procedure.
	Manipulated *Reference `json:"manipulated,omitempty"`
}

// SearchParameter is A search parameter that defines a named search item that can be used to search/filter on a resource.
type SearchParameter struct {
	_ *DomainResource
	// A natural language name identifying the search parameter. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// Explaination of why this search parameter is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// The type of value a search parameter refers to, and how the content is interpreted.
	Type string `json:"type"`
	// Where this search parameter is originally defined. If a derivedFrom is provided, then the details in the search parameter must be consistent with the definition from which it is defined. I.e. the parameter should have the same meaning, and (usually) the functionality should be a proper subset of the underlying search parameter.
	DerivedFrom string `json:"derivedFrom"`
	// An absolute URI that is used to identify this search parameter when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this search parameter is (or will be) published. The URL SHOULD include the major version of the search parameter. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// A legal or geographic region in which the search parameter is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// A free text natural language description of the search parameter from a consumer's perspective. and how it used.
	Description string `json:"description"`
	// Extensions for xpathUsage
	XpathUsage_ext *Element `json:"_xpathUsage"`
	// Extensions for target
	Target_ext []*Element `json:"_target"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate search parameter instances.
	UseContext []*UsageContext `json:"useContext"`
	// Extensions for comparator
	Comparator_ext []*Element `json:"_comparator"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The status of this search parameter. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// Extensions for base
	Base_ext []*Element `json:"_base"`
	// Extensions for derivedFrom
	DerivedFrom_ext *Element `json:"_derivedFrom"`
	// A FHIRPath expression that returns a set of elements for the search parameter.
	Expression string `json:"expression"`
	// An XPath expression that returns a set of elements for the search parameter.
	Xpath string `json:"xpath"`
	// How the search parameter relates to the set of elements returned by evaluating the xpath query.
	XpathUsage string `json:"xpathUsage"`
	// The date  (and optionally time) when the search parameter was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the search parameter changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Types of resource (if a resource is referenced).
	Target []string `json:"target"`
	// A modifier supported for the search parameter.
	Modifier []string `json:"modifier"`
	// Extensions for modifier
	Modifier_ext []*Element `json:"_modifier"`
	// Contains the names of any search parameters which may be chained to the containing search parameter. Chained parameters may be added to search parameters of type reference, and specify that resources will only be returned if they contain a reference to a resource which matches the chained parameter value. Values for this field should be drawn from SearchParameter.code for a parameter on the target resource type.
	Chain []string `json:"chain"`
	// This is a SearchParameter resource
	ResourceType string `json:"resourceType,omitempty"`
	// The identifier that is used to identify this version of the search parameter when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the search parameter author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// The base resource type(s) that this search parameter can be used against.
	Base []string `json:"base"`
	// Extensions for chain
	Chain_ext []*Element `json:"_chain"`
	// A boolean value to indicate that this search parameter is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for xpath
	Xpath_ext *Element `json:"_xpath"`
	// Comparators supported for the search parameter.
	Comparator []string `json:"comparator"`
	// Used to define the parts of a composite search parameter.
	Component []*SearchParameter_Component `json:"component"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// The name of the individual or organization that published the search parameter.
	Publisher string `json:"publisher"`
	// The code used in the URL or the parameter name in a parameters resource for this search parameter.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Extensions for expression
	Expression_ext *Element `json:"_expression"`
}

func NewSearchParameter() *SearchParameter { return &SearchParameter{ResourceType: "SearchParameter"} }

// Specimen is A sample to be used for analysis.
type Specimen struct {
	_ *DomainResource
	// To communicate any details or issues about the specimen or during the specimen collection. (for example: broken vial, sent with patient, frozen).
	Note []*Annotation `json:"note"`
	// Id for specimen.
	Identifier []*Identifier `json:"identifier"`
	// The kind of material that forms the specimen.
	Type *CodeableConcept `json:"type"`
	// Where the specimen came from. This may be from the patient(s) or from the environment or a device.
	Subject *Reference `json:"subject,omitempty"`
	// The container holding the specimen.  The recursive nature of containers; i.e. blood in tube in tray in rack is not addressed here.
	Container []*Specimen_Container `json:"container"`
	// This is a Specimen resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for receivedTime
	ReceivedTime_ext *Element `json:"_receivedTime"`
	// Reference to the parent (source) specimen which is used when the specimen was either derived from or a component of another specimen.
	Parent []*Reference `json:"parent"`
	// Details concerning a test or procedure request that required a specimen to be collected.
	Request []*Reference `json:"request"`
	// Details concerning processing and processing steps for the specimen.
	Processing []*Specimen_Processing `json:"processing"`
	// Details concerning the specimen collection.
	Collection *Specimen_Collection `json:"collection"`
	// The identifier assigned by the lab when accessioning specimen(s). This is not necessarily the same as the specimen identifier, depending on local lab procedures.
	AccessionIdentifier *Identifier `json:"accessionIdentifier"`
	// The availability of the specimen.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Time when specimen was received for processing or testing.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ReceivedTime time.Time `json:"receivedTime"`
}

func NewSpecimen() *Specimen { return &Specimen{ResourceType: "Specimen"} }

// SupplyDelivery_SuppliedItem is Record of delivery of what is supplied.
type SupplyDelivery_SuppliedItem struct {
	_ *BackboneElement
	// The amount of supply that has been dispensed. Includes unit of measure.
	Quantity *Quantity `json:"quantity"`
	// Identifies the medication, substance or device being dispensed. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept"`
	// Identifies the medication, substance or device being dispensed. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	ItemReference *Reference `json:"itemReference"`
}

// CodeableConcept is A concept that may be defined by a formal reference to a terminology or ontology or may be provided by text.
type CodeableConcept struct {
	_ *Element
	// A reference to a code defined by a terminology system.
	Coding []*Coding `json:"coding"`
	// A human language representation of the concept as seen/selected/uttered by the user who entered the data and/or which represents the intended meaning of the user.
	Text string `json:"text"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
}

// ElementDefinition_Discriminator is Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition_Discriminator struct {
	_ *BackboneElement
	// How the element value is interpreted when discrimination is evaluated.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// A FHIRPath expression, using a restricted subset of FHIRPath, that is used to identify the element on which discrimination is based.
	Path string `json:"path"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
}

// Location_Position is Details and position information for a physical place where services are provided  and resources and participants may be stored, found, contained or accommodated.
type Location_Position struct {
	_ *BackboneElement
	// Extensions for altitude
	Altitude_ext *Element `json:"_altitude"`
	// Longitude. The value domain and the interpretation are the same as for the text of the longitude element in KML (see notes below).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Longitude float64 `json:"longitude"`
	// Extensions for longitude
	Longitude_ext *Element `json:"_longitude"`
	// Latitude. The value domain and the interpretation are the same as for the text of the latitude element in KML (see notes below).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Latitude float64 `json:"latitude"`
	// Extensions for latitude
	Latitude_ext *Element `json:"_latitude"`
	// Altitude. The value domain and the interpretation are the same as for the text of the altitude element in KML (see notes below).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Altitude float64 `json:"altitude"`
}

// QuestionnaireResponse is A structured set of questions and their answers. The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to.
type QuestionnaireResponse struct {
	_ *DomainResource
	// A procedure or observation that this questionnaire was performed as part of the execution of.  For example, the surgery a checklist was executed as part of.
	Parent []*Reference `json:"parent"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The subject of the questionnaire response.  This could be a patient, organization, practitioner, device, etc.  This is who/what the answers apply to, but is not necessarily the source of information.
	Subject *Reference `json:"subject"`
	// The person who answered the questions about the subject.
	Source *Reference `json:"source"`
	// The order, proposal or plan that is fulfilled in whole or in part by this QuestionnaireResponse.  For example, a ProcedureRequest seeking an intake assessment or a decision support recommendation to assess for post-partum depression.
	BasedOn []*Reference `json:"basedOn"`
	// The position of the questionnaire response within its overall lifecycle.
	Status string `json:"status"`
	// The encounter or episode of care with primary association to the questionnaire response.
	Context *Reference `json:"context"`
	// Extensions for authored
	Authored_ext *Element `json:"_authored"`
	// Person who received the answers to the questions in the QuestionnaireResponse and recorded them in the system.
	Author *Reference `json:"author"`
	// This is a QuestionnaireResponse resource
	ResourceType string `json:"resourceType,omitempty"`
	// A business identifier assigned to a particular completed (or partially completed) questionnaire.
	Identifier *Identifier `json:"identifier"`
	// The date and/or time that this set of answers were last changed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Authored time.Time `json:"authored"`
	// The Questionnaire that defines and organizes the questions for which answers are being provided.
	Questionnaire *Reference `json:"questionnaire"`
	// A group or question item from the original questionnaire for which answers are provided.
	Item []*QuestionnaireResponse_Item `json:"item"`
}

func NewQuestionnaireResponse() *QuestionnaireResponse {
	return &QuestionnaireResponse{ResourceType: "QuestionnaireResponse"}
}

// QuestionnaireResponse_Item is A structured set of questions and their answers. The questions are ordered and grouped into coherent subsets, corresponding to the structure of the grouping of the questionnaire being responded to.
type QuestionnaireResponse_Item struct {
	_ *BackboneElement
	// The item from the Questionnaire that corresponds to this item in the QuestionnaireResponse resource.
	LinkId string `json:"linkId"`
	// A reference to an [[[ElementDefinition]]] that provides the details for the item.
	Definition string `json:"definition"`
	// Extensions for definition
	Definition_ext *Element `json:"_definition"`
	// More specific subject this section's answers are about, details the subject given in QuestionnaireResponse.
	Subject *Reference `json:"subject"`
	// Questions or sub-groups nested beneath a question or group.
	Item []*QuestionnaireResponse_Item `json:"item"`
	// Extensions for linkId
	LinkId_ext *Element `json:"_linkId"`
	// Text that is displayed above the contents of the group or as the text of the question being answered.
	Text string `json:"text"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// The respondent's answer(s) to the question.
	Answer []*QuestionnaireResponse_Answer `json:"answer"`
}

// TestScript_Origin is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Origin struct {
	_ *BackboneElement
	// The type of origin profile the test system supports.
	Profile *Coding `json:"profile,omitempty"`
	// Abstract name given to an origin server in this test script.  The name is provided as a number starting at 1.
	// pattern -?([0]|([1-9][0-9]*))
	Index int64 `json:"index"`
	// Extensions for index
	Index_ext *Element `json:"_index"`
}

// TestScript_Ruleset1 is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Ruleset1 struct {
	_ *BackboneElement
	// The TestScript.ruleset id value this assert will evaluate.
	// pattern [A-Za-z0-9\-\.]{1,64}
	RulesetId string `json:"rulesetId"`
	// Extensions for rulesetId
	RulesetId_ext *Element `json:"_rulesetId"`
	// The referenced rule within the external ruleset template.
	Rule []*TestScript_Rule3 `json:"rule"`
}

// Address is An address expressed using postal conventions (as opposed to GPS or other location definition formats).  This data type may be used to convey addresses for use in delivering mail as well as for visiting locations which might not be valid for mail delivery.  There are a variety of postal address formats defined around the world.
type Address struct {
	_ *Element
	// Extensions for line
	Line_ext []*Element `json:"_line"`
	// The name of the administrative area (county).
	District string `json:"district"`
	// Extensions for postalCode
	PostalCode_ext *Element `json:"_postalCode"`
	// Country - a nation as commonly understood or generally accepted.
	Country string `json:"country"`
	// Extensions for country
	Country_ext *Element `json:"_country"`
	// Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.
	Type string `json:"type"`
	// A full text representation of the address.
	Text string `json:"text"`
	// This component contains the house number, apartment number, street name, street direction,  P.O. Box number, delivery hints, and similar address information.
	Line []string `json:"line"`
	// A postal code designating a region defined by the postal service.
	PostalCode string `json:"postalCode"`
	// The purpose of this address.
	Use string `json:"use"`
	// Extensions for city
	City_ext *Element `json:"_city"`
	// Extensions for state
	State_ext *Element `json:"_state"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The name of the city, town, village or other community or delivery center.
	City string `json:"city"`
	// Time period when address was/is in use.
	Period *Period `json:"period"`
	// Sub-unit of a country with limited sovereignty in a federally organized country. A code may be used if codes are in common use (i.e. US 2 letter state codes).
	State string `json:"state"`
	// Extensions for use
	Use_ext *Element `json:"_use"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// Extensions for district
	District_ext *Element `json:"_district"`
}

// EligibilityResponse_Financial is This resource provides eligibility and plan details from the processing of an Eligibility resource.
type EligibilityResponse_Financial struct {
	_ *BackboneElement
	// Extensions for usedUnsignedInt
	UsedUnsignedInt_ext *Element `json:"_usedUnsignedInt"`
	// Benefits used.
	UsedMoney *Money `json:"usedMoney"`
	// Deductable, visits, benefit amount.
	Type *CodeableConcept `json:"type,omitempty"`
	// Benefits allowed.
	AllowedString string `json:"allowedString"`
	// Benefits used.
	// pattern [0]|([1-9][0-9]*)
	UsedUnsignedInt uint64 `json:"usedUnsignedInt"`
	// Benefits allowed.
	AllowedMoney *Money `json:"allowedMoney"`
	// Benefits allowed.
	// pattern [0]|([1-9][0-9]*)
	AllowedUnsignedInt uint64 `json:"allowedUnsignedInt"`
	// Extensions for allowedUnsignedInt
	AllowedUnsignedInt_ext *Element `json:"_allowedUnsignedInt"`
	// Extensions for allowedString
	AllowedString_ext *Element `json:"_allowedString"`
}

// TestReport_Action2 is A summary of information based on the results of executing a TestScript.
type TestReport_Action2 struct {
	_ *BackboneElement
	// An operation would involve a REST request to a server.
	Operation *TestReport_Operation `json:"operation,omitempty"`
}

// TestScript_Param2 is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Param2 struct {
	_ *BackboneElement
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The value for the parameter that will be passed on to the external rule template.
	Value string `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// Descriptive name for this parameter that matches the external assert rule parameter name.
	Name string `json:"name"`
}

// CapabilityStatement_Interaction1 is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Interaction1 struct {
	_ *BackboneElement
	// A coded identifier of the operation, supported by the system.
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// Guidance specific to the implementation of this operation, such as limitations on the kind of transactions allowed, or information about system wide search is implemented.
	Documentation string `json:"documentation"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
}

// DeviceUseStatement is A record of a device being used by a patient where the record is the result of a report from the patient or another clinician.
type DeviceUseStatement struct {
	_ *DomainResource
	// The patient who used the device.
	Subject *Reference `json:"subject,omitempty"`
	// The time period over which the device was used.
	WhenUsed *Period `json:"whenUsed"`
	// How often the device was used.
	TimingPeriod *Period `json:"timingPeriod"`
	// The time at which the statement was made/recorded.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	RecordedOn time.Time `json:"recordedOn"`
	// Extensions for recordedOn
	RecordedOn_ext *Element `json:"_recordedOn"`
	// Who reported the device was being used by the patient.
	Source *Reference `json:"source"`
	// Indicates the site on the subject's body where the device was used ( i.e. the target site).
	BodySite *CodeableConcept `json:"bodySite"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// How often the device was used.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	TimingDateTime time.Time `json:"timingDateTime"`
	// Reason or justification for the use of the device.
	Indication []*CodeableConcept `json:"indication"`
	// This is a DeviceUseStatement resource
	ResourceType string `json:"resourceType,omitempty"`
	// An external identifier for this statement such as an IRI.
	Identifier []*Identifier `json:"identifier"`
	// A code representing the patient or other source's judgment about the state of the device used that this statement is about.  Generally this will be active or completed.
	Status string `json:"status"`
	// How often the device was used.
	TimingTiming *Timing `json:"timingTiming"`
	// Extensions for timingDateTime
	TimingDateTime_ext *Element `json:"_timingDateTime"`
	// The details of the device used.
	Device *Reference `json:"device,omitempty"`
	// Details about the device statement that were not represented at all or sufficiently in one of the attributes provided in a class. These may include for example a comment, an instruction, or a note associated with the statement.
	Note []*Annotation `json:"note"`
}

func NewDeviceUseStatement() *DeviceUseStatement {
	return &DeviceUseStatement{ResourceType: "DeviceUseStatement"}
}

// NutritionOrder_Administration is A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrder_Administration struct {
	_ *BackboneElement
	// The time period and frequency at which the enteral formula should be delivered to the patient.
	Schedule *Timing `json:"schedule"`
	// The volume of formula to provide to the patient per the specified administration schedule.
	Quantity *Quantity `json:"quantity"`
	// The rate of administration of formula via a feeding pump, e.g. 60 mL per hour, according to the specified schedule.
	RateSimpleQuantity *Quantity `json:"rateSimpleQuantity"`
	// The rate of administration of formula via a feeding pump, e.g. 60 mL per hour, according to the specified schedule.
	RateRatio *Ratio `json:"rateRatio"`
}

// Observation_Related is Measurements and simple assertions made about a patient, device or other subject.
type Observation_Related struct {
	_ *BackboneElement
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// A reference to the observation or [[[QuestionnaireResponse]]] resource that is related to this observation.
	Target *Reference `json:"target,omitempty"`
	// A code specifying the kind of relationship that exists with the target resource.
	Type string `json:"type"`
}

// Parameters_Parameter is This special resource type is used to represent an operation request and response (operations.html). It has no other use, and there is no RESTful endpoint associated with it.
type Parameters_Parameter struct {
	_ *BackboneElement
	// If the parameter is a data type.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	ValueOid string `json:"valueOid"`
	// If the parameter is a data type.
	ValueExtension *Extension `json:"valueExtension"`
	// If the parameter is a data type.
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition"`
	// Extensions for valueInteger
	ValueInteger_ext *Element `json:"_valueInteger"`
	// Extensions for valueBase64Binary
	ValueBase64Binary_ext *Element `json:"_valueBase64Binary"`
	// Extensions for valueDate
	ValueDate_ext *Element `json:"_valueDate"`
	// If the parameter is a data type.
	// pattern urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	ValueUuid string `json:"valueUuid"`
	// Extensions for valueId
	ValueId_ext *Element `json:"_valueId"`
	// If the parameter is a data type.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// If the parameter is a data type.
	ValueAge *Age `json:"valueAge"`
	// If the parameter is a data type.
	ValueElementDefinition *ElementDefinition `json:"valueElementDefinition"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// If the parameter is a data type.
	// pattern -?([0]|([1-9][0-9]*))
	ValueInteger int64 `json:"valueInteger"`
	// Extensions for valueMarkdown
	ValueMarkdown_ext *Element `json:"_valueMarkdown"`
	// If the parameter is a data type.
	ValueAttachment *Attachment `json:"valueAttachment"`
	// Extensions for valueUuid
	ValueUuid_ext *Element `json:"_valueUuid"`
	// If the parameter is a data type.
	// pattern [0]|([1-9][0-9]*)
	ValueUnsignedInt uint64 `json:"valueUnsignedInt"`
	// If the parameter is a data type.
	ValueTiming *Timing `json:"valueTiming"`
	// If the parameter is a data type.
	ValueInstant string `json:"valueInstant"`
	// Extensions for valueOid
	ValueOid_ext *Element `json:"_valueOid"`
	// Extensions for valueInstant
	ValueInstant_ext *Element `json:"_valueInstant"`
	// Extensions for valueUnsignedInt
	ValueUnsignedInt_ext *Element `json:"_valueUnsignedInt"`
	// If the parameter is a data type.
	ValueCount *Count `json:"valueCount"`
	// If the parameter is a data type.
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition"`
	// A named part of a multi-part parameter.
	Part []*Parameters_Parameter `json:"part"`
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// Extensions for valueTime
	ValueTime_ext *Element `json:"_valueTime"`
	// If the parameter is a data type.
	// pattern [1-9][0-9]*
	ValuePositiveInt uint64 `json:"valuePositiveInt"`
	// If the parameter is a data type.
	ValueBackboneElement *BackboneElement `json:"valueBackboneElement"`
	// If the parameter is a data type.
	ValueDistance *Distance `json:"valueDistance"`
	// If the parameter is a data type.
	ValueHumanName *HumanName `json:"valueHumanName"`
	// If the parameter is a data type.
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement"`
	// If the parameter is a whole resource.
	Resource interface{} `json:"resource"`
	// If the parameter is a data type.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ValueDateTime time.Time `json:"valueDateTime"`
	// If the parameter is a data type.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	ValueTime time.Time `json:"valueTime"`
	// If the parameter is a data type.
	ValueDuration *Duration `json:"valueDuration"`
	// If the parameter is a data type.
	ValueDosage *Dosage `json:"valueDosage"`
	// If the parameter is a data type.
	// pattern [^\s]+([\s]?[^\s]+)*
	ValueCode string `json:"valueCode"`
	// If the parameter is a data type.
	ValueNarrative *Narrative `json:"valueNarrative"`
	// If the parameter is a data type.
	ValueAnnotation *Annotation `json:"valueAnnotation"`
	// If the parameter is a data type.
	ValueContributor *Contributor `json:"valueContributor"`
	// The name of the parameter (reference to the operation definition).
	Name string `json:"name"`
	// Extensions for valueDecimal
	ValueDecimal_ext *Element `json:"_valueDecimal"`
	// If the parameter is a data type.
	ValueBase64Binary string `json:"valueBase64Binary"`
	// Extensions for valuePositiveInt
	ValuePositiveInt_ext *Element `json:"_valuePositiveInt"`
	// If the parameter is a data type.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	// If the parameter is a data type.
	ValueRange *Range `json:"valueRange"`
	// If the parameter is a data type.
	ValueElement *Element `json:"valueElement"`
	// If the parameter is a data type.
	ValueSignature *Signature `json:"valueSignature"`
	// If the parameter is a data type.
	ValueIdentifier *Identifier `json:"valueIdentifier"`
	// If the parameter is a data type.
	ValueCoding *Coding `json:"valueCoding"`
	// If the parameter is a data type.
	ValueRatio *Ratio `json:"valueRatio"`
	// If the parameter is a data type.
	ValueUsageContext *UsageContext `json:"valueUsageContext"`
	// Extensions for valueUri
	ValueUri_ext *Element `json:"_valueUri"`
	// If the parameter is a data type.
	ValueReference *Reference `json:"valueReference"`
	// If the parameter is a data type.
	ValueMeta *Meta `json:"valueMeta"`
	// If the parameter is a data type.
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact"`
	// If the parameter is a data type.
	ValueBoolean bool `json:"valueBoolean"`
	// If the parameter is a data type.
	ValueString string `json:"valueString"`
	// Extensions for valueDateTime
	ValueDateTime_ext *Element `json:"_valueDateTime"`
	// If the parameter is a data type.
	// pattern [A-Za-z0-9\-\.]{1,64}
	ValueId string `json:"valueId"`
	// If the parameter is a data type.
	ValueContactPoint *ContactPoint `json:"valueContactPoint"`
	// If the parameter is a data type.
	ValueContactDetail *ContactDetail `json:"valueContactDetail"`
	// If the parameter is a data type.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	ValueDecimal float64 `json:"valueDecimal"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// If the parameter is a data type.
	ValueUri string `json:"valueUri"`
	// Extensions for valueCode
	ValueCode_ext *Element `json:"_valueCode"`
	// If the parameter is a data type.
	ValueMarkdown string `json:"valueMarkdown"`
	// If the parameter is a data type.
	ValueSimpleQuantity *Quantity `json:"valueSimpleQuantity"`
	// If the parameter is a data type.
	ValuePeriod *Period `json:"valuePeriod"`
	// If the parameter is a data type.
	ValueSampledData *SampledData `json:"valueSampledData"`
	// If the parameter is a data type.
	ValueAddress *Address `json:"valueAddress"`
	// If the parameter is a data type.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ValueDate time.Time `json:"valueDate"`
	// If the parameter is a data type.
	ValueMoney *Money `json:"valueMoney"`
}

// TestScript_Param1 is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Param1 struct {
	_ *BackboneElement
	// Descriptive name for this parameter that matches the external assert ruleset rule parameter name.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The value for the parameter that will be passed on to the external ruleset rule template.
	Value string `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
}

// Consent_Actor1 is A record of a healthcare consumer's policy choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent_Actor1 struct {
	_ *BackboneElement
	// How the individual is involved in the resources content that is described in the exception.
	Role *CodeableConcept `json:"role,omitempty"`
	// The resource that identifies the actor. To identify a actors by type, use group to identify a set of actors by some property they share (e.g. 'admitting officers').
	Reference *Reference `json:"reference,omitempty"`
}

// EligibilityResponse_Insurance is This resource provides eligibility and plan details from the processing of an Eligibility resource.
type EligibilityResponse_Insurance struct {
	_ *BackboneElement
	// A suite of updated or additional Coverages from the Insurer.
	Coverage *Reference `json:"coverage"`
	// The contract resource which may provide more detailed information.
	Contract *Reference `json:"contract"`
	// Benefits and optionally current balances by Category.
	BenefitBalance []*EligibilityResponse_BenefitBalance `json:"benefitBalance"`
}

// ImmunizationRecommendation_Recommendation is A patient's point-in-time immunization and recommendation (i.e. forecasting a patient's immunization eligibility according to a published schedule) with optional supporting justification.
type ImmunizationRecommendation_Recommendation struct {
	_ *BackboneElement
	// The targeted disease for the recommendation.
	TargetDisease *CodeableConcept `json:"targetDisease"`
	// The next recommended dose number (e.g. dose 2 is the next recommended dose).
	// pattern [1-9][0-9]*
	DoseNumber uint64 `json:"doseNumber"`
	// Extensions for doseNumber
	DoseNumber_ext *Element `json:"_doseNumber"`
	// Vaccine administration status.
	ForecastStatus *CodeableConcept `json:"forecastStatus,omitempty"`
	// Immunization event history that supports the status and recommendation.
	SupportingImmunization []*Reference `json:"supportingImmunization"`
	// Patient Information that supports the status and recommendation.  This includes patient observations, adverse reactions and allergy/intolerance information.
	SupportingPatientInformation []*Reference `json:"supportingPatientInformation"`
	// The date the immunization recommendation was created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Vaccine that pertains to the recommendation.
	VaccineCode *CodeableConcept `json:"vaccineCode"`
	// Vaccine date recommendations.  For example, earliest date to administer, latest date to administer, etc.
	DateCriterion []*ImmunizationRecommendation_DateCriterion `json:"dateCriterion"`
	// Contains information about the protocol under which the vaccine was administered.
	Protocol *ImmunizationRecommendation_Protocol `json:"protocol"`
}

// MeasureReport_Stratifier is The MeasureReport resource contains the results of evaluating a measure.
type MeasureReport_Stratifier struct {
	_ *BackboneElement
	// The identifier of this stratifier, as defined in the measure definition.
	Identifier *Identifier `json:"identifier"`
	// This element contains the results for a single stratum within the stratifier. For example, when stratifying on administrative gender, there will be four strata, one for each possible gender value.
	Stratum []*MeasureReport_Stratum `json:"stratum"`
}

// DiagnosticReport is The findings and interpretation of diagnostic  tests performed on patients, groups of patients, devices, and locations, and/or specimens derived from these. The report includes clinical context such as requesting and provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports.
type DiagnosticReport struct {
	_ *DomainResource
	// Indicates who or what participated in producing the report.
	Performer []*DiagnosticReport_Performer `json:"performer"`
	// One or more links to full details of any imaging performed during the diagnostic investigation. Typically, this is imaging performed by DICOM enabled modalities, but this is not required. A fully enabled PACS viewer can use this information to provide views of the source images.
	ImagingStudy []*Reference `json:"imagingStudy"`
	// Concise and clinically contextualized impression / summary of the diagnostic report.
	Conclusion string `json:"conclusion"`
	// Codes for the conclusion.
	CodedDiagnosis []*CodeableConcept `json:"codedDiagnosis"`
	// Details concerning a test or procedure requested.
	BasedOn []*Reference `json:"basedOn"`
	// A code or name that describes this diagnostic report.
	Code *CodeableConcept `json:"code,omitempty"`
	// Observations that are part of this diagnostic report. Observations can be simple name/value pairs (e.g. "atomic" results), or they can be grouping observations that include references to other members of the group (e.g. "panels").
	Result []*Reference `json:"result"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for effectiveDateTime
	EffectiveDateTime_ext *Element `json:"_effectiveDateTime"`
	// The subject of the report. Usually, but not always, this is a patient. However diagnostic services also perform analyses on specimens collected from a variety of other sources.
	Subject *Reference `json:"subject"`
	// The healthcare event  (e.g. a patient and healthcare provider interaction) which this DiagnosticReport per is about.
	Context *Reference `json:"context"`
	// The date and time that this version of the report was released from the source diagnostic service.
	Issued string `json:"issued"`
	// Extensions for issued
	Issued_ext *Element `json:"_issued"`
	// Details about the specimens on which this diagnostic report is based.
	Specimen []*Reference `json:"specimen"`
	// A list of key images associated with this report. The images are generally created during the diagnostic process, and may be directly of the patient, or of treated specimens (i.e. slides of interest).
	Image []*DiagnosticReport_Image `json:"image"`
	// Identifiers assigned to this report by the performer or other systems.
	Identifier []*Identifier `json:"identifier"`
	// The status of the diagnostic report as a whole.
	Status string `json:"status"`
	// Rich text representation of the entire result as issued by the diagnostic service. Multiple formats are allowed but they SHALL be semantically equivalent.
	PresentedForm []*Attachment `json:"presentedForm"`
	// The time or time-period the observed values are related to. When the subject of the report is a patient, this is usually either the time of the procedure or of specimen collection(s), but very often the source of the date/time is not known, only the date/time itself.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	EffectiveDateTime time.Time `json:"effectiveDateTime"`
	// The time or time-period the observed values are related to. When the subject of the report is a patient, this is usually either the time of the procedure or of specimen collection(s), but very often the source of the date/time is not known, only the date/time itself.
	EffectivePeriod *Period `json:"effectivePeriod"`
	// Extensions for conclusion
	Conclusion_ext *Element `json:"_conclusion"`
	// This is a DiagnosticReport resource
	ResourceType string `json:"resourceType,omitempty"`
	// A code that classifies the clinical discipline, department or diagnostic service that created the report (e.g. cardiology, biochemistry, hematology, MRI). This is used for searching, sorting and display purposes.
	Category *CodeableConcept `json:"category"`
}

func NewDiagnosticReport() *DiagnosticReport {
	return &DiagnosticReport{ResourceType: "DiagnosticReport"}
}

// ExplanationOfBenefit_SubDetail is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_SubDetail struct {
	_ *BackboneElement
	// A service line number.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// The type of reveneu or cost center providing the product and/or service.
	Revenue *CodeableConcept `json:"revenue"`
	// Health Care Service Type Codes  to identify the classification of service or benefits.
	Category *CodeableConcept `json:"category"`
	// The number of repetitions of a service or product.
	Quantity *Quantity `json:"quantity"`
	// List of Unique Device Identifiers associated with this line item.
	Udi []*Reference `json:"udi"`
	// Extensions for factor
	Factor_ext *Element `json:"_factor"`
	// The quantity times the unit price for an addittional service or product or charge. For example, the formula: unit Quantity * unit Price (Cost per Point) * factor Number  * points = net Amount. Quantity, factor and points are assumed to be 1 if not supplied.
	Net *Money `json:"net"`
	// A list of note references to the notes provided below.
	NoteNumber []uint64 `json:"noteNumber"`
	// The adjudications results.
	Adjudication []*ExplanationOfBenefit_Adjudication `json:"adjudication"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Factor float64 `json:"factor"`
	// Extensions for noteNumber
	NoteNumber_ext []*Element `json:"_noteNumber"`
	// The type of product or service.
	Type *CodeableConcept `json:"type,omitempty"`
	// A code to indicate the Professional Service or Product supplied (eg. CTP, HCPCS,USCLS,ICD10, NCPDP,DIN,ACHI,CCI).
	Service *CodeableConcept `json:"service"`
	// Item typification or modifiers codes, eg for Oral whether the treatment is cosmetic or associated with TMJ, or for medical whether the treatment was outside the clinic or out of office hours.
	Modifier []*CodeableConcept `json:"modifier"`
	// For programs which require reson codes for the inclusion, covering, of this billed item under the program or sub-program.
	ProgramCode []*CodeableConcept `json:"programCode"`
	// The fee for an addittional service or product or charge.
	UnitPrice *Money `json:"unitPrice"`
}

// Attachment is For referring to data content defined in other formats.
type Attachment struct {
	_ *Element
	// The number of bytes of data that make up this attachment (before base64 encoding, if that is done).
	// pattern [0]|([1-9][0-9]*)
	Size uint64 `json:"size"`
	// Extensions for creation
	Creation_ext *Element `json:"_creation"`
	// Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. Includes mime type parameters such as charset where appropriate.
	// pattern [^\s]+([\s]?[^\s]+)*
	ContentType string `json:"contentType"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
	// The actual data of the attachment - a sequence of bytes. In XML, represented using base64.
	Data string `json:"data"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// The calculated hash of the data using SHA-1. Represented using base64.
	Hash string `json:"hash"`
	// A label or set of text to display in place of the data.
	Title string `json:"title"`
	// The human language of the content. The value can be any valid value according to BCP 47.
	// pattern [^\s]+([\s]?[^\s]+)*
	Language string `json:"language"`
	// Extensions for data
	Data_ext *Element `json:"_data"`
	// An alternative location where the data can be accessed.
	Url string `json:"url"`
	// Extensions for size
	Size_ext *Element `json:"_size"`
	// Extensions for contentType
	ContentType_ext *Element `json:"_contentType"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for hash
	Hash_ext *Element `json:"_hash"`
	// The date that the attachment was first created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Creation time.Time `json:"creation"`
}

// CodeSystem_Filter is A code system resource specifies a set of codes drawn from one or more code systems.
type CodeSystem_Filter struct {
	_ *BackboneElement
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// A description of how or why the filter is used.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A list of operators that can be used with the filter.
	Operator []string `json:"operator"`
	// Extensions for operator
	Operator_ext []*Element `json:"_operator"`
	// A description of what the value for the filter should be.
	Value string `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// The code that identifies this filter when it is used in the instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
}

// CompartmentDefinition_Resource is A compartment definition that defines how resources are accessed on a server.
type CompartmentDefinition_Resource struct {
	_ *BackboneElement
	// The name of a resource supported by the server.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The name of a search parameter that represents the link to the compartment. More than one may be listed because a resource may be linked to a compartment in more than one way,.
	Param []string `json:"param"`
	// Extensions for param
	Param_ext []*Element `json:"_param"`
	// Additional documentation about the resource and compartment.
	Documentation string `json:"documentation"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
}

// Group is Represents a defined collection of entities that may be discussed or acted upon collectively but which are not expected to act collectively and are not formally or legally recognized; i.e. a collection of entities that isn't an Organization.
type Group struct {
	_ *DomainResource
	// Indicates whether the record for the group is available for use or is merely being retained for historical purposes.
	Active bool `json:"active"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// A unique business identifier for this group.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for quantity
	Quantity_ext *Element `json:"_quantity"`
	// Identifies the traits shared by members of the group.
	Characteristic []*Group_Characteristic `json:"characteristic"`
	// Identifies the resource instances that are members of the group.
	Member []*Group_Member `json:"member"`
	// This is a Group resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for actual
	Actual_ext *Element `json:"_actual"`
	// A label assigned to the group for human identification and communication.
	Name string `json:"name"`
	// A count of the number of resource instances that are part of the group.
	// pattern [0]|([1-9][0-9]*)
	Quantity uint64 `json:"quantity"`
	// If true, indicates that the resource refers to a specific group of real individuals.  If false, the group defines a set of intended individuals.
	Actual bool `json:"actual"`
	// Identifies the broad classification of the kind of resources the group includes.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Provides a specific type of resource the group includes; e.g. "cow", "syringe", etc.
	Code *CodeableConcept `json:"code"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
}

func NewGroup() *Group { return &Group{ResourceType: "Group"} }

// Task is A task to be performed.
type Task struct {
	_ *DomainResource
	// A name or code (or both) briefly describing what the task involves.
	Code *CodeableConcept `json:"code"`
	// The current status of the task.
	Status string `json:"status"`
	// Extensions for priority
	Priority_ext *Element `json:"_priority"`
	// The date and time of last modification to this task.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	LastModified time.Time `json:"lastModified"`
	// Extensions for lastModified
	LastModified_ext *Element `json:"_lastModified"`
	// Individual organization or Device currently responsible for task execution.
	Owner *Reference `json:"owner"`
	// A free-text description of what is to be performed.
	Description string `json:"description"`
	// The request being actioned or the resource being manipulated by this task.
	Focus *Reference `json:"focus"`
	// A description or code indicating why this task needs to be performed.
	Reason *CodeableConcept `json:"reason"`
	// Identifies the time action was first taken against the task (start) and/or the time final action was taken against the task prior to marking it as completed (end).
	ExecutionPeriod *Period `json:"executionPeriod"`
	// The type of participant that can execute the task.
	PerformerType []*CodeableConcept `json:"performerType"`
	// Contains business-specific nuances of the business state.
	BusinessStatus *CodeableConcept `json:"businessStatus"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Free-text information captured about the task as it progresses.
	Note []*Annotation `json:"note"`
	// Links to Provenance records for past versions of this Task that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the task.
	RelevantHistory []*Reference `json:"relevantHistory"`
	// If the Task.focus is a request resource and the task is seeking fulfillment (i.e is asking for the request to be actioned), this element identifies any limitations on what parts of the referenced request should be actioned.
	Restriction *Task_Restriction `json:"restriction"`
	// Additional information that may be needed in the execution of the task.
	Input []*Task_Input `json:"input"`
	// A reference to a formal or informal definition of the task.  For example, a protocol, a step within a defined workflow definition, etc.
	DefinitionReference *Reference `json:"definitionReference"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// An explanation as to why this task is held, failed, was refused, etc.
	StatusReason *CodeableConcept `json:"statusReason"`
	// The healthcare event  (e.g. a patient and healthcare provider interaction) during which this task was created.
	Context *Reference `json:"context"`
	// A reference to a formal or informal definition of the task.  For example, a protocol, a step within a defined workflow definition, etc.
	DefinitionUri string `json:"definitionUri"`
	// BasedOn refers to a higher-level authorization that triggered the creation of the task.  It references a "request" resource such as a ProcedureRequest, MedicationRequest, ProcedureRequest, CarePlan, etc. which is distinct from the "request" resource the task is seeking to fulfil.  This latter resource is referenced by FocusOn.  For example, based on a ProcedureRequest (= BasedOn), a task is created to fulfil a procedureRequest ( = FocusOn ) to collect a specimen from a patient.
	BasedOn []*Reference `json:"basedOn"`
	// Extensions for definitionUri
	DefinitionUri_ext *Element `json:"_definitionUri"`
	// Task that this particular task is part of.
	PartOf []*Reference `json:"partOf"`
	// Indicates the "level" of actionability associated with the Task.  I.e. Is this a proposed task, a planned task, an actionable task, etc.
	// pattern [^\s]+([\s]?[^\s]+)*
	Intent string `json:"intent"`
	// Extensions for intent
	Intent_ext *Element `json:"_intent"`
	// Indicates how quickly the Task should be addressed with respect to other requests.
	// pattern [^\s]+([\s]?[^\s]+)*
	Priority string `json:"priority"`
	// The date and time this task was created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AuthoredOn time.Time `json:"authoredOn"`
	// This is a Task resource
	ResourceType string `json:"resourceType,omitempty"`
	// The business identifier for this task.
	Identifier []*Identifier `json:"identifier"`
	// The entity who benefits from the performance of the service specified in the task (e.g., the patient).
	For *Reference `json:"for"`
	// The creator of the task.
	Requester *Task_Requester `json:"requester"`
	// Outputs produced by the Task.
	Output []*Task_Output `json:"output"`
	// An identifier that links together multiple tasks and other requests that were created in the same context.
	GroupIdentifier *Identifier `json:"groupIdentifier"`
	// Extensions for authoredOn
	AuthoredOn_ext *Element `json:"_authoredOn"`
}

func NewTask() *Task { return &Task{ResourceType: "Task"} }

// TestScript_Fixture is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Fixture struct {
	_ *BackboneElement
	// Whether or not to implicitly create the fixture during setup. If true, the fixture is automatically created on each server being tested during setup, therefore no create operation is required for this fixture in the TestScript.setup section.
	Autocreate bool `json:"autocreate"`
	// Extensions for autocreate
	Autocreate_ext *Element `json:"_autocreate"`
	// Whether or not to implicitly delete the fixture during teardown. If true, the fixture is automatically deleted on each server being tested during teardown, therefore no delete operation is required for this fixture in the TestScript.teardown section.
	Autodelete bool `json:"autodelete"`
	// Extensions for autodelete
	Autodelete_ext *Element `json:"_autodelete"`
	// Reference to the resource (containing the contents of the resource needed for operations).
	Resource *Reference `json:"resource"`
}

// Distance is A length - a value with a unit that is a physical distance.
type Distance struct {
	_ *Quantity
}

// Composition_Section is A set of healthcare-related information that is assembled together into a single logical document that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. While a Composition defines the structure, it does not actually contain the content: rather the full content of a document is contained in a Bundle, of which the Composition is the first resource contained.
type Composition_Section struct {
	_ *BackboneElement
	// Specifies the order applied to the items in the section entries.
	OrderedBy *CodeableConcept `json:"orderedBy"`
	// A reference to the actual resource from which the narrative in the section is derived.
	Entry []*Reference `json:"entry"`
	// If the section is empty, why the list is empty. An empty section typically has some text explaining the empty reason.
	EmptyReason *CodeableConcept `json:"emptyReason"`
	// A human-readable narrative that contains the attested content of the section, used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative.
	Text *Narrative `json:"text"`
	// How the entry list was prepared - whether it is a working list that is suitable for being maintained on an ongoing basis, or if it represents a snapshot of a list of items from another source, or whether it is a prepared list where items may be marked as added, modified or deleted.
	// pattern [^\s]+([\s]?[^\s]+)*
	Mode string `json:"mode"`
	// A code identifying the kind of content contained within the section. This must be consistent with the section title.
	Code *CodeableConcept `json:"code"`
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// A nested sub-section within this section.
	Section []*Composition_Section `json:"section"`
	// The label for this particular section.  This will be part of the rendered content for the document, and is often used to build a table of contents.
	Title string `json:"title"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
}

// ConceptMap_Target is A statement of relationships from one set of concepts to one or more other concepts - either code systems or data elements, or classes in class models.
type ConceptMap_Target struct {
	_ *BackboneElement
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// A set of additional dependencies for this mapping to hold. This mapping is only applicable if the specified element can be resolved, and it has the specified value.
	DependsOn []*ConceptMap_DependsOn `json:"dependsOn"`
	// A set of additional outcomes from this mapping to other elements. To properly execute this mapping, the specified element must be mapped to some data element or source that is in context. The mapping may still be useful without a place for the additional data elements, but the equivalence cannot be relied on.
	Product []*ConceptMap_DependsOn `json:"product"`
	// Identity (code or path) or the element/item that the map refers to.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for display
	Display_ext *Element `json:"_display"`
	// A description of status/issues in mapping that conveys additional information not represented in  the structured data.
	Comment string `json:"comment"`
	// Extensions for equivalence
	Equivalence_ext *Element `json:"_equivalence"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The display for the code. The display is only provided to help editors when editing the concept map.
	Display string `json:"display"`
	// The equivalence between the source and target concepts (counting for the dependencies and products). The equivalence is read from target to source (e.g. the target is 'wider' than the source).
	Equivalence string `json:"equivalence"`
}

// Organization_Contact is A formally or informally recognized grouping of people or organizations formed for the purpose of achieving some form of collective action.  Includes companies, institutions, corporations, departments, community groups, healthcare practice groups, etc.
type Organization_Contact struct {
	_ *BackboneElement
	// Indicates a purpose for which the contact can be reached.
	Purpose *CodeableConcept `json:"purpose"`
	// A name associated with the contact.
	Name *HumanName `json:"name"`
	// A contact detail (e.g. a telephone number or an email address) by which the party may be contacted.
	Telecom []*ContactPoint `json:"telecom"`
	// Visiting or postal addresses for the contact.
	Address *Address `json:"address"`
}

// TestReport_Action1 is A summary of information based on the results of executing a TestScript.
type TestReport_Action1 struct {
	_ *BackboneElement
	// An operation would involve a REST request to a server.
	Operation *TestReport_Operation `json:"operation"`
	// The results of the assertion performed on the previous operations.
	Assert *TestReport_Assert `json:"assert"`
}

// AppointmentResponse is A reply to an appointment request for a patient and/or practitioner(s), such as a confirmation or rejection.
type AppointmentResponse struct {
	_ *DomainResource
	// Extensions for participantStatus
	ParticipantStatus_ext *Element `json:"_participantStatus"`
	// This is a AppointmentResponse resource
	ResourceType string `json:"resourceType,omitempty"`
	// This records identifiers associated with this appointment response concern that are defined by business processes and/ or used to refer to it when a direct URL reference to the resource itself is not appropriate.
	Identifier []*Identifier `json:"identifier"`
	// Appointment that this response is replying to.
	Appointment *Reference `json:"appointment,omitempty"`
	// Extensions for start
	Start_ext *Element `json:"_start"`
	// This may be either the same as the appointment request to confirm the details of the appointment, or alternately a new time to request a re-negotiation of the end time.
	End string `json:"end"`
	// Role of participant in the appointment.
	ParticipantType []*CodeableConcept `json:"participantType"`
	// Participation status of the participant. When the status is declined or tentative if the start/end times are different to the appointment, then these times should be interpreted as a requested time change. When the status is accepted, the times can either be the time of the appointment (as a confirmation of the time) or can be empty.
	// pattern [^\s]+([\s]?[^\s]+)*
	ParticipantStatus string `json:"participantStatus"`
	// Date/Time that the appointment is to take place, or requested new start time.
	Start string `json:"start"`
	// Extensions for end
	End_ext *Element `json:"_end"`
	// A Person, Location/HealthcareService or Device that is participating in the appointment.
	Actor *Reference `json:"actor"`
	// Additional comments about the appointment.
	Comment string `json:"comment"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
}

func NewAppointmentResponse() *AppointmentResponse {
	return &AppointmentResponse{ResourceType: "AppointmentResponse"}
}

// AuditEvent_Entity is A record of an event made for purposes of maintaining a security log. Typical uses include detection of intrusion attempts and monitoring for inappropriate usage.
type AuditEvent_Entity struct {
	_ *BackboneElement
	// Identifier for the data life-cycle stage for the entity.
	Lifecycle *Coding `json:"lifecycle"`
	// Security labels for the identified entity.
	SecurityLabel []*Coding `json:"securityLabel"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Text that describes the entity in more detail.
	Description string `json:"description"`
	// Tagged value pairs for conveying additional information about the entity.
	Detail []*AuditEvent_Detail `json:"detail"`
	// Extensions for query
	Query_ext *Element `json:"_query"`
	// Identifies a specific instance of the entity. The reference should always be version specific.
	Identifier *Identifier `json:"identifier"`
	// Identifies a specific instance of the entity. The reference should be version specific.
	Reference *Reference `json:"reference"`
	// The type of the object that was involved in this audit event.
	Type *Coding `json:"type"`
	// Code representing the role the entity played in the event being audited.
	Role *Coding `json:"role"`
	// A name of the entity in the audit event.
	Name string `json:"name"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The query parameters for a query-type entities.
	Query string `json:"query"`
}

// MessageDefinition is Defines the characteristics of a message that can be shared between systems, including the type of event that initiates the message, the content to be transmitted and what response(s), if any, are permitted.
type MessageDefinition struct {
	_ *DomainResource
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// A MessageDefinition that is superseded by this definition.
	Replaces []*Reference `json:"replaces"`
	// Indicates whether a response is required for this message.
	ResponseRequired bool `json:"responseRequired"`
	// This is a MessageDefinition resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// A legal or geographic region in which the message definition is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// The MessageDefinition that is the basis for the contents of this resource.
	Base *Reference `json:"base"`
	// Identifies a protocol or workflow that this MessageDefinition represents a step in.
	Parent []*Reference `json:"parent"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// The identifier that is used to identify this version of the message definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the message definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// A short, descriptive, user-friendly title for the message definition.
	Title string `json:"title"`
	// The status of this message definition. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A boolean value to indicate that this message definition is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// A formal identifier that is used to identify this message definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier *Identifier `json:"identifier"`
	// The date  (and optionally time) when the message definition was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the message definition changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// A copyright statement relating to the message definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the message definition.
	Copyright string `json:"copyright"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate message definition instances.
	UseContext []*UsageContext `json:"useContext"`
	// Extensions for responseRequired
	ResponseRequired_ext *Element `json:"_responseRequired"`
	// A free text natural language description of the message definition from a consumer's perspective.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for category
	Category_ext *Element `json:"_category"`
	// Indicates what types of messages may be sent as an application-level response to this message.
	AllowedResponse []*MessageDefinition_AllowedResponse `json:"allowedResponse"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// A coded identifier of a supported messaging event.
	Event *Coding `json:"event,omitempty"`
	// A natural language name identifying the message definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// Identifies the resource (or resources) that are being addressed by the event.  For example, the Encounter for an admit message or two Account records for a merge.
	Focus []*MessageDefinition_Focus `json:"focus"`
	// The name of the individual or organization that published the message definition.
	Publisher string `json:"publisher"`
	// Explaination of why this message definition is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// The impact of the content of the message.
	// pattern [^\s]+([\s]?[^\s]+)*
	Category string `json:"category"`
	// An absolute URI that is used to identify this message definition when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this message definition is (or will be) published. The URL SHOULD include the major version of the message definition. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
}

func NewMessageDefinition() *MessageDefinition {
	return &MessageDefinition{ResourceType: "MessageDefinition"}
}

// AuditEvent_Agent is A record of an event made for purposes of maintaining a security log. Typical uses include detection of intrusion attempts and monitoring for inappropriate usage.
type AuditEvent_Agent struct {
	_ *BackboneElement
	// Indicator that the user is or is not the requestor, or initiator, for the event being audited.
	Requestor bool `json:"requestor"`
	// The policy or plan that authorized the activity being recorded. Typically, a single activity may have multiple applicable policies, such as patient consent, guarantor funding, etc. The policy would also indicate the security token used.
	Policy []string `json:"policy"`
	// Type of media involved. Used when the event is about exporting/importing onto media.
	Media *Coding `json:"media"`
	// The security role that the user was acting under, that come from local codes defined by the access control security system (e.g. RBAC, ABAC) used in the local context.
	Role []*CodeableConcept `json:"role"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Where the event occurred.
	Location *Reference `json:"location"`
	// Logical network location for application activity, if the activity has a network location.
	Network *AuditEvent_Network `json:"network"`
	// Unique identifier for the user actively participating in the event.
	UserId *Identifier `json:"userId"`
	// Extensions for altId
	AltId_ext *Element `json:"_altId"`
	// Extensions for policy
	Policy_ext []*Element `json:"_policy"`
	// The reason (purpose of use), specific to this agent, that was used during the event being recorded.
	PurposeOfUse []*CodeableConcept `json:"purposeOfUse"`
	// Direct reference to a resource that identifies the agent.
	Reference *Reference `json:"reference"`
	// Human-meaningful name for the agent.
	Name string `json:"name"`
	// Extensions for requestor
	Requestor_ext *Element `json:"_requestor"`
	// Alternative agent Identifier. For a human, this should be a user identifier text string from authentication system. This identifier would be one known to a common authentication system (e.g. single sign-on), if available.
	AltId string `json:"altId"`
}

// DocumentManifest_Related is A collection of documents compiled for a purpose together with metadata that applies to the collection.
type DocumentManifest_Related struct {
	_ *BackboneElement
	// Related identifier to this DocumentManifest.  For example, Order numbers, accession numbers, XDW workflow numbers.
	Identifier *Identifier `json:"identifier"`
	// Related Resource to this DocumentManifest. For example, Order, ProcedureRequest,  Procedure, EligibilityRequest, etc.
	Ref *Reference `json:"ref"`
}

// DocumentReference_RelatesTo is A reference to a document.
type DocumentReference_RelatesTo struct {
	_ *BackboneElement
	// The type of relationship that this document has with anther document.
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The target document of this relationship.
	Target *Reference `json:"target,omitempty"`
}

// StructureMap_Input is A Map of relationships between 2 structures that can be used to transform data.
type StructureMap_Input struct {
	_ *BackboneElement
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// Documentation for this instance of data.
	Documentation string `json:"documentation"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// Name for this instance of data.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Type for this instance of data.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Mode for this instance of data.
	Mode string `json:"mode"`
}

// OperationDefinition_Binding is A formal computable definition of an operation (on the RESTful interface) or a named query (using the search interaction).
type OperationDefinition_Binding struct {
	_ *BackboneElement
	// Indicates the degree of conformance expectations associated with this binding - that is, the degree to which the provided value set must be adhered to in the instances.
	Strength string `json:"strength"`
	// Extensions for strength
	Strength_ext *Element `json:"_strength"`
	// Points to the value set or external definition (e.g. implicit value set) that identifies the set of codes to be used.
	ValueSetUri string `json:"valueSetUri"`
	// Extensions for valueSetUri
	ValueSetUri_ext *Element `json:"_valueSetUri"`
	// Points to the value set or external definition (e.g. implicit value set) that identifies the set of codes to be used.
	ValueSetReference *Reference `json:"valueSetReference"`
}

// ExplanationOfBenefit_CareTeam is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_CareTeam struct {
	_ *BackboneElement
	// Extensions for responsible
	Responsible_ext *Element `json:"_responsible"`
	// The lead, assisting or supervising practitioner and their discipline if a multidisiplinary team.
	Role *CodeableConcept `json:"role"`
	// The qualification which is applicable for this service.
	Qualification *CodeableConcept `json:"qualification"`
	// Sequence of careteam which serves to order and provide a link.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// The members of the team who provided the overall service.
	Provider *Reference `json:"provider,omitempty"`
	// The practitioner who is billing and responsible for the claimed services rendered to the patient.
	Responsible bool `json:"responsible"`
}

// ClaimResponse_Insurance is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse_Insurance struct {
	_ *BackboneElement
	// A service line item.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// Reference to the program or plan identification, underwriter or payor.
	Coverage *Reference `json:"coverage,omitempty"`
	// The contract number of a business agreement which describes the terms and conditions.
	BusinessArrangement string `json:"businessArrangement"`
	// Extensions for businessArrangement
	BusinessArrangement_ext *Element `json:"_businessArrangement"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// The instance number of the Coverage which is the focus for adjudication. The Coverage against which the claim is to be adjudicated.
	Focal bool `json:"focal"`
	// Extensions for focal
	Focal_ext *Element `json:"_focal"`
	// A list of references from the Insurer to which these services pertain.
	PreAuthRef []string `json:"preAuthRef"`
	// Extensions for preAuthRef
	PreAuthRef_ext []*Element `json:"_preAuthRef"`
	// The Coverages adjudication details.
	ClaimResponse *Reference `json:"claimResponse"`
}

// Flag is Prospective warnings of potential issues when providing care to the patient.
type Flag struct {
	_ *DomainResource
	// The person, organization or device that created the flag.
	Author *Reference `json:"author"`
	// Supports basic workflow.
	Status string `json:"status"`
	// Allows an flag to be divided into different categories like clinical, administrative etc. Intended to be used as a means of filtering which flags are displayed to particular user or in a given context.
	Category *CodeableConcept `json:"category"`
	// The coded value or textual component of the flag to display to the user.
	Code *CodeableConcept `json:"code,omitempty"`
	// The patient, location, group , organization , or practitioner, etc. this is about record this flag is associated with.
	Subject *Reference `json:"subject,omitempty"`
	// This alert is only relevant during the encounter.
	Encounter *Reference `json:"encounter"`
	// This is a Flag resource
	ResourceType string `json:"resourceType,omitempty"`
	// Identifier assigned to the flag for external use (outside the FHIR environment).
	Identifier []*Identifier `json:"identifier"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The period of time from the activation of the flag to inactivation of the flag. If the flag is active, the end of the period should be unspecified.
	Period *Period `json:"period"`
}

func NewFlag() *Flag { return &Flag{ResourceType: "Flag"} }

// MedicationAdministration_Performer is Describes the event of a patient consuming or otherwise being administered a medication.  This may be as simple as swallowing a tablet or it may be a long running infusion.  Related resources tie this event to the authorizing prescription, and the specific encounter between patient and health care practitioner.
type MedicationAdministration_Performer struct {
	_ *BackboneElement
	// The device, practitioner, etc. who performed the action.
	Actor *Reference `json:"actor,omitempty"`
	// The organization the device or practitioner was acting on behalf of.
	OnBehalfOf *Reference `json:"onBehalfOf"`
}

// PractitionerRole_NotAvailable is A specific set of Roles/Locations/specialties/services that a practitioner may perform at an organization for a period of time.
type PractitionerRole_NotAvailable struct {
	_ *BackboneElement
	// The reason that can be presented to the user as to why this time is not available.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Service is not available (seasonally or for a public holiday) from this date.
	During *Period `json:"during"`
}

// RequestGroup_Condition is A group of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
type RequestGroup_Condition struct {
	_ *BackboneElement
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The media type of the language for the expression.
	Language string `json:"language"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
	// An expression that returns true or false, indicating whether or not the condition is satisfied.
	Expression string `json:"expression"`
	// Extensions for expression
	Expression_ext *Element `json:"_expression"`
	// The kind of condition.
	// pattern [^\s]+([\s]?[^\s]+)*
	Kind string `json:"kind"`
	// Extensions for kind
	Kind_ext *Element `json:"_kind"`
	// A brief, natural language description of the condition that effectively communicates the intended semantics.
	Description string `json:"description"`
}

// Immunization_Practitioner is Describes the event of a patient being administered a vaccination or a record of a vaccination as reported by a patient, a clinician or another party and may include vaccine reaction information and what vaccination protocol was followed.
type Immunization_Practitioner struct {
	_ *BackboneElement
	// Describes the type of performance (e.g. ordering provider, administering provider, etc.).
	Role *CodeableConcept `json:"role"`
	// The device, practitioner, etc. who performed the action.
	Actor *Reference `json:"actor,omitempty"`
}

// RelatedArtifact is Related artifacts such as additional documentation, justification, or bibliographic references.
type RelatedArtifact struct {
	_ *Element
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// A bibliographic citation for the related artifact. This text SHOULD be formatted according to an accepted citation format.
	Citation string `json:"citation"`
	// The document being referenced, represented as an attachment. This is exclusive with the resource element.
	Document *Attachment `json:"document"`
	// The related resource, such as a library, value set, profile, or other knowledge resource.
	Resource *Reference `json:"resource"`
	// The type of relationship to the related artifact.
	Type string `json:"type"`
	// A brief description of the document or knowledge resource being referenced, suitable for display to a consumer.
	Display string `json:"display"`
	// Extensions for display
	Display_ext *Element `json:"_display"`
	// Extensions for citation
	Citation_ext *Element `json:"_citation"`
	// A url for the artifact that can be followed to access the actual content.
	Url string `json:"url"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
}

// DocumentReference_Related is A reference to a document.
type DocumentReference_Related struct {
	_ *BackboneElement
	// Related identifier to this DocumentReference. If both id and ref are present they shall refer to the same thing.
	Identifier *Identifier `json:"identifier"`
	// Related Resource to this DocumentReference. If both id and ref are present they shall refer to the same thing.
	Ref *Reference `json:"ref"`
}

// Linkage is Identifies two or more records (resource instances) that are referring to the same real-world "occurrence".
type Linkage struct {
	_ *DomainResource
	// This is a Linkage resource
	ResourceType string `json:"resourceType,omitempty"`
	// Indicates whether the asserted set of linkages are considered to be "in effect".
	Active bool `json:"active"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
	// Identifies the user or organization responsible for asserting the linkages and who establishes the context for evaluating the nature of each linkage.
	Author *Reference `json:"author"`
	// Identifies one of the records that is considered to refer to the same real-world occurrence as well as how the items hould be evaluated within the collection of linked items.
	Item []*Linkage_Item `json:"item,omitempty"`
}

func NewLinkage() *Linkage { return &Linkage{ResourceType: "Linkage"} }

// TestScript_Link is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Link struct {
	_ *BackboneElement
	// URL to a particular requirement or feature within the FHIR specification.
	Url string `json:"url"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Short description of the link.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
}

// Claim is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim struct {
	_ *DomainResource
	// A finer grained suite of claim subtype codes which may convey Inpatient vs Outpatient and/or a specialty service. In the US the BillType.
	SubType []*CodeableConcept `json:"subType"`
	// Facility where the services were provided.
	Facility *Reference `json:"facility"`
	// Ordered list of patient procedures performed to support the adjudication.
	Procedure []*Claim_Procedure `json:"procedure"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for use
	Use_ext *Element `json:"_use"`
	// Other claims which are related to this claim such as prior claim versions or for related services.
	Related []*Claim_Related `json:"related"`
	// The referral resource which lists the date, practitioner, reason and other supporting information.
	Referral *Reference `json:"referral"`
	// The start and optional end dates of when the patient was precluded from working due to the treatable condition(s).
	EmploymentImpacted *Period `json:"employmentImpacted"`
	// First tier of goods and services.
	Item []*Claim_Item `json:"item"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// Person who created the invoice/claim/pre-determination or pre-authorization.
	Enterer *Reference `json:"enterer"`
	// The category of claim, eg, oral, pharmacy, vision, insitutional, professional.
	Type *CodeableConcept `json:"type"`
	// The date when the enclosed suite of services were performed or completed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// The Insurer who is target of the request.
	Insurer *Reference `json:"insurer"`
	// The provider which is responsible for the bill, claim pre-determination, pre-authorization.
	Provider *Reference `json:"provider"`
	// In the case of a Pre-Determination/Pre-Authorization the provider may request that funds in the amount of the expected Benefit be reserved ('Patient' or 'Provider') to pay for the Benefits determined on the subsequent claim(s). 'None' explicitly indicates no funds reserving is requested.
	FundsReserve *CodeableConcept `json:"fundsReserve"`
	// List of patient diagnosis for which care is sought.
	Diagnosis []*Claim_Diagnosis `json:"diagnosis"`
	// This is a Claim resource
	ResourceType string `json:"resourceType,omitempty"`
	// The business identifier for the instance: claim number, pre-determination or pre-authorization number.
	Identifier []*Identifier `json:"identifier"`
	// Patient Resource.
	Patient *Reference `json:"patient"`
	// The organization which is responsible for the bill, claim pre-determination, pre-authorization.
	Organization *Reference `json:"organization"`
	// Immediate (STAT), best effort (NORMAL), deferred (DEFER).
	Priority *CodeableConcept `json:"priority"`
	// Original prescription which has been superceded by this prescription to support the dispensing of pharmacy services, medications or products. For example, a physician may prescribe a medication which the pharmacy determines is contraindicated, or for which the patient has an intolerance, and therefor issues a new precription for an alternate medication which has the same theraputic intent. The prescription from the pharmacy becomes the 'prescription' and that from the physician becomes the 'original prescription'.
	OriginalPrescription *Reference `json:"originalPrescription"`
	// Complete (Bill or Claim), Proposed (Pre-Authorization), Exploratory (Pre-determination).
	Use string `json:"use"`
	// The billable period for which charges are being submitted.
	BillablePeriod *Period `json:"billablePeriod"`
	// Prescription to support the dispensing of Pharmacy or Vision products.
	Prescription *Reference `json:"prescription"`
	// An accident which resulted in the need for healthcare services.
	Accident *Claim_Accident `json:"accident"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// The party to be reimbursed for the services.
	Payee *Claim_Payee `json:"payee"`
	// The members of the team who provided the overall service as well as their role and whether responsible and qualifications.
	CareTeam []*Claim_CareTeam `json:"careTeam"`
	// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues. Often there are mutiple jurisdiction specific valuesets which are required.
	Information []*Claim_Information `json:"information"`
	// Financial instrument by which payment information for health care.
	Insurance []*Claim_Insurance `json:"insurance"`
	// The start and optional end dates of when the patient was confined to a treatment center.
	Hospitalization *Period `json:"hospitalization"`
	// The total value of the claim.
	Total *Money `json:"total"`
}

func NewClaim() *Claim { return &Claim{ResourceType: "Claim"} }

// EpisodeOfCare_Diagnosis is An association between a patient and an organization / healthcare provider(s) during which time encounters may occur. The managing organization assumes a level of responsibility for the patient during this time.
type EpisodeOfCare_Diagnosis struct {
	_ *BackboneElement
	// Role that this diagnosis has within the episode of care (e.g. admission, billing, discharge ...).
	Role *CodeableConcept `json:"role"`
	// Ranking of the diagnosis (for each role type).
	// pattern [1-9][0-9]*
	Rank uint64 `json:"rank"`
	// Extensions for rank
	Rank_ext *Element `json:"_rank"`
	// A list of conditions/problems/diagnoses that this episode of care is intended to be providing care for.
	Condition *Reference `json:"condition,omitempty"`
}

// ExplanationOfBenefit_Payment is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Payment struct {
	_ *BackboneElement
	// Adjustment to the payment of this transaction which is not related to adjudication of this transaction.
	Adjustment *Money `json:"adjustment"`
	// Reason for the payment adjustment.
	AdjustmentReason *CodeableConcept `json:"adjustmentReason"`
	// Estimated payment date.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Payable less any payment adjustment.
	Amount *Money `json:"amount"`
	// Payment identifer.
	Identifier *Identifier `json:"identifier"`
	// Whether this represents partial or complete payment of the claim.
	Type *CodeableConcept `json:"type"`
}

// ImplementationGuide_Global is A set of rules of how FHIR is used to solve a particular problem. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuide_Global struct {
	_ *BackboneElement
	// The type of resource that all instances must conform to.
	// pattern [^\s]+([\s]?[^\s]+)*
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// A reference to the profile that all instances must conform to.
	Profile *Reference `json:"profile,omitempty"`
}

// NutritionOrder_EnteralFormula is A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrder_EnteralFormula struct {
	_ *BackboneElement
	// Extensions for baseFormulaProductName
	BaseFormulaProductName_ext *Element `json:"_baseFormulaProductName"`
	// Indicates the type of modular component such as protein, carbohydrate, fat or fiber to be provided in addition to or mixed with the base formula.
	AdditiveType *CodeableConcept `json:"additiveType"`
	// Extensions for additiveProductName
	AdditiveProductName_ext *Element `json:"_additiveProductName"`
	// Formula administration instructions as structured data.  This repeating structure allows for changing the administration rate or volume over time for both bolus and continuous feeding.  An example of this would be an instruction to increase the rate of continuous feeding every 2 hours.
	Administration []*NutritionOrder_Administration `json:"administration"`
	// The maximum total quantity of formula that may be administered to a subject over the period of time, e.g. 1440 mL over 24 hours.
	MaxVolumeToDeliver *Quantity `json:"maxVolumeToDeliver"`
	// The type of enteral or infant formula such as an adult standard formula with fiber or a soy-based infant formula.
	BaseFormulaType *CodeableConcept `json:"baseFormulaType"`
	// The product or brand name of the enteral or infant formula product such as "ACME Adult Standard Formula".
	BaseFormulaProductName string `json:"baseFormulaProductName"`
	// The product or brand name of the type of modular component to be added to the formula.
	AdditiveProductName string `json:"additiveProductName"`
	// The amount of energy (calories) that the formula should provide per specified volume, typically per mL or fluid oz.  For example, an infant may require a formula that provides 24 calories per fluid ounce or an adult may require an enteral formula that provides 1.5 calorie/mL.
	CaloricDensity *Quantity `json:"caloricDensity"`
	// The route or physiological path of administration into the patient's gastrointestinal  tract for purposes of providing the formula feeding, e.g. nasogastric tube.
	RouteofAdministration *CodeableConcept `json:"routeofAdministration"`
	// Free text formula administration, feeding instructions or additional instructions or information.
	AdministrationInstruction string `json:"administrationInstruction"`
	// Extensions for administrationInstruction
	AdministrationInstruction_ext *Element `json:"_administrationInstruction"`
}

// ReferralRequest is Used to record and send details about a request for referral service or transfer of a patient to the care of another provider or provider organization.
type ReferralRequest struct {
	_ *DomainResource
	// This is a ReferralRequest resource
	ResourceType string `json:"resourceType,omitempty"`
	// An indication of the urgency of referral (or where applicable the type of transfer of care) request.
	// pattern [^\s]+([\s]?[^\s]+)*
	Priority string `json:"priority"`
	// The patient who is the subject of a referral or transfer of care request.
	Subject *Reference `json:"subject,omitempty"`
	// Extensions for priority
	Priority_ext *Element `json:"_priority"`
	// Indication of the clinical domain or discipline to which the referral or transfer of care request is sent.  For example: Cardiology Gastroenterology Diabetology.
	Specialty *CodeableConcept `json:"specialty"`
	// The healthcare provider(s) or provider organization(s) who/which is to receive the referral/transfer of care request.
	Recipient []*Reference `json:"recipient"`
	// The period of time within which the services identified in the referral/transfer of care is specified or required to occur.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OccurrenceDateTime time.Time `json:"occurrenceDateTime"`
	// Any additional (administrative, financial or clinical) information required to support request for referral or transfer of care.  For example: Presenting problems/chief complaints Medical History Family History Alerts Allergy/Intolerance and Adverse Reactions Medications Observations/Assessments (may include cognitive and fundtional assessments) Diagnostic Reports Care Plan.
	SupportingInfo []*Reference `json:"supportingInfo"`
	// Business identifier that uniquely identifies the referral/care transfer request instance.
	Identifier []*Identifier `json:"identifier"`
	// The business identifier of the logical "grouping" request/order that this referral is a part of.
	GroupIdentifier *Identifier `json:"groupIdentifier"`
	// The status of the authorization/intention reflected by the referral request record.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// An indication of the type of referral (or where applicable the type of transfer of care) request.
	Type *CodeableConcept `json:"type"`
	// The encounter at which the request for referral or transfer of care is initiated.
	Context *Reference `json:"context"`
	// Extensions for occurrenceDateTime
	OccurrenceDateTime_ext *Element `json:"_occurrenceDateTime"`
	// The individual who initiated the request and has responsibility for its activation.
	Requester *ReferralRequest_Requester `json:"requester"`
	// Comments made about the referral request by any of the participants.
	Note []*Annotation `json:"note"`
	// A protocol, guideline, orderset or other definition that is adhered to in whole or in part by this request.
	Definition []*Reference `json:"definition"`
	// Extensions for authoredOn
	AuthoredOn_ext *Element `json:"_authoredOn"`
	// Indicates another resource whose existence justifies this request.
	ReasonReference []*Reference `json:"reasonReference"`
	// The reason element gives a short description of why the referral is being made, the description expands on this to support a more complete clinical summary.
	Description string `json:"description"`
	// Links to Provenance records for past versions of this resource or fulfilling request or event resources that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the resource.
	RelevantHistory []*Reference `json:"relevantHistory"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for intent
	Intent_ext *Element `json:"_intent"`
	// Description of clinical condition indicating why referral/transfer of care is requested.  For example:  Pathological Anomalies, Disabled (physical or mental),  Behavioral Management.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Indicates any plans, proposals or orders that this request is intended to satisfy - in whole or in part.
	BasedOn []*Reference `json:"basedOn"`
	// Completed or terminated request(s) whose function is taken by this new request.
	Replaces []*Reference `json:"replaces"`
	// Distinguishes the "level" of authorization/demand implicit in this request.
	// pattern [^\s]+([\s]?[^\s]+)*
	Intent string `json:"intent"`
	// The service(s) that is/are requested to be provided to the patient.  For example: cardiac pacemaker insertion.
	ServiceRequested []*CodeableConcept `json:"serviceRequested"`
	// The period of time within which the services identified in the referral/transfer of care is specified or required to occur.
	OccurrencePeriod *Period `json:"occurrencePeriod"`
	// Date/DateTime of creation for draft requests and date of activation for active requests.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AuthoredOn time.Time `json:"authoredOn"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
}

func NewReferralRequest() *ReferralRequest { return &ReferralRequest{ResourceType: "ReferralRequest"} }

// RequestGroup is A group of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
type RequestGroup struct {
	_ *DomainResource
	// Indicates how quickly the request should be addressed with respect to other requests.
	// pattern [^\s]+([\s]?[^\s]+)*
	Priority string `json:"priority"`
	// The subject for which the request group was created.
	Subject *Reference `json:"subject"`
	// Extensions for authoredOn
	AuthoredOn_ext *Element `json:"_authoredOn"`
	// Provides a reference to the author of the request group.
	Author *Reference `json:"author"`
	// Indicates the reason the request group was created. This is typically provided as a parameter to the evaluation and echoed by the service, although for some use cases, such as subscription- or event-based scenarios, it may provide an indication of the cause for the response.
	ReasonCodeableConcept *CodeableConcept `json:"reasonCodeableConcept"`
	// The actions, if any, produced by the evaluation of the artifact.
	Action []*RequestGroup_Action `json:"action"`
	// A protocol, guideline, orderset or other definition that is adhered to in whole or in part by this request.
	Definition []*Reference `json:"definition"`
	// Indicates the level of authority/intentionality associated with the request and where the request fits into the workflow chain.
	// pattern [^\s]+([\s]?[^\s]+)*
	Intent string `json:"intent"`
	// Extensions for priority
	Priority_ext *Element `json:"_priority"`
	// Indicates the reason the request group was created. This is typically provided as a parameter to the evaluation and echoed by the service, although for some use cases, such as subscription- or event-based scenarios, it may provide an indication of the cause for the response.
	ReasonReference *Reference `json:"reasonReference"`
	// Provides a mechanism to communicate additional information about the response.
	Note []*Annotation `json:"note"`
	// The current state of the request. For request groups, the status reflects the status of all the requests in the group.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// A plan, proposal or order that is fulfilled in whole or in part by this request.
	BasedOn []*Reference `json:"basedOn"`
	// Completed or terminated request(s) whose function is taken by this new request.
	Replaces []*Reference `json:"replaces"`
	// A shared identifier common to all requests that were authorized more or less simultaneously by a single author, representing the identifier of the requisition, prescription or similar form.
	GroupIdentifier *Identifier `json:"groupIdentifier"`
	// Extensions for intent
	Intent_ext *Element `json:"_intent"`
	// Indicates when the request group was created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AuthoredOn time.Time `json:"authoredOn"`
	// This is a RequestGroup resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Describes the context of the request group, if any.
	Context *Reference `json:"context"`
	// Allows a service to provide a unique, business identifier for the request.
	Identifier []*Identifier `json:"identifier"`
}

func NewRequestGroup() *RequestGroup { return &RequestGroup{ResourceType: "RequestGroup"} }

// CapabilityStatement_Certificate is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Certificate struct {
	_ *BackboneElement
	// Mime type for a certificate.
	// pattern [^\s]+([\s]?[^\s]+)*
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Actual certificate.
	Blob string `json:"blob"`
	// Extensions for blob
	Blob_ext *Element `json:"_blob"`
}

// Money is An amount of economic utility in some recognized currency.
type Money struct {
	_ *Quantity
}

// CapabilityStatement_Rest is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Rest struct {
	_ *BackboneElement
	// An absolute URI which is a reference to the definition of a compartment that the system supports. The reference is to a CompartmentDefinition resource by its canonical URL .
	Compartment []string `json:"compartment"`
	// Information about the system's restful capabilities that apply across all applications, such as security.
	Documentation string `json:"documentation"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// A specification of the restful capabilities of the solution for a specific resource type.
	Resource []*CapabilityStatement_Resource `json:"resource"`
	// A specification of restful operations supported by the system.
	Interaction []*CapabilityStatement_Interaction1 `json:"interaction"`
	// Search parameters that are supported for searching all resources for implementations to support and/or make use of - either references to ones defined in the specification, or additional ones defined for/by the implementation.
	SearchParam []*CapabilityStatement_SearchParam `json:"searchParam"`
	// Identifies whether this portion of the statement is describing the ability to initiate or receive restful operations.
	Mode string `json:"mode"`
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// Information about security implementation from an interface perspective - what a client needs to know.
	Security *CapabilityStatement_Security `json:"security"`
	// Definition of an operation or a named query together with its parameters and their meaning and type.
	Operation []*CapabilityStatement_Operation `json:"operation"`
	// Extensions for compartment
	Compartment_ext []*Element `json:"_compartment"`
}

// DiagnosticReport_Image is The findings and interpretation of diagnostic  tests performed on patients, groups of patients, devices, and locations, and/or specimens derived from these. The report includes clinical context such as requesting and provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports.
type DiagnosticReport_Image struct {
	_ *BackboneElement
	// A comment about the image. Typically, this is used to provide an explanation for why the image is included, or to draw the viewer's attention to important features.
	Comment string `json:"comment"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// Reference to the image source.
	Link *Reference `json:"link,omitempty"`
}

// Claim_Accident is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_Accident struct {
	_ *BackboneElement
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Type of accident: work, auto, etc.
	Type *CodeableConcept `json:"type"`
	// Accident Place.
	LocationAddress *Address `json:"locationAddress"`
	// Accident Place.
	LocationReference *Reference `json:"locationReference"`
	// Date of an accident which these services are addressing.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	Date time.Time `json:"date"`
}

// Condition_Stage is A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
type Condition_Stage struct {
	_ *BackboneElement
	// Reference to a formal record of the evidence on which the staging assessment is based.
	Assessment []*Reference `json:"assessment"`
	// A simple summary of the stage such as "Stage 3". The determination of the stage is disease-specific.
	Summary *CodeableConcept `json:"summary"`
}

// Contract_Legal is A formal agreement between parties regarding the conduct of business, exchange of information or other matters.
type Contract_Legal struct {
	_ *BackboneElement
	// Contract legal text in human renderable form.
	ContentAttachment *Attachment `json:"contentAttachment"`
	// Contract legal text in human renderable form.
	ContentReference *Reference `json:"contentReference"`
}

// DocumentReference is A reference to a document.
type DocumentReference struct {
	_ *DomainResource
	// This is a DocumentReference resource
	ResourceType string `json:"resourceType,omitempty"`
	// The status of this document reference.
	Status string `json:"status"`
	// Identifies who is responsible for adding the information to the document.
	Author []*Reference `json:"author"`
	// Which person or organization authenticates that this document is valid.
	Authenticator *Reference `json:"authenticator"`
	// Identifies the organization or group who is responsible for ongoing maintenance of and access to the document.
	Custodian *Reference `json:"custodian"`
	// Relationships that this document has with other document references that already exist.
	RelatesTo []*DocumentReference_RelatesTo `json:"relatesTo"`
	// A set of Security-Tag codes specifying the level of privacy/security of the Document. Note that DocumentReference.meta.security contains the security labels of the "reference" to the document, while DocumentReference.securityLabel contains a snapshot of the security labels on the document the reference refers to.
	SecurityLabel []*CodeableConcept `json:"securityLabel"`
	// The clinical context in which the document was prepared.
	Context *DocumentReference_Context `json:"context"`
	// Other identifiers associated with the document, including version independent identifiers.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The status of the underlying document.
	// pattern [^\s]+([\s]?[^\s]+)*
	DocStatus string `json:"docStatus"`
	// Who or what the document is about. The document can be about a person, (patient or healthcare practitioner), a device (e.g. a machine) or even a group of subjects (such as a document about a herd of farm animals, or a set of patients that share a common exposure).
	Subject *Reference `json:"subject"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A categorization for the type of document referenced - helps for indexing and searching. This may be implied by or derived from the code specified in the DocumentReference.type.
	Class *CodeableConcept `json:"class"`
	// Document identifier as assigned by the source of the document. This identifier is specific to this version of the document. This unique identifier may be used elsewhere to identify this version of the document.
	MasterIdentifier *Identifier `json:"masterIdentifier"`
	// Extensions for docStatus
	DocStatus_ext *Element `json:"_docStatus"`
	// Specifies the particular kind of document referenced  (e.g. History and Physical, Discharge Summary, Progress Note). This usually equates to the purpose of making the document referenced.
	Type *CodeableConcept `json:"type,omitempty"`
	// When the document was created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// When the document reference was created.
	Indexed string `json:"indexed"`
	// Extensions for indexed
	Indexed_ext *Element `json:"_indexed"`
	// Human-readable description of the source document. This is sometimes known as the "title".
	Description string `json:"description"`
	// The document and format referenced. There may be multiple content element repetitions, each with a different format.
	Content []*DocumentReference_Content `json:"content,omitempty"`
}

func NewDocumentReference() *DocumentReference {
	return &DocumentReference{ResourceType: "DocumentReference"}
}

// GraphDefinition_Compartment is A formal computable definition of a graph of resources - that is, a coherent set of resources that form a graph by following references. The Graph Definition resource defines a set and makes rules about the set.
type GraphDefinition_Compartment struct {
	_ *BackboneElement
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// identical | matching | different | no-rule | custom.
	Rule string `json:"rule"`
	// Extensions for rule
	Rule_ext *Element `json:"_rule"`
	// Custom rule, as a FHIRPath expression.
	Expression string `json:"expression"`
	// Extensions for expression
	Expression_ext *Element `json:"_expression"`
	// Documentation for FHIRPath expression.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Identifies the compartment.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
}

// SupplyRequest_OrderedItem is A record of a request for a medication, substance or device used in the healthcare setting.
type SupplyRequest_OrderedItem struct {
	_ *BackboneElement
	// The item that is requested to be supplied. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept"`
	// The item that is requested to be supplied. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	ItemReference *Reference `json:"itemReference"`
	// The amount that is being ordered of the indicated item.
	Quantity *Quantity `json:"quantity,omitempty"`
}

// Account_Guarantor is A financial tool for tracking value accrued for a particular purpose.  In the healthcare field, used to track charges for a patient, cost centers, etc.
type Account_Guarantor struct {
	_ *BackboneElement
	// The entity who is responsible.
	Party *Reference `json:"party,omitempty"`
	// A guarantor may be placed on credit hold or otherwise have their role temporarily suspended.
	OnHold bool `json:"onHold"`
	// Extensions for onHold
	OnHold_ext *Element `json:"_onHold"`
	// The timeframe during which the guarantor accepts responsibility for the account.
	Period *Period `json:"period"`
}

// Immunization_Reaction is Describes the event of a patient being administered a vaccination or a record of a vaccination as reported by a patient, a clinician or another party and may include vaccine reaction information and what vaccination protocol was followed.
type Immunization_Reaction struct {
	_ *BackboneElement
	// Date of reaction to the immunization.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Details of the reaction.
	Detail *Reference `json:"detail"`
	// Self-reported indicator.
	Reported bool `json:"reported"`
	// Extensions for reported
	Reported_ext *Element `json:"_reported"`
}

// ConceptMap_DependsOn is A statement of relationships from one set of concepts to one or more other concepts - either code systems or data elements, or classes in class models.
type ConceptMap_DependsOn struct {
	_ *BackboneElement
	// Extensions for property
	Property_ext *Element `json:"_property"`
	// An absolute URI that identifies the code system of the dependency code (if the source/dependency is a value set that crosses code systems).
	System string `json:"system"`
	// Extensions for system
	System_ext *Element `json:"_system"`
	// Identity (code or path) or the element/item/ValueSet that the map depends on / refers to.
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The display for the code. The display is only provided to help editors when editing the concept map.
	Display string `json:"display"`
	// Extensions for display
	Display_ext *Element `json:"_display"`
	// A reference to an element that holds a coded value that corresponds to a code system property. The idea is that the information model carries an element somwhere that is labeled to correspond with a code system property.
	Property string `json:"property"`
}

// ProcessRequest is This resource provides the target, request and response, and action details for an action to be performed by the target on or about existing resources.
type ProcessRequest struct {
	_ *DomainResource
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The date when this resource was created.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// Extensions for nullify
	Nullify_ext *Element `json:"_nullify"`
	// Extensions for include
	Include_ext []*Element `json:"_include"`
	// Extensions for exclude
	Exclude_ext []*Element `json:"_exclude"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Extensions for action
	Action_ext *Element `json:"_action"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// Reference of resource which is the target or subject of this action.
	Request *Reference `json:"request"`
	// Names of resource types to include.
	Include []string `json:"include"`
	// Names of resource types to exclude.
	Exclude []string `json:"exclude"`
	// A period of time during which the fulfilling resources would have been created.
	Period *Period `json:"period"`
	// The type of processing action being requested, for example Reversal, Readjudication, StatusRequest,PendedRequest.
	Action string `json:"action"`
	// The organization which is responsible for the action speccified in this request.
	Organization *Reference `json:"organization"`
	// If true remove all history excluding audit.
	Nullify bool `json:"nullify"`
	// Extensions for reference
	Reference_ext *Element `json:"_reference"`
	// The practitioner who is responsible for the action specified in this request.
	Provider *Reference `json:"provider"`
	// The ProcessRequest business identifier.
	Identifier []*Identifier `json:"identifier"`
	// The organization which is the target of the request.
	Target *Reference `json:"target"`
	// Reference of a prior response to resource which is the target or subject of this action.
	Response *Reference `json:"response"`
	// A reference to supply which authenticates the process.
	Reference string `json:"reference"`
	// List of top level items to be re-adjudicated, if none specified then the entire submission is re-adjudicated.
	Item []*ProcessRequest_Item `json:"item"`
	// This is a ProcessRequest resource
	ResourceType string `json:"resourceType,omitempty"`
}

func NewProcessRequest() *ProcessRequest { return &ProcessRequest{ResourceType: "ProcessRequest"} }

// TestScript_Assert is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Assert struct {
	_ *BackboneElement
	// Extensions for label
	Label_ext *Element `json:"_label"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The request method or HTTP operation code to compare against that used by the client system under test.
	RequestMethod string `json:"requestMethod"`
	// The value of the HTTP response code to be tested.
	ResponseCode string `json:"responseCode"`
	// Id of the source fixture used as the contents to be evaluated by either the "source/expression" or "sourceId/path" definition.
	CompareToSourceId string `json:"compareToSourceId"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// Extensions for responseCode
	ResponseCode_ext *Element `json:"_responseCode"`
	// The ID of the Profile to validate against.
	// pattern [A-Za-z0-9\-\.]{1,64}
	ValidateProfileId string `json:"validateProfileId"`
	// The direction to use for the assertion.
	Direction string `json:"direction"`
	// XPath or JSONPath expression to evaluate against the source fixture. When compareToSourceId is defined, either compareToSourceExpression or compareToSourcePath must be defined, but not both.
	CompareToSourcePath string `json:"compareToSourcePath"`
	// Extensions for compareToSourcePath
	CompareToSourcePath_ext *Element `json:"_compareToSourcePath"`
	// Extensions for expression
	Expression_ext *Element `json:"_expression"`
	// Extensions for resource
	Resource_ext *Element `json:"_resource"`
	// Whether or not the test execution will produce a warning only on error for this assert.
	WarningOnly bool `json:"warningOnly"`
	// Extensions for contentType
	ContentType_ext *Element `json:"_contentType"`
	// The ID of a fixture.  Asserts that the response contains at a minimum the fixture specified by minimumId.
	MinimumId string `json:"minimumId"`
	// Extensions for requestURL
	RequestURL_ext *Element `json:"_requestURL"`
	// Fixture to evaluate the XPath/JSONPath expression or the headerField  against.
	// pattern [A-Za-z0-9\-\.]{1,64}
	SourceId string `json:"sourceId"`
	// Extensions for sourceId
	SourceId_ext *Element `json:"_sourceId"`
	// Extensions for operator
	Operator_ext *Element `json:"_operator"`
	// Extensions for requestMethod
	RequestMethod_ext *Element `json:"_requestMethod"`
	// The type of the resource.  See http://build.fhir.org/resourcelist.html.
	// pattern [^\s]+([\s]?[^\s]+)*
	Resource string `json:"resource"`
	// The TestScript.rule this assert will evaluate.
	Rule *TestScript_Rule2 `json:"rule"`
	// Extensions for validateProfileId
	ValidateProfileId_ext *Element `json:"_validateProfileId"`
	// Whether or not the test execution performs validation on the bundle navigation links.
	NavigationLinks bool `json:"navigationLinks"`
	// Extensions for response
	Response_ext *Element `json:"_response"`
	// The value to compare to.
	Value string `json:"value"`
	// The fluentpath expression to evaluate against the source fixture. When compareToSourceId is defined, either compareToSourceExpression or compareToSourcePath must be defined, but not both.
	CompareToSourceExpression string `json:"compareToSourceExpression"`
	// Extensions for compareToSourceExpression
	CompareToSourceExpression_ext *Element `json:"_compareToSourceExpression"`
	// The content-type or mime-type to use for RESTful operation in the 'Content-Type' header.
	ContentType string `json:"contentType"`
	// The HTTP header field name e.g. 'Location'.
	HeaderField string `json:"headerField"`
	// Extensions for minimumId
	MinimumId_ext *Element `json:"_minimumId"`
	// The value to use in a comparison against the request URL path string.
	RequestURL string `json:"requestURL"`
	// okay | created | noContent | notModified | bad | forbidden | notFound | methodNotAllowed | conflict | gone | preconditionFailed | unprocessable.
	Response string `json:"response"`
	// Extensions for warningOnly
	WarningOnly_ext *Element `json:"_warningOnly"`
	// The description would be used by test engines for tracking and reporting purposes.
	Description string `json:"description"`
	// The fluentpath expression to be evaluated against the request or response message contents - HTTP headers and payload.
	Expression string `json:"expression"`
	// Extensions for navigationLinks
	NavigationLinks_ext *Element `json:"_navigationLinks"`
	// The operator type defines the conditional behavior of the assert. If not defined, the default is equals.
	Operator string `json:"operator"`
	// The XPath or JSONPath expression to be evaluated against the fixture representing the response received from server.
	Path string `json:"path"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// The label would be used for tracking/logging purposes by test engines.
	Label string `json:"label"`
	// Extensions for direction
	Direction_ext *Element `json:"_direction"`
	// Extensions for compareToSourceId
	CompareToSourceId_ext *Element `json:"_compareToSourceId"`
	// Extensions for headerField
	HeaderField_ext *Element `json:"_headerField"`
	// The TestScript.ruleset this assert will evaluate.
	Ruleset *TestScript_Ruleset1 `json:"ruleset"`
}

// MedicationDispense is Indicates that a medication product is to be or has been dispensed for a named person/patient.  This includes a description of the medication product (supply) provided and the instructions for administering the medication.  The medication dispense is the result of a pharmacy system responding to a medication order.
type MedicationDispense struct {
	_ *DomainResource
	// A link to a resource representing the person or the group to whom the medication will be given.
	Subject *Reference `json:"subject"`
	// The encounter or episode of care that establishes the context for this event.
	Context *Reference `json:"context"`
	// Extra information about the dispense that could not be conveyed in the other attributes.
	Note []*Annotation `json:"note"`
	// This is a MedicationDispense resource
	ResourceType string `json:"resourceType,omitempty"`
	// The procedure that the dispense is done because of.
	PartOf []*Reference `json:"partOf"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Indicates the reason why a dispense was not performed.
	NotDoneReasonCodeableConcept *CodeableConcept `json:"notDoneReasonCodeableConcept"`
	// A summary of the events of interest that have occurred, such as when the dispense was verified.
	EventHistory []*Reference `json:"eventHistory"`
	// The time when the dispensed product was packaged and reviewed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	WhenPrepared time.Time `json:"whenPrepared"`
	// Extensions for whenHandedOver
	WhenHandedOver_ext *Element `json:"_whenHandedOver"`
	// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. Drug-drug interaction, duplicate therapy, dosage alert etc.
	DetectedIssue []*Reference `json:"detectedIssue"`
	// The time the dispensed product was provided to the patient or their representative.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	WhenHandedOver time.Time `json:"whenHandedOver"`
	// Indicates how the medication is to be used by the patient.
	DosageInstruction []*Dosage `json:"dosageInstruction"`
	// True if the dispense was not performed for some reason.
	NotDone bool `json:"notDone"`
	// Indicates who or what performed the event.  It should be assumed that the performer is the dispenser of the medication.
	Performer []*MedicationDispense_Performer `json:"performer"`
	// Indicates the medication order that is being dispensed against.
	AuthorizingPrescription []*Reference `json:"authorizingPrescription"`
	// Indicates whether or not substitution was made as part of the dispense.  In some cases substitution will be expected but does not happen, in other cases substitution is not expected but does happen.  This block explains what substitution did or did not happen and why.  If nothing is specified, substitution was not done.
	Substitution *MedicationDispense_Substitution `json:"substitution"`
	// Extensions for notDone
	NotDone_ext *Element `json:"_notDone"`
	// Indicates the reason why a dispense was not performed.
	NotDoneReasonReference *Reference `json:"notDoneReasonReference"`
	// Identifier assigned by the dispensing facility - this is an identifier assigned outside FHIR.
	Identifier []*Identifier `json:"identifier"`
	// A code specifying the state of the set of dispense events.
	Status string `json:"status"`
	// Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	MedicationCodeableConcept *CodeableConcept `json:"medicationCodeableConcept"`
	// Indicates type of medication dispense and where the medication is expected to be consumed or administered.
	Category *CodeableConcept `json:"category"`
	// Extensions for whenPrepared
	WhenPrepared_ext *Element `json:"_whenPrepared"`
	// Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	MedicationReference *Reference `json:"medicationReference"`
	// Additional information that supports the medication being dispensed.
	SupportingInformation []*Reference `json:"supportingInformation"`
	// The amount of medication that has been dispensed. Includes unit of measure.
	Quantity *Quantity `json:"quantity"`
	// Identifies the person who picked up the medication.  This will usually be a patient or their caregiver, but some cases exist where it can be a healthcare professional.
	Receiver []*Reference `json:"receiver"`
	// Indicates the type of dispensing event that is performed. For example, Trial Fill, Completion of Trial, Partial Fill, Emergency Fill, Samples, etc.
	Type *CodeableConcept `json:"type"`
	// The amount of medication expressed as a timing amount.
	DaysSupply *Quantity `json:"daysSupply"`
	// Identification of the facility/location where the medication was shipped to, as part of the dispense event.
	Destination *Reference `json:"destination"`
}

func NewMedicationDispense() *MedicationDispense {
	return &MedicationDispense{ResourceType: "MedicationDispense"}
}

// Extension is Optional Extension Element - found in all resources.
type Extension struct {
	_ *Element
	// Extensions for valueMarkdown
	ValueMarkdown_ext *Element `json:"_valueMarkdown"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueSignature *Signature `json:"valueSignature"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueHumanName *HumanName `json:"valueHumanName"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueMeta *Meta `json:"valueMeta"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueContactDetail *ContactDetail `json:"valueContactDetail"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern [0]|([1-9][0-9]*)
	ValueUnsignedInt uint64 `json:"valueUnsignedInt"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern [1-9][0-9]*
	ValuePositiveInt uint64 `json:"valuePositiveInt"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueAnnotation *Annotation `json:"valueAnnotation"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueAttachment *Attachment `json:"valueAttachment"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueCount *Count `json:"valueCount"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueReference *Reference `json:"valueReference"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	ValueTime time.Time `json:"valueTime"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// Extensions for valueUri
	ValueUri_ext *Element `json:"_valueUri"`
	// Extensions for valueDateTime
	ValueDateTime_ext *Element `json:"_valueDateTime"`
	// Extensions for valuePositiveInt
	ValuePositiveInt_ext *Element `json:"_valuePositiveInt"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern -?([0]|([1-9][0-9]*))
	ValueInteger int64 `json:"valueInteger"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueInstant string `json:"valueInstant"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueMarkdown string `json:"valueMarkdown"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueSampledData *SampledData `json:"valueSampledData"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueContactPoint *ContactPoint `json:"valueContactPoint"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact"`
	// Extensions for valueInteger
	ValueInteger_ext *Element `json:"_valueInteger"`
	// Extensions for valueOid
	ValueOid_ext *Element `json:"_valueOid"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueElement *Element `json:"valueElement"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueNarrative *Narrative `json:"valueNarrative"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueIdentifier *Identifier `json:"valueIdentifier"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueRange *Range `json:"valueRange"`
	// Extensions for valueDate
	ValueDate_ext *Element `json:"_valueDate"`
	// Extensions for valueUuid
	ValueUuid_ext *Element `json:"_valueUuid"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueDistance *Distance `json:"valueDistance"`
	// Extensions for valueTime
	ValueTime_ext *Element `json:"_valueTime"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	ValueUuid string `json:"valueUuid"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueAddress *Address `json:"valueAddress"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ValueDateTime time.Time `json:"valueDateTime"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueRatio *Ratio `json:"valueRatio"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueUsageContext *UsageContext `json:"valueUsageContext"`
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern [A-Za-z0-9\-\.]{1,64}
	ValueId string `json:"valueId"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern [^\s]+([\s]?[^\s]+)*
	ValueCode string `json:"valueCode"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueDuration *Duration `json:"valueDuration"`
	// Extensions for valueId
	ValueId_ext *Element `json:"_valueId"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueBoolean bool `json:"valueBoolean"`
	// Extensions for valueInstant
	ValueInstant_ext *Element `json:"_valueInstant"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueSimpleQuantity *Quantity `json:"valueSimpleQuantity"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueAge *Age `json:"valueAge"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ValueDate time.Time `json:"valueDate"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueQuantity *Quantity `json:"valueQuantity"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueElementDefinition *ElementDefinition `json:"valueElementDefinition"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueDosage *Dosage `json:"valueDosage"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueBackboneElement *BackboneElement `json:"valueBackboneElement"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueCoding *Coding `json:"valueCoding"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueUri string `json:"valueUri"`
	// Extensions for valueCode
	ValueCode_ext *Element `json:"_valueCode"`
	// Extensions for valueUnsignedInt
	ValueUnsignedInt_ext *Element `json:"_valueUnsignedInt"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueExtension *Extension `json:"valueExtension"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueTiming *Timing `json:"valueTiming"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	ValueDecimal float64 `json:"valueDecimal"`
	// Extensions for valueDecimal
	ValueDecimal_ext *Element `json:"_valueDecimal"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValuePeriod *Period `json:"valuePeriod"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueContributor *Contributor `json:"valueContributor"`
	// Source of the definition for the extension code - a logical name or a URL.
	Url string `json:"url"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueMoney *Money `json:"valueMoney"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueString string `json:"valueString"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	ValueOid string `json:"valueOid"`
	// Value of extension - may be a resource or one of a constrained set of the data types (see Extensibility in the spec for list).
	ValueBase64Binary string `json:"valueBase64Binary"`
	// Extensions for valueBase64Binary
	ValueBase64Binary_ext *Element `json:"_valueBase64Binary"`
}

// Account is A financial tool for tracking value accrued for a particular purpose.  In the healthcare field, used to track charges for a patient, cost centers, etc.
type Account struct {
	_ *DomainResource
	// This is a Account resource
	ResourceType string `json:"resourceType,omitempty"`
	// Indicates whether the account is presently used/usable or not.
	Status string `json:"status"`
	// Categorizes the account for reporting and searching purposes.
	Type *CodeableConcept `json:"type"`
	// Name used for the account when displaying it to humans in reports, etc.
	Name string `json:"name"`
	// Identifies the period of time the account applies to; e.g. accounts created per fiscal year, quarter, etc.
	Period *Period `json:"period"`
	// Unique identifier used to reference the account.  May or may not be intended for human use (e.g. credit card number).
	Identifier []*Identifier `json:"identifier"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Provides additional information about what the account tracks and how it is used.
	Description string `json:"description"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Represents the sum of all credits less all debits associated with the account.  Might be positive, zero or negative.
	Balance *Money `json:"balance"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Parties financially responsible for the account.
	Guarantor []*Account_Guarantor `json:"guarantor"`
	// Identifies the patient, device, practitioner, location or other object the account is associated with.
	Subject *Reference `json:"subject"`
	// Indicates the period of time over which the account is allowed to have transactions posted to it.
	// This period may be different to the coveragePeriod which is the duration of time that services may occur.
	Active *Period `json:"active"`
	// The party(s) that are responsible for covering the payment of this account, and what order should they be applied to the account.
	Coverage []*Account_Coverage `json:"coverage"`
	// Indicates the organization, department, etc. with responsibility for the account.
	Owner *Reference `json:"owner"`
}

func NewAccount() *Account { return &Account{ResourceType: "Account"} }

// AllergyIntolerance_Reaction is Risk of harmful or undesirable, physiological response which is unique to an individual and associated with exposure to a substance.
type AllergyIntolerance_Reaction struct {
	_ *BackboneElement
	// Clinical symptoms and/or signs that are observed or associated with the adverse reaction event.
	Manifestation []*CodeableConcept `json:"manifestation,omitempty"`
	// Text description about the reaction as a whole, including details of the manifestation if required.
	Description string `json:"description"`
	// Identification of the route by which the subject was exposed to the substance.
	ExposureRoute *CodeableConcept `json:"exposureRoute"`
	// Additional text about the adverse reaction event not captured in other fields.
	Note []*Annotation `json:"note"`
	// Identification of the specific substance (or pharmaceutical product) considered to be responsible for the Adverse Reaction event. Note: the substance for a specific reaction may be different from the substance identified as the cause of the risk, but it must be consistent with it. For instance, it may be a more specific substance (e.g. a brand medication) or a composite product that includes the identified substance. It must be clinically safe to only process the 'code' and ignore the 'reaction.substance'.
	Substance *CodeableConcept `json:"substance"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Record of the date and/or time of the onset of the Reaction.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Onset time.Time `json:"onset"`
	// Extensions for onset
	Onset_ext *Element `json:"_onset"`
	// Clinical assessment of the severity of the reaction event as a whole, potentially considering multiple different manifestations.
	Severity string `json:"severity"`
	// Extensions for severity
	Severity_ext *Element `json:"_severity"`
}

// Person is Demographics and administrative information about a person independent of a specific health-related context.
type Person struct {
	_ *DomainResource
	// One or more addresses for the person.
	Address []*Address `json:"address"`
	// An image that can be displayed as a thumbnail of the person to enhance the identification of the individual.
	Photo *Attachment `json:"photo"`
	// Whether this person's record is in active use.
	Active bool `json:"active"`
	// Identifier for a person within a particular scope.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for gender
	Gender_ext *Element `json:"_gender"`
	// The birth date for the person.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	BirthDate time.Time `json:"birthDate"`
	// A name associated with the person.
	Name []*HumanName `json:"name"`
	// Administrative Gender.
	Gender string `json:"gender"`
	// The organization that is the custodian of the person record.
	ManagingOrganization *Reference `json:"managingOrganization"`
	// Link to a resource that concerns the same actual person.
	Link []*Person_Link `json:"link"`
	// A contact detail for the person, e.g. a telephone number or an email address.
	Telecom []*ContactPoint `json:"telecom"`
	// Extensions for birthDate
	BirthDate_ext *Element `json:"_birthDate"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
	// This is a Person resource
	ResourceType string `json:"resourceType,omitempty"`
}

func NewPerson() *Person { return &Person{ResourceType: "Person"} }

// GuidanceResponse is A guidance response is the formal response to a guidance request, including any output parameters returned by the evaluation, as well as the description of any proposed actions to be taken.
type GuidanceResponse struct {
	_ *DomainResource
	// The id of the request associated with this response. If an id was given as part of the request, it will be reproduced here to enable the requester to more easily identify the response in a multi-request scenario.
	// pattern [A-Za-z0-9\-\.]{1,64}
	RequestId string `json:"requestId"`
	// Allows the context of the guidance response to be provided if available. In a service context, this would likely be unavailable.
	Context *Reference `json:"context"`
	// Indicates the reason the request was initiated. This is typically provided as a parameter to the evaluation and echoed by the service, although for some use cases, such as subscription- or event-based scenarios, it may provide an indication of the cause for the response.
	ReasonReference *Reference `json:"reasonReference"`
	// Provides a mechanism to communicate additional information about the response.
	Note []*Annotation `json:"note"`
	// The actions, if any, produced by the evaluation of the artifact.
	Result *Reference `json:"result"`
	// Extensions for requestId
	RequestId_ext *Element `json:"_requestId"`
	// A reference to the knowledge module that was invoked.
	Module *Reference `json:"module,omitempty"`
	// The status of the response. If the evaluation is completed successfully, the status will indicate success. However, in order to complete the evaluation, the engine may require more information. In this case, the status will be data-required, and the response will contain a description of the additional required information. If the evaluation completed successfully, but the engine determines that a potentially more accurate response could be provided if more data was available, the status will be data-requested, and the response will contain a description of the additional requested information.
	Status string `json:"status"`
	// The patient for which the request was processed.
	Subject *Reference `json:"subject"`
	// Messages resulting from the evaluation of the artifact or artifacts. As part of evaluating the request, the engine may produce informational or warning messages. These messages will be provided by this element.
	EvaluationMessage []*Reference `json:"evaluationMessage"`
	// This is a GuidanceResponse resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Indicates when the guidance response was processed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OccurrenceDateTime time.Time `json:"occurrenceDateTime"`
	// Indicates the reason the request was initiated. This is typically provided as a parameter to the evaluation and echoed by the service, although for some use cases, such as subscription- or event-based scenarios, it may provide an indication of the cause for the response.
	ReasonCodeableConcept *CodeableConcept `json:"reasonCodeableConcept"`
	// If the evaluation could not be completed due to lack of information, or additional information would potentially result in a more accurate response, this element will a description of the data required in order to proceed with the evaluation. A subsequent request to the service should include this data.
	DataRequirement []*DataRequirement `json:"dataRequirement"`
	// Allows a service to provide a unique, business identifier for the response.
	Identifier *Identifier `json:"identifier"`
	// Extensions for occurrenceDateTime
	OccurrenceDateTime_ext *Element `json:"_occurrenceDateTime"`
	// Provides a reference to the device that performed the guidance.
	Performer *Reference `json:"performer"`
	// The output parameters of the evaluation, if any. Many modules will result in the return of specific resources such as procedure or communication requests that are returned as part of the operation result. However, modules may define specific outputs that would be returned as the result of the evaluation, and these would be returned in this element.
	OutputParameters *Reference `json:"outputParameters"`
}

func NewGuidanceResponse() *GuidanceResponse {
	return &GuidanceResponse{ResourceType: "GuidanceResponse"}
}

// OperationDefinition_Parameter is A formal computable definition of an operation (on the RESTful interface) or a named query (using the search interaction).
type OperationDefinition_Parameter struct {
	_ *BackboneElement
	// The minimum number of times this parameter SHALL appear in the request or response.
	// pattern -?([0]|([1-9][0-9]*))
	Min int64 `json:"min"`
	// Extensions for min
	Min_ext *Element `json:"_min"`
	// The maximum number of times this element is permitted to appear in the request or response.
	Max string `json:"max"`
	// How the parameter is understood as a search parameter. This is only used if the parameter type is 'string'.
	SearchType string `json:"searchType"`
	// The parts of a nested Parameter.
	Part []*OperationDefinition_Parameter `json:"part"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The type for this parameter.
	// pattern [^\s]+([\s]?[^\s]+)*
	Type string `json:"type"`
	// Whether this is an input or an output parameter.
	Use string `json:"use"`
	// Extensions for use
	Use_ext *Element `json:"_use"`
	// Describes the meaning or use of this parameter.
	Documentation string `json:"documentation"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// A profile the specifies the rules that this parameter must conform to.
	Profile *Reference `json:"profile"`
	// Binds to a value set if this parameter is coded (code, Coding, CodeableConcept).
	Binding *OperationDefinition_Binding `json:"binding"`
	// The name of used to identify the parameter.
	// pattern [^\s]+([\s]?[^\s]+)*
	Name string `json:"name"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// Extensions for searchType
	SearchType_ext *Element `json:"_searchType"`
	// Extensions for max
	Max_ext *Element `json:"_max"`
}

// Procedure_Performer is An action that is or was performed on a patient. This can be a physical intervention like an operation, or less invasive like counseling or hypnotherapy.
type Procedure_Performer struct {
	_ *BackboneElement
	// For example: surgeon, anaethetist, endoscopist.
	Role *CodeableConcept `json:"role"`
	// The practitioner who was involved in the procedure.
	Actor *Reference `json:"actor,omitempty"`
	// The organization the device or practitioner was acting on behalf of.
	OnBehalfOf *Reference `json:"onBehalfOf"`
}

// StructureDefinition is A definition of a FHIR structure. This resource is used to describe the underlying resources, data types defined in FHIR, and also for describing extensions and constraints on resources and data types.
type StructureDefinition struct {
	_ *DomainResource
	// Extensions for contextType
	ContextType_ext *Element `json:"_contextType"`
	// Extensions for context
	Context_ext []*Element `json:"_context"`
	// The identifier that is used to identify this version of the structure definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the structure definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate structure definition instances.
	UseContext []*UsageContext `json:"useContext"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// If this is an extension, Identifies the context within FHIR resources where the extension can be used.
	ContextType string `json:"contextType"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// Whether structure this definition describes is abstract or not  - that is, whether the structure is not intended to be instantiated. For Resources and Data types, abstract types will never be exchanged  between systems.
	Abstract bool `json:"abstract"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// A snapshot view is expressed in a stand alone form that can be used and interpreted without considering the base StructureDefinition.
	Snapshot *StructureDefinition_Snapshot `json:"snapshot"`
	// Extensions for contextInvariant
	ContextInvariant_ext []*Element `json:"_contextInvariant"`
	// Extensions for derivation
	Derivation_ext *Element `json:"_derivation"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// A natural language name identifying the structure definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// The date  (and optionally time) when the structure definition was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the structure definition changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for abstract
	Abstract_ext *Element `json:"_abstract"`
	// A legal or geographic region in which the structure definition is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Explaination of why this structure definition is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// A copyright statement relating to the structure definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the structure definition.
	Copyright string `json:"copyright"`
	// This is a StructureDefinition resource
	ResourceType string `json:"resourceType,omitempty"`
	// A short, descriptive, user-friendly title for the structure definition.
	Title string `json:"title"`
	// A boolean value to indicate that this structure definition is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for kind
	Kind_ext *Element `json:"_kind"`
	// An external specification that the content is mapped to.
	Mapping []*StructureDefinition_Mapping `json:"mapping"`
	// How the type relates to the baseDefinition.
	Derivation string `json:"derivation"`
	// An absolute URI that is used to identify this structure definition when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this structure definition is (or will be) published. The URL SHOULD include the major version of the structure definition. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// The status of this structure definition. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// A free text natural language description of the structure definition from a consumer's perspective.
	Description string `json:"description"`
	// The version of the FHIR specification on which this StructureDefinition is based - this is the formal version of the specification, without the revision number, e.g. [publication].[major].[minor], which is 3.0.1 for this version.
	// pattern [A-Za-z0-9\-\.]{1,64}
	FhirVersion string `json:"fhirVersion"`
	// Extensions for baseDefinition
	BaseDefinition_ext *Element `json:"_baseDefinition"`
	// A formal identifier that is used to identify this structure definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// Defines the kind of structure that this definition is describing.
	Kind string `json:"kind"`
	// An absolute URI that is the base structure from which this type is derived, either by specialization or constraint.
	BaseDefinition string `json:"baseDefinition"`
	// The name of the individual or organization that published the structure definition.
	Publisher string `json:"publisher"`
	// Extensions for fhirVersion
	FhirVersion_ext *Element `json:"_fhirVersion"`
	// Identifies the types of resource or data type elements to which the extension can be applied.
	Context []string `json:"context"`
	// A set of rules as Fluent Invariants about when the extension can be used (e.g. co-occurrence variants for the extension).
	ContextInvariant []string `json:"contextInvariant"`
	// The type this structure describes. If the derivation kind is 'specialization' then this is the master definition for a type, and there is always one of these (a data type, an extension, a resource, including abstract ones). Otherwise the structure definition is a constraint on the stated type (and in this case, the type cannot be an abstract type).
	// pattern [^\s]+([\s]?[^\s]+)*
	Type string `json:"type"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// A set of key words or terms from external terminologies that may be used to assist with indexing and searching of templates.
	Keyword []*Coding `json:"keyword"`
	// A differential view is expressed relative to the base StructureDefinition - a statement of differences that it applies.
	Differential *StructureDefinition_Differential `json:"differential"`
}

func NewStructureDefinition() *StructureDefinition {
	return &StructureDefinition{ResourceType: "StructureDefinition"}
}

// Claim_CareTeam is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_CareTeam struct {
	_ *BackboneElement
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// Member of the team who provided the overall service.
	Provider *Reference `json:"provider,omitempty"`
	// The party who is billing and responsible for the claimed good or service rendered to the patient.
	Responsible bool `json:"responsible"`
	// Extensions for responsible
	Responsible_ext *Element `json:"_responsible"`
	// The lead, assisting or supervising practitioner and their discipline if a multidisiplinary team.
	Role *CodeableConcept `json:"role"`
	// The qualification which is applicable for this service.
	Qualification *CodeableConcept `json:"qualification"`
	// Sequence of the careTeam which serves to order and provide a link.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
}

// ExplanationOfBenefit_Detail1 is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Detail1 struct {
	_ *BackboneElement
	// A code to indicate the Professional Service or Product supplied (eg. CTP, HCPCS,USCLS,ICD10, NCPDP,DIN,ACHI,CCI).
	Service *CodeableConcept `json:"service"`
	// Item typification or modifiers codes, eg for Oral whether the treatment is cosmetic or associated with TMJ, or for medical whether the treatment was outside the clinic or out of office hours.
	Modifier []*CodeableConcept `json:"modifier"`
	// The fee charged for the professional service or product.
	Fee *Money `json:"fee"`
	// A list of note references to the notes provided below.
	NoteNumber []uint64 `json:"noteNumber"`
	// Extensions for noteNumber
	NoteNumber_ext []*Element `json:"_noteNumber"`
	// The adjudications results.
	Adjudication []*ExplanationOfBenefit_Adjudication `json:"adjudication"`
	// The type of reveneu or cost center providing the product and/or service.
	Revenue *CodeableConcept `json:"revenue"`
	// Health Care Service Type Codes  to identify the classification of service or benefits.
	Category *CodeableConcept `json:"category"`
}

// ImmunizationRecommendation_Protocol is A patient's point-in-time immunization and recommendation (i.e. forecasting a patient's immunization eligibility according to a published schedule) with optional supporting justification.
type ImmunizationRecommendation_Protocol struct {
	_ *BackboneElement
	// Indicates the nominal position in a series of the next dose.  This is the recommended dose number as per a specified protocol.
	// pattern [1-9][0-9]*
	DoseSequence uint64 `json:"doseSequence"`
	// Extensions for doseSequence
	DoseSequence_ext *Element `json:"_doseSequence"`
	// Contains the description about the protocol under which the vaccine was administered.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Indicates the authority who published the protocol.  For example, ACIP.
	Authority *Reference `json:"authority"`
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series string `json:"series"`
	// Extensions for series
	Series_ext *Element `json:"_series"`
}

// StructureMap_Group is A Map of relationships between 2 structures that can be used to transform data.
type StructureMap_Group struct {
	_ *BackboneElement
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// A name assigned to an instance of data. The instance must be provided when the mapping is invoked.
	Input []*StructureMap_Input `json:"input,omitempty"`
	// A unique name for the group for the convenience of human readers.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Additional supporting documentation that explains the purpose of the group and the types of mappings within it.
	Documentation string `json:"documentation"`
	// Extensions for typeMode
	TypeMode_ext *Element `json:"_typeMode"`
	// Transform Rule from source to target.
	Rule []*StructureMap_Rule `json:"rule,omitempty"`
	// Another group that this group adds rules to.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Extends string `json:"extends"`
	// Extensions for extends
	Extends_ext *Element `json:"_extends"`
	// If this is the default rule set to apply for thie source type, or this combination of types.
	TypeMode string `json:"typeMode"`
}

// TestReport_Participant is A summary of information based on the results of executing a TestScript.
type TestReport_Participant struct {
	_ *BackboneElement
	// The display name of the participant.
	Display string `json:"display"`
	// Extensions for display
	Display_ext *Element `json:"_display"`
	// The type of participant.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The uri of the participant. An absolute URL is preferred.
	Uri string `json:"uri"`
	// Extensions for uri
	Uri_ext *Element `json:"_uri"`
}

// Consent_Data is A record of a healthcare consumer's policy choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent_Data struct {
	_ *BackboneElement
	// Extensions for meaning
	Meaning_ext *Element `json:"_meaning"`
	// A reference to a specific resource that defines which resources are covered by this consent.
	Reference *Reference `json:"reference,omitempty"`
	// How the resource reference is interpreted when testing consent restrictions.
	Meaning string `json:"meaning"`
}

// EligibilityResponse_BenefitBalance is This resource provides eligibility and plan details from the processing of an Eligibility resource.
type EligibilityResponse_BenefitBalance struct {
	_ *BackboneElement
	// Extensions for excluded
	Excluded_ext *Element `json:"_excluded"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Unit designation: individual or family.
	Unit *CodeableConcept `json:"unit"`
	// Benefits Used to date.
	Financial []*EligibilityResponse_Financial `json:"financial"`
	// Network designation.
	Network *CodeableConcept `json:"network"`
	// The term or period of the values such as 'maximum lifetime benefit' or 'maximum annual vistis'.
	Term *CodeableConcept `json:"term"`
	// Dental, Vision, Medical, Pharmacy, Rehab etc.
	Category *CodeableConcept `json:"category,omitempty"`
	// Dental: basic, major, ortho; Vision exam, glasses, contacts; etc.
	SubCategory *CodeableConcept `json:"subCategory"`
	// True if the indicated class of service is excluded from the plan, missing or False indicated the service is included in the coverage.
	Excluded bool `json:"excluded"`
	// A short name or tag for the benefit, for example MED01, or DENT2.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// A richer description of the benefit, for example 'DENT2 covers 100% of basic, 50% of major but exclused Ortho, Implants and Costmetic services'.
	Description string `json:"description"`
}

// PlanDefinition is This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinition struct {
	_ *DomainResource
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// The status of this plan definition. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for approvalDate
	ApprovalDate_ext *Element `json:"_approvalDate"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// An absolute URI that is used to identify this plan definition when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this plan definition is (or will be) published. The URL SHOULD include the major version of the plan definition. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The date  (and optionally time) when the plan definition was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the plan definition changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// A copyright statement relating to the plan definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the plan definition.
	Copyright string `json:"copyright"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// The date on which the resource content was last reviewed. Review happens periodically after approval, but doesn't change the original approval date.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	LastReviewDate time.Time `json:"lastReviewDate"`
	// Descriptive topics related to the content of the plan definition. Topics provide a high-level categorization of the definition that can be useful for filtering and searching.
	Topic []*CodeableConcept `json:"topic"`
	// This is a PlanDefinition resource
	ResourceType string `json:"resourceType,omitempty"`
	// A formal identifier that is used to identify this plan definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []*Identifier `json:"identifier"`
	// The identifier that is used to identify this version of the plan definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the plan definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	Version string `json:"version"`
	// A boolean value to indicate that this plan definition is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Explaination of why this plan definition is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// Extensions for usage
	Usage_ext *Element `json:"_usage"`
	// Goals that describe what the activities within the plan are intended to achieve. For example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc.
	Goal []*PlanDefinition_Goal `json:"goal"`
	// A reference to a Library resource containing any formal logic used by the plan definition.
	Library []*Reference `json:"library"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// A short, descriptive, user-friendly title for the plan definition.
	Title string `json:"title"`
	// The type of asset the plan definition represents, e.g. an order set, protocol, or event-condition-action rule.
	Type *CodeableConcept `json:"type"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// A detailed description of how the asset is used from a clinical perspective.
	Usage string `json:"usage"`
	// The period during which the plan definition content was or is planned to be in active use.
	EffectivePeriod *Period `json:"effectivePeriod"`
	// A free text natural language description of the plan definition from a consumer's perspective.
	Description string `json:"description"`
	// A legal or geographic region in which the plan definition is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Extensions for lastReviewDate
	LastReviewDate_ext *Element `json:"_lastReviewDate"`
	// A contributor to the content of the asset, including authors, editors, reviewers, and endorsers.
	Contributor []*Contributor `json:"contributor"`
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []*RelatedArtifact `json:"relatedArtifact"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate plan definition instances.
	UseContext []*UsageContext `json:"useContext"`
	// An action to be taken as part of the plan.
	Action []*PlanDefinition_Action `json:"action"`
	// A natural language name identifying the plan definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The name of the individual or organization that published the plan definition.
	Publisher string `json:"publisher"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ApprovalDate time.Time `json:"approvalDate"`
}

func NewPlanDefinition() *PlanDefinition { return &PlanDefinition{ResourceType: "PlanDefinition"} }

// PlanDefinition_Action is This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinition_Action struct {
	_ *BackboneElement
	// Extensions for label
	Label_ext *Element `json:"_label"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Defines the selection behavior for the action and its children.
	SelectionBehavior string `json:"selectionBehavior"`
	// The title of the action displayed to a user.
	Title string `json:"title"`
	// A text equivalent of the action to be performed. This provides a human-interpretable description of the action when the definition is consumed by a system that may not be capable of interpreting it dynamically.
	TextEquivalent string `json:"textEquivalent"`
	// Customizations that should be applied to the statically defined resource. For example, if the dosage of a medication must be computed based on the patient's weight, a customization would be used to specify an expression that calculated the weight, and the path on the resource that would contain the result.
	DynamicValue []*PlanDefinition_DynamicValue `json:"dynamicValue"`
	// A user-visible label for the action.
	Label string `json:"label"`
	// Extensions for goalId
	GoalId_ext []*Element `json:"_goalId"`
	// An expression that describes applicability criteria, or start/stop conditions for the action.
	Condition []*PlanDefinition_Condition `json:"condition"`
	// Defines input data requirements for the action.
	Input []*DataRequirement `json:"input"`
	// Defines the outputs of the action, if any.
	Output []*DataRequirement `json:"output"`
	// Sub actions that are contained within the action. The behavior of this action determines the functionality of the sub-actions. For example, a selection behavior of at-most-one indicates that of the sub-actions, at most one may be chosen as part of realizing the action definition.
	Action []*PlanDefinition_Action `json:"action"`
	// Extensions for textEquivalent
	TextEquivalent_ext *Element `json:"_textEquivalent"`
	// A description of why this action is necessary or appropriate.
	Reason []*CodeableConcept `json:"reason"`
	// The type of action to perform (create, update, remove).
	Type *Coding `json:"type"`
	// Defines whether the action should usually be preselected.
	PrecheckBehavior string `json:"precheckBehavior"`
	// Extensions for groupingBehavior
	GroupingBehavior_ext *Element `json:"_groupingBehavior"`
	// Defines whether the action can be selected multiple times.
	CardinalityBehavior string `json:"cardinalityBehavior"`
	// A reference to an ActivityDefinition that describes the action to be taken in detail, or a PlanDefinition that describes a series of actions to be taken.
	Definition *Reference `json:"definition"`
	// A short description of the action used to provide a summary to display to the user.
	Description string `json:"description"`
	// A code that provides meaning for the action or action group. For example, a section may have a LOINC code for a the section of a documentation template.
	Code []*CodeableConcept `json:"code"`
	// A description of when the action should be triggered.
	TriggerDefinition []*TriggerDefinition `json:"triggerDefinition"`
	// An optional value describing when the action should be performed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	TimingDateTime time.Time `json:"timingDateTime"`
	// An optional value describing when the action should be performed.
	TimingDuration *Duration `json:"timingDuration"`
	// Extensions for cardinalityBehavior
	CardinalityBehavior_ext *Element `json:"_cardinalityBehavior"`
	// Extensions for selectionBehavior
	SelectionBehavior_ext *Element `json:"_selectionBehavior"`
	// A reference to a StructureMap resource that defines a transform that can be executed to produce the intent resource using the ActivityDefinition instance as the input.
	Transform *Reference `json:"transform"`
	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient. Information resources can include inline text commentary and links to web resources.
	Documentation []*RelatedArtifact `json:"documentation"`
	// A relationship to another action such as "before" or "30-60 minutes after start of".
	RelatedAction []*PlanDefinition_RelatedAction `json:"relatedAction"`
	// Extensions for timingDateTime
	TimingDateTime_ext *Element `json:"_timingDateTime"`
	// An optional value describing when the action should be performed.
	TimingPeriod *Period `json:"timingPeriod"`
	// An optional value describing when the action should be performed.
	TimingRange *Range `json:"timingRange"`
	// Indicates who should participate in performing the action described.
	Participant []*PlanDefinition_Participant `json:"participant"`
	// Extensions for precheckBehavior
	PrecheckBehavior_ext *Element `json:"_precheckBehavior"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Identifies goals that this action supports. The reference must be to a goal element defined within this plan definition.
	GoalId []string `json:"goalId"`
	// An optional value describing when the action should be performed.
	TimingTiming *Timing `json:"timingTiming"`
	// Defines the grouping behavior for the action and its children.
	GroupingBehavior string `json:"groupingBehavior"`
	// Defines the requiredness behavior for the action.
	RequiredBehavior string `json:"requiredBehavior"`
	// Extensions for requiredBehavior
	RequiredBehavior_ext *Element `json:"_requiredBehavior"`
}

// SearchParameter_Component is A search parameter that defines a named search item that can be used to search/filter on a resource.
type SearchParameter_Component struct {
	_ *BackboneElement
	// The definition of the search parameter that describes this part.
	Definition *Reference `json:"definition,omitempty"`
	// A sub-expression that defines how to extract values for this component from the output of the main SearchParameter.expression.
	Expression string `json:"expression"`
	// Extensions for expression
	Expression_ext *Element `json:"_expression"`
}

// DomainResource is A resource that includes narrative, extensions, and contained resources.
type DomainResource struct {
	_ *Resource
	// A human-readable narrative that contains a summary of the resource, and may be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *Narrative `json:"text"`
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []interface{} `json:"contained"`
	// May be used to represent additional information that is not part of the basic definition of the resource. In order to make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []*Extension `json:"extension"`
	// May be used to represent additional information that is not part of the basic definition of the resource, and that modifies the understanding of the element that contains it. Usually modifier elements provide negation or qualification. In order to make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	ModifierExtension []*Extension `json:"modifierExtension"`
}

// ValueSet_Concept is A value set specifies a set of codes drawn from one or more code systems.
type ValueSet_Concept struct {
	_ *BackboneElement
	// Extensions for display
	Display_ext *Element `json:"_display"`
	// Additional representations for this concept when used in this value set - other languages, aliases, specialized purposes, used for particular purposes, etc.
	Designation []*ValueSet_Designation `json:"designation"`
	// Specifies a code for the concept to be included or excluded.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The text to display to the user for this concept in the context of this valueset. If no display is provided, then applications using the value set use the display specified for the code by the system.
	Display string `json:"display"`
}

// Signature is A digital signature along with supporting context. The signature may be electronic/cryptographic in nature, or a graphical image representing a hand-written signature, or a signature process. Different signature approaches have different utilities.
type Signature struct {
	_ *Element
	// Extensions for contentType
	ContentType_ext *Element `json:"_contentType"`
	// The base64 encoding of the Signature content. When signature is not recorded electronically this element would be empty.
	Blob string `json:"blob"`
	// Extensions for blob
	Blob_ext *Element `json:"_blob"`
	// Extensions for onBehalfOfUri
	OnBehalfOfUri_ext *Element `json:"_onBehalfOfUri"`
	// A reference to an application-usable description of the identity that is represented by the signature.
	OnBehalfOfReference *Reference `json:"onBehalfOfReference"`
	// Extensions for when
	When_ext *Element `json:"_when"`
	// A reference to an application-usable description of the identity that signed  (e.g. the signature used their private key).
	WhoUri string `json:"whoUri"`
	// Extensions for whoUri
	WhoUri_ext *Element `json:"_whoUri"`
	// A reference to an application-usable description of the identity that signed  (e.g. the signature used their private key).
	WhoReference *Reference `json:"whoReference"`
	// A reference to an application-usable description of the identity that is represented by the signature.
	OnBehalfOfUri string `json:"onBehalfOfUri"`
	// A mime type that indicates the technical format of the signature. Important mime types are application/signature+xml for X ML DigSig, application/jwt for JWT, and image/* for a graphical image of a signature, etc.
	// pattern [^\s]+([\s]?[^\s]+)*
	ContentType string `json:"contentType"`
	// An indication of the reason that the entity signed this document. This may be explicitly included as part of the signature information and can be used when determining accountability for various actions concerning the document.
	Type []*Coding `json:"type,omitempty"`
	// When the digital signature was signed.
	When string `json:"when"`
}

// ClaimResponse_Payment is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse_Payment struct {
	_ *BackboneElement
	// Payment identifier.
	Identifier *Identifier `json:"identifier"`
	// Whether this represents partial or complete payment of the claim.
	Type *CodeableConcept `json:"type"`
	// Adjustment to the payment of this transaction which is not related to adjudication of this transaction.
	Adjustment *Money `json:"adjustment"`
	// Reason for the payment adjustment.
	AdjustmentReason *CodeableConcept `json:"adjustmentReason"`
	// Estimated payment data.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Payable less any payment adjustment.
	Amount *Money `json:"amount"`
}

// Task_Restriction is A task to be performed.
type Task_Restriction struct {
	_ *BackboneElement
	// Over what time-period is fulfillment sought.
	Period *Period `json:"period"`
	// For requests that are targeted to more than on potential recipient/target, for whom is fulfillment sought?
	Recipient []*Reference `json:"recipient"`
	// Indicates the number of times the requested action should occur.
	// pattern [1-9][0-9]*
	Repetitions uint64 `json:"repetitions"`
	// Extensions for repetitions
	Repetitions_ext *Element `json:"_repetitions"`
}

// ExpansionProfile_Include is Resource to define constraints on the Expansion of a FHIR ValueSet.
type ExpansionProfile_Include struct {
	_ *BackboneElement
	// A data group for each designation to be included.
	Designation []*ExpansionProfile_Designation1 `json:"designation"`
}

// Organization is A formally or informally recognized grouping of people or organizations formed for the purpose of achieving some form of collective action.  Includes companies, institutions, corporations, departments, community groups, healthcare practice groups, etc.
type Organization struct {
	_ *DomainResource
	// Identifier for the organization that is used to identify the organization across multiple disparate systems.
	Identifier []*Identifier `json:"identifier"`
	// Whether the organization's record is still in active use.
	Active bool `json:"active"`
	// Extensions for active
	Active_ext *Element `json:"_active"`
	// Contact for the organization for a certain purpose.
	Contact []*Organization_Contact `json:"contact"`
	// This is a Organization resource
	ResourceType string `json:"resourceType,omitempty"`
	// A name associated with the organization.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// A list of alternate names that the organization is known as, or was known as in the past.
	Alias []string `json:"alias"`
	// A contact detail for the organization.
	Telecom []*ContactPoint `json:"telecom"`
	// Technical endpoints providing access to services operated for the organization.
	Endpoint []*Reference `json:"endpoint"`
	// The kind(s) of organization that this is.
	Type []*CodeableConcept `json:"type"`
	// Extensions for alias
	Alias_ext []*Element `json:"_alias"`
	// An address for the organization.
	Address []*Address `json:"address"`
	// The organization of which this organization forms a part.
	PartOf *Reference `json:"partOf"`
}

func NewOrganization() *Organization { return &Organization{ResourceType: "Organization"} }

// Procedure is An action that is or was performed on a patient. This can be a physical intervention like an operation, or less invasive like counseling or hypnotherapy.
type Procedure struct {
	_ *DomainResource
	// This is a Procedure resource
	ResourceType string `json:"resourceType,omitempty"`
	// Any complications that occurred during the procedure, or in the immediate post-performance period. These are generally tracked separately from the notes, which will typically describe the procedure itself rather than any 'post procedure' issues.
	Complication []*CodeableConcept `json:"complication"`
	// This records identifiers associated with this procedure that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []*Identifier `json:"identifier"`
	// The outcome of the procedure - did it resolve reasons for the procedure being performed?
	Outcome *CodeableConcept `json:"outcome"`
	// This could be a histology result, pathology report, surgical report, etc..
	Report []*Reference `json:"report"`
	// Any other notes about the procedure.  E.g. the operative notes.
	Note []*Annotation `json:"note"`
	// A protocol, guideline, orderset or other definition that was adhered to in whole or in part by this procedure.
	Definition []*Reference `json:"definition"`
	// Set this to true if the record is saying that the procedure was NOT performed.
	NotDone bool `json:"notDone"`
	// The date(time)/period over which the procedure was performed. Allows a period to support complex procedures that span more than one date, and also allows for the length of the procedure to be captured.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	PerformedDateTime time.Time `json:"performedDateTime"`
	// A reference to a resource that contains details of the request for this procedure.
	BasedOn []*Reference `json:"basedOn"`
	// A code specifying the state of the procedure. Generally this will be in-progress or completed state.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Identifies medications, devices and any other substance used as part of the procedure.
	UsedReference []*Reference `json:"usedReference"`
	// Limited to 'real' people rather than equipment.
	Performer []*Procedure_Performer `json:"performer"`
	// The coded reason why the procedure was performed. This may be coded entity of some type, or may simply be present as text.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// The condition that is the reason why the procedure was performed.
	ReasonReference []*Reference `json:"reasonReference"`
	// Detailed and structured anatomical location information. Multiple locations are allowed - e.g. multiple punch biopsies of a lesion.
	BodySite []*CodeableConcept `json:"bodySite"`
	// A larger event of which this particular procedure is a component or step.
	PartOf []*Reference `json:"partOf"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The date(time)/period over which the procedure was performed. Allows a period to support complex procedures that span more than one date, and also allows for the length of the procedure to be captured.
	PerformedPeriod *Period `json:"performedPeriod"`
	// If the procedure required specific follow up - e.g. removal of sutures. The followup may be represented as a simple note, or could potentially be more complex in which case the CarePlan resource can be used.
	FollowUp []*CodeableConcept `json:"followUp"`
	// A code indicating why the procedure was not performed.
	NotDoneReason *CodeableConcept `json:"notDoneReason"`
	// The person, animal or group on which the procedure was performed.
	Subject *Reference `json:"subject,omitempty"`
	// Extensions for performedDateTime
	PerformedDateTime_ext *Element `json:"_performedDateTime"`
	// A device that is implanted, removed or otherwise manipulated (calibration, battery replacement, fitting a prosthesis, attaching a wound-vac, etc.) as a focal portion of the Procedure.
	FocalDevice []*Procedure_FocalDevice `json:"focalDevice"`
	// Identifies coded items that were used as part of the procedure.
	UsedCode []*CodeableConcept `json:"usedCode"`
	// Extensions for notDone
	NotDone_ext *Element `json:"_notDone"`
	// The location where the procedure actually happened.  E.g. a newborn at home, a tracheostomy at a restaurant.
	Location *Reference `json:"location"`
	// Any complications that occurred during the procedure, or in the immediate post-performance period.
	ComplicationDetail []*Reference `json:"complicationDetail"`
	// A code that classifies the procedure for searching, sorting and display purposes (e.g. "Surgical Procedure").
	Category *CodeableConcept `json:"category"`
	// The specific procedure that is performed. Use text if the exact nature of the procedure cannot be coded (e.g. "Laparoscopic Appendectomy").
	Code *CodeableConcept `json:"code"`
	// The encounter during which the procedure was performed.
	Context *Reference `json:"context"`
}

func NewProcedure() *Procedure { return &Procedure{ResourceType: "Procedure"} }

// CapabilityStatement_Resource is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Resource struct {
	_ *BackboneElement
	// A code that indicates how the server supports conditional read.
	ConditionalRead string `json:"conditionalRead"`
	// Extensions for referencePolicy
	ReferencePolicy_ext []*Element `json:"_referencePolicy"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// Extensions for conditionalUpdate
	ConditionalUpdate_ext *Element `json:"_conditionalUpdate"`
	// A flag that indicates that the server supports conditional create.
	ConditionalCreate bool `json:"conditionalCreate"`
	// A type of resource exposed via the restful interface.
	// pattern [^\s]+([\s]?[^\s]+)*
	Type string `json:"type"`
	// Extensions for versioning
	Versioning_ext *Element `json:"_versioning"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Extensions for conditionalCreate
	ConditionalCreate_ext *Element `json:"_conditionalCreate"`
	// Extensions for searchRevInclude
	SearchRevInclude_ext []*Element `json:"_searchRevInclude"`
	// Extensions for readHistory
	ReadHistory_ext *Element `json:"_readHistory"`
	// A flag to indicate that the server allows or needs to allow the client to create new identities on the server (e.g. that is, the client PUTs to a location where there is no existing resource). Allowing this operation means that the server allows the client to create new identities on the server.
	UpdateCreate bool `json:"updateCreate"`
	// A flag that indicates that the server supports conditional update.
	ConditionalUpdate bool `json:"conditionalUpdate"`
	// A list of _include values supported by the server.
	SearchInclude []string `json:"searchInclude"`
	// Extensions for conditionalDelete
	ConditionalDelete_ext *Element `json:"_conditionalDelete"`
	// A set of flags that defines how references are supported.
	ReferencePolicy []string `json:"referencePolicy"`
	// Extensions for searchInclude
	SearchInclude_ext []*Element `json:"_searchInclude"`
	// Search parameters for implementations to support and/or make use of - either references to ones defined in the specification, or additional ones defined for/by the implementation.
	SearchParam []*CapabilityStatement_SearchParam `json:"searchParam"`
	// Extensions for conditionalRead
	ConditionalRead_ext *Element `json:"_conditionalRead"`
	// A specification of the profile that describes the solution's overall support for the resource, including any constraints on cardinality, bindings, lengths or other limitations. See further discussion in [Using Profiles](profiling.html#profile-uses).
	Profile *Reference `json:"profile"`
	// Additional information about the resource type used by the system.
	Documentation string `json:"documentation"`
	// Identifies a restful operation supported by the solution.
	Interaction []*CapabilityStatement_Interaction `json:"interaction,omitempty"`
	// A flag for whether the server is able to return past versions as part of the vRead operation.
	ReadHistory bool `json:"readHistory"`
	// A code that indicates how the server supports conditional delete.
	ConditionalDelete string `json:"conditionalDelete"`
	// This field is set to no-version to specify that the system does not support (server) or use (client) versioning for this resource type. If this has some other value, the server must at least correctly track and populate the versionId meta-property on resources. If the value is 'versioned-update', then the server supports all the versioning features, including using e-tags for version integrity in the API.
	Versioning string `json:"versioning"`
	// Extensions for updateCreate
	UpdateCreate_ext *Element `json:"_updateCreate"`
	// A list of _revinclude (reverse include) values supported by the server.
	SearchRevInclude []string `json:"searchRevInclude"`
}

// CapabilityStatement_Document is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Document struct {
	_ *BackboneElement
	// Mode of this document declaration - whether an application is a producer or consumer.
	Mode string `json:"mode"`
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// A description of how the application supports or uses the specified document profile.  For example, when documents are created, what action is taken with consumed documents, etc.
	Documentation string `json:"documentation"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// A constraint on a resource used in the document.
	Profile *Reference `json:"profile,omitempty"`
}

// Contract_Agent1 is A formal agreement between parties regarding the conduct of business, exchange of information or other matters.
type Contract_Agent1 struct {
	_ *BackboneElement
	// The agent assigned a role in this Contract Provision.
	Actor *Reference `json:"actor,omitempty"`
	// Role played by the agent assigned this role in the execution of this Contract Provision.
	Role []*CodeableConcept `json:"role"`
}

// PlanDefinition_RelatedAction is This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinition_RelatedAction struct {
	_ *BackboneElement
	// The element id of the related action.
	// pattern [A-Za-z0-9\-\.]{1,64}
	ActionId string `json:"actionId"`
	// Extensions for actionId
	ActionId_ext *Element `json:"_actionId"`
	// The relationship of this action to the related action.
	Relationship string `json:"relationship"`
	// Extensions for relationship
	Relationship_ext *Element `json:"_relationship"`
	// A duration or range of durations to apply to the relationship. For example, 30-60 minutes before.
	OffsetDuration *Duration `json:"offsetDuration"`
	// A duration or range of durations to apply to the relationship. For example, 30-60 minutes before.
	OffsetRange *Range `json:"offsetRange"`
}

// TestReport_Setup is A summary of information based on the results of executing a TestScript.
type TestReport_Setup struct {
	_ *BackboneElement
	// Action would contain either an operation or an assertion.
	Action []*TestReport_Action `json:"action,omitempty"`
}

// TestReport_Test is A summary of information based on the results of executing a TestScript.
type TestReport_Test struct {
	_ *BackboneElement
	// The name of this test used for tracking/logging purposes by test engines.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// A short description of the test used by test engines for tracking and reporting purposes.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Action would contain either an operation or an assertion.
	Action []*TestReport_Action1 `json:"action,omitempty"`
}

// ValueSet_Designation is A value set specifies a set of codes drawn from one or more code systems.
type ValueSet_Designation struct {
	_ *BackboneElement
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// The language this designation is defined for.
	// pattern [^\s]+([\s]?[^\s]+)*
	Language string `json:"language"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
	// A code that details how this designation would be used.
	Use *Coding `json:"use"`
	// The text value for this designation.
	Value string `json:"value"`
}

// CapabilityStatement_Messaging is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Messaging struct {
	_ *BackboneElement
	// Length if the receiver's reliable messaging cache in minutes (if a receiver) or how long the cache length on the receiver should be (if a sender).
	// pattern [0]|([1-9][0-9]*)
	ReliableCache uint64 `json:"reliableCache"`
	// Extensions for reliableCache
	ReliableCache_ext *Element `json:"_reliableCache"`
	// Documentation about the system's messaging capabilities for this endpoint not otherwise documented by the capability statement.  For example, the process for becoming an authorized messaging exchange partner.
	Documentation string `json:"documentation"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// References to message definitions for messages this system can send or receive.
	SupportedMessage []*CapabilityStatement_SupportedMessage `json:"supportedMessage"`
	// A description of the solution's support for an event at this end-point.
	Event []*CapabilityStatement_Event `json:"event"`
	// An endpoint (network accessible address) to which messages and/or replies are to be sent.
	Endpoint []*CapabilityStatement_Endpoint `json:"endpoint"`
}

// Contract_ValuedItem1 is A formal agreement between parties regarding the conduct of business, exchange of information or other matters.
type Contract_ValuedItem1 struct {
	_ *BackboneElement
	// Specific type of Contract Provision Valued Item that may be priced.
	EntityCodeableConcept *CodeableConcept `json:"entityCodeableConcept"`
	// Extensions for effectiveTime
	EffectiveTime_ext *Element `json:"_effectiveTime"`
	// A Contract Provision Valued Item unit valuation measure.
	UnitPrice *Money `json:"unitPrice"`
	// A real number that represents a multiplier used in determining the overall value of the Contract Provision Valued Item delivered. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Factor float64 `json:"factor"`
	// Extensions for factor
	Factor_ext *Element `json:"_factor"`
	// An amount that expresses the weighting (based on difficulty, cost and/or resource intensiveness) associated with the Contract Provision Valued Item delivered. The concept of Points allows for assignment of point values for a Contract ProvisionValued Item, such that a monetary amount can be assigned to each point.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Points float64 `json:"points"`
	// Specific type of Contract Provision Valued Item that may be priced.
	EntityReference *Reference `json:"entityReference"`
	// Identifies a Contract Provision Valued Item instance.
	Identifier *Identifier `json:"identifier"`
	// Indicates the time during which this Contract Term ValuedItem information is effective.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	EffectiveTime time.Time `json:"effectiveTime"`
	// Specifies the units by which the Contract Provision Valued Item is measured or counted, and quantifies the countable or measurable Contract Term Valued Item instances.
	Quantity *Quantity `json:"quantity"`
	// Extensions for points
	Points_ext *Element `json:"_points"`
	// Expresses the product of the Contract Provision Valued Item unitQuantity and the unitPriceAmt. For example, the formula: unit Quantity * unit Price (Cost per Point) * factor Number  * points = net Amount. Quantity, factor and points are assumed to be 1 if not supplied.
	Net *Money `json:"net"`
}

// GraphDefinition_Link is A formal computable definition of a graph of resources - that is, a coherent set of resources that form a graph by following references. The Graph Definition resource defines a set and makes rules about the set.
type GraphDefinition_Link struct {
	_ *BackboneElement
	// Minimum occurrences for this link.
	// pattern -?([0]|([1-9][0-9]*))
	Min int64 `json:"min"`
	// Extensions for max
	Max_ext *Element `json:"_max"`
	// Information about why this link is of interest in this graph definition.
	Description string `json:"description"`
	// Potential target for the link.
	Target []*GraphDefinition_Target `json:"target,omitempty"`
	// Path in the resource that contains the link.
	Path string `json:"path"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// Which slice (if profiled).
	SliceName string `json:"sliceName"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for sliceName
	SliceName_ext *Element `json:"_sliceName"`
	// Extensions for min
	Min_ext *Element `json:"_min"`
	// Maximum occurrences for this link.
	Max string `json:"max"`
}

// NutritionOrder_Texture is A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrder_Texture struct {
	_ *BackboneElement
	// Any texture modifications (for solid foods) that should be made, e.g. easy to chew, chopped, ground, and pureed.
	Modifier *CodeableConcept `json:"modifier"`
	// The food type(s) (e.g. meats, all foods)  that the texture modification applies to.  This could be all foods types.
	FoodType *CodeableConcept `json:"foodType"`
}

// Communication_Payload is An occurrence of information being transmitted; e.g. an alert that was sent to a responsible provider, a public health agency was notified about a reportable condition.
type Communication_Payload struct {
	_ *BackboneElement
	// A communicated content (or for multi-part communications, one portion of the communication).
	ContentString string `json:"contentString"`
	// Extensions for contentString
	ContentString_ext *Element `json:"_contentString"`
	// A communicated content (or for multi-part communications, one portion of the communication).
	ContentAttachment *Attachment `json:"contentAttachment"`
	// A communicated content (or for multi-part communications, one portion of the communication).
	ContentReference *Reference `json:"contentReference"`
}

// DeviceMetric is Describes a measurement, calculation or setting capability of a medical device.
type DeviceMetric struct {
	_ *DomainResource
	// Describes the unit that an observed value determined for this metric will have. For example: Percent, Seconds, etc.
	Unit *CodeableConcept `json:"unit"`
	// Describes the color representation for the metric. This is often used to aid clinicians to track and identify parameter types by color. In practice, consider a Patient Monitor that has ECG/HR and Pleth for example; the parameters are displayed in different characteristic colors, such as HR-blue, BP-green, and PR and SpO2- magenta.
	Color string `json:"color"`
	// Describes the calibrations that have been performed or that are required to be performed.
	Calibration []*DeviceMetric_Calibration `json:"calibration"`
	// Describes the unique identification of this metric that has been assigned by the device or gateway software. For example: handle ID.  It should be noted that in order to make the identifier unique, the system element of the identifier should be set to the unique identifier of the device.
	Identifier *Identifier `json:"identifier,omitempty"`
	// Extensions for category
	Category_ext *Element `json:"_category"`
	// Describes the link to the  Device that this DeviceMetric belongs to and that contains administrative device information such as manufacturer, serial number, etc.
	Source *Reference `json:"source"`
	// Describes the link to the  DeviceComponent that this DeviceMetric belongs to and that provide information about the location of this DeviceMetric in the containment structure of the parent Device. An example would be a DeviceComponent that represents a Channel. This reference can be used by a client application to distinguish DeviceMetrics that have the same type, but should be interpreted based on their containment location.
	Parent *Reference `json:"parent"`
	// Describes the measurement repetition time. This is not necessarily the same as the update period. The measurement repetition time can range from milliseconds up to hours. An example for a measurement repetition time in the range of milliseconds is the sampling rate of an ECG. An example for a measurement repetition time in the range of hours is a NIBP that is triggered automatically every hour. The update period may be different than the measurement repetition time, if the device does not update the published observed value with the same frequency as it was measured.
	MeasurementPeriod *Timing `json:"measurementPeriod"`
	// Extensions for operationalStatus
	OperationalStatus_ext *Element `json:"_operationalStatus"`
	// Extensions for color
	Color_ext *Element `json:"_color"`
	// Indicates the category of the observation generation process. A DeviceMetric can be for example a setting, measurement, or calculation.
	Category string `json:"category"`
	// This is a DeviceMetric resource
	ResourceType string `json:"resourceType,omitempty"`
	// Describes the type of the metric. For example: Heart Rate, PEEP Setting, etc.
	Type *CodeableConcept `json:"type,omitempty"`
	// Indicates current operational state of the device. For example: On, Off, Standby, etc.
	OperationalStatus string `json:"operationalStatus"`
}

func NewDeviceMetric() *DeviceMetric { return &DeviceMetric{ResourceType: "DeviceMetric"} }

// DiagnosticReport_Performer is The findings and interpretation of diagnostic  tests performed on patients, groups of patients, devices, and locations, and/or specimens derived from these. The report includes clinical context such as requesting and provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports.
type DiagnosticReport_Performer struct {
	_ *BackboneElement
	// Describes the type of participation (e.g.  a responsible party, author, or verifier).
	Role *CodeableConcept `json:"role"`
	// The reference to the  practitioner or organization involved in producing the report. For example, the diagnostic service that is responsible for issuing the report.
	Actor *Reference `json:"actor,omitempty"`
}

// TestScript_Capability is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Capability struct {
	_ *BackboneElement
	// Whether or not the test execution will validate the given capabilities of the server in order for this test script to execute.
	Validated bool `json:"validated"`
	// Extensions for validated
	Validated_ext *Element `json:"_validated"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Which server these requirements apply to.
	// pattern -?([0]|([1-9][0-9]*))
	Destination int64 `json:"destination"`
	// Extensions for link
	Link_ext []*Element `json:"_link"`
	// Whether or not the test execution will require the given capabilities of the server in order for this test script to execute.
	Required bool `json:"required"`
	// Extensions for required
	Required_ext *Element `json:"_required"`
	// Description of the capabilities that this test script is requiring the server to support.
	Description string `json:"description"`
	// Which origin server these requirements apply to.
	Origin []int64 `json:"origin"`
	// Extensions for origin
	Origin_ext []*Element `json:"_origin"`
	// Extensions for destination
	Destination_ext *Element `json:"_destination"`
	// Links to the FHIR specification that describes this interaction and the resources involved in more detail.
	Link []string `json:"link"`
	// Minimum capabilities required of server for test script to execute successfully.   If server does not meet at a minimum the referenced capability statement, then all tests in this script are skipped.
	Capabilities *Reference `json:"capabilities,omitempty"`
}

// VisionPrescription is An authorization for the supply of glasses and/or contact lenses to a patient.
type VisionPrescription struct {
	_ *DomainResource
	// Business identifier which may be used by other parties to reference or identify the prescription.
	Identifier []*Identifier `json:"identifier"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// The date (and perhaps time) when the prescription was written.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	DateWritten time.Time `json:"dateWritten"`
	// The healthcare professional responsible for authorizing the prescription.
	Prescriber *Reference `json:"prescriber"`
	// This is a VisionPrescription resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A link to a resource representing the person to whom the vision products will be supplied.
	Patient *Reference `json:"patient"`
	// A link to a resource that identifies the particular occurrence of contact between patient and health care provider.
	Encounter *Reference `json:"encounter"`
	// Extensions for dateWritten
	DateWritten_ext *Element `json:"_dateWritten"`
	// Can be the reason or the indication for writing the prescription.
	ReasonCodeableConcept *CodeableConcept `json:"reasonCodeableConcept"`
	// Can be the reason or the indication for writing the prescription.
	ReasonReference *Reference `json:"reasonReference"`
	// Deals with details of the dispense part of the supply specification.
	Dispense []*VisionPrescription_Dispense `json:"dispense"`
}

func NewVisionPrescription() *VisionPrescription {
	return &VisionPrescription{ResourceType: "VisionPrescription"}
}

// ChargeItem_Participant is The resource ChargeItem describes the provision of healthcare provider products for a certain patient, therefore referring not only to the product, but containing in addition details of the provision, like date, time, amounts and participating organizations and persons. Main Usage of the ChargeItem is to enable the billing process and internal cost allocation.
type ChargeItem_Participant struct {
	_ *BackboneElement
	// Describes the type of performance or participation(e.g. primary surgeon, anaesthesiologiest, etc.).
	Role *CodeableConcept `json:"role"`
	// The device, practitioner, etc. who performed or participated in the service.
	Actor *Reference `json:"actor,omitempty"`
}

// Contract_Signer is A formal agreement between parties regarding the conduct of business, exchange of information or other matters.
type Contract_Signer struct {
	_ *BackboneElement
	// Role of this Contract signer, e.g. notary, grantee.
	Type *Coding `json:"type,omitempty"`
	// Party which is a signator to this Contract.
	Party *Reference `json:"party,omitempty"`
	// Legally binding Contract DSIG signature contents in Base64.
	Signature []*Signature `json:"signature,omitempty"`
}

// Observation_Component is Measurements and simple assertions made about a patient, device or other subject.
type Observation_Component struct {
	_ *BackboneElement
	// Guidance on how to interpret the value by comparison to a normal or recommended range.
	ReferenceRange []*Observation_ReferenceRange `json:"referenceRange"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueRatio *Ratio `json:"valueRatio"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValuePeriod *Period `json:"valuePeriod"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	// Extensions for valueTime
	ValueTime_ext *Element `json:"_valueTime"`
	// The information determined as a result of making the observation, if the information has a simple value.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ValueDateTime time.Time `json:"valueDateTime"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueString string `json:"valueString"`
	// The information determined as a result of making the observation, if the information has a simple value.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	ValueTime time.Time `json:"valueTime"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueSampledData *SampledData `json:"valueSampledData"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueAttachment *Attachment `json:"valueAttachment"`
	// Extensions for valueDateTime
	ValueDateTime_ext *Element `json:"_valueDateTime"`
	// Provides a reason why the expected value in the element Observation.value[x] is missing.
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason"`
	// The assessment made based on the result of the observation.  Intended as a simple compact code often placed adjacent to the result value in reports and flow sheets to signal the meaning/normalcy status of the result. Otherwise known as abnormal flag.
	Interpretation *CodeableConcept `json:"interpretation"`
	// Describes what was observed. Sometimes this is called the observation "code".
	Code *CodeableConcept `json:"code,omitempty"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueRange *Range `json:"valueRange"`
}

// ResearchStudy is A process where a researcher or organization plans and then executes a series of steps intended to increase the field of healthcare-related knowledge.  This includes studies of safety, efficacy, comparative effectiveness and other information about medications, devices, therapies and other interventional and investigative techniques.  A ResearchStudy involves the gathering of information about human or animal subjects.
type ResearchStudy struct {
	_ *DomainResource
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A description and/or code explaining the premature termination of the study.
	ReasonStopped *CodeableConcept `json:"reasonStopped"`
	// Comments made about the event by the performer, subject or other participants.
	Note []*Annotation `json:"note"`
	// A short, descriptive user-friendly label for the study.
	Title string `json:"title"`
	// Codes categorizing the type of study such as investigational vs. observational, type of blinding, type of randomization, safety vs. efficacy, etc.
	Category []*CodeableConcept `json:"category"`
	// Reference to a Group that defines the criteria for and quantity of subjects participating in the study.  E.g. " 200 female Europeans between the ages of 20 and 45 with early onset diabetes".
	Enrollment []*Reference `json:"enrollment"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Contact details to assist a user in learning more about or engaging with the study.
	Contact []*ContactDetail `json:"contact"`
	// The current state of the study.
	Status string `json:"status"`
	// Indicates a country, state or other region where the study is taking place.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Identifies the start date and the expected (or actual, depending on status) end date for the study.
	Period *Period `json:"period"`
	// The organization responsible for the execution of the study.
	Sponsor *Reference `json:"sponsor"`
	// Indicates the individual who has primary oversite of the execution of the study.
	PrincipalInvestigator *Reference `json:"principalInvestigator"`
	// Describes an expected sequence of events for one of the participants of a study.  E.g. Exposure to drug A, wash-out, exposure to drug B, wash-out, follow-up.
	Arm []*ResearchStudy_Arm `json:"arm"`
	// Identifiers assigned to this research study by the sponsor or other systems.
	Identifier []*Identifier `json:"identifier"`
	// A larger research study of which this particular study is a component or step.
	PartOf []*Reference `json:"partOf"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The condition(s), medication(s), food(s), therapy(ies), device(s) or other concerns or interventions that the study is seeking to gain more information about.
	Focus []*CodeableConcept `json:"focus"`
	// Citations, references and other related documents.
	RelatedArtifact []*RelatedArtifact `json:"relatedArtifact"`
	// Key terms to aid in searching for or filtering the study.
	Keyword []*CodeableConcept `json:"keyword"`
	// A full description of how the study is being conducted.
	Description string `json:"description"`
	// Clinic, hospital or other healthcare location that is participating in the study.
	Site []*Reference `json:"site"`
	// This is a ResearchStudy resource
	ResourceType string `json:"resourceType,omitempty"`
	// The set of steps expected to be performed as part of the execution of the study.
	Protocol []*Reference `json:"protocol"`
}

func NewResearchStudy() *ResearchStudy { return &ResearchStudy{ResourceType: "ResearchStudy"} }

// ContactPoint is Details for all kinds of technology mediated contact points for a person or organization, including telephone, email, etc.
type ContactPoint struct {
	_ *Element
	// Extensions for use
	Use_ext *Element `json:"_use"`
	// Extensions for rank
	Rank_ext *Element `json:"_rank"`
	// Telecommunications form for contact point - what communications system is required to make use of the contact.
	System string `json:"system"`
	// Extensions for system
	System_ext *Element `json:"_system"`
	// The actual contact point details, in a form that is meaningful to the designated communication system (i.e. phone number or email address).
	Value string `json:"value"`
	// Identifies the purpose for the contact point.
	Use string `json:"use"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// Specifies a preferred order in which to use a set of contacts. Contacts are ranked with lower values coming before higher values.
	// pattern [1-9][0-9]*
	Rank uint64 `json:"rank"`
	// Time period when the contact point was/is in use.
	Period *Period `json:"period"`
}

// Account_Coverage is A financial tool for tracking value accrued for a particular purpose.  In the healthcare field, used to track charges for a patient, cost centers, etc.
type Account_Coverage struct {
	_ *BackboneElement
	// The priority of the coverage in the context of this account.
	// pattern [1-9][0-9]*
	Priority uint64 `json:"priority"`
	// Extensions for priority
	Priority_ext *Element `json:"_priority"`
	// The party(s) that are responsible for payment (or part of) of charges applied to this account (including self-pay).
	//
	// A coverage may only be resposible for specific types of charges, and the sequence of the coverages in the account could be important when processing billing.
	Coverage *Reference `json:"coverage,omitempty"`
}

// Measure is The Measure resource provides the definition of a quality measure.
type Measure struct {
	_ *DomainResource
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// A short, descriptive, user-friendly title for the measure.
	Title string `json:"title"`
	// Descriptive topics related to the content of the measure. Topics provide a high-level categorization of the type of the measure that can be useful for filtering and searching.
	Topic []*CodeableConcept `json:"topic"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// A description of the risk adjustment factors that may impact the resulting score for the measure and how they may be accounted for when computing and reporting measure results.
	RiskAdjustment string `json:"riskAdjustment"`
	// Extensions for riskAdjustment
	RiskAdjustment_ext *Element `json:"_riskAdjustment"`
	// Extensions for rateAggregation
	RateAggregation_ext *Element `json:"_rateAggregation"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// A detailed description of how the measure is used from a clinical perspective.
	Usage string `json:"usage"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate measure instances.
	UseContext []*UsageContext `json:"useContext"`
	// Information on whether an increase or decrease in score is the preferred result (e.g., a higher score indicates better quality OR a lower score indicates better quality OR quality is whthin a range).
	ImprovementNotation string `json:"improvementNotation"`
	// Extensions for improvementNotation
	ImprovementNotation_ext *Element `json:"_improvementNotation"`
	// Additional guidance for the measure including how it can be used in a clinical context, and the intent of the measure.
	Guidance string `json:"guidance"`
	// Indicates how the calculation is performed for the measure, including proportion, ratio, continuous variable, and cohort. The value set is extensible, allowing additional measure scoring types to be represented.
	Scoring *CodeableConcept `json:"scoring"`
	// Provides a succint statement of the need for the measure. Usually includes statements pertaining to importance criterion: impact, gap in care, and evidence.
	Rationale string `json:"rationale"`
	// Extensions for definition
	Definition_ext []*Element `json:"_definition"`
	// Extensions for usage
	Usage_ext *Element `json:"_usage"`
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ApprovalDate time.Time `json:"approvalDate"`
	// Notices and disclaimers regarding the use of the measure, or related to intellectual property (such as code systems) referenced by the measure.
	Disclaimer string `json:"disclaimer"`
	// Extensions for approvalDate
	ApprovalDate_ext *Element `json:"_approvalDate"`
	// A reference to a Library resource containing the formal logic used by the measure.
	Library []*Reference `json:"library"`
	// A formal identifier that is used to identify this measure when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []*Identifier `json:"identifier"`
	// The name of the individual or organization that published the measure.
	Publisher string `json:"publisher"`
	// The date  (and optionally time) when the measure was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the measure changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// A group of population criteria for the measure.
	Group []*Measure_Group `json:"group"`
	// Extensions for lastReviewDate
	LastReviewDate_ext *Element `json:"_lastReviewDate"`
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []*RelatedArtifact `json:"relatedArtifact"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// Explaination of why this measure is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// The date on which the resource content was last reviewed. Review happens periodically after approval, but doesn't change the original approval date.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	LastReviewDate time.Time `json:"lastReviewDate"`
	// A legal or geographic region in which the measure is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Extensions for set
	Set_ext *Element `json:"_set"`
	// Indicates whether the measure is used to examine a process, an outcome over time, a patient-reported outcome, or a structure measure such as utilization.
	Type []*CodeableConcept `json:"type"`
	// Extensions for rationale
	Rationale_ext *Element `json:"_rationale"`
	// Extensions for clinicalRecommendationStatement
	ClinicalRecommendationStatement_ext *Element `json:"_clinicalRecommendationStatement"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// A copyright statement relating to the measure and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the measure.
	Copyright string `json:"copyright"`
	// The status of this measure. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for guidance
	Guidance_ext *Element `json:"_guidance"`
	// Extensions for disclaimer
	Disclaimer_ext *Element `json:"_disclaimer"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// A boolean value to indicate that this measure is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The supplemental data criteria for the measure report, specified as either the name of a valid CQL expression within a referenced library, or a valid FHIR Resource Path.
	SupplementalData []*Measure_SupplementalData `json:"supplementalData"`
	// This is a Measure resource
	ResourceType string `json:"resourceType,omitempty"`
	// The identifier that is used to identify this version of the measure when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the measure author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	Version string `json:"version"`
	// Describes how to combine the information calculated, based on logic in each of several populations, into one summarized result.
	RateAggregation string `json:"rateAggregation"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Provides a description of an individual term used within the measure.
	Definition []string `json:"definition"`
	// Provides a summary of relevant clinical guidelines or other clinical recommendations supporting the measure.
	ClinicalRecommendationStatement string `json:"clinicalRecommendationStatement"`
	// An absolute URI that is used to identify this measure when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this measure is (or will be) published. The URL SHOULD include the major version of the measure. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// A natural language name identifying the measure. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// A contributor to the content of the measure, including authors, editors, reviewers, and endorsers.
	Contributor []*Contributor `json:"contributor"`
	// If this is a composite measure, the scoring method used to combine the component measures to determine the composite score.
	CompositeScoring *CodeableConcept `json:"compositeScoring"`
	// The measure set, e.g. Preventive Care and Screening.
	Set string `json:"set"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A free text natural language description of the measure from a consumer's perspective.
	Description string `json:"description"`
	// The period during which the measure content was or is planned to be in active use.
	EffectivePeriod *Period `json:"effectivePeriod"`
}

func NewMeasure() *Measure { return &Measure{ResourceType: "Measure"} }

// Person_Link is Demographics and administrative information about a person independent of a specific health-related context.
type Person_Link struct {
	_ *BackboneElement
	// The resource to which this actual person is associated.
	Target *Reference `json:"target,omitempty"`
	// Level of assurance that this link is actually associated with the target resource.
	Assurance string `json:"assurance"`
	// Extensions for assurance
	Assurance_ext *Element `json:"_assurance"`
}

// PlanDefinition_Condition is This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinition_Condition struct {
	_ *BackboneElement
	// A brief, natural language description of the condition that effectively communicates the intended semantics.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The media type of the language for the expression.
	Language string `json:"language"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
	// An expression that returns true or false, indicating whether or not the condition is satisfied.
	Expression string `json:"expression"`
	// Extensions for expression
	Expression_ext *Element `json:"_expression"`
	// The kind of condition.
	Kind string `json:"kind"`
	// Extensions for kind
	Kind_ext *Element `json:"_kind"`
}

// Slot is A slot of time on a schedule that may be available for booking appointments.
type Slot struct {
	_ *DomainResource
	// The specialty of a practitioner that would be required to perform the service requested in this appointment.
	Specialty []*CodeableConcept `json:"specialty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Date/Time that the slot is to begin.
	Start string `json:"start"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// A broad categorisation of the service that is to be performed during this appointment.
	ServiceCategory *CodeableConcept `json:"serviceCategory"`
	// The type of appointments that can be booked into this slot (ideally this would be an identifiable service - which is at a location, rather than the location itself). If provided then this overrides the value provided on the availability resource.
	ServiceType []*CodeableConcept `json:"serviceType"`
	// Date/Time that the slot is to conclude.
	End string `json:"end"`
	// Comments on the slot to describe any extended information. Such as custom constraints on the slot.
	Comment string `json:"comment"`
	// External Ids for this item.
	Identifier []*Identifier `json:"identifier"`
	// The style of appointment or patient that may be booked in the slot (not service type).
	AppointmentType *CodeableConcept `json:"appointmentType"`
	// This slot has already been overbooked, appointments are unlikely to be accepted for this time.
	Overbooked bool `json:"overbooked"`
	// Extensions for overbooked
	Overbooked_ext *Element `json:"_overbooked"`
	// This is a Slot resource
	ResourceType string `json:"resourceType,omitempty"`
	// The schedule resource that this slot defines an interval of status information.
	Schedule *Reference `json:"schedule,omitempty"`
	// busy | free | busy-unavailable | busy-tentative | entered-in-error.
	Status string `json:"status"`
	// Extensions for start
	Start_ext *Element `json:"_start"`
	// Extensions for end
	End_ext *Element `json:"_end"`
}

func NewSlot() *Slot { return &Slot{ResourceType: "Slot"} }

// StructureMap is A Map of relationships between 2 structures that can be used to transform data.
type StructureMap struct {
	_ *DomainResource
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate structure map instances.
	UseContext []*UsageContext `json:"useContext"`
	// Explaination of why this structure map is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// An absolute URI that is used to identify this structure map when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this structure map is (or will be) published. The URL SHOULD include the major version of the structure map. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// A formal identifier that is used to identify this structure map when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []*Identifier `json:"identifier"`
	// The status of this structure map. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// The date  (and optionally time) when the structure map was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the structure map changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// A structure definition used by this map. The structure definition may describe instances that are converted, or the instances that are produced.
	Structure []*StructureMap_Structure `json:"structure"`
	// The name of the individual or organization that published the structure map.
	Publisher string `json:"publisher"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for import
	Import_ext []*Element `json:"_import"`
	// This is a StructureMap resource
	ResourceType string `json:"resourceType,omitempty"`
	// A natural language name identifying the structure map. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// A short, descriptive, user-friendly title for the structure map.
	Title string `json:"title"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A free text natural language description of the structure map from a consumer's perspective.
	Description string `json:"description"`
	// A legal or geographic region in which the structure map is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// A copyright statement relating to the structure map and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the structure map.
	Copyright string `json:"copyright"`
	// The identifier that is used to identify this version of the structure map when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the structure map author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// A boolean value to indicate that this structure map is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Other maps used by this map (canonical URLs).
	Import []string `json:"import"`
	// Organizes the mapping into managable chunks for human review/ease of maintenance.
	Group []*StructureMap_Group `json:"group,omitempty"`
}

func NewStructureMap() *StructureMap { return &StructureMap{ResourceType: "StructureMap"} }

// Subscription is The subscription resource is used to define a push based subscription from a server to another system. Once a subscription is registered with the server, the server checks every resource that is created or updated, and if the resource matches the given criteria, it sends a message on the defined "channel" so that another system is able to take an appropriate action.
type Subscription struct {
	_ *DomainResource
	// The status of the subscription, which marks the server state for managing the subscription.
	Status string `json:"status"`
	// A record of the last error that occurred when the server processed a notification.
	Error string `json:"error"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A description of why this subscription is defined.
	Reason string `json:"reason"`
	// The rules that the server should use to determine when to generate notifications for this subscription.
	Criteria string `json:"criteria"`
	// Extensions for criteria
	Criteria_ext *Element `json:"_criteria"`
	// Extensions for error
	Error_ext *Element `json:"_error"`
	// This is a Subscription resource
	ResourceType string `json:"resourceType,omitempty"`
	// The time for the server to turn the subscription off.
	End string `json:"end"`
	// Extensions for end
	End_ext *Element `json:"_end"`
	// A tag to add to any resource that matches the criteria, after the subscription is processed.
	Tag []*Coding `json:"tag"`
	// Contact details for a human to contact about the subscription. The primary use of this for system administrator troubleshooting.
	Contact []*ContactPoint `json:"contact"`
	// Extensions for reason
	Reason_ext *Element `json:"_reason"`
	// Details where to send notifications when resources are received that meet the criteria.
	Channel *Subscription_Channel `json:"channel,omitempty"`
}

func NewSubscription() *Subscription { return &Subscription{ResourceType: "Subscription"} }

// TestReport_Assert is A summary of information based on the results of executing a TestScript.
type TestReport_Assert struct {
	_ *BackboneElement
	// An explanatory message associated with the result.
	Message string `json:"message"`
	// Extensions for message
	Message_ext *Element `json:"_message"`
	// A link to further details on the result.
	Detail string `json:"detail"`
	// Extensions for detail
	Detail_ext *Element `json:"_detail"`
	// The result of this assertion.
	Result string `json:"result"`
	// Extensions for result
	Result_ext *Element `json:"_result"`
}

// Encounter_Hospitalization is An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type Encounter_Hospitalization struct {
	_ *BackboneElement
	// Pre-admission identifier.
	PreAdmissionIdentifier *Identifier `json:"preAdmissionIdentifier"`
	// Special courtesies (VIP, board member).
	SpecialCourtesy []*CodeableConcept `json:"specialCourtesy"`
	// Any special requests that have been made for this hospitalization encounter, such as the provision of specific equipment or other things.
	SpecialArrangement []*CodeableConcept `json:"specialArrangement"`
	// Location to which the patient is discharged.
	Destination *Reference `json:"destination"`
	// Category or kind of location after discharge.
	DischargeDisposition *CodeableConcept `json:"dischargeDisposition"`
	// The location from which the patient came before admission.
	Origin *Reference `json:"origin"`
	// From where patient was admitted (physician referral, transfer).
	AdmitSource *CodeableConcept `json:"admitSource"`
	// Whether this hospitalization is a readmission and why if known.
	ReAdmission *CodeableConcept `json:"reAdmission"`
	// Diet preferences reported by the patient.
	DietPreference []*CodeableConcept `json:"dietPreference"`
}

// MedicationRequest_Substitution is An order or request for both supply of the medication and the instructions for administration of the medication to a patient. The resource is called "MedicationRequest" rather than "MedicationPrescription" or "MedicationOrder" to generalize the use across inpatient and outpatient settings, including care plans, etc., and to harmonize with workflow patterns.
type MedicationRequest_Substitution struct {
	_ *BackboneElement
	// True if the prescriber allows a different drug to be dispensed from what was prescribed.
	Allowed bool `json:"allowed"`
	// Extensions for allowed
	Allowed_ext *Element `json:"_allowed"`
	// Indicates the reason for the substitution, or why substitution must or must not be performed.
	Reason *CodeableConcept `json:"reason"`
}

// MessageDefinition_Focus is Defines the characteristics of a message that can be shared between systems, including the type of event that initiates the message, the content to be transmitted and what response(s), if any, are permitted.
type MessageDefinition_Focus struct {
	_ *BackboneElement
	// Identifies the maximum number of resources of this type that must be pointed to by a message in order for it to be valid against this MessageDefinition.
	Max string `json:"max"`
	// Extensions for max
	Max_ext *Element `json:"_max"`
	// The kind of resource that must be the focus for this message.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// A profile that reflects constraints for the focal resource (and potentially for related resources).
	Profile *Reference `json:"profile"`
	// Identifies the minimum number of resources of this type that must be pointed to by a message in order for it to be valid against this MessageDefinition.
	// pattern [0]|([1-9][0-9]*)
	Min uint64 `json:"min"`
	// Extensions for min
	Min_ext *Element `json:"_min"`
}

// Sequence_Variant is Raw data describing a biological sequence.
type Sequence_Variant struct {
	_ *BackboneElement
	// Extensions for start
	Start_ext *Element `json:"_start"`
	// Extensions for observedAllele
	ObservedAllele_ext *Element `json:"_observedAllele"`
	// A pointer to an Observation containing variant information.
	VariantPointer *Reference `json:"variantPointer"`
	// Extended CIGAR string for aligning the sequence with reference bases. See detailed documentation [here](http://support.illumina.com/help/SequencingAnalysisWorkflow/Content/Vault/Informatics/Sequencing_Analysis/CASAVA/swSEQ_mCA_ExtendedCIGARFormat.htm).
	Cigar string `json:"cigar"`
	// Extensions for cigar
	Cigar_ext *Element `json:"_cigar"`
	// Start position of the variant on the  reference sequence.If the coordinate system is either 0-based or 1-based, then start position is inclusive.
	// pattern -?([0]|([1-9][0-9]*))
	Start int64 `json:"start"`
	// End position of the variant on the reference sequence.If the coordinate system is 0-based then end is is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
	// pattern -?([0]|([1-9][0-9]*))
	End int64 `json:"end"`
	// Extensions for end
	End_ext *Element `json:"_end"`
	// An allele is one of a set of coexisting sequence variants of a gene ([SO:0001023](http://www.sequenceontology.org/browser/current_svn/term/SO:0001023)).  Nucleotide(s)/amino acids from start position of sequence to stop position of sequence on the positive (+) strand of the observed  sequence. When the sequence  type is DNA, it should be the sequence on the positive (+) strand. This will lay in the range between variant.start and variant.end.
	ObservedAllele string `json:"observedAllele"`
	// An allele is one of a set of coexisting sequence variants of a gene ([SO:0001023](http://www.sequenceontology.org/browser/current_svn/term/SO:0001023)). Nucleotide(s)/amino acids from start position of sequence to stop position of sequence on the positive (+) strand of the reference sequence. When the sequence  type is DNA, it should be the sequence on the positive (+) strand. This will lay in the range between variant.start and variant.end.
	ReferenceAllele string `json:"referenceAllele"`
	// Extensions for referenceAllele
	ReferenceAllele_ext *Element `json:"_referenceAllele"`
}

// ServiceDefinition is The ServiceDefinition describes a unit of decision support functionality that is made available as a service, such as immunization modules or drug-drug interaction checking.
type ServiceDefinition struct {
	_ *DomainResource
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for lastReviewDate
	LastReviewDate_ext *Element `json:"_lastReviewDate"`
	// A legal or geographic region in which the service definition is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Descriptive topics related to the module. Topics provide a high-level categorization of the module that can be useful for filtering and searching.
	Topic []*CodeableConcept `json:"topic"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// A formal identifier that is used to identify this service definition when it is represented in other formats, or referenced in a specification, model, design or an instance. This is used for CMS or NQF identifiers for a measure artifact. Note that at least one identifier is required for non-experimental active artifacts.
	Identifier []*Identifier `json:"identifier"`
	// A short, descriptive, user-friendly title for the service definition.
	Title string `json:"title"`
	// Extensions for approvalDate
	ApprovalDate_ext *Element `json:"_approvalDate"`
	// The date on which the resource content was last reviewed. Review happens periodically after approval, but doesn't change the original approval date.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	LastReviewDate time.Time `json:"lastReviewDate"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate service definition instances.
	UseContext []*UsageContext `json:"useContext"`
	// A contributor to the content of the module, including authors, editors, reviewers, and endorsers.
	Contributor []*Contributor `json:"contributor"`
	// The period during which the service definition content was or is planned to be in active use.
	EffectivePeriod *Period `json:"effectivePeriod"`
	// The trigger element defines when the rule should be invoked. This information is used by consumers of the rule to determine how to integrate the rule into a specific workflow.
	Trigger []*TriggerDefinition `json:"trigger"`
	// A reference to the operation that is used to invoke this service.
	OperationDefinition *Reference `json:"operationDefinition"`
	// This is a ServiceDefinition resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// The date  (and optionally time) when the service definition was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the service definition changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Explaination of why this service definition is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// A detailed description of how the module is used from a clinical perspective.
	Usage string `json:"usage"`
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ApprovalDate time.Time `json:"approvalDate"`
	// Data requirements are a machine processable description of the data required by the module in order to perform a successful evaluation.
	DataRequirement []*DataRequirement `json:"dataRequirement"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// A natural language name identifying the service definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// The status of this service definition. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// A free text natural language description of the service definition from a consumer's perspective.
	Description string `json:"description"`
	// A copyright statement relating to the service definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the service definition.
	Copyright string `json:"copyright"`
	// An absolute URI that is used to identify this service definition when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this service definition is (or will be) published. The URL SHOULD include the major version of the service definition. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The name of the individual or organization that published the service definition.
	Publisher string `json:"publisher"`
	// The identifier that is used to identify this version of the service definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the service definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// A boolean value to indicate that this service definition is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for usage
	Usage_ext *Element `json:"_usage"`
	// Related resources such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []*RelatedArtifact `json:"relatedArtifact"`
}

func NewServiceDefinition() *ServiceDefinition {
	return &ServiceDefinition{ResourceType: "ServiceDefinition"}
}

// PaymentReconciliation is This resource provides payment details and claim references supporting a bulk payment.
type PaymentReconciliation struct {
	_ *DomainResource
	// Total payment amount.
	Total *Money `json:"total"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// The organization which is responsible for the services rendered to the patient.
	RequestOrganization *Reference `json:"requestOrganization"`
	// List of individual settlement amounts and the corresponding transaction.
	Detail []*PaymentReconciliation_Detail `json:"detail"`
	// The form to be used for printing the content.
	Form *CodeableConcept `json:"form"`
	// The period of time for which payments have been gathered into this bulk payment for settlement.
	Period *Period `json:"period"`
	// Transaction status: error, complete.
	Outcome *CodeableConcept `json:"outcome"`
	// A description of the status of the adjudication.
	Disposition string `json:"disposition"`
	// The date when the enclosed suite of services were performed or completed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// Original request resource reference.
	Request *Reference `json:"request"`
	// The practitioner who is responsible for the services rendered to the patient.
	RequestProvider *Reference `json:"requestProvider"`
	// Suite of notes.
	ProcessNote []*PaymentReconciliation_ProcessNote `json:"processNote"`
	// The Insurer who produced this adjudicated response.
	Organization *Reference `json:"organization"`
	// Extensions for disposition
	Disposition_ext *Element `json:"_disposition"`
	// This is a PaymentReconciliation resource
	ResourceType string `json:"resourceType,omitempty"`
	// The Response business identifier.
	Identifier []*Identifier `json:"identifier"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
}

func NewPaymentReconciliation() *PaymentReconciliation {
	return &PaymentReconciliation{ResourceType: "PaymentReconciliation"}
}

// ElementDefinition is Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition struct {
	_ *Element
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// Extensions for fixedInteger
	FixedInteger_ext *Element `json:"_fixedInteger"`
	// Extensions for fixedDateTime
	FixedDateTime_ext *Element `json:"_fixedDateTime"`
	// The minimum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	MinValueDecimal float64 `json:"minValueDecimal"`
	// Formal constraints such as co-occurrence and other constraints that can be computationally evaluated within the context of the instance.
	Constraint []*ElementDefinition_Constraint `json:"constraint"`
	// The minimum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	MinValueTime time.Time `json:"minValueTime"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueElement *Element `json:"defaultValueElement"`
	// Extensions for fixedBoolean
	FixedBoolean_ext *Element `json:"_fixedBoolean"`
	// Extensions for patternDateTime
	PatternDateTime_ext *Element `json:"_patternDateTime"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	PatternTime time.Time `json:"patternTime"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternPeriod *Period `json:"patternPeriod"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternRelatedArtifact *RelatedArtifact `json:"patternRelatedArtifact"`
	// If true, the value of this element affects the interpretation of the element or resource that contains it, and the value of the element cannot be ignored. Typically, this is used for status, negation and qualification codes. The effect of this is that the element cannot be ignored by systems: they SHALL either recognize the element and process it, and/or a pre-determination has been made that it is not relevant to their particular system.
	IsModifier bool `json:"isModifier"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueMarkdown string `json:"defaultValueMarkdown"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueAge *Age `json:"defaultValueAge"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	FixedOid string `json:"fixedOid"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedParameterDefinition *ParameterDefinition `json:"fixedParameterDefinition"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternAnnotation *Annotation `json:"patternAnnotation"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternAttachment *Attachment `json:"patternAttachment"`
	// The maximum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	MaxValueDecimal float64 `json:"maxValueDecimal"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueInstant string `json:"defaultValueInstant"`
	// Extensions for defaultValuePositiveInt
	DefaultValuePositiveInt_ext *Element `json:"_defaultValuePositiveInt"`
	// Extensions for fixedCode
	FixedCode_ext *Element `json:"_fixedCode"`
	// Extensions for fixedMarkdown
	FixedMarkdown_ext *Element `json:"_fixedMarkdown"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternUri string `json:"patternUri"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternRange *Range `json:"patternRange"`
	// Indicates that the element is sliced into a set of alternative definitions (i.e. in a structure definition, there are multiple different constraints on a single element in the base resource). Slicing can be used in any resource that has cardinality ..* on the base resource, or any resource with a choice of types. The set of slices is any elements that come after this in the element sequence that have the same path, until a shorter path occurs (the shorter path terminates the set).
	Slicing *ElementDefinition_Slicing `json:"slicing"`
	// If present, indicates that the order of the repeating element has meaning and describes what that meaning is.  If absent, it means that the order of the element has no meaning.
	OrderMeaning string `json:"orderMeaning"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	PatternDate time.Time `json:"patternDate"`
	// Extensions for minValueDecimal
	MinValueDecimal_ext *Element `json:"_minValueDecimal"`
	// The maximum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	MaxValueDateTime time.Time `json:"maxValueDateTime"`
	// Extensions for alias
	Alias_ext []*Element `json:"_alias"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueExtension *Extension `json:"defaultValueExtension"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedString string `json:"fixedString"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedExtension *Extension `json:"fixedExtension"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternTriggerDefinition *TriggerDefinition `json:"patternTriggerDefinition"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern -?([0]|([1-9][0-9]*))
	FixedInteger int64 `json:"fixedInteger"`
	// Extensions for patternBase64Binary
	PatternBase64Binary_ext *Element `json:"_patternBase64Binary"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternContactDetail *ContactDetail `json:"patternContactDetail"`
	// The maximum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	MaxValueQuantity *Quantity `json:"maxValueQuantity"`
	// Extensions for condition
	Condition_ext []*Element `json:"_condition"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	FixedDate time.Time `json:"fixedDate"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternQuantity *Quantity `json:"patternQuantity"`
	// The minimum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	MinValueInstant string `json:"minValueInstant"`
	// Extensions for maxValuePositiveInt
	MaxValuePositiveInt_ext *Element `json:"_maxValuePositiveInt"`
	// Identifies the identity of an element defined elsewhere in the profile whose content rules should be applied to the current element.
	ContentReference string `json:"contentReference"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueNarrative *Narrative `json:"defaultValueNarrative"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueAddress *Address `json:"defaultValueAddress"`
	// Extensions for fixedDate
	FixedDate_ext *Element `json:"_fixedDate"`
	// Extensions for patternString
	PatternString_ext *Element `json:"_patternString"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern [A-Za-z0-9\-\.]{1,64}
	PatternId string `json:"patternId"`
	// The name of this element definition slice, when slicing is working. The name must be a token with no dots or spaces. This is a unique name referring to a specific set of constraints applied to this element, used to provide a name to different slices of the same element.
	SliceName string `json:"sliceName"`
	// Extensions for defaultValueMarkdown
	DefaultValueMarkdown_ext *Element `json:"_defaultValueMarkdown"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternUsageContext *UsageContext `json:"patternUsageContext"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueDataRequirement *DataRequirement `json:"defaultValueDataRequirement"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern [0]|([1-9][0-9]*)
	FixedUnsignedInt uint64 `json:"fixedUnsignedInt"`
	// Extensions for mustSupport
	MustSupport_ext *Element `json:"_mustSupport"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedDuration *Duration `json:"fixedDuration"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedDataRequirement *DataRequirement `json:"fixedDataRequirement"`
	// Extensions for patternId
	PatternId_ext *Element `json:"_patternId"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern [1-9][0-9]*
	PatternPositiveInt uint64 `json:"patternPositiveInt"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueString string `json:"defaultValueString"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	DefaultValueUuid string `json:"defaultValueUuid"`
	// Extensions for defaultValueUuid
	DefaultValueUuid_ext *Element `json:"_defaultValueUuid"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	FixedUuid string `json:"fixedUuid"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedRange *Range `json:"fixedRange"`
	// Extensions for patternDecimal
	PatternDecimal_ext *Element `json:"_patternDecimal"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternAddress *Address `json:"patternAddress"`
	// The maximum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern -?([0]|([1-9][0-9]*))
	MaxValueInteger int64 `json:"maxValueInteger"`
	// Codes that define how this element is represented in instances, when the deviation varies from the normal case.
	Representation []string `json:"representation"`
	// Extensions for meaningWhenMissing
	MeaningWhenMissing_ext *Element `json:"_meaningWhenMissing"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternMarkdown string `json:"patternMarkdown"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternCount *Count `json:"patternCount"`
	// Extensions for maxValueInteger
	MaxValueInteger_ext *Element `json:"_maxValueInteger"`
	// Identifies additional names by which this element might also be known.
	Alias []string `json:"alias"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueDistance *Distance `json:"defaultValueDistance"`
	// Extensions for fixedUri
	FixedUri_ext *Element `json:"_fixedUri"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedMarkdown string `json:"fixedMarkdown"`
	// Extensions for patternUuid
	PatternUuid_ext *Element `json:"_patternUuid"`
	// The Implicit meaning that is to be understood when this element is missing (e.g. 'when this element is missing, the period is ongoing'.
	MeaningWhenMissing string `json:"meaningWhenMissing"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	PatternUuid string `json:"patternUuid"`
	// Extensions for isSummary
	IsSummary_ext *Element `json:"_isSummary"`
	// A code that has the same meaning as the element in a particular terminology.
	Code []*Coding `json:"code"`
	// Extensions for definition
	Definition_ext *Element `json:"_definition"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedElement *Element `json:"fixedElement"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedCount *Count `json:"fixedCount"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	PatternOid string `json:"patternOid"`
	// Extensions for patternOid
	PatternOid_ext *Element `json:"_patternOid"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedContactPoint *ContactPoint `json:"fixedContactPoint"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternBoolean bool `json:"patternBoolean"`
	// Extensions for short
	Short_ext *Element `json:"_short"`
	// Extensions for requirements
	Requirements_ext *Element `json:"_requirements"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueSimpleQuantity *Quantity `json:"defaultValueSimpleQuantity"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueSampledData *SampledData `json:"defaultValueSampledData"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	FixedDateTime time.Time `json:"fixedDateTime"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedCoding *Coding `json:"fixedCoding"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternHumanName *HumanName `json:"patternHumanName"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueBase64Binary string `json:"defaultValueBase64Binary"`
	// Extensions for defaultValueBase64Binary
	DefaultValueBase64Binary_ext *Element `json:"_defaultValueBase64Binary"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueContributor *Contributor `json:"defaultValueContributor"`
	// Indicates the maximum length in characters that is permitted to be present in conformant instances and which is expected to be supported by conformant consumers that support the element.
	// pattern -?([0]|([1-9][0-9]*))
	MaxLength int64 `json:"maxLength"`
	// Extensions for isModifier
	IsModifier_ext *Element `json:"_isModifier"`
	// The data type or resource that the value of this element is permitted to be.
	Type []*ElementDefinition_Type `json:"type"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedBoolean bool `json:"fixedBoolean"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	FixedDecimal float64 `json:"fixedDecimal"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternParameterDefinition *ParameterDefinition `json:"patternParameterDefinition"`
	// Extensions for maxLength
	MaxLength_ext *Element `json:"_maxLength"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern [A-Za-z0-9\-\.]{1,64}
	FixedId string `json:"fixedId"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedSimpleQuantity *Quantity `json:"fixedSimpleQuantity"`
	// Extensions for patternTime
	PatternTime_ext *Element `json:"_patternTime"`
	// Extensions for patternUnsignedInt
	PatternUnsignedInt_ext *Element `json:"_patternUnsignedInt"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern [^\s]+([\s]?[^\s]+)*
	DefaultValueCode string `json:"defaultValueCode"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueMoney *Money `json:"defaultValueMoney"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedAddress *Address `json:"fixedAddress"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternBase64Binary string `json:"patternBase64Binary"`
	// Extensions for defaultValueInteger
	DefaultValueInteger_ext *Element `json:"_defaultValueInteger"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedReference *Reference `json:"fixedReference"`
	// Extensions for minValuePositiveInt
	MinValuePositiveInt_ext *Element `json:"_minValuePositiveInt"`
	// Extensions for maxValueInstant
	MaxValueInstant_ext *Element `json:"_maxValueInstant"`
	// The maximum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	MaxValueTime time.Time `json:"maxValueTime"`
	// Extensions for defaultValueUri
	DefaultValueUri_ext *Element `json:"_defaultValueUri"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueHumanName *HumanName `json:"defaultValueHumanName"`
	// Extensions for fixedUnsignedInt
	FixedUnsignedInt_ext *Element `json:"_fixedUnsignedInt"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternCodeableConcept *CodeableConcept `json:"patternCodeableConcept"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternReference *Reference `json:"patternReference"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedElementDefinition *ElementDefinition `json:"fixedElementDefinition"`
	// Extensions for patternDate
	PatternDate_ext *Element `json:"_patternDate"`
	// Extensions for defaultValueDate
	DefaultValueDate_ext *Element `json:"_defaultValueDate"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern [1-9][0-9]*
	FixedPositiveInt uint64 `json:"fixedPositiveInt"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueReference *Reference `json:"defaultValueReference"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternDataRequirement *DataRequirement `json:"patternDataRequirement"`
	// Whether the element should be included if a client requests a search with the parameter _summary=true.
	IsSummary bool `json:"isSummary"`
	// Extensions for minValueDate
	MinValueDate_ext *Element `json:"_minValueDate"`
	// The minimum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	MinValueQuantity *Quantity `json:"minValueQuantity"`
	// The minimum number of times this element SHALL appear in the instance.
	// pattern [0]|([1-9][0-9]*)
	Min uint64 `json:"min"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedUri string `json:"fixedUri"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedDistance *Distance `json:"fixedDistance"`
	// Extensions for patternInteger
	PatternInteger_ext *Element `json:"_patternInteger"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternDistance *Distance `json:"patternDistance"`
	// A sample value for this element demonstrating the type of information that would typically be found in the element.
	Example []*ElementDefinition_Example `json:"example"`
	// Extensions for fixedId
	FixedId_ext *Element `json:"_fixedId"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedNarrative *Narrative `json:"fixedNarrative"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternElement *Element `json:"patternElement"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternContributor *Contributor `json:"patternContributor"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	DefaultValueDecimal float64 `json:"defaultValueDecimal"`
	// Extensions for defaultValueUnsignedInt
	DefaultValueUnsignedInt_ext *Element `json:"_defaultValueUnsignedInt"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueBackboneElement *BackboneElement `json:"defaultValueBackboneElement"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedQuantity *Quantity `json:"fixedQuantity"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedDosage *Dosage `json:"fixedDosage"`
	// Extensions for patternUri
	PatternUri_ext *Element `json:"_patternUri"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueAttachment *Attachment `json:"defaultValueAttachment"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedRelatedArtifact *RelatedArtifact `json:"fixedRelatedArtifact"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternInstant string `json:"patternInstant"`
	// Extensions for minValueTime
	MinValueTime_ext *Element `json:"_minValueTime"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	DefaultValueOid string `json:"defaultValueOid"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern -?([0]|([1-9][0-9]*))
	PatternInteger int64 `json:"patternInteger"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternElementDefinition *ElementDefinition `json:"patternElementDefinition"`
	// Extensions for minValueInteger
	MinValueInteger_ext *Element `json:"_minValueInteger"`
	// The maximum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	MaxValueInstant string `json:"maxValueInstant"`
	// Explanatory notes and implementation guidance about the data element, including notes about how to use the data properly, exceptions to proper use, etc.
	Comment string `json:"comment"`
	// Extensions for defaultValueDecimal
	DefaultValueDecimal_ext *Element `json:"_defaultValueDecimal"`
	// Extensions for fixedString
	FixedString_ext *Element `json:"_fixedString"`
	// Extensions for maxValueUnsignedInt
	MaxValueUnsignedInt_ext *Element `json:"_maxValueUnsignedInt"`
	// Extensions for contentReference
	ContentReference_ext *Element `json:"_contentReference"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueParameterDefinition *ParameterDefinition `json:"defaultValueParameterDefinition"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	PatternDateTime time.Time `json:"patternDateTime"`
	// The minimum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern [0]|([1-9][0-9]*)
	MinValueUnsignedInt uint64 `json:"minValueUnsignedInt"`
	// Provides a complete explanation of the meaning of the data element for human readability.  For the case of elements derived from existing elements (e.g. constraints), the definition SHALL be consistent with the base definition, but convey the meaning of the element in the particular context of use of the resource.
	Definition string `json:"definition"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueCoding *Coding `json:"defaultValueCoding"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedSignature *Signature `json:"fixedSignature"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedContactDetail *ContactDetail `json:"fixedContactDetail"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern [0]|([1-9][0-9]*)
	PatternUnsignedInt uint64 `json:"patternUnsignedInt"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternNarrative *Narrative `json:"patternNarrative"`
	// Extensions for min
	Min_ext *Element `json:"_min"`
	// Information about the base definition of the element, provided to make it unnecessary for tools to trace the deviation of the element through the derived and related profiles. This information is provided when the element definition is not the original definition of an element - i.g. either in a constraint on another type, or for elements from a super type in a snap shot.
	Base *ElementDefinition_Base `json:"base"`
	// Extensions for defaultValueTime
	DefaultValueTime_ext *Element `json:"_defaultValueTime"`
	// The minimum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	MinValueDateTime time.Time `json:"minValueDateTime"`
	// Extensions for defaultValueInstant
	DefaultValueInstant_ext *Element `json:"_defaultValueInstant"`
	// Extensions for orderMeaning
	OrderMeaning_ext *Element `json:"_orderMeaning"`
	// Extensions for fixedInstant
	FixedInstant_ext *Element `json:"_fixedInstant"`
	// Extensions for patternCode
	PatternCode_ext *Element `json:"_patternCode"`
	// Extensions for defaultValueId
	DefaultValueId_ext *Element `json:"_defaultValueId"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueRange *Range `json:"defaultValueRange"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueRatio *Ratio `json:"defaultValueRatio"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternCoding *Coding `json:"patternCoding"`
	// The maximum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	MaxValueDate time.Time `json:"maxValueDate"`
	// Extensions for maxValueDateTime
	MaxValueDateTime_ext *Element `json:"_maxValueDateTime"`
	// A concise description of what this element means (e.g. for use in autogenerated summaries).
	Short string `json:"short"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedAttachment *Attachment `json:"fixedAttachment"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedHumanName *HumanName `json:"fixedHumanName"`
	// Extensions for patternMarkdown
	PatternMarkdown_ext *Element `json:"_patternMarkdown"`
	// The minimum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern [1-9][0-9]*
	MinValuePositiveInt uint64 `json:"minValuePositiveInt"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// Extensions for defaultValueOid
	DefaultValueOid_ext *Element `json:"_defaultValueOid"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedAnnotation *Annotation `json:"fixedAnnotation"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern [0]|([1-9][0-9]*)
	DefaultValueUnsignedInt uint64 `json:"defaultValueUnsignedInt"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueDuration *Duration `json:"defaultValueDuration"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedAge *Age `json:"fixedAge"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueContactPoint *ContactPoint `json:"defaultValueContactPoint"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueDosage *Dosage `json:"defaultValueDosage"`
	// Extensions for fixedOid
	FixedOid_ext *Element `json:"_fixedOid"`
	// Extensions for maxValueTime
	MaxValueTime_ext *Element `json:"_maxValueTime"`
	// Extensions for fixedBase64Binary
	FixedBase64Binary_ext *Element `json:"_fixedBase64Binary"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedSampledData *SampledData `json:"fixedSampledData"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern [^\s]+([\s]?[^\s]+)*
	PatternCode string `json:"patternCode"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternRatio *Ratio `json:"patternRatio"`
	// The maximum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern [1-9][0-9]*
	MaxValuePositiveInt uint64 `json:"maxValuePositiveInt"`
	// Extensions for defaultValueBoolean
	DefaultValueBoolean_ext *Element `json:"_defaultValueBoolean"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueSignature *Signature `json:"defaultValueSignature"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedBackboneElement *BackboneElement `json:"fixedBackboneElement"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	PatternDecimal float64 `json:"patternDecimal"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternString string `json:"patternString"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternSignature *Signature `json:"patternSignature"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedCodeableConcept *CodeableConcept `json:"fixedCodeableConcept"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternMeta *Meta `json:"patternMeta"`
	// The maximum number of times this element is permitted to appear in the instance.
	Max string `json:"max"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	DefaultValueDate time.Time `json:"defaultValueDate"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern [A-Za-z0-9\-\.]{1,64}
	DefaultValueId string `json:"defaultValueId"`
	// A reference to an invariant that may make additional statements about the cardinality or value in the instance.
	Condition []string `json:"condition"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	DefaultValueTime time.Time `json:"defaultValueTime"`
	// Extensions for fixedPositiveInt
	FixedPositiveInt_ext *Element `json:"_fixedPositiveInt"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternIdentifier *Identifier `json:"patternIdentifier"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternContactPoint *ContactPoint `json:"patternContactPoint"`
	// The minimum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern -?([0]|([1-9][0-9]*))
	MinValueInteger int64 `json:"minValueInteger"`
	// Extensions for maxValueDecimal
	MaxValueDecimal_ext *Element `json:"_maxValueDecimal"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternExtension *Extension `json:"patternExtension"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternDuration *Duration `json:"patternDuration"`
	// Extensions for max
	Max_ext *Element `json:"_max"`
	// Extensions for defaultValueString
	DefaultValueString_ext *Element `json:"_defaultValueString"`
	// Extensions for fixedDecimal
	FixedDecimal_ext *Element `json:"_fixedDecimal"`
	// Extensions for fixedUuid
	FixedUuid_ext *Element `json:"_fixedUuid"`
	// Extensions for patternInstant
	PatternInstant_ext *Element `json:"_patternInstant"`
	// Extensions for patternPositiveInt
	PatternPositiveInt_ext *Element `json:"_patternPositiveInt"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternTiming *Timing `json:"patternTiming"`
	// Extensions for minValueInstant
	MinValueInstant_ext *Element `json:"_minValueInstant"`
	// The maximum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern [0]|([1-9][0-9]*)
	MaxValueUnsignedInt uint64 `json:"maxValueUnsignedInt"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedInstant string `json:"fixedInstant"`
	// Extensions for fixedTime
	FixedTime_ext *Element `json:"_fixedTime"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedPeriod *Period `json:"fixedPeriod"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedMeta *Meta `json:"fixedMeta"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternSimpleQuantity *Quantity `json:"patternSimpleQuantity"`
	// Extensions for defaultValueDateTime
	DefaultValueDateTime_ext *Element `json:"_defaultValueDateTime"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueUsageContext *UsageContext `json:"defaultValueUsageContext"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	FixedTime time.Time `json:"fixedTime"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedTiming *Timing `json:"fixedTiming"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedUsageContext *UsageContext `json:"fixedUsageContext"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedTriggerDefinition *TriggerDefinition `json:"fixedTriggerDefinition"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueElementDefinition *ElementDefinition `json:"defaultValueElementDefinition"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueContactDetail *ContactDetail `json:"defaultValueContactDetail"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternDosage *Dosage `json:"patternDosage"`
	// The path identifies the element and is expressed as a "."-separated list of ancestor elements, beginning with the name of the resource or extension.
	Path string `json:"path"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueUri string `json:"defaultValueUri"`
	// Extensions for defaultValueCode
	DefaultValueCode_ext *Element `json:"_defaultValueCode"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern [1-9][0-9]*
	DefaultValuePositiveInt uint64 `json:"defaultValuePositiveInt"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueCount *Count `json:"defaultValueCount"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedRatio *Ratio `json:"fixedRatio"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueCodeableConcept *CodeableConcept `json:"defaultValueCodeableConcept"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueRelatedArtifact *RelatedArtifact `json:"defaultValueRelatedArtifact"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternMoney *Money `json:"patternMoney"`
	// Extensions for minValueUnsignedInt
	MinValueUnsignedInt_ext *Element `json:"_minValueUnsignedInt"`
	// This element is for traceability of why the element was created and why the constraints exist as they do. This may be used to point to source materials or specifications that drove the structure of this element.
	Requirements string `json:"requirements"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueBoolean bool `json:"defaultValueBoolean"`
	// If true, implementations that produce or consume resources SHALL provide "support" for the element in some meaningful way.  If false, the element may be ignored and not supported.
	MustSupport bool `json:"mustSupport"`
	// Identifies a concept from an external specification that roughly corresponds to this element.
	Mapping []*ElementDefinition_Mapping `json:"mapping"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedContributor *Contributor `json:"fixedContributor"`
	// Extensions for representation
	Representation_ext []*Element `json:"_representation"`
	// Extensions for sliceName
	SliceName_ext *Element `json:"_sliceName"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern -?([0]|([1-9][0-9]*))
	DefaultValueInteger int64 `json:"defaultValueInteger"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueTiming *Timing `json:"defaultValueTiming"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedBase64Binary string `json:"fixedBase64Binary"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	// pattern [^\s]+([\s]?[^\s]+)*
	FixedCode string `json:"fixedCode"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	DefaultValueDateTime time.Time `json:"defaultValueDateTime"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueTriggerDefinition *TriggerDefinition `json:"defaultValueTriggerDefinition"`
	// Extensions for label
	Label_ext *Element `json:"_label"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedIdentifier *Identifier `json:"fixedIdentifier"`
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	FixedMoney *Money `json:"fixedMoney"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternBackboneElement *BackboneElement `json:"patternBackboneElement"`
	// A single preferred label which is the text to display beside the element indicating its meaning or to use to prompt for the element in a user display or form.
	Label string `json:"label"`
	// The minimum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	MinValueDate time.Time `json:"minValueDate"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueIdentifier *Identifier `json:"defaultValueIdentifier"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternSampledData *SampledData `json:"patternSampledData"`
	// Extensions for maxValueDate
	MaxValueDate_ext *Element `json:"_maxValueDate"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueAnnotation *Annotation `json:"defaultValueAnnotation"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueQuantity *Quantity `json:"defaultValueQuantity"`
	// Extensions for patternBoolean
	PatternBoolean_ext *Element `json:"_patternBoolean"`
	// Extensions for minValueDateTime
	MinValueDateTime_ext *Element `json:"_minValueDateTime"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValueMeta *Meta `json:"defaultValueMeta"`
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.  The values of elements present in the pattern must match exactly (case-sensitive, accent-sensitive, etc.).
	PatternAge *Age `json:"patternAge"`
	// Binds to a value set if this element is coded (code, Coding, CodeableConcept, Quantity), or the data types (string, uri).
	Binding *ElementDefinition_Binding `json:"binding"`
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValuePeriod *Period `json:"defaultValuePeriod"`
}

// Bundle_Link is A container for a collection of resources.
type Bundle_Link struct {
	_ *BackboneElement
	// A name which details the functional use for this link - see [http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1](http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1).
	Relation string `json:"relation"`
	// Extensions for relation
	Relation_ext *Element `json:"_relation"`
	// The reference details for the link.
	Url string `json:"url"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
}

// ClaimResponse_AddItem is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse_AddItem struct {
	_ *BackboneElement
	// The type of reveneu or cost center providing the product and/or service.
	Revenue *CodeableConcept `json:"revenue"`
	// The adjudications results.
	Adjudication []*ClaimResponse_Adjudication `json:"adjudication"`
	// Health Care Service Type Codes  to identify the classification of service or benefits.
	Category *CodeableConcept `json:"category"`
	// A code to indicate the Professional Service or Product supplied.
	Service *CodeableConcept `json:"service"`
	// Item typification or modifiers codes, eg for Oral whether the treatment is cosmetic or associated with TMJ, or for medical whether the treatment was outside the clinic or out of office hours.
	Modifier []*CodeableConcept `json:"modifier"`
	// The fee charged for the professional service or product..
	Fee *Money `json:"fee"`
	// A list of note references to the notes provided below.
	NoteNumber []uint64 `json:"noteNumber"`
	// Extensions for noteNumber
	NoteNumber_ext []*Element `json:"_noteNumber"`
	// List of input service items which this service line is intended to replace.
	SequenceLinkId []uint64 `json:"sequenceLinkId"`
	// Extensions for sequenceLinkId
	SequenceLinkId_ext []*Element `json:"_sequenceLinkId"`
	// The second tier service adjudications for payor added services.
	Detail []*ClaimResponse_Detail1 `json:"detail"`
}

// Library is The Library resource is a general-purpose container for knowledge asset definitions. It can be used to describe and expose existing knowledge assets such as logic libraries and information model descriptions, as well as to describe a collection of knowledge assets.
type Library struct {
	_ *DomainResource
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ApprovalDate time.Time `json:"approvalDate"`
	// The date on which the resource content was last reviewed. Review happens periodically after approval, but doesn't change the original approval date.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	LastReviewDate time.Time `json:"lastReviewDate"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate library instances.
	UseContext []*UsageContext `json:"useContext"`
	// Descriptive topics related to the content of the library. Topics provide a high-level categorization of the library that can be useful for filtering and searching.
	Topic []*CodeableConcept `json:"topic"`
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []*RelatedArtifact `json:"relatedArtifact"`
	// The identifier that is used to identify this version of the library when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the library author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	Version string `json:"version"`
	// A natural language name identifying the library. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// The name of the individual or organization that published the library.
	Publisher string `json:"publisher"`
	// A short, descriptive, user-friendly title for the library.
	Title string `json:"title"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for lastReviewDate
	LastReviewDate_ext *Element `json:"_lastReviewDate"`
	// Identifies the type of library such as a Logic Library, Model Definition, Asset Collection, or Module Definition.
	Type *CodeableConcept `json:"type,omitempty"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A formal identifier that is used to identify this library when it is represented in other formats, or referenced in a specification, model, design or an instance. e.g. CMS or NQF identifiers for a measure artifact. Note that at least one identifier is required for non-experimental active artifacts.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// A detailed description of how the library is used from a clinical perspective.
	Usage string `json:"usage"`
	// The period during which the library content was or is planned to be in active use.
	EffectivePeriod *Period `json:"effectivePeriod"`
	// A legal or geographic region in which the library is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// The parameter element defines parameters used by the library.
	Parameter []*ParameterDefinition `json:"parameter"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Describes a set of data that must be provided in order to be able to successfully perform the computations defined by the library.
	DataRequirement []*DataRequirement `json:"dataRequirement"`
	// An absolute URI that is used to identify this library when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this library is (or will be) published. The URL SHOULD include the major version of the library. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// The date  (and optionally time) when the library was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the library changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for usage
	Usage_ext *Element `json:"_usage"`
	// Explaination of why this library is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// Extensions for approvalDate
	ApprovalDate_ext *Element `json:"_approvalDate"`
	// A contributor to the content of the library, including authors, editors, reviewers, and endorsers.
	Contributor []*Contributor `json:"contributor"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// This is a Library resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// A free text natural language description of the library from a consumer's perspective.
	Description string `json:"description"`
	// A copyright statement relating to the library and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the library.
	Copyright string `json:"copyright"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// The content of the library as an Attachment. The content may be a reference to a url, or may be directly embedded as a base-64 string. Either way, the contentType of the attachment determines how to interpret the content.
	Content []*Attachment `json:"content"`
	// The status of this library. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// A boolean value to indicate that this library is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
}

func NewLibrary() *Library { return &Library{ResourceType: "Library"} }

// DataRequirement_CodeFilter is Describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type DataRequirement_CodeFilter struct {
	_ *BackboneElement
	// The valueset for the code filter. The valueSet and value elements are exclusive. If valueSet is specified, the filter will return only those data items for which the value of the code-valued element specified in the path is a member of the specified valueset.
	ValueSetString string `json:"valueSetString"`
	// Extensions for valueSetString
	ValueSetString_ext *Element `json:"_valueSetString"`
	// The valueset for the code filter. The valueSet and value elements are exclusive. If valueSet is specified, the filter will return only those data items for which the value of the code-valued element specified in the path is a member of the specified valueset.
	ValueSetReference *Reference `json:"valueSetReference"`
	// Extensions for valueCode
	ValueCode_ext []*Element `json:"_valueCode"`
	// The code-valued attribute of the filter. The specified path must be resolvable from the type of the required data. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements. Note that the index must be an integer constant. The path must resolve to an element of type code, Coding, or CodeableConcept.
	Path string `json:"path"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// The codes for the code filter. Only one of valueSet, valueCode, valueCoding, or valueCodeableConcept may be specified. If values are given, the filter will return only those data items for which the code-valued attribute specified by the path has a value that is one of the specified codes.
	ValueCode []string `json:"valueCode"`
	// The Codings for the code filter. Only one of valueSet, valueCode, valueConding, or valueCodeableConcept may be specified. If values are given, the filter will return only those data items for which the code-valued attribute specified by the path has a value that is one of the specified Codings.
	ValueCoding []*Coding `json:"valueCoding"`
	// The CodeableConcepts for the code filter. Only one of valueSet, valueCode, valueConding, or valueCodeableConcept may be specified. If values are given, the filter will return only those data items for which the code-valued attribute specified by the path has a value that is one of the specified CodeableConcepts.
	ValueCodeableConcept []*CodeableConcept `json:"valueCodeableConcept"`
}

// ChargeItem is The resource ChargeItem describes the provision of healthcare provider products for a certain patient, therefore referring not only to the product, but containing in addition details of the provision, like date, time, amounts and participating organizations and persons. Main Usage of the ChargeItem is to enable the billing process and internal cost allocation.
type ChargeItem struct {
	_ *DomainResource
	// The organization requesting the service.
	PerformingOrganization *Reference `json:"performingOrganization"`
	// The individual or set of individuals the action is being or was performed on.
	Subject *Reference `json:"subject,omitempty"`
	// Extensions for factorOverride
	FactorOverride_ext *Element `json:"_factorOverride"`
	// Extensions for enteredDate
	EnteredDate_ext *Element `json:"_enteredDate"`
	// Date/time(s) or duration when the charged service was applied.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OccurrenceDateTime time.Time `json:"occurrenceDateTime"`
	// The encounter or episode of care that establishes the context for this event.
	Context *Reference `json:"context"`
	// Date/time(s) or duration when the charged service was applied.
	OccurrencePeriod *Period `json:"occurrencePeriod"`
	// Further information supporting the this charge.
	SupportingInformation []*Reference `json:"supportingInformation"`
	// References the source of pricing information, rules of application for the code this ChargeItem uses.
	Definition []string `json:"definition"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Date/time(s) or duration when the charged service was applied.
	OccurrenceTiming *Timing `json:"occurrenceTiming"`
	// Quantity of which the charge item has been serviced.
	Quantity *Quantity `json:"quantity"`
	// The anatomical location where the related service has been applied.
	Bodysite []*CodeableConcept `json:"bodysite"`
	// Factor overriding the factor determined by the rules associated with the code.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	FactorOverride float64 `json:"factorOverride"`
	// The device, practitioner, etc. who entered the charge item.
	Enterer *Reference `json:"enterer"`
	// Comments made about the event by the performer, subject or other participants.
	Note []*Annotation `json:"note"`
	// Identifiers assigned to this event performer or other systems.
	Identifier *Identifier `json:"identifier"`
	// Total price of the charge overriding the list price associated with the code.
	PriceOverride *Money `json:"priceOverride"`
	// Indicates who or what performed or participated in the charged service.
	Participant []*ChargeItem_Participant `json:"participant"`
	// Extensions for overrideReason
	OverrideReason_ext *Element `json:"_overrideReason"`
	// Date the charge item was entered.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	EnteredDate time.Time `json:"enteredDate"`
	// Describes why the event occurred in coded or textual form.
	Reason []*CodeableConcept `json:"reason"`
	// Indicated the rendered service that caused this charge.
	Service []*Reference `json:"service"`
	// ChargeItems can be grouped to larger ChargeItems covering the whole set.
	PartOf []*Reference `json:"partOf"`
	// The organization performing the service.
	RequestingOrganization *Reference `json:"requestingOrganization"`
	// The current state of the ChargeItem.
	Status string `json:"status"`
	// Extensions for definition
	Definition_ext []*Element `json:"_definition"`
	// A code that identifies the charge, like a billing code.
	Code *CodeableConcept `json:"code,omitempty"`
	// Extensions for occurrenceDateTime
	OccurrenceDateTime_ext *Element `json:"_occurrenceDateTime"`
	// If the list price or the rule based factor associated with the code is overridden, this attribute can capture a text to indicate the  reason for this action.
	OverrideReason string `json:"overrideReason"`
	// Account into which this ChargeItems belongs.
	Account []*Reference `json:"account"`
	// This is a ChargeItem resource
	ResourceType string `json:"resourceType,omitempty"`
}

func NewChargeItem() *ChargeItem { return &ChargeItem{ResourceType: "ChargeItem"} }

// StructureMap_Structure is A Map of relationships between 2 structures that can be used to transform data.
type StructureMap_Structure struct {
	_ *BackboneElement
	// Documentation that describes how the structure is used in the mapping.
	Documentation string `json:"documentation"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// The canonical URL that identifies the structure.
	Url string `json:"url"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// How the referenced structure is used in this mapping.
	Mode string `json:"mode"`
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// The name used for this type in the map.
	Alias string `json:"alias"`
	// Extensions for alias
	Alias_ext *Element `json:"_alias"`
}

// ClaimResponse_Item is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse_Item struct {
	_ *BackboneElement
	// A service line number.
	// pattern [1-9][0-9]*
	SequenceLinkId uint64 `json:"sequenceLinkId"`
	// Extensions for sequenceLinkId
	SequenceLinkId_ext *Element `json:"_sequenceLinkId"`
	// A list of note references to the notes provided below.
	NoteNumber []uint64 `json:"noteNumber"`
	// Extensions for noteNumber
	NoteNumber_ext []*Element `json:"_noteNumber"`
	// The adjudication results.
	Adjudication []*ClaimResponse_Adjudication `json:"adjudication"`
	// The second tier service adjudications for submitted services.
	Detail []*ClaimResponse_Detail `json:"detail"`
}

// Coverage_Grouping is Financial instrument which may be used to reimburse or pay for health care products and services.
type Coverage_Grouping struct {
	_ *BackboneElement
	// Extensions for subGroupDisplay
	SubGroupDisplay_ext *Element `json:"_subGroupDisplay"`
	// A short description for the subplan.
	SubPlanDisplay string `json:"subPlanDisplay"`
	// Identifies a style or collective of coverage issues by the underwriter, for example may be used to identify a class of coverage such as a level of deductables or co-payment.
	Class string `json:"class"`
	// Extensions for subClass
	SubClass_ext *Element `json:"_subClass"`
	// Extensions for subClassDisplay
	SubClassDisplay_ext *Element `json:"_subClassDisplay"`
	// Extensions for group
	Group_ext *Element `json:"_group"`
	// Identifies a style or collective of coverage issued by the underwriter, for example may be used to identify a subset of an employer group.
	SubGroup string `json:"subGroup"`
	// A short description for the subgroup.
	SubGroupDisplay string `json:"subGroupDisplay"`
	// Extensions for plan
	Plan_ext *Element `json:"_plan"`
	// Extensions for subPlanDisplay
	SubPlanDisplay_ext *Element `json:"_subPlanDisplay"`
	// Identifies a sub-style or sub-collective of coverage issues by the underwriter, for example may be used to identify a subclass of coverage such as a sub-level of deductables or co-payment.
	SubClass string `json:"subClass"`
	// Extensions for groupDisplay
	GroupDisplay_ext *Element `json:"_groupDisplay"`
	// Identifies a style or collective of coverage issued by the underwriter, for example may be used to identify a collection of benefits provided to employees. May be referred to as a Section or Division ID.
	Plan string `json:"plan"`
	// Identifies a sub-style or sub-collective of coverage issued by the underwriter, for example may be used to identify a subset of a collection of benefits provided to employees.
	SubPlan string `json:"subPlan"`
	// Extensions for classDisplay
	ClassDisplay_ext *Element `json:"_classDisplay"`
	// Identifies a style or collective of coverage issued by the underwriter, for example may be used to identify an employer group. May also be referred to as a Policy or Group ID.
	Group string `json:"group"`
	// Extensions for subGroup
	SubGroup_ext *Element `json:"_subGroup"`
	// A short description for the plan.
	PlanDisplay string `json:"planDisplay"`
	// Extensions for planDisplay
	PlanDisplay_ext *Element `json:"_planDisplay"`
	// Extensions for subPlan
	SubPlan_ext *Element `json:"_subPlan"`
	// Extensions for class
	Class_ext *Element `json:"_class"`
	// A short description for the class.
	ClassDisplay string `json:"classDisplay"`
	// A short description for the subclass.
	SubClassDisplay string `json:"subClassDisplay"`
	// A short description for the group.
	GroupDisplay string `json:"groupDisplay"`
}

// ImagingStudy_Series is Representation of the content produced in a DICOM imaging study. A study comprises a set of series, each of which includes a set of Service-Object Pair Instances (SOP Instances - images or other data) acquired or produced in a common context.  A series is of only one modality (e.g. X-ray, CT, MR, ultrasound), but a study may have multiple series of different modalities.
type ImagingStudy_Series struct {
	_ *BackboneElement
	// Formal identifier for this series.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	Uid string `json:"uid"`
	// Extensions for number
	Number_ext *Element `json:"_number"`
	// The laterality of the (possibly paired) anatomic structures examined. E.g., the left knee, both lungs, or unpaired abdomen. If present, shall be consistent with any laterality information indicated in ImagingStudy.series.bodySite.
	Laterality *Coding `json:"laterality"`
	// Extensions for started
	Started_ext *Element `json:"_started"`
	// The physician or operator (often the radiology technician)  who performed the series. The performer is recorded at the series level, since each series in a study may be performed by a different practitioner, at different times, and using different devices. A series may be performed by multiple practitioners.
	Performer []*Reference `json:"performer"`
	// Extensions for uid
	Uid_ext *Element `json:"_uid"`
	// The numeric identifier of this series in the study.
	// pattern [0]|([1-9][0-9]*)
	Number uint64 `json:"number"`
	// The modality of this series sequence.
	Modality *Coding `json:"modality,omitempty"`
	// A description of the series.
	Description string `json:"description"`
	// Availability of series (online, offline or nearline).
	Availability string `json:"availability"`
	// The network service providing access (e.g., query, view, or retrieval) for this series. See implementation notes for information about using DICOM endpoints. A series-level endpoint, if present, has precedence over a study-level endpoint with the same Endpoint.type.
	Endpoint []*Reference `json:"endpoint"`
	// The anatomic structures examined. See DICOM Part 16 Annex L (http://dicom.nema.org/medical/dicom/current/output/chtml/part16/chapter_L.html) for DICOM to SNOMED-CT mappings. The bodySite may indicate the laterality of body part imaged; if so, it shall be consistent with any content of ImagingStudy.series.laterality.
	BodySite *Coding `json:"bodySite"`
	// A single SOP instance within the series, e.g. an image, or presentation state.
	Instance []*ImagingStudy_Instance `json:"instance"`
	// Extensions for numberOfInstances
	NumberOfInstances_ext *Element `json:"_numberOfInstances"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Number of SOP Instances in the Study. The value given may be larger than the number of instance elements this resource contains due to resource availability, security, or other factors. This element should be present if any instance elements are present.
	// pattern [0]|([1-9][0-9]*)
	NumberOfInstances uint64 `json:"numberOfInstances"`
	// Extensions for availability
	Availability_ext *Element `json:"_availability"`
	// The date and time the series was started.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Started time.Time `json:"started"`
}

// Bundle_Response is A container for a collection of resources.
type Bundle_Response struct {
	_ *BackboneElement
	// The status code returned by processing this entry. The status SHALL start with a 3 digit HTTP code (e.g. 404) and may contain the standard HTTP description associated with the status code.
	Status string `json:"status"`
	// Extensions for location
	Location_ext *Element `json:"_location"`
	// The date/time that the resource was modified on the server.
	LastModified string `json:"lastModified"`
	// Extensions for lastModified
	LastModified_ext *Element `json:"_lastModified"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The location header created by processing this operation.
	Location string `json:"location"`
	// The etag for the resource, it the operation for the entry produced a versioned resource (see [Resource Metadata and Versioning](http.html#versioning) and [Managing Resource Contention](http.html#concurrency)).
	Etag string `json:"etag"`
	// Extensions for etag
	Etag_ext *Element `json:"_etag"`
	// An OperationOutcome containing hints and warnings produced as part of processing this entry in a batch or transaction.
	Outcome interface{} `json:"outcome"`
}

// ImplementationGuide_Page is A set of rules of how FHIR is used to solve a particular problem. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuide_Page struct {
	_ *BackboneElement
	// The source address for the page.
	Source string `json:"source"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// The kind of page that this is. Some pages are autogenerated (list, example), and other kinds are of interest so that tools can navigate the user to the page of interest.
	Kind string `json:"kind"`
	// For constructed pages, a list of packages to include in the page (or else empty for everything).
	Package []string `json:"package"`
	// Extensions for package
	Package_ext []*Element `json:"_package"`
	// The format of the page.
	// pattern [^\s]+([\s]?[^\s]+)*
	Format string `json:"format"`
	// Extensions for format
	Format_ext *Element `json:"_format"`
	// Nested Pages/Sections under this page.
	Page []*ImplementationGuide_Page `json:"page"`
	// Extensions for source
	Source_ext *Element `json:"_source"`
	// A short title used to represent this page in navigational structures such as table of contents, bread crumbs, etc.
	Title string `json:"title"`
	// Extensions for kind
	Kind_ext *Element `json:"_kind"`
	// For constructed pages, what kind of resources to include in the list.
	Type []string `json:"type"`
	// Extensions for type
	Type_ext []*Element `json:"_type"`
}

// SupplyDelivery is Record of delivery of what is supplied.
type SupplyDelivery struct {
	_ *DomainResource
	// A link to a resource representing the person whom the delivered item is for.
	Patient *Reference `json:"patient"`
	// The individual responsible for dispensing the medication, supplier or device.
	Supplier *Reference `json:"supplier"`
	// Identification of the facility/location where the Supply was shipped to, as part of the dispense event.
	Destination *Reference `json:"destination"`
	// A plan, proposal or order that is fulfilled in whole or in part by this event.
	BasedOn []*Reference `json:"basedOn"`
	// A larger event of which this particular event is a component or step.
	PartOf []*Reference `json:"partOf"`
	// A code specifying the state of the dispense event.
	Status string `json:"status"`
	// This is a SupplyDelivery resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The date or time(s) the activity occurred.
	OccurrencePeriod *Period `json:"occurrencePeriod"`
	// Indicates the type of dispensing event that is performed. Examples include: Trial Fill, Completion of Trial, Partial Fill, Emergency Fill, Samples, etc.
	Type *CodeableConcept `json:"type"`
	// Extensions for occurrenceDateTime
	OccurrenceDateTime_ext *Element `json:"_occurrenceDateTime"`
	// The date or time(s) the activity occurred.
	OccurrenceTiming *Timing `json:"occurrenceTiming"`
	// Identifies the person who picked up the Supply.
	Receiver []*Reference `json:"receiver"`
	// Identifier assigned by the dispensing facility when the item(s) is dispensed.
	Identifier *Identifier `json:"identifier"`
	// The item that is being delivered or has been supplied.
	SuppliedItem *SupplyDelivery_SuppliedItem `json:"suppliedItem"`
	// The date or time(s) the activity occurred.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OccurrenceDateTime time.Time `json:"occurrenceDateTime"`
}

func NewSupplyDelivery() *SupplyDelivery { return &SupplyDelivery{ResourceType: "SupplyDelivery"} }

// AuditEvent is A record of an event made for purposes of maintaining a security log. Typical uses include detection of intrusion attempts and monitoring for inappropriate usage.
type AuditEvent struct {
	_ *DomainResource
	// Identifier for the category of event.
	Subtype []*Coding `json:"subtype"`
	// Extensions for outcomeDesc
	OutcomeDesc_ext *Element `json:"_outcomeDesc"`
	// Specific instances of data or objects that have been accessed.
	Entity []*AuditEvent_Entity `json:"entity"`
	// Indicator for type of action performed during the event that generated the audit.
	Action string `json:"action"`
	// A free text description of the outcome of the event.
	OutcomeDesc string `json:"outcomeDesc"`
	// An actor taking an active role in the event or activity that is logged.
	Agent []*AuditEvent_Agent `json:"agent,omitempty"`
	// This is a AuditEvent resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for action
	Action_ext *Element `json:"_action"`
	// The time when the event occurred on the source.
	Recorded string `json:"recorded"`
	// Extensions for outcome
	Outcome_ext *Element `json:"_outcome"`
	// Identifier for a family of the event.  For example, a menu item, program, rule, policy, function code, application name or URL. It identifies the performed function.
	Type *Coding `json:"type,omitempty"`
	// Extensions for recorded
	Recorded_ext *Element `json:"_recorded"`
	// Indicates whether the event succeeded or failed.
	Outcome string `json:"outcome"`
	// The purposeOfUse (reason) that was used during the event being recorded.
	PurposeOfEvent []*CodeableConcept `json:"purposeOfEvent"`
	// The system that is reporting the event.
	Source *AuditEvent_Source `json:"source,omitempty"`
}

func NewAuditEvent() *AuditEvent { return &AuditEvent{ResourceType: "AuditEvent"} }

// AuditEvent_Detail is A record of an event made for purposes of maintaining a security log. Typical uses include detection of intrusion attempts and monitoring for inappropriate usage.
type AuditEvent_Detail struct {
	_ *BackboneElement
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// The type of extra detail provided in the value.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The details, base64 encoded. Used to carry bulk information.
	Value string `json:"value"`
}

// DeviceMetric_Calibration is Describes a measurement, calculation or setting capability of a medical device.
type DeviceMetric_Calibration struct {
	_ *BackboneElement
	// Describes the type of the calibration method.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Describes the state of the calibration.
	State string `json:"state"`
	// Extensions for state
	State_ext *Element `json:"_state"`
	// Describes the time last calibration has been performed.
	Time string `json:"time"`
	// Extensions for time
	Time_ext *Element `json:"_time"`
}

// ExplanationOfBenefit_Procedure is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Procedure struct {
	_ *BackboneElement
	// The procedure code.
	ProcedureReference *Reference `json:"procedureReference"`
	// Sequence of procedures which serves to order and provide a link.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// Date and optionally time the procedure was performed .
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// The procedure code.
	ProcedureCodeableConcept *CodeableConcept `json:"procedureCodeableConcept"`
}

// Goal is Describes the intended objective(s) for a patient, group or organization care, for example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc.
type Goal struct {
	_ *DomainResource
	// Details of what's changed (or not changed).
	OutcomeReference []*Reference `json:"outcomeReference"`
	// This records identifiers associated with this care plan that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []*Identifier `json:"identifier"`
	// Identifies the patient, group or organization for whom the goal is being established.
	Subject *Reference `json:"subject"`
	// The date or event after which the goal should begin being pursued.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	StartDate time.Time `json:"startDate"`
	// Extensions for statusDate
	StatusDate_ext *Element `json:"_statusDate"`
	// Identifies the change (or lack of change) at the point when the status of the goal is assessed.
	OutcomeCode []*CodeableConcept `json:"outcomeCode"`
	// Indicates a category the goal falls within.
	Category []*CodeableConcept `json:"category"`
	// Extensions for startDate
	StartDate_ext *Element `json:"_startDate"`
	// The date or event after which the goal should begin being pursued.
	StartCodeableConcept *CodeableConcept `json:"startCodeableConcept"`
	// Indicates what should be done by when.
	Target *Goal_Target `json:"target"`
	// The identified conditions and other health record elements that are intended to be addressed by the goal.
	Addresses []*Reference `json:"addresses"`
	// Indicates whether the goal has been reached and is still considered relevant.
	Status string `json:"status"`
	// Identifies the mutually agreed level of importance associated with reaching/sustaining the goal.
	Priority *CodeableConcept `json:"priority"`
	// Captures the reason for the current status.
	StatusReason string `json:"statusReason"`
	// Extensions for statusReason
	StatusReason_ext *Element `json:"_statusReason"`
	// Any comments related to the goal.
	Note []*Annotation `json:"note"`
	// This is a Goal resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Human-readable and/or coded description of a specific desired objective of care, such as "control blood pressure" or "negotiate an obstacle course" or "dance with child at wedding".
	Description *CodeableConcept `json:"description,omitempty"`
	// Identifies when the current status.  I.e. When initially created, when achieved, when cancelled, etc.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	StatusDate time.Time `json:"statusDate"`
	// Indicates whose goal this is - patient goal, practitioner goal, etc.
	ExpressedBy *Reference `json:"expressedBy"`
}

func NewGoal() *Goal { return &Goal{ResourceType: "Goal"} }

// Location is Details and position information for a physical place where services are provided  and resources and participants may be stored, found, contained or accommodated.
type Location struct {
	_ *DomainResource
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Physical form of the location, e.g. building, room, vehicle, road.
	PhysicalType *CodeableConcept `json:"physicalType"`
	// Another Location which this Location is physically part of.
	PartOf *Reference `json:"partOf"`
	// Technical endpoints providing access to services operated for the location.
	Endpoint []*Reference `json:"endpoint"`
	// Name of the location as used by humans. Does not need to be unique.
	Name string `json:"name"`
	// The status property covers the general availability of the resource, not the current value which may be covered by the operationStatus, or by a schedule/slots if they are configured for the location.
	Status string `json:"status"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Description of the Location, which helps in finding or referencing the place.
	Description string `json:"description"`
	// Indicates whether a resource instance represents a specific location or a class of locations.
	Mode string `json:"mode"`
	// The contact details of communication devices available at the location. This can include phone numbers, fax numbers, mobile numbers, email addresses and web sites.
	Telecom []*ContactPoint `json:"telecom"`
	// The absolute geographic location of the Location, expressed using the WGS84 datum (This is the same co-ordinate system used in KML).
	Position *Location_Position `json:"position"`
	// The organization responsible for the provisioning and upkeep of the location.
	ManagingOrganization *Reference `json:"managingOrganization"`
	// This is a Location resource
	ResourceType string `json:"resourceType,omitempty"`
	// The Operational status covers operation values most relevant to beds (but can also apply to rooms/units/chair/etc such as an isolation unit/dialisys chair). This typically covers concepts such as contamination, housekeeping and other activities like maintenance.
	OperationalStatus *Coding `json:"operationalStatus"`
	// A list of alternate names that the location is known as, or was known as in the past.
	Alias []string `json:"alias"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for alias
	Alias_ext []*Element `json:"_alias"`
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// Indicates the type of function performed at the location.
	Type *CodeableConcept `json:"type"`
	// Physical location.
	Address *Address `json:"address"`
	// Unique code or number identifying the location to its users.
	Identifier []*Identifier `json:"identifier"`
}

func NewLocation() *Location { return &Location{ResourceType: "Location"} }

// TestScript_Action is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Action struct {
	_ *BackboneElement
	// The operation to perform.
	Operation *TestScript_Operation `json:"operation"`
	// Evaluates the results of previous operations to determine if the server under test behaves appropriately.
	Assert *TestScript_Assert `json:"assert"`
}

// Endpoint is The technical details of an endpoint that can be used for electronic services, such as for web services providing XDS.b or a REST endpoint for another FHIR server. This may include any security context information.
type Endpoint struct {
	_ *DomainResource
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The mime type to send the payload in - e.g. application/fhir+xml, application/fhir+json. If the mime type is not specified, then the sender could send any content (including no content depending on the connectionType).
	PayloadMimeType []string `json:"payloadMimeType"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The organization that manages this endpoint (even if technically another organisation is hosting this in the cloud, it is the organisation associated with the data).
	ManagingOrganization *Reference `json:"managingOrganization"`
	// The payload type describes the acceptable content that can be communicated on the endpoint.
	PayloadType []*CodeableConcept `json:"payloadType,omitempty"`
	// Extensions for payloadMimeType
	PayloadMimeType_ext []*Element `json:"_payloadMimeType"`
	// The uri that describes the actual end-point to connect to.
	Address string `json:"address"`
	// Additional headers / information to send as part of the notification.
	Header []string `json:"header"`
	// A friendly name that this endpoint can be referred to with.
	Name string `json:"name"`
	// Identifier for the organization that is used to identify the endpoint across multiple disparate systems.
	Identifier []*Identifier `json:"identifier"`
	// active | suspended | error | off | test.
	Status string `json:"status"`
	// A coded value that represents the technical details of the usage of this endpoint, such as what WSDLs should be used in what way. (e.g. XDS.b/DICOM/cds-hook).
	ConnectionType *Coding `json:"connectionType,omitempty"`
	// Contact details for a human to contact about the subscription. The primary use of this for system administrator troubleshooting.
	Contact []*ContactPoint `json:"contact"`
	// The interval during which the endpoint is expected to be operational.
	Period *Period `json:"period"`
	// Extensions for address
	Address_ext *Element `json:"_address"`
	// Extensions for header
	Header_ext []*Element `json:"_header"`
	// This is a Endpoint resource
	ResourceType string `json:"resourceType,omitempty"`
}

func NewEndpoint() *Endpoint { return &Endpoint{ResourceType: "Endpoint"} }

// OperationDefinition_Overload is A formal computable definition of an operation (on the RESTful interface) or a named query (using the search interaction).
type OperationDefinition_Overload struct {
	_ *BackboneElement
	// Name of parameter to include in overload.
	ParameterName []string `json:"parameterName"`
	// Extensions for parameterName
	ParameterName_ext []*Element `json:"_parameterName"`
	// Comments to go on overload.
	Comment string `json:"comment"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
}

// Contributor is A contributor to the content of a knowledge asset, including authors, editors, reviewers, and endorsers.
type Contributor struct {
	_ *Element
	// The type of contributor.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The name of the individual or organization responsible for the contribution.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Contact details to assist a user in finding and communicating with the contributor.
	Contact []*ContactDetail `json:"contact"`
}

// CapabilityStatement_SearchParam is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_SearchParam struct {
	_ *BackboneElement
	// Extensions for definition
	Definition_ext *Element `json:"_definition"`
	// The type of value a search parameter refers to, and how the content is interpreted.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// This allows documentation of any distinct behaviors about how the search parameter is used.  For example, text matching algorithms.
	Documentation string `json:"documentation"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// The name of the search parameter used in the interface.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// An absolute URI that is a formal reference to where this parameter was first defined, so that a client can be confident of the meaning of the search parameter (a reference to [[[SearchParameter.url]]]).
	Definition string `json:"definition"`
}

// Claim_Insurance is A provider issued list of services and products provided, or to be provided, to a patient which is provided to an insurer for payment recovery.
type Claim_Insurance struct {
	_ *BackboneElement
	// A list of references from the Insurer to which these services pertain.
	PreAuthRef []string `json:"preAuthRef"`
	// Extensions for preAuthRef
	PreAuthRef_ext []*Element `json:"_preAuthRef"`
	// The Coverages adjudication details.
	ClaimResponse *Reference `json:"claimResponse"`
	// A flag to indicate that this Coverage is the focus for adjudication. The Coverage against which the claim is to be adjudicated.
	Focal bool `json:"focal"`
	// Extensions for focal
	Focal_ext *Element `json:"_focal"`
	// Reference to the program or plan identification, underwriter or payor.
	Coverage *Reference `json:"coverage,omitempty"`
	// Extensions for businessArrangement
	BusinessArrangement_ext *Element `json:"_businessArrangement"`
	// Sequence of coverage which serves to provide a link and convey coordination of benefit order.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// The contract number of a business agreement which describes the terms and conditions.
	BusinessArrangement string `json:"businessArrangement"`
}

// DeviceRequest_Requester is Represents a request for a patient to employ a medical device. The device may be an implantable device, or an external assistive device, such as a walker.
type DeviceRequest_Requester struct {
	_ *BackboneElement
	// The organization the device or practitioner was acting on behalf of.
	OnBehalfOf *Reference `json:"onBehalfOf"`
	// The device, practitioner, etc. who initiated the request.
	Agent *Reference `json:"agent,omitempty"`
}

// EpisodeOfCare is An association between a patient and an organization / healthcare provider(s) during which time encounters may occur. The managing organization assumes a level of responsibility for the patient during this time.
type EpisodeOfCare struct {
	_ *DomainResource
	// planned | waitlist | active | onhold | finished | cancelled.
	Status string `json:"status"`
	// The organization that has assumed the specific responsibilities for the specified duration.
	ManagingOrganization *Reference `json:"managingOrganization"`
	// The list of practitioners that may be facilitating this episode of care for specific purposes.
	Team []*Reference `json:"team"`
	// The EpisodeOfCare may be known by different identifiers for different contexts of use, such as when an external agency is tracking the Episode for funding purposes.
	Identifier []*Identifier `json:"identifier"`
	// A classification of the type of episode of care; e.g. specialist referral, disease management, type of funded care.
	Type []*CodeableConcept `json:"type"`
	// The list of diagnosis relevant to this episode of care.
	Diagnosis []*EpisodeOfCare_Diagnosis `json:"diagnosis"`
	// The set of accounts that may be used for billing for this EpisodeOfCare.
	Account []*Reference `json:"account"`
	// The history of statuses that the EpisodeOfCare has been through (without requiring processing the history of the resource).
	StatusHistory []*EpisodeOfCare_StatusHistory `json:"statusHistory"`
	// The patient who is the focus of this episode of care.
	Patient *Reference `json:"patient,omitempty"`
	// The practitioner that is the care manager/care co-ordinator for this patient.
	CareManager *Reference `json:"careManager"`
	// Referral Request(s) that are fulfilled by this EpisodeOfCare, incoming referrals.
	ReferralRequest []*Reference `json:"referralRequest"`
	// This is a EpisodeOfCare resource
	ResourceType string `json:"resourceType,omitempty"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The interval during which the managing organization assumes the defined responsibility.
	Period *Period `json:"period"`
}

func NewEpisodeOfCare() *EpisodeOfCare { return &EpisodeOfCare{ResourceType: "EpisodeOfCare"} }

// ExpansionProfile_ExcludedSystem is Resource to define constraints on the Expansion of a FHIR ValueSet.
type ExpansionProfile_ExcludedSystem struct {
	_ *BackboneElement
	// The version of the code system from which codes in the expansion should be excluded.
	Version string `json:"version"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// An absolute URI which is the code system to be excluded.
	System string `json:"system"`
	// Extensions for system
	System_ext *Element `json:"_system"`
}

// Measure_Population is The Measure resource provides the definition of a quality measure.
type Measure_Population struct {
	_ *BackboneElement
	// Extensions for criteria
	Criteria_ext *Element `json:"_criteria"`
	// A unique identifier for the population criteria. This identifier is used to report data against this criteria within the measure report.
	Identifier *Identifier `json:"identifier"`
	// The type of population criteria.
	Code *CodeableConcept `json:"code"`
	// Optional name or short description of this population.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The human readable description of this population criteria.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The name of a valid referenced CQL expression (may be namespaced) that defines this population criteria.
	Criteria string `json:"criteria"`
}

// CapabilityStatement_Implementation is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Implementation struct {
	_ *BackboneElement
	// Information about the specific installation that this capability statement relates to.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// An absolute base URL for the implementation.  This forms the base for REST interfaces as well as the mailbox and document interfaces.
	Url string `json:"url"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
}

// DeviceRequest is Represents a request for a patient to employ a medical device. The device may be an implantable device, or an external assistive device, such as a walker.
type DeviceRequest struct {
	_ *DomainResource
	// The timing schedule for the use of the device. The Schedule data type allows many different expressions, for example. "Every 8 hours"; "Three times a day"; "1/2 an hour before breakfast for 10 days from 23-Dec 2011:"; "15 Oct 2013, 17 Oct 2013 and 1 Nov 2013".
	OccurrencePeriod *Period `json:"occurrencePeriod"`
	// The desired perfomer for doing the diagnostic testing.
	Performer *Reference `json:"performer"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Whether the request is a proposal, plan, an original order or a reflex order.
	Intent *CodeableConcept `json:"intent,omitempty"`
	// Extensions for authoredOn
	AuthoredOn_ext *Element `json:"_authoredOn"`
	// Desired type of performer for doing the diagnostic testing.
	PerformerType *CodeableConcept `json:"performerType"`
	// This is a DeviceRequest resource
	ResourceType string `json:"resourceType,omitempty"`
	// The request takes the place of the referenced completed or terminated request(s).
	PriorRequest []*Reference `json:"priorRequest"`
	// An encounter that provides additional context in which this request is made.
	Context *Reference `json:"context"`
	// Reason or justification for the use of this device.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Protocol or definition followed by this request. For example: The proposed act must be performed if the indicated conditions occur, e.g.., shortness of breath, SpO2 less than x%.
	Definition []*Reference `json:"definition"`
	// The details of the device to be used.
	CodeCodeableConcept *CodeableConcept `json:"codeCodeableConcept"`
	// The timing schedule for the use of the device. The Schedule data type allows many different expressions, for example. "Every 8 hours"; "Three times a day"; "1/2 an hour before breakfast for 10 days from 23-Dec 2011:"; "15 Oct 2013, 17 Oct 2013 and 1 Nov 2013".
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OccurrenceDateTime time.Time `json:"occurrenceDateTime"`
	// Plan/proposal/order fulfilled by this request.
	BasedOn []*Reference `json:"basedOn"`
	// Indicates how quickly the {{title}} should be addressed with respect to other requests.
	// pattern [^\s]+([\s]?[^\s]+)*
	Priority string `json:"priority"`
	// Extensions for priority
	Priority_ext *Element `json:"_priority"`
	// The individual who initiated the request and has responsibility for its activation.
	Requester *DeviceRequest_Requester `json:"requester"`
	// Identifiers assigned to this order by the orderer or by the receiver.
	Identifier []*Identifier `json:"identifier"`
	// When the request transitioned to being actionable.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AuthoredOn time.Time `json:"authoredOn"`
	// The patient who will use the device.
	Subject *Reference `json:"subject,omitempty"`
	// The details of the device to be used.
	CodeReference *Reference `json:"codeReference"`
	// Reason or justification for the use of this device.
	ReasonReference []*Reference `json:"reasonReference"`
	// Additional clinical information about the patient that may influence the request fulfilment.  For example, this may includes body where on the subject's the device will be used ( i.e. the target site).
	SupportingInfo []*Reference `json:"supportingInfo"`
	// Details about this request that were not represented at all or sufficiently in one of the attributes provided in a class. These may include for example a comment, an instruction, or a note associated with the statement.
	Note []*Annotation `json:"note"`
	// The status of the request.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Extensions for occurrenceDateTime
	OccurrenceDateTime_ext *Element `json:"_occurrenceDateTime"`
	// The timing schedule for the use of the device. The Schedule data type allows many different expressions, for example. "Every 8 hours"; "Three times a day"; "1/2 an hour before breakfast for 10 days from 23-Dec 2011:"; "15 Oct 2013, 17 Oct 2013 and 1 Nov 2013".
	OccurrenceTiming *Timing `json:"occurrenceTiming"`
	// Key events in the history of the request.
	RelevantHistory []*Reference `json:"relevantHistory"`
	// Composite request this is part of.
	GroupIdentifier *Identifier `json:"groupIdentifier"`
}

func NewDeviceRequest() *DeviceRequest { return &DeviceRequest{ResourceType: "DeviceRequest"} }

// ExplanationOfBenefit_Information is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Information struct {
	_ *BackboneElement
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	ValueAttachment *Attachment `json:"valueAttachment"`
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	ValueReference *Reference `json:"valueReference"`
	// Sequence of the information element which serves to provide a link.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// The general class of the information supplied: information; exception; accident, employment; onset, etc.
	Category *CodeableConcept `json:"category,omitempty"`
	// The date when or period to which this information refers.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	TimingDate time.Time `json:"timingDate"`
	// Extensions for timingDate
	TimingDate_ext *Element `json:"_timingDate"`
	// The date when or period to which this information refers.
	TimingPeriod *Period `json:"timingPeriod"`
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	ValueString string `json:"valueString"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// For example, provides the reason for: the additional stay, or missing tooth or any other situation where a reason code is required in addition to the content.
	Reason *Coding `json:"reason"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// System and code pertaining to the specific information regarding special conditions relating to the setting, treatment or patient  for which care is sought which may influence the adjudication.
	Code *CodeableConcept `json:"code"`
}

// Immunization_Explanation is Describes the event of a patient being administered a vaccination or a record of a vaccination as reported by a patient, a clinician or another party and may include vaccine reaction information and what vaccination protocol was followed.
type Immunization_Explanation struct {
	_ *BackboneElement
	// Reasons why a vaccine was administered.
	Reason []*CodeableConcept `json:"reason"`
	// Reason why a vaccine was not administered.
	ReasonNotGiven []*CodeableConcept `json:"reasonNotGiven"`
}

// ElementDefinition_Constraint is Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition_Constraint struct {
	_ *BackboneElement
	// Identifies the impact constraint violation has on the conformance of the instance.
	Severity string `json:"severity"`
	// Extensions for severity
	Severity_ext *Element `json:"_severity"`
	// A [FHIRPath](http://hl7.org/fluentpath) expression of constraint that can be executed to see if this constraint is met.
	Expression string `json:"expression"`
	// An XPath expression of constraint that can be executed to see if this constraint is met.
	Xpath string `json:"xpath"`
	// Extensions for expression
	Expression_ext *Element `json:"_expression"`
	// A reference to the original source of the constraint, for traceability purposes.
	Source string `json:"source"`
	// Extensions for source
	Source_ext *Element `json:"_source"`
	// Allows identification of which elements have their cardinalities impacted by the constraint.  Will not be referenced for constraints that do not affect cardinality.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Key string `json:"key"`
	// Extensions for key
	Key_ext *Element `json:"_key"`
	// Extensions for requirements
	Requirements_ext *Element `json:"_requirements"`
	// Extensions for human
	Human_ext *Element `json:"_human"`
	// Extensions for xpath
	Xpath_ext *Element `json:"_xpath"`
	// Description of why this constraint is necessary or appropriate.
	Requirements string `json:"requirements"`
	// Text that can be used to describe the constraint in messages identifying that the constraint has been violated.
	Human string `json:"human"`
}

// DataRequirement_DateFilter is Describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type DataRequirement_DateFilter struct {
	_ *BackboneElement
	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration from now.
	ValueDuration *Duration `json:"valueDuration"`
	// The date-valued attribute of the filter. The specified path must be resolvable from the type of the required data. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements. Note that the index must be an integer constant. The path must resolve to an element of type dateTime, Period, Schedule, or Timing.
	Path string `json:"path"`
	// Extensions for path
	Path_ext *Element `json:"_path"`
	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration from now.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ValueDateTime time.Time `json:"valueDateTime"`
	// Extensions for valueDateTime
	ValueDateTime_ext *Element `json:"_valueDateTime"`
	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration from now.
	ValuePeriod *Period `json:"valuePeriod"`
}

// Contract_Term is A formal agreement between parties regarding the conduct of business, exchange of information or other matters.
type Contract_Term struct {
	_ *BackboneElement
	// Unique identifier for this particular Contract Provision.
	Identifier *Identifier `json:"identifier"`
	// Extensions for issued
	Issued_ext *Element `json:"_issued"`
	// Relevant time or time-period when this Contract Provision is applicable.
	Applies *Period `json:"applies"`
	// Type of Contract Provision such as specific requirements, purposes for actions, obligations, prohibitions, e.g. life time maximum benefit.
	Type *CodeableConcept `json:"type"`
	// Reason or purpose for the action stipulated by this Contract Provision.
	ActionReason []*CodeableConcept `json:"actionReason"`
	// Nested group of Contract Provisions.
	Group []*Contract_Term `json:"group"`
	// When this Contract Provision was issued.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Issued time.Time `json:"issued"`
	// Action stipulated by this Contract Provision.
	Action []*CodeableConcept `json:"action"`
	// A set of security labels that define which terms are controlled by this condition.
	SecurityLabel []*Coding `json:"securityLabel"`
	// An actor taking a role in an activity for which it can be assigned some degree of responsibility for the activity taking place.
	Agent []*Contract_Agent1 `json:"agent"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// Subtype of this Contract Provision, e.g. life time maximum payment for a contract term for specific valued item, e.g. disability payment.
	SubType *CodeableConcept `json:"subType"`
	// The matter of concern in the context of this provision of the agrement.
	Topic []*Reference `json:"topic"`
	// Human readable form of this Contract Provision.
	Text string `json:"text"`
	// Contract Provision Valued Item List.
	ValuedItem []*Contract_ValuedItem1 `json:"valuedItem"`
}

// ImmunizationRecommendation_DateCriterion is A patient's point-in-time immunization and recommendation (i.e. forecasting a patient's immunization eligibility according to a published schedule) with optional supporting justification.
type ImmunizationRecommendation_DateCriterion struct {
	_ *BackboneElement
	// Date classification of recommendation.  For example, earliest date to give, latest date to give, etc.
	Code *CodeableConcept `json:"code,omitempty"`
	// The date whose meaning is specified by dateCriterion.code.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Value time.Time `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
}

// MessageHeader_Destination is The header for a message exchange that is either requesting or responding to an action.  The reference(s) that are the subject of the action as well as other information related to the action are typically transmitted in a bundle in which the MessageHeader resource instance is the first resource in the bundle.
type MessageHeader_Destination struct {
	_ *BackboneElement
	// Extensions for endpoint
	Endpoint_ext *Element `json:"_endpoint"`
	// Human-readable name for the target system.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Identifies the target end system in situations where the initial message transmission is to an intermediary system.
	Target *Reference `json:"target"`
	// Indicates where the message should be routed to.
	Endpoint string `json:"endpoint"`
}

// ProcessRequest_Item is This resource provides the target, request and response, and action details for an action to be performed by the target on or about existing resources.
type ProcessRequest_Item struct {
	_ *BackboneElement
	// A service line number.
	// pattern -?([0]|([1-9][0-9]*))
	SequenceLinkId int64 `json:"sequenceLinkId"`
	// Extensions for sequenceLinkId
	SequenceLinkId_ext *Element `json:"_sequenceLinkId"`
}

// StructureMap_Parameter is A Map of relationships between 2 structures that can be used to transform data.
type StructureMap_Parameter struct {
	_ *BackboneElement
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// Extensions for valueDecimal
	ValueDecimal_ext *Element `json:"_valueDecimal"`
	// Extensions for valueId
	ValueId_ext *Element `json:"_valueId"`
	// Parameter value - variable or literal.
	ValueBoolean bool `json:"valueBoolean"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// Parameter value - variable or literal.
	// pattern -?([0]|([1-9][0-9]*))
	ValueInteger int64 `json:"valueInteger"`
	// Extensions for valueInteger
	ValueInteger_ext *Element `json:"_valueInteger"`
	// Parameter value - variable or literal.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	ValueDecimal float64 `json:"valueDecimal"`
	// Parameter value - variable or literal.
	// pattern [A-Za-z0-9\-\.]{1,64}
	ValueId string `json:"valueId"`
	// Parameter value - variable or literal.
	ValueString string `json:"valueString"`
}

// Duration is A length of time.
type Duration struct {
	_ *Quantity
}

// ElementDefinition_Mapping is Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition_Mapping struct {
	_ *BackboneElement
	// Identifies the computable language in which mapping.map is expressed.
	// pattern [^\s]+([\s]?[^\s]+)*
	Language string `json:"language"`
	// Extensions for language
	Language_ext *Element `json:"_language"`
	// Expresses what part of the target specification corresponds to this element.
	Map string `json:"map"`
	// Extensions for map
	Map_ext *Element `json:"_map"`
	// Comments that provide information about the mapping or its use.
	Comment string `json:"comment"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// An internal reference to the definition of a mapping.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Identity string `json:"identity"`
	// Extensions for identity
	Identity_ext *Element `json:"_identity"`
}

// CommunicationRequest_Payload is A request to convey information; e.g. the CDS system proposes that an alert be sent to a responsible provider, the CDS system proposes that the public health agency be notified about a reportable condition.
type CommunicationRequest_Payload struct {
	_ *BackboneElement
	// Extensions for contentString
	ContentString_ext *Element `json:"_contentString"`
	// The communicated content (or for multi-part communications, one portion of the communication).
	ContentAttachment *Attachment `json:"contentAttachment"`
	// The communicated content (or for multi-part communications, one portion of the communication).
	ContentReference *Reference `json:"contentReference"`
	// The communicated content (or for multi-part communications, one portion of the communication).
	ContentString string `json:"contentString"`
}

// PlanDefinition_Target is This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinition_Target struct {
	_ *BackboneElement
	// The parameter whose value is to be tracked, e.g. body weigth, blood pressure, or hemoglobin A1c level.
	Measure *CodeableConcept `json:"measure"`
	// The target value of the measure to be achieved to signify fulfillment of the goal, e.g. 150 pounds or 7.0%. Either the high or low or both values of the range can be specified. Whan a low value is missing, it indicates that the goal is achieved at any value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any value at or above the low value.
	DetailQuantity *Quantity `json:"detailQuantity"`
	// The target value of the measure to be achieved to signify fulfillment of the goal, e.g. 150 pounds or 7.0%. Either the high or low or both values of the range can be specified. Whan a low value is missing, it indicates that the goal is achieved at any value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any value at or above the low value.
	DetailRange *Range `json:"detailRange"`
	// The target value of the measure to be achieved to signify fulfillment of the goal, e.g. 150 pounds or 7.0%. Either the high or low or both values of the range can be specified. Whan a low value is missing, it indicates that the goal is achieved at any value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any value at or above the low value.
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept"`
	// Indicates the timeframe after the start of the goal in which the goal should be met.
	Due *Duration `json:"due"`
}

// Questionnaire is A structured set of questions intended to guide the collection of answers from end-users. Questionnaires provide detailed control over order, presentation, phraseology and grouping to allow coherent, consistent data collection.
type Questionnaire struct {
	_ *DomainResource
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// A formal identifier that is used to identify this questionnaire when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ApprovalDate time.Time `json:"approvalDate"`
	// A particular question, question grouping or display text that is part of the questionnaire.
	Item []*Questionnaire_Item `json:"item"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The status of this questionnaire. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// The date  (and optionally time) when the questionnaire was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the questionnaire changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// The period during which the questionnaire content was or is planned to be in active use.
	EffectivePeriod *Period `json:"effectivePeriod"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// A natural language name identifying the questionnaire. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// The identifier that is used to identify this version of the questionnaire when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the questionnaire author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// Extensions for approvalDate
	ApprovalDate_ext *Element `json:"_approvalDate"`
	// The date on which the resource content was last reviewed. Review happens periodically after approval, but doesn't change the original approval date.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	LastReviewDate time.Time `json:"lastReviewDate"`
	// A legal or geographic region in which the questionnaire is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// An identifier for this question or group of questions in a particular terminology such as LOINC.
	Code []*Coding `json:"code"`
	// The types of subjects that can be the subject of responses created for the questionnaire.
	SubjectType []string `json:"subjectType"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// A free text natural language description of the questionnaire from a consumer's perspective.
	Description string `json:"description"`
	// Explaination of why this questionnaire is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// Extensions for lastReviewDate
	LastReviewDate_ext *Element `json:"_lastReviewDate"`
	// An absolute URI that is used to identify this questionnaire when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this questionnaire is (or will be) published. The URL SHOULD include the major version of the questionnaire. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// A boolean value to indicate that this questionnaire is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Extensions for subjectType
	SubjectType_ext []*Element `json:"_subjectType"`
	// This is a Questionnaire resource
	ResourceType string `json:"resourceType,omitempty"`
	// A short, descriptive, user-friendly title for the questionnaire.
	Title string `json:"title"`
	// The name of the individual or organization that published the questionnaire.
	Publisher string `json:"publisher"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate questionnaire instances.
	UseContext []*UsageContext `json:"useContext"`
	// A copyright statement relating to the questionnaire and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the questionnaire.
	Copyright string `json:"copyright"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
}

func NewQuestionnaire() *Questionnaire { return &Questionnaire{ResourceType: "Questionnaire"} }

// TestScript_Param3 is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Param3 struct {
	_ *BackboneElement
	// Descriptive name for this parameter that matches the external assert ruleset rule parameter name.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The value for the parameter that will be passed on to the external ruleset rule template.
	Value string `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
}

// Appointment is A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time. This may result in one or more Encounter(s).
type Appointment struct {
	_ *DomainResource
	// The reason that this appointment is being scheduled. This is more clinical than administrative.
	Reason []*CodeableConcept `json:"reason"`
	// The brief description of the appointment as would be shown on a subject line in a meeting request, or appointment list. Detailed or expanded information should be put in the comment field.
	Description string `json:"description"`
	// A set of date ranges (potentially including times) that the appointment is preferred to be scheduled within. When using these values, the minutes duration should be provided to indicate the length of the appointment to fill and populate the start/end times for the actual allocated time.
	RequestedPeriod []*Period `json:"requestedPeriod"`
	// Number of minutes that the appointment is to take. This can be less than the duration between the start and end times (where actual time of appointment is only an estimate or is a planned appointment request).
	// pattern [1-9][0-9]*
	MinutesDuration uint64 `json:"minutesDuration"`
	// The slots from the participants' schedules that will be filled by the appointment.
	Slot []*Reference `json:"slot"`
	// The referral request this appointment is allocated to assess (incoming referral).
	IncomingReferral []*Reference `json:"incomingReferral"`
	// This is a Appointment resource
	ResourceType string `json:"resourceType,omitempty"`
	// This records identifiers associated with this appointment concern that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []*Identifier `json:"identifier"`
	// A broad categorisation of the service that is to be performed during this appointment.
	ServiceCategory *CodeableConcept `json:"serviceCategory"`
	// The date that this appointment was initially created. This could be different to the meta.lastModified value on the initial entry, as this could have been before the resource was created on the FHIR server, and should remain unchanged over the lifespan of the appointment.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// Additional comments about the appointment.
	Comment string `json:"comment"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// The specific service that is to be performed during this appointment.
	ServiceType []*CodeableConcept `json:"serviceType"`
	// The specialty of a practitioner that would be required to perform the service requested in this appointment.
	Specialty []*CodeableConcept `json:"specialty"`
	// The priority of the appointment. Can be used to make informed decisions if needing to re-prioritize appointments. (The iCal Standard specifies 0 as undefined, 1 as highest, 9 as lowest priority).
	// pattern [0]|([1-9][0-9]*)
	Priority uint64 `json:"priority"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Date/Time that the appointment is to take place.
	Start string `json:"start"`
	// The overall status of the Appointment. Each of the participants has their own participation status which indicates their involvement in the process, however this status indicates the shared status.
	Status string `json:"status"`
	// The style of appointment or patient that has been booked in the slot (not service type).
	AppointmentType *CodeableConcept `json:"appointmentType"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Reason the appointment has been scheduled to take place, as specified using information from another resource. When the patient arrives and the encounter begins it may be used as the admission diagnosis. The indication will typically be a Condition (with other resources referenced in the evidence.detail), or a Procedure.
	Indication []*Reference `json:"indication"`
	// Additional information to support the appointment provided when making the appointment.
	SupportingInformation []*Reference `json:"supportingInformation"`
	// Extensions for start
	Start_ext *Element `json:"_start"`
	// Date/Time that the appointment is to conclude.
	End string `json:"end"`
	// List of participants involved in the appointment.
	Participant []*Appointment_Participant `json:"participant,omitempty"`
	// Extensions for priority
	Priority_ext *Element `json:"_priority"`
	// Extensions for end
	End_ext *Element `json:"_end"`
	// Extensions for minutesDuration
	MinutesDuration_ext *Element `json:"_minutesDuration"`
}

func NewAppointment() *Appointment { return &Appointment{ResourceType: "Appointment"} }

// ExpansionProfile_Exclude is Resource to define constraints on the Expansion of a FHIR ValueSet.
type ExpansionProfile_Exclude struct {
	_ *BackboneElement
	// A data group for each designation to be excluded.
	Designation []*ExpansionProfile_Designation2 `json:"designation"`
}

// Goal_Target is Describes the intended objective(s) for a patient, group or organization care, for example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc.
type Goal_Target struct {
	_ *BackboneElement
	// The target value of the focus to be achieved to signify the fulfillment of the goal, e.g. 150 pounds, 7.0%. Either the high or low or both values of the range can be specified. When a low value is missing, it indicates that the goal is achieved at any focus value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any focus value at or above the low value.
	DetailQuantity *Quantity `json:"detailQuantity"`
	// The target value of the focus to be achieved to signify the fulfillment of the goal, e.g. 150 pounds, 7.0%. Either the high or low or both values of the range can be specified. When a low value is missing, it indicates that the goal is achieved at any focus value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any focus value at or above the low value.
	DetailRange *Range `json:"detailRange"`
	// The target value of the focus to be achieved to signify the fulfillment of the goal, e.g. 150 pounds, 7.0%. Either the high or low or both values of the range can be specified. When a low value is missing, it indicates that the goal is achieved at any focus value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any focus value at or above the low value.
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept"`
	// Indicates either the date or the duration after start by which the goal should be met.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	DueDate time.Time `json:"dueDate"`
	// Extensions for dueDate
	DueDate_ext *Element `json:"_dueDate"`
	// Indicates either the date or the duration after start by which the goal should be met.
	DueDuration *Duration `json:"dueDuration"`
	// The parameter whose value is being tracked, e.g. body weight, blood pressure, or hemoglobin A1c level.
	Measure *CodeableConcept `json:"measure"`
}

// TestScript_Rule3 is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Rule3 struct {
	_ *BackboneElement
	// Each rule template can take one or more parameters for rule evaluation.
	Param []*TestScript_Param3 `json:"param"`
	// Id of the referenced rule within the external ruleset template.
	// pattern [A-Za-z0-9\-\.]{1,64}
	RuleId string `json:"ruleId"`
	// Extensions for ruleId
	RuleId_ext *Element `json:"_ruleId"`
}

// List_Entry is A set of information summarized from a list of other resources.
type List_Entry struct {
	_ *BackboneElement
	// Extensions for deleted
	Deleted_ext *Element `json:"_deleted"`
	// When this item was added to the list.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// A reference to the actual resource from which data was derived.
	Item *Reference `json:"item,omitempty"`
	// The flag allows the system constructing the list to indicate the role and significance of the item in the list.
	Flag *CodeableConcept `json:"flag"`
	// True if this item is marked as deleted in the list.
	Deleted bool `json:"deleted"`
}

// ResearchSubject is A process where a researcher or organization plans and then executes a series of steps intended to increase the field of healthcare-related knowledge.  This includes studies of safety, efficacy, comparative effectiveness and other information about medications, devices, therapies and other interventional and investigative techniques.  A ResearchStudy involves the gathering of information about human or animal subjects.
type ResearchSubject struct {
	_ *DomainResource
	// Extensions for actualArm
	ActualArm_ext *Element `json:"_actualArm"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The record of the person or animal who is involved in the study.
	Individual *Reference `json:"individual,omitempty"`
	// The current state of the subject.
	Status string `json:"status"`
	// The dates the subject began and ended their participation in the study.
	Period *Period `json:"period"`
	// Reference to the study the subject is participating in.
	Study *Reference `json:"study,omitempty"`
	// The name of the arm in the study the subject is expected to follow as part of this study.
	AssignedArm string `json:"assignedArm"`
	// Extensions for assignedArm
	AssignedArm_ext *Element `json:"_assignedArm"`
	// The name of the arm in the study the subject actually followed as part of this study.
	ActualArm string `json:"actualArm"`
	// This is a ResearchSubject resource
	ResourceType string `json:"resourceType,omitempty"`
	// Identifiers assigned to this research study by the sponsor or other systems.
	Identifier *Identifier `json:"identifier"`
	// A record of the patient's informed agreement to participate in the study.
	Consent *Reference `json:"consent"`
}

func NewResearchSubject() *ResearchSubject { return &ResearchSubject{ResourceType: "ResearchSubject"} }

// CareTeam is The Care Team includes all the people and organizations who plan to participate in the coordination and delivery of care for a patient.
type CareTeam struct {
	_ *DomainResource
	// Comments made about the CareTeam.
	Note []*Annotation `json:"note"`
	// The encounter or episode of care that establishes the context for this care team.
	Context *Reference `json:"context"`
	// The organization responsible for the care team.
	ManagingOrganization []*Reference `json:"managingOrganization"`
	// A label for human use intended to distinguish like teams.  E.g. the "red" vs. "green" trauma teams.
	Name string `json:"name"`
	// Identifies the patient or group whose intended care is handled by the team.
	Subject *Reference `json:"subject"`
	// Describes why the care team exists.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Condition(s) that this care team addresses.
	ReasonReference []*Reference `json:"reasonReference"`
	// Indicates the current state of the care team.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// This is a CareTeam resource
	ResourceType string `json:"resourceType,omitempty"`
	// Identifies what kind of team.  This is to support differentiation between multiple co-existing teams, such as care plan team, episode of care team, longitudinal care team.
	Category []*CodeableConcept `json:"category"`
	// Indicates when the team did (or is intended to) come into effect and end.
	Period *Period `json:"period"`
	// Identifies all people and organizations who are expected to be involved in the care team.
	Participant []*CareTeam_Participant `json:"participant"`
	// This records identifiers associated with this care team that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
}

func NewCareTeam() *CareTeam { return &CareTeam{ResourceType: "CareTeam"} }

// ImagingStudy_Instance is Representation of the content produced in a DICOM imaging study. A study comprises a set of series, each of which includes a set of Service-Object Pair Instances (SOP Instances - images or other data) acquired or produced in a common context.  A series is of only one modality (e.g. X-ray, CT, MR, ultrasound), but a study may have multiple series of different modalities.
type ImagingStudy_Instance struct {
	_ *BackboneElement
	// DICOM instance  type.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	SopClass string `json:"sopClass"`
	// Extensions for sopClass
	SopClass_ext *Element `json:"_sopClass"`
	// The description of the instance.
	Title string `json:"title"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Formal identifier for this image or other content.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	Uid string `json:"uid"`
	// Extensions for uid
	Uid_ext *Element `json:"_uid"`
	// The number of instance in the series.
	// pattern [0]|([1-9][0-9]*)
	Number uint64 `json:"number"`
	// Extensions for number
	Number_ext *Element `json:"_number"`
}

// RiskAssessment_Prediction is An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
type RiskAssessment_Prediction struct {
	_ *BackboneElement
	// Indicates the period of time or age range of the subject to which the specified probability applies.
	WhenPeriod *Period `json:"whenPeriod"`
	// How likely is the outcome (in the specified timeframe).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	ProbabilityDecimal float64 `json:"probabilityDecimal"`
	// Extensions for probabilityDecimal
	ProbabilityDecimal_ext *Element `json:"_probabilityDecimal"`
	// How likely is the outcome (in the specified timeframe), expressed as a qualitative value (e.g. low, medium, high).
	QualitativeRisk *CodeableConcept `json:"qualitativeRisk"`
	// Extensions for relativeRisk
	RelativeRisk_ext *Element `json:"_relativeRisk"`
	// Additional information explaining the basis for the prediction.
	Rationale string `json:"rationale"`
	// Extensions for rationale
	Rationale_ext *Element `json:"_rationale"`
	// One of the potential outcomes for the patient (e.g. remission, death,  a particular condition).
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	// How likely is the outcome (in the specified timeframe).
	ProbabilityRange *Range `json:"probabilityRange"`
	// Indicates the risk for this particular subject (with their specific characteristics) divided by the risk of the population in general.  (Numbers greater than 1 = higher risk than the population, numbers less than 1 = lower risk.).
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	RelativeRisk float64 `json:"relativeRisk"`
	// Indicates the period of time or age range of the subject to which the specified probability applies.
	WhenRange *Range `json:"whenRange"`
}

// ClaimResponse is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse struct {
	_ *DomainResource
	// The Response business identifier.
	Identifier []*Identifier `json:"identifier"`
	// Mutually exclusive with Services Provided (Item).
	Error []*ClaimResponse_Error `json:"error"`
	// The amount of deductible applied which was not allocated to any particular service line.
	UnallocDeductable *Money `json:"unallocDeductable"`
	// Status of funds reservation (For provider, for Patient, None).
	Reserved *Coding `json:"reserved"`
	// The first tier service adjudications for submitted services.
	Item []*ClaimResponse_Item `json:"item"`
	// Note text.
	ProcessNote []*ClaimResponse_ProcessNote `json:"processNote"`
	// Financial instrument by which payment information for health care.
	Insurance []*ClaimResponse_Insurance `json:"insurance"`
	// The practitioner who is responsible for the services rendered to the patient.
	RequestProvider *Reference `json:"requestProvider"`
	// Extensions for disposition
	Disposition_ext *Element `json:"_disposition"`
	// Payment details for the claim if the claim has been paid.
	Payment *ClaimResponse_Payment `json:"payment"`
	// The form to be used for printing the content.
	Form *CodeableConcept `json:"form"`
	// Processing outcome errror, partial or complete processing.
	Outcome *CodeableConcept `json:"outcome"`
	// Party to be reimbursed: Subscriber, provider, other.
	PayeeType *CodeableConcept `json:"payeeType"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// This is a ClaimResponse resource
	ResourceType string `json:"resourceType,omitempty"`
	// Total amount of benefit payable (Equal to sum of the Benefit amounts from all detail lines and additions less the Unallocated Deductible).
	TotalBenefit *Money `json:"totalBenefit"`
	// Patient Resource.
	Patient *Reference `json:"patient"`
	// The date when the enclosed suite of services were performed or completed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// Original request resource referrence.
	Request *Reference `json:"request"`
	// Request for additional supporting or authorizing information, such as: documents, images or resources.
	CommunicationRequest []*Reference `json:"communicationRequest"`
	// The Insurer who produced this adjudicated response.
	Insurer *Reference `json:"insurer"`
	// A description of the status of the adjudication.
	Disposition string `json:"disposition"`
	// The first tier service adjudications for payor added services.
	AddItem []*ClaimResponse_AddItem `json:"addItem"`
	// The total cost of the services reported.
	TotalCost *Money `json:"totalCost"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// The organization which is responsible for the services rendered to the patient.
	RequestOrganization *Reference `json:"requestOrganization"`
}

func NewClaimResponse() *ClaimResponse { return &ClaimResponse{ResourceType: "ClaimResponse"} }

// DataElement_Mapping is The formal description of a single piece of information that can be gathered and reported.
type DataElement_Mapping struct {
	_ *BackboneElement
	// An internal id that is used to identify this mapping set when specific mappings are made on a per-element basis.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Identity string `json:"identity"`
	// Extensions for identity
	Identity_ext *Element `json:"_identity"`
	// An absolute URI that identifies the specification that this mapping is expressed to.
	Uri string `json:"uri"`
	// Extensions for uri
	Uri_ext *Element `json:"_uri"`
	// A name for the specification that is being mapped to.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Comments about this mapping, including version notes, issues, scope limitations, and other important notes for usage.
	Comment string `json:"comment"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
}

// Encounter is An interaction between a patient and healthcare provider(s) for the purpose of providing healthcare service(s) or assessing the health status of a patient.
type Encounter struct {
	_ *DomainResource
	// List of locations where  the patient has been during this encounter.
	Location []*Encounter_Location `json:"location"`
	// The start and end time of the encounter.
	Period *Period `json:"period"`
	// The set of accounts that may be used for billing for this Encounter.
	Account []*Reference `json:"account"`
	// The class history permits the tracking of the encounters transitions without needing to go  through the resource history.
	//
	// This would be used for a case where an admission starts of as an emergency encounter, then transisions into an inpatient scenario. Doing this and not restarting a new encounter ensures that any lab/diagnostic results can more easily follow the patient and not require re-processing and not get lost or cancelled during a kindof discharge from emergency to inpatient.
	ClassHistory []*Encounter_ClassHistory `json:"classHistory"`
	// The referral request this encounter satisfies (incoming referral).
	IncomingReferral []*Reference `json:"incomingReferral"`
	// The list of people responsible for providing the service.
	Participant []*Encounter_Participant `json:"participant"`
	// Reason the encounter takes place, expressed as a code. For admissions, this can be used for a coded admission diagnosis.
	Reason []*CodeableConcept `json:"reason"`
	// This is a Encounter resource
	ResourceType string `json:"resourceType,omitempty"`
	// inpatient | outpatient | ambulatory | emergency +.
	Class *Coding `json:"class"`
	// Quantity of time the encounter lasted. This excludes the time during leaves of absence.
	Length *Duration `json:"length"`
	// The list of diagnosis relevant to this encounter.
	Diagnosis []*Encounter_Diagnosis `json:"diagnosis"`
	// Details about the admission to a healthcare service.
	Hospitalization *Encounter_Hospitalization `json:"hospitalization"`
	// Another Encounter of which this encounter is a part of (administratively or in time).
	PartOf *Reference `json:"partOf"`
	// Identifier(s) by which this encounter is known.
	Identifier []*Identifier `json:"identifier"`
	// Specific type of encounter (e.g. e-mail consultation, surgical day-care, skilled nursing, rehabilitation).
	Type []*CodeableConcept `json:"type"`
	// The status history permits the encounter resource to contain the status history without needing to read through the historical versions of the resource, or even have the server store them.
	StatusHistory []*Encounter_StatusHistory `json:"statusHistory"`
	// Indicates the urgency of the encounter.
	Priority *CodeableConcept `json:"priority"`
	// The patient ro group present at the encounter.
	Subject *Reference `json:"subject"`
	// Where a specific encounter should be classified as a part of a specific episode(s) of care this field should be used. This association can facilitate grouping of related encounters together for a specific purpose, such as government reporting, issue tracking, association via a common problem.  The association is recorded on the encounter as these are typically created after the episode of care, and grouped on entry rather than editing the episode of care to append another encounter to it (the episode of care could span years).
	EpisodeOfCare []*Reference `json:"episodeOfCare"`
	// The appointment that scheduled this encounter.
	Appointment *Reference `json:"appointment"`
	// An organization that is in charge of maintaining the information of this Encounter (e.g. who maintains the report or the master service catalog item, etc.). This MAY be the same as the organization on the Patient record, however it could be different. This MAY not be not the Service Delivery Location's Organization.
	ServiceProvider *Reference `json:"serviceProvider"`
	// planned | arrived | triaged | in-progress | onleave | finished | cancelled +.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
}

func NewEncounter() *Encounter { return &Encounter{ResourceType: "Encounter"} }

// Task_Input is A task to be performed.
type Task_Input struct {
	_ *BackboneElement
	// The value of the input parameter as a basic type.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ValueDateTime time.Time `json:"valueDateTime"`
	// The value of the input parameter as a basic type.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	ValueTime time.Time `json:"valueTime"`
	// Extensions for valueUnsignedInt
	ValueUnsignedInt_ext *Element `json:"_valueUnsignedInt"`
	// The value of the input parameter as a basic type.
	ValueAttachment *Attachment `json:"valueAttachment"`
	// The value of the input parameter as a basic type.
	ValuePeriod *Period `json:"valuePeriod"`
	// The value of the input parameter as a basic type.
	ValueContributor *Contributor `json:"valueContributor"`
	// The value of the input parameter as a basic type.
	ValueAnnotation *Annotation `json:"valueAnnotation"`
	// Extensions for valueUuid
	ValueUuid_ext *Element `json:"_valueUuid"`
	// Extensions for valueInteger
	ValueInteger_ext *Element `json:"_valueInteger"`
	// The value of the input parameter as a basic type.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	ValueOid string `json:"valueOid"`
	// The value of the input parameter as a basic type.
	ValueSignature *Signature `json:"valueSignature"`
	// The value of the input parameter as a basic type.
	ValueRatio *Ratio `json:"valueRatio"`
	// The value of the input parameter as a basic type.
	ValueSampledData *SampledData `json:"valueSampledData"`
	// The value of the input parameter as a basic type.
	ValueContactDetail *ContactDetail `json:"valueContactDetail"`
	// The value of the input parameter as a basic type.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	ValueDate time.Time `json:"valueDate"`
	// Extensions for valueId
	ValueId_ext *Element `json:"_valueId"`
	// The value of the input parameter as a basic type.
	ValueNarrative *Narrative `json:"valueNarrative"`
	// The value of the input parameter as a basic type.
	ValueSimpleQuantity *Quantity `json:"valueSimpleQuantity"`
	// The value of the input parameter as a basic type.
	ValueAddress *Address `json:"valueAddress"`
	// Extensions for valueInstant
	ValueInstant_ext *Element `json:"_valueInstant"`
	// The value of the input parameter as a basic type.
	ValueString string `json:"valueString"`
	// Extensions for valueUri
	ValueUri_ext *Element `json:"_valueUri"`
	// Extensions for valueDate
	ValueDate_ext *Element `json:"_valueDate"`
	// The value of the input parameter as a basic type.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	// The value of the input parameter as a basic type.
	ValueRange *Range `json:"valueRange"`
	// The value of the input parameter as a basic type.
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact"`
	// The value of the input parameter as a basic type.
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition"`
	// A code or description indicating how the input is intended to be used as part of the task execution.
	Type *CodeableConcept `json:"type,omitempty"`
	// The value of the input parameter as a basic type.
	// pattern [^\s]+([\s]?[^\s]+)*
	ValueCode string `json:"valueCode"`
	// The value of the input parameter as a basic type.
	// pattern [A-Za-z0-9\-\.]{1,64}
	ValueId string `json:"valueId"`
	// The value of the input parameter as a basic type.
	// pattern [1-9][0-9]*
	ValuePositiveInt uint64 `json:"valuePositiveInt"`
	// The value of the input parameter as a basic type.
	ValueAge *Age `json:"valueAge"`
	// The value of the input parameter as a basic type.
	ValueElementDefinition *ElementDefinition `json:"valueElementDefinition"`
	// The value of the input parameter as a basic type.
	ValueBase64Binary string `json:"valueBase64Binary"`
	// The value of the input parameter as a basic type.
	ValueElement *Element `json:"valueElement"`
	// The value of the input parameter as a basic type.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// The value of the input parameter as a basic type.
	ValueDistance *Distance `json:"valueDistance"`
	// The value of the input parameter as a basic type.
	ValueInstant string `json:"valueInstant"`
	// Extensions for valueTime
	ValueTime_ext *Element `json:"_valueTime"`
	// The value of the input parameter as a basic type.
	ValueHumanName *HumanName `json:"valueHumanName"`
	// The value of the input parameter as a basic type.
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement"`
	// The value of the input parameter as a basic type.
	ValueBackboneElement *BackboneElement `json:"valueBackboneElement"`
	// The value of the input parameter as a basic type.
	ValueContactPoint *ContactPoint `json:"valueContactPoint"`
	// Extensions for valueOid
	ValueOid_ext *Element `json:"_valueOid"`
	// The value of the input parameter as a basic type.
	// pattern urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	ValueUuid string `json:"valueUuid"`
	// Extensions for valuePositiveInt
	ValuePositiveInt_ext *Element `json:"_valuePositiveInt"`
	// The value of the input parameter as a basic type.
	ValueDuration *Duration `json:"valueDuration"`
	// The value of the input parameter as a basic type.
	ValueTiming *Timing `json:"valueTiming"`
	// The value of the input parameter as a basic type.
	ValueDosage *Dosage `json:"valueDosage"`
	// The value of the input parameter as a basic type.
	// pattern -?([0]|([1-9][0-9]*))
	ValueInteger int64 `json:"valueInteger"`
	// Extensions for valueDecimal
	ValueDecimal_ext *Element `json:"_valueDecimal"`
	// Extensions for valueBase64Binary
	ValueBase64Binary_ext *Element `json:"_valueBase64Binary"`
	// The value of the input parameter as a basic type.
	ValueUri string `json:"valueUri"`
	// Extensions for valueCode
	ValueCode_ext *Element `json:"_valueCode"`
	// The value of the input parameter as a basic type.
	ValueMoney *Money `json:"valueMoney"`
	// The value of the input parameter as a basic type.
	// pattern [0]|([1-9][0-9]*)
	ValueUnsignedInt uint64 `json:"valueUnsignedInt"`
	// The value of the input parameter as a basic type.
	ValueMarkdown string `json:"valueMarkdown"`
	// Extensions for valueMarkdown
	ValueMarkdown_ext *Element `json:"_valueMarkdown"`
	// The value of the input parameter as a basic type.
	ValueExtension *Extension `json:"valueExtension"`
	// The value of the input parameter as a basic type.
	ValueCount *Count `json:"valueCount"`
	// The value of the input parameter as a basic type.
	ValueReference *Reference `json:"valueReference"`
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// The value of the input parameter as a basic type.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	ValueDecimal float64 `json:"valueDecimal"`
	// Extensions for valueDateTime
	ValueDateTime_ext *Element `json:"_valueDateTime"`
	// The value of the input parameter as a basic type.
	ValueIdentifier *Identifier `json:"valueIdentifier"`
	// The value of the input parameter as a basic type.
	ValueCoding *Coding `json:"valueCoding"`
	// The value of the input parameter as a basic type.
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition"`
	// The value of the input parameter as a basic type.
	ValueBoolean bool `json:"valueBoolean"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// The value of the input parameter as a basic type.
	ValueMeta *Meta `json:"valueMeta"`
	// The value of the input parameter as a basic type.
	ValueUsageContext *UsageContext `json:"valueUsageContext"`
}

// ValueSet_Contains is A value set specifies a set of codes drawn from one or more code systems.
type ValueSet_Contains struct {
	_ *BackboneElement
	// Extensions for system
	System_ext *Element `json:"_system"`
	// Extensions for inactive
	Inactive_ext *Element `json:"_inactive"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// Additional representations for this item - other languages, aliases, specialized purposes, used for particular purposes, etc. These are relevant when the conditions of the expansion do not fix to a single correct representation.
	Designation []*ValueSet_Designation `json:"designation"`
	// Other codes and entries contained under this entry in the hierarchy.
	Contains []*ValueSet_Contains `json:"contains"`
	// If the concept is inactive in the code system that defines it. Inactive codes are those that are no longer to be used, but are maintained by the code system for understanding legacy data.
	Inactive bool `json:"inactive"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// If true, this entry is included in the expansion for navigational purposes, and the user cannot select the code directly as a proper value.
	Abstract bool `json:"abstract"`
	// The version of this code system that defined this code and/or display. This should only be used with code systems that do not enforce concept permanence.
	Version string `json:"version"`
	// The code for this item in the expansion hierarchy. If this code is missing the entry in the hierarchy is a place holder (abstract) and does not represent a valid code in the value set.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// The recommended display for this item in the expansion.
	Display string `json:"display"`
	// An absolute URI which is the code system in which the code for this item in the expansion is defined.
	System string `json:"system"`
	// Extensions for abstract
	Abstract_ext *Element `json:"_abstract"`
	// Extensions for display
	Display_ext *Element `json:"_display"`
}

// CapabilityStatement_Interaction is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Interaction struct {
	_ *BackboneElement
	// Coded identifier of the operation, supported by the system resource.
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// Guidance specific to the implementation of this operation, such as 'delete is a logical delete' or 'updates are only allowed with version id' or 'creates permitted from pre-authorized certificates only'.
	Documentation string `json:"documentation"`
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
}

// Composition_Attester is A set of healthcare-related information that is assembled together into a single logical document that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. While a Composition defines the structure, it does not actually contain the content: rather the full content of a document is contained in a Bundle, of which the Composition is the first resource contained.
type Composition_Attester struct {
	_ *BackboneElement
	// When the composition was attested by the party.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Time time.Time `json:"time"`
	// Extensions for time
	Time_ext *Element `json:"_time"`
	// Who attested the composition in the specified way.
	Party *Reference `json:"party"`
	// The type of attestation the authenticator offers.
	Mode []string `json:"mode"`
	// Extensions for mode
	Mode_ext []*Element `json:"_mode"`
}

// ExplanationOfBenefit_AddItem is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_AddItem struct {
	_ *BackboneElement
	// Extensions for sequenceLinkId
	SequenceLinkId_ext []*Element `json:"_sequenceLinkId"`
	// Health Care Service Type Codes  to identify the classification of service or benefits.
	Category *CodeableConcept `json:"category"`
	// If this is an actual service or product line, ie. not a Group, then use code to indicate the Professional Service or Product supplied (eg. CTP, HCPCS,USCLS,ICD10, NCPDP,DIN,ACHI,CCI). If a grouping item then use a group code to indicate the type of thing being grouped eg. 'glasses' or 'compound'.
	Service *CodeableConcept `json:"service"`
	// Item typification or modifiers codes, eg for Oral whether the treatment is cosmetic or associated with TMJ, or for medical whether the treatment was outside the clinic or out of office hours.
	Modifier []*CodeableConcept `json:"modifier"`
	// The second tier service adjudications for payor added services.
	Detail []*ExplanationOfBenefit_Detail1 `json:"detail"`
	// List of input service items which this service line is intended to replace.
	SequenceLinkId []uint64 `json:"sequenceLinkId"`
	// The type of reveneu or cost center providing the product and/or service.
	Revenue *CodeableConcept `json:"revenue"`
	// The fee charged for the professional service or product.
	Fee *Money `json:"fee"`
	// A list of note references to the notes provided below.
	NoteNumber []uint64 `json:"noteNumber"`
	// Extensions for noteNumber
	NoteNumber_ext []*Element `json:"_noteNumber"`
	// The adjudications results.
	Adjudication []*ExplanationOfBenefit_Adjudication `json:"adjudication"`
}

// Provenance_Entity is Provenance of a resource is a record that describes entities and processes involved in producing and delivering or otherwise influencing that resource. Provenance provides a critical foundation for assessing authenticity, enabling trust, and allowing reproducibility. Provenance assertions are a form of contextual metadata and can themselves become important records with their own provenance. Provenance statement indicates clinical significance in terms of confidence in authenticity, reliability, and trustworthiness, integrity, and stage in lifecycle (e.g. Document Completion - has the artifact been legally authenticated), all of which may impact security, privacy, and trust policies.
type Provenance_Entity struct {
	_ *BackboneElement
	// Extensions for whatUri
	WhatUri_ext *Element `json:"_whatUri"`
	// Identity of the  Entity used. May be a logical or physical uri and maybe absolute or relative.
	WhatReference *Reference `json:"whatReference"`
	// Identity of the  Entity used. May be a logical or physical uri and maybe absolute or relative.
	WhatIdentifier *Identifier `json:"whatIdentifier"`
	// The entity is attributed to an agent to express the agent's responsibility for that entity, possibly along with other agents. This description can be understood as shorthand for saying that the agent was responsible for the activity which generated the entity.
	Agent []*Provenance_Agent `json:"agent"`
	// How the entity was used during the activity.
	Role string `json:"role"`
	// Extensions for role
	Role_ext *Element `json:"_role"`
	// Identity of the  Entity used. May be a logical or physical uri and maybe absolute or relative.
	WhatUri string `json:"whatUri"`
}

// CapabilityStatement_Operation is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Operation struct {
	_ *BackboneElement
	// The name of the operation or query. For an operation, this is the name  prefixed with $ and used in the URL. For a query, this is the name used in the _query parameter when the query is called.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Where the formal definition can be found.
	Definition *Reference `json:"definition,omitempty"`
}

// ExpansionProfile is Resource to define constraints on the Expansion of a FHIR ValueSet.
type ExpansionProfile struct {
	_ *DomainResource
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate expansion profile instances.
	UseContext []*UsageContext `json:"useContext"`
	// Specifies the language to be used for description in the expansions i.e. the language to be used for ValueSet.expansion.contains.display.
	// pattern [^\s]+([\s]?[^\s]+)*
	DisplayLanguage string `json:"displayLanguage"`
	// If the value set being expanded is incomplete (because it is too big to expand), return a limited expansion (a subset) with an indicator that expansion is incomplete, using the extension [http://hl7.org/fhir/StructureDefinition/valueset-toocostly](extension-valueset-toocostly.html).
	LimitedExpansion bool `json:"limitedExpansion"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The status of this expansion profile. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// The date  (and optionally time) when the expansion profile was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the expansion profile changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// A legal or geographic region in which the expansion profile is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Code system, or a particular version of a code system to be excluded from value set expansions.
	ExcludedSystem *ExpansionProfile_ExcludedSystem `json:"excludedSystem"`
	// Controls whether inactive concepts are included or excluded in value set expansions.
	ActiveOnly bool `json:"activeOnly"`
	// Extensions for excludePostCoordinated
	ExcludePostCoordinated_ext *Element `json:"_excludePostCoordinated"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Controls whether the value set definition is included or excluded in value set expansions.
	IncludeDefinition bool `json:"includeDefinition"`
	// Extensions for excludeNested
	ExcludeNested_ext *Element `json:"_excludeNested"`
	// Extensions for displayLanguage
	DisplayLanguage_ext *Element `json:"_displayLanguage"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// Controls whether concept designations are to be included or excluded in value set expansions.
	IncludeDesignations bool `json:"includeDesignations"`
	// Extensions for activeOnly
	ActiveOnly_ext *Element `json:"_activeOnly"`
	// A formal identifier that is used to identify this expansion profile when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier *Identifier `json:"identifier"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Fix use of a particular code system to a particular version.
	FixedVersion []*ExpansionProfile_FixedVersion `json:"fixedVersion"`
	// Extensions for includeDefinition
	IncludeDefinition_ext *Element `json:"_includeDefinition"`
	// A boolean value to indicate that this expansion profile is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// Extensions for includeDesignations
	IncludeDesignations_ext *Element `json:"_includeDesignations"`
	// Extensions for excludeNotForUI
	ExcludeNotForUI_ext *Element `json:"_excludeNotForUI"`
	// Extensions for limitedExpansion
	LimitedExpansion_ext *Element `json:"_limitedExpansion"`
	// This is a ExpansionProfile resource
	ResourceType string `json:"resourceType,omitempty"`
	// A natural language name identifying the expansion profile. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// The name of the individual or organization that published the expansion profile.
	Publisher string `json:"publisher"`
	// A free text natural language description of the expansion profile from a consumer's perspective.
	Description string `json:"description"`
	// Controls whether or not the value set expansion includes codes which cannot be displayed in user interfaces.
	ExcludeNotForUI bool `json:"excludeNotForUI"`
	// An absolute URI that is used to identify this expansion profile when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this expansion profile is (or will be) published. The URL SHOULD include the major version of the expansion profile. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// The identifier that is used to identify this version of the expansion profile when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the expansion profile author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// A set of criteria that provide the constraints imposed on the value set expansion by including or excluding designations.
	Designation *ExpansionProfile_Designation `json:"designation"`
	// Controls whether or not the value set expansion nests codes or not (i.e. ValueSet.expansion.contains.contains).
	ExcludeNested bool `json:"excludeNested"`
	// Controls whether or not the value set expansion includes post coordinated codes.
	ExcludePostCoordinated bool `json:"excludePostCoordinated"`
}

func NewExpansionProfile() *ExpansionProfile {
	return &ExpansionProfile{ResourceType: "ExpansionProfile"}
}

// ExplanationOfBenefit_Related is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Related struct {
	_ *BackboneElement
	// Other claims which are related to this claim such as prior claim versions or for related services.
	Claim *Reference `json:"claim"`
	// For example prior or umbrella.
	Relationship *CodeableConcept `json:"relationship"`
	// An alternate organizational reference to the case or file to which this particular claim pertains - eg Property/Casualy insurer claim # or Workers Compensation case # .
	Reference *Identifier `json:"reference"`
}

// ReferralRequest_Requester is Used to record and send details about a request for referral service or transfer of a patient to the care of another provider or provider organization.
type ReferralRequest_Requester struct {
	_ *BackboneElement
	// The device, practitioner, etc. who initiated the request.
	Agent *Reference `json:"agent,omitempty"`
	// The organization the device or practitioner was acting on behalf of.
	OnBehalfOf *Reference `json:"onBehalfOf"`
}

// ImplementationGuide_Package is A set of rules of how FHIR is used to solve a particular problem. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuide_Package struct {
	_ *BackboneElement
	// The name for the group, as used in page.package.
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Human readable text describing the package.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A resource that is part of the implementation guide. Conformance resources (value set, structure definition, capability statements etc.) are obvious candidates for inclusion, but any kind of resource can be included as an example resource.
	Resource []*ImplementationGuide_Resource `json:"resource,omitempty"`
}

// TestScript is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript struct {
	_ *DomainResource
	// An absolute URI that is used to identify this test script when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this test script is (or will be) published. The URL SHOULD include the major version of the test script. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// A formal identifier that is used to identify this test script when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier *Identifier `json:"identifier"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// A free text natural language description of the test script from a consumer's perspective.
	Description string `json:"description"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// An abstract server used in operations within this test script in the origin element.
	Origin []*TestScript_Origin `json:"origin"`
	// Reference to the profile to be used for validation.
	Profile []*Reference `json:"profile"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// A series of operations required to clean up after the all the tests are executed (successfully or otherwise).
	Teardown *TestScript_Teardown `json:"teardown"`
	// This is a TestScript resource
	ResourceType string `json:"resourceType,omitempty"`
	// The identifier that is used to identify this version of the test script when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the test script author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// The status of this test script. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// A boolean value to indicate that this test script is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Explaination of why this test script is needed and why it has been designed as it has.
	Purpose string `json:"purpose"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The name of the individual or organization that published the test script.
	Publisher string `json:"publisher"`
	// A legal or geographic region in which the test script is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// A series of required setup operations before tests are executed.
	Setup *TestScript_Setup `json:"setup"`
	// A short, descriptive, user-friendly title for the test script.
	Title string `json:"title"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate test script instances.
	UseContext []*UsageContext `json:"useContext"`
	// An abstract server used in operations within this test script in the destination element.
	Destination []*TestScript_Destination `json:"destination"`
	// Variable is set based either on element value in response body or on header field value in the response headers.
	Variable []*TestScript_Variable `json:"variable"`
	// Contains one or more rules.  Offers a way to group rules so assertions could reference the group of rules and have them all applied.
	Ruleset []*TestScript_Ruleset `json:"ruleset"`
	// A natural language name identifying the test script. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A copyright statement relating to the test script and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the test script.
	Copyright string `json:"copyright"`
	// The required capability must exist and are assumed to function correctly on the FHIR server being tested.
	Metadata *TestScript_Metadata `json:"metadata"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// Fixture in the test script - by reference (uri). All fixtures are required for the test script to execute.
	Fixture []*TestScript_Fixture `json:"fixture"`
	// Assert rule to be used in one or more asserts within the test script.
	Rule []*TestScript_Rule `json:"rule"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// The date  (and optionally time) when the test script was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the test script changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// Extensions for purpose
	Purpose_ext *Element `json:"_purpose"`
	// A test in this script.
	Test []*TestScript_Test `json:"test"`
}

func NewTestScript() *TestScript { return &TestScript{ResourceType: "TestScript"} }

// ImplementationGuide is A set of rules of how FHIR is used to solve a particular problem. This resource is used to gather all the parts of an implementation guide into a logical whole and to publish a computable definition of all the parts.
type ImplementationGuide struct {
	_ *DomainResource
	// A page / section in the implementation guide. The root page is the implementation guide home page.
	Page *ImplementationGuide_Page `json:"page"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// The date  (and optionally time) when the implementation guide was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the implementation guide changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Another implementation guide that this implementation depends on. Typically, an implementation guide uses value sets, profiles etc.defined in other implementation guides.
	Dependency []*ImplementationGuide_Dependency `json:"dependency"`
	// An absolute URI that is used to identify this implementation guide when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this implementation guide is (or will be) published. The URL SHOULD include the major version of the implementation guide. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The name of the individual or organization that published the implementation guide.
	Publisher string `json:"publisher"`
	// A copyright statement relating to the implementation guide and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the implementation guide.
	Copyright string `json:"copyright"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// The identifier that is used to identify this version of the implementation guide when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the implementation guide author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// A set of profiles that all resources covered by this implementation guide must conform to.
	Global []*ImplementationGuide_Global `json:"global"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate implementation guide instances.
	UseContext []*UsageContext `json:"useContext"`
	// This is a ImplementationGuide resource
	ResourceType string `json:"resourceType,omitempty"`
	// A free text natural language description of the implementation guide from a consumer's perspective.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// A binary file that is included in the  implementation guide when it is published.
	Binary []string `json:"binary"`
	// Extensions for fhirVersion
	FhirVersion_ext *Element `json:"_fhirVersion"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// A natural language name identifying the implementation guide. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// The status of this implementation guide. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// A boolean value to indicate that this implementation guide is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// A legal or geographic region in which the implementation guide is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// The version of the FHIR specification on which this ImplementationGuide is based - this is the formal version of the specification, without the revision number, e.g. [publication].[major].[minor], which is 3.0.1 for this version.
	// pattern [A-Za-z0-9\-\.]{1,64}
	FhirVersion string `json:"fhirVersion"`
	// A logical group of resources. Logical groups can be used when building pages.
	Package []*ImplementationGuide_Package `json:"package"`
	// Extensions for binary
	Binary_ext []*Element `json:"_binary"`
}

func NewImplementationGuide() *ImplementationGuide {
	return &ImplementationGuide{ResourceType: "ImplementationGuide"}
}

// MeasureReport is The MeasureReport resource contains the results of evaluating a measure.
type MeasureReport struct {
	_ *DomainResource
	// The report status. No data will be available until the report status is complete.
	Status string `json:"status"`
	// A reference to a Bundle containing the Resources that were used in the evaluation of this report.
	EvaluatedResources *Reference `json:"evaluatedResources"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Reporting Organization.
	ReportingOrganization *Reference `json:"reportingOrganization"`
	// This is a MeasureReport resource
	ResourceType string `json:"resourceType,omitempty"`
	// The type of measure report. This may be an individual report, which provides a single patient's score for the measure; a patient listing, which returns the list of patients that meet the various criteria in the measure; or a summary report, which returns a population count for each of the criteria in the measure.
	Type string `json:"type"`
	// Optional Patient if the report was requested for a single patient.
	Patient *Reference `json:"patient"`
	// The date this measure report was generated.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// The reporting period for which the report was calculated.
	Period *Period `json:"period,omitempty"`
	// The results of the calculation, one for each population group in the measure.
	Group []*MeasureReport_Group `json:"group"`
	// A formal identifier that is used to identify this report when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier *Identifier `json:"identifier"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// A reference to the Measure that was evaluated to produce this report.
	Measure *Reference `json:"measure,omitempty"`
}

func NewMeasureReport() *MeasureReport { return &MeasureReport{ResourceType: "MeasureReport"} }

// Coding is A reference to a code defined by a terminology system.
type Coding struct {
	_ *Element
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// A symbol in syntax defined by the system. The symbol may be a predefined code or an expression in a syntax defined by the coding system (e.g. post-coordination).
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// Indicates that this coding was chosen by a user directly - i.e. off a pick list of available items (codes or displays).
	UserSelected bool `json:"userSelected"`
	// The identification of the code system that defines the meaning of the symbol in the code.
	System string `json:"system"`
	// Extensions for system
	System_ext *Element `json:"_system"`
	// The version of the code system which was used when choosing this code. Note that a well-maintained code system does not need the version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured. and when the meaning is not guaranteed to be consistent, the version SHOULD be exchanged.
	Version string `json:"version"`
	// A representation of the meaning of the code in the system, following the rules of the system.
	Display string `json:"display"`
	// Extensions for display
	Display_ext *Element `json:"_display"`
	// Extensions for userSelected
	UserSelected_ext *Element `json:"_userSelected"`
}

// CapabilityStatement_Endpoint is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_Endpoint struct {
	_ *BackboneElement
	// A list of the messaging transport protocol(s) identifiers, supported by this endpoint.
	Protocol *Coding `json:"protocol,omitempty"`
	// The network address of the end-point. For solutions that do not use network addresses for routing, it can be just an identifier.
	Address string `json:"address"`
	// Extensions for address
	Address_ext *Element `json:"_address"`
}

// MedicationStatement is A record of a medication that is being consumed by a patient.   A MedicationStatement may indicate that the patient may be taking the medication now, or has taken the medication in the past or will be taking the medication in the future.  The source of this information can be the patient, significant other (such as a family member or spouse), or a clinician.  A common scenario where this information is captured is during the history taking process during a patient visit or stay.   The medication information may come from sources such as the patient's memory, from a prescription bottle,  or from a list of medications the patient, clinician or other party maintains
//
// The primary difference between a medication statement and a medication administration is that the medication administration has complete administration information and is based on actual administration information from the person who administered the medication.  A medication statement is often, if not always, less specific.  There is no required date/time when the medication was administered, in fact we only know that a source has reported the patient is taking this medication, where details such as time, quantity, or rate or even medication product may be incomplete or missing or less precise.  As stated earlier, the medication statement information may come from the patient's memory, from a prescription bottle or from a list of medications the patient, clinician or other party maintains.  Medication administration is more formal and is not missing detailed information.
type MedicationStatement struct {
	_ *DomainResource
	// This is a MedicationStatement resource
	ResourceType string `json:"resourceType,omitempty"`
	// A larger event of which this particular event is a component or step.
	PartOf []*Reference `json:"partOf"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for dateAsserted
	DateAsserted_ext *Element `json:"_dateAsserted"`
	// The person or organization that provided the information about the taking of this medication. Note: Use derivedFrom when a MedicationStatement is derived from other resources, e.g Claim or MedicationRequest.
	InformationSource *Reference `json:"informationSource"`
	// Provides extra information about the medication statement that is not conveyed by the other attributes.
	Note []*Annotation `json:"note"`
	// External identifier - FHIR will generate its own internal identifiers (probably URLs) which do not need to be explicitly managed by the resource.  The identifier here is one that would be used by another non-FHIR system - for example an automated medication pump would provide a record each time it operated; an administration while the patient was off the ward might be made with a different system and entered after the event.  Particularly important if these records have to be updated.
	Identifier []*Identifier `json:"identifier"`
	// The interval of time during which it is being asserted that the patient was taking the medication (or was not taking, when the wasNotGiven element is true).
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	EffectiveDateTime time.Time `json:"effectiveDateTime"`
	// The interval of time during which it is being asserted that the patient was taking the medication (or was not taking, when the wasNotGiven element is true).
	EffectivePeriod *Period `json:"effectivePeriod"`
	// Allows linking the MedicationStatement to the underlying MedicationRequest, or to other information that supports or is used to derive the MedicationStatement.
	DerivedFrom []*Reference `json:"derivedFrom"`
	// A code indicating why the medication was not taken.
	ReasonNotTaken []*CodeableConcept `json:"reasonNotTaken"`
	// A plan, proposal or order that is fulfilled in whole or in part by this event.
	BasedOn []*Reference `json:"basedOn"`
	// A code representing the patient or other source's judgment about the state of the medication used that this statement is about.  Generally this will be active or completed.
	Status string `json:"status"`
	// Indicates where type of medication statement and where the medication is expected to be consumed or administered.
	Category *CodeableConcept `json:"category"`
	// Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	MedicationReference *Reference `json:"medicationReference"`
	// Extensions for effectiveDateTime
	EffectiveDateTime_ext *Element `json:"_effectiveDateTime"`
	// The date when the medication statement was asserted by the information source.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	DateAsserted time.Time `json:"dateAsserted"`
	// The person, animal or group who is/was taking the medication.
	Subject *Reference `json:"subject,omitempty"`
	// The encounter or episode of care that establishes the context for this MedicationStatement.
	Context *Reference `json:"context"`
	// Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	MedicationCodeableConcept *CodeableConcept `json:"medicationCodeableConcept"`
	// Indicator of the certainty of whether the medication was taken by the patient.
	Taken string `json:"taken"`
	// Extensions for taken
	Taken_ext *Element `json:"_taken"`
	// A reason for why the medication is being/was taken.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Condition or observation that supports why the medication is being/was taken.
	ReasonReference []*Reference `json:"reasonReference"`
	// Indicates how the medication is/was or should be taken by the patient.
	Dosage []*Dosage `json:"dosage"`
}

func NewMedicationStatement() *MedicationStatement {
	return &MedicationStatement{ResourceType: "MedicationStatement"}
}

// ProcedureRequest is A record of a request for diagnostic investigations, treatments, or operations to be performed.
type ProcedureRequest struct {
	_ *DomainResource
	// The desired perfomer for doing the diagnostic testing.  For example, the surgeon, dermatopathologist, endoscopist, etc.
	Performer *Reference `json:"performer"`
	// An explanation or justification for why this diagnostic investigation is being requested in coded or textual form.   This is often for billing purposes.  May relate to the resources referred to in supportingInformation.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Indicates another resource that provides a justification for why this diagnostic investigation is being requested.   May relate to the resources referred to in supportingInformation.
	ReasonReference []*Reference `json:"reasonReference"`
	// Plan/proposal/order fulfilled by this request.
	BasedOn []*Reference `json:"basedOn"`
	// Indicates how quickly the ProcedureRequest should be addressed with respect to other requests.
	// pattern [^\s]+([\s]?[^\s]+)*
	Priority string `json:"priority"`
	// The date/time at which the diagnostic testing should occur.
	OccurrenceTiming *Timing `json:"occurrenceTiming"`
	// An encounter or episode of care that provides additional information about the healthcare context in which this request is made.
	Context *Reference `json:"context"`
	// The date/time at which the diagnostic testing should occur.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OccurrenceDateTime time.Time `json:"occurrenceDateTime"`
	// Key events in the history of the request.
	RelevantHistory []*Reference `json:"relevantHistory"`
	// When the request transitioned to being actionable.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AuthoredOn time.Time `json:"authoredOn"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Set this to true if the record is saying that the procedure should NOT be performed.
	DoNotPerform bool `json:"doNotPerform"`
	// If a CodeableConcept is present, it indicates the pre-condition for performing the procedure.  For example "pain", "on flare-up", etc.
	AsNeededCodeableConcept *CodeableConcept `json:"asNeededCodeableConcept"`
	// Additional clinical information about the patient or specimen that may influence the procedure or diagnostics or their interpretations.     This information includes diagnosis, clinical findings and other observations.  In laboratory ordering these are typically referred to as "ask at order entry questions (AOEs)".  This includes observations explicitly requested by the producer (filler) to provide context or supporting information needed to complete the order. For example,  reporting the amount of inspired oxygen for blood gas measurements.
	SupportingInfo []*Reference `json:"supportingInfo"`
	// Any other notes and comments made about the service request. For example, letting provider know that "patient hates needles" or other provider instructions.
	Note []*Annotation `json:"note"`
	// Extensions for priority
	Priority_ext *Element `json:"_priority"`
	// A code that classifies the procedure for searching, sorting and display purposes (e.g. "Surgical Procedure").
	Category []*CodeableConcept `json:"category"`
	// On whom or what the procedure or diagnostic is to be performed. This is usually a human patient, but can also be requested on animals, groups of humans or animals, devices such as dialysis machines, or even locations (typically for environmental scans).
	Subject *Reference `json:"subject,omitempty"`
	// The date/time at which the diagnostic testing should occur.
	OccurrencePeriod *Period `json:"occurrencePeriod"`
	// Extensions for asNeededBoolean
	AsNeededBoolean_ext *Element `json:"_asNeededBoolean"`
	// The request takes the place of the referenced completed or terminated request(s).
	Replaces []*Reference `json:"replaces"`
	// Whether the request is a proposal, plan, an original order or a reflex order.
	// pattern [^\s]+([\s]?[^\s]+)*
	Intent string `json:"intent"`
	// Extensions for intent
	Intent_ext *Element `json:"_intent"`
	// The individual who initiated the request and has responsibility for its activation.
	Requester *ProcedureRequest_Requester `json:"requester"`
	// Anatomic location where the procedure should be performed. This is the target site.
	BodySite []*CodeableConcept `json:"bodySite"`
	// A code that identifies a particular procedure, diagnostic investigation, or panel of investigations, that have been requested.
	Code *CodeableConcept `json:"code,omitempty"`
	// Extensions for authoredOn
	AuthoredOn_ext *Element `json:"_authoredOn"`
	// This is a ProcedureRequest resource
	ResourceType string `json:"resourceType,omitempty"`
	// Identifiers assigned to this order instance by the orderer and/or the receiver and/or order fulfiller.
	Identifier []*Identifier `json:"identifier"`
	// A shared identifier common to all procedure or diagnostic requests that were authorized more or less simultaneously by a single author, representing the composite or group identifier.
	Requisition *Identifier `json:"requisition"`
	// Desired type of performer for doing the diagnostic testing.
	PerformerType *CodeableConcept `json:"performerType"`
	// Protocol or definition followed by this request.
	Definition []*Reference `json:"definition"`
	// The status of the order.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Extensions for doNotPerform
	DoNotPerform_ext *Element `json:"_doNotPerform"`
	// Extensions for occurrenceDateTime
	OccurrenceDateTime_ext *Element `json:"_occurrenceDateTime"`
	// If a CodeableConcept is present, it indicates the pre-condition for performing the procedure.  For example "pain", "on flare-up", etc.
	AsNeededBoolean bool `json:"asNeededBoolean"`
	// One or more specimens that the laboratory procedure will use.
	Specimen []*Reference `json:"specimen"`
}

func NewProcedureRequest() *ProcedureRequest {
	return &ProcedureRequest{ResourceType: "ProcedureRequest"}
}

// ValueSet_Compose is A value set specifies a set of codes drawn from one or more code systems.
type ValueSet_Compose struct {
	_ *BackboneElement
	// If a locked date is defined, then the Content Logical Definition must be evaluated using the current version as of the locked date for referenced code system(s) and value set instances where ValueSet.compose.include.version is not defined.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	LockedDate time.Time `json:"lockedDate"`
	// Extensions for lockedDate
	LockedDate_ext *Element `json:"_lockedDate"`
	// Whether inactive codes - codes that are not approved for current use - are in the value set. If inactive = true, inactive codes are to be included in the expansion, if inactive = false, the inactive codes will not be included in the expansion. If absent, the behavior is determined by the implementation, or by the applicable ExpansionProfile (but generally, inactive codes would be expected to be included).
	Inactive bool `json:"inactive"`
	// Extensions for inactive
	Inactive_ext *Element `json:"_inactive"`
	// Include one or more codes from a code system or other value set(s).
	Include []*ValueSet_Include `json:"include,omitempty"`
	// Exclude one or more codes from the value set based on code system filters and/or other value sets.
	Exclude []*ValueSet_Include `json:"exclude"`
}

// Appointment_Participant is A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time. This may result in one or more Encounter(s).
type Appointment_Participant struct {
	_ *BackboneElement
	// Is this participant required to be present at the meeting. This covers a use-case where 2 doctors need to meet to discuss the results for a specific patient, and the patient is not required to be present.
	Required string `json:"required"`
	// Extensions for required
	Required_ext *Element `json:"_required"`
	// Participation status of the actor.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Role of participant in the appointment.
	Type []*CodeableConcept `json:"type"`
	// A Person, Location/HealthcareService or Device that is participating in the appointment.
	Actor *Reference `json:"actor"`
}

// Composition is A set of healthcare-related information that is assembled together into a single logical document that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. While a Composition defines the structure, it does not actually contain the content: rather the full content of a document is contained in a Bundle, of which the Composition is the first resource contained.
type Composition struct {
	_ *DomainResource
	// Logical identifier for the composition, assigned when created. This identifier stays constant as the composition is changed over time.
	Identifier *Identifier `json:"identifier"`
	// Specifies the particular kind of composition (e.g. History and Physical, Discharge Summary, Progress Note). This usually equates to the purpose of making the composition.
	Type *CodeableConcept `json:"type,omitempty"`
	// Who or what the composition is about. The composition can be about a person, (patient or healthcare practitioner), a device (e.g. a machine) or even a group of subjects (such as a document about a herd of livestock, or a set of patients that share a common exposure).
	Subject *Reference `json:"subject,omitempty"`
	// The root of the sections that make up the composition.
	Section []*Composition_Section `json:"section"`
	// A participant who has attested to the accuracy of the composition/document.
	Attester []*Composition_Attester `json:"attester"`
	// Relationships that this composition has with other compositions or documents that already exist.
	RelatesTo []*Composition_RelatesTo `json:"relatesTo"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Describes the clinical encounter or type of care this documentation is associated with.
	Encounter *Reference `json:"encounter"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// Official human-readable label for the composition.
	Title string `json:"title"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// Extensions for confidentiality
	Confidentiality_ext *Element `json:"_confidentiality"`
	// This is a Composition resource
	ResourceType string `json:"resourceType,omitempty"`
	// A categorization for the type of the composition - helps for indexing and searching. This may be implied by or derived from the code specified in the Composition Type.
	Class *CodeableConcept `json:"class"`
	// The composition editing time, when the composition was last logically changed by the author.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Identifies who is responsible for the information in the composition, not necessarily who typed it in.
	Author []*Reference `json:"author,omitempty"`
	// The code specifying the level of confidentiality of the Composition.
	// pattern [^\s]+([\s]?[^\s]+)*
	Confidentiality string `json:"confidentiality"`
	// The workflow/clinical status of this composition. The status is a marker for the clinical standing of the document.
	Status string `json:"status"`
	// Identifies the organization or group who is responsible for ongoing maintenance of and access to the composition/document information.
	Custodian *Reference `json:"custodian"`
	// The clinical service, such as a colonoscopy or an appendectomy, being documented.
	Event []*Composition_Event `json:"event"`
}

func NewComposition() *Composition { return &Composition{ResourceType: "Composition"} }

// Composition_RelatesTo is A set of healthcare-related information that is assembled together into a single logical document that provides a single coherent statement of meaning, establishes its own context and that has clinical attestation with regard to who is making the statement. While a Composition defines the structure, it does not actually contain the content: rather the full content of a document is contained in a Bundle, of which the Composition is the first resource contained.
type Composition_RelatesTo struct {
	_ *BackboneElement
	// The type of relationship that this composition has with anther composition or document.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// The target composition/document of this relationship.
	TargetIdentifier *Identifier `json:"targetIdentifier"`
	// The target composition/document of this relationship.
	TargetReference *Reference `json:"targetReference"`
}

// Media is A photo, video, or audio recording acquired or used in healthcare. The actual content may be inline or provided by direct reference.
type Media struct {
	_ *DomainResource
	// Extensions for occurrenceDateTime
	OccurrenceDateTime_ext *Element `json:"_occurrenceDateTime"`
	// Describes why the event occurred in coded or textual form.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Indicates the site on the subject's body where the media was collected (i.e. the target site).
	BodySite *CodeableConcept `json:"bodySite"`
	// Extensions for duration
	Duration_ext *Element `json:"_duration"`
	// This is a Media resource
	ResourceType string `json:"resourceType,omitempty"`
	// A procedure that is fulfilled in whole or in part by the creation of this media.
	BasedOn []*Reference `json:"basedOn"`
	// Details of the type of the media - usually, how it was acquired (what type of device). If images sourced from a DICOM system, are wrapped in a Media resource, then this is the modality.
	Subtype *CodeableConcept `json:"subtype"`
	// The name of the imaging view e.g. Lateral or Antero-posterior (AP).
	View *CodeableConcept `json:"view"`
	// The person who administered the collection of the image.
	Operator *Reference `json:"operator"`
	// The device used to collect the media.
	Device *Reference `json:"device"`
	// The duration of the recording in seconds - for audio and video.
	// pattern [0]|([1-9][0-9]*)
	Duration uint64 `json:"duration"`
	// Whether the media is a photo (still image), an audio recording, or a video recording.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Who/What this Media is a record of.
	Subject *Reference `json:"subject"`
	// The date and time(s) at which the media was collected.
	OccurrencePeriod *Period `json:"occurrencePeriod"`
	// Extensions for height
	Height_ext *Element `json:"_height"`
	// Extensions for width
	Width_ext *Element `json:"_width"`
	// Extensions for frames
	Frames_ext *Element `json:"_frames"`
	// Comments made about the media by the performer, subject or other participants.
	Note []*Annotation `json:"note"`
	// Identifiers associated with the image - these may include identifiers for the image itself, identifiers for the context of its collection (e.g. series ids) and context ids such as accession numbers or other workflow identifiers.
	Identifier []*Identifier `json:"identifier"`
	// The encounter or episode of care that establishes the context for this media.
	Context *Reference `json:"context"`
	// The date and time(s) at which the media was collected.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OccurrenceDateTime time.Time `json:"occurrenceDateTime"`
	// Height of the image in pixels (photo/video).
	// pattern [1-9][0-9]*
	Height uint64 `json:"height"`
	// Width of the image in pixels (photo/video).
	// pattern [1-9][0-9]*
	Width uint64 `json:"width"`
	// The number of frames in a photo. This is used with a multi-page fax, or an imaging acquisition context that takes multiple slices in a single image, or an animated gif. If there is more than one frame, this SHALL have a value in order to alert interface software that a multi-frame capable rendering widget is required.
	// pattern [1-9][0-9]*
	Frames uint64 `json:"frames"`
	// The actual content of the media - inline or by direct reference to the media source file.
	Content *Attachment `json:"content,omitempty"`
}

func NewMedia() *Media { return &Media{ResourceType: "Media"} }

// TestScript_Action1 is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Action1 struct {
	_ *BackboneElement
	// An operation would involve a REST request to a server.
	Operation *TestScript_Operation `json:"operation"`
	// Evaluates the results of previous operations to determine if the server under test behaves appropriately.
	Assert *TestScript_Assert `json:"assert"`
}

// CapabilityStatement_SupportedMessage is A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement_SupportedMessage struct {
	_ *BackboneElement
	// The mode of this event declaration - whether application is sender or receiver.
	Mode string `json:"mode"`
	// Extensions for mode
	Mode_ext *Element `json:"_mode"`
	// Points to a message definition that identifies the messaging event, message structure, allowed responses, etc.
	Definition *Reference `json:"definition,omitempty"`
}

// StructureMap_Rule is A Map of relationships between 2 structures that can be used to transform data.
type StructureMap_Rule struct {
	_ *BackboneElement
	// Extensions for documentation
	Documentation_ext *Element `json:"_documentation"`
	// Name of the rule for internal references.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Name string `json:"name"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// Source inputs to the mapping.
	Source []*StructureMap_Source `json:"source,omitempty"`
	// Content to create because of this mapping rule.
	Target []*StructureMap_Target `json:"target"`
	// Rules contained in this rule.
	Rule []*StructureMap_Rule `json:"rule"`
	// Which other rules to apply in the context of this rule.
	Dependent []*StructureMap_Dependent `json:"dependent"`
	// Documentation for this instance of data.
	Documentation string `json:"documentation"`
}

// EnrollmentResponse is This resource provides enrollment and plan details from the processing of an Enrollment resource.
type EnrollmentResponse struct {
	_ *DomainResource
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// A description of the status of the adjudication.
	Disposition string `json:"disposition"`
	// Extensions for disposition
	Disposition_ext *Element `json:"_disposition"`
	// The date when the enclosed suite of services were performed or completed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// The practitioner who is responsible for the services rendered to the patient.
	RequestProvider *Reference `json:"requestProvider"`
	// This is a EnrollmentResponse resource
	ResourceType string `json:"resourceType,omitempty"`
	// The Response business identifier.
	Identifier []*Identifier `json:"identifier"`
	// Original request resource reference.
	Request *Reference `json:"request"`
	// Processing status: error, complete.
	Outcome *CodeableConcept `json:"outcome"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// The Insurer who produced this adjudicated response.
	Organization *Reference `json:"organization"`
	// The organization which is responsible for the services rendered to the patient.
	RequestOrganization *Reference `json:"requestOrganization"`
}

func NewEnrollmentResponse() *EnrollmentResponse {
	return &EnrollmentResponse{ResourceType: "EnrollmentResponse"}
}

// ExplanationOfBenefit_Detail is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Detail struct {
	_ *BackboneElement
	// A service line number.
	// pattern [1-9][0-9]*
	Sequence uint64 `json:"sequence"`
	// Item typification or modifiers codes, eg for Oral whether the treatment is cosmetic or associated with TMJ, or for medical whether the treatment was outside the clinic or out of office hours.
	Modifier []*CodeableConcept `json:"modifier"`
	// For programs which require reson codes for the inclusion, covering, of this billed item under the program or sub-program.
	ProgramCode []*CodeableConcept `json:"programCode"`
	// A list of note references to the notes provided below.
	NoteNumber []uint64 `json:"noteNumber"`
	// The type of product or service.
	Type *CodeableConcept `json:"type,omitempty"`
	// The number of repetitions of a service or product.
	Quantity *Quantity `json:"quantity"`
	// If the item is a node then this is the fee for the product or service, otherwise this is the total of the fees for the children of the group.
	UnitPrice *Money `json:"unitPrice"`
	// Extensions for factor
	Factor_ext *Element `json:"_factor"`
	// The quantity times the unit price for an addittional service or product or charge. For example, the formula: unit Quantity * unit Price (Cost per Point) * factor Number  * points = net Amount. Quantity, factor and points are assumed to be 1 if not supplied.
	Net *Money `json:"net"`
	// The type of reveneu or cost center providing the product and/or service.
	Revenue *CodeableConcept `json:"revenue"`
	// Health Care Service Type Codes  to identify the classification of service or benefits.
	Category *CodeableConcept `json:"category"`
	// If this is an actual service or product line, ie. not a Group, then use code to indicate the Professional Service or Product supplied (eg. CTP, HCPCS,USCLS,ICD10, NCPDP,DIN,ACHI,CCI). If a grouping item then use a group code to indicate the type of thing being grouped eg. 'glasses' or 'compound'.
	Service *CodeableConcept `json:"service"`
	// Third tier of goods and services.
	SubDetail []*ExplanationOfBenefit_SubDetail `json:"subDetail"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Factor float64 `json:"factor"`
	// List of Unique Device Identifiers associated with this line item.
	Udi []*Reference `json:"udi"`
	// Extensions for noteNumber
	NoteNumber_ext []*Element `json:"_noteNumber"`
	// The adjudications results.
	Adjudication []*ExplanationOfBenefit_Adjudication `json:"adjudication"`
}

// PlanDefinition_Participant is This resource allows for the definition of various types of plans as a sharable, consumable, and executable artifact. The resource is general enough to support the description of a broad range of clinical artifacts such as clinical decision support rules, order sets and protocols.
type PlanDefinition_Participant struct {
	_ *BackboneElement
	// The type of participant in the action.
	Type string `json:"type"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// The role the participant should play in performing the described action.
	Role *CodeableConcept `json:"role"`
}

// RiskAssessment is An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
type RiskAssessment struct {
	_ *DomainResource
	// The type of the risk assessment performed.
	Code *CodeableConcept `json:"code"`
	// Extensions for occurrenceDateTime
	OccurrenceDateTime_ext *Element `json:"_occurrenceDateTime"`
	// Describes the expected outcome for the subject.
	Prediction []*RiskAssessment_Prediction `json:"prediction"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// A reference to the request that is fulfilled by this risk assessment.
	BasedOn *Reference `json:"basedOn"`
	// The status of the RiskAssessment, using the same statuses as an Observation.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// The patient or group the risk assessment applies to.
	Subject *Reference `json:"subject"`
	// The date (and possibly time) the risk assessment was performed.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OccurrenceDateTime time.Time `json:"occurrenceDateTime"`
	// For assessments or prognosis specific to a particular condition, indicates the condition being assessed.
	Condition *Reference `json:"condition"`
	// The reason the risk assessment was performed.
	ReasonCodeableConcept *CodeableConcept `json:"reasonCodeableConcept"`
	// A description of the steps that might be taken to reduce the identified risk(s).
	Mitigation string `json:"mitigation"`
	// A reference to a resource that this risk assessment is part of, such as a Procedure.
	Parent *Reference `json:"parent"`
	// The algorithm, process or mechanism used to evaluate the risk.
	Method *CodeableConcept `json:"method"`
	// The reason the risk assessment was performed.
	ReasonReference *Reference `json:"reasonReference"`
	// Additional comments about the risk assessment.
	Comment string `json:"comment"`
	// Business identifier assigned to the risk assessment.
	Identifier *Identifier `json:"identifier"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The date (and possibly time) the risk assessment was performed.
	OccurrencePeriod *Period `json:"occurrencePeriod"`
	// The provider or software application that performed the assessment.
	Performer *Reference `json:"performer"`
	// Indicates the source data considered as part of the assessment (FamilyHistory, Observations, Procedures, Conditions, etc.).
	Basis []*Reference `json:"basis"`
	// Extensions for mitigation
	Mitigation_ext *Element `json:"_mitigation"`
	// This is a RiskAssessment resource
	ResourceType string `json:"resourceType,omitempty"`
	// The encounter where the assessment was performed.
	Context *Reference `json:"context"`
}

func NewRiskAssessment() *RiskAssessment { return &RiskAssessment{ResourceType: "RiskAssessment"} }

// TestScript_RequestHeader is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_RequestHeader struct {
	_ *BackboneElement
	// The HTTP header field e.g. "Accept".
	Field string `json:"field"`
	// Extensions for field
	Field_ext *Element `json:"_field"`
	// The value of the header e.g. "application/fhir+xml".
	Value string `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
}

// Quantity is A measured amount (or an amount that can potentially be measured). Note that measured amounts include amounts that are not precisely quantified, including amounts involving arbitrary units and floating currencies.
type Quantity struct {
	_ *Element
	// Extensions for code
	Code_ext *Element `json:"_code"`
	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator string `json:"comparator"`
	// Extensions for comparator
	Comparator_ext *Element `json:"_comparator"`
	// A human-readable form of the unit.
	Unit string `json:"unit"`
	// Extensions for unit
	Unit_ext *Element `json:"_unit"`
	// The identification of the system that provides the coded form of the unit.
	System string `json:"system"`
	// A computer processable form of the unit in some unit representation system.
	// pattern [^\s]+([\s]?[^\s]+)*
	Code string `json:"code"`
	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	Value float64 `json:"value"`
	// Extensions for value
	Value_ext *Element `json:"_value"`
	// Extensions for system
	System_ext *Element `json:"_system"`
}

// Dosage is Indicates how the medication is/was taken or should be taken by the patient.
type Dosage struct {
	_ *Element
	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction string `json:"patientInstruction"`
	// When medication should be administered.
	Timing *Timing `json:"timing"`
	// Extensions for asNeededBoolean
	AsNeededBoolean_ext *Element `json:"_asNeededBoolean"`
	// Body site to administer to.
	Site *CodeableConcept `json:"site"`
	// How drug should enter body.
	Route *CodeableConcept `json:"route"`
	// Amount of medication per dose.
	DoseRange *Range `json:"doseRange"`
	// Amount of medication per dose.
	DoseSimpleQuantity *Quantity `json:"doseSimpleQuantity"`
	// Indicates the order in which the dosage instructions should be applied or interpreted.
	// pattern -?([0]|([1-9][0-9]*))
	Sequence int64 `json:"sequence"`
	// Amount of medication per unit of time.
	RateRatio *Ratio `json:"rateRatio"`
	// Amount of medication per unit of time.
	RateRange *Range `json:"rateRange"`
	// Upper limit on medication per lifetime of the patient.
	MaxDosePerLifetime *Quantity `json:"maxDosePerLifetime"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// Extensions for patientInstruction
	PatientInstruction_ext *Element `json:"_patientInstruction"`
	// Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option), or it indicates the precondition for taking the Medication (CodeableConcept).
	AsNeededCodeableConcept *CodeableConcept `json:"asNeededCodeableConcept"`
	// Upper limit on medication per unit of time.
	MaxDosePerPeriod *Ratio `json:"maxDosePerPeriod"`
	// Free text dosage instructions e.g. SIG.
	Text string `json:"text"`
	// Technique for administering medication.
	Method *CodeableConcept `json:"method"`
	// Amount of medication per unit of time.
	RateSimpleQuantity *Quantity `json:"rateSimpleQuantity"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option), or it indicates the precondition for taking the Medication (CodeableConcept).
	AsNeededBoolean bool `json:"asNeededBoolean"`
	// Upper limit on medication per administration.
	MaxDosePerAdministration *Quantity `json:"maxDosePerAdministration"`
	// Supplemental instruction - e.g. "with meals".
	AdditionalInstruction []*CodeableConcept `json:"additionalInstruction"`
}

// DataElement is The formal description of a single piece of information that can be gathered and reported.
type DataElement struct {
	_ *DomainResource
	// The identifier that is used to identify this version of the data element when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the data element author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version string `json:"version"`
	// The name of the individual or organization that published the data element.
	Publisher string `json:"publisher"`
	// Identifies a specification (other than a terminology) that the elements which make up the DataElement have some correspondence with.
	Mapping []*DataElement_Mapping `json:"mapping"`
	// This is a DataElement resource
	ResourceType string `json:"resourceType,omitempty"`
	// A formal identifier that is used to identify this data element when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []*Identifier `json:"identifier"`
	// Extensions for copyright
	Copyright_ext *Element `json:"_copyright"`
	// Identifies how precise the data element is in its definition.
	Stringency string `json:"stringency"`
	// Extensions for stringency
	Stringency_ext *Element `json:"_stringency"`
	// The status of this data element. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// A boolean value to indicate that this data element is authored for testing purposes (or education/evaluation/marketing), and is not intended to be used for genuine usage.
	Experimental bool `json:"experimental"`
	// A natural language name identifying the data element. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// A short, descriptive, user-friendly title for the data element.
	Title string `json:"title"`
	// Extensions for title
	Title_ext *Element `json:"_title"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate data element instances.
	UseContext []*UsageContext `json:"useContext"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for version
	Version_ext *Element `json:"_version"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The date  (and optionally time) when the data element was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the data element changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
	// A legal or geographic region in which the data element is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// A copyright statement relating to the data element and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the data element.
	Copyright string `json:"copyright"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// An absolute URI that is used to identify this data element when it is referenced in a specification, model, design or an instance. This SHALL be a URL, SHOULD be globally unique, and SHOULD be an address at which this data element is (or will be) published. The URL SHOULD include the major version of the data element. For more information see [Technical and Business Versions](resource.html#versions).
	Url string `json:"url"`
	// Extensions for experimental
	Experimental_ext *Element `json:"_experimental"`
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// Defines the structure, type, allowed values and other constraining characteristics of the data element.
	Element []*ElementDefinition `json:"element,omitempty"`
}

func NewDataElement() *DataElement { return &DataElement{ResourceType: "DataElement"} }

// ExplanationOfBenefit_Payee is This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit_Payee struct {
	_ *BackboneElement
	// Type of Party to be reimbursed: Subscriber, provider, other.
	Type *CodeableConcept `json:"type"`
	// organization | patient | practitioner | relatedperson.
	ResourceType *CodeableConcept `json:"resourceType"`
	// Party to be reimbursed: Subscriber, provider, other.
	Party *Reference `json:"party"`
}

// StructureMap_Source is A Map of relationships between 2 structures that can be used to transform data.
type StructureMap_Source struct {
	_ *BackboneElement
	// Specified minimum cardinality for the element. This is optional; if present, it acts an implicit check on the input content.
	// pattern -?([0]|([1-9][0-9]*))
	Min int64 `json:"min"`
	// A value to use if there is no existing value in the source object.
	// pattern -?([0]|([1-9][0-9]*))
	DefaultValueInteger int64 `json:"defaultValueInteger"`
	// A value to use if there is no existing value in the source object.
	DefaultValueDistance *Distance `json:"defaultValueDistance"`
	// A value to use if there is no existing value in the source object.
	DefaultValueHumanName *HumanName `json:"defaultValueHumanName"`
	// Named context for field, if a field is specified.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Variable string `json:"variable"`
	// Extensions for defaultValueInstant
	DefaultValueInstant_ext *Element `json:"_defaultValueInstant"`
	// A value to use if there is no existing value in the source object.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1]))?)?
	DefaultValueDate time.Time `json:"defaultValueDate"`
	// Extensions for element
	Element_ext *Element `json:"_element"`
	// A value to use if there is no existing value in the source object.
	DefaultValueString string `json:"defaultValueString"`
	// A value to use if there is no existing value in the source object.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	DefaultValueDateTime time.Time `json:"defaultValueDateTime"`
	// Extensions for defaultValuePositiveInt
	DefaultValuePositiveInt_ext *Element `json:"_defaultValuePositiveInt"`
	// A value to use if there is no existing value in the source object.
	DefaultValueBackboneElement *BackboneElement `json:"defaultValueBackboneElement"`
	// A value to use if there is no existing value in the source object.
	DefaultValueSimpleQuantity *Quantity `json:"defaultValueSimpleQuantity"`
	// A value to use if there is no existing value in the source object.
	DefaultValueSampledData *SampledData `json:"defaultValueSampledData"`
	// A value to use if there is no existing value in the source object.
	DefaultValueContactDetail *ContactDetail `json:"defaultValueContactDetail"`
	// FHIRPath expression  - must be true or the rule does not apply.
	Condition string `json:"condition"`
	// Extensions for defaultValueId
	DefaultValueId_ext *Element `json:"_defaultValueId"`
	// A value to use if there is no existing value in the source object.
	DefaultValueCodeableConcept *CodeableConcept `json:"defaultValueCodeableConcept"`
	// A value to use if there is no existing value in the source object.
	DefaultValueAddress *Address `json:"defaultValueAddress"`
	// Extensions for defaultValueDateTime
	DefaultValueDateTime_ext *Element `json:"_defaultValueDateTime"`
	// Extensions for defaultValueUuid
	DefaultValueUuid_ext *Element `json:"_defaultValueUuid"`
	// Extensions for defaultValueMarkdown
	DefaultValueMarkdown_ext *Element `json:"_defaultValueMarkdown"`
	// A value to use if there is no existing value in the source object.
	DefaultValueAge *Age `json:"defaultValueAge"`
	// A value to use if there is no existing value in the source object.
	DefaultValueTriggerDefinition *TriggerDefinition `json:"defaultValueTriggerDefinition"`
	// Extensions for listMode
	ListMode_ext *Element `json:"_listMode"`
	// Extensions for defaultValueBoolean
	DefaultValueBoolean_ext *Element `json:"_defaultValueBoolean"`
	// Extensions for defaultValueInteger
	DefaultValueInteger_ext *Element `json:"_defaultValueInteger"`
	// Extensions for defaultValueBase64Binary
	DefaultValueBase64Binary_ext *Element `json:"_defaultValueBase64Binary"`
	// A value to use if there is no existing value in the source object.
	DefaultValueInstant string `json:"defaultValueInstant"`
	// A value to use if there is no existing value in the source object.
	DefaultValueMarkdown string `json:"defaultValueMarkdown"`
	// A value to use if there is no existing value in the source object.
	DefaultValueElement *Element `json:"defaultValueElement"`
	// A value to use if there is no existing value in the source object.
	DefaultValueElementDefinition *ElementDefinition `json:"defaultValueElementDefinition"`
	// A value to use if there is no existing value in the source object.
	DefaultValueContactPoint *ContactPoint `json:"defaultValueContactPoint"`
	// A value to use if there is no existing value in the source object.
	DefaultValueDataRequirement *DataRequirement `json:"defaultValueDataRequirement"`
	// A value to use if there is no existing value in the source object.
	DefaultValueParameterDefinition *ParameterDefinition `json:"defaultValueParameterDefinition"`
	// Optional field for this source.
	Element string `json:"element"`
	// Extensions for defaultValueTime
	DefaultValueTime_ext *Element `json:"_defaultValueTime"`
	// Extensions for defaultValueUnsignedInt
	DefaultValueUnsignedInt_ext *Element `json:"_defaultValueUnsignedInt"`
	// A value to use if there is no existing value in the source object.
	DefaultValueIdentifier *Identifier `json:"defaultValueIdentifier"`
	// A value to use if there is no existing value in the source object.
	DefaultValueMeta *Meta `json:"defaultValueMeta"`
	// Type or variable this rule applies to.
	// pattern [A-Za-z0-9\-\.]{1,64}
	Context string `json:"context"`
	// Extensions for defaultValueString
	DefaultValueString_ext *Element `json:"_defaultValueString"`
	// Extensions for defaultValueDate
	DefaultValueDate_ext *Element `json:"_defaultValueDate"`
	// A value to use if there is no existing value in the source object.
	// pattern [1-9][0-9]*
	DefaultValuePositiveInt uint64 `json:"defaultValuePositiveInt"`
	// A value to use if there is no existing value in the source object.
	DefaultValueAnnotation *Annotation `json:"defaultValueAnnotation"`
	// A value to use if there is no existing value in the source object.
	DefaultValueQuantity *Quantity `json:"defaultValueQuantity"`
	// A value to use if there is no existing value in the source object.
	DefaultValueReference *Reference `json:"defaultValueReference"`
	// Specified type for the element. This works as a condition on the mapping - use for polymorphic elements.
	Type string `json:"type"`
	// A value to use if there is no existing value in the source object.
	DefaultValueBoolean bool `json:"defaultValueBoolean"`
	// A value to use if there is no existing value in the source object.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	DefaultValueTime time.Time `json:"defaultValueTime"`
	// A value to use if there is no existing value in the source object.
	// pattern urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
	DefaultValueUuid string `json:"defaultValueUuid"`
	// A value to use if there is no existing value in the source object.
	DefaultValueDuration *Duration `json:"defaultValueDuration"`
	// A value to use if there is no existing value in the source object.
	DefaultValueTiming *Timing `json:"defaultValueTiming"`
	// Extensions for type
	Type_ext *Element `json:"_type"`
	// Extensions for defaultValueDecimal
	DefaultValueDecimal_ext *Element `json:"_defaultValueDecimal"`
	// Extensions for defaultValueUri
	DefaultValueUri_ext *Element `json:"_defaultValueUri"`
	// Extensions for variable
	Variable_ext *Element `json:"_variable"`
	// Extensions for defaultValueOid
	DefaultValueOid_ext *Element `json:"_defaultValueOid"`
	// A value to use if there is no existing value in the source object.
	DefaultValueDosage *Dosage `json:"defaultValueDosage"`
	// A value to use if there is no existing value in the source object.
	DefaultValueRelatedArtifact *RelatedArtifact `json:"defaultValueRelatedArtifact"`
	// How to handle the list mode for this element.
	ListMode string `json:"listMode"`
	// Extensions for check
	Check_ext *Element `json:"_check"`
	// Specified maximum cardinality for the element - a number or a "*". This is optional; if present, it acts an implicit check on the input content (* just serves as documentation; it's the default value).
	Max string `json:"max"`
	// A value to use if there is no existing value in the source object.
	DefaultValueNarrative *Narrative `json:"defaultValueNarrative"`
	// A value to use if there is no existing value in the source object.
	DefaultValueCount *Count `json:"defaultValueCount"`
	// A value to use if there is no existing value in the source object.
	DefaultValueMoney *Money `json:"defaultValueMoney"`
	// Extensions for min
	Min_ext *Element `json:"_min"`
	// A value to use if there is no existing value in the source object.
	DefaultValueUri string `json:"defaultValueUri"`
	// A value to use if there is no existing value in the source object.
	// pattern [^\s]+([\s]?[^\s]+)*
	DefaultValueCode string `json:"defaultValueCode"`
	// A value to use if there is no existing value in the source object.
	// pattern [0]|([1-9][0-9]*)
	DefaultValueUnsignedInt uint64 `json:"defaultValueUnsignedInt"`
	// A value to use if there is no existing value in the source object.
	DefaultValueAttachment *Attachment `json:"defaultValueAttachment"`
	// A value to use if there is no existing value in the source object.
	DefaultValuePeriod *Period `json:"defaultValuePeriod"`
	// Extensions for context
	Context_ext *Element `json:"_context"`
	// A value to use if there is no existing value in the source object.
	// pattern -?([0]|([1-9][0-9]*))(\.[0-9]+)?
	DefaultValueDecimal float64 `json:"defaultValueDecimal"`
	// A value to use if there is no existing value in the source object.
	DefaultValueRatio *Ratio `json:"defaultValueRatio"`
	// A value to use if there is no existing value in the source object.
	DefaultValueSignature *Signature `json:"defaultValueSignature"`
	// A value to use if there is no existing value in the source object.
	DefaultValueUsageContext *UsageContext `json:"defaultValueUsageContext"`
	// Extensions for condition
	Condition_ext *Element `json:"_condition"`
	// FHIRPath expression  - must be true or the mapping engine throws an error instead of completing.
	Check string `json:"check"`
	// A value to use if there is no existing value in the source object.
	DefaultValueRange *Range `json:"defaultValueRange"`
	// Extensions for max
	Max_ext *Element `json:"_max"`
	// A value to use if there is no existing value in the source object.
	DefaultValueBase64Binary string `json:"defaultValueBase64Binary"`
	// Extensions for defaultValueCode
	DefaultValueCode_ext *Element `json:"_defaultValueCode"`
	// A value to use if there is no existing value in the source object.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	DefaultValueOid string `json:"defaultValueOid"`
	// A value to use if there is no existing value in the source object.
	// pattern [A-Za-z0-9\-\.]{1,64}
	DefaultValueId string `json:"defaultValueId"`
	// A value to use if there is no existing value in the source object.
	DefaultValueExtension *Extension `json:"defaultValueExtension"`
	// A value to use if there is no existing value in the source object.
	DefaultValueCoding *Coding `json:"defaultValueCoding"`
	// A value to use if there is no existing value in the source object.
	DefaultValueContributor *Contributor `json:"defaultValueContributor"`
}

// CarePlan_Detail is Describes the intention of how one or more practitioners intend to deliver care for a particular patient, group or community for a period of time, possibly limited to care for a specific condition or set of conditions.
type CarePlan_Detail struct {
	_ *BackboneElement
	// The period, timing or frequency upon which the described activity is to occur.
	ScheduledPeriod *Period `json:"scheduledPeriod"`
	// The period, timing or frequency upon which the described activity is to occur.
	ScheduledString string `json:"scheduledString"`
	// Identifies who's expected to be involved in the activity.
	Performer []*Reference `json:"performer"`
	// This provides a textual description of constraints on the intended activity occurrence, including relation to other activities.  It may also include objectives, pre-conditions and end-conditions.  Finally, it may convey specifics about the activity such as body site, method, route, etc.
	Description string `json:"description"`
	// Internal reference that identifies the goals that this activity is intended to contribute towards meeting.
	Goal []*Reference `json:"goal"`
	// Provides reason why the activity isn't yet started, is on hold, was cancelled, etc.
	StatusReason string `json:"statusReason"`
	// If true, indicates that the described activity is one that must NOT be engaged in when following the plan.  If false, indicates that the described activity is one that should be engaged in when following the plan.
	Prohibited bool `json:"prohibited"`
	// Extensions for prohibited
	Prohibited_ext *Element `json:"_prohibited"`
	// Extensions for scheduledString
	ScheduledString_ext *Element `json:"_scheduledString"`
	// Identifies the food, drug or other product to be consumed or supplied in the activity.
	ProductReference *Reference `json:"productReference"`
	// Detailed description of the type of planned activity; e.g. What lab test, what procedure, what kind of encounter.
	Code *CodeableConcept `json:"code"`
	// Identifies what progress is being made for the specific activity.
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// High-level categorization of the type of activity in a care plan.
	Category *CodeableConcept `json:"category"`
	// Extensions for statusReason
	StatusReason_ext *Element `json:"_statusReason"`
	// Identifies the quantity expected to be consumed in a given day.
	DailyAmount *Quantity `json:"dailyAmount"`
	// The period, timing or frequency upon which the described activity is to occur.
	ScheduledTiming *Timing `json:"scheduledTiming"`
	// Identifies the facility where the activity will occur; e.g. home, hospital, specific clinic, etc.
	Location *Reference `json:"location"`
	// Identifies the food, drug or other product to be consumed or supplied in the activity.
	ProductCodeableConcept *CodeableConcept `json:"productCodeableConcept"`
	// Identifies the quantity expected to be supplied, administered or consumed by the subject.
	Quantity *Quantity `json:"quantity"`
	// Identifies the protocol, questionnaire, guideline or other specification the planned activity should be conducted in accordance with.
	Definition *Reference `json:"definition"`
	// Provides the rationale that drove the inclusion of this particular activity as part of the plan or the reason why the activity was prohibited.
	ReasonCode []*CodeableConcept `json:"reasonCode"`
	// Provides the health condition(s) that drove the inclusion of this particular activity as part of the plan.
	ReasonReference []*Reference `json:"reasonReference"`
}

// DeviceComponent_ProductionSpecification is The characteristics, operational status and capabilities of a medical-related component of a medical device.
type DeviceComponent_ProductionSpecification struct {
	_ *BackboneElement
	// The specification type, such as, serial number, part number, hardware revision, software revision, etc.
	SpecType *CodeableConcept `json:"specType"`
	// The internal component unique identification. This is a provision for manufacture specific standard components using a private OID. 11073-10101 has a partition for private OID semantic that the manufacturer can make use of.
	ComponentId *Identifier `json:"componentId"`
	// The printable string defining the component.
	ProductionSpec string `json:"productionSpec"`
	// Extensions for productionSpec
	ProductionSpec_ext *Element `json:"_productionSpec"`
}

// DocumentReference_Content is A reference to a document.
type DocumentReference_Content struct {
	_ *BackboneElement
	// The document or URL of the document along with critical metadata to prove content has integrity.
	Attachment *Attachment `json:"attachment,omitempty"`
	// An identifier of the document encoding, structure, and template that the document conforms to beyond the base format indicated in the mimeType.
	Format *Coding `json:"format"`
}

// ElementDefinition_Binding is Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition_Binding struct {
	_ *BackboneElement
	// Extensions for strength
	Strength_ext *Element `json:"_strength"`
	// Describes the intended use of this particular set of codes.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Points to the value set or external definition (e.g. implicit value set) that identifies the set of codes to be used. If the binding refers to an explicit value set - the normal case - then use a Reference(ValueSet) preferably containing the canonical URL for the value set. If the reference is to an implicit value set - usually, an IETF RFC that defines a grammar, such as mime types - then use a uri.
	ValueSetUri string `json:"valueSetUri"`
	// Extensions for valueSetUri
	ValueSetUri_ext *Element `json:"_valueSetUri"`
	// Points to the value set or external definition (e.g. implicit value set) that identifies the set of codes to be used. If the binding refers to an explicit value set - the normal case - then use a Reference(ValueSet) preferably containing the canonical URL for the value set. If the reference is to an implicit value set - usually, an IETF RFC that defines a grammar, such as mime types - then use a uri.
	ValueSetReference *Reference `json:"valueSetReference"`
	// Indicates the degree of conformance expectations associated with this binding - that is, the degree to which the provided value set must be adhered to in the instances.
	Strength string `json:"strength"`
}

// ImagingManifest_Series is A text description of the DICOM SOP instances selected in the ImagingManifest; or the reason for, or significance of, the selection.
type ImagingManifest_Series struct {
	_ *BackboneElement
	// Identity and locating information of the selected DICOM SOP instances.
	Instance []*ImagingManifest_Instance `json:"instance,omitempty"`
	// Series instance UID of the SOP instances in the selection.
	// pattern urn:oid:(0|[1-9][0-9]*)(\.(0|[1-9][0-9]*))*
	Uid string `json:"uid"`
	// Extensions for uid
	Uid_ext *Element `json:"_uid"`
	// The network service providing access (e.g., query, view, or retrieval) for this series. See implementation notes for information about using DICOM endpoints. A series-level endpoint, if present, has precedence over a study-level endpoint with the same Endpoint.type.
	Endpoint []*Reference `json:"endpoint"`
}

// MedicationRequest_Requester is An order or request for both supply of the medication and the instructions for administration of the medication to a patient. The resource is called "MedicationRequest" rather than "MedicationPrescription" or "MedicationOrder" to generalize the use across inpatient and outpatient settings, including care plans, etc., and to harmonize with workflow patterns.
type MedicationRequest_Requester struct {
	_ *BackboneElement
	// The healthcare professional responsible for authorizing the initial prescription.
	Agent *Reference `json:"agent,omitempty"`
	// The organization the device or practitioner was acting on behalf of.
	OnBehalfOf *Reference `json:"onBehalfOf"`
}

// NamingSystem is A curated namespace that issues unique symbols within that namespace for the identification of concepts, people, devices, etc.  Represents a "System" used within the Identifier and Coding data types.
type NamingSystem struct {
	_ *DomainResource
	// Extensions for publisher
	Publisher_ext *Element `json:"_publisher"`
	// Extensions for responsible
	Responsible_ext *Element `json:"_responsible"`
	// For naming systems that are retired, indicates the naming system that should be used in their place (if any).
	ReplacedBy *Reference `json:"replacedBy"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// Indicates the purpose for the naming system - what kinds of things does it make unique?
	Kind string `json:"kind"`
	// Extensions for kind
	Kind_ext *Element `json:"_kind"`
	// Extensions for date
	Date_ext *Element `json:"_date"`
	// The name of the organization that is responsible for issuing identifiers or codes for this namespace and ensuring their non-collision.
	Responsible string `json:"responsible"`
	// Indicates how the system may be identified when referenced in electronic exchange.
	UniqueId []*NamingSystem_UniqueId `json:"uniqueId,omitempty"`
	// A natural language name identifying the naming system. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name string `json:"name"`
	// The name of the individual or organization that published the naming system.
	Publisher string `json:"publisher"`
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []*ContactDetail `json:"contact"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// The content was developed with a focus and intent of supporting the contexts that are listed. These terms may be used to assist with indexing and searching for appropriate naming system instances.
	UseContext []*UsageContext `json:"useContext"`
	// This is a NamingSystem resource
	ResourceType string `json:"resourceType,omitempty"`
	// The status of this naming system. Enables tracking the life-cycle of the content.
	Status string `json:"status"`
	// The date  (and optionally time) when the naming system was published. The date must change if and when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the naming system changes.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Date time.Time `json:"date"`
	// Categorizes a naming system for easier search by grouping related naming systems.
	Type *CodeableConcept `json:"type"`
	// A free text natural language description of the naming system from a consumer's perspective. Details about what the namespace identifies including scope, granularity, version labeling, etc.
	Description string `json:"description"`
	// A legal or geographic region in which the naming system is intended to be used.
	Jurisdiction []*CodeableConcept `json:"jurisdiction"`
	// Provides guidance on the use of the namespace, including the handling of formatting characters, use of upper vs. lower case, etc.
	Usage string `json:"usage"`
	// Extensions for usage
	Usage_ext *Element `json:"_usage"`
	// Extensions for name
	Name_ext *Element `json:"_name"`
}

func NewNamingSystem() *NamingSystem { return &NamingSystem{ResourceType: "NamingSystem"} }

// TestScript_Setup is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Setup struct {
	_ *BackboneElement
	// Action would contain either an operation or an assertion.
	Action []*TestScript_Action `json:"action,omitempty"`
}

// Condition is A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
type Condition struct {
	_ *DomainResource
	// Encounter during which the condition was first asserted.
	Context *Reference `json:"context"`
	// Supporting Evidence / manifestations that are the basis on which this condition is suspected or confirmed.
	Evidence []*Condition_Evidence `json:"evidence"`
	// Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
	OnsetString string `json:"onsetString"`
	// The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
	AbatementRange *Range `json:"abatementRange"`
	// Extensions for assertedDate
	AssertedDate_ext *Element `json:"_assertedDate"`
	// The anatomical location where this condition manifests itself.
	BodySite []*CodeableConcept `json:"bodySite"`
	// Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
	OnsetPeriod *Period `json:"onsetPeriod"`
	// The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
	AbatementString string `json:"abatementString"`
	// Extensions for verificationStatus
	VerificationStatus_ext *Element `json:"_verificationStatus"`
	// The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
	AbatementPeriod *Period `json:"abatementPeriod"`
	// A category assigned to the condition.
	Category []*CodeableConcept `json:"category"`
	// Extensions for onsetString
	OnsetString_ext *Element `json:"_onsetString"`
	// Extensions for abatementString
	AbatementString_ext *Element `json:"_abatementString"`
	// This is a Condition resource
	ResourceType string `json:"resourceType,omitempty"`
	// The verification status to support the clinical status of the condition.
	VerificationStatus string `json:"verificationStatus"`
	// Clinical stage or grade of a condition. May include formal severity assessments.
	Stage *Condition_Stage `json:"stage"`
	// Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
	OnsetAge *Age `json:"onsetAge"`
	// Extensions for abatementDateTime
	AbatementDateTime_ext *Element `json:"_abatementDateTime"`
	// The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
	AbatementAge *Age `json:"abatementAge"`
	// This records identifiers associated with this condition that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []*Identifier `json:"identifier"`
	// A subjective assessment of the severity of the condition as evaluated by the clinician.
	Severity *CodeableConcept `json:"severity"`
	// Indicates the patient or group who the condition record is associated with.
	Subject *Reference `json:"subject,omitempty"`
	// Extensions for onsetDateTime
	OnsetDateTime_ext *Element `json:"_onsetDateTime"`
	// The date on which the existance of the Condition was first asserted or acknowledged.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AssertedDate time.Time `json:"assertedDate"`
	// Extensions for clinicalStatus
	ClinicalStatus_ext *Element `json:"_clinicalStatus"`
	// Identification of the condition, problem or diagnosis.
	Code *CodeableConcept `json:"code"`
	// Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
	OnsetRange *Range `json:"onsetRange"`
	// The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	AbatementDateTime time.Time `json:"abatementDateTime"`
	// The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
	AbatementBoolean bool `json:"abatementBoolean"`
	// Extensions for abatementBoolean
	AbatementBoolean_ext *Element `json:"_abatementBoolean"`
	// Individual who is making the condition statement.
	Asserter *Reference `json:"asserter"`
	// Additional information about the Condition. This is a general notes/comments entry  for description of the Condition, its diagnosis and prognosis.
	Note []*Annotation `json:"note"`
	// The clinical status of the condition.
	// pattern [^\s]+([\s]?[^\s]+)*
	ClinicalStatus string `json:"clinicalStatus"`
	// Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	OnsetDateTime time.Time `json:"onsetDateTime"`
}

func NewCondition() *Condition { return &Condition{ResourceType: "Condition"} }

// TestScript_Operation is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Operation struct {
	_ *BackboneElement
	// Id of fixture used for extracting the [id],  [type], and [vid] for GET requests.
	// pattern [A-Za-z0-9\-\.]{1,64}
	TargetId string `json:"targetId"`
	// Extensions for url
	Url_ext *Element `json:"_url"`
	// Extensions for resource
	Resource_ext *Element `json:"_resource"`
	// Extensions for accept
	Accept_ext *Element `json:"_accept"`
	// The server where the request message is destined for.  Must be one of the server numbers listed in TestScript.destination section.
	// pattern -?([0]|([1-9][0-9]*))
	Destination int64 `json:"destination"`
	// Whether or not to implicitly send the request url in encoded format. The default is true to match the standard RESTful client behavior. Set to false when communicating with a server that does not support encoded url paths.
	EncodeRequestUrl bool `json:"encodeRequestUrl"`
	// Extensions for requestId
	RequestId_ext *Element `json:"_requestId"`
	// The fixture id (maybe new) to map to the response.
	// pattern [A-Za-z0-9\-\.]{1,64}
	ResponseId string `json:"responseId"`
	// The description would be used by test engines for tracking and reporting purposes.
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for encodeRequestUrl
	EncodeRequestUrl_ext *Element `json:"_encodeRequestUrl"`
	// Extensions for params
	Params_ext *Element `json:"_params"`
	// The content-type or mime-type to use for RESTful operation in the 'Content-Type' header.
	ContentType string `json:"contentType"`
	// Extensions for contentType
	ContentType_ext *Element `json:"_contentType"`
	// The server where the request message originates from.  Must be one of the server numbers listed in TestScript.origin section.
	// pattern -?([0]|([1-9][0-9]*))
	Origin int64 `json:"origin"`
	// Extensions for responseId
	ResponseId_ext *Element `json:"_responseId"`
	// Extensions for targetId
	TargetId_ext *Element `json:"_targetId"`
	// The type of the resource.  See http://build.fhir.org/resourcelist.html.
	// pattern [^\s]+([\s]?[^\s]+)*
	Resource string `json:"resource"`
	// The label would be used for tracking/logging purposes by test engines.
	Label string `json:"label"`
	// Extensions for sourceId
	SourceId_ext *Element `json:"_sourceId"`
	// Extensions for origin
	Origin_ext *Element `json:"_origin"`
	// The id of the fixture used as the body of a PUT or POST request.
	// pattern [A-Za-z0-9\-\.]{1,64}
	SourceId string `json:"sourceId"`
	// Complete request URL.
	Url string `json:"url"`
	// Server interaction or operation type.
	Type *Coding `json:"type"`
	// The content-type or mime-type to use for RESTful operation in the 'Accept' header.
	Accept string `json:"accept"`
	// The fixture id (maybe new) to map to the request.
	// pattern [A-Za-z0-9\-\.]{1,64}
	RequestId string `json:"requestId"`
	// Path plus parameters after [type].  Used to set parts of the request URL explicitly.
	Params string `json:"params"`
	// Extensions for label
	Label_ext *Element `json:"_label"`
	// Extensions for destination
	Destination_ext *Element `json:"_destination"`
	// Header elements would be used to set HTTP headers.
	RequestHeader []*TestScript_RequestHeader `json:"requestHeader"`
}

// Element is Base definition for all elements in a resource.
type Element struct {
	// unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id string `json:"id"`
	// Extensions for id
	Id_ext *Element `json:"_id"`
	// May be used to represent additional information that is not part of the basic definition of the element. In order to make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []*Extension `json:"extension"`
}

// Reference is A reference from one resource to another.
type Reference struct {
	_ *Element
	// A reference to a location at which the other resource is found. The reference may be a relative reference, in which case it is relative to the service base URL, or an absolute URL that resolves to the location where the resource is found. The reference may be version specific or not. If the reference is not to a FHIR RESTful server, then it should be assumed to be version specific. Internal fragment references (start with '#') refer to contained resources.
	Reference string `json:"reference"`
	// Extensions for reference
	Reference_ext *Element `json:"_reference"`
	// An identifier for the other resource. This is used when there is no way to reference the other resource directly, either because the entity is not available through a FHIR server, or because there is no way for the author of the resource to convert a known identifier to an actual location. There is no requirement that a Reference.identifier point to something that is actually exposed as a FHIR instance, but it SHALL point to a business concept that would be expected to be exposed as a FHIR instance, and that instance would need to be of a FHIR resource type allowed by the reference.
	Identifier *Identifier `json:"identifier"`
	// Plain text narrative that identifies the resource in addition to the resource reference.
	Display string `json:"display"`
	// Extensions for display
	Display_ext *Element `json:"_display"`
}

// ClaimResponse_Detail1 is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse_Detail1 struct {
	_ *BackboneElement
	// Health Care Service Type Codes  to identify the classification of service or benefits.
	Category *CodeableConcept `json:"category"`
	// A code to indicate the Professional Service or Product supplied.
	Service *CodeableConcept `json:"service"`
	// Item typification or modifiers codes, eg for Oral whether the treatment is cosmetic or associated with TMJ, or for medical whether the treatment was outside the clinic or out of office hours.
	Modifier []*CodeableConcept `json:"modifier"`
	// The fee charged for the professional service or product..
	Fee *Money `json:"fee"`
	// A list of note references to the notes provided below.
	NoteNumber []uint64 `json:"noteNumber"`
	// Extensions for noteNumber
	NoteNumber_ext []*Element `json:"_noteNumber"`
	// The adjudications results.
	Adjudication []*ClaimResponse_Adjudication `json:"adjudication"`
	// The type of reveneu or cost center providing the product and/or service.
	Revenue *CodeableConcept `json:"revenue"`
}

// ConceptMap_Group is A statement of relationships from one set of concepts to one or more other concepts - either code systems or data elements, or classes in class models.
type ConceptMap_Group struct {
	_ *BackboneElement
	// Mappings for an individual concept in the source to one or more concepts in the target.
	Element []*ConceptMap_Element `json:"element,omitempty"`
	// An absolute URI that identifies the Code System (if the source is a value set that crosses more than one code system).
	Source string `json:"source"`
	// The specific version of the code system, as determined by the code system authority.
	SourceVersion string `json:"sourceVersion"`
	// Extensions for targetVersion
	TargetVersion_ext *Element `json:"_targetVersion"`
	// Extensions for target
	Target_ext *Element `json:"_target"`
	// The specific version of the code system, as determined by the code system authority.
	TargetVersion string `json:"targetVersion"`
	// What to do when there is no match in the mappings in the group.
	Unmapped *ConceptMap_Unmapped `json:"unmapped"`
	// Extensions for source
	Source_ext *Element `json:"_source"`
	// Extensions for sourceVersion
	SourceVersion_ext *Element `json:"_sourceVersion"`
	// An absolute URI that identifies the code system of the target code (if the target is a value set that cross code systems).
	Target string `json:"target"`
}

// Consent_Actor is A record of a healthcare consumer's policy choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent_Actor struct {
	_ *BackboneElement
	// How the individual is involved in the resources content that is described in the consent.
	Role *CodeableConcept `json:"role,omitempty"`
	// The resource that identifies the actor. To identify a actors by type, use group to identify a set of actors by some property they share (e.g. 'admitting officers').
	Reference *Reference `json:"reference,omitempty"`
}

// DocumentManifest is A collection of documents compiled for a purpose together with metadata that applies to the collection.
type DocumentManifest struct {
	_ *DomainResource
	// When the document manifest was created for submission to the server (not necessarily the same thing as the actual resource last modified time, since it may be modified, replicated, etc.).
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	Created time.Time `json:"created"`
	// Identifies the source system, application, or software that produced the document manifest.
	Source string `json:"source"`
	// Extensions for source
	Source_ext *Element `json:"_source"`
	// Other identifiers associated with the document manifest, including version independent  identifiers.
	Identifier []*Identifier `json:"identifier"`
	// Specifies the kind of this set of documents (e.g. Patient Summary, Discharge Summary, Prescription, etc.). The type of a set of documents may be the same as one of the documents in it - especially if there is only one - but it may be wider.
	Type *CodeableConcept `json:"type"`
	// The list of Documents included in the manifest.
	Content []*DocumentManifest_Content `json:"content,omitempty"`
	// A single identifier that uniquely identifies this manifest. Principally used to refer to the manifest in non-FHIR contexts.
	MasterIdentifier *Identifier `json:"masterIdentifier"`
	// Identifies who is responsible for creating the manifest, and adding  documents to it.
	Author []*Reference `json:"author"`
	// Human-readable description of the source document. This is sometimes known as the "title".
	Description string `json:"description"`
	// Extensions for description
	Description_ext *Element `json:"_description"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The status of this document manifest.
	Status string `json:"status"`
	// Who or what the set of documents is about. The documents can be about a person, (patient or healthcare practitioner), a device (i.e. machine) or even a group of subjects (such as a document about a herd of farm animals, or a set of patients that share a common exposure). If the documents cross more than one subject, then more than one subject is allowed here (unusual use case).
	Subject *Reference `json:"subject"`
	// Extensions for created
	Created_ext *Element `json:"_created"`
	// A patient, practitioner, or organization for which this set of documents is intended.
	Recipient []*Reference `json:"recipient"`
	// Related identifiers or resources associated with the DocumentManifest.
	Related []*DocumentManifest_Related `json:"related"`
	// This is a DocumentManifest resource
	ResourceType string `json:"resourceType,omitempty"`
}

func NewDocumentManifest() *DocumentManifest {
	return &DocumentManifest{ResourceType: "DocumentManifest"}
}

// ClaimResponse_ProcessNote is This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse_ProcessNote struct {
	_ *BackboneElement
	// An integer associated with each note which may be referred to from each service line item.
	// pattern [1-9][0-9]*
	Number uint64 `json:"number"`
	// Extensions for number
	Number_ext *Element `json:"_number"`
	// The note purpose: Print/Display.
	Type *CodeableConcept `json:"type"`
	// The note text.
	Text string `json:"text"`
	// Extensions for text
	Text_ext *Element `json:"_text"`
	// The ISO-639-1 alpha 2 code in lower case for the language, optionally followed by a hyphen and the ISO-3166-1 alpha 2 code for the region in upper case; e.g. "en" for English, or "en-US" for American English versus "en-EN" for England English.
	Language *CodeableConcept `json:"language"`
}

// Coverage is Financial instrument which may be used to reimburse or pay for health care products and services.
type Coverage struct {
	_ *DomainResource
	// This is a Coverage resource
	ResourceType string `json:"resourceType,omitempty"`
	// Time period during which the coverage is in force. A missing start date indicates the start date isn't known, a missing end date means the coverage is continuing to be in force.
	Period *Period `json:"period"`
	// Extensions for sequence
	Sequence_ext *Element `json:"_sequence"`
	// The policy(s) which constitute this insurance coverage.
	Contract []*Reference `json:"contract"`
	// A unique identifier for a dependent under the coverage.
	Dependent string `json:"dependent"`
	// The insurer assigned ID for the Subscriber.
	SubscriberId string `json:"subscriberId"`
	// A suite of underwrite specific classifiers, for example may be used to identify a class of coverage or employer group, Policy, Plan.
	Grouping *Coverage_Grouping `json:"grouping"`
	// Extensions for dependent
	Dependent_ext *Element `json:"_dependent"`
	// The order of applicability of this coverage relative to other coverages which are currently inforce. Note, there may be gaps in the numbering and this does not imply primary, secondard etc. as the specific positioning of coverages depends upon the episode of care.
	// pattern [1-9][0-9]*
	Order uint64 `json:"order"`
	// The status of the resource instance.
	// pattern [^\s]+([\s]?[^\s]+)*
	Status string `json:"status"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The party who 'owns' the insurance policy,  may be an individual, corporation or the subscriber's employer.
	PolicyHolder *Reference `json:"policyHolder"`
	// The party who has signed-up for or 'owns' the contractual relationship to the policy or to whom the benefit of the policy for services rendered to them or their family is due.
	Subscriber *Reference `json:"subscriber"`
	// The insurer-specific identifier for the insurer-defined network of providers to which the beneficiary may seek treatment which will be covered at the 'in-network' rate, otherwise 'out of network' terms and conditions apply.
	Network string `json:"network"`
	// Extensions for network
	Network_ext *Element `json:"_network"`
	// The relationship of beneficiary (patient) to the subscriber.
	Relationship *CodeableConcept `json:"relationship"`
	// The program or plan underwriter or payor including both insurance and non-insurance agreements, such as patient-pay agreements. May provide multiple identifiers such as insurance company identifier or business identifier (BIN number).
	Payor []*Reference `json:"payor"`
	// An optional counter for a particular instance of the identified coverage which increments upon each renewal.
	Sequence string `json:"sequence"`
	// Extensions for order
	Order_ext *Element `json:"_order"`
	// The main (and possibly only) identifier for the coverage - often referred to as a Member Id, Certificate number, Personal Health Number or Case ID. May be constructed as the concatination of the Coverage.SubscriberID and the Coverage.dependant.
	Identifier []*Identifier `json:"identifier"`
	// The type of coverage: social program, medical plan, accident coverage (workers compensation, auto), group health or payment by an individual or organization.
	Type *CodeableConcept `json:"type"`
	// Extensions for subscriberId
	SubscriberId_ext *Element `json:"_subscriberId"`
	// The party who benefits from the insurance coverage., the patient when services are provided.
	Beneficiary *Reference `json:"beneficiary"`
}

func NewCoverage() *Coverage { return &Coverage{ResourceType: "Coverage"} }

// FamilyMemberHistory_Condition is Significant health events and conditions for a person related to the patient relevant in the context of care for the patient.
type FamilyMemberHistory_Condition struct {
	_ *BackboneElement
	// An area where general notes can be placed about this specific condition.
	Note []*Annotation `json:"note"`
	// The actual condition specified. Could be a coded condition (like MI or Diabetes) or a less specific string like 'cancer' depending on how much is known about the condition and the capabilities of the creating system.
	Code *CodeableConcept `json:"code,omitempty"`
	// Indicates what happened as a result of this condition.  If the condition resulted in death, deceased date is captured on the relation.
	Outcome *CodeableConcept `json:"outcome"`
	// Either the age of onset, range of approximate age or descriptive string can be recorded.  For conditions with multiple occurrences, this describes the first known occurrence.
	OnsetAge *Age `json:"onsetAge"`
	// Either the age of onset, range of approximate age or descriptive string can be recorded.  For conditions with multiple occurrences, this describes the first known occurrence.
	OnsetRange *Range `json:"onsetRange"`
	// Either the age of onset, range of approximate age or descriptive string can be recorded.  For conditions with multiple occurrences, this describes the first known occurrence.
	OnsetPeriod *Period `json:"onsetPeriod"`
	// Either the age of onset, range of approximate age or descriptive string can be recorded.  For conditions with multiple occurrences, this describes the first known occurrence.
	OnsetString string `json:"onsetString"`
	// Extensions for onsetString
	OnsetString_ext *Element `json:"_onsetString"`
}

// Observation is Measurements and simple assertions made about a patient, device or other subject.
type Observation struct {
	_ *DomainResource
	// The time or time-period the observed value is asserted as being true. For biological subjects - e.g. human patients - this is usually called the "physiologically relevant time". This is usually either the time of the procedure or of specimen collection, but very often the source of the date/time is not known, only the date/time itself.
	EffectivePeriod *Period `json:"effectivePeriod"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueQuantity *Quantity `json:"valueQuantity"`
	// Extensions for valueString
	ValueString_ext *Element `json:"_valueString"`
	// The information determined as a result of making the observation, if the information has a simple value.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	ValueDateTime time.Time `json:"valueDateTime"`
	// Extensions for comment
	Comment_ext *Element `json:"_comment"`
	// A code that classifies the general type of observation being made.
	Category []*CodeableConcept `json:"category"`
	// Describes what was observed. Sometimes this is called the observation "name".
	Code *CodeableConcept `json:"code,omitempty"`
	// Provides a reason why the expected value in the element Observation.value[x] is missing.
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason"`
	// Some observations have multiple component observations.  These component observations are expressed as separate code value pairs that share the same attributes.  Examples include systolic and diastolic component observations for blood pressure measurement and multiple component observations for genetics observations.
	Component []*Observation_Component `json:"component"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueString string `json:"valueString"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValuePeriod *Period `json:"valuePeriod"`
	// Extensions for issued
	Issued_ext *Element `json:"_issued"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueBoolean bool `json:"valueBoolean"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueSampledData *SampledData `json:"valueSampledData"`
	// The information determined as a result of making the observation, if the information has a simple value.
	// pattern ([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?
	ValueTime time.Time `json:"valueTime"`
	// A  reference to another resource (usually another Observation) whose relationship is defined by the relationship type code.
	Related []*Observation_Related `json:"related"`
	// The patient, or group of patients, location, or device whose characteristics (direct or indirect) are described by the observation and into whose record the observation is placed.  Comments: Indirect characteristics may be those of a specimen, fetus, donor,  other observer (for example a relative or EMT), or any observation made about the subject.
	Subject *Reference `json:"subject"`
	// The date and time this observation was made available to providers, typically after the results have been reviewed and verified.
	Issued string `json:"issued"`
	// Who was responsible for asserting the observed value as "true".
	Performer []*Reference `json:"performer"`
	// May include statements about significant, unexpected or unreliable values, or information about the source of the value where this may be relevant to the interpretation of the result.
	Comment string `json:"comment"`
	// Guidance on how to interpret the value by comparison to a normal or recommended range.
	ReferenceRange []*Observation_ReferenceRange `json:"referenceRange"`
	// A unique identifier assigned to this observation.
	Identifier []*Identifier `json:"identifier"`
	// The healthcare event  (e.g. a patient and healthcare provider interaction) during which this observation is made.
	Context *Reference `json:"context"`
	// Indicates the mechanism used to perform the observation.
	Method *CodeableConcept `json:"method"`
	// A plan, proposal or order that is fulfilled in whole or in part by this event.
	BasedOn []*Reference `json:"basedOn"`
	// The assessment made based on the result of the observation.  Intended as a simple compact code often placed adjacent to the result value in reports and flow sheets to signal the meaning/normalcy status of the result. Otherwise known as abnormal flag.
	Interpretation *CodeableConcept `json:"interpretation"`
	// Extensions for valueBoolean
	ValueBoolean_ext *Element `json:"_valueBoolean"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueAttachment *Attachment `json:"valueAttachment"`
	// Extensions for valueDateTime
	ValueDateTime_ext *Element `json:"_valueDateTime"`
	// The device used to generate the observation data.
	Device *Reference `json:"device"`
	// Extensions for effectiveDateTime
	EffectiveDateTime_ext *Element `json:"_effectiveDateTime"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept"`
	// The time or time-period the observed value is asserted as being true. For biological subjects - e.g. human patients - this is usually called the "physiologically relevant time". This is usually either the time of the procedure or of specimen collection, but very often the source of the date/time is not known, only the date/time itself.
	// pattern -?[0-9]{4}(-(0[1-9]|1[0-2])(-(0[0-9]|[1-2][0-9]|3[0-1])(T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\.[0-9]+)?(Z|(\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00)))?)?)?
	EffectiveDateTime time.Time `json:"effectiveDateTime"`
	// Extensions for valueTime
	ValueTime_ext *Element `json:"_valueTime"`
	// Extensions for status
	Status_ext *Element `json:"_status"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueRange *Range `json:"valueRange"`
	// The information determined as a result of making the observation, if the information has a simple value.
	ValueRatio *Ratio `json:"valueRatio"`
	// Indicates the site on the subject's body where the observation was made (i.e. the target site).
	BodySite *CodeableConcept `json:"bodySite"`
	// The specimen that was used when this observation was made.
	Specimen *Reference `json:"specimen"`
	// This is a Observation resource
	ResourceType string `json:"resourceType,omitempty"`
	// The status of the result value.
	Status string `json:"status"`
}

func NewObservation() *Observation { return &Observation{ResourceType: "Observation"} }

// StructureDefinition_Differential is A definition of a FHIR structure. This resource is used to describe the underlying resources, data types defined in FHIR, and also for describing extensions and constraints on resources and data types.
type StructureDefinition_Differential struct {
	_ *BackboneElement
	// Captures constraints on each element within the resource.
	Element []*ElementDefinition `json:"element,omitempty"`
}

// TestScript_Action2 is A structured set of tests against a FHIR server implementation to determine compliance against the FHIR specification.
type TestScript_Action2 struct {
	_ *BackboneElement
	// An operation would involve a REST request to a server.
	Operation *TestScript_Operation `json:"operation,omitempty"`
}
