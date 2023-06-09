// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.5
// source: api/msg_gateway/msg_gateway.proto

package msg_gateway

import (
	protocol "github.com/openlinkz/openlink/api/protocol"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PushMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SID      string             `protobuf:"bytes,1,opt,name=SID,proto3" json:"SID,omitempty"`
	Protocol *protocol.Protocol `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *PushMsgReq) Reset() {
	*x = PushMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_gateway_msg_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReq) ProtoMessage() {}

func (x *PushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_gateway_msg_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReq.ProtoReflect.Descriptor instead.
func (*PushMsgReq) Descriptor() ([]byte, []int) {
	return file_api_msg_gateway_msg_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *PushMsgReq) GetSID() string {
	if x != nil {
		return x.SID
	}
	return ""
}

func (x *PushMsgReq) GetProtocol() *protocol.Protocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

type PushBatchMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SIDs     []string           `protobuf:"bytes,1,rep,name=SIDs,proto3" json:"SIDs,omitempty"`
	Protocol *protocol.Protocol `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *PushBatchMsgReq) Reset() {
	*x = PushBatchMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_gateway_msg_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushBatchMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBatchMsgReq) ProtoMessage() {}

func (x *PushBatchMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_gateway_msg_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBatchMsgReq.ProtoReflect.Descriptor instead.
func (*PushBatchMsgReq) Descriptor() ([]byte, []int) {
	return file_api_msg_gateway_msg_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *PushBatchMsgReq) GetSIDs() []string {
	if x != nil {
		return x.SIDs
	}
	return nil
}

func (x *PushBatchMsgReq) GetProtocol() *protocol.Protocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

type PushMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PushMsgReply) Reset() {
	*x = PushMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_gateway_msg_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReply) ProtoMessage() {}

func (x *PushMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_gateway_msg_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReply.ProtoReflect.Descriptor instead.
func (*PushMsgReply) Descriptor() ([]byte, []int) {
	return file_api_msg_gateway_msg_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *PushMsgReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PushBatchMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessTotal int32 `protobuf:"varint,1,opt,name=successTotal,proto3" json:"successTotal,omitempty"`
	FailedTotal  int32 `protobuf:"varint,2,opt,name=failedTotal,proto3" json:"failedTotal,omitempty"`
}

func (x *PushBatchMsgReply) Reset() {
	*x = PushBatchMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_gateway_msg_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushBatchMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBatchMsgReply) ProtoMessage() {}

func (x *PushBatchMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_gateway_msg_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBatchMsgReply.ProtoReflect.Descriptor instead.
func (*PushBatchMsgReply) Descriptor() ([]byte, []int) {
	return file_api_msg_gateway_msg_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *PushBatchMsgReply) GetSuccessTotal() int32 {
	if x != nil {
		return x.SuccessTotal
	}
	return 0
}

func (x *PushBatchMsgReply) GetFailedTotal() int32 {
	if x != nil {
		return x.FailedTotal
	}
	return 0
}

var File_api_msg_gateway_msg_gateway_proto protoreflect.FileDescriptor

var file_api_msg_gateway_msg_gateway_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x56, 0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x53, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x49, 0x44, 0x12,
	0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x5d, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x49,
	0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x53, 0x49, 0x44, 0x73, 0x12, 0x36,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x59, 0x0a, 0x11, 0x50, 0x75, 0x73, 0x68, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xf2, 0x01, 0x0a, 0x11,
	0x4d, 0x73, 0x67, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x63, 0x0a, 0x07, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x22, 0x11, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x5f,
	0x6d, 0x73, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70,
	0x75, 0x73, 0x68, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x73, 0x67, 0x3a, 0x01, 0x2a,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x7a, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6c, 0x69, 0x6e,
	0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_msg_gateway_msg_gateway_proto_rawDescOnce sync.Once
	file_api_msg_gateway_msg_gateway_proto_rawDescData = file_api_msg_gateway_msg_gateway_proto_rawDesc
)

func file_api_msg_gateway_msg_gateway_proto_rawDescGZIP() []byte {
	file_api_msg_gateway_msg_gateway_proto_rawDescOnce.Do(func() {
		file_api_msg_gateway_msg_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_msg_gateway_msg_gateway_proto_rawDescData)
	})
	return file_api_msg_gateway_msg_gateway_proto_rawDescData
}

var file_api_msg_gateway_msg_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_msg_gateway_msg_gateway_proto_goTypes = []interface{}{
	(*PushMsgReq)(nil),        // 0: api.msg.gateway.PushMsgReq
	(*PushBatchMsgReq)(nil),   // 1: api.msg.gateway.PushBatchMsgReq
	(*PushMsgReply)(nil),      // 2: api.msg.gateway.PushMsgReply
	(*PushBatchMsgReply)(nil), // 3: api.msg.gateway.PushBatchMsgReply
	(*protocol.Protocol)(nil), // 4: api.msg.protocol.Protocol
}
var file_api_msg_gateway_msg_gateway_proto_depIdxs = []int32{
	4, // 0: api.msg.gateway.PushMsgReq.protocol:type_name -> api.msg.protocol.Protocol
	4, // 1: api.msg.gateway.PushBatchMsgReq.protocol:type_name -> api.msg.protocol.Protocol
	0, // 2: api.msg.gateway.MsgGatewayService.PushMsg:input_type -> api.msg.gateway.PushMsgReq
	1, // 3: api.msg.gateway.MsgGatewayService.PushBatchMsg:input_type -> api.msg.gateway.PushBatchMsgReq
	2, // 4: api.msg.gateway.MsgGatewayService.PushMsg:output_type -> api.msg.gateway.PushMsgReply
	3, // 5: api.msg.gateway.MsgGatewayService.PushBatchMsg:output_type -> api.msg.gateway.PushBatchMsgReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_msg_gateway_msg_gateway_proto_init() }
func file_api_msg_gateway_msg_gateway_proto_init() {
	if File_api_msg_gateway_msg_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_msg_gateway_msg_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReq); i {
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
		file_api_msg_gateway_msg_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushBatchMsgReq); i {
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
		file_api_msg_gateway_msg_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReply); i {
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
		file_api_msg_gateway_msg_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushBatchMsgReply); i {
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
			RawDescriptor: file_api_msg_gateway_msg_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_msg_gateway_msg_gateway_proto_goTypes,
		DependencyIndexes: file_api_msg_gateway_msg_gateway_proto_depIdxs,
		MessageInfos:      file_api_msg_gateway_msg_gateway_proto_msgTypes,
	}.Build()
	File_api_msg_gateway_msg_gateway_proto = out.File
	file_api_msg_gateway_msg_gateway_proto_rawDesc = nil
	file_api_msg_gateway_msg_gateway_proto_goTypes = nil
	file_api_msg_gateway_msg_gateway_proto_depIdxs = nil
}
