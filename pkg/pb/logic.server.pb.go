// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logic.server.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("logic.server.proto", fileDescriptor_61926e2a885cc182) }

var fileDescriptor_61926e2a885cc182 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xc9, 0x4f, 0xcf,
	0x4c, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2a, 0x48, 0x92, 0x82, 0x8a, 0x27, 0xe7, 0x64, 0xa6, 0xe6, 0x95, 0x40, 0xc4, 0x8d, 0x3c, 0xb8,
	0xf8, 0x7c, 0x40, 0xa2, 0xc1, 0x60, 0xc5, 0xae, 0x15, 0x25, 0x42, 0x66, 0x5c, 0xdc, 0xc1, 0xa9,
	0x79, 0x29, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x42, 0x7a, 0x05, 0x49, 0x7a, 0x48,
	0x02, 0x41, 0xa9, 0x85, 0x52, 0xc2, 0x18, 0x62, 0xc5, 0x05, 0x49, 0x6c, 0x60, 0x03, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xa1, 0xfd, 0x1b, 0x7e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogicServerExtClient is the client API for LogicServerExt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicServerExtClient interface {
	// 发送消息
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
}

type logicServerExtClient struct {
	cc *grpc.ClientConn
}

func NewLogicServerExtClient(cc *grpc.ClientConn) LogicServerExtClient {
	return &logicServerExtClient{cc}
}

func (c *logicServerExtClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error) {
	out := new(SendMessageResp)
	err := c.cc.Invoke(ctx, "/pb.LogicServerExt/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicServerExtServer is the server API for LogicServerExt service.
type LogicServerExtServer interface {
	// 发送消息
	SendMessage(context.Context, *SendMessageReq) (*SendMessageResp, error)
}

// UnimplementedLogicServerExtServer can be embedded to have forward compatible implementations.
type UnimplementedLogicServerExtServer struct {
}

func (*UnimplementedLogicServerExtServer) SendMessage(ctx context.Context, req *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterLogicServerExtServer(s *grpc.Server, srv LogicServerExtServer) {
	s.RegisterService(&_LogicServerExt_serviceDesc, srv)
}

func _LogicServerExt_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServerExtServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicServerExt/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServerExtServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogicServerExt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LogicServerExt",
	HandlerType: (*LogicServerExtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _LogicServerExt_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic.server.proto",
}
