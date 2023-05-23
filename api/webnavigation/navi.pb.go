// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/webnavigation/navi.proto

package webnavigation

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateNaviRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateNaviRequest) Reset() {
	*x = CreateNaviRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNaviRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNaviRequest) ProtoMessage() {}

func (x *CreateNaviRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNaviRequest.ProtoReflect.Descriptor instead.
func (*CreateNaviRequest) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNaviRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateNaviReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateNaviReply) Reset() {
	*x = CreateNaviReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNaviReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNaviReply) ProtoMessage() {}

func (x *CreateNaviReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNaviReply.ProtoReflect.Descriptor instead.
func (*CreateNaviReply) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNaviReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NaviType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Index int32  `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *NaviType) Reset() {
	*x = NaviType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NaviType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NaviType) ProtoMessage() {}

func (x *NaviType) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NaviType.ProtoReflect.Descriptor instead.
func (*NaviType) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{2}
}

func (x *NaviType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NaviType) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type DeleteNaviRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteNaviRequest) Reset() {
	*x = DeleteNaviRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNaviRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNaviRequest) ProtoMessage() {}

func (x *DeleteNaviRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNaviRequest.ProtoReflect.Descriptor instead.
func (*DeleteNaviRequest) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNaviRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteNaviReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteNaviReply) Reset() {
	*x = DeleteNaviReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNaviReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNaviReply) ProtoMessage() {}

func (x *DeleteNaviReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNaviReply.ProtoReflect.Descriptor instead.
func (*DeleteNaviReply) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{4}
}

type SortNaviRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"` // 把新的顺序完整的传过来
}

func (x *SortNaviRequest) Reset() {
	*x = SortNaviRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortNaviRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortNaviRequest) ProtoMessage() {}

func (x *SortNaviRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortNaviRequest.ProtoReflect.Descriptor instead.
func (*SortNaviRequest) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{5}
}

func (x *SortNaviRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type SortNaviReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SortNaviReply) Reset() {
	*x = SortNaviReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortNaviReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortNaviReply) ProtoMessage() {}

func (x *SortNaviReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortNaviReply.ProtoReflect.Descriptor instead.
func (*SortNaviReply) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{6}
}

type ListNaviRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListNaviRequest) Reset() {
	*x = ListNaviRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNaviRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNaviRequest) ProtoMessage() {}

func (x *ListNaviRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNaviRequest.ProtoReflect.Descriptor instead.
func (*ListNaviRequest) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{7}
}

type ListNaviReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*NaviType `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListNaviReply) Reset() {
	*x = ListNaviReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNaviReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNaviReply) ProtoMessage() {}

func (x *ListNaviReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNaviReply.ProtoReflect.Descriptor instead.
func (*ListNaviReply) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{8}
}

func (x *ListNaviReply) GetItems() []*NaviType {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetGuestSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip     string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`         // 根据cookie和IP判断，是否首次登陆
	Cookie string `protobuf:"bytes,2,opt,name=cookie,proto3" json:"cookie,omitempty"` // 根据cookie和IP判断，是否首次登陆
}

func (x *GetGuestSettingsRequest) Reset() {
	*x = GetGuestSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGuestSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGuestSettingsRequest) ProtoMessage() {}

func (x *GetGuestSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGuestSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetGuestSettingsRequest) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{9}
}

func (x *GetGuestSettingsRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *GetGuestSettingsRequest) GetCookie() string {
	if x != nil {
		return x.Cookie
	}
	return ""
}

type GetGuestSettingsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsNotRemindInfo bool `protobuf:"varint,1,opt,name=is_not_remind_info,json=isNotRemindInfo,proto3" json:"is_not_remind_info,omitempty"` ////根据cookie和IP判断，是否首次登陆
}

func (x *GetGuestSettingsReply) Reset() {
	*x = GetGuestSettingsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGuestSettingsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGuestSettingsReply) ProtoMessage() {}

func (x *GetGuestSettingsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGuestSettingsReply.ProtoReflect.Descriptor instead.
func (*GetGuestSettingsReply) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{10}
}

