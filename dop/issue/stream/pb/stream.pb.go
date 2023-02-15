// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: stream.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	pb "github.com/erda-project/erda-proto-go/common/pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommentIssueStreamCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssueID int64          `protobuf:"varint,1,opt,name=issueID,proto3" json:"issueID,omitempty"`
	Type    string         `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Content string         `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	MrInfo  *MRCommentInfo `protobuf:"bytes,4,opt,name=mrInfo,proto3" json:"mrInfo,omitempty"`
	UserID  string         `protobuf:"bytes,5,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *CommentIssueStreamCreateRequest) Reset() {
	*x = CommentIssueStreamCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentIssueStreamCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentIssueStreamCreateRequest) ProtoMessage() {}

func (x *CommentIssueStreamCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentIssueStreamCreateRequest.ProtoReflect.Descriptor instead.
func (*CommentIssueStreamCreateRequest) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{0}
}

func (x *CommentIssueStreamCreateRequest) GetIssueID() int64 {
	if x != nil {
		return x.IssueID
	}
	return 0
}

func (x *CommentIssueStreamCreateRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CommentIssueStreamCreateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentIssueStreamCreateRequest) GetMrInfo() *MRCommentInfo {
	if x != nil {
		return x.MrInfo
	}
	return nil
}

func (x *CommentIssueStreamCreateRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type MRCommentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID   int64  `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
	MrID    int64  `protobuf:"varint,2,opt,name=mrID,proto3" json:"mrID,omitempty"`
	MrTitle string `protobuf:"bytes,3,opt,name=mrTitle,proto3" json:"mrTitle,omitempty"`
}

func (x *MRCommentInfo) Reset() {
	*x = MRCommentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MRCommentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MRCommentInfo) ProtoMessage() {}

func (x *MRCommentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MRCommentInfo.ProtoReflect.Descriptor instead.
func (*MRCommentInfo) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{1}
}

func (x *MRCommentInfo) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *MRCommentInfo) GetMrID() int64 {
	if x != nil {
		return x.MrID
	}
	return 0
}

func (x *MRCommentInfo) GetMrTitle() string {
	if x != nil {
		return x.MrTitle
	}
	return ""
}

type CommentIssueStreamBatchCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssueStreams []*CommentIssueStreamCreateRequest `protobuf:"bytes,1,rep,name=issueStreams,proto3" json:"issueStreams,omitempty"`
}

func (x *CommentIssueStreamBatchCreateRequest) Reset() {
	*x = CommentIssueStreamBatchCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentIssueStreamBatchCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentIssueStreamBatchCreateRequest) ProtoMessage() {}

func (x *CommentIssueStreamBatchCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentIssueStreamBatchCreateRequest.ProtoReflect.Descriptor instead.
func (*CommentIssueStreamBatchCreateRequest) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{2}
}

func (x *CommentIssueStreamBatchCreateRequest) GetIssueStreams() []*CommentIssueStreamCreateRequest {
	if x != nil {
		return x.IssueStreams
	}
	return nil
}

type CommentIssueStreamBatchCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommentIssueStreamBatchCreateResponse) Reset() {
	*x = CommentIssueStreamBatchCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentIssueStreamBatchCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentIssueStreamBatchCreateResponse) ProtoMessage() {}

func (x *CommentIssueStreamBatchCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentIssueStreamBatchCreateResponse.ProtoReflect.Descriptor instead.
func (*CommentIssueStreamBatchCreateResponse) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{3}
}

type IssueStreamCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssueID      int64            `protobuf:"varint,1,opt,name=issueID,proto3" json:"issueID,omitempty"`
	Type         string           `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Content      string           `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	MrInfo       *MRCommentInfo   `protobuf:"bytes,4,opt,name=mrInfo,proto3" json:"mrInfo,omitempty"`
	IdentityInfo *pb.IdentityInfo `protobuf:"bytes,5,opt,name=identityInfo,json=-,proto3" json:"identityInfo,omitempty"`
	Id           string           `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IssueStreamCreateRequest) Reset() {
	*x = IssueStreamCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueStreamCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueStreamCreateRequest) ProtoMessage() {}

func (x *IssueStreamCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueStreamCreateRequest.ProtoReflect.Descriptor instead.
func (*IssueStreamCreateRequest) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{4}
}

func (x *IssueStreamCreateRequest) GetIssueID() int64 {
	if x != nil {
		return x.IssueID
	}
	return 0
}

func (x *IssueStreamCreateRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IssueStreamCreateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *IssueStreamCreateRequest) GetMrInfo() *MRCommentInfo {
	if x != nil {
		return x.MrInfo
	}
	return nil
}

func (x *IssueStreamCreateRequest) GetIdentityInfo() *pb.IdentityInfo {
	if x != nil {
		return x.IdentityInfo
	}
	return nil
}

func (x *IssueStreamCreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IssueStreamCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data int64 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *IssueStreamCreateResponse) Reset() {
	*x = IssueStreamCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueStreamCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueStreamCreateResponse) ProtoMessage() {}

func (x *IssueStreamCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueStreamCreateResponse.ProtoReflect.Descriptor instead.
func (*IssueStreamCreateResponse) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{5}
}

func (x *IssueStreamCreateResponse) GetData() int64 {
	if x != nil {
		return x.Data
	}
	return 0
}

type PagingIssueStreamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssueID  uint64 `protobuf:"varint,1,opt,name=issueID,proto3" json:"issueID,omitempty"`
	PageNo   uint64 `protobuf:"varint,2,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize uint64 `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Id       string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PagingIssueStreamsRequest) Reset() {
	*x = PagingIssueStreamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingIssueStreamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingIssueStreamsRequest) ProtoMessage() {}

func (x *PagingIssueStreamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingIssueStreamsRequest.ProtoReflect.Descriptor instead.
func (*PagingIssueStreamsRequest) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{6}
}

func (x *PagingIssueStreamsRequest) GetIssueID() uint64 {
	if x != nil {
		return x.IssueID
	}
	return 0
}

func (x *PagingIssueStreamsRequest) GetPageNo() uint64 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *PagingIssueStreamsRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PagingIssueStreamsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PagingIssueStreamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *IssueStreamPagingResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	UserIDs []string                       `protobuf:"bytes,2,rep,name=userIDs,proto3" json:"userIDs,omitempty"`
}

func (x *PagingIssueStreamsResponse) Reset() {
	*x = PagingIssueStreamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingIssueStreamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingIssueStreamsResponse) ProtoMessage() {}

func (x *PagingIssueStreamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingIssueStreamsResponse.ProtoReflect.Descriptor instead.
func (*PagingIssueStreamsResponse) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{7}
}

func (x *PagingIssueStreamsResponse) GetData() *IssueStreamPagingResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PagingIssueStreamsResponse) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

type IssueStreamPagingResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*IssueStream `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *IssueStreamPagingResponseData) Reset() {
	*x = IssueStreamPagingResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueStreamPagingResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueStreamPagingResponseData) ProtoMessage() {}

func (x *IssueStreamPagingResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueStreamPagingResponseData.ProtoReflect.Descriptor instead.
func (*IssueStreamPagingResponseData) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{8}
}

func (x *IssueStreamPagingResponseData) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *IssueStreamPagingResponseData) GetList() []*IssueStream {
	if x != nil {
		return x.List
	}
	return nil
}

type IssueStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IssueID    int64                  `protobuf:"varint,2,opt,name=issueID,proto3" json:"issueID,omitempty"`
	Operator   string                 `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	StreamType string                 `protobuf:"bytes,4,opt,name=streamType,proto3" json:"streamType,omitempty"`
	Content    string                 `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	MrInfo     *MRCommentInfo         `protobuf:"bytes,6,opt,name=mrInfo,proto3" json:"mrInfo,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *IssueStream) Reset() {
	*x = IssueStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueStream) ProtoMessage() {}

