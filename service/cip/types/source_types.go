package types

type ListSourcesResponse struct {
	Sources []SourcesResponse `json:"sources"`
}

type SourceFilters struct {
	FilterType string `json:"filterType"`
	Name       string `json:"name"`
	Regexp     string `json:"regexp"`
}

type SourcesResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	SourceType string `json:"sourceType"`
	Alive      bool   `json:"alive"`
}
