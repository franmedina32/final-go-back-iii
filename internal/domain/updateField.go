package domain

type UpdateFieldRequest struct {
	FieldName string `json:"field_name" example:"name"` // Field name to update
	Value     string `json:"value" example:"John"`      // New value for the field
}