func (x *IssueStream) ProtoReflect() protoreflect.Message {
	mi := &file_stream_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueStream.ProtoReflect.Descriptor instead.
func (*IssueStream) Descriptor() ([]byte, []int) {
	return file_stream_proto_rawDescGZIP(), []int{9}
}

func (x *IssueStream) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IssueStream) GetIssueID() int64 {
	if x != nil {
		return x.IssueID
	}
	return 0
}

func (x *IssueStream) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *IssueStream) GetStreamType() string {
	if x != nil {
		return x.StreamType
	}
	return ""
}

func (x *IssueStream) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *IssueStream) GetMrInfo() *MRCommentInfo {
	if x != nil {
		return x.MrInfo
	}
	return nil
}

func (x *IssueStream) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IssueStream) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_stream_proto protoreflect.FileDescriptor

var file_stream_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x1f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a,
	0x06, 0x6d, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4d, 0x52, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6d, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x22, 0x53, 0x0a, 0x0d, 0x4d, 0x52, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6d, 0x72, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x72, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x72, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x24, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x5a, 0x0a, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64,
	0x6f, 0x70, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x27, 0x0a,
	0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe4, 0x01, 0x0a, 0x18, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x6d,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x72,
	0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x4d, 0x52, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x6d, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x0c, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x01, 0x2d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a,
	0x19, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x79,
	0x0a, 0x19, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x1a, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f,
	0x70, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x22, 0x6d, 0x0a, 0x1d,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xbf, 0x02, 0x0a, 0x0b,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x6d,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x72,
	0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x4d, 0x52, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x6d, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xa9, 0x05,
	0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x02, 0x0a, 0x16,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3b, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f,
	0x70, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x22, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x2d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xfa, 0x81, 0xf9, 0x1b, 0x31, 0x0a,
	0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0xb7, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2f, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f,
	0x70, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64,
	0x6f, 0x70, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0xfa, 0x81, 0xf9, 0x1b, 0x1a,
	0x0a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0xba, 0x01, 0x0a, 0x12, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x12, 0x30, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0xfa, 0x81, 0xf9, 0x1b, 0x1a, 0x0a, 0x18, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0xc2, 0xc4, 0xcb, 0x1c, 0x0b, 0x22, 0x03,
	0x64, 0x6f, 0x70, 0x32, 0x04, 0x10, 0x01, 0x20, 0x01, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x67, 0x6f, 0x2f, 0x64, 0x6f, 0x70, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_proto_rawDescOnce sync.Once
	file_stream_proto_rawDescData = file_stream_proto_rawDesc
)

func file_stream_proto_rawDescGZIP() []byte {
	file_stream_proto_rawDescOnce.Do(func() {
		file_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_proto_rawDescData)
	})
	return file_stream_proto_rawDescData
}

