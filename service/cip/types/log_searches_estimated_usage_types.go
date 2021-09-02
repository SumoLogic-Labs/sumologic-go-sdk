package types

type AutoCompleteValueSyncDefinition struct {
	// The label of the autocomplete value.
	Label string `json:"label"`
	// The value of the autocomplete value.
	Value string `json:"value"`
}

type EstimatedUsageDetails struct {
	// Amount of data scanned in bytes, to run the query.
	DataScannedInBytes int64 `json:"dataScannedInBytes,omitempty"`
}

type EstimatedUsageDetailsWithTier struct {
	// Name of the data tier. Supported Values are Continuous, Frequent, Infrequent
	Tier string `json:"tier,omitempty"`
	// Amount of data scanned in bytes, to run the query.
	DataScannedInBytes int64 `json:"dataScannedInBytes,omitempty"`
}

type LogSearchEstimatedUsageByTierDefinition struct {
	// Time zone to get the estimated usage details. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone              string                          `json:"timezone"`
	EstimatedUsageDetails []EstimatedUsageDetailsWithTier `json:"estimatedUsageDetails"`
}

type LogSearchEstimatedUsageDefinition struct {
	// Time zone to get the estimated usage details. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone              string                 `json:"timezone"`
	EstimatedUsageDetails *EstimatedUsageDetails `json:"estimatedUsageDetails"`
}

type LogSearchEstimatedUsageRequest struct {
	// Define the parsing mode to scan the JSON format log messages. Possible values are:   1. `AutoParse`   2. `Manual` In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid=0011).
	ParsingMode string `json:"parsingMode,omitempty"`
	// Time zone to get the estimated usage details. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone string `json:"timezone"`
}

type LogSearchEstimatedUsageRequestV2 struct {
	// Query to perform.
	QueryString string               `json:"queryString"`
	TimeRange   *ResolvableTimeRange `json:"timeRange"`
	// This has the value `true` if the search is to be run by receipt time and `false` if it is to be run by message time.
	RunByReceiptTime bool `json:"runByReceiptTime,omitempty"`
	// Definition of the query parameters.
	QueryParameters []QueryParameterSyncDefinition `json:"queryParameters,omitempty"`
	// Time zone to get the estimated usage details. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone string `json:"timezone"`
}

type ParameterAutoCompleteSyncDefinition struct {
	// The autocomplete parameter type. Supported values are:   1. `SKIP_AUTOCOMPLETE`   2. `CSV_AUTOCOMPLETE`   3. `AUTOCOMPLETE_KEY`   4. `VALUE_ONLY_AUTOCOMPLETE`   5. `VALUE_ONLY_LOOKUP_AUTOCOMPLETE`   6. `LABEL_VALUE_LOOKUP_AUTOCOMPLETE`
	AutoCompleteType string `json:"autoCompleteType"`
	// The autocomplete key to be used to fetch autocomplete values.
	AutoCompleteKey string `json:"autoCompleteKey,omitempty"`
	// The array of values of the corresponding autocomplete parameter.
	AutoCompleteValues []AutoCompleteValueSyncDefinition `json:"autoCompleteValues,omitempty"`
	// The lookup file to use as a source for autocomplete values.
	LookupFileName string `json:"lookupFileName,omitempty"`
	// The column from the lookup file to use for autocomplete labels.
	LookupLabelColumn string `json:"lookupLabelColumn,omitempty"`
	// The column from the lookup file to fill the actual value when a particular label is selected.
	LookupValueColumn string `json:"lookupValueColumn,omitempty"`
}

type QueryParameterSyncDefinition struct {
	// The name of the parameter.
	Name string `json:"name"`
	// The label of the parameter.
	Label string `json:"label"`
	// A description of the parameter.
	Description string `json:"description"`
	// The data type of the parameter. Supported values are:   1. `NUMBER`   2. `STRING`   3. `QUERY_FRAGMENT`   4. `SEARCH_KEYWORD`
	DataType string `json:"dataType"`
	// A value for the parameter. Should be compatible with the type set in dataType field.
	Value        string                               `json:"value"`
	AutoComplete *ParameterAutoCompleteSyncDefinition `json:"autoComplete"`
}
