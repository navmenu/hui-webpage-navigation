// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: api/webnavigation/navi_lvl2.proto

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

type CreateNaviLvl2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NaviName string `protobuf:"bytes,1,opt,name=navi_name,json=naviName,proto3" json:"navi_name,omitempty"`  // 分类名称：你要添加在哪个分类里
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                          // 名称：按钮上的简洁文字
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`                          // 中转页面内容部分包含以下元素 文案：后台自定义，比如“这是汇旺担保群，点击直达：”，文案单独一行；
	Link     string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`                          // 中转页面内容部分包含以下元素 链接：后台自定义链接文案，自定义链接的地址；字号比前面一行大两号；
	IsEscrow bool   `protobuf:"varint,5,opt,name=is_escrow,json=isEscrow,proto3" json:"is_escrow,omitempty"` // 需求-添加分类内容时可以勾选担保标志
}

func (x *CreateNaviLvl2Request) Reset() {
	*x = CreateNaviLvl2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNaviLvl2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNaviLvl2Request) ProtoMessage() {}

func (x *CreateNaviLvl2Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNaviLvl2Request.ProtoReflect.Descriptor instead.
func (*CreateNaviLvl2Request) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_lvl2_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNaviLvl2Request) GetNaviName() string {
	if x != nil {
		return x.NaviName
	}
	return ""
}

func (x *CreateNaviLvl2Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNaviLvl2Request) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateNaviLvl2Request) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *CreateNaviLvl2Request) GetIsEscrow() bool {
	if x != nil {
		return x.IsEscrow
	}
	return false
}

type CreateNaviLvl2Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NaviName string `protobuf:"bytes,1,opt,name=navi_name,json=naviName,proto3" json:"navi_name,omitempty"` // 在哪个分类下添加（由于分类不会重命名，也就没必要用ID表示）
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                         // 内容名称
}

func (x *CreateNaviLvl2Reply) Reset() {
	*x = CreateNaviLvl2Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNaviLvl2Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNaviLvl2Reply) ProtoMessage() {}

func (x *CreateNaviLvl2Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNaviLvl2Reply.ProtoReflect.Descriptor instead.
func (*CreateNaviLvl2Reply) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_lvl2_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNaviLvl2Reply) GetNaviName() string {
	if x != nil {
		return x.NaviName
	}
	return ""
}

func (x *CreateNaviLvl2Reply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NaviLvl2Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NaviName string `protobuf:"bytes,1,opt,name=navi_name,json=naviName,proto3" json:"navi_name,omitempty"`  // 分类名称
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                          // 名称：按钮上的简洁文字
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`                          // 中转页面内容部分包含以下元素 文案：后台自定义，比如“这是汇旺担保群，点击直达：”，文案单独一行；
	Link     string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`                          // 中转页面内容部分包含以下元素 链接：后台自定义链接文案，自定义链接的地址；字号比前面一行大两号；
	IsEscrow bool   `protobuf:"varint,5,opt,name=is_escrow,json=isEscrow,proto3" json:"is_escrow,omitempty"` // 需求-添加分类内容时可以勾选担保标志
	Sort     int32  `protobuf:"varint,6,opt,name=sort,proto3" json:"sort,omitempty"`                         // 当前序号
}

func (x *NaviLvl2Type) Reset() {
	*x = NaviLvl2Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NaviLvl2Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NaviLvl2Type) ProtoMessage() {}

func (x *NaviLvl2Type) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NaviLvl2Type.ProtoReflect.Descriptor instead.
func (*NaviLvl2Type) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_lvl2_proto_rawDescGZIP(), []int{2}
}

func (x *NaviLvl2Type) GetNaviName() string {
	if x != nil {
		return x.NaviName
	}
	return ""
}

func (x *NaviLvl2Type) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NaviLvl2Type) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *NaviLvl2Type) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *NaviLvl2Type) GetIsEscrow() bool {
	if x != nil {
		return x.IsEscrow
	}
	return false
}

