// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unit21/pkg/models/shared"
)

// ListEntitiesMessageGeneralResponse - OK
type ListEntitiesMessageGeneralResponse struct {
	// Error message
	ErrorCode *string `json:"error_code,omitempty"`
	// Detailed message
	Message *string `json:"message,omitempty"`
}

type ListEntitiesResponse struct {
	ContentType string
	// OK
	MessageGeneralResponse *ListEntitiesMessageGeneralResponse
	StatusCode             int
	RawResponse            *http.Response
	// OK
	ListEntityResponse *shared.ListEntityResponse
}
