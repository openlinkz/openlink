// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.5
// source: api/msg_api/msg_exchange.proto

package msg_api

import (
	protocol "github.com/openlinkz/openlink/api/protocol"
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

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   string             `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	SID      string             `protobuf:"bytes,2,opt,name=SID,proto3" json:"SID,omitempty"`
	UID      string             `protobuf:"bytes,3,opt,name=UID,proto3" json:"UID,omitempty"`
	Platform protocol.Platform  `protobuf:"varint,4,opt,name=platform,proto3,enum=msg.protocol.Platform" json:"platform,omitempty"`
	Protocol *protocol.Protocol `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *Msg) GetSID() string {
	if x != nil {
		return x.SID
	}
	return ""
}

func (x *Msg) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *Msg) GetPlatform() protocol.Platform {
	if x != nil {
		return x.Platform
	}
	return protocol.Platform(0)
}

func (x *Msg) GetProtocol() *protocol.Protocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

type UserConnectionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserConnectionInfo) Reset() {
	*x = UserConnectionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserConnectionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConnectionInfo) ProtoMessage() {}

func (x *UserConnectionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConnectionInfo.ProtoReflect.Descriptor instead.
func (*UserConnectionInfo) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_exchange_proto_rawDescGZIP(), []int{1}
}

type SendMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMsgReply) Reset() {
	*x = SendMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgReply) ProtoMessage() {}

func (x *SendMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgReply.ProtoReflect.Descriptor instead.
func (*SendMsgReply) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_exchange_proto_rawDescGZIP(), []int{2}
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	SID      string `protobuf:"bytes,2,opt,name=SID,proto3" json:"SID,omitempty"`
	UID      string `protobuf:"bytes,3,opt,name=UID,proto3" json:"UID,omitempty"`
	Platform string `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *ConnectRequest) GetSID() string {
	if x != nil {
		return x.SID
	}
	return ""
}

func (x *ConnectRequest) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *ConnectRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

type ConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectReply) Reset() {
	*x = ConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_exchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReply) ProtoMessage() {}

func (x *ConnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_exchange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectReply.ProtoReflect.Descriptor instead.
func (*ConnectReply) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_exchange_proto_rawDescGZIP(), []int{4}
}

type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	SID    string `protobuf:"bytes,2,opt,name=SID,proto3" json:"SID,omitempty"`
	UID    string `protobuf:"bytes,3,opt,name=UID,proto3" json:"UID,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_exchange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_exchange_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_exchange_proto_rawDescGZIP(), []int{5}
}

func (x *DisconnectRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *DisconnectRequest) GetSID() string {
	if x != nil {
		return x.SID
	}
	return ""
}

func (x *DisconnectRequest) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

type DisconnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectReply) Reset() {
	*x = DisconnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_exchange_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectReply) ProtoMessage() {}

func (x *DisconnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_exchange_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectReply.ProtoReflect.Descriptor instead.
func (*DisconnectReply) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_exchange_proto_rawDescGZIP(), []int{6}
}

type KeepAliveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	SID      string `protobuf:"bytes,2,opt,name=SID,proto3" json:"SID,omitempty"`
	UID      string `protobuf:"bytes,3,opt,name=UID,proto3" json:"UID,omitempty"`
	Platform string `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
}

func (x *KeepAliveRequest) Reset() {
	*x = KeepAliveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_exchange_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepAliveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAliveRequest) ProtoMessage() {}

func (x *KeepAliveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_exchange_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepAliveRequest.ProtoReflect.Descriptor instead.
func (*KeepAliveRequest) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_exchange_proto_rawDescGZIP(), []int{7}
}

func (x *KeepAliveRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *KeepAliveRequest) GetSID() string {
	if x != nil {
		return x.SID
	}
	return ""
}

func (x *KeepAliveRequest) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *KeepAliveRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

type KeepAliveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *KeepAliveReply) Reset() {
	*x = KeepAliveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_exchange_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepAliveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAliveReply) ProtoMessage() {}

func (x *KeepAliveReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_exchange_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepAliveReply.ProtoReflect.Descriptor instead.
func (*KeepAliveReply) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_exchange_proto_rawDescGZIP(), []int{8}
}

func (x *KeepAliveReply) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_api_msg_api_msg_exchange_proto protoreflect.FileDescriptor

var file_api_msg_api_msg_exchange_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73,
	0x67, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6d, 0x73, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x53, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x14,
	0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x68, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x53, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x49, 0x44,
	0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x0e,
	0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x4f,
	0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x53,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x49, 0x44, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x49, 0x44, 0x22,
	0x11, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x6a, 0x0a, 0x10, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x53, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x49, 0x44,
	0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x28,
	0x0a, 0x0e, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0x84, 0x02, 0x0a, 0x12, 0x4d, 0x73, 0x67,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x0c, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x73, 0x67, 0x1a, 0x15, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x39, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x42, 0x0a, 0x0a, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3f,
	0x0a, 0x09, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x7a, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_msg_api_msg_exchange_proto_rawDescOnce sync.Once
	file_api_msg_api_msg_exchange_proto_rawDescData = file_api_msg_api_msg_exchange_proto_rawDesc
)

func file_api_msg_api_msg_exchange_proto_rawDescGZIP() []byte {
	file_api_msg_api_msg_exchange_proto_rawDescOnce.Do(func() {
		file_api_msg_api_msg_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_msg_api_msg_exchange_proto_rawDescData)
	})
	return file_api_msg_api_msg_exchange_proto_rawDescData
}

var file_api_msg_api_msg_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_msg_api_msg_exchange_proto_goTypes = []interface{}{
	(*Msg)(nil),                // 0: msg.api.Msg
	(*UserConnectionInfo)(nil), // 1: msg.api.UserConnectionInfo
	(*SendMsgReply)(nil),       // 2: msg.api.SendMsgReply
	(*ConnectRequest)(nil),     // 3: msg.api.ConnectRequest
	(*ConnectReply)(nil),       // 4: msg.api.ConnectReply
	(*DisconnectRequest)(nil),  // 5: msg.api.DisconnectRequest
	(*DisconnectReply)(nil),    // 6: msg.api.DisconnectReply
	(*KeepAliveRequest)(nil),   // 7: msg.api.KeepAliveRequest
	(*KeepAliveReply)(nil),     // 8: msg.api.KeepAliveReply
	(protocol.Platform)(0),     // 9: msg.protocol.Platform
	(*protocol.Protocol)(nil),  // 10: msg.protocol.Protocol
}
var file_api_msg_api_msg_exchange_proto_depIdxs = []int32{
	9,  // 0: msg.api.Msg.platform:type_name -> msg.protocol.Platform
	10, // 1: msg.api.Msg.protocol:type_name -> msg.protocol.Protocol
	0,  // 2: msg.api.MsgExchangeService.SendMsg:input_type -> msg.api.Msg
	3,  // 3: msg.api.MsgExchangeService.Connect:input_type -> msg.api.ConnectRequest
	5,  // 4: msg.api.MsgExchangeService.Disconnect:input_type -> msg.api.DisconnectRequest
	7,  // 5: msg.api.MsgExchangeService.KeepAlive:input_type -> msg.api.KeepAliveRequest
	2,  // 6: msg.api.MsgExchangeService.SendMsg:output_type -> msg.api.SendMsgReply
	4,  // 7: msg.api.MsgExchangeService.Connect:output_type -> msg.api.ConnectReply
	6,  // 8: msg.api.MsgExchangeService.Disconnect:output_type -> msg.api.DisconnectReply
	8,  // 9: msg.api.MsgExchangeService.KeepAlive:output_type -> msg.api.KeepAliveReply
	6,  // [6:10] is the sub-list for method output_type
	2,  // [2:6] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_msg_api_msg_exchange_proto_init() }
func file_api_msg_api_msg_exchange_proto_init() {
	if File_api_msg_api_msg_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_msg_api_msg_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_api_msg_api_msg_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserConnectionInfo); i {
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
		file_api_msg_api_msg_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgReply); i {
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
		file_api_msg_api_msg_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_api_msg_api_msg_exchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectReply); i {
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
		file_api_msg_api_msg_exchange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
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
		file_api_msg_api_msg_exchange_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectReply); i {
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
		file_api_msg_api_msg_exchange_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepAliveRequest); i {
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
		file_api_msg_api_msg_exchange_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepAliveReply); i {
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
			RawDescriptor: file_api_msg_api_msg_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_msg_api_msg_exchange_proto_goTypes,
		DependencyIndexes: file_api_msg_api_msg_exchange_proto_depIdxs,
		MessageInfos:      file_api_msg_api_msg_exchange_proto_msgTypes,
	}.Build()
	File_api_msg_api_msg_exchange_proto = out.File
	file_api_msg_api_msg_exchange_proto_rawDesc = nil
	file_api_msg_api_msg_exchange_proto_goTypes = nil
	file_api_msg_api_msg_exchange_proto_depIdxs = nil
}
