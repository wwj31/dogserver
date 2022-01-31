// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: error.proto

package outer

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ERROR int32

const (
	ERROR_SUCCESS                   ERROR = 0
	ERROR_FAILED                    ERROR = 1
	ERROR_SECURITYCODE_CHECK_FAILED ERROR = 2
	ERROR_ITEM_NOT_ENOUGH           ERROR = 3
	ERROR_GOLD_NOT_ENOUGH           ERROR = 4
	ERROR_LEVEL_NOT_ENGOUTH         ERROR = 5
	ERROR_MAIL_REPEAT_RECV_ITEM     ERROR = 6
	ERROR_CLIENT_WRONG_PARAM        ERROR = 9
	ERROR_CFG_NO_THIS_PARAM         ERROR = 10
	ERROR_NAME_LEN_OUTRANGE         ERROR = 13
)

var ERROR_name = map[int32]string{
	0:  "SUCCESS",
	1:  "FAILED",
	2:  "SECURITYCODE_CHECK_FAILED",
	3:  "ITEM_NOT_ENOUGH",
	4:  "GOLD_NOT_ENOUGH",
	5:  "LEVEL_NOT_ENGOUTH",
	6:  "MAIL_REPEAT_RECV_ITEM",
	9:  "CLIENT_WRONG_PARAM",
	10: "CFG_NO_THIS_PARAM",
	13: "NAME_LEN_OUTRANGE",
}

var ERROR_value = map[string]int32{
	"SUCCESS":                   0,
	"FAILED":                    1,
	"SECURITYCODE_CHECK_FAILED": 2,
	"ITEM_NOT_ENOUGH":           3,
	"GOLD_NOT_ENOUGH":           4,
	"LEVEL_NOT_ENGOUTH":         5,
	"MAIL_REPEAT_RECV_ITEM":     6,
	"CLIENT_WRONG_PARAM":        9,
	"CFG_NO_THIS_PARAM":         10,
	"NAME_LEN_OUTRANGE":         13,
}

func (x ERROR) String() string {
	return proto.EnumName(ERROR_name, int32(x))
}

func (ERROR) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0579b252106fcf4a, []int{0}
}

func init() {
	proto.RegisterEnum("outer.ERROR", ERROR_name, ERROR_value)
}

func init() { proto.RegisterFile("error.proto", fileDescriptor_0579b252106fcf4a) }

var fileDescriptor_0579b252106fcf4a = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0xf1, 0x04, 0x48, 0x10, 0xaf, 0x42, 0x18, 0xa3, 0x22, 0x75, 0xc0, 0x62, 0x66, 0x80,
	0x81, 0x13, 0x18, 0xe7, 0xd5, 0xb1, 0x70, 0xec, 0xca, 0x71, 0x8a, 0x60, 0x79, 0x12, 0x52, 0xe7,
	0xa0, 0xa8, 0xdc, 0x83, 0x63, 0x31, 0x76, 0x64, 0x84, 0xe4, 0x22, 0x28, 0x55, 0x86, 0xae, 0xbf,
	0x4f, 0xfa, 0x86, 0x3f, 0xcc, 0x36, 0x5d, 0xd7, 0x76, 0xf7, 0x1f, 0x5d, 0xbb, 0x6d, 0x79, 0xd6,
	0x7e, 0x6e, 0x37, 0xdd, 0xdd, 0x5f, 0x0a, 0x19, 0x86, 0xe0, 0x03, 0x9f, 0xc1, 0x69, 0xdd, 0x28,
	0x85, 0x75, 0xcd, 0x12, 0x0e, 0x90, 0x2f, 0xa5, 0xb1, 0x58, 0xb0, 0x94, 0xdf, 0xc0, 0xa2, 0x46,
	0xd5, 0x04, 0x13, 0x5f, 0x95, 0x2f, 0x90, 0x54, 0x89, 0xea, 0x99, 0xa6, 0xf9, 0x88, 0x5f, 0xc1,
	0x85, 0x89, 0x58, 0x91, 0xf3, 0x91, 0xd0, 0xf9, 0x46, 0x97, 0xec, 0x78, 0x44, 0xed, 0x6d, 0x71,
	0x88, 0x27, 0x7c, 0x0e, 0x97, 0x16, 0xd7, 0x68, 0x27, 0xd5, 0xbe, 0x89, 0x25, 0xcb, 0xf8, 0x02,
	0xe6, 0x95, 0x34, 0x96, 0x02, 0xae, 0x50, 0x46, 0x0a, 0xa8, 0xd6, 0x34, 0x3e, 0xb2, 0x9c, 0x5f,
	0x03, 0x57, 0xd6, 0xa0, 0x8b, 0xf4, 0x12, 0xbc, 0xd3, 0xb4, 0x92, 0x41, 0x56, 0xec, 0x6c, 0x7c,
	0x52, 0x4b, 0x4d, 0xce, 0x53, 0x2c, 0x4d, 0x3d, 0x31, 0x8c, 0xec, 0x64, 0x85, 0x64, 0xd1, 0x91,
	0x6f, 0x62, 0x90, 0x4e, 0x23, 0x3b, 0x7f, 0xba, 0xfd, 0xee, 0x45, 0xba, 0xeb, 0x45, 0xfa, 0xdb,
	0x8b, 0xf4, 0x6b, 0x10, 0xc9, 0x6e, 0x10, 0xc9, 0xcf, 0x20, 0x92, 0xb7, 0xfc, 0x61, 0x5f, 0xe1,
	0x3d, 0xdf, 0x37, 0x79, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x40, 0xbd, 0x9e, 0x72, 0x22, 0x01,
	0x00, 0x00,
}
