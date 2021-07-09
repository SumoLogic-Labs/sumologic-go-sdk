package types

import (
	"github.com/antihax/optional"
	"time"
)

type AccessKey struct {
	// Identifier of the access key.
	Id string `json:"id"`
	// The name of the access key.
	Label string `json:"label"`
	// An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don't match any entry in the allowlist.
	CorsHeaders []string `json:"corsHeaders,omitempty"`
	// Indicates whether the access key is disabled or not.
	Disabled bool `json:"disabled"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the access key.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// The key for the created access key. This field will have values only in the response for an access key create request. The value will be an empty string while listing all keys.
	Key string `json:"key"`
}

type AccessKeyCreateRequest struct {
	// A name for the access key to be created.
	Label string `json:"label"`
	// An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don't match any entry in the allowlist.
	CorsHeaders []string `json:"corsHeaders,omitempty"`
}

type AccessKeyManagementApiListAccessKeysOpts struct {
	Limit optional.Int32
	Token optional.String
}

type AccessKeyPublic struct {
	// Identifier of the access key.
	Id string `json:"id"`
	// The name of the access key.
	Label string `json:"label"`
	// An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don't match any entry in the allowlist.
	CorsHeaders []string `json:"corsHeaders,omitempty"`
	// Indicates whether the access key is disabled or not.
	Disabled bool `json:"disabled"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the access key.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
}

type ListAccessKeysResult struct {
	// An array of access keys.
	Data []AccessKeyPublic `json:"data"`
}

type PaginatedListAccessKeysResult struct {
	// An array of access keys.
	Data []AccessKeyPublic `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}

type AccessKeyUpdateRequest struct {
	// Indicates whether the access key is disabled or not.
	Disabled bool `json:"disabled"`
	// An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don't match any entry in the allowlist.
	CorsHeaders []string `json:"corsHeaders,omitempty"`
}
