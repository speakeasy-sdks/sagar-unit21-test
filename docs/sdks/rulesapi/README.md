# RulesAPI

## Overview

Rules are the model logic that will find fraudulent and suspicious transactions and actions. Rules create alerts that can turn into cases with flagged entities, transactions and instruments. The `/rules` endpoint can get and list rules. 


### Available Operations

* [ExportRules](#exportrules) - Bulk export rules
* [ListRules](#listrules) - List rules
* [ReadOneRule](#readonerule) - Get a rule

## ExportRules

Initiates an email and dashboard export of rules. The export will be as a CSV file.

Either the agent `ID` or `email` is required to begin the export.

Either the `filters` or the list of `rule IDs` are required for the export.

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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.RulesAPI.ExportRules(ctx, operations.ExportRulesRequestBody{
        Filters: &operations.ExportRulesRequestBodyFilters{
            AgentIds: []int64{
                14,
            },
            EndDate: sdk.String("2021-11-05 04:13:46"),
            RuleIds: []int64{
                14,
                14,
                14,
            },
            StartDate: sdk.String("2019-11-05 04:13:46"),
            Status: sdk.String("OPEN"),
            Statuses: []ExportRulesRequestBodyFiltersStatuses{
                operations.ExportRulesRequestBodyFiltersStatusesInactive,
                operations.ExportRulesRequestBodyFiltersStatusesActive,
                operations.ExportRulesRequestBodyFiltersStatusesInactive,
                operations.ExportRulesRequestBodyFiltersStatusesActive,
            },
            TagIds: []int64{
                9,
                9,
                9,
                9,
            },
        },
        RuleIds: []int64{
            367562,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ExportRulesRequestBody](../../models/operations/exportrulesrequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ExportRulesResponse](../../models/operations/exportrulesresponse.md), error**


## ListRules

Returns paginated list of of top-level information about rules.     
Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 50)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record. The offset is relative to the number of pages (not the total count of objects).

The `total_count` field contains the total number of rules where the  `response_count` field contains the number of rules included in the response.

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
    res, err := s.RulesAPI.ListRules(ctx, shared.ListRequest{
        Limit: sdk.Int64(2),
        Offset: sdk.Int64(97260),
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
| `request`                                                | [shared.ListRequest](../../models/shared/listrequest.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


### Response

**[*operations.ListRulesResponse](../../models/operations/listrulesresponse.md), error**


## ReadOneRule

Returns all data objects belonging to a single rule.

This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the rule is first created.

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
    res, err := s.RulesAPI.ReadOneRule(ctx, operations.ReadOneRuleRequest{
        Unit21ID: "iure",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReadOneRule200ApplicationJSONAny != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ReadOneRuleRequest](../../models/operations/readonerulerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ReadOneRuleResponse](../../models/operations/readoneruleresponse.md), error**

