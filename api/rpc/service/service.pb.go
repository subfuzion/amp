// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/service/service.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/service/service.proto

It has these top-level messages:
	ServiceCreateRequest
	ServiceCreateResponse
	ServiceSpec
	PublishSpec
*/
package service

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

type ServiceCreateRequest struct {
	ServiceSpec *ServiceSpec `protobuf:"bytes,1,opt,name=service_spec,json=serviceSpec" json:"service_spec,omitempty"`
}

func (m *ServiceCreateRequest) Reset()                    { *m = ServiceCreateRequest{} }
func (m *ServiceCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceCreateRequest) ProtoMessage()               {}
func (*ServiceCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServiceCreateRequest) GetServiceSpec() *ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

type ServiceCreateResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ServiceCreateResponse) Reset()                    { *m = ServiceCreateResponse{} }
func (m *ServiceCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*ServiceCreateResponse) ProtoMessage()               {}
func (*ServiceCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ServiceSpec struct {
	Image        string            `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Name         string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Replicas     uint64            `protobuf:"varint,3,opt,name=replicas" json:"replicas,omitempty"`
	Env          map[string]string `protobuf:"bytes,4,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Labels       map[string]string `protobuf:"bytes,5,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	PublishSpecs []*PublishSpec    `protobuf:"bytes,6,rep,name=publish_specs,json=publishSpecs" json:"publish_specs,omitempty"`
}

func (m *ServiceSpec) Reset()                    { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string            { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()               {}
func (*ServiceSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServiceSpec) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *ServiceSpec) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ServiceSpec) GetPublishSpecs() []*PublishSpec {
	if m != nil {
		return m.PublishSpecs
	}
	return nil
}

type PublishSpec struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Protocol     string `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
	PublishPort  uint32 `protobuf:"varint,3,opt,name=publish_port,json=publishPort" json:"publish_port,omitempty"`
	InternalPort uint32 `protobuf:"varint,4,opt,name=internal_port,json=internalPort" json:"internal_port,omitempty"`
}

func (m *PublishSpec) Reset()                    { *m = PublishSpec{} }
func (m *PublishSpec) String() string            { return proto.CompactTextString(m) }
func (*PublishSpec) ProtoMessage()               {}
func (*PublishSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*ServiceCreateRequest)(nil), "service.ServiceCreateRequest")
	proto.RegisterType((*ServiceCreateResponse)(nil), "service.ServiceCreateResponse")
	proto.RegisterType((*ServiceSpec)(nil), "service.ServiceSpec")
	proto.RegisterType((*PublishSpec)(nil), "service.PublishSpec")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Service service

type ServiceClient interface {
	Create(ctx context.Context, in *ServiceCreateRequest, opts ...grpc.CallOption) (*ServiceCreateResponse, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Create(ctx context.Context, in *ServiceCreateRequest, opts ...grpc.CallOption) (*ServiceCreateResponse, error) {
	out := new(ServiceCreateResponse)
	err := grpc.Invoke(ctx, "/service.Service/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceServer interface {
	Create(context.Context, *ServiceCreateRequest) (*ServiceCreateResponse, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Service/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Create(ctx, req.(*ServiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Service_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/service/service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x25, 0xd9, 0xed, 0xb6, 0x4c, 0x76, 0x11, 0xb2, 0x16, 0x29, 0x8a, 0x04, 0x0a, 0xe1, 0xc0,
	0x9e, 0x12, 0x69, 0x91, 0xa0, 0xcb, 0x15, 0xf5, 0x80, 0x84, 0x44, 0x95, 0x72, 0xaf, 0x1c, 0x77,
	0xd4, 0x5a, 0x38, 0xb1, 0xb1, 0x9d, 0x48, 0xfd, 0x00, 0xfe, 0x96, 0x8f, 0x40, 0xb1, 0x9d, 0xdd,
	0x2d, 0x94, 0x03, 0xa7, 0xcc, 0x1b, 0xbf, 0x79, 0x7a, 0xf3, 0x32, 0xb0, 0xbb, 0xe5, 0xf6, 0xae,
	0x6f, 0x4a, 0x26, 0xdb, 0x8a, 0x2a, 0xc5, 0x50, 0xa0, 0xa6, 0x56, 0xea, 0x8a, 0xb6, 0xaa, 0xa2,
	0x8a, 0x57, 0x5a, 0xb1, 0xca, 0xa0, 0x1e, 0x38, 0xc3, 0xe9, 0x5b, 0x2a, 0x2d, 0xad, 0x24, 0xa7,
	0x01, 0x16, 0x5f, 0x61, 0x7d, 0xe5, 0xcb, 0x4f, 0x1a, 0xa9, 0xc5, 0x1a, 0x7f, 0xf4, 0x68, 0x2c,
	0xf9, 0x00, 0xcb, 0x40, 0xb9, 0x36, 0x0a, 0x59, 0x1a, 0xe5, 0xd1, 0x26, 0xd9, 0xae, 0xcb, 0x49,
	0x26, 0x0c, 0x5d, 0x29, 0x64, 0x75, 0x62, 0x0e, 0xa0, 0x78, 0x0b, 0x2f, 0xfe, 0x10, 0x34, 0x4a,
	0x76, 0x06, 0xc9, 0x33, 0x88, 0xf9, 0x8d, 0xd3, 0x79, 0x5a, 0xc7, 0xfc, 0xa6, 0xf8, 0x15, 0x43,
	0x72, 0xa4, 0x42, 0xd6, 0x70, 0xc2, 0x5b, 0x7a, 0x8b, 0x81, 0xe2, 0x01, 0x21, 0x30, 0xef, 0x68,
	0x8b, 0x69, 0xec, 0x9a, 0xae, 0x26, 0x19, 0x9c, 0x69, 0x54, 0x82, 0x33, 0x6a, 0xd2, 0x59, 0x1e,
	0x6d, 0xe6, 0xf5, 0x1e, 0x93, 0x0a, 0x66, 0xd8, 0x0d, 0xe9, 0x3c, 0x9f, 0x6d, 0x92, 0xed, 0xcb,
	0xc7, 0xec, 0x96, 0x17, 0xdd, 0x70, 0xd1, 0x59, 0x7d, 0x5f, 0x8f, 0x4c, 0x72, 0x0e, 0x0b, 0x41,
	0x1b, 0x14, 0x26, 0x3d, 0x71, 0x33, 0xf9, 0xa3, 0x33, 0x5f, 0x1c, 0xc5, 0x8f, 0x05, 0x3e, 0xd9,
	0xc1, 0x4a, 0xf5, 0x8d, 0xe0, 0xe6, 0xce, 0x45, 0x64, 0xd2, 0x85, 0x13, 0x38, 0x64, 0x74, 0xe9,
	0x5f, 0x5d, 0x46, 0x4b, 0x75, 0x00, 0x26, 0x7b, 0x0f, 0x67, 0x93, 0x0b, 0xf2, 0x1c, 0x66, 0xdf,
	0xf1, 0x3e, 0x6c, 0x3d, 0x96, 0x63, 0x12, 0x03, 0x15, 0xfd, 0xb4, 0xb4, 0x07, 0x1f, 0xe3, 0xf3,
	0x28, 0xdb, 0x41, 0x72, 0xe4, 0xe4, 0x7f, 0x46, 0x8b, 0x9f, 0x11, 0x24, 0x47, 0x86, 0xf6, 0xc1,
	0x46, 0x0f, 0x83, 0x75, 0xe7, 0xc1, 0xa4, 0x08, 0x02, 0x7b, 0x4c, 0x5e, 0xc3, 0xb4, 0xc2, 0xb5,
	0x92, 0xda, 0xba, 0xe0, 0x57, 0x75, 0x12, 0x7a, 0x97, 0x52, 0x5b, 0xf2, 0x06, 0x56, 0xbc, 0xb3,
	0xa8, 0x3b, 0x2a, 0x3c, 0x67, 0xee, 0x38, 0xcb, 0xa9, 0x39, 0x92, 0xb6, 0xdf, 0xe0, 0x34, 0x04,
	0x4b, 0x3e, 0xc3, 0xc2, 0xdf, 0x08, 0xf9, 0xeb, 0x47, 0x3d, 0x38, 0xc6, 0xec, 0xd5, 0xbf, 0x9e,
	0xfd, 0x69, 0x15, 0x4f, 0x9a, 0x85, 0xf3, 0xf9, 0xee, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac,
	0x5d, 0x00, 0x21, 0x13, 0x03, 0x00, 0x00,
}