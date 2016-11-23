// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/cmd/swarm-server/servercore/swarm-server.proto
// DO NOT EDIT!

/*
Package servercore is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/cmd/swarm-server/servercore/swarm-server.proto

It has these top-level messages:
	AgentMes
	ClientMes
	DeclareRequest
	ServerRet
*/
package servercore

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

type AgentMes struct {
	Function string            `protobuf:"bytes,1,opt,name=function" json:"function,omitempty"`
	Args     map[string]string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AgentMes) Reset()                    { *m = AgentMes{} }
func (m *AgentMes) String() string            { return proto.CompactTextString(m) }
func (*AgentMes) ProtoMessage()               {}
func (*AgentMes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AgentMes) GetArgs() map[string]string {
	if m != nil {
		return m.Args
	}
	return nil
}

type ClientMes struct {
	Function string            `protobuf:"bytes,1,opt,name=function" json:"function,omitempty"`
	Args     map[string]string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ClientMes) Reset()                    { *m = ClientMes{} }
func (m *ClientMes) String() string            { return proto.CompactTextString(m) }
func (*ClientMes) ProtoMessage()               {}
func (*ClientMes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientMes) GetArgs() map[string]string {
	if m != nil {
		return m.Args
	}
	return nil
}

type DeclareRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeclareRequest) Reset()                    { *m = DeclareRequest{} }
func (m *DeclareRequest) String() string            { return proto.CompactTextString(m) }
func (*DeclareRequest) ProtoMessage()               {}
func (*DeclareRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ServerRet struct {
	Ack bool `protobuf:"varint,1,opt,name=ack" json:"ack,omitempty"`
}

func (m *ServerRet) Reset()                    { *m = ServerRet{} }
func (m *ServerRet) String() string            { return proto.CompactTextString(m) }
func (*ServerRet) ProtoMessage()               {}
func (*ServerRet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*AgentMes)(nil), "servercore.AgentMes")
	proto.RegisterType((*ClientMes)(nil), "servercore.ClientMes")
	proto.RegisterType((*DeclareRequest)(nil), "servercore.DeclareRequest")
	proto.RegisterType((*ServerRet)(nil), "servercore.ServerRet")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for SwarmServerService service

type SwarmServerServiceClient interface {
	GetAgentStream(ctx context.Context, opts ...grpc.CallOption) (SwarmServerService_GetAgentStreamClient, error)
	GetClientStream(ctx context.Context, opts ...grpc.CallOption) (SwarmServerService_GetClientStreamClient, error)
	DeclareAgent(ctx context.Context, in *DeclareRequest, opts ...grpc.CallOption) (*ServerRet, error)
}

type swarmServerServiceClient struct {
	cc *grpc.ClientConn
}

func NewSwarmServerServiceClient(cc *grpc.ClientConn) SwarmServerServiceClient {
	return &swarmServerServiceClient{cc}
}

func (c *swarmServerServiceClient) GetAgentStream(ctx context.Context, opts ...grpc.CallOption) (SwarmServerService_GetAgentStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SwarmServerService_serviceDesc.Streams[0], c.cc, "/servercore.SwarmServerService/GetAgentStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmServerServiceGetAgentStreamClient{stream}
	return x, nil
}

type SwarmServerService_GetAgentStreamClient interface {
	Send(*AgentMes) error
	Recv() (*AgentMes, error)
	grpc.ClientStream
}

type swarmServerServiceGetAgentStreamClient struct {
	grpc.ClientStream
}

