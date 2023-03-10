// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: msp_apm_exception.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*EXCEPTION_EVENT_ID_LIST_Request)(nil)
var _ json.Unmarshaler = (*EXCEPTION_EVENT_ID_LIST_Request)(nil)
var _ json.Marshaler = (*EXCEPTION_EVENT_Request)(nil)
var _ json.Unmarshaler = (*EXCEPTION_EVENT_Request)(nil)
var _ json.Marshaler = (*EXCEPTION_LIST_Request)(nil)
var _ json.Unmarshaler = (*EXCEPTION_LIST_Request)(nil)

// EXCEPTION_EVENT_ID_LIST_Request implement json.Marshaler.
func (m *EXCEPTION_EVENT_ID_LIST_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// EXCEPTION_EVENT_ID_LIST_Request implement json.Marshaler.
func (m *EXCEPTION_EVENT_ID_LIST_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// EXCEPTION_EVENT_Request implement json.Marshaler.
func (m *EXCEPTION_EVENT_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// EXCEPTION_EVENT_Request implement json.Marshaler.
func (m *EXCEPTION_EVENT_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// EXCEPTION_LIST_Request implement json.Marshaler.
func (m *EXCEPTION_LIST_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// EXCEPTION_LIST_Request implement json.Marshaler.
func (m *EXCEPTION_LIST_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
