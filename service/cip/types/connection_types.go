package types

import (
	"github.com/antihax/optional"
	"time"
)

type Connection struct {
	// Type of connection. Valid values are `WebhookConnection`, `ServiceNowConnection`.
	Type_ string `json:"type"`
	// Unique identifier for the connection.
	Id string `json:"id"`
	// Name of the connection.
	Name string `json:"name"`
	// Description of the connection.
	Description string `json:"description"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
}

type ConnectionDefinition struct {
	// Type of connection. Valid values are `WebhookDefinition`, `ServiceNowDefinition`.
	Type_ string `json:"type"`
	// Name of connection. Name should be a valid alphanumeric value.
	Name string `json:"name"`
	// Description of the connection.
	Description string `json:"description,omitempty"`
}

type ConnectionsOpts struct {
	Limit optional.Int32
	Token optional.String
}

type Header struct {
	// Name of the header field.
	Name string `json:"name"`
	// Value of the header field.
	Value string `json:"value"`
}

type ListConnectionsResponse struct {
	// List of connections.
	Data []Connection `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type ServiceNowConnection struct {
	// Type of connection. Valid values are `WebhookConnection`, `ServiceNowConnection`.
	Type_ string `json:"type"`
	// Unique identifier for the connection.
	Id string `json:"id"`
	// Name of the connection.
	Name string `json:"name"`
	// Description of the connection.
	Description string `json:"description"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// URL for the ServiceNow connection.
	Url string `json:"url"`
	// User name for the ServiceNow connection.
	Username string `json:"username"`
}

type ServiceNowDefinition struct {
	// Type of connection. Valid values are `WebhookDefinition`, `ServiceNowDefinition`.
	Type_ string `json:"type"`
	// Name of connection. Name should be a valid alphanumeric value.
	Name string `json:"name"`
	// Description of the connection.
	Description string `json:"description,omitempty"`
	// URL for the ServiceNow connection.
	Url string `json:"url"`
	// User name for the ServiceNow connection.
	Username string `json:"username"`
	// User password for the ServiceNow connection.
	Password string `json:"password"`
}

type TestConnectionResponse struct {
	// Status code of the response of the connection test.
	StatusCode int32 `json:"statusCode"`
	// Content of the response of the connection test.
	ResponseContent string `json:"responseContent"`
}

type WebhookConnection struct {
	// Type of connection. Valid values are `WebhookConnection`, `ServiceNowConnection`.
	Type_ string `json:"type"`
	// Unique identifier for the connection.
	Id string `json:"id"`
	// Name of the connection.
	Name string `json:"name"`
	// Description of the connection.
	Description string `json:"description"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// URL for the webhook connection.
	Url string `json:"url"`
	// List of access authorization headers.
	Headers []Header `json:"headers"`
	// List of custom webhook headers.
	CustomHeaders []Header `json:"customHeaders"`
	// Default payload of the webhook.
	DefaultPayload string `json:"defaultPayload"`
	// Type of webhook. Valid values are `AWSLambda`, `Azure`, `Datadog`, `HipChat`, `NewRelic`, `Opsgenie`, `PagerDuty`, `Slack`, `MicrosoftTeams` and `Webhook`. `MicrosoftTeams` webhooks are in beta and not available until given access. Please reach out to your Sumo Logic representative to add new webhook types.
	WebhookType string `json:"webhookType"`
	// Webhook endpoint warning for incorrect variable names and syntax.
	Warnings []string `json:"warnings,omitempty"`
}

type WebhookDefinition struct {
	// Type of connection. Valid values are `WebhookDefinition`, `ServiceNowDefinition`.
	Type_ string `json:"type"`
	// Name of connection. Name should be a valid alphanumeric value.
	Name string `json:"name"`
	// Description of the connection.
	Description string `json:"description,omitempty"`
	// URL for the webhook connection.
	Url string `json:"url"`
	// List of access authorization headers.
	Headers []Header `json:"headers,omitempty"`
	// List of custom webhook headers.
	CustomHeaders []Header `json:"customHeaders,omitempty"`
	// Default payload of the webhook.
	DefaultPayload string `json:"defaultPayload"`
	// Type of webhook. Valid values are `AWSLambda`, `Azure`, `Datadog`, `HipChat`, `NewRelic`, `Opsgenie`, `PagerDuty`, `Slack`, `MicrosoftTeams` and `Webhook`. `MicrosoftTeams` webhooks are in beta and not available until given access. Please reach out to your Sumo Logic representative to add new webhook types.
	WebhookType string `json:"webhookType,omitempty"`
}
