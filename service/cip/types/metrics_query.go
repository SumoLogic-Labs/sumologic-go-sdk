package types

type MetricDefinition struct {
	// Name of the metric returning the timeseries.
	Metric string `json:"metric,omitempty"`
	// Metric dimensions / metadata related to each timeseries.
	Dimensions map[string]string `json:"dimensions,omitempty"`
}

type MetricsQueryRequest struct {
	// A list of metrics queries.
	Queries   []MetricsQueryRow    `json:"queries"`
	TimeRange *ResolvableTimeRange `json:"timeRange"`
}

type MetricsQueryResponse struct {
	// A list of the time series returned by metric query.
	QueryResult []TimeSeriesRow `json:"queryResult,omitempty"`
	Errors      *ErrorResponse  `json:"errors"`
}

type MetricsQueryRow struct {
	// Row id for the query row, A to Z letter.
	RowId string `json:"rowId"`
	// A metric query consists of a metric, one or more filters and optionally, one or more [Metrics Operators](https://help.sumologic.com/?cid=10144). Strictly speaking, both filters and operators are optional.  Most of the [Metrics Operators](https://help.sumologic.com/?cid=10144) are allowed in the query string except `fillmissing`, `outlier`, `quantize` and `timeshift`.    * `fillmissing`: Not supported in API.   * `outlier`: Not supported in API.   * `quantize`: Only supported through `quantization` param.   * `timeshift`: Only supported through `timeshift` param.   In practice, your metric queries will almost always contain filters that narrow the scope of your query. For more information about the query language see [Metrics Queries](http://help.sumologic.com/?cid=1079).
	Query string `json:"query"`
	// Segregates time series data by time period. This allows you to create aggregated results in buckets of fixed intervals (for example, 5-minute intervals). The value is in milliseconds.
	Quantization int64 `json:"quantization,omitempty"`
	// We use the term rollup to refer to the aggregation function Sumo Logic uses when quantizing metrics. Can be `Avg`, `Sum`, `Min`, `Max`, `Count` or `None`.
	Rollup string `json:"rollup,omitempty"`
	// Shifts the time series from your metrics query by the specified amount of time. This can help when comparing a time series across multiple time periods. Specified as a signed duration in milliseconds.
	Timeshift int64 `json:"timeshift,omitempty"`
	// Determines if the row should be returned in the response. Can be used in conjunction with a join, if only the result of the join is needed, and not the intermediate rows.
	Transient bool `json:"transient,omitempty"`
}

type Points struct {
	// Array of timestamps of datapoints in milliseconds.
	Timestamps []int64 `json:"timestamps"`
	// Array of values of datapoints corresponding to timestamp array.
	Values []float64 `json:"values"`
}

type TimeSeries struct {
	MetricDefinition *MetricDefinition `json:"metricDefinition"`
	Points           *Points           `json:"points"`
}

type TimeSeriesList struct {
	// A list of timeseries returned by corresponding query.
	TimeSeries []TimeSeries `json:"timeSeries"`
	// Unit of the query.
	Unit string `json:"unit,omitempty"`
	// Time shift value if specified in request in human readable format.
	TimeShiftLabel string `json:"timeShiftLabel,omitempty"`
}

type TimeSeriesRow struct {
	// Row id for the query row as specified in the request.
	RowId          string          `json:"rowId"`
	TimeSeriesList *TimeSeriesList `json:"timeSeriesList"`
}
