// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: mail.proto

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

type Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string          `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	CreateAt     int64           `protobuf:"varint,2,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`                                                                                    // 邮件发送时间
	SenderRoleId string          `protobuf:"bytes,3,opt,name=SenderRoleId,proto3" json:"SenderRoleId,omitempty"`                                                                             // 发送者RoleId
	Name         string          `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`                                                                                             // 发送者名字
	Title        string          `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`                                                                                           // 标题
	Content      string          `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"`                                                                                       // 正文
	Items        map[int64]int64 `protobuf:"bytes,7,rep,name=Items,proto3" json:"Items,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 附件道具
	Status       int32           `protobuf:"varint,8,opt,name=Status,proto3" json:"Status,omitempty"`                                                                                        // 状态 0.未读 1.已读 2.已领
}

func (x *Mail) Reset() {
	*x = Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mail) ProtoMessage() {}

func (x *Mail) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mail.ProtoReflect.Descriptor instead.
func (*Mail) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{0}
}

func (x *Mail) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Mail) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Mail) GetSenderRoleId() string {
	if x != nil {
		return x.SenderRoleId
	}
	return ""
}

func (x *Mail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Mail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Mail) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Mail) GetItems() map[int64]int64 {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Mail) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type MailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mails map[string]*Mail `protobuf:"bytes,1,rep,name=Mails,proto3" json:"Mails,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 所有邮件
}

func (x *MailInfo) Reset() {
	*x = MailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailInfo) ProtoMessage() {}

func (x *MailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailInfo.ProtoReflect.Descriptor instead.
func (*MailInfo) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{1}
}

func (x *MailInfo) GetMails() map[string]*Mail {
	if x != nil {
		return x.Mails
	}
	return nil
}

type MailListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *MailListReq) Reset() {
	*x = MailListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailListReq) ProtoMessage() {}

func (x *MailListReq) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailListReq.ProtoReflect.Descriptor instead.
func (*MailListReq) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{2}
}

func (x *MailListReq) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type MailListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mails []*Mail `protobuf:"bytes,1,rep,name=Mails,proto3" json:"Mails,omitempty"`
}

func (x *MailListRsp) Reset() {
	*x = MailListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailListRsp) ProtoMessage() {}

func (x *MailListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailListRsp.ProtoReflect.Descriptor instead.
func (*MailListRsp) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{3}
}

func (x *MailListRsp) GetMails() []*Mail {
	if x != nil {
		return x.Mails
	}
	return nil
}

type AddMailNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
}

func (x *AddMailNotify) Reset() {
	*x = AddMailNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMailNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMailNotify) ProtoMessage() {}

func (x *AddMailNotify) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMailNotify.ProtoReflect.Descriptor instead.
func (*AddMailNotify) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{4}
}

func (x *AddMailNotify) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ReadMailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
}

func (x *ReadMailReq) Reset() {
	*x = ReadMailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadMailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadMailReq) ProtoMessage() {}

func (x *ReadMailReq) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadMailReq.ProtoReflect.Descriptor instead.
func (*ReadMailReq) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{5}
}

func (x *ReadMailReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ReadMailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
}

func (x *ReadMailRsp) Reset() {
	*x = ReadMailRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadMailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadMailRsp) ProtoMessage() {}

func (x *ReadMailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadMailRsp.ProtoReflect.Descriptor instead.
func (*ReadMailRsp) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{6}
}

func (x *ReadMailRsp) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ReceiveMailItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
}

func (x *ReceiveMailItemReq) Reset() {
	*x = ReceiveMailItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMailItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMailItemReq) ProtoMessage() {}

func (x *ReceiveMailItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMailItemReq.ProtoReflect.Descriptor instead.
func (*ReceiveMailItemReq) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{7}
}

func (x *ReceiveMailItemReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ReceiveMailItemRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
}

func (x *ReceiveMailItemRsp) Reset() {
	*x = ReceiveMailItemRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMailItemRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMailItemRsp) ProtoMessage() {}

func (x *ReceiveMailItemRsp) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMailItemRsp.ProtoReflect.Descriptor instead.
func (*ReceiveMailItemRsp) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{8}
}

func (x *ReceiveMailItemRsp) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type DeleteMailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuids []string `protobuf:"bytes,1,rep,name=Uuids,proto3" json:"Uuids,omitempty"`
}

func (x *DeleteMailReq) Reset() {
	*x = DeleteMailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMailReq) ProtoMessage() {}

func (x *DeleteMailReq) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMailReq.ProtoReflect.Descriptor instead.
func (*DeleteMailReq) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteMailReq) GetUuids() []string {
	if x != nil {
		return x.Uuids
	}
	return nil
}

var File_mail_proto protoreflect.FileDescriptor

var file_mail_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x22, 0x9e, 0x02, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6c,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x30, 0x0a, 0x05, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x4d, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x45, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x23, 0x0a, 0x0b, 0x4d, 0x61,
	0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x30, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x21,
	0x0a, 0x05, 0x4d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x4d, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0x23, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x61,
	0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x12,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64,
	0x22, 0x25, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x55, 0x75, 0x69, 0x64, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mail_proto_rawDescOnce sync.Once
	file_mail_proto_rawDescData = file_mail_proto_rawDesc
)

func file_mail_proto_rawDescGZIP() []byte {
	file_mail_proto_rawDescOnce.Do(func() {
		file_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_mail_proto_rawDescData)
	})
	return file_mail_proto_rawDescData
}

var file_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_mail_proto_goTypes = []interface{}{
	(*Mail)(nil),               // 0: outer.Mail
	(*MailInfo)(nil),           // 1: outer.MailInfo
	(*MailListReq)(nil),        // 2: outer.MailListReq
	(*MailListRsp)(nil),        // 3: outer.MailListRsp
	(*AddMailNotify)(nil),      // 4: outer.AddMailNotify
	(*ReadMailReq)(nil),        // 5: outer.ReadMailReq
	(*ReadMailRsp)(nil),        // 6: outer.ReadMailRsp
	(*ReceiveMailItemReq)(nil), // 7: outer.ReceiveMailItemReq
	(*ReceiveMailItemRsp)(nil), // 8: outer.ReceiveMailItemRsp
	(*DeleteMailReq)(nil),      // 9: outer.DeleteMailReq
	nil,                        // 10: outer.Mail.ItemsEntry
	nil,                        // 11: outer.MailInfo.MailsEntry
}
var file_mail_proto_depIdxs = []int32{
	10, // 0: outer.Mail.Items:type_name -> outer.Mail.ItemsEntry
	11, // 1: outer.MailInfo.Mails:type_name -> outer.MailInfo.MailsEntry
	0,  // 2: outer.MailListRsp.Mails:type_name -> outer.Mail
	0,  // 3: outer.MailInfo.MailsEntry.value:type_name -> outer.Mail
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_mail_proto_init() }
func file_mail_proto_init() {
	if File_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailListRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMailNotify); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadMailReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mail_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadMailRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mail_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveMailItemReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mail_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveMailItemRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mail_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMailReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mail_proto_goTypes,
		DependencyIndexes: file_mail_proto_depIdxs,
		MessageInfos:      file_mail_proto_msgTypes,
	}.Build()
	File_mail_proto = out.File
	file_mail_proto_rawDesc = nil
	file_mail_proto_goTypes = nil
	file_mail_proto_depIdxs = nil
}
