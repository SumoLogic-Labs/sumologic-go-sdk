package types

import "github.com/antihax/optional"

type availableBuildsDefinition struct {
	Version string `json:"version"`
	Latest  bool   `json:"latest"`
}

type DeleteOfflineCollectorsOpts struct {
	AliveBeforeDays optional.Int32
}

type GetUpgradableCollectorsOpts struct {
	Limit     optional.Int32  `json:"limit"`
	Offset    optional.Int32  `json:"offset"`
	ToVersion optional.String `json:"toVersion"`
}

type GetAvailableBuildsModel struct {
	Targets []availableBuildsDefinition `json:"targets"`
}

type ListCollectorsOpts struct {
	Filter optional.String
	Limit  optional.Int32
	Offset optional.Int32
}

type ListCollectorsOfflineOpts struct {
	AliveBeforeDays optional.Int32
	Limit           optional.Int32
	Offset          optional.Int32
}

type CollectorModel struct {
	// Alive collector is online or offline
	Alive bool `json:"alive"`
	// Category overrides the categories on sources
	Category string `json:"category,omitempty"`
	// CollectorType hosted or installed
	CollectorType string `json:"collectorType"`
	// CollectorVersion of the collector
	CollectorVersion string `json:"collectorVersion,omitempty"`
	// CutoffRelativeTime provides a relative offset from the current time
	CutoffRelativeTime string `json:"cutoffRelativeTime,omitempty"`
	// CutoffTimestamp only collect data from files with a modified date more recent than this timestamp, specified as milliseconds since epoch
	CutoffTimestamp int `json:"cutoffTimestamp,omitempty"`
	// Description for the collector
	Description string `json:"description,omitempty"`
	// Ephemeral when true, the collector will be deleted after 12 hours of inactivity
	Ephemeral bool `json:"ephemeral"`
	// Fields JSON map of key value pairs
	Fields map[string]string `json:"fields,omitempty"`
	//Links sources applied to collector
	Links []collectorLinks `json:"links,omitempty"`
	// HostName of the collector
	HostName string `json:"hostName,omitempty"`
	// Id of the collector
	Id string `json:"id"`
	// LastSeenAlive the last time the Sumo Logic service received an active heartbeat from the Collector, specified as milliseconds since epoch
	LastSeenAlive int `json:"lastSeenAlive,omitempty"`
	// Name of the collector
	Name string `json:"name"`
	// SourceSyncMode whether the collector is using local source or cloud management
	SourceSyncMode string `json:"sourceSyncMode,omitempty"`
	// TimeZone of the collector
	TimeZone string `json:"timeZone,omitempty"`
	// TargetCpu CPU limit for the collector
	TargetCpu int `json:"targetCpu,omitempty"`
	// OsName which the collector is installed on
	OsName string `json:"osName,omitempty"`
	// OsVersion which the collector is installed on
	OsVersion string `json:"osVersion,omitempty"`
	// OsArch which the collector is installed on
	OsArch string `json:"osArch,omitempty"`
	// OsTime of the machine the collector is installed on in milliseconds
	OsTime int `json:"osTime,omitempty"`
}

type collectorLinks struct {
	// Rel generally set to source
	Rel string `json:"rel"`
	// Href URL for source
	Href string `json:"href"`
}

type CreateCollectorRequestDefinition struct {
	// CollectorType Only "Hosted" collectors can be created via the API
	CollectorType string `json:"collectorType"`
	// Name of the collector
	Name string `json:"name"`
	// Description of the collector
	Description string `json:"description,omitempty"`
	// Category overrides the categories on sources
	Category string `json:"category,omitempty"`
	// Fields JSON map of key value pairs
	Fields map[string]string `json:"fields,omitempty"`
}

type CreateCollectorRequest struct {
	// Collector specify the configuration for the collector
	Collector CreateCollectorRequestDefinition `json:"collector"`
}

type ListCollectorsModel struct {
	// Collectors array of types.CollectorModel
	Collectors []CollectorModel `json:"collectors"`
}

type UpdateHostedCollectorDefinition struct {
	// Id of the collector
	Id string `json:"id"`
	// Name of the collector
	Name string `json:"name" json:"name"`
	// Description of the collector
	Description string `json:"description,omitempty"`
	// Category overrides the categories on sources
	Category string `json:"category,omitempty"`
	// CutOffTimestamp collects data from files with a modified date more recent than this timestamp (specified as milliseconds since epoch).
	CutOffTimestamp int32 `json:"cutoffTimestamp"`
	// Fields JSON map of key value pairs
	Fields map[string]string `json:"fields,omitempty"`
	// TimeZone of the collector
	TimeZone string `json:"timeZone"`
}

type UpdateInstalledCollectorDefinition struct {
	// Id of the collector
	Id string `json:"id"`
	// Name of the collector
	Name string `json:"name" json:"name"`
	// Description of the collector
	Description string `json:"description,omitempty"`
	// Category overrides the categories on sources
	Category string `json:"category,omitempty"`
	// CutOffTimestamp collects data from files with a modified date more recent than this timestamp (specified as milliseconds since epoch).
	CutOffTimestamp int32 `json:"cutoffTimestamp"`
	// Fields JSON map of key value pairs
	Fields map[string]string `json:"fields,omitempty"`
	// Ephemeral whether the collector will be removed after 12 hours of no activity
	Ephemeral bool `json:"ephemeral"`
	// HostName of the OS
	HostName string `json:"hostName"`
	// SourceSyncMode how the sources are managed (can be either "UI" or "JSON")
	SourceSyncMode string `json:"sourceSyncMode"`
	// TimeZone of the collector
	TimeZone string `json:"timeZone"`
	// TargetCPU CPU utilization threshold
	TargetCPU int32 `json:"targetCpu"`
}

type CreateUpgradeOrDowngradeRequest struct {
	CollectorId int    `json:"collectorId"`
	ToVersion   string `json:"toVersion"`
}

type upgradeOrDowngradeTaskLinkDefinition struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type UpgradeOrDowngradeTaskModel struct {
	Id   string                               `json:"id"`
	Link upgradeOrDowngradeTaskLinkDefinition `json:"link"`
}

type UpgradeOrDowngradeTaskStatusModel struct {
	Upgrade upgradeOrDowngradeTaskStatusDefinition `json:"upgrade"`
}

type upgradeOrDowngradeTaskStatusDefinition struct {
	Id          int    `json:"id"`
	CollectorId int    `json:"collectorId"`
	ToVersion   string `json:"toVersion"`
	RequestTime int    `json:"requestTime"`
	Status      int    `json:"status"`
	Message     string `json:"message"`
}
