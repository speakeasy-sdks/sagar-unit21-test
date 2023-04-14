# unit21

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/sagar-unit21-test
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "unit21"
    "unit21/pkg/models/shared"
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### AgentsAPI

* `DeactivateAgent` - Deactivate an agent
* `ListAgents` - List agents

### AlertsAPI

* `CreateAlert` - Create alerts
* `ExportAlerts` - Bulk export alerts
* `GetAlertByUnit21ID` - Get an alert
* `LinkMediaToAlert` - Add media to an alert
* `ListAlerts` - List alerts
* `UpdateAlert` - Update alert

### BlacklistsAPI

* `AddBlacklistValues` - Add items to a blacklist
* `CreateBlacklist` - Create a blacklist
* `ListBlacklists` - List blacklists

### CasesAPI

* `CreateCase` - Create a case
* `ExportCases` - Bulk export cases
* `GetCaseByUnit21ID` - Get a case
* `LinkMediaToCase` - Add media to a case
* `ListCases` - List cases
* `UpdateCase` - Update case

### DatafilesAPI

* `CreateDatafiles` - Upload datafiles
* `GetDatafileByUnit21ID` - Get datafile
* `GetDatafileMappings` - Retrieve datafile mappings

### DevicesAPI

* `CreateDevice` - Create a device
* `ExportDevices` - Bulk export devices
* `GetDeviceByExternal` - Get a device using external ID
* `ListDevices` - List devices
* `UpdateDeviceByExternal` - Update device using external ID

### EntitiesAPI

* `AddInstruments` - Add instruments to entity
* `CreateEntity` - Create an entity
* `DelMediaEntity` - Delete entity media
* `ExportEntities` - Bulk export entities
* `GetEntity` - Get an entity
* `LinkMediaToEntity` - Add media to an entity
* `ListEntities` - List entities
* `UpdateEntity` - Update entity

### EntityVerificationAPI

* `AddVerificationResultToEntity` - Link external verification
* `GetEntityVerificationWorkflowExecutions` - Get entity verification workflow IDs
* `GetVerificationResult` - Get verification results by result id
* `GetVerificationResultFromWorkflowExecution` - Get verification results from workflow
* `GetVerificationWorkflowExecution` - Get verification workflow execution details
* `RunVerificationsWorkflowThroughExternalID` - Verify an entity
* `UpdateContinuousMonitoring` - Update continuous monitoring
* `UpdateSuppressedProviderEntities` - Suppress provider entity

### EventsAPI

* `CreateEvent` - Create an event
* `ExportEvents` - Bulk export events
* `ExportTransactions` - Bulk export transactions
* `GetEvent` - Get an event
* `ListEvents` - List events
* `UpdateEvent` - Update event

### ExportsAPI

* `DownloadFileExport` - Download export
* `ListExports` - List exports

### ImportAPI

* `DatafileStatus` - Retrieve datafile status
* `GetPreSignedURL` - Get pre-signed URL
* `ListDatafiles` - Retrieve datafiles list
* `UploadDatafiles` - Upload data to URL

### InstrumentsAPI

* `CreateInstrument` - Create an instrument
* `ExportInstruments` - Bulk export instruments
* `GetInstrument` - Get an instrument
* `ListInstruments` - List instruments
* `UpdateInstrument` - Update instrument

### RulesAPI

* `ExportRules` - Bulk export rules
* `ListRules` - List rules
* `ReadOneRule` - Get a rule

### SarsAPI

* `ExportSars` - Bulk export sars
* `ListSars` - List sars
* `ReadOneSar` - Get a sars

### TagAssociationsAPI

* `ListTags` - List tags

### VerificationFormsAPI

* `CreateVerificationForm` - Verification Forms API

### WebhooksAPI

* `UpdateWebhook` - Update webhook URL
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
