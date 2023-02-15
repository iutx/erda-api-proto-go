// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: status.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusEnum int32

const (
	StatusEnum_unknown_status StatusEnum = 0
	StatusEnum_not_found      StatusEnum = 404
)

// Enum value maps for StatusEnum.
var (
	StatusEnum_name = map[int32]string{
		0:   "unknown_status",
		404: "not_found",
	}
	StatusEnum_value = map[string]int32{
		"unknown_status": 0,
		"not_found":      404,
	}
)

func (x StatusEnum) Enum() *StatusEnum {
	p := new(StatusEnum)
	*p = x
	return p
}

func (x StatusEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_status_proto_enumTypes[0].Descriptor()
}

func (StatusEnum) Type() protoreflect.EnumType {
	return &file_status_proto_enumTypes[0]
}

func (x StatusEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusEnum.Descriptor instead.
func (StatusEnum) EnumDescriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{0}
}

var File_status_proto protoreflect.FileDescriptor

var file_status_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0x30, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x0e, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x94, 0x03, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61,
	0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_status_proto_rawDescOnce sync.Once
	file_status_proto_rawDescData = file_status_proto_rawDesc
)

func file_status_proto_rawDescGZIP() []byte {
	file_status_proto_rawDescOnce.Do(func() {
		file_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_status_proto_rawDescData)
	})
	return file_status_proto_rawDescData
}

var file_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_status_proto_goTypes = []interface{}{
	(StatusEnum)(0), // 0: erda.common.StatusEnum
}
var file_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_status_proto_init() }
func file_status_proto_init() {
	if File_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_status_proto_goTypes,
		DependencyIndexes: file_status_proto_depIdxs,
		EnumInfos:         file_status_proto_enumTypes,
	}.Build()
	File_status_proto = out.File
	file_status_proto_rawDesc = nil
	file_status_proto_goTypes = nil
	file_status_proto_depIdxs = nil
}