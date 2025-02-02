// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// RecordingStatus - The status of the recording process of this stream session.
type RecordingStatus string

const (
	RecordingStatusWaiting RecordingStatus = "waiting"
	RecordingStatusReady   RecordingStatus = "ready"
	RecordingStatusNone    RecordingStatus = "none"
)

func (e RecordingStatus) ToPointer() *RecordingStatus {
	return &e
}

func (e *RecordingStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "waiting":
		fallthrough
	case "ready":
		fallthrough
	case "none":
		*e = RecordingStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RecordingStatus: %v", v)
	}
}

type Session struct {
	ID *string `json:"id,omitempty"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Kind *string `json:"kind,omitempty"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	UserID             *string  `json:"userId,omitempty"`
	Name               string   `json:"name"`
	LastSeen           *float64 `json:"lastSeen,omitempty"`
	SourceSegments     *float64 `json:"sourceSegments,omitempty"`
	TranscodedSegments *float64 `json:"transcodedSegments,omitempty"`
	// Duration of all the source segments, sec
	SourceSegmentsDuration *float64 `json:"sourceSegmentsDuration,omitempty"`
	// Duration of all the transcoded segments, sec
	TranscodedSegmentsDuration *float64 `json:"transcodedSegmentsDuration,omitempty"`
	SourceBytes                *float64 `json:"sourceBytes,omitempty"`
	TranscodedBytes            *float64 `json:"transcodedBytes,omitempty"`
	// Rate at which sourceBytes increases (bytes/second)
	IngestRate *float64 `json:"ingestRate,omitempty"`
	// Rate at which transcodedBytes increases (bytes/second)
	OutgoingRate *float64 `json:"outgoingRate,omitempty"`
	// Indicates whether the stream is healthy or not.
	IsHealthy *bool `json:"isHealthy,omitempty"`
	// A string array of human-readable errors describing issues affecting the stream, if any.
	Issues []string `json:"issues,omitempty"`
	// Timestamp (in milliseconds) at which stream object was created
	CreatedAt *float64 `json:"createdAt,omitempty"`
	// Points to parent stream object
	ParentID *string `json:"parentId,omitempty"`
	// Whether the stream should be recorded. Uses default settings. For more customization, create and configure an object store.
	//
	Record *bool `json:"record,omitempty"`
	// The status of the recording process of this stream session.
	RecordingStatus *RecordingStatus `json:"recordingStatus,omitempty"`
	// URL for accessing the recording of this stream session.
	RecordingURL *string `json:"recordingUrl,omitempty"`
	// The URL for the stream session recording packaged in an MP4.
	Mp4URL *string `json:"mp4Url,omitempty"`
	// The playback ID to use with the Playback Info endpoint to retrieve playback URLs.
	PlaybackID *string         `json:"playbackId,omitempty"`
	Profiles   []FfmpegProfile `json:"profiles,omitempty"`
}

func (o *Session) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Session) GetKind() *string {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *Session) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *Session) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Session) GetLastSeen() *float64 {
	if o == nil {
		return nil
	}
	return o.LastSeen
}

func (o *Session) GetSourceSegments() *float64 {
	if o == nil {
		return nil
	}
	return o.SourceSegments
}

func (o *Session) GetTranscodedSegments() *float64 {
	if o == nil {
		return nil
	}
	return o.TranscodedSegments
}

func (o *Session) GetSourceSegmentsDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.SourceSegmentsDuration
}

func (o *Session) GetTranscodedSegmentsDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.TranscodedSegmentsDuration
}

func (o *Session) GetSourceBytes() *float64 {
	if o == nil {
		return nil
	}
	return o.SourceBytes
}

func (o *Session) GetTranscodedBytes() *float64 {
	if o == nil {
		return nil
	}
	return o.TranscodedBytes
}

func (o *Session) GetIngestRate() *float64 {
	if o == nil {
		return nil
	}
	return o.IngestRate
}

func (o *Session) GetOutgoingRate() *float64 {
	if o == nil {
		return nil
	}
	return o.OutgoingRate
}

func (o *Session) GetIsHealthy() *bool {
	if o == nil {
		return nil
	}
	return o.IsHealthy
}

func (o *Session) GetIssues() []string {
	if o == nil {
		return nil
	}
	return o.Issues
}

func (o *Session) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Session) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *Session) GetRecord() *bool {
	if o == nil {
		return nil
	}
	return o.Record
}

func (o *Session) GetRecordingStatus() *RecordingStatus {
	if o == nil {
		return nil
	}
	return o.RecordingStatus
}

func (o *Session) GetRecordingURL() *string {
	if o == nil {
		return nil
	}
	return o.RecordingURL
}

func (o *Session) GetMp4URL() *string {
	if o == nil {
		return nil
	}
	return o.Mp4URL
}

func (o *Session) GetPlaybackID() *string {
	if o == nil {
		return nil
	}
	return o.PlaybackID
}

func (o *Session) GetProfiles() []FfmpegProfile {
	if o == nil {
		return nil
	}
	return o.Profiles
}
