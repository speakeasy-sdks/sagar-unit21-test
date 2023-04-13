// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ReadOneSarRequest struct {
	// A Unit21 internally-assigned unique identifier for an object within the Unit21 system. Depending on the endpoint, `unit21_id` can refer to an entity, an event, a case, an alert, a device, etc.
	Unit21ID string `pathParam:"style=simple,explode=false,name=unit21_id"`
}

type ReadOneSarResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ReadOneSar200ApplicationJSONAny interface{}
}
