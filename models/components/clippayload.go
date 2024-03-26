// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type ClipPayload struct {
	// The playback ID of the stream or stream recording to clip. Asset playback IDs are not supported yet.
	PlaybackID string `json:"playbackId"`
	// The start timestamp of the clip in Unix milliseconds. _See the ClipTrigger in the UI Kit for an example of how this is calculated (for HLS, it uses `Program Date-Time` tags, and for WebRTC, it uses the latency from server to client at stream startup)._
	StartTime float64 `json:"startTime"`
	// The end timestamp of the clip in Unix milliseconds. _See the ClipTrigger in the UI Kit for an example of how this is calculated (for HLS, it uses `Program Date-Time` tags, and for WebRTC, it uses the latency from server to client at stream startup)._
	EndTime *float64 `json:"endTime,omitempty"`
	// The optional friendly name of the clip to create.
	Name *string `json:"name,omitempty"`
	// The optional session ID of the stream to clip. This can be used to clip _recordings_ - if it is not specified, it will clip the ongoing livestream.
	SessionID *string `json:"sessionId,omitempty"`
}

func (o *ClipPayload) GetPlaybackID() string {
	if o == nil {
		return ""
	}
	return o.PlaybackID
}

func (o *ClipPayload) GetStartTime() float64 {
	if o == nil {
		return 0.0
	}
	return o.StartTime
}

func (o *ClipPayload) GetEndTime() *float64 {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *ClipPayload) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ClipPayload) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}
