// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.0
// source: error.proto

package outer

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

type ERROR int32

const (
	ERROR_FAILED                     ERROR = 0 //协议执行失败，原因模糊
	ERROR_SECURITY_CODE_CHECK_FAILED ERROR = 1 //安全码校验失败
	ERROR_LOGIN_TOKEN_INVALID        ERROR = 2 //登录token过期
	ERROR_REPEAT_LOGIN               ERROR = 3 //被顶号
	ERROR_GOLD_NOT_ENOUGH            ERROR = 4 //金币不足
	ERROR_INVALID_PHONE              ERROR = 5 //无效的电话
	ERROR_SMS_SEND_FAILED            ERROR = 6 //sms校验失败
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0: "FAILED",
		1: "SECURITY_CODE_CHECK_FAILED",
		2: "LOGIN_TOKEN_INVALID",
		3: "REPEAT_LOGIN",
		4: "GOLD_NOT_ENOUGH",
		5: "INVALID_PHONE",
		6: "SMS_SEND_FAILED",
	}
	ERROR_value = map[string]int32{
		"FAILED":                     0,
		"SECURITY_CODE_CHECK_FAILED": 1,
		"LOGIN_TOKEN_INVALID":        2,
		"REPEAT_LOGIN":               3,
		"GOLD_NOT_ENOUGH":            4,
		"INVALID_PHONE":              5,
		"SMS_SEND_FAILED":            6,
	}
)

func (x ERROR) Enum() *ERROR {
	p := new(ERROR)
	*p = x
	return p
}

func (x ERROR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERROR) Descriptor() protoreflect.EnumDescriptor {
	return file_error_proto_enumTypes[0].Descriptor()
}

func (ERROR) Type() protoreflect.EnumType {
	return &file_error_proto_enumTypes[0]
}

func (x ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERROR.Descriptor instead.
func (ERROR) EnumDescriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

type FailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error ERROR  `protobuf:"varint,1,opt,name=Error,proto3,enum=outer.ERROR" json:"Error,omitempty"`
	Info  string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *FailRsp) Reset() {
	*x = FailRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailRsp) ProtoMessage() {}

func (x *FailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailRsp.ProtoReflect.Descriptor instead.
func (*FailRsp) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

func (x *FailRsp) GetError() ERROR {
	if x != nil {
		return x.Error
	}
	return ERROR_FAILED
}

func (x *FailRsp) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type Unknown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Unknown) Reset() {
	*x = Unknown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unknown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unknown) ProtoMessage() {}

func (x *Unknown) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unknown.ProtoReflect.Descriptor instead.
func (*Unknown) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{1}
}

var File_error_proto protoreflect.FileDescriptor

var file_error_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x12,
	0x22, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x52, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x09, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x2a, 0x9b, 0x01, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x45, 0x43, 0x55,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x4f, 0x47, 0x49,
	0x4e, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x02, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x49,
	0x4e, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x4f, 0x4c, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x4d, 0x53, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x06,
	0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_error_proto_rawDescOnce sync.Once
	file_error_proto_rawDescData = file_error_proto_rawDesc
)

func file_error_proto_rawDescGZIP() []byte {
	file_error_proto_rawDescOnce.Do(func() {
		file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_proto_rawDescData)
	})
	return file_error_proto_rawDescData
}

var file_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_error_proto_goTypes = []interface{}{
	(ERROR)(0),      // 0: outer.ERROR
	(*FailRsp)(nil), // 1: outer.FailRsp
	(*Unknown)(nil), // 2: outer.Unknown
}
var file_error_proto_depIdxs = []int32{
	0, // 0: outer.FailRsp.Error:type_name -> outer.ERROR
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_error_proto_init() }
func file_error_proto_init() {
	if File_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailRsp); i {
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
		file_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unknown); i {
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
			RawDescriptor: file_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_proto_goTypes,
		DependencyIndexes: file_error_proto_depIdxs,
		EnumInfos:         file_error_proto_enumTypes,
		MessageInfos:      file_error_proto_msgTypes,
	}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}
