// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: task_error.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*ErrorLogListRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ErrorLogListResponseData)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ErrorLog)(nil)

// ErrorLogListRequest implement urlenc.URLValuesUnmarshaler.
func (m *ErrorLogListRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scopeType":
				m.ScopeType = vals[0]
			case "scopeId":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ScopeId = val
			case "resourceType":
				m.ResourceType = vals[0]
			case "resourceId":
				m.ResourceId = vals[0]
			case "startTime":
				m.StartTime = vals[0]
			}
		}
	}
	return nil
}

// ErrorLogListResponseData implement urlenc.URLValuesUnmarshaler.
func (m *ErrorLogListResponseData) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ErrorLog implement urlenc.URLValuesUnmarshaler.
func (m *ErrorLog) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "level":
				m.Level = vals[0]
			case "resourceType":
				m.ResourceType = vals[0]
			case "resourceId":
				m.ResourceId = vals[0]
			case "occurrenceTime":
				m.OccurrenceTime = vals[0]
			case "humanLog":
				m.HumanLog = vals[0]
			case "primevalLog":
				m.PrimevalLog = vals[0]
			case "deDupId":
				m.DeDupId = vals[0]
			}
		}
	}
	return nil
}
