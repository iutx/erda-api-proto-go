// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: issuerelation.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*IssueRelation)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateIssueRelationRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateIssueRelationResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteIssueRelationRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteIssueRelationResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListIssueRelationRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListIssueRelationResponse)(nil)

// IssueRelation implement urlenc.URLValuesUnmarshaler.
func (m *IssueRelation) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ID":
				m.ID = vals[0]
			case "timeCreated":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
			case "timeCreated.seconds":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeCreated.Seconds = val
			case "timeCreated.nanos":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.TimeCreated.Nanos = int32(val)
			case "timeUpdated":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "timeUpdated.seconds":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeUpdated.Seconds = val
			case "timeUpdated.nanos":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.TimeUpdated.Nanos = int32(val)
			case "softDeletedAt":
				if m.SoftDeletedAt == nil {
					m.SoftDeletedAt = &timestamppb.Timestamp{}
				}
			case "softDeletedAt.seconds":
				if m.SoftDeletedAt == nil {
					m.SoftDeletedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.SoftDeletedAt.Seconds = val
			case "softDeletedAt.nanos":
				if m.SoftDeletedAt == nil {
					m.SoftDeletedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.SoftDeletedAt.Nanos = int32(val)
			case "relation":
				m.Relation = vals[0]
			case "issueID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.IssueID = val
			case "type":
				m.Type = vals[0]
			case "extra":
				m.Extra = vals[0]
			}
		}
	}
	return nil
}

// CreateIssueRelationRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateIssueRelationRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "relation":
				m.Relation = vals[0]
			case "issueID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.IssueID = val
			case "type":
				m.Type = vals[0]
			case "extra":
				m.Extra = vals[0]
			}
		}
	}
	return nil
}

// CreateIssueRelationResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateIssueRelationResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "issueRelation":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
			case "issueRelation.ID":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				m.IssueRelation.ID = vals[0]
			case "issueRelation.timeCreated":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				if m.IssueRelation.TimeCreated == nil {
					m.IssueRelation.TimeCreated = &timestamppb.Timestamp{}
				}
			case "issueRelation.timeCreated.seconds":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				if m.IssueRelation.TimeCreated == nil {
					m.IssueRelation.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.IssueRelation.TimeCreated.Seconds = val
			case "issueRelation.timeCreated.nanos":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				if m.IssueRelation.TimeCreated == nil {
					m.IssueRelation.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.IssueRelation.TimeCreated.Nanos = int32(val)
			case "issueRelation.timeUpdated":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				if m.IssueRelation.TimeUpdated == nil {
					m.IssueRelation.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "issueRelation.timeUpdated.seconds":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				if m.IssueRelation.TimeUpdated == nil {
					m.IssueRelation.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.IssueRelation.TimeUpdated.Seconds = val
			case "issueRelation.timeUpdated.nanos":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				if m.IssueRelation.TimeUpdated == nil {
					m.IssueRelation.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.IssueRelation.TimeUpdated.Nanos = int32(val)
			case "issueRelation.softDeletedAt":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				if m.IssueRelation.SoftDeletedAt == nil {
					m.IssueRelation.SoftDeletedAt = &timestamppb.Timestamp{}
				}
			case "issueRelation.softDeletedAt.seconds":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				if m.IssueRelation.SoftDeletedAt == nil {
					m.IssueRelation.SoftDeletedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.IssueRelation.SoftDeletedAt.Seconds = val
			case "issueRelation.softDeletedAt.nanos":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				if m.IssueRelation.SoftDeletedAt == nil {
					m.IssueRelation.SoftDeletedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.IssueRelation.SoftDeletedAt.Nanos = int32(val)
			case "issueRelation.relation":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				m.IssueRelation.Relation = vals[0]
			case "issueRelation.issueID":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.IssueRelation.IssueID = val
			case "issueRelation.type":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				m.IssueRelation.Type = vals[0]
			case "issueRelation.extra":
				if m.IssueRelation == nil {
					m.IssueRelation = &IssueRelation{}
				}
				m.IssueRelation.Extra = vals[0]
			}
		}
	}
	return nil
}

// DeleteIssueRelationRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteIssueRelationRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "relationID":
				m.RelationID = vals[0]
			}
		}
	}
	return nil
}

// DeleteIssueRelationResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteIssueRelationResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ListIssueRelationRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListIssueRelationRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "relations":
				m.Relations = vals
			case "issueIDs":
				list := make([]uint64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseUint(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.IssueIDs = list
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// ListIssueRelationResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListIssueRelationResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
