// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: base.proto

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

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId int32  `protobuf:"varint,1,opt,name=MsgId,proto3" json:"MsgId,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

func (x *Base) GetMsgId() int32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *Base) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RID        string   `protobuf:"bytes,1,opt,name=RID,proto3" json:"RID,omitempty"`                                // 玩家id
	ShortId    int64    `protobuf:"varint,2,opt,name=ShortId,proto3" json:"ShortId,omitempty"`                       // 短id
	Name       string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`                              // 名字
	Icon       string   `protobuf:"bytes,4,opt,name=Icon,proto3" json:"Icon,omitempty"`                              // 头像
	Gender     int32    `protobuf:"varint,5,opt,name=Gender,proto3" json:"Gender,omitempty"`                         // 性别
	AllianceId int32    `protobuf:"varint,6,opt,name=AllianceId,proto3" json:"AllianceId,omitempty"`                 // 联盟id
	Position   Position `protobuf:"varint,7,opt,name=Position,proto3,enum=outer.Position" json:"Position,omitempty"` // 职位
	LoginAt    int64    `protobuf:"varint,8,opt,name=LoginAt,proto3" json:"LoginAt,omitempty"`                       // 登录时间(秒)
	LogoutAt   int64    `protobuf:"varint,9,opt,name=LogoutAt,proto3" json:"LogoutAt,omitempty"`                     // 离线时间(秒)
	UpShortId  int64    `protobuf:"varint,10,opt,name=UpShortId,proto3" json:"UpShortId,omitempty"`                  // 上级id
	Gold       int64    `protobuf:"varint,11,opt,name=Gold,proto3" json:"Gold,omitempty"`                            // 金币
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerInfo) GetRID() string {
	if x != nil {
		return x.RID
	}
	return ""
}

func (x *PlayerInfo) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

func (x *PlayerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerInfo) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *PlayerInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *PlayerInfo) GetAllianceId() int32 {
	if x != nil {
		return x.AllianceId
	}
	return 0
}

func (x *PlayerInfo) GetPosition() Position {
	if x != nil {
		return x.Position
	}
	return Position_NoAlliance
}

func (x *PlayerInfo) GetLoginAt() int64 {
	if x != nil {
		return x.LoginAt
	}
	return 0
}

func (x *PlayerInfo) GetLogoutAt() int64 {
	if x != nil {
		return x.LogoutAt
	}
	return 0
}

func (x *PlayerInfo) GetUpShortId() int64 {
	if x != nil {
		return x.UpShortId
	}
	return 0
}

func (x *PlayerInfo) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

// 游戏参数
type GameParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mahjong *MahjongParams `protobuf:"bytes,1,opt,name=Mahjong,proto3" json:"Mahjong,omitempty"`
	NiuNiu  *NiuNiuParams  `protobuf:"bytes,2,opt,name=NiuNiu,proto3" json:"NiuNiu,omitempty"`
}

func (x *GameParams) Reset() {
	*x = GameParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameParams) ProtoMessage() {}

func (x *GameParams) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameParams.ProtoReflect.Descriptor instead.
func (*GameParams) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *GameParams) GetMahjong() *MahjongParams {
	if x != nil {
		return x.Mahjong
	}
	return nil
}

func (x *GameParams) GetNiuNiu() *NiuNiuParams {
	if x != nil {
		return x.NiuNiu
	}
	return nil
}

