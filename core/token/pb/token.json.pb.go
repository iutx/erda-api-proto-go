// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: token.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*GetTokenRequest)(nil)
var _ json.Unmarshaler = (*GetTokenRequest)(nil)
var _ json.Marshaler = (*GetTokenResponse)(nil)
var _ json.Unmarshaler = (*GetTokenResponse)(nil)
var _ json.Marshaler = (*Token)(nil)
var _ json.Unmarshaler = (*Token)(nil)
var _ json.Marshaler = (*QueryTokensRequest)(nil)
var _ json.Unmarshaler = (*QueryTokensRequest)(nil)
var _ json.Marshaler = (*QueryTokensResponse)(nil)
var _ json.Unmarshaler = (*QueryTokensResponse)(nil)
var _ json.Marshaler = (*CreateTokenRequest)(nil)
var _ json.Unmarshaler = (*CreateTokenRequest)(nil)
var _ json.Marshaler = (*CreateTokenResponse)(nil)
var _ json.Unmarshaler = (*CreateTokenResponse)(nil)
var _ json.Marshaler = (*UpdateTokenRequest)(nil)
var _ json.Unmarshaler = (*UpdateTokenRequest)(nil)
var _ json.Marshaler = (*UpdateTokenResponse)(nil)
var _ json.Unmarshaler = (*UpdateTokenResponse)(nil)
var _ json.Marshaler = (*DeleteTokenRequest)(nil)
var _ json.Unmarshaler = (*DeleteTokenRequest)(nil)
var _ json.Marshaler = (*DeleteTokenResponse)(nil)
var _ json.Unmarshaler = (*DeleteTokenResponse)(nil)
var _ json.Marshaler = (*ScopeEnum)(nil)
var _ json.Unmarshaler = (*ScopeEnum)(nil)

// GetTokenRequest implement json.Marshaler.
func (m *GetTokenRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetTokenRequest implement json.Marshaler.
func (m *GetTokenRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetTokenResponse implement json.Marshaler.
func (m *GetTokenResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetTokenResponse implement json.Marshaler.
func (m *GetTokenResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Token implement json.Marshaler.
func (m *Token) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Token implement json.Marshaler.
func (m *Token) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryTokensRequest implement json.Marshaler.
func (m *QueryTokensRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryTokensRequest implement json.Marshaler.
func (m *QueryTokensRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryTokensResponse implement json.Marshaler.
func (m *QueryTokensResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryTokensResponse implement json.Marshaler.
func (m *QueryTokensResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateTokenRequest implement json.Marshaler.
func (m *CreateTokenRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateTokenRequest implement json.Marshaler.
func (m *CreateTokenRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateTokenResponse implement json.Marshaler.
func (m *CreateTokenResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateTokenResponse implement json.Marshaler.
func (m *CreateTokenResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateTokenRequest implement json.Marshaler.
func (m *UpdateTokenRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateTokenRequest implement json.Marshaler.
func (m *UpdateTokenRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateTokenResponse implement json.Marshaler.
func (m *UpdateTokenResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateTokenResponse implement json.Marshaler.
func (m *UpdateTokenResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteTokenRequest implement json.Marshaler.
func (m *DeleteTokenRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteTokenRequest implement json.Marshaler.
func (m *DeleteTokenRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteTokenResponse implement json.Marshaler.
func (m *DeleteTokenResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteTokenResponse implement json.Marshaler.
func (m *DeleteTokenResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ScopeEnum implement json.Marshaler.
func (m *ScopeEnum) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ScopeEnum implement json.Marshaler.
func (m *ScopeEnum) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
