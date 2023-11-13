// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"livepeer/internal/utils"
)

// AssetType - Type of the asset.
type AssetType string

const (
	AssetTypeVideo AssetType = "video"
	AssetTypeAudio AssetType = "audio"
)

func (e AssetType) ToPointer() *AssetType {
	return &e
}

func (e *AssetType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "video":
		fallthrough
	case "audio":
		*e = AssetType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssetType: %v", v)
	}
}

type AssetSchemasSource3Type string

const (
	AssetSchemasSource3TypeDirectUpload AssetSchemasSource3Type = "directUpload"
	AssetSchemasSource3TypeClip         AssetSchemasSource3Type = "clip"
)

func (e AssetSchemasSource3Type) ToPointer() *AssetSchemasSource3Type {
	return &e
}

func (e *AssetSchemasSource3Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "directUpload":
		fallthrough
	case "clip":
		*e = AssetSchemasSource3Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssetSchemasSource3Type: %v", v)
	}
}

type Asset3 struct {
	Type       AssetSchemasSource3Type `json:"type"`
	Encryption *EncryptionOutput       `json:"encryption,omitempty"`
}

func (o *Asset3) GetType() AssetSchemasSource3Type {
	if o == nil {
		return AssetSchemasSource3Type("")
	}
	return o.Type
}

func (o *Asset3) GetEncryption() *EncryptionOutput {
	if o == nil {
		return nil
	}
	return o.Encryption
}

type AssetSchemasSourceType string

const (
	AssetSchemasSourceTypeRecording AssetSchemasSourceType = "recording"
)

func (e AssetSchemasSourceType) ToPointer() *AssetSchemasSourceType {
	return &e
}

func (e *AssetSchemasSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "recording":
		*e = AssetSchemasSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssetSchemasSourceType: %v", v)
	}
}

type Two struct {
	Type AssetSchemasSourceType `json:"type"`
	// ID of the session from which this asset was created
	SessionID string `json:"sessionId"`
}

func (o *Two) GetType() AssetSchemasSourceType {
	if o == nil {
		return AssetSchemasSourceType("")
	}
	return o.Type
}

func (o *Two) GetSessionID() string {
	if o == nil {
		return ""
	}
	return o.SessionID
}

type AssetSchemasType string

const (
	AssetSchemasTypeURL AssetSchemasType = "url"
)

func (e AssetSchemasType) ToPointer() *AssetSchemasType {
	return &e
}

func (e *AssetSchemasType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "url":
		*e = AssetSchemasType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssetSchemasType: %v", v)
	}
}

type Asset1Output struct {
	Type AssetSchemasType `json:"type"`
	// URL from which the asset was uploaded
	URL string `json:"url"`
	// Gateway URL from asset if parsed from provided URL on upload.
	GatewayURL *string           `json:"gatewayUrl,omitempty"`
	Encryption *EncryptionOutput `json:"encryption,omitempty"`
}

func (o *Asset1Output) GetType() AssetSchemasType {
	if o == nil {
		return AssetSchemasType("")
	}
	return o.Type
}

func (o *Asset1Output) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *Asset1Output) GetGatewayURL() *string {
	if o == nil {
		return nil
	}
	return o.GatewayURL
}

func (o *Asset1Output) GetEncryption() *EncryptionOutput {
	if o == nil {
		return nil
	}
	return o.Encryption
}

type AssetSourceType string

const (
	AssetSourceTypeAsset1Output AssetSourceType = "asset_1_output"
	AssetSourceTypeTwo          AssetSourceType = "2"
	AssetSourceTypeAsset3       AssetSourceType = "asset_3"
)

type AssetSource struct {
	Asset1Output *Asset1Output
	Two          *Two
	Asset3       *Asset3

	Type AssetSourceType
}

func CreateAssetSourceAsset1Output(asset1Output Asset1Output) AssetSource {
	typ := AssetSourceTypeAsset1Output

	return AssetSource{
		Asset1Output: &asset1Output,
		Type:         typ,
	}
}

func CreateAssetSourceTwo(two Two) AssetSource {
	typ := AssetSourceTypeTwo

	return AssetSource{
		Two:  &two,
		Type: typ,
	}
}

func CreateAssetSourceAsset3(asset3 Asset3) AssetSource {
	typ := AssetSourceTypeAsset3

	return AssetSource{
		Asset3: &asset3,
		Type:   typ,
	}
}

func (u *AssetSource) UnmarshalJSON(data []byte) error {

	two := Two{}
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = AssetSourceTypeTwo
		return nil
	}

	asset3 := Asset3{}
	if err := utils.UnmarshalJSON(data, &asset3, "", true, true); err == nil {
		u.Asset3 = &asset3
		u.Type = AssetSourceTypeAsset3
		return nil
	}

	asset1Output := Asset1Output{}
	if err := utils.UnmarshalJSON(data, &asset1Output, "", true, true); err == nil {
		u.Asset1Output = &asset1Output
		u.Type = AssetSourceTypeAsset1Output
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AssetSource) MarshalJSON() ([]byte, error) {
	if u.Asset1Output != nil {
		return utils.MarshalJSON(u.Asset1Output, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	if u.Asset3 != nil {
		return utils.MarshalJSON(u.Asset3, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// AssetNftMetadataTemplate - Name of the NFT metadata template to export. 'player'
// will embed the Livepeer Player on the NFT while 'file'
// will reference only the immutable MP4 files.
type AssetNftMetadataTemplate string

const (
	AssetNftMetadataTemplateFile   AssetNftMetadataTemplate = "file"
	AssetNftMetadataTemplatePlayer AssetNftMetadataTemplate = "player"
)

func (e AssetNftMetadataTemplate) ToPointer() *AssetNftMetadataTemplate {
	return &e
}

func (e *AssetNftMetadataTemplate) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "file":
		fallthrough
	case "player":
		*e = AssetNftMetadataTemplate(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssetNftMetadataTemplate: %v", v)
	}
}

// AssetNftMetadata - Additional data to add to the NFT metadata exported to
// IPFS. Will be deep merged with the default metadata
// exported.
type AssetNftMetadata struct {
}

type AssetSpec struct {
	// Name of the NFT metadata template to export. 'player'
	// will embed the Livepeer Player on the NFT while 'file'
	// will reference only the immutable MP4 files.
	//
	NftMetadataTemplate *AssetNftMetadataTemplate `default:"file" json:"nftMetadataTemplate"`
	// Additional data to add to the NFT metadata exported to
	// IPFS. Will be deep merged with the default metadata
	// exported.
	//
	NftMetadata *AssetNftMetadata `json:"nftMetadata,omitempty"`
}

func (a AssetSpec) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssetSpec) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AssetSpec) GetNftMetadataTemplate() *AssetNftMetadataTemplate {
	if o == nil {
		return nil
	}
	return o.NftMetadataTemplate
}

func (o *AssetSpec) GetNftMetadata() *AssetNftMetadata {
	if o == nil {
		return nil
	}
	return o.NftMetadata
}

type AssetIpfs struct {
	Spec *AssetSpec `json:"spec,omitempty"`
	// CID of the file on IPFS
	Cid *string `json:"cid,omitempty"`
	// URL with IPFS scheme for the file
	URL *string `json:"url,omitempty"`
	// URL to access file via HTTP through an IPFS gateway
	GatewayURL  *string       `json:"gatewayUrl,omitempty"`
	NftMetadata *IpfsFileInfo `json:"nftMetadata,omitempty"`
	// Timestamp (in milliseconds) at which IPFS export task was
	// updated
	//
	UpdatedAt *float64 `json:"updatedAt,omitempty"`
}

func (o *AssetIpfs) GetSpec() *AssetSpec {
	if o == nil {
		return nil
	}
	return o.Spec
}

func (o *AssetIpfs) GetCid() *string {
	if o == nil {
		return nil
	}
	return o.Cid
}

func (o *AssetIpfs) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *AssetIpfs) GetGatewayURL() *string {
	if o == nil {
		return nil
	}
	return o.GatewayURL
}

func (o *AssetIpfs) GetNftMetadata() *IpfsFileInfo {
	if o == nil {
		return nil
	}
	return o.NftMetadata
}

func (o *AssetIpfs) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type AssetStorage struct {
	Ipfs   *AssetIpfs     `json:"ipfs,omitempty"`
	Status *StorageStatus `json:"status,omitempty"`
}

func (o *AssetStorage) GetIpfs() *AssetIpfs {
	if o == nil {
		return nil
	}
	return o.Ipfs
}

func (o *AssetStorage) GetStatus() *StorageStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

// AssetPhase - Phase of the asset
type AssetPhase string

const (
	AssetPhaseUploading  AssetPhase = "uploading"
	AssetPhaseWaiting    AssetPhase = "waiting"
	AssetPhaseProcessing AssetPhase = "processing"
	AssetPhaseReady      AssetPhase = "ready"
	AssetPhaseFailed     AssetPhase = "failed"
)

func (e AssetPhase) ToPointer() *AssetPhase {
	return &e
}

func (e *AssetPhase) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "uploading":
		fallthrough
	case "waiting":
		fallthrough
	case "processing":
		fallthrough
	case "ready":
		fallthrough
	case "failed":
		*e = AssetPhase(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssetPhase: %v", v)
	}
}

