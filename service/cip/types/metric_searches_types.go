package types

import "time"

type MetricsSearchInstance struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Item title in the content library.
	Title string `json:"title"`
	// Item description in the content library.
	Description string               `json:"description"`
	TimeRange   *ResolvableTimeRange `json:"timeRange"`
	// Log query used to add an overlay to the chart.
	LogQuery string `json:"logQuery,omitempty"`
	// Metrics queries, up to the maximum of six.
	MetricsQueries []MetricsSearchQuery `json:"metricsQueries"`
	// Desired quantization in seconds.
	DesiredQuantizationInSecs int32 `json:"desiredQuantizationInSecs,omitempty"`
	// Chart properties, like line width, color palette, and the fill missing data method. Leave this field empty to use the defaults. This property contains JSON object encoded as a string.
	Properties string `json:"properties,omitempty"`
	// Identifier of the metrics search.
	Id string `json:"id"`
	// Identifier of the parent element in the content library, such as folder.
	ParentId string `json:"parentId,omitempty"`
}

type MetricsSearchQuery struct {
	// Row identifier. All row IDs are represented by subsequent upper case letters starting with `A`.
	RowId string `json:"rowId"`
	// Metrics query.
	Query string `json:"query"`
}

type SaveMetricsSearchRequest struct {
	// Item title in the content library.
	Title string `json:"title"`
	// Item description in the content library.
	Description string               `json:"description"`
	TimeRange   *ResolvableTimeRange `json:"timeRange"`
	// Log query used to add an overlay to the chart.
	LogQuery string `json:"logQuery,omitempty"`
	// Metrics queries, up to the maximum of six.
	MetricsQueries []MetricsSearchQuery `json:"metricsQueries"`
	// Desired quantization in seconds.
	DesiredQuantizationInSecs int32 `json:"desiredQuantizationInSecs,omitempty"`
	// Chart properties, like line width, color palette, and the fill missing data method. Leave this field empty to use the defaults. This property contains JSON object encoded as a string.
	Properties string `json:"properties,omitempty"`
	// Identifier of a folder to which the metrics search should be added.
	ParentId string `json:"parentId"`
}

type MetricsSearchV1 struct {
	// Item title in the content library.
	Title string `json:"title"`
	// Item description in the content library.
	Description string               `json:"description"`
	TimeRange   *ResolvableTimeRange `json:"timeRange"`
	// Log query used to add an overlay to the chart.
	LogQuery string `json:"logQuery,omitempty"`
	// Metrics queries, up to the maximum of six.
	MetricsQueries []MetricsSearchQuery `json:"metricsQueries"`
	// Desired quantization in seconds.
	DesiredQuantizationInSecs int32 `json:"desiredQuantizationInSecs,omitempty"`
	// Chart properties, like line width, color palette, and the fill missing data method. Leave this field empty to use the defaults. This property contains JSON object encoded as a string.
	Properties string `json:"properties,omitempty"`
}
