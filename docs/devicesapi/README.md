# DevicesAPI

## Overview

Devices representing any computer or physical device used to execute an event. Devices are most suitable when events can be traced back to a specific device fingerprint. The `/devices` endpoint can create, list, and update instruments. 


### Available Operations

* [CreateDevice](#createdevice) - Create a device
* [ExportDevices](#exportdevices) - Bulk export devices
* [GetDeviceByExternal](#getdevicebyexternal) - Get a device using external ID
* [ListDevices](#listdevices) - List devices
* [UpdateDeviceByExternal](#updatedevicebyexternal) - Update device using external ID

## CreateDevice

Creates a new device, sending device data in the request body.

Recommended values for `device_type` include: laptop, home_computer, work_laptop, phone, smartphone, scanner.

If the `/devices/create` API is called for an device that already exists in our system (i.e.  has an existing `device_id`, it is treated it as an  [upsert](https://docs.unit21.ai/u21/reference/should-i-update-or-upsert) and an update on the existing device is performed. The response to the request will then contain the entry `previously_existed: true`. 

Unit21 selectively ignores upserts if the request is identical to a previous request.  The response to any isgnored upsert will  contain the field `ignored: true`.

Updates to an device's `device_id` are not allowed.

Follow the links for more information:
  - [Endpoint options](https://docs.unit21.ai/u21/reference/endpoint-options)
  - [Custom data](https://docs.unit21.ai/u21/reference/best-practices-for-custom-data)
  - [Batch uploads](https://docs.unit21.ai/u21/reference/batch-request-examples)
  - [Modifying tags](https://docs.unit21.ai/u21/reference/modifying-tags)


The response will consist of the following fields:

  | Field                    | Type     | Description                                            |
  |--------------------------|----------|--------------------------------------------------------|
  | `device_id`	             | String   | 	Unique identifier of the device on your platform     |
  | `unit21_id`	             | String   | 	Internal ID of the device within Unit21's system     |
  | `previously_existed`	   | Boolean  | 	If entity (with the same `device_id`) already exists |


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
    res, err := s.DevicesAPI.CreateDevice(ctx, operations.CreateDeviceDeviceData{
        AppVersion: sdk.String("5.44.4"),
        DeviceID: "11b72726-18d6-43b3-a0bf-b4adf6dfd2da",
        DeviceManufacturer: sdk.String("samsung"),
        DeviceModel: sdk.String("SM-N970U"),
        DeviceSubtype: sdk.String("android"),
        DeviceType: "mobile",
        NetworkCarrier: sdk.String("T-Mobile"),
        NetworkCellular: sdk.Bool(true),
        Options: &shared.Options{
            MergeCustomData: sdk.Bool(false),
            ResolveGeoip: sdk.Bool(true),
            UpsertOnConflict: sdk.Bool(true),
        },
        OsName: sdk.String("Android"),
        OsVersion: sdk.String("10.1.x"),
        PhoneNumbers: []string{
            "+14159627132",
            "+14159627132",
            "+14159627132",
        },
        RegisteredAt: sdk.Int64(1623365011),
        Status: sdk.String("active"),
        Tags: []string{
            "Sector:Europe",
            "Sector:Europe",
        },
        Timezone: sdk.String("UTC-7"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDeviceResponse != nil {
        // handle response
    }
}
```

## ExportDevices

Initiates an email and dashboard export of devices. The export will be as a CSV file.

Either the agent `ID` or `email` is required to begin the export.

Either the `filters` or the list of `device IDs` are required for the export.


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
    res, err := s.DevicesAPI.ExportDevices(ctx, operations.ExportDevicesRequestBody{
        AgentEmail: sdk.String("foobar@unit21.ai"),
        AgentID: sdk.String("295"),
        DeviceIds: []int64{
            778346,
        },
        Filters: &operations.ExportDevicesRequestBodyFilters{
            ChildOrgIds: []int64{
                85,
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

## GetDeviceByExternal

Returns all data objects belonging to a single device.

This endpoint requires the `device_id` which is a unique ID created by your organization to identify the device. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.

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
    res, err := s.DevicesAPI.GetDeviceByExternal(ctx, operations.GetDeviceByExternalRequest{
        DeviceID: "tenetur",
        OrgName: "ipsam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## ListDevices

Returns paginated list of of top-level information about devices. 
Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
* `limit`  indicates how many objects the request returns (the page maximum is 50)
* `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record. The offset is relative to the number of pages (not the total count of objects).

The `total_count` field contains the total number of devices where the  `response_count` field contains the number of devices included in the response.


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
            APIKeyAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.DevicesAPI.ListDevices(ctx, shared.ListRequest{
        Limit: sdk.Int64(2),
        Offset: sdk.Int64(662527),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListResponse != nil {
        // handle response
    }
}
```

## UpdateDeviceByExternal

Updates an device's information using the `device_id` from your platform. 

Updating a device has no required fields. You MAY send any  subset of the fields that the `devices/create` endpoint accepts.

This endpoint requires the `device_id` which is a unique ID created by your organization to identify the device. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.

Note that you can also update a device using an upsert through `/devices/create`.

Follow the links for more information:
  - [Endpoint options](https://docs.unit21.ai/u21/reference/endpoint-options)
  - [Custom data](https://docs.unit21.ai/u21/reference/best-practices-for-custom-data)
  - [Batch uploads](https://docs.unit21.ai/u21/reference/batch-request-examples)
  - [Modifying tags](https://docs.unit21.ai/u21/reference/modifying-tags)


The response will consist of the following fields:

  | Field                    | Type     | Description                                            |
  |--------------------------|----------|--------------------------------------------------------|
  | `device_id`	             | String   | 	Unique identifier of the device on your platform     |
  | `unit21_id`	             | String   | 	Internal ID of the device within Unit21's system     |
  | `previously_existed`	   | Boolean  | 	If entity (with the same `device_id`) already exists |

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
    res, err := s.DevicesAPI.UpdateDeviceByExternal(ctx, operations.UpdateDeviceByExternalRequest{
        RequestBody: &shared.CustomData{
            CustomData: map[string]interface{}{
                "quasi": "error",
            },
        },
        DeviceID: "temporibus",
        OrgName: "laborum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
