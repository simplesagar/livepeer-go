// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// UploadOutput1 - Parameters for the upload task
type UploadOutput1 struct {
	// URL of the asset to "upload"
	URL        *string           `json:"url,omitempty"`
	Encryption *EncryptionOutput `json:"encryption,omitempty"`
	// ID of the original recorded session to avoid re-transcoding
	// of the same content. Only used for import task.
	//
	RecordedSessionID *string `json:"recordedSessionId,omitempty"`
}

func (o *UploadOutput1) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *UploadOutput1) GetEncryption() *EncryptionOutput {
	if o == nil {
		return nil
	}
	return o.Encryption
}

func (o *UploadOutput1) GetRecordedSessionID() *string {
	if o == nil {
		return nil
	}
	return o.RecordedSessionID
}
