// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: account/service/v1/account_user.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccountUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AccountUser) Reset() {
	*x = AccountUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_v1_account_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountUser) ProtoMessage() {}

func (x *AccountUser) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_v1_account_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountUser.ProtoReflect.Descriptor instead.
func (*AccountUser) Descriptor() ([]byte, []int) {
	return file_account_service_v1_account_user_proto_rawDescGZIP(), []int{0}
}

func (x *AccountUser) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateAccountUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAccountUserReq) Reset() {
	*x = CreateAccountUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_v1_account_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountUserReq) ProtoMessage() {}

func (x *CreateAccountUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_v1_account_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountUserReq.ProtoReflect.Descriptor instead.
func (*CreateAccountUserReq) Descriptor() ([]byte, []int) {
	return file_account_service_v1_account_user_proto_rawDescGZIP(), []int{1}
}

type UpdateAccountUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAccountUserReq) Reset() {
	*x = UpdateAccountUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_v1_account_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountUserReq) ProtoMessage() {}

func (x *UpdateAccountUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_v1_account_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountUserReq.ProtoReflect.Descriptor instead.
func (*UpdateAccountUserReq) Descriptor() ([]byte, []int) {
	return file_account_service_v1_account_user_proto_rawDescGZIP(), []int{2}
}

type DeleteAccountUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAccountUserReq) Reset() {
	*x = DeleteAccountUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_v1_account_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountUserReq) ProtoMessage() {}

