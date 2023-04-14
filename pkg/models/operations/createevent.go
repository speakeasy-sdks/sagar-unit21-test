// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unit21/pkg/models/shared"
)

// CreateEventEventOptionsActionData - Information about any notable actions that a user takes on your system. Examples of actions that may be worth tracking are:
//   - Password changes from new IP addresses
//   - Logins from disparate locations
//   - Linking or unlinking instruments at an unusual frequency
//   - Finding users frequently using referral codes, potentially signally fake referral schemes.
type CreateEventEventOptionsActionData struct {
	// Extra information that provides more context about the event.
	ActionDetails *string `json:"action_details,omitempty"`
	// A keyword (of your choosing) that describes an action event, e.g. `LOGIN`, `SIGNUP`, `PASSWORD_RESET`.
	//
	ActionType *string `json:"action_type,omitempty"`
	// Identifier of the entity on your platform that is associated with this action event─should correspond to the `entity_id` field in an [entity's `general_data` section](ref:create_entity).
	EntityID string `json:"entity_id"`
	// Describes a user such as `employee` or `business`
	EntityType string `json:"entity_type"`
	// Identifier of the transaction instrument on your platform that is associated with this action event - should correspond to the `instrument_id` field in an [entity's `instrument_data` section](ref:create_entity)
	InstrumentID *string `json:"instrument_id,omitempty"`
}

// CreateEventEventOptionsDigitalData - Associated digital properties - IP, device, browser, client info etc.
type CreateEventEventOptionsDigitalData struct {
	// Either IPv4 or IPv6
	IPAddress *string `json:"ip_address,omitempty"`
}

type CreateEventEventOptionsGeneralDataParents struct {
	// Unique identifier of the event on your platform
	EventID *string `json:"event_id,omitempty"`
	// `transaction` for monetary flows, `action` for other state changes, like new logins.
	//
	EventType *shared.EventTypeEnum `json:"event_type,omitempty"`
}

// CreateEventEventOptionsGeneralData - General data is required for any request made to the v1/events/create endpoint. This defines any pieces of information that allows you to link up any event on Unit21's system to transactions or user activities on your platform.
type CreateEventEventOptionsGeneralData struct {
	// Unique identifier of the event on your platform
	EventID string `json:"event_id"`
	// Extra info about how your organization classifies the event. You MAY enter any value. Useful for granular categories, e.g. if you have two types of products and a transaction can be associated with either.
	//
	EventSubtype *string `json:"event_subtype,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	EventTime int64 `json:"event_time"`
	// `transaction` for monetary flows, `action` for other state changes, like new logins.
	//
	EventType shared.EventTypeEnum `json:"event_type"`
	// The parent object consists of two fields─`event_id` for parent unique identifier  and `event_type`.
	Parents []CreateEventEventOptionsGeneralDataParents `json:"parents,omitempty"`
	// Status of the object on your system. You MAY enter any string value.
	Status *string `json:"status,omitempty"`
	// List of string tags, in the format `keyString:valueString` (note that the Key strings are NOT enclosed in `"`)
	Tags []string `json:"tags,omitempty"`
}

// CreateEventEventOptionsOptionsLinkedEntityEnum - Possible values are `sender`, `receiver`, and `both`. Defaults to `both`. If `link_digital_data_to_entity` is flagged on transaction events, this specifies which entities to associate the `digital_data` to. If there is no `digital_data` or entities, no exception is thrown.
type CreateEventEventOptionsOptionsLinkedEntityEnum string

const (
	CreateEventEventOptionsOptionsLinkedEntityEnumSender   CreateEventEventOptionsOptionsLinkedEntityEnum = "sender"
	CreateEventEventOptionsOptionsLinkedEntityEnumReceiver CreateEventEventOptionsOptionsLinkedEntityEnum = "receiver"
	CreateEventEventOptionsOptionsLinkedEntityEnumBoth     CreateEventEventOptionsOptionsLinkedEntityEnum = "both"
)

func (e *CreateEventEventOptionsOptionsLinkedEntityEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "sender":
		fallthrough
	case "receiver":
		fallthrough
	case "both":
		*e = CreateEventEventOptionsOptionsLinkedEntityEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateEventEventOptionsOptionsLinkedEntityEnum: %s", s)
	}
}

type CreateEventEventOptionsOptions struct {
	// Whether or not to link the included `digital_data` with the related entities. Includes geoip information if resolve_geoip is enabled as well. On action events, defaults to `true`
	LinkDigitalDataToEntity *bool `json:"link_digital_data_to_entity,omitempty"`
	// Possible values are `sender`, `receiver`, and `both`. Defaults to `both`. If `link_digital_data_to_entity` is flagged on transaction events, this specifies which entities to associate the `digital_data` to. If there is no `digital_data` or entities, no exception is thrown.
	LinkedEntity *CreateEventEventOptionsOptionsLinkedEntityEnum `json:"linked_entity,omitempty"`
	// Only relevant for updates/upserts, ignored otherwise. See [custom data merge strategy](doc:how-data-merges-on-updates#custom-data-merge-strategy) for more details. **Default**: `false`
	MergeCustomData *bool `json:"merge_custom_data,omitempty"`
	// Whether or not to monitor this event (defaults to `true`). Typically used to signal Unit21 to not flag such events or include them in calculations i.e. to prevent double counting, or to ignore applying monitoring to unimportant events that you still want to associate with users
	Monitor *bool `json:"monitor,omitempty"`
	// If `false`, does not resolve the geographic location from the provided IP. If `true` and `digital_data.ip_addresses` is empty, Unit21 ignores the field. **Default**: `true`
	ResolveGeoip *bool `json:"resolve_geoip,omitempty"`
	// If POST request includes an object that already exists when  `upsert_on_conflict` is `false`, API returns a 409 error code and the object is not overwritten. **Default**: `true`
	UpsertOnConflict *bool `json:"upsert_on_conflict,omitempty"`
}

