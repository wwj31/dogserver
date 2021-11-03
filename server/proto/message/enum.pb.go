// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0--rc2
// source: enum.proto

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

//前后端通信协议 id段 100000-1000000
type MSG int32

const (
	MSG_UNKNOWN MSG = 0
	MSG_PING    MSG = 100001
	MSG_PONG    MSG = 100002
	//--------------------------------------------------------------- login proto 200001-300000
	// 消息段begin
	MSG_LOGIN_SEGMENT_BEGIN MSG = 200001
	// 请求登录
	MSG_LOGIN_REQ MSG = 200002 // LoginReq
	MSG_LOGIN_RSP MSG = 200003 // LoginRsp
	// 消息段end
	MSG_LOGIN_SEGMENT_END MSG = 300000
	//--------------------------------------------------------------- game proto 300001-310000
	// 消息段begin
	MSG_GAME_SEGMENT_BEGIN MSG = 300001
	// 请求登录
	MSG_ENTER_GAME_REQ MSG = 300010 // EnterGameReq
	MSG_ENTER_GAME_RSP MSG = 300011 // EnterGameRsp
	// 消息段end
	MSG_GAME_SEGMENT_END MSG = 310000
)

// Enum value maps for MSG.
var (
	MSG_name = map[int32]string{
		0:      "UNKNOWN",
		100001: "PING",
		100002: "PONG",
		200001: "LOGIN_SEGMENT_BEGIN",
		200002: "LOGIN_REQ",
		200003: "LOGIN_RSP",
		300000: "LOGIN_SEGMENT_END",
		300001: "GAME_SEGMENT_BEGIN",
		300010: "ENTER_GAME_REQ",
		300011: "ENTER_GAME_RSP",
		310000: "GAME_SEGMENT_END",
	}
	MSG_value = map[string]int32{
		"UNKNOWN":             0,
		"PING":                100001,
		"PONG":                100002,
		"LOGIN_SEGMENT_BEGIN": 200001,
		"LOGIN_REQ":           200002,
		"LOGIN_RSP":           200003,
		"LOGIN_SEGMENT_END":   300000,
		"GAME_SEGMENT_BEGIN":  300001,
		"ENTER_GAME_REQ":      300010,
		"ENTER_GAME_RSP":      300011,
		"GAME_SEGMENT_END":    310000,
	}
)

func (x MSG) Enum() *MSG {
	p := new(MSG)
	*p = x
	return p
}

func (x MSG) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSG) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[0].Descriptor()
}

func (MSG) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[0]
}

func (x MSG) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG.Descriptor instead.
func (MSG) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{0}
}

var File_enum_proto protoreflect.FileDescriptor

var file_enum_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0xde, 0x01, 0x0a, 0x03, 0x4d, 0x53, 0x47, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x04, 0x50, 0x49,
	0x4e, 0x47, 0x10, 0xa1, 0x8d, 0x06, 0x12, 0x0a, 0x0a, 0x04, 0x50, 0x4f, 0x4e, 0x47, 0x10, 0xa2,
	0x8d, 0x06, 0x12, 0x19, 0x0a, 0x13, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x53, 0x45, 0x47, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x42, 0x45, 0x47, 0x49, 0x4e, 0x10, 0xc1, 0x9a, 0x0c, 0x12, 0x0f, 0x0a,
	0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xc2, 0x9a, 0x0c, 0x12, 0x0f,
	0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x53, 0x50, 0x10, 0xc3, 0x9a, 0x0c, 0x12,
	0x17, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x45, 0x4e, 0x44, 0x10, 0xe0, 0xa7, 0x12, 0x12, 0x18, 0x0a, 0x12, 0x47, 0x41, 0x4d, 0x45,
	0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x45, 0x47, 0x49, 0x4e, 0x10, 0xe1,
	0xa7, 0x12, 0x12, 0x14, 0x0a, 0x0e, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x47, 0x41, 0x4d, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x10, 0xea, 0xa7, 0x12, 0x12, 0x14, 0x0a, 0x0e, 0x45, 0x4e, 0x54, 0x45,
	0x52, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x53, 0x50, 0x10, 0xeb, 0xa7, 0x12, 0x12, 0x16,
	0x0a, 0x10, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x45,
	0x4e, 0x44, 0x10, 0xf0, 0xf5, 0x12, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_enum_proto_goTypes = []interface{}{
	(MSG)(0), // 0: message.MSG
}
var file_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enum_proto_init() }
func file_enum_proto_init() {
	if File_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_proto_goTypes,
		DependencyIndexes: file_enum_proto_depIdxs,
		EnumInfos:         file_enum_proto_enumTypes,
	}.Build()
	File_enum_proto = out.File
	file_enum_proto_rawDesc = nil
	file_enum_proto_goTypes = nil
	file_enum_proto_depIdxs = nil
}
