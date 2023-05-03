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
    res, err := s.AgentsAPI.DeactivateAgent(ctx, operations.DeactivateAgentRequest{
        AgentEmail: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeactivateAgent200ApplicationJSONAny != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->