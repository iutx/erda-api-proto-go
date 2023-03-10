// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: notify.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*CreateNotifyHistoryRequest)(nil)
var _ json.Unmarshaler = (*CreateNotifyHistoryRequest)(nil)
var _ json.Marshaler = (*NotifyTarget)(nil)
var _ json.Unmarshaler = (*NotifyTarget)(nil)
var _ json.Marshaler = (*Target)(nil)
var _ json.Unmarshaler = (*Target)(nil)
var _ json.Marshaler = (*NotifySource)(nil)
var _ json.Unmarshaler = (*NotifySource)(nil)
var _ json.Marshaler = (*CreateNotifyHistoryResponse)(nil)
var _ json.Unmarshaler = (*CreateNotifyHistoryResponse)(nil)
var _ json.Marshaler = (*QueryNotifyHistoriesRequest)(nil)
var _ json.Unmarshaler = (*QueryNotifyHistoriesRequest)(nil)
var _ json.Marshaler = (*QueryNotifyHistoriesResponse)(nil)
var _ json.Unmarshaler = (*QueryNotifyHistoriesResponse)(nil)
var _ json.Marshaler = (*QueryNotifyHistoryData)(nil)
var _ json.Unmarshaler = (*QueryNotifyHistoryData)(nil)
var _ json.Marshaler = (*NotifyHistory)(nil)
var _ json.Unmarshaler = (*NotifyHistory)(nil)
var _ json.Marshaler = (*GetNotifyStatusRequest)(nil)
var _ json.Unmarshaler = (*GetNotifyStatusRequest)(nil)
var _ json.Marshaler = (*GetNotifyStatusResponse)(nil)
var _ json.Unmarshaler = (*GetNotifyStatusResponse)(nil)
var _ json.Marshaler = (*GetNotifyHistogramRequest)(nil)
var _ json.Unmarshaler = (*GetNotifyHistogramRequest)(nil)
var _ json.Marshaler = (*GetNotifyHistogramResponse)(nil)
var _ json.Unmarshaler = (*GetNotifyHistogramResponse)(nil)
var _ json.Marshaler = (*NotifyHistogramData)(nil)
var _ json.Unmarshaler = (*NotifyHistogramData)(nil)
var _ json.Marshaler = (*StatisticValue)(nil)
var _ json.Unmarshaler = (*StatisticValue)(nil)
var _ json.Marshaler = (*QueryAlertNotifyHistoriesRequest)(nil)
var _ json.Unmarshaler = (*QueryAlertNotifyHistoriesRequest)(nil)
var _ json.Marshaler = (*QueryAlertNotifyHistoriesResponse)(nil)
var _ json.Unmarshaler = (*QueryAlertNotifyHistoriesResponse)(nil)
var _ json.Marshaler = (*AlertNotifyHistories)(nil)
var _ json.Unmarshaler = (*AlertNotifyHistories)(nil)
var _ json.Marshaler = (*AlertNotifyIndex)(nil)
var _ json.Unmarshaler = (*AlertNotifyIndex)(nil)
var _ json.Marshaler = (*GetAlertNotifyDetailRequest)(nil)
var _ json.Unmarshaler = (*GetAlertNotifyDetailRequest)(nil)
var _ json.Marshaler = (*GetAlertNotifyDetailResponse)(nil)
var _ json.Unmarshaler = (*GetAlertNotifyDetailResponse)(nil)
var _ json.Marshaler = (*AlertNotifyDetail)(nil)
var _ json.Unmarshaler = (*AlertNotifyDetail)(nil)
var _ json.Marshaler = (*GetTypeNotifyHistogramRequest)(nil)
var _ json.Unmarshaler = (*GetTypeNotifyHistogramRequest)(nil)
var _ json.Marshaler = (*GetTypeNotifyHistogramResponse)(nil)
var _ json.Unmarshaler = (*GetTypeNotifyHistogramResponse)(nil)
var _ json.Marshaler = (*TypeNotifyHistogram)(nil)
var _ json.Unmarshaler = (*TypeNotifyHistogram)(nil)

