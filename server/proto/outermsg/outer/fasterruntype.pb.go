// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: fasterruntype.proto

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

// 跑得快 游戏状态机
type FasterRunState int32

const (
	FasterRunState_FasterRunStateReady      FasterRunState = 0 // 准备中状态
	FasterRunState_FasterRunStateDeal       FasterRunState = 1 // 游戏开始发牌状态
	FasterRunState_FasterRunStatePlaying    FasterRunState = 2 // 游戏状态
	FasterRunState_FasterRunStateSettlement FasterRunState = 3 // 结算
)

// Enum value maps for FasterRunState.
var (
	FasterRunState_name = map[int32]string{
		0: "FasterRunStateReady",
		1: "FasterRunStateDeal",
		2: "FasterRunStatePlaying",
		3: "FasterRunStateSettlement",
	}
	FasterRunState_value = map[string]int32{
		"FasterRunStateReady":      0,
		"FasterRunStateDeal":       1,
		"FasterRunStatePlaying":    2,
		"FasterRunStateSettlement": 3,
	}
)

func (x FasterRunState) Enum() *FasterRunState {
	p := new(FasterRunState)
	*p = x
	return p
}

func (x FasterRunState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FasterRunState) Descriptor() protoreflect.EnumDescriptor {
	return file_fasterruntype_proto_enumTypes[0].Descriptor()
}

func (FasterRunState) Type() protoreflect.EnumType {
	return &file_fasterruntype_proto_enumTypes[0]
}

func (x FasterRunState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FasterRunState.Descriptor instead.
func (FasterRunState) EnumDescriptor() ([]byte, []int) {
	return file_fasterruntype_proto_rawDescGZIP(), []int{0}
}

var File_fasterruntype_proto protoreflect.FileDescriptor

var file_fasterruntype_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x61, 0x73, 0x74, 0x65, 0x72, 0x72, 0x75, 0x6e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2a, 0x7a, 0x0a, 0x0e,
	0x46, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x0a, 0x13, 0x46, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x65, 0x61, 0x6c, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x46, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fasterruntype_proto_rawDescOnce sync.Once
	file_fasterruntype_proto_rawDescData = file_fasterruntype_proto_rawDesc
)

func file_fasterruntype_proto_rawDescGZIP() []byte {
	file_fasterruntype_proto_rawDescOnce.Do(func() {
		file_fasterruntype_proto_rawDescData = protoimpl.X.CompressGZIP(file_fasterruntype_proto_rawDescData)
	})
	return file_fasterruntype_proto_rawDescData
}

var file_fasterruntype_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fasterruntype_proto_goTypes = []interface{}{
	(FasterRunState)(0), // 0: outer.FasterRunState
}
var file_fasterruntype_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fasterruntype_proto_init() }
func file_fasterruntype_proto_init() {
	if File_fasterruntype_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fasterruntype_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fasterruntype_proto_goTypes,
		DependencyIndexes: file_fasterruntype_proto_depIdxs,
		EnumInfos:         file_fasterruntype_proto_enumTypes,
	}.Build()
	File_fasterruntype_proto = out.File
	file_fasterruntype_proto_rawDesc = nil
	file_fasterruntype_proto_goTypes = nil
	file_fasterruntype_proto_depIdxs = nil
}