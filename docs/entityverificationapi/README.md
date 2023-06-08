# EntityVerificationAPI

## Overview

Unit21 can execute entity verifications according to steps defined in system workflows. They can run when a new entity is added to the system, an update is made to an existing entity, or by directly triggering an API endpoint.


### Available Operations

* [AddVerificationResultToEntity](#addverificationresulttoentity) - Link external verification
* [GetEntityVerificationWorkflowExecutions](#getentityverificationworkflowexecutions) - Get entity verification workflow IDs
* [GetVerificationResult](#getverificationresult) - Get verification results by result id
* [GetVerificationResultFromWorkflowExecution](#getverificationresultfromworkflowexecution) - Get verification results from workflow
* [GetVerificationWorkflowExecution](#getverificationworkflowexecution) - Get verification workflow execution details
* [RunVerificationsWorkflowThroughExternalID](#runverificationsworkflowthroughexternalid) - Verify an entity
* [UpdateContinuousMonitoring](#updatecontinuousmonitoring) - Update continuous monitoring
* [UpdateSuppressedProviderEntities](#updatesuppressedproviderentities) - Suppress provider entity

## AddVerificationResultToEntity

Add the verification result from an external ID provider to an entity on the Unit21 system.
You can only send 1 result per request.

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
    res, err := s.EntityVerificationAPI.AddVerificationResultToEntity(ctx, operations.AddVerificationResultToEntityRequest{
        LinkVerificationResult: &shared.LinkVerificationResult{
            Content: map[string]interface{}{
                "explicabo": "deserunt",
                "distinctio": "quibusdam",
                "labore": "modi",
                "qui": "aliquid",
            },
            ProviderName: "FAKE_PROVIDER",
            VerificationTimestamp: sdk.Int64(1623365011),
            VerificationType: shared.LinkVerificationResultVerificationTypeWatchlistScreening,
        },
        Unit21ID: "quos",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkVerificationResponse != nil {
        // handle response
    }
}
```

## GetEntityVerificationWorkflowExecutions

Returns the verification workflow IDs for an entity.

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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.EntityVerificationAPI.GetEntityVerificationWorkflowExecutions(ctx, operations.GetEntityVerificationWorkflowExecutionsRequest{
        Unit21ID: "perferendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetVerificationResult

Returns all the information from the verification of a specific entity.

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
    res, err := s.EntityVerificationAPI.GetVerificationResult(ctx, operations.GetVerificationResultRequest{
        ResultID: "magni",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetVerificationResultFromWorkflowExecution

Returns all the information from the verification workflow execution for a specific entity.

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
    res, err := s.EntityVerificationAPI.GetVerificationResultFromWorkflowExecution(ctx, operations.GetVerificationResultFromWorkflowExecutionRequest{
        VerificationWorkflowExecutionID: 828940,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetVerificationWorkflowExecution

Returns all the data associated with a verification_workflow_execution_id

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
    res, err := s.EntityVerificationAPI.GetVerificationWorkflowExecution(ctx, operations.GetVerificationWorkflowExecutionRequest{
        VerificationWorkflowExecutionID: 369808,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## RunVerificationsWorkflowThroughExternalID

Run a verification workflow on an entity using the `entity_id` from your platform. 

Requires a `workflow_id`. You can create a verification workflow from the Unit21 dashboard.

This endpoint requires the `entity_id` which is a unique ID created by your organization to identify the entity. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.


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
    res, err := s.EntityVerificationAPI.RunVerificationsWorkflowThroughExternalID(ctx, operations.RunVerificationsWorkflowThroughExternalIDRequest{
        EntityVerification: shared.EntityVerification{
            IncludeFullResponse: sdk.Bool(false),
            SynchronousResponse: sdk.Bool(false),
            WorkflowID: "sanctions_check_1",
        },
        EntityID: "u-3593dece-6642-4cdc-8547-aafc1454e0a0",
        OrgName: "alias",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RunVerificationsWorkflowThroughExternalID200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## UpdateContinuousMonitoring

Fetch status and enables/disables Socure continuous monitoring for an entity.

For asynchronous continuous monitoring, the endpoint will always return a 200 success status response.

For synchronous continous monitoring, the endpoint will always return a 200 success status response  but you should look at the `is_success = true` field to check if the result was actually successful: 

`
  {
    "error_message": "This entity has no existing continuous monitoring subscriptions to disable.",
    "is_success": true
  }
`


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
    res, err := s.EntityVerificationAPI.UpdateContinuousMonitoring(ctx, operations.UpdateContinuousMonitoringRequest{
        ContinuousMonitoring: &shared.ContinuousMonitoring{
            ContinuousMonitoring: false,
            SynchronousResponse: sdk.Bool(false),
        },
        Unit21ID: "fugit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## UpdateSuppressedProviderEntities

Mute Socure continuous monitoring for an entity. 1 - Suppress 0 - Unsuppress


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
    res, err := s.EntityVerificationAPI.UpdateSuppressedProviderEntities(ctx, operations.UpdateSuppressedProviderEntitiesRequest{
        SuppressProviderEntity: &shared.SuppressProviderEntity{
            ProviderEntityIds: []string{
                "excepturi",
                "tempora",
                "facilis",
            },
            Suppress: false,
            SynchronousResponse: sdk.Bool(false),
        },
        Unit21ID: "tempore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
