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

type SetMemberReq struct {
	Players []*PlayerInfo `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (m *SetMemberReq) Reset()         { *m = SetMemberReq{} }
func (m *SetMemberReq) String() string { return proto.CompactTextString(m) }
func (*SetMemberReq) ProtoMessage()    {}
func (*SetMemberReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d63a1f01d82564b, []int{0}
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
	return fileDescriptor_9d63a1f01d82564b, []int{1}
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

func init() {
	proto.RegisterType((*SetMemberReq)(nil), "inner.SetMemberReq")
	proto.RegisterType((*SetMemberRsp)(nil), "inner.SetMemberRsp")
}

func init() { proto.RegisterFile("alliance.proto", fileDescriptor_9d63a1f01d82564b) }

var fileDescriptor_9d63a1f01d82564b = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcc, 0xc9, 0xc9,
	0x4c, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcc, 0xcb, 0x4b,
	0x2d, 0x92, 0xe2, 0x4a, 0x4a, 0x2c, 0x86, 0x0a, 0x29, 0x59, 0x73, 0xf1, 0x04, 0xa7, 0x96, 0xf8,
	0xa6, 0xe6, 0x26, 0xa5, 0x16, 0x05, 0xa5, 0x16, 0x0a, 0x69, 0x73, 0xb1, 0x17, 0xe4, 0x24, 0x56,
	0xa6, 0x16, 0x15, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x09, 0xea, 0x81, 0x35, 0xe9, 0x05,
	0x80, 0x45, 0x3d, 0xf3, 0xd2, 0xf2, 0x83, 0x60, 0x2a, 0x94, 0xf8, 0x90, 0x35, 0x17, 0x17, 0x38,
	0x29, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0x9b, 0x3e, 0xd8, 0x94,
	0x24, 0x36, 0xb0, 0xad, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xe6, 0x9e, 0x66, 0x9a,
	0x00, 0x00, 0x00,
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

func sovAlliance(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAlliance(x uint64) (n int) {
	return sovAlliance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