func (x *NaviLvl2Type) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type DeleteNaviLvl2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NaviName string `protobuf:"bytes,1,opt,name=navi_name,json=naviName,proto3" json:"navi_name,omitempty"` // 分类名称
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                         // 内容名称
}

func (x *DeleteNaviLvl2Request) Reset() {
	*x = DeleteNaviLvl2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNaviLvl2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNaviLvl2Request) ProtoMessage() {}

func (x *DeleteNaviLvl2Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNaviLvl2Request.ProtoReflect.Descriptor instead.
func (*DeleteNaviLvl2Request) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_lvl2_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNaviLvl2Request) GetNaviName() string {
	if x != nil {
		return x.NaviName
	}
	return ""
}

func (x *DeleteNaviLvl2Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteNaviLvl2Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteNaviLvl2Reply) Reset() {
	*x = DeleteNaviLvl2Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNaviLvl2Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNaviLvl2Reply) ProtoMessage() {}

func (x *DeleteNaviLvl2Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNaviLvl2Reply.ProtoReflect.Descriptor instead.
func (*DeleteNaviLvl2Reply) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_lvl2_proto_rawDescGZIP(), []int{4}
}

type SortNaviLvl2Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`  // 名称，不检查名字是否存在
	Sort int32  `protobuf:"varint,2,opt,name=sort,proto3" json:"sort,omitempty"` // 新序号，非负数，不检查序号重复和冲突
}

func (x *SortNaviLvl2Item) Reset() {
	*x = SortNaviLvl2Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortNaviLvl2Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortNaviLvl2Item) ProtoMessage() {}

func (x *SortNaviLvl2Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortNaviLvl2Item.ProtoReflect.Descriptor instead.
func (*SortNaviLvl2Item) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_lvl2_proto_rawDescGZIP(), []int{5}
}

func (x *SortNaviLvl2Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SortNaviLvl2Item) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type SortNaviLvl2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NaviName string              `protobuf:"bytes,1,opt,name=navi_name,json=naviName,proto3" json:"navi_name,omitempty"` // 分类名称
	Items    []*SortNaviLvl2Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`                       // 把新的顺序传过来
}

func (x *SortNaviLvl2Request) Reset() {
	*x = SortNaviLvl2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortNaviLvl2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortNaviLvl2Request) ProtoMessage() {}

func (x *SortNaviLvl2Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortNaviLvl2Request.ProtoReflect.Descriptor instead.
func (*SortNaviLvl2Request) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_lvl2_proto_rawDescGZIP(), []int{6}
}

func (x *SortNaviLvl2Request) GetNaviName() string {
	if x != nil {
		return x.NaviName
	}
	return ""
}

func (x *SortNaviLvl2Request) GetItems() []*SortNaviLvl2Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type SortNaviLvl2Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SortNaviLvl2Reply) Reset() {
	*x = SortNaviLvl2Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortNaviLvl2Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortNaviLvl2Reply) ProtoMessage() {}

func (x *SortNaviLvl2Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortNaviLvl2Reply.ProtoReflect.Descriptor instead.
func (*SortNaviLvl2Reply) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_lvl2_proto_rawDescGZIP(), []int{7}
}

type ListNaviLvl2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NaviName string `protobuf:"bytes,1,opt,name=navi_name,json=naviName,proto3" json:"navi_name,omitempty"` // 分类名称
}

func (x *ListNaviLvl2Request) Reset() {
	*x = ListNaviLvl2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNaviLvl2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNaviLvl2Request) ProtoMessage() {}

func (x *ListNaviLvl2Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNaviLvl2Request.ProtoReflect.Descriptor instead.
func (*ListNaviLvl2Request) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_lvl2_proto_rawDescGZIP(), []int{8}
}

func (x *ListNaviLvl2Request) GetNaviName() string {
	if x != nil {
		return x.NaviName
	}
	return ""
}

