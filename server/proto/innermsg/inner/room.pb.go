// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: room.proto

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

type RoomInfo struct {
	RoomId   int32         `protobuf:"varint,1,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	GameType int32         `protobuf:"varint,2,opt,name=GameType,proto3" json:"GameType,omitempty"`
	Players  []*PlayerInfo `protobuf:"bytes,3,rep,name=Players,proto3" json:"Players,omitempty"`
}

func (m *RoomInfo) Reset()         { *m = RoomInfo{} }
func (m *RoomInfo) String() string { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()    {}
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{0}
}
func (m *RoomInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoomInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInfo.Merge(m, src)
}
func (m *RoomInfo) XXX_Size() int {
	return m.Size()
}
func (m *RoomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInfo proto.InternalMessageInfo

func (m *RoomInfo) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RoomInfo) GetGameType() int32 {
	if m != nil {
		return m.GameType
	}
	return 0
}

func (m *RoomInfo) GetPlayers() []*PlayerInfo {
	if m != nil {
		return m.Players
	}
	return nil
}

// 创建房间
type CreateRoomReq struct {
	CreatorShortId int64 `protobuf:"varint,1,opt,name=CreatorShortId,proto3" json:"CreatorShortId,omitempty"`
}

func (m *CreateRoomReq) Reset()         { *m = CreateRoomReq{} }
func (m *CreateRoomReq) String() string { return proto.CompactTextString(m) }
func (*CreateRoomReq) ProtoMessage()    {}
func (*CreateRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{1}
}
func (m *CreateRoomReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateRoomReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomReq.Merge(m, src)
}
func (m *CreateRoomReq) XXX_Size() int {
	return m.Size()
}
func (m *CreateRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomReq proto.InternalMessageInfo

func (m *CreateRoomReq) GetCreatorShortId() int64 {
	if m != nil {
		return m.CreatorShortId
	}
	return 0
}

type CreateRoomRsp struct {
	RoomInfo *RoomInfo `protobuf:"bytes,1,opt,name=RoomInfo,proto3" json:"RoomInfo,omitempty"`
}

func (m *CreateRoomRsp) Reset()         { *m = CreateRoomRsp{} }
func (m *CreateRoomRsp) String() string { return proto.CompactTextString(m) }
func (*CreateRoomRsp) ProtoMessage()    {}
func (*CreateRoomRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{2}
}
func (m *CreateRoomRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateRoomRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateRoomRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateRoomRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomRsp.Merge(m, src)
}
func (m *CreateRoomRsp) XXX_Size() int {
	return m.Size()
}
func (m *CreateRoomRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomRsp proto.InternalMessageInfo

func (m *CreateRoomRsp) GetRoomInfo() *RoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

// 获得房间基础信息
type RoomInfoReq struct {
}

func (m *RoomInfoReq) Reset()         { *m = RoomInfoReq{} }
func (m *RoomInfoReq) String() string { return proto.CompactTextString(m) }
func (*RoomInfoReq) ProtoMessage()    {}
func (*RoomInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{3}
}
func (m *RoomInfoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoomInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoomInfoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoomInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInfoReq.Merge(m, src)
}
func (m *RoomInfoReq) XXX_Size() int {
	return m.Size()
}
func (m *RoomInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInfoReq proto.InternalMessageInfo

type RoomInfoRsp struct {
	RoomInfo *RoomInfo `protobuf:"bytes,1,opt,name=RoomInfo,proto3" json:"RoomInfo,omitempty"`
}

func (m *RoomInfoRsp) Reset()         { *m = RoomInfoRsp{} }
func (m *RoomInfoRsp) String() string { return proto.CompactTextString(m) }
func (*RoomInfoRsp) ProtoMessage()    {}
func (*RoomInfoRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{4}
}
func (m *RoomInfoRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoomInfoRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoomInfoRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoomInfoRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInfoRsp.Merge(m, src)
}
func (m *RoomInfoRsp) XXX_Size() int {
	return m.Size()
}
func (m *RoomInfoRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInfoRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInfoRsp proto.InternalMessageInfo

func (m *RoomInfoRsp) GetRoomInfo() *RoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

// 加入房间
type JoinRoomReq struct {
	Player *PlayerInfo `protobuf:"bytes,1,opt,name=Player,proto3" json:"Player,omitempty"`
}

func (m *JoinRoomReq) Reset()         { *m = JoinRoomReq{} }
func (m *JoinRoomReq) String() string { return proto.CompactTextString(m) }
func (*JoinRoomReq) ProtoMessage()    {}
func (*JoinRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{5}
}
func (m *JoinRoomReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinRoomReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomReq.Merge(m, src)
}
func (m *JoinRoomReq) XXX_Size() int {
	return m.Size()
}
func (m *JoinRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomReq proto.InternalMessageInfo

func (m *JoinRoomReq) GetPlayer() *PlayerInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

type JoinRoomRsp struct {
}

func (m *JoinRoomRsp) Reset()         { *m = JoinRoomRsp{} }
func (m *JoinRoomRsp) String() string { return proto.CompactTextString(m) }
func (*JoinRoomRsp) ProtoMessage()    {}
func (*JoinRoomRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{6}
}
func (m *JoinRoomRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinRoomRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinRoomRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinRoomRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomRsp.Merge(m, src)
}
func (m *JoinRoomRsp) XXX_Size() int {
	return m.Size()
}
func (m *JoinRoomRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomRsp.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomRsp proto.InternalMessageInfo

// 离开房间
type LeaveRoomReq struct {
	ShortId int64 `protobuf:"varint,1,opt,name=ShortId,proto3" json:"ShortId,omitempty"`
}

func (m *LeaveRoomReq) Reset()         { *m = LeaveRoomReq{} }
func (m *LeaveRoomReq) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomReq) ProtoMessage()    {}
func (*LeaveRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{7}
}
func (m *LeaveRoomReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaveRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaveRoomReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaveRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomReq.Merge(m, src)
}
func (m *LeaveRoomReq) XXX_Size() int {
	return m.Size()
}
func (m *LeaveRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomReq proto.InternalMessageInfo

func (m *LeaveRoomReq) GetShortId() int64 {
	if m != nil {
		return m.ShortId
	}
	return 0
}

type LeaveRoomRsp struct {
}

func (m *LeaveRoomRsp) Reset()         { *m = LeaveRoomRsp{} }
func (m *LeaveRoomRsp) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomRsp) ProtoMessage()    {}
func (*LeaveRoomRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{8}
}
func (m *LeaveRoomRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LeaveRoomRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LeaveRoomRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LeaveRoomRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomRsp.Merge(m, src)
}
func (m *LeaveRoomRsp) XXX_Size() int {
	return m.Size()
}
func (m *LeaveRoomRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomRsp.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomRsp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RoomInfo)(nil), "inner.RoomInfo")
	proto.RegisterType((*CreateRoomReq)(nil), "inner.CreateRoomReq")
	proto.RegisterType((*CreateRoomRsp)(nil), "inner.CreateRoomRsp")
	proto.RegisterType((*RoomInfoReq)(nil), "inner.RoomInfoReq")
	proto.RegisterType((*RoomInfoRsp)(nil), "inner.RoomInfoRsp")
	proto.RegisterType((*JoinRoomReq)(nil), "inner.JoinRoomReq")
	proto.RegisterType((*JoinRoomRsp)(nil), "inner.JoinRoomRsp")
	proto.RegisterType((*LeaveRoomReq)(nil), "inner.LeaveRoomReq")
	proto.RegisterType((*LeaveRoomRsp)(nil), "inner.LeaveRoomRsp")
}

func init() { proto.RegisterFile("room.proto", fileDescriptor_c5fd27dd97284ef4) }

var fileDescriptor_c5fd27dd97284ef4 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xca, 0xcf, 0xcf,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcc, 0xcb, 0x4b, 0x2d, 0x92, 0xe2, 0x4a,
	0x4a, 0x2c, 0x4e, 0x85, 0x08, 0x29, 0x65, 0x73, 0x71, 0x04, 0xe5, 0xe7, 0xe7, 0x7a, 0xe6, 0xa5,
	0xe5, 0x0b, 0x89, 0x71, 0xb1, 0x81, 0xd9, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50,
	0x9e, 0x90, 0x14, 0x17, 0x87, 0x7b, 0x62, 0x6e, 0x6a, 0x48, 0x65, 0x41, 0xaa, 0x04, 0x13, 0x58,
	0x06, 0xce, 0x17, 0xd2, 0xe6, 0x62, 0x0f, 0xc8, 0x49, 0xac, 0x4c, 0x2d, 0x2a, 0x96, 0x60, 0x56,
	0x60, 0xd6, 0xe0, 0x36, 0x12, 0xd4, 0x03, 0x5b, 0xa2, 0x07, 0x11, 0x05, 0x99, 0x1b, 0x04, 0x53,
	0xa1, 0x64, 0xce, 0xc5, 0xeb, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x0a, 0x32, 0x38, 0x28, 0xb5, 0x50,
	0x48, 0x8d, 0x8b, 0x0f, 0x2c, 0x90, 0x5f, 0x14, 0x9c, 0x91, 0x5f, 0x54, 0x02, 0xb5, 0x99, 0x39,
	0x08, 0x4d, 0x54, 0xc9, 0x06, 0x45, 0x63, 0x71, 0x81, 0x90, 0x36, 0xc2, 0xd9, 0x60, 0x2d, 0xdc,
	0x46, 0xfc, 0x50, 0x7b, 0x61, 0xc2, 0x41, 0x70, 0x05, 0x4a, 0xbc, 0x5c, 0xdc, 0x70, 0xd1, 0xd4,
	0x42, 0x25, 0x2b, 0x24, 0x2e, 0xa9, 0x46, 0x59, 0x70, 0x71, 0x7b, 0xe5, 0x67, 0xe6, 0xc1, 0xdc,
	0xaf, 0xc9, 0xc5, 0x06, 0xf1, 0x1b, 0x54, 0x27, 0x16, 0xcf, 0x43, 0x15, 0x80, 0x1c, 0x01, 0xd7,
	0x59, 0x5c, 0xa0, 0xa4, 0xc1, 0xc5, 0xe3, 0x93, 0x9a, 0x58, 0x06, 0x0f, 0x09, 0x09, 0x2e, 0x76,
	0xd4, 0x20, 0x80, 0x71, 0x95, 0xf8, 0x90, 0x55, 0x16, 0x17, 0x38, 0x29, 0x9c, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c,
	0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0x9b, 0x3e, 0xd8, 0xf6, 0x24, 0x36, 0x70, 0xd4, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xc6, 0x74, 0x4f, 0xfb, 0x01, 0x00, 0x00,
}

func (m *RoomInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoomInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoomInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintRoom(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.GameType != 0 {
		i = encodeVarintRoom(dAtA, i, uint64(m.GameType))
		i--
		dAtA[i] = 0x10
	}
	if m.RoomId != 0 {
		i = encodeVarintRoom(dAtA, i, uint64(m.RoomId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CreateRoomReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRoomReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRoomReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatorShortId != 0 {
		i = encodeVarintRoom(dAtA, i, uint64(m.CreatorShortId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CreateRoomRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRoomRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRoomRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RoomInfo != nil {
		{
			size, err := m.RoomInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRoom(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RoomInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoomInfoReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoomInfoReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *RoomInfoRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoomInfoRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoomInfoRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RoomInfo != nil {
		{
			size, err := m.RoomInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRoom(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JoinRoomReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinRoomReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinRoomReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Player != nil {
		{
			size, err := m.Player.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRoom(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JoinRoomRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinRoomRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JoinRoomRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *LeaveRoomReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaveRoomReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaveRoomReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ShortId != 0 {
		i = encodeVarintRoom(dAtA, i, uint64(m.ShortId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LeaveRoomRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LeaveRoomRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LeaveRoomRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintRoom(dAtA []byte, offset int, v uint64) int {
	offset -= sovRoom(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RoomInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoomId != 0 {
		n += 1 + sovRoom(uint64(m.RoomId))
	}
	if m.GameType != 0 {
		n += 1 + sovRoom(uint64(m.GameType))
	}
	if len(m.Players) > 0 {
		for _, e := range m.Players {
			l = e.Size()
			n += 1 + l + sovRoom(uint64(l))
		}
	}
	return n
}

func (m *CreateRoomReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreatorShortId != 0 {
		n += 1 + sovRoom(uint64(m.CreatorShortId))
	}
	return n
}

func (m *CreateRoomRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoomInfo != nil {
		l = m.RoomInfo.Size()
		n += 1 + l + sovRoom(uint64(l))
	}
	return n
}

func (m *RoomInfoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *RoomInfoRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoomInfo != nil {
		l = m.RoomInfo.Size()
		n += 1 + l + sovRoom(uint64(l))
	}
	return n
}

func (m *JoinRoomReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Player != nil {
		l = m.Player.Size()
		n += 1 + l + sovRoom(uint64(l))
	}
	return n
}

func (m *JoinRoomRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *LeaveRoomReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ShortId != 0 {
		n += 1 + sovRoom(uint64(m.ShortId))
	}
	return n
}

func (m *LeaveRoomRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovRoom(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRoom(x uint64) (n int) {
	return sovRoom(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RoomInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoom
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
			return fmt.Errorf("proto: RoomInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoomInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomId", wireType)
			}
			m.RoomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoomId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameType", wireType)
			}
			m.GameType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Players", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
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
				return ErrInvalidLengthRoom
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoom
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
			skippy, err := skipRoom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoom
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
func (m *CreateRoomReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoom
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
			return fmt.Errorf("proto: CreateRoomReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRoomReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatorShortId", wireType)
			}
			m.CreatorShortId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatorShortId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRoom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoom
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
func (m *CreateRoomRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoom
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
			return fmt.Errorf("proto: CreateRoomRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRoomRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
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
				return ErrInvalidLengthRoom
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RoomInfo == nil {
				m.RoomInfo = &RoomInfo{}
			}
			if err := m.RoomInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoom
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
func (m *RoomInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoom
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
			return fmt.Errorf("proto: RoomInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoomInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRoom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoom
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
func (m *RoomInfoRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoom
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
			return fmt.Errorf("proto: RoomInfoRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoomInfoRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
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
				return ErrInvalidLengthRoom
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RoomInfo == nil {
				m.RoomInfo = &RoomInfo{}
			}
			if err := m.RoomInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoom
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
func (m *JoinRoomReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoom
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
			return fmt.Errorf("proto: JoinRoomReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinRoomReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
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
				return ErrInvalidLengthRoom
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRoom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Player == nil {
				m.Player = &PlayerInfo{}
			}
			if err := m.Player.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoom
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
func (m *JoinRoomRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoom
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
			return fmt.Errorf("proto: JoinRoomRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinRoomRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRoom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoom
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
func (m *LeaveRoomReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoom
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
			return fmt.Errorf("proto: LeaveRoomReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaveRoomReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortId", wireType)
			}
			m.ShortId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
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
			skippy, err := skipRoom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoom
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
func (m *LeaveRoomRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoom
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
			return fmt.Errorf("proto: LeaveRoomRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaveRoomRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRoom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoom
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
func skipRoom(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoom
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
					return 0, ErrIntOverflowRoom
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
					return 0, ErrIntOverflowRoom
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
				return 0, ErrInvalidLengthRoom
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRoom
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRoom
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRoom        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoom          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRoom = fmt.Errorf("proto: unexpected end of group")
)