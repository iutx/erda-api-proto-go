// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: hepa.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*DomainListRequest)(nil)
var _ json.Unmarshaler = (*DomainListRequest)(nil)
var _ json.Marshaler = (*DomainListResponse)(nil)
var _ json.Unmarshaler = (*DomainListResponse)(nil)
var _ json.Marshaler = (*HEPA_DOMAINS_GET_Request)(nil)
var _ json.Unmarshaler = (*HEPA_DOMAINS_GET_Request)(nil)
var _ json.Marshaler = (*HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request)(nil)
var _ json.Unmarshaler = (*HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request)(nil)

// DomainListRequest implement json.Marshaler.
func (m *DomainListRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DomainListRequest implement json.Marshaler.
func (m *DomainListRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DomainListResponse implement json.Marshaler.
func (m *DomainListResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DomainListResponse implement json.Marshaler.
func (m *DomainListResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// HEPA_DOMAINS_GET_Request implement json.Marshaler.
func (m *HEPA_DOMAINS_GET_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// HEPA_DOMAINS_GET_Request implement json.Marshaler.
func (m *HEPA_DOMAINS_GET_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request implement json.Marshaler.
func (m *HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request implement json.Marshaler.
func (m *HEPA_RUNTIME_SERVICE_DOMAIN_UPDATE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
