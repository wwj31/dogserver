//Code generated by msgidgen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
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
	Msg_IdUnknown                      Msg = 0
	Msg_IdAgentMembersReq              Msg = 1894584925 // dispatch to player
	Msg_IdAgentMembersRsp              Msg = 1592587688
	Msg_IdAllianceInfoNtf              Msg = 672428706
	Msg_IdDisbandAllianceReq           Msg = 650636160 // dispatch to player
	Msg_IdDisbandAllianceRsp           Msg = 952633399
	Msg_IdSearchPlayerInfoReq          Msg = 688893070 // dispatch to player
	Msg_IdSearchPlayerInfoRsp          Msg = 990890275
	Msg_IdInviteAllianceReq            Msg = 710800074 // dispatch to player
	Msg_IdInviteAllianceRsp            Msg = 543023757
	Msg_IdSetMemberPositionReq         Msg = 1891047110 // dispatch to player
	Msg_IdSetMemberPositionRsp         Msg = 1656160475
	Msg_IdKickOutMemberReq             Msg = 1475199389 // dispatch to player
	Msg_IdKickOutMemberRsp             Msg = 1173202408
	Msg_IdFailRsp                      Msg = 160109657
	Msg_IdSetRoleInfoReq               Msg = 58140470 // dispatch to player
	Msg_IdSetRoleInfoRsp               Msg = 1903626880
	Msg_IdBindPhoneReq                 Msg = 403639403 // dispatch to player
	Msg_IdBindPhoneRsp                 Msg = 34531592
	Msg_IdModifyPasswordReq            Msg = 655389589 // dispatch to player
	Msg_IdModifyPasswordRsp            Msg = 890276254
	Msg_IdHeartReq                     Msg = 31040569
	Msg_IdHeartRsp                     Msg = 400148124
	Msg_IdLoginReq                     Msg = 1510628254 // dispatch to login
	Msg_IdLoginRsp                     Msg = 1275741587
	Msg_IdEnterGameReq                 Msg = 117622385 // dispatch to player
	Msg_IdEnterGameRsp                 Msg = 2097329717
	Msg_IdMahjongBTEReadyReq           Msg = 39989867 // dispatch to gambling
	Msg_IdMahjongBTEReadyRsp           Msg = 207765930
	Msg_IdMahjongBTEExchange3Req       Msg = 983022052 // dispatch to gambling
	Msg_IdMahjongBTEExchange3Rsp       Msg = 681024845
	Msg_IdMahjongBTEDecideIgnoreReq    Msg = 1993774143 // dispatch to gambling
	Msg_IdMahjongBTEDecideIgnoreRsp    Msg = 1758887476
	Msg_IdMahjongBTEPlayCardReq        Msg = 2138714649 // dispatch to gambling
	Msg_IdMahjongBTEPlayCardRsp        Msg = 1970938586
	Msg_IdMahjongBTEOperateReq         Msg = 1792910030 // dispatch to gambling
	Msg_IdMahjongBTEOperateRsp         Msg = 2094907235
	Msg_IdMahjongBTEPlayerReadyNtf     Msg = 176653714
	Msg_IdMahjongBTEReadyNtf           Msg = 23903615
	Msg_IdMahjongBTEDecideMasterNtf    Msg = 701290466
	Msg_IdMahjongBTEDealNtf            Msg = 881890379
	Msg_IdMahjongBTEExchange3Ntf       Msg = 735992092
	Msg_IdMahjongBTEExchange3EndNtf    Msg = 595136049
	Msg_IdMahjongBTEDecideIgnoreNtf    Msg = 1745361179
	Msg_IdMahjongBTEDecideIgnoreEndNtf Msg = 1299626231
	Msg_IdMahjongBTEPlayingNtf         Msg = 116867344
	Msg_IdMahjongBTETurnNtf            Msg = 1486690182
	Msg_IdMahjongBTEOperaNtf           Msg = 1682707103
	Msg_IdMahjongBTESettlementNtf      Msg = 1808167222
	Msg_IdMailListReq                  Msg = 1443645576 // dispatch to player
	Msg_IdMailListRsp                  Msg = 1812753385
	Msg_IdReadMailReq                  Msg = 1963047789 // dispatch to player
	Msg_IdReadMailRsp                  Msg = 1661050550
	Msg_IdReceiveMailItemReq           Msg = 1679531254 // dispatch to player
	Msg_IdReceiveMailItemRsp           Msg = 1981528235
	Msg_IdDeleteMailReq                Msg = 1054436458 // dispatch to player
	Msg_IdCreateRoomReq                Msg = 2001014089 // dispatch to player
	Msg_IdCreateRoomRsp                Msg = 222637993
	Msg_IdDisbandRoomReq               Msg = 1782747149 // dispatch to player
	Msg_IdDisbandRoomRsp               Msg = 2017634040
	Msg_IdRoomListReq                  Msg = 21167689 // dispatch to player
	Msg_IdRoomListRsp                  Msg = 390275242
	Msg_IdJoinRoomReq                  Msg = 1530633489 // dispatch to player
	Msg_IdJoinRoomRsp                  Msg = 1899741298
	Msg_IdLeaveRoomReq                 Msg = 1875949722 // dispatch to player
	Msg_IdLeaveRoomRsp                 Msg = 97573884
	Msg_IdRoomPlayerEnterNtf           Msg = 867132366
	Msg_IdRoomPlayerLeaveNtf           Msg = 721409295
	Msg_IdRoomPlayerOnlineNtf          Msg = 1386944282
)

