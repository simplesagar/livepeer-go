// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/livepeer/livepeer-go/internal/utils"
	"github.com/livepeer/livepeer-go/models/components"
	"github.com/livepeer/livepeer-go/models/sdkerrors"
	"time"
)

type QueryParamFromType string

const (
	QueryParamFromTypeDateTime QueryParamFromType = "date-time"
	QueryParamFromTypeInteger  QueryParamFromType = "integer"
)

// QueryParamFrom - Start timestamp for the query range (inclusive)
type QueryParamFrom struct {
	DateTime *time.Time
	Integer  *int64

	Type QueryParamFromType
}

func CreateQueryParamFromDateTime(dateTime time.Time) QueryParamFrom {
	typ := QueryParamFromTypeDateTime

	return QueryParamFrom{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateQueryParamFromInteger(integer int64) QueryParamFrom {
	typ := QueryParamFromTypeInteger

	return QueryParamFrom{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *QueryParamFrom) UnmarshalJSON(data []byte) error {

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = QueryParamFromTypeDateTime
		return nil
	}

	integer := int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = QueryParamFromTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u QueryParamFrom) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type QueryParamToType string

const (
	QueryParamToTypeDateTime QueryParamToType = "date-time"
	QueryParamToTypeInteger  QueryParamToType = "integer"
)

// QueryParamTo - End timestamp for the query range (exclusive)
type QueryParamTo struct {
	DateTime *time.Time
	Integer  *int64

	Type QueryParamToType
}

func CreateQueryParamToDateTime(dateTime time.Time) QueryParamTo {
	typ := QueryParamToTypeDateTime

	return QueryParamTo{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateQueryParamToInteger(integer int64) QueryParamTo {
	typ := QueryParamToTypeInteger

	return QueryParamTo{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *QueryParamTo) UnmarshalJSON(data []byte) error {

	dateTime := time.Time{}
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = &dateTime
		u.Type = QueryParamToTypeDateTime
		return nil
	}

	integer := int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = QueryParamToTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u QueryParamTo) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// QueryParamTimeStep - The time step to aggregate viewership metrics by
type QueryParamTimeStep string

const (
	QueryParamTimeStepHour  QueryParamTimeStep = "hour"
	QueryParamTimeStepDay   QueryParamTimeStep = "day"
	QueryParamTimeStepWeek  QueryParamTimeStep = "week"
	QueryParamTimeStepMonth QueryParamTimeStep = "month"
	QueryParamTimeStepYear  QueryParamTimeStep = "year"
)

func (e QueryParamTimeStep) ToPointer() *QueryParamTimeStep {
	return &e
}

func (e *QueryParamTimeStep) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "hour":
		fallthrough
	case "day":
		fallthrough
	case "week":
		fallthrough
	case "month":
		fallthrough
	case "year":
		*e = QueryParamTimeStep(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamTimeStep: %v", v)
	}
}

type QueryParamBreakdownBy string

const (
	QueryParamBreakdownByDeviceType    QueryParamBreakdownBy = "deviceType"
	QueryParamBreakdownByDevice        QueryParamBreakdownBy = "device"
	QueryParamBreakdownByCPU           QueryParamBreakdownBy = "cpu"
	QueryParamBreakdownByOs            QueryParamBreakdownBy = "os"
	QueryParamBreakdownByBrowser       QueryParamBreakdownBy = "browser"
	QueryParamBreakdownByBrowserEngine QueryParamBreakdownBy = "browserEngine"
	QueryParamBreakdownByContinent     QueryParamBreakdownBy = "continent"
	QueryParamBreakdownByCountry       QueryParamBreakdownBy = "country"
	QueryParamBreakdownBySubdivision   QueryParamBreakdownBy = "subdivision"
	QueryParamBreakdownByTimezone      QueryParamBreakdownBy = "timezone"
	QueryParamBreakdownByViewerID      QueryParamBreakdownBy = "viewerId"
)

func (e QueryParamBreakdownBy) ToPointer() *QueryParamBreakdownBy {
	return &e
}

func (e *QueryParamBreakdownBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deviceType":
		fallthrough
	case "device":
		fallthrough
	case "cpu":
		fallthrough
	case "os":
		fallthrough
	case "browser":
		fallthrough
	case "browserEngine":
		fallthrough
	case "continent":
		fallthrough
	case "country":
		fallthrough
	case "subdivision":
		fallthrough
	case "timezone":
		fallthrough
	case "viewerId":
		*e = QueryParamBreakdownBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamBreakdownBy: %v", v)
	}
}

type GetDataViewsQueryCreatorRequest struct {
	// Start timestamp for the query range (inclusive)
	From *QueryParamFrom `queryParam:"style=form,explode=true,name=from"`
	// End timestamp for the query range (exclusive)
	To *QueryParamTo `queryParam:"style=form,explode=true,name=to"`
	// The time step to aggregate viewership metrics by
	TimeStep *QueryParamTimeStep `queryParam:"style=form,explode=true,name=timeStep"`
	// The asset ID to filter metrics for
	AssetID *string `queryParam:"style=form,explode=true,name=assetId"`
	// The stream ID to filter metrics for
	StreamID *string `queryParam:"style=form,explode=true,name=streamId"`
	// The creator ID to filter the query results
	CreatorID *string `queryParam:"style=form,explode=true,name=creatorId"`
	// The list of fields to break down the query results. Specify this
	// query-string multiple times to break down by multiple fields.
	//
	BreakdownBy []QueryParamBreakdownBy `queryParam:"style=form,explode=true,name=breakdownBy[]"`
}

func (o *GetDataViewsQueryCreatorRequest) GetFrom() *QueryParamFrom {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *GetDataViewsQueryCreatorRequest) GetTo() *QueryParamTo {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *GetDataViewsQueryCreatorRequest) GetTimeStep() *QueryParamTimeStep {
	if o == nil {
		return nil
	}
	return o.TimeStep
}

func (o *GetDataViewsQueryCreatorRequest) GetAssetID() *string {
	if o == nil {
		return nil
	}
	return o.AssetID
}

func (o *GetDataViewsQueryCreatorRequest) GetStreamID() *string {
	if o == nil {
		return nil
	}
	return o.StreamID
}

func (o *GetDataViewsQueryCreatorRequest) GetCreatorID() *string {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *GetDataViewsQueryCreatorRequest) GetBreakdownBy() []QueryParamBreakdownBy {
	if o == nil {
		return nil
	}
	return o.BreakdownBy
}

type GetDataViewsQueryCreatorResponse struct {
	HTTPMeta components.HTTPMetadata
	// A list of Metric objects
	ViewershipMetrics []components.ViewershipMetric
	// Error
	Error *sdkerrors.Error
}

func (o *GetDataViewsQueryCreatorResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetDataViewsQueryCreatorResponse) GetViewershipMetrics() []components.ViewershipMetric {
	if o == nil {
		return nil
	}
	return o.ViewershipMetrics
}

func (o *GetDataViewsQueryCreatorResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
