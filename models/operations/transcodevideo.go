// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type TranscodeVideoResponse struct {
	HTTPMeta components.HTTPMetadata
	// Success
	Task *components.Task
	// Error
	Error *sdkerrors.Error
}

func (o *TranscodeVideoResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *TranscodeVideoResponse) GetTask() *components.Task {
	if o == nil {
		return nil
	}
	return o.Task
}

func (o *TranscodeVideoResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
