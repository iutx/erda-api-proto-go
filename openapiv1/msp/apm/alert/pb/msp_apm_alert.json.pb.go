// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: msp_apm_alert.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*APM_ALERTS_RULES_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERTS_RULES_Request)(nil)
var _ json.Marshaler = (*APM_ALERTS_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERTS_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_CREATE_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_CREATE_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_DELETE_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_DELETE_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_RECORDS_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_RECORDS_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_RECORD_ATTRS_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_RECORD_ATTRS_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_RECORD_HISTORIES_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_RECORD_HISTORIES_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_RECORD_ISSUE_CREATE_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_RECORD_ISSUE_CREATE_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_RECORD_ISSUE_UPDATE_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_RECORD_ISSUE_UPDATE_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_RECORD_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_RECORD_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_SWITCH_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_SWITCH_Request)(nil)
var _ json.Marshaler = (*APM_ALERT_UPDATE_Request)(nil)
var _ json.Unmarshaler = (*APM_ALERT_UPDATE_Request)(nil)
var _ json.Marshaler = (*APM_CUSTOMIZE_ALERTS_Request)(nil)
var _ json.Unmarshaler = (*APM_CUSTOMIZE_ALERTS_Request)(nil)
var _ json.Marshaler = (*APM_CUSTOMIZE_ALERT_CREATE_Request)(nil)
var _ json.Unmarshaler = (*APM_CUSTOMIZE_ALERT_CREATE_Request)(nil)
var _ json.Marshaler = (*APM_CUSTOMIZE_ALERT_DASH_PREVIEW_Request)(nil)
var _ json.Unmarshaler = (*APM_CUSTOMIZE_ALERT_DASH_PREVIEW_Request)(nil)
var _ json.Marshaler = (*APM_CUSTOMIZE_ALERT_DELETE_Request)(nil)
var _ json.Unmarshaler = (*APM_CUSTOMIZE_ALERT_DELETE_Request)(nil)
var _ json.Marshaler = (*APM_CUSTOMIZE_ALERT_METRICS_Request)(nil)
var _ json.Unmarshaler = (*APM_CUSTOMIZE_ALERT_METRICS_Request)(nil)
var _ json.Marshaler = (*APM_CUSTOMIZE_ALERT_NOTIFY_TARGET_Request)(nil)
var _ json.Unmarshaler = (*APM_CUSTOMIZE_ALERT_NOTIFY_TARGET_Request)(nil)
var _ json.Marshaler = (*APM_CUSTOMIZE_ALERT_Request)(nil)
var _ json.Unmarshaler = (*APM_CUSTOMIZE_ALERT_Request)(nil)
var _ json.Marshaler = (*APM_CUSTOMIZE_ALERT_SWITCH_Request)(nil)
var _ json.Unmarshaler = (*APM_CUSTOMIZE_ALERT_SWITCH_Request)(nil)
var _ json.Marshaler = (*APM_CUSTOMIZE_ALERT_UPDATE_Request)(nil)
var _ json.Unmarshaler = (*APM_CUSTOMIZE_ALERT_UPDATE_Request)(nil)

// APM_ALERTS_RULES_Request implement json.Marshaler.
func (m *APM_ALERTS_RULES_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERTS_RULES_Request implement json.Marshaler.
func (m *APM_ALERTS_RULES_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERTS_Request implement json.Marshaler.
func (m *APM_ALERTS_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERTS_Request implement json.Marshaler.
func (m *APM_ALERTS_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_CREATE_Request implement json.Marshaler.
func (m *APM_ALERT_CREATE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_CREATE_Request implement json.Marshaler.
func (m *APM_ALERT_CREATE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_DELETE_Request implement json.Marshaler.
func (m *APM_ALERT_DELETE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_DELETE_Request implement json.Marshaler.
func (m *APM_ALERT_DELETE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_RECORDS_Request implement json.Marshaler.
func (m *APM_ALERT_RECORDS_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_RECORDS_Request implement json.Marshaler.
func (m *APM_ALERT_RECORDS_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_RECORD_ATTRS_Request implement json.Marshaler.
func (m *APM_ALERT_RECORD_ATTRS_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_RECORD_ATTRS_Request implement json.Marshaler.
func (m *APM_ALERT_RECORD_ATTRS_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_RECORD_HISTORIES_Request implement json.Marshaler.
func (m *APM_ALERT_RECORD_HISTORIES_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_RECORD_HISTORIES_Request implement json.Marshaler.
func (m *APM_ALERT_RECORD_HISTORIES_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_RECORD_ISSUE_CREATE_Request implement json.Marshaler.
func (m *APM_ALERT_RECORD_ISSUE_CREATE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_RECORD_ISSUE_CREATE_Request implement json.Marshaler.
func (m *APM_ALERT_RECORD_ISSUE_CREATE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_RECORD_ISSUE_UPDATE_Request implement json.Marshaler.
func (m *APM_ALERT_RECORD_ISSUE_UPDATE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_RECORD_ISSUE_UPDATE_Request implement json.Marshaler.
func (m *APM_ALERT_RECORD_ISSUE_UPDATE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_RECORD_Request implement json.Marshaler.
func (m *APM_ALERT_RECORD_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_RECORD_Request implement json.Marshaler.
func (m *APM_ALERT_RECORD_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_Request implement json.Marshaler.
func (m *APM_ALERT_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_Request implement json.Marshaler.
func (m *APM_ALERT_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_SWITCH_Request implement json.Marshaler.
func (m *APM_ALERT_SWITCH_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_SWITCH_Request implement json.Marshaler.
func (m *APM_ALERT_SWITCH_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_ALERT_UPDATE_Request implement json.Marshaler.
func (m *APM_ALERT_UPDATE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_ALERT_UPDATE_Request implement json.Marshaler.
func (m *APM_ALERT_UPDATE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_CUSTOMIZE_ALERTS_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERTS_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_CUSTOMIZE_ALERTS_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERTS_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_CUSTOMIZE_ALERT_CREATE_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_CREATE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_CUSTOMIZE_ALERT_CREATE_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_CREATE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_CUSTOMIZE_ALERT_DASH_PREVIEW_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_DASH_PREVIEW_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_CUSTOMIZE_ALERT_DASH_PREVIEW_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_DASH_PREVIEW_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_CUSTOMIZE_ALERT_DELETE_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_DELETE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_CUSTOMIZE_ALERT_DELETE_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_DELETE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_CUSTOMIZE_ALERT_METRICS_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_METRICS_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_CUSTOMIZE_ALERT_METRICS_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_METRICS_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_CUSTOMIZE_ALERT_NOTIFY_TARGET_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_NOTIFY_TARGET_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_CUSTOMIZE_ALERT_NOTIFY_TARGET_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_NOTIFY_TARGET_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_CUSTOMIZE_ALERT_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_CUSTOMIZE_ALERT_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_CUSTOMIZE_ALERT_SWITCH_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_SWITCH_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_CUSTOMIZE_ALERT_SWITCH_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_SWITCH_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// APM_CUSTOMIZE_ALERT_UPDATE_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_UPDATE_Request) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// APM_CUSTOMIZE_ALERT_UPDATE_Request implement json.Marshaler.
func (m *APM_CUSTOMIZE_ALERT_UPDATE_Request) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
