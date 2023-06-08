// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: alliance.proto

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

// 通知联盟信息发送变化、加入联盟、职位变更、联盟解散
type AllianceInfoNtf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllianceId int32 `protobuf:"varint,1,opt,name=AllianceId,proto3" json:"AllianceId,omitempty"`
	Position   int32 `protobuf:"varint,2,opt,name=Position,proto3" json:"Position,omitempty"`
}

func (x *AllianceInfoNtf) Reset() {
	*x = AllianceInfoNtf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alliance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllianceInfoNtf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllianceInfoNtf) ProtoMessage() {}

func (x *AllianceInfoNtf) ProtoReflect() protoreflect.Message {
	mi := &file_alliance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllianceInfoNtf.ProtoReflect.Descriptor instead.
func (*AllianceInfoNtf) Descriptor() ([]byte, []int) {
	return file_alliance_proto_rawDescGZIP(), []int{0}
}

func (x *AllianceInfoNtf) GetAllianceId() int32 {
	if x != nil {
		return x.AllianceId
	}
	return 0
}

func (x *AllianceInfoNtf) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

// 盟主请求解散联盟
type DisbandAllianceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisbandAllianceReq) Reset() {
	*x = DisbandAllianceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alliance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisbandAllianceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisbandAllianceReq) ProtoMessage() {}

func (x *DisbandAllianceReq) ProtoReflect() protoreflect.Message {
	mi := &file_alliance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisbandAllianceReq.ProtoReflect.Descriptor instead.
func (*DisbandAllianceReq) Descriptor() ([]byte, []int) {
	return file_alliance_proto_rawDescGZIP(), []int{1}
}

type DisbandAllianceRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisbandAllianceRsp) Reset() {
	*x = DisbandAllianceRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alliance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisbandAllianceRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisbandAllianceRsp) ProtoMessage() {}

func (x *DisbandAllianceRsp) ProtoReflect() protoreflect.Message {
	mi := &file_alliance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisbandAllianceRsp.ProtoReflect.Descriptor instead.
func (*DisbandAllianceRsp) Descriptor() ([]byte, []int) {
	return file_alliance_proto_rawDescGZIP(), []int{2}
}

// 邀请加入联盟&绑定上下级关系
type InviteAllianceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId int64 `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"` // 邀请的玩家短id
}

func (x *InviteAllianceReq) Reset() {
	*x = InviteAllianceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alliance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteAllianceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteAllianceReq) ProtoMessage() {}

func (x *InviteAllianceReq) ProtoReflect() protoreflect.Message {
	mi := &file_alliance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteAllianceReq.ProtoReflect.Descriptor instead.
func (*InviteAllianceReq) Descriptor() ([]byte, []int) {
	return file_alliance_proto_rawDescGZIP(), []int{3}
}

func (x *InviteAllianceReq) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

type InviteAllianceRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InviteAllianceRsp) Reset() {
	*x = InviteAllianceRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alliance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteAllianceRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteAllianceRsp) ProtoMessage() {}

func (x *InviteAllianceRsp) ProtoReflect() protoreflect.Message {
	mi := &file_alliance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteAllianceRsp.ProtoReflect.Descriptor instead.
func (*InviteAllianceRsp) Descriptor() ([]byte, []int) {
	return file_alliance_proto_rawDescGZIP(), []int{4}
}

// 设置成员职位级别
type SetMemberPositionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId  int64    `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"` // 被设置成员短Id
	Position Position `protobuf:"varint,2,opt,name=Position,proto3,enum=outer.Position" json:"Position,omitempty"`
}

func (x *SetMemberPositionReq) Reset() {
	*x = SetMemberPositionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alliance_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMemberPositionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMemberPositionReq) ProtoMessage() {}

func (x *SetMemberPositionReq) ProtoReflect() protoreflect.Message {
	mi := &file_alliance_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMemberPositionReq.ProtoReflect.Descriptor instead.
func (*SetMemberPositionReq) Descriptor() ([]byte, []int) {
	return file_alliance_proto_rawDescGZIP(), []int{5}
}

func (x *SetMemberPositionReq) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

func (x *SetMemberPositionReq) GetPosition() Position {
	if x != nil {
		return x.Position
	}
	return Position_NoAlliance
}

type SetMemberPositionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId  int64    `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"` // 被设置成员短Id
	Position Position `protobuf:"varint,2,opt,name=Position,proto3,enum=outer.Position" json:"Position,omitempty"`
}

func (x *SetMemberPositionRsp) Reset() {
	*x = SetMemberPositionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alliance_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMemberPositionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMemberPositionRsp) ProtoMessage() {}

