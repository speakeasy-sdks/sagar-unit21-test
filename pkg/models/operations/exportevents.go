// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// ExportEventsRequestBodyFilters - Filter to narrow down events in export
type ExportEventsRequestBodyFilters struct {
	// Currency used in the transaction.
	Currency *string `json:"currency,omitempty"`
	// Event creation date end.
	EndDate *string `json:"end_date,omitempty"`
	// Numerical IDs of the entities in the event.
	EntityIds []int64 `json:"entity_ids,omitempty"`
	// Maximum amount in the event transaction.
	MaximumAmount *int64 `json:"maximum_amount,omitempty"`
	// Minimum amount in the event transaction.
	MinimumAmount *int64 `json:"minimum_amount,omitempty"`
	// Event creation date start.
	StartDate *string `json:"start_date,omitempty"`
	// Status of the event.
	Status *string `json:"status,omitempty"`
	// Status for the events.
	Statuses []string `json:"statuses,omitempty"`
	// Numerical IDs of the tags.
	TagIds []int64 `json:"tag_ids,omitempty"`
}

type ExportEventsRequestBody struct {
	// Array of the unique identifiers of the event IDs.
	EventIds []int64 `json:"event_ids,omitempty"`
	// Filter to narrow down events in export
	Filters *ExportEventsRequestBodyFilters `json:"filters,omitempty"`
}

type ExportEventsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
