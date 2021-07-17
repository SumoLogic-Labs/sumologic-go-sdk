package types

type CreateEventHubSourceRequest struct {
	// ApiVersion version of the api
	ApiVersion string `json:"api.version,omitempty"`
	// Source type of source
	Source EventHubSource `json:"source"`
}

type EventHubSource struct {
	// SchemaRef type of source
	SchemaRef EventHubSourceSchema `json:"schemaRef"`
	// Config Event Hub source configuration
	Config EventHubSourceConfigurationDefinition `json:"config"`
	//SourceType type of source (use "Universal")
	SourceType string `json:"sourceType"`
}

type EventHubSourceConfigurationDefinition struct {
	// Name of the source
	Name string `json:"name"`
	// Description of the source
	Description string `json:"description,omitempty"`
	// Namespace Azure Event Hub namespace name
	Namespace string `json:"namespace"`
	// HubName Azure Event Hub name
	HubName string `json:"hub_name"`
	// AccessPolicyName Azure Event Hub Authorization Rule name
	AccessPolicyName string `json:"access_policy_name"`
	// AccessPolicyKey Azure Event Hub Access Key
	AccessPolicyKey string `json:"access_policy_key"`
	// ConsumerGroup specify a custom consumer group name if needed. If you don't need a custom group use "$Default"
	ConsumerGroup string `json:"consumer_group"`
	// Fields to apply to the source
	Fields map[string]string `json:"fields,omitempty"`
	// Category of the source
	Category string `json:"category"`
	// ReceiveWithLatestOffset receive data with the latest offset or from the timestamp
	ReceiveWithLatestOffset bool `json:"receive_with_latest_offset,omitempty"`
	// ReceiveFromTimestamp set to true when ReceiveWithLatestOffset is false
	ReceiveFromTimestamp bool `json:"receive_from_timestamp,omitempty"`
}

type EventHubModel struct {
	// ApiVersion version of the api
	ApiVersion string `json:"api.version"`
	// Source type of source
	Source EventHubSourceModel `json:"source"`
}

type EventHubSourceModel struct {
	// Id of source
	Id string `json:"id"`
	// SchemaRef type of source
	SchemaRef EventHubSourceSchema `json:"schemaRef"`
	// Config Event Hub source configuration
	Config EventHubSourceConfigurationDefinition `json:"config"`
	//SourceType type of source ("Universal")
	SourceType string `json:"sourceType"`
	// State of the Event Hub source
	State EventHubSourceStateModel `json:"state"`
}

type EventHubSourceStateModel struct {
	// State of the Event Hub source
	State string `json:"state"`
}

type EventHubSourceSchema struct {
	// Type of source (use "Azure Event Hubs")
	Type string `json:"type"`
}
