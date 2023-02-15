// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: report.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ListTasksRequest)(nil)
var _ json.Unmarshaler = (*ListTasksRequest)(nil)
var _ json.Marshaler = (*ListTasksResponse)(nil)
var _ json.Unmarshaler = (*ListTasksResponse)(nil)
var _ json.Marshaler = (*CreateTaskRequest)(nil)
var _ json.Unmarshaler = (*CreateTaskRequest)(nil)
var _ json.Marshaler = (*CreateTaskResponse)(nil)
var _ json.Unmarshaler = (*CreateTaskResponse)(nil)
var _ json.Marshaler = (*UpdateTaskRequest)(nil)
var _ json.Unmarshaler = (*UpdateTaskRequest)(nil)
var _ json.Marshaler = (*SwitchTaskRequest)(nil)
var _ json.Unmarshaler = (*SwitchTaskRequest)(nil)
var _ json.Marshaler = (*GetTaskRequest)(nil)
var _ json.Unmarshaler = (*GetTaskRequest)(nil)
var _ json.Marshaler = (*DeleteTaskRequest)(nil)
var _ json.Unmarshaler = (*DeleteTaskRequest)(nil)
var _ json.Marshaler = (*ListTypesRequest)(nil)
var _ json.Unmarshaler = (*ListTypesRequest)(nil)
var _ json.Marshaler = (*ListTypesResponse)(nil)
var _ json.Unmarshaler = (*ListTypesResponse)(nil)
var _ json.Marshaler = (*Type)(nil)
var _ json.Unmarshaler = (*Type)(nil)
var _ json.Marshaler = (*ListHistoriesRequest)(nil)
var _ json.Unmarshaler = (*ListHistoriesRequest)(nil)
var _ json.Marshaler = (*ListHistoriesResponse)(nil)
var _ json.Unmarshaler = (*ListHistoriesResponse)(nil)
var _ json.Marshaler = (*ReportHistoryDTO)(nil)
var _ json.Unmarshaler = (*ReportHistoryDTO)(nil)
var _ json.Marshaler = (*CreateHistoryRequest)(nil)
var _ json.Unmarshaler = (*CreateHistoryRequest)(nil)
var _ json.Marshaler = (*CreateHistoryResponse)(nil)
var _ json.Unmarshaler = (*CreateHistoryResponse)(nil)
var _ json.Marshaler = (*GetHistoryRequest)(nil)
var _ json.Unmarshaler = (*GetHistoryRequest)(nil)
var _ json.Marshaler = (*GetHistoryResponse)(nil)
var _ json.Unmarshaler = (*GetHistoryResponse)(nil)
var _ json.Marshaler = (*DeleteHistoryRequest)(nil)
var _ json.Unmarshaler = (*DeleteHistoryRequest)(nil)
var _ json.Marshaler = (*ReportTaskDTO)(nil)
var _ json.Unmarshaler = (*ReportTaskDTO)(nil)
var _ json.Marshaler = (*ReportTaskOnly)(nil)
var _ json.Unmarshaler = (*ReportTaskOnly)(nil)
var _ json.Marshaler = (*Notify)(nil)
var _ json.Unmarshaler = (*Notify)(nil)

// ListTasksRequest implement json.Marshaler.
func (m *ListTasksRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListTasksRequest implement json.Marshaler.
func (m *ListTasksRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListTasksResponse implement json.Marshaler.
func (m *ListTasksResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListTasksResponse implement json.Marshaler.
func (m *ListTasksResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateTaskRequest implement json.Marshaler.
func (m *CreateTaskRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateTaskRequest implement json.Marshaler.
func (m *CreateTaskRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateTaskResponse implement json.Marshaler.
func (m *CreateTaskResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateTaskResponse implement json.Marshaler.
func (m *CreateTaskResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateTaskRequest implement json.Marshaler.
func (m *UpdateTaskRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateTaskRequest implement json.Marshaler.
func (m *UpdateTaskRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SwitchTaskRequest implement json.Marshaler.
func (m *SwitchTaskRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SwitchTaskRequest implement json.Marshaler.
func (m *SwitchTaskRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetTaskRequest implement json.Marshaler.
func (m *GetTaskRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetTaskRequest implement json.Marshaler.
func (m *GetTaskRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteTaskRequest implement json.Marshaler.
func (m *DeleteTaskRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteTaskRequest implement json.Marshaler.
func (m *DeleteTaskRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListTypesRequest implement json.Marshaler.
func (m *ListTypesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListTypesRequest implement json.Marshaler.
func (m *ListTypesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListTypesResponse implement json.Marshaler.
func (m *ListTypesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListTypesResponse implement json.Marshaler.
func (m *ListTypesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Type implement json.Marshaler.
func (m *Type) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Type implement json.Marshaler.
func (m *Type) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListHistoriesRequest implement json.Marshaler.
func (m *ListHistoriesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListHistoriesRequest implement json.Marshaler.
func (m *ListHistoriesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListHistoriesResponse implement json.Marshaler.
func (m *ListHistoriesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListHistoriesResponse implement json.Marshaler.
func (m *ListHistoriesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReportHistoryDTO implement json.Marshaler.
func (m *ReportHistoryDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReportHistoryDTO implement json.Marshaler.
func (m *ReportHistoryDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateHistoryRequest implement json.Marshaler.
func (m *CreateHistoryRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateHistoryRequest implement json.Marshaler.
func (m *CreateHistoryRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateHistoryResponse implement json.Marshaler.
func (m *CreateHistoryResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateHistoryResponse implement json.Marshaler.
func (m *CreateHistoryResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetHistoryRequest implement json.Marshaler.
func (m *GetHistoryRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetHistoryRequest implement json.Marshaler.
func (m *GetHistoryRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetHistoryResponse implement json.Marshaler.
func (m *GetHistoryResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetHistoryResponse implement json.Marshaler.
func (m *GetHistoryResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteHistoryRequest implement json.Marshaler.
func (m *DeleteHistoryRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteHistoryRequest implement json.Marshaler.
func (m *DeleteHistoryRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReportTaskDTO implement json.Marshaler.
func (m *ReportTaskDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReportTaskDTO implement json.Marshaler.
func (m *ReportTaskDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ReportTaskOnly implement json.Marshaler.
func (m *ReportTaskOnly) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ReportTaskOnly implement json.Marshaler.
func (m *ReportTaskOnly) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Notify implement json.Marshaler.
func (m *Notify) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Notify implement json.Marshaler.
func (m *Notify) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
