// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enum.proto

package outer

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

//前后端通信协议 id段 100000-1000000
type MSG int32

const (
	MSG_UNKNOWN MSG = 0
	MSG_PING    MSG = 100001
	MSG_PONG    MSG = 100002
	MSG_FAIL    MSG = 100003
	MSG_OK      MSG = 100004
	//--------------------------------------------------------------- login proto 200001-300000
	// 消息段begin
	MSG_LOGIN_SEGMENT_BEGIN MSG = 200001
	// 请求登录
	MSG_LOGIN_REQ  MSG = 200002
	MSG_LOGIN_RESP MSG = 200003
	// 消息段end
	MSG_LOGIN_SEGMENT_END MSG = 300000
	//--------------------------------------------------------------- game proto 300001-310000
	// 消息段begin
	MSG_GAME_SEGMENT_BEGIN MSG = 300001
	// 请求登录
	MSG_ENTER_GAME_REQ  MSG = 300100
	MSG_ROLE_INFO_PUSH  MSG = 300101
	MSG_ITEM_INFO_PUSH  MSG = 300102
	MSG_ENTER_GAME_RESP MSG = 300201
	// 道具
	MSG_USE_ITEM_REQ       MSG = 300301
	MSG_USE_ITEM_RESP      MSG = 300302
	MSG_ITEM_CHANGE_NOTIFY MSG = 300303
	// 邮件
	MSG_MAIL_LIST_REQ          MSG = 300401
	MSG_MAIL_LIST_RESP         MSG = 300402
	MSG_ADD_MAIL_NOTIFY        MSG = 300403
	MSG_READ_MAIL_REQ          MSG = 300404
	MSG_READ_MAIL_RESP         MSG = 300405
	MSG_RECEIVE_MAIL_ITEM_REQ  MSG = 300406
	MSG_RECEIVE_MAIL_ITEM_RESP MSG = 300407
	MSG_DELETE_MAIL_REQ        MSG = 300408
	// 消息段end
	MSG_GAME_SEGMENT_END MSG = 310000
)

var MSG_name = map[int32]string{
	0:      "UNKNOWN",
	100001: "PING",
	100002: "PONG",
	100003: "FAIL",
	100004: "OK",
	200001: "LOGIN_SEGMENT_BEGIN",
	200002: "LOGIN_REQ",
	200003: "LOGIN_RESP",
	300000: "LOGIN_SEGMENT_END",
	300001: "GAME_SEGMENT_BEGIN",
	300100: "ENTER_GAME_REQ",
	300101: "ROLE_INFO_PUSH",
	300102: "ITEM_INFO_PUSH",
	300201: "ENTER_GAME_RESP",
	300301: "USE_ITEM_REQ",
	300302: "USE_ITEM_RESP",
	300303: "ITEM_CHANGE_NOTIFY",
	300401: "MAIL_LIST_REQ",
	300402: "MAIL_LIST_RESP",
	300403: "ADD_MAIL_NOTIFY",
	300404: "READ_MAIL_REQ",
	300405: "READ_MAIL_RESP",
	300406: "RECEIVE_MAIL_ITEM_REQ",
	300407: "RECEIVE_MAIL_ITEM_RESP",
	300408: "DELETE_MAIL_REQ",
	310000: "GAME_SEGMENT_END",
}

var MSG_value = map[string]int32{
	"UNKNOWN":                0,
	"PING":                   100001,
	"PONG":                   100002,
	"FAIL":                   100003,
	"OK":                     100004,
	"LOGIN_SEGMENT_BEGIN":    200001,
	"LOGIN_REQ":              200002,
	"LOGIN_RESP":             200003,
	"LOGIN_SEGMENT_END":      300000,
	"GAME_SEGMENT_BEGIN":     300001,
	"ENTER_GAME_REQ":         300100,
	"ROLE_INFO_PUSH":         300101,
	"ITEM_INFO_PUSH":         300102,
	"ENTER_GAME_RESP":        300201,
	"USE_ITEM_REQ":           300301,
	"USE_ITEM_RESP":          300302,
	"ITEM_CHANGE_NOTIFY":     300303,
	"MAIL_LIST_REQ":          300401,
	"MAIL_LIST_RESP":         300402,
	"ADD_MAIL_NOTIFY":        300403,
	"READ_MAIL_REQ":          300404,
	"READ_MAIL_RESP":         300405,
	"RECEIVE_MAIL_ITEM_REQ":  300406,
	"RECEIVE_MAIL_ITEM_RESP": 300407,
	"DELETE_MAIL_REQ":        300408,
	"GAME_SEGMENT_END":       310000,
}

func (x MSG) String() string {
	return proto.EnumName(MSG_name, int32(x))
}

func (MSG) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{0}
}

type Unknown struct {
}