// AssetStatus - Status of the asset
type AssetStatus struct {
	// Phase of the asset
	Phase AssetPhase `json:"phase"`
	// Timestamp (in milliseconds) at which the asset was last updated
	UpdatedAt float64 `json:"updatedAt"`
	// Current progress of the task creating this asset.
	Progress *float64 `json:"progress,omitempty"`
	// Error message if the asset creation failed.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

func (o *AssetStatus) GetPhase() AssetPhase {
	if o == nil {
		return AssetPhase("")
	}
	return o.Phase
}

func (o *AssetStatus) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

func (o *AssetStatus) GetProgress() *float64 {
	if o == nil {
		return nil
	}
	return o.Progress
}

func (o *AssetStatus) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

type Hash struct {
	// Hash of the asset
	Hash *string `json:"hash,omitempty"`
	// Hash algorithm used to compute the hash
	Algorithm *string `json:"algorithm,omitempty"`
}

func (o *Hash) GetHash() *string {
	if o == nil {
		return nil
	}
	return o.Hash
}

func (o *Hash) GetAlgorithm() *string {
	if o == nil {
		return nil
	}
	return o.Algorithm
}

// AssetSchemasVideoSpecType - type of track
type AssetSchemasVideoSpecType string

const (
	AssetSchemasVideoSpecTypeVideo AssetSchemasVideoSpecType = "video"
	AssetSchemasVideoSpecTypeAudio AssetSchemasVideoSpecType = "audio"
)

func (e AssetSchemasVideoSpecType) ToPointer() *AssetSchemasVideoSpecType {
	return &e
}

func (e *AssetSchemasVideoSpecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "video":
		fallthrough
	case "audio":
		*e = AssetSchemasVideoSpecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssetSchemasVideoSpecType: %v", v)
	}
}

type Tracks struct {
	// type of track
	Type AssetSchemasVideoSpecType `json:"type"`
	// Codec of the track
	Codec string `json:"codec"`
	// Start time of the track in seconds
	StartTime *float64 `json:"startTime,omitempty"`
	// Duration of the track in seconds
	Duration *float64 `json:"duration,omitempty"`
	// Bitrate of the track in bits per second
	Bitrate *float64 `json:"bitrate,omitempty"`
	// Width of the track - only for video tracks
	Width *float64 `json:"width,omitempty"`
	// Height of the track - only for video tracks
	Height *float64 `json:"height,omitempty"`
	// Pixel format of the track - only for video tracks
	PixelFormat *string `json:"pixelFormat,omitempty"`
	// Frame rate of the track - only for video tracks
	Fps *float64 `json:"fps,omitempty"`
	// Amount of audio channels in the track
	Channels *float64 `json:"channels,omitempty"`
	// Sample rate of the track in samples per second - only for
	// audio tracks
	//
	SampleRate *float64 `json:"sampleRate,omitempty"`
	// Bit depth of the track - only for audio tracks
	BitDepth *float64 `json:"bitDepth,omitempty"`
}

func (o *Tracks) GetType() AssetSchemasVideoSpecType {
	if o == nil {
		return AssetSchemasVideoSpecType("")
	}
	return o.Type
}

func (o *Tracks) GetCodec() string {
	if o == nil {
		return ""
	}
	return o.Codec
}

func (o *Tracks) GetStartTime() *float64 {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *Tracks) GetDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *Tracks) GetBitrate() *float64 {
	if o == nil {
		return nil
	}
	return o.Bitrate
}

func (o *Tracks) GetWidth() *float64 {
	if o == nil {
		return nil
	}
	return o.Width
}

func (o *Tracks) GetHeight() *float64 {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *Tracks) GetPixelFormat() *string {
	if o == nil {
		return nil
	}
	return o.PixelFormat
}

func (o *Tracks) GetFps() *float64 {
	if o == nil {
		return nil
	}
	return o.Fps
}

func (o *Tracks) GetChannels() *float64 {
	if o == nil {
		return nil
	}
	return o.Channels
}

func (o *Tracks) GetSampleRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SampleRate
}

func (o *Tracks) GetBitDepth() *float64 {
	if o == nil {
		return nil
	}
	return o.BitDepth
}

