// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: hall.proto

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

// /////////////////////// 金币变动记录 /////////////////////////
// 变动原因
type ReasonType int32

const (
	ReasonType_ReasonUnknown  ReasonType = 0
	ReasonType_UpModifyGold   ReasonType = 1 // 被上级上\下分
	ReasonType_ModifyDownGold ReasonType = 2 // 对下级上\下分
	ReasonType_GameWinOrLose  ReasonType = 3 // 游戏结算输赢
	ReasonType_Rebate         ReasonType = 4 // 领取返利
)

// Enum value maps for ReasonType.
var (
	ReasonType_name = map[int32]string{
		0: "ReasonUnknown",
		1: "UpModifyGold",
		2: "ModifyDownGold",
		3: "GameWinOrLose",
		4: "Rebate",
	}
	ReasonType_value = map[string]int32{
		"ReasonUnknown":  0,
		"UpModifyGold":   1,
		"ModifyDownGold": 2,
		"GameWinOrLose":  3,
		"Rebate":         4,
	}
)

func (x ReasonType) Enum() *ReasonType {
	p := new(ReasonType)
	*p = x
	return p
}

func (x ReasonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReasonType) Descriptor() protoreflect.EnumDescriptor {
	return file_hall_proto_enumTypes[0].Descriptor()
}

func (ReasonType) Type() protoreflect.EnumType {
	return &file_hall_proto_enumTypes[0]
}

func (x ReasonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReasonType.Descriptor instead.
func (ReasonType) EnumDescriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{0}
}

// 设置玩家基础信息
type SetRoleInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Icon   string `protobuf:"bytes,1,opt,name=Icon,proto3" json:"Icon,omitempty"`      // 头像
	Gender int32  `protobuf:"varint,2,opt,name=Gender,proto3" json:"Gender,omitempty"` // 性别 0.男、1.女
	Name   string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`      // 服务器限制最长20个字符
}

func (x *SetRoleInfoReq) Reset() {
	*x = SetRoleInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRoleInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRoleInfoReq) ProtoMessage() {}

func (x *SetRoleInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRoleInfoReq.ProtoReflect.Descriptor instead.
func (*SetRoleInfoReq) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{0}
}

func (x *SetRoleInfoReq) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *SetRoleInfoReq) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *SetRoleInfoReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SetRoleInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Icon   string `protobuf:"bytes,1,opt,name=Icon,proto3" json:"Icon,omitempty"`      // 头像
	Gender int32  `protobuf:"varint,2,opt,name=Gender,proto3" json:"Gender,omitempty"` // 性别
	Name   string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`      //
}

func (x *SetRoleInfoRsp) Reset() {
	*x = SetRoleInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRoleInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRoleInfoRsp) ProtoMessage() {}

func (x *SetRoleInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRoleInfoRsp.ProtoReflect.Descriptor instead.
func (*SetRoleInfoRsp) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{1}
}

