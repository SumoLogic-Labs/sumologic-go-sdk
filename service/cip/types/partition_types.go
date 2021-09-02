package types

import (
	"github.com/antihax/optional"
	"time"
)

type CreatePartitionDefinition struct {
	// The name of the partition.
	Name string `json:"name"`
	// The query that defines the data to be included in the partition.
	RoutingExpression string `json:"routingExpression"`
	// The Data Tier where the data in the partition will reside. Possible values are:               1. `continuous`               2. `frequent`               3. `infrequent` Note: The \"infrequent\" and \"frequent\" tiers are only to available Cloud Flex Credits Enterprise Suite accounts.
	AnalyticsTier string `json:"analyticsTier,omitempty"`
	// The number of days to retain data in the partition, or -1 to use the default value for your account.  Only relevant if your account has variable retention enabled.
	RetentionPeriod int32 `json:"retentionPeriod,omitempty"`
	// Whether the partition is compliant or not. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition is marked compliant. A partition once marked compliant, cannot be marked non-compliant later.
	IsCompliant bool `json:"isCompliant,omitempty"`
}

type ListPartitionsResponse struct {
	// List of partitions.
	Data []Partition `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type Partition struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// The name of the partition.
	Name string `json:"name"`
	// The query that defines the data to be included in the partition.
	RoutingExpression string `json:"routingExpression"`
	// The Data Tier where the data in the partition will reside. Possible values are:               1. `continuous`               2. `frequent`               3. `infrequent` Note: The \"infrequent\" and \"frequent\" tiers are only to available Cloud Flex Credits Enterprise Suite accounts.
	AnalyticsTier string `json:"analyticsTier,omitempty"`
	// The number of days to retain data in the partition, or -1 to use the default value for your account.  Only relevant if your account has variable retention enabled.
	RetentionPeriod int32 `json:"retentionPeriod,omitempty"`
	// Whether the partition is compliant or not. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition is marked compliant. A partition once marked compliant, cannot be marked non-compliant later.
	IsCompliant bool `json:"isCompliant,omitempty"`
	// Unique identifier for the partition.
	Id string `json:"id"`
	// Size of data in partition in bytes.
	TotalBytes int64 `json:"totalBytes"`
	// This has the value `true` if the partition is active and `false` if it has been decommissioned.
	IsActive bool `json:"isActive,omitempty"`
	// If the retentionPeriod is scheduled to be updated in the future (i.e., if retentionPeriod is previously reduced with value of reduceRetentionPeriodImmediately as false), this property gives the future value of retentionPeriod while retentionPeriod gives the current value. retentionPeriod will take up the value of newRetentionPeriod after the scheduled time.
	NewRetentionPeriod int32 `json:"newRetentionPeriod,omitempty"`
	// This has the value `DefaultIndex`, `AuditIndex`or `Partition` depending upon the type of partition.
	IndexType string `json:"indexType,omitempty"`
	// When the newRetentionPeriod will become effective in UTC format.
	RetentionEffectiveAt time.Time `json:"retentionEffectiveAt,omitempty"`
	// Id of the data forwarding configuration to be used by the partition.
	DataForwardingId string `json:"dataForwardingId,omitempty"`
}

type PartitionOpts struct {
	Limit     optional.Int32
	Token     optional.String
	ViewTypes optional.Interface
}

type UpdatePartitionDefinition struct {
	// The number of days to retain data in the partition, or -1 to use the default value for your account. Only relevant if your account has variable retention enabled.
	RetentionPeriod int32 `json:"retentionPeriod,omitempty"`
	// This is required if the newly specified `retentionPeriod` is less than the existing retention period.  In such a situation, a value of `true` says that data between the existing retention period and the new retention period should be deleted immediately; if `false`, such data will be deleted after seven days. This property is optional and ignored if the specified `retentionPeriod` is greater than or equal to the current retention period.
	ReduceRetentionPeriodImmediately bool `json:"reduceRetentionPeriodImmediately,omitempty"`
	// Whether to mark a partition as compliant. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition marked as compliant. A partition once marked compliant, cannot be marked non-compliant later.
	IsCompliant bool `json:"isCompliant,omitempty"`
	// The query that defines the data to be included in the partition.
	RoutingExpression string `json:"routingExpression,omitempty"`
}
