// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: apm_diagnotor.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*ListServicesRequest)(nil)
var _ json.Unmarshaler = (*ListServicesRequest)(nil)
var _ json.Marshaler = (*ListServicesResponse)(nil)
var _ json.Unmarshaler = (*ListServicesResponse)(nil)
var _ json.Marshaler = (*ServiceInfo)(nil)
var _ json.Unmarshaler = (*ServiceInfo)(nil)
var _ json.Marshaler = (*InstanceInfo)(nil)
var _ json.Unmarshaler = (*InstanceInfo)(nil)
var _ json.Marshaler = (*StartDiagnosisRequest)(nil)
var _ json.Unmarshaler = (*StartDiagnosisRequest)(nil)
var _ json.Marshaler = (*StartDiagnosisResponse)(nil)
var _ json.Unmarshaler = (*StartDiagnosisResponse)(nil)
var _ json.Marshaler = (*QueryDiagnosisStatusRequest)(nil)
var _ json.Unmarshaler = (*QueryDiagnosisStatusRequest)(nil)
var _ json.Marshaler = (*QueryDiagnosisStatusResponse)(nil)
var _ json.Unmarshaler = (*QueryDiagnosisStatusResponse)(nil)
var _ json.Marshaler = (*StopDiagnosisRequest)(nil)
var _ json.Unmarshaler = (*StopDiagnosisRequest)(nil)
var _ json.Marshaler = (*StopDiagnosisResponse)(nil)
var _ json.Unmarshaler = (*StopDiagnosisResponse)(nil)
var _ json.Marshaler = (*ListProcessesRequest)(nil)
var _ json.Unmarshaler = (*ListProcessesRequest)(nil)
var _ json.Marshaler = (*ListProcessesResponse)(nil)
var _ json.Unmarshaler = (*ListProcessesResponse)(nil)

// ListServicesRequest implement json.Marshaler.
func (m *ListServicesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListServicesRequest implement json.Marshaler.
func (m *ListServicesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListServicesResponse implement json.Marshaler.
func (m *ListServicesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListServicesResponse implement json.Marshaler.
func (m *ListServicesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ServiceInfo implement json.Marshaler.
func (m *ServiceInfo) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ServiceInfo implement json.Marshaler.
func (m *ServiceInfo) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// InstanceInfo implement json.Marshaler.
func (m *InstanceInfo) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// InstanceInfo implement json.Marshaler.
func (m *InstanceInfo) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

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
