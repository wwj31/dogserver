// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: room.proto

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

// 房间信息
type RoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32         `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	GameType GameType      `protobuf:"varint,2,opt,name=GameType,proto3,enum=outer.GameType" json:"GameType,omitempty"`
	Players  []*PlayerInfo `protobuf:"bytes,3,rep,name=Players,proto3" json:"Players,omitempty"` // 房间内的成员信息
}

func (x *RoomInfo) Reset() {
	*x = RoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomInfo) ProtoMessage() {}

func (x *RoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomInfo.ProtoReflect.Descriptor instead.
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{0}
}

func (x *RoomInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoomInfo) GetGameType() GameType {
	if x != nil {
		return x.GameType
	}
	return GameType_Mahjong
}

func (x *RoomInfo) GetPlayers() []*PlayerInfo {
	if x != nil {
		return x.Players
	}
	return nil
}

// 创建房间
type CreateRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameType GameType `protobuf:"varint,1,opt,name=GameType,proto3,enum=outer.GameType" json:"GameType,omitempty"`
}

func (x *CreateRoomReq) Reset() {
	*x = CreateRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomReq) ProtoMessage() {}

func (x *CreateRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomReq.ProtoReflect.Descriptor instead.
func (*CreateRoomReq) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoomReq) GetGameType() GameType {
	if x != nil {
		return x.GameType
	}
	return GameType_Mahjong
}

type CreateRoomRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *RoomInfo `protobuf:"bytes,1,opt,name=Room,proto3" json:"Room,omitempty"`
}

func (x *CreateRoomRsp) Reset() {
	*x = CreateRoomRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRsp) ProtoMessage() {}

func (x *CreateRoomRsp) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRsp.ProtoReflect.Descriptor instead.
func (*CreateRoomRsp) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoomRsp) GetRoom() *RoomInfo {
	if x != nil {
		return x.Room
	}
	return nil
}

// 解散房间
type DisbandRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DisbandRoomReq) Reset() {
	*x = DisbandRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisbandRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisbandRoomReq) ProtoMessage() {}

func (x *DisbandRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisbandRoomReq.ProtoReflect.Descriptor instead.
func (*DisbandRoomReq) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{3}
}

func (x *DisbandRoomReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DisbandRoomRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DisbandRoomRsp) Reset() {
	*x = DisbandRoomRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisbandRoomRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisbandRoomRsp) ProtoMessage() {}

func (x *DisbandRoomRsp) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisbandRoomRsp.ProtoReflect.Descriptor instead.
func (*DisbandRoomRsp) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{4}
}

func (x *DisbandRoomRsp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 房间列表
type RoomListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoomListReq) Reset() {
	*x = RoomListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomListReq) ProtoMessage() {}

func (x *RoomListReq) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomListReq.ProtoReflect.Descriptor instead.
func (*RoomListReq) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{5}
}

type RoomListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomList []*RoomInfo `protobuf:"bytes,1,rep,name=RoomList,proto3" json:"RoomList,omitempty"`
}

func (x *RoomListRsp) Reset() {
	*x = RoomListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomListRsp) ProtoMessage() {}

func (x *RoomListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomListRsp.ProtoReflect.Descriptor instead.
func (*RoomListRsp) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{6}
}

func (x *RoomListRsp) GetRoomList() []*RoomInfo {
	if x != nil {
		return x.RoomList
	}
	return nil
}

// 加入房间
type JoinRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int32 `protobuf:"varint,1,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
}

func (x *JoinRoomReq) Reset() {
	*x = JoinRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomReq) ProtoMessage() {}

func (x *JoinRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomReq.ProtoReflect.Descriptor instead.
func (*JoinRoomReq) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{7}
}

func (x *JoinRoomReq) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type JoinRoomRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *RoomInfo `protobuf:"bytes,1,opt,name=Room,proto3" json:"Room,omitempty"`
}

func (x *JoinRoomRsp) Reset() {
	*x = JoinRoomRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRsp) ProtoMessage() {}

func (x *JoinRoomRsp) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRsp.ProtoReflect.Descriptor instead.
func (*JoinRoomRsp) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{8}
}

func (x *JoinRoomRsp) GetRoom() *RoomInfo {
	if x != nil {
		return x.Room
	}
	return nil
}

// 离开房间
type LeaveRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaveRoomReq) Reset() {
	*x = LeaveRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRoomReq) ProtoMessage() {}

func (x *LeaveRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRoomReq.ProtoReflect.Descriptor instead.
func (*LeaveRoomReq) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{9}
}

type LeaveRoomRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LeaveRoomRsp) Reset() {
	*x = LeaveRoomRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRoomRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRoomRsp) ProtoMessage() {}

func (x *LeaveRoomRsp) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRoomRsp.ProtoReflect.Descriptor instead.
func (*LeaveRoomRsp) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{10}
}

// 通知类消息
type RoomPlayerEnterNtf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *PlayerInfo `protobuf:"bytes,1,opt,name=Player,proto3" json:"Player,omitempty"`
}

func (x *RoomPlayerEnterNtf) Reset() {
	*x = RoomPlayerEnterNtf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomPlayerEnterNtf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomPlayerEnterNtf) ProtoMessage() {}

func (x *RoomPlayerEnterNtf) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomPlayerEnterNtf.ProtoReflect.Descriptor instead.
func (*RoomPlayerEnterNtf) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{11}
}

func (x *RoomPlayerEnterNtf) GetPlayer() *PlayerInfo {
	if x != nil {
		return x.Player
	}
	return nil
}

