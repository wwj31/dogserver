// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: type.proto

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

type Position int32

const (
	Position_NoAlliance   Position = 0
	Position_Normal       Position = 1 // 普通成员
	Position_Captain      Position = 2 // 队长
	Position_Manager      Position = 3 // 管理员
	Position_DeputyMaster Position = 4 // 副盟主
	Position_Master       Position = 5 // 盟主
)

// Enum value maps for Position.
var (
	Position_name = map[int32]string{
		0: "NoAlliance",
		1: "Normal",
		2: "Captain",
		3: "Manager",
		4: "DeputyMaster",
		5: "Master",
	}
	Position_value = map[string]int32{
		"NoAlliance":   0,
		"Normal":       1,
		"Captain":      2,
		"Manager":      3,
		"DeputyMaster": 4,
		"Master":       5,
	}
)

func (x Position) Enum() *Position {
	p := new(Position)
	*p = x
	return p
}

func (x Position) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Position) Descriptor() protoreflect.EnumDescriptor {
	return file_type_proto_enumTypes[0].Descriptor()
}

func (Position) Type() protoreflect.EnumType {
	return &file_type_proto_enumTypes[0]
}

func (x Position) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Position.Descriptor instead.
func (Position) EnumDescriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{0}
}

var File_type_proto protoreflect.FileDescriptor

var file_type_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x2a, 0x5e, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x0a, 0x4e, 0x6f, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x61, 0x70, 0x74, 0x61, 0x69, 0x6e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x75, 0x74, 0x79, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x10, 0x05, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_type_proto_rawDescOnce sync.Once
	file_type_proto_rawDescData = file_type_proto_rawDesc
)

func file_type_proto_rawDescGZIP() []byte {
	file_type_proto_rawDescOnce.Do(func() {
		file_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_type_proto_rawDescData)
	})
	return file_type_proto_rawDescData
}

var file_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_type_proto_goTypes = []interface{}{
	(Position)(0), // 0: outer.Position
}
var file_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_type_proto_init() }
func file_type_proto_init() {
	if File_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_type_proto_goTypes,
		DependencyIndexes: file_type_proto_depIdxs,
		EnumInfos:         file_type_proto_enumTypes,
	}.Build()
	File_type_proto = out.File
	file_type_proto_rawDesc = nil
	file_type_proto_goTypes = nil
	file_type_proto_depIdxs = nil
}
