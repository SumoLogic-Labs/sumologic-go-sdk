package types

type ColoringRule struct {
	// Regex string to match queries to apply coloring to.
	Scope string `json:"scope"`
	// Function to aggregate one series into one single value.
	SingleSeriesAggregateFunction string `json:"singleSeriesAggregateFunction"`
	// Function to aggregate the aggregate values of multiple time series into one single value.
	MultipleSeriesAggregateFunction string `json:"multipleSeriesAggregateFunction"`
	// Color thresholds.
	ColorThresholds []ColoringThreshold `json:"colorThresholds,omitempty"`
}

type ColoringThreshold struct {
	// Color for the threshold.
	Color string `json:"color"`
	// Absolute inclusive threshold to color by.
	Min float64 `json:"min,omitempty"`
	// Absolute exclusive threshold to color by.
	Max float64 `json:"max,omitempty"`
}

type Dashboard struct {
	// Title of the dashboard.
	Title string `json:"title"`
	// Description of the dashboard.
	Description string `json:"description,omitempty"`
	// The identifier of the folder to save the dashboard in. By default it is saved in your personal folder.
	FolderId         string           `json:"folderId,omitempty"`
	TopologyLabelMap TopologyLabelMap `json:"topologyLabelMap,omitempty"`
	// If set denotes that the dashboard concerns a given domain (e.g. `aws`, `k8s`, `app`).
	Domain string `json:"domain,omitempty"`
	// Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard. This functionality is currently not supported.
	RefreshInterval int32               `json:"refreshInterval,omitempty"`
	TimeRange       ResolvableTimeRange `json:"timeRange"`
	// Panels in the dashboard.
	Panels []Panel `json:"panels,omitempty"`
	Layout Layout  `json:"layout,omitempty"`
	// Variables to apply to the panels.
	Variables []Variable `json:"variables,omitempty"`
	// Theme for the dashboard. Either `Light` or `Dark`.
	Theme string `json:"theme,omitempty"`
	// Rules to set the color of data. This is an internal field and is not current supported by UI.
	ColoringRules []ColoringRule `json:"coloringRules,omitempty"`
	// Unique identifier for the dashboard.
	Id string `json:"id,omitempty"`
}

type DashboardRequest struct {
	// Name of the dashboard
	Name string `json:"name"`
	// Title of the dashboard.
	Title string `json:"title"`
	// Type of dashboard
	Type string `json:"type"`
	// Description of the dashboard.
	Description string `json:"description"`
	// The identifier of the folder to save the dashboard in. By default it is saved in your personal folder.
	FolderId         string           `json:"folderId,omitempty"`
	TopologyLabelMap TopologyLabelMap `json:"topologyLabelMap,omitempty"`
	// If set denotes that the dashboard concerns a given domain (e.g. `aws`, `k8s`, `app`).
	Domain string `json:"domain,omitempty"`
	// Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard. This functionality is currently not supported.
	RefreshInterval int32               `json:"refreshInterval"`
	TimeRange       ResolvableTimeRange `json:"timeRange"`
	// Panels in the dashboard.
	Panels []Panel `json:"panels,omitempty"`
	Layout *Layout `json:"layout,omitempty"`
	// Variables to apply to the panels.
	Variables []Variable `json:"variables,omitempty"`
	// Theme for the dashboard. Either `Light` or `Dark`.
	Theme string `json:"theme"`
	// Rules to set the color of data. This is an internal field and is not current supported by UI.
	ColoringRules []ColoringRule `json:"coloringRules,omitempty"`
}

type TimeRangeBoundary struct {
	Type         string `json:"type,omitempty"`
	RelativeTime string `json:"relativeTime,omitempty"`
	EpochMillis  int64  `json:"epochMillis,omitempty"`
	Iso8601Time  string `json:"iso8601Time,omitempty"`
	RangeName    string `json:"rangeName,omitempty"`
}

type Layout struct {
	// The type of panel layout on the Dashboard. For example, Grid, Tabs, or Hierarchical. Currently supports `Grid` only.
	LayoutType string `json:"layoutType"`
	// Layout structures for the panel childen.
	LayoutStructures []LayoutStructure `json:"layoutStructures"`
}

type LayoutStructure struct {
	// The identifier of the panel that this structure applies to.
	Key string `json:"key"`
	// The structure of a panel.
	Structure string `json:"structure"`
}

type Panel struct {
	// Unique identifier for the panel.
	Id string `json:"id,omitempty"`
	// Key for the panel. Used to create searches for the queries in the panel and configure the layout of the panel in the dashboard.
	Key string `json:"key"`
	// Title of the panel.
	Title string `json:"title,omitempty"`
	// Visual settings of the panel.
	VisualSettings string `json:"visualSettings,omitempty"`
	// Keeps the visual settings, like series colors, consistent with the settings of the parent panel.
	KeepVisualSettingsConsistentWithParent bool `json:"keepVisualSettingsConsistentWithParent,omitempty"`
	// Type of panel.
	PanelType string `json:"panelType"`
}

type ResolvableTimeRange struct {
	// Type of the time range. Value must be either `CompleteLiteralTimeRange` or `BeginBoundedTimeRange`.
	Type_ string            `json:"type"`
	From  TimeRangeBoundary `json:"from"`
	To    TimeRangeBoundary `json:"to,omitempty"`
}

type Variable struct {
	// Unique identifier for the variable.
	Id string `json:"id,omitempty"`
	// Name of the variable. The variable name is case-insensitive. Only alphanumeric, and underscores are allowed in the variable name.
	Name string `json:"name"`
	// Display name of the variable shown in the UI. If this field is empty, the name field will be used. The display name is case-insensitive. Only numbers, and underscores are allowed in the variable name. This field is not yet supported by the UI.
	DisplayName string `json:"displayName,omitempty"`
	// Default value of the variable.
	DefaultValue     string                   `json:"defaultValue,omitempty"`
	SourceDefinition VariableSourceDefinition `json:"sourceDefinition"`
	// Allow multiple selections in the values dropdown.
	AllowMultiSelect bool `json:"allowMultiSelect,omitempty"`
	// Include an \"All\" option at the top of the variable's values dropdown.
	IncludeAllOption bool `json:"includeAllOption,omitempty"`
	// Hide the variable in the dashboard UI.
	HideFromUI bool `json:"hideFromUI,omitempty"`
}

type VariableSourceDefinition struct {
	// Source type of the variable values.
	VariableSourceType string `json:"variableSourceType"`
}

type TopologyLabelMap struct {
	// Map from topology labels to `TopologyLabelValuesList`.
	Data map[string][]string `json:"data"`
}
