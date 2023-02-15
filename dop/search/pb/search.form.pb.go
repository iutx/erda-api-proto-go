// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: search.proto

package pb

import (
	json "encoding/json"
	url "net/url"
	strconv "strconv"
	strings "strings"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	pb "github.com/erda-project/erda-proto-go/common/pb"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*SearchRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*SearchResultItem)(nil)
var _ urlenc.URLValuesUnmarshaler = (*SearchResultContent)(nil)
var _ urlenc.URLValuesUnmarshaler = (*SearchResponse)(nil)

// SearchRequest implement urlenc.URLValuesUnmarshaler.
func (m *SearchRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.OrgID = val
			case "projectID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			case "query":
				m.Query = vals[0]
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
			}
		}
	}
	return nil
}

// SearchResultItem implement urlenc.URLValuesUnmarshaler.
func (m *SearchResultItem) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "item":
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Item = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Item = val
					} else {
						m.Item = structpb.NewStringValue(vals[0])
					}
				}
			case "link":
				m.Link = vals[0]
			}
		}
	}
	return nil
}

// SearchResultContent implement urlenc.URLValuesUnmarshaler.
func (m *SearchResultContent) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// SearchResponse implement urlenc.URLValuesUnmarshaler.
func (m *SearchResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}