package types

type AsyncJobStatus struct {
	// Whether or not the request is in progress (`InProgress`), has completed successfully (`Success`), or has completed with an error (`Failed`).
	Status string `json:"status"`
	// Additional status message generated if the status is not `Failed`.
	StatusMessage string            `json:"statusMessage,omitempty"`
	Error_        *ErrorDescription `json:"error,omitempty"`
}

type BeginAsyncJobResponse struct {
	// Identifier to get the status of an asynchronous job.
	Id string `json:"id"`
}