func (x *DeleteAccountUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_v1_account_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountUserReq.ProtoReflect.Descriptor instead.
func (*DeleteAccountUserReq) Descriptor() ([]byte, []int) {
	return file_account_service_v1_account_user_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAccountUserReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAccountUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAccountUserReq) Reset() {
	*x = GetAccountUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_v1_account_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountUserReq) ProtoMessage() {}

func (x *GetAccountUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_v1_account_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountUserReq.ProtoReflect.Descriptor instead.
func (*GetAccountUserReq) Descriptor() ([]byte, []int) {
	return file_account_service_v1_account_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetAccountUserReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListAccountUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAccountUserReq) Reset() {
	*x = ListAccountUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_v1_account_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountUserReq) ProtoMessage() {}

func (x *ListAccountUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_v1_account_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountUserReq.ProtoReflect.Descriptor instead.
func (*ListAccountUserReq) Descriptor() ([]byte, []int) {
	return file_account_service_v1_account_user_proto_rawDescGZIP(), []int{5}
}

type ListAccountUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*AccountUser `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListAccountUserResp) Reset() {
	*x = ListAccountUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_v1_account_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountUserResp) ProtoMessage() {}

func (x *ListAccountUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_v1_account_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountUserResp.ProtoReflect.Descriptor instead.
func (*ListAccountUserResp) Descriptor() ([]byte, []int) {
	return file_account_service_v1_account_user_proto_rawDescGZIP(), []int{6}
}

func (x *ListAccountUserResp) GetList() []*AccountUser {
	if x != nil {
		return x.List
	}
	return nil
}

type PageListAccountUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int64 `protobuf:"varint,1,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize int64 `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *PageListAccountUserReq) Reset() {
	*x = PageListAccountUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_v1_account_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageListAccountUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageListAccountUserReq) ProtoMessage() {}

func (x *PageListAccountUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_v1_account_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageListAccountUserReq.ProtoReflect.Descriptor instead.
func (*PageListAccountUserReq) Descriptor() ([]byte, []int) {
	return file_account_service_v1_account_user_proto_rawDescGZIP(), []int{7}
}

func (x *PageListAccountUserReq) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *PageListAccountUserReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PageListAccountUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List     []*AccountUser `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total    int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page     int64          `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64          `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *PageListAccountUserResp) Reset() {
	*x = PageListAccountUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_v1_account_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageListAccountUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageListAccountUserResp) ProtoMessage() {}

func (x *PageListAccountUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_v1_account_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageListAccountUserResp.ProtoReflect.Descriptor instead.
func (*PageListAccountUserResp) Descriptor() ([]byte, []int) {
	return file_account_service_v1_account_user_proto_rawDescGZIP(), []int{8}
}

func (x *PageListAccountUserResp) GetList() []*AccountUser {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *PageListAccountUserResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageListAccountUserResp) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageListAccountUserResp) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

var File_account_service_v1_account_user_proto protoreflect.FileDescriptor

var file_account_service_v1_account_user_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22,
	0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x4a, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x33, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x16, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x17, 0x50, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x33, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x32, 0xd9, 0x04,
	0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x58, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x25, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x62, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x27, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x6e, 0x0a, 0x13, 0x50, 0x61, 0x67,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x2a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x6c, 0x69, 0x66, 0x65, 0x69, 0x65,
	0x72, 0x2f, 0x76, 0x76, 0x67, 0x6f, 0x2d, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_service_v1_account_user_proto_rawDescOnce sync.Once
	file_account_service_v1_account_user_proto_rawDescData = file_account_service_v1_account_user_proto_rawDesc
)

func file_account_service_v1_account_user_proto_rawDescGZIP() []byte {
	file_account_service_v1_account_user_proto_rawDescOnce.Do(func() {
		file_account_service_v1_account_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_service_v1_account_user_proto_rawDescData)
	})
	return file_account_service_v1_account_user_proto_rawDescData
}

var file_account_service_v1_account_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_account_service_v1_account_user_proto_goTypes = []interface{}{
	(*AccountUser)(nil),             // 0: account.service.v1.AccountUser
	(*CreateAccountUserReq)(nil),    // 1: account.service.v1.CreateAccountUserReq
	(*UpdateAccountUserReq)(nil),    // 2: account.service.v1.UpdateAccountUserReq
	(*DeleteAccountUserReq)(nil),    // 3: account.service.v1.DeleteAccountUserReq
	(*GetAccountUserReq)(nil),       // 4: account.service.v1.GetAccountUserReq
	(*ListAccountUserReq)(nil),      // 5: account.service.v1.ListAccountUserReq
	(*ListAccountUserResp)(nil),     // 6: account.service.v1.ListAccountUserResp
	(*PageListAccountUserReq)(nil),  // 7: account.service.v1.PageListAccountUserReq
	(*PageListAccountUserResp)(nil), // 8: account.service.v1.PageListAccountUserResp
	(*emptypb.Empty)(nil),           // 9: google.protobuf.Empty
}
var file_account_service_v1_account_user_proto_depIdxs = []int32{
	0, // 0: account.service.v1.ListAccountUserResp.list:type_name -> account.service.v1.AccountUser
	0, // 1: account.service.v1.PageListAccountUserResp.list:type_name -> account.service.v1.AccountUser
	1, // 2: account.service.v1.AccountUserService.CreateAccountUser:input_type -> account.service.v1.CreateAccountUserReq
	2, // 3: account.service.v1.AccountUserService.UpdateAccountUser:input_type -> account.service.v1.UpdateAccountUserReq
	3, // 4: account.service.v1.AccountUserService.DeleteAccountUser:input_type -> account.service.v1.DeleteAccountUserReq
	4, // 5: account.service.v1.AccountUserService.GetAccountUser:input_type -> account.service.v1.GetAccountUserReq
	5, // 6: account.service.v1.AccountUserService.ListAccountUser:input_type -> account.service.v1.ListAccountUserReq
	7, // 7: account.service.v1.AccountUserService.PageListAccountUser:input_type -> account.service.v1.PageListAccountUserReq
	0, // 8: account.service.v1.AccountUserService.CreateAccountUser:output_type -> account.service.v1.AccountUser
	0, // 9: account.service.v1.AccountUserService.UpdateAccountUser:output_type -> account.service.v1.AccountUser
	9, // 10: account.service.v1.AccountUserService.DeleteAccountUser:output_type -> google.protobuf.Empty
	0, // 11: account.service.v1.AccountUserService.GetAccountUser:output_type -> account.service.v1.AccountUser
	6, // 12: account.service.v1.AccountUserService.ListAccountUser:output_type -> account.service.v1.ListAccountUserResp
	8, // 13: account.service.v1.AccountUserService.PageListAccountUser:output_type -> account.service.v1.PageListAccountUserResp
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_account_service_v1_account_user_proto_init() }
func file_account_service_v1_account_user_proto_init() {
	if File_account_service_v1_account_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_service_v1_account_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountUser); i {
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
		file_account_service_v1_account_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountUserReq); i {
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
		file_account_service_v1_account_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountUserReq); i {
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
		file_account_service_v1_account_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountUserReq); i {
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
		file_account_service_v1_account_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountUserReq); i {
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
		file_account_service_v1_account_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountUserReq); i {
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
		file_account_service_v1_account_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountUserResp); i {
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
		file_account_service_v1_account_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageListAccountUserReq); i {
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
		file_account_service_v1_account_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageListAccountUserResp); i {
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
			RawDescriptor: file_account_service_v1_account_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_service_v1_account_user_proto_goTypes,
		DependencyIndexes: file_account_service_v1_account_user_proto_depIdxs,
		MessageInfos:      file_account_service_v1_account_user_proto_msgTypes,
	}.Build()
	File_account_service_v1_account_user_proto = out.File
	file_account_service_v1_account_user_proto_rawDesc = nil
	file_account_service_v1_account_user_proto_goTypes = nil
	file_account_service_v1_account_user_proto_depIdxs = nil
}
