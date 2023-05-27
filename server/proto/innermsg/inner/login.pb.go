// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: login.proto

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

type BindSessionWithRID struct {
	GateSession string `protobuf:"bytes,1,opt,name=GateSession,proto3" json:"GateSession,omitempty"`
	RID         string `protobuf:"bytes,2,opt,name=RID,proto3" json:"RID,omitempty"`
}

func (m *BindSessionWithRID) Reset()         { *m = BindSessionWithRID{} }
func (m *BindSessionWithRID) String() string { return proto.CompactTextString(m) }
func (*BindSessionWithRID) ProtoMessage()    {}
func (*BindSessionWithRID) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
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
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
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
	return fileDescriptor_67c21677aa7f4e4f, []int{2}
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
	return fileDescriptor_67c21677aa7f4e4f, []int{3}
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

func init() {
	proto.RegisterType((*BindSessionWithRID)(nil), "inner.BindSessionWithRID")
	proto.RegisterType((*KickOutReq)(nil), "inner.KickOutReq")
	proto.RegisterType((*KickOutRsp)(nil), "inner.KickOutRsp")
	proto.RegisterType((*GSessionClosed)(nil), "inner.GSessionClosed")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x4f, 0xcf,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcc, 0xcb, 0x4b, 0x2d, 0x52, 0xf2,
	0xe0, 0x12, 0x72, 0xca, 0xcc, 0x4b, 0x09, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x0b, 0xcf, 0x2c,
	0xc9, 0x08, 0xf2, 0x74, 0x11, 0x52, 0xe0, 0xe2, 0x76, 0x4f, 0x2c, 0x49, 0x85, 0x8a, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x0b, 0x09, 0x09, 0x70, 0x31, 0x07, 0x79, 0xba, 0x48, 0x30,
	0x81, 0x65, 0x40, 0x4c, 0x25, 0x07, 0x2e, 0x2e, 0xef, 0xcc, 0xe4, 0x6c, 0xff, 0xd2, 0x92, 0xa0,
	0xd4, 0x42, 0xb2, 0x4c, 0xe0, 0x41, 0x98, 0x50, 0x5c, 0xa0, 0x64, 0xc4, 0xc5, 0xe7, 0x0e, 0x55,
	0xeb, 0x9c, 0x93, 0x5f, 0x9c, 0x9a, 0x42, 0xd8, 0x4c, 0x27, 0x85, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0x62, 0xd3, 0x07, 0xfb, 0x37, 0x89, 0x0d, 0xec, 0x7b, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x73, 0x3d, 0xab, 0x0c, 0x01, 0x00, 0x00,
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
		i = encodeVarintLogin(dAtA, i, uint64(len(m.RID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GateSession) > 0 {
		i -= len(m.GateSession)
		copy(dAtA[i:], m.GateSession)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.GateSession)))
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
		i = encodeVarintLogin(dAtA, i, uint64(len(m.RID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GateSession) > 0 {
		i -= len(m.GateSession)
		copy(dAtA[i:], m.GateSession)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.GateSession)))
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
		i = encodeVarintLogin(dAtA, i, uint64(len(m.GateSession)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLogin(dAtA []byte, offset int, v uint64) int {
	offset -= sovLogin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BindSessionWithRID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GateSession)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	l = len(m.RID)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
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
		n += 1 + l + sovLogin(uint64(l))
	}
	l = len(m.RID)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
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
		n += 1 + l + sovLogin(uint64(l))
	}
	return n
}

func sovLogin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLogin(x uint64) (n int) {
	return sovLogin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BindSessionWithRID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
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
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
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
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLogin
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
				return ErrIntOverflowLogin
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
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
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
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLogin
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
				return ErrIntOverflowLogin
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
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLogin
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
				return ErrIntOverflowLogin
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
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GateSession = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLogin
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
func skipLogin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogin
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
					return 0, ErrIntOverflowLogin
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
					return 0, ErrIntOverflowLogin
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
				return 0, ErrInvalidLengthLogin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLogin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLogin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLogin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLogin = fmt.Errorf("proto: unexpected end of group")
)
