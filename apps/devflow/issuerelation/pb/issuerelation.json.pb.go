// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: issuerelation.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*IssueRelation)(nil)
var _ json.Unmarshaler = (*IssueRelation)(nil)
var _ json.Marshaler = (*CreateIssueRelationRequest)(nil)
var _ json.Unmarshaler = (*CreateIssueRelationRequest)(nil)
var _ json.Marshaler = (*CreateIssueRelationResponse)(nil)
var _ json.Unmarshaler = (*CreateIssueRelationResponse)(nil)
var _ json.Marshaler = (*DeleteIssueRelationRequest)(nil)
var _ json.Unmarshaler = (*DeleteIssueRelationRequest)(nil)
var _ json.Marshaler = (*DeleteIssueRelationResponse)(nil)
var _ json.Unmarshaler = (*DeleteIssueRelationResponse)(nil)
var _ json.Marshaler = (*ListIssueRelationRequest)(nil)
var _ json.Unmarshaler = (*ListIssueRelationRequest)(nil)
var _ json.Marshaler = (*ListIssueRelationResponse)(nil)
var _ json.Unmarshaler = (*ListIssueRelationResponse)(nil)

// IssueRelation implement json.Marshaler.
func (m *IssueRelation) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// IssueRelation implement json.Marshaler.
func (m *IssueRelation) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateIssueRelationRequest implement json.Marshaler.
func (m *CreateIssueRelationRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateIssueRelationRequest implement json.Marshaler.
func (m *CreateIssueRelationRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateIssueRelationResponse implement json.Marshaler.
func (m *CreateIssueRelationResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateIssueRelationResponse implement json.Marshaler.
func (m *CreateIssueRelationResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteIssueRelationRequest implement json.Marshaler.
func (m *DeleteIssueRelationRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteIssueRelationRequest implement json.Marshaler.
func (m *DeleteIssueRelationRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteIssueRelationResponse implement json.Marshaler.
func (m *DeleteIssueRelationResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteIssueRelationResponse implement json.Marshaler.
func (m *DeleteIssueRelationResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListIssueRelationRequest implement json.Marshaler.
func (m *ListIssueRelationRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListIssueRelationRequest implement json.Marshaler.
func (m *ListIssueRelationRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListIssueRelationResponse implement json.Marshaler.
func (m *ListIssueRelationResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListIssueRelationResponse implement json.Marshaler.
func (m *ListIssueRelationResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
