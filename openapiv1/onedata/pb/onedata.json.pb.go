// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: onedata.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"

	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*AtomicIndexDTO)(nil)
var _ json.Unmarshaler = (*AtomicIndexDTO)(nil)
var _ json.Marshaler = (*AttrData)(nil)
var _ json.Unmarshaler = (*AttrData)(nil)
var _ json.Marshaler = (*BaseParam)(nil)
var _ json.Unmarshaler = (*BaseParam)(nil)
var _ json.Marshaler = (*BusinessDomainDTO)(nil)
var _ json.Unmarshaler = (*BusinessDomainDTO)(nil)
var _ json.Marshaler = (*BusinessProcessDTO)(nil)
var _ json.Unmarshaler = (*BusinessProcessDTO)(nil)
var _ json.Marshaler = (*BusinessProcessData)(nil)
var _ json.Unmarshaler = (*BusinessProcessData)(nil)
var _ json.Marshaler = (*DataDomainDTO)(nil)
var _ json.Unmarshaler = (*DataDomainDTO)(nil)
var _ json.Marshaler = (*DimDTO)(nil)
var _ json.Unmarshaler = (*DimDTO)(nil)
var _ json.Marshaler = (*ExtBaseParam)(nil)
var _ json.Unmarshaler = (*ExtBaseParam)(nil)
var _ json.Marshaler = (*MarketDomainDTO)(nil)
var _ json.Unmarshaler = (*MarketDomainDTO)(nil)
var _ json.Marshaler = (*OneDataAnalysisBussProcRequest)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisBussProcRequest)(nil)
var _ json.Marshaler = (*OneDataAnalysisBussProcResponse)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisBussProcResponse)(nil)
var _ json.Marshaler = (*OneDataAnalysisBussProcsRequest)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisBussProcsRequest)(nil)
var _ json.Marshaler = (*OneDataAnalysisBussProcsResponse)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisBussProcsResponse)(nil)
var _ json.Marshaler = (*OneDataAnalysisDimRequest)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisDimRequest)(nil)
var _ json.Marshaler = (*OneDataAnalysisFuzzyAttrsRequest)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisFuzzyAttrsRequest)(nil)
var _ json.Marshaler = (*OneDataAnalysisFuzzyAttrsResponse)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisFuzzyAttrsResponse)(nil)
var _ json.Marshaler = (*OneDataAnalysisOutputTablesRequest)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisOutputTablesRequest)(nil)
var _ json.Marshaler = (*OneDataAnalysisOutputTablesResponse)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisOutputTablesResponse)(nil)
var _ json.Marshaler = (*OneDataAnalysisRequest)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisRequest)(nil)
var _ json.Marshaler = (*OneDataAnalysisResponse)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisResponse)(nil)
var _ json.Marshaler = (*OneDataAnalysisStarRequest)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisStarRequest)(nil)
var _ json.Marshaler = (*OneDataAnalysisStarResponse)(nil)
var _ json.Unmarshaler = (*OneDataAnalysisStarResponse)(nil)
var _ json.Marshaler = (*OneDataDTO)(nil)
var _ json.Unmarshaler = (*OneDataDTO)(nil)
var _ json.Marshaler = (*OutputTableData)(nil)
var _ json.Unmarshaler = (*OutputTableData)(nil)
var _ json.Marshaler = (*RelationDTO)(nil)
var _ json.Unmarshaler = (*RelationDTO)(nil)
var _ json.Marshaler = (*StarDTO)(nil)
var _ json.Unmarshaler = (*StarDTO)(nil)

