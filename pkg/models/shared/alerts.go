// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AlertsAlertTypeEnum - Either transaction monitoring, `tm`, or know-your-customer `kyc`. Default is `tm`
type AlertsAlertTypeEnum string

const (
	AlertsAlertTypeEnumTm          AlertsAlertTypeEnum = "tm"
	AlertsAlertTypeEnumKyc         AlertsAlertTypeEnum = "kyc"
	AlertsAlertTypeEnumChainalysis AlertsAlertTypeEnum = "chainalysis"
)

func (e AlertsAlertTypeEnum) ToPointer() *AlertsAlertTypeEnum {
	return &e
}

func (e *AlertsAlertTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "tm":
		fallthrough
	case "kyc":
		fallthrough
	case "chainalysis":
		*e = AlertsAlertTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for AlertsAlertTypeEnum: %s", s)
	}
}

type Alerts struct {
	// Unique identifier of the alert on the customer's platform.
	AlertID string `json:"alert_id"`
	// Either transaction monitoring, `tm`, or know-your-customer `kyc`. Default is `tm`
	AlertType AlertsAlertTypeEnum `json:"alert_type"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	CreatedAt int64 `json:"created_at"`
	// Description of the alert
	Description *string `json:"description,omitempty"`
	// Labels that describe the outcome of an alert or case investigation. More info in [this knowledge base article about dispositions](https://docs.unit21.ai/docs/concept-dispositions).
	Disposition *string `json:"disposition,omitempty"`
	// Free form text documenting reasoning and investigation notes
	DispositionNotes *string `json:"disposition_notes,omitempty"`
	// Investigation status, either `OPEN` or `ClOSED`
	Status InvestigationStatusEnum `json:"status"`
	// List of string tags, in the format `keyString:valueString` (note that the Key strings are NOT enclosed in `"`)
	Tags []string `json:"tags,omitempty"`
	// Title of the alert
	Title string `json:"title"`
	// Integer value greater than or equal to 1. Used when `alert_type` is `kyc`.
	VerificationResultID *int64 `json:"verification_result_id,omitempty"`
}