func (x *GetGuestSettingsReply) GetIsNotRemindInfo() bool {
	if x != nil {
		return x.IsNotRemindInfo
	}
	return false
}

type SetNotRemindInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip     string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`         // 根据cookie和IP判断，是否首次登陆
	Cookie string `protobuf:"bytes,2,opt,name=cookie,proto3" json:"cookie,omitempty"` // 根据cookie和IP判断，是否首次登陆
}

func (x *SetNotRemindInfoRequest) Reset() {
	*x = SetNotRemindInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNotRemindInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNotRemindInfoRequest) ProtoMessage() {}

func (x *SetNotRemindInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNotRemindInfoRequest.ProtoReflect.Descriptor instead.
func (*SetNotRemindInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{11}
}

func (x *SetNotRemindInfoRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *SetNotRemindInfoRequest) GetCookie() string {
	if x != nil {
		return x.Cookie
	}
	return ""
}

type SetNotRemindInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetNotRemindInfoReply) Reset() {
	*x = SetNotRemindInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNotRemindInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNotRemindInfoReply) ProtoMessage() {}

func (x *SetNotRemindInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNotRemindInfoReply.ProtoReflect.Descriptor instead.
func (*SetNotRemindInfoReply) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_proto_rawDescGZIP(), []int{12}
}

var File_api_webnavigation_navi_proto protoreflect.FileDescriptor

var file_api_webnavigation_navi_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x61, 0x76, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x34, 0x0a, 0x08, 0x4e, 0x61, 0x76, 0x69, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x27, 0x0a, 0x0f, 0x53, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x76, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x0f, 0x0a,
	0x0d, 0x53, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x42, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x41, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x22, 0x44, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2b, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69,
	0x73, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x41,
	0x0a, 0x17, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69,
	0x65, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xd1, 0x06, 0x0a, 0x04, 0x4e,
	0x61, 0x76, 0x69, 0x12, 0x85, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x76, 0x69, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x76,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77,
	0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65,
	0x62, 0x2d, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x12, 0x85, 0x01, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22,
	0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2d, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6e,
	0x61, 0x76, 0x69, 0x12, 0x7d, 0x0a, 0x08, 0x53, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x12,
	0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x76, 0x69,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a,
	0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2d, 0x6e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61,
	0x76, 0x69, 0x12, 0x7a, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x12, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2d, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x12, 0x9b,
	0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2b, 0x12, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2d, 0x6e, 0x61, 0x76, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x9f, 0x01, 0x0a,
	0x10, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a,
	0x01, 0x2a, 0x22, 0x2a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2d, 0x6e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x5f, 0x6e,
	0x6f, 0x74, 0x5f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x4d,
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x36, 0x68, 0x75, 0x69, 0x2d, 0x77, 0x65, 0x62, 0x70, 0x61,
	0x67, 0x65, 0x2d, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b,
	0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_webnavigation_navi_proto_rawDescOnce sync.Once
	file_api_webnavigation_navi_proto_rawDescData = file_api_webnavigation_navi_proto_rawDesc
)

func file_api_webnavigation_navi_proto_rawDescGZIP() []byte {
	file_api_webnavigation_navi_proto_rawDescOnce.Do(func() {
		file_api_webnavigation_navi_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_webnavigation_navi_proto_rawDescData)
	})
	return file_api_webnavigation_navi_proto_rawDescData
}

var file_api_webnavigation_navi_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_webnavigation_navi_proto_goTypes = []interface{}{
	(*CreateNaviRequest)(nil),       // 0: api.webnavigation.CreateNaviRequest
	(*CreateNaviReply)(nil),         // 1: api.webnavigation.CreateNaviReply
	(*NaviType)(nil),                // 2: api.webnavigation.NaviType
	(*DeleteNaviRequest)(nil),       // 3: api.webnavigation.DeleteNaviRequest
	(*DeleteNaviReply)(nil),         // 4: api.webnavigation.DeleteNaviReply
	(*SortNaviRequest)(nil),         // 5: api.webnavigation.SortNaviRequest
	(*SortNaviReply)(nil),           // 6: api.webnavigation.SortNaviReply
	(*ListNaviRequest)(nil),         // 7: api.webnavigation.ListNaviRequest
	(*ListNaviReply)(nil),           // 8: api.webnavigation.ListNaviReply
	(*GetGuestSettingsRequest)(nil), // 9: api.webnavigation.GetGuestSettingsRequest
	(*GetGuestSettingsReply)(nil),   // 10: api.webnavigation.GetGuestSettingsReply
	(*SetNotRemindInfoRequest)(nil), // 11: api.webnavigation.SetNotRemindInfoRequest
	(*SetNotRemindInfoReply)(nil),   // 12: api.webnavigation.SetNotRemindInfoReply
}
var file_api_webnavigation_navi_proto_depIdxs = []int32{
	2,  // 0: api.webnavigation.ListNaviReply.items:type_name -> api.webnavigation.NaviType
	0,  // 1: api.webnavigation.Navi.CreateNavi:input_type -> api.webnavigation.CreateNaviRequest
	3,  // 2: api.webnavigation.Navi.DeleteNavi:input_type -> api.webnavigation.DeleteNaviRequest
	5,  // 3: api.webnavigation.Navi.SortNavi:input_type -> api.webnavigation.SortNaviRequest
	7,  // 4: api.webnavigation.Navi.ListNavi:input_type -> api.webnavigation.ListNaviRequest
	9,  // 5: api.webnavigation.Navi.GetGuestSettings:input_type -> api.webnavigation.GetGuestSettingsRequest
	11, // 6: api.webnavigation.Navi.SetNotRemindInfo:input_type -> api.webnavigation.SetNotRemindInfoRequest
	1,  // 7: api.webnavigation.Navi.CreateNavi:output_type -> api.webnavigation.CreateNaviReply
	4,  // 8: api.webnavigation.Navi.DeleteNavi:output_type -> api.webnavigation.DeleteNaviReply
	6,  // 9: api.webnavigation.Navi.SortNavi:output_type -> api.webnavigation.SortNaviReply
	8,  // 10: api.webnavigation.Navi.ListNavi:output_type -> api.webnavigation.ListNaviReply
	10, // 11: api.webnavigation.Navi.GetGuestSettings:output_type -> api.webnavigation.GetGuestSettingsReply
	12, // 12: api.webnavigation.Navi.SetNotRemindInfo:output_type -> api.webnavigation.SetNotRemindInfoReply
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_webnavigation_navi_proto_init() }
func file_api_webnavigation_navi_proto_init() {
	if File_api_webnavigation_navi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_webnavigation_navi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNaviRequest); i {
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
		file_api_webnavigation_navi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNaviReply); i {
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
		file_api_webnavigation_navi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NaviType); i {
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
		file_api_webnavigation_navi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNaviRequest); i {
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
		file_api_webnavigation_navi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNaviReply); i {
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
		file_api_webnavigation_navi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortNaviRequest); i {
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
		file_api_webnavigation_navi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortNaviReply); i {
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
		file_api_webnavigation_navi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNaviRequest); i {
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
		file_api_webnavigation_navi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNaviReply); i {
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
		file_api_webnavigation_navi_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGuestSettingsRequest); i {
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
		file_api_webnavigation_navi_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGuestSettingsReply); i {
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
		file_api_webnavigation_navi_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNotRemindInfoRequest); i {
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
		file_api_webnavigation_navi_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNotRemindInfoReply); i {
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
			RawDescriptor: file_api_webnavigation_navi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_webnavigation_navi_proto_goTypes,
		DependencyIndexes: file_api_webnavigation_navi_proto_depIdxs,
		MessageInfos:      file_api_webnavigation_navi_proto_msgTypes,
	}.Build()
	File_api_webnavigation_navi_proto = out.File
	file_api_webnavigation_navi_proto_rawDesc = nil
	file_api_webnavigation_navi_proto_goTypes = nil
	file_api_webnavigation_navi_proto_depIdxs = nil
}