// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alliance.proto

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

// 创建联盟
type CreateAllianceReq struct {
	MasterShortId int64 `protobuf:"varint,1,opt,name=MasterShortId,proto3" json:"MasterShortId,omitempty"`
}

func (m *CreateAllianceReq) Reset()         { *m = CreateAllianceReq{} }
func (m *CreateAllianceReq) String() string { return proto.CompactTextString(m) }
func (*CreateAllianceReq) ProtoMessage()    {}
func (*CreateAllianceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d63a1f01d82564b, []int{0}
}
func (m *CreateAllianceReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateAllianceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateAllianceReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateAllianceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAllianceReq.Merge(m, src)
}
func (m *CreateAllianceReq) XXX_Size() int {
	return m.Size()
}
func (m *CreateAllianceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAllianceReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAllianceReq proto.InternalMessageInfo

func (m *CreateAllianceReq) GetMasterShortId() int64 {
	if m != nil {
		return m.MasterShortId
	}
	return 0
}

type CreateAllianceRsp struct {
	AllianceId int32 `protobuf:"varint,1,opt,name=AllianceId,proto3" json:"AllianceId,omitempty"`
}

func (m *CreateAllianceRsp) Reset()         { *m = CreateAllianceRsp{} }
func (m *CreateAllianceRsp) String() string { return proto.CompactTextString(m) }
func (*CreateAllianceRsp) ProtoMessage()    {}
func (*CreateAllianceRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d63a1f01d82564b, []int{1}
}
func (m *CreateAllianceRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateAllianceRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateAllianceRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateAllianceRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAllianceRsp.Merge(m, src)
}
func (m *CreateAllianceRsp) XXX_Size() int {
	return m.Size()
}
func (m *CreateAllianceRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAllianceRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAllianceRsp proto.InternalMessageInfo

func (m *CreateAllianceRsp) GetAllianceId() int32 {
	if m != nil {
		return m.AllianceId
	}
	return 0
}

// 设置成员信息
type SetMemberReq struct {
	Players []*PlayerInfo `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (m *SetMemberReq) Reset()         { *m = SetMemberReq{} }
func (m *SetMemberReq) String() string { return proto.CompactTextString(m) }
func (*SetMemberReq) ProtoMessage()    {}
func (*SetMemberReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d63a1f01d82564b, []int{2}
}
func (m *SetMemberReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetMemberReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetMemberReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetMemberReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMemberReq.Merge(m, src)
}
func (m *SetMemberReq) XXX_Size() int {
	return m.Size()
}
func (m *SetMemberReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMemberReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetMemberReq proto.InternalMessageInfo

func (m *SetMemberReq) GetPlayers() []*PlayerInfo {
	if m != nil {
		return m.Players
	}
	return nil
}

type SetMemberRsp struct {
}

func (m *SetMemberRsp) Reset()         { *m = SetMemberRsp{} }
func (m *SetMemberRsp) String() string { return proto.CompactTextString(m) }
func (*SetMemberRsp) ProtoMessage()    {}
func (*SetMemberRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d63a1f01d82564b, []int{3}
}
func (m *SetMemberRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetMemberRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetMemberRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetMemberRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMemberRsp.Merge(m, src)
}
func (m *SetMemberRsp) XXX_Size() int {
	return m.Size()
}
func (m *SetMemberRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMemberRsp.DiscardUnknown(m)
}

var xxx_messageInfo_SetMemberRsp proto.InternalMessageInfo

// 登录时，绑定角色session，并请求自己在联盟中的成员最新信息
type MemberInfoOnLoginReq struct {
	GateSession string `protobuf:"bytes,1,opt,name=GateSession,proto3" json:"GateSession,omitempty"`
	RID         string `protobuf:"bytes,2,opt,name=RID,proto3" json:"RID,omitempty"`
}

func (m *MemberInfoOnLoginReq) Reset()         { *m = MemberInfoOnLoginReq{} }
func (m *MemberInfoOnLoginReq) String() string { return proto.CompactTextString(m) }
func (*MemberInfoOnLoginReq) ProtoMessage()    {}
func (*MemberInfoOnLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d63a1f01d82564b, []int{4}
}
func (m *MemberInfoOnLoginReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MemberInfoOnLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberInfoOnLoginReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MemberInfoOnLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberInfoOnLoginReq.Merge(m, src)
}
func (m *MemberInfoOnLoginReq) XXX_Size() int {
	return m.Size()
}
func (m *MemberInfoOnLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberInfoOnLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_MemberInfoOnLoginReq proto.InternalMessageInfo

func (m *MemberInfoOnLoginReq) GetGateSession() string {
	if m != nil {
		return m.GateSession
	}
	return ""
}

func (m *MemberInfoOnLoginReq) GetRID() string {
	if m != nil {
		return m.RID
	}
	return ""
}

type MemberInfoOnLoginRsp struct {
	AllianceId int32 `protobuf:"varint,1,opt,name=allianceId,proto3" json:"allianceId,omitempty"`
	Position   int32 `protobuf:"varint,2,opt,name=Position,proto3" json:"Position,omitempty"`
}

func (m *MemberInfoOnLoginRsp) Reset()         { *m = MemberInfoOnLoginRsp{} }
func (m *MemberInfoOnLoginRsp) String() string { return proto.CompactTextString(m) }
func (*MemberInfoOnLoginRsp) ProtoMessage()    {}
func (*MemberInfoOnLoginRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d63a1f01d82564b, []int{5}
}
func (m *MemberInfoOnLoginRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MemberInfoOnLoginRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberInfoOnLoginRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MemberInfoOnLoginRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberInfoOnLoginRsp.Merge(m, src)
}
func (m *MemberInfoOnLoginRsp) XXX_Size() int {
	return m.Size()
}
func (m *MemberInfoOnLoginRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberInfoOnLoginRsp.DiscardUnknown(m)
}

var xxx_messageInfo_MemberInfoOnLoginRsp proto.InternalMessageInfo

func (m *MemberInfoOnLoginRsp) GetAllianceId() int32 {
	if m != nil {
		return m.AllianceId
	}
	return 0
}

func (m *MemberInfoOnLoginRsp) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

// 下线后，通知联盟
type MemberInfoOnLogoutReq struct {
	GateSession string `protobuf:"bytes,1,opt,name=GateSession,proto3" json:"GateSession,omitempty"`
	RID         string `protobuf:"bytes,2,opt,name=RID,proto3" json:"RID,omitempty"`
}

func (m *MemberInfoOnLogoutReq) Reset()         { *m = MemberInfoOnLogoutReq{} }
func (m *MemberInfoOnLogoutReq) String() string { return proto.CompactTextString(m) }
func (*MemberInfoOnLogoutReq) ProtoMessage()    {}
func (*MemberInfoOnLogoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d63a1f01d82564b, []int{6}
}
func (m *MemberInfoOnLogoutReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MemberInfoOnLogoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MemberInfoOnLogoutReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MemberInfoOnLogoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberInfoOnLogoutReq.Merge(m, src)
}
func (m *MemberInfoOnLogoutReq) XXX_Size() int {
	return m.Size()
}
func (m *MemberInfoOnLogoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberInfoOnLogoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_MemberInfoOnLogoutReq proto.InternalMessageInfo

func (m *MemberInfoOnLogoutReq) GetGateSession() string {
	if m != nil {
		return m.GateSession
	}
	return ""
}

func (m *MemberInfoOnLogoutReq) GetRID() string {
	if m != nil {
		return m.RID
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateAllianceReq)(nil), "inner.CreateAllianceReq")
	proto.RegisterType((*CreateAllianceRsp)(nil), "inner.CreateAllianceRsp")
	proto.RegisterType((*SetMemberReq)(nil), "inner.SetMemberReq")
	proto.RegisterType((*SetMemberRsp)(nil), "inner.SetMemberRsp")
	proto.RegisterType((*MemberInfoOnLoginReq)(nil), "inner.MemberInfoOnLoginReq")
	proto.RegisterType((*MemberInfoOnLoginRsp)(nil), "inner.MemberInfoOnLoginRsp")
	proto.RegisterType((*MemberInfoOnLogoutReq)(nil), "inner.MemberInfoOnLogoutReq")
}

func init() { proto.RegisterFile("alliance.proto", fileDescriptor_9d63a1f01d82564b) }

var fileDescriptor_9d63a1f01d82564b = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcc, 0xc9, 0xc9,
	0x4c, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcc, 0xcb, 0x4b,
	0x2d, 0x92, 0xe2, 0x4a, 0x4a, 0x2c, 0x86, 0x0a, 0x29, 0x59, 0x72, 0x09, 0x3a, 0x17, 0xa5, 0x26,
	0x96, 0xa4, 0x3a, 0x42, 0x95, 0x06, 0xa5, 0x16, 0x0a, 0xa9, 0x70, 0xf1, 0xfa, 0x26, 0x16, 0x97,
	0xa4, 0x16, 0x05, 0x67, 0xe4, 0x17, 0x95, 0x78, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07,
	0xa1, 0x0a, 0x2a, 0x19, 0x63, 0x68, 0x2d, 0x2e, 0x10, 0x92, 0xe3, 0xe2, 0x82, 0x71, 0xa1, 0xfa,
	0x58, 0x83, 0x90, 0x44, 0x94, 0xac, 0xb9, 0x78, 0x82, 0x53, 0x4b, 0x7c, 0x53, 0x73, 0x93, 0x52,
	0x8b, 0x40, 0x56, 0x69, 0x73, 0xb1, 0x17, 0xe4, 0x24, 0x56, 0xa6, 0x16, 0x15, 0x4b, 0x30, 0x2a,
	0x30, 0x6b, 0x70, 0x1b, 0x09, 0xea, 0x81, 0x1d, 0xa9, 0x17, 0x00, 0x16, 0xf5, 0xcc, 0x4b, 0xcb,
	0x0f, 0x82, 0xa9, 0x50, 0xe2, 0x43, 0xd6, 0x5c, 0x5c, 0xa0, 0xe4, 0xc5, 0x25, 0x02, 0xe1, 0x80,
	0x94, 0xf9, 0xe7, 0xf9, 0xe4, 0xa7, 0x67, 0xe6, 0x81, 0x0c, 0x55, 0xe0, 0xe2, 0x76, 0x4f, 0x2c,
	0x49, 0x0d, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x03, 0xbb, 0x82, 0x33, 0x08, 0x59, 0x48, 0x48,
	0x80, 0x8b, 0x39, 0xc8, 0xd3, 0x45, 0x82, 0x09, 0x2c, 0x03, 0x62, 0x2a, 0x05, 0x61, 0x33, 0x0b,
	0xe2, 0xa1, 0x44, 0x0c, 0x0f, 0x21, 0x44, 0x84, 0xa4, 0xb8, 0x38, 0x02, 0xf2, 0x8b, 0x33, 0x4b,
	0x40, 0x16, 0x31, 0x81, 0x65, 0xe1, 0x7c, 0x25, 0x6f, 0x2e, 0x51, 0x34, 0x33, 0xf3, 0x4b, 0x4b,
	0xc8, 0x74, 0xa0, 0x93, 0xc2, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xb1,
	0xe9, 0x83, 0x83, 0x2c, 0x89, 0x0d, 0x1c, 0xa5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5,
	0x4f, 0xfa, 0xfd, 0xf7, 0x01, 0x00, 0x00,
}

func (m *CreateAllianceReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateAllianceReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateAllianceReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MasterShortId != 0 {
		i = encodeVarintAlliance(dAtA, i, uint64(m.MasterShortId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CreateAllianceRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateAllianceRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateAllianceRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllianceId != 0 {
		i = encodeVarintAlliance(dAtA, i, uint64(m.AllianceId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SetMemberReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetMemberReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetMemberReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Players) > 0 {
		for iNdEx := len(m.Players) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Players[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAlliance(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SetMemberRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetMemberRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetMemberRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MemberInfoOnLoginReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberInfoOnLoginReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberInfoOnLoginReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RID) > 0 {
		i -= len(m.RID)
		copy(dAtA[i:], m.RID)
		i = encodeVarintAlliance(dAtA, i, uint64(len(m.RID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GateSession) > 0 {
		i -= len(m.GateSession)
		copy(dAtA[i:], m.GateSession)
		i = encodeVarintAlliance(dAtA, i, uint64(len(m.GateSession)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MemberInfoOnLoginRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberInfoOnLoginRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberInfoOnLoginRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Position != 0 {
		i = encodeVarintAlliance(dAtA, i, uint64(m.Position))
		i--
		dAtA[i] = 0x10
	}
	if m.AllianceId != 0 {
		i = encodeVarintAlliance(dAtA, i, uint64(m.AllianceId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MemberInfoOnLogoutReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MemberInfoOnLogoutReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MemberInfoOnLogoutReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RID) > 0 {
		i -= len(m.RID)
		copy(dAtA[i:], m.RID)
		i = encodeVarintAlliance(dAtA, i, uint64(len(m.RID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GateSession) > 0 {
		i -= len(m.GateSession)
		copy(dAtA[i:], m.GateSession)
		i = encodeVarintAlliance(dAtA, i, uint64(len(m.GateSession)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAlliance(dAtA []byte, offset int, v uint64) int {
	offset -= sovAlliance(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateAllianceReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MasterShortId != 0 {
		n += 1 + sovAlliance(uint64(m.MasterShortId))
	}
	return n
}

func (m *CreateAllianceRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AllianceId != 0 {
		n += 1 + sovAlliance(uint64(m.AllianceId))
	}
	return n
}

func (m *SetMemberReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Players) > 0 {
		for _, e := range m.Players {
			l = e.Size()
			n += 1 + l + sovAlliance(uint64(l))
		}
	}
	return n
}

func (m *SetMemberRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MemberInfoOnLoginReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GateSession)
	if l > 0 {
		n += 1 + l + sovAlliance(uint64(l))
	}
	l = len(m.RID)
	if l > 0 {
		n += 1 + l + sovAlliance(uint64(l))
	}
	return n
}

func (m *MemberInfoOnLoginRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AllianceId != 0 {
		n += 1 + sovAlliance(uint64(m.AllianceId))
	}
	if m.Position != 0 {
		n += 1 + sovAlliance(uint64(m.Position))
	}
	return n
}

func (m *MemberInfoOnLogoutReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GateSession)
	if l > 0 {
		n += 1 + l + sovAlliance(uint64(l))
	}
	l = len(m.RID)
	if l > 0 {
		n += 1 + l + sovAlliance(uint64(l))
	}
	return n
}

func sovAlliance(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAlliance(x uint64) (n int) {
	return sovAlliance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateAllianceReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlliance
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
			return fmt.Errorf("proto: CreateAllianceReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateAllianceReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterShortId", wireType)
			}
			m.MasterShortId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlliance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MasterShortId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAlliance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlliance
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
func (m *CreateAllianceRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlliance
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
			return fmt.Errorf("proto: CreateAllianceRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateAllianceRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllianceId", wireType)
			}
			m.AllianceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlliance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllianceId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAlliance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlliance
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
func (m *SetMemberReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlliance
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
			return fmt.Errorf("proto: SetMemberReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetMemberReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Players", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlliance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAlliance
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAlliance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Players = append(m.Players, &PlayerInfo{})
			if err := m.Players[len(m.Players)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlliance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlliance
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
func (m *SetMemberRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlliance
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
			return fmt.Errorf("proto: SetMemberRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetMemberRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAlliance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlliance
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
func (m *MemberInfoOnLoginReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlliance
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
			return fmt.Errorf("proto: MemberInfoOnLoginReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberInfoOnLoginReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GateSession", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlliance
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
				return ErrInvalidLengthAlliance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlliance
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
					return ErrIntOverflowAlliance
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
				return ErrInvalidLengthAlliance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlliance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlliance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlliance
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
func (m *MemberInfoOnLoginRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlliance
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
			return fmt.Errorf("proto: MemberInfoOnLoginRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberInfoOnLoginRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllianceId", wireType)
			}
			m.AllianceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlliance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllianceId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			m.Position = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlliance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Position |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAlliance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlliance
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
func (m *MemberInfoOnLogoutReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAlliance
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
			return fmt.Errorf("proto: MemberInfoOnLogoutReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberInfoOnLogoutReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GateSession", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAlliance
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
				return ErrInvalidLengthAlliance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlliance
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
					return ErrIntOverflowAlliance
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
				return ErrInvalidLengthAlliance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAlliance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAlliance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAlliance
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
func skipAlliance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAlliance
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
					return 0, ErrIntOverflowAlliance
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
					return 0, ErrIntOverflowAlliance
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
				return 0, ErrInvalidLengthAlliance
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAlliance
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAlliance
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAlliance        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAlliance          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAlliance = fmt.Errorf("proto: unexpected end of group")
)
