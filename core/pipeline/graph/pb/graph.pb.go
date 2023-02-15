// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: graph.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/erda-project/erda-proto-go/common/pb"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/base/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PipelineYmlGraphRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineYmlContent        string            `protobuf:"bytes,1,opt,name=pipelineYmlContent,proto3" json:"pipelineYmlContent,omitempty"`
	GlobalSnippetConfigLabels map[string]string `protobuf:"bytes,2,rep,name=globalSnippetConfigLabels,proto3" json:"globalSnippetConfigLabels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SnippetConfig             *pb.SnippetConfig `protobuf:"bytes,3,opt,name=snippetConfig,proto3" json:"snippetConfig,omitempty"`
}

func (x *PipelineYmlGraphRequest) Reset() {
	*x = PipelineYmlGraphRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineYmlGraphRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineYmlGraphRequest) ProtoMessage() {}

func (x *PipelineYmlGraphRequest) ProtoReflect() protoreflect.Message {
	mi := &file_graph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineYmlGraphRequest.ProtoReflect.Descriptor instead.
func (*PipelineYmlGraphRequest) Descriptor() ([]byte, []int) {
	return file_graph_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineYmlGraphRequest) GetPipelineYmlContent() string {
	if x != nil {
		return x.PipelineYmlContent
	}
	return ""
}

func (x *PipelineYmlGraphRequest) GetGlobalSnippetConfigLabels() map[string]string {
	if x != nil {
		return x.GlobalSnippetConfigLabels
	}
	return nil
}

func (x *PipelineYmlGraphRequest) GetSnippetConfig() *pb.SnippetConfig {
	if x != nil {
		return x.SnippetConfig
	}
	return nil
}

type PipelineYmlGraphResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *pb.PipelineYml `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PipelineYmlGraphResponse) Reset() {
	*x = PipelineYmlGraphResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineYmlGraphResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineYmlGraphResponse) ProtoMessage() {}

func (x *PipelineYmlGraphResponse) ProtoReflect() protoreflect.Message {
	mi := &file_graph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineYmlGraphResponse.ProtoReflect.Descriptor instead.
func (*PipelineYmlGraphResponse) Descriptor() ([]byte, []int) {
	return file_graph_proto_rawDescGZIP(), []int{1}
}

func (x *PipelineYmlGraphResponse) GetData() *pb.PipelineYml {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_graph_proto protoreflect.FileDescriptor

var file_graph_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65,
	0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x1a, 0x1d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x17, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d,
	0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x12, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x8e, 0x01,
	0x0a, 0x19, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x50, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x19, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4c,
	0x0a, 0x0d, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x73,
	0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x4c, 0x0a, 0x1e,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x54, 0x0a, 0x18, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x84, 0x02, 0x0a, 0x0c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xdc, 0x01, 0x0a, 0x10, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d,
	0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x31, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x22, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x79, 0x6d, 0x6c, 0x2d, 0x67, 0x72, 0x61, 0x70, 0x68,
	0xfa, 0x81, 0xf9, 0x1b, 0x2b, 0x0a, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2d, 0x79, 0x6d, 0x6c, 0x2d, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x1a, 0x15, 0xc2, 0xc4, 0xcb, 0x1c, 0x10, 0x22, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x32, 0x04, 0x10, 0x01, 0x20, 0x01, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_graph_proto_rawDescOnce sync.Once
	file_graph_proto_rawDescData = file_graph_proto_rawDesc
)

func file_graph_proto_rawDescGZIP() []byte {
	file_graph_proto_rawDescOnce.Do(func() {
		file_graph_proto_rawDescData = protoimpl.X.CompressGZIP(file_graph_proto_rawDescData)
	})
	return file_graph_proto_rawDescData
}

var file_graph_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_graph_proto_goTypes = []interface{}{
	(*PipelineYmlGraphRequest)(nil),  // 0: erda.core.pipeline.graph.PipelineYmlGraphRequest
	(*PipelineYmlGraphResponse)(nil), // 1: erda.core.pipeline.graph.PipelineYmlGraphResponse
	nil,                              // 2: erda.core.pipeline.graph.PipelineYmlGraphRequest.GlobalSnippetConfigLabelsEntry
	(*pb.SnippetConfig)(nil),         // 3: erda.core.pipeline.base.SnippetConfig
	(*pb.PipelineYml)(nil),           // 4: erda.core.pipeline.base.PipelineYml
}
var file_graph_proto_depIdxs = []int32{
	2, // 0: erda.core.pipeline.graph.PipelineYmlGraphRequest.globalSnippetConfigLabels:type_name -> erda.core.pipeline.graph.PipelineYmlGraphRequest.GlobalSnippetConfigLabelsEntry
	3, // 1: erda.core.pipeline.graph.PipelineYmlGraphRequest.snippetConfig:type_name -> erda.core.pipeline.base.SnippetConfig
	4, // 2: erda.core.pipeline.graph.PipelineYmlGraphResponse.data:type_name -> erda.core.pipeline.base.PipelineYml
	0, // 3: erda.core.pipeline.graph.GraphService.PipelineYmlGraph:input_type -> erda.core.pipeline.graph.PipelineYmlGraphRequest
	1, // 4: erda.core.pipeline.graph.GraphService.PipelineYmlGraph:output_type -> erda.core.pipeline.graph.PipelineYmlGraphResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_graph_proto_init() }
func file_graph_proto_init() {
	if File_graph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineYmlGraphRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_graph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineYmlGraphResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_graph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_graph_proto_goTypes,
		DependencyIndexes: file_graph_proto_depIdxs,
		MessageInfos:      file_graph_proto_msgTypes,
	}.Build()
	File_graph_proto = out.File
	file_graph_proto_rawDesc = nil
	file_graph_proto_goTypes = nil
	file_graph_proto_depIdxs = nil
}
