// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/livepeer/livepeer-go/internal/utils"
)

// InputType - Type of service. This is optional and defaults to `url` if
// ŚURL field is provided.
type InputType string

const (
	InputTypeS3 InputType = "s3"
)

func (e InputType) ToPointer() *InputType {
	return &e
}

func (e *InputType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "s3":
		*e = InputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputType: %v", v)
	}
}

// Credentials for the private input video storage
type Credentials struct {
	// Access Key ID
	AccessKeyID string `json:"accessKeyId"`
	// Secret Access Key
	SecretAccessKey string `json:"secretAccessKey"`
}

func (o *Credentials) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *Credentials) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

// Input2 - S3-like storage input video
type Input2 struct {
	// Type of service. This is optional and defaults to `url` if
	// ŚURL field is provided.
	//
	Type InputType `json:"type"`
	// Service endpoint URL (AWS S3 endpoint list: https://docs.aws.amazon.com/general/latest/gr/s3.html, GCP S3 endpoint: https://storage.googleapis.com, Storj: https://gateway.storjshare.io)
	Endpoint string `json:"endpoint"`
	// Bucket with input file
	Bucket string `json:"bucket"`
	// Path to the input file inside the bucket
	Path string `json:"path"`
	// Credentials for the private input video storage
	Credentials Credentials `json:"credentials"`
}

func (o *Input2) GetType() InputType {
	if o == nil {
		return InputType("")
	}
	return o.Type
}

func (o *Input2) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *Input2) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *Input2) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *Input2) GetCredentials() Credentials {
	if o == nil {
		return Credentials{}
	}
	return o.Credentials
}

// Input1 - URL input video
type Input1 struct {
	// URL of the video to transcode
	URL string `json:"url"`
}

func (o *Input1) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type InputUnionType string

const (
	InputUnionTypeInput1 InputUnionType = "input_1"
	InputUnionTypeInput2 InputUnionType = "input_2"
)

type Input struct {
	Input1 *Input1
	Input2 *Input2

	Type InputUnionType
}

func CreateInputInput1(input1 Input1) Input {
	typ := InputUnionTypeInput1

	return Input{
		Input1: &input1,
		Type:   typ,
	}
}

func CreateInputInput2(input2 Input2) Input {
	typ := InputUnionTypeInput2

	return Input{
		Input2: &input2,
		Type:   typ,
	}
}

func (u *Input) UnmarshalJSON(data []byte) error {

	input1 := Input1{}
	if err := utils.UnmarshalJSON(data, &input1, "", true, true); err == nil {
		u.Input1 = &input1
		u.Type = InputUnionTypeInput1
		return nil
	}

	input2 := Input2{}
	if err := utils.UnmarshalJSON(data, &input2, "", true, true); err == nil {
		u.Input2 = &input2
		u.Type = InputUnionTypeInput2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Input) MarshalJSON() ([]byte, error) {
	if u.Input1 != nil {
		return utils.MarshalJSON(u.Input1, "", true)
	}

	if u.Input2 != nil {
		return utils.MarshalJSON(u.Input2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// TranscodePayloadStorageType - Type of service used for output files
type TranscodePayloadStorageType string

const (
	TranscodePayloadStorageTypeWeb3Storage TranscodePayloadStorageType = "web3.storage"
)

func (e TranscodePayloadStorageType) ToPointer() *TranscodePayloadStorageType {
	return &e
}

func (e *TranscodePayloadStorageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web3.storage":
		*e = TranscodePayloadStorageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TranscodePayloadStorageType: %v", v)
	}
}

// TranscodePayloadStorageCredentials - Delegation proof for Livepeer to be able to upload to
// web3.storage
type TranscodePayloadStorageCredentials struct {
	// Base64 encoded UCAN delegation proof
	Proof string `json:"proof"`
}

func (o *TranscodePayloadStorageCredentials) GetProof() string {
	if o == nil {
		return ""
	}
	return o.Proof
}

// Storage2 - Storage for the output files
type Storage2 struct {
	// Type of service used for output files
	Type TranscodePayloadStorageType `json:"type"`
	// Delegation proof for Livepeer to be able to upload to
	// web3.storage
	//
	Credentials TranscodePayloadStorageCredentials `json:"credentials"`
}

func (o *Storage2) GetType() TranscodePayloadStorageType {
	if o == nil {
		return TranscodePayloadStorageType("")
	}
	return o.Type
}

func (o *Storage2) GetCredentials() TranscodePayloadStorageCredentials {
	if o == nil {
		return TranscodePayloadStorageCredentials{}
	}
	return o.Credentials
}

// StorageType - Type of service used for output files
type StorageType string

const (
	StorageTypeS3 StorageType = "s3"
)

func (e StorageType) ToPointer() *StorageType {
	return &e
}

func (e *StorageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "s3":
		*e = StorageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StorageType: %v", v)
	}
}

// StorageCredentials - Credentials for the output video storage
type StorageCredentials struct {
	// Access Key ID
	AccessKeyID string `json:"accessKeyId"`
	// Secret Access Key
	SecretAccessKey string `json:"secretAccessKey"`
}

func (o *StorageCredentials) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *StorageCredentials) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

// Storage1 - Storage for the output files
type Storage1 struct {
	// Type of service used for output files
	Type StorageType `json:"type"`
	// Service endpoint URL (AWS S3 endpoint list: https://docs.aws.amazon.com/general/latest/gr/s3.html, GCP S3 endpoint: https://storage.googleapis.com, Storj: https://gateway.storjshare.io)
	Endpoint string `json:"endpoint"`
	// Bucket with output files
	Bucket string `json:"bucket"`
	// Credentials for the output video storage
	Credentials StorageCredentials `json:"credentials"`
}

func (o *Storage1) GetType() StorageType {
	if o == nil {
		return StorageType("")
	}
	return o.Type
}

func (o *Storage1) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *Storage1) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *Storage1) GetCredentials() StorageCredentials {
	if o == nil {
		return StorageCredentials{}
	}
	return o.Credentials
}

