package types

import (
	"github.com/antihax/optional"
	"time"
)

type CreateScheduledViewDefinition struct {
	// The query that defines the data to be included in the scheduled view.
	Query string `json:"query"`
	// Name of the index for the scheduled view.
	IndexName string `json:"indexName"`
	// Start timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	StartTime time.Time `json:"startTime"`
	// The number of days to retain data in the scheduled view, or -1 to use the default value for your account.  Only relevant if your account has multi-retention. enabled.
	RetentionPeriod int32 `json:"retentionPeriod,omitempty"`
	// An optional ID of a data forwarding configuration to be used by the scheduled view.
	DataForwardingId string `json:"dataForwardingId,omitempty"`
	// Define the parsing mode to scan the JSON format log messages. Possible values are:   1. `AutoParse`   2. `Manual` In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid=0011).
	ParsingMode string `json:"parsingMode,omitempty"`
}

type ListScheduledViewsResponse struct {
	// List of scheduled views.
	Data []ScheduledView `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type ScheduledView struct {
	// The query that defines the data to be included in the scheduled view.
	Query string `json:"query"`
	// Name of the index for the scheduled view.
	IndexName string `json:"indexName"`
	// Start timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	StartTime time.Time `json:"startTime"`
	// The number of days to retain data in the scheduled view, or -1 to use the default value for your account.  Only relevant if your account has multi-retention. enabled.
	RetentionPeriod int32 `json:"retentionPeriod,omitempty"`
	// An optional ID of a data forwarding configuration to be used by the scheduled view.
	DataForwardingId string `json:"dataForwardingId,omitempty"`
	// Define the parsing mode to scan the JSON format log messages. Possible values are:   1. `AutoParse`   2. `Manual` In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid=0011).
	ParsingMode string `json:"parsingMode,omitempty"`
	// Identifier for the scheduled view.
	Id string `json:"id"`
	// Creation timestamp in UTC.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// If the scheduled view is created by OptimizeIt.
	CreatedByOptimizeIt bool `json:"createdByOptimizeIt,omitempty"`
	// Errors related to the scheduled view.
	Error_ string `json:"error,omitempty"`
	// Status of the scheduled view.
	Status string `json:"status,omitempty"`
	// Total storage consumed by the scheduled view.
	TotalBytes int64 `json:"totalBytes,omitempty"`
	// Total number of messages for the scheduled view.
	TotalMessageCount int64 `json:"totalMessageCount,omitempty"`
	// Identifier of the user who created the scheduled view.
	CreatedBy string `json:"createdBy,omitempty"`
}

type ScheduledViewsOpts struct {
	Limit optional.Int32
	Token optional.String
}

type UpdateScheduledViewDefinition struct {
	// An optional ID of a data forwarding configuration to be used by the scheduled view.
	DataForwardingId string `json:"dataForwardingId,omitempty"`
	// The number of days to retain data in the scheduled view, or -1 to use the default value for your account.  Only relevant if your account has multi-retention. enabled.
	RetentionPeriod int32 `json:"retentionPeriod,omitempty"`
	// This is required if the newly specified `retentionPeriod` is less than the existing retention period.  In such a situation, a value of `true` says that data between the existing retention period and the new retention period should be deleted immediately; if `false`, such data will be deleted after seven days. This property is optional and ignored if the specified `retentionPeriod` is greater than or equal to the current retention period.
	ReduceRetentionPeriodImmediately bool `json:"reduceRetentionPeriodImmediately,omitempty"`
}
