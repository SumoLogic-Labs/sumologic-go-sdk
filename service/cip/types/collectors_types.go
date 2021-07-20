package types

type CollectorModel struct {
	// Alive collector is online or offline
	Alive bool `json:"alive"`
	// Category overrides the categories on sources
	Category string `json:"category,omitempty"`
	// CollectorType hosted or installed
	CollectorType string `json:"collectorType"`
	// CollectorVersion of the collector
	CollectorVersion string `json:"collectorVersion"`
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
