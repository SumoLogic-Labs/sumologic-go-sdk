package types

import (
	"github.com/antihax/optional"
	"time"
)

type Content struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Identifier of the content item.
	Id string `json:"id"`
	// The name of the content item.
	Name string `json:"name"`
	// Type of the content item. Supported values are:   1. Folder   2. Search   3. Report (for old dashboards)   4. Dashboard (for new dashboards)   5. Lookups
	ItemType string `json:"itemType"`
	// Identifier of the parent content item.
	ParentId string `json:"parentId"`
	// List of permissions the user has on the content item.
	Permissions []string `json:"permissions"`
}

type ContentList struct {
	// A list of the content items.
	Data []Content `json:"data"`
}

type ContentPath struct {
	// Path of the content item.
	Path string `json:"path"`
}

type ContentSyncDefinition struct {
	// The item type. Dashboard links are not supported.
	Type_ string `json:"type"`
	// The name of the item.
	Name string `json:"name"`
}

type ContentManagementApiAsyncCopyStatusOpts struct {
	IsAdminMode optional.String
}

type ContentManagementApiBeginAsyncCopyOpts struct {
	IsAdminMode optional.String
}

type ContentManagementApiBeginAsyncDeleteOpts struct {
	IsAdminMode optional.String
}

type ContentManagementApiBeginAsyncExportOpts struct {
	IsAdminMode optional.String
}

type ContentManagementApiBeginAsyncImportOpts struct {
	IsAdminMode optional.String
	Overwrite   optional.Bool
}

type ContentManagementApiGetAsyncDeleteStatusOpts struct {
	IsAdminMode optional.String
}

type ContentManagementApiGetAsyncExportResultOpts struct {
	IsAdminMode optional.String
}

type ContentManagementApiGetAsyncExportStatusOpts struct {
	IsAdminMode optional.String
}

type ContentManagementApiGetAsyncImportStatusOpts struct {
	IsAdminMode optional.String
}

type ContentManagementApiMoveItemOpts struct {
	IsAdminMode optional.String
}

type DashboardSyncDefinition struct {
	Type        string                      `json:"type"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	DetailLevel int                         `json:"detailLevel"`
	Properties  string                      `json:"properties"`
	Panels      []reportPanelSyncDefinition `json:"panels"`
	Filters     []filtersSyncDefinition     `json:"filters"`
}

type FolderSyncDefinition struct {
	Type        string                  `json:"type"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Children    []contentSyncDefinition `json:"children"`
}