// 麻将
type MahjongParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayCountLimit       int32         `protobuf:"varint,1,opt,name=playCountLimit,proto3" json:"playCountLimit,omitempty"`              // 必须打满多少局才能退出
	BaseScore            int32         `protobuf:"varint,2,opt,name=BaseScore,proto3" json:"BaseScore,omitempty"`                        // 底分
	BaseScoreTimes       float32       `protobuf:"fixed32,3,opt,name=BaseScoreTimes,proto3" json:"BaseScoreTimes,omitempty"`             // 倍数 0.1、0.5、1、2
	ZiMoJia              int32         `protobuf:"varint,4,opt,name=ZiMoJia,proto3" json:"ZiMoJia,omitempty"`                            // 自摸 0.自摸加番、1.自摸加底
	FanUpLimit           int32         `protobuf:"varint,5,opt,name=FanUpLimit,proto3" json:"FanUpLimit,omitempty"`                      // 翻数封顶 1、2、3、4、5
	DianGangHua          int32         `protobuf:"varint,6,opt,name=DianGangHua,proto3" json:"DianGangHua,omitempty"`                    // 点杠花 0.点炮、1.自摸
	HuanSanZhang         int32         `protobuf:"varint,7,opt,name=HuanSanZhang,proto3" json:"HuanSanZhang,omitempty"`                  // 换三张 0.无换三张、1.同花色换三张、1.任意换三张
	YaoJiuDui            bool          `protobuf:"varint,8,opt,name=YaoJiuDui,proto3" json:"YaoJiuDui,omitempty"`                        // 幺九将对
	MenQingZhongZhang    bool          `protobuf:"varint,9,opt,name=MenQingZhongZhang,proto3" json:"MenQingZhongZhang,omitempty"`        // 门清中张
	TianDiHu             bool          `protobuf:"varint,10,opt,name=TianDiHu,proto3" json:"TianDiHu,omitempty"`                         // 天地胡
	DianPaoPingHu        bool          `protobuf:"varint,11,opt,name=DianPaoPingHu,proto3" json:"DianPaoPingHu,omitempty"`               // 点炮可平胡
	JiaXinWu             bool          `protobuf:"varint,12,opt,name=JiaXinWu,proto3" json:"JiaXinWu,omitempty"`                         // 夹心五
	GangImmediatelyScore bool          `protobuf:"varint,13,opt,name=GangImmediatelyScore,proto3" json:"GangImmediatelyScore,omitempty"` // 杠分实时结算
	HuImmediatelyScore   bool          `protobuf:"varint,14,opt,name=HuImmediatelyScore,proto3" json:"HuImmediatelyScore,omitempty"`     // 胡牌实时结算
	AllowScoreSmallZero  bool          `protobuf:"varint,15,opt,name=AllowScoreSmallZero,proto3" json:"AllowScoreSmallZero,omitempty"`   // 是否允许负分
	BigWinner            bool          `protobuf:"varint,16,opt,name=BigWinner,proto3" json:"BigWinner,omitempty"`                       // 是否开启大赢家模式
	ReBate               *RebateParams `protobuf:"bytes,17,opt,name=ReBate,proto3" json:"ReBate,omitempty"`                              // 抽水相关设置
}

func (x *MahjongParams) Reset() {
	*x = MahjongParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MahjongParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MahjongParams) ProtoMessage() {}

func (x *MahjongParams) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MahjongParams.ProtoReflect.Descriptor instead.
func (*MahjongParams) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *MahjongParams) GetPlayCountLimit() int32 {
	if x != nil {
		return x.PlayCountLimit
	}
	return 0
}

func (x *MahjongParams) GetBaseScore() int32 {
	if x != nil {
		return x.BaseScore
	}
	return 0
}

func (x *MahjongParams) GetBaseScoreTimes() float32 {
	if x != nil {
		return x.BaseScoreTimes
	}
	return 0
}

func (x *MahjongParams) GetZiMoJia() int32 {
	if x != nil {
		return x.ZiMoJia
	}
	return 0
}

func (x *MahjongParams) GetFanUpLimit() int32 {
	if x != nil {
		return x.FanUpLimit
	}
	return 0
}

func (x *MahjongParams) GetDianGangHua() int32 {
	if x != nil {
		return x.DianGangHua
	}
	return 0
}

func (x *MahjongParams) GetHuanSanZhang() int32 {
	if x != nil {
		return x.HuanSanZhang
	}
	return 0
}

func (x *MahjongParams) GetYaoJiuDui() bool {
	if x != nil {
		return x.YaoJiuDui
	}
	return false
}

func (x *MahjongParams) GetMenQingZhongZhang() bool {
	if x != nil {
		return x.MenQingZhongZhang
	}
	return false
}

func (x *MahjongParams) GetTianDiHu() bool {
	if x != nil {
		return x.TianDiHu
	}
	return false
}

func (x *MahjongParams) GetDianPaoPingHu() bool {
	if x != nil {
		return x.DianPaoPingHu
	}
	return false
}

func (x *MahjongParams) GetJiaXinWu() bool {
	if x != nil {
		return x.JiaXinWu
	}
	return false
}

func (x *MahjongParams) GetGangImmediatelyScore() bool {
	if x != nil {
		return x.GangImmediatelyScore
	}
	return false
}

func (x *MahjongParams) GetHuImmediatelyScore() bool {
	if x != nil {
		return x.HuImmediatelyScore
	}
	return false
}

func (x *MahjongParams) GetAllowScoreSmallZero() bool {
	if x != nil {
		return x.AllowScoreSmallZero
	}
	return false
}

func (x *MahjongParams) GetBigWinner() bool {
	if x != nil {
		return x.BigWinner
	}
	return false
}

func (x *MahjongParams) GetReBate() *RebateParams {
	if x != nil {
		return x.ReBate
	}
	return nil
}

// 牛牛
type NiuNiuParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NiuNiuParams) Reset() {
	*x = NiuNiuParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NiuNiuParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NiuNiuParams) ProtoMessage() {}

