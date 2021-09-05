package types

type ListSourcesResponse struct {
	Sources []SourcesResponse `json:"sources"`
}

type SourcesResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	SourceType string `json:"sourceType"`
	Alive      bool   `json:"alive"`
}
