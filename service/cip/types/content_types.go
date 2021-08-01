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
