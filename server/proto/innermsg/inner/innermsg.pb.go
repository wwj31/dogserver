// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: innermsg.proto

package inner

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

type Error struct {
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_04a0580320a2f95b, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Error.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return m.Size()
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

type GateMsgWrapper struct {
	GateSession string `protobuf:"bytes,1,opt,name=GateSession,proto3" json:"GateSession,omitempty"`
	MsgName     string `protobuf:"bytes,2,opt,name=MsgName,proto3" json:"MsgName,omitempty"`
	Data        []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *GateMsgWrapper) Reset()         { *m = GateMsgWrapper{} }
func (m *GateMsgWrapper) String() string { return proto.CompactTextString(m) }
func (*GateMsgWrapper) ProtoMessage()    {}
func (*GateMsgWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_04a0580320a2f95b, []int{1}
}
func (m *GateMsgWrapper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GateMsgWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GateMsgWrapper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GateMsgWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GateMsgWrapper.Merge(m, src)
}
func (m *GateMsgWrapper) XXX_Size() int {
	return m.Size()
}
func (m *GateMsgWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_GateMsgWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_GateMsgWrapper proto.InternalMessageInfo

func (m *GateMsgWrapper) GetGateSession() string {
	if m != nil {
		return m.GateSession
	}
	return ""
}

func (m *GateMsgWrapper) GetMsgName() string {
	if m != nil {
		return m.MsgName
	}
	return ""
}

func (m *GateMsgWrapper) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BindSessionWithRID struct {
	GateSession string `protobuf:"bytes,1,opt,name=GateSession,proto3" json:"GateSession,omitempty"`
	RID         string `protobuf:"bytes,2,opt,name=RID,proto3" json:"RID,omitempty"`
}

func (m *BindSessionWithRID) Reset()         { *m = BindSessionWithRID{} }
func (m *BindSessionWithRID) String() string { return proto.CompactTextString(m) }
func (*BindSessionWithRID) ProtoMessage()    {}
func (*BindSessionWithRID) Descriptor() ([]byte, []int) {
	return fileDescriptor_04a0580320a2f95b, []int{2}
}
func (m *BindSessionWithRID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BindSessionWithRID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BindSessionWithRID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BindSessionWithRID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindSessionWithRID.Merge(m, src)
}
func (m *BindSessionWithRID) XXX_Size() int {
	return m.Size()
}
func (m *BindSessionWithRID) XXX_DiscardUnknown() {
	xxx_messageInfo_BindSessionWithRID.DiscardUnknown(m)
}

var xxx_messageInfo_BindSessionWithRID proto.InternalMessageInfo

func (m *BindSessionWithRID) GetGateSession() string {
	if m != nil {
		return m.GateSession
	}
	return ""
}

func (m *BindSessionWithRID) GetRID() string {
	if m != nil {
		return m.RID
	}
	return ""
}

type KickOutReq struct {
	GateSession string `protobuf:"bytes,1,opt,name=GateSession,proto3" json:"GateSession,omitempty"`
	RID         string `protobuf:"bytes,2,opt,name=RID,proto3" json:"RID,omitempty"`
}

func (m *KickOutReq) Reset()         { *m = KickOutReq{} }
func (m *KickOutReq) String() string { return proto.CompactTextString(m) }
func (*KickOutReq) ProtoMessage()    {}
func (*KickOutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_04a0580320a2f95b, []int{3}
}
func (m *KickOutReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KickOutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KickOutReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KickOutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickOutReq.Merge(m, src)
}
func (m *KickOutReq) XXX_Size() int {
	return m.Size()
}
func (m *KickOutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_KickOutReq.DiscardUnknown(m)
}

var xxx_messageInfo_KickOutReq proto.InternalMessageInfo

func (m *KickOutReq) GetGateSession() string {
	if m != nil {
		return m.GateSession
	}
	return ""
}

func (m *KickOutReq) GetRID() string {
	if m != nil {
		return m.RID
	}
	return ""
}

type KickOutRsp struct {
}

func (m *KickOutRsp) Reset()         { *m = KickOutRsp{} }
func (m *KickOutRsp) String() string { return proto.CompactTextString(m) }
func (*KickOutRsp) ProtoMessage()    {}
func (*KickOutRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_04a0580320a2f95b, []int{4}
}
func (m *KickOutRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KickOutRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KickOutRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KickOutRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickOutRsp.Merge(m, src)
}
func (m *KickOutRsp) XXX_Size() int {
	return m.Size()
}
func (m *KickOutRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_KickOutRsp.DiscardUnknown(m)
}

var xxx_messageInfo_KickOutRsp proto.InternalMessageInfo

// session断开,gate通知player
type GSessionClosed struct {
	GateSession string `protobuf:"bytes,1,opt,name=GateSession,proto3" json:"GateSession,omitempty"`
}

func (m *GSessionClosed) Reset()         { *m = GSessionClosed{} }
func (m *GSessionClosed) String() string { return proto.CompactTextString(m) }
func (*GSessionClosed) ProtoMessage()    {}
func (*GSessionClosed) Descriptor() ([]byte, []int) {
	return fileDescriptor_04a0580320a2f95b, []int{5}
}
func (m *GSessionClosed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GSessionClosed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GSessionClosed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GSessionClosed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GSessionClosed.Merge(m, src)
}
func (m *GSessionClosed) XXX_Size() int {
	return m.Size()
}
func (m *GSessionClosed) XXX_DiscardUnknown() {
	xxx_messageInfo_GSessionClosed.DiscardUnknown(m)
}

var xxx_messageInfo_GSessionClosed proto.InternalMessageInfo

func (m *GSessionClosed) GetGateSession() string {
	if m != nil {
		return m.GateSession
	}
	return ""
}

type PullPlayer struct {
	RID     string `protobuf:"bytes,1,opt,name=RID,proto3" json:"RID,omitempty"`
	ShortId int64  `protobuf:"varint,2,opt,name=ShortId,proto3" json:"ShortId,omitempty"`
}

func (m *PullPlayer) Reset()         { *m = PullPlayer{} }
func (m *PullPlayer) String() string { return proto.CompactTextString(m) }
func (*PullPlayer) ProtoMessage()    {}
func (*PullPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_04a0580320a2f95b, []int{6}
}
func (m *PullPlayer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PullPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PullPlayer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PullPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullPlayer.Merge(m, src)
}
func (m *PullPlayer) XXX_Size() int {
	return m.Size()
}
func (m *PullPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_PullPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_PullPlayer proto.InternalMessageInfo

func (m *PullPlayer) GetRID() string {
	if m != nil {
		return m.RID
	}
	return ""
}

func (m *PullPlayer) GetShortId() int64 {
	if m != nil {
		return m.ShortId
	}
	return 0
}

func init() {
	proto.RegisterType((*Error)(nil), "inner.Error")
	proto.RegisterType((*GateMsgWrapper)(nil), "inner.GateMsgWrapper")
	proto.RegisterType((*BindSessionWithRID)(nil), "inner.BindSessionWithRID")
	proto.RegisterType((*KickOutReq)(nil), "inner.KickOutReq")
	proto.RegisterType((*KickOutRsp)(nil), "inner.KickOutRsp")
	proto.RegisterType((*GSessionClosed)(nil), "inner.GSessionClosed")
	proto.RegisterType((*PullPlayer)(nil), "inner.PullPlayer")
}

func init() { proto.RegisterFile("innermsg.proto", fileDescriptor_04a0580320a2f95b) }

var fileDescriptor_04a0580320a2f95b = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0xcb, 0x4b,
	0x2d, 0xca, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xf3, 0x95, 0xd8,
	0xb9, 0x58, 0x5d, 0x8b, 0x8a, 0xf2, 0x8b, 0x94, 0x12, 0xb8, 0xf8, 0xdc, 0x13, 0x4b, 0x52, 0x7d,
	0x8b, 0xd3, 0xc3, 0x8b, 0x12, 0x0b, 0x0a, 0x52, 0x8b, 0x84, 0x14, 0xb8, 0xb8, 0x41, 0x22, 0xc1,
	0xa9, 0xc5, 0xc5, 0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x42,
	0x12, 0x5c, 0xec, 0xbe, 0xc5, 0xe9, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x59, 0x18, 0x57,
	0x48, 0x88, 0x8b, 0xc5, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x59, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc,
	0x56, 0xf2, 0xe0, 0x12, 0x72, 0xca, 0xcc, 0x4b, 0x81, 0x6a, 0x0e, 0xcf, 0x2c, 0xc9, 0x08, 0xf2,
	0x74, 0x21, 0xc2, 0x16, 0x01, 0x2e, 0xe6, 0x20, 0x4f, 0x17, 0xa8, 0x0d, 0x20, 0xa6, 0x92, 0x03,
	0x17, 0x97, 0x77, 0x66, 0x72, 0xb6, 0x7f, 0x69, 0x49, 0x50, 0x6a, 0x21, 0x59, 0x26, 0xf0, 0x20,
	0x4c, 0x28, 0x2e, 0x50, 0x32, 0xe2, 0xe2, 0x73, 0x87, 0xaa, 0x75, 0xce, 0xc9, 0x2f, 0x4e, 0x4d,
	0x21, 0x6c, 0xa6, 0x92, 0x05, 0x17, 0x57, 0x40, 0x69, 0x4e, 0x4e, 0x40, 0x4e, 0x62, 0x65, 0x6a,
	0x11, 0xcc, 0x06, 0x46, 0xb8, 0x0d, 0xa0, 0xb0, 0x09, 0xce, 0xc8, 0x2f, 0x2a, 0xf1, 0x4c, 0x01,
	0xdb, 0xcb, 0x1c, 0x04, 0xe3, 0x3a, 0x29, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x14, 0x9b, 0x3e, 0x38, 0x52, 0x92, 0xd8, 0xc0, 0x51, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x55, 0xf6, 0x2c, 0x42, 0xb4, 0x01, 0x00, 0x00,
}

func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Error) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GateMsgWrapper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GateMsgWrapper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GateMsgWrapper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintInnermsg(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MsgName) > 0 {
		i -= len(m.MsgName)
		copy(dAtA[i:], m.MsgName)
		i = encodeVarintInnermsg(dAtA, i, uint64(len(m.MsgName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GateSession) > 0 {
		i -= len(m.GateSession)
		copy(dAtA[i:], m.GateSession)
		i = encodeVarintInnermsg(dAtA, i, uint64(len(m.GateSession)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BindSessionWithRID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BindSessionWithRID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BindSessionWithRID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RID) > 0 {
		i -= len(m.RID)
		copy(dAtA[i:], m.RID)
		i = encodeVarintInnermsg(dAtA, i, uint64(len(m.RID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GateSession) > 0 {
		i -= len(m.GateSession)
		copy(dAtA[i:], m.GateSession)
		i = encodeVarintInnermsg(dAtA, i, uint64(len(m.GateSession)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KickOutReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KickOutReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KickOutReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RID) > 0 {
		i -= len(m.RID)
		copy(dAtA[i:], m.RID)
		i = encodeVarintInnermsg(dAtA, i, uint64(len(m.RID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GateSession) > 0 {
		i -= len(m.GateSession)
		copy(dAtA[i:], m.GateSession)
		i = encodeVarintInnermsg(dAtA, i, uint64(len(m.GateSession)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KickOutRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KickOutRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KickOutRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GSessionClosed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GSessionClosed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GSessionClosed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GateSession) > 0 {
		i -= len(m.GateSession)
		copy(dAtA[i:], m.GateSession)
		i = encodeVarintInnermsg(dAtA, i, uint64(len(m.GateSession)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PullPlayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PullPlayer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PullPlayer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ShortId != 0 {
		i = encodeVarintInnermsg(dAtA, i, uint64(m.ShortId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RID) > 0 {
		i -= len(m.RID)
		copy(dAtA[i:], m.RID)
		i = encodeVarintInnermsg(dAtA, i, uint64(len(m.RID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInnermsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovInnermsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GateMsgWrapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GateSession)
	if l > 0 {
		n += 1 + l + sovInnermsg(uint64(l))
	}
	l = len(m.MsgName)
	if l > 0 {
		n += 1 + l + sovInnermsg(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovInnermsg(uint64(l))
	}
	return n
}

func (m *BindSessionWithRID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GateSession)
	if l > 0 {
		n += 1 + l + sovInnermsg(uint64(l))
	}
	l = len(m.RID)
	if l > 0 {
		n += 1 + l + sovInnermsg(uint64(l))
	}
	return n
}

func (m *KickOutReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GateSession)
	if l > 0 {
		n += 1 + l + sovInnermsg(uint64(l))
	}
	l = len(m.RID)
	if l > 0 {
		n += 1 + l + sovInnermsg(uint64(l))
	}
	return n
}

func (m *KickOutRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GSessionClosed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GateSession)
	if l > 0 {
		n += 1 + l + sovInnermsg(uint64(l))
	}
	return n
}

func (m *PullPlayer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RID)
	if l > 0 {
		n += 1 + l + sovInnermsg(uint64(l))
	}
	if m.ShortId != 0 {
		n += 1 + sovInnermsg(uint64(m.ShortId))
	}
	return n
}

func sovInnermsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInnermsg(x uint64) (n int) {
	return sovInnermsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInnermsg
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInnermsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInnermsg
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
func (m *GateMsgWrapper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInnermsg
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
			return fmt.Errorf("proto: GateMsgWrapper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GateMsgWrapper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GateSession", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInnermsg
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
				return ErrInvalidLengthInnermsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInnermsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GateSession = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInnermsg
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
				return ErrInvalidLengthInnermsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInnermsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInnermsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInnermsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInnermsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInnermsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInnermsg
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
func (m *BindSessionWithRID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInnermsg
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
			return fmt.Errorf("proto: BindSessionWithRID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BindSessionWithRID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GateSession", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInnermsg
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
				return ErrInvalidLengthInnermsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInnermsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GateSession = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInnermsg
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
				return ErrInvalidLengthInnermsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInnermsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInnermsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInnermsg
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
func (m *KickOutReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInnermsg
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
			return fmt.Errorf("proto: KickOutReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KickOutReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GateSession", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInnermsg
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
				return ErrInvalidLengthInnermsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInnermsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GateSession = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInnermsg
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
				return ErrInvalidLengthInnermsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInnermsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInnermsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInnermsg
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
func (m *KickOutRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInnermsg
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
			return fmt.Errorf("proto: KickOutRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KickOutRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInnermsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInnermsg
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
func (m *GSessionClosed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInnermsg
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
			return fmt.Errorf("proto: GSessionClosed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GSessionClosed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GateSession", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInnermsg
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
				return ErrInvalidLengthInnermsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInnermsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GateSession = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInnermsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInnermsg
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
func (m *PullPlayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInnermsg
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
			return fmt.Errorf("proto: PullPlayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PullPlayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInnermsg
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
				return ErrInvalidLengthInnermsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInnermsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortId", wireType)
			}
			m.ShortId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInnermsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShortId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInnermsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInnermsg
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
func skipInnermsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInnermsg
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
					return 0, ErrIntOverflowInnermsg
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
					return 0, ErrIntOverflowInnermsg
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
				return 0, ErrInvalidLengthInnermsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInnermsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInnermsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInnermsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInnermsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInnermsg = fmt.Errorf("proto: unexpected end of group")
)
