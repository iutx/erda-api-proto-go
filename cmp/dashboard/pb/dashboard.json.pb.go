// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: dashboard.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*GetClustersResourcesRequest)(nil)
var _ json.Unmarshaler = (*GetClustersResourcesRequest)(nil)
var _ json.Marshaler = (*GetClusterResourcesResponse)(nil)
var _ json.Unmarshaler = (*GetClusterResourcesResponse)(nil)
var _ json.Marshaler = (*ClusterResourceDetail)(nil)
var _ json.Unmarshaler = (*ClusterResourceDetail)(nil)
var _ json.Marshaler = (*HostResourceDetail)(nil)
var _ json.Unmarshaler = (*HostResourceDetail)(nil)
var _ json.Marshaler = (*GetNamespacesResourcesRequest)(nil)
var _ json.Unmarshaler = (*GetNamespacesResourcesRequest)(nil)
var _ json.Marshaler = (*ClusterNamespacePair)(nil)
var _ json.Unmarshaler = (*ClusterNamespacePair)(nil)
var _ json.Marshaler = (*GetNamespacesResourcesResponse)(nil)
var _ json.Unmarshaler = (*GetNamespacesResourcesResponse)(nil)
var _ json.Marshaler = (*ClusterResourceItem)(nil)
var _ json.Unmarshaler = (*ClusterResourceItem)(nil)
var _ json.Marshaler = (*NamespaceResourceDetail)(nil)
var _ json.Unmarshaler = (*NamespaceResourceDetail)(nil)
var _ json.Marshaler = (*GetPodsByLabelsRequest)(nil)
var _ json.Unmarshaler = (*GetPodsByLabelsRequest)(nil)
var _ json.Marshaler = (*GetPodsByLabelsResponse)(nil)
var _ json.Unmarshaler = (*GetPodsByLabelsResponse)(nil)
var _ json.Marshaler = (*GetPodsByLabelsItem)(nil)
var _ json.Unmarshaler = (*GetPodsByLabelsItem)(nil)

// GetClustersResourcesRequest implement json.Marshaler.
func (m *GetClustersResourcesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetClustersResourcesRequest implement json.Marshaler.
func (m *GetClustersResourcesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetClusterResourcesResponse implement json.Marshaler.
func (m *GetClusterResourcesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetClusterResourcesResponse implement json.Marshaler.
func (m *GetClusterResourcesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ClusterResourceDetail implement json.Marshaler.
func (m *ClusterResourceDetail) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ClusterResourceDetail implement json.Marshaler.
func (m *ClusterResourceDetail) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// HostResourceDetail implement json.Marshaler.
func (m *HostResourceDetail) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// HostResourceDetail implement json.Marshaler.
func (m *HostResourceDetail) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNamespacesResourcesRequest implement json.Marshaler.
func (m *GetNamespacesResourcesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNamespacesResourcesRequest implement json.Marshaler.
func (m *GetNamespacesResourcesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ClusterNamespacePair implement json.Marshaler.
func (m *ClusterNamespacePair) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ClusterNamespacePair implement json.Marshaler.
func (m *ClusterNamespacePair) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNamespacesResourcesResponse implement json.Marshaler.
func (m *GetNamespacesResourcesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNamespacesResourcesResponse implement json.Marshaler.
func (m *GetNamespacesResourcesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ClusterResourceItem implement json.Marshaler.
func (m *ClusterResourceItem) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ClusterResourceItem implement json.Marshaler.
func (m *ClusterResourceItem) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NamespaceResourceDetail implement json.Marshaler.
func (m *NamespaceResourceDetail) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NamespaceResourceDetail implement json.Marshaler.
func (m *NamespaceResourceDetail) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetPodsByLabelsRequest implement json.Marshaler.
func (m *GetPodsByLabelsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetPodsByLabelsRequest implement json.Marshaler.
func (m *GetPodsByLabelsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetPodsByLabelsResponse implement json.Marshaler.
func (m *GetPodsByLabelsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetPodsByLabelsResponse implement json.Marshaler.
func (m *GetPodsByLabelsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetPodsByLabelsItem implement json.Marshaler.
func (m *GetPodsByLabelsItem) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetPodsByLabelsItem implement json.Marshaler.
func (m *GetPodsByLabelsItem) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
