//Code generated by msgidgen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.0
// source: msgid.proto

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

type Msg int32

const (
	Msg_IdUnknown              Msg = 0
	Msg_IdAgentMembersReq      Msg = 1894584925 // dispatch to player
	Msg_IdAgentMembersRsp      Msg = 1592587688
	Msg_IdSetMemberPositionReq Msg = 1891047110
	Msg_IdSetMemberPositionRsp Msg = 1656160475
	Msg_IdFailRsp              Msg = 160109657
	Msg_IdSetRoleInfoReq       Msg = 58140470 // dispatch to player
	Msg_IdSetRoleInfoRsp       Msg = 1903626880
	Msg_IdBindPhoneReq         Msg = 403639403 // dispatch to player
	Msg_IdBindPhoneRsp         Msg = 34531592
	Msg_IdModifyPasswordReq    Msg = 655389589 // dispatch to player
	Msg_IdModifyPasswordRsp    Msg = 890276254
	Msg_IdHeartReq             Msg = 31040569
	Msg_IdHeartRsp             Msg = 400148124
	Msg_IdLoginReq             Msg = 1510628254 // dispatch to login
	Msg_IdLoginRsp             Msg = 1275741587
	Msg_IdEnterGameReq         Msg = 117622385 // dispatch to player
	Msg_IdEnterGameRsp         Msg = 2097329717
	Msg_IdMailListReq          Msg = 1443645576 // dispatch to player
	Msg_IdMailListRsp          Msg = 1812753385
	Msg_IdReadMailReq          Msg = 1963047789 // dispatch to player
	Msg_IdReadMailRsp          Msg = 1661050550
	Msg_IdReceiveMailItemReq   Msg = 1679531254 // dispatch to player
	Msg_IdReceiveMailItemRsp   Msg = 1981528235
	Msg_IdDeleteMailReq        Msg = 1054436458 // dispatch to player
)

// Enum value maps for Msg.
var (
	Msg_name = map[int32]string{
		0:          "IdUnknown",
		1894584925: "IdAgentMembersReq",
		1592587688: "IdAgentMembersRsp",
		1891047110: "IdSetMemberPositionReq",
		1656160475: "IdSetMemberPositionRsp",
		160109657:  "IdFailRsp",
		58140470:   "IdSetRoleInfoReq",
		1903626880: "IdSetRoleInfoRsp",
		403639403:  "IdBindPhoneReq",
		34531592:   "IdBindPhoneRsp",
		655389589:  "IdModifyPasswordReq",
		890276254:  "IdModifyPasswordRsp",
		31040569:   "IdHeartReq",
		400148124:  "IdHeartRsp",
		1510628254: "IdLoginReq",
		1275741587: "IdLoginRsp",
		117622385:  "IdEnterGameReq",
		2097329717: "IdEnterGameRsp",
		1443645576: "IdMailListReq",
		1812753385: "IdMailListRsp",
		1963047789: "IdReadMailReq",
		1661050550: "IdReadMailRsp",
		1679531254: "IdReceiveMailItemReq",
		1981528235: "IdReceiveMailItemRsp",
		1054436458: "IdDeleteMailReq",
	}
	Msg_value = map[string]int32{
		"IdUnknown":              0,
		"IdAgentMembersReq":      1894584925,
		"IdAgentMembersRsp":      1592587688,
		"IdSetMemberPositionReq": 1891047110,
		"IdSetMemberPositionRsp": 1656160475,
		"IdFailRsp":              160109657,
		"IdSetRoleInfoReq":       58140470,
		"IdSetRoleInfoRsp":       1903626880,
		"IdBindPhoneReq":         403639403,
		"IdBindPhoneRsp":         34531592,
		"IdModifyPasswordReq":    655389589,
		"IdModifyPasswordRsp":    890276254,
		"IdHeartReq":             31040569,
		"IdHeartRsp":             400148124,
		"IdLoginReq":             1510628254,
		"IdLoginRsp":             1275741587,
		"IdEnterGameReq":         117622385,
		"IdEnterGameRsp":         2097329717,
		"IdMailListReq":          1443645576,
		"IdMailListRsp":          1812753385,
		"IdReadMailReq":          1963047789,
		"IdReadMailRsp":          1661050550,
		"IdReceiveMailItemReq":   1679531254,
		"IdReceiveMailItemRsp":   1981528235,
		"IdDeleteMailReq":        1054436458,
	}
)

func (x Msg) Enum() *Msg {
	p := new(Msg)
	*p = x
	return p
}

func (x Msg) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Msg) Descriptor() protoreflect.EnumDescriptor {
	return file_msgid_proto_enumTypes[0].Descriptor()
}

func (Msg) Type() protoreflect.EnumType {
	return &file_msgid_proto_enumTypes[0]
}

func (x Msg) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Msg.Descriptor instead.
func (Msg) EnumDescriptor() ([]byte, []int) {
	return file_msgid_proto_rawDescGZIP(), []int{0}
}

var File_msgid_proto protoreflect.FileDescriptor

