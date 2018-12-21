package models

// TODO: It might be nice to move all of these constants into the generated code

type CapabilityStatementKind string
type ConditionalDeleteStatus string
type ConditionalReadStatus string
type MimeType string
type PublicationStatus string
type ReferenceHandlingPolicy string
type ResourceVersionPolicy string
type RestfulCapabilityMode string
type RestfulSecurityService string
type SearchParamType string
type SystemRestfulInteraction string
type TypeRestfulInteraction string
type UnknownContentCode string

const (
	TypeRestfulInteractionCreate          TypeRestfulInteraction = "create"
	TypeRestfulInteractionDelete          TypeRestfulInteraction = "delete"
	TypeRestfulInteractionHistoryInstance TypeRestfulInteraction = "history-instance"
	TypeRestfulInteractionHistoryType     TypeRestfulInteraction = "history-type"
	TypeRestfulInteractionPatch           TypeRestfulInteraction = "patch"
	TypeRestfulInteractionRead            TypeRestfulInteraction = "read"
	TypeRestfulInteractionSearchType      TypeRestfulInteraction = "search-type"
	TypeRestfulInteractionUpdate          TypeRestfulInteraction = "update"
	TypeRestfulInteractionVRead           TypeRestfulInteraction = "vread"

	PublicationStatusActive  PublicationStatus = "active"
	PublicationStatusDraft   PublicationStatus = "draft"
	PublicationStatusRetired PublicationStatus = "retired"
	PublicationStatusUnknown PublicationStatus = "unknown"

	ConditionalDeleteStatusNotSupported ConditionalDeleteStatus = "not-supported"
	ConditionalDeleteStatusSingle       ConditionalDeleteStatus = "single"
	ConditionalDeleteStatusMultiple     ConditionalDeleteStatus = "multiple"

	CapabilityStatementKindInstance     CapabilityStatementKind = "instance"
	CapabilityStatementKindCapability   CapabilityStatementKind = "capability"
	CapabilityStatementKindRequirements CapabilityStatementKind = "requirements"

	UnknownContentCodeNo         UnknownContentCode = "no"
	UnknownContentCodeExtensions UnknownContentCode = "extensions"
	UnknownContentCodeElements   UnknownContentCode = "elements"
	UnknownContentCodeBoth       UnknownContentCode = "both"

	MimeTypeXML  MimeType = "xml"
	MimeTypeJSON MimeType = "json"
	MimeTypeTTL  MimeType = "ttl"

	RestfulCapabilityModeClient RestfulCapabilityMode = "client"
	RestfulCapabilityModeServer RestfulCapabilityMode = "server"

	RestfulSecurityServiceOAuth        RestfulSecurityService = "OAuth"
	RestfulSecurityServiceSMART        RestfulSecurityService = "SMART-on-FHIR"
	RestfulSecurityServiceNTLM         RestfulSecurityService = "NTLM"
	RestfulSecurityServiceKerberos     RestfulSecurityService = "Kerberos"
	RestfulSecurityServiceCertificates RestfulSecurityService = "Certificates"

	ResourceVersionPolicyNoVersion       ResourceVersionPolicy = "no-version"
	ResourceVersionPolicyVersioned       ResourceVersionPolicy = "versioned"
	ResourceVersionPolicyVersionedUpdate ResourceVersionPolicy = "versioned-update"

	ConditionalReadStatusNotSupported  ConditionalReadStatus = "not-supported"
	ConditionalReadStatusModifiedSince ConditionalReadStatus = "modified-since"
	ConditionalReadStatusNotMatch      ConditionalReadStatus = "not-match"
	ConditionalReadStatusFullSupport   ConditionalReadStatus = "full-support"

	ReferenceHandlingPolicyLiteral  ReferenceHandlingPolicy = "literal"
	ReferenceHandlingPolicyLogical  ReferenceHandlingPolicy = "logical"
	ReferenceHandlingPolicyResolves ReferenceHandlingPolicy = "resolves"
	ReferenceHandlingPolicyEnforced ReferenceHandlingPolicy = "enforced"
	ReferenceHandlingPolicyLocal    ReferenceHandlingPolicy = "local"

	SearchParamTypeNumber    SearchParamType = "number"
	SearchParamTypeDate      SearchParamType = "date"
	SearchParamTypeString    SearchParamType = "string"
	SearchParamTypeToken     SearchParamType = "token"
	SearchParamTypeReference SearchParamType = "reference"
	SearchParamTypeComposite SearchParamType = "composite"
	SearchParamTypeQuantity  SearchParamType = "quantity"
	SearchParamTypeURI       SearchParamType = "uri"

	SystemRestfulInteractionTransaction   SystemRestfulInteraction = "transaction"
	SystemRestfulInteractionBatch         SystemRestfulInteraction = "batch"
	SystemRestfulInteractionSearchSystem  SystemRestfulInteraction = "search-system"
	SystemRestfulInteractionHistorySystem SystemRestfulInteraction = "history-system"
)

type ResourceValidationMode string

const (
	ResourceValidationModeCreate ResourceValidationMode = "create"
	ResourceValidationModeUpdate ResourceValidationMode = "update"
	ResourceValidationModeDelete ResourceValidationMode = "delete"
)

type IssueSeverity = string

const (
	IssueSeverityFatal       IssueSeverity = "fatal"
	IssueSeverityError       IssueSeverity = "error"
	IssueSeverityWarning     IssueSeverity = "warning"
	IssueSeverityInformation IssueSeverity = "information"
)

type IssueType = string

// https://www.hl7.org/fhir/valueset-issue-type.html
const (
	IssueTypeInvalid       IssueType = "invalid"
	IssueTypeInformational IssueType = "informational"
)
