// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: release.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ReleaseList)(nil)
var _ json.Unmarshaler = (*ReleaseList)(nil)
var _ json.Marshaler = (*Mode)(nil)
var _ json.Unmarshaler = (*Mode)(nil)
var _ json.Marshaler = (*ReleaseCreateRequest)(nil)
var _ json.Unmarshaler = (*ReleaseCreateRequest)(nil)
var _ json.Marshaler = (*ReleaseResource)(nil)
var _ json.Unmarshaler = (*ReleaseResource)(nil)
var _ json.Marshaler = (*ReleaseCreateResponse)(nil)
var _ json.Unmarshaler = (*ReleaseCreateResponse)(nil)
var _ json.Marshaler = (*ReleaseCreateResponseData)(nil)
var _ json.Unmarshaler = (*ReleaseCreateResponseData)(nil)
var _ json.Marshaler = (*ReleaseUpdateRequest)(nil)
var _ json.Unmarshaler = (*ReleaseUpdateRequest)(nil)
var _ json.Marshaler = (*ReleaseDataResponse)(nil)
var _ json.Unmarshaler = (*ReleaseDataResponse)(nil)
var _ json.Marshaler = (*ReleaseUserDataResponse)(nil)
var _ json.Unmarshaler = (*ReleaseUserDataResponse)(nil)
var _ json.Marshaler = (*ReleaseUpdateResponse)(nil)
var _ json.Unmarshaler = (*ReleaseUpdateResponse)(nil)
var _ json.Marshaler = (*ReleaseReferenceUpdateRequest)(nil)
var _ json.Unmarshaler = (*ReleaseReferenceUpdateRequest)(nil)
var _ json.Marshaler = (*GetIosPlistRequest)(nil)
var _ json.Unmarshaler = (*GetIosPlistRequest)(nil)
var _ json.Marshaler = (*GetIosPlistResponse)(nil)
var _ json.Unmarshaler = (*GetIosPlistResponse)(nil)
var _ json.Marshaler = (*ReleaseDeleteRequest)(nil)
var _ json.Unmarshaler = (*ReleaseDeleteRequest)(nil)
var _ json.Marshaler = (*ReleaseDeleteResponse)(nil)
var _ json.Unmarshaler = (*ReleaseDeleteResponse)(nil)
var _ json.Marshaler = (*ReleaseGetRequest)(nil)
var _ json.Unmarshaler = (*ReleaseGetRequest)(nil)
var _ json.Marshaler = (*ReleaseGetResponse)(nil)
var _ json.Unmarshaler = (*ReleaseGetResponse)(nil)
var _ json.Marshaler = (*ModeSummary)(nil)
var _ json.Unmarshaler = (*ModeSummary)(nil)
var _ json.Marshaler = (*ReleaseSummaryArray)(nil)
var _ json.Unmarshaler = (*ReleaseSummaryArray)(nil)
var _ json.Marshaler = (*Tag)(nil)
var _ json.Unmarshaler = (*Tag)(nil)
var _ json.Marshaler = (*ReleaseGetResponseData)(nil)
var _ json.Unmarshaler = (*ReleaseGetResponseData)(nil)
var _ json.Marshaler = (*AddonInfo)(nil)
var _ json.Unmarshaler = (*AddonInfo)(nil)
var _ json.Marshaler = (*ServiceImagePair)(nil)
var _ json.Unmarshaler = (*ServiceImagePair)(nil)
var _ json.Marshaler = (*ApplicationReleaseSummary)(nil)
var _ json.Unmarshaler = (*ApplicationReleaseSummary)(nil)
var _ json.Marshaler = (*ReleaseListRequest)(nil)
var _ json.Unmarshaler = (*ReleaseListRequest)(nil)
var _ json.Marshaler = (*ReleaseListResponse)(nil)
var _ json.Unmarshaler = (*ReleaseListResponse)(nil)
var _ json.Marshaler = (*ListReleaseNameRequest)(nil)
var _ json.Unmarshaler = (*ListReleaseNameRequest)(nil)
var _ json.Marshaler = (*ListReleaseNameResponse)(nil)
var _ json.Unmarshaler = (*ListReleaseNameResponse)(nil)
var _ json.Marshaler = (*ReleaseListResponseData)(nil)
var _ json.Unmarshaler = (*ReleaseListResponseData)(nil)
var _ json.Marshaler = (*ReleaseData)(nil)
var _ json.Unmarshaler = (*ReleaseData)(nil)
var _ json.Marshaler = (*ReleaseNameListRequest)(nil)
var _ json.Unmarshaler = (*ReleaseNameListRequest)(nil)
var _ json.Marshaler = (*ReleaseNameListResponse)(nil)
var _ json.Unmarshaler = (*ReleaseNameListResponse)(nil)
var _ json.Marshaler = (*GetLatestReleasesRequest)(nil)
var _ json.Unmarshaler = (*GetLatestReleasesRequest)(nil)
var _ json.Marshaler = (*GetLatestReleasesResponse)(nil)
var _ json.Unmarshaler = (*GetLatestReleasesResponse)(nil)
var _ json.Marshaler = (*GetLatestReleasesResponseData)(nil)
var _ json.Unmarshaler = (*GetLatestReleasesResponseData)(nil)
var _ json.Marshaler = (*ReleaseGCRequest)(nil)
var _ json.Unmarshaler = (*ReleaseGCRequest)(nil)
var _ json.Marshaler = (*ReleaseUploadRequest)(nil)
var _ json.Unmarshaler = (*ReleaseUploadRequest)(nil)
var _ json.Marshaler = (*ReleaseUploadResponse)(nil)
var _ json.Unmarshaler = (*ReleaseUploadResponse)(nil)
var _ json.Marshaler = (*ParseReleaseFileRequest)(nil)
var _ json.Unmarshaler = (*ParseReleaseFileRequest)(nil)
var _ json.Marshaler = (*ParseReleaseFileResponse)(nil)
var _ json.Unmarshaler = (*ParseReleaseFileResponse)(nil)
var _ json.Marshaler = (*ParseReleaseFileResponseData)(nil)
var _ json.Unmarshaler = (*ParseReleaseFileResponseData)(nil)
var _ json.Marshaler = (*FormalReleaseRequest)(nil)
var _ json.Unmarshaler = (*FormalReleaseRequest)(nil)
var _ json.Marshaler = (*FormalReleaseResponse)(nil)
var _ json.Unmarshaler = (*FormalReleaseResponse)(nil)
var _ json.Marshaler = (*FormalReleasesRequest)(nil)
var _ json.Unmarshaler = (*FormalReleasesRequest)(nil)
var _ json.Marshaler = (*FormalReleasesResponse)(nil)
var _ json.Unmarshaler = (*FormalReleasesResponse)(nil)
var _ json.Marshaler = (*ReleasesDeleteRequest)(nil)
var _ json.Unmarshaler = (*ReleasesDeleteRequest)(nil)
var _ json.Marshaler = (*ReleasesDeleteResponse)(nil)
var _ json.Unmarshaler = (*ReleasesDeleteResponse)(nil)
var _ json.Marshaler = (*CheckVersionRequest)(nil)
var _ json.Unmarshaler = (*CheckVersionRequest)(nil)
var _ json.Marshaler = (*CheckVersionResponse)(nil)
var _ json.Unmarshaler = (*CheckVersionResponse)(nil)
var _ json.Marshaler = (*CheckVersionResponseData)(nil)
var _ json.Unmarshaler = (*CheckVersionResponseData)(nil)
var _ json.Marshaler = (*UpdateGalleryInfoRequest)(nil)
var _ json.Unmarshaler = (*UpdateGalleryInfoRequest)(nil)
var _ json.Marshaler = (*ReleasePutOnRequest)(nil)
var _ json.Unmarshaler = (*ReleasePutOnRequest)(nil)
var _ json.Marshaler = (*ReleasePutOnResponse)(nil)
var _ json.Unmarshaler = (*ReleasePutOnResponse)(nil)
var _ json.Marshaler = (*ReleasePutOffRequest)(nil)
var _ json.Unmarshaler = (*ReleasePutOffRequest)(nil)
var _ json.Marshaler = (*ReleasePutOffResponse)(nil)
var _ json.Unmarshaler = (*ReleasePutOffResponse)(nil)

