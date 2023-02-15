// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: rule.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*FireRequest)(nil)
var _ json.Unmarshaler = (*FireRequest)(nil)
var _ json.Marshaler = (*Config)(nil)
var _ json.Unmarshaler = (*Config)(nil)
var _ json.Marshaler = (*FireResponse)(nil)
var _ json.Unmarshaler = (*FireResponse)(nil)
var _ json.Marshaler = (*CreateRuleRequest)(nil)
var _ json.Unmarshaler = (*CreateRuleRequest)(nil)
var _ json.Marshaler = (*CreateRuleResponse)(nil)
var _ json.Unmarshaler = (*CreateRuleResponse)(nil)
var _ json.Marshaler = (*GetRuleRequest)(nil)
var _ json.Unmarshaler = (*GetRuleRequest)(nil)
var _ json.Marshaler = (*GetRuleResponse)(nil)
var _ json.Unmarshaler = (*GetRuleResponse)(nil)
var _ json.Marshaler = (*Rule)(nil)
var _ json.Unmarshaler = (*Rule)(nil)
var _ json.Marshaler = (*ActionParams)(nil)
var _ json.Unmarshaler = (*ActionParams)(nil)
var _ json.Marshaler = (*ActionNode)(nil)
var _ json.Unmarshaler = (*ActionNode)(nil)
var _ json.Marshaler = (*DingTalkConfig)(nil)
var _ json.Unmarshaler = (*DingTalkConfig)(nil)
var _ json.Marshaler = (*UpdateRuleRequest)(nil)
var _ json.Unmarshaler = (*UpdateRuleRequest)(nil)
var _ json.Marshaler = (*UpdateRuleResponse)(nil)
var _ json.Unmarshaler = (*UpdateRuleResponse)(nil)
var _ json.Marshaler = (*ListRulesRequest)(nil)
var _ json.Unmarshaler = (*ListRulesRequest)(nil)
var _ json.Marshaler = (*ListRulesResponse)(nil)
var _ json.Unmarshaler = (*ListRulesResponse)(nil)
var _ json.Marshaler = (*ListRulesResponseData)(nil)
var _ json.Unmarshaler = (*ListRulesResponseData)(nil)
var _ json.Marshaler = (*DeleteRuleRequest)(nil)
var _ json.Unmarshaler = (*DeleteRuleRequest)(nil)
var _ json.Marshaler = (*DeleteRuleResponse)(nil)
var _ json.Unmarshaler = (*DeleteRuleResponse)(nil)
var _ json.Marshaler = (*ListRuleExecHistoryRequest)(nil)
var _ json.Unmarshaler = (*ListRuleExecHistoryRequest)(nil)
var _ json.Marshaler = (*ListRuleExecHistoryResponse)(nil)
var _ json.Unmarshaler = (*ListRuleExecHistoryResponse)(nil)
var _ json.Marshaler = (*ListRuleExecHistoryResponseData)(nil)
var _ json.Unmarshaler = (*ListRuleExecHistoryResponseData)(nil)
var _ json.Marshaler = (*RuleExecHistory)(nil)
var _ json.Unmarshaler = (*RuleExecHistory)(nil)

// FireRequest implement json.Marshaler.
func (m *FireRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// FireRequest implement json.Marshaler.
func (m *FireRequest) UnmarshalJSON(b []byte) error {
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

// FireResponse implement json.Marshaler.
func (m *FireResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// FireResponse implement json.Marshaler.
func (m *FireResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateRuleRequest implement json.Marshaler.
func (m *CreateRuleRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateRuleRequest implement json.Marshaler.
func (m *CreateRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateRuleResponse implement json.Marshaler.
func (m *CreateRuleResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateRuleResponse implement json.Marshaler.
func (m *CreateRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetRuleRequest implement json.Marshaler.
func (m *GetRuleRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetRuleRequest implement json.Marshaler.
func (m *GetRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetRuleResponse implement json.Marshaler.
func (m *GetRuleResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetRuleResponse implement json.Marshaler.
func (m *GetRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Rule implement json.Marshaler.
func (m *Rule) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Rule implement json.Marshaler.
func (m *Rule) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ActionParams implement json.Marshaler.
func (m *ActionParams) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ActionParams implement json.Marshaler.
func (m *ActionParams) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ActionNode implement json.Marshaler.
func (m *ActionNode) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ActionNode implement json.Marshaler.
func (m *ActionNode) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DingTalkConfig implement json.Marshaler.
func (m *DingTalkConfig) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DingTalkConfig implement json.Marshaler.
func (m *DingTalkConfig) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateRuleRequest implement json.Marshaler.
func (m *UpdateRuleRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateRuleRequest implement json.Marshaler.
func (m *UpdateRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateRuleResponse implement json.Marshaler.
func (m *UpdateRuleResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateRuleResponse implement json.Marshaler.
func (m *UpdateRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListRulesRequest implement json.Marshaler.
func (m *ListRulesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListRulesRequest implement json.Marshaler.
func (m *ListRulesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListRulesResponse implement json.Marshaler.
func (m *ListRulesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListRulesResponse implement json.Marshaler.
func (m *ListRulesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListRulesResponseData implement json.Marshaler.
func (m *ListRulesResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListRulesResponseData implement json.Marshaler.
func (m *ListRulesResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteRuleRequest implement json.Marshaler.
func (m *DeleteRuleRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteRuleRequest implement json.Marshaler.
func (m *DeleteRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteRuleResponse implement json.Marshaler.
func (m *DeleteRuleResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteRuleResponse implement json.Marshaler.
func (m *DeleteRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListRuleExecHistoryRequest implement json.Marshaler.
func (m *ListRuleExecHistoryRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListRuleExecHistoryRequest implement json.Marshaler.
func (m *ListRuleExecHistoryRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListRuleExecHistoryResponse implement json.Marshaler.
func (m *ListRuleExecHistoryResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListRuleExecHistoryResponse implement json.Marshaler.
func (m *ListRuleExecHistoryResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ListRuleExecHistoryResponseData implement json.Marshaler.
func (m *ListRuleExecHistoryResponseData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ListRuleExecHistoryResponseData implement json.Marshaler.
func (m *ListRuleExecHistoryResponseData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RuleExecHistory implement json.Marshaler.
func (m *RuleExecHistory) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RuleExecHistory implement json.Marshaler.
func (m *RuleExecHistory) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
