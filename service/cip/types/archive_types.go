package types

import (
	"github.com/antihax/optional"
	"time"
)

type ArchiveJob struct {
	// The name of the ingestion job.
	Name string `json:"name"`
	// The starting timestamp of the ingestion job.
	StartTime time.Time `json:"startTime"`
	// The ending timestamp of the ingestion job.
	EndTime time.Time `json:"endTime"`
	// The unique identifier of the ingestion job.
	Id string `json:"id"`
	// The total number of objects scanned by the ingestion job.
	TotalObjectsScanned int64 `json:"totalObjectsScanned"`
	// The total number of objects ingested by the ingestion job.
	TotalObjectsIngested int64 `json:"totalObjectsIngested"`
	// The total bytes ingested by the ingestion job.
	TotalBytesIngested int64 `json:"totalBytesIngested"`
	// The status of the ingestion job, either `Pending`,`Scanning`,`Ingesting`,`Failed`, or `Succeeded`.
	Status string `json:"status"`
	// The creation timestamp in UTC of the ingestion job.
	CreatedAt time.Time `json:"createdAt"`
	// The identifier of the user who created the ingestion job.
	CreatedBy string `json:"createdBy"`
}

type ArchiveJobsCount struct {
	// Identifier for the archive source.
	SourceId string `json:"sourceId"`
	// The total number of archive jobs with pending status for the archive source.
	Pending int64 `json:"pending"`
	// The total number of archive jobs with scanning status for the archive source.
	Scanning int64 `json:"scanning"`
	// The total number of archive jobs with ingesting status for the archive source.
	Ingesting int64 `json:"ingesting"`
	// The total number of archive jobs with failed status for the archive source.
	Failed int64 `json:"failed"`
	// The total number of archive jobs with succeeded status for the archive source.
	Succeeded int64 `json:"succeeded"`
}

type ArchiveManagementApiListArchiveJobsBySourceIdOpts struct {
	Limit optional.Int32
	Token optional.String
}

type CreateArchiveJobRequest struct {
	// The name of the ingestion job.
	Name string `json:"name"`
	// The starting timestamp of the ingestion job.
	StartTime time.Time `json:"startTime"`
	// The ending timestamp of the ingestion job.
	EndTime time.Time `json:"endTime"`
}

type ListArchiveJobsCount struct {
	// List of archive sources with count of jobs having various statuses.
	Data []ArchiveJobsCount `json:"data"`
}

type ListArchiveJobsResponse struct {
	// List of Archive Jobs.
	Data []ArchiveJob `json:"data"`
	// Next continuation token.
	Next string `json:"next,omitempty"`
}
