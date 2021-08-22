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
	// Resource type. Supported types are - `Collector`, `Source`, `IngestBudget`, `Organisation`, `LogsToMetricsRule`.
	Type_ string `json:"type"`
	// The identifier the collector this source belongs to (default is "Unknown") Only required when Type_ is set to `CollectorId`.
	CollectorId string `json:"collectorId,omitempty"`
	// The name of the collector this source belongs to (default is "Unknown"). Only required when Type_ is set to `Collector`.
	CollectorName string `json:"collectorName,omitempty"`
	// The unique field value of the ingest budget v1. This will be empty for v2 budgets. (default is "Unknown"). Only required when Type_ is set to `IngestBudget`.
	IngestBudgetFieldValue string `json:"ingestBudgetFieldValue"`
	// The scope of the ingest budget v2. This will be empty for v1 budgets. (default is "Unknown"). Only required when Type_ is set to `IngestBudget`.
	Scope string `json:"scope"`
}

type TrackerIdentity struct {
	// Name that uniquely identifies the health event. It focuses on what happened rather than why.
	TrackerId string `json:"trackerId"`
	// Description of the underlying reason for the event change.
	Error_ string `json:"error"`
	// A more elaborate description of why the event occurred.
	Description string `json:"description"`
}
