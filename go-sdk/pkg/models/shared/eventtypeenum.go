// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// EventTypeEnum - `transaction` for monetary flows, `action` for other state changes, like new logins.
type EventTypeEnum string

const (
	EventTypeEnumTransaction EventTypeEnum = "transaction"
	EventTypeEnumAction      EventTypeEnum = "action"
)

func (e *EventTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "transaction":
		fallthrough
	case "action":
		*e = EventTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for EventTypeEnum: %s", s)
	}
}
