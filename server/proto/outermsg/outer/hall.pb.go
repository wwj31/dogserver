// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
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

// 设置玩家基础信息
type SetRoleInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Icon   string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`      // 头像
	Gender int32  `protobuf:"varint,2,opt,name=gender,proto3" json:"gender,omitempty"` // 性别 0.男、1.女
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

	Icon   string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`      // 头像
	Gender int32  `protobuf:"varint,2,opt,name=gender,proto3" json:"gender,omitempty"` // 性别
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

	Phone    string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
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

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
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

	Password    string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
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

var File_hall_proto protoreflect.FileDescriptor

var file_hall_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x0c, 0x42, 0x69, 0x6e, 0x64, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x42, 0x69, 0x6e,
	0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22,
	0x51, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x73, 0x70, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_hall_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hall_proto_goTypes = []interface{}{
	(*SetRoleInfoReq)(nil),    // 0: outer.SetRoleInfoReq
	(*SetRoleInfoRsp)(nil),    // 1: outer.SetRoleInfoRsp
	(*BindPhoneReq)(nil),      // 2: outer.BindPhoneReq
	(*BindPhoneRsp)(nil),      // 3: outer.BindPhoneRsp
	(*ModifyPasswordReq)(nil), // 4: outer.ModifyPasswordReq
	(*ModifyPasswordRsp)(nil), // 5: outer.ModifyPasswordRsp
}
var file_hall_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hall_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hall_proto_goTypes,
		DependencyIndexes: file_hall_proto_depIdxs,
		MessageInfos:      file_hall_proto_msgTypes,
	}.Build()
	File_hall_proto = out.File
	file_hall_proto_rawDesc = nil
	file_hall_proto_goTypes = nil
	file_hall_proto_depIdxs = nil
}
