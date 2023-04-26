# BlacklistsAPI

## Overview

Blacklists comprise one of the following categories:
  * of entities (users or business)
  * IPs (single or ranges)
  * strings


### Available Operations

* [AddBlacklistValues](#addblacklistvalues) - Add items to a blacklist
* [CreateBlacklist](#createblacklist) - Create a blacklist
* [ListBlacklists](#listblacklists) - List blacklists

## AddBlacklistValues

Add items to a blacklist, according to the blacklist's `type`.

Each request must specify at least **1** object to blacklist. You may add up to **100**  values to an existing blacklist at once.

The `/blacklists/<blacklist-id>/add-values` API will ignore entries provided that already exist  in the blacklist. No error will be thrown when this occurs.

The response will consist of the following fields:

  | Type       | Description                                                              | Example                           |
  |------------|--------------------------------------------------------------------------|-----------------------------------|
  | `STRING`	 | Plain strings to match against any text-type field.                      | 		"blacklist_value": "abcde"    |
  | `IP_INET`	 | IPv4 or IPv6 IP addresses to blacklist.                                  | 	"ip_address": "255.255.255.255" |
  | `IP_CIDR`	 | Classless Inter-Domain Routing (CIDR) [notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) IP address ranges to blacklist.  | 	"cidr": "255.255.255.255/32" |
  | `USER`	   | 	Series of fields that a Unit21 user entity will be matched against.     | 	user_data object                |
  | `BUSINESS` | Series of fields that a Unit21 business entity will be matched against.  | 	business_data object            |


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
    req := operations.AddBlacklistValuesRequest{
        RequestBody: []shared.BlacklistCIDR{
            shared.BlacklistCIDR{
                Cidr: "255.255.255.255/24",
                Source: sdk.String("USA"),
            },
        },
        Unit21ID: "iste",
    }

    res, err := s.BlacklistsAPI.AddBlacklistValues(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateBlacklist

Create a new blacklist sending blacklist data in the request body. 

Unit21 currently supports 5 types of blacklists:

  * `STRING`: Plain strings to match against any text-type field.
  * `IP_INET`: IPv4 or IPv6 IP addresses to blacklist.
  * `IP_CIDR`: [Classless Inter-Domain Routing (CIDR) notation IP address ranges](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) to blacklist, 
  * `USER`: Series of fields that a Unit21 user entity will be matched against.
  * `BUSINESS`: Series of fields that a Unit21 business entity will be matched against.


If the `/blacklists/create` API is called multiple times, it will create a new blacklist each time.  This endpoint does not support updates/upserts.

This endpoint does not support batch uploads.

The response will consist of the following fields:

  | Field           | Type     | Description                                           |
  |-----------------|----------|-------------------------------------------------------|
  | `blacklist_id`  | String   | 	Unique identifier of the entity on your platform     |


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
    req := shared.CreateBlacklist{
        Description: sdk.String("Unit21 rules will check against items in this list"),
        Name: sdk.String("New Blacklist"),
        Type: shared.CreateBlacklistTypeEnumIPInet.ToPointer(),
    }

    res, err := s.BlacklistsAPI.CreateBlacklist(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBlacklist200ApplicationJSONAny != nil {
        // handle response
    }
}
```

## ListBlacklists

Returns an array of blacklist in your environment. 

Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 50)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record.

The `total_count` field contains the total number of blacklists where the  `response_count` field contains the number of blacklists included in the response.


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
    req := shared.ListRequest{
        Limit: sdk.Int64(2),
        Offset: sdk.Int64(616934),
    }

    res, err := s.BlacklistsAPI.ListBlacklists(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListBlacklists200ApplicationJSONObject != nil {
        // handle response
    }
}
```
