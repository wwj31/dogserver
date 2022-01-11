// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: error.proto

package message

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
	ERROR_SUCCESS                   ERROR = 0  //执行成功
	ERROR_FAILED                    ERROR = 1  //协议执行失败，原因模糊
	ERROR_SECURITYCODE_CHECK_FAILED ERROR = 2  //安全码校验失败
	ERROR_ITEM_NOT_ENOUGH           ERROR = 3  //道具不足
	ERROR_GOLD_NOT_ENOUGH           ERROR = 4  //金币不足
	ERROR_LEVEL_NOT_ENGOUTH         ERROR = 5  //等级不足
	ERROR_MAIL_REPEAT_RECV_ITEM     ERROR = 6  //邮件道具重复领取
	ERROR_CLIENT_WRONG_PARAM        ERROR = 9  //客户端错误参数
	ERROR_CFG_NO_THIS_PARAM         ERROR = 10 //配置表错误
	ERROR_NAME_LEN_OUTRANGE         ERROR = 13 //命名超过长度限制
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0:  "SUCCESS",
		1:  "FAILED",
		2:  "SECURITYCODE_CHECK_FAILED",
		3:  "ITEM_NOT_ENOUGH",
		4:  "GOLD_NOT_ENOUGH",
		5:  "LEVEL_NOT_ENGOUTH",
		6:  "MAIL_REPEAT_RECV_ITEM",
		9:  "CLIENT_WRONG_PARAM",
		10: "CFG_NO_THIS_PARAM",
		13: "NAME_LEN_OUTRANGE",
	}
	ERROR_value = map[string]int32{
		"SUCCESS":                   0,
		"FAILED":                    1,
		"SECURITYCODE_CHECK_FAILED": 2,
		"ITEM_NOT_ENOUGH":           3,
		"GOLD_NOT_ENOUGH":           4,
		"LEVEL_NOT_ENGOUTH":         5,
		"MAIL_REPEAT_RECV_ITEM":     6,
		"CLIENT_WRONG_PARAM":        9,
		"CFG_NO_THIS_PARAM":         10,
		"NAME_LEN_OUTRANGE":         13,
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

var File_error_proto protoreflect.FileDescriptor

var file_error_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0xe1, 0x01, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x45, 0x43,
	0x55, 0x52, 0x49, 0x54, 0x59, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x54, 0x45, 0x4d,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x10, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x47, 0x4f, 0x4c, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48,
	0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x45, 0x4e, 0x47, 0x4f, 0x55, 0x54, 0x48, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x41, 0x49,
	0x4c, 0x5f, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x5f, 0x52, 0x45, 0x43, 0x56, 0x5f, 0x49, 0x54,
	0x45, 0x4d, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x57,
	0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11,
	0x43, 0x46, 0x47, 0x5f, 0x4e, 0x4f, 0x5f, 0x54, 0x48, 0x49, 0x53, 0x5f, 0x50, 0x41, 0x52, 0x41,
	0x4d, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4c, 0x45, 0x4e, 0x5f,
	0x4f, 0x55, 0x54, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x0d, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_error_proto_goTypes = []interface{}{
	(ERROR)(0), // 0: message.ERROR
}
var file_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_proto_init() }
func file_error_proto_init() {
	if File_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_proto_goTypes,
		DependencyIndexes: file_error_proto_depIdxs,
		EnumInfos:         file_error_proto_enumTypes,
	}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}
