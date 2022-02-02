package types

type CreateHttpSourceRequest struct {
	// ApiVersion version of the api
	ApiVersion string `json:"api.version"`
	// Source type of source
	Source HttpSource `json:"source"`
}

type HttpSource struct {
	// SourceType of the source
	SourceType string `json:"sourceType"`
	// Name of the source
	Name string `json:"name"`
	// Category of the source
	Category string `json:"category"`
	// Fields of the source
	Fields map[string]string `json:"fields"`
	// MessagePerRequest enables one message per request
	MessagePerRequest bool `json:"messagePerRequest"`
	// MultilineProcessingEnabled processes multiple line messages
	MultilineProcessingEnabled bool `json:"multilineProcessingEnabled"`
}

type HttpSourceDefinition struct {
	Id                         int               `json:"id"`
	Name                       string            `json:"name"`
	HostName                   string            `json:"hostName,omitempty"`
	AutomaticDateParsing       bool              `json:"automaticDateParsing"`
	MultilineProcessingEnabled bool              `json:"multilineProcessingEnabled"`
	UseAutolineMatching        bool              `json:"useAutolineMatching"`
	ForceTimezone              bool              `json:"forceTimezome"`
	Filters                    []SourceFilters   `json:"filters"`
	CutoffTimestamp            int               `json:"cutoffTimestamp"`
	Encoding                   string            `json:"encoding"`
	Interval                   int               `json:"interval"`
	Metrics                    []string          `json:"metrics"`
	SourceType                 string            `json:"sourceType"`
	Alive                      bool              `json:"alive"`
	Url                        string            `json:"url,omitempty"`
	Category                   string            `json:"category,omitempty"`
	Fields                     map[string]string `json:"fields,omitempty"`
	MessagePerRequest          bool              `json:"messagePerRequest"`
}

type HttpSourceModel struct {
	// Source http source definition
	Source HttpSourceDefinition `json:"source"`
}

type UpdateHttpSourceRequest struct {
	// Source http source definition
	Source HttpSourceDefinition `json:"source"`
}
