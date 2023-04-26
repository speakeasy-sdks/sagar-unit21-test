# ImportAPI

## Overview

If you want to upload a file with a lot of information in it, you can send them via a POST to a custom URL. You can request this URL via a POST to the `/datafiles` endpoint.  For the fastest processing, the file SHOULD be JSON or CSV.


### Available Operations

* [DatafileStatus](#datafilestatus) - Retrieve datafile status
* [GetPreSignedURL](#getpresignedurl) - Get pre-signed URL
* [ListDatafiles](#listdatafiles) - Retrieve datafiles list
* [UploadDatafiles](#uploaddatafiles) - Upload data to URL

## DatafileStatus

Retrieve datafile status.

Note `file_id` will be included in the initial request to get a presigned_url.

Programmatically check on the status of files and take action should errors occur.

Files uploaded and processed have the following `status` value with the following definitions:

  | Error code               | Definition                                                                                                                                 |
  |--------------------------|--------------------------------------------------------------------------------------------------------------------------------------------|
  | `PENDING_UPLOAD`	       | Customer is programmatically uploading via API, but the file has not landed (or not yet been detected as landed) in S3                     |
  | `ADDED`	                 | Customer manually uploaded to the UI, but has not attempted to process the file yet                                                        |
  | `QUEUED`	               | File is in a queue waiting to process                                                                                                      |
  | `PROCESSING`	           | File is presently being processed                                                                                                          |
  | `FINISHED`	             | The File finished successfully. Note that this does not mean all data is processed successfully as referenced by Hard error Handling below |
  | `FAILED`	               | File hit a hard failure case and was unable to process data                                                                                |


Hard errors refer to unprocessable datafiles, aka files that whose status end up `FAILED` are accompanied by an error code from this list:

  | Error code               | Definition                                                                                                                                                                                                                                           |
  |--------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
  | `unparseable_file`	     | File failed to read (i.e. file contents was not actually csv or the contents use a non-traditional delimiter character)                                                                                                                              |
  | `invalid_schema`	       | By the stream configuration, the file had unexpected column headers or values present and thus the system did not process the data                                                                                                                   |
  | `stream_not_configured`	 | This error means that the stream has not yet been configured with all the necessary settings to ingest the data yet. Generally this should only happen if you are testing uploaded datafiles in advance of having defined landing the data in Unit21 |
  | `unknown`	               | This is akin to 500 server error, and Unit21 does not have a specific known cause at this time                                                                                                                                                       |


### Example Usage

```go
package main

import(
	"context"
	"log"
	"unit21"
	"unit21/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.DatafileStatusRequest{
        FileID: 501324,
    }

    res, err := s.ImportAPI.DatafileStatus(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetPreSignedURL

Get details your unique URL you can use to import data into the Unit21 system.

The response will include a URL which you must use to upload your datafile. In the example below, your datafile will be uploaded to https://local-tm-uploads.s3.amazonaws.com/.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"unit21"
	"unit21/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetPreSignedURLRequestBody{
        FileName: "custom_data.csv",
        Md5Hash: sdk.String("352bfecf-ce8e-4c3d-64b9-ba0707fc2496"),
        StreamName: "insts",
    }

    res, err := s.ImportAPI.GetPreSignedURL(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## ListDatafiles

Retrieve list of datafiles.

Note `file_id` will be included in the initial request to get a presigned_url.

This route will be limited to 1000 records ordered by upload time.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"unit21"
	"unit21/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ListDatafilesRequestBody{
        CreatedAfter: sdk.Int64(1623365011),
        CreatedBefore: sdk.Int64(1621365011),
    }

    res, err := s.ImportAPI.ListDatafiles(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## UploadDatafiles

Upload your file to our S3 bucket using the pre signed AWS URL you received in the `create` endpoint.

**PLEASE NOTE THAT THE URL FOR THIS POST REQUEST IS THE `pre_signed_url` ONLY.** YOU MUST REMOVE `https://sandbox1-api.unit21.com/v1`.

Documentation shows an incorrect example: https://sandbox1-api.unit21.com/v1/local-tm-uploads.s3.amazonaws.com/.

The correct URL would be: https://local-tm-uploads.s3.amazonaws.com/

Only one file can be uploaded in a request, with a file size maximum of 1GB. Please add a waiting time of two seconds between requests.

Use `--form file` to specify the file.

We support JSON, JSONL, CSV format only.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"unit21"
	"unit21/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.UploadDatafilesRequest{
        RequestBody: &operations.UploadDatafilesRequestBody{
            File: &operations.UploadDatafilesRequestBodyFile{
                Content: []byte("deleniti"),
                File: "sapiente",
            },
        },
        PreSignedURL: "amet",
    }

    res, err := s.ImportAPI.UploadDatafiles(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
