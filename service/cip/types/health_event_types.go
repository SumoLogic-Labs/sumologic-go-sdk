package types

import (
	"github.com/antihax/optional"
	"time"
)

type HealthEvent struct {
	// The unique identifier of the event.
	EventId string `json:"eventId"`
	// The name of the event.
	EventName        string            `json:"eventName"`
	Details          *TrackerIdentity  `json:"details"`
	ResourceIdentity *ResourceIdentity `json:"resourceIdentity"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	EventTime time.Time `json:"eventTime"`
	// The product area of the event.
	Subsystem string `json:"subsystem"`
	// The criticality of the event. It is either `Error` or `Warning`
	SeverityLevel string `json:"severityLevel"`
}

type HealthEventsOpts struct {
	Limit optional.Int32
	Token optional.String
}

type IngestBudgetIdentities struct {
	// A list of source resources
	Data []IngestBudgetIdentity `json:"data"`
}

type IngestBudgetIdentity struct {
	// The unique identifier of the resource
	Id string `json:"id"`
	// The name of the resource
	Name string `json:"name"`
	// Resource type. Supported type for this type is `IngestBudget`.
	Type_ string `json:"type"`
	// The unique field value of the ingest budget v1. This will be empty for v2 budgets. (default is "Unknown")
	IngestBudgetFieldValue string `json:"ingestBudgetFieldValue"`
	// The scope of the ingest budget v2. This will be empty for v1 budgets. (default is "Unknown")
	Scope string `json:"scope"`
}

type ListHealthEventResponse struct {
	// List of health events.
	Data []HealthEvent `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type ResourceIdentities struct {
	// A list of the resources.
	Data []ResourceIdentity `json:"data"`
}

type ResourceIdentity struct {
	// The unique identifier of the resource.
	Id string `json:"id"`
	// The name of the resource.
	Name string `json:"name,omitempty"`
	// Resource type. Supported types are - `Collector`, `Organisation`, `LogsToMetricsRule`.
	Type_ string `json:"type"`
}

type SourceIdentities struct {
	// A list of source resources
	Data []SourceIdentity `json:"data"`
}

type SourceIdentity struct {
	// The unique identifier of the resource
	Id string `json:"id"`
	// The name of the resource
	Name string `json:"name"`
	// Resource type. Supported type for this type is `Source`.
	Type_ string `json:"type"`
	// The identifier the collector this source belongs to (default is "Unknown")
	CollectorId string `json:"collectorId,omitempty"`
	// The name of the collector this source belongs to (default is "Unknown")
	CollectorName string `json:"collectorName,omitempty"`
}

type TrackerIdentity struct {
	// Name that uniquely identifies the health event. It focuses on what happened rather than why.
	TrackerId string `json:"trackerId"`
	// Description of the underlying reason for the event change.
	Error_ string `json:"error"`
	// A more elaborate description of why the event occurred.
	Description string `json:"description"`
}
