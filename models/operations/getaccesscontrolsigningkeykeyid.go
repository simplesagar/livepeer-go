// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type GetAccessControlSigningKeyKeyIDRequest struct {
	// ID of the signing key
	KeyID string `pathParam:"style=simple,explode=false,name=keyId"`
}

func (o *GetAccessControlSigningKeyKeyIDRequest) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

type GetAccessControlSigningKeyKeyIDResponse struct {
	HTTPMeta components.HTTPMetadata
	// Success
	SigningKey *components.SigningKey
	// Error
	Error *sdkerrors.Error
}

func (o *GetAccessControlSigningKeyKeyIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAccessControlSigningKeyKeyIDResponse) GetSigningKey() *components.SigningKey {
	if o == nil {
		return nil
	}
	return o.SigningKey
}

func (o *GetAccessControlSigningKeyKeyIDResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
