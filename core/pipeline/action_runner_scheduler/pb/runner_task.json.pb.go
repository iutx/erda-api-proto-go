// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: runner_task.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*RunnerTask)(nil)
var _ json.Unmarshaler = (*RunnerTask)(nil)
var _ json.Marshaler = (*RunnerTaskQueryRequest)(nil)
var _ json.Unmarshaler = (*RunnerTaskQueryRequest)(nil)
var _ json.Marshaler = (*RunnerTaskQueryResponse)(nil)
var _ json.Unmarshaler = (*RunnerTaskQueryResponse)(nil)
var _ json.Marshaler = (*RunnerTaskFetchRequest)(nil)
var _ json.Unmarshaler = (*RunnerTaskFetchRequest)(nil)
var _ json.Marshaler = (*RunnerTaskFetchResponse)(nil)
var _ json.Unmarshaler = (*RunnerTaskFetchResponse)(nil)
var _ json.Marshaler = (*RunnerTaskCreateRequest)(nil)
var _ json.Unmarshaler = (*RunnerTaskCreateRequest)(nil)
var _ json.Marshaler = (*RunnerTaskCreateResponse)(nil)
var _ json.Unmarshaler = (*RunnerTaskCreateResponse)(nil)
var _ json.Marshaler = (*RunnerTaskUpdateRequest)(nil)
var _ json.Unmarshaler = (*RunnerTaskUpdateRequest)(nil)
var _ json.Marshaler = (*RunnerTaskUpdateResponse)(nil)
var _ json.Unmarshaler = (*RunnerTaskUpdateResponse)(nil)
var _ json.Marshaler = (*LogCollectRequest)(nil)
var _ json.Unmarshaler = (*LogCollectRequest)(nil)
var _ json.Marshaler = (*LogCollectResponse)(nil)
var _ json.Unmarshaler = (*LogCollectResponse)(nil)

// RunnerTask implement json.Marshaler.
func (m *RunnerTask) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RunnerTask implement json.Marshaler.
func (m *RunnerTask) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RunnerTaskQueryRequest implement json.Marshaler.
func (m *RunnerTaskQueryRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RunnerTaskQueryRequest implement json.Marshaler.
func (m *RunnerTaskQueryRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RunnerTaskQueryResponse implement json.Marshaler.
func (m *RunnerTaskQueryResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RunnerTaskQueryResponse implement json.Marshaler.
func (m *RunnerTaskQueryResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RunnerTaskFetchRequest implement json.Marshaler.
func (m *RunnerTaskFetchRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RunnerTaskFetchRequest implement json.Marshaler.
func (m *RunnerTaskFetchRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RunnerTaskFetchResponse implement json.Marshaler.
func (m *RunnerTaskFetchResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RunnerTaskFetchResponse implement json.Marshaler.
func (m *RunnerTaskFetchResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RunnerTaskCreateRequest implement json.Marshaler.
func (m *RunnerTaskCreateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RunnerTaskCreateRequest implement json.Marshaler.
func (m *RunnerTaskCreateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RunnerTaskCreateResponse implement json.Marshaler.
func (m *RunnerTaskCreateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RunnerTaskCreateResponse implement json.Marshaler.
func (m *RunnerTaskCreateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RunnerTaskUpdateRequest implement json.Marshaler.
func (m *RunnerTaskUpdateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RunnerTaskUpdateRequest implement json.Marshaler.
func (m *RunnerTaskUpdateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RunnerTaskUpdateResponse implement json.Marshaler.
func (m *RunnerTaskUpdateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RunnerTaskUpdateResponse implement json.Marshaler.
func (m *RunnerTaskUpdateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// LogCollectRequest implement json.Marshaler.
func (m *LogCollectRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// LogCollectRequest implement json.Marshaler.
func (m *LogCollectRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// LogCollectResponse implement json.Marshaler.
func (m *LogCollectResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// LogCollectResponse implement json.Marshaler.
func (m *LogCollectResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
