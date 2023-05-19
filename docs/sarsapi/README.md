# SarsAPI

## Overview

Sars are cases that have been investigated and turned into a Suspicious Activity report with the intent to file it to FinCen. The `/sars` endpoint can get and list sars. 


### Available Operations

* [ExportSars](#exportsars) - Bulk export sars
* [ListSars](#listsars) - List sars
* [ReadOneSar](#readonesar) - Get a sars

## ExportSars

Initiates an email and dashboard export of sars. The export will be as a CSV file.

Either the agent `ID` or `email` is required to begin the export.

Either the `filters` or the list of `sar IDs` are required for the export.

Custom data filters are not supported for bulk exports at this time.


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
    res, err := s.SarsAPI.ExportSars(ctx, operations.ExportSarsRequestBody{
        Filters: &operations.ExportSarsRequestBodyFilters{
            CreatedAtEnd: sdk.String("2021-11-05 04:13:46"),
            CreatedAtStart: sdk.String("2019-11-05 04:13:46"),
            FiledAtEnd: sdk.String("2021-11-05 04:13:46"),
            FiledAtStart: sdk.String("2019-11-05 04:13:46"),
            ReportType: operations.ExportSarsRequestBodyFiltersReportTypeFincenSar.ToPointer(),
            Status: sdk.String("active"),
            Statuses: []ExportSarsRequestBodyFiltersStatuses{
                operations.ExportSarsRequestBodyFiltersStatusesFincenStatusOther,
                operations.ExportSarsRequestBodyFiltersStatusesPrefilingValidationSuccess,
                operations.ExportSarsRequestBodyFiltersStatusesFincenValidationFailed,
                operations.ExportSarsRequestBodyFiltersStatusesSentToFincen,
            },
            Submitted: sdk.Bool(true),
            TagIds: []int64{
                9,
                9,
                9,
            },
            UpdatedAtEnd: sdk.String("2021-11-05 04:13:46"),
            UpdatedAtStart: sdk.String("2019-11-05 04:13:46"),
        },
        SarIds: []int64{
            100226,
            99569,
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

## ListSars

Returns paginated list of of top-level information about paths/sars@list.yaml    

Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 50)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record. The offset is relative to the number of pages (not the total count of objects).

To narrow down your sars search, we provide filter parameters to this endpoint. Note that all list inputs function as an "or" filter, as in any one of the values must match the selected sar(s):


  | Field                   | Type        | Description                                                                                                       |
  | ----------------------- | ----------- | ----------------------------------------------------------------------------------------------------------------- |
  | `created_after`         | Numeric     | SARs created on or after this unix timestamp                                                                      |
  | `created_before`        | Numeric     | SARs created before this unix timestamp                                                                           |
  | `tag_filters`           | String[]    | List of string tags (`key:value`) or keys to associate this SARs with (e.g. `sars_type:high_velocity` or `sars_type`). If only the key is provided, we will match against all tags with that key        |
  | `limit`                 | Numeric     | A limit on the number of objects to be returned. Limit can range between 1 and 50, and the default is 10          |
  | `offset`                | Numeric     | The offset for pagination. Default is 1                                                                           |
  | `options`               | Object      | Options for the data included in the returned SARs. Removing unneeded options can improve response speed          |


The `total_count` field contains the total number of sars where the  `response_count` field contains the number of sars included in the response.

Follow the links for more information:
  - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
  

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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.SarsAPI.ListSars(ctx, shared.ListRequest{
        Limit: sdk.Int64(2),
        Offset: sdk.Int64(919483),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## ReadOneSar

Returns all data objects belonging to a single SAR.

This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the sar is first created.

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
    res, err := s.SarsAPI.ReadOneSar(ctx, operations.ReadOneSarRequest{
        Unit21ID: "ullam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReadOneSar200ApplicationJSONAny != nil {
        // handle response
    }
}
```