// AtomicIndexDTO implement json.Marshaler.
func (m *AtomicIndexDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// AtomicIndexDTO implement json.Marshaler.
func (m *AtomicIndexDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// AttrData implement json.Marshaler.
func (m *AttrData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// AttrData implement json.Marshaler.
func (m *AttrData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// BaseParam implement json.Marshaler.
func (m *BaseParam) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// BaseParam implement json.Marshaler.
func (m *BaseParam) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// BusinessDomainDTO implement json.Marshaler.
func (m *BusinessDomainDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// BusinessDomainDTO implement json.Marshaler.
func (m *BusinessDomainDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// BusinessProcessDTO implement json.Marshaler.
func (m *BusinessProcessDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// BusinessProcessDTO implement json.Marshaler.
func (m *BusinessProcessDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// BusinessProcessData implement json.Marshaler.
func (m *BusinessProcessData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// BusinessProcessData implement json.Marshaler.
func (m *BusinessProcessData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DataDomainDTO implement json.Marshaler.
func (m *DataDomainDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DataDomainDTO implement json.Marshaler.
func (m *DataDomainDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DimDTO implement json.Marshaler.
func (m *DimDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DimDTO implement json.Marshaler.
func (m *DimDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// ExtBaseParam implement json.Marshaler.
func (m *ExtBaseParam) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// ExtBaseParam implement json.Marshaler.
func (m *ExtBaseParam) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// MarketDomainDTO implement json.Marshaler.
func (m *MarketDomainDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// MarketDomainDTO implement json.Marshaler.
func (m *MarketDomainDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisBussProcRequest implement json.Marshaler.
func (m *OneDataAnalysisBussProcRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisBussProcRequest implement json.Marshaler.
func (m *OneDataAnalysisBussProcRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisBussProcResponse implement json.Marshaler.
func (m *OneDataAnalysisBussProcResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisBussProcResponse implement json.Marshaler.
func (m *OneDataAnalysisBussProcResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisBussProcsRequest implement json.Marshaler.
func (m *OneDataAnalysisBussProcsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisBussProcsRequest implement json.Marshaler.
func (m *OneDataAnalysisBussProcsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisBussProcsResponse implement json.Marshaler.
func (m *OneDataAnalysisBussProcsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisBussProcsResponse implement json.Marshaler.
func (m *OneDataAnalysisBussProcsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisDimRequest implement json.Marshaler.
func (m *OneDataAnalysisDimRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisDimRequest implement json.Marshaler.
func (m *OneDataAnalysisDimRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisFuzzyAttrsRequest implement json.Marshaler.
func (m *OneDataAnalysisFuzzyAttrsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisFuzzyAttrsRequest implement json.Marshaler.
func (m *OneDataAnalysisFuzzyAttrsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisFuzzyAttrsResponse implement json.Marshaler.
func (m *OneDataAnalysisFuzzyAttrsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisFuzzyAttrsResponse implement json.Marshaler.
func (m *OneDataAnalysisFuzzyAttrsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisOutputTablesRequest implement json.Marshaler.
func (m *OneDataAnalysisOutputTablesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisOutputTablesRequest implement json.Marshaler.
func (m *OneDataAnalysisOutputTablesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisOutputTablesResponse implement json.Marshaler.
func (m *OneDataAnalysisOutputTablesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisOutputTablesResponse implement json.Marshaler.
func (m *OneDataAnalysisOutputTablesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisRequest implement json.Marshaler.
func (m *OneDataAnalysisRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisRequest implement json.Marshaler.
func (m *OneDataAnalysisRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisResponse implement json.Marshaler.
func (m *OneDataAnalysisResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisResponse implement json.Marshaler.
func (m *OneDataAnalysisResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisStarRequest implement json.Marshaler.
func (m *OneDataAnalysisStarRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisStarRequest implement json.Marshaler.
func (m *OneDataAnalysisStarRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataAnalysisStarResponse implement json.Marshaler.
func (m *OneDataAnalysisStarResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataAnalysisStarResponse implement json.Marshaler.
func (m *OneDataAnalysisStarResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OneDataDTO implement json.Marshaler.
func (m *OneDataDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OneDataDTO implement json.Marshaler.
func (m *OneDataDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// OutputTableData implement json.Marshaler.
func (m *OutputTableData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// OutputTableData implement json.Marshaler.
func (m *OutputTableData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// RelationDTO implement json.Marshaler.
func (m *RelationDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// RelationDTO implement json.Marshaler.
func (m *RelationDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// StarDTO implement json.Marshaler.
func (m *StarDTO) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// StarDTO implement json.Marshaler.
func (m *StarDTO) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
