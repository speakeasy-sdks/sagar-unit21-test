# EventsAPI

## Overview

Events have two types, _transaction events_  and _action events_:
* Transaction events are any monetary flow that is sent or received by an entity on your system.
* Action events are non-monetary changes of state that occur on your system, e.g. user logins. The `/events` endpoint sends and receives data about significant actions that occur with an entity or instrument on your system.  


### Available Operations

* [CreateEvent](#createevent) - Create an event
* [ExportEvents](#exportevents) - Bulk export events
* [ExportTransactions](#exporttransactions) - Bulk export transactions
* [GetEvent](#getevent) - Get an event
* [ListEvents](#listevents) - List events
* [UpdateEvent](#updateevent) - Update event

## CreateEvent

Creates a new event, sending event data in the request body. 

Two objects are required: `general_data` and either `transaction_data` or `action_data`. `general_data` requires the fields: `event_id`, `event_type`, and `event_time`. `transaction_data` requires only the field `amount`.

Unlike entities, events on our system are cannot be explicitly updated. However, they can be  overwritten in a naive upsert-overwrite fashion. 

If we receive a request to create an event for an `event_id` that already exists in our system,  we will simply overwrite that previous entry with the newly provided data if this transaction  is not already associated with other articles in our system. 

For instance, if a transaction is flagged in an alert and we receive a request to overwrite  the details of this transaction, we will respond with a **409 error code** indicating that this  event cannot be overwritten.

Updates to an event's `general_data.event_id` are not allowed.

Follow the links for more information:
  - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
  - [Verification options](https://docs.unit21.ai/reference/identity-verification-options)
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
    res, err := s.EventsAPI.CreateEvent(ctx, operations.CreateEventEventOptions{
        ActionData: &operations.CreateEventEventOptionsActionData{
            ActionDetails: sdk.String("Through mobile app v0.8.8"),
            ActionType: sdk.String("LOGIN"),
            EntityID: "371c4d7b-0563-4685-aab1",
            EntityType: "user",
            InstrumentID: sdk.String("238938823-34347686-39405443"),
        },
        CustomData: map[string]interface{}{
            "delectus": "eum",
            "non": "eligendi",
        },
        DigitalData: &operations.CreateEventEventOptionsDigitalData{
            IPAddress: sdk.String("255.255.255.255"),
        },
        GeneralData: &operations.CreateEventEventOptionsGeneralData{
            EventID: "t-9daaebad-478d-4ea3-bbf9-e6320d3f1cea",
            EventSubtype: sdk.String("payment"),
            EventTime: 1623365011,
            EventType: shared.EventTypeTransaction,
            Parents: []CreateEventEventOptionsGeneralDataParents{
                operations.CreateEventEventOptionsGeneralDataParents{
                    EventID: sdk.String("t-9daaebad-478d-4ea3-bbf9-e6320d3f1cea"),
                    EventType: shared.EventTypeTransaction.ToPointer(),
                },
                operations.CreateEventEventOptionsGeneralDataParents{
                    EventID: sdk.String("t-9daaebad-478d-4ea3-bbf9-e6320d3f1cea"),
                    EventType: shared.EventTypeTransaction.ToPointer(),
                },
                operations.CreateEventEventOptionsGeneralDataParents{
                    EventID: sdk.String("t-9daaebad-478d-4ea3-bbf9-e6320d3f1cea"),
                    EventType: shared.EventTypeTransaction.ToPointer(),
                },
            },
            Status: sdk.String("active"),
            Tags: []string{
                "Sector:Europe",
                "Sector:Europe",
            },
        },
        LocationData: &shared.LocationDataProperties{
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
        Options: &operations.CreateEventEventOptionsOptions{
            LinkDigitalDataToEntity: sdk.Bool(true),
            LinkedEntity: operations.CreateEventEventOptionsOptionsLinkedEntitySender.ToPointer(),
            MergeCustomData: sdk.Bool(false),
            Monitor: sdk.Bool(false),
            ResolveGeoip: sdk.Bool(true),
            UpsertOnConflict: sdk.Bool(true),
        },
        TransactionData: &operations.CreateEventEventOptionsTransactionData{
            Amount: 13562.83,
            ExchangeRate: sdk.Float64(1.0012),
            ExternalFee: sdk.Float64(5),
            InternalFee: sdk.Float64(5),
            ReceivedAmount: sdk.Float64(13562.83),
            ReceivedCurrency: sdk.String("USD"),
            ReceiverEntityID: sdk.String("u-d8e1d453-c205-4996-a935-ff08be01bcd1"),
            ReceiverEntityType: sdk.String("business"),
            ReceiverInstrumentID: sdk.String("t376839428256371"),
            SenderEntityID: sdk.String("u-371b5091-da30-48a8-936b-7600f9983e80"),
            SenderEntityType: sdk.String("user"),
            SenderInstrumentID: sdk.String("instrumentA-3967112"),
            SentAmount: sdk.Float64(13562.83),
            SentCurrency: sdk.String("USD"),
            TransactionHash: sdk.String("af830da0919f9d3ebbc413040460708c4107e77c83c1d1a56c6bc76a48d"),
            UsdConversionNotes: sdk.String("pulled from forex.com at 1572672226"),
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

## ExportEvents

Initiates an email and dashboard export of events. The export will be as a CSV file.

Either the agent `ID` or `email` is required to begin the export.

Either the `filters` or the list of `event IDs` are required for the export.

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
    res, err := s.EventsAPI.ExportEvents(ctx, operations.ExportEventsRequestBody{
        EventIds: []int64{
            896039,
            572252,
            638921,
        },
        Filters: &operations.ExportEventsRequestBodyFilters{
            Currency: sdk.String("USD"),
            EndDate: sdk.String("2021-11-05 04:13:46"),
            EntityIds: []int64{
                85,
            },
            MaximumAmount: sdk.Int64(100000),
            MinimumAmount: sdk.Int64(1000),
            StartDate: sdk.String("2019-11-05 04:13:46"),
            Status: sdk.String("active"),
            Statuses: []string{
                "new",
                "new",
                "new",
                "new",
            },
            TagIds: []int64{
                9,
                9,
                9,
                9,
            },
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

## ExportTransactions

Initiates an email and dashboard export of events. The export will be as a CSV file.

Either the agent `ID` or `email` is required to begin the export.

Either the `filters` or the list of `event IDs` are required for the export.

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
    res, err := s.EventsAPI.ExportTransactions(ctx, operations.ExportTransactionsRequestBody{
        EventIds: []int64{
            447125,
            449198,
            846409,
        },
        Filters: &operations.ExportTransactionsRequestBodyFilters{
            Currency: sdk.String("USD"),
            EndDate: sdk.String("2021-11-05 04:13:46"),
            EntityIds: []int64{
                85,
                85,
                85,
                85,
            },
            MaximumAmount: sdk.Int64(100000),
            MinimumAmount: sdk.Int64(1000),
            StartDate: sdk.String("2019-11-05 04:13:46"),
            Status: sdk.String("active"),
            Statuses: []string{
                "new",
                "new",
                "new",
            },
            TagIds: []int64{
                9,
            },
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

## GetEvent

Returns all data objects belonging to a single event.

This endpoint requires the `events_id` which is a unique ID created by your organization to identify the event. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.

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
    res, err := s.EventsAPI.GetEvent(ctx, operations.GetEventRequest{
        EventID: "magnam",
        OrgName: "cumque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEvent200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## ListEvents

Returns an array of top-level information about events in your environment.

Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 50)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record.
* `case_id`  is a filter. Only events with the associated case ID will be shown.
* `alert_id` is a filter. Only events with the associated alert ID will be shown.
* `start_date`  is a filter. Only events that started on or after this date will be shown.
* `end_date` is a filter. Only events that ended on or before this date will be shown.

The `total_count` field contains the total number of events where the  `response_count` field contains the number of events included in the response.


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
    res, err := s.EventsAPI.ListEvents(ctx, shared.ListDateRequest{
        AlertID: sdk.Int64(73913),
        CaseID: sdk.Int64(10107269),
        EndDate: sdk.String("2021-11-05T04:13:46.000Z"),
        Limit: sdk.Int64(2),
        Offset: sdk.Int64(813798),
        StartDate: sdk.String("2019-11-05T04:13:46.000Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListResponse != nil {
        // handle response
    }
}
```

## UpdateEvent

Update an event using the `event_id` from your platform. 

Updating an event has no required fields. You MAY send any  subset of the fields that the events/create endpoint accepts.

This endpoint requires the `event_id` which is a unique ID created by your organization to identify the event. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.

Note that you can also update an event using an upsert through `/events/create`.

Follow the links for more information:
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
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.EventsAPI.UpdateEvent(ctx, operations.UpdateEventRequest{
        RequestBody: &operations.UpdateEventEventOptions{
            Options: &operations.UpdateEventEventOptionsOptions{
                LinkDigitalDataToEntity: sdk.Bool(true),
                LinkedEntity: operations.UpdateEventEventOptionsOptionsLinkedEntitySender.ToPointer(),
                MergeCustomData: sdk.Bool(false),
                Monitor: sdk.Bool(false),
                ResolveGeoip: sdk.Bool(true),
                UpsertOnConflict: sdk.Bool(true),
            },
        },
        EventID: "ea",
        OrgName: "aliquid",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
