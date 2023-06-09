// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: user.proto

package user

import (
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

type AvatarEnum int32

const (
	AvatarEnum_ZERO AvatarEnum = 0
	AvatarEnum_ONE  AvatarEnum = 1
)

// Enum value maps for AvatarEnum.
var (
	AvatarEnum_name = map[int32]string{
		0: "ZERO",
		1: "ONE",
	}
	AvatarEnum_value = map[string]int32{
		"ZERO": 0,
		"ONE":  1,
	}
)

func (x AvatarEnum) Enum() *AvatarEnum {
	p := new(AvatarEnum)
	*p = x
	return p
}

func (x AvatarEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvatarEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (AvatarEnum) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x AvatarEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvatarEnum.Descriptor instead.
func (AvatarEnum) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type GenderEnum int32

const (
	GenderEnum_GIRL GenderEnum = 0
	GenderEnum_BOY  GenderEnum = 1
)

// Enum value maps for GenderEnum.
var (
	GenderEnum_name = map[int32]string{
		0: "GIRL",
		1: "BOY",
	}
	GenderEnum_value = map[string]int32{
		"GIRL": 0,
		"BOY":  1,
	}
)

func (x GenderEnum) Enum() *GenderEnum {
	p := new(GenderEnum)
	*p = x
	return p
}

func (x GenderEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenderEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[1].Descriptor()
}

func (GenderEnum) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[1]
}

func (x GenderEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenderEnum.Descriptor instead.
func (GenderEnum) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

// GetMemberList
type GetMemberListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64           `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Message    string          `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Data       *MemberListInfo `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetMemberListRes) Reset() {
	*x = GetMemberListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberListRes) ProtoMessage() {}

func (x *GetMemberListRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberListRes.ProtoReflect.Descriptor instead.
func (*GetMemberListRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetMemberListRes) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetMemberListRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMemberListRes) GetData() *MemberListInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type MemberListInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberList []*MemberList `protobuf:"bytes,1,rep,name=memberList,proto3" json:"memberList,omitempty"`
}

func (x *MemberListInfo) Reset() {
	*x = MemberListInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberListInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberListInfo) ProtoMessage() {}

func (x *MemberListInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberListInfo.ProtoReflect.Descriptor instead.
func (*MemberListInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *MemberListInfo) GetMemberList() []*MemberList {
	if x != nil {
		return x.MemberList
	}
	return nil
}

type MemberList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account        string     `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password       string     `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Avatar         AvatarEnum `protobuf:"varint,3,opt,name=avatar,proto3,enum=user.AvatarEnum" json:"avatar,omitempty"`
	Gender         GenderEnum `protobuf:"varint,4,opt,name=gender,proto3,enum=user.GenderEnum" json:"gender,omitempty"`
	LastLoginTime  int64      `protobuf:"varint,5,opt,name=lastLoginTime,proto3" json:"lastLoginTime,omitempty"`
	LastLogoutTime int64      `protobuf:"varint,6,opt,name=lastLogoutTime,proto3" json:"lastLogoutTime,omitempty"`
	Balance        int64      `protobuf:"varint,7,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *MemberList) Reset() {
	*x = MemberList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberList) ProtoMessage() {}

func (x *MemberList) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberList.ProtoReflect.Descriptor instead.
func (*MemberList) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *MemberList) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *MemberList) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *MemberList) GetAvatar() AvatarEnum {
	if x != nil {
		return x.Avatar
	}
	return AvatarEnum_ZERO
}

func (x *MemberList) GetGender() GenderEnum {
	if x != nil {
		return x.Gender
	}
	return GenderEnum_GIRL
}

func (x *MemberList) GetLastLoginTime() int64 {
	if x != nil {
		return x.LastLoginTime
	}
	return 0
}

func (x *MemberList) GetLastLogoutTime() int64 {
	if x != nil {
		return x.LastLogoutTime
	}
	return 0
}

func (x *MemberList) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

// SetMember
type SetMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string     `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password string     `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Gender   AvatarEnum `protobuf:"varint,3,opt,name=Gender,proto3,enum=user.AvatarEnum" json:"Gender,omitempty"`
	Avatar   GenderEnum `protobuf:"varint,4,opt,name=Avatar,proto3,enum=user.GenderEnum" json:"Avatar,omitempty"`
	Balance  int64      `protobuf:"varint,5,opt,name=Balance,proto3" json:"Balance,omitempty"`
}

