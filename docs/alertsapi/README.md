# AlertsAPI

## Overview

Alerts have two origins. Typically, alerts are generated whenever a Unit21 detection tool (like a rule) flags an object, like an entity. However, your organization can also send alerts generated from your in-house detection systems. The `/alerts` endpoint can create, list, and update alerts. 


### Available Operations

* [CreateAlert](#createalert) - Create alerts
* [ExportAlerts](#exportalerts) - Bulk export alerts
* [GetAlertByUnit21ID](#getalertbyunit21id) - Get an alert
* [LinkMediaToAlert](#linkmediatoalert) - Add media to an alert
* [ListAlerts](#listalerts) - List alerts
* [UpdateAlert](#updatealert) - Update alert

## CreateAlert

Creates a new alert, sending alert data in the request body. 
To create an Alert, you MUST include the following fields: `alert_id`, `alert_type`, `created_at`, `title`, and `status`. The other top-level fields are optional.

If we receive a request to create an alert for an `alert_id` that already exists in our system,  we will respond with a **409 error code** indicating that this alert cannot be created/updated. You must use the `/alert/update` endpoint to update an alert.

You can add the following objects to an alert:

  | Field                    | Type     | Description                                                                                                           |
  |--------------------------|----------|-----------------------------------------------------------------------------------------------------------------------|
  | `rules`	                 | String[] | Unique identifier of the rules/triggers/scenarios that triggered this alert                                           |
  | `events`	               | Object[] | Transactions affiliated with the alert. Each object must consist of a `event_id` and `event_type` field               |
  | `entities`	             | Object[] | Users or businesses affiliated with the alert. Each object must consist of an `entity_id` and `entity_type` field     |
  | `instruments`	           | String[] | Unique identifiers of any instruments affiliated with the alert                                                       |


Please note that if `verification_result_id` is included, it will link the entity that is associated  with the verification result with the alert.

Updates to an alert's `alert_id` are not allowed.

Follow the links for more information:
  - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
  - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
  - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)


The response will consist of the following fields:

  | Field                    | Type     | Description                                             |
  |--------------------------|----------|---------------------------------------------------------|
  | `alert_id`	             | String   | Unique identifier of the alert on your platform         |
  | `unit21_id`	             | String   | Internal ID of the alert within Unit21's system         |
  | `previously_existed`	   | Boolean  | If alert (with the same `alert_id`) already exists      |


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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AlertsAPI.CreateAlert(ctx, operations.CreateAlertRules{
        AlertID: "alertA-028eb01a-f8d3-42fb-b398-785b596ee4cb",
        AlertType: operations.CreateAlertRulesAlertTypeTm,
        CreatedAt: 1623365011,
        Description: sdk.String("Flagged 2 transactions in last hour that were 3 standard deviations outside 3 month mean."),
        Disposition: sdk.String("TRUE_POSITIVE"),
        DispositionNotes: sdk.String("User confirmed that they did not initiate transaction"),
        Entities: []CreateAlertRulesEntities{
            operations.CreateAlertRulesEntities{
                EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
            },
            operations.CreateAlertRulesEntities{
                EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
            },
            operations.CreateAlertRulesEntities{
                EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
            },
        },
        Events: []CreateAlertRulesEvents{
            operations.CreateAlertRulesEvents{
                EventID: "event-1063e4e3e1",
                EventType: shared.EventTypeTransaction,
            },
            operations.CreateAlertRulesEvents{
                EventID: "event-1063e4e3e1",
                EventType: shared.EventTypeTransaction,
            },
            operations.CreateAlertRulesEvents{
                EventID: "event-1063e4e3e1",
                EventType: shared.EventTypeTransaction,
            },
            operations.CreateAlertRulesEvents{
                EventID: "event-1063e4e3e1",
                EventType: shared.EventTypeTransaction,
            },
        },
        Instruments: []string{
            "3234-sdghfdf-3332",
            "3234-sdghfdf-3332",
            "3234-sdghfdf-3332",
        },
        Rules: []string{
            "r3-0ddfn3",
            "r3-0ddfn3",
            "r3-0ddfn3",
            "r3-0ddfn3",
        },
        Status: shared.InvestigationStatusOpen,
        Tags: []string{
            "Sector:Europe",
            "Sector:Europe",
            "Sector:Europe",
        },
        Title: "Account deviation for user T18029",
        VerificationResultID: sdk.Int64(42),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAlertResponse != nil {
        // handle response
    }
}
```

## ExportAlerts

Initiates an email and dashboard export of alerts. The export will be as a CSV file.

Either the agent `ID` or `email` is required to begin the export.

Either the `filters` or the list of `alert IDs` are required for the export.

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
    res, err := s.AlertsAPI.ExportAlerts(ctx, operations.ExportAlertsRequestBody{
        AlertIds: []int64{
            423655,
            623564,
            645894,
            384382,
        },
        Filters: &operations.ExportAlertsRequestBodyFilters{
            AgentIds: []int64{
                14,
                14,
            },
            AlertIds: []int64{
                45,
                45,
            },
            AlertQueueIds: []int64{
                125,
                125,
                125,
                125,
            },
            AlertTypes: []ExportAlertsRequestBodyFiltersAlertTypes{
                operations.ExportAlertsRequestBodyFiltersAlertTypesCar,
            },
            AlertingOrgIds: []int64{
                45,
                45,
            },
            Disposition: sdk.String("CLOSED"),
            DispositionEndDate: sdk.String("2021-11-05 04:13:46"),
            DispositionStartDate: sdk.String("2019-11-05 04:13:46"),
            EndDate: sdk.String("2021-11-05 04:13:46"),
            EntityIds: []int64{
                15453219,
                15453219,
            },
            MaximumAmount: sdk.Int64(100000),
            MinimumAmount: sdk.Int64(1000),
            Phrase: sdk.String("Case 37765"),
            RuleIds: []int64{
                371,
                371,
            },
            Sources: []ExportAlertsRequestBodyFiltersSources{
                operations.ExportAlertsRequestBodyFiltersSourcesExternal,
                operations.ExportAlertsRequestBodyFiltersSourcesExternal,
                operations.ExportAlertsRequestBodyFiltersSourcesInternal,
                operations.ExportAlertsRequestBodyFiltersSourcesExternal,
            },
            StartDate: sdk.String("2019-11-05 04:13:46"),
            Status: sdk.String("OPEN"),
            Statuses: []ExportAlertsRequestBodyFiltersStatuses{
                operations.ExportAlertsRequestBodyFiltersStatusesClosed,
                operations.ExportAlertsRequestBodyFiltersStatusesClosed,
            },
            SubdispositionIds: []int64{
                2,
            },
            SubdispositionOptionIds: []int64{
                1,
                1,
            },
            TagIds: []int64{
                9,
            },
            TeamIds: []int64{
                6,
                6,
                6,
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

## GetAlertByUnit21ID

Returns all data objects belonging to a single alert.

This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the entity is first created.

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
    res, err := s.AlertsAPI.GetAlertByUnit21ID(ctx, operations.GetAlertByUnit21IDRequest{
        Unit21ID: "perferendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAlertByUnit21ID200ApplicationJSONAny != nil {
        // handle response
    }
}
```

## LinkMediaToAlert

Adds rich media objects (images, videos, etc.) to an existing alert. 

This endpoint is useful for sending in rich media such as profile pictures, ID card scans, official documents etc.  that you want available for investigative and verification purposes.

Supported file types are: txt, pdf, video (mp4, mov, wmv, avi, mkv), images (png, jpg, tiff, gif, raw, eps).

The payload to this endpoint can either be a **form-data** or a **base64** encoded media file via the requests JSON body.

**Form-data** sent to this endpoint must use the key `media_key` and the `value` as the media file.  If you wish to provide optional information, use the `media_key` and provide stringified JSON data as the value.  There are no required fields in each media file's supplementary form data. However, if a recognized `media_type` value is provided,  the Unit21 system will be able to use the media object for purposes such as document verification.

```
    --form 'document_front=@/src/103031/images/document_front.jpg' \
    --form 'document_front={"media_type": "IMAGE_ID_CARD_FRONT", "source": "passport_app", "timestamp": 1572673229}'
```

**Base64** encoded media objects must follow the format:

```json
  {
    "media": "iVBORw0KGgoAAAANSUhEUgAAAQMAAADCCAYAAABNEqduAAAgAElEQVR4Aey9CbgmV1Xv...",
    "name": "Drivers_License.png",
    "media_type": "IMAGE_DRIVERS_LICENSE_FRONT"
  }
```
    
`media` and `name` are the only required fields for each media object. The `name`` must include the file extension such a `File.pdf`. Supplementary form data is sent through the optional `custom_data` object.

Recognized values of `media_type` are: 

  | media_type                  |
  |-----------------------------|
  | IMAGE_PROFILE_PICTURE       |
  | IMAGE_DRIVERS_LICENSE_FRONT |
  | IMAGE_DRIVERS_LICENSE_BACK  |
  | IMAGE_PASSPORT_FRONT        |
  | IMAGE_ID_CARD_FRONT         |
  | IMAGE_ID_CARD_BACK          |
  | IMAGE_FACE_IMAGE            |

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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AlertsAPI.LinkMediaToAlert(ctx, operations.LinkMediaToAlertRequest{
        RequestBody: &operations.LinkMediaToAlertRequestBody{},
        Unit21ID: "ipsam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## ListAlerts

Returns an array of top-level information about alerts in your environment.

Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 50)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record.

To narrow down your alert search, we provide filter parameters to this endpoint. Note that all list inputs function as an "or" filter, as in any one of the values must match the selected alert(s):


  | Field                   | Type        | Description                                                                                                       |
  | ----------------------- | ----------- | ----------------------------------------------------------------------------------------------------------------- |
  | `case_id`               | Numeric     | Only alerts with the associated case ID will be shown.                                                            |
  | `types`                 | String[]    | Must be list of alert types: `tm`, `kyc`, `chainalysis`                                                           |
  | `created_after`         | Numeric     | Alerts created on or after this unix timestamp                                                                    |
  | `created_before`        | Numeric     | Alerts created before this unix timestamp                                                                         |
  | `dispositions`          | String[]    | List of alert disposition states (defined on an integration basis)                                                |
  | `dispositioned_after`   | Numeric     | Alerts with a disposition most recently updated after this unix timestamp                                         |
  | `dispositioned_before`  | Numeric     | Alerts with a disposition most recently updated before this unix timestamp                                        |
  | `dispositioned_by`      | String[]    | List of agent emails. Returns alerts with a disposition most recently changed by agents in the list               |
  | `rules`                 | Numeric[]   | List of Unit21 rule ids that are associated with the alert                                                        |
  | `associated_entities`   | Numeric[]   | List of Unit21 entity ids associated with this alert                                                              |
  | `associated_events`     | Numeric[]   | List of Unit21 event ids associated with this alert                                                               |
  | `associated_instruments`| Numeric[]   | List of Unit21 instrument ids associated with this alert                                                          |
  | `sources`               | String[]    | Must be list of alert sources: `INTERNAL`, `EXTERNAL`                                                             |
  | `statuses`              | String[]    | Must be list of alert statuses: `OPEN`, `CLOSED`                                                                  |
  | `tag_filters`           | String[]    | List of string tags (`key:value`) or keys to associate this alert with (e.g. `alert_type:high_velocity` or `alert_type`). If only the key is provided, we will match against all tags with that key        |
  | `limit`                 | Numeric     | A limit on the number of objects to be returned. Limit can range between 1 and 50, and the default is 10          |
  | `offset`                | Numeric     | The offset for pagination. Default is 1                                                                           |
  | `options`               | Object      | Options for the data included in the returned alerts. Removing unneeded options can improve response speed        |


The `total_count` field contains the total number of alerts where the  `response_count` field contains the number of alerts included in the response.

Follow the links for more information:
  - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
  


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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AlertsAPI.ListAlerts(ctx, operations.ListAlertsRequestBody{
        AssociatedEntities: []int64{
            957156,
            778157,
            140350,
            870013,
        },
        AssociatedEvents: []int64{
            978619,
            473608,
            799159,
            800911,
        },
        AssociatedInstruments: []int64{
            520478,
            780529,
        },
        CaseID: sdk.Int64(10107269),
        CreatedAfter: sdk.Int64(1623365011),
        CreatedBefore: sdk.Int64(1623365011),
        DispositionedAfter: sdk.Int64(1623365011),
        DispositionedBefore: sdk.Int64(1623365011),
        DispositionedBy: []string{
            "agent1@example.com",
            "agent1@example.com",
            "agent1@example.com",
        },
        Dispositions: []string{
            "TRUE_POSITIVE",
        },
        Limit: sdk.Int64(2),
        Offset: sdk.Int64(720633),
        Options: &operations.ListAlertsRequestBodyOptions{
            IncludeActions: sdk.Bool(false),
            IncludeAssociations: sdk.Bool(false),
        },
        Rules: []int64{
            582020,
            143353,
            537373,
        },
        Sources: []shared.SourceArray{
            shared.SourceArrayExternal,
            shared.SourceArrayExternal,
            shared.SourceArrayInternal,
            shared.SourceArrayInternal,
        },
        Statuses: []shared.InvestigationStatus{
            shared.InvestigationStatusOpen,
            shared.InvestigationStatusOpen,
        },
        TagFilters: []string{
            "Sector:Europe",
            "Sector:Europe",
        },
        Types: []ListAlertsRequestBodyTypes{
            operations.ListAlertsRequestBodyTypesChainalysis,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListResponse != nil {
        // handle response
    }
}
```

## UpdateAlert

Updates an alert's information using the `unit21_id`. ONLY EXTERNAL ALERTS CAN BE UPDATED!

Updating an alert has no required fields. You MAY send any subset of the fields that the `alerts/create` endpoint accepts.

This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the entity is first created.

Please note that if `verification_result_id` is included, it will link the entity that is associated with the  verification result with the alert.

Follow the links for more information:
  - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
  - [Custom data](https://docs.unit21.ai/reference/best-practices-for-custom-data)
  - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
  - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)


The response will consist of the following fields:

  | Field                    | Type     | Description                                             |
  |--------------------------|----------|---------------------------------------------------------|
  | `alert_id`	             | String   | Unique identifier of the alert on your platform         |
  | `unit21_id`	             | String   | Internal ID of the alert within Unit21's system         |
  | `previously_existed`	   | Boolean  | If alert (with the same `alert_id`) already exists      |

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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AlertsAPI.UpdateAlert(ctx, operations.UpdateAlertRequest{
        RequestBody: &shared.CustomData{
            CustomData: map[string]interface{}{
                "ipsum": "excepturi",
                "aspernatur": "perferendis",
            },
        },
        Unit21ID: "ad",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateAlert200ApplicationJSONAny != nil {
        // handle response
    }
}
```
