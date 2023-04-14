// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unit21/pkg/models/shared"
)

type UpdateInstrumentRequest struct {
	RequestBody interface{} `request:"mediaType=application/json"`
	// Unique identifier of the instrument on your platform
	InstrumentID string `pathParam:"style=simple,explode=false,name=instrument_id"`
	// Name of organization in your environment
	OrgName string `pathParam:"style=simple,explode=false,name=org_name"`
}

// UpdateInstrumentMessageGeneralResponse - OK
type UpdateInstrumentMessageGeneralResponse struct {
	// Error message
	ErrorCode *string `json:"error_code,omitempty"`
	// Detailed message
	Message *string `json:"message,omitempty"`
}

type UpdateInstrumentResponse struct {
	ContentType string
	// OK
	MessageGeneralResponse *UpdateInstrumentMessageGeneralResponse
	StatusCode             int
	RawResponse            *http.Response
	// OK
	UpdateInstrumentResponse *shared.UpdateInstrumentResponse
}