func (x *NiuNiuParams) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NiuNiuParams.ProtoReflect.Descriptor instead.
func (*NiuNiuParams) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{4}
}

type RangeParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid            bool  `protobuf:"varint,1,opt,name=Valid,proto3" json:"Valid,omitempty"`                       // 是否勾选生效
	Min              int64 `protobuf:"varint,2,opt,name=Min,proto3" json:"Min,omitempty"`                           // 区间最小值
	Max              int64 `protobuf:"varint,3,opt,name=Max,proto3" json:"Max,omitempty"`                           // 区间最大值
	RebateRatio      int64 `protobuf:"varint,4,opt,name=RebateRatio,proto3" json:"RebateRatio,omitempty"`           // 抽水百分比,发5就是抽5%
	MinimumGuarantee int64 `protobuf:"varint,5,opt,name=MinimumGuarantee,proto3" json:"MinimumGuarantee,omitempty"` // 保底          注：最终抽水=保底+(盈利*抽水百分比)
	MinimumRebate    int64 `protobuf:"varint,6,opt,name=MinimumRebate,proto3" json:"MinimumRebate,omitempty"`       // 最低起抽       注：最少要赢这么多分才抽,否则就不抽
}

func (x *RangeParams) Reset() {
	*x = RangeParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RangeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeParams) ProtoMessage() {}

func (x *RangeParams) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeParams.ProtoReflect.Descriptor instead.
func (*RangeParams) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{5}
}

func (x *RangeParams) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *RangeParams) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *RangeParams) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *RangeParams) GetRebateRatio() int64 {
	if x != nil {
		return x.RebateRatio
	}
	return 0
}

func (x *RangeParams) GetMinimumGuarantee() int64 {
	if x != nil {
		return x.MinimumGuarantee
	}
	return 0
}

func (x *RangeParams) GetMinimumRebate() int64 {
	if x != nil {
		return x.MinimumRebate
	}
	return 0
}

// 返利参数
type RebateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RangeL1 *RangeParams `protobuf:"bytes,1,opt,name=RangeL1,proto3" json:"RangeL1,omitempty"` // 第1区间
	RangeL2 *RangeParams `protobuf:"bytes,2,opt,name=RangeL2,proto3" json:"RangeL2,omitempty"` // 第2区间
	RangeL3 *RangeParams `protobuf:"bytes,3,opt,name=RangeL3,proto3" json:"RangeL3,omitempty"` // 第3区间
	RangeL4 *RangeParams `protobuf:"bytes,4,opt,name=RangeL4,proto3" json:"RangeL4,omitempty"` // 第4区间
}

func (x *RebateParams) Reset() {
	*x = RebateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebateParams) ProtoMessage() {}

func (x *RebateParams) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebateParams.ProtoReflect.Descriptor instead.
func (*RebateParams) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{6}
}

func (x *RebateParams) GetRangeL1() *RangeParams {
	if x != nil {
		return x.RangeL1
	}
	return nil
}

func (x *RebateParams) GetRangeL2() *RangeParams {
	if x != nil {
		return x.RangeL2
	}
	return nil
}

func (x *RebateParams) GetRangeL3() *RangeParams {
	if x != nil {
		return x.RangeL3
	}
	return nil
}