// CreateEventEventOptionsTransactionData - In addition to the required `amount` field, must include at least one of:
// * `sender_entity_id`
// * `sender_instrument_id`
// * `receiver_entity_id`
// * `receiver_instrument_id`
type CreateEventEventOptionsTransactionData struct {
	// The normalized monetary value of the transaction in customer's home currency. This value MUST be greater than or equal to zero. Direction of monetary flow should be entirely denoted by the specifying sender/receiver entities and instruments appropriately.
	//
	Amount float64 `json:"amount"`
	// If both `sent_currency` and `received_currency` are defined, and neither of them are in the home currency, then this is defined as `sent_amount/received_amount`.  If either `sent_currency` or `received_currency` are in the home currency  (but not both), then `exchange_rate` is defined as  `home_currency_amount/non_home_currency_amount` (i.e exchange rate to USD). If both `sent_currency` and `received_currency` are in the home currency,  this is not required.
	//
	ExchangeRate *float64 `json:"exchange_rate,omitempty"`
	// The sum of all external fees associated with the transaction, specified in USD. External fees are paid out to external parties independent of your platform e.g. payment network fees, shipping fees, cleaning fees. All other amount fields should not include the value of these fees
	//
	ExternalFee *float64 `json:"external_fee,omitempty"`
	// The sum of all internal fees associated with the transaction, specified in USD. Internal fees are collected by your platform e.g. if users pay a fee to you for facilitating a transaction from user to user. All other field amounts should not include the value of these fees.
	//
	InternalFee *float64 `json:"internal_fee,omitempty"`
	// The monetary value of the transaction, specified in terms of the currency set in the `received_currency` field. This value must be positive (greater or equal to zero). Direction of monetary flow should be entirely denoted by the specifying sender/receiver entities and instruments appropriately.
	ReceivedAmount *float64 `json:"received_amount,omitempty"`
	// The currency that the receiver party received
	ReceivedCurrency *string `json:"received_currency,omitempty"`
	// Identifier of the receiver entity on your platform
	ReceiverEntityID *string `json:"receiver_entity_id,omitempty"`
	// The entity type of the receiver as defined in Entity's general_data section.
	//
	ReceiverEntityType *string `json:"receiver_entity_type,omitempty"`
	// Identifier of the receiver's transaction instrument on your platform.
	//
	ReceiverInstrumentID *string `json:"receiver_instrument_id,omitempty"`
	// Identifier of the sender entity on your platform.
	//
	SenderEntityID *string `json:"sender_entity_id,omitempty"`
	// The entity type of the sender as defined in Entity's `general_data` section.
	//
	SenderEntityType *string `json:"sender_entity_type,omitempty"`
	// Identifier of the sender's transaction instrument on your platform.
	//
	SenderInstrumentID *string `json:"sender_instrument_id,omitempty"`
	// The monetary value of the transaction, specified in terms of the currency set in the sent_currency field. This value must be positive (greater or equal to zero). Denote by the specifying sender/receiver entities and instruments appropriately.
	//
	SentAmount *float64 `json:"sent_amount,omitempty"`
	// The currency that the sender sent
	SentCurrency *string `json:"sent_currency,omitempty"`
	// Used for chainalysis. Chanalysis transaction hash.
	TransactionHash *string `json:"transaction_hash,omitempty"`
	// Readable information associated with the exchange rate(s) used, e.g. when the exchange rate was pulled, from what source it was obtained
	UsdConversionNotes *string `json:"usd_conversion_notes,omitempty"`
}

type CreateEventEventOptions struct {
	// Information about any notable actions that a user takes on your system. Examples of actions that may be worth tracking are:
	//   * Password changes from new IP addresses
	//   * Logins from disparate locations
	//   * Linking or unlinking instruments at an unusual frequency
	//   * Finding users frequently using referral codes, potentially signally fake referral schemes.
	//
	ActionData *CreateEventEventOptionsActionData `json:"action_data,omitempty"`
	// Any custom information that you wish our system to associate with this object  (accepts any valid JSON object -- up to 3 layers deep -- in key:value format --  `string:string`, `string:int`, `string:bool` -- no arrays)
	CustomData map[string]interface{} `json:"custom_data,omitempty"`
	// Associated digital properties - IP, device, browser, client info etc.
	DigitalData *CreateEventEventOptionsDigitalData `json:"digital_data,omitempty"`
	// General data is required for any request made to the v1/events/create endpoint. This defines any pieces of information that allows you to link up any event on Unit21's system to transactions or user activities on your platform.
	//
	GeneralData *CreateEventEventOptionsGeneralData `json:"general_data,omitempty"`
	// Address/location data
	LocationData *shared.LocationDataProperties  `json:"location_data,omitempty"`
	Options      *CreateEventEventOptionsOptions `json:"options,omitempty"`
	// In addition to the required `amount` field, must include at least one of:
	// * `sender_entity_id`
	// * `sender_instrument_id`
	// * `receiver_entity_id`
	// * `receiver_instrument_id`
	//
	TransactionData *CreateEventEventOptionsTransactionData `json:"transaction_data,omitempty"`
}

type CreateEventResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
