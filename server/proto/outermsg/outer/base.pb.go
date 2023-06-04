// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
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
	LoginAt    string   `protobuf:"bytes,8,opt,name=LoginAt,proto3" json:"LoginAt,omitempty"`                        // 登录时间
	LogoutAt   string   `protobuf:"bytes,9,opt,name=LogoutAt,proto3" json:"LogoutAt,omitempty"`                      // 离线时间
	UpShortId  int64    `protobuf:"varint,10,opt,name=UpShortId,proto3" json:"UpShortId,omitempty"`                  // 上级id
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

func (x *PlayerInfo) GetLoginAt() string {
	if x != nil {
		return x.LoginAt
	}
	return ""
}

func (x *PlayerInfo) GetLogoutAt() string {
	if x != nil {
		return x.LogoutAt
	}
	return ""
}

func (x *PlayerInfo) GetUpShortId() int64 {
	if x != nil {
		return x.UpShortId
	}
	return 0
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x1a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x30, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x99, 0x02, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
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
	0x69, 0x6e, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x55, 0x70, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x42, 0x08, 0x5a,
	0x06, 0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_base_proto_goTypes = []interface{}{
	(*Base)(nil),       // 0: outer.Base
	(*PlayerInfo)(nil), // 1: outer.PlayerInfo
	(Position)(0),      // 2: outer.Position
}
var file_base_proto_depIdxs = []int32{
	2, // 0: outer.PlayerInfo.Position:type_name -> outer.Position
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
