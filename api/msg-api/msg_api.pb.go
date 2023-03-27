// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.5
// source: api/msg-api/msg_api.proto

package msg_api

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

type BizMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizCode    string `protobuf:"bytes,1,opt,name=bizCode,proto3" json:"bizCode,omitempty"`
	BizDesc    string `protobuf:"bytes,2,opt,name=bizDesc,proto3" json:"bizDesc,omitempty"`
	BizPayload string `protobuf:"bytes,3,opt,name=bizPayload,proto3" json:"bizPayload,omitempty"`
	Uid        uint64 `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"` // 推送的用户标识
}

func (x *BizMsg) Reset() {
	*x = BizMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BizMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BizMsg) ProtoMessage() {}

func (x *BizMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BizMsg.ProtoReflect.Descriptor instead.
func (*BizMsg) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_api_proto_rawDescGZIP(), []int{0}
}

func (x *BizMsg) GetBizCode() string {
	if x != nil {
		return x.BizCode
	}
	return ""
}

func (x *BizMsg) GetBizDesc() string {
	if x != nil {
		return x.BizDesc
	}
	return ""
}

func (x *BizMsg) GetBizPayload() string {
	if x != nil {
		return x.BizPayload
	}
	return ""
}

func (x *BizMsg) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type PushBizMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessCount int64 `protobuf:"varint,1,opt,name=successCount,proto3" json:"successCount,omitempty"` // 推送成功用户数量
}

func (x *PushBizMsgReply) Reset() {
	*x = PushBizMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_msg_api_msg_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushBizMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBizMsgReply) ProtoMessage() {}

func (x *PushBizMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_msg_api_msg_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBizMsgReply.ProtoReflect.Descriptor instead.
func (*PushBizMsgReply) Descriptor() ([]byte, []int) {
	return file_api_msg_api_msg_api_proto_rawDescGZIP(), []int{1}
}

func (x *PushBizMsgReply) GetSuccessCount() int64 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

var File_api_msg_api_msg_api_proto protoreflect.FileDescriptor

var file_api_msg_api_msg_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73,
	0x67, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x73, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x22, 0x6e, 0x0a, 0x06, 0x42, 0x69, 0x7a, 0x4d, 0x73, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x69, 0x7a, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x69, 0x7a, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x69, 0x7a, 0x44,
	0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x69, 0x7a, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x69, 0x7a, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x69, 0x7a, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x42, 0x69, 0x7a, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x44, 0x0a, 0x0d, 0x4d,
	0x73, 0x67, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04,
	0x50, 0x75, 0x73, 0x68, 0x12, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42,
	0x69, 0x7a, 0x4d, 0x73, 0x67, 0x1a, 0x18, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x42, 0x69, 0x7a, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x7a, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6c, 0x69,
	0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x73, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_msg_api_msg_api_proto_rawDescOnce sync.Once
	file_api_msg_api_msg_api_proto_rawDescData = file_api_msg_api_msg_api_proto_rawDesc
)

func file_api_msg_api_msg_api_proto_rawDescGZIP() []byte {
	file_api_msg_api_msg_api_proto_rawDescOnce.Do(func() {
		file_api_msg_api_msg_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_msg_api_msg_api_proto_rawDescData)
	})
	return file_api_msg_api_msg_api_proto_rawDescData
}

var file_api_msg_api_msg_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_msg_api_msg_api_proto_goTypes = []interface{}{
	(*BizMsg)(nil),          // 0: msg.api.BizMsg
	(*PushBizMsgReply)(nil), // 1: msg.api.PushBizMsgReply
}
var file_api_msg_api_msg_api_proto_depIdxs = []int32{
	0, // 0: msg.api.MsgAPIService.Push:input_type -> msg.api.BizMsg
	1, // 1: msg.api.MsgAPIService.Push:output_type -> msg.api.PushBizMsgReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_msg_api_msg_api_proto_init() }
func file_api_msg_api_msg_api_proto_init() {
	if File_api_msg_api_msg_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_msg_api_msg_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BizMsg); i {
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
		file_api_msg_api_msg_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushBizMsgReply); i {
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
			RawDescriptor: file_api_msg_api_msg_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_msg_api_msg_api_proto_goTypes,
		DependencyIndexes: file_api_msg_api_msg_api_proto_depIdxs,
		MessageInfos:      file_api_msg_api_msg_api_proto_msgTypes,
	}.Build()
	File_api_msg_api_msg_api_proto = out.File
	file_api_msg_api_msg_api_proto_rawDesc = nil
	file_api_msg_api_msg_api_proto_goTypes = nil
	file_api_msg_api_msg_api_proto_depIdxs = nil
}