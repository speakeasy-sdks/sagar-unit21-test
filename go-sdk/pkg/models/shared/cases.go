// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Cases - Top-level case data
type Cases struct {
	// Unique identifier of the case on the customer's platform
	CaseID string `json:"case_id"`
	// Description of the case
	Description *string `json:"description,omitempty"`
	// Labels that describe the outcome of an alert or case investigation. More info in [this knowledge base article about dispositions](https://docs.unit21.ai/docs/concept-dispositions).
	Disposition *string `json:"disposition,omitempty"`
	// Free form text documenting reasoning and investigation notes
	DispositionNotes *string `json:"disposition_notes,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	EndDate *int64 `json:"end_date,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	StartDate int64 `json:"start_date"`
	// Investigation status, either `OPEN` or `ClOSED`
	Status *InvestigationStatusEnum `json:"status,omitempty"`
	// List of string tags, in the format `keyString:valueString` (note that the Key strings are NOT enclosed in `"`)
	Tags []string `json:"tags,omitempty"`
	// Title of the case
	Title string `json:"title"`
}
