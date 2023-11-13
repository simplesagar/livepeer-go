// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"livepeer/models/components"
	"net/http"
)

type GetSessionRequest struct {
	// ID of the session
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSessionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSessionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Session *components.Session
}

func (o *GetSessionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSessionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSessionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSessionResponse) GetSession() *components.Session {
	if o == nil {
		return nil
	}
	return o.Session
}
