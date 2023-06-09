// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"unit21/pkg/models/operations"
	"unit21/pkg/models/shared"
	"unit21/pkg/utils"
)

// eventsAPI - Events have two types, _transaction events_  and _action events_:
// * Transaction events are any monetary flow that is sent or received by an entity on your system.
// * Action events are non-monetary changes of state that occur on your system, e.g. user logins. The `/events` endpoint sends and receives data about significant actions that occur with an entity or instrument on your system.
type eventsAPI struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newEventsAPI(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *eventsAPI {
	return &eventsAPI{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateEvent - Create an event
// Creates a new event, sending event data in the request body.
//
// Two objects are required: `general_data` and either `transaction_data` or `action_data`. `general_data` requires the fields: `event_id`, `event_type`, and `event_time`. `transaction_data` requires only the field `amount`.
//
// Unlike entities, events on our system are cannot be explicitly updated. However, they can be  overwritten in a naive upsert-overwrite fashion.
//
// If we receive a request to create an event for an `event_id` that already exists in our system,  we will simply overwrite that previous entry with the newly provided data if this transaction  is not already associated with other articles in our system.
//
// For instance, if a transaction is flagged in an alert and we receive a request to overwrite  the details of this transaction, we will respond with a **409 error code** indicating that this  event cannot be overwritten.
//
// Updates to an event's `general_data.event_id` are not allowed.
//
// Follow the links for more information:
//   - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
//   - [Verification options](https://docs.unit21.ai/reference/identity-verification-options)
//   - [Custom data](https://docs.unit21.ai/reference/best-practices-for-custom-data)
//   - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
//   - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)
func (s *eventsAPI) CreateEvent(ctx context.Context, request operations.CreateEventEventOptions) (*operations.CreateEventResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/events/create"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
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

	res := &operations.CreateEventResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// ExportEvents - Bulk export events
// Initiates an email and dashboard export of events. The export will be as a CSV file.
//
// Either the agent `ID` or `email` is required to begin the export.
//
// Either the `filters` or the list of `event IDs` are required for the export.
//
// Custom data filters are not supported for bulk exports at this time.
func (s *eventsAPI) ExportEvents(ctx context.Context, request operations.ExportEventsRequestBody) (*operations.ExportEventsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/events/bulk-export"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
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

	res := &operations.ExportEventsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// ExportTransactions - Bulk export transactions
// Initiates an email and dashboard export of events. The export will be as a CSV file.
//
// Either the agent `ID` or `email` is required to begin the export.
//
// Either the `filters` or the list of `event IDs` are required for the export.
//
// Custom data filters are not supported for bulk exports at this time.
func (s *eventsAPI) ExportTransactions(ctx context.Context, request operations.ExportTransactionsRequestBody) (*operations.ExportTransactionsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/txn-events/bulk-export"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
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

	res := &operations.ExportTransactionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetEvent - Get an event
// Returns all data objects belonging to a single event.
//
// This endpoint requires the `events_id` which is a unique ID created by your organization to identify the event. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.
func (s *eventsAPI) GetEvent(ctx context.Context, request operations.GetEventRequest) (*operations.GetEventResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/{org_name}/events/{event_id}", request, nil)

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

	res := &operations.GetEventResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetEvent200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetEvent200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListEvents - List events
// Returns an array of top-level information about events in your environment.
//
// Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
// * `limit`  indicates how many objects the request returns (the page maximum is 50)
// * `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record.
// * `case_id`  is a filter. Only events with the associated case ID will be shown.
// * `alert_id` is a filter. Only events with the associated alert ID will be shown.
// * `start_date`  is a filter. Only events that started on or after this date will be shown.
// * `end_date` is a filter. Only events that ended on or before this date will be shown.
//
// The `total_count` field contains the total number of events where the  `response_count` field contains the number of events included in the response.
func (s *eventsAPI) ListEvents(ctx context.Context, request shared.ListDateRequest) (*operations.ListEventsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/events/list"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
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

	res := &operations.ListEventsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ListResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListResponse = out
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 409:
		fallthrough
	case httpRes.StatusCode == 423:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		fallthrough
	case httpRes.StatusCode == 503:
	}

	return res, nil
}

// UpdateEvent - Update event
// Update an event using the `event_id` from your platform.
//
// Updating an event has no required fields. You MAY send any  subset of the fields that the events/create endpoint accepts.
//
// This endpoint requires the `event_id` which is a unique ID created by your organization to identify the event. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.
//
// Note that you can also update an event using an upsert through `/events/create`.
//
// Follow the links for more information:
//   - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
//   - [Custom data](https://docs.unit21.ai/reference/best-practices-for-custom-data)
//   - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
//   - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)
func (s *eventsAPI) UpdateEvent(ctx context.Context, request operations.UpdateEventRequest) (*operations.UpdateEventResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/{org_name}/events/{event_id}/update", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
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

	res := &operations.UpdateEventResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
