// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.5
// source: api/msg_gateway/msg_gateway.proto

package msg_gateway

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MsgGatewayServiceClient is the client API for MsgGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgGatewayServiceClient interface {
	PushMsg(ctx context.Context, in *PushMsgReq, opts ...grpc.CallOption) (*PushMsgReply, error)
}

type msgGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgGatewayServiceClient(cc grpc.ClientConnInterface) MsgGatewayServiceClient {
	return &msgGatewayServiceClient{cc}
}

func (c *msgGatewayServiceClient) PushMsg(ctx context.Context, in *PushMsgReq, opts ...grpc.CallOption) (*PushMsgReply, error) {
	out := new(PushMsgReply)
	err := c.cc.Invoke(ctx, "/api.msg.gateway.MsgGatewayService/PushMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgGatewayServiceServer is the server API for MsgGatewayService service.
// All implementations must embed UnimplementedMsgGatewayServiceServer
// for forward compatibility
type MsgGatewayServiceServer interface {
	PushMsg(context.Context, *PushMsgReq) (*PushMsgReply, error)
	mustEmbedUnimplementedMsgGatewayServiceServer()
}

// UnimplementedMsgGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMsgGatewayServiceServer struct {
}

func (UnimplementedMsgGatewayServiceServer) PushMsg(context.Context, *PushMsgReq) (*PushMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMsg not implemented")
}
func (UnimplementedMsgGatewayServiceServer) mustEmbedUnimplementedMsgGatewayServiceServer() {}

// UnsafeMsgGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgGatewayServiceServer will
// result in compilation errors.
type UnsafeMsgGatewayServiceServer interface {
	mustEmbedUnimplementedMsgGatewayServiceServer()
}

func RegisterMsgGatewayServiceServer(s grpc.ServiceRegistrar, srv MsgGatewayServiceServer) {
	s.RegisterService(&MsgGatewayService_ServiceDesc, srv)
}

func _MsgGatewayService_PushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgGatewayServiceServer).PushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.msg.gateway.MsgGatewayService/PushMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgGatewayServiceServer).PushMsg(ctx, req.(*PushMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MsgGatewayService_ServiceDesc is the grpc.ServiceDesc for MsgGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MsgGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.msg.gateway.MsgGatewayService",
	HandlerType: (*MsgGatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMsg",
			Handler:    _MsgGatewayService_PushMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/msg_gateway/msg_gateway.proto",
}
