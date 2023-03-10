// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: rule.proto

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
var _ urlenc.URLValuesUnmarshaler = (*FireRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Config)(nil)
var _ urlenc.URLValuesUnmarshaler = (*FireResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Rule)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ActionParams)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ActionNode)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DingTalkConfig)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*UpdateRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListRulesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListRulesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListRulesResponseData)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteRuleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteRuleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListRuleExecHistoryRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListRuleExecHistoryResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListRuleExecHistoryResponseData)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RuleExecHistory)(nil)

// FireRequest implement urlenc.URLValuesUnmarshaler.
func (m *FireRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "eventType":
				m.EventType = vals[0]
			}
		}
	}
	return nil
}

// Config implement urlenc.URLValuesUnmarshaler.
func (m *Config) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "code":
				m.Code = &vals[0]
			}
		}
	}
	return nil
}

// FireResponse implement urlenc.URLValuesUnmarshaler.
func (m *FireResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "output":
				list := make([]bool, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseBool(text)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.Output = list
			}
		}
	}
	return nil
}

// CreateRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "eventType":
				m.EventType = vals[0]
			case "code":
				m.Code = &vals[0]
			case "name":
				m.Name = vals[0]
			case "params":
				if m.Params == nil {
					m.Params = &ActionParams{}
				}
			case "enabled":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Enabled = val
			}
		}
	}
	return nil
}

// CreateRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
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

// GetRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Rule{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				m.Data.Id = vals[0]
			case "data.name":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				m.Data.Name = vals[0]
			case "data.scope":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				m.Data.Scope = vals[0]
			case "data.scopeID":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				m.Data.ScopeID = vals[0]
			case "data.eventType":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				m.Data.EventType = vals[0]
			case "data.code":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				m.Data.Code = vals[0]
			case "data.params":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				if m.Data.Params == nil {
					m.Data.Params = &ActionParams{}
				}
			case "data.enabled":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.Enabled = val
			case "data.actor":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				m.Data.Actor = vals[0]
			case "data.createdAt":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
			case "data.createdAt.seconds":
				if m.Data == nil {
					m.Data = &Rule{}
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
					m.Data = &Rule{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.CreatedAt.Nanos = int32(val)
			case "data.updatedAt":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "data.updatedAt.seconds":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.UpdatedAt.Seconds = val
			case "data.updatedAt.nanos":
				if m.Data == nil {
					m.Data = &Rule{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.UpdatedAt.Nanos = int32(val)
			case "userIDs":
				m.UserIDs = vals
			}
		}
	}
	return nil
}

// Rule implement urlenc.URLValuesUnmarshaler.
func (m *Rule) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "name":
				m.Name = vals[0]
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "eventType":
				m.EventType = vals[0]
			case "code":
				m.Code = vals[0]
			case "params":
				if m.Params == nil {
					m.Params = &ActionParams{}
				}
			case "enabled":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Enabled = val
			case "actor":
				m.Actor = vals[0]
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

// ActionParams implement urlenc.URLValuesUnmarshaler.
func (m *ActionParams) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ActionNode implement urlenc.URLValuesUnmarshaler.
func (m *ActionNode) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "dingTalk":
				if m.DingTalk == nil {
					m.DingTalk = &DingTalkConfig{}
				}
			case "dingTalk.webhook":
				if m.DingTalk == nil {
					m.DingTalk = &DingTalkConfig{}
				}
				m.DingTalk.Webhook = vals[0]
			case "dingTalk.signature":
				if m.DingTalk == nil {
					m.DingTalk = &DingTalkConfig{}
				}
				m.DingTalk.Signature = vals[0]
			case "snippet":
				m.Snippet = vals[0]
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// DingTalkConfig implement urlenc.URLValuesUnmarshaler.
func (m *DingTalkConfig) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "webhook":
				m.Webhook = vals[0]
			case "signature":
				m.Signature = vals[0]
			}
		}
	}
	return nil
}

// UpdateRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *UpdateRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "code":
				m.Code = vals[0]
			case "eventType":
				m.EventType = vals[0]
			case "enabled":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Enabled = &val
			case "name":
				m.Name = vals[0]
			case "params":
				if m.Params == nil {
					m.Params = &ActionParams{}
				}
			case "actor":
				m.Actor = vals[0]
			}
		}
	}
	return nil
}

// UpdateRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *UpdateRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ListRulesRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListRulesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "eventType":
				m.EventType = vals[0]
			case "enabled":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Enabled = &val
			case "name":
				m.Name = vals[0]
			case "actor":
				m.Actor = vals[0]
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
			}
		}
	}
	return nil
}

// ListRulesResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListRulesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &ListRulesResponseData{}
				}
			case "data.total":
				if m.Data == nil {
					m.Data = &ListRulesResponseData{}
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

// ListRulesResponseData implement urlenc.URLValuesUnmarshaler.
func (m *ListRulesResponseData) UnmarshalURLValues(prefix string, values url.Values) error {
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

// DeleteRuleRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteRuleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
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

// DeleteRuleResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteRuleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ListRuleExecHistoryRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListRuleExecHistoryRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "eventType":
				m.EventType = vals[0]
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
			case "ruleID":
				m.RuleID = vals[0]
			case "succeed":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Succeed = &val
			}
		}
	}
	return nil
}

// ListRuleExecHistoryResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListRuleExecHistoryResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &ListRuleExecHistoryResponseData{}
				}
			case "data.total":
				if m.Data == nil {
					m.Data = &ListRuleExecHistoryResponseData{}
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

// ListRuleExecHistoryResponseData implement urlenc.URLValuesUnmarshaler.
func (m *ListRuleExecHistoryResponseData) UnmarshalURLValues(prefix string, values url.Values) error {
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

// RuleExecHistory implement urlenc.URLValuesUnmarshaler.
func (m *RuleExecHistory) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
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
			case "scope":
				m.Scope = vals[0]
			case "scopeID":
				m.ScopeID = vals[0]
			case "ruleID":
				m.RuleID = vals[0]
			case "code":
				m.Code = vals[0]
			case "env":
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
					m.Env = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Env = val
					} else {
						m.Env = structpb.NewStringValue(vals[0])
					}
				}
			case "succeed":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Succeed = val
			case "actionOutput":
				m.ActionOutput = vals[0]
			case "actor":
				m.Actor = vals[0]
			}
		}
	}
	return nil
}
