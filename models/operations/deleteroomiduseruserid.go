// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type DeleteRoomIDUserUserIDRequest struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID string `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *DeleteRoomIDUserUserIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteRoomIDUserUserIDRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type DeleteRoomIDUserUserIDResponse struct {
	HTTPMeta components.HTTPMetadata
	// Error
	Error *sdkerrors.Error
}

func (o *DeleteRoomIDUserUserIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteRoomIDUserUserIDResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
