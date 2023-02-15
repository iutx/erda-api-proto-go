// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: graph.proto

package pb

import (
	url "net/url"
	strconv "strconv"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/base/pb"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*PipelineYmlGraphRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineYmlGraphResponse)(nil)

// PipelineYmlGraphRequest implement urlenc.URLValuesUnmarshaler.
func (m *PipelineYmlGraphRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "pipelineYmlContent":
				m.PipelineYmlContent = vals[0]
			case "snippetConfig":
				if m.SnippetConfig == nil {
					m.SnippetConfig = &pb.SnippetConfig{}
				}
			case "snippetConfig.source":
				if m.SnippetConfig == nil {
					m.SnippetConfig = &pb.SnippetConfig{}
				}
				m.SnippetConfig.Source = vals[0]
			case "snippetConfig.name":
				if m.SnippetConfig == nil {
					m.SnippetConfig = &pb.SnippetConfig{}
				}
				m.SnippetConfig.Name = vals[0]
			}
		}
	}
	return nil
}

// PipelineYmlGraphResponse implement urlenc.URLValuesUnmarshaler.
func (m *PipelineYmlGraphResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "data":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
			case "data.version":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				m.Data.Version = vals[0]
			case "data.cron":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				m.Data.Cron = vals[0]
			case "data.cronCompensator":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.CronCompensator == nil {
					m.Data.CronCompensator = &pb.CronCompensator{}
				}
			case "data.cronCompensator.enable":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.CronCompensator == nil {
					m.Data.CronCompensator = &pb.CronCompensator{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.CronCompensator.Enable = val
			case "data.cronCompensator.latestFirst":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.CronCompensator == nil {
					m.Data.CronCompensator = &pb.CronCompensator{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.CronCompensator.LatestFirst = val
			case "data.cronCompensator.stopIfLatterExecuted":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.CronCompensator == nil {
					m.Data.CronCompensator = &pb.CronCompensator{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.CronCompensator.StopIfLatterExecuted = val
			case "data.stages":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.Stages == nil {
					m.Data.Stages = &structpb.ListValue{}
				}
			case "data.needUpgrade":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Data.NeedUpgrade = val
			case "data.ymlContent":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				m.Data.YmlContent = vals[0]
			case "data.on":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.On == nil {
					m.Data.On = &pb.TriggerConfig{}
				}
			case "data.on.push":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.On == nil {
					m.Data.On = &pb.TriggerConfig{}
				}
				if m.Data.On.Push == nil {
					m.Data.On.Push = &pb.PushTrigger{}
				}
			case "data.on.push.branches":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.On == nil {
					m.Data.On = &pb.TriggerConfig{}
				}
				if m.Data.On.Push == nil {
					m.Data.On.Push = &pb.PushTrigger{}
				}
				m.Data.On.Push.Branches = vals
			case "data.on.push.tags":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.On == nil {
					m.Data.On = &pb.TriggerConfig{}
				}
				if m.Data.On.Push == nil {
					m.Data.On.Push = &pb.PushTrigger{}
				}
				m.Data.On.Push.Tags = vals
			case "data.on.merge":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.On == nil {
					m.Data.On = &pb.TriggerConfig{}
				}
				if m.Data.On.Merge == nil {
					m.Data.On.Merge = &pb.MergeTrigger{}
				}
			case "data.on.merge.branches":
				if m.Data == nil {
					m.Data = &pb.PipelineYml{}
				}
				if m.Data.On == nil {
					m.Data.On = &pb.TriggerConfig{}
				}
				if m.Data.On.Merge == nil {
					m.Data.On.Merge = &pb.MergeTrigger{}
				}
				m.Data.On.Merge.Branches = vals
			}
		}
	}
	return nil
}
