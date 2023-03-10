// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: cicdcms.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*CICDCmsUpdateRequest)(nil)
var _ json.Unmarshaler = (*CICDCmsUpdateRequest)(nil)
var _ json.Marshaler = (*CICDCmsUpdateResponse)(nil)
var _ json.Unmarshaler = (*CICDCmsUpdateResponse)(nil)
var _ json.Marshaler = (*CICDCmsCreateRequest)(nil)
var _ json.Unmarshaler = (*CICDCmsCreateRequest)(nil)
var _ json.Marshaler = (*CICDCmsCreateResponse)(nil)
var _ json.Unmarshaler = (*CICDCmsCreateResponse)(nil)
var _ json.Marshaler = (*CICDCmsDeleteRequest)(nil)
var _ json.Unmarshaler = (*CICDCmsDeleteRequest)(nil)
var _ json.Marshaler = (*CICDCmsDeleteResponse)(nil)
var _ json.Unmarshaler = (*CICDCmsDeleteResponse)(nil)
var _ json.Marshaler = (*Config)(nil)
var _ json.Unmarshaler = (*Config)(nil)

// CICDCmsUpdateRequest implement json.Marshaler.
func (m *CICDCmsUpdateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CICDCmsUpdateRequest implement json.Marshaler.
func (m *CICDCmsUpdateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CICDCmsUpdateResponse implement json.Marshaler.
func (m *CICDCmsUpdateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CICDCmsUpdateResponse implement json.Marshaler.
func (m *CICDCmsUpdateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CICDCmsCreateRequest implement json.Marshaler.
func (m *CICDCmsCreateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CICDCmsCreateRequest implement json.Marshaler.
func (m *CICDCmsCreateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CICDCmsCreateResponse implement json.Marshaler.
func (m *CICDCmsCreateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CICDCmsCreateResponse implement json.Marshaler.
func (m *CICDCmsCreateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CICDCmsDeleteRequest implement json.Marshaler.
func (m *CICDCmsDeleteRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CICDCmsDeleteRequest implement json.Marshaler.
func (m *CICDCmsDeleteRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CICDCmsDeleteResponse implement json.Marshaler.
func (m *CICDCmsDeleteResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CICDCmsDeleteResponse implement json.Marshaler.
func (m *CICDCmsDeleteResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Config implement json.Marshaler.
func (m *Config) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Config implement json.Marshaler.
func (m *Config) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
