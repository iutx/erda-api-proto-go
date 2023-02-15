// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: org.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*CreateOrgRequest)(nil)
var _ json.Unmarshaler = (*CreateOrgRequest)(nil)
var _ json.Marshaler = (*CreateOrgResponse)(nil)
var _ json.Unmarshaler = (*CreateOrgResponse)(nil)
var _ json.Marshaler = (*UpdateOrgRequest)(nil)
var _ json.Unmarshaler = (*UpdateOrgRequest)(nil)
var _ json.Marshaler = (*UpdateOrgResponse)(nil)
var _ json.Unmarshaler = (*UpdateOrgResponse)(nil)
var _ json.Marshaler = (*GetOrgRequest)(nil)
var _ json.Unmarshaler = (*GetOrgRequest)(nil)
var _ json.Marshaler = (*GetOrgResponse)(nil)
var _ json.Unmarshaler = (*GetOrgResponse)(nil)
var _ json.Marshaler = (*DeleteOrgRequest)(nil)
var _ json.Unmarshaler = (*DeleteOrgRequest)(nil)
var _ json.Marshaler = (*DeleteOrgResponse)(nil)
var _ json.Unmarshaler = (*DeleteOrgResponse)(nil)
var _ json.Marshaler = (*ListOrgRequest)(nil)
var _ json.Unmarshaler = (*ListOrgRequest)(nil)
var _ json.Marshaler = (*ListOrgResponse)(nil)
var _ json.Unmarshaler = (*ListOrgResponse)(nil)
var _ json.Marshaler = (*GetOrgByDomainRequest)(nil)
var _ json.Unmarshaler = (*GetOrgByDomainRequest)(nil)
var _ json.Marshaler = (*GetOrgByDomainResponse)(nil)
var _ json.Unmarshaler = (*GetOrgByDomainResponse)(nil)
var _ json.Marshaler = (*ChangeCurrentOrgRequest)(nil)
var _ json.Unmarshaler = (*ChangeCurrentOrgRequest)(nil)
var _ json.Marshaler = (*ChangeCurrentOrgResponse)(nil)
var _ json.Unmarshaler = (*ChangeCurrentOrgResponse)(nil)
var _ json.Marshaler = (*OrgClusterRelationCreateRequest)(nil)
var _ json.Unmarshaler = (*OrgClusterRelationCreateRequest)(nil)
var _ json.Marshaler = (*OrgClusterRelationCreateResponse)(nil)
var _ json.Unmarshaler = (*OrgClusterRelationCreateResponse)(nil)
var _ json.Marshaler = (*ListOrgClusterRelationRequest)(nil)
var _ json.Unmarshaler = (*ListOrgClusterRelationRequest)(nil)
var _ json.Marshaler = (*ListOrgClusterRelationResponse)(nil)
var _ json.Unmarshaler = (*ListOrgClusterRelationResponse)(nil)
var _ json.Marshaler = (*OrgClusterRelation)(nil)
var _ json.Unmarshaler = (*OrgClusterRelation)(nil)
var _ json.Marshaler = (*SetReleaseCrossClusterRequest)(nil)
var _ json.Unmarshaler = (*SetReleaseCrossClusterRequest)(nil)
var _ json.Marshaler = (*SetReleaseCrossClusterResponse)(nil)
var _ json.Unmarshaler = (*SetReleaseCrossClusterResponse)(nil)
var _ json.Marshaler = (*GenVerifyCodeRequest)(nil)
var _ json.Unmarshaler = (*GenVerifyCodeRequest)(nil)
var _ json.Marshaler = (*GenVerifyCodeResponse)(nil)
var _ json.Unmarshaler = (*GenVerifyCodeResponse)(nil)
var _ json.Marshaler = (*SetNotifyConfigRequest)(nil)
var _ json.Unmarshaler = (*SetNotifyConfigRequest)(nil)
var _ json.Marshaler = (*SetNotifyConfigResponse)(nil)
var _ json.Unmarshaler = (*SetNotifyConfigResponse)(nil)
var _ json.Marshaler = (*GetNotifyConfigRequest)(nil)
var _ json.Unmarshaler = (*GetNotifyConfigRequest)(nil)
var _ json.Marshaler = (*GetNotifyConfigResponse)(nil)
var _ json.Unmarshaler = (*GetNotifyConfigResponse)(nil)
var _ json.Marshaler = (*NotifyConfig)(nil)
var _ json.Unmarshaler = (*NotifyConfig)(nil)
var _ json.Marshaler = (*GetOrgClusterRelationsByOrgRequest)(nil)
var _ json.Unmarshaler = (*GetOrgClusterRelationsByOrgRequest)(nil)
var _ json.Marshaler = (*GetOrgClusterRelationsByOrgResponse)(nil)
var _ json.Unmarshaler = (*GetOrgClusterRelationsByOrgResponse)(nil)
var _ json.Marshaler = (*OrgConfig)(nil)
var _ json.Unmarshaler = (*OrgConfig)(nil)
var _ json.Marshaler = (*DereferenceClusterRequest)(nil)
var _ json.Unmarshaler = (*DereferenceClusterRequest)(nil)
var _ json.Marshaler = (*DereferenceClusterResponse)(nil)
var _ json.Unmarshaler = (*DereferenceClusterResponse)(nil)
var _ json.Marshaler = (*AuditMessage)(nil)
var _ json.Unmarshaler = (*AuditMessage)(nil)
var _ json.Marshaler = (*BlockoutConfig)(nil)
var _ json.Unmarshaler = (*BlockoutConfig)(nil)
var _ json.Marshaler = (*Org)(nil)
var _ json.Unmarshaler = (*Org)(nil)

