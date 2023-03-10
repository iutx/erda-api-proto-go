// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: apm_diagnotor.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	pb "github.com/erda-project/erda-proto-go/core/monitor/diagnotor/pb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*ListServicesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListServicesResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ServiceInfo)(nil)
var _ urlenc.URLValuesUnmarshaler = (*InstanceInfo)(nil)
var _ urlenc.URLValuesUnmarshaler = (*StartDiagnosisRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*StartDiagnosisResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryDiagnosisStatusRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*QueryDiagnosisStatusResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*StopDiagnosisRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*StopDiagnosisResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListProcessesRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ListProcessesResponse)(nil)

// ListServicesRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListServicesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "terminusKey":
				m.TerminusKey = vals[0]
			}
		}
	}
	return nil
}

// ListServicesResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListServicesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// ServiceInfo implement urlenc.URLValuesUnmarshaler.
func (m *ServiceInfo) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "orgName":
				m.OrgName = vals[0]
			case "orgID":
				m.OrgID = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			case "projectName":
				m.ProjectName = vals[0]
			case "projectID":
				m.ProjectID = vals[0]
			case "applicationName":
				m.ApplicationName = vals[0]
			case "applicationID":
				m.ApplicationID = vals[0]
			case "service":
				m.Service = vals[0]
			}
		}
	}
	return nil
}

// InstanceInfo implement urlenc.URLValuesUnmarshaler.
func (m *InstanceInfo) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "podName":
				m.PodName = vals[0]
			case "namespace":
				m.Namespace = vals[0]
			case "hostIP":
				m.HostIP = vals[0]
			case "ip":
				m.Ip = vals[0]
			case "runtimeName":
				m.RuntimeName = vals[0]
			case "runtimeID":
				m.RuntimeID = vals[0]
			}
		}
	}
	return nil
}

// StartDiagnosisRequest implement urlenc.URLValuesUnmarshaler.
func (m *StartDiagnosisRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "terminusKey":
				m.TerminusKey = vals[0]
			case "instanceIP":
				m.InstanceIP = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			}
		}
	}
	return nil
}

// StartDiagnosisResponse implement urlenc.URLValuesUnmarshaler.
func (m *StartDiagnosisResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
			case "data.clusterName":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.ClusterName = vals[0]
			case "data.namespace":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.Namespace = vals[0]
			case "data.podName":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.PodName = vals[0]
			case "data.hostIP":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.HostIP = vals[0]
			case "data.podIP":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.PodIP = vals[0]
			case "data.status":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.Status = vals[0]
			case "data.message":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.Message = vals[0]
			}
		}
	}
	return nil
}

// QueryDiagnosisStatusRequest implement urlenc.URLValuesUnmarshaler.
func (m *QueryDiagnosisStatusRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "terminusKey":
				m.TerminusKey = vals[0]
			case "instanceIP":
				m.InstanceIP = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			}
		}
	}
	return nil
}

// QueryDiagnosisStatusResponse implement urlenc.URLValuesUnmarshaler.
func (m *QueryDiagnosisStatusResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
			case "data.clusterName":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.ClusterName = vals[0]
			case "data.namespace":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.Namespace = vals[0]
			case "data.podName":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.PodName = vals[0]
			case "data.hostIP":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.HostIP = vals[0]
			case "data.podIP":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.PodIP = vals[0]
			case "data.status":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.Status = vals[0]
			case "data.message":
				if m.Data == nil {
					m.Data = &pb.DiagnosisInstance{}
				}
				m.Data.Message = vals[0]
			}
		}
	}
	return nil
}

// StopDiagnosisRequest implement urlenc.URLValuesUnmarshaler.
func (m *StopDiagnosisRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "terminusKey":
				m.TerminusKey = vals[0]
			case "instanceIP":
				m.InstanceIP = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			}
		}
	}
	return nil
}

// StopDiagnosisResponse implement urlenc.URLValuesUnmarshaler.
func (m *StopDiagnosisResponse) UnmarshalURLValues(prefix string, values url.Values) error {
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

// ListProcessesRequest implement urlenc.URLValuesUnmarshaler.
func (m *ListProcessesRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "terminusKey":
				m.TerminusKey = vals[0]
			case "instanceIP":
				m.InstanceIP = vals[0]
			case "clusterName":
				m.ClusterName = vals[0]
			}
		}
	}
	return nil
}

// ListProcessesResponse implement urlenc.URLValuesUnmarshaler.
func (m *ListProcessesResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &pb.HostProcessStatus{}
				}
			case "data.totalMemory":
				if m.Data == nil {
					m.Data = &pb.HostProcessStatus{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.TotalMemory = val
			case "data.memoryUsed":
				if m.Data == nil {
					m.Data = &pb.HostProcessStatus{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.MemoryUsed = val
			case "data.memoryUsedPercent":
				if m.Data == nil {
					m.Data = &pb.HostProcessStatus{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.MemoryUsedPercent = val
			case "data.totalCpuCores":
				if m.Data == nil {
					m.Data = &pb.HostProcessStatus{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.TotalCpuCores = val
			case "data.cpuUsedCores":
				if m.Data == nil {
					m.Data = &pb.HostProcessStatus{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.CpuUsedCores = val
			case "data.cpuUsedPercent":
				if m.Data == nil {
					m.Data = &pb.HostProcessStatus{}
				}
				val, err := strconv.ParseFloat(vals[0], 64)
				if err != nil {
					return err
				}
				m.Data.CpuUsedPercent = val
			case "data.connections":
				if m.Data == nil {
					m.Data = &pb.HostProcessStatus{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Data.Connections = val
			}
		}
	}
	return nil
}
