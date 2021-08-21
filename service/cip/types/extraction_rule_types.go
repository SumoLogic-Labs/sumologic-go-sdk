package types

import (
	"github.com/antihax/optional"
	"time"
)

type DroppedField struct {
	// Field name.
	FieldName string `json:"fieldName"`
}

type ExtractionRule struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Is the field extraction rule enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Unique identifier for the field extraction rule.
	Id string `json:"id"`
	// List of extracted fields from \"parseExpression\".
	FieldNames []string `json:"fieldNames,omitempty"`
}

type ExtractionRuleDefinition struct {
	// Name of the field extraction rule. Use a name that makes it easy to identify the rule.
	Name string `json:"name"`
	// Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You'll use the Scope to run a search against the rule.
	Scope string `json:"scope"`
	// Describes the fields to be parsed.
	ParseExpression string `json:"parseExpression"`
	// Is the field extraction rule enabled.
	Enabled bool `json:"enabled,omitempty"`
}

type ExtractionRuleManagementApiListExtractionRulesOpts struct {
	Limit optional.Int32
	Token optional.String
}

type ListExtractionRulesResponse struct {
	// List of field extraction rules.
	Data []ExtractionRule `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type UpdateExtractionRuleDefinition struct {
	// Name of the field extraction rule. Use a name that makes it easy to identify the rule.
	Name string `json:"name"`
	// Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You'll use the Scope to run a search against the rule.
	Scope string `json:"scope"`
	// Describes the fields to be parsed.
	ParseExpression string `json:"parseExpression"`
	// Is the field extraction rule enabled.
	Enabled bool `json:"enabled"`
}
