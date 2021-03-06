// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: channel.proto

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

type ChatReq struct {
	Content     string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	ChannelType int32  `protobuf:"varint,2,opt,name=ChannelType,proto3" json:"ChannelType,omitempty"`
}

func (m *ChatReq) Reset()         { *m = ChatReq{} }
func (m *ChatReq) String() string { return proto.CompactTextString(m) }
func (*ChatReq) ProtoMessage()    {}
func (*ChatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}
func (m *ChatReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChatReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatReq.Merge(m, src)
}
func (m *ChatReq) XXX_Size() int {
	return m.Size()
}
func (m *ChatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChatReq proto.InternalMessageInfo

func (m *ChatReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ChatReq) GetChannelType() int32 {
	if m != nil {
		return m.ChannelType
	}
	return 0
}

type ChatResp struct {
	Err ERROR `protobuf:"varint,1,opt,name=err,proto3,enum=outer.ERROR" json:"err,omitempty"`
}

func (m *ChatResp) Reset()         { *m = ChatResp{} }
func (m *ChatResp) String() string { return proto.CompactTextString(m) }
func (*ChatResp) ProtoMessage()    {}
func (*ChatResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}
func (m *ChatResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChatResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChatResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChatResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatResp.Merge(m, src)
}
func (m *ChatResp) XXX_Size() int {
	return m.Size()
}
func (m *ChatResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatResp.DiscardUnknown(m)
}

var xxx_messageInfo_ChatResp proto.InternalMessageInfo

func (m *ChatResp) GetErr() ERROR {
	if m != nil {
		return m.Err
	}
	return ERROR_SUCCESS
}

type ChatNotify struct {
	SenderId    uint64 `protobuf:"varint,1,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	ChannelType int32  `protobuf:"varint,4,opt,name=ChannelType,proto3" json:"ChannelType,omitempty"`
}

func (m *ChatNotify) Reset()         { *m = ChatNotify{} }
func (m *ChatNotify) String() string { return proto.CompactTextString(m) }
func (*ChatNotify) ProtoMessage()    {}
func (*ChatNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}
func (m *ChatNotify) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChatNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChatNotify.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChatNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatNotify.Merge(m, src)
}
func (m *ChatNotify) XXX_Size() int {
	return m.Size()
}
func (m *ChatNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatNotify.DiscardUnknown(m)
}

var xxx_messageInfo_ChatNotify proto.InternalMessageInfo

func (m *ChatNotify) GetSenderId() uint64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *ChatNotify) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChatNotify) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ChatNotify) GetChannelType() int32 {
	if m != nil {
		return m.ChannelType
	}
	return 0
}

func init() {
	proto.RegisterType((*ChatReq)(nil), "outer.ChatReq")
	proto.RegisterType((*ChatResp)(nil), "outer.ChatResp")
	proto.RegisterType((*ChatNotify)(nil), "outer.ChatNotify")
}

func init() { proto.RegisterFile("channel.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0x2f, 0x2d, 0x49, 0x2d, 0x92, 0xe2, 0x4e,
	0x2d, 0x2a, 0xca, 0x2f, 0x82, 0x88, 0x29, 0xb9, 0x72, 0xb1, 0x3b, 0x67, 0x24, 0x96, 0x04, 0xa5,
	0x16, 0x0a, 0x49, 0x70, 0xb1, 0x3b, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x0a, 0x5c, 0xdc, 0xce, 0x19, 0x89, 0x79, 0x79, 0xa9, 0x39,
	0x21, 0x95, 0x05, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0xc8, 0x42, 0x4a, 0x5a, 0x5c,
	0x1c, 0x10, 0x63, 0x8a, 0x0b, 0x84, 0xe4, 0xb8, 0x98, 0x53, 0x8b, 0x8a, 0xc0, 0x66, 0xf0, 0x19,
	0xf1, 0xe8, 0x81, 0x2d, 0xd5, 0x73, 0x0d, 0x0a, 0xf2, 0x0f, 0x0a, 0x02, 0x49, 0x28, 0x55, 0x70,
	0x71, 0x81, 0xd4, 0xfa, 0xe5, 0x97, 0x64, 0xa6, 0x55, 0x0a, 0x49, 0x71, 0x71, 0x04, 0xa7, 0xe6,
	0xa5, 0xa4, 0x16, 0x79, 0xa6, 0x80, 0xb5, 0xb0, 0x04, 0xc1, 0xf9, 0x42, 0x42, 0x5c, 0x2c, 0x7e,
	0x89, 0xb9, 0x10, 0x0b, 0x39, 0x83, 0xc0, 0x6c, 0x64, 0x57, 0x32, 0xe3, 0x75, 0x25, 0x0b, 0x86,
	0x2b, 0x9d, 0x14, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x8a, 0x4d, 0x1f,
	0xec, 0xcc, 0x24, 0x36, 0x70, 0xa8, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38, 0x80, 0x1b,
	0x2b, 0x37, 0x01, 0x00, 0x00,
}

func (m *ChatReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChatReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ChannelType != 0 {
		i = encodeVarintChat(dAtA, i, uint64(m.ChannelType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintChat(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChatResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChatResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Err != 0 {
		i = encodeVarintChat(dAtA, i, uint64(m.Err))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ChatNotify) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChatNotify) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChatNotify) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ChannelType != 0 {
		i = encodeVarintChat(dAtA, i, uint64(m.ChannelType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintChat(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintChat(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.SenderId != 0 {
		i = encodeVarintChat(dAtA, i, uint64(m.SenderId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintChat(dAtA []byte, offset int, v uint64) int {
	offset -= sovChat(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChatReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovChat(uint64(l))
	}
	if m.ChannelType != 0 {
		n += 1 + sovChat(uint64(m.ChannelType))
	}
	return n
}

func (m *ChatResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Err != 0 {
		n += 1 + sovChat(uint64(m.Err))
	}
	return n
}

func (m *ChatNotify) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SenderId != 0 {
		n += 1 + sovChat(uint64(m.SenderId))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovChat(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovChat(uint64(l))
	}
	if m.ChannelType != 0 {
		n += 1 + sovChat(uint64(m.ChannelType))
	}
	return n
}

func sovChat(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChat(x uint64) (n int) {
	return sovChat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChatReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: ChatReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelType", wireType)
			}
			m.ChannelType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChannelType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChat
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
func (m *ChatResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: ChatResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			m.Err = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Err |= ERROR(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChat
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
func (m *ChatNotify) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: ChatNotify: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChatNotify: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderId", wireType)
			}
			m.SenderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SenderId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelType", wireType)
			}
			m.ChannelType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChannelType |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChat
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
func skipChat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChat
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
					return 0, ErrIntOverflowChat
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
					return 0, ErrIntOverflowChat
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
				return 0, ErrInvalidLengthChat
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChat
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChat
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChat        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChat          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChat = fmt.Errorf("proto: unexpected end of group")
)
