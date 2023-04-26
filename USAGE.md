<!-- Start SDK Example Usage -->
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
        AgentEmail: "corrupti",
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
<!-- End SDK Example Usage -->