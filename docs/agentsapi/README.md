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
            APIKeyAuth: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.DeactivateAgentRequest{
        AgentEmail: "provident",
    }

    res, err := s.AgentsAPI.DeactivateAgent(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeactivateAgent200ApplicationJSONAny != nil {
        // handle response
    }
}
```

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
            APIKeyAuth: "YOUR_API_KEY_HERE",
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