func (m *Unknown) Reset()         { *m = Unknown{} }
func (m *Unknown) String() string { return proto.CompactTextString(m) }
func (*Unknown) ProtoMessage()    {}
func (*Unknown) Descriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{0}
}
func (m *Unknown) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Unknown) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Unknown.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Unknown) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unknown.Merge(m, src)
}
func (m *Unknown) XXX_Size() int {
	return m.Size()
}
func (m *Unknown) XXX_DiscardUnknown() {
	xxx_messageInfo_Unknown.DiscardUnknown(m)
}

var xxx_messageInfo_Unknown proto.InternalMessageInfo

type Ok struct {
}

func (m *Ok) Reset()         { *m = Ok{} }
func (m *Ok) String() string { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()    {}
func (*Ok) Descriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{1}
}
func (m *Ok) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ok) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ok.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ok) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ok.Merge(m, src)
}
func (m *Ok) XXX_Size() int {
	return m.Size()
}
func (m *Ok) XXX_DiscardUnknown() {
	xxx_messageInfo_Ok.DiscardUnknown(m)
}

var xxx_messageInfo_Ok proto.InternalMessageInfo

type Fail struct {
	Error ERROR  `protobuf:"varint,1,opt,name=Error,proto3,enum=outer.ERROR" json:"Error,omitempty"`
	Info  string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (m *Fail) Reset()         { *m = Fail{} }
func (m *Fail) String() string { return proto.CompactTextString(m) }
func (*Fail) ProtoMessage()    {}
func (*Fail) Descriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{2}
}
func (m *Fail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fail.Merge(m, src)
}
func (m *Fail) XXX_Size() int {
	return m.Size()
}
func (m *Fail) XXX_DiscardUnknown() {
	xxx_messageInfo_Fail.DiscardUnknown(m)
}

var xxx_messageInfo_Fail proto.InternalMessageInfo

func (m *Fail) GetError() ERROR {
	if m != nil {
		return m.Error
	}
	return ERROR_SUCCESS
}

func (m *Fail) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterEnum("outer.MSG", MSG_name, MSG_value)
	proto.RegisterType((*Unknown)(nil), "outer.Unknown")
	proto.RegisterType((*Ok)(nil), "outer.Ok")
	proto.RegisterType((*Fail)(nil), "outer.Fail")
}

func init() { proto.RegisterFile("enum.proto", fileDescriptor_13a9f1b5947140c8) }