var file_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_stream_proto_goTypes = []interface{}{
	(*CommentIssueStreamCreateRequest)(nil),       // 0: erda.dop.issue.stream.CommentIssueStreamCreateRequest
	(*MRCommentInfo)(nil),                         // 1: erda.dop.issue.stream.MRCommentInfo
	(*CommentIssueStreamBatchCreateRequest)(nil),  // 2: erda.dop.issue.stream.CommentIssueStreamBatchCreateRequest
	(*CommentIssueStreamBatchCreateResponse)(nil), // 3: erda.dop.issue.stream.CommentIssueStreamBatchCreateResponse
	(*IssueStreamCreateRequest)(nil),              // 4: erda.dop.issue.stream.IssueStreamCreateRequest
	(*IssueStreamCreateResponse)(nil),             // 5: erda.dop.issue.stream.IssueStreamCreateResponse
	(*PagingIssueStreamsRequest)(nil),             // 6: erda.dop.issue.stream.PagingIssueStreamsRequest
	(*PagingIssueStreamsResponse)(nil),            // 7: erda.dop.issue.stream.PagingIssueStreamsResponse
	(*IssueStreamPagingResponseData)(nil),         // 8: erda.dop.issue.stream.IssueStreamPagingResponseData
	(*IssueStream)(nil),                           // 9: erda.dop.issue.stream.IssueStream
	(*pb.IdentityInfo)(nil),                       // 10: erda.common.IdentityInfo
	(*timestamppb.Timestamp)(nil),                 // 11: google.protobuf.Timestamp
}
var file_stream_proto_depIdxs = []int32{
	1,  // 0: erda.dop.issue.stream.CommentIssueStreamCreateRequest.mrInfo:type_name -> erda.dop.issue.stream.MRCommentInfo
	0,  // 1: erda.dop.issue.stream.CommentIssueStreamBatchCreateRequest.issueStreams:type_name -> erda.dop.issue.stream.CommentIssueStreamCreateRequest
	1,  // 2: erda.dop.issue.stream.IssueStreamCreateRequest.mrInfo:type_name -> erda.dop.issue.stream.MRCommentInfo
	10, // 3: erda.dop.issue.stream.IssueStreamCreateRequest.identityInfo:type_name -> erda.common.IdentityInfo
	8,  // 4: erda.dop.issue.stream.PagingIssueStreamsResponse.data:type_name -> erda.dop.issue.stream.IssueStreamPagingResponseData
	9,  // 5: erda.dop.issue.stream.IssueStreamPagingResponseData.list:type_name -> erda.dop.issue.stream.IssueStream
	1,  // 6: erda.dop.issue.stream.IssueStream.mrInfo:type_name -> erda.dop.issue.stream.MRCommentInfo
	11, // 7: erda.dop.issue.stream.IssueStream.createdAt:type_name -> google.protobuf.Timestamp
	11, // 8: erda.dop.issue.stream.IssueStream.updatedAt:type_name -> google.protobuf.Timestamp
	2,  // 9: erda.dop.issue.stream.CommentIssueStreamService.BatchCreateIssueStream:input_type -> erda.dop.issue.stream.CommentIssueStreamBatchCreateRequest
	4,  // 10: erda.dop.issue.stream.CommentIssueStreamService.CreateIssueStream:input_type -> erda.dop.issue.stream.IssueStreamCreateRequest
	6,  // 11: erda.dop.issue.stream.CommentIssueStreamService.PagingIssueStreams:input_type -> erda.dop.issue.stream.PagingIssueStreamsRequest
	3,  // 12: erda.dop.issue.stream.CommentIssueStreamService.BatchCreateIssueStream:output_type -> erda.dop.issue.stream.CommentIssueStreamBatchCreateResponse
	5,  // 13: erda.dop.issue.stream.CommentIssueStreamService.CreateIssueStream:output_type -> erda.dop.issue.stream.IssueStreamCreateResponse
	7,  // 14: erda.dop.issue.stream.CommentIssueStreamService.PagingIssueStreams:output_type -> erda.dop.issue.stream.PagingIssueStreamsResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_stream_proto_init() }
func file_stream_proto_init() {
	if File_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentIssueStreamCreateRequest); i {
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
		file_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MRCommentInfo); i {
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
		file_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentIssueStreamBatchCreateRequest); i {
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
		file_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentIssueStreamBatchCreateResponse); i {
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
		file_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueStreamCreateRequest); i {
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
		file_stream_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueStreamCreateResponse); i {
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
		file_stream_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingIssueStreamsRequest); i {
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
		file_stream_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingIssueStreamsResponse); i {
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
		file_stream_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueStreamPagingResponseData); i {
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
		file_stream_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueStream); i {
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
			RawDescriptor: file_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_proto_goTypes,
		DependencyIndexes: file_stream_proto_depIdxs,
		MessageInfos:      file_stream_proto_msgTypes,
	}.Build()
	File_stream_proto = out.File
	file_stream_proto_rawDesc = nil
	file_stream_proto_goTypes = nil
	file_stream_proto_depIdxs = nil
}
