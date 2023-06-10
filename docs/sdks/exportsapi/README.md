# ExportsAPI

## Overview

If you want to download files requested via the dashboard or the bulk export endpoints. The `/file-exports` endpoint can get and list sars. 


### Available Operations

* [DownloadFileExport](#downloadfileexport) - Download export
* [ListExports](#listexports) - List exports

## DownloadFileExport

Returns a signed url to download the file.

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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ExportsAPI.DownloadFileExport(ctx, operations.DownloadFileExportRequest{
        FileExportID: 675439,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DownloadFileExportRequest](../../models/operations/downloadfileexportrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DownloadFileExportResponse](../../models/operations/downloadfileexportresponse.md), error**


## ListExports

Returns paginated list exports. It will only show the exports initiated by the requester (The requester is the creator of the API key)

**This endpoint omits any exports from the "System" source (generated directly from the Dashboard instead of the API).**

Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 50)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record. The offset is relative to the number of pages (not the total count of objects).

The `total_count` field contains the total number of exports where the  `response_count` field contains the number of exports included in the response.

The `statuses` for exports address:

  | Status                   | Description                                             |
  |--------------------------|---------------------------------------------------------|
  | READY_FOR_DOWNLOAD	     | File is ready for download                              |
  | GENERATING	             | File is generating                                      |
  | FAILED                   | File export failed                                      |
  | REQUESTED	               | File exort has been requested                           |

### Example Usage

```go
package main

import(
	"context"
	"log"
	"unit21"
	"unit21/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.ExportsAPI.ListExports(ctx, shared.ListExports{
        FileExportIds: []int64{
            249796,
            581273,
            313218,
            881736,
        },
        Limit: sdk.Int64(2),
        Offset: sdk.Int64(965417),
        Statuses: []string{
            "provident",
            "nam",
            "id",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [shared.ListExports](../../models/shared/listexports.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.ListExportsResponse](../../models/operations/listexportsresponse.md), error**