type TranscodePayloadStorageUnionType string

const (
	TranscodePayloadStorageUnionTypeStorage1 TranscodePayloadStorageUnionType = "storage_1"
	TranscodePayloadStorageUnionTypeStorage2 TranscodePayloadStorageUnionType = "storage_2"
)

type TranscodePayloadStorage struct {
	Storage1 *Storage1
	Storage2 *Storage2

	Type TranscodePayloadStorageUnionType
}

func CreateTranscodePayloadStorageStorage1(storage1 Storage1) TranscodePayloadStorage {
	typ := TranscodePayloadStorageUnionTypeStorage1

	return TranscodePayloadStorage{
		Storage1: &storage1,
		Type:     typ,
	}
}

func CreateTranscodePayloadStorageStorage2(storage2 Storage2) TranscodePayloadStorage {
	typ := TranscodePayloadStorageUnionTypeStorage2

	return TranscodePayloadStorage{
		Storage2: &storage2,
		Type:     typ,
	}
}

func (u *TranscodePayloadStorage) UnmarshalJSON(data []byte) error {

	storage2 := Storage2{}
	if err := utils.UnmarshalJSON(data, &storage2, "", true, true); err == nil {
		u.Storage2 = &storage2
		u.Type = TranscodePayloadStorageUnionTypeStorage2
		return nil
	}

	storage1 := Storage1{}
	if err := utils.UnmarshalJSON(data, &storage1, "", true, true); err == nil {
		u.Storage1 = &storage1
		u.Type = TranscodePayloadStorageUnionTypeStorage1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TranscodePayloadStorage) MarshalJSON() ([]byte, error) {
	if u.Storage1 != nil {
		return utils.MarshalJSON(u.Storage1, "", true)
	}

	if u.Storage2 != nil {
		return utils.MarshalJSON(u.Storage2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// Hls - HLS output format
type Hls struct {
	// Path for the HLS output
	Path string `json:"path"`
}

func (o *Hls) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

// Mp4 - MP4 output format
type Mp4 struct {
	// Path for the MP4 output
	Path string `json:"path"`
}

func (o *Mp4) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

// Fmp4 - FMP4 output format
type Fmp4 struct {
	// Path for the FMP4 output
	Path string `json:"path"`
}

func (o *Fmp4) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

// Outputs - Output formats
type Outputs struct {
	// HLS output format
	Hls *Hls `json:"hls,omitempty"`
	// MP4 output format
	Mp4 *Mp4 `json:"mp4,omitempty"`
	// FMP4 output format
	Fmp4 *Fmp4 `json:"fmp4,omitempty"`
}

func (o *Outputs) GetHls() *Hls {
	if o == nil {
		return nil
	}
	return o.Hls
}

func (o *Outputs) GetMp4() *Mp4 {
	if o == nil {
		return nil
	}
	return o.Mp4
}

func (o *Outputs) GetFmp4() *Fmp4 {
	if o == nil {
		return nil
	}
	return o.Fmp4
}

type TranscodePayload struct {
	Input   Input                   `json:"input"`
	Storage TranscodePayloadStorage `json:"storage"`
	// Output formats
	Outputs  Outputs            `json:"outputs"`
	Profiles []TranscodeProfile `json:"profiles,omitempty"`
	// How many seconds the duration of each output segment should be
	TargetSegmentSizeSecs *float64        `json:"targetSegmentSizeSecs,omitempty"`
	CreatorID             *InputCreatorID `json:"creatorId,omitempty"`
	// Decides if the output video should include C2PA signature
	C2pa *bool `json:"c2pa,omitempty"`
}

func (o *TranscodePayload) GetInput() Input {
	if o == nil {
		return Input{}
	}
	return o.Input
}

func (o *TranscodePayload) GetStorage() TranscodePayloadStorage {
	if o == nil {
		return TranscodePayloadStorage{}
	}
	return o.Storage
}

func (o *TranscodePayload) GetOutputs() Outputs {
	if o == nil {
		return Outputs{}
	}
	return o.Outputs
}

func (o *TranscodePayload) GetProfiles() []TranscodeProfile {
	if o == nil {
		return nil
	}
	return o.Profiles
}

func (o *TranscodePayload) GetTargetSegmentSizeSecs() *float64 {
	if o == nil {
		return nil
	}
	return o.TargetSegmentSizeSecs
}

func (o *TranscodePayload) GetCreatorID() *InputCreatorID {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *TranscodePayload) GetC2pa() *bool {
	if o == nil {
		return nil
	}
	return o.C2pa
}