var fileDescriptor_13a9f1b5947140c8 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xe3, 0x92, 0xa6, 0xe4, 0x34, 0xb4, 0xc3, 0x29, 0x2d, 0xe1, 0x22, 0x2b, 0xca, 0xaa,
	0x62, 0x11, 0x24, 0xd8, 0x23, 0xb9, 0xcd, 0x89, 0x3b, 0xaa, 0x3d, 0x36, 0x33, 0x0e, 0x08, 0x36,
	0x16, 0x48, 0x41, 0xaa, 0x0a, 0x36, 0xb2, 0x5a, 0xf1, 0x14, 0x01, 0xd6, 0xd9, 0x01, 0x95, 0xda,
	0xfa, 0x2d, 0xb8, 0x8a, 0x65, 0x97, 0x2c, 0x4b, 0xf2, 0x02, 0xdc, 0x0a, 0x2c, 0x91, 0x67, 0x0a,
	0x24, 0x88, 0xdd, 0xd1, 0xf7, 0xcf, 0xf9, 0xe6, 0x97, 0x3d, 0x00, 0xbd, 0x64, 0xfb, 0x41, 0xeb,
	0x61, 0x96, 0x6e, 0xa5, 0x38, 0x9d, 0x6e, 0x6f, 0xf5, 0xb2, 0xf3, 0xb3, 0xbd, 0x2c, 0x4b, 0x33,
	0xc3, 0x9a, 0x55, 0x98, 0xe9, 0x26, 0x9b, 0x49, 0xfa, 0x28, 0x69, 0x96, 0x61, 0x2a, 0xd8, 0x6c,
	0x5e, 0x83, 0x72, 0xe7, 0xce, 0xc6, 0x7d, 0x6c, 0xc2, 0x34, 0x15, 0xe7, 0xea, 0x56, 0xc3, 0x5a,
	0x9e, 0xbb, 0x52, 0x6b, 0xe9, 0xe5, 0x16, 0x49, 0x19, 0x48, 0x69, 0x22, 0x44, 0x28, 0x6f, 0x24,
	0xf7, 0xd2, 0xfa, 0x54, 0xc3, 0x5a, 0xae, 0x4a, 0x3d, 0x5f, 0xda, 0x29, 0xc3, 0x09, 0x5f, 0xb9,
	0x38, 0x0b, 0x33, 0x5d, 0xb1, 0x2e, 0x82, 0x9b, 0x82, 0x95, 0x10, 0xa0, 0x1c, 0x72, 0xe1, 0xb2,
	0x67, 0xfd, 0x8a, 0x9e, 0x03, 0xe1, 0xb2, 0xe7, 0x66, 0xee, 0x38, 0xdc, 0x63, 0x2f, 0xfa, 0x15,
	0x3c, 0x09, 0x53, 0xc1, 0x3a, 0xdb, 0xe9, 0x57, 0xf0, 0x1c, 0x2c, 0x78, 0x81, 0xcb, 0x45, 0xac,
	0xc8, 0xf5, 0x49, 0x44, 0xf1, 0x0a, 0xb9, 0x5c, 0xb0, 0x97, 0x83, 0x1a, 0xce, 0x43, 0xd5, 0x44,
	0x92, 0xae, 0xb3, 0x57, 0x83, 0x1a, 0x32, 0x80, 0xdf, 0x40, 0x85, 0xec, 0xf5, 0xa0, 0x86, 0x67,
	0xe1, 0xf4, 0xe4, 0x36, 0x89, 0x36, 0x3b, 0xdc, 0x45, 0xac, 0x03, 0xba, 0x8e, 0x4f, 0xff, 0x58,
	0x3f, 0xee, 0x22, 0x9e, 0x81, 0x39, 0x12, 0x11, 0xc9, 0x58, 0xe7, 0x85, 0xfa, 0xcd, 0x9e, 0xa6,
	0x32, 0xf0, 0x28, 0xe6, 0xa2, 0x13, 0xc4, 0x61, 0x57, 0xad, 0xb1, 0xb7, 0x86, 0xf2, 0x88, 0xfc,
	0x31, 0xfa, 0x6e, 0x0f, 0x71, 0x11, 0xe6, 0x27, 0x0c, 0x2a, 0x64, 0xfb, 0xfb, 0x88, 0x08, 0xb5,
	0xae, 0xa2, 0x58, 0x2f, 0x14, 0xda, 0x7e, 0x8e, 0xb8, 0x00, 0xa7, 0xc6, 0x98, 0x0a, 0xd9, 0xe3,
	0x5c, 0x77, 0xd3, 0x60, 0x75, 0xcd, 0x11, 0x2e, 0xc5, 0x22, 0x88, 0x78, 0xe7, 0x16, 0x7b, 0x62,
	0x8e, 0xfb, 0x0e, 0xf7, 0x62, 0x8f, 0xab, 0x48, 0x3b, 0x3e, 0xe7, 0xba, 0xc4, 0x38, 0x54, 0x21,
	0xfb, 0x92, 0xeb, 0x12, 0x4e, 0xbb, 0x1d, 0xeb, 0xe4, 0xd8, 0xf0, 0xd5, 0x18, 0x24, 0x39, 0xc7,
	0xbc, 0x30, 0x7c, 0x33, 0x86, 0x71, 0xa8, 0x42, 0x76, 0x94, 0x23, 0x5e, 0x80, 0x45, 0x49, 0xab,
	0xc4, 0x6f, 0x90, 0x09, 0xfe, 0x14, 0xff, 0x9e, 0x23, 0x5e, 0x84, 0xa5, 0xff, 0x85, 0x2a, 0x64,
	0x3f, 0xcc, 0xe5, 0x6d, 0xf2, 0x28, 0xa2, 0xbf, 0xf7, 0xfc, 0xcc, 0x11, 0x97, 0x80, 0x4d, 0x7c,
	0xf4, 0xe2, 0x67, 0x7c, 0x3a, 0xc2, 0x95, 0xc6, 0xfb, 0xa1, 0x6d, 0x1d, 0x0c, 0x6d, 0xeb, 0x70,
	0x68, 0x5b, 0x4f, 0x47, 0x76, 0xe9, 0x60, 0x64, 0x97, 0x3e, 0x8c, 0xec, 0xd2, 0xed, 0xca, 0x65,
	0xfd, 0xd2, 0xee, 0x56, 0xf4, 0x03, 0xbd, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xad, 0x51,
	0x42, 0xc2, 0x02, 0x00, 0x00,
}

func (m *Unknown) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unknown) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Unknown) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Ok) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ok) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ok) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Fail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintEnum(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x12
	}
	if m.Error != 0 {
		i = encodeVarintEnum(dAtA, i, uint64(m.Error))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEnum(dAtA []byte, offset int, v uint64) int {
	offset -= sovEnum(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Unknown) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Ok) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Fail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Error != 0 {
		n += 1 + sovEnum(uint64(m.Error))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovEnum(uint64(l))
	}
	return n
}

func sovEnum(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEnum(x uint64) (n int) {
	return sovEnum(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Unknown) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnum
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Unknown: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unknown: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEnum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEnum
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Ok) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnum
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ok: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ok: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEnum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEnum
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Fail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnum
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Fail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			m.Error = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Error |= ERROR(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEnum
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEnum(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnum
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnum
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEnum
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEnum
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEnum
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEnum
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEnum        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnum          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEnum = fmt.Errorf("proto: unexpected end of group")
)
