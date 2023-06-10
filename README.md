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
            APIKeyAuth: "",
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


### [AgentsAPI](docs/sdks/agentsapi/README.md)

* [DeactivateAgent](docs/sdks/agentsapi/README.md#deactivateagent) - Deactivate an agent
* [ListAgents](docs/sdks/agentsapi/README.md#listagents) - List agents

### [AlertsAPI](docs/sdks/alertsapi/README.md)

* [CreateAlert](docs/sdks/alertsapi/README.md#createalert) - Create alerts
* [ExportAlerts](docs/sdks/alertsapi/README.md#exportalerts) - Bulk export alerts
* [GetAlertByUnit21ID](docs/sdks/alertsapi/README.md#getalertbyunit21id) - Get an alert
* [LinkMediaToAlert](docs/sdks/alertsapi/README.md#linkmediatoalert) - Add media to an alert
* [ListAlerts](docs/sdks/alertsapi/README.md#listalerts) - List alerts
* [UpdateAlert](docs/sdks/alertsapi/README.md#updatealert) - Update alert

### [BlacklistsAPI](docs/sdks/blacklistsapi/README.md)

* [AddBlacklistValues](docs/sdks/blacklistsapi/README.md#addblacklistvalues) - Add items to a blacklist
* [CreateBlacklist](docs/sdks/blacklistsapi/README.md#createblacklist) - Create a blacklist
* [ListBlacklists](docs/sdks/blacklistsapi/README.md#listblacklists) - List blacklists

### [CasesAPI](docs/sdks/casesapi/README.md)

* [CreateCase](docs/sdks/casesapi/README.md#createcase) - Create a case
* [ExportCases](docs/sdks/casesapi/README.md#exportcases) - Bulk export cases
* [GetCaseByUnit21ID](docs/sdks/casesapi/README.md#getcasebyunit21id) - Get a case
* [LinkMediaToCase](docs/sdks/casesapi/README.md#linkmediatocase) - Add media to a case
* [ListCases](docs/sdks/casesapi/README.md#listcases) - List cases
* [UpdateCase](docs/sdks/casesapi/README.md#updatecase) - Update case

### [DatafilesAPI](docs/sdks/datafilesapi/README.md)

* [CreateDatafiles](docs/sdks/datafilesapi/README.md#createdatafiles) - Upload datafiles
* [GetDatafileByUnit21ID](docs/sdks/datafilesapi/README.md#getdatafilebyunit21id) - Get datafile
* [GetDatafileMappings](docs/sdks/datafilesapi/README.md#getdatafilemappings) - Retrieve datafile mappings

### [DevicesAPI](docs/sdks/devicesapi/README.md)

* [CreateDevice](docs/sdks/devicesapi/README.md#createdevice) - Create a device
* [ExportDevices](docs/sdks/devicesapi/README.md#exportdevices) - Bulk export devices
* [GetDeviceByExternal](docs/sdks/devicesapi/README.md#getdevicebyexternal) - Get a device using external ID
* [ListDevices](docs/sdks/devicesapi/README.md#listdevices) - List devices
* [UpdateDeviceByExternal](docs/sdks/devicesapi/README.md#updatedevicebyexternal) - Update device using external ID

### [EntitiesAPI](docs/sdks/entitiesapi/README.md)

* [AddInstruments](docs/sdks/entitiesapi/README.md#addinstruments) - Add instruments to entity
* [CreateEntity](docs/sdks/entitiesapi/README.md#createentity) - Create an entity
* [DelMediaEntity](docs/sdks/entitiesapi/README.md#delmediaentity) - Delete entity media
* [ExportEntities](docs/sdks/entitiesapi/README.md#exportentities) - Bulk export entities
* [GetEntity](docs/sdks/entitiesapi/README.md#getentity) - Get an entity
* [LinkMediaToEntity](docs/sdks/entitiesapi/README.md#linkmediatoentity) - Add media to an entity
* [ListEntities](docs/sdks/entitiesapi/README.md#listentities) - List entities
* [UpdateEntity](docs/sdks/entitiesapi/README.md#updateentity) - Update entity

### [EntityVerificationAPI](docs/sdks/entityverificationapi/README.md)

* [AddVerificationResultToEntity](docs/sdks/entityverificationapi/README.md#addverificationresulttoentity) - Link external verification
* [GetEntityVerificationWorkflowExecutions](docs/sdks/entityverificationapi/README.md#getentityverificationworkflowexecutions) - Get entity verification workflow IDs
* [GetVerificationResult](docs/sdks/entityverificationapi/README.md#getverificationresult) - Get verification results by result id
* [GetVerificationResultFromWorkflowExecution](docs/sdks/entityverificationapi/README.md#getverificationresultfromworkflowexecution) - Get verification results from workflow
* [GetVerificationWorkflowExecution](docs/sdks/entityverificationapi/README.md#getverificationworkflowexecution) - Get verification workflow execution details
* [RunVerificationsWorkflowThroughExternalID](docs/sdks/entityverificationapi/README.md#runverificationsworkflowthroughexternalid) - Verify an entity
* [UpdateContinuousMonitoring](docs/sdks/entityverificationapi/README.md#updatecontinuousmonitoring) - Update continuous monitoring
* [UpdateSuppressedProviderEntities](docs/sdks/entityverificationapi/README.md#updatesuppressedproviderentities) - Suppress provider entity

### [EventsAPI](docs/sdks/eventsapi/README.md)

* [CreateEvent](docs/sdks/eventsapi/README.md#createevent) - Create an event
* [ExportEvents](docs/sdks/eventsapi/README.md#exportevents) - Bulk export events
* [ExportTransactions](docs/sdks/eventsapi/README.md#exporttransactions) - Bulk export transactions
* [GetEvent](docs/sdks/eventsapi/README.md#getevent) - Get an event
* [ListEvents](docs/sdks/eventsapi/README.md#listevents) - List events
* [UpdateEvent](docs/sdks/eventsapi/README.md#updateevent) - Update event

### [ExportsAPI](docs/sdks/exportsapi/README.md)

* [DownloadFileExport](docs/sdks/exportsapi/README.md#downloadfileexport) - Download export
* [ListExports](docs/sdks/exportsapi/README.md#listexports) - List exports

### [ImportAPI](docs/sdks/importapi/README.md)

* [DatafileStatus](docs/sdks/importapi/README.md#datafilestatus) - Retrieve datafile status
* [GetPreSignedURL](docs/sdks/importapi/README.md#getpresignedurl) - Get pre-signed URL
* [ListDatafiles](docs/sdks/importapi/README.md#listdatafiles) - Retrieve datafiles list
* [UploadDatafiles](docs/sdks/importapi/README.md#uploaddatafiles) - Upload data to URL

### [InstrumentsAPI](docs/sdks/instrumentsapi/README.md)

* [CreateInstrument](docs/sdks/instrumentsapi/README.md#createinstrument) - Create an instrument
* [ExportInstruments](docs/sdks/instrumentsapi/README.md#exportinstruments) - Bulk export instruments
* [GetInstrument](docs/sdks/instrumentsapi/README.md#getinstrument) - Get an instrument
* [ListInstruments](docs/sdks/instrumentsapi/README.md#listinstruments) - List instruments
* [UpdateInstrument](docs/sdks/instrumentsapi/README.md#updateinstrument) - Update instrument

### [RulesAPI](docs/sdks/rulesapi/README.md)

* [ExportRules](docs/sdks/rulesapi/README.md#exportrules) - Bulk export rules
* [ListRules](docs/sdks/rulesapi/README.md#listrules) - List rules
* [ReadOneRule](docs/sdks/rulesapi/README.md#readonerule) - Get a rule

### [SarsAPI](docs/sdks/sarsapi/README.md)

* [ExportSars](docs/sdks/sarsapi/README.md#exportsars) - Bulk export sars
* [ListSars](docs/sdks/sarsapi/README.md#listsars) - List sars
* [ReadOneSar](docs/sdks/sarsapi/README.md#readonesar) - Get a sars

### [TagAssociationsAPI](docs/sdks/tagassociationsapi/README.md)

* [ListTags](docs/sdks/tagassociationsapi/README.md#listtags) - List tags

### [VerificationFormsAPI](docs/sdks/verificationformsapi/README.md)

* [CreateVerificationForm](docs/sdks/verificationformsapi/README.md#createverificationform) - Verification Forms API

### [WebhooksAPI](docs/sdks/webhooksapi/README.md)

* [UpdateWebhook](docs/sdks/webhooksapi/README.md#updatewebhook) - Update webhook URL
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
