// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type PutWebhookIDRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PutWebhookIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type PutWebhookIDResponse struct {
	HTTPMeta components.HTTPMetadata
	// Success
	Webhook *components.Webhook
	// Error
	Error *sdkerrors.Error
}

func (o *PutWebhookIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PutWebhookIDResponse) GetWebhook() *components.Webhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}

func (o *PutWebhookIDResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