// VideoSpec - Video metadata
type VideoSpec struct {
	// Format of the asset
	Format *string `json:"format,omitempty"`
	// Duration of the asset in seconds (float)
	Duration *float64 `json:"duration,omitempty"`
	// Bitrate of the video in bits per second
	Bitrate *float64 `json:"bitrate,omitempty"`
	// List of tracks associated with the asset when the format
	// contemplates them (e.g. mp4)
	//
	Tracks []Tracks `json:"tracks,omitempty"`
}

func (o *VideoSpec) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *VideoSpec) GetDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *VideoSpec) GetBitrate() *float64 {
	if o == nil {
		return nil
	}
	return o.Bitrate
}

func (o *VideoSpec) GetTracks() []Tracks {
	if o == nil {
		return nil
	}
	return o.Tracks
}

type Asset struct {
	ID string `json:"id"`
	// Type of the asset.
	Type *AssetType `json:"type,omitempty"`
	// Used to form playback URL and storage folder
	PlaybackID *string `json:"playbackId,omitempty"`
	// URL for HLS playback
	PlaybackURL *string `json:"playbackUrl,omitempty"`
	// URL to manually download the asset if desired
	DownloadURL *string `json:"downloadUrl,omitempty"`
	// Whether the playback policy for a asset or stream is public or signed
	PlaybackPolicy *PlaybackPolicy `json:"playbackPolicy,omitempty"`
	Source         AssetSource     `json:"source"`
	CreatorID      *CreatorID      `json:"creatorId,omitempty"`
	Storage        *AssetStorage   `json:"storage,omitempty"`
	// Status of the asset
	Status *AssetStatus `json:"status,omitempty"`
	// Name of the asset. This is not necessarily the filename, can be a
	// custom name or title
	//
	Name string `json:"name"`
	// Timestamp (in milliseconds) at which asset was created
	CreatedAt *float64 `json:"createdAt,omitempty"`
	// Size of the asset in bytes
	Size *float64 `json:"size,omitempty"`
	// Hash of the asset
	Hash []Hash `json:"hash,omitempty"`
	// Video metadata
	VideoSpec *VideoSpec `json:"videoSpec,omitempty"`
}

func (o *Asset) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Asset) GetType() *AssetType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Asset) GetPlaybackID() *string {
	if o == nil {
		return nil
	}
	return o.PlaybackID
}

func (o *Asset) GetPlaybackURL() *string {
	if o == nil {
		return nil
	}
	return o.PlaybackURL
}

func (o *Asset) GetDownloadURL() *string {
	if o == nil {
		return nil
	}
	return o.DownloadURL
}

func (o *Asset) GetPlaybackPolicy() *PlaybackPolicy {
	if o == nil {
		return nil
	}
	return o.PlaybackPolicy
}

func (o *Asset) GetSource() AssetSource {
	if o == nil {
		return AssetSource{}
	}
	return o.Source
}

func (o *Asset) GetCreatorID() *CreatorID {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *Asset) GetStorage() *AssetStorage {
	if o == nil {
		return nil
	}
	return o.Storage
}

func (o *Asset) GetStatus() *AssetStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Asset) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Asset) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Asset) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *Asset) GetHash() []Hash {
	if o == nil {
		return nil
	}
	return o.Hash
}

func (o *Asset) GetVideoSpec() *VideoSpec {
	if o == nil {
		return nil
	}
	return o.VideoSpec
}

type Three struct {
	Type       AssetSchemasSource3Type `json:"type"`
	Encryption *Encryption             `json:"encryption,omitempty"`
}

func (o *Three) GetType() AssetSchemasSource3Type {
	if o == nil {
		return AssetSchemasSource3Type("")
	}
	return o.Type
}

func (o *Three) GetEncryption() *Encryption {
	if o == nil {
		return nil
	}
	return o.Encryption
}

type Asset1 struct {
	Type AssetSchemasType `json:"type"`
	// URL from which the asset was uploaded
	URL string `json:"url"`
	// Gateway URL from asset if parsed from provided URL on upload.
	GatewayURL *string     `json:"gatewayUrl,omitempty"`
	Encryption *Encryption `json:"encryption,omitempty"`
}

func (o *Asset1) GetType() AssetSchemasType {
	if o == nil {
		return AssetSchemasType("")
	}
	return o.Type
}

func (o *Asset1) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *Asset1) GetGatewayURL() *string {
	if o == nil {
		return nil
	}
	return o.GatewayURL
}