type ListNaviLvl2Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*NaviLvl2Type `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // 内容是有序的
}

func (x *ListNaviLvl2Reply) Reset() {
	*x = ListNaviLvl2Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNaviLvl2Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNaviLvl2Reply) ProtoMessage() {}

func (x *ListNaviLvl2Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_webnavigation_navi_lvl2_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNaviLvl2Reply.ProtoReflect.Descriptor instead.
func (*ListNaviLvl2Reply) Descriptor() ([]byte, []int) {
	return file_api_webnavigation_navi_lvl2_proto_rawDescGZIP(), []int{9}
}

func (x *ListNaviLvl2Reply) GetItems() []*NaviLvl2Type {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_api_webnavigation_navi_lvl2_proto protoreflect.FileDescriptor

var file_api_webnavigation_navi_lvl2_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6c, 0x76, 0x6c, 0x32, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x76, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x65, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x22, 0x46, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x61, 0x76, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x98, 0x01, 0x0a, 0x0c, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x76, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x65, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x48, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x76, 0x69, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61,
	0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3a, 0x0a, 0x10, 0x53,
	0x6f, 0x72, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x6d, 0x0a, 0x13, 0x53, 0x6f, 0x72, 0x74, 0x4e,
	0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x61, 0x76, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x6f, 0x72, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x6f, 0x72, 0x74, 0x4e, 0x61,
	0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x32, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x76, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x4a, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xdb, 0x04, 0x0a, 0x08,
	0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x12, 0x96, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x12, 0x28, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e,
	0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x32, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77,
	0x65, 0x62, 0x2d, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6c, 0x76, 0x6c,
	0x32, 0x12, 0x96, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69,
	0x4c, 0x76, 0x6c, 0x32, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61,
	0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c,
	0x32, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01,
	0x2a, 0x22, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2d, 0x6e, 0x61, 0x76, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6c, 0x76, 0x6c, 0x32, 0x12, 0x8e, 0x01, 0x0a, 0x0c, 0x53,
	0x6f, 0x72, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x12, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x76, 0x69,
	0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2d, 0x6e,
	0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x72,
	0x74, 0x5f, 0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6c, 0x76, 0x6c, 0x32, 0x12, 0x8b, 0x01, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x12, 0x26, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61,
	0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x76,
	0x69, 0x4c, 0x76, 0x6c, 0x32, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x27, 0x12, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2d, 0x6e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x76, 0x69, 0x5f, 0x6c, 0x76, 0x6c, 0x32, 0x42, 0x4d, 0x0a, 0x11, 0x61, 0x70, 0x69,
	0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x01,
	0x5a, 0x36, 0x68, 0x75, 0x69, 0x2d, 0x77, 0x65, 0x62, 0x70, 0x61, 0x67, 0x65, 0x2d, 0x6e, 0x61,
	0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62,
	0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x77, 0x65, 0x62, 0x6e, 0x61,
	0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_webnavigation_navi_lvl2_proto_rawDescOnce sync.Once
	file_api_webnavigation_navi_lvl2_proto_rawDescData = file_api_webnavigation_navi_lvl2_proto_rawDesc
)

func file_api_webnavigation_navi_lvl2_proto_rawDescGZIP() []byte {
	file_api_webnavigation_navi_lvl2_proto_rawDescOnce.Do(func() {
		file_api_webnavigation_navi_lvl2_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_webnavigation_navi_lvl2_proto_rawDescData)
	})
	return file_api_webnavigation_navi_lvl2_proto_rawDescData
}

var file_api_webnavigation_navi_lvl2_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_webnavigation_navi_lvl2_proto_goTypes = []interface{}{
	(*CreateNaviLvl2Request)(nil), // 0: api.webnavigation.CreateNaviLvl2Request
	(*CreateNaviLvl2Reply)(nil),   // 1: api.webnavigation.CreateNaviLvl2Reply
	(*NaviLvl2Type)(nil),          // 2: api.webnavigation.NaviLvl2Type
	(*DeleteNaviLvl2Request)(nil), // 3: api.webnavigation.DeleteNaviLvl2Request
	(*DeleteNaviLvl2Reply)(nil),   // 4: api.webnavigation.DeleteNaviLvl2Reply
	(*SortNaviLvl2Item)(nil),      // 5: api.webnavigation.SortNaviLvl2Item
	(*SortNaviLvl2Request)(nil),   // 6: api.webnavigation.SortNaviLvl2Request
	(*SortNaviLvl2Reply)(nil),     // 7: api.webnavigation.SortNaviLvl2Reply
	(*ListNaviLvl2Request)(nil),   // 8: api.webnavigation.ListNaviLvl2Request
	(*ListNaviLvl2Reply)(nil),     // 9: api.webnavigation.ListNaviLvl2Reply
}
var file_api_webnavigation_navi_lvl2_proto_depIdxs = []int32{
	5, // 0: api.webnavigation.SortNaviLvl2Request.items:type_name -> api.webnavigation.SortNaviLvl2Item
	2, // 1: api.webnavigation.ListNaviLvl2Reply.items:type_name -> api.webnavigation.NaviLvl2Type
	0, // 2: api.webnavigation.NaviLvl2.CreateNaviLvl2:input_type -> api.webnavigation.CreateNaviLvl2Request
	3, // 3: api.webnavigation.NaviLvl2.DeleteNaviLvl2:input_type -> api.webnavigation.DeleteNaviLvl2Request
	6, // 4: api.webnavigation.NaviLvl2.SortNaviLvl2:input_type -> api.webnavigation.SortNaviLvl2Request
	8, // 5: api.webnavigation.NaviLvl2.ListNaviLvl2:input_type -> api.webnavigation.ListNaviLvl2Request
	1, // 6: api.webnavigation.NaviLvl2.CreateNaviLvl2:output_type -> api.webnavigation.CreateNaviLvl2Reply
	4, // 7: api.webnavigation.NaviLvl2.DeleteNaviLvl2:output_type -> api.webnavigation.DeleteNaviLvl2Reply
	7, // 8: api.webnavigation.NaviLvl2.SortNaviLvl2:output_type -> api.webnavigation.SortNaviLvl2Reply
	9, // 9: api.webnavigation.NaviLvl2.ListNaviLvl2:output_type -> api.webnavigation.ListNaviLvl2Reply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_webnavigation_navi_lvl2_proto_init() }
func file_api_webnavigation_navi_lvl2_proto_init() {
	if File_api_webnavigation_navi_lvl2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_webnavigation_navi_lvl2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNaviLvl2Request); i {
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
		file_api_webnavigation_navi_lvl2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNaviLvl2Reply); i {
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
		file_api_webnavigation_navi_lvl2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NaviLvl2Type); i {
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
		file_api_webnavigation_navi_lvl2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNaviLvl2Request); i {
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
		file_api_webnavigation_navi_lvl2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNaviLvl2Reply); i {
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
		file_api_webnavigation_navi_lvl2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortNaviLvl2Item); i {
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
		file_api_webnavigation_navi_lvl2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortNaviLvl2Request); i {
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
		file_api_webnavigation_navi_lvl2_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortNaviLvl2Reply); i {
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
		file_api_webnavigation_navi_lvl2_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNaviLvl2Request); i {
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
		file_api_webnavigation_navi_lvl2_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNaviLvl2Reply); i {
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
			RawDescriptor: file_api_webnavigation_navi_lvl2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_webnavigation_navi_lvl2_proto_goTypes,
		DependencyIndexes: file_api_webnavigation_navi_lvl2_proto_depIdxs,
		MessageInfos:      file_api_webnavigation_navi_lvl2_proto_msgTypes,
	}.Build()
	File_api_webnavigation_navi_lvl2_proto = out.File
	file_api_webnavigation_navi_lvl2_proto_rawDesc = nil
	file_api_webnavigation_navi_lvl2_proto_goTypes = nil
	file_api_webnavigation_navi_lvl2_proto_depIdxs = nil
}
