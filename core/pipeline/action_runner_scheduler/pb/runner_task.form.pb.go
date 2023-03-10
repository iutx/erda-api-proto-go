// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: runner_task.proto

package pb

import (
	base64 "encoding/base64"
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*RunnerTask)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RunnerTaskQueryRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RunnerTaskQueryResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RunnerTaskFetchRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RunnerTaskFetchResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RunnerTaskCreateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RunnerTaskCreateResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RunnerTaskUpdateRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*RunnerTaskUpdateResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*LogCollectRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*LogCollectResponse)(nil)

// RunnerTask implement urlenc.URLValuesUnmarshaler.
func (m *RunnerTask) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ID":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ID = val
			case "jobID":
				m.JobID = vals[0]
			case "status":
				m.Status = vals[0]
			case "contextDataUrl":
				m.ContextDataUrl = vals[0]
			case "openApiToken":
				m.OpenApiToken = vals[0]
			case "resultDataUrl":
				m.ResultDataUrl = vals[0]
			case "commands":
				m.Commands = vals
			case "targets":
				m.Targets = vals
			case "workdir":
				m.Workdir = vals[0]
			}
		}
	}
	return nil
}

// RunnerTaskQueryRequest implement urlenc.URLValuesUnmarshaler.
func (m *RunnerTaskQueryRequest) UnmarshalURLValues(prefix string, values url.Values) error {
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

// RunnerTaskQueryResponse implement urlenc.URLValuesUnmarshaler.
func (m *RunnerTaskQueryResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &RunnerTask{}
				}
			case "data.ID":
				if m.Data == nil {
					m.Data = &RunnerTask{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.ID = val
			case "data.jobID":
				if m.Data == nil {
					m.Data = &RunnerTask{}
				}
				m.Data.JobID = vals[0]
			case "data.status":
				if m.Data == nil {
					m.Data = &RunnerTask{}
				}
				m.Data.Status = vals[0]
			case "data.contextDataUrl":
				if m.Data == nil {
					m.Data = &RunnerTask{}
				}
				m.Data.ContextDataUrl = vals[0]
			case "data.openApiToken":
				if m.Data == nil {
					m.Data = &RunnerTask{}
				}
				m.Data.OpenApiToken = vals[0]
			case "data.resultDataUrl":
				if m.Data == nil {
					m.Data = &RunnerTask{}
				}
				m.Data.ResultDataUrl = vals[0]
			case "data.commands":
				if m.Data == nil {
					m.Data = &RunnerTask{}
				}
				m.Data.Commands = vals
			case "data.targets":
				if m.Data == nil {
					m.Data = &RunnerTask{}
				}
				m.Data.Targets = vals
			case "data.workdir":
				if m.Data == nil {
					m.Data = &RunnerTask{}
				}
				m.Data.Workdir = vals[0]
			}
		}
	}
	return nil
}

// RunnerTaskFetchRequest implement urlenc.URLValuesUnmarshaler.
func (m *RunnerTaskFetchRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// RunnerTaskFetchResponse implement urlenc.URLValuesUnmarshaler.
func (m *RunnerTaskFetchResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// RunnerTaskCreateRequest implement urlenc.URLValuesUnmarshaler.
func (m *RunnerTaskCreateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "jobID":
				m.JobID = vals[0]
			case "contextDataUrl":
				m.ContextDataUrl = vals[0]
			case "commands":
				m.Commands = vals
			case "targets":
				m.Targets = vals
			case "workdir":
				m.Workdir = vals[0]
			}
		}
	}
	return nil
}

// RunnerTaskCreateResponse implement urlenc.URLValuesUnmarshaler.
func (m *RunnerTaskCreateResponse) UnmarshalURLValues(prefix string, values url.Values) error {
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

// RunnerTaskUpdateRequest implement urlenc.URLValuesUnmarshaler.
func (m *RunnerTaskUpdateRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			case "taskID":
				m.TaskID = vals[0]
			case "status":
				m.Status = vals[0]
			case "contextDataUrl":
				m.ContextDataUrl = vals[0]
			case "resultDataUrl":
				m.ResultDataUrl = vals[0]
			}
		}
	}
	return nil
}

// RunnerTaskUpdateResponse implement urlenc.URLValuesUnmarshaler.
func (m *RunnerTaskUpdateResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// LogCollectRequest implement urlenc.URLValuesUnmarshaler.
func (m *LogCollectRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "source":
				m.Source = vals[0]
			case "content":
				val, err := base64.StdEncoding.DecodeString(vals[0])
				if err != nil {
					return err
				}
				m.Content = val
			}
		}
	}
	return nil
}

// LogCollectResponse implement urlenc.URLValuesUnmarshaler.
func (m *LogCollectResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
