// generated by openapi-gen-protobuf

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: msp_tenant.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/erda-project/erda-proto-go/common/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CREATE_TENANTS_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CREATE_TENANTS_Request) Reset() {
	*x = CREATE_TENANTS_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msp_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CREATE_TENANTS_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CREATE_TENANTS_Request) ProtoMessage() {}

func (x *CREATE_TENANTS_Request) ProtoReflect() protoreflect.Message {
	mi := &file_msp_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CREATE_TENANTS_Request.ProtoReflect.Descriptor instead.
func (*CREATE_TENANTS_Request) Descriptor() ([]byte, []int) {
	return file_msp_tenant_proto_rawDescGZIP(), []int{0}
}

type GET_TENANTS_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GET_TENANTS_Request) Reset() {
	*x = GET_TENANTS_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msp_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GET_TENANTS_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GET_TENANTS_Request) ProtoMessage() {}

func (x *GET_TENANTS_Request) ProtoReflect() protoreflect.Message {
	mi := &file_msp_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GET_TENANTS_Request.ProtoReflect.Descriptor instead.
func (*GET_TENANTS_Request) Descriptor() ([]byte, []int) {
	return file_msp_tenant_proto_rawDescGZIP(), []int{1}
}

var File_msp_tenant_proto protoreflect.FileDescriptor

var file_msp_tenant_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x73, 0x70, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2e, 0x6d, 0x73, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x18, 0x0a, 0x16, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e,
	0x54, 0x53, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x45,
	0x54, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x53, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0xcd, 0x02, 0x0a, 0x0a, 0x6d, 0x73, 0x70, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x12, 0x9c, 0x01, 0x0a, 0x0e, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x45, 0x4e, 0x41,
	0x4e, 0x54, 0x53, 0x12, 0x2a, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f,
	0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x53, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x46, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22,
	0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0xfa, 0x81, 0xf9, 0x1b, 0x2a, 0x0a, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x32, 0x04, 0x10, 0x01, 0x20, 0x01, 0x3a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x4d, 0x53, 0x50, 0x20, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12,
	0x93, 0x01, 0x0a, 0x0b, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x53, 0x12,
	0x27, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x47, 0x45, 0x54, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x53,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x43, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x73, 0x70, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0xfa, 0x81, 0xf9, 0x1b, 0x27, 0x0a, 0x0f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x32,
	0x04, 0x10, 0x01, 0x20, 0x01, 0x3a, 0x0e, 0x47, 0x45, 0x54, 0x20, 0x4d, 0x53, 0x50, 0x20, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x1a, 0x0a, 0xc2, 0xc4, 0xcb, 0x1c, 0x05, 0x22, 0x03, 0x6d, 0x73,
	0x70, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64,
	0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msp_tenant_proto_rawDescOnce sync.Once
	file_msp_tenant_proto_rawDescData = file_msp_tenant_proto_rawDesc
)

func file_msp_tenant_proto_rawDescGZIP() []byte {
	file_msp_tenant_proto_rawDescOnce.Do(func() {
		file_msp_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_msp_tenant_proto_rawDescData)
	})
	return file_msp_tenant_proto_rawDescData
}

var file_msp_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msp_tenant_proto_goTypes = []interface{}{
	(*CREATE_TENANTS_Request)(nil), // 0: erda.openapiv1.msp.CREATE_TENANTS_Request
	(*GET_TENANTS_Request)(nil),    // 1: erda.openapiv1.msp.GET_TENANTS_Request
	(*emptypb.Empty)(nil),          // 2: google.protobuf.Empty
}
var file_msp_tenant_proto_depIdxs = []int32{
	0, // 0: erda.openapiv1.msp.msp_tenant.CREATE_TENANTS:input_type -> erda.openapiv1.msp.CREATE_TENANTS_Request
	1, // 1: erda.openapiv1.msp.msp_tenant.GET_TENANTS:input_type -> erda.openapiv1.msp.GET_TENANTS_Request
	2, // 2: erda.openapiv1.msp.msp_tenant.CREATE_TENANTS:output_type -> google.protobuf.Empty
	2, // 3: erda.openapiv1.msp.msp_tenant.GET_TENANTS:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msp_tenant_proto_init() }
func file_msp_tenant_proto_init() {
	if File_msp_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msp_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CREATE_TENANTS_Request); i {
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
		file_msp_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GET_TENANTS_Request); i {
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
			RawDescriptor: file_msp_tenant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msp_tenant_proto_goTypes,
		DependencyIndexes: file_msp_tenant_proto_depIdxs,
		MessageInfos:      file_msp_tenant_proto_msgTypes,
	}.Build()
	File_msp_tenant_proto = out.File
	file_msp_tenant_proto_rawDesc = nil
	file_msp_tenant_proto_goTypes = nil
	file_msp_tenant_proto_depIdxs = nil
}
