// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: msgtype.proto

package message

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//前后端通信协议 id段 100000-1000000
type MSG int32

const (
	MSG_ERROR MSG = 0      // TODO
	MSG_PING  MSG = 100001 // Ping
	MSG_PONG  MSG = 100002 // Pong
	//--------------------------------------------------------------- login proto 200001-300000
	// 消息段begin
	MSG_LOGIN_SEGMENT_BEGIN MSG = 200001
	// 请求登录
	MSG_LOGIN_REQ MSG = 200002 // LoginReq
	MSG_LOGIN_RES MSG = 200003 // LoginRes
	// 消息段end
	MSG_LOGIN_SEGMENT_END MSG = 300000
	//--------------------------------------------------------------- game proto 300001-310000
	// 消息段begin
	MSG_GAME_SEGMENT_BEGIN MSG = 300001
	//进入游戏
	MSG_ENTER_GAME_REQ MSG = 300002 // EnterGameReq
	MSG_ENTER_GAME_RES MSG = 300003 // EnterGameRes
	MSG_GM_COMMAND_REQ MSG = 300205 // GmCommandReq
	MSG_GM_COMMAND_RES MSG = 300206 // GmCommandRes
	// 消息段end
	MSG_GAME_SEGMENT_END MSG = 310000
	//--------------------------------------------------------------- world proto 320000-350000
	// 消息段begin
	MSG_WORLD_SEGMENT_BEGIN  MSG = 320000
	MSG_WORLD_WATCH_POSITION MSG = 320001 // WorldWatchPosition
	MSG_WORLD_UPDATE_ENTITY  MSG = 320002 // WorldUpdateEntity
	// 操作某个实体移动
	MSG_WORLD_OPERATE_ENTITY MSG = 330001 // WorldOperateEntity
	// 消息段end
	MSG_WORLD_SEGMENT_END MSG = 350000
)

// Enum value maps for MSG.
var (
	MSG_name = map[int32]string{
		0:      "ERROR",
		100001: "PING",
		100002: "PONG",
		200001: "LOGIN_SEGMENT_BEGIN",
		200002: "LOGIN_REQ",
		200003: "LOGIN_RES",
		300000: "LOGIN_SEGMENT_END",
		300001: "GAME_SEGMENT_BEGIN",
		300002: "ENTER_GAME_REQ",
		300003: "ENTER_GAME_RES",
		300205: "GM_COMMAND_REQ",
		300206: "GM_COMMAND_RES",
		310000: "GAME_SEGMENT_END",
		320000: "WORLD_SEGMENT_BEGIN",
		320001: "WORLD_WATCH_POSITION",
		320002: "WORLD_UPDATE_ENTITY",
		330001: "WORLD_OPERATE_ENTITY",
		350000: "WORLD_SEGMENT_END",
	}
	MSG_value = map[string]int32{
		"ERROR":                0,
		"PING":                 100001,
		"PONG":                 100002,
		"LOGIN_SEGMENT_BEGIN":  200001,
		"LOGIN_REQ":            200002,
		"LOGIN_RES":            200003,
		"LOGIN_SEGMENT_END":    300000,
		"GAME_SEGMENT_BEGIN":   300001,
		"ENTER_GAME_REQ":       300002,
		"ENTER_GAME_RES":       300003,
		"GM_COMMAND_REQ":       300205,
		"GM_COMMAND_RES":       300206,
		"GAME_SEGMENT_END":     310000,
		"WORLD_SEGMENT_BEGIN":  320000,
		"WORLD_WATCH_POSITION": 320001,
		"WORLD_UPDATE_ENTITY":  320002,
		"WORLD_OPERATE_ENTITY": 330001,
		"WORLD_SEGMENT_END":    350000,
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
	return file_msgtype_proto_enumTypes[0].Descriptor()
}

func (MSG) Type() protoreflect.EnumType {
	return &file_msgtype_proto_enumTypes[0]
}

func (x MSG) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG.Descriptor instead.
func (MSG) EnumDescriptor() ([]byte, []int) {
	return file_msgtype_proto_rawDescGZIP(), []int{0}
}

type MSG_RESULT int32

const (
	MSG_RESULT_SUCCESS                   MSG_RESULT = 0  //执行成功
	MSG_RESULT_FAILED                    MSG_RESULT = 1  //协议执行失败，原因模糊
	MSG_RESULT_SECURITYCODE_CHECK_FAILED MSG_RESULT = 2  //安全码校验失败
	MSG_RESULT_ITEM_NOT_ENOUGH           MSG_RESULT = 3  //道具不足
	MSG_RESULT_GOLD_NOT_ENOUGH           MSG_RESULT = 4  //金币不足
	MSG_RESULT_LEVEL_NOT_ENGOUTH         MSG_RESULT = 5  //等级不足
	MSG_RESULT_CLIENT_WRONG_PARAM        MSG_RESULT = 9  //客户端错误参数
	MSG_RESULT_CFG_NO_THIS_PARAM         MSG_RESULT = 10 //配置表错误
	MSG_RESULT_NAME_LEN_OUTRANGE         MSG_RESULT = 13 //命名超过长度限制
)

// Enum value maps for MSG_RESULT.
var (
	MSG_RESULT_name = map[int32]string{
		0:  "SUCCESS",
		1:  "FAILED",
		2:  "SECURITYCODE_CHECK_FAILED",
		3:  "ITEM_NOT_ENOUGH",
		4:  "GOLD_NOT_ENOUGH",
		5:  "LEVEL_NOT_ENGOUTH",
		9:  "CLIENT_WRONG_PARAM",
		10: "CFG_NO_THIS_PARAM",
		13: "NAME_LEN_OUTRANGE",
	}
	MSG_RESULT_value = map[string]int32{
		"SUCCESS":                   0,
		"FAILED":                    1,
		"SECURITYCODE_CHECK_FAILED": 2,
		"ITEM_NOT_ENOUGH":           3,
		"GOLD_NOT_ENOUGH":           4,
		"LEVEL_NOT_ENGOUTH":         5,
		"CLIENT_WRONG_PARAM":        9,
		"CFG_NO_THIS_PARAM":         10,
		"NAME_LEN_OUTRANGE":         13,
	}
)

func (x MSG_RESULT) Enum() *MSG_RESULT {
	p := new(MSG_RESULT)
	*p = x
	return p
}

func (x MSG_RESULT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSG_RESULT) Descriptor() protoreflect.EnumDescriptor {
	return file_msgtype_proto_enumTypes[1].Descriptor()
}

func (MSG_RESULT) Type() protoreflect.EnumType {
	return &file_msgtype_proto_enumTypes[1]
}

func (x MSG_RESULT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG_RESULT.Descriptor instead.
func (MSG_RESULT) EnumDescriptor() ([]byte, []int) {
	return file_msgtype_proto_rawDescGZIP(), []int{1}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgtype_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_msgtype_proto_msgTypes[0]
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
	return file_msgtype_proto_rawDescGZIP(), []int{0}
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientTimestamp int64 `protobuf:"varint,1,opt,name=ClientTimestamp,json=clientTimestamp,proto3" json:"ClientTimestamp"` // 毫秒时间戳
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgtype_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_msgtype_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_msgtype_proto_rawDescGZIP(), []int{1}
}

func (x *Ping) GetClientTimestamp() int64 {
	if x != nil {
		return x.ClientTimestamp
	}
	return 0
}

//[Sync(SyncPingResult)]
type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientTimestamp int64 `protobuf:"varint,1,opt,name=ClientTimestamp,json=clientTimestamp,proto3" json:"ClientTimestamp"` //
	ServerTimestamp int64 `protobuf:"varint,2,opt,name=ServerTimestamp,json=serverTimestamp,proto3" json:"ServerTimestamp"` // 服务器本地时间 毫秒时间戳
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msgtype_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_msgtype_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_msgtype_proto_rawDescGZIP(), []int{2}
}

func (x *Pong) GetClientTimestamp() int64 {
	if x != nil {
		return x.ClientTimestamp
	}
	return 0
}

func (x *Pong) GetServerTimestamp() int64 {
	if x != nil {
		return x.ServerTimestamp
	}
	return 0
}

var File_msgtype_proto protoreflect.FileDescriptor

var file_msgtype_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x30, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x5a, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a,
	0x8f, 0x03, 0x0a, 0x03, 0x4d, 0x53, 0x47, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10, 0xa1, 0x8d, 0x06, 0x12, 0x0a,
	0x0a, 0x04, 0x50, 0x4f, 0x4e, 0x47, 0x10, 0xa2, 0x8d, 0x06, 0x12, 0x19, 0x0a, 0x13, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x45, 0x47, 0x49,
	0x4e, 0x10, 0xc1, 0x9a, 0x0c, 0x12, 0x0f, 0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52,
	0x45, 0x51, 0x10, 0xc2, 0x9a, 0x0c, 0x12, 0x0f, 0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f,
	0x52, 0x45, 0x53, 0x10, 0xc3, 0x9a, 0x0c, 0x12, 0x17, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x49, 0x4e,
	0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0xe0, 0xa7, 0x12,
	0x12, 0x18, 0x0a, 0x12, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x42, 0x45, 0x47, 0x49, 0x4e, 0x10, 0xe1, 0xa7, 0x12, 0x12, 0x14, 0x0a, 0x0e, 0x45, 0x4e,
	0x54, 0x45, 0x52, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xe2, 0xa7, 0x12,
	0x12, 0x14, 0x0a, 0x0e, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x52,
	0x45, 0x53, 0x10, 0xe3, 0xa7, 0x12, 0x12, 0x14, 0x0a, 0x0e, 0x47, 0x4d, 0x5f, 0x43, 0x4f, 0x4d,
	0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xad, 0xa9, 0x12, 0x12, 0x14, 0x0a, 0x0e,
	0x47, 0x4d, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x10, 0xae,
	0xa9, 0x12, 0x12, 0x16, 0x0a, 0x10, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0xf0, 0xf5, 0x12, 0x12, 0x19, 0x0a, 0x13, 0x57, 0x4f,
	0x52, 0x4c, 0x44, 0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x45, 0x47, 0x49,
	0x4e, 0x10, 0x80, 0xc4, 0x13, 0x12, 0x1a, 0x0a, 0x14, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x5f, 0x57,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x81, 0xc4,
	0x13, 0x12, 0x19, 0x0a, 0x13, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x82, 0xc4, 0x13, 0x12, 0x1a, 0x0a, 0x14,
	0x57, 0x4f, 0x52, 0x4c, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x10, 0x91, 0x92, 0x14, 0x12, 0x17, 0x0a, 0x11, 0x57, 0x4f, 0x52, 0x4c,
	0x44, 0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0xb0, 0xae,
	0x15, 0x2a, 0xcb, 0x01, 0x0a, 0x0a, 0x4d, 0x53, 0x47, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x45, 0x43,
	0x55, 0x52, 0x49, 0x54, 0x59, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x54, 0x45, 0x4d,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x10, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x47, 0x4f, 0x4c, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48,
	0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x45, 0x4e, 0x47, 0x4f, 0x55, 0x54, 0x48, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x5f, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10,
	0x09, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x46, 0x47, 0x5f, 0x4e, 0x4f, 0x5f, 0x54, 0x48, 0x49, 0x53,
	0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x41, 0x4d, 0x45,
	0x5f, 0x4c, 0x45, 0x4e, 0x5f, 0x4f, 0x55, 0x54, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x0d, 0x42,
	0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_msgtype_proto_rawDescOnce sync.Once
	file_msgtype_proto_rawDescData = file_msgtype_proto_rawDesc
)

func file_msgtype_proto_rawDescGZIP() []byte {
	file_msgtype_proto_rawDescOnce.Do(func() {
		file_msgtype_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgtype_proto_rawDescData)
	})
	return file_msgtype_proto_rawDescData
}

var file_msgtype_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_msgtype_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_msgtype_proto_goTypes = []interface{}{
	(MSG)(0),        // 0: message.MSG
	(MSG_RESULT)(0), // 1: message.MSG_RESULT
	(*Error)(nil),   // 2: message.Error
	(*Ping)(nil),    // 3: message.Ping
	(*Pong)(nil),    // 4: message.Pong
}
var file_msgtype_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msgtype_proto_init() }
func file_msgtype_proto_init() {
	if File_msgtype_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msgtype_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_msgtype_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_msgtype_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
			RawDescriptor: file_msgtype_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msgtype_proto_goTypes,
		DependencyIndexes: file_msgtype_proto_depIdxs,
		EnumInfos:         file_msgtype_proto_enumTypes,
		MessageInfos:      file_msgtype_proto_msgTypes,
	}.Build()
	File_msgtype_proto = out.File
	file_msgtype_proto_rawDesc = nil
	file_msgtype_proto_goTypes = nil
	file_msgtype_proto_depIdxs = nil
}
