// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: admin.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ADMIN_APPROVE_GET_Request)(nil)
var _ json.Unmarshaler = (*ADMIN_APPROVE_GET_Request)(nil)
var _ json.Marshaler = (*ADMIN_CLUSTER_GET_Request)(nil)
var _ json.Unmarshaler = (*ADMIN_CLUSTER_GET_Request)(nil)
var _ json.Marshaler = (*ADMIN_NOTICE_DELETE_Request)(nil)
var _ json.Unmarshaler = (*ADMIN_NOTICE_DELETE_Request)(nil)
var _ json.Marshaler = (*ADMIN_NOTICE_PUBLISH_Request)(nil)
var _ json.Unmarshaler = (*ADMIN_NOTICE_PUBLISH_Request)(nil)
var _ json.Marshaler = (*ADMIN_NOTICE_UNPUBLISH_Request)(nil)
var _ json.Unmarshaler = (*ADMIN_NOTICE_UNPUBLISH_Request)(nil)
var _ json.Marshaler = (*ApproveDTO)(nil)
var _ json.Unmarshaler = (*ApproveDTO)(nil)
var _ json.Marshaler = (*ApproveDetailResponse)(nil)
var _ json.Unmarshaler = (*ApproveDetailResponse)(nil)
var _ json.Marshaler = (*ApproveListRequest)(nil)
var _ json.Unmarshaler = (*ApproveListRequest)(nil)
var _ json.Marshaler = (*ApproveUpdateRequest)(nil)
var _ json.Unmarshaler = (*ApproveUpdateRequest)(nil)
var _ json.Marshaler = (*ApproveUpdateResponse)(nil)
var _ json.Unmarshaler = (*ApproveUpdateResponse)(nil)
var _ json.Marshaler = (*AuditsListRequest)(nil)
var _ json.Unmarshaler = (*AuditsListRequest)(nil)
var _ json.Marshaler = (*AuditsListResponse)(nil)
var _ json.Unmarshaler = (*AuditsListResponse)(nil)
var _ json.Marshaler = (*ClusterInfo)(nil)
var _ json.Unmarshaler = (*ClusterInfo)(nil)
var _ json.Marshaler = (*ClusterListRequest)(nil)
var _ json.Unmarshaler = (*ClusterListRequest)(nil)
var _ json.Marshaler = (*ClusterListResponse)(nil)
var _ json.Unmarshaler = (*ClusterListResponse)(nil)
var _ json.Marshaler = (*Notice)(nil)
var _ json.Unmarshaler = (*Notice)(nil)
var _ json.Marshaler = (*NoticeCreateRequest)(nil)
var _ json.Unmarshaler = (*NoticeCreateRequest)(nil)
var _ json.Marshaler = (*NoticeCreateResponse)(nil)
var _ json.Unmarshaler = (*NoticeCreateResponse)(nil)
var _ json.Marshaler = (*NoticeDeleteResponse)(nil)
var _ json.Unmarshaler = (*NoticeDeleteResponse)(nil)
var _ json.Marshaler = (*NoticeListRequest)(nil)
var _ json.Unmarshaler = (*NoticeListRequest)(nil)
var _ json.Marshaler = (*NoticeListResponse)(nil)
var _ json.Unmarshaler = (*NoticeListResponse)(nil)
var _ json.Marshaler = (*NoticeListResponseData)(nil)
var _ json.Unmarshaler = (*NoticeListResponseData)(nil)
var _ json.Marshaler = (*NoticePublishResponse)(nil)
var _ json.Unmarshaler = (*NoticePublishResponse)(nil)
var _ json.Marshaler = (*NoticeUpdateRequest)(nil)
var _ json.Unmarshaler = (*NoticeUpdateRequest)(nil)
var _ json.Marshaler = (*NoticeUpdateResponse)(nil)
var _ json.Unmarshaler = (*NoticeUpdateResponse)(nil)
var _ json.Marshaler = (*PagingApproveDTO)(nil)
var _ json.Unmarshaler = (*PagingApproveDTO)(nil)
var _ json.Marshaler = (*UserInfo)(nil)
var _ json.Unmarshaler = (*UserInfo)(nil)
var _ json.Marshaler = (*UserListRequest)(nil)
var _ json.Unmarshaler = (*UserListRequest)(nil)
var _ json.Marshaler = (*UserListResponse)(nil)
var _ json.Unmarshaler = (*UserListResponse)(nil)
var _ json.Marshaler = (*UserListResponseData)(nil)
var _ json.Unmarshaler = (*UserListResponseData)(nil)

