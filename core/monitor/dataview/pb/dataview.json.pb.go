// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: dataview.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*CustomDashboardHistory)(nil)
var _ json.Unmarshaler = (*CustomDashboardHistory)(nil)
var _ json.Marshaler = (*ListCustomDashboardHistoryRequest)(nil)
var _ json.Unmarshaler = (*ListCustomDashboardHistoryRequest)(nil)
var _ json.Marshaler = (*ListCustomDashboardHistoryResponse)(nil)
var _ json.Unmarshaler = (*ListCustomDashboardHistoryResponse)(nil)
var _ json.Marshaler = (*ExportCustomViewRequest)(nil)
var _ json.Unmarshaler = (*ExportCustomViewRequest)(nil)
var _ json.Marshaler = (*ListSystemViewsRequest)(nil)
var _ json.Unmarshaler = (*ListSystemViewsRequest)(nil)
var _ json.Marshaler = (*ListSystemViewsResponse)(nil)
var _ json.Unmarshaler = (*ListSystemViewsResponse)(nil)
var _ json.Marshaler = (*GetSystemViewRequest)(nil)
var _ json.Unmarshaler = (*GetSystemViewRequest)(nil)
var _ json.Marshaler = (*GetSystemViewResponse)(nil)
var _ json.Unmarshaler = (*GetSystemViewResponse)(nil)
var _ json.Marshaler = (*ListCustomViewsRequest)(nil)
var _ json.Unmarshaler = (*ListCustomViewsRequest)(nil)
var _ json.Marshaler = (*GetCustomViewsCreatorRequest)(nil)
var _ json.Unmarshaler = (*GetCustomViewsCreatorRequest)(nil)
var _ json.Marshaler = (*ListCustomViewsResponse)(nil)
var _ json.Unmarshaler = (*ListCustomViewsResponse)(nil)
var _ json.Marshaler = (*Creator)(nil)
var _ json.Unmarshaler = (*Creator)(nil)
var _ json.Marshaler = (*GetCustomViewsCreatorResponse)(nil)
var _ json.Unmarshaler = (*GetCustomViewsCreatorResponse)(nil)
var _ json.Marshaler = (*GetCustomViewRequest)(nil)
var _ json.Unmarshaler = (*GetCustomViewRequest)(nil)
var _ json.Marshaler = (*GetCustomViewResponse)(nil)
var _ json.Unmarshaler = (*GetCustomViewResponse)(nil)
var _ json.Marshaler = (*CreateCustomViewRequest)(nil)
var _ json.Unmarshaler = (*CreateCustomViewRequest)(nil)
var _ json.Marshaler = (*CreateCustomViewResponse)(nil)
var _ json.Unmarshaler = (*CreateCustomViewResponse)(nil)
var _ json.Marshaler = (*UpdateCustomViewRequest)(nil)
var _ json.Unmarshaler = (*UpdateCustomViewRequest)(nil)
var _ json.Marshaler = (*UpdateCustomViewResponse)(nil)
var _ json.Unmarshaler = (*UpdateCustomViewResponse)(nil)
var _ json.Marshaler = (*DeleteCustomViewRequest)(nil)
var _ json.Unmarshaler = (*DeleteCustomViewRequest)(nil)
var _ json.Marshaler = (*DeleteCustomViewResponse)(nil)
var _ json.Unmarshaler = (*DeleteCustomViewResponse)(nil)
var _ json.Marshaler = (*ViewList)(nil)
var _ json.Unmarshaler = (*ViewList)(nil)
var _ json.Marshaler = (*View)(nil)
var _ json.Unmarshaler = (*View)(nil)
var _ json.Marshaler = (*Block)(nil)
var _ json.Unmarshaler = (*Block)(nil)
var _ json.Marshaler = (*Chart)(nil)
var _ json.Unmarshaler = (*Chart)(nil)
var _ json.Marshaler = (*DataItem)(nil)
var _ json.Unmarshaler = (*DataItem)(nil)

