// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetVerificationWorkflowExecutionRequest struct {
	// A Unit21 internally-assigned unique identifier for the verification workflow execution.
	VerificationWorkflowExecutionID int64 `pathParam:"style=simple,explode=false,name=verification_workflow_execution_id"`
}

type GetVerificationWorkflowExecutionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
