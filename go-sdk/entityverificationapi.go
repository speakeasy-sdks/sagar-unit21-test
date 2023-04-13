// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"net/http"
	"unit21/pkg/models/operations"
	"unit21/pkg/models/shared"
	"unit21/pkg/utils"
)

// entityVerificationAPI - Unit21 can execute entity verifications according to steps defined in system workflows. They can run when a new entity is added to the system, an update is made to an existing entity, or by directly triggering an API endpoint.
type entityVerificationAPI struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newEntityVerificationAPI(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *entityVerificationAPI {
	return &entityVerificationAPI{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// AddVerificationResultToEntity - Link external verification
// Add the verification result from an external ID provider to an entity on the Unit21 system.
// You can only send 1 result per request.
func (s *entityVerificationAPI) AddVerificationResultToEntity(ctx context.Context, request operations.AddVerificationResultToEntityRequest) (*operations.AddVerificationResultToEntityResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/entities/{unit21_id}/link-verification-result", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "LinkVerificationResult", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.AddVerificationResultToEntityResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LinkVerificationResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LinkVerificationResponse = out
		}
	}

	return res, nil
}

// GetEntityVerificationWorkflowExecutions - Get entity verification workflow IDs
// Returns the verification workflow IDs for an entity.
//
// This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the entity is first created.
func (s *entityVerificationAPI) GetEntityVerificationWorkflowExecutions(ctx context.Context, request operations.GetEntityVerificationWorkflowExecutionsRequest) (*operations.GetEntityVerificationWorkflowExecutionsResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/entities/{unit21_id}/verification_workflow_executions", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetEntityVerificationWorkflowExecutionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetVerificationResult - Get verification results by result id
// Returns all the information from the verification of a specific entity.
func (s *entityVerificationAPI) GetVerificationResult(ctx context.Context, request operations.GetVerificationResultRequest) (*operations.GetVerificationResultResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/verification/result/{result_id}", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetVerificationResultResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetVerificationResultFromWorkflowExecution - Get verification results from workflow
// Returns all the information from the verification workflow execution for a specific entity.
func (s *entityVerificationAPI) GetVerificationResultFromWorkflowExecution(ctx context.Context, request operations.GetVerificationResultFromWorkflowExecutionRequest) (*operations.GetVerificationResultFromWorkflowExecutionResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/verification/verification-workflow-execution/{verification_workflow_execution_id}/results", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetVerificationResultFromWorkflowExecutionResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetVerificationWorkflowExecution - Get verification workflow execution details
// Returns all the data associated with a verification_workflow_execution_id
func (s *entityVerificationAPI) GetVerificationWorkflowExecution(ctx context.Context, request operations.GetVerificationWorkflowExecutionRequest) (*operations.GetVerificationWorkflowExecutionResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/verification/verification-workflow-execution/{verification_workflow_execution_id}", request, nil)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetVerificationWorkflowExecutionResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// RunVerificationsWorkflowThroughExternalID - Verify an entity
// Run a verification workflow on an entity using the `entity_id` from your platform.
//
// Requires a `workflow_id`. You can create a verification workflow from the Unit21 dashboard.
//
// This endpoint requires the `entity_id` which is a unique ID created by your organization to identify the entity. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.
func (s *entityVerificationAPI) RunVerificationsWorkflowThroughExternalID(ctx context.Context, request operations.RunVerificationsWorkflowThroughExternalIDRequest) (*operations.RunVerificationsWorkflowThroughExternalIDResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/{org_name}/entities/{entity_id}/verify", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "EntityVerification", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.RunVerificationsWorkflowThroughExternalIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.RunVerificationsWorkflowThroughExternalID200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.RunVerificationsWorkflowThroughExternalID200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// UpdateContinuousMonitoring - Update continuous monitoring
// Fetch status and enables/disables Socure continuous monitoring for an entity.
//
// For asynchronous continuous monitoring, the endpoint will always return a 200 success status response.
//
// For synchronous continous monitoring, the endpoint will always return a 200 success status response  but you should look at the `is_success = true` field to check if the result was actually successful:
//
// `
//
//	{
//	  "error_message": "This entity has no existing continuous monitoring subscriptions to disable.",
//	  "is_success": true
//	}
//
// `
func (s *entityVerificationAPI) UpdateContinuousMonitoring(ctx context.Context, request operations.UpdateContinuousMonitoringRequest) (*operations.UpdateContinuousMonitoringResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/entities/{unit21_id}/continuous-monitoring", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "ContinuousMonitoring", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateContinuousMonitoringResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// UpdateSuppressedProviderEntities - Suppress provider entity
// Mute Socure continuous monitoring for an entity. 1 - Suppress 0 - Unsuppress
func (s *entityVerificationAPI) UpdateSuppressedProviderEntities(ctx context.Context, request operations.UpdateSuppressedProviderEntitiesRequest) (*operations.UpdateSuppressedProviderEntitiesResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/entities/{unit21_id}/suppress-provider-entity", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "SuppressProviderEntity", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateSuppressedProviderEntitiesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
