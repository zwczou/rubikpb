// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: enum.proto

package rbkpb

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

type Error_Code int32

const (
	Error_ERR_OK                 Error_Code = 0
	Error_ERR_NOT_FOUND          Error_Code = 1
	Error_ERR_METHOD_NOT_ALLOWED Error_Code = 2
	Error_ERR_INTERNAL_SERVER    Error_Code = 3
	Error_ERR_UNAUTHORIZED       Error_Code = 4
	Error_ERR_FORBIDDEN          Error_Code = 5
	Error_ERR_RATE_LIMIT         Error_Code = 6
	Error_ERR_ILLEGAL_ACCESS     Error_Code = 7
	Error_ERR_DATABASE           Error_Code = 8
	Error_ERR_INVALID_ARGS       Error_Code = 9
	Error_ERR_INVALID_METHOD     Error_Code = 10
)

// Enum value maps for Error_Code.
var (
	Error_Code_name = map[int32]string{
		0:  "ERR_OK",
		1:  "ERR_NOT_FOUND",
		2:  "ERR_METHOD_NOT_ALLOWED",
		3:  "ERR_INTERNAL_SERVER",
		4:  "ERR_UNAUTHORIZED",
		5:  "ERR_FORBIDDEN",
		6:  "ERR_RATE_LIMIT",
		7:  "ERR_ILLEGAL_ACCESS",
		8:  "ERR_DATABASE",
		9:  "ERR_INVALID_ARGS",
		10: "ERR_INVALID_METHOD",
	}
	Error_Code_value = map[string]int32{
		"ERR_OK":                 0,
		"ERR_NOT_FOUND":          1,
		"ERR_METHOD_NOT_ALLOWED": 2,
		"ERR_INTERNAL_SERVER":    3,
		"ERR_UNAUTHORIZED":       4,
		"ERR_FORBIDDEN":          5,
		"ERR_RATE_LIMIT":         6,
		"ERR_ILLEGAL_ACCESS":     7,
		"ERR_DATABASE":           8,
		"ERR_INVALID_ARGS":       9,
		"ERR_INVALID_METHOD":     10,
	}
)

func (x Error_Code) Enum() *Error_Code {
	p := new(Error_Code)
	*p = x
	return p
}

func (x Error_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[0].Descriptor()
}

func (Error_Code) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[0]
}

func (x Error_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_Code.Descriptor instead.
func (Error_Code) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{0, 0}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=rbkpb.Error_Code" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_enum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() Error_Code {
	if x != nil {
		return x.Code
	}
	return Error_ERR_OK
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_enum_proto protoreflect.FileDescriptor

var file_enum_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x62,
	0x6b, 0x70, 0x62, 0x22, 0xba, 0x02, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x72, 0x62,
	0x6b, 0x70, 0x62, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xef,
	0x01, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x52, 0x52, 0x5f, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x5f, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e,
	0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x52, 0x52, 0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10,
	0x04, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44,
	0x45, 0x4e, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x5f, 0x52, 0x41, 0x54, 0x45,
	0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f,
	0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x07,
	0x12, 0x10, 0x0a, 0x0c, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45,
	0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x41, 0x52, 0x47, 0x53, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x10, 0x0a,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a,
	0x77, 0x63, 0x7a, 0x6f, 0x75, 0x2f, 0x72, 0x75, 0x62, 0x69, 0x6b, 0x70, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x62, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_enum_proto_rawDescOnce sync.Once
	file_enum_proto_rawDescData = file_enum_proto_rawDesc
)

func file_enum_proto_rawDescGZIP() []byte {
	file_enum_proto_rawDescOnce.Do(func() {
		file_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_enum_proto_rawDescData)
	})
	return file_enum_proto_rawDescData
}

var file_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_enum_proto_goTypes = []interface{}{
	(Error_Code)(0), // 0: rbkpb.Error.Code
	(*Error)(nil),   // 1: rbkpb.Error
}
var file_enum_proto_depIdxs = []int32{
	0, // 0: rbkpb.Error.code:type_name -> rbkpb.Error.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_enum_proto_init() }
func file_enum_proto_init() {
	if File_enum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_proto_goTypes,
		DependencyIndexes: file_enum_proto_depIdxs,
		EnumInfos:         file_enum_proto_enumTypes,
		MessageInfos:      file_enum_proto_msgTypes,
	}.Build()
	File_enum_proto = out.File
	file_enum_proto_rawDesc = nil
	file_enum_proto_goTypes = nil
	file_enum_proto_depIdxs = nil
}
