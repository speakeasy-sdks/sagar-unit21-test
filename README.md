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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AgentsAPI](docs/agentsapi/README.md)

* [DeactivateAgent](docs/agentsapi/README.md#deactivateagent) - Deactivate an agent
* [ListAgents](docs/agentsapi/README.md#listagents) - List agents

### [AlertsAPI](docs/alertsapi/README.md)

* [CreateAlert](docs/alertsapi/README.md#createalert) - Create alerts
* [ExportAlerts](docs/alertsapi/README.md#exportalerts) - Bulk export alerts
* [GetAlertByUnit21ID](docs/alertsapi/README.md#getalertbyunit21id) - Get an alert
* [LinkMediaToAlert](docs/alertsapi/README.md#linkmediatoalert) - Add media to an alert
* [ListAlerts](docs/alertsapi/README.md#listalerts) - List alerts
* [UpdateAlert](docs/alertsapi/README.md#updatealert) - Update alert

### [BlacklistsAPI](docs/blacklistsapi/README.md)

* [AddBlacklistValues](docs/blacklistsapi/README.md#addblacklistvalues) - Add items to a blacklist
* [CreateBlacklist](docs/blacklistsapi/README.md#createblacklist) - Create a blacklist
* [ListBlacklists](docs/blacklistsapi/README.md#listblacklists) - List blacklists

### [CasesAPI](docs/casesapi/README.md)

* [CreateCase](docs/casesapi/README.md#createcase) - Create a case
* [ExportCases](docs/casesapi/README.md#exportcases) - Bulk export cases
* [GetCaseByUnit21ID](docs/casesapi/README.md#getcasebyunit21id) - Get a case
* [LinkMediaToCase](docs/casesapi/README.md#linkmediatocase) - Add media to a case
* [ListCases](docs/casesapi/README.md#listcases) - List cases
* [UpdateCase](docs/casesapi/README.md#updatecase) - Update case

### [DatafilesAPI](docs/datafilesapi/README.md)

* [CreateDatafiles](docs/datafilesapi/README.md#createdatafiles) - Upload datafiles
* [GetDatafileByUnit21ID](docs/datafilesapi/README.md#getdatafilebyunit21id) - Get datafile
* [GetDatafileMappings](docs/datafilesapi/README.md#getdatafilemappings) - Retrieve datafile mappings

### [DevicesAPI](docs/devicesapi/README.md)

* [CreateDevice](docs/devicesapi/README.md#createdevice) - Create a device
* [ExportDevices](docs/devicesapi/README.md#exportdevices) - Bulk export devices
* [GetDeviceByExternal](docs/devicesapi/README.md#getdevicebyexternal) - Get a device using external ID
* [ListDevices](docs/devicesapi/README.md#listdevices) - List devices
* [UpdateDeviceByExternal](docs/devicesapi/README.md#updatedevicebyexternal) - Update device using external ID

### [EntitiesAPI](docs/entitiesapi/README.md)

* [AddInstruments](docs/entitiesapi/README.md#addinstruments) - Add instruments to entity
* [CreateEntity](docs/entitiesapi/README.md#createentity) - Create an entity
* [DelMediaEntity](docs/entitiesapi/README.md#delmediaentity) - Delete entity media
* [ExportEntities](docs/entitiesapi/README.md#exportentities) - Bulk export entities
* [GetEntity](docs/entitiesapi/README.md#getentity) - Get an entity
* [LinkMediaToEntity](docs/entitiesapi/README.md#linkmediatoentity) - Add media to an entity
* [ListEntities](docs/entitiesapi/README.md#listentities) - List entities
* [UpdateEntity](docs/entitiesapi/README.md#updateentity) - Update entity

### [EntityVerificationAPI](docs/entityverificationapi/README.md)

* [AddVerificationResultToEntity](docs/entityverificationapi/README.md#addverificationresulttoentity) - Link external verification
* [GetEntityVerificationWorkflowExecutions](docs/entityverificationapi/README.md#getentityverificationworkflowexecutions) - Get entity verification workflow IDs
* [GetVerificationResult](docs/entityverificationapi/README.md#getverificationresult) - Get verification results by result id
* [GetVerificationResultFromWorkflowExecution](docs/entityverificationapi/README.md#getverificationresultfromworkflowexecution) - Get verification results from workflow
* [GetVerificationWorkflowExecution](docs/entityverificationapi/README.md#getverificationworkflowexecution) - Get verification workflow execution details
* [RunVerificationsWorkflowThroughExternalID](docs/entityverificationapi/README.md#runverificationsworkflowthroughexternalid) - Verify an entity
* [UpdateContinuousMonitoring](docs/entityverificationapi/README.md#updatecontinuousmonitoring) - Update continuous monitoring
* [UpdateSuppressedProviderEntities](docs/entityverificationapi/README.md#updatesuppressedproviderentities) - Suppress provider entity

### [EventsAPI](docs/eventsapi/README.md)

* [CreateEvent](docs/eventsapi/README.md#createevent) - Create an event
* [ExportEvents](docs/eventsapi/README.md#exportevents) - Bulk export events
* [ExportTransactions](docs/eventsapi/README.md#exporttransactions) - Bulk export transactions
* [GetEvent](docs/eventsapi/README.md#getevent) - Get an event
* [ListEvents](docs/eventsapi/README.md#listevents) - List events
* [UpdateEvent](docs/eventsapi/README.md#updateevent) - Update event

### [ExportsAPI](docs/exportsapi/README.md)

* [DownloadFileExport](docs/exportsapi/README.md#downloadfileexport) - Download export
* [ListExports](docs/exportsapi/README.md#listexports) - List exports

### [ImportAPI](docs/importapi/README.md)

* [DatafileStatus](docs/importapi/README.md#datafilestatus) - Retrieve datafile status
* [GetPreSignedURL](docs/importapi/README.md#getpresignedurl) - Get pre-signed URL
* [ListDatafiles](docs/importapi/README.md#listdatafiles) - Retrieve datafiles list
* [UploadDatafiles](docs/importapi/README.md#uploaddatafiles) - Upload data to URL

### [InstrumentsAPI](docs/instrumentsapi/README.md)

* [CreateInstrument](docs/instrumentsapi/README.md#createinstrument) - Create an instrument
* [ExportInstruments](docs/instrumentsapi/README.md#exportinstruments) - Bulk export instruments
* [GetInstrument](docs/instrumentsapi/README.md#getinstrument) - Get an instrument
* [ListInstruments](docs/instrumentsapi/README.md#listinstruments) - List instruments
* [UpdateInstrument](docs/instrumentsapi/README.md#updateinstrument) - Update instrument

### [RulesAPI](docs/rulesapi/README.md)

* [ExportRules](docs/rulesapi/README.md#exportrules) - Bulk export rules
* [ListRules](docs/rulesapi/README.md#listrules) - List rules
* [ReadOneRule](docs/rulesapi/README.md#readonerule) - Get a rule

### [SarsAPI](docs/sarsapi/README.md)

* [ExportSars](docs/sarsapi/README.md#exportsars) - Bulk export sars
* [ListSars](docs/sarsapi/README.md#listsars) - List sars
* [ReadOneSar](docs/sarsapi/README.md#readonesar) - Get a sars

### [TagAssociationsAPI](docs/tagassociationsapi/README.md)

* [ListTags](docs/tagassociationsapi/README.md#listtags) - List tags

### [VerificationFormsAPI](docs/verificationformsapi/README.md)

* [CreateVerificationForm](docs/verificationformsapi/README.md#createverificationform) - Verification Forms API

### [WebhooksAPI](docs/webhooksapi/README.md)

* [UpdateWebhook](docs/webhooksapi/README.md#updatewebhook) - Update webhook URL
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
