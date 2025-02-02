// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Events string

const (
	EventsStreamStarted           Events = "stream.started"
	EventsStreamDetection         Events = "stream.detection"
	EventsStreamIdle              Events = "stream.idle"
	EventsRecordingReady          Events = "recording.ready"
	EventsRecordingStarted        Events = "recording.started"
	EventsRecordingWaiting        Events = "recording.waiting"
	EventsMultistreamConnected    Events = "multistream.connected"
	EventsMultistreamError        Events = "multistream.error"
	EventsMultistreamDisconnected Events = "multistream.disconnected"
	EventsPlaybackUserNew         Events = "playback.user.new"
	EventsPlaybackAccessControl   Events = "playback.accessControl"
	EventsAssetCreated            Events = "asset.created"
	EventsAssetUpdated            Events = "asset.updated"
	EventsAssetFailed             Events = "asset.failed"
	EventsAssetReady              Events = "asset.ready"
	EventsAssetDeleted            Events = "asset.deleted"
	EventsTaskSpawned             Events = "task.spawned"
	EventsTaskUpdated             Events = "task.updated"
	EventsTaskCompleted           Events = "task.completed"
	EventsTaskFailed              Events = "task.failed"
)

func (e Events) ToPointer() *Events {
	return &e
}

func (e *Events) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stream.started":
		fallthrough
	case "stream.detection":
		fallthrough
	case "stream.idle":
		fallthrough
	case "recording.ready":
		fallthrough
	case "recording.started":
		fallthrough
	case "recording.waiting":
		fallthrough
	case "multistream.connected":
		fallthrough
	case "multistream.error":
		fallthrough
	case "multistream.disconnected":
		fallthrough
	case "playback.user.new":
		fallthrough
	case "playback.accessControl":
		fallthrough
	case "asset.created":
		fallthrough
	case "asset.updated":
		fallthrough
	case "asset.failed":
		fallthrough
	case "asset.ready":
		fallthrough
	case "asset.deleted":
		fallthrough
	case "task.spawned":
		fallthrough
	case "task.updated":
		fallthrough
	case "task.completed":
		fallthrough
	case "task.failed":
		*e = Events(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Events: %v", v)
	}
}

// LastFailure - failure timestamp and error message with status code
type LastFailure struct {
	// Timestamp (in milliseconds) at which the webhook last failed
	Timestamp *float64 `json:"timestamp,omitempty"`
	// Webhook failure error message
	Error *string `json:"error,omitempty"`
	// Webhook failure response
	Response *string `json:"response,omitempty"`
	// Webhook failure status code
	StatusCode *float64 `json:"statusCode,omitempty"`
}

func (o *LastFailure) GetTimestamp() *float64 {
	if o == nil {
		return nil
	}
	return o.Timestamp
}

func (o *LastFailure) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *LastFailure) GetResponse() *string {
	if o == nil {
		return nil
	}
	return o.Response
}

func (o *LastFailure) GetStatusCode() *float64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

// Status - status of webhook
type Status struct {
	// failure timestamp and error message with status code
	LastFailure *LastFailure `json:"lastFailure,omitempty"`
	// Timestamp (in milliseconds) at which the webhook last was
	// triggered
	//
	LastTriggeredAt *float64 `json:"lastTriggeredAt,omitempty"`
}

func (o *Status) GetLastFailure() *LastFailure {
	if o == nil {
		return nil
	}
	return o.LastFailure
}

func (o *Status) GetLastTriggeredAt() *float64 {
	if o == nil {
		return nil
	}
	return o.LastTriggeredAt
}

type Webhook struct {
	ID   *string `json:"id,omitempty"`
	Name string  `json:"name"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Kind *string `json:"kind,omitempty"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	UserID *string `json:"userId,omitempty"`
	// Timestamp (in milliseconds) at which stream object was created
	CreatedAt *float64 `json:"createdAt,omitempty"`
	Events    []Events `json:"events,omitempty"`
	URL       string   `json:"url"`
	// streamId of the stream on which the webhook is applied
	StreamID *string `json:"streamId,omitempty"`
	// status of webhook
	Status *Status `json:"status,omitempty"`
}

func (o *Webhook) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Webhook) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Webhook) GetKind() *string {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *Webhook) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *Webhook) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Webhook) GetEvents() []Events {
	if o == nil {
		return nil
	}
	return o.Events
}

func (o *Webhook) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *Webhook) GetStreamID() *string {
	if o == nil {
		return nil
	}
	return o.StreamID
}

func (o *Webhook) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type WebhookInput struct {
	Name   string   `json:"name"`
	Events []Events `json:"events,omitempty"`
	URL    string   `json:"url"`
	// shared secret used to sign the webhook payload
	SharedSecret *string `json:"sharedSecret,omitempty"`
	// streamId of the stream on which the webhook is applied
	StreamID *string `json:"streamId,omitempty"`
}

func (o *WebhookInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *WebhookInput) GetEvents() []Events {
	if o == nil {
		return nil
	}
	return o.Events
}

func (o *WebhookInput) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *WebhookInput) GetSharedSecret() *string {
	if o == nil {
		return nil
	}
	return o.SharedSecret
}

func (o *WebhookInput) GetStreamID() *string {
	if o == nil {
		return nil
	}
	return o.StreamID
}
