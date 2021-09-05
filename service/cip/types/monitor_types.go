package types

import (
	"github.com/antihax/optional"
	"time"
)

type ContentCopyParams struct {
	// Identifier of the parent folder to copy to.
	ParentId string `json:"parentId"`
	// Optionally provide a new name.
	Name string `json:"name,omitempty"`
	// Optionally provide a new description.
	Description string `json:"description,omitempty"`
}

type MonitorsLibraryBase struct {
	// Name of the monitor or folder.
	Name string `json:"name"`
	// Description of the monitor or folder.
	Description string `json:"description,omitempty"`
	// Type of the object model. Valid values:   1) MonitorsLibraryMonitor   2) MonitorsLibraryFolder
	Type_ string `json:"type"`
}

type MonitorsLibraryBaseExport struct {
	// Name of the monitor or folder.
	Name string `json:"name"`
	// Description of the monitor or folder.
	Description string `json:"description,omitempty"`
	// Type of the object model.
	Type_ string `json:"type"`
}

type MonitorsLibraryBaseResponse struct {
	// Identifier of the monitor or folder.
	Id string `json:"id"`
	// Identifier of the monitor or folder.
	Name string `json:"name"`
	// Description of the monitor or folder.
	Description string `json:"description"`
	// Version of the monitor or folder.
	Version int64 `json:"version"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Identifier of the parent folder.
	ParentId string `json:"parentId"`
	// Type of the content. Valid values:   1) Monitor   2) Folder
	ContentType string `json:"contentType"`
	// Type of the object model.
	Type_ string `json:"type"`
	// System objects are objects provided by Sumo Logic. System objects can only be localized. Non-local fields can't be updated.
	IsSystem bool `json:"isSystem"`
	// Immutable objects are \"READ-ONLY\".
	IsMutable bool `json:"isMutable"`
	// Aggregated permission summary for the calling user. If detailed permission statements are required, please call list permissions endpoint.
	Permissions []string `json:"permissions,omitempty"`
}

type MonitorsLibraryBaseUpdate struct {
	// The name of the monitor or folder.
	Name string `json:"name"`
	// The description of the monitor or folder.
	Description string `json:"description,omitempty"`
	// The version of the monitor or folder.
	Version int64 `json:"version"`
	// Type of the object model.
	Type_ string `json:"type"`
}

type MonitorsLibraryFolderResponse struct {
	// Identifier of the monitor or folder.
	Id string `json:"id"`
	// Identifier of the monitor or folder.
	Name string `json:"name"`
	// Description of the monitor or folder.
	Description string `json:"description"`
	// Version of the monitor or folder.
	Version int64 `json:"version"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Identifier of the parent folder.
	ParentId string `json:"parentId"`
	// Type of the content. Valid values:   1) Monitor   2) Folder
	ContentType string `json:"contentType"`
	// Type of the object model.
	Type_ string `json:"type"`
	// System objects are objects provided by Sumo Logic. System objects can only be localized. Non-local fields can't be updated.
	IsSystem bool `json:"isSystem"`
	// Immutable objects are \"READ-ONLY\".
	IsMutable bool `json:"isMutable"`
	// Aggregated permission summary for the calling user. If detailed permission statements are required, please call list permissions endpoint.
	Permissions []string `json:"permissions"`
	// Children of the folder. NOTE: Permissions field will not be filled (empty list) for children.
	Children []MonitorsLibraryBaseResponse `json:"children"`
}

type MonitorsLibraryItemWithPath struct {
	Item *MonitorsLibraryBaseResponse `json:"item"`
	// Path of the monitor or folder.
	Path string `json:"path"`
}

type MonitorsSearchOpts struct {
	Limit  optional.Int32
	Offset optional.Int32
}

type MonitorUsage struct {
	// The type of monitor usage info (Logs or Metrics).
	MonitorType string `json:"monitorType,omitempty"`
	// Current number of active Logs/Metrics monitors.
	Usage int32 `json:"usage,omitempty"`
	// The limit of active Logs/Metrics monitors.
	Limit int32 `json:"limit,omitempty"`
	// The total number of monitors created. (Including both active and disabled Logs/Metrics monitors)
	Total int32 `json:"total,omitempty"`
}

type MonitorQueries struct {
	// The type of monitor. Valid values:   1. `Logs`: A logs query monitor.   2. `Metrics`: A metrics query monitor.
	MonitorType string `json:"monitorType"`
	// The relative time range of the monitor.
	TimeRange string `json:"timeRange"`
	// Queries to be validated.
	Queries []UnvalidatedMonitorQuery `json:"queries"`
}

type MonitorQueriesValidationResult struct {
	// Whether or not if queries are valid.
	IsValid bool `json:"isValid,omitempty"`
	// Message from validation.
	Message string `json:"message,omitempty"`
}

type Path struct {
	// Elements of the path.
	PathItems []PathItem `json:"pathItems"`
	// String representation of the path.
	Path string `json:"path"`
}

type PathItem struct {
	// Identifier of the path element.
	Id string `json:"id"`
	// Name of the path element.
	Name string `json:"name"`
}

type UnvalidatedMonitorQuery struct {
	// The unique identifier of the row. Defaults to sequential capital letters, `A`, `B`, `C`, etc.
	RowId string `json:"rowId"`
	// The logs or metrics query that defines the stream of data the monitor runs on.
	Query string `json:"query"`
}
