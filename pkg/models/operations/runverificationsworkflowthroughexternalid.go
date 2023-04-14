// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unit21/pkg/models/shared"
)

type RunVerificationsWorkflowThroughExternalIDRequest struct {
	EntityVerification shared.EntityVerification `request:"mediaType=application/json"`
	EntityID           string                    `pathParam:"style=simple,explode=false,name=entity_id"`
	// Name of organization in your environment
	OrgName string `pathParam:"style=simple,explode=false,name=org_name"`
}

// RunVerificationsWorkflowThroughExternalID200ApplicationJSON - OK
type RunVerificationsWorkflowThroughExternalID200ApplicationJSON struct {
	// The end result of the workflow, one of a set of end results defined by the workflow
	EndAction *string `json:"end_action,omitempty"`
	// The raw response from the verification provider. Contents depend on the verification source and type.
	FullResponse map[string]interface{} `json:"full_response,omitempty"`
	// Whether or not the workflow successful completed. Workflows rely on external services, which may at times fail
	IsSuccess *bool `json:"is_success,omitempty"`
	// Object mapping from the executed verifications (e.g. IDOLOGY:DOC_VERIFICATION)
	Results map[string]interface{} `json:"results,omitempty"`
}

type RunVerificationsWorkflowThroughExternalIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	RunVerificationsWorkflowThroughExternalID200ApplicationJSONObject *RunVerificationsWorkflowThroughExternalID200ApplicationJSON
}