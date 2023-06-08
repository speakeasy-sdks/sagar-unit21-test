# CasesAPI

## Overview

Cases are usually active investigations, which may span multiple events, entities and documents. They can be directly escalated into a suspicious activity report. The `/cases` endpoint can create, list, and update cases. 


### Available Operations

* [CreateCase](#createcase) - Create a case
* [ExportCases](#exportcases) - Bulk export cases
* [GetCaseByUnit21ID](#getcasebyunit21id) - Get a case
* [LinkMediaToCase](#linkmediatocase) - Add media to a case
* [ListCases](#listcases) - List cases
* [UpdateCase](#updatecase) - Update case

## CreateCase

Creates a new case, sending case data in the request body. 
To create a case, you MUST include the following fields: `case_id`, `title`, and `created_at`.  The other top-level fields are optional.

If we receive a request to create a case for an `case_id` that already exists in our system,  we will respond with a **409 error code** indicating that this case cannot be created/updated. You must use the `/case/update` endpoint to update a case.

You can add the following objects to a case:

  | Field                    | Type     | Description                                                             |
  |--------------------------|----------|-------------------------------------------------------------------------|
  | `alerts`	               | Array[]  | Alerts that are associated with this case. Consists of `alert_id`s      |
  | `events`	               | Array[]  | Transactions affiliated with the case. Consists of `event_id`s          |
  | `entities`	             | Array[]  | Entities affiliated with the case. Consists of `entity_id`s             |

Updates to a cases's `case_id` are not allowed.

Follow the links for more information:
  - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
  - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
  - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)


The response will consist of the following fields:

  | Field                    | Type     | Description                                          |
  |--------------------------|----------|------------------------------------------------------|
  | `case_id`	               | String   | 	Unique identifier of the case on your platform     |
  | `previously_existed`	   | Boolean  | 	If case (with the same `case_id`) already exists   |


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
    res, err := s.CasesAPI.CreateCase(ctx, operations.CreateCaseCaseData{
        AlertIds: []int64{
            943749,
            902599,
        },
        CaseID: "CaseA-123",
        Description: sdk.String("suspected money laundering example"),
        Disposition: sdk.String("TRUE_POSITIVE"),
        DispositionNotes: sdk.String("User confirmed that they did not initiate transaction"),
        EndDate: sdk.Int64(1623365011),
        EntityIds: []int64{
            449950,
            359508,
            613064,
        },
        EventIds: []int64{
            902349,
            697631,
        },
        StartDate: 1623365011,
        Status: shared.InvestigationStatusOpen.ToPointer(),
        Tags: []string{
            "Sector:Europe",
        },
        Title: "Active fraud investigation",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCaseResponse != nil {
        // handle response
    }
}
```

## ExportCases

Initiates an email and dashboard export of cases. The export will be as a CSV file.

Either the agent `ID` or `email` is required to begin the export.

Either the `filters` or the list of `case IDs` are required for the export.

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
    res, err := s.CasesAPI.ExportCases(ctx, operations.ExportCasesRequestBody{
        CaseIds: []int64{
            969810,
        },
        Filters: &operations.ExportCasesRequestBodyFilters{
            AgentIds: []int64{
                14,
                14,
                14,
            },
            CaseIds: []int64{
                45,
                45,
                45,
            },
            EndDate: sdk.String("2021-11-05 04:13:46"),
            MaximumAmount: sdk.Int64(100000),
            MinimumAmount: sdk.Int64(1000),
            RuleIds: []int64{
                371,
                371,
                371,
            },
            StartDate: sdk.String("2019-11-05 04:13:46"),
            Status: sdk.String("OPEN"),
            Statuses: []ExportCasesRequestBodyFiltersStatuses{
                operations.ExportCasesRequestBodyFiltersStatusesOpen,
            },
            TagIds: []int64{
                9,
                9,
            },
            TeamIds: []int64{
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

## GetCaseByUnit21ID

Returns all data objects belonging to a single case.

This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the case is first created.

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
    res, err := s.CasesAPI.GetCaseByUnit21ID(ctx, operations.GetCaseByUnit21IDRequest{
        Unit21ID: "nobis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCaseByUnit21ID200ApplicationJSONAny != nil {
        // handle response
    }
}
```

## LinkMediaToCase

Adds rich media objects (images, videos, etc.) to an existing case. 

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

  | media_type                    |
  |-------------------------------|
  | IMAGE_PROFILE_PICTURE         |
  | IMAGE_DRIVERS_LICENSE_FRONT   |
  | IMAGE_DRIVERS_LICENSE_BACK    |
  | IMAGE_PASSPORT_FRONT          |
  | IMAGE_ID_CARD_FRONT           |
  | IMAGE_ID_CARD_BACK            |
  | IMAGE_FACE_IMAGE              |

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
    res, err := s.CasesAPI.LinkMediaToCase(ctx, operations.LinkMediaToCaseRequest{
        RequestBody: &operations.LinkMediaToCaseRequestBody{},
        Unit21ID: "enim",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkMediaToCase200ApplicationJSONAny != nil {
        // handle response
    }
}
```

## ListCases

Returns an array of top-level information about cases in your environment.

Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 50)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record.

To narrow down your case search, we provide filter parameters to this endpoint. Note that all list inputs function as an "or" filter, as in any one of the values must match the selected case(s):


  | Field                   | Type        | Description                                                                                                       |
  | ----------------------- | ----------- | ----------------------------------------------------------------------------------------------------------------- |
  | `created_after`         | Numeric     | Cases created on or after this unix timestamp                                                                     |
  | `created_before`        | Numeric     | Cases created before this unix timestamp                                                                          |
  | `dispositions`          | String[]    | List of case disposition states (defined on an integration basis)                                                 |
  | `dispositioned_after`   | Numeric     | Cases with a disposition most recently updated after this unix timestamp                                          |
  | `dispositioned_before`  | Numeric     | Cases with a disposition most recently updated before this unix timestamp                                         |
  | `dispositioned_by`      | String[]    | List of agent emails. Returns alerts with a disposition most recently changed by agents in the list               |
  | `rules`                 | Numeric[]   | List of Unit21 rule ids that are associated with the case                                                         |
  | `associated_entities`   | Numeric[]   | List of Unit21 entity ids associated with this case                                                               |
  | `associated_events`     | Numeric[]   | List of Unit21 event ids associated with this case                                                                |
  | `associated_alerts`     | Numeric[]   | List of Unit21 alert ids associated with this case                                                                |
  | `sources`               | String[]    | Must be list of alert sources: `INTERNAL`, `EXTERNAL`                                                             |
  | `statuses`              | String[]    | Must be list of alert statuses: `OPEN`, `CLOSED`                                                                  |
  | `tag_filters`           | String[]    | List of string tags (`key:value`) or keys to associate this case with (e.g. `case_type:high_velocity` or `case_type`). If only the key is provided, we will match against all tags with that key        |
  | `limit`                 | Numeric     | A limit on the number of objects to be returned. Limit can range between 1 and 50, and the default is 10          |
  | `offset`                | Numeric     | The offset for pagination. Default is 1                                                                           |
  | `options`               | Object      | Options for the data included in the returned cases. Removing unneeded options can improve response speed         |


The `total_count` field contains the total number of case where the  `response_count` field contains the number of cases included in the response.

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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.CasesAPI.ListCases(ctx, operations.ListCasesRequestBody{
        AssociatedEntities: []int64{
            363711,
            325047,
            570197,
        },
        AssociatedEvents: []int64{
            438601,
        },
        AssociatedInstruments: []int64{
            988374,
            958950,
            102044,
        },
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
        Offset: sdk.Int64(635059),
        Options: &operations.ListCasesRequestBodyOptions{
            IncludeActions: sdk.Bool(false),
            IncludeAssociations: sdk.Bool(false),
        },
        Rules: []int64{
            995300,
        },
        Sources: []shared.SourceArray{
            shared.SourceArrayExternal,
            shared.SourceArrayInternal,
            shared.SourceArrayInternal,
        },
        Statuses: []shared.InvestigationStatus{
            shared.InvestigationStatusOpen,
            shared.InvestigationStatusOpen,
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

## UpdateCase

Update a case through its `unit21_id`. ONLY EXTERNAL CASES CAN BE UPDATED!
    

Updating a case has no required fields. You MAY send any subset of the fields that the `cases/create` endpoint accepts.

This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the case is first created.

Note that you can also update an alert using an upsert through `/cases/create`.

Follow the links for more information:
  - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
  - [Custom data](https://docs.unit21.ai/reference/best-practices-for-custom-data)
  - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
  - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)


The response will consist of the following fields:

  | Field                    | Type     | Description                                          |
  |--------------------------|----------|------------------------------------------------------|
  | `case_id`	               | String   | 	Unique identifier of the case on your platform     |
  | `previously_existed`	   | Boolean  | 	If entity (with the same `case_id`) already exists |

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
    res, err := s.CasesAPI.UpdateCase(ctx, operations.UpdateCaseRequest{
        RequestBody: &shared.Cases{
            CaseID: "CaseA-123",
            Description: sdk.String("suspected money laundering example"),
            Disposition: sdk.String("TRUE_POSITIVE"),
            DispositionNotes: sdk.String("User confirmed that they did not initiate transaction"),
            EndDate: sdk.Int64(1623365011),
            StartDate: 1623365011,
            Status: shared.InvestigationStatusOpen.ToPointer(),
            Tags: []string{
                "Sector:Europe",
            },
            Title: "Active fraud investigation",
        },
        Unit21ID: "error",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateCase200ApplicationJSONAny != nil {
        // handle response
    }
}
```