func (x *SetRoleInfoRsp) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *SetRoleInfoRsp) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *SetRoleInfoRsp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 手机号作为账号，建立账号密码
type BindPhoneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone    string `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *BindPhoneReq) Reset() {
	*x = BindPhoneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindPhoneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindPhoneReq) ProtoMessage() {}

func (x *BindPhoneReq) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindPhoneReq.ProtoReflect.Descriptor instead.
func (*BindPhoneReq) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{2}
}

func (x *BindPhoneReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *BindPhoneReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type BindPhoneRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *BindPhoneRsp) Reset() {
	*x = BindPhoneRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindPhoneRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindPhoneRsp) ProtoMessage() {}

func (x *BindPhoneRsp) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindPhoneRsp.ProtoReflect.Descriptor instead.
func (*BindPhoneRsp) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{3}
}

func (x *BindPhoneRsp) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

// 修改密码
type ModifyPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password    string `protobuf:"bytes,1,opt,name=Password,proto3" json:"Password,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=NewPassword,proto3" json:"NewPassword,omitempty"`
}

func (x *ModifyPasswordReq) Reset() {
	*x = ModifyPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyPasswordReq) ProtoMessage() {}

func (x *ModifyPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyPasswordReq.ProtoReflect.Descriptor instead.
func (*ModifyPasswordReq) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{4}
}

func (x *ModifyPasswordReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ModifyPasswordReq) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ModifyPasswordRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ModifyPasswordRsp) Reset() {
	*x = ModifyPasswordRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyPasswordRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyPasswordRsp) ProtoMessage() {}

func (x *ModifyPasswordRsp) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyPasswordRsp.ProtoReflect.Descriptor instead.
func (*ModifyPasswordRsp) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{5}
}

type UpdateGoldNtf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gold int64 `protobuf:"varint,1,opt,name=Gold,proto3" json:"Gold,omitempty"` // 最新金币
}

func (x *UpdateGoldNtf) Reset() {
	*x = UpdateGoldNtf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoldNtf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoldNtf) ProtoMessage() {}

func (x *UpdateGoldNtf) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoldNtf.ProtoReflect.Descriptor instead.
func (*UpdateGoldNtf) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateGoldNtf) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

// 金币变动历史记录信息
type GoldRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoldUpdateType ReasonType `protobuf:"varint,1,opt,name=GoldUpdateType,proto3,enum=outer.ReasonType" json:"GoldUpdateType,omitempty"` // 变动愿意
	Gold           int64      `protobuf:"varint,2,opt,name=Gold,proto3" json:"Gold,omitempty"`                                           // 变动值(正数为增加，负数为减少)
	AfterGold      int64      `protobuf:"varint,3,opt,name=AfterGold,proto3" json:"AfterGold,omitempty"`                                 // 变动后的值
	UpShortId      int64      `protobuf:"varint,4,opt,name=UpShortId,proto3" json:"UpShortId,omitempty"`                                 // 上级ID 用于"被"上级 上下分
	DownShortId    int64      `protobuf:"varint,5,opt,name=DownShortId,proto3" json:"DownShortId,omitempty"`                             // 下级ID 用于"对"下级 上下分
	GameType       int32      `protobuf:"varint,6,opt,name=GameType,proto3" json:"GameType,omitempty"`                                   // 游戏类型 0.麻将 1.跑得快 用于游戏中输赢
	OccurAt        int64      `protobuf:"varint,7,opt,name=OccurAt,proto3" json:"OccurAt,omitempty"`                                     // 发生时间(毫秒) 值小于等于0，都表示无效值
}

func (x *GoldRecords) Reset() {
	*x = GoldRecords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoldRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoldRecords) ProtoMessage() {}

func (x *GoldRecords) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoldRecords.ProtoReflect.Descriptor instead.
func (*GoldRecords) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{7}
}

func (x *GoldRecords) GetGoldUpdateType() ReasonType {
	if x != nil {
		return x.GoldUpdateType
	}
	return ReasonType_ReasonUnknown
}

func (x *GoldRecords) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

func (x *GoldRecords) GetAfterGold() int64 {
	if x != nil {
		return x.AfterGold
	}
	return 0
}

func (x *GoldRecords) GetUpShortId() int64 {
	if x != nil {
		return x.UpShortId
	}
	return 0
}

func (x *GoldRecords) GetDownShortId() int64 {
	if x != nil {
		return x.DownShortId
	}
	return 0
}

func (x *GoldRecords) GetGameType() int32 {
	if x != nil {
		return x.GameType
	}
	return 0
}

func (x *GoldRecords) GetOccurAt() int64 {
	if x != nil {
		return x.OccurAt
	}
	return 0
}

// 请求金币变动记录
type GoldRecordsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 例如，一页显示20条，start-end分别传0-19，第二页20-39
	StartIndex int64 `protobuf:"varint,1,opt,name=StartIndex,proto3" json:"StartIndex,omitempty"` // 开始位置
	EndIndex   int64 `protobuf:"varint,2,opt,name=EndIndex,proto3" json:"EndIndex,omitempty"`     // 结束位置
}

func (x *GoldRecordsReq) Reset() {
	*x = GoldRecordsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoldRecordsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoldRecordsReq) ProtoMessage() {}

func (x *GoldRecordsReq) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoldRecordsReq.ProtoReflect.Descriptor instead.
func (*GoldRecordsReq) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{8}
}

func (x *GoldRecordsReq) GetStartIndex() int64 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

func (x *GoldRecordsReq) GetEndIndex() int64 {
	if x != nil {
		return x.EndIndex
	}
	return 0
}

type GoldRecordsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalLen int64          `protobuf:"varint,1,opt,name=TotalLen,proto3" json:"TotalLen,omitempty"` // 总长度
	Records  []*GoldRecords `protobuf:"bytes,2,rep,name=Records,proto3" json:"Records,omitempty"`    // 金币变化记录
}

func (x *GoldRecordsRsp) Reset() {
	*x = GoldRecordsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hall_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoldRecordsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoldRecordsRsp) ProtoMessage() {}

func (x *GoldRecordsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_hall_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoldRecordsRsp.ProtoReflect.Descriptor instead.
func (*GoldRecordsRsp) Descriptor() ([]byte, []int) {
	return file_hall_proto_rawDescGZIP(), []int{9}
}

func (x *GoldRecordsRsp) GetTotalLen() int64 {
	if x != nil {
		return x.TotalLen
	}
	return 0
}

func (x *GoldRecordsRsp) GetRecords() []*GoldRecords {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_hall_proto protoreflect.FileDescriptor

var file_hall_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x0c, 0x42, 0x69, 0x6e, 0x64, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x42, 0x69, 0x6e,
	0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22,
	0x51, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x73, 0x70, 0x22, 0x23, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x6f, 0x6c, 0x64, 0x4e, 0x74, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x6f, 0x6c, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x22, 0xf0, 0x01, 0x0a,
	0x0b, 0x47, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0e,
	0x47, 0x6f, 0x6c, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x47, 0x6f, 0x6c, 0x64, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x47, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x47, 0x6f, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x6f, 0x77, 0x6e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x44, 0x6f,
	0x77, 0x6e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x61, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x47, 0x61, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x41, 0x74, 0x22,
	0x4c, 0x0a, 0x0e, 0x47, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x45, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x5a, 0x0a,
	0x0e, 0x47, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x73, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x07, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x07, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2a, 0x64, 0x0a, 0x0a, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x70,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x47, 0x6f, 0x6c, 0x64, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x47, 0x6f, 0x6c, 0x64, 0x10, 0x02,
	0x12, 0x11, 0x0a, 0x0d, 0x47, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x4f, 0x72, 0x4c, 0x6f, 0x73,
	0x65, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x10, 0x04, 0x42,
	0x08, 0x5a, 0x06, 0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_hall_proto_rawDescOnce sync.Once
	file_hall_proto_rawDescData = file_hall_proto_rawDesc
)

func file_hall_proto_rawDescGZIP() []byte {
	file_hall_proto_rawDescOnce.Do(func() {
		file_hall_proto_rawDescData = protoimpl.X.CompressGZIP(file_hall_proto_rawDescData)
	})
	return file_hall_proto_rawDescData
}

var file_hall_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hall_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_hall_proto_goTypes = []interface{}{
	(ReasonType)(0),           // 0: outer.ReasonType
	(*SetRoleInfoReq)(nil),    // 1: outer.SetRoleInfoReq
	(*SetRoleInfoRsp)(nil),    // 2: outer.SetRoleInfoRsp
	(*BindPhoneReq)(nil),      // 3: outer.BindPhoneReq
	(*BindPhoneRsp)(nil),      // 4: outer.BindPhoneRsp
	(*ModifyPasswordReq)(nil), // 5: outer.ModifyPasswordReq
	(*ModifyPasswordRsp)(nil), // 6: outer.ModifyPasswordRsp
	(*UpdateGoldNtf)(nil),     // 7: outer.UpdateGoldNtf
	(*GoldRecords)(nil),       // 8: outer.GoldRecords
	(*GoldRecordsReq)(nil),    // 9: outer.GoldRecordsReq
	(*GoldRecordsRsp)(nil),    // 10: outer.GoldRecordsRsp
}
var file_hall_proto_depIdxs = []int32{
	0, // 0: outer.GoldRecords.GoldUpdateType:type_name -> outer.ReasonType
	8, // 1: outer.GoldRecordsRsp.Records:type_name -> outer.GoldRecords
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hall_proto_init() }
func file_hall_proto_init() {
	if File_hall_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hall_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRoleInfoReq); i {
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
		file_hall_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRoleInfoRsp); i {
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
		file_hall_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindPhoneReq); i {
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
		file_hall_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindPhoneRsp); i {
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
		file_hall_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyPasswordReq); i {
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
		file_hall_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyPasswordRsp); i {
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
		file_hall_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoldNtf); i {
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
		file_hall_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoldRecords); i {
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
		file_hall_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoldRecordsReq); i {
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
		file_hall_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoldRecordsRsp); i {
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
			RawDescriptor: file_hall_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hall_proto_goTypes,
		DependencyIndexes: file_hall_proto_depIdxs,
		EnumInfos:         file_hall_proto_enumTypes,
		MessageInfos:      file_hall_proto_msgTypes,
	}.Build()
	File_hall_proto = out.File
	file_hall_proto_rawDesc = nil
	file_hall_proto_goTypes = nil
	file_hall_proto_depIdxs = nil
}
