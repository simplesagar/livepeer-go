// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"livepeer/internal/utils"
	"livepeer/models/components"
	"net/http"
	"time"
)

type QueryParamFromType string

const (
	QueryParamFromTypeDateTime QueryParamFromType = "date-time"
	QueryParamFromTypeInteger  QueryParamFromType = "integer"
)

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

type GetCreatorMetricsRequest struct {
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

func (o *GetCreatorMetricsRequest) GetFrom() *QueryParamFrom {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *GetCreatorMetricsRequest) GetTo() *QueryParamTo {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *GetCreatorMetricsRequest) GetTimeStep() *QueryParamTimeStep {
	if o == nil {
		return nil
	}
	return o.TimeStep
}

func (o *GetCreatorMetricsRequest) GetAssetID() *string {
	if o == nil {
		return nil
	}
	return o.AssetID
}

func (o *GetCreatorMetricsRequest) GetStreamID() *string {
	if o == nil {
		return nil
	}
	return o.StreamID
}

func (o *GetCreatorMetricsRequest) GetCreatorID() *string {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *GetCreatorMetricsRequest) GetBreakdownBy() []QueryParamBreakdownBy {
	if o == nil {
		return nil
	}
	return o.BreakdownBy
}

type GetCreatorMetricsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A list of Metric objects
	Data []components.ViewershipMetric
}

func (o *GetCreatorMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCreatorMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCreatorMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCreatorMetricsResponse) GetData() []components.ViewershipMetric {
	if o == nil {
		return nil
	}
	return o.Data
}
