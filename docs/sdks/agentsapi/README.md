# AgentsAPI

## Overview

Agents are your organization's members who use the Unit21 system to investigate suspicious objects and events. The `/agents` endpoint can list your agents. 


### Available Operations

* [DeactivateAgent](#deactivateagent) - Deactivate an agent
* [ListAgents](#listagents) - List agents

## DeactivateAgent

Archives an agent so that he/she can no longer log in to the dashboard.

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
    res, err := s.AgentsAPI.DeactivateAgent(ctx, operations.DeactivateAgentRequest{
        AgentEmail: "provident",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeactivateAgent200ApplicationJSONAny != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeactivateAgentRequest](../../models/operations/deactivateagentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeactivateAgentResponse](../../models/operations/deactivateagentresponse.md), error**


## ListAgents

Returns an array of all agents in your organization who are using the environment.
There are no options or filters for this endpoint. The request will return ALL agents.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"unit21"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AgentsAPI.ListAgents(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListAgentsResponse](../../models/operations/listagentsresponse.md), error**