func (x *swarmServerServiceGetAgentStreamClient) Send(m *AgentMes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *swarmServerServiceGetAgentStreamClient) Recv() (*AgentMes, error) {
	m := new(AgentMes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swarmServerServiceClient) GetClientStream(ctx context.Context, opts ...grpc.CallOption) (SwarmServerService_GetClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SwarmServerService_serviceDesc.Streams[1], c.cc, "/servercore.SwarmServerService/GetClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &swarmServerServiceGetClientStreamClient{stream}
	return x, nil
}

type SwarmServerService_GetClientStreamClient interface {
	Send(*ClientMes) error
	Recv() (*ClientMes, error)
	grpc.ClientStream
}

type swarmServerServiceGetClientStreamClient struct {
	grpc.ClientStream
}

func (x *swarmServerServiceGetClientStreamClient) Send(m *ClientMes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *swarmServerServiceGetClientStreamClient) Recv() (*ClientMes, error) {
	m := new(ClientMes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swarmServerServiceClient) DeclareAgent(ctx context.Context, in *DeclareRequest, opts ...grpc.CallOption) (*ServerRet, error) {
	out := new(ServerRet)
	err := grpc.Invoke(ctx, "/servercore.SwarmServerService/DeclareAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SwarmServerService service

type SwarmServerServiceServer interface {
	GetAgentStream(SwarmServerService_GetAgentStreamServer) error
	GetClientStream(SwarmServerService_GetClientStreamServer) error
	DeclareAgent(context.Context, *DeclareRequest) (*ServerRet, error)
}

func RegisterSwarmServerServiceServer(s *grpc.Server, srv SwarmServerServiceServer) {
	s.RegisterService(&_SwarmServerService_serviceDesc, srv)
}

func _SwarmServerService_GetAgentStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SwarmServerServiceServer).GetAgentStream(&swarmServerServiceGetAgentStreamServer{stream})
}

type SwarmServerService_GetAgentStreamServer interface {
	Send(*AgentMes) error
	Recv() (*AgentMes, error)
	grpc.ServerStream
}

type swarmServerServiceGetAgentStreamServer struct {
	grpc.ServerStream
}

func (x *swarmServerServiceGetAgentStreamServer) Send(m *AgentMes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *swarmServerServiceGetAgentStreamServer) Recv() (*AgentMes, error) {
	m := new(AgentMes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SwarmServerService_GetClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SwarmServerServiceServer).GetClientStream(&swarmServerServiceGetClientStreamServer{stream})
}

type SwarmServerService_GetClientStreamServer interface {
	Send(*ClientMes) error
	Recv() (*ClientMes, error)
	grpc.ServerStream
}

type swarmServerServiceGetClientStreamServer struct {
	grpc.ServerStream
}

func (x *swarmServerServiceGetClientStreamServer) Send(m *ClientMes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *swarmServerServiceGetClientStreamServer) Recv() (*ClientMes, error) {
	m := new(ClientMes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SwarmServerService_DeclareAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwarmServerServiceServer).DeclareAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servercore.SwarmServerService/DeclareAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwarmServerServiceServer).DeclareAgent(ctx, req.(*DeclareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SwarmServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "servercore.SwarmServerService",
	HandlerType: (*SwarmServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeclareAgent",
			Handler:    _SwarmServerService_DeclareAgent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAgentStream",
			Handler:       _SwarmServerService_GetAgentStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetClientStream",
			Handler:       _SwarmServerService_GetClientStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/cmd/swarm-server/servercore/swarm-server.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xbb, 0x6d, 0x95, 0x66, 0x94, 0x2a, 0x4b, 0x85, 0x12, 0x50, 0x4b, 0xf0, 0xd0, 0x8b,
	0x89, 0xb4, 0x07, 0xc5, 0x5b, 0xad, 0xa5, 0x20, 0x78, 0x49, 0x9e, 0x60, 0xbb, 0x8e, 0x31, 0x34,
	0xff, 0xdc, 0xdd, 0x54, 0xfa, 0x1c, 0xe2, 0x1b, 0xfa, 0x20, 0x92, 0x4d, 0x4c, 0x1b, 0xc8, 0xc1,
	0x83, 0x97, 0x64, 0xf6, 0xdb, 0x99, 0x1f, 0xdf, 0x7c, 0x09, 0x3c, 0xf9, 0x81, 0x7a, 0xcb, 0x56,
	0x36, 0x4f, 0x22, 0x87, 0xa5, 0x29, 0xc7, 0x10, 0x05, 0x53, 0x89, 0x70, 0x58, 0x94, 0x3a, 0x3c,
	0x7a, 0x71, 0xe4, 0x07, 0x13, 0xd1, 0xb5, 0x44, 0xb1, 0x41, 0xe1, 0x14, 0x2f, 0x9e, 0x08, 0xac,
	0xe9, 0x76, 0x2a, 0x12, 0x95, 0x50, 0xd8, 0x5d, 0x5b, 0x9f, 0x04, 0x7a, 0x33, 0x1f, 0x63, 0xf5,
	0x8c, 0x92, 0x9a, 0xd0, 0x7b, 0xcd, 0x62, 0xae, 0x82, 0x24, 0x1e, 0x92, 0x11, 0x19, 0x1b, 0x6e,
	0x75, 0xa6, 0x13, 0xe8, 0x32, 0xe1, 0xcb, 0x61, 0x7b, 0xd4, 0x19, 0x1f, 0x4d, 0x2e, 0xec, 0x1d,
	0xc3, 0xfe, 0x9d, 0xb7, 0x67, 0xc2, 0x97, 0x8b, 0x58, 0x89, 0xad, 0xab, 0x7b, 0xcd, 0x5b, 0x30,
	0x2a, 0x89, 0x9e, 0x42, 0x67, 0x8d, 0xdb, 0x92, 0x9b, 0x97, 0x74, 0x00, 0x07, 0x1b, 0x16, 0x66,
	0x38, 0x6c, 0x6b, 0xad, 0x38, 0xdc, 0xb7, 0xef, 0x88, 0xf5, 0x45, 0xc0, 0x98, 0x87, 0xc1, 0x1f,
	0x6c, 0x4d, 0x6b, 0xb6, 0x2e, 0xf7, 0x6d, 0x55, 0x80, 0xff, 0xf3, 0x75, 0x05, 0xfd, 0x47, 0xe4,
	0x21, 0x13, 0xe8, 0xe2, 0x7b, 0x86, 0x52, 0x51, 0x0a, 0xdd, 0x98, 0x45, 0x58, 0x8e, 0xeb, 0xda,
	0x3a, 0x07, 0xc3, 0xd3, 0x36, 0x5c, 0x54, 0x39, 0x9e, 0xf1, 0xb5, 0xbe, 0xef, 0xb9, 0x79, 0x39,
	0xf9, 0x26, 0x40, 0xbd, 0xfc, 0xab, 0x14, 0x4d, 0xf9, 0x33, 0xe0, 0x48, 0x1f, 0xa0, 0xbf, 0x44,
	0xa5, 0xb3, 0xf4, 0x94, 0x40, 0x16, 0xd1, 0x41, 0x53, 0xc8, 0x66, 0xa3, 0x6a, 0xb5, 0xc6, 0xe4,
	0x86, 0xd0, 0x05, 0x9c, 0x2c, 0x51, 0x15, 0x8b, 0x97, 0x90, 0xb3, 0xc6, 0x48, 0xcc, 0x66, 0xb9,
	0xc4, 0xcc, 0xe1, 0xb8, 0x5c, 0x53, 0xf3, 0xa9, 0xb9, 0xdf, 0x5c, 0x0f, 0xa0, 0x0e, 0xaa, 0xd6,
	0xb6, 0x5a, 0xab, 0x43, 0xfd, 0xb3, 0x4d, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x34, 0x3f,
	0x42, 0xba, 0x02, 0x00, 0x00,
}