func (x *SetMemberPositionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_alliance_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMemberPositionRsp.ProtoReflect.Descriptor instead.
func (*SetMemberPositionRsp) Descriptor() ([]byte, []int) {
	return file_alliance_proto_rawDescGZIP(), []int{6}
}

func (x *SetMemberPositionRsp) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

func (x *SetMemberPositionRsp) GetPosition() Position {
	if x != nil {
		return x.Position
	}
	return Position_NoAlliance
}

// 踢出联盟
type KickOutMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortId int64 `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"` // 被设置成员短Id
}

func (x *KickOutMemberReq) Reset() {
	*x = KickOutMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alliance_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickOutMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickOutMemberReq) ProtoMessage() {}

func (x *KickOutMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_alliance_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickOutMemberReq.ProtoReflect.Descriptor instead.
func (*KickOutMemberReq) Descriptor() ([]byte, []int) {
	return file_alliance_proto_rawDescGZIP(), []int{7}
}

func (x *KickOutMemberReq) GetShortId() int64 {
	if x != nil {
		return x.ShortId
	}
	return 0
}

type KickOutMemberRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KickOutMemberRsp) Reset() {
	*x = KickOutMemberRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alliance_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickOutMemberRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickOutMemberRsp) ProtoMessage() {}

func (x *KickOutMemberRsp) ProtoReflect() protoreflect.Message {
	mi := &file_alliance_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickOutMemberRsp.ProtoReflect.Descriptor instead.
func (*KickOutMemberRsp) Descriptor() ([]byte, []int) {
	return file_alliance_proto_rawDescGZIP(), []int{8}
}

var File_alliance_proto protoreflect.FileDescriptor

var file_alliance_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x1a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x4e, 0x74, 0x66, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x41, 0x6c, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x62, 0x61, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x62,
	0x61, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x73, 0x70, 0x22, 0x2d,
	0x0a, 0x11, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x13, 0x0a,
	0x11, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x73, 0x70, 0x22, 0x5d, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x5d, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x2c, 0x0a, 0x10, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x12,
	0x0a, 0x10, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x73, 0x70, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alliance_proto_rawDescOnce sync.Once
	file_alliance_proto_rawDescData = file_alliance_proto_rawDesc
)

func file_alliance_proto_rawDescGZIP() []byte {
	file_alliance_proto_rawDescOnce.Do(func() {
		file_alliance_proto_rawDescData = protoimpl.X.CompressGZIP(file_alliance_proto_rawDescData)
	})
	return file_alliance_proto_rawDescData
}

var file_alliance_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_alliance_proto_goTypes = []interface{}{
	(*AllianceInfoNtf)(nil),      // 0: outer.AllianceInfoNtf
	(*DisbandAllianceReq)(nil),   // 1: outer.DisbandAllianceReq
	(*DisbandAllianceRsp)(nil),   // 2: outer.DisbandAllianceRsp
	(*InviteAllianceReq)(nil),    // 3: outer.InviteAllianceReq
	(*InviteAllianceRsp)(nil),    // 4: outer.InviteAllianceRsp
	(*SetMemberPositionReq)(nil), // 5: outer.SetMemberPositionReq
	(*SetMemberPositionRsp)(nil), // 6: outer.SetMemberPositionRsp
	(*KickOutMemberReq)(nil),     // 7: outer.KickOutMemberReq
	(*KickOutMemberRsp)(nil),     // 8: outer.KickOutMemberRsp
	(Position)(0),                // 9: outer.Position
}
var file_alliance_proto_depIdxs = []int32{
	9, // 0: outer.SetMemberPositionReq.Position:type_name -> outer.Position
	9, // 1: outer.SetMemberPositionRsp.Position:type_name -> outer.Position
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_alliance_proto_init() }
func file_alliance_proto_init() {
	if File_alliance_proto != nil {
		return
	}
	file_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alliance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllianceInfoNtf); i {
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
		file_alliance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisbandAllianceReq); i {
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
		file_alliance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisbandAllianceRsp); i {
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
		file_alliance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteAllianceReq); i {
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
		file_alliance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteAllianceRsp); i {
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
		file_alliance_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMemberPositionReq); i {
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
		file_alliance_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMemberPositionRsp); i {
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
		file_alliance_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickOutMemberReq); i {
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
		file_alliance_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickOutMemberRsp); i {
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
			RawDescriptor: file_alliance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alliance_proto_goTypes,
		DependencyIndexes: file_alliance_proto_depIdxs,
		MessageInfos:      file_alliance_proto_msgTypes,
	}.Build()
	File_alliance_proto = out.File
	file_alliance_proto_rawDesc = nil
	file_alliance_proto_goTypes = nil
	file_alliance_proto_depIdxs = nil
}
