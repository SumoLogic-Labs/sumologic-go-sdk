package types

import "github.com/antihax/optional"

type DynamicRule struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt string `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt string `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Name of the dynamic parsing rule. Use a name that makes it easy to identify the rule.
	Name string `json:"name"`
	// Scope of the dynamic parsing rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You'll use the Scope to run a search against the rule.
	Scope string `json:"scope"`
	// Is the dynamic parsing rule enabled.
	Enabled bool `json:"enabled"`
	// Unique identifier for the dynamic parsing rule.
	Id string `json:"id"`
}

type DynamicRuleDefinition struct {
	// Name of the dynamic parsing rule. Use a name that makes it easy to identify the rule.
	Name string `json:"name"`
	// Scope of the dynamic parsing rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You'll use the Scope to run a search against the rule.
	Scope string `json:"scope"`
	// Is the dynamic parsing rule enabled.
	Enabled bool `json:"enabled"`
}

type DynamicParsingRuleManagementApiListDynamicParsingRulesOpts struct {
	Limit optional.Int32
	Token optional.String
}

type ListDynamicRulesResponse struct {
	// List of dynamic parsing rules.
	Data []DynamicRule `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}
