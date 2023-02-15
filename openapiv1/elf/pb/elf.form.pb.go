// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: elf.proto

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
var _ urlenc.URLValuesUnmarshaler = (*DependencyPackageList)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DependencyPackageListResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DependencyPackageSpec)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DependencyPackageType)(nil)
var _ urlenc.URLValuesUnmarshaler = (*DependencyPackageTypeItem)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ELF_ENVIROMENT_DELETE_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ELF_ENVIROMENT_GET_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ELF_ENVIROMENT_LIST_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ELF_NOTEBOOK_DELETE_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ELF_NOTEBOOK_GET_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ELF_NOTEBOOK_LIST_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ELF_PACKAGE_LIST_Request)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ElfMetadata)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ElfResource)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Environment)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EnvironmentList)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EnvironmentListResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EnvironmentListSpec)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EnvironmentResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*EnvironmentSpec)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListMetadata)(nil)
var _ urlenc.URLValuesUnmarshaler = (*MetadataField)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NoteBookList)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NoteBookListSpec)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Notebook)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotebookListResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotebookResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotebookSpec)(nil)
var _ urlenc.URLValuesUnmarshaler = (*NotebookStatus)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Package)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Require)(nil)

// DependencyPackageList implement urlenc.URLValuesUnmarshaler.
func (m *DependencyPackageList) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "listMetadata":
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
					m.ListMetadata = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.ListMetadata = val
					} else {
						m.ListMetadata = structpb.NewStringValue(vals[0])
					}
				}
			case "dependencyPackageSpec":
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
					m.DependencyPackageSpec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.DependencyPackageSpec = val
					} else {
						m.DependencyPackageSpec = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// DependencyPackageListResponse implement urlenc.URLValuesUnmarshaler.
func (m *DependencyPackageListResponse) UnmarshalURLValues(prefix string, values url.Values) error {
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

// DependencyPackageSpec implement urlenc.URLValuesUnmarshaler.
func (m *DependencyPackageSpec) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// DependencyPackageType implement urlenc.URLValuesUnmarshaler.
func (m *DependencyPackageType) UnmarshalURLValues(prefix string, values url.Values) error {
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

// DependencyPackageTypeItem implement urlenc.URLValuesUnmarshaler.
func (m *DependencyPackageTypeItem) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "version":
				m.Version = vals
			}
		}
	}
	return nil
}

// ELF_ENVIROMENT_DELETE_Request implement urlenc.URLValuesUnmarshaler.
func (m *ELF_ENVIROMENT_DELETE_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ELF_ENVIROMENT_GET_Request implement urlenc.URLValuesUnmarshaler.
func (m *ELF_ENVIROMENT_GET_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ELF_ENVIROMENT_LIST_Request implement urlenc.URLValuesUnmarshaler.
func (m *ELF_ENVIROMENT_LIST_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ELF_NOTEBOOK_DELETE_Request implement urlenc.URLValuesUnmarshaler.
func (m *ELF_NOTEBOOK_DELETE_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ELF_NOTEBOOK_GET_Request implement urlenc.URLValuesUnmarshaler.
func (m *ELF_NOTEBOOK_GET_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ELF_NOTEBOOK_LIST_Request implement urlenc.URLValuesUnmarshaler.
func (m *ELF_NOTEBOOK_LIST_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ELF_PACKAGE_LIST_Request implement urlenc.URLValuesUnmarshaler.
func (m *ELF_PACKAGE_LIST_Request) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ElfMetadata implement urlenc.URLValuesUnmarshaler.
func (m *ElfMetadata) UnmarshalURLValues(prefix string, values url.Values) error {
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
			case "description":
				m.Description = vals[0]
			case "workspace":
				m.Workspace = vals[0]
			case "ownerName":
				m.OwnerName = vals[0]
			case "ownerID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.OwnerID = val
			case "organizationID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.OrganizationID = val
			case "organizationName":
				m.OrganizationName = vals[0]
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

// ElfResource implement urlenc.URLValuesUnmarshaler.
func (m *ElfResource) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "cPU":
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.CPU = val
			case "memory":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Memory = val
			}
		}
	}
	return nil
}

// Environment implement urlenc.URLValuesUnmarshaler.
func (m *Environment) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "elfMetadata":
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
					m.ElfMetadata = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.ElfMetadata = val
					} else {
						m.ElfMetadata = structpb.NewStringValue(vals[0])
					}
				}
			case "environmentSpec":
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
					m.EnvironmentSpec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.EnvironmentSpec = val
					} else {
						m.EnvironmentSpec = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// EnvironmentList implement urlenc.URLValuesUnmarshaler.
func (m *EnvironmentList) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "listMetadata":
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
					m.ListMetadata = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.ListMetadata = val
					} else {
						m.ListMetadata = structpb.NewStringValue(vals[0])
					}
				}
			case "environmentListSpec":
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
					m.EnvironmentListSpec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.EnvironmentListSpec = val
					} else {
						m.EnvironmentListSpec = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// EnvironmentListResponse implement urlenc.URLValuesUnmarshaler.
func (m *EnvironmentListResponse) UnmarshalURLValues(prefix string, values url.Values) error {
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

// EnvironmentListSpec implement urlenc.URLValuesUnmarshaler.
func (m *EnvironmentListSpec) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// EnvironmentResponse implement urlenc.URLValuesUnmarshaler.
func (m *EnvironmentResponse) UnmarshalURLValues(prefix string, values url.Values) error {
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

// EnvironmentSpec implement urlenc.URLValuesUnmarshaler.
func (m *EnvironmentSpec) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "dBRequires":
				m.DBRequires = vals[0]
			case "dBLabels":
				m.DBLabels = vals[0]
			case "notebookCount":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.NotebookCount = val
			}
		}
	}
	return nil
}

