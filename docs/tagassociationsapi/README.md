# TagAssociationsAPI

## Overview

Tags provide a flexible means of linking objects together in the Unit21 system. You can use the `/tag-associations` endpoint to explore these associations.


### Available Operations

* [ListTags](#listtags) - List tags

## ListTags

Returns an array of objects associated with a set of tags in your environment.     

Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 1000)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record.

To narrow down your tag association search, we provide filter parameters to this endpoint.

  | Field                   | Type        | Description                                                                                                       |
  | ----------------------- | ----------- | ----------------------------------------------------------------------------------------------------------------- |
  | `created_after`         | Numeric     | Tag associations created on or after this unix timestamp                                                          |
  | `created_before`        | Numeric     | Tag associations created before this unix timestamp                                                               |
  | `object_types`          | String[]    | List of object types to match against. Supported values are `alert`, `case`, `sar`, `rule`, `agent`, `event`, `entity`, and `instrument`. Specifying [`entity`, `alert`] means that we will only match against tags associated with entities and alerts in the system, and will not return results of tags associated with rules, events etc. If more than one value is provided to `object_types` and `object_id` is specified, an error will be thrown.     |
  | `object_id`             | String      | String representing the unit21 ID of the object you want to get tag associations for. If this is specified and `object_types` contains more than one value, an error will be thrown.                    |
  | `tag_filters`           | String[]    | List of string tags (`key:value`) or keys to associate this case with (e.g. `case_type:high_velocity` or `case_type`). If only the key is provided, we will match against all tags with that key        |
  | `limit`                 | Numeric     | A limit on the number of objects to be returned. Limit can range between 1 and 50, and the default is 10          |
  | `offset`                | Numeric     | The offset for pagination. Default is 1                                                                           |
  | `options`               | Object      | Options for the data included in the returned cases. Removing unneeded options can improve response speed         |


The `total_count` field contains the total number of tags where the  `response_count` field contains the number of tags included in the response.


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
    res, err := s.TagAssociationsAPI.ListTags(ctx, operations.ListTagsRequestBody{
        CreatedAfter: sdk.Int64(1623365011),
        CreatedBefore: sdk.Int64(1623365011),
        Limit: sdk.Int64(2),
        ObjectID: sdk.Int64(714242),
        ObjectTypes: []ListTagsRequestBodyObjectTypesEnum{
            operations.ListTagsRequestBodyObjectTypesEnumInstrument,
            operations.ListTagsRequestBodyObjectTypesEnumEntity,
        },
        Offset: sdk.Int64(149448),
        TagFilters: []string{
            "Sector:Europe",
            "Sector:Europe",
            "Sector:Europe",
            "Sector:Europe",
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
