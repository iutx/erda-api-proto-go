// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: core-project.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*CheckProjectExistReq)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CheckProjectExistResp)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetProjectByIDReq)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ProjectDto)(nil)

// CheckProjectExistReq implement urlenc.URLValuesUnmarshaler.
func (m *CheckProjectExistReq) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			}
		}
	}
	return nil
}

// CheckProjectExistResp implement urlenc.URLValuesUnmarshaler.
func (m *CheckProjectExistResp) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ok":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Ok = val
			}
		}
	}
	return nil
}

// GetProjectByIDReq implement urlenc.URLValuesUnmarshaler.
func (m *GetProjectByIDReq) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "userID":
				m.UserID = &vals[0]
			}
		}
	}
	return nil
}

// ProjectDto implement urlenc.URLValuesUnmarshaler.
func (m *ProjectDto) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "name":
				m.Name = vals[0]
			case "displayName":
				m.DisplayName = vals[0]
			case "orgID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.OrgID = val
			case "creatorID":
				m.CreatorID = vals[0]
			case "logo":
				m.Logo = vals[0]
			case "desc":
				m.Desc = vals[0]
			case "activeTime":
				if m.ActiveTime == nil {
					m.ActiveTime = &timestamppb.Timestamp{}
				}
			case "activeTime.seconds":
				if m.ActiveTime == nil {
					m.ActiveTime = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ActiveTime.Seconds = val
			case "activeTime.nanos":
				if m.ActiveTime == nil {
					m.ActiveTime = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.ActiveTime.Nanos = int32(val)
			case "isPublic":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsPublic = val
			case "createdTime":
				if m.CreatedTime == nil {
					m.CreatedTime = &timestamppb.Timestamp{}
				}
			case "createdTime.seconds":
				if m.CreatedTime == nil {
					m.CreatedTime = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreatedTime.Seconds = val
			case "createdTime.nanos":
				if m.CreatedTime == nil {
					m.CreatedTime = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.CreatedTime.Nanos = int32(val)
			case "updatedTime":
				if m.UpdatedTime == nil {
					m.UpdatedTime = &timestamppb.Timestamp{}
				}
			case "updatedTime.seconds":
				if m.UpdatedTime == nil {
					m.UpdatedTime = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.UpdatedTime.Seconds = val
			case "updatedTime.nanos":
				if m.UpdatedTime == nil {
					m.UpdatedTime = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.UpdatedTime.Nanos = int32(val)
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}
