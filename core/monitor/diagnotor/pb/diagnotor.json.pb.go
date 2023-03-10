// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: diagnotor.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*StartDiagnosisRequest)(nil)
var _ json.Unmarshaler = (*StartDiagnosisRequest)(nil)
var _ json.Marshaler = (*StartDiagnosisResponse)(nil)
var _ json.Unmarshaler = (*StartDiagnosisResponse)(nil)
var _ json.Marshaler = (*QueryDiagnosisStatusRequest)(nil)
var _ json.Unmarshaler = (*QueryDiagnosisStatusRequest)(nil)
var _ json.Marshaler = (*QueryDiagnosisStatusResponse)(nil)
var _ json.Unmarshaler = (*QueryDiagnosisStatusResponse)(nil)
var _ json.Marshaler = (*DiagnosisInstance)(nil)
var _ json.Unmarshaler = (*DiagnosisInstance)(nil)
var _ json.Marshaler = (*StopDiagnosisRequest)(nil)
var _ json.Unmarshaler = (*StopDiagnosisRequest)(nil)
var _ json.Marshaler = (*StopDiagnosisResponse)(nil)
var _ json.Unmarshaler = (*StopDiagnosisResponse)(nil)
var _ json.Marshaler = (*ListProcessesRequest)(nil)
var _ json.Unmarshaler = (*ListProcessesRequest)(nil)
var _ json.Marshaler = (*ListProcessesResponse)(nil)
var _ json.Unmarshaler = (*ListProcessesResponse)(nil)
var _ json.Marshaler = (*HostProcessStatus)(nil)
var _ json.Unmarshaler = (*HostProcessStatus)(nil)
var _ json.Marshaler = (*Process)(nil)
var _ json.Unmarshaler = (*Process)(nil)
var _ json.Marshaler = (*ProcessRLimit)(nil)
var _ json.Unmarshaler = (*ProcessRLimit)(nil)
var _ json.Marshaler = (*ProcessRLimitStatus)(nil)
var _ json.Unmarshaler = (*ProcessRLimitStatus)(nil)
var _ json.Marshaler = (*ProcessMemoryStatus)(nil)
var _ json.Unmarshaler = (*ProcessMemoryStatus)(nil)
var _ json.Marshaler = (*ProcessCPUStatus)(nil)
var _ json.Unmarshaler = (*ProcessCPUStatus)(nil)
var _ json.Marshaler = (*ProcessIOStatus)(nil)
var _ json.Unmarshaler = (*ProcessIOStatus)(nil)
var _ json.Marshaler = (*ProcessContextSwitches)(nil)
var _ json.Unmarshaler = (*ProcessContextSwitches)(nil)

// StartDiagnosisRequest implement json.Marshaler.
func (m *StartDiagnosisRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// StartDiagnosisRequest implement json.Marshaler.
func (m *StartDiagnosisRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// StartDiagnosisResponse implement json.Marshaler.
func (m *StartDiagnosisResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// StartDiagnosisResponse implement json.Marshaler.
func (m *StartDiagnosisResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryDiagnosisStatusRequest implement json.Marshaler.
func (m *QueryDiagnosisStatusRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryDiagnosisStatusRequest implement json.Marshaler.
func (m *QueryDiagnosisStatusRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryDiagnosisStatusResponse implement json.Marshaler.
func (m *QueryDiagnosisStatusResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryDiagnosisStatusResponse implement json.Marshaler.
func (m *QueryDiagnosisStatusResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DiagnosisInstance implement json.Marshaler.
func (m *DiagnosisInstance) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DiagnosisInstance implement json.Marshaler.
func (m *DiagnosisInstance) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// StopDiagnosisRequest implement json.Marshaler.
func (m *StopDiagnosisRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// StopDiagnosisRequest implement json.Marshaler.
func (m *StopDiagnosisRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// StopDiagnosisResponse implement json.Marshaler.
func (m *StopDiagnosisResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// StopDiagnosisResponse implement json.Marshaler.
func (m *StopDiagnosisResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListProcessesRequest implement json.Marshaler.
func (m *ListProcessesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListProcessesRequest implement json.Marshaler.
func (m *ListProcessesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListProcessesResponse implement json.Marshaler.
func (m *ListProcessesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListProcessesResponse implement json.Marshaler.
func (m *ListProcessesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// HostProcessStatus implement json.Marshaler.
func (m *HostProcessStatus) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// HostProcessStatus implement json.Marshaler.
func (m *HostProcessStatus) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Process implement json.Marshaler.
func (m *Process) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Process implement json.Marshaler.
func (m *Process) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProcessRLimit implement json.Marshaler.
func (m *ProcessRLimit) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProcessRLimit implement json.Marshaler.
func (m *ProcessRLimit) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProcessRLimitStatus implement json.Marshaler.
func (m *ProcessRLimitStatus) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProcessRLimitStatus implement json.Marshaler.
func (m *ProcessRLimitStatus) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProcessMemoryStatus implement json.Marshaler.
func (m *ProcessMemoryStatus) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProcessMemoryStatus implement json.Marshaler.
func (m *ProcessMemoryStatus) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProcessCPUStatus implement json.Marshaler.
func (m *ProcessCPUStatus) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProcessCPUStatus implement json.Marshaler.
func (m *ProcessCPUStatus) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProcessIOStatus implement json.Marshaler.
func (m *ProcessIOStatus) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProcessIOStatus implement json.Marshaler.
func (m *ProcessIOStatus) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ProcessContextSwitches implement json.Marshaler.
func (m *ProcessContextSwitches) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ProcessContextSwitches implement json.Marshaler.
func (m *ProcessContextSwitches) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
