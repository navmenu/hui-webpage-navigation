// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/webnavigation/errorreason.proto

package webnavigation

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_UNKNOWN        ErrorReason = 0
	ErrorReason_MISSING_PARAM  ErrorReason = 10000
	ErrorReason_BAD_PARAM      ErrorReason = 10001
	ErrorReason_USER_NOT_FOUND ErrorReason = 10002
	ErrorReason_ALREADY_EXIST  ErrorReason = 10003
	ErrorReason_NOT_EXIST      ErrorReason = 10004
	ErrorReason_DB_ERROR       ErrorReason = 10005
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:     "UNKNOWN",
		10000: "MISSING_PARAM",
		10001: "BAD_PARAM",
		10002: "USER_NOT_FOUND",
		10003: "ALREADY_EXIST",
		10004: "NOT_EXIST",
		10005: "DB_ERROR",
	}
	ErrorReason_value = map[string]int32{
		"UNKNOWN":        0,
		"MISSING_PARAM":  10000,
		"BAD_PARAM":      10001,
		"USER_NOT_FOUND": 10002,
		"ALREADY_EXIST":  10003,
		"NOT_EXIST":      10004,
		"DB_ERROR":       10005,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_webnavigation_errorreason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_webnavigation_errorreason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_webnavigation_errorreason_proto_rawDescGZIP(), []int{0}
}

var File_api_webnavigation_errorreason_proto protoreflect.FileDescriptor

var file_api_webnavigation_errorreason_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65, 0x62, 0x6e, 0x61,
	0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb6, 0x01,
	0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03,
	0x12, 0x18, 0x0a, 0x0d, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x41,
	0x4d, 0x10, 0x90, 0x4e, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x14, 0x0a, 0x09, 0x42, 0x41,
	0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0x91, 0x4e, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03,
	0x12, 0x19, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x92, 0x4e, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x18, 0x0a, 0x0d, 0x41,
	0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x93, 0x4e, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x14, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49,
	0x53, 0x54, 0x10, 0x94, 0x4e, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x13, 0x0a, 0x08, 0x44,
	0x42, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x95, 0x4e, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03,
	0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x4d, 0x0a, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x65,
	0x62, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x36, 0x68,
	0x75, 0x69, 0x2d, 0x77, 0x65, 0x62, 0x70, 0x61, 0x67, 0x65, 0x2d, 0x6e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76,
	0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x77, 0x65, 0x62, 0x6e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_webnavigation_errorreason_proto_rawDescOnce sync.Once
	file_api_webnavigation_errorreason_proto_rawDescData = file_api_webnavigation_errorreason_proto_rawDesc
)

func file_api_webnavigation_errorreason_proto_rawDescGZIP() []byte {
	file_api_webnavigation_errorreason_proto_rawDescOnce.Do(func() {
		file_api_webnavigation_errorreason_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_webnavigation_errorreason_proto_rawDescData)
	})
	return file_api_webnavigation_errorreason_proto_rawDescData
}

var file_api_webnavigation_errorreason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_webnavigation_errorreason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: api.webnavigation.ErrorReason
}
var file_api_webnavigation_errorreason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_webnavigation_errorreason_proto_init() }
func file_api_webnavigation_errorreason_proto_init() {
	if File_api_webnavigation_errorreason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_webnavigation_errorreason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_webnavigation_errorreason_proto_goTypes,
		DependencyIndexes: file_api_webnavigation_errorreason_proto_depIdxs,
		EnumInfos:         file_api_webnavigation_errorreason_proto_enumTypes,
	}.Build()
	File_api_webnavigation_errorreason_proto = out.File
	file_api_webnavigation_errorreason_proto_rawDesc = nil
	file_api_webnavigation_errorreason_proto_goTypes = nil
	file_api_webnavigation_errorreason_proto_depIdxs = nil
}
