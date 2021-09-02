package types

import (
	"github.com/antihax/optional"
	"time"
)

type IngestBudgetV2 struct {
	// Display name of the ingest budget.
	Name string `json:"name"`
	// A scope is a constraint that will be used to identify the messages on which budget needs to be applied. A scope is consists of key and value separated by =. The field must be enabled in the fields table. Value supports wildcard. e.g. _sourceCategory=*prod*payment*, cluster=kafka. If the scope is defined _sourceCategory=*nginx* in this budget will be applied on messages having fields _sourceCategory=prod/nginx, _sourceCategory=dev/nginx, or _sourceCategory=dev/nginx/error
	Scope string `json:"scope"`
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
	// The creation timestamp in UTC of the Ingest Budget.
	CreatedAt time.Time `json:"createdAt"`
	// The identifier of the user who created the Ingest Budget.
	CreatedBy string `json:"createdBy"`
	// The modified timestamp in UTC of the Ingest Budget.
	ModifiedAt time.Time `json:"modifiedAt"`
	// The identifier of the user who modified the Ingest Budget.
	ModifiedBy string `json:"modifiedBy"`
	// The version of the Ingest Budget
	BudgetVersion int32 `json:"budgetVersion,omitempty"`
}

type IngestBudgetDefinitionV2 struct {
	// Display name of the ingest budget.
	Name string `json:"name"`
	// A scope is a constraint that will be used to identify the messages on which budget needs to be applied. A scope is consists of key and value separated by =. The field must be enabled in the fields table. Value supports wildcard. e.g. _sourceCategory=*prod*payment*, cluster=kafka. If the scope is defined _sourceCategory=*nginx* in this budget will be applied on messages having fields _sourceCategory=prod/nginx, _sourceCategory=dev/nginx, or _sourceCategory=dev/nginx/error
	Scope string `json:"scope"`
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

type ListIngestBudgetV2Opts struct {
	Limit optional.Int32
	Token optional.String
}

type ListIngestBudgetsResponseV2 struct {
	// List of ingest budgets.
	Data []IngestBudgetV2 `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}
