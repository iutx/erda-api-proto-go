// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: unittest.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*TestRecordPagingRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestRecordPagingResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestRecordPagingResult)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestRecordGetRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestRecordGetResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestRecord)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestTypeRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestTypeResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestCallBackRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestResult)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestError)(nil)
var _ urlenc.URLValuesUnmarshaler = (*Test)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestTotal)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestSuite)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ToolTip)(nil)
var _ urlenc.URLValuesUnmarshaler = (*CodeCoverageNode)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ReportCounter)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ReportPackage)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ReportClass)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ReportMethod)(nil)
var _ urlenc.URLValuesUnmarshaler = (*TestCallBackResponse)(nil)

// TestRecordPagingRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestRecordPagingRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
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
			case "applicationId":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ApplicationId = val
			}
		}
	}
	return nil
}

// TestRecordPagingResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestRecordPagingResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &TestRecordPagingResult{}
				}
			case "data.total":
				if m.Data == nil {
					m.Data = &TestRecordPagingResult{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Total = val
			}
		}
	}
	return nil
}

// TestRecordPagingResult implement urlenc.URLValuesUnmarshaler.
func (m *TestRecordPagingResult) UnmarshalURLValues(prefix string, values url.Values) error {
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

// TestRecordGetRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestRecordGetRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
			}
		}
	}
	return nil
}

// TestRecordGetResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestRecordGetResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
			case "data.id":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Id = val
			case "data.createdAt":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				if m.Data.CreatedAt == nil {
					m.Data.CreatedAt = &timestamppb.Timestamp{}
				}
			case "data.createdAt.seconds":
				if m.Data == nil {
					m.Data = &TestRecord{}
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
					m.Data = &TestRecord{}
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
					m.Data = &TestRecord{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
			case "data.updatedAt.seconds":
				if m.Data == nil {
					m.Data = &TestRecord{}
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
					m.Data = &TestRecord{}
				}
				if m.Data.UpdatedAt == nil {
					m.Data.UpdatedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Data.UpdatedAt.Nanos = int32(val)
			case "data.applicationId":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.ApplicationId = val
			case "data.projectId":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.ProjectId = val
			case "data.buildId":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.BuildId = val
			case "data.name":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.Name = vals[0]
			case "data.uuid":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.Uuid = vals[0]
			case "data.applicationName":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.ApplicationName = vals[0]
			case "data.output":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.Output = vals[0]
			case "data.desc":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.Desc = vals[0]
			case "data.operatorId":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.OperatorId = vals[0]
			case "data.operatorName":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.OperatorName = vals[0]
			case "data.commitId":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.CommitId = vals[0]
			case "data.branch":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.Branch = vals[0]
			case "data.gitRepo":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.GitRepo = vals[0]
			case "data.caseDir":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.CaseDir = vals[0]
			case "data.application":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.Application = vals[0]
			case "data.type":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.Type = vals[0]
			case "data.totals":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				if m.Data.Totals == nil {
					m.Data.Totals = &TestTotal{}
				}
			case "data.totals.tests":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				if m.Data.Totals == nil {
					m.Data.Totals = &TestTotal{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Totals.Tests = val
			case "data.totals.duration":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				if m.Data.Totals == nil {
					m.Data.Totals = &TestTotal{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Totals.Duration = val
			case "data.parserType":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.ParserType = vals[0]
			case "data.workspace":
				if m.Data == nil {
					m.Data = &TestRecord{}
				}
				m.Data.Workspace = vals[0]
			}
		}
	}
	return nil
}

// TestRecord implement urlenc.URLValuesUnmarshaler.
func (m *TestRecord) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "id":
				val, err := strconv.ParseUint(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Id = val
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
			case "applicationId":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ApplicationId = val
			case "projectId":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectId = val
			case "buildId":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.BuildId = val
			case "name":
				m.Name = vals[0]
			case "uuid":
				m.Uuid = vals[0]
			case "applicationName":
				m.ApplicationName = vals[0]
			case "output":
				m.Output = vals[0]
			case "desc":
				m.Desc = vals[0]
			case "operatorId":
				m.OperatorId = vals[0]
			case "operatorName":
				m.OperatorName = vals[0]
			case "commitId":
				m.CommitId = vals[0]
			case "branch":
				m.Branch = vals[0]
			case "gitRepo":
				m.GitRepo = vals[0]
			case "caseDir":
				m.CaseDir = vals[0]
			case "application":
				m.Application = vals[0]
			case "type":
				m.Type = vals[0]
			case "totals":
				if m.Totals == nil {
					m.Totals = &TestTotal{}
				}
			case "totals.tests":
				if m.Totals == nil {
					m.Totals = &TestTotal{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Totals.Tests = val
			case "totals.duration":
				if m.Totals == nil {
					m.Totals = &TestTotal{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Totals.Duration = val
			case "parserType":
				m.ParserType = vals[0]
			case "workspace":
				m.Workspace = vals[0]
			}
		}
	}
	return nil
}

// TestTypeRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestTypeRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// TestTypeResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestTypeResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				m.Data = vals
			}
		}
	}
	return nil
}