var file_msgid_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x73, 0x67, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2a, 0xe7, 0x04, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x0d, 0x0a, 0x09,
	0x49, 0x64, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x11, 0x49,
	0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x10, 0xdd, 0xa4, 0xb4, 0x87, 0x07, 0x12, 0x19, 0x0a, 0x11, 0x49, 0x64, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x73, 0x70, 0x10, 0xa8, 0xeb, 0xb3, 0xf7,
	0x05, 0x12, 0x1e, 0x0a, 0x16, 0x49, 0x64, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x10, 0xc6, 0xad, 0xdc, 0x85,
	0x07, 0x12, 0x1e, 0x0a, 0x16, 0x49, 0x64, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x10, 0xdb, 0x81, 0xdc, 0x95,
	0x06, 0x12, 0x10, 0x0a, 0x09, 0x49, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x10, 0xd9,
	0xa8, 0xac, 0x4c, 0x12, 0x17, 0x0a, 0x10, 0x49, 0x64, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x10, 0xb6, 0xce, 0xdc, 0x1b, 0x12, 0x18, 0x0a, 0x10,
	0x49, 0x64, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70,
	0x10, 0x80, 0x95, 0xdc, 0x8b, 0x07, 0x12, 0x16, 0x0a, 0x0e, 0x49, 0x64, 0x42, 0x69, 0x6e, 0x64,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x10, 0xeb, 0x98, 0xbc, 0xc0, 0x01, 0x12, 0x15,
	0x0a, 0x0e, 0x49, 0x64, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x73, 0x70,
	0x10, 0x88, 0xd2, 0xbb, 0x10, 0x12, 0x1b, 0x0a, 0x13, 0x49, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x10, 0x95, 0xe7, 0xc1,
	0xb8, 0x02, 0x12, 0x1b, 0x0a, 0x13, 0x49, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x73, 0x70, 0x10, 0x9e, 0x93, 0xc2, 0xa8, 0x03, 0x12,
	0x11, 0x0a, 0x0a, 0x49, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x10, 0xb9, 0xc8,
	0xe6, 0x0e, 0x12, 0x12, 0x0a, 0x0a, 0x49, 0x64, 0x48, 0x65, 0x61, 0x72, 0x74, 0x52, 0x73, 0x70,
	0x10, 0x9c, 0x8d, 0xe7, 0xbe, 0x01, 0x12, 0x12, 0x0a, 0x0a, 0x49, 0x64, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x10, 0x9e, 0xb7, 0xa9, 0xd0, 0x05, 0x12, 0x12, 0x0a, 0x0a, 0x49, 0x64,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x10, 0x93, 0x8b, 0xa9, 0xe0, 0x04, 0x12, 0x15,
	0x0a, 0x0e, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x10, 0xf1, 0x8c, 0x8b, 0x38, 0x12, 0x16, 0x0a, 0x0e, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x73, 0x70, 0x10, 0xb5, 0xec, 0x8a, 0xe8, 0x07, 0x12, 0x15, 0x0a,
	0x0d, 0x49, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x10, 0x88,
	0x91, 0xb1, 0xb0, 0x05, 0x12, 0x15, 0x0a, 0x0d, 0x49, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x73, 0x70, 0x10, 0xe9, 0xd7, 0xb1, 0xe0, 0x06, 0x12, 0x15, 0x0a, 0x0d, 0x49,
	0x64, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x10, 0xed, 0xf6, 0x86,
	0xa8, 0x07, 0x12, 0x15, 0x0a, 0x0d, 0x49, 0x64, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x69, 0x6c,
	0x52, 0x73, 0x70, 0x10, 0xb6, 0xbd, 0x86, 0x98, 0x06, 0x12, 0x1c, 0x0a, 0x14, 0x49, 0x64, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x10, 0xf6, 0xb9, 0xee, 0xa0, 0x06, 0x12, 0x1c, 0x0a, 0x14, 0x49, 0x64, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x73, 0x70, 0x10,
	0xab, 0xf1, 0xee, 0xb0, 0x07, 0x12, 0x17, 0x0a, 0x0f, 0x49, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x10, 0xea, 0xd8, 0xe5, 0xf6, 0x03, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msgid_proto_rawDescOnce sync.Once
	file_msgid_proto_rawDescData = file_msgid_proto_rawDesc
)

func file_msgid_proto_rawDescGZIP() []byte {
	file_msgid_proto_rawDescOnce.Do(func() {
		file_msgid_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgid_proto_rawDescData)
	})
	return file_msgid_proto_rawDescData
}

var file_msgid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msgid_proto_goTypes = []interface{}{
	(Msg)(0), // 0: outer.Msg
}
var file_msgid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msgid_proto_init() }
func file_msgid_proto_init() {
	if File_msgid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msgid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msgid_proto_goTypes,
		DependencyIndexes: file_msgid_proto_depIdxs,
		EnumInfos:         file_msgid_proto_enumTypes,
	}.Build()
	File_msgid_proto = out.File
	file_msgid_proto_rawDesc = nil
	file_msgid_proto_goTypes = nil
	file_msgid_proto_depIdxs = nil
}