type RoomPlayerLeaveNtf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId int64 `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"`
}

func (x *RoomPlayerLeaveNtf) Reset() {
	*x = RoomPlayerLeaveNtf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomPlayerLeaveNtf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomPlayerLeaveNtf) ProtoMessage() {}

func (x *RoomPlayerLeaveNtf) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomPlayerLeaveNtf.ProtoReflect.Descriptor instead.
func (*RoomPlayerLeaveNtf) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{12}
}

func (x *RoomPlayerLeaveNtf) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

type RoomPlayerOnlineNtf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId int64 `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"`
	Online  bool  `protobuf:"varint,2,opt,name=Online,proto3" json:"Online,omitempty"` // true在线、false离线
}

func (x *RoomPlayerOnlineNtf) Reset() {
	*x = RoomPlayerOnlineNtf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomPlayerOnlineNtf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomPlayerOnlineNtf) ProtoMessage() {}

func (x *RoomPlayerOnlineNtf) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomPlayerOnlineNtf.ProtoReflect.Descriptor instead.
func (*RoomPlayerOnlineNtf) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{13}
}

func (x *RoomPlayerOnlineNtf) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

func (x *RoomPlayerOnlineNtf) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

var File_room_proto protoreflect.FileDescriptor

var file_room_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x08, 0x52,
	0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x47, 0x61, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x22, 0x3c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x12, 0x2b, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x34, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x73, 0x70,
	0x12, 0x23, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x62, 0x61, 0x6e, 0x64,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x62, 0x61,
	0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x6f, 0x6f,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x3a, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x52, 0x6f, 0x6f, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x0b, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x04, 0x52, 0x6f,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x22,
	0x0e, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x22,
	0x0e, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x73, 0x70, 0x22,
	0x3f, 0x0a, 0x12, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x4e, 0x74, 0x66, 0x12, 0x29, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x22, 0x2e, 0x0a, 0x12, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x4e, 0x74, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64,
	0x22, 0x47, 0x0a, 0x13, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x74, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_room_proto_rawDescOnce sync.Once
	file_room_proto_rawDescData = file_room_proto_rawDesc
)

func file_room_proto_rawDescGZIP() []byte {
	file_room_proto_rawDescOnce.Do(func() {
		file_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_proto_rawDescData)
	})
	return file_room_proto_rawDescData
}

var file_room_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_room_proto_goTypes = []interface{}{
	(*RoomInfo)(nil),            // 0: outer.RoomInfo
	(*CreateRoomReq)(nil),       // 1: outer.CreateRoomReq
	(*CreateRoomRsp)(nil),       // 2: outer.CreateRoomRsp
	(*DisbandRoomReq)(nil),      // 3: outer.DisbandRoomReq
	(*DisbandRoomRsp)(nil),      // 4: outer.DisbandRoomRsp
	(*RoomListReq)(nil),         // 5: outer.RoomListReq
	(*RoomListRsp)(nil),         // 6: outer.RoomListRsp
	(*JoinRoomReq)(nil),         // 7: outer.JoinRoomReq
	(*JoinRoomRsp)(nil),         // 8: outer.JoinRoomRsp
	(*LeaveRoomReq)(nil),        // 9: outer.LeaveRoomReq
	(*LeaveRoomRsp)(nil),        // 10: outer.LeaveRoomRsp
	(*RoomPlayerEnterNtf)(nil),  // 11: outer.RoomPlayerEnterNtf
	(*RoomPlayerLeaveNtf)(nil),  // 12: outer.RoomPlayerLeaveNtf
	(*RoomPlayerOnlineNtf)(nil), // 13: outer.RoomPlayerOnlineNtf
	(GameType)(0),               // 14: outer.GameType
	(*PlayerInfo)(nil),          // 15: outer.PlayerInfo
}
var file_room_proto_depIdxs = []int32{
	14, // 0: outer.RoomInfo.GameType:type_name -> outer.GameType
	15, // 1: outer.RoomInfo.Players:type_name -> outer.PlayerInfo
	14, // 2: outer.CreateRoomReq.GameType:type_name -> outer.GameType
	0,  // 3: outer.CreateRoomRsp.Room:type_name -> outer.RoomInfo
	0,  // 4: outer.RoomListRsp.RoomList:type_name -> outer.RoomInfo
	0,  // 5: outer.JoinRoomRsp.Room:type_name -> outer.RoomInfo
	15, // 6: outer.RoomPlayerEnterNtf.Player:type_name -> outer.PlayerInfo
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_room_proto_init() }
func file_room_proto_init() {
	if File_room_proto != nil {
		return
	}
	file_base_proto_init()
	file_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomInfo); i {
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
		file_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomReq); i {
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
		file_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomRsp); i {
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
		file_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisbandRoomReq); i {
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
		file_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisbandRoomRsp); i {
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
		file_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomListReq); i {
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
		file_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomListRsp); i {
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
		file_room_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomReq); i {
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
		file_room_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomRsp); i {
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
		file_room_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRoomReq); i {
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
		file_room_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRoomRsp); i {
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
		file_room_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomPlayerEnterNtf); i {
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
		file_room_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomPlayerLeaveNtf); i {
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
		file_room_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomPlayerOnlineNtf); i {
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
			RawDescriptor: file_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_room_proto_goTypes,
		DependencyIndexes: file_room_proto_depIdxs,
		MessageInfos:      file_room_proto_msgTypes,
	}.Build()
	File_room_proto = out.File
	file_room_proto_rawDesc = nil
	file_room_proto_goTypes = nil
	file_room_proto_depIdxs = nil
}