// TestCallBackRequest implement urlenc.URLValuesUnmarshaler.
func (m *TestCallBackRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "results":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
			case "results.applicationId":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Results.ApplicationId = val
			case "results.buildId":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Results.BuildId = val
			case "results.projectId":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Results.ProjectId = val
			case "results.type":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.Type = vals[0]
			case "results.name":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.Name = vals[0]
			case "results.applicationName":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.ApplicationName = vals[0]
			case "results.branch":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.Branch = vals[0]
			case "results.gitRepo":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.GitRepo = vals[0]
			case "results.commitId":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.CommitId = vals[0]
			case "results.operatorName":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.OperatorName = vals[0]
			case "results.operatorId":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.OperatorId = vals[0]
			case "results.status":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.Status = vals[0]
			case "results.workspace":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.Workspace = vals[0]
			case "results.parserType":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.ParserType = vals[0]
			case "results.uuid":
				if m.Results == nil {
					m.Results = &TestResult{}
				}
				m.Results.Uuid = vals[0]
			case "totals":
				if m.Totals == nil {
					m.Totals = &TestTotal{}
				}
			case "totals.tests":
				if m.Totals == nil {
					m.Totals = &TestTotal{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Totals.Tests = val
			case "totals.duration":
				if m.Totals == nil {
					m.Totals = &TestTotal{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Totals.Duration = val
			}
		}
	}
	return nil
}

// TestResult implement urlenc.URLValuesUnmarshaler.
func (m *TestResult) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "applicationId":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ApplicationId = val
			case "buildId":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.BuildId = val
			case "projectId":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.ProjectId = val
			case "type":
				m.Type = vals[0]
			case "name":
				m.Name = vals[0]
			case "applicationName":
				m.ApplicationName = vals[0]
			case "branch":
				m.Branch = vals[0]
			case "gitRepo":
				m.GitRepo = vals[0]
			case "commitId":
				m.CommitId = vals[0]
			case "operatorName":
				m.OperatorName = vals[0]
			case "operatorId":
				m.OperatorId = vals[0]
			case "status":
				m.Status = vals[0]
			case "workspace":
				m.Workspace = vals[0]
			case "parserType":
				m.ParserType = vals[0]
			case "uuid":
				m.Uuid = vals[0]
			}
		}
	}
	return nil
}

// TestError implement urlenc.URLValuesUnmarshaler.
func (m *TestError) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "message":
				m.Message = vals[0]
			case "type":
				m.Type = vals[0]
			case "body":
				m.Body = vals[0]
			}
		}
	}
	return nil
}

// Test implement urlenc.URLValuesUnmarshaler.
func (m *Test) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "classname":
				m.Classname = vals[0]
			case "duration":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Duration = val
			case "status":
				m.Status = vals[0]
			case "error":
				if m.Error == nil {
					m.Error = &TestError{}
				}
			case "error.message":
				if m.Error == nil {
					m.Error = &TestError{}
				}
				m.Error.Message = vals[0]
			case "error.type":
				if m.Error == nil {
					m.Error = &TestError{}
				}
				m.Error.Type = vals[0]
			case "error.body":
				if m.Error == nil {
					m.Error = &TestError{}
				}
				m.Error.Body = vals[0]
			case "stdout":
				m.Stdout = vals[0]
			case "stderr":
				m.Stderr = vals[0]
			}
		}
	}
	return nil
}

// TestTotal implement urlenc.URLValuesUnmarshaler.
func (m *TestTotal) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "tests":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Tests = val
			case "duration":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Duration = val
			}
		}
	}
	return nil
}

// TestSuite implement urlenc.URLValuesUnmarshaler.
func (m *TestSuite) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "package":
				m.Package = vals[0]
			case "stdout":
				m.Stdout = vals[0]
			case "stderr":
				m.Stderr = vals[0]
			case "totals":
				if m.Totals == nil {
					m.Totals = &TestTotal{}
				}
			case "totals.tests":
				if m.Totals == nil {
					m.Totals = &TestTotal{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Totals.Tests = val
			case "totals.duration":
				if m.Totals == nil {
					m.Totals = &TestTotal{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Totals.Duration = val
			}
		}
	}
	return nil
}

// ToolTip implement urlenc.URLValuesUnmarshaler.
func (m *ToolTip) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "formatter":
				m.Formatter = vals[0]
			}
		}
	}
	return nil
}

// CodeCoverageNode implement urlenc.URLValuesUnmarshaler.
func (m *CodeCoverageNode) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "value":
				list := make([]float32, 0, len(vals))
				for _, text := range vals {
					val, err := strconv.ParseFloat(text, 32)
					if err != nil {
						return err
					}
					list = append(list, float32(val))
				}
				m.Value = list
			case "name":
				m.Name = vals[0]
			case "path":
				m.Path = vals[0]
			case "tooltip":
				if m.Tooltip == nil {
					m.Tooltip = &ToolTip{}
				}
			case "tooltip.formatter":
				if m.Tooltip == nil {
					m.Tooltip = &ToolTip{}
				}
				m.Tooltip.Formatter = vals[0]
			}
		}
	}
	return nil
}

// ReportCounter implement urlenc.URLValuesUnmarshaler.
func (m *ReportCounter) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "covered":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Covered = val
			case "missed":
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Missed = val
			case "type":
				m.Type = vals[0]
			}
		}
	}
	return nil
}

// ReportPackage implement urlenc.URLValuesUnmarshaler.
func (m *ReportPackage) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			}
		}
	}
	return nil
}

// ReportClass implement urlenc.URLValuesUnmarshaler.
func (m *ReportClass) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "sourcefilename":
				m.Sourcefilename = vals[0]
			}
		}
	}
	return nil
}

// ReportMethod implement urlenc.URLValuesUnmarshaler.
func (m *ReportMethod) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "desc":
				m.Desc = vals[0]
			case "line":
				m.Line = vals[0]
			}
		}
	}
	return nil
}

// TestCallBackResponse implement urlenc.URLValuesUnmarshaler.
func (m *TestCallBackResponse) UnmarshalURLValues(prefix string, values url.Values) error {
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
