// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package otgrpc_testing is a generated protocol buffer package.

It is generated from these files:

	test.prot
	test.proto

e top-level messages:

	SimpleRe
It has these top-level messages:
	SimpleRequest
	SimpleResponse
*/
package otgrpc_testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SimpleRequest struct {
	Payload int32 `protobuf:"varint,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *SimpleRequest) Reset()                    { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string            { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()               {}
func (*SimpleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SimpleRequest) GetPayload() int32 {
	if m != nil {
		return m.Payload
	}
	return 0
}

type SimpleResponse struct {
	Payload int32 `protobuf:"varint,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *SimpleResponse) Reset()                    { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string            { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()               {}
func (*SimpleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SimpleResponse) GetPayload() int32 {
	if m != nil {
		return m.Payload
	}
	return 0
}

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "otgrpc.testing.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "otgrpc.testing.SimpleResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	StreamingOutputCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error)
	StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error)
	StreamingBidirectionalCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingBidirectionalCallClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/otgrpc.testing.TestService/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) StreamingOutputCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/otgrpc.testing.TestService/StreamingOutputCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingOutputCallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_StreamingOutputCallClient interface {
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type testServiceStreamingOutputCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingOutputCallClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[1], c.cc, "/otgrpc.testing.TestService/StreamingInputCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingInputCallClient{stream}
	return x, nil
}

type TestService_StreamingInputCallClient interface {
	Send(*SimpleRequest) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type testServiceStreamingInputCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingInputCallClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceStreamingInputCallClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) StreamingBidirectionalCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingBidirectionalCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[2], c.cc, "/otgrpc.testing.TestService/StreamingBidirectionalCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingBidirectionalCallClient{stream}
	return x, nil
}

type TestService_StreamingBidirectionalCallClient interface {
	Send(*SimpleRequest) error
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type testServiceStreamingBidirectionalCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingBidirectionalCallClient) Send(m *SimpleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceStreamingBidirectionalCallClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	StreamingOutputCall(*SimpleRequest, TestService_StreamingOutputCallServer) error
	StreamingInputCall(TestService_StreamingInputCallServer) error
	StreamingBidirectionalCall(TestService_StreamingBidirectionalCallServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otgrpc.testing.TestService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_StreamingOutputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).StreamingOutputCall(m, &testServiceStreamingOutputCallServer{stream})
}

type TestService_StreamingOutputCallServer interface {
	Send(*SimpleResponse) error
	grpc.ServerStream
}

type testServiceStreamingOutputCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingOutputCallServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_StreamingInputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).StreamingInputCall(&testServiceStreamingInputCallServer{stream})
}

type TestService_StreamingInputCallServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type testServiceStreamingInputCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingInputCallServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceStreamingInputCallServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_StreamingBidirectionalCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).StreamingBidirectionalCall(&testServiceStreamingBidirectionalCallServer{stream})
}

type TestService_StreamingBidirectionalCallServer interface {
	Send(*SimpleResponse) error
	Recv() (*SimpleRequest, error)
	grpc.ServerStream
}

type testServiceStreamingBidirectionalCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingBidirectionalCallServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceStreamingBidirectionalCallServer) Recv() (*SimpleRequest, error) {
	m := new(SimpleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "otgrpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingOutputCall",
			Handler:       _TestService_StreamingOutputCall_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamingInputCall",
			Handler:       _TestService_StreamingInputCall_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamingBidirectionalCall",
			Handler:       _TestService_StreamingBidirectionalCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0x2f, 0x49, 0x2f, 0x2a, 0x48, 0xd6, 0x03,
	0x09, 0x65, 0xe6, 0xa5, 0x2b, 0x69, 0x72, 0xf1, 0x06, 0x67, 0xe6, 0x16, 0xe4, 0xa4, 0x06, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0xb1, 0x17, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0xc1, 0xb8, 0x4a, 0x5a, 0x5c, 0x7c, 0x30, 0xa5, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0xb8, 0xd5, 0x1a, 0xbd, 0x64, 0xe2, 0xe2, 0x0e, 0x49, 0x2d, 0x2e,
	0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xf2, 0xe2, 0xe2, 0x0c, 0xcd, 0x4b, 0x2c, 0xaa,
	0x74, 0x4e, 0xcc, 0xc9, 0x11, 0x92, 0xd5, 0x43, 0x75, 0x84, 0x1e, 0x8a, 0x0b, 0xa4, 0xe4, 0x70,
	0x49, 0x43, 0x6d, 0x0d, 0xe3, 0x12, 0x0e, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xcd, 0xcc, 0x4b, 0xf7,
	0x2f, 0x2d, 0x29, 0x28, 0x2d, 0xa1, 0x82, 0xa9, 0x06, 0x8c, 0x42, 0xa1, 0x5c, 0x42, 0x70, 0x73,
	0x3d, 0xf3, 0xa8, 0x63, 0xac, 0x06, 0xa3, 0x50, 0x3c, 0x97, 0x14, 0xdc, 0x58, 0xa7, 0xcc, 0x94,
	0xcc, 0xa2, 0xd4, 0xe4, 0x92, 0xcc, 0xfc, 0xbc, 0xc4, 0x1c, 0xaa, 0x18, 0x6f, 0xc0, 0x98, 0xc4,
	0x06, 0x8e, 0x59, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x50, 0x3e, 0xe4, 0xe7, 0x01,
	0x00, 0x00,
}
