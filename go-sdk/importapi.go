// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"unit21/pkg/models/operations"
	"unit21/pkg/utils"
)

// importAPI - If you want to upload a file with a lot of information in it, you can send them via a POST to a custom URL. You can request this URL via a POST to the `/datafiles` endpoint.  For the fastest processing, the file SHOULD be JSON or CSV.
type importAPI struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newImportAPI(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *importAPI {
	return &importAPI{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// DatafileStatus - Retrieve datafile status
// Retrieve datafile status.
//
// Note `file_id` will be included in the initial request to get a presigned_url.
//
// Programmatically check on the status of files and take action should errors occur.
//
// Files uploaded and processed have the following `status` value with the following definitions:
//
//	| Error code               | Definition                                                                                                                                 |
//	|--------------------------|--------------------------------------------------------------------------------------------------------------------------------------------|
//	| `PENDING_UPLOAD`	       | Customer is programmatically uploading via API, but the file has not landed (or not yet been detected as landed) in S3                     |
//	| `ADDED`	                 | Customer manually uploaded to the UI, but has not attempted to process the file yet                                                        |
//	| `QUEUED`	               | File is in a queue waiting to process                                                                                                      |
//	| `PROCESSING`	           | File is presently being processed                                                                                                          |
//	| `FINISHED`	             | The File finished successfully. Note that this does not mean all data is processed successfully as referenced by Hard error Handling below |
//	| `FAILED`	               | File hit a hard failure case and was unable to process data                                                                                |
//
// Hard errors refer to unprocessable datafiles, aka files that whose status end up `FAILED` are accompanied by an error code from this list:
//
//	| Error code               | Definition                                                                                                                                                                                                                                           |
//	|--------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
//	| `unparseable_file`	     | File failed to read (i.e. file contents was not actually csv or the contents use a non-traditional delimiter character)                                                                                                                              |
//	| `invalid_schema`	       | By the stream configuration, the file had unexpected column headers or values present and thus the system did not process the data                                                                                                                   |
//	| `stream_not_configured`	 | This error means that the stream has not yet been configured with all the necessary settings to ingest the data yet. Generally this should only happen if you are testing uploaded datafiles in advance of having defined landing the data in Unit21 |
//	| `unknown`	               | This is akin to 500 server error, and Unit21 does not have a specific known cause at this time                                                                                                                                                       |
func (s *importAPI) DatafileStatus(ctx context.Context, request operations.DatafileStatusRequest) (*operations.DatafileStatusResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/imports/{file_id}", request, nil)

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

	res := &operations.DatafileStatusResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetPreSignedURL - Get pre-signed URL
// Get details your unique URL you can use to import data into the Unit21 system.
//
// The response will include a URL which you must use to upload your datafile. In the example below, your datafile will be uploaded to https://local-tm-uploads.s3.amazonaws.com/.
func (s *importAPI) GetPreSignedURL(ctx context.Context, request operations.GetPreSignedURLRequestBody) (*operations.GetPreSignedURLResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/imports/pre-signed-url/create"

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

	res := &operations.GetPreSignedURLResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// ListDatafiles - Retrieve datafiles list
// Retrieve list of datafiles.
//
// Note `file_id` will be included in the initial request to get a presigned_url.
//
// This route will be limited to 1000 records ordered by upload time.
func (s *importAPI) ListDatafiles(ctx context.Context, request operations.ListDatafilesRequestBody) (*operations.ListDatafilesResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/imports/list"

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

	res := &operations.ListDatafilesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// UploadDatafiles - Upload data to URL
// Upload your file to our S3 bucket using the pre signed AWS URL you received in the `create` endpoint.
//
// **PLEASE NOTE THAT THE URL FOR THIS POST REQUEST IS THE `pre_signed_url` ONLY.** YOU MUST REMOVE `https://sandbox1-api.unit21.com/v1`.
//
// Documentation shows an incorrect example: https://sandbox1-api.unit21.com/v1/local-tm-uploads.s3.amazonaws.com/.
//
// The correct URL would be: https://local-tm-uploads.s3.amazonaws.com/
//
// Only one file can be uploaded in a request, with a file size maximum of 1GB. Please add a waiting time of two seconds between requests.
//
// Use `--form file` to specify the file.
//
// We support JSON, JSONL, CSV format only.
func (s *importAPI) UploadDatafiles(ctx context.Context, request operations.UploadDatafilesRequest) (*operations.UploadDatafilesResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/{pre_signed_url}", request, nil)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "multipart")
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

	res := &operations.UploadDatafilesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}
