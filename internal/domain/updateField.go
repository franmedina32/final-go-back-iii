package domain

type UpdateFieldRequest struct {
	FieldName string `json:"field_name" example:"nombre"` // Field name to update
	Value     string `json:"value" example:"John"`      // New value for the field
}
