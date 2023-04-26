# EntitiesAPI

## Overview

_Entities_ are typically businesses or users that have transactions on your platform.  The `/entities` endpoint can create, list, and update entities.


### Available Operations

* [AddInstruments](#addinstruments) - Add instruments to entity
* [CreateEntity](#createentity) - Create an entity
* [DelMediaEntity](#delmediaentity) - Delete entity media
* [ExportEntities](#exportentities) - Bulk export entities
* [GetEntity](#getentity) - Get an entity
* [LinkMediaToEntity](#linkmediatoentity) - Add media to an entity
* [ListEntities](#listentities) - List entities
* [UpdateEntity](#updateentity) - Update entity

## AddInstruments

Associate an entity with an array of instruments.

Specify the `instrument_id` of the instrument when associating instruments.

If we do not find any instruments with a corresponding `instrument_id` in our system, we will create a [placeholder](https://docs.unit21.ai/reference/placeholder-objects) for it.

Instrument details can then be supplemented through the `/instruments/create` or `/instruments/update` endpoints.

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
    req := operations.AddInstrumentsRequest{
        EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
        LinkInstrument: &shared.LinkInstrument{
            InstrumentIds: []string{
                "reiciendis",
            },
        },
        OrgName: "voluptatibus",
    }

    res, err := s.EntitiesAPI.AddInstruments(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateEntity

Creates a new entity, sending entity data in the request body. 

If the `/entities/create` API is called for an entity that already exists in our system, it is treated it as an  [upsert](https://docs.unit21.ai/reference/should-i-update-or-upsert) and an update on the existing entity is performed. The response to the request will then contain the entry `previously_existed: true`. 

Unit21 selectively ignores upserts if the request is identical to a previous request. The response to any  ignored upsert will contain the field `ignored: true`. 

If you want to perform strict validation and not perform an upsert on conflict, specifying  `options.upsert_on_conflict: false` will result in the API responding with a **409 error code** indicating  that this entity cannot be overwritten.

Updates to an entity's `general_data.entity_id` are not allowed.

Instruments can be associated with entities by providing the IDs of these  instruments within the `instrument_ids` section of the request. If the instrument doesn't already exist,  Unit21 creates a [placeholder](https://docs.unit21.ai/reference/placeholder-objects) instrument.

Whitelisted entities cannot be updated through the `/entities/create` endpoint. 

We recommend that you create entities prior to running verification. In the event you wish to run a  verification on an entity immediately, Unit21 recommends that you wait at-least 2 minutes for your entity  data to be fully processed. You will receive a **423 error code** if an entity is *busy*. When a 200 response is received,  the data has been successfully stored on the Unit21 backend; however, it may take a few additional seconds to process that data  so that it becomes available in subsequent API calls, in the frontend UI, and/or for verification purposes.

Follow the links for more information:
  - [Relationships](https://docs.unit21.ai/reference/relationships)
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
	"unit21/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.CreateEntityRequest{
        BusinessData: &shared.BusinessDataProperties{
            AccountHolderName: sdk.String("John Smith"),
            BusinessName: sdk.String("Acme"),
            CorporateTaxID: sdk.String("434-455-3166"),
            DoingBusinessAs: sdk.String("Global Liquids"),
            RegisteredCountry: sdk.String("US"),
            RegisteredState: sdk.String("CA"),
            Website: sdk.String("www.google.com"),
        },
        CommunicationData: &shared.CommunicationDataProperties{
            EmailAddresses: []string{
                "JohnJay@example.com",
                "JohnJay@example.com",
                "JohnJay@example.com",
                "JohnJay@example.com",
            },
            PhoneNumbers: []string{
                "+14159627132",
                "+14159627132",
            },
        },
        CustomData: map[string]interface{}{
            "voluptatibus": "ipsa",
            "omnis": "voluptate",
            "cum": "perferendis",
        },
        DigitalData: &shared.DigitalDataArrayProperties{
            ClientFingerprints: []string{
                "5454-ahfd-4531d-f4il",
            },
            IPAddresses: []string{
                "255.255.255.255",
                "255.255.255.255",
            },
        },
        DocumentData: []shared.DocumentDataProperties{
            shared.DocumentDataProperties{
                Country: sdk.String("US"),
                DocumentID: "G3352403F",
                DocumentType: "drivers_license",
                ExpiresAt: sdk.Int64(1572673227),
                IssuedAt: sdk.Int64(1572673226),
                State: sdk.String("CA"),
                VerifiedOn: sdk.Int64(1572673227),
            },
            shared.DocumentDataProperties{
                Country: sdk.String("US"),
                DocumentID: "G3352403F",
                DocumentType: "drivers_license",
                ExpiresAt: sdk.Int64(1572673227),
                IssuedAt: sdk.Int64(1572673226),
                State: sdk.String("CA"),
                VerifiedOn: sdk.Int64(1572673227),
            },
        },
        GeneralData: shared.GeneralEntitiesProperties{
            EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
            EntitySubtype: sdk.String("contractor"),
            EntityType: sdk.String("user"),
            Parents: []shared.GeneralEntitiesPropertiesParents{
                shared.GeneralEntitiesPropertiesParents{
                    EntityID: sdk.String("entity-5500"),
                    EntityType: sdk.String("user"),
                },
                shared.GeneralEntitiesPropertiesParents{
                    EntityID: sdk.String("entity-5500"),
                    EntityType: sdk.String("user"),
                },
                shared.GeneralEntitiesPropertiesParents{
                    EntityID: sdk.String("entity-5500"),
                    EntityType: sdk.String("user"),
                },
                shared.GeneralEntitiesPropertiesParents{
                    EntityID: sdk.String("entity-5500"),
                    EntityType: sdk.String("user"),
                },
            },
            RegisteredAt: sdk.Int64(1623365011),
            Status: sdk.String("active"),
            Tags: []string{
                "Sector:Europe",
            },
        },
        InstrumentIds: []string{
            "3234-sdghfdf-3332",
            "3234-sdghfdf-3332",
        },
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
        },
        Options: &shared.EntityOptionsProperties{
            IdentityVerifications: &shared.EntityOptionsPropertiesIdentityVerifications{
                IncludeFullResponse: sdk.Bool(false),
                RunVerifications: sdk.Bool(false),
                SocureOverrideAPIKey: sdk.String("4543h90fdm3-02r8fgdgfd-78gf7tdg"),
                SynchronousResponse: sdk.Bool(false),
                WorkflowID: sdk.String("sanctions_check_1"),
            },
            MergeCustomData: sdk.Bool(false),
            ResolveGeoip: sdk.Bool(true),
            UpsertOnConflict: sdk.Bool(true),
        },
        Relationships: &shared.RelationshipsDataProperties{
            Children: []shared.RelatedEntity{
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
            Instruments: []shared.RelatedInstrument{
                shared.RelatedInstrument{
                    InstrumentID: sdk.String("3234-sdghfdf-3332"),
                    RelationshipID: sdk.String("friend"),
                },
            },
            Parents: []shared.RelatedEntity{
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
            Undirected: []shared.RelatedEntity{
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
        },
        RiskScores: []shared.RiskScoresProperties{
            shared.RiskScoresProperties{
                Denominator: sdk.Int64(100),
                Name: "risk_score1",
                Score: 65,
            },
            shared.RiskScoresProperties{
                Denominator: sdk.Int64(100),
                Name: "risk_score1",
                Score: 65,
            },
            shared.RiskScoresProperties{
                Denominator: sdk.Int64(100),
                Name: "risk_score1",
                Score: 65,
            },
            shared.RiskScoresProperties{
                Denominator: sdk.Int64(100),
                Name: "risk_score1",
                Score: 65,
            },
        },
        UserData: &shared.UserDataProperties{
            DayOfBirth: sdk.Int64(23),
            FirstName: sdk.String("John"),
            Gender: shared.UserDataPropertiesGenderEnumFemale.ToPointer(),
            LastName: sdk.String("Smith"),
            MiddleName: sdk.String("Joseph"),
            MonthOfBirth: sdk.Int64(12),
            Ssn: sdk.String("733-99-5921"),
            YearOfBirth: sdk.Int64(1990),
        },
    }

    res, err := s.EntitiesAPI.CreateEntity(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateEntityResponse != nil {
        // handle response
    }
}
```

## DelMediaEntity

Deletes rich media objects (images, videos, etc.) to an existing entity. 

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
    req := operations.DelMediaEntityRequest{
        EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
        OrgName: "commodi",
    }

    res, err := s.EntitiesAPI.DelMediaEntity(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## ExportEntities

Initiates an email and dashboard export of entities. The export will be as a CSV file.

Either the agent `ID` or `email` is required to begin the export.

Either the `filters` or the list of `entity IDs` are required for the export.

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
    req := operations.ExportEntitiesRequestBody{
        EntityIds: []int64{
            64147,
            216822,
            692472,
            565189,
        },
        Filters: &operations.ExportEntitiesRequestBodyFilters{
            ChildOrgIds: []int64{
                85,
                85,
                85,
            },
            EndDate: sdk.String("2021-11-05 04:13:46"),
            EntityIds: []int64{
                1894,
                1894,
                1894,
                1894,
            },
            EntityType: sdk.String("user"),
            InternalEntityType: []string{
                "employer",
                "employer",
            },
            RuleIds: []int64{
                371,
                371,
                371,
            },
            StartDate: sdk.String("2019-11-05 04:13:46"),
            Status: sdk.String("active"),
            Statuses: []ExportEntitiesRequestBodyFiltersStatusesEnum{
                operations.ExportEntitiesRequestBodyFiltersStatusesEnumInactive,
                operations.ExportEntitiesRequestBodyFiltersStatusesEnumActive,
                operations.ExportEntitiesRequestBodyFiltersStatusesEnumInactive,
            },
            TagIds: []int64{
                9,
                9,
                9,
            },
        },
    }

    res, err := s.EntitiesAPI.ExportEntities(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.MessageResponse != nil {
        // handle response
    }
}
```

## GetEntity

Returns all data objects belonging to a single entity, including `general_data`, `document_data`, etc.

This endpoint requires the `entity_id` which is a unique ID created by your organization to identify the entity. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.

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
    req := operations.GetEntityRequest{
        EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
        OrgName: "veritatis",
    }

    res, err := s.EntitiesAPI.GetEntity(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.EntityList != nil {
        // handle response
    }
}
```

## LinkMediaToEntity

Adds rich media objects (images, videos, etc.) to an existing entity. 

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
    "media_type": "IMAGE_DRIVERS_LICENSE_FRONT",
    "custom_data": {
      "internal_notes": "Reviewed by Mitchell on 31 June 2019",
      "reviewers": 3,
      "login": 1638384860,
      "timestamp": "2012-03-40 05:12:41.000Z",
      "daily_email": true,
      "employees": ["John", "Anna", "Peter"],
      "socure_device_session_id": "12121212121212112"
    }
  }
```
    
`media` and `name` are the only required fields for each media object. The `name`` must include the file extension such a `File.pdf`. Supplementary form data is sent through the optional `custom_data` object.

For verification purposes, recognized values of `media_type` are: 


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
    req := operations.LinkMediaToEntityRequest{
        RequestBody: &operations.LinkMediaToEntityRequestBody{},
        EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
        OrgName: "itaque",
    }

    res, err := s.EntitiesAPI.LinkMediaToEntity(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## ListEntities

Returns paginated list of of top-level information about entities. 

NOTICE: Entity Type is will be optional soon.

Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 50)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record. The offset is relative to the number of pages (not the total count of objects).
* `case_id`  is a filter. Only entities with the associated case ID will be shown.
* `alert_id` is a filter. Only entities with the associated alert ID will be shown.

The `total_count` field contains the total number of entities where the  `response_count` field contains the number of entities included in the response.

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
    req := shared.ListEntityRequest{
        AlertID: sdk.Int64(73913),
        CaseID: sdk.Int64(10107269),
        Limit: 2,
        Offset: 277718,
    }

    res, err := s.EntitiesAPI.ListEntities(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListEntityResponse != nil {
        // handle response
    }
}
```

## UpdateEntity

Updates an entity's information using the `entity_id` from your platform. 

Updating an entity has no required fields. You MAY send any  subset of the fields that the entities/create endpoint accepts.

This endpoint requires the `entity_id` which is a unique ID created by your organization to identify the entity. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.

Note that you can also update an entity using an upsert through `/entities/create`.

Follow the links for more information:
  - [Relationships](https://docs.unit21.ai/reference/relationships)
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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.UpdateEntityRequest{
        RequestBody: shared.DocumentDataArray{
            DocumentData: []shared.DocumentDataProperties{
                shared.DocumentDataProperties{
                    Country: sdk.String("US"),
                    DocumentID: "G3352403F",
                    DocumentType: "drivers_license",
                    ExpiresAt: sdk.Int64(1572673227),
                    IssuedAt: sdk.Int64(1572673226),
                    State: sdk.String("CA"),
                    VerifiedOn: sdk.Int64(1572673227),
                },
            },
        },
        EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
        OrgName: "est",
    }

    res, err := s.EntitiesAPI.UpdateEntity(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateEntityResponse != nil {
        // handle response
    }
}
```
