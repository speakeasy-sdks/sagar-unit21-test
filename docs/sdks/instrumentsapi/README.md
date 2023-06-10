# InstrumentsAPI

## Overview

Instruments represent any physical, digital, or logical intermediary between an entity and a transaction event. The `/instruments` endpoint can create, list, and update instruments. 


### Available Operations

* [CreateInstrument](#createinstrument) - Create an instrument
* [ExportInstruments](#exportinstruments) - Bulk export instruments
* [GetInstrument](#getinstrument) - Get an instrument
* [ListInstruments](#listinstruments) - List instruments
* [UpdateInstrument](#updateinstrument) - Update instrument

## CreateInstrument

Creates a new instrument, sending instrument data in the request body.

Recommended values for `instrument_type` include: account, crypto_address,  digital_wallet, credit_card, debit_card, gift_card, voucher, check.

If the `/instruments/create` API is called for an instrument that already exists in our system (i.e.  has an existing `instrument_id`, it is treated it as an  [upsert](https://docs.unit21.ai/reference/should-i-update-or-upsert) and an update on the existing  instrument is performed. The response to the request will then contain the entry `previously_existed: true`. 

Unit21 selectively ignores upserts if the request is identical to a previous request. The response to any  ignored upsert will contain the field `ignored: true`. 

Updates to an instrument's `instrument_id` are not allowed.

Follow the links for more information:
  - [Relationships](https://docs.unit21.ai/reference/relationships)
  - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
  - [Custom data](https://docs.unit21.ai/reference/best-practices-for-custom-data)
  - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
  - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)
  


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
    res, err := s.InstrumentsAPI.CreateInstrument(ctx, shared.CreateInstrumentRequest{
        CustomData: map[string]interface{}{
            "nisi": "vel",
            "natus": "omnis",
            "molestiae": "perferendis",
        },
        DigitalData: &shared.DigitalDataArrayProperties{
            ClientFingerprints: []string{
                "5454-ahfd-4531d-f4il",
                "5454-ahfd-4531d-f4il",
            },
            IPAddresses: []string{
                "255.255.255.255",
                "255.255.255.255",
            },
        },
        Entities: []shared.RelatedEntity{
            shared.RelatedEntity{
                EntityID: sdk.String("roster-1894"),
                EntityType: sdk.String("user"),
                RelationshipID: sdk.String("friend"),
            },
            shared.RelatedEntity{
                EntityID: sdk.String("roster-1894"),
                EntityType: sdk.String("user"),
                RelationshipID: sdk.String("friend"),
            },
            shared.RelatedEntity{
                EntityID: sdk.String("roster-1894"),
                EntityType: sdk.String("user"),
                RelationshipID: sdk.String("friend"),
            },
        },
        InstrumentID: "3234-sdghfdf-3332",
        InstrumentSubtype: sdk.String("visa"),
        InstrumentType: sdk.String("account"),
        LocationData: []shared.LocationDataProperties{
            shared.LocationDataProperties{
                BuildingNumber: sdk.String("6c"),
                City: sdk.String("Redmond"),
                Country: sdk.String("US"),
                PostalCode: sdk.String("07710-0001"),
                State: sdk.String("CA"),
                StreetName: sdk.String("California Ave"),
                Type: sdk.String("SHIPPING"),
                UnitNumber: sdk.String("22a"),
                VerifiedOn: sdk.Int64(1572673227),
            },
            shared.LocationDataProperties{
                BuildingNumber: sdk.String("6c"),
                City: sdk.String("Redmond"),
                Country: sdk.String("US"),
                PostalCode: sdk.String("07710-0001"),
                State: sdk.String("CA"),
                StreetName: sdk.String("California Ave"),
                Type: sdk.String("SHIPPING"),
                UnitNumber: sdk.String("22a"),
                VerifiedOn: sdk.Int64(1572673227),
            },
            shared.LocationDataProperties{
                BuildingNumber: sdk.String("6c"),
                City: sdk.String("Redmond"),
                Country: sdk.String("US"),
                PostalCode: sdk.String("07710-0001"),
                State: sdk.String("CA"),
                StreetName: sdk.String("California Ave"),
                Type: sdk.String("SHIPPING"),
                UnitNumber: sdk.String("22a"),
                VerifiedOn: sdk.Int64(1572673227),
            },
        },
        Options: &shared.Options{
            MergeCustomData: sdk.Bool(false),
            ResolveGeoip: sdk.Bool(true),
            UpsertOnConflict: sdk.Bool(true),
        },
        ParentInstrumentID: sdk.String("3234-sdghfdf-3331"),
        RegisteredAt: sdk.Int64(1623365011),
        Source: shared.CreateInstrumentRequestSourceInternal.ToPointer(),
        Status: sdk.String("active"),
        Tags: []string{
            "Sector:Europe",
            "Sector:Europe",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateInstrumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.CreateInstrumentRequest](../../models/shared/createinstrumentrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateInstrumentResponse](../../models/operations/createinstrumentresponse.md), error**


## ExportInstruments

Initiates an email and dashboard export of instruments. The export will be as a CSV file.

The agent making the request will need to have the correct permissions for the export to ensure success.

Either the `filters` or the list of `instrument IDs` are required for the export.

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
    res, err := s.InstrumentsAPI.ExportInstruments(ctx, operations.ExportInstrumentsRequestBody{
        Filters: &operations.ExportInstrumentsRequestBodyFilters{
            EndDate: sdk.String("2021-11-05 04:13:46"),
            EntityIds: []int64{
                15453219,
                15453219,
            },
            InstrumentType: sdk.String("wallet"),
            RuleIds: []int64{
                371,
                371,
            },
            StartDate: sdk.String("2019-11-05 04:13:46"),
            Status: sdk.String("active"),
            Statuses: []ExportInstrumentsRequestBodyFiltersStatuses{
                operations.ExportInstrumentsRequestBodyFiltersStatusesInactive,
                operations.ExportInstrumentsRequestBodyFiltersStatusesActive,
                operations.ExportInstrumentsRequestBodyFiltersStatusesInactive,
            },
            TagIds: []int64{
                9,
            },
        },
        InstrumentIds: []string{
            "3234-sdghfdf-3332",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BulkExportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ExportInstrumentsRequestBody](../../models/operations/exportinstrumentsrequestbody.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.ExportInstrumentsResponse](../../models/operations/exportinstrumentsresponse.md), error**


## GetInstrument

Returns all data objects belonging to a single instrument.

This endpoint requires the `instrument_id` which is a unique ID created by your organization to identify the instrument. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.

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
    res, err := s.InstrumentsAPI.GetInstrument(ctx, operations.GetInstrumentRequest{
        InstrumentID: "magnam",
        OrgName: "et",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InstrumentList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetInstrumentRequest](../../models/operations/getinstrumentrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetInstrumentResponse](../../models/operations/getinstrumentresponse.md), error**


## ListInstruments

Returns paginated list of of top-level information about instruments. 
Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 50)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record. The offset is relative to the number of pages (not the total count of objects).
* `alert_id` is a filter. Only instruments with the associated alert ID will be shown.

The `total_count` field contains the total number of instruments where the  `response_count` field contains the number of instruments included in the response.


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
    res, err := s.InstrumentsAPI.ListInstruments(ctx, shared.ListAlertRequest{
        AlertID: sdk.Int64(73913),
        Limit: 2,
        Offset: 569965,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListInstrumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.ListAlertRequest](../../models/shared/listalertrequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.ListInstrumentsResponse](../../models/operations/listinstrumentsresponse.md), error**


## UpdateInstrument

Updates an instrument's information using the `instrument_id` from your platform. 

Updating an instrument has no required fields. You MAY send any  subset of the fields that the `instruments/create` endpoint accepts.

This endpoint requires the `instrument_id` which is a unique ID created by your organization to identify the instrument. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.

Note that you can also update an instrument using an upsert through `/instruments/create`.

Follow the links for more information:
  - [Relationships](https://docs.unit21.ai/reference/relationships)
  - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
  - [Custom data](https://docs.unit21.ai/reference/best-practices-for-custom-data)
  - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
  - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)

### Example Usage

```go
package main

import(
	"context"
	"log"
	"unit21"
	"unit21/pkg/models/operations"
	"unit21/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.InstrumentsAPI.UpdateInstrument(ctx, operations.UpdateInstrumentRequest{
        RequestBody: &shared.EntityRelations{
            Entities: []shared.EntityRelationsEntities{
                shared.EntityRelationsEntities{
                    EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
                    EntityType: sdk.String("user"),
                    RelationshipID: sdk.String("friend"),
                },
                shared.EntityRelationsEntities{
                    EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
                    EntityType: sdk.String("user"),
                    RelationshipID: sdk.String("friend"),
                },
                shared.EntityRelationsEntities{
                    EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
                    EntityType: sdk.String("user"),
                    RelationshipID: sdk.String("friend"),
                },
            },
        },
        InstrumentID: "quos",
        OrgName: "sint",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateInstrumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateInstrumentRequest](../../models/operations/updateinstrumentrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateInstrumentResponse](../../models/operations/updateinstrumentresponse.md), error**

