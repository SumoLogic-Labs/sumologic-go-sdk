package types

import (
	"github.com/antihax/optional"
	"time"
)

type CreateRoleDefinition struct {
	// Name of the role.
	Name string `json:"name"`
	// Description of the role.
	Description string `json:"description,omitempty"`
	// A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using '!_sourceCategory=billing' as a filter predicate will prevent users assigned to the role from viewing logs from the source category named 'billing'.
	FilterPredicate string `json:"filterPredicate,omitempty"`
	// List of user identifiers to assign the role to.
	Users []string `json:"users,omitempty"`
	// List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - dataVolumeIndex   - viewConnections   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - manageDataStreams  ### Entity management   - manageEntityTypeConfig  ### Metrics   - manageMonitors   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Monitors (New)   - viewMonitorsV2   - manageMonitorsV2
	Capabilities []string `json:"capabilities,omitempty"`
	// Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies.
	AutofillDependencies bool `json:"autofillDependencies,omitempty"`
}

type ListRoleModelsResponse struct {
	// List of roles.
	Data []RoleModel `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type ListRolesOpts struct {
	Limit  optional.Int32
	Token  optional.String
	SortBy optional.String
	Name   optional.String
}

type RoleModel struct {
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Name of the role.
	Name string `json:"name"`
	// Description of the role.
	Description string `json:"description,omitempty"`
	// A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using '!_sourceCategory=billing' as a filter predicate will prevent users assigned to the role from viewing logs from the source category named 'billing'.
	FilterPredicate string `json:"filterPredicate,omitempty"`
	// List of user identifiers to assign the role to.
	Users []string `json:"users,omitempty"`
	// List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - dataVolumeIndex   - viewConnections   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - manageDataStreams  ### Entity management   - manageEntityTypeConfig  ### Metrics   - manageMonitors   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Monitors (New)   - viewMonitorsV2   - manageMonitorsV2
	Capabilities []string `json:"capabilities,omitempty"`
	// Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies.
	AutofillDependencies bool `json:"autofillDependencies,omitempty"`
	// Unique identifier for the role.
	Id string `json:"id"`
	// Role is system or user defined.
	SystemDefined bool `json:"systemDefined,omitempty"`
}

type UpdateRoleDefinition struct {
	// Name of the role.
	Name string `json:"name"`
	// Description of the role.
	Description string `json:"description"`
	// A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using '!_sourceCategory=billing' as a filter predicate will prevent users assigned to the role from viewing logs from the source category named 'billing'.
	FilterPredicate string `json:"filterPredicate"`
	// List of user identifiers to assign the role to.
	Users []string `json:"users"`
	// List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - dataVolumeIndex   - viewConnections   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - manageDataStreams  ### Entity management   - manageEntityTypeConfig  ### Metrics   - manageMonitors   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Monitors (New)   - viewMonitorsV2   - manageMonitorsV2
	Capabilities []string `json:"capabilities"`
	// Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies.
	AutofillDependencies bool `json:"autofillDependencies,omitempty"`
}
