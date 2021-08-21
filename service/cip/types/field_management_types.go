package types

type BuiltinField struct {
	// Field name.
	FieldName string `json:"fieldName"`
	// Identifier of the field.
	FieldId string `json:"fieldId"`
	// Field type. Possible values are `String`, `Long`, `Int`, `Double`, and `Boolean`.
	DataType string `json:"dataType"`
	// Indicates whether the field is enabled and its values are being accepted. Possible values are `Enabled` and `Disabled`.
	State string `json:"state"`
}

type CustomField struct {
	// Field name.
	FieldName string `json:"fieldName"`
	// Identifier of the field.
	FieldId string `json:"fieldId"`
	// Field type. Possible values are `String`, `Long`, `Int`, `Double`, and `Boolean`.
	DataType string `json:"dataType"`
	// Indicates whether the field is enabled and its values are being accepted. Possible values are `Enabled` and `Disabled`.
	State string `json:"state"`
}

type FieldName struct {
	// Field name.
	FieldName string `json:"fieldName"`
}

type FieldQuotaUsage struct {
	// Maximum number of fields available.
	Quota int32 `json:"quota"`
	// Current number of fields available.
	Remaining int32 `json:"remaining"`
}

type ListBuiltinFieldsResponse struct {
	// List of built-in fields.
	Data []BuiltinField `json:"data"`
}

type ListCustomFieldsResponse struct {
	// List of custom fields.
	Data []CustomField `json:"data"`
}

type ListDroppedFieldsResponse struct {
	// List of dropped fields.
	Data []DroppedField `json:"data"`
}
