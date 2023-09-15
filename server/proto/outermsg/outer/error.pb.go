// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
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
	ERROR_OK                        ERROR = 0  //
	ERROR_FAILED                    ERROR = 1  // 协议执行失败，原因模糊
	ERROR_MSG_REQ_PARAM_INVALID     ERROR = 2  // 请求参数错误
	ERROR_LOGIN_TOKEN_INVALID       ERROR = 3  // 登录token过期
	ERROR_REPEAT_LOGIN              ERROR = 4  // 被顶号
	ERROR_INVALID_PHONE_FORMAT      ERROR = 5  // 无效的电话格式
	ERROR_INVALID_PASSWORD_FORMAT   ERROR = 6  // 无效的密码格式
	ERROR_PHONE_WAS_BOUND           ERROR = 7  // 手机被绑定
	ERROR_PHONE_PASSWORD_IS_EMPTY   ERROR = 8  // 绑定手机账号，密码不能为空
	ERROR_PHONE_PASSWORD_ERROR      ERROR = 9  // 手机账号登录密码错误
	ERROR_PHONE_NOT_FOUND           ERROR = 10 // 未找到账号
	ERROR_NEW_ACCOUNT_FAILED        ERROR = 11 // 创建账号失败
	ERROR_GOLD_NOT_ENOUGH           ERROR = 12 // 金币不足
	ERROR_NAME_LEN_OUT_OF_RANGE     ERROR = 13 // 名字太长
	ERROR_MODIFY_PASSWORD_NOT_PHONE ERROR = 14 // 修改失败，未绑定手机号
	ERROR_CAN_NOT_FIND_PLAYER_INFO  ERROR = 15 // 无法通过短id找到玩家
	// 联盟、上下级
	ERROR_PLAYER_NOT_IN_ALLIANCE             ERROR = 56 // 玩家没有联盟
	ERROR_PLAYER_POSITION_LIMIT              ERROR = 57 // 玩家职位权限不够
	ERROR_PLAYER_NOT_IN_CORRECT_ALLIANCE     ERROR = 58 // 被设置的玩家不在本联盟中
	ERROR_CAN_NOT_SET_HIGHER_POSITION        ERROR = 59 // 不能对职位更高的人设置职位
	ERROR_CAN_NOT_SET_NOT_IN_DOWN_POSITION   ERROR = 60 // 不能对非直属下级设置职位
	ERROR_TARGET_IS_NOT_DOWN                 ERROR = 61 // 不是直属上下级关系
	ERROR_PLAYER_ALREADY_IN_ALLIANCE         ERROR = 62 // 玩家已有联盟
	ERROR_PLAYER_ALREADY_HAS_UP              ERROR = 63 // 玩家已有上级
	ERROR_AGENT_SET_REBATE_ONLY_HIGHER       ERROR = 64 // 设置的点位不能小于已有的点位
	ERROR_AGENT_SET_REBATE_ONLY_OUT_OF_RANGE ERROR = 65 // 下级总点位和操作玩家点位
	// 房间
	ERROR_PLAYER_ALREADY_IN_ROOM           ERROR = 104 // 玩家已经在房间内，不能重复进入房间
	ERROR_ROOM_HAS_PLAYER_CAN_NOT_DISBAND  ERROR = 105 // 房间还有人，不能解散
	ERROR_ROOM_WAS_FULL_CAN_NOT_ENTER      ERROR = 106 // 房间已满
	ERROR_PLAYER_NOT_IN_ROOM               ERROR = 107 // 玩家没有房间
	ERROR_ROOM_CAN_NOT_ENTER               ERROR = 108 // 房间当前不能进入
	ERROR_ROOM_CAN_NOT_LEAVE               ERROR = 109 // 房间当前不能退出
	ERROR_ROOM_CAN_NOT_READY               ERROR = 100 // 房间当前不能切换准备状态
	ERROR_ROOM_CAN_NOT_SET_GOLD            ERROR = 101 // 房间当前状态不允许玩家上下分
	ERROR_ROOM_PLAYER_NOT_IN_GAME          ERROR = 102 // 玩家未参与游戏
	ERROR_ROOM_CANNOT_ENTER_WITH_GOLD_LINE ERROR = 103 // 玩家金币为达到警戒线，不能进房间
	// 血战到底
	ERROR_MAHJONG_EXCHANGE3_LEN_ERROR         ERROR = 200 // 请求换三张长度不为3
	ERROR_MAHJONG_EXCHANGE3_INDEX_ERROR       ERROR = 201 // 请求换三张下标越界长度不为0-12
	ERROR_MAHJONG_EXCHANGE3_INDEX_EQUAL       ERROR = 202 // 请求换三张下标重复
	ERROR_MAHJONG_EXCHANGE3_OPERATED          ERROR = 203 // 请求换三张已经操作过
	ERROR_MAHJONG_EXCHANGE3_COLOR_ERROR       ERROR = 204 // 同花色换三张，花色不同
	ERROR_MAHJONG_ACTION_PLAYER_NOT_MATCH     ERROR = 205 // 当前未轮到该玩家行动
	ERROR_MAHJONG_ACTION_PLAYER_NOT_OPERA     ERROR = 206 // 当前玩家可以行动，但是执行的行动无效
	ERROR_MAHJONG_HU_INVALID                  ERROR = 207 // 玩家当前牌型不能胡牌
	ERROR_MAHJONG_MUST_OUT_IGNORE_COLOR       ERROR = 208 // 还存在定缺的花色的牌,必须打该花色的牌
	ERROR_MAHJONG_STATE_MSG_INVALID           ERROR = 209 // 当前状态不受理此消息
	ERROR_MAHJONG_SPARE_CARDS_WAS_EMPTY       ERROR = 210 // 牌堆没牌了
	ERROR_MAHJONG_REBATE_PARAM_MINMAX_INVALID ERROR = 220 // 麻将抽水范围参数设置错误
	ERROR_MAHJONG_REBATE_PARAM_REBATE_INVALID ERROR = 221 // 麻将抽水百分比参数设置错误
	ERROR_MAHJONG_REBATE_PARAM_LIMIT_INVALID  ERROR = 222 // 麻将抽水参数设置错误
	ERROR_MAHJONG_REBATE_PARAM_RANGE_INVALID  ERROR = 223 // 麻将抽水范围参数设置有交叉
	// 跑得快
	ERROR_FASTERRUN_STATE_MSG_INVALID             ERROR = 300 // 当前状态不受理此消息
	ERROR_FASTERRUN_PLAY_CARDS_LEN_EMPTY          ERROR = 301 // 请求打牌，发的牌是0张
	ERROR_FASTERRUN_PLAY_CARDS_MISS               ERROR = 302 // 传入手牌中不存在的牌
	ERROR_FASTERRUN_PLAY_NOT_YOUR_TURN            ERROR = 303 // 没有轮到你出牌
	ERROR_FASTERRUN_PLAY_CARDS_INVALID            ERROR = 304 // 无效的出牌牌型
	ERROR_FASTERRUN_PLAY_CARDS_SHOULD_BE_FOLLOW   ERROR = 305 // 出牌牌型不匹配
	ERROR_FASTERRUN_PLAY_CARDS_SHOULD_BE_BIGGER   ERROR = 306 // 必须出更大的牌
	ERROR_FASTERRUN_PLAY_CARDS_SIDE_CARD_LEN_ERR  ERROR = 307 // 副牌数量不对
	ERROR_FASTERRUN_PLAY_EXIST_BIGGER_CANNOT_PASS ERROR = 308 // 存在更大的牌，不能过
	ERROR_FASTERRUN_PLAY_FIRST_SPADE3_LIMIT       ERROR = 309 // 开启了手牌必带黑桃3，请打黑桃3
	ERROR_FASTERRUN_PLAY_IS_YOUR_TURN             ERROR = 310 // 你的牌权出牌，不能过
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0:   "OK",
		1:   "FAILED",
		2:   "MSG_REQ_PARAM_INVALID",
		3:   "LOGIN_TOKEN_INVALID",
		4:   "REPEAT_LOGIN",
		5:   "INVALID_PHONE_FORMAT",
		6:   "INVALID_PASSWORD_FORMAT",
		7:   "PHONE_WAS_BOUND",
		8:   "PHONE_PASSWORD_IS_EMPTY",
		9:   "PHONE_PASSWORD_ERROR",
		10:  "PHONE_NOT_FOUND",
		11:  "NEW_ACCOUNT_FAILED",
		12:  "GOLD_NOT_ENOUGH",
		13:  "NAME_LEN_OUT_OF_RANGE",
		14:  "MODIFY_PASSWORD_NOT_PHONE",
		15:  "CAN_NOT_FIND_PLAYER_INFO",
		56:  "PLAYER_NOT_IN_ALLIANCE",
		57:  "PLAYER_POSITION_LIMIT",
		58:  "PLAYER_NOT_IN_CORRECT_ALLIANCE",
		59:  "CAN_NOT_SET_HIGHER_POSITION",
		60:  "CAN_NOT_SET_NOT_IN_DOWN_POSITION",
		61:  "TARGET_IS_NOT_DOWN",
		62:  "PLAYER_ALREADY_IN_ALLIANCE",
		63:  "PLAYER_ALREADY_HAS_UP",
		64:  "AGENT_SET_REBATE_ONLY_HIGHER",
		65:  "AGENT_SET_REBATE_ONLY_OUT_OF_RANGE",
		104: "PLAYER_ALREADY_IN_ROOM",
		105: "ROOM_HAS_PLAYER_CAN_NOT_DISBAND",
		106: "ROOM_WAS_FULL_CAN_NOT_ENTER",
		107: "PLAYER_NOT_IN_ROOM",
		108: "ROOM_CAN_NOT_ENTER",
		109: "ROOM_CAN_NOT_LEAVE",
		100: "ROOM_CAN_NOT_READY",
		101: "ROOM_CAN_NOT_SET_GOLD",
		102: "ROOM_PLAYER_NOT_IN_GAME",
		103: "ROOM_CANNOT_ENTER_WITH_GOLD_LINE",
		200: "MAHJONG_EXCHANGE3_LEN_ERROR",
		201: "MAHJONG_EXCHANGE3_INDEX_ERROR",
		202: "MAHJONG_EXCHANGE3_INDEX_EQUAL",
		203: "MAHJONG_EXCHANGE3_OPERATED",
		204: "MAHJONG_EXCHANGE3_COLOR_ERROR",
		205: "MAHJONG_ACTION_PLAYER_NOT_MATCH",
		206: "MAHJONG_ACTION_PLAYER_NOT_OPERA",
		207: "MAHJONG_HU_INVALID",
		208: "MAHJONG_MUST_OUT_IGNORE_COLOR",
		209: "MAHJONG_STATE_MSG_INVALID",
		210: "MAHJONG_SPARE_CARDS_WAS_EMPTY",
		220: "MAHJONG_REBATE_PARAM_MINMAX_INVALID",
		221: "MAHJONG_REBATE_PARAM_REBATE_INVALID",
		222: "MAHJONG_REBATE_PARAM_LIMIT_INVALID",
		223: "MAHJONG_REBATE_PARAM_RANGE_INVALID",
		300: "FASTERRUN_STATE_MSG_INVALID",
		301: "FASTERRUN_PLAY_CARDS_LEN_EMPTY",
		302: "FASTERRUN_PLAY_CARDS_MISS",
		303: "FASTERRUN_PLAY_NOT_YOUR_TURN",
		304: "FASTERRUN_PLAY_CARDS_INVALID",
		305: "FASTERRUN_PLAY_CARDS_SHOULD_BE_FOLLOW",
		306: "FASTERRUN_PLAY_CARDS_SHOULD_BE_BIGGER",
		307: "FASTERRUN_PLAY_CARDS_SIDE_CARD_LEN_ERR",
		308: "FASTERRUN_PLAY_EXIST_BIGGER_CANNOT_PASS",
		309: "FASTERRUN_PLAY_FIRST_SPADE3_LIMIT",
		310: "FASTERRUN_PLAY_IS_YOUR_TURN",
	}
	ERROR_value = map[string]int32{
		"OK":                                      0,
		"FAILED":                                  1,
		"MSG_REQ_PARAM_INVALID":                   2,
		"LOGIN_TOKEN_INVALID":                     3,
		"REPEAT_LOGIN":                            4,
		"INVALID_PHONE_FORMAT":                    5,
		"INVALID_PASSWORD_FORMAT":                 6,
		"PHONE_WAS_BOUND":                         7,
		"PHONE_PASSWORD_IS_EMPTY":                 8,
		"PHONE_PASSWORD_ERROR":                    9,
		"PHONE_NOT_FOUND":                         10,
		"NEW_ACCOUNT_FAILED":                      11,
		"GOLD_NOT_ENOUGH":                         12,
		"NAME_LEN_OUT_OF_RANGE":                   13,
		"MODIFY_PASSWORD_NOT_PHONE":               14,
		"CAN_NOT_FIND_PLAYER_INFO":                15,
		"PLAYER_NOT_IN_ALLIANCE":                  56,
		"PLAYER_POSITION_LIMIT":                   57,
		"PLAYER_NOT_IN_CORRECT_ALLIANCE":          58,
		"CAN_NOT_SET_HIGHER_POSITION":             59,
		"CAN_NOT_SET_NOT_IN_DOWN_POSITION":        60,
		"TARGET_IS_NOT_DOWN":                      61,
		"PLAYER_ALREADY_IN_ALLIANCE":              62,
		"PLAYER_ALREADY_HAS_UP":                   63,
		"AGENT_SET_REBATE_ONLY_HIGHER":            64,
		"AGENT_SET_REBATE_ONLY_OUT_OF_RANGE":      65,
		"PLAYER_ALREADY_IN_ROOM":                  104,
		"ROOM_HAS_PLAYER_CAN_NOT_DISBAND":         105,
		"ROOM_WAS_FULL_CAN_NOT_ENTER":             106,
		"PLAYER_NOT_IN_ROOM":                      107,
		"ROOM_CAN_NOT_ENTER":                      108,
		"ROOM_CAN_NOT_LEAVE":                      109,
		"ROOM_CAN_NOT_READY":                      100,
		"ROOM_CAN_NOT_SET_GOLD":                   101,
		"ROOM_PLAYER_NOT_IN_GAME":                 102,
		"ROOM_CANNOT_ENTER_WITH_GOLD_LINE":        103,
		"MAHJONG_EXCHANGE3_LEN_ERROR":             200,
		"MAHJONG_EXCHANGE3_INDEX_ERROR":           201,
		"MAHJONG_EXCHANGE3_INDEX_EQUAL":           202,
		"MAHJONG_EXCHANGE3_OPERATED":              203,
		"MAHJONG_EXCHANGE3_COLOR_ERROR":           204,
		"MAHJONG_ACTION_PLAYER_NOT_MATCH":         205,
		"MAHJONG_ACTION_PLAYER_NOT_OPERA":         206,
		"MAHJONG_HU_INVALID":                      207,
		"MAHJONG_MUST_OUT_IGNORE_COLOR":           208,
		"MAHJONG_STATE_MSG_INVALID":               209,
		"MAHJONG_SPARE_CARDS_WAS_EMPTY":           210,
		"MAHJONG_REBATE_PARAM_MINMAX_INVALID":     220,
		"MAHJONG_REBATE_PARAM_REBATE_INVALID":     221,
		"MAHJONG_REBATE_PARAM_LIMIT_INVALID":      222,
		"MAHJONG_REBATE_PARAM_RANGE_INVALID":      223,
		"FASTERRUN_STATE_MSG_INVALID":             300,
		"FASTERRUN_PLAY_CARDS_LEN_EMPTY":          301,
		"FASTERRUN_PLAY_CARDS_MISS":               302,
		"FASTERRUN_PLAY_NOT_YOUR_TURN":            303,
		"FASTERRUN_PLAY_CARDS_INVALID":            304,
		"FASTERRUN_PLAY_CARDS_SHOULD_BE_FOLLOW":   305,
		"FASTERRUN_PLAY_CARDS_SHOULD_BE_BIGGER":   306,
		"FASTERRUN_PLAY_CARDS_SIDE_CARD_LEN_ERR":  307,
		"FASTERRUN_PLAY_EXIST_BIGGER_CANNOT_PASS": 308,
		"FASTERRUN_PLAY_FIRST_SPADE3_LIMIT":       309,
		"FASTERRUN_PLAY_IS_YOUR_TURN":             310,
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
	return ERROR_OK
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
	0x77, 0x6e, 0x2a, 0xa9, 0x0f, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x19, 0x0a, 0x15, 0x4d, 0x53, 0x47, 0x5f, 0x52, 0x45, 0x51, 0x5f, 0x50, 0x41, 0x52, 0x41,
	0x4d, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x5f, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x05,
	0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x53, 0x53,
	0x57, 0x4f, 0x52, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x06, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x57, 0x41, 0x53, 0x5f, 0x42, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0x07, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53,
	0x57, 0x4f, 0x52, 0x44, 0x5f, 0x49, 0x53, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x08, 0x12,
	0x18, 0x0a, 0x14, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52,
	0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x48, 0x4f,
	0x4e, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x0a, 0x12, 0x16,
	0x0a, 0x12, 0x4e, 0x45, 0x57, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x4f, 0x4c, 0x44, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x4c, 0x45, 0x4e, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x10, 0x0d, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59,
	0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x48,
	0x4f, 0x4e, 0x45, 0x10, 0x0e, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x49, 0x4e, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46,
	0x4f, 0x10, 0x0f, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x38, 0x12,
	0x19, 0x0a, 0x15, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x39, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x4c,
	0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x52, 0x52,
	0x45, 0x43, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x3a, 0x12, 0x1f,
	0x0a, 0x1b, 0x43, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x48, 0x49,
	0x47, 0x48, 0x45, 0x52, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x3b, 0x12,
	0x24, 0x0a, 0x20, 0x43, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x3c, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x3d, 0x12, 0x1e, 0x0a,
	0x1a, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f,
	0x49, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x3e, 0x12, 0x19, 0x0a,
	0x15, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f,
	0x48, 0x41, 0x53, 0x5f, 0x55, 0x50, 0x10, 0x3f, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x47, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x42, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e, 0x4c,
	0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x45, 0x52, 0x10, 0x40, 0x12, 0x26, 0x0a, 0x22, 0x41, 0x47,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x42, 0x41, 0x54, 0x45, 0x5f, 0x4f,
	0x4e, 0x4c, 0x59, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45,
	0x10, 0x41, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x68, 0x12, 0x23,
	0x0a, 0x1f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45,
	0x52, 0x5f, 0x43, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x42, 0x41, 0x4e,
	0x44, 0x10, 0x69, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x57, 0x41, 0x53, 0x5f,
	0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x43, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54,
	0x45, 0x52, 0x10, 0x6a, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x6b, 0x12, 0x16, 0x0a, 0x12,
	0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x43, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54,
	0x45, 0x52, 0x10, 0x6c, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x43, 0x41, 0x4e,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x10, 0x6d, 0x12, 0x16, 0x0a, 0x12,
	0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x43, 0x41, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x10, 0x64, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x43, 0x41, 0x4e,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x47, 0x4f, 0x4c, 0x44, 0x10, 0x65, 0x12,
	0x1b, 0x0a, 0x17, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10, 0x66, 0x12, 0x24, 0x0a, 0x20,
	0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x45,
	0x52, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x47, 0x4f, 0x4c, 0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x45,
	0x10, 0x67, 0x12, 0x20, 0x0a, 0x1b, 0x4d, 0x41, 0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f, 0x45, 0x58,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x33, 0x5f, 0x4c, 0x45, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0xc8, 0x01, 0x12, 0x22, 0x0a, 0x1d, 0x4d, 0x41, 0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f,
	0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x33, 0x5f, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xc9, 0x01, 0x12, 0x22, 0x0a, 0x1d, 0x4d, 0x41, 0x48, 0x4a,
	0x4f, 0x4e, 0x47, 0x5f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x33, 0x5f, 0x49, 0x4e,
	0x44, 0x45, 0x58, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0xca, 0x01, 0x12, 0x1f, 0x0a, 0x1a,
	0x4d, 0x41, 0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x33, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x10, 0xcb, 0x01, 0x12, 0x22, 0x0a,
	0x1d, 0x4d, 0x41, 0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x33, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xcc,
	0x01, 0x12, 0x24, 0x0a, 0x1f, 0x4d, 0x41, 0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d,
	0x41, 0x54, 0x43, 0x48, 0x10, 0xcd, 0x01, 0x12, 0x24, 0x0a, 0x1f, 0x4d, 0x41, 0x48, 0x4a, 0x4f,
	0x4e, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x10, 0xce, 0x01, 0x12, 0x17, 0x0a,
	0x12, 0x4d, 0x41, 0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f, 0x48, 0x55, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0xcf, 0x01, 0x12, 0x22, 0x0a, 0x1d, 0x4d, 0x41, 0x48, 0x4a, 0x4f, 0x4e,
	0x47, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x49, 0x47, 0x4e, 0x4f, 0x52,
	0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x10, 0xd0, 0x01, 0x12, 0x1e, 0x0a, 0x19, 0x4d, 0x41,
	0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x53, 0x47, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xd1, 0x01, 0x12, 0x22, 0x0a, 0x1d, 0x4d, 0x41,
	0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f, 0x53, 0x50, 0x41, 0x52, 0x45, 0x5f, 0x43, 0x41, 0x52, 0x44,
	0x53, 0x5f, 0x57, 0x41, 0x53, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0xd2, 0x01, 0x12, 0x28,
	0x0a, 0x23, 0x4d, 0x41, 0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x42, 0x41, 0x54, 0x45,
	0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x4d, 0x49, 0x4e, 0x4d, 0x41, 0x58, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xdc, 0x01, 0x12, 0x28, 0x0a, 0x23, 0x4d, 0x41, 0x48, 0x4a,
	0x4f, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x42, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x5f, 0x52, 0x45, 0x42, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0xdd, 0x01, 0x12, 0x27, 0x0a, 0x22, 0x4d, 0x41, 0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f, 0x52, 0x45,
	0x42, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xde, 0x01, 0x12, 0x27, 0x0a, 0x22, 0x4d,
	0x41, 0x48, 0x4a, 0x4f, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x42, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x41,
	0x52, 0x41, 0x4d, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0xdf, 0x01, 0x12, 0x20, 0x0a, 0x1b, 0x46, 0x41, 0x53, 0x54, 0x45, 0x52, 0x52, 0x55,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0xac, 0x02, 0x12, 0x23, 0x0a, 0x1e, 0x46, 0x41, 0x53, 0x54, 0x45, 0x52,
	0x52, 0x55, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x53, 0x5f, 0x4c,
	0x45, 0x4e, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0xad, 0x02, 0x12, 0x1e, 0x0a, 0x19, 0x46,
	0x41, 0x53, 0x54, 0x45, 0x52, 0x52, 0x55, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x41,
	0x52, 0x44, 0x53, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x10, 0xae, 0x02, 0x12, 0x21, 0x0a, 0x1c, 0x46,
	0x41, 0x53, 0x54, 0x45, 0x52, 0x52, 0x55, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x59, 0x4f, 0x55, 0x52, 0x5f, 0x54, 0x55, 0x52, 0x4e, 0x10, 0xaf, 0x02, 0x12, 0x21,
	0x0a, 0x1c, 0x46, 0x41, 0x53, 0x54, 0x45, 0x52, 0x52, 0x55, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59,
	0x5f, 0x43, 0x41, 0x52, 0x44, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xb0,
	0x02, 0x12, 0x2a, 0x0a, 0x25, 0x46, 0x41, 0x53, 0x54, 0x45, 0x52, 0x52, 0x55, 0x4e, 0x5f, 0x50,
	0x4c, 0x41, 0x59, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x53, 0x5f, 0x53, 0x48, 0x4f, 0x55, 0x4c, 0x44,
	0x5f, 0x42, 0x45, 0x5f, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0xb1, 0x02, 0x12, 0x2a, 0x0a,
	0x25, 0x46, 0x41, 0x53, 0x54, 0x45, 0x52, 0x52, 0x55, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f,
	0x43, 0x41, 0x52, 0x44, 0x53, 0x5f, 0x53, 0x48, 0x4f, 0x55, 0x4c, 0x44, 0x5f, 0x42, 0x45, 0x5f,
	0x42, 0x49, 0x47, 0x47, 0x45, 0x52, 0x10, 0xb2, 0x02, 0x12, 0x2b, 0x0a, 0x26, 0x46, 0x41, 0x53,
	0x54, 0x45, 0x52, 0x52, 0x55, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x41, 0x52, 0x44,
	0x53, 0x5f, 0x53, 0x49, 0x44, 0x45, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x5f, 0x4c, 0x45, 0x4e, 0x5f,
	0x45, 0x52, 0x52, 0x10, 0xb3, 0x02, 0x12, 0x2c, 0x0a, 0x27, 0x46, 0x41, 0x53, 0x54, 0x45, 0x52,
	0x52, 0x55, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x5f, 0x42,
	0x49, 0x47, 0x47, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x41, 0x53,
	0x53, 0x10, 0xb4, 0x02, 0x12, 0x26, 0x0a, 0x21, 0x46, 0x41, 0x53, 0x54, 0x45, 0x52, 0x52, 0x55,
	0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x53, 0x50, 0x41,
	0x44, 0x45, 0x33, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0xb5, 0x02, 0x12, 0x20, 0x0a, 0x1b,
	0x46, 0x41, 0x53, 0x54, 0x45, 0x52, 0x52, 0x55, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x49,
	0x53, 0x5f, 0x59, 0x4f, 0x55, 0x52, 0x5f, 0x54, 0x55, 0x52, 0x4e, 0x10, 0xb6, 0x02, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