// ReleaseList implement json.Marshaler.
func (m *ReleaseList) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseList implement json.Marshaler.
func (m *ReleaseList) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Mode implement json.Marshaler.
func (m *Mode) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Mode implement json.Marshaler.
func (m *Mode) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseCreateRequest implement json.Marshaler.
func (m *ReleaseCreateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseCreateRequest implement json.Marshaler.
func (m *ReleaseCreateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseResource implement json.Marshaler.
func (m *ReleaseResource) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseResource implement json.Marshaler.
func (m *ReleaseResource) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseCreateResponse implement json.Marshaler.
func (m *ReleaseCreateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseCreateResponse implement json.Marshaler.
func (m *ReleaseCreateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseCreateResponseData implement json.Marshaler.
func (m *ReleaseCreateResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseCreateResponseData implement json.Marshaler.
func (m *ReleaseCreateResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseUpdateRequest implement json.Marshaler.
func (m *ReleaseUpdateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseUpdateRequest implement json.Marshaler.
func (m *ReleaseUpdateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseDataResponse implement json.Marshaler.
func (m *ReleaseDataResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseDataResponse implement json.Marshaler.
func (m *ReleaseDataResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseUserDataResponse implement json.Marshaler.
func (m *ReleaseUserDataResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseUserDataResponse implement json.Marshaler.
func (m *ReleaseUserDataResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseUpdateResponse implement json.Marshaler.
func (m *ReleaseUpdateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseUpdateResponse implement json.Marshaler.
func (m *ReleaseUpdateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseReferenceUpdateRequest implement json.Marshaler.
func (m *ReleaseReferenceUpdateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseReferenceUpdateRequest implement json.Marshaler.
func (m *ReleaseReferenceUpdateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetIosPlistRequest implement json.Marshaler.
func (m *GetIosPlistRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetIosPlistRequest implement json.Marshaler.
func (m *GetIosPlistRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetIosPlistResponse implement json.Marshaler.
func (m *GetIosPlistResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetIosPlistResponse implement json.Marshaler.
func (m *GetIosPlistResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseDeleteRequest implement json.Marshaler.
func (m *ReleaseDeleteRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseDeleteRequest implement json.Marshaler.
func (m *ReleaseDeleteRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseDeleteResponse implement json.Marshaler.
func (m *ReleaseDeleteResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseDeleteResponse implement json.Marshaler.
func (m *ReleaseDeleteResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseGetRequest implement json.Marshaler.
func (m *ReleaseGetRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseGetRequest implement json.Marshaler.
func (m *ReleaseGetRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseGetResponse implement json.Marshaler.
func (m *ReleaseGetResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseGetResponse implement json.Marshaler.
func (m *ReleaseGetResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ModeSummary implement json.Marshaler.
func (m *ModeSummary) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ModeSummary implement json.Marshaler.
func (m *ModeSummary) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseSummaryArray implement json.Marshaler.
func (m *ReleaseSummaryArray) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseSummaryArray implement json.Marshaler.
func (m *ReleaseSummaryArray) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Tag implement json.Marshaler.
func (m *Tag) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Tag implement json.Marshaler.
func (m *Tag) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseGetResponseData implement json.Marshaler.
func (m *ReleaseGetResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseGetResponseData implement json.Marshaler.
func (m *ReleaseGetResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// AddonInfo implement json.Marshaler.
func (m *AddonInfo) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// AddonInfo implement json.Marshaler.
func (m *AddonInfo) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ServiceImagePair implement json.Marshaler.
func (m *ServiceImagePair) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ServiceImagePair implement json.Marshaler.
func (m *ServiceImagePair) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ApplicationReleaseSummary implement json.Marshaler.
func (m *ApplicationReleaseSummary) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ApplicationReleaseSummary implement json.Marshaler.
func (m *ApplicationReleaseSummary) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseListRequest implement json.Marshaler.
func (m *ReleaseListRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseListRequest implement json.Marshaler.
func (m *ReleaseListRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseListResponse implement json.Marshaler.
func (m *ReleaseListResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseListResponse implement json.Marshaler.
func (m *ReleaseListResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListReleaseNameRequest implement json.Marshaler.
func (m *ListReleaseNameRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListReleaseNameRequest implement json.Marshaler.
func (m *ListReleaseNameRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListReleaseNameResponse implement json.Marshaler.
func (m *ListReleaseNameResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListReleaseNameResponse implement json.Marshaler.
func (m *ListReleaseNameResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseListResponseData implement json.Marshaler.
func (m *ReleaseListResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseListResponseData implement json.Marshaler.
func (m *ReleaseListResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseData implement json.Marshaler.
func (m *ReleaseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseData implement json.Marshaler.
func (m *ReleaseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseNameListRequest implement json.Marshaler.
func (m *ReleaseNameListRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseNameListRequest implement json.Marshaler.
func (m *ReleaseNameListRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseNameListResponse implement json.Marshaler.
func (m *ReleaseNameListResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseNameListResponse implement json.Marshaler.
func (m *ReleaseNameListResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetLatestReleasesRequest implement json.Marshaler.
func (m *GetLatestReleasesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetLatestReleasesRequest implement json.Marshaler.
func (m *GetLatestReleasesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetLatestReleasesResponse implement json.Marshaler.
func (m *GetLatestReleasesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetLatestReleasesResponse implement json.Marshaler.
func (m *GetLatestReleasesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetLatestReleasesResponseData implement json.Marshaler.
func (m *GetLatestReleasesResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetLatestReleasesResponseData implement json.Marshaler.
func (m *GetLatestReleasesResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseGCRequest implement json.Marshaler.
func (m *ReleaseGCRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseGCRequest implement json.Marshaler.
func (m *ReleaseGCRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseUploadRequest implement json.Marshaler.
func (m *ReleaseUploadRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseUploadRequest implement json.Marshaler.
func (m *ReleaseUploadRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleaseUploadResponse implement json.Marshaler.
func (m *ReleaseUploadResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleaseUploadResponse implement json.Marshaler.
func (m *ReleaseUploadResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ParseReleaseFileRequest implement json.Marshaler.
func (m *ParseReleaseFileRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ParseReleaseFileRequest implement json.Marshaler.
func (m *ParseReleaseFileRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ParseReleaseFileResponse implement json.Marshaler.
func (m *ParseReleaseFileResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ParseReleaseFileResponse implement json.Marshaler.
func (m *ParseReleaseFileResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ParseReleaseFileResponseData implement json.Marshaler.
func (m *ParseReleaseFileResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ParseReleaseFileResponseData implement json.Marshaler.
func (m *ParseReleaseFileResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// FormalReleaseRequest implement json.Marshaler.
func (m *FormalReleaseRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// FormalReleaseRequest implement json.Marshaler.
func (m *FormalReleaseRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// FormalReleaseResponse implement json.Marshaler.
func (m *FormalReleaseResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// FormalReleaseResponse implement json.Marshaler.
func (m *FormalReleaseResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// FormalReleasesRequest implement json.Marshaler.
func (m *FormalReleasesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// FormalReleasesRequest implement json.Marshaler.
func (m *FormalReleasesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// FormalReleasesResponse implement json.Marshaler.
func (m *FormalReleasesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// FormalReleasesResponse implement json.Marshaler.
func (m *FormalReleasesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleasesDeleteRequest implement json.Marshaler.
func (m *ReleasesDeleteRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleasesDeleteRequest implement json.Marshaler.
func (m *ReleasesDeleteRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleasesDeleteResponse implement json.Marshaler.
func (m *ReleasesDeleteResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleasesDeleteResponse implement json.Marshaler.
func (m *ReleasesDeleteResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CheckVersionRequest implement json.Marshaler.
func (m *CheckVersionRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CheckVersionRequest implement json.Marshaler.
func (m *CheckVersionRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CheckVersionResponse implement json.Marshaler.
func (m *CheckVersionResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CheckVersionResponse implement json.Marshaler.
func (m *CheckVersionResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CheckVersionResponseData implement json.Marshaler.
func (m *CheckVersionResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CheckVersionResponseData implement json.Marshaler.
func (m *CheckVersionResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateGalleryInfoRequest implement json.Marshaler.
func (m *UpdateGalleryInfoRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateGalleryInfoRequest implement json.Marshaler.
func (m *UpdateGalleryInfoRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleasePutOnRequest implement json.Marshaler.
func (m *ReleasePutOnRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleasePutOnRequest implement json.Marshaler.
func (m *ReleasePutOnRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleasePutOnResponse implement json.Marshaler.
func (m *ReleasePutOnResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleasePutOnResponse implement json.Marshaler.
func (m *ReleasePutOnResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleasePutOffRequest implement json.Marshaler.
func (m *ReleasePutOffRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleasePutOffRequest implement json.Marshaler.
func (m *ReleasePutOffRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReleasePutOffResponse implement json.Marshaler.
func (m *ReleasePutOffResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReleasePutOffResponse implement json.Marshaler.
func (m *ReleasePutOffResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
