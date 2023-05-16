// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"net/http"
	"time"
	"unit21/pkg/models/shared"
	"unit21/pkg/utils"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Sandbox1 server for testing
	"https://sandbox1-api.unit21.com/v1",
	// Sandbox2 server for testing
	"https://sandbox2-api.unit21.com/v1",
	// Main production server
	"https://api.unit21.com/v1",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// SDK - Every endpoint available to consumers of the Unit21 API.
type SDK struct {
	// AgentsAPI - Agents are your organization's members who use the Unit21 system to investigate suspicious objects and events. The `/agents` endpoint can list your agents.
	//
	AgentsAPI *agentsAPI
	// AlertsAPI - Alerts have two origins. Typically, alerts are generated whenever a Unit21 detection tool (like a rule) flags an object, like an entity. However, your organization can also send alerts generated from your in-house detection systems. The `/alerts` endpoint can create, list, and update alerts.
	//
	AlertsAPI *alertsAPI
	// BlacklistsAPI - Blacklists comprise one of the following categories:
	//   * of entities (users or business)
	//   * IPs (single or ranges)
	//   * strings
	//
	BlacklistsAPI *blacklistsAPI
	// CasesAPI - Cases are usually active investigations, which may span multiple events, entities and documents. They can be directly escalated into a suspicious activity report. The `/cases` endpoint can create, list, and update cases.
	//
	CasesAPI *casesAPI
	// DatafilesAPI - If you want to bulk upload multiple objects, you can send them via a POST to the `/datafiles` endpoint. For the fastest processing, the datafile SHOULD be a JSON file in the format of a typical POST request to this API.
	//
	DatafilesAPI *datafilesAPI
	// DevicesAPI - Devices representing any computer or physical device used to execute an event. Devices are most suitable when events can be traced back to a specific device fingerprint. The `/devices` endpoint can create, list, and update instruments.
	//
	DevicesAPI *devicesAPI
	// EntitiesAPI - _Entities_ are typically businesses or users that have transactions on your platform.  The `/entities` endpoint can create, list, and update entities.
	//
	EntitiesAPI *entitiesAPI
	// EntityVerificationAPI - Unit21 can execute entity verifications according to steps defined in system workflows. They can run when a new entity is added to the system, an update is made to an existing entity, or by directly triggering an API endpoint.
	//
	EntityVerificationAPI *entityVerificationAPI
	// EventsAPI - Events have two types, _transaction events_  and _action events_:
	// * Transaction events are any monetary flow that is sent or received by an entity on your system.
	// * Action events are non-monetary changes of state that occur on your system, e.g. user logins. The `/events` endpoint sends and receives data about significant actions that occur with an entity or instrument on your system.
	//
	EventsAPI *eventsAPI
	// ExportsAPI - If you want to download files requested via the dashboard or the bulk export endpoints. The `/file-exports` endpoint can get and list sars.
	//
	ExportsAPI *exportsAPI
	// ImportAPI - If you want to upload a file with a lot of information in it, you can send them via a POST to a custom URL. You can request this URL via a POST to the `/datafiles` endpoint.  For the fastest processing, the file SHOULD be JSON or CSV.
	//
	ImportAPI *importAPI
	// InstrumentsAPI - Instruments represent any physical, digital, or logical intermediary between an entity and a transaction event. The `/instruments` endpoint can create, list, and update instruments.
	//
	InstrumentsAPI *instrumentsAPI
	// RulesAPI - Rules are the model logic that will find fraudulent and suspicious transactions and actions. Rules create alerts that can turn into cases with flagged entities, transactions and instruments. The `/rules` endpoint can get and list rules.
	//
	RulesAPI *rulesAPI
	// SarsAPI - Sars are cases that have been investigated and turned into a Suspicious Activity report with the intent to file it to FinCen. The `/sars` endpoint can get and list sars.
	//
	SarsAPI *sarsAPI
	// TagAssociationsAPI - Tags provide a flexible means of linking objects together in the Unit21 system. You can use the `/tag-associations` endpoint to explore these associations.
	//
	TagAssociationsAPI *tagAssociationsAPI
	// VerificationFormsAPI - With verification forms, you can automate ID verification and user collection. To gather user input, the `verification forms` endpoint creates a URL. This URL is only valid for a specified period of time.
	//
	VerificationFormsAPI *verificationFormsAPI
	// WebhooksAPI - Whenever an event happens on the Unit21 platform, Unit21 can send a webhook about the event to whatever URL you configure. Such events include entity verification results, generated alerts, case re-openings and closings, etcetera.
	//
	WebhooksAPI *webhooksAPI

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk._serverURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk._defaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SDK) {
		sdk._security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		_language:   "go",
		_sdkVersion: "0.4.0",
		_genVersion: "2.28.0",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {
		if sdk._security != nil {
			sdk._securityClient = utils.ConfigureSecurityClient(sdk._defaultClient, sdk._security)
		} else {
			sdk._securityClient = sdk._defaultClient
		}
	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.AgentsAPI = newAgentsAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.AlertsAPI = newAlertsAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.BlacklistsAPI = newBlacklistsAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.CasesAPI = newCasesAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.DatafilesAPI = newDatafilesAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.DevicesAPI = newDevicesAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.EntitiesAPI = newEntitiesAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.EntityVerificationAPI = newEntityVerificationAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.EventsAPI = newEventsAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.ExportsAPI = newExportsAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.ImportAPI = newImportAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.InstrumentsAPI = newInstrumentsAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.RulesAPI = newRulesAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.SarsAPI = newSarsAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.TagAssociationsAPI = newTagAssociationsAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.VerificationFormsAPI = newVerificationFormsAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.WebhooksAPI = newWebhooksAPI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
