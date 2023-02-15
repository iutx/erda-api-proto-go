// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: dop_filetree.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*CMDB_PROJECT_FILETREE_CREATE_Request)(nil)
var _ json.Unmarshaler = (*CMDB_PROJECT_FILETREE_CREATE_Request)(nil)
var _ json.Marshaler = (*CMDB_PROJECT_FILETREE_DELETE_Request)(nil)
var _ json.Unmarshaler = (*CMDB_PROJECT_FILETREE_DELETE_Request)(nil)
var _ json.Marshaler = (*CMDB_PROJECT_FILETREE_FIND_ANCESTORS_Request)(nil)
var _ json.Unmarshaler = (*CMDB_PROJECT_FILETREE_FIND_ANCESTORS_Request)(nil)
var _ json.Marshaler = (*CMDB_PROJECT_FILETREE_FUZZY_SEARCH_Request)(nil)
var _ json.Unmarshaler = (*CMDB_PROJECT_FILETREE_FUZZY_SEARCH_Request)(nil)
var _ json.Marshaler = (*CMDB_PROJECT_FILETREE_GET_Request)(nil)
var _ json.Unmarshaler = (*CMDB_PROJECT_FILETREE_GET_Request)(nil)
var _ json.Marshaler = (*CMDB_PROJECT_FILETREE_LIST_Request)(nil)
var _ json.Unmarshaler = (*CMDB_PROJECT_FILETREE_LIST_Request)(nil)

// CMDB_PROJECT_FILETREE_CREATE_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_CREATE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CMDB_PROJECT_FILETREE_CREATE_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_CREATE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CMDB_PROJECT_FILETREE_DELETE_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_DELETE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CMDB_PROJECT_FILETREE_DELETE_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_DELETE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CMDB_PROJECT_FILETREE_FIND_ANCESTORS_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_FIND_ANCESTORS_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CMDB_PROJECT_FILETREE_FIND_ANCESTORS_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_FIND_ANCESTORS_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CMDB_PROJECT_FILETREE_FUZZY_SEARCH_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_FUZZY_SEARCH_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CMDB_PROJECT_FILETREE_FUZZY_SEARCH_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_FUZZY_SEARCH_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CMDB_PROJECT_FILETREE_GET_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_GET_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CMDB_PROJECT_FILETREE_GET_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_GET_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CMDB_PROJECT_FILETREE_LIST_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_LIST_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CMDB_PROJECT_FILETREE_LIST_Request implement json.Marshaler.
func (m *CMDB_PROJECT_FILETREE_LIST_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}