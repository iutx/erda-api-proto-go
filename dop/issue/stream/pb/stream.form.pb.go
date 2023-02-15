// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: stream.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	pb "github.com/erda-project/erda-proto-go/common/pb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*CommentIssueStreamCreateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*MRCommentInfo)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CommentIssueStreamBatchCreateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CommentIssueStreamBatchCreateResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*IssueStreamCreateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*IssueStreamCreateResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PagingIssueStreamsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PagingIssueStreamsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*IssueStreamPagingResponseData)(nil)
var _ urlenc.URLValuesUnmarshaler = (*IssueStream)(nil)

// CommentIssueStreamCreateRequest implement urlenc.URLValuesUnmarshaler.
func (m *CommentIssueStreamCreateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "issueID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.IssueID = val
			case "type":
				m.Type = vals[0]
			case "content":
				m.Content = vals[0]
			case "mrInfo":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
			case "mrInfo.appID":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MrInfo.AppID = val
			case "mrInfo.mrID":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MrInfo.MrID = val
			case "mrInfo.mrTitle":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
				m.MrInfo.MrTitle = vals[0]
			case "userID":
				m.UserID = vals[0]
			}
		}
	}
	return nil
}

// MRCommentInfo implement urlenc.URLValuesUnmarshaler.
func (m *MRCommentInfo) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "appID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.AppID = val
			case "mrID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MrID = val
			case "mrTitle":
				m.MrTitle = vals[0]
			}
		}
	}
	return nil
}

// CommentIssueStreamBatchCreateRequest implement urlenc.URLValuesUnmarshaler.
func (m *CommentIssueStreamBatchCreateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CommentIssueStreamBatchCreateResponse implement urlenc.URLValuesUnmarshaler.
func (m *CommentIssueStreamBatchCreateResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// IssueStreamCreateRequest implement urlenc.URLValuesUnmarshaler.
func (m *IssueStreamCreateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "issueID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.IssueID = val
			case "type":
				m.Type = vals[0]
			case "content":
				m.Content = vals[0]
			case "mrInfo":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
			case "mrInfo.appID":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MrInfo.AppID = val
			case "mrInfo.mrID":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MrInfo.MrID = val
			case "mrInfo.mrTitle":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
				m.MrInfo.MrTitle = vals[0]
			case "identityInfo":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
			case "identityInfo.userID":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
				m.IdentityInfo.UserID = vals[0]
			case "identityInfo.internalClient":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
				m.IdentityInfo.InternalClient = vals[0]
			case "identityInfo.orgID":
				if m.IdentityInfo == nil {
					m.IdentityInfo = &pb.IdentityInfo{}
				}
				m.IdentityInfo.OrgID = vals[0]
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// IssueStreamCreateResponse implement urlenc.URLValuesUnmarshaler.
func (m *IssueStreamCreateResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// PagingIssueStreamsRequest implement urlenc.URLValuesUnmarshaler.
func (m *PagingIssueStreamsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "issueID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.IssueID = val
			case "pageNo":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageNo = val
			case "pageSize":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageSize = val
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// PagingIssueStreamsResponse implement urlenc.URLValuesUnmarshaler.
func (m *PagingIssueStreamsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &IssueStreamPagingResponseData{}
				}
			case "data.total":
				if m.Data == nil {
					m.Data = &IssueStreamPagingResponseData{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Total = val
			case "userIDs":
				m.UserIDs = vals
			}
		}
	}
	return nil
}

// IssueStreamPagingResponseData implement urlenc.URLValuesUnmarshaler.
func (m *IssueStreamPagingResponseData) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "total":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Total = val
			}
		}
	}
	return nil
}

// IssueStream implement urlenc.URLValuesUnmarshaler.
func (m *IssueStream) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "issueID":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.IssueID = val
			case "operator":
				m.Operator = vals[0]
			case "streamType":
				m.StreamType = vals[0]
			case "content":
				m.Content = vals[0]
			case "mrInfo":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
			case "mrInfo.appID":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MrInfo.AppID = val
			case "mrInfo.mrID":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MrInfo.MrID = val
			case "mrInfo.mrTitle":
				if m.MrInfo == nil {
					m.MrInfo = &MRCommentInfo{}
				}
				m.MrInfo.MrTitle = vals[0]
			case "createdAt":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
			case "createdAt.seconds":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CreatedAt.Seconds = val
			case "createdAt.nanos":
				if m.CreatedAt == nil {
					m.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.CreatedAt.Nanos = int32(val)
			case "updatedAt":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "updatedAt.seconds":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.UpdatedAt.Seconds = val
			case "updatedAt.nanos":
				if m.UpdatedAt == nil {
					m.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.UpdatedAt.Nanos = int32(val)
			}
		}
	}
	return nil
}