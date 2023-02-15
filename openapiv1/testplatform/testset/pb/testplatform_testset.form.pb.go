// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: testplatform_testset.proto

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
var _ urlenc.URLValuesUnmarshaler = (*GET_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSet)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetCleanFromRecycleBinRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetCleanFromRecycleBinResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetCopyRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetCopyResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetCreateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetCreateResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetGetResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetListRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetListResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetRecoverFromRecycleBinRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetRecoverFromRecycleBinResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetRecycleRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetRecycleResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSetUpdateRequest)(nil)

// GET_Request implement urlenc.URLValuesUnmarshaler.
func (m *GET_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "testSetID":
				m.TestSetID = vals[0]
			}
		}
	}
	return nil
}

// TestSet implement urlenc.URLValuesUnmarshaler.
func (m *TestSet) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ID = val
			case "name":
				m.Name = vals[0]
			case "projectID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = val
			case "parentID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ParentID = val
			case "recycled":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Recycled = val
			case "directory":
				m.Directory = vals[0]
			case "order":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Order = val
			case "creatorID":
				m.CreatorID = vals[0]
			case "updaterID":
				m.UpdaterID = vals[0]
			}
		}
	}
	return nil
}

// TestSetCleanFromRecycleBinRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestSetCleanFromRecycleBinRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "testSetID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TestSetID = val
			}
		}
	}
	return nil
}

// TestSetCleanFromRecycleBinResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestSetCleanFromRecycleBinResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// TestSetCopyRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestSetCopyRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "copyToTestSetID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.CopyToTestSetID = val
			case "testSetID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TestSetID = val
			}
		}
	}
	return nil
}

// TestSetCopyResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestSetCopyResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data = val
			}
		}
	}
	return nil
}

// TestSetCreateRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestSetCreateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "projectID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = &val
			case "parentID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ParentID = &val
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// TestSetCreateResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestSetCreateResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
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
			}
		}
	}
	return nil
}

// TestSetGetResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestSetGetResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
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
			}
		}
	}
	return nil
}

// TestSetListRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestSetListRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "recycled":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Recycled = val
			case "parentID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ParentID = &val
			case "projectID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectID = &val
			case "testSetIDs":
				list := make([]uint64, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseUint(text, 10, 64)
					if err != nil {
						return err
					}
					list = append(list, val)
				}
				m.TestSetIDs = list
			case "noSubTestSets":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.NoSubTestSets = val
			}
		}
	}
	return nil
}

// TestSetListResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestSetListResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// TestSetRecoverFromRecycleBinRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestSetRecoverFromRecycleBinRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "testSetID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TestSetID = val
			case "recoverToTestSetID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.RecoverToTestSetID = &val
			}
		}
	}
	return nil
}

// TestSetRecoverFromRecycleBinResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestSetRecoverFromRecycleBinResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// TestSetRecycleRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestSetRecycleRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "testSetID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TestSetID = val
			case "isRoot":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsRoot = val
			}
		}
	}
	return nil
}

// TestSetRecycleResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestSetRecycleResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// TestSetUpdateRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestSetUpdateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "testSetID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TestSetID = val
			case "name":
				m.Name = &vals[0]
			case "moveToParentID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.MoveToParentID = &val
			}
		}
	}
	return nil
}