// ADMIN_APPROVE_GET_Request implement json.Marshaler.
func (m *ADMIN_APPROVE_GET_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ADMIN_APPROVE_GET_Request implement json.Marshaler.
func (m *ADMIN_APPROVE_GET_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ADMIN_CLUSTER_GET_Request implement json.Marshaler.
func (m *ADMIN_CLUSTER_GET_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ADMIN_CLUSTER_GET_Request implement json.Marshaler.
func (m *ADMIN_CLUSTER_GET_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ADMIN_NOTICE_DELETE_Request implement json.Marshaler.
func (m *ADMIN_NOTICE_DELETE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ADMIN_NOTICE_DELETE_Request implement json.Marshaler.
func (m *ADMIN_NOTICE_DELETE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ADMIN_NOTICE_PUBLISH_Request implement json.Marshaler.
func (m *ADMIN_NOTICE_PUBLISH_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ADMIN_NOTICE_PUBLISH_Request implement json.Marshaler.
func (m *ADMIN_NOTICE_PUBLISH_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ADMIN_NOTICE_UNPUBLISH_Request implement json.Marshaler.
func (m *ADMIN_NOTICE_UNPUBLISH_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ADMIN_NOTICE_UNPUBLISH_Request implement json.Marshaler.
func (m *ADMIN_NOTICE_UNPUBLISH_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ApproveDTO implement json.Marshaler.
func (m *ApproveDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ApproveDTO implement json.Marshaler.
func (m *ApproveDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ApproveDetailResponse implement json.Marshaler.
func (m *ApproveDetailResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ApproveDetailResponse implement json.Marshaler.
func (m *ApproveDetailResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ApproveListRequest implement json.Marshaler.
func (m *ApproveListRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ApproveListRequest implement json.Marshaler.
func (m *ApproveListRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ApproveUpdateRequest implement json.Marshaler.
func (m *ApproveUpdateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ApproveUpdateRequest implement json.Marshaler.
func (m *ApproveUpdateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ApproveUpdateResponse implement json.Marshaler.
func (m *ApproveUpdateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ApproveUpdateResponse implement json.Marshaler.
func (m *ApproveUpdateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// AuditsListRequest implement json.Marshaler.
func (m *AuditsListRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// AuditsListRequest implement json.Marshaler.
func (m *AuditsListRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// AuditsListResponse implement json.Marshaler.
func (m *AuditsListResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// AuditsListResponse implement json.Marshaler.
func (m *AuditsListResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ClusterInfo implement json.Marshaler.
func (m *ClusterInfo) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ClusterInfo implement json.Marshaler.
func (m *ClusterInfo) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ClusterListRequest implement json.Marshaler.
func (m *ClusterListRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ClusterListRequest implement json.Marshaler.
func (m *ClusterListRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ClusterListResponse implement json.Marshaler.
func (m *ClusterListResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ClusterListResponse implement json.Marshaler.
func (m *ClusterListResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Notice implement json.Marshaler.
func (m *Notice) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Notice implement json.Marshaler.
func (m *Notice) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NoticeCreateRequest implement json.Marshaler.
func (m *NoticeCreateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NoticeCreateRequest implement json.Marshaler.
func (m *NoticeCreateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NoticeCreateResponse implement json.Marshaler.
func (m *NoticeCreateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NoticeCreateResponse implement json.Marshaler.
func (m *NoticeCreateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NoticeDeleteResponse implement json.Marshaler.
func (m *NoticeDeleteResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NoticeDeleteResponse implement json.Marshaler.
func (m *NoticeDeleteResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NoticeListRequest implement json.Marshaler.
func (m *NoticeListRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NoticeListRequest implement json.Marshaler.
func (m *NoticeListRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NoticeListResponse implement json.Marshaler.
func (m *NoticeListResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NoticeListResponse implement json.Marshaler.
func (m *NoticeListResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NoticeListResponseData implement json.Marshaler.
func (m *NoticeListResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NoticeListResponseData implement json.Marshaler.
func (m *NoticeListResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NoticePublishResponse implement json.Marshaler.
func (m *NoticePublishResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NoticePublishResponse implement json.Marshaler.
func (m *NoticePublishResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NoticeUpdateRequest implement json.Marshaler.
func (m *NoticeUpdateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NoticeUpdateRequest implement json.Marshaler.
func (m *NoticeUpdateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NoticeUpdateResponse implement json.Marshaler.
func (m *NoticeUpdateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NoticeUpdateResponse implement json.Marshaler.
func (m *NoticeUpdateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// PagingApproveDTO implement json.Marshaler.
func (m *PagingApproveDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// PagingApproveDTO implement json.Marshaler.
func (m *PagingApproveDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UserInfo implement json.Marshaler.
func (m *UserInfo) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UserInfo implement json.Marshaler.
func (m *UserInfo) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UserListRequest implement json.Marshaler.
func (m *UserListRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UserListRequest implement json.Marshaler.
func (m *UserListRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UserListResponse implement json.Marshaler.
func (m *UserListResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UserListResponse implement json.Marshaler.
func (m *UserListResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UserListResponseData implement json.Marshaler.
func (m *UserListResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UserListResponseData implement json.Marshaler.
func (m *UserListResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
