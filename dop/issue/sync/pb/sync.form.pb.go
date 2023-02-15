// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: sync.proto

package pb

import (
	json "encoding/json"
	url "net/url"
	strconv "strconv"
	strings "strings"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*IssueSyncRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Fields)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Value)(nil)
var _ urlenc.URLValuesUnmarshaler = (*IssueSyncResponse)(nil)

// IssueSyncRequest implement urlenc.URLValuesUnmarshaler.
func (m *IssueSyncRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			}
		}
	}
	return nil
}

// Fields implement urlenc.URLValuesUnmarshaler.
func (m *Fields) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "updateType":
			case "field":
				m.Field = vals[0]
			case "value":
				if m.Value == nil {
					m.Value = &Value{}
				}
			case "value.content":
				if m.Value == nil {
					m.Value = &Value{}
				}
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
					m.Value.Content = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Value.Content = val
					} else {
						m.Value.Content = structpb.NewStringValue(vals[0])
					}
				}
			case "value.content.null_value":
				if m.Value == nil {
					m.Value = &Value{}
				}
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
					m.Value.Content = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Value.Content = val
					} else {
						m.Value.Content = structpb.NewStringValue(vals[0])
					}
				}
			case "value.content.number_value":
				if m.Value == nil {
					m.Value = &Value{}
				}
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
					m.Value.Content = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Value.Content = val
					} else {
						m.Value.Content = structpb.NewStringValue(vals[0])
					}
				}
			case "value.content.string_value":
				if m.Value == nil {
					m.Value = &Value{}
				}
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
					m.Value.Content = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Value.Content = val
					} else {
						m.Value.Content = structpb.NewStringValue(vals[0])
					}
				}
			case "value.content.bool_value":
				if m.Value == nil {
					m.Value = &Value{}
				}
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
					m.Value.Content = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Value.Content = val
					} else {
						m.Value.Content = structpb.NewStringValue(vals[0])
					}
				}
			case "value.content.struct_value":
				if m.Value == nil {
					m.Value = &Value{}
				}
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
					m.Value.Content = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Value.Content = val
					} else {
						m.Value.Content = structpb.NewStringValue(vals[0])
					}
				}
			case "value.content.list_value":
				if m.Value == nil {
					m.Value = &Value{}
				}
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
					m.Value.Content = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Value.Content = val
					} else {
						m.Value.Content = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// Value implement urlenc.URLValuesUnmarshaler.
func (m *Value) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "content":
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
					m.Content = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Content = val
					} else {
						m.Content = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// IssueSyncResponse implement urlenc.URLValuesUnmarshaler.
func (m *IssueSyncResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
