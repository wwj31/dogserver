// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: agent.proto

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

// 获取上、下级基础信息
type AgentMembersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AgentMembersReq) Reset() {
	*x = AgentMembersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentMembersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentMembersReq) ProtoMessage() {}

func (x *AgentMembersReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentMembersReq.ProtoReflect.Descriptor instead.
func (*AgentMembersReq) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

type AgentMembersRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpMember    *PlayerInfo   `protobuf:"bytes,1,opt,name=UpMember,proto3" json:"UpMember,omitempty"`
	DownMembers []*PlayerInfo `protobuf:"bytes,2,rep,name=DownMembers,proto3" json:"DownMembers,omitempty"`
}

func (x *AgentMembersRsp) Reset() {
	*x = AgentMembersRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentMembersRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentMembersRsp) ProtoMessage() {}

func (x *AgentMembersRsp) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentMembersRsp.ProtoReflect.Descriptor instead.
func (*AgentMembersRsp) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

func (x *AgentMembersRsp) GetUpMember() *PlayerInfo {
	if x != nil {
		return x.UpMember
	}
	return nil
}

func (x *AgentMembersRsp) GetDownMembers() []*PlayerInfo {
	if x != nil {
		return x.DownMembers
	}
	return nil
}

// 获取下级每日游戏信息
type AgentDownDailyStatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortIds []int64 `protobuf:"varint,1,rep,packed,name=ShortIds,proto3" json:"ShortIds,omitempty"` // 查询的下级ID
}

func (x *AgentDownDailyStatReq) Reset() {
	*x = AgentDownDailyStatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentDownDailyStatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentDownDailyStatReq) ProtoMessage() {}

func (x *AgentDownDailyStatReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentDownDailyStatReq.ProtoReflect.Descriptor instead.
func (*AgentDownDailyStatReq) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{2}
}

func (x *AgentDownDailyStatReq) GetShortIds() []int64 {
	if x != nil {
		return x.ShortIds
	}
	return nil
}

type AgentDownDailyStatRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DownDailyStats map[int64]*PlayerDailyStat `protobuf:"bytes,1,rep,name=DownDailyStats,proto3" json:"DownDailyStats,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AgentDownDailyStatRsp) Reset() {
	*x = AgentDownDailyStatRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentDownDailyStatRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentDownDailyStatRsp) ProtoMessage() {}

func (x *AgentDownDailyStatRsp) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentDownDailyStatRsp.ProtoReflect.Descriptor instead.
func (*AgentDownDailyStatRsp) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{3}
}

func (x *AgentDownDailyStatRsp) GetDownDailyStats() map[int64]*PlayerDailyStat {
	if x != nil {
		return x.DownDailyStats
	}
	return nil
}

// 获取返利以及下级返利点位相关信息
type AgentRebateInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AgentRebateInfoReq) Reset() {
	*x = AgentRebateInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRebateInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRebateInfoReq) ProtoMessage() {}

func (x *AgentRebateInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRebateInfoReq.ProtoReflect.Descriptor instead.
func (*AgentRebateInfoReq) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{4}
}

type AgentRebateInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnRebatePoints int32           `protobuf:"varint,1,opt,name=OwnRebatePoints,proto3" json:"OwnRebatePoints,omitempty"`                                                                                // 自己的返利点位 0~100%
	DownPoints      map[int64]int32 `protobuf:"bytes,2,rep,name=DownPoints,proto3" json:"DownPoints,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 每一位下级的返利点位 map<shortId,Points>
}

func (x *AgentRebateInfoRsp) Reset() {
	*x = AgentRebateInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentRebateInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentRebateInfoRsp) ProtoMessage() {}

func (x *AgentRebateInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentRebateInfoRsp.ProtoReflect.Descriptor instead.
func (*AgentRebateInfoRsp) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{5}
}

func (x *AgentRebateInfoRsp) GetOwnRebatePoints() int32 {
	if x != nil {
		return x.OwnRebatePoints
	}
	return 0
}

func (x *AgentRebateInfoRsp) GetDownPoints() map[int64]int32 {
	if x != nil {
		return x.DownPoints
	}
	return nil
}

// 给下级分配点位
type SetAgentDownRebateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId int64 `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"` // 下级的短id
	Rebate  int32 `protobuf:"varint,2,opt,name=Rebate,proto3" json:"Rebate,omitempty"`   // 分配给他的点位
}

func (x *SetAgentDownRebateReq) Reset() {
	*x = SetAgentDownRebateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAgentDownRebateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAgentDownRebateReq) ProtoMessage() {}

func (x *SetAgentDownRebateReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAgentDownRebateReq.ProtoReflect.Descriptor instead.
func (*SetAgentDownRebateReq) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{6}
}

func (x *SetAgentDownRebateReq) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

func (x *SetAgentDownRebateReq) GetRebate() int32 {
	if x != nil {
		return x.Rebate
	}
	return 0
}

type SetAgentDownRebateRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId int64 `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"` // 下级的短id
	Rebate  int32 `protobuf:"varint,2,opt,name=Rebate,proto3" json:"Rebate,omitempty"`   // 最新点位
}

func (x *SetAgentDownRebateRsp) Reset() {
	*x = SetAgentDownRebateRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAgentDownRebateRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAgentDownRebateRsp) ProtoMessage() {}

func (x *SetAgentDownRebateRsp) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAgentDownRebateRsp.ProtoReflect.Descriptor instead.
func (*SetAgentDownRebateRsp) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{7}
}

func (x *SetAgentDownRebateRsp) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

func (x *SetAgentDownRebateRsp) GetRebate() int32 {
	if x != nil {
		return x.Rebate
	}
	return 0
}

// 查看当前自己的返利分
type RebateScoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RebateScoreReq) Reset() {
	*x = RebateScoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebateScoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebateScoreReq) ProtoMessage() {}

func (x *RebateScoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebateScoreReq.ProtoReflect.Descriptor instead.
func (*RebateScoreReq) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{8}
}

type RebateScoreRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gold        int64 `protobuf:"varint,1,opt,name=Gold,proto3" json:"Gold,omitempty"`               // 最新可领取的返利分
	GoldOfToday int64 `protobuf:"varint,2,opt,name=GoldOfToday,proto3" json:"GoldOfToday,omitempty"` // 今日累计返利
	GoldOfWeek  int64 `protobuf:"varint,3,opt,name=GoldOfWeek,proto3" json:"GoldOfWeek,omitempty"`   // 本周累计返利
}

func (x *RebateScoreRsp) Reset() {
	*x = RebateScoreRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebateScoreRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebateScoreRsp) ProtoMessage() {}

func (x *RebateScoreRsp) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebateScoreRsp.ProtoReflect.Descriptor instead.
func (*RebateScoreRsp) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{9}
}

func (x *RebateScoreRsp) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

func (x *RebateScoreRsp) GetGoldOfToday() int64 {
	if x != nil {
		return x.GoldOfToday
	}
	return 0
}

func (x *RebateScoreRsp) GetGoldOfWeek() int64 {
	if x != nil {
		return x.GoldOfWeek
	}
	return 0
}

// 领取返利分
type ClaimRebateScoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClaimRebateScoreReq) Reset() {
	*x = ClaimRebateScoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimRebateScoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimRebateScoreReq) ProtoMessage() {}

func (x *ClaimRebateScoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimRebateScoreReq.ProtoReflect.Descriptor instead.
func (*ClaimRebateScoreReq) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{10}
}

type ClaimRebateScoreRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gold int64 `protobuf:"varint,1,opt,name=Gold,proto3" json:"Gold,omitempty"` // 领取了多少分
}

func (x *ClaimRebateScoreRsp) Reset() {
	*x = ClaimRebateScoreRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimRebateScoreRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimRebateScoreRsp) ProtoMessage() {}

func (x *ClaimRebateScoreRsp) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimRebateScoreRsp.ProtoReflect.Descriptor instead.
func (*ClaimRebateScoreRsp) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{11}
}

func (x *ClaimRebateScoreRsp) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

// 给下级上\下分
type SetScoreForDownReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId int64 `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"` // 下级的短id
	Gold    int64 `protobuf:"varint,2,opt,name=Gold,proto3" json:"Gold,omitempty"`       // 上分发正数，下分发负数 NOTE:(用真实分数比如:10.12发10120)
}

func (x *SetScoreForDownReq) Reset() {
	*x = SetScoreForDownReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScoreForDownReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScoreForDownReq) ProtoMessage() {}

func (x *SetScoreForDownReq) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScoreForDownReq.ProtoReflect.Descriptor instead.
func (*SetScoreForDownReq) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{12}
}

func (x *SetScoreForDownReq) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

func (x *SetScoreForDownReq) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

type SetScoreForDownRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId int64 `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"`
	Gold    int64 `protobuf:"varint,2,opt,name=Gold,proto3" json:"Gold,omitempty"` // 玩家设置后最新分数
}

func (x *SetScoreForDownRsp) Reset() {
	*x = SetScoreForDownRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetScoreForDownRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetScoreForDownRsp) ProtoMessage() {}

func (x *SetScoreForDownRsp) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetScoreForDownRsp.ProtoReflect.Descriptor instead.
func (*SetScoreForDownRsp) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{13}
}

func (x *SetScoreForDownRsp) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

