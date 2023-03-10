// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: eventbox.proto

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
var _ urlenc.URLValuesUnmarshaler = (*CreateMessageRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateMessageResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PrefixGetRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PrefixGetResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PrefixValue)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PutRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PutResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DelRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DelResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetVersionRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetVersionResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetSMTPInfoRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*GetSMTPInfoResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*MailSubscriberInfo)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListHooksRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListHooksResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*InspectHookRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*InspectHookResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateHookRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CreateHookResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EditHookRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EditHookResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PingHookRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PingHookResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteHookRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DeleteHookResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListHookEventsRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListHookEventsResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*HookEvents)(nil)
var _ urlenc.URLValuesUnmarshaler = (*StatRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*StatResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Hook)(nil)

// CreateMessageRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateMessageRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "sender":
				m.Sender = vals[0]
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
			case "time":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Time = val
			case "originContent":
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
					m.OriginContent = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.OriginContent = val
					} else {
						m.OriginContent = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// CreateMessageResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateMessageResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// PrefixGetRequest implement urlenc.URLValuesUnmarshaler.
func (m *PrefixGetRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			}
		}
	}
	return nil
}

// PrefixGetResponse implement urlenc.URLValuesUnmarshaler.
func (m *PrefixGetResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// PrefixValue implement urlenc.URLValuesUnmarshaler.
func (m *PrefixValue) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// PutRequest implement urlenc.URLValuesUnmarshaler.
func (m *PutRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			}
		}
	}
	return nil
}

// PutResponse implement urlenc.URLValuesUnmarshaler.
func (m *PutResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// DelRequest implement urlenc.URLValuesUnmarshaler.
func (m *DelRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			}
		}
	}
	return nil
}

// DelResponse implement urlenc.URLValuesUnmarshaler.
func (m *DelResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// GetVersionRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetVersionRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetVersionResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetVersionResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// GetSMTPInfoRequest implement urlenc.URLValuesUnmarshaler.
func (m *GetSMTPInfoRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// GetSMTPInfoResponse implement urlenc.URLValuesUnmarshaler.
func (m *GetSMTPInfoResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &MailSubscriberInfo{}
				}
			case "data.host":
				if m.Data == nil {
					m.Data = &MailSubscriberInfo{}
				}
				m.Data.Host = vals[0]
			case "data.port":
				if m.Data == nil {
					m.Data = &MailSubscriberInfo{}
				}
				m.Data.Port = vals[0]
			case "data.user":
				if m.Data == nil {
					m.Data = &MailSubscriberInfo{}
				}
				m.Data.User = vals[0]
			case "data.password":
				if m.Data == nil {
					m.Data = &MailSubscriberInfo{}
				}
				m.Data.Password = vals[0]
			case "data.displayUser":
				if m.Data == nil {
					m.Data = &MailSubscriberInfo{}
				}
				m.Data.DisplayUser = vals[0]
			case "data.isSSL":
				if m.Data == nil {
					m.Data = &MailSubscriberInfo{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.IsSSL = val
			case "data.insecureSkipVerify":
				if m.Data == nil {
					m.Data = &MailSubscriberInfo{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.InsecureSkipVerify = val
			}
		}
	}
	return nil
}

// MailSubscriberInfo implement urlenc.URLValuesUnmarshaler.
func (m *MailSubscriberInfo) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "host":
				m.Host = vals[0]
			case "port":
				m.Port = vals[0]
			case "user":
				m.User = vals[0]
			case "password":
				m.Password = vals[0]
			case "displayUser":
				m.DisplayUser = vals[0]
			case "isSSL":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsSSL = val
			case "insecureSkipVerify":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.InsecureSkipVerify = val
			}
		}
	}
	return nil
}

// ListHooksRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListHooksRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgId":
				m.OrgId = vals[0]
			case "projectId":
				m.ProjectId = vals[0]
			case "applicationId":
				m.ApplicationId = vals[0]
			case "env":
				m.Env = vals[0]
			}
		}
	}
	return nil
}

// ListHooksResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListHooksResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// InspectHookRequest implement urlenc.URLValuesUnmarshaler.
func (m *InspectHookRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgId":
				m.OrgId = vals[0]
			case "projectId":
				m.ProjectId = vals[0]
			case "applicationId":
				m.ApplicationId = vals[0]
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// InspectHookResponse implement urlenc.URLValuesUnmarshaler.
func (m *InspectHookResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &Hook{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.Id = vals[0]
			case "data.updatedAt":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.UpdatedAt = vals[0]
			case "data.createdAt":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.CreatedAt = vals[0]
			case "data.secret":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.Secret = vals[0]
			case "data.name":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.Name = vals[0]
			case "data.events":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.Events = vals
			case "data.url":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.Url = vals[0]
			case "data.active":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.Active = val
			case "data.orgID":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.OrgID = vals[0]
			case "data.projectID":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.ProjectID = vals[0]
			case "data.applicationID":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.ApplicationID = vals[0]
			case "data.env":
				if m.Data == nil {
					m.Data = &Hook{}
				}
				m.Data.Env = vals
			}
		}
	}
	return nil
}

// CreateHookRequest implement urlenc.URLValuesUnmarshaler.
func (m *CreateHookRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "events":
				m.Events = vals
			case "url":
				m.Url = vals[0]
			case "active":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Active = val
			case "org":
				m.Org = vals[0]
			case "project":
				m.Project = vals[0]
			case "application":
				m.Application = vals[0]
			case "env":
				m.Env = vals
			}
		}
	}
	return nil
}

// CreateHookResponse implement urlenc.URLValuesUnmarshaler.
func (m *CreateHookResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// EditHookRequest implement urlenc.URLValuesUnmarshaler.
func (m *EditHookRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "events":
				m.Events = vals
			case "removeEvents":
				m.RemoveEvents = vals
			case "addEvents":
				m.AddEvents = vals
			case "url":
				m.Url = vals[0]
			case "active":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Active = val
			case "id":
				m.Id = vals[0]
			case "orgId":
				m.OrgId = vals[0]
			case "projectId":
				m.ProjectId = vals[0]
			case "applicationId":
				m.ApplicationId = vals[0]
			}
		}
	}
	return nil
}

// EditHookResponse implement urlenc.URLValuesUnmarshaler.
func (m *EditHookResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// PingHookRequest implement urlenc.URLValuesUnmarshaler.
func (m *PingHookRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgId":
				m.OrgId = vals[0]
			case "projectId":
				m.ProjectId = vals[0]
			case "applicationId":
				m.ApplicationId = vals[0]
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// PingHookResponse implement urlenc.URLValuesUnmarshaler.
func (m *PingHookResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// DeleteHookRequest implement urlenc.URLValuesUnmarshaler.
func (m *DeleteHookRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgId":
				m.OrgId = vals[0]
			case "projectId":
				m.ProjectId = vals[0]
			case "applicationId":
				m.ApplicationId = vals[0]
			case "id":
				m.Id = vals[0]
			}
		}
	}
	return nil
}

// DeleteHookResponse implement urlenc.URLValuesUnmarshaler.
func (m *DeleteHookResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// ListHookEventsRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListHookEventsRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgId":
				m.OrgId = vals[0]
			case "projectId":
				m.ProjectId = vals[0]
			case "applicationId":
				m.ApplicationId = vals[0]
			}
		}
	}
	return nil
}

// ListHookEventsResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListHookEventsResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// HookEvents implement urlenc.URLValuesUnmarshaler.
func (m *HookEvents) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "key":
				m.Key = vals[0]
			case "title":
				m.Title = vals[0]
			case "desc":
				m.Desc = vals[0]
			}
		}
	}
	return nil
}

// StatRequest implement urlenc.URLValuesUnmarshaler.
func (m *StatRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// StatResponse implement urlenc.URLValuesUnmarshaler.
func (m *StatResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals[0]
			}
		}
	}
	return nil
}

// Hook implement urlenc.URLValuesUnmarshaler.
func (m *Hook) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				m.Id = vals[0]
			case "updatedAt":
				m.UpdatedAt = vals[0]
			case "createdAt":
				m.CreatedAt = vals[0]
			case "secret":
				m.Secret = vals[0]
			case "name":
				m.Name = vals[0]
			case "events":
				m.Events = vals
			case "url":
				m.Url = vals[0]
			case "active":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Active = val
			case "orgID":
				m.OrgID = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "applicationID":
				m.ApplicationID = vals[0]
			case "env":
				m.Env = vals
			}
		}
	}
	return nil
}