// Enum value maps for Msg.
var (
	Msg_name = map[int32]string{
		0:          "IdUnknown",
		1894584925: "IdAgentMembersReq",
		1592587688: "IdAgentMembersRsp",
		672428706:  "IdAllianceInfoNtf",
		650636160:  "IdDisbandAllianceReq",
		952633399:  "IdDisbandAllianceRsp",
		688893070:  "IdSearchPlayerInfoReq",
		990890275:  "IdSearchPlayerInfoRsp",
		710800074:  "IdInviteAllianceReq",
		543023757:  "IdInviteAllianceRsp",
		1891047110: "IdSetMemberPositionReq",
		1656160475: "IdSetMemberPositionRsp",
		1475199389: "IdKickOutMemberReq",
		1173202408: "IdKickOutMemberRsp",
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
		39989867:   "IdMahjongBTEReadyReq",
		207765930:  "IdMahjongBTEReadyRsp",
		983022052:  "IdMahjongBTEExchange3Req",
		681024845:  "IdMahjongBTEExchange3Rsp",
		1993774143: "IdMahjongBTEDecideIgnoreReq",
		1758887476: "IdMahjongBTEDecideIgnoreRsp",
		2138714649: "IdMahjongBTEPlayCardReq",
		1970938586: "IdMahjongBTEPlayCardRsp",
		1792910030: "IdMahjongBTEOperateReq",
		2094907235: "IdMahjongBTEOperateRsp",
		176653714:  "IdMahjongBTEPlayerReadyNtf",
		23903615:   "IdMahjongBTEReadyNtf",
		701290466:  "IdMahjongBTEDecideMasterNtf",
		881890379:  "IdMahjongBTEDealNtf",
		735992092:  "IdMahjongBTEExchange3Ntf",
		595136049:  "IdMahjongBTEExchange3EndNtf",
		1745361179: "IdMahjongBTEDecideIgnoreNtf",
		1299626231: "IdMahjongBTEDecideIgnoreEndNtf",
		116867344:  "IdMahjongBTEPlayingNtf",
		1486690182: "IdMahjongBTETurnNtf",
		1682707103: "IdMahjongBTEOperaNtf",
		1808167222: "IdMahjongBTESettlementNtf",
		1443645576: "IdMailListReq",
		1812753385: "IdMailListRsp",
		1963047789: "IdReadMailReq",
		1661050550: "IdReadMailRsp",
		1679531254: "IdReceiveMailItemReq",
		1981528235: "IdReceiveMailItemRsp",
		1054436458: "IdDeleteMailReq",
		2001014089: "IdCreateRoomReq",
		222637993:  "IdCreateRoomRsp",
		1782747149: "IdDisbandRoomReq",
		2017634040: "IdDisbandRoomRsp",
		21167689:   "IdRoomListReq",
		390275242:  "IdRoomListRsp",
		1530633489: "IdJoinRoomReq",
		1899741298: "IdJoinRoomRsp",
		1875949722: "IdLeaveRoomReq",
		97573884:   "IdLeaveRoomRsp",
		867132366:  "IdRoomPlayerEnterNtf",
		721409295:  "IdRoomPlayerLeaveNtf",
		1386944282: "IdRoomPlayerOnlineNtf",
	}
	Msg_value = map[string]int32{
		"IdUnknown":                      0,
		"IdAgentMembersReq":              1894584925,
		"IdAgentMembersRsp":              1592587688,
		"IdAllianceInfoNtf":              672428706,
		"IdDisbandAllianceReq":           650636160,
		"IdDisbandAllianceRsp":           952633399,
		"IdSearchPlayerInfoReq":          688893070,
		"IdSearchPlayerInfoRsp":          990890275,
		"IdInviteAllianceReq":            710800074,
		"IdInviteAllianceRsp":            543023757,
		"IdSetMemberPositionReq":         1891047110,
		"IdSetMemberPositionRsp":         1656160475,
		"IdKickOutMemberReq":             1475199389,
		"IdKickOutMemberRsp":             1173202408,
		"IdFailRsp":                      160109657,
		"IdSetRoleInfoReq":               58140470,
		"IdSetRoleInfoRsp":               1903626880,
		"IdBindPhoneReq":                 403639403,
		"IdBindPhoneRsp":                 34531592,
		"IdModifyPasswordReq":            655389589,
		"IdModifyPasswordRsp":            890276254,
		"IdHeartReq":                     31040569,
		"IdHeartRsp":                     400148124,
		"IdLoginReq":                     1510628254,
		"IdLoginRsp":                     1275741587,
		"IdEnterGameReq":                 117622385,
		"IdEnterGameRsp":                 2097329717,
		"IdMahjongBTEReadyReq":           39989867,
		"IdMahjongBTEReadyRsp":           207765930,
		"IdMahjongBTEExchange3Req":       983022052,
		"IdMahjongBTEExchange3Rsp":       681024845,
		"IdMahjongBTEDecideIgnoreReq":    1993774143,
		"IdMahjongBTEDecideIgnoreRsp":    1758887476,
		"IdMahjongBTEPlayCardReq":        2138714649,
		"IdMahjongBTEPlayCardRsp":        1970938586,
		"IdMahjongBTEOperateReq":         1792910030,
		"IdMahjongBTEOperateRsp":         2094907235,
		"IdMahjongBTEPlayerReadyNtf":     176653714,
		"IdMahjongBTEReadyNtf":           23903615,
		"IdMahjongBTEDecideMasterNtf":    701290466,
		"IdMahjongBTEDealNtf":            881890379,
		"IdMahjongBTEExchange3Ntf":       735992092,
		"IdMahjongBTEExchange3EndNtf":    595136049,
		"IdMahjongBTEDecideIgnoreNtf":    1745361179,
		"IdMahjongBTEDecideIgnoreEndNtf": 1299626231,
		"IdMahjongBTEPlayingNtf":         116867344,
		"IdMahjongBTETurnNtf":            1486690182,
		"IdMahjongBTEOperaNtf":           1682707103,
		"IdMahjongBTESettlementNtf":      1808167222,
		"IdMailListReq":                  1443645576,
		"IdMailListRsp":                  1812753385,
		"IdReadMailReq":                  1963047789,
		"IdReadMailRsp":                  1661050550,
		"IdReceiveMailItemReq":           1679531254,
		"IdReceiveMailItemRsp":           1981528235,
		"IdDeleteMailReq":                1054436458,
		"IdCreateRoomReq":                2001014089,
		"IdCreateRoomRsp":                222637993,
		"IdDisbandRoomReq":               1782747149,
		"IdDisbandRoomRsp":               2017634040,
		"IdRoomListReq":                  21167689,
		"IdRoomListRsp":                  390275242,
		"IdJoinRoomReq":                  1530633489,
		"IdJoinRoomRsp":                  1899741298,
		"IdLeaveRoomReq":                 1875949722,
		"IdLeaveRoomRsp":                 97573884,
		"IdRoomPlayerEnterNtf":           867132366,
		"IdRoomPlayerLeaveNtf":           721409295,
		"IdRoomPlayerOnlineNtf":          1386944282,
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
	0x75, 0x74, 0x65, 0x72, 0x2a, 0x95, 0x0f, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x0d, 0x0a, 0x09,
	0x49, 0x64, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x11, 0x49,
	0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x10, 0xdd, 0xa4, 0xb4, 0x87, 0x07, 0x12, 0x19, 0x0a, 0x11, 0x49, 0x64, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x73, 0x70, 0x10, 0xa8, 0xeb, 0xb3, 0xf7,
	0x05, 0x12, 0x19, 0x0a, 0x11, 0x49, 0x64, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x4e, 0x74, 0x66, 0x10, 0xa2, 0xe5, 0xd1, 0xc0, 0x02, 0x12, 0x1c, 0x0a, 0x14,
	0x49, 0x64, 0x44, 0x69, 0x73, 0x62, 0x61, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x10, 0x80, 0xd7, 0x9f, 0xb6, 0x02, 0x12, 0x1c, 0x0a, 0x14, 0x49, 0x64,
	0x44, 0x69, 0x73, 0x62, 0x61, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x73, 0x70, 0x10, 0xb7, 0x90, 0xa0, 0xc6, 0x03, 0x12, 0x1d, 0x0a, 0x15, 0x49, 0x64, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x10, 0x8e, 0xd9, 0xbe, 0xc8, 0x02, 0x12, 0x1d, 0x0a, 0x15, 0x49, 0x64, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70,
	0x10, 0xa3, 0x92, 0xbf, 0xd8, 0x03, 0x12, 0x1b, 0x0a, 0x13, 0x49, 0x64, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x10, 0xca, 0xe5,
	0xf7, 0xd2, 0x02, 0x12, 0x1b, 0x0a, 0x13, 0x49, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x41,
	0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x73, 0x70, 0x10, 0x8d, 0xc5, 0xf7, 0x82, 0x02,
	0x12, 0x1e, 0x0a, 0x16, 0x49, 0x64, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x10, 0xc6, 0xad, 0xdc, 0x85, 0x07,
	0x12, 0x1e, 0x0a, 0x16, 0x49, 0x64, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x10, 0xdb, 0x81, 0xdc, 0x95, 0x06,
	0x12, 0x1a, 0x0a, 0x12, 0x49, 0x64, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x10, 0x9d, 0x83, 0xb7, 0xbf, 0x05, 0x12, 0x1a, 0x0a, 0x12,
	0x49, 0x64, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x73, 0x70, 0x10, 0xe8, 0xcb, 0xb6, 0xaf, 0x04, 0x12, 0x10, 0x0a, 0x09, 0x49, 0x64, 0x46, 0x61,
	0x69, 0x6c, 0x52, 0x73, 0x70, 0x10, 0xd9, 0xa8, 0xac, 0x4c, 0x12, 0x17, 0x0a, 0x10, 0x49, 0x64,
	0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x10, 0xb6,
	0xce, 0xdc, 0x1b, 0x12, 0x18, 0x0a, 0x10, 0x49, 0x64, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x10, 0x80, 0x95, 0xdc, 0x8b, 0x07, 0x12, 0x16, 0x0a,
	0x0e, 0x49, 0x64, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x10,
	0xeb, 0x98, 0xbc, 0xc0, 0x01, 0x12, 0x15, 0x0a, 0x0e, 0x49, 0x64, 0x42, 0x69, 0x6e, 0x64, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x52, 0x73, 0x70, 0x10, 0x88, 0xd2, 0xbb, 0x10, 0x12, 0x1b, 0x0a, 0x13,
	0x49, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x10, 0x95, 0xe7, 0xc1, 0xb8, 0x02, 0x12, 0x1b, 0x0a, 0x13, 0x49, 0x64, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x73, 0x70,
	0x10, 0x9e, 0x93, 0xc2, 0xa8, 0x03, 0x12, 0x11, 0x0a, 0x0a, 0x49, 0x64, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x10, 0xb9, 0xc8, 0xe6, 0x0e, 0x12, 0x12, 0x0a, 0x0a, 0x49, 0x64, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x52, 0x73, 0x70, 0x10, 0x9c, 0x8d, 0xe7, 0xbe, 0x01, 0x12, 0x12, 0x0a,
	0x0a, 0x49, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x10, 0x9e, 0xb7, 0xa9, 0xd0,
	0x05, 0x12, 0x12, 0x0a, 0x0a, 0x49, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x10,
	0x93, 0x8b, 0xa9, 0xe0, 0x04, 0x12, 0x15, 0x0a, 0x0e, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x10, 0xf1, 0x8c, 0x8b, 0x38, 0x12, 0x16, 0x0a, 0x0e,
	0x49, 0x64, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x73, 0x70, 0x10, 0xb5,
	0xec, 0x8a, 0xe8, 0x07, 0x12, 0x1b, 0x0a, 0x14, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e,
	0x67, 0x42, 0x54, 0x45, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x10, 0xeb, 0xe4, 0x88,
	0x13, 0x12, 0x1b, 0x0a, 0x14, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54,
	0x45, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x73, 0x70, 0x10, 0xaa, 0x83, 0x89, 0x63, 0x12, 0x20,
	0x0a, 0x18, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x33, 0x52, 0x65, 0x71, 0x10, 0xe4, 0xf3, 0xde, 0xd4, 0x03,
	0x12, 0x20, 0x0a, 0x18, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x33, 0x52, 0x73, 0x70, 0x10, 0xcd, 0xba, 0xde,
	0xc4, 0x02, 0x12, 0x23, 0x0a, 0x1b, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42,
	0x54, 0x45, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x10, 0xbf, 0xa8, 0xda, 0xb6, 0x07, 0x12, 0x23, 0x0a, 0x1b, 0x49, 0x64, 0x4d, 0x61, 0x68,
	0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x49, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x52, 0x73, 0x70, 0x10, 0xb4, 0xfc, 0xd9, 0xc6, 0x06, 0x12, 0x1f, 0x0a, 0x17,
	0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x50, 0x6c, 0x61, 0x79,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x10, 0x99, 0xe4, 0xe8, 0xfb, 0x07, 0x12, 0x1f, 0x0a,
	0x17, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x50, 0x6c, 0x61,
	0x79, 0x43, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x10, 0xda, 0xc5, 0xe8, 0xab, 0x07, 0x12, 0x1e,
	0x0a, 0x16, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x10, 0xce, 0xc5, 0xf6, 0xd6, 0x06, 0x12, 0x1e,
	0x0a, 0x16, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x10, 0xe3, 0xfe, 0xf6, 0xe6, 0x07, 0x12, 0x21,
	0x0a, 0x1a, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x79, 0x4e, 0x74, 0x66, 0x10, 0x92, 0x8b, 0x9e,
	0x54, 0x12, 0x1b, 0x0a, 0x14, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54,
	0x45, 0x52, 0x65, 0x61, 0x64, 0x79, 0x4e, 0x74, 0x66, 0x10, 0xff, 0xfa, 0xb2, 0x0b, 0x12, 0x23,
	0x0a, 0x1b, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x44, 0x65,
	0x63, 0x69, 0x64, 0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x74, 0x66, 0x10, 0xe2, 0xaf,
	0xb3, 0xce, 0x02, 0x12, 0x1b, 0x0a, 0x13, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67,
	0x42, 0x54, 0x45, 0x44, 0x65, 0x61, 0x6c, 0x4e, 0x74, 0x66, 0x10, 0xcb, 0xa8, 0xc2, 0xa4, 0x03,
	0x12, 0x20, 0x0a, 0x18, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x33, 0x4e, 0x74, 0x66, 0x10, 0x9c, 0xb2, 0xf9,
	0xde, 0x02, 0x12, 0x23, 0x0a, 0x1b, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42,
	0x54, 0x45, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x33, 0x45, 0x6e, 0x64, 0x4e, 0x74,
	0x66, 0x10, 0xb1, 0x9c, 0xe4, 0x9b, 0x02, 0x12, 0x23, 0x0a, 0x1b, 0x49, 0x64, 0x4d, 0x61, 0x68,
	0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x44, 0x65, 0x63, 0x69, 0x64, 0x65, 0x49, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x4e, 0x74, 0x66, 0x10, 0x9b, 0xb2, 0xa0, 0xc0, 0x06, 0x12, 0x26, 0x0a, 0x1e,
	0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x44, 0x65, 0x63, 0x69,
	0x64, 0x65, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x64, 0x4e, 0x74, 0x66, 0x10, 0xf7,
	0xf1, 0xda, 0xeb, 0x04, 0x12, 0x1d, 0x0a, 0x16, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e,
	0x67, 0x42, 0x54, 0x45, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x4e, 0x74, 0x66, 0x10, 0x90,
	0x82, 0xdd, 0x37, 0x12, 0x1b, 0x0a, 0x13, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67,
	0x42, 0x54, 0x45, 0x54, 0x75, 0x72, 0x6e, 0x4e, 0x74, 0x66, 0x10, 0x86, 0xaf, 0xf4, 0xc4, 0x05,
	0x12, 0x1c, 0x0a, 0x14, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x4e, 0x74, 0x66, 0x10, 0x9f, 0xa5, 0xb0, 0xa2, 0x06, 0x12, 0x21,
	0x0a, 0x19, 0x49, 0x64, 0x4d, 0x61, 0x68, 0x6a, 0x6f, 0x6e, 0x67, 0x42, 0x54, 0x45, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x74, 0x66, 0x10, 0xb6, 0xe2, 0x99, 0xde,
	0x06, 0x12, 0x15, 0x0a, 0x0d, 0x49, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x10, 0x88, 0x91, 0xb1, 0xb0, 0x05, 0x12, 0x15, 0x0a, 0x0d, 0x49, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x10, 0xe9, 0xd7, 0xb1, 0xe0, 0x06, 0x12,
	0x15, 0x0a, 0x0d, 0x49, 0x64, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x10, 0xed, 0xf6, 0x86, 0xa8, 0x07, 0x12, 0x15, 0x0a, 0x0d, 0x49, 0x64, 0x52, 0x65, 0x61, 0x64,
	0x4d, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x10, 0xb6, 0xbd, 0x86, 0x98, 0x06, 0x12, 0x1c, 0x0a,
	0x14, 0x49, 0x64, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x10, 0xf6, 0xb9, 0xee, 0xa0, 0x06, 0x12, 0x1c, 0x0a, 0x14, 0x49,
	0x64, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x73, 0x70, 0x10, 0xab, 0xf1, 0xee, 0xb0, 0x07, 0x12, 0x17, 0x0a, 0x0f, 0x49, 0x64, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x10, 0xea, 0xd8, 0xe5,
	0xf6, 0x03, 0x12, 0x17, 0x0a, 0x0f, 0x49, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x10, 0xc9, 0x9a, 0x94, 0xba, 0x07, 0x12, 0x16, 0x0a, 0x0f, 0x49,
	0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x73, 0x70, 0x10, 0xa9,
	0xdf, 0x94, 0x6a, 0x12, 0x18, 0x0a, 0x10, 0x49, 0x64, 0x44, 0x69, 0x73, 0x62, 0x61, 0x6e, 0x64,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x10, 0x8d, 0xa0, 0x8a, 0xd2, 0x06, 0x12, 0x18, 0x0a,
	0x10, 0x49, 0x64, 0x44, 0x69, 0x73, 0x62, 0x61, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x73,
	0x70, 0x10, 0xf8, 0xcd, 0x8a, 0xc2, 0x07, 0x12, 0x14, 0x0a, 0x0d, 0x49, 0x64, 0x52, 0x6f, 0x6f,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x10, 0xc9, 0xfc, 0x8b, 0x0a, 0x12, 0x15, 0x0a,
	0x0d, 0x49, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x10, 0xaa,
	0xc1, 0x8c, 0xba, 0x01, 0x12, 0x15, 0x0a, 0x0d, 0x49, 0x64, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x71, 0x10, 0x91, 0xba, 0xee, 0xd9, 0x05, 0x12, 0x15, 0x0a, 0x0d, 0x49,
	0x64, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x73, 0x70, 0x10, 0xf2, 0x80, 0xef,
	0x89, 0x07, 0x12, 0x16, 0x0a, 0x0e, 0x49, 0x64, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x71, 0x10, 0x9a, 0xf1, 0xc2, 0xfe, 0x06, 0x12, 0x15, 0x0a, 0x0e, 0x49, 0x64,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x73, 0x70, 0x10, 0xfc, 0xb7, 0xc3,
	0x2e, 0x12, 0x1c, 0x0a, 0x14, 0x49, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x74, 0x66, 0x10, 0xce, 0xc7, 0xbd, 0x9d, 0x03, 0x12,
	0x1c, 0x0a, 0x14, 0x49, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x4e, 0x74, 0x66, 0x10, 0x8f, 0xaa, 0xff, 0xd7, 0x02, 0x12, 0x1d, 0x0a,
	0x15, 0x49, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x4e, 0x74, 0x66, 0x10, 0x9a, 0xae, 0xac, 0x95, 0x05, 0x42, 0x08, 0x5a, 0x06,
	0x2f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
