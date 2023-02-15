// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: fdp.proto

package pb

import (
	url "net/url"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*CDP_DELETE_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CDP_GET_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CDP_PATCH_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CDP_POST_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CDP_PUT_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*FDP_WEBSOCKET_Request)(nil)

// CDP_DELETE_Request implement urlenc.URLValuesUnmarshaler.
func (m *CDP_DELETE_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CDP_GET_Request implement urlenc.URLValuesUnmarshaler.
func (m *CDP_GET_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CDP_PATCH_Request implement urlenc.URLValuesUnmarshaler.
func (m *CDP_PATCH_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CDP_POST_Request implement urlenc.URLValuesUnmarshaler.
func (m *CDP_POST_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// CDP_PUT_Request implement urlenc.URLValuesUnmarshaler.
func (m *CDP_PUT_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// FDP_WEBSOCKET_Request implement urlenc.URLValuesUnmarshaler.
func (m *FDP_WEBSOCKET_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}