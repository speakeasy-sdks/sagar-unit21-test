// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InstrumentListSourceEnum - LEGACY. If your platform owns and administers this instrument, the instrument is `internal`, otherwise it is `external`.
type InstrumentListSourceEnum string

const (
	InstrumentListSourceEnumInternal InstrumentListSourceEnum = "internal"
	InstrumentListSourceEnumExternal InstrumentListSourceEnum = "external"
)

func (e *InstrumentListSourceEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "internal":
		fallthrough
	case "external":
		*e = InstrumentListSourceEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InstrumentListSourceEnum: %s", s)
	}
}

// InstrumentList - Top-level data about an instrument.
type InstrumentList struct {
	CustomData *CustomData `json:"custom_data,omitempty"`
	// Unique identifier of the instrument. Intrument IDs must be unique and only comprise of the characters -_:.@a-zA-Z0-9!#$%&*+/=?^`{'
	InstrumentID *string `json:"instrument_id,omitempty"`
	// Type of instrument. Common examples include `bank`, `cash`, and `credit card`
	InstrumentType *string `json:"instrument_type,omitempty"`
	// Is the instrument buffered?
	IsSpooled *bool `json:"is_spooled,omitempty"`
	// Unique identifier of the parent instrument.
	ParentInstrumentID *string `json:"parent_instrument_id,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	RegisteredAt *int64 `json:"registered_at,omitempty"`
	// LEGACY. If your platform owns and administers this instrument, the instrument is `internal`, otherwise it is `external`.
	Source *InstrumentListSourceEnum `json:"source,omitempty"`
	// Status of the object on your system. You MAY enter any string value.
	Status *string `json:"status,omitempty"`
	// List of string tags, in the format `keyString:valueString` (note that the Key strings are NOT enclosed in `"`)
	Tags []string `json:"tags,omitempty"`
	// A Unit21 internally-assigned unique identifier for an object within the Unit21 system.
	Unit21ID *string `json:"unit21_id,omitempty"`
}