// ListMetadata implement urlenc.URLValuesUnmarshaler.
func (m *ListMetadata) UnmarshalURLValues(prefix string, values url.Values) error {
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

// MetadataField implement urlenc.URLValuesUnmarshaler.
func (m *MetadataField) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "value":
				m.Value = vals[0]
			case "type":
				m.Type = vals[0]
			case "optional":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Optional = val
			case "level":
				m.Level = vals[0]
			}
		}
	}
	return nil
}

// NoteBookList implement urlenc.URLValuesUnmarshaler.
func (m *NoteBookList) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "listMetadata":
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
					m.ListMetadata = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.ListMetadata = val
					} else {
						m.ListMetadata = structpb.NewStringValue(vals[0])
					}
				}
			case "noteBookListSpec":
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
					m.NoteBookListSpec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.NoteBookListSpec = val
					} else {
						m.NoteBookListSpec = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// NoteBookListSpec implement urlenc.URLValuesUnmarshaler.
func (m *NoteBookListSpec) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// Notebook implement urlenc.URLValuesUnmarshaler.
func (m *Notebook) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "notebookSpec":
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
					m.NotebookSpec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.NotebookSpec = val
					} else {
						m.NotebookSpec = structpb.NewStringValue(vals[0])
					}
				}
			case "notebookStatus":
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
					m.NotebookStatus = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.NotebookStatus = val
					} else {
						m.NotebookStatus = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// NotebookListResponse implement urlenc.URLValuesUnmarshaler.
func (m *NotebookListResponse) UnmarshalURLValues(prefix string, values url.Values) error {
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

// NotebookResponse implement urlenc.URLValuesUnmarshaler.
func (m *NotebookResponse) UnmarshalURLValues(prefix string, values url.Values) error {
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

// NotebookSpec implement urlenc.URLValuesUnmarshaler.
func (m *NotebookSpec) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "clusterName":
				m.ClusterName = vals[0]
			case "projectName":
				m.ProjectName = vals[0]
			case "applicationName":
				m.ApplicationName = vals[0]
			case "dBEnvs":
				m.DBEnvs = vals[0]
			case "image":
				m.Image = vals[0]
			case "requirementEnvID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.RequirementEnvID = val
			case "dataSourceID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.DataSourceID = val
			case "genericDomain":
				m.GenericDomain = vals[0]
			case "clusterDomain":
				m.ClusterDomain = vals[0]
			case "elfResource":
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
					m.ElfResource = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.ElfResource = val
					} else {
						m.ElfResource = structpb.NewStringValue(vals[0])
					}
				}
			}
		}
	}
	return nil
}

// NotebookStatus implement urlenc.URLValuesUnmarshaler.
func (m *NotebookStatus) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "startedAt":
				if m.StartedAt == nil {
					m.StartedAt = &timestamppb.Timestamp{}
				}
			case "startedAt.seconds":
				if m.StartedAt == nil {
					m.StartedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.StartedAt.Seconds = val
			case "startedAt.nanos":
				if m.StartedAt == nil {
					m.StartedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.StartedAt.Nanos = int32(val)
			case "state":
				m.State = vals[0]
			}
		}
	}
	return nil
}

// Package implement urlenc.URLValuesUnmarshaler.
func (m *Package) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "version":
				m.Version = vals[0]
			}
		}
	}
	return nil
}

// Require implement urlenc.URLValuesUnmarshaler.
func (m *Require) UnmarshalURLValues(prefix string, values url.Values) error {
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
