package types

import "time"

type ListTokensBaseResponse struct {
	// List of tokens.
	Data []TokenBaseResponse `json:"data"`
}

type TokenBaseDefinition struct {
	// Name of the token.
	Name string `json:"name"`
	// Description of the token.
	Description string `json:"description,omitempty"`
	// Status of the token. Can be `Active`, or `Inactive`.
	Status string `json:"status"`
	// Type of the token. Valid values: 1) CollectorRegistration
	Type_ string `json:"type"`
}

type TokenBaseResponse struct {
	// Identifier of the token.
	Id string `json:"id"`
	// Name of the token.
	Name string `json:"name"`
	// Description of the token.
	Description string `json:"description"`
	// Status of the token. Can be `Active`, or `Inactive`.
	Status string `json:"status"`
	// Type of the token. Valid values: 1) CollectorRegistrationTokenResponse
	Type_ string `json:"type"`
	// Version of the token.
	Version int64 `json:"version"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
}

type TokenBaseDefinitionUpdate struct {
	// Name of the token.
	Name string `json:"name"`
	// Description of the token.
	Description string `json:"description,omitempty"`
	// Status of the token. Can be `Active`, or `Inactive`.
	Status string `json:"status"`
	// Type of the token. Valid values: 1) CollectorRegistration
	Type_ string `json:"type"`
	// Version of the token.
	Version int64 `json:"version"`
}
