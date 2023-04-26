// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// InstrumentBasicSourceEnum - LEGACY. If your platform owns and administers this instrument, the instrument is `internal`, otherwise it is `external`.
type InstrumentBasicSourceEnum string

const (
	InstrumentBasicSourceEnumInternal InstrumentBasicSourceEnum = "internal"
	InstrumentBasicSourceEnumExternal InstrumentBasicSourceEnum = "external"
)

func (e InstrumentBasicSourceEnum) ToPointer() *InstrumentBasicSourceEnum {
	return &e
}

func (e *InstrumentBasicSourceEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "internal":
		fallthrough
	case "external":
		*e = InstrumentBasicSourceEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for InstrumentBasicSourceEnum: %s", s)
	}
}

// InstrumentBasic - Top-level data about an instrument.
type InstrumentBasic struct {
	// A more detailed type of instrument. Common examples include `visa`, `mastercard` or `american-express` if type is `cc`
	InstrumentSubtype *string `json:"instrument_subtype,omitempty"`
	// Type of instrument. Common examples include `bank`, `cash`, and `credit card`
	InstrumentType *string  `json:"instrument_type,omitempty"`
	Options        *Options `json:"options,omitempty"`
	// Unique identifier of the parent instrument.
	ParentInstrumentID *string `json:"parent_instrument_id,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	RegisteredAt *int64 `json:"registered_at,omitempty"`
	// LEGACY. If your platform owns and administers this instrument, the instrument is `internal`, otherwise it is `external`.
	Source *InstrumentBasicSourceEnum `json:"source,omitempty"`
	// Status of the object on your system. You MAY enter any string value.
	Status *string `json:"status,omitempty"`
	// List of string tags, in the format `keyString:valueString` (note that the Key strings are NOT enclosed in `"`)
	Tags []string `json:"tags,omitempty"`
}
