// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CreateDatafilesRequestBodyDatafile struct {
	Content  []byte `multipartForm:"content"`
	Datafile string `multipartForm:"name=datafile"`
}

type CreateDatafilesRequestBody struct {
	// Path to datafile
	Datafile *CreateDatafilesRequestBodyDatafile `multipartForm:"file"`
	// Whether to run U21 rules on the datafile after processing
	RunRules *bool `multipartForm:"name=run_rules"`
}

type CreateDatafilesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