type LookupTableSyncDefinition struct {
	Type            string   `json:"type"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Fields          []fields `json:"fields"`
	PrimaryKeys     []string `json:"primaryKeys"`
	TTL             int      `json:"ttl"`
	SizeLimitAction string   `json:"sizeLimitAction"`
}

type MetricsSearchSyncDefinition struct {
	Type           string              `json:"type"`
	Name           string              `json:"name"`
	TimeRange      timeRangeDefinition `json:"timeRange"`
	Description    string              `json:"description"`
	Queries        []queries           `json:"queries"`
	VisualSettings string              `json:"visualSettings"`
}

type MetricsSavedSearchSyncDefinition struct {
	Type                      string                     `json:"type"`
	Name                      string                     `json:"name"`
	Description               string                     `json:"description"`
	TimeRange                 timeRangeDefinition        `json:"timeRange"`
	LogQuery                  string                     `json:"logQuery"`
	MetricsQueries            []metricsQueriesDefinition `json:"metricsQueries"`
	DesiredQuantizationInSecs int                        `json:"desiredQuantizationInSecs"`
	Properties                string                     `json:"properties"`
}

type MewboardSyncDefinition struct {
	Type             string                  `json:"type"`
	Name             string                  `json:"name"`
	Description      string                  `json:"description"`
	Title            string                  `json:"title"`
	RootPanel        rootPanelDefinition     `json:"rootPanel"`
	Theme            string                  `json:"theme"`
	TopologyLabelMap topologyLabelMap        `json:"topologyLabelMap"`
	RefreshInterval  int                     `json:"refreshInterval"`
	TimeRange        timeRangeDefinition     `json:"timeRange"`
	Layout           layout                  `json:"layout"`
	Panels           panelsDefinition        `json:"panels"`
	Variables        variablesDefinition     `json:"variables"`
	ColoringRules    coloringRulesDefinition `json:"coloringRules"`
}

type SavedSearchWithScheduleSyncDefinition struct {
	Type           string         `json:"type"`
	Name           string         `json:"name"`
	Search         search         `json:"search"`
	SearchSchedule searchSchedule `json:"searchSchedule"`
	Description    string         `json:"description"`
}

type autoComplete struct {
	AutoCompleteType   string               `json:"autoCompleteType"`
	AutoCompleteKey    string               `json:"autoCompleteKey"`
	AutoCompleteValues []autoCompleteValues `json:"autoCompleteValues"`
	LookupFileName     string               `json:"lookupFileName"`
	LookupLabelColumn  string               `json:"lookupLabelColumn"`
	LookupValueColumn  string               `json:"lookupValueColumn"`
}

type autoCompleteValues struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type autoParsingInfo struct {
	Mode string `json:"mode"`
}

type coloringRulesDefinition struct {
	Scope                           string          `json:"scope"`
	SingleSeriesAggregateFunction   string          `json:"singleSeriesAggregateFunction"`
	MultipleSeriesAggregateFunction string          `json:"multipleSeriesAggregateFunction"`
	ColorThresholds                 colorThresholds `json:"colorThresholds"`
}

type colorThresholds struct {
	Color string `json:"color"`
	Min   int    `json:"min"`
	Max   int    `json:"max"`
}

type contentSyncDefinition struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type fields struct {
	FieldName string `json:"fieldName"`
	FieldType string `json:"fieldType"`
}

type filtersSyncDefinition struct {
	FieldName    string   `json:"fieldName"`
	Label        string   `json:"label"`
	DefaultValue string   `json:"defaultValue"`
	FilterType   string   `json:"filterType"`
	Properties   string   `json:"properties"`
	PanelIds     []string `json:"panelIds"`
}

type layout struct {
	LayoutType       string            `json:"layoutType"`
	LayoutStructures []layoutStructure `json:"layoutStructures"`
}

type layoutStructure struct {
	Key       string `json:"key"`
	Structure string `json:"structure"`
}

type metricsQueriesDefinition struct {
	Query string `json:"query"`
	RowId string `json:"rowId"`
}

type panelsDefinition struct {
	Id                                     string `json:"id"`
	Key                                    string `json:"key"`
	Title                                  string `json:"title"`
	visualSettings                         string `json:"visualSettings"`
	KeepVisualSettingsConsistentWithParent bool   `json:"keepVisualSettingsConsistentWithParent"`
	PanelType                              string `json:"panelType"`
}

type queries struct {
	QueryString string `json:"queryString"`
	QueryType   string `json:"queryType"`
	QueryKey    string `json:"queryKey"`
}

type queryParameters struct {
	Name         string       `json:"name"`
	Label        string       `json:"label"`
	Description  string       `json:"description"`
	DataType     string       `json:"dataType"`
	Value        string       `json:"value"`
	AutoComplete autoComplete `json:"autoComplete"`
}

type reportPanelSyncDefinition struct {
	Name                      string                     `json:"name"`
	ViewerType                string                     `json:"viewerType"`
	DetailLevel               int                        `json:"detailLevel"`
	QueryString               string                     `json:"queryString"`
	MetricsQueries            []metricsQueriesDefinition `json:"metricsQueries"`
	TimeRange                 timeRangeDefinition        `json:"timeRange"`
	X                         int                        `json:"x"`
	Y                         int                        `json:"y"`
	Width                     int                        `json:"width"`
	Height                    int                        `json:"height"`
	Properties                string                     `json:"properties"`
	Id                        string                     `json:"id"`
	DesiredQuantizationInSecs int                        `json:"desiredQuantizationInSecs"`
	QueryParameters           []queryParameters          `json:"queryParameters"`
	AutoParsingInfo           autoParsingInfo            `json:"autoParsingInfo"`
}

type rootPanelDefinition struct {
	Id                                     string                    `json:"id"`
	Key                                    string                    `json:"key"`
	Title                                  string                    `json:"title"`
	VisualSettings                         string                    `json:"visualSettings"`
	KeepVisualSettingsConsistentWithParent bool                      `json:"keepVisualSettingsConsistentWithParent"`
	PanelType                              string                    `json:"panelType"`
	Layout                                 layout                    `json:"layout"`
	Panels                                 []panelsDefinition        `json:"panels"`
	Variables                              []variablesDefinition     `json:"variables"`
	ColoringRules                          []coloringRulesDefinition `json:"coloringRules"`
}

type search struct {
	QueryText        string            `json:"queryText"`
	DefaultTimeRange string            `json:"defaultTimeRange"`
	ByReceiptTime    bool              `json:"byReceiptTime"`
	ViewName         string            `json:"viewName"`
	ViewStartTime    string            `json:"viewStartTime"`
	QueryParameters  []queryParameters `json:"queryParameters"`
	ParsingMode      string            `json:"parsingMode"`
}

type searchSchedule struct {
	CronExpression       string                     `json:"cronExpression"`
	DisplayableTimeRange string                     `json:"displayableTimeRange"`
	ParseableTimeRange   timeRangeDefinition        `json:"parseableTimeRange"`
	TimeZone             string                     `json:"timeZone"`
	Threshold            searchScheduleThreshold    `json:"threshold"`
	Notification         searchScheduleNotification `json:"notification"`
	ScheduleType         string                     `json:"scheduleType"`
	MuteErrorEmails      bool                       `json:"muteErrorEmails"`
	Parameters           []searchScheduleParamters  `json:"parameters"`
}

type searchScheduleNotification struct {
	TaskType string `json:"taskType"`
}

type searchScheduleParamters struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type searchScheduleThreshold struct {
	ThresholdType string `json:"thresholdType"`
	Operator      string `json:"operator"`
	Count         int    `json:"count"`
}

type timeRangeDefinition struct {
	Type string                  `json:"type"`
	From timeRangeFromDefinition `json:"from"`
}

type timeRangeFromDefinition struct {
	Type         string `json:"type"`
	RelativeTime string `json:"relativeTime"`
}

type topologyLabelMap struct {
	Service []string `json:"service"`
}

type variablesDefinition struct {
	Id               string                    `json:"id"`
	Name             string                    `json:"name"`
	DisplayName      string                    `json:"displayName"`
	DefaultValue     string                    `json:"defaultValue"`
	SourceDefinition variablesSourceDefinition `json:"sourceDefinition"`
	AllowMultiSelect bool                      `json:"allowMultiSelect"`
	IncludeAllOption bool                      `json:"includeAllOption"`
	HideFromUI       bool                      `json:"hideFromUI"`
}

type variablesSourceDefinition struct {
	VariableSourceType string `json:"variableSourceType"`
}
