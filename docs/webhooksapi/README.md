# WebhooksAPI

## Overview

Whenever an event happens on the Unit21 platform, Unit21 can send a webhook about the event to whatever URL you configure. Such events include entity verification results, generated alerts, case re-openings and closings, etcetera.


### Available Operations

* [UpdateWebhook](#updatewebhook) - Update webhook URL

## UpdateWebhook

Change the URL of an existing webhook from Unit21.

This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the webhook is first created.

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
    res, err := s.WebhooksAPI.UpdateWebhook(ctx, operations.UpdateWebhookRequest{
        RequestBody: &operations.UpdateWebhookRequestBody{
            URL: sdk.String("https://example.com"),
        },
        Unit21ID: "accusantium",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
