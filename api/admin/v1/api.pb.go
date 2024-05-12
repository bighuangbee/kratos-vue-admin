// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetApiRequest) Reset() {
	*x = GetApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiRequest) ProtoMessage() {}

func (x *GetApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiRequest.ProtoReflect.Descriptor instead.
func (*GetApiRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetApiRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *ApiData `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
}

func (x *GetApiReply) Reset() {
	*x = GetApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiReply) ProtoMessage() {}

func (x *GetApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiReply.ProtoReflect.Descriptor instead.
func (*GetApiReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetApiReply) GetApi() *ApiData {
	if x != nil {
		return x.Api
	}
	return nil
}

type ListApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum     int32  `protobuf:"varint,1,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize    int32  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Path        string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	ApiGroup    string `protobuf:"bytes,4,opt,name=apiGroup,proto3" json:"apiGroup,omitempty"`
	Method      string `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ListApiRequest) Reset() {
	*x = ListApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiRequest) ProtoMessage() {}

func (x *ListApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiRequest.ProtoReflect.Descriptor instead.
func (*ListApiRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListApiRequest) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListApiRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListApiRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListApiRequest) GetApiGroup() string {
	if x != nil {
		return x.ApiGroup
	}
	return ""
}

func (x *ListApiRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ListApiRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ListApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int32      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageNum  int32      `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize int32      `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Data     []*ApiData `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListApiReply) Reset() {
	*x = ListApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiReply) ProtoMessage() {}

func (x *ListApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiReply.ProtoReflect.Descriptor instead.
func (*ListApiReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *ListApiReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListApiReply) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListApiReply) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListApiReply) GetData() []*ApiData {
	if x != nil {
		return x.Data
	}
	return nil
}

type AllApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AllApiRequest) Reset() {
	*x = AllApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllApiRequest) ProtoMessage() {}

