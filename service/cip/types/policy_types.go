package types

type AuditPolicy struct {
	// Whether the Audit policy is enabled.
	Enabled bool `json:"enabled"`
}

type DataAccessLevelPolicy struct {
	// Whether the Data Access Level policy is enabled.
	Enabled bool `json:"enabled"`
}

type MaxUserSessionTimeoutPolicy struct {
	// Maximum web session timeout users are able to configure within their user preferences. Valid values are: `5m`, `15m`, `30m`, `1h`, `2h`, `6h`, `12h`, `1d`, `2d`, `3d`, `5d`, or `7d`
	MaxUserSessionTimeout string `json:"maxUserSessionTimeout"`
}

type SearchAuditPolicy struct {
	// Whether the Search Audit policy is enabled.
	Enabled bool `json:"enabled"`
}

type ShareDashboardsOutsideOrganizationPolicy struct {
	// Whether the Share Dashboards Outside Organization policy is enabled.
	Enabled bool `json:"enabled"`
}

type UserConcurrentSessionsLimitPolicy struct {
	// Whether the User Concurrent Sessions Limit policy is enabled.
	Enabled bool `json:"enabled"`
	// Maximum number of concurrent sessions a user may have.
	MaxConcurrentSessions int32 `json:"maxConcurrentSessions,omitempty"`
}