// CustomDashboardHistory implement json.Marshaler.
func (m *CustomDashboardHistory) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CustomDashboardHistory implement json.Marshaler.
func (m *CustomDashboardHistory) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListCustomDashboardHistoryRequest implement json.Marshaler.
func (m *ListCustomDashboardHistoryRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListCustomDashboardHistoryRequest implement json.Marshaler.
func (m *ListCustomDashboardHistoryRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListCustomDashboardHistoryResponse implement json.Marshaler.
func (m *ListCustomDashboardHistoryResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListCustomDashboardHistoryResponse implement json.Marshaler.
func (m *ListCustomDashboardHistoryResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ExportCustomViewRequest implement json.Marshaler.
func (m *ExportCustomViewRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ExportCustomViewRequest implement json.Marshaler.
func (m *ExportCustomViewRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListSystemViewsRequest implement json.Marshaler.
func (m *ListSystemViewsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListSystemViewsRequest implement json.Marshaler.
func (m *ListSystemViewsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListSystemViewsResponse implement json.Marshaler.
func (m *ListSystemViewsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListSystemViewsResponse implement json.Marshaler.
func (m *ListSystemViewsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetSystemViewRequest implement json.Marshaler.
func (m *GetSystemViewRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetSystemViewRequest implement json.Marshaler.
func (m *GetSystemViewRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetSystemViewResponse implement json.Marshaler.
func (m *GetSystemViewResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetSystemViewResponse implement json.Marshaler.
func (m *GetSystemViewResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListCustomViewsRequest implement json.Marshaler.
func (m *ListCustomViewsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListCustomViewsRequest implement json.Marshaler.
func (m *ListCustomViewsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCustomViewsCreatorRequest implement json.Marshaler.
func (m *GetCustomViewsCreatorRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCustomViewsCreatorRequest implement json.Marshaler.
func (m *GetCustomViewsCreatorRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListCustomViewsResponse implement json.Marshaler.
func (m *ListCustomViewsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListCustomViewsResponse implement json.Marshaler.
func (m *ListCustomViewsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Creator implement json.Marshaler.
func (m *Creator) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Creator implement json.Marshaler.
func (m *Creator) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCustomViewsCreatorResponse implement json.Marshaler.
func (m *GetCustomViewsCreatorResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCustomViewsCreatorResponse implement json.Marshaler.
func (m *GetCustomViewsCreatorResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCustomViewRequest implement json.Marshaler.
func (m *GetCustomViewRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCustomViewRequest implement json.Marshaler.
func (m *GetCustomViewRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCustomViewResponse implement json.Marshaler.
func (m *GetCustomViewResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCustomViewResponse implement json.Marshaler.
func (m *GetCustomViewResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateCustomViewRequest implement json.Marshaler.
func (m *CreateCustomViewRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateCustomViewRequest implement json.Marshaler.
func (m *CreateCustomViewRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateCustomViewResponse implement json.Marshaler.
func (m *CreateCustomViewResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateCustomViewResponse implement json.Marshaler.
func (m *CreateCustomViewResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateCustomViewRequest implement json.Marshaler.
func (m *UpdateCustomViewRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateCustomViewRequest implement json.Marshaler.
func (m *UpdateCustomViewRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateCustomViewResponse implement json.Marshaler.
func (m *UpdateCustomViewResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateCustomViewResponse implement json.Marshaler.
func (m *UpdateCustomViewResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteCustomViewRequest implement json.Marshaler.
func (m *DeleteCustomViewRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteCustomViewRequest implement json.Marshaler.
func (m *DeleteCustomViewRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteCustomViewResponse implement json.Marshaler.
func (m *DeleteCustomViewResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteCustomViewResponse implement json.Marshaler.
func (m *DeleteCustomViewResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ViewList implement json.Marshaler.
func (m *ViewList) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ViewList implement json.Marshaler.
func (m *ViewList) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// View implement json.Marshaler.
func (m *View) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// View implement json.Marshaler.
func (m *View) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Block implement json.Marshaler.
func (m *Block) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Block implement json.Marshaler.
func (m *Block) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Chart implement json.Marshaler.
func (m *Chart) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Chart implement json.Marshaler.
func (m *Chart) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DataItem implement json.Marshaler.
func (m *DataItem) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DataItem implement json.Marshaler.
func (m *DataItem) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
