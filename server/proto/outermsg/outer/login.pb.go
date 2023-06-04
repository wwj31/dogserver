// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.2
// source: login.proto

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

type HeartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientTimestamp int64 `protobuf:"varint,1,opt,name=ClientTimestamp,proto3" json:"ClientTimestamp,omitempty"`
}

func (x *HeartReq) Reset() {
	*x = HeartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartReq) ProtoMessage() {}

func (x *HeartReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartReq.ProtoReflect.Descriptor instead.
func (*HeartReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *HeartReq) GetClientTimestamp() int64 {
	if x != nil {
		return x.ClientTimestamp
	}
	return 0
}

type HeartRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientTimestamp int64 `protobuf:"varint,1,opt,name=ClientTimestamp,proto3" json:"ClientTimestamp,omitempty"`
	ServerTimestamp int64 `protobuf:"varint,2,opt,name=ServerTimestamp,proto3" json:"ServerTimestamp,omitempty"`
}

func (x *HeartRsp) Reset() {
	*x = HeartRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartRsp) ProtoMessage() {}

func (x *HeartRsp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartRsp.ProtoReflect.Descriptor instead.
func (*HeartRsp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *HeartRsp) GetClientTimestamp() int64 {
	if x != nil {
		return x.ClientTimestamp
	}
	return 0
}

func (x *HeartRsp) GetServerTimestamp() int64 {
	if x != nil {
		return x.ServerTimestamp
	}
	return 0
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginType     int32  `protobuf:"varint,1,opt,name=LoginType,proto3" json:"LoginType,omitempty"`        // 登录方式 1.游客、2.phone、3.微信、4.token
	WeiXinOpenID  string `protobuf:"bytes,2,opt,name=WeiXinOpenID,proto3" json:"WeiXinOpenID,omitempty"`   // 微信的Openid
	DeviceID      string `protobuf:"bytes,3,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`           // 设备ID
	Phone         string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`                 // 电话登录
	PhonePassword string `protobuf:"bytes,5,opt,name=PhonePassword,proto3" json:"PhonePassword,omitempty"` // 绑定的密码
	Token         string `protobuf:"bytes,6,opt,name=Token,proto3" json:"Token,omitempty"`                 // token
	OS            string `protobuf:"bytes,7,opt,name=OS,proto3" json:"OS,omitempty"`                       // 系统
	ClientVersion string `protobuf:"bytes,8,opt,name=ClientVersion,proto3" json:"ClientVersion,omitempty"` // 版本号
	UpShortId     int64  `protobuf:"varint,9,opt,name=UpShortId,proto3" json:"UpShortId,omitempty"`        // 上级shortId
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

func (x *LoginReq) GetLoginType() int32 {
	if x != nil {
		return x.LoginType
	}
	return 0
}

func (x *LoginReq) GetWeiXinOpenID() string {
	if x != nil {
		return x.WeiXinOpenID
	}
	return ""
}

func (x *LoginReq) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *LoginReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *LoginReq) GetPhonePassword() string {
	if x != nil {
		return x.PhonePassword
	}
	return ""
}

func (x *LoginReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginReq) GetOS() string {
	if x != nil {
		return x.OS
	}
	return ""
}

func (x *LoginReq) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

func (x *LoginReq) GetUpShortId() int64 {
	if x != nil {
		return x.UpShortId
	}
	return 0
}

type LoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RID       string `protobuf:"bytes,1,opt,name=RID,proto3" json:"RID,omitempty"` // role id
	NewPlayer bool   `protobuf:"varint,2,opt,name=NewPlayer,proto3" json:"NewPlayer,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	Checksum  string `protobuf:"bytes,4,opt,name=Checksum,proto3" json:"Checksum,omitempty"`
}

func (x *LoginRsp) Reset() {
	*x = LoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRsp) ProtoMessage() {}

func (x *LoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRsp.ProtoReflect.Descriptor instead.
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRsp) GetRID() string {
	if x != nil {
		return x.RID
	}
	return ""
}

func (x *LoginRsp) GetNewPlayer() bool {
	if x != nil {
		return x.NewPlayer
	}
	return false
}

func (x *LoginRsp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginRsp) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

type EnterGameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RID       string `protobuf:"bytes,1,opt,name=RID,proto3" json:"RID,omitempty"`
	NewPlayer bool   `protobuf:"varint,2,opt,name=NewPlayer,proto3" json:"NewPlayer,omitempty"`
	Checksum  string `protobuf:"bytes,3,opt,name=Checksum,proto3" json:"Checksum,omitempty"`
}

func (x *EnterGameReq) Reset() {
	*x = EnterGameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterGameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterGameReq) ProtoMessage() {}

