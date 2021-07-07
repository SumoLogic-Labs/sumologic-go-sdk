package types

import "github.com/antihax/optional"

type Folder struct {
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
	// The description of the folder.
	Description string `json:"description,omitempty"`
	// A list of the content items.
	Children []Content `json:"children,omitempty"`
}

type FolderDefinition struct {
	// The name of the folder.
	Name string `json:"name"`
	// The description of the folder.
	Description string `json:"description,omitempty"`
	// The identifier of the parent folder.
	ParentId string `json:"parentId"`
}

type FolderManagementApiCreateFolderOpts struct {
	IsAdminMode optional.String
}

type FolderManagementApiGetAdminRecommendedFolderAsyncOpts struct {
	IsAdminMode optional.String
}

type FolderManagementApiGetFolderOpts struct {
	IsAdminMode optional.String
}

type FolderManagementApiGetGlobalFolderAsyncOpts struct {
	IsAdminMode optional.String
}

type FolderManagementApiUpdateFolderOpts struct {
	IsAdminMode optional.String
}

type UpdateFolderRequest struct {
	// The name of the folder.
	Name string `json:"name"`
	// The description of the folder.
	Description string `json:"description,omitempty"`
}
