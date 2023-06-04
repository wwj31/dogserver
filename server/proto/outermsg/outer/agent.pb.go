// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
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

// 获取上、下级信息
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
	0x6f, 0x77, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_agent_proto_goTypes = []interface{}{
	(*AgentMembersReq)(nil), // 0: outer.AgentMembersReq
	(*AgentMembersRsp)(nil), // 1: outer.AgentMembersRsp
	(*PlayerInfo)(nil),      // 2: outer.PlayerInfo
}
var file_agent_proto_depIdxs = []int32{
	2, // 0: outer.AgentMembersRsp.UpMember:type_name -> outer.PlayerInfo
	2, // 1: outer.AgentMembersRsp.DownMembers:type_name -> outer.PlayerInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
