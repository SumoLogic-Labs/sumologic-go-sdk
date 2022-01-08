package types

type ErrorDescription struct {
	// An error code describing the type of error.
	Code string `json:"code"`
	// A short English-language description of the error.
	Message string `json:"message"`
	// An optional fuller English-language description of the error.
	Detail string `json:"detail,omitempty"`
	// An optional list of metadata about the error.
	Meta ErrorMeta `json:"meta,omitempty"`
}

type ErrorMeta struct {
	Reason string `json:"reason,omitempty"`
}

type ErrorResponse struct {
	// An identifier for the error; this is unique to the specific API request.
	Id string `json:"id"`
	// A list of one or more causes of the error.
	Errors []ErrorDescription `json:"errors"`
}

type LegacyErrorResponse struct {
	// Error code
	Code string `json:"code"`
	// An identifier for the error; this is unique to the specific API request.
	Id string `json:"id"`
	// Error message
	Message string `json:"message"`
	// Error status
	Status int `json:"status"`
}
