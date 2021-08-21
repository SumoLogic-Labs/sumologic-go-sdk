package types

import (
	"github.com/antihax/optional"
	"time"
)

type IngestBudget struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt     time.Time `json:"createdAt"`
	CreatedByUser *UserInfo `json:"createdByUser"`
	// Last modification timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	ModifiedAt     time.Time `json:"modifiedAt"`
	ModifiedByUser *UserInfo `json:"modifiedByUser"`
	// Display name of the ingest budget.
	Name string `json:"name"`
	// Custom field value that is used to assign Collectors to the ingest budget.
	FieldValue string `json:"fieldValue"`
	// Capacity of the ingest budget, in bytes. It takes a few minutes for Collectors to stop collecting when capacity is reached. We recommend setting a soft limit that is lower than your needed hard limit.
	CapacityBytes int64 `json:"capacityBytes"`
	// Time zone of the reset time for the ingest budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone string `json:"timezone"`
	// Reset time of the ingest budget in HH:MM format.
	ResetTime string `json:"resetTime"`
	// Description of the ingest budget.
	Description string `json:"description,omitempty"`
	// Action to take when ingest budget's capacity is reached. All actions are audited. Supported values are:   * `stopCollecting`   * `keepCollecting`
	Action string `json:"action"`
	// The threshold as a percentage of when an ingest budget's capacity usage is logged in the Audit Index.
	AuditThreshold int32 `json:"auditThreshold,omitempty"`
	// Unique identifier for the ingest budget.
	Id string `json:"id"`
	// Current usage since the last reset, in bytes.
	UsageBytes int64 `json:"usageBytes,omitempty"`
	// Status of the current usage. Can be `Normal`, `Approaching`, `Exceeded`, or `Unknown` (unable to retrieve usage).
	UsageStatus string `json:"usageStatus,omitempty"`
	// Number of collectors assigned to the ingest budget.
	NumberOfCollectors int64 `json:"numberOfCollectors,omitempty"`
}

type IngestBudgetDefinition struct {
	// Display name of the ingest budget.
	Name string `json:"name"`
	// Custom field value that is used to assign Collectors to the ingest budget.
	FieldValue string `json:"fieldValue"`
	// Capacity of the ingest budget, in bytes. It takes a few minutes for Collectors to stop collecting when capacity is reached. We recommend setting a soft limit that is lower than your needed hard limit.
	CapacityBytes int64 `json:"capacityBytes"`
	// Time zone of the reset time for the ingest budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone string `json:"timezone"`
	// Reset time of the ingest budget in HH:MM format.
	ResetTime string `json:"resetTime"`
	// Description of the ingest budget.
	Description string `json:"description,omitempty"`
	// Action to take when ingest budget's capacity is reached. All actions are audited. Supported values are:   * `stopCollecting`   * `keepCollecting`
	Action string `json:"action"`
	// The threshold as a percentage of when an ingest budget's capacity usage is logged in the Audit Index.
	AuditThreshold int32 `json:"auditThreshold,omitempty"`
}

type ListIngestBudgetsResponse struct {
	// List of ingest budgets.
	Data []IngestBudget `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type ListIngestBudgetV1Opts struct {
	Limit optional.Int32
	Token optional.String
}