func (x *RebateParams) GetRangeL4() *RangeParams {
	if x != nil {
		return x.RangeL4
	}
	return nil
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x1a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x30, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x22, 0xad, 0x02, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x10, 0x0a, 0x03, 0x52, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x52,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x49, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x55, 0x70, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x47, 0x6f, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x47, 0x6f, 0x6c,
	0x64, 0x22, 0x69, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x2e, 0x0a, 0x07, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x12,
	0x2b, 0x0a, 0x06, 0x4e, 0x69, 0x75, 0x4e, 0x69, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x69, 0x75, 0x4e, 0x69, 0x75, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x4e, 0x69, 0x75, 0x4e, 0x69, 0x75, 0x22, 0x88, 0x05, 0x0a,
	0x0d, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26,
	0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x42, 0x61, 0x73, 0x65, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x42, 0x61, 0x73, 0x65, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x42, 0x61,
	0x73, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x5a, 0x69, 0x4d, 0x6f, 0x4a, 0x69, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x5a,
	0x69, 0x4d, 0x6f, 0x4a, 0x69, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x61, 0x6e, 0x55, 0x70, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x46, 0x61, 0x6e, 0x55,
	0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x61, 0x6e, 0x47, 0x61,
	0x6e, 0x67, 0x48, 0x75, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x44, 0x69, 0x61,
	0x6e, 0x47, 0x61, 0x6e, 0x67, 0x48, 0x75, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x48, 0x75, 0x61, 0x6e,
	0x53, 0x61, 0x6e, 0x5a, 0x68, 0x61, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x48, 0x75, 0x61, 0x6e, 0x53, 0x61, 0x6e, 0x5a, 0x68, 0x61, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09,
	0x59, 0x61, 0x6f, 0x4a, 0x69, 0x75, 0x44, 0x75, 0x69, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x59, 0x61, 0x6f, 0x4a, 0x69, 0x75, 0x44, 0x75, 0x69, 0x12, 0x2c, 0x0a, 0x11, 0x4d, 0x65,
	0x6e, 0x51, 0x69, 0x6e, 0x67, 0x5a, 0x68, 0x6f, 0x6e, 0x67, 0x5a, 0x68, 0x61, 0x6e, 0x67, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x4d, 0x65, 0x6e, 0x51, 0x69, 0x6e, 0x67, 0x5a, 0x68,
	0x6f, 0x6e, 0x67, 0x5a, 0x68, 0x61, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x61, 0x6e,
	0x44, 0x69, 0x48, 0x75, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x54, 0x69, 0x61, 0x6e,
	0x44, 0x69, 0x48, 0x75, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x69, 0x61, 0x6e, 0x50, 0x61, 0x6f, 0x50,
	0x69, 0x6e, 0x67, 0x48, 0x75, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x44, 0x69, 0x61,
	0x6e, 0x50, 0x61, 0x6f, 0x50, 0x69, 0x6e, 0x67, 0x48, 0x75, 0x12, 0x1a, 0x0a, 0x08, 0x4a, 0x69,
	0x61, 0x58, 0x69, 0x6e, 0x57, 0x75, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x4a, 0x69,
	0x61, 0x58, 0x69, 0x6e, 0x57, 0x75, 0x12, 0x32, 0x0a, 0x14, 0x47, 0x61, 0x6e, 0x67, 0x49, 0x6d,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x47, 0x61, 0x6e, 0x67, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x48, 0x75,
	0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x48, 0x75, 0x49, 0x6d, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x65, 0x6c, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x5a, 0x65, 0x72,
	0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x5a, 0x65, 0x72, 0x6f, 0x12, 0x1c, 0x0a, 0x09,
	0x42, 0x69, 0x67, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x42, 0x69, 0x67, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x06, 0x52, 0x65,
	0x42, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x06, 0x52, 0x65, 0x42, 0x61, 0x74, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x4e, 0x69, 0x75, 0x4e, 0x69,
	0x75, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x4d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4d, 0x69, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x4d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4d, 0x61,
	0x78, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x74, 0x69, 0x6f, 0x12, 0x2a, 0x0a, 0x10, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x47, 0x75,
	0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x4d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x47, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x52,
	0x65, 0x62, 0x61, 0x74, 0x65, 0x22, 0xc6, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x62, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4c,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x4c, 0x31, 0x12, 0x2c, 0x0a, 0x07, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x4c, 0x32, 0x12, 0x2c, 0x0a, 0x07, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x33, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x33,
	0x12, 0x2c, 0x0a, 0x07, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x34, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_base_proto_goTypes = []interface{}{
	(*Base)(nil),          // 0: outer.Base
	(*PlayerInfo)(nil),    // 1: outer.PlayerInfo
	(*GameParams)(nil),    // 2: outer.GameParams
	(*MahjongParams)(nil), // 3: outer.MahjongParams
	(*NiuNiuParams)(nil),  // 4: outer.NiuNiuParams
	(*RangeParams)(nil),   // 5: outer.RangeParams
	(*RebateParams)(nil),  // 6: outer.RebateParams
	(Position)(0),         // 7: outer.Position
}
var file_base_proto_depIdxs = []int32{
	7, // 0: outer.PlayerInfo.Position:type_name -> outer.Position
	3, // 1: outer.GameParams.Mahjong:type_name -> outer.MahjongParams
	4, // 2: outer.GameParams.NiuNiu:type_name -> outer.NiuNiuParams
	6, // 3: outer.MahjongParams.ReBate:type_name -> outer.RebateParams
	5, // 4: outer.RebateParams.RangeL1:type_name -> outer.RangeParams
	5, // 5: outer.RebateParams.RangeL2:type_name -> outer.RangeParams
	5, // 6: outer.RebateParams.RangeL3:type_name -> outer.RangeParams
	5, // 7: outer.RebateParams.RangeL4:type_name -> outer.RangeParams
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	file_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
		file_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfo); i {
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
		file_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameParams); i {
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
		file_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MahjongParams); i {
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
		file_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NiuNiuParams); i {
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
		file_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RangeParams); i {
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
		file_base_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RebateParams); i {
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
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
