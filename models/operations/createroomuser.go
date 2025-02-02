// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type CreateRoomUserRequest struct {
	ID              string                     `pathParam:"style=simple,explode=false,name=id"`
	RoomUserPayload components.RoomUserPayload `request:"mediaType=application/json"`
}

func (o *CreateRoomUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateRoomUserRequest) GetRoomUserPayload() components.RoomUserPayload {
	if o == nil {
		return components.RoomUserPayload{}
	}
	return o.RoomUserPayload
}

type CreateRoomUserResponse struct {
	HTTPMeta components.HTTPMetadata
	// Success
	RoomUserResponse *components.RoomUserResponse
	// Error
	Error *sdkerrors.Error
}

func (o *CreateRoomUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateRoomUserResponse) GetRoomUserResponse() *components.RoomUserResponse {
	if o == nil {
		return nil
	}
	return o.RoomUserResponse
}

func (o *CreateRoomUserResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