func (x *EnterGameReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterGameReq.ProtoReflect.Descriptor instead.
func (*EnterGameReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{4}
}

func (x *EnterGameReq) GetRID() string {
	if x != nil {
		return x.RID
	}
	return ""
}

func (x *EnterGameReq) GetNewPlayer() bool {
	if x != nil {
		return x.NewPlayer
	}
	return false
}

func (x *EnterGameReq) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

type EnterGameRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否是新玩家
	NewPlayer bool      `protobuf:"varint,1,opt,name=NewPlayer,proto3" json:"NewPlayer,omitempty"`
	RoleInfo  *RoleInfo `protobuf:"bytes,2,opt,name=roleInfo,proto3" json:"roleInfo,omitempty"` // 角色信息
}

func (x *EnterGameRsp) Reset() {
	*x = EnterGameRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterGameRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterGameRsp) ProtoMessage() {}

func (x *EnterGameRsp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterGameRsp.ProtoReflect.Descriptor instead.
func (*EnterGameRsp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{5}
}

func (x *EnterGameRsp) GetNewPlayer() bool {
	if x != nil {
		return x.NewPlayer
	}
	return false
}

func (x *EnterGameRsp) GetRoleInfo() *RoleInfo {
	if x != nil {
		return x.RoleInfo
	}
	return nil
}

// 角色信息
type RoleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RID        string   `protobuf:"bytes,1,opt,name=RID,proto3" json:"RID,omitempty"`          // 玩家id
	ShortId    int64    `protobuf:"varint,2,opt,name=ShortId,proto3" json:"ShortId,omitempty"` // 短id
	Phone      string   `protobuf:"bytes,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Name       string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Icon       string   `protobuf:"bytes,5,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Gender     int32    `protobuf:"varint,6,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Gold       int64    `protobuf:"varint,7,opt,name=Gold,proto3" json:"Gold,omitempty"`
	AllianceId int32    `protobuf:"varint,8,opt,name=AllianceId,proto3" json:"AllianceId,omitempty"`                 // 联盟id
	Position   Position `protobuf:"varint,9,opt,name=Position,proto3,enum=outer.Position" json:"Position,omitempty"` // 职位
}

func (x *RoleInfo) Reset() {
	*x = RoleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleInfo) ProtoMessage() {}

func (x *RoleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleInfo.ProtoReflect.Descriptor instead.
func (*RoleInfo) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{6}
}

func (x *RoleInfo) GetRID() string {
	if x != nil {
		return x.RID
	}
	return ""
}

func (x *RoleInfo) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

func (x *RoleInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RoleInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleInfo) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *RoleInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *RoleInfo) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

func (x *RoleInfo) GetAllianceId() int32 {
	if x != nil {
		return x.AllianceId
	}
	return 0
}

func (x *RoleInfo) GetPosition() Position {
	if x != nil {
		return x.Position
	}
	return Position_NoAlliance
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x1a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x34, 0x0a, 0x08, 0x48, 0x65, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x0f,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x5e, 0x0a, 0x08, 0x48, 0x65, 0x61, 0x72, 0x74, 0x52,
	0x73, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x0f,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x8e, 0x02, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x57, 0x65, 0x69, 0x58, 0x69, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x57, 0x65, 0x69, 0x58, 0x69, 0x6e, 0x4f,
	0x70, 0x65, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x53, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x4f, 0x53, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x52, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x75, 0x6d, 0x22, 0x5a, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x52, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x4e, 0x65, 0x77, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75,
	0x6d, 0x22, 0x59, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x73,
	0x70, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x2b, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xed, 0x01, 0x0a,
	0x08, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x52, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49,
	0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x47,
	0x6f, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x47, 0x6f, 0x6c, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x2b, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x5a, 0x06,
	0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_login_proto_goTypes = []interface{}{
	(*HeartReq)(nil),     // 0: outer.HeartReq
	(*HeartRsp)(nil),     // 1: outer.HeartRsp
	(*LoginReq)(nil),     // 2: outer.LoginReq
	(*LoginRsp)(nil),     // 3: outer.LoginRsp
	(*EnterGameReq)(nil), // 4: outer.EnterGameReq
	(*EnterGameRsp)(nil), // 5: outer.EnterGameRsp
	(*RoleInfo)(nil),     // 6: outer.RoleInfo
	(Position)(0),        // 7: outer.Position
}
var file_login_proto_depIdxs = []int32{
	6, // 0: outer.EnterGameRsp.roleInfo:type_name -> outer.RoleInfo
	7, // 1: outer.RoleInfo.Position:type_name -> outer.Position
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	file_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartReq); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartRsp); i {
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
		file_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRsp); i {
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
		file_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterGameReq); i {
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
		file_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterGameRsp); i {
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
		file_login_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleInfo); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
