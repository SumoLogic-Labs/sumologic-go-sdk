package types

import (
	"github.com/antihax/optional"
	"time"
)

type LookupAsyncJobStatus struct {
	// An identifier returned in response to an asynchronous request.
	JobId string `json:"jobId"`
	// Whether or not the request is pending (`Pending`), in progress (`InProgress`), has completed successfully (`Success`), has completed partially with warnings (`PartialSuccess`) or has completed with an error (`Failed`).
	Status string `json:"status"`
	// Additional status messages generated if any if the status is `Success`.
	StatusMessages []string `json:"statusMessages,omitempty"`
	// More information about the failures, if the status is `Failed`.
	Errors []ErrorDescription `json:"errors,omitempty"`
	// More information about the warnings, if the status is `PartialSuccess`.
	Warnings []WarningDescription `json:"warnings,omitempty"`
	// Content id of lookup table on which this operation was performed.
	LookupContentId string `json:"lookupContentId"`
	// Name of lookup table on which this operation was performed.
	LookupName string `json:"lookupName"`
	// Content path of lookup table on which this operation was performed.
	LookupContentPath string `json:"lookupContentPath"`
	// Type of asynchronous request made:   - `BulkMerge`   - `BulkReplace`   - `Truncate`
	RequestType string `json:"requestType,omitempty"`
	// User id of user who initiated this operation.
	UserId string `json:"userId"`
	// Creation time of this job in UTC.
	CreatedAt time.Time `json:"createdAt"`
	// Timestamp in UTC when status was last updated.
	ModifiedAt time.Time `json:"modifiedAt"`
}

type LookupRequestToken struct {
	// The identifier used to track the request.
	Id string `json:"id"`
}

type LookupTable struct {
	// The name of the lookup table.
	Name string `json:"name"`
	// The parent-folder-path identifier of the lookup table in the Library.
	ParentFolderId string `json:"parentFolderId"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Identifier of the lookup table as a content item.
	Id string `json:"id"`
	// Address/path of the parent folder of this lookup table in content library. For example, a lookup table existing  in the personal/lookupTable folder for user johndoe would be: /Library/Users/johndoe@acme.com/lookupTable
	ContentPath string `json:"contentPath,omitempty"`
	// The current size of the lookup table in bytes
	Size int64 `json:"size,omitempty"`
}

type LookupTableDefinition struct {
	// The description of the lookup table.
	Description string `json:"description"`
	// The list of fields in the lookup table.
	Fields []LookupTableField `json:"fields"`
	// The names of the fields that make up the primary key for the lookup table. These will be a subset of the fields that the table will contain.
	PrimaryKeys []string `json:"primaryKeys"`
	// A time to live for each entry in the lookup table (in minutes). 365 days is the maximum time to live for each entry that you can specify. Setting it to 0 means that the records will not expire automatically.
	Ttl int32 `json:"ttl,omitempty"`
	// The action that needs to be taken when the size limit is reached for the table. The possible values can be `StopIncomingMessages` or `DeleteOldData`. DeleteOldData will start deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached.
	SizeLimitAction string `json:"sizeLimitAction,omitempty"`
	// The name of the lookup table.
	Name string `json:"name"`
	// The parent-folder-path identifier of the lookup table in the Library.
	ParentFolderId string `json:"parentFolderId"`
}

type LookupTableField struct {
	// The name of the field.
	FieldName string `json:"fieldName"`
	// The data type of the field. Supported types:   - `boolean`   - `int`   - `long`   - `double`   - `string`
	FieldType string `json:"fieldType"`
}

type LookupTableUploadFileOpts struct {
	Merge        optional.Bool
	FileEncoding optional.String
}

type LookupUpdateDefinition struct {
	// A time to live for each entry in the lookup table (in minutes). 0 is a special value. A TTL of 0 implies entry will never be deleted from the table.
	Ttl int32 `json:"ttl"`
	// The description of the lookup table. The description cannot be blank.
	Description string `json:"description"`
	// The action that needs to be taken when the size limit is reached for the table. The possible values can be `StopIncomingMessages` or `DeleteOldData`. DeleteOldData will starting deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached.
	SizeLimitAction string `json:"sizeLimitAction,omitempty"`
}

type RowDeleteDefinition struct {
	// A list of all primary key field identifiers and their corresponding values.
	PrimaryKey []TableRow `json:"primaryKey"`
}

type RowUpdateDefinition struct {
	// A list of all the field identifiers and their corresponding values.
	Row []TableRow `json:"row"`
}

type TableRow struct {
	// Name of the column of the table.
	ColumnName string `json:"columnName"`
	// Value of the specified column.
	ColumnValue string `json:"columnValue"`
}

type WarningDescription struct {
	// Description of the warning.
	Message string `json:"message"`
	// An optional cause of this warning.
	Cause string `json:"cause,omitempty"`
}