func (x *AllApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllApiRequest.ProtoReflect.Descriptor instead.
func (*AllApiRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

type AllApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ApiData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AllApiReply) Reset() {
	*x = AllApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllApiReply) ProtoMessage() {}

func (x *AllApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllApiReply.ProtoReflect.Descriptor instead.
func (*AllApiReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *AllApiReply) GetData() []*ApiData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path        string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ApiGroup    string `protobuf:"bytes,3,opt,name=apiGroup,proto3" json:"apiGroup,omitempty"`
	Method      string `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *CreateApiRequest) Reset() {
	*x = CreateApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiRequest) ProtoMessage() {}

func (x *CreateApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiRequest.ProtoReflect.Descriptor instead.
func (*CreateApiRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *CreateApiRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateApiRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateApiRequest) GetApiGroup() string {
	if x != nil {
		return x.ApiGroup
	}
	return ""
}

func (x *CreateApiRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type CreateApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateApiReply) Reset() {
	*x = CreateApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiReply) ProtoMessage() {}

func (x *CreateApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiReply.ProtoReflect.Descriptor instead.
func (*CreateApiReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

type UpdateApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Path        string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	ApiGroup    string `protobuf:"bytes,6,opt,name=apiGroup,proto3" json:"apiGroup,omitempty"`
	Method      string `protobuf:"bytes,7,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *UpdateApiRequest) Reset() {
	*x = UpdateApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApiRequest) ProtoMessage() {}

func (x *UpdateApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApiRequest.ProtoReflect.Descriptor instead.
func (*UpdateApiRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateApiRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateApiRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UpdateApiRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateApiRequest) GetApiGroup() string {
	if x != nil {
		return x.ApiGroup
	}
	return ""
}

func (x *UpdateApiRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type UpdateApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateApiReply) Reset() {
	*x = UpdateApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApiReply) ProtoMessage() {}

func (x *UpdateApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApiReply.ProtoReflect.Descriptor instead.
func (*UpdateApiReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

type DeleteApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteApiRequest) Reset() {
	*x = DeleteApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApiRequest) ProtoMessage() {}

func (x *DeleteApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApiRequest.ProtoReflect.Descriptor instead.
func (*DeleteApiRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteApiRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteApiReply) Reset() {
	*x = DeleteApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApiReply) ProtoMessage() {}

func (x *DeleteApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApiReply.ProtoReflect.Descriptor instead.
func (*DeleteApiReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{11}
}

type GetPolicyPathByRoleKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleKey string `protobuf:"bytes,1,opt,name=roleKey,proto3" json:"roleKey,omitempty"`
}

func (x *GetPolicyPathByRoleKeyRequest) Reset() {
	*x = GetPolicyPathByRoleKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyPathByRoleKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyPathByRoleKeyRequest) ProtoMessage() {}

func (x *GetPolicyPathByRoleKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyPathByRoleKeyRequest.ProtoReflect.Descriptor instead.
func (*GetPolicyPathByRoleKeyRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{12}
}

func (x *GetPolicyPathByRoleKeyRequest) GetRoleKey() string {
	if x != nil {
		return x.RoleKey
	}
	return ""
}

type GetPolicyPathByRoleKeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apis []*ApiBase `protobuf:"bytes,1,rep,name=apis,proto3" json:"apis,omitempty"`
}

func (x *GetPolicyPathByRoleKeyReply) Reset() {
	*x = GetPolicyPathByRoleKeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyPathByRoleKeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyPathByRoleKeyReply) ProtoMessage() {}

func (x *GetPolicyPathByRoleKeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyPathByRoleKeyReply.ProtoReflect.Descriptor instead.
func (*GetPolicyPathByRoleKeyReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{13}
}

func (x *GetPolicyPathByRoleKeyReply) GetApis() []*ApiBase {
	if x != nil {
		return x.Apis
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x70, 0x69, 0x44, 0x61, 0x74, 0x61, 0x52, 0x03, 0x61, 0x70, 0x69, 0x22, 0xb0, 0x01, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x69,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x85, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x6c, 0x6c, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x7c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x8c, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x39, 0x0a, 0x1d, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x61, 0x74, 0x68, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f,
	0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c,
	0x65, 0x4b, 0x65, 0x79, 0x22, 0x48, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x50, 0x61, 0x74, 0x68, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x04, 0x61, 0x70, 0x69, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x69, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x61, 0x70, 0x69, 0x73, 0x32, 0xe4,
	0x05, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x5d, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70,
	0x69, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x06, 0x41, 0x6c, 0x6c, 0x41, 0x70, 0x69, 0x12,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x6c, 0x6c, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x41,
	0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12,
	0x0f, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x6c,
	0x12, 0x61, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x3a, 0x01, 0x2a, 0x12, 0x61, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69,
	0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x1a, 0x0b, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x3a, 0x01, 0x2a, 0x12, 0x9b, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x50, 0x61, 0x74, 0x68, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x61, 0x74, 0x68, 0x42, 0x79,
	0x52, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x61, 0x74, 0x68, 0x42, 0x79, 0x52, 0x6f, 0x6c,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x23, 0x12, 0x21, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x61, 0x74, 0x68, 0x42, 0x79, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x5a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x12, 0x1b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x63, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x19, 0x5a, 0x17, 0x66, 0x65, 0x6e, 0x67, 0x79, 0x69, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_proto_goTypes = []interface{}{
	(*GetApiRequest)(nil),                 // 0: api.admin.v1.GetApiRequest
	(*GetApiReply)(nil),                   // 1: api.admin.v1.GetApiReply
	(*ListApiRequest)(nil),                // 2: api.admin.v1.ListApiRequest
	(*ListApiReply)(nil),                  // 3: api.admin.v1.ListApiReply
	(*AllApiRequest)(nil),                 // 4: api.admin.v1.AllApiRequest
	(*AllApiReply)(nil),                   // 5: api.admin.v1.AllApiReply
	(*CreateApiRequest)(nil),              // 6: api.admin.v1.CreateApiRequest
	(*CreateApiReply)(nil),                // 7: api.admin.v1.CreateApiReply
	(*UpdateApiRequest)(nil),              // 8: api.admin.v1.UpdateApiRequest
	(*UpdateApiReply)(nil),                // 9: api.admin.v1.UpdateApiReply
	(*DeleteApiRequest)(nil),              // 10: api.admin.v1.DeleteApiRequest
	(*DeleteApiReply)(nil),                // 11: api.admin.v1.DeleteApiReply
	(*GetPolicyPathByRoleKeyRequest)(nil), // 12: api.admin.v1.GetPolicyPathByRoleKeyRequest
	(*GetPolicyPathByRoleKeyReply)(nil),   // 13: api.admin.v1.GetPolicyPathByRoleKeyReply
	(*ApiData)(nil),                       // 14: api.admin.v1.ApiData
	(*ApiBase)(nil),                       // 15: api.admin.v1.ApiBase
}
var file_api_proto_depIdxs = []int32{
	14, // 0: api.admin.v1.GetApiReply.api:type_name -> api.admin.v1.ApiData
	14, // 1: api.admin.v1.ListApiReply.data:type_name -> api.admin.v1.ApiData
	14, // 2: api.admin.v1.AllApiReply.data:type_name -> api.admin.v1.ApiData
	15, // 3: api.admin.v1.GetPolicyPathByRoleKeyReply.apis:type_name -> api.admin.v1.ApiBase
	2,  // 4: api.admin.v1.Api.ListApi:input_type -> api.admin.v1.ListApiRequest
	4,  // 5: api.admin.v1.Api.AllApi:input_type -> api.admin.v1.AllApiRequest
	6,  // 6: api.admin.v1.Api.CreateApi:input_type -> api.admin.v1.CreateApiRequest
	8,  // 7: api.admin.v1.Api.UpdateApi:input_type -> api.admin.v1.UpdateApiRequest
	12, // 8: api.admin.v1.Api.GetPolicyPathByRoleKey:input_type -> api.admin.v1.GetPolicyPathByRoleKeyRequest
	0,  // 9: api.admin.v1.Api.GetApi:input_type -> api.admin.v1.GetApiRequest
	10, // 10: api.admin.v1.Api.DeleteApi:input_type -> api.admin.v1.DeleteApiRequest
	3,  // 11: api.admin.v1.Api.ListApi:output_type -> api.admin.v1.ListApiReply
	5,  // 12: api.admin.v1.Api.AllApi:output_type -> api.admin.v1.AllApiReply
	7,  // 13: api.admin.v1.Api.CreateApi:output_type -> api.admin.v1.CreateApiReply
	9,  // 14: api.admin.v1.Api.UpdateApi:output_type -> api.admin.v1.UpdateApiReply
	13, // 15: api.admin.v1.Api.GetPolicyPathByRoleKey:output_type -> api.admin.v1.GetPolicyPathByRoleKeyReply
	1,  // 16: api.admin.v1.Api.GetApi:output_type -> api.admin.v1.GetApiReply
	11, // 17: api.admin.v1.Api.DeleteApi:output_type -> api.admin.v1.DeleteApiReply
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApiRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApiReply); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApiRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApiReply); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllApiRequest); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllApiReply); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiRequest); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiReply); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateApiRequest); i {
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
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateApiReply); i {
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
		file_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApiRequest); i {
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
		file_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApiReply); i {
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
		file_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyPathByRoleKeyRequest); i {
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
		file_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyPathByRoleKeyReply); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
