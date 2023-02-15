// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: action.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*PipelineActionListRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineActionSaveRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineActionDeleteRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineActionDeleteResponse)(nil)

// PipelineActionListRequest implement urlenc.URLValuesUnmarshaler.
func (m *PipelineActionListRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "all":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.All = val
			case "labels":
				m.Labels = vals[0]
			}
		}
	}
	return nil
}

// PipelineActionSaveRequest implement urlenc.URLValuesUnmarshaler.
func (m *PipelineActionSaveRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "readme":
				m.Readme = vals[0]
			case "dice":
				m.Dice = vals[0]
			case "spec":
				m.Spec = vals[0]
			}
		}
	}
	return nil
}

// PipelineActionDeleteRequest implement urlenc.URLValuesUnmarshaler.
func (m *PipelineActionDeleteRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "version":
				m.Version = vals[0]
			case "location":
				m.Location = vals[0]
			}
		}
	}
	return nil
}

// PipelineActionDeleteResponse implement urlenc.URLValuesUnmarshaler.
func (m *PipelineActionDeleteResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
