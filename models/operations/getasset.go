// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
)

type GetAssetRequest struct {
	// ID of the asset
	AssetID string `pathParam:"style=simple,explode=false,name=assetId"`
}

func (o *GetAssetRequest) GetAssetID() string {
	if o == nil {
		return ""
	}
	return o.AssetID
}

type GetAssetResponse struct {
	HTTPMeta components.HTTPMetadata
	// Success
	Asset *components.Asset
	// Error
	Error *sdkerrors.Error
}

func (o *GetAssetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAssetResponse) GetAsset() *components.Asset {
	if o == nil {
		return nil
	}
	return o.Asset
}

func (o *GetAssetResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