func (o *Asset1) GetEncryption() *Encryption {
	if o == nil {
		return nil
	}
	return o.Encryption
}

type SourceType string

const (
	SourceTypeAsset1 SourceType = "asset_1"
	SourceTypeTwo    SourceType = "2"
	SourceTypeThree  SourceType = "3"
)

type Source struct {
	Asset1 *Asset1
	Two    *Two
	Three  *Three

	Type SourceType
}

func CreateSourceAsset1(asset1 Asset1) Source {
	typ := SourceTypeAsset1

	return Source{
		Asset1: &asset1,
		Type:   typ,
	}
}

func CreateSourceTwo(two Two) Source {
	typ := SourceTypeTwo

	return Source{
		Two:  &two,
		Type: typ,
	}
}

func CreateSourceThree(three Three) Source {
	typ := SourceTypeThree

	return Source{
		Three: &three,
		Type:  typ,
	}
}

func (u *Source) UnmarshalJSON(data []byte) error {

	two := Two{}
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = SourceTypeTwo
		return nil
	}

	three := Three{}
	if err := utils.UnmarshalJSON(data, &three, "", true, true); err == nil {
		u.Three = &three
		u.Type = SourceTypeThree
		return nil
	}

	asset1 := Asset1{}
	if err := utils.UnmarshalJSON(data, &asset1, "", true, true); err == nil {
		u.Asset1 = &asset1
		u.Type = SourceTypeAsset1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Source) MarshalJSON() ([]byte, error) {
	if u.Asset1 != nil {
		return utils.MarshalJSON(u.Asset1, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	if u.Three != nil {
		return utils.MarshalJSON(u.Three, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type AssetIpfsInput struct {
	Spec *AssetSpec `json:"spec,omitempty"`
	// CID of the file on IPFS
	Cid         *string            `json:"cid,omitempty"`
	NftMetadata *IpfsFileInfoInput `json:"nftMetadata,omitempty"`
}

func (o *AssetIpfsInput) GetSpec() *AssetSpec {
	if o == nil {
		return nil
	}
	return o.Spec
}

func (o *AssetIpfsInput) GetCid() *string {
	if o == nil {
		return nil
	}
	return o.Cid
}

func (o *AssetIpfsInput) GetNftMetadata() *IpfsFileInfoInput {
	if o == nil {
		return nil
	}
	return o.NftMetadata
}

type AssetStorageInput struct {
	Ipfs *AssetIpfsInput `json:"ipfs,omitempty"`
}

func (o *AssetStorageInput) GetIpfs() *AssetIpfsInput {
	if o == nil {
		return nil
	}
	return o.Ipfs
}

type AssetInput struct {
	// Type of the asset.
	Type *AssetType `json:"type,omitempty"`
	// Used to form playback URL and storage folder
	PlaybackID *string `json:"playbackId,omitempty"`
	// Whether to generate MP4s for the asset.
	StaticMp4 *bool `json:"staticMp4,omitempty"`
	// Whether the playback policy for a asset or stream is public or signed
	PlaybackPolicy *PlaybackPolicy    `json:"playbackPolicy,omitempty"`
	Source         Source             `json:"source"`
	CreatorID      *CreatorID         `json:"creatorId,omitempty"`
	Storage        *AssetStorageInput `json:"storage,omitempty"`
	// Name of the asset. This is not necessarily the filename, can be a
	// custom name or title
	//
	Name string `json:"name"`
	// Hash of the asset
	Hash []Hash `json:"hash,omitempty"`
}

func (o *AssetInput) GetType() *AssetType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AssetInput) GetPlaybackID() *string {
	if o == nil {
		return nil
	}
	return o.PlaybackID
}

func (o *AssetInput) GetStaticMp4() *bool {
	if o == nil {
		return nil
	}
	return o.StaticMp4
}

func (o *AssetInput) GetPlaybackPolicy() *PlaybackPolicy {
	if o == nil {
		return nil
	}
	return o.PlaybackPolicy
}

func (o *AssetInput) GetSource() Source {
	if o == nil {
		return Source{}
	}
	return o.Source
}

func (o *AssetInput) GetCreatorID() *CreatorID {
	if o == nil {
		return nil
	}
	return o.CreatorID
}

func (o *AssetInput) GetStorage() *AssetStorageInput {
	if o == nil {
		return nil
	}
	return o.Storage
}

func (o *AssetInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AssetInput) GetHash() []Hash {
	if o == nil {
		return nil
	}
	return o.Hash
}