// CreateOrgRequest implement json.Marshaler.
func (m *CreateOrgRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateOrgRequest implement json.Marshaler.
func (m *CreateOrgRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateOrgResponse implement json.Marshaler.
func (m *CreateOrgResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateOrgResponse implement json.Marshaler.
func (m *CreateOrgResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateOrgRequest implement json.Marshaler.
func (m *UpdateOrgRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateOrgRequest implement json.Marshaler.
func (m *UpdateOrgRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateOrgResponse implement json.Marshaler.
func (m *UpdateOrgResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateOrgResponse implement json.Marshaler.
func (m *UpdateOrgResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetOrgRequest implement json.Marshaler.
func (m *GetOrgRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetOrgRequest implement json.Marshaler.
func (m *GetOrgRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetOrgResponse implement json.Marshaler.
func (m *GetOrgResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetOrgResponse implement json.Marshaler.
func (m *GetOrgResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteOrgRequest implement json.Marshaler.
func (m *DeleteOrgRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteOrgRequest implement json.Marshaler.
func (m *DeleteOrgRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteOrgResponse implement json.Marshaler.
func (m *DeleteOrgResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteOrgResponse implement json.Marshaler.
func (m *DeleteOrgResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListOrgRequest implement json.Marshaler.
func (m *ListOrgRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListOrgRequest implement json.Marshaler.
func (m *ListOrgRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListOrgResponse implement json.Marshaler.
func (m *ListOrgResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListOrgResponse implement json.Marshaler.
func (m *ListOrgResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetOrgByDomainRequest implement json.Marshaler.
func (m *GetOrgByDomainRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetOrgByDomainRequest implement json.Marshaler.
func (m *GetOrgByDomainRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetOrgByDomainResponse implement json.Marshaler.
func (m *GetOrgByDomainResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetOrgByDomainResponse implement json.Marshaler.
func (m *GetOrgByDomainResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ChangeCurrentOrgRequest implement json.Marshaler.
func (m *ChangeCurrentOrgRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ChangeCurrentOrgRequest implement json.Marshaler.
func (m *ChangeCurrentOrgRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ChangeCurrentOrgResponse implement json.Marshaler.
func (m *ChangeCurrentOrgResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ChangeCurrentOrgResponse implement json.Marshaler.
func (m *ChangeCurrentOrgResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OrgClusterRelationCreateRequest implement json.Marshaler.
func (m *OrgClusterRelationCreateRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OrgClusterRelationCreateRequest implement json.Marshaler.
func (m *OrgClusterRelationCreateRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OrgClusterRelationCreateResponse implement json.Marshaler.
func (m *OrgClusterRelationCreateResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OrgClusterRelationCreateResponse implement json.Marshaler.
func (m *OrgClusterRelationCreateResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListOrgClusterRelationRequest implement json.Marshaler.
func (m *ListOrgClusterRelationRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListOrgClusterRelationRequest implement json.Marshaler.
func (m *ListOrgClusterRelationRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListOrgClusterRelationResponse implement json.Marshaler.
func (m *ListOrgClusterRelationResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListOrgClusterRelationResponse implement json.Marshaler.
func (m *ListOrgClusterRelationResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OrgClusterRelation implement json.Marshaler.
func (m *OrgClusterRelation) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OrgClusterRelation implement json.Marshaler.
func (m *OrgClusterRelation) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SetReleaseCrossClusterRequest implement json.Marshaler.
func (m *SetReleaseCrossClusterRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SetReleaseCrossClusterRequest implement json.Marshaler.
func (m *SetReleaseCrossClusterRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SetReleaseCrossClusterResponse implement json.Marshaler.
func (m *SetReleaseCrossClusterResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SetReleaseCrossClusterResponse implement json.Marshaler.
func (m *SetReleaseCrossClusterResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GenVerifyCodeRequest implement json.Marshaler.
func (m *GenVerifyCodeRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GenVerifyCodeRequest implement json.Marshaler.
func (m *GenVerifyCodeRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GenVerifyCodeResponse implement json.Marshaler.
func (m *GenVerifyCodeResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GenVerifyCodeResponse implement json.Marshaler.
func (m *GenVerifyCodeResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SetNotifyConfigRequest implement json.Marshaler.
func (m *SetNotifyConfigRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SetNotifyConfigRequest implement json.Marshaler.
func (m *SetNotifyConfigRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// SetNotifyConfigResponse implement json.Marshaler.
func (m *SetNotifyConfigResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// SetNotifyConfigResponse implement json.Marshaler.
func (m *SetNotifyConfigResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyConfigRequest implement json.Marshaler.
func (m *GetNotifyConfigRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyConfigRequest implement json.Marshaler.
func (m *GetNotifyConfigRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyConfigResponse implement json.Marshaler.
func (m *GetNotifyConfigResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyConfigResponse implement json.Marshaler.
func (m *GetNotifyConfigResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyConfig implement json.Marshaler.
func (m *NotifyConfig) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyConfig implement json.Marshaler.
func (m *NotifyConfig) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetOrgClusterRelationsByOrgRequest implement json.Marshaler.
func (m *GetOrgClusterRelationsByOrgRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetOrgClusterRelationsByOrgRequest implement json.Marshaler.
func (m *GetOrgClusterRelationsByOrgRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetOrgClusterRelationsByOrgResponse implement json.Marshaler.
func (m *GetOrgClusterRelationsByOrgResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetOrgClusterRelationsByOrgResponse implement json.Marshaler.
func (m *GetOrgClusterRelationsByOrgResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OrgConfig implement json.Marshaler.
func (m *OrgConfig) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OrgConfig implement json.Marshaler.
func (m *OrgConfig) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DereferenceClusterRequest implement json.Marshaler.
func (m *DereferenceClusterRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DereferenceClusterRequest implement json.Marshaler.
func (m *DereferenceClusterRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DereferenceClusterResponse implement json.Marshaler.
func (m *DereferenceClusterResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DereferenceClusterResponse implement json.Marshaler.
func (m *DereferenceClusterResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// AuditMessage implement json.Marshaler.
func (m *AuditMessage) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// AuditMessage implement json.Marshaler.
func (m *AuditMessage) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// BlockoutConfig implement json.Marshaler.
func (m *BlockoutConfig) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// BlockoutConfig implement json.Marshaler.
func (m *BlockoutConfig) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Org implement json.Marshaler.
func (m *Org) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Org implement json.Marshaler.
func (m *Org) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
