// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: token.proto

package pb

import (
	json "encoding/json"
	url "net/url"
	strconv "strconv"
	strings "strings"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*GetTokenRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetTokenResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Token)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryTokensRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryTokensResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateTokenRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateTokenResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateTokenRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateTokenResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteTokenRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteTokenResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ScopeEnum)(nil)

// GetTokenRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetTokenRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// GetTokenResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetTokenResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Token{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.Id = vals[0]
			case "data.secretKey":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.SecretKey = vals[0]
			case "data.accessKey":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.AccessKey = vals[0]
			case "data.data":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.null_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.number_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.string_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.bool_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.struct_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.list_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.status":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.Status = vals[0]
			case "data.description":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.Description = vals[0]
			case "data.scope":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.Scope = vals[0]
			case "data.scopeId":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.ScopeId = vals[0]
			case "data.createdAt":
				if m.Data == nil {
					m.Data = &Token{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
			case "data.createdAt.seconds":
				if m.Data == nil {
					m.Data = &Token{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Seconds = val
			case "data.createdAt.nanos":
				if m.Data == nil {
					m.Data = &Token{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Nanos = int32(val)
			case "data.creatorId":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.CreatorId = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.Type = vals[0]
			case "data.expiresIn":
				if m.Data == nil {
					m.Data = &Token{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.ExpiresIn = val
			}
		}
	}
	return nil
}

// Token implement urlenc.URLValuesUnmarshaler.
func (m *Token) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "secretKey":
				m.SecretKey = vals[0]
			case "accessKey":
				m.AccessKey = vals[0]
			case "data":
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
					m.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data = val
					} else {
						m.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "status":
				m.Status = vals[0]
			case "description":
				m.Description = vals[0]
			case "scope":
				m.Scope = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
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
			case "creatorId":
				m.CreatorId = vals[0]
			case "type":
				m.Type = vals[0]
			case "expiresIn":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ExpiresIn = val
			}
		}
	}
	return nil
}

// QueryTokensRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueryTokensRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
			case "type":
				m.Type = vals[0]
			case "access":
				m.Access = vals[0]
			case "status":
				m.Status = vals[0]
			case "pageNo":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageNo = val
			case "pageSize":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.PageSize = val
			case "creatorId":
				m.CreatorId = vals[0]
			}
		}
	}
	return nil
}

// QueryTokensResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueryTokensResponse) UnmarshalURLValues(prefix string, values url.Values) error {
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

// CreateTokenRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateTokenRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "type":
				m.Type = vals[0]
			case "description":
				m.Description = vals[0]
			case "scope":
				m.Scope = vals[0]
			case "scopeId":
				m.ScopeId = vals[0]
			case "creatorId":
				m.CreatorId = vals[0]
			}
		}
	}
	return nil
}

// CreateTokenResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateTokenResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Token{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.Id = vals[0]
			case "data.secretKey":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.SecretKey = vals[0]
			case "data.accessKey":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.AccessKey = vals[0]
			case "data.data":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.null_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.number_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.string_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.bool_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.struct_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.data.list_value":
				if m.Data == nil {
					m.Data = &Token{}
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
					m.Data.Data = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Data.Data = val
					} else {
						m.Data.Data = structpb.NewStringValue(vals[0])
					}
				}
			case "data.status":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.Status = vals[0]
			case "data.description":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.Description = vals[0]
			case "data.scope":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.Scope = vals[0]
			case "data.scopeId":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.ScopeId = vals[0]
			case "data.createdAt":
				if m.Data == nil {
					m.Data = &Token{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
			case "data.createdAt.seconds":
				if m.Data == nil {
					m.Data = &Token{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Seconds = val
			case "data.createdAt.nanos":
				if m.Data == nil {
					m.Data = &Token{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Nanos = int32(val)
			case "data.creatorId":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.CreatorId = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &Token{}
				}
				m.Data.Type = vals[0]
			case "data.expiresIn":
				if m.Data == nil {
					m.Data = &Token{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.ExpiresIn = val
			}
		}
	}
	return nil
}

// UpdateTokenRequest implement urlenc.URLValuesUnmarshaler.
func (m *UpdateTokenRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "status":
				m.Status = vals[0]
			case "description":
				m.Description = vals[0]
			}
		}
	}
	return nil
}

// UpdateTokenResponse implement urlenc.URLValuesUnmarshaler.
func (m *UpdateTokenResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// DeleteTokenRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteTokenRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// DeleteTokenResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteTokenResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ScopeEnum implement urlenc.URLValuesUnmarshaler.
func (m *ScopeEnum) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