func (x *SetMember) Reset() {
	*x = SetMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMember) ProtoMessage() {}

func (x *SetMember) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMember.ProtoReflect.Descriptor instead.
func (*SetMember) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *SetMember) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SetMember) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SetMember) GetGender() AvatarEnum {
	if x != nil {
		return x.Gender
	}
	return AvatarEnum_ZERO
}

func (x *SetMember) GetAvatar() GenderEnum {
	if x != nil {
		return x.Avatar
	}
	return GenderEnum_GIRL
}

func (x *SetMember) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type SetMemberRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *SetMemberRes) Reset() {
	*x = SetMemberRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMemberRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMemberRes) ProtoMessage() {}

func (x *SetMemberRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMemberRes.ProtoReflect.Descriptor instead.
func (*SetMemberRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *SetMemberRes) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *SetMemberRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// KickOutMember
type KickOutMemberInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Balance int64  `protobuf:"varint,2,opt,name=Balance,proto3" json:"Balance,omitempty"`
}

func (x *KickOutMemberInfo) Reset() {
	*x = KickOutMemberInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickOutMemberInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickOutMemberInfo) ProtoMessage() {}

func (x *KickOutMemberInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickOutMemberInfo.ProtoReflect.Descriptor instead.
func (*KickOutMemberInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *KickOutMemberInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *KickOutMemberInfo) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type KickOutMemberRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *KickOutMemberRes) Reset() {
	*x = KickOutMemberRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickOutMemberRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickOutMemberRes) ProtoMessage() {}

func (x *KickOutMemberRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickOutMemberRes.ProtoReflect.Descriptor instead.
func (*KickOutMemberRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *KickOutMemberRes) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *KickOutMemberRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x22, 0x76,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x42, 0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0a,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xfe, 0x01, 0x0a, 0x0a, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x28, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x28, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x09,
	0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x28, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x06, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x48, 0x0a,
	0x0c, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x4b, 0x69, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x4c, 0x0a, 0x10, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x1f,
	0x0a, 0x0a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04,
	0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x2a,
	0x1f, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a,
	0x04, 0x47, 0x49, 0x52, 0x4c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x4f, 0x59, 0x10, 0x01,
	0x32, 0xb7, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x34, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0d, 0x4b, 0x69, 0x63, 0x6b,
	0x4f, 0x75, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_proto_goTypes = []interface{}{
	(AvatarEnum)(0),           // 0: user.AvatarEnum
	(GenderEnum)(0),           // 1: user.GenderEnum
	(*EmptyReq)(nil),          // 2: user.EmptyReq
	(*GetMemberListRes)(nil),  // 3: user.GetMemberListRes
	(*MemberListInfo)(nil),    // 4: user.MemberListInfo
	(*MemberList)(nil),        // 5: user.MemberList
	(*SetMember)(nil),         // 6: user.SetMember
	(*SetMemberRes)(nil),      // 7: user.SetMemberRes
	(*KickOutMemberInfo)(nil), // 8: user.KickOutMemberInfo
	(*KickOutMemberRes)(nil),  // 9: user.KickOutMemberRes
}
var file_user_proto_depIdxs = []int32{
	4, // 0: user.GetMemberListRes.Data:type_name -> user.MemberListInfo
	5, // 1: user.MemberListInfo.memberList:type_name -> user.MemberList
	0, // 2: user.MemberList.avatar:type_name -> user.AvatarEnum
	1, // 3: user.MemberList.gender:type_name -> user.GenderEnum
	0, // 4: user.SetMember.Gender:type_name -> user.AvatarEnum
	1, // 5: user.SetMember.Avatar:type_name -> user.GenderEnum
	2, // 6: user.User.GetMemberList:input_type -> user.EmptyReq
	6, // 7: user.User.SetMemberData:input_type -> user.SetMember
	8, // 8: user.User.KickOutMember:input_type -> user.KickOutMemberInfo
	3, // 9: user.User.GetMemberList:output_type -> user.GetMemberListRes
	7, // 10: user.User.SetMemberData:output_type -> user.SetMemberRes
	9, // 11: user.User.KickOutMember:output_type -> user.KickOutMemberRes
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemberListRes); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberListInfo); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberList); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMember); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMemberRes); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickOutMemberInfo); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickOutMemberRes); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