// CreateNotifyHistoryRequest implement json.Marshaler.
func (m *CreateNotifyHistoryRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateNotifyHistoryRequest implement json.Marshaler.
func (m *CreateNotifyHistoryRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyTarget implement json.Marshaler.
func (m *NotifyTarget) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyTarget implement json.Marshaler.
func (m *NotifyTarget) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// Target implement json.Marshaler.
func (m *Target) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// Target implement json.Marshaler.
func (m *Target) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifySource implement json.Marshaler.
func (m *NotifySource) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifySource implement json.Marshaler.
func (m *NotifySource) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateNotifyHistoryResponse implement json.Marshaler.
func (m *CreateNotifyHistoryResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateNotifyHistoryResponse implement json.Marshaler.
func (m *CreateNotifyHistoryResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryNotifyHistoriesRequest implement json.Marshaler.
func (m *QueryNotifyHistoriesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryNotifyHistoriesRequest implement json.Marshaler.
func (m *QueryNotifyHistoriesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryNotifyHistoriesResponse implement json.Marshaler.
func (m *QueryNotifyHistoriesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryNotifyHistoriesResponse implement json.Marshaler.
func (m *QueryNotifyHistoriesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryNotifyHistoryData implement json.Marshaler.
func (m *QueryNotifyHistoryData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryNotifyHistoryData implement json.Marshaler.
func (m *QueryNotifyHistoryData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyHistory implement json.Marshaler.
func (m *NotifyHistory) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyHistory implement json.Marshaler.
func (m *NotifyHistory) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyStatusRequest implement json.Marshaler.
func (m *GetNotifyStatusRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyStatusRequest implement json.Marshaler.
func (m *GetNotifyStatusRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyStatusResponse implement json.Marshaler.
func (m *GetNotifyStatusResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyStatusResponse implement json.Marshaler.
func (m *GetNotifyStatusResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyHistogramRequest implement json.Marshaler.
func (m *GetNotifyHistogramRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyHistogramRequest implement json.Marshaler.
func (m *GetNotifyHistogramRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetNotifyHistogramResponse implement json.Marshaler.
func (m *GetNotifyHistogramResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetNotifyHistogramResponse implement json.Marshaler.
func (m *GetNotifyHistogramResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// NotifyHistogramData implement json.Marshaler.
func (m *NotifyHistogramData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// NotifyHistogramData implement json.Marshaler.
func (m *NotifyHistogramData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// StatisticValue implement json.Marshaler.
func (m *StatisticValue) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// StatisticValue implement json.Marshaler.
func (m *StatisticValue) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryAlertNotifyHistoriesRequest implement json.Marshaler.
func (m *QueryAlertNotifyHistoriesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryAlertNotifyHistoriesRequest implement json.Marshaler.
func (m *QueryAlertNotifyHistoriesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryAlertNotifyHistoriesResponse implement json.Marshaler.
func (m *QueryAlertNotifyHistoriesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryAlertNotifyHistoriesResponse implement json.Marshaler.
func (m *QueryAlertNotifyHistoriesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// AlertNotifyHistories implement json.Marshaler.
func (m *AlertNotifyHistories) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// AlertNotifyHistories implement json.Marshaler.
func (m *AlertNotifyHistories) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// AlertNotifyIndex implement json.Marshaler.
func (m *AlertNotifyIndex) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// AlertNotifyIndex implement json.Marshaler.
func (m *AlertNotifyIndex) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertNotifyDetailRequest implement json.Marshaler.
func (m *GetAlertNotifyDetailRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertNotifyDetailRequest implement json.Marshaler.
func (m *GetAlertNotifyDetailRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertNotifyDetailResponse implement json.Marshaler.
func (m *GetAlertNotifyDetailResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertNotifyDetailResponse implement json.Marshaler.
func (m *GetAlertNotifyDetailResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// AlertNotifyDetail implement json.Marshaler.
func (m *AlertNotifyDetail) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// AlertNotifyDetail implement json.Marshaler.
func (m *AlertNotifyDetail) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetTypeNotifyHistogramRequest implement json.Marshaler.
func (m *GetTypeNotifyHistogramRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetTypeNotifyHistogramRequest implement json.Marshaler.
func (m *GetTypeNotifyHistogramRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetTypeNotifyHistogramResponse implement json.Marshaler.
func (m *GetTypeNotifyHistogramResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetTypeNotifyHistogramResponse implement json.Marshaler.
func (m *GetTypeNotifyHistogramResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// TypeNotifyHistogram implement json.Marshaler.
func (m *TypeNotifyHistogram) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// TypeNotifyHistogram implement json.Marshaler.
func (m *TypeNotifyHistogram) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
