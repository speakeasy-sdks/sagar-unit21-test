# VerificationFormsAPI

## Overview

With verification forms, you can automate ID verification and user collection. To gather user input, the `verification forms` endpoint creates a URL. This URL is only valid for a specified period of time.


### Available Operations

* [CreateVerificationForm](#createverificationform) - Verification Forms API

## CreateVerificationForm

If you are verifying IDs and collecting user data, this endpoint creates a temporary URL to which you can redirect users.

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
    res, err := s.VerificationFormsAPI.CreateVerificationForm(ctx, operations.CreateVerificationFormRequestBody{
        SessionLengthMinutes: sdk.Int64(868126),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateVerificationForm200ApplicationJSONObject != nil {
        // handle response
    }
}
```