func (x *SetScoreForDownRsp) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x11, 0x0a, 0x0f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x22, 0x75, 0x0a, 0x0f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x52, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x08, 0x55, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x55, 0x70, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x0b, 0x44, 0x6f, 0x77, 0x6e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x44,
	0x6f, 0x77, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x33, 0x0a, 0x15, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x73, 0x22,
	0xcc, 0x01, 0x0a, 0x15, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x52, 0x73, 0x70, 0x12, 0x58, 0x0a, 0x0e, 0x44, 0x6f, 0x77,
	0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44,
	0x6f, 0x77, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x52, 0x73, 0x70, 0x2e,
	0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0e, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x1a, 0x59, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x14,
	0x0a, 0x12, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x22, 0xc8, 0x01, 0x0a, 0x12, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x62, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x4f,
	0x77, 0x6e, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x4f, 0x77, 0x6e, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x49, 0x0a, 0x0a, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x73, 0x70, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x1a, 0x3d, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x49, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x62, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x22, 0x49, 0x0a, 0x15, 0x53, 0x65,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65,
	0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52,
	0x65, 0x62, 0x61, 0x74, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x22, 0x66, 0x0a, 0x0e, 0x52, 0x65, 0x62, 0x61, 0x74,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x6f, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x47, 0x6f, 0x6c, 0x64, 0x4f, 0x66, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x47, 0x6f, 0x6c, 0x64, 0x4f, 0x66, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x47, 0x6f, 0x6c, 0x64, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x47, 0x6f, 0x6c, 0x64, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x22,
	0x15, 0x0a, 0x13, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x22, 0x29, 0x0a, 0x13, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52,
	0x65, 0x62, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x47, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x47, 0x6f, 0x6c,
	0x64, 0x22, 0x42, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x46, 0x6f, 0x72,
	0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x47, 0x6f, 0x6c, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x46, 0x6f, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData = file_agent_proto_rawDesc
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_proto_rawDescData)
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_agent_proto_goTypes = []interface{}{
	(*AgentMembersReq)(nil),       // 0: outer.AgentMembersReq
	(*AgentMembersRsp)(nil),       // 1: outer.AgentMembersRsp
	(*AgentDownDailyStatReq)(nil), // 2: outer.AgentDownDailyStatReq
	(*AgentDownDailyStatRsp)(nil), // 3: outer.AgentDownDailyStatRsp
	(*AgentRebateInfoReq)(nil),    // 4: outer.AgentRebateInfoReq
	(*AgentRebateInfoRsp)(nil),    // 5: outer.AgentRebateInfoRsp
	(*SetAgentDownRebateReq)(nil), // 6: outer.SetAgentDownRebateReq
	(*SetAgentDownRebateRsp)(nil), // 7: outer.SetAgentDownRebateRsp
	(*RebateScoreReq)(nil),        // 8: outer.RebateScoreReq
	(*RebateScoreRsp)(nil),        // 9: outer.RebateScoreRsp
	(*ClaimRebateScoreReq)(nil),   // 10: outer.ClaimRebateScoreReq
	(*ClaimRebateScoreRsp)(nil),   // 11: outer.ClaimRebateScoreRsp
	(*SetScoreForDownReq)(nil),    // 12: outer.SetScoreForDownReq
	(*SetScoreForDownRsp)(nil),    // 13: outer.SetScoreForDownRsp
	nil,                           // 14: outer.AgentDownDailyStatRsp.DownDailyStatsEntry
	nil,                           // 15: outer.AgentRebateInfoRsp.DownPointsEntry
	(*PlayerInfo)(nil),            // 16: outer.PlayerInfo
	(*PlayerDailyStat)(nil),       // 17: outer.PlayerDailyStat
}
var file_agent_proto_depIdxs = []int32{
	16, // 0: outer.AgentMembersRsp.UpMember:type_name -> outer.PlayerInfo
	16, // 1: outer.AgentMembersRsp.DownMembers:type_name -> outer.PlayerInfo
	14, // 2: outer.AgentDownDailyStatRsp.DownDailyStats:type_name -> outer.AgentDownDailyStatRsp.DownDailyStatsEntry
	15, // 3: outer.AgentRebateInfoRsp.DownPoints:type_name -> outer.AgentRebateInfoRsp.DownPointsEntry
	17, // 4: outer.AgentDownDailyStatRsp.DownDailyStatsEntry.value:type_name -> outer.PlayerDailyStat
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentMembersReq); i {
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
		file_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentMembersRsp); i {
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
		file_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentDownDailyStatReq); i {
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
		file_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentDownDailyStatRsp); i {
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
		file_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRebateInfoReq); i {
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
		file_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentRebateInfoRsp); i {
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
		file_agent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAgentDownRebateReq); i {
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
		file_agent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAgentDownRebateRsp); i {
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
		file_agent_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RebateScoreReq); i {
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
		file_agent_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RebateScoreRsp); i {
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
		file_agent_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimRebateScoreReq); i {
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
		file_agent_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimRebateScoreRsp); i {
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
		file_agent_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScoreForDownReq); i {
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
		file_agent_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetScoreForDownRsp); i {
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
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_rawDesc = nil
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}
