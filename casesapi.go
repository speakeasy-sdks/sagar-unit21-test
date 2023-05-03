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

// casesAPI - Cases are usually active investigations, which may span multiple events, entities and documents. They can be directly escalated into a suspicious activity report. The `/cases` endpoint can create, list, and update cases.
type casesAPI struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newCasesAPI(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *casesAPI {
	return &casesAPI{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateCase - Create a case
// Creates a new case, sending case data in the request body.
// To create a case, you MUST include the following fields: `case_id`, `title`, and `created_at`.  The other top-level fields are optional.
//
// If we receive a request to create a case for an `case_id` that already exists in our system,  we will respond with a **409 error code** indicating that this case cannot be created/updated. You must use the `/case/update` endpoint to update a case.
//
// You can add the following objects to a case:
//
//   | Field                    | Type     | Description                                                             |
//   |--------------------------|----------|-------------------------------------------------------------------------|
//   | `alerts`	               | Array[]  | Alerts that are associated with this case. Consists of `alert_id`s      |
//   | `events`	               | Array[]  | Transactions affiliated with the case. Consists of `event_id`s          |
//   | `entities`	             | Array[]  | Entities affiliated with the case. Consists of `entity_id`s             |
//
// Updates to a cases's `case_id` are not allowed.
//
// Follow the links for more information:
//   - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
//   - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
//   - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)
//
//
// The response will consist of the following fields:
//
//   | Field                    | Type     | Description                                          |
//   |--------------------------|----------|------------------------------------------------------|
//   | `case_id`	               | String   | 	Unique identifier of the case on your platform     |
//   | `previously_existed`	   | Boolean  | 	If case (with the same `case_id`) already exists   |
//

func (s *casesAPI) CreateCase(ctx context.Context, request operations.CreateCaseCaseData) (*operations.CreateCaseResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/cases/create"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
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

	res := &operations.CreateCaseResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.CreateCaseResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateCaseResponse = out
		}
	}

	return res, nil
}

// ExportCases - Bulk export cases
// Initiates an email and dashboard export of cases. The export will be as a CSV file.
//
// Either the agent `ID` or `email` is required to begin the export.
//
// Either the `filters` or the list of `case IDs` are required for the export.
//
// Custom data filters are not supported for bulk exports at this time.
//

func (s *casesAPI) ExportCases(ctx context.Context, request operations.ExportCasesRequestBody) (*operations.ExportCasesResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/cases/bulk-export"

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

	res := &operations.ExportCasesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetCaseByUnit21ID - Get a case
// Returns all data objects belonging to a single case.
//
// This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the case is first created.

func (s *casesAPI) GetCaseByUnit21ID(ctx context.Context, request operations.GetCaseByUnit21IDRequest) (*operations.GetCaseByUnit21IDResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/cases/{unit21_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

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

	res := &operations.GetCaseByUnit21IDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetCaseByUnit21ID200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// LinkMediaToCase - Add media to a case
// Adds rich media objects (images, videos, etc.) to an existing case.
//
// This endpoint is useful for sending in rich media such as profile pictures, ID card scans, official documents etc.  that you want available for investigative and verification purposes.
//
// Supported file types are: txt, pdf, video (mp4, mov, wmv, avi, mkv), images (png, jpg, tiff, gif, raw, eps).
//
// The payload to this endpoint can either be a **form-data** or a **base64** encoded media file via the requests JSON body.
//
// **Form-data** sent to this endpoint must use the key `media_key` and the `value` as the media file.  If you wish to provide optional information, use the `media_key` and provide stringified JSON data as the value.  There are no required fields in each media file's supplementary form data. However, if a recognized `media_type` value is provided,  the Unit21 system will be able to use the media object for purposes such as document verification.
//
// ```
//     --form 'document_front=@/src/103031/images/document_front.jpg' \
//     --form 'document_front={"media_type": "IMAGE_ID_CARD_FRONT", "source": "passport_app", "timestamp": 1572673229}'
// ```
//
// **Base64** encoded media objects must follow the format:
//
// ```json
//   {
//     "media": "iVBORw0KGgoAAAANSUhEUgAAAQMAAADCCAYAAABNEqduAAAgAElEQVR4Aey9CbgmV1Xv...",
//     "name": "Drivers_License.png",
//     "media_type": "IMAGE_DRIVERS_LICENSE_FRONT"
//   }
// ```
//
// `media` and `name` are the only required fields for each media object. The `name`` must include the file extension such a `File.pdf`. Supplementary form data is sent through the optional `custom_data` object.
//
// Recognized values of `media_type` are:
//
//   | media_type                    |
//   |-------------------------------|
//   | IMAGE_PROFILE_PICTURE         |
//   | IMAGE_DRIVERS_LICENSE_FRONT   |
//   | IMAGE_DRIVERS_LICENSE_BACK    |
//   | IMAGE_PASSPORT_FRONT          |
//   | IMAGE_ID_CARD_FRONT           |
//   | IMAGE_ID_CARD_BACK            |
//   | IMAGE_FACE_IMAGE              |

func (s *casesAPI) LinkMediaToCase(ctx context.Context, request operations.LinkMediaToCaseRequest) (*operations.LinkMediaToCaseResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/cases/{unit21_id}/link-media", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

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

	res := &operations.LinkMediaToCaseResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LinkMediaToCase200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// ListCases - List cases
// Returns an array of top-level information about cases in your environment.
//
// Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
// * `limit`  indicates how many objects the request returns (the page maximum is 50)
// * `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record.
//
// To narrow down your case search, we provide filter parameters to this endpoint. Note that all list inputs function as an "or" filter, as in any one of the values must match the selected case(s):
//
//
//   | Field                   | Type        | Description                                                                                                       |
//   | ----------------------- | ----------- | ----------------------------------------------------------------------------------------------------------------- |
//   | `created_after`         | Numeric     | Cases created on or after this unix timestamp                                                                     |
//   | `created_before`        | Numeric     | Cases created before this unix timestamp                                                                          |
//   | `dispositions`          | String[]    | List of case disposition states (defined on an integration basis)                                                 |
//   | `dispositioned_after`   | Numeric     | Cases with a disposition most recently updated after this unix timestamp                                          |
//   | `dispositioned_before`  | Numeric     | Cases with a disposition most recently updated before this unix timestamp                                         |
//   | `dispositioned_by`      | String[]    | List of agent emails. Returns alerts with a disposition most recently changed by agents in the list               |
//   | `rules`                 | Numeric[]   | List of Unit21 rule ids that are associated with the case                                                         |
//   | `associated_entities`   | Numeric[]   | List of Unit21 entity ids associated with this case                                                               |
//   | `associated_events`     | Numeric[]   | List of Unit21 event ids associated with this case                                                                |
//   | `associated_alerts`     | Numeric[]   | List of Unit21 alert ids associated with this case                                                                |
//   | `sources`               | String[]    | Must be list of alert sources: `INTERNAL`, `EXTERNAL`                                                             |
//   | `statuses`              | String[]    | Must be list of alert statuses: `OPEN`, `CLOSED`                                                                  |
//   | `tag_filters`           | String[]    | List of string tags (`key:value`) or keys to associate this case with (e.g. `case_type:high_velocity` or `case_type`). If only the key is provided, we will match against all tags with that key        |
//   | `limit`                 | Numeric     | A limit on the number of objects to be returned. Limit can range between 1 and 50, and the default is 10          |
//   | `offset`                | Numeric     | The offset for pagination. Default is 1                                                                           |
//   | `options`               | Object      | Options for the data included in the returned cases. Removing unneeded options can improve response speed         |
//
//
// The `total_count` field contains the total number of case where the  `response_count` field contains the number of cases included in the response.
//
// Follow the links for more information:
//   - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
//
//

func (s *casesAPI) ListCases(ctx context.Context, request operations.ListCasesRequestBody) (*operations.ListCasesResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/cases/list"

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

	res := &operations.ListCasesResponse{
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
	}

	return res, nil
}

// UpdateCase - Update case
// Update a case through its `unit21_id`. ONLY EXTERNAL CASES CAN BE UPDATED!
//
//
// Updating a case has no required fields. You MAY send any subset of the fields that the `cases/create` endpoint accepts.
//
// This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the case is first created.
//
// Note that you can also update an alert using an upsert through `/cases/create`.
//
// Follow the links for more information:
//   - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
//   - [Custom data](https://docs.unit21.ai/reference/best-practices-for-custom-data)
//   - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
//   - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)
//
//
// The response will consist of the following fields:
//
//   | Field                    | Type     | Description                                          |
//   |--------------------------|----------|------------------------------------------------------|
//   | `case_id`	               | String   | 	Unique identifier of the case on your platform     |
//   | `previously_existed`	   | Boolean  | 	If entity (with the same `case_id`) already exists |

func (s *casesAPI) UpdateCase(ctx context.Context, request operations.UpdateCaseRequest) (*operations.UpdateCaseResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/cases/{unit21_id}/update", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

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

	res := &operations.UpdateCaseResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateCase200ApplicationJSONAny = out
		}
	}

	return res, nil
}
