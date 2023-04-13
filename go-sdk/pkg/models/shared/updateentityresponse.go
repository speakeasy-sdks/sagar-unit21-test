// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateEntityResponse - OK
type UpdateEntityResponse struct {
	// Unique identifier of the entity. Entity IDs must be unique and only comprise of the characters -_:.@a-zA-Z0-9!#$%&*+/=?^`{'
	EntityID *string `json:"entity_id,omitempty"`
	// A Unit21 internally-assigned unique identifier for an object within the Unit21 system.
	Unit21ID *string `json:"unit21_id,omitempty"`
}
