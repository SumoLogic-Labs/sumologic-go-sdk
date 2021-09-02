package types

import "github.com/antihax/optional"

type ContentPermissionAssignment struct {
	// Content permission name. Valid values are: `View`, `GrantView`, `Edit`, `GrantEdit`, `Manage`, and `GrantManage`.
	PermissionName string `json:"permissionName"`
	// Type of source for the permission. Valid values are: `user`, `role`, and `org`.
	SourceType string `json:"sourceType"`
	// An identifier that belongs to the source type chosen above. For e.g. if the sourceType is set to \"user\", sourceId should be identifier of a user (same goes for `role` and `org` sourceType)
	SourceId string `json:"sourceId"`
	// Unique identifier for the content item.
	ContentId string `json:"contentId"`
}

type ContentPermissionsOpts struct {
	IsAdminMode optional.String
}

type ContentPermissionResult struct {
	// Explicitly assigned content permissions.
	ExplicitPermissions []ContentPermissionAssignment `json:"explicitPermissions"`
	// Implicitly inherited content permissions.
	ImplicitPermissions []ContentPermissionAssignment `json:"implicitPermissions,omitempty"`
}

type ContentPermissionUpdateRequest struct {
	// Content permissions to be updated.
	ContentPermissionAssignments []ContentPermissionAssignment `json:"contentPermissionAssignments"`
	// Set this to \"true\" to notify the users who had a permission update.
	NotifyRecipients bool `json:"notifyRecipients"`
	// The notification message sent to the users who had a permission update.
	NotificationMessage string `json:"notificationMessage"`
}

type GetContentPermissionsOpts struct {
	ExplicitOnly optional.Bool
	IsAdminMode  optional.String
}
