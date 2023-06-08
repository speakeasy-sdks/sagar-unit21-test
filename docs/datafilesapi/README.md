# DatafilesAPI

## Overview

If you want to bulk upload multiple objects, you can send them via a POST to the `/datafiles` endpoint. For the fastest processing, the datafile SHOULD be a JSON file in the format of a typical POST request to this API.


### Available Operations

* [CreateDatafiles](#createdatafiles) - Upload datafiles
* [GetDatafileByUnit21ID](#getdatafilebyunit21id) - Get datafile
* [GetDatafileMappings](#getdatafilemappings) - Retrieve datafile mappings

## CreateDatafiles

Bulk upload multiple objects of the same type. Can be entities, events, or instruments.

Only one file can be uploaded in a request, with a file size maximum of 1GB. Please add a waiting time of two seconds between requests.

Use `--form datafile` to specify the datafile, and `run_rules` to configure whether to run Unit21 rules on the datafile after it's processed.

We support JSON format only.


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
    res, err := s.DatafilesAPI.CreateDatafiles(ctx, operations.CreateDatafilesRequestBody{
        Datafile: &operations.CreateDatafilesRequestBodyDatafile{
            Content: []byte("quia"),
            Datafile: "quis",
        },
        RunRules: sdk.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetDatafileByUnit21ID

Get details about a datafile.

This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the datafile is first created.

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
    res, err := s.DatafilesAPI.GetDatafileByUnit21ID(ctx, operations.GetDatafileByUnit21IDRequest{
        Unit21ID: "vitae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetDatafileMappings

Retrieve datafile mapping of objects uploaded with a datafile. 

Includes 3 arrays: entities, events, instruments. Each list is limited to a maximum of 500 items. 

The total number of items can be retrieved from the GET endpoint.

Please note that an empty response `{}` will be returned if the datafile is not yet processed.


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
    res, err := s.DatafilesAPI.GetDatafileMappings(ctx, operations.GetDatafileMappingsRequest{
        RequestBody: &operations.GetDatafileMappingsRequestBody{
            Offset: sdk.Int64(1),
        },
        Unit21ID: "laborum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
