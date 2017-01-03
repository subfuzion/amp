// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/function/function.proto
// DO NOT EDIT!

/*
Package function is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/function/function.proto

It has these top-level messages:
	FunctionEntry
	FunctionCall
	FunctionReturn
	CreateRequest
	CreateReply
	ListRequest
	ListReply
	DeleteRequest
	DeleteReply
*/
package function

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

type FunctionEntry struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Image string `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
}

func (m *FunctionEntry) Reset()                    { *m = FunctionEntry{} }
func (m *FunctionEntry) String() string            { return proto.CompactTextString(m) }
func (*FunctionEntry) ProtoMessage()               {}
func (*FunctionEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FunctionEntry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FunctionEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FunctionEntry) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type FunctionCall struct {
	CallID   string         `protobuf:"bytes,1,opt,name=callID" json:"callID,omitempty"`
	Input    []byte         `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	Function *FunctionEntry `protobuf:"bytes,3,opt,name=function" json:"function,omitempty"`
	ReturnTo string         `protobuf:"bytes,4,opt,name=returnTo" json:"returnTo,omitempty"`
}

func (m *FunctionCall) Reset()                    { *m = FunctionCall{} }
func (m *FunctionCall) String() string            { return proto.CompactTextString(m) }
func (*FunctionCall) ProtoMessage()               {}
func (*FunctionCall) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FunctionCall) GetCallID() string {
	if m != nil {
		return m.CallID
	}
	return ""
}

func (m *FunctionCall) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *FunctionCall) GetFunction() *FunctionEntry {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *FunctionCall) GetReturnTo() string {
	if m != nil {
		return m.ReturnTo
	}
	return ""
}

type FunctionReturn struct {
	CallID string `protobuf:"bytes,1,opt,name=callID" json:"callID,omitempty"`
	Output []byte `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (m *FunctionReturn) Reset()                    { *m = FunctionReturn{} }
func (m *FunctionReturn) String() string            { return proto.CompactTextString(m) }
func (*FunctionReturn) ProtoMessage()               {}
func (*FunctionReturn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FunctionReturn) GetCallID() string {
	if m != nil {
		return m.CallID
	}
	return ""
}

func (m *FunctionReturn) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

type CreateRequest struct {
	Function *FunctionEntry `protobuf:"bytes,1,opt,name=function" json:"function,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateRequest) GetFunction() *FunctionEntry {
	if m != nil {
		return m.Function
	}
	return nil
}

type CreateReply struct {
	Function *FunctionEntry `protobuf:"bytes,1,opt,name=function" json:"function,omitempty"`
}

func (m *CreateReply) Reset()                    { *m = CreateReply{} }
func (m *CreateReply) String() string            { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()               {}
func (*CreateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateReply) GetFunction() *FunctionEntry {
	if m != nil {
		return m.Function
	}
	return nil
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListReply struct {
	Functions []*FunctionEntry `protobuf:"bytes,1,rep,name=functions" json:"functions,omitempty"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListReply) GetFunctions() []*FunctionEntry {
	if m != nil {
		return m.Functions
	}
	return nil
}

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteReply struct {
}

func (m *DeleteReply) Reset()                    { *m = DeleteReply{} }
func (m *DeleteReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteReply) ProtoMessage()               {}
func (*DeleteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*FunctionEntry)(nil), "function.FunctionEntry")
	proto.RegisterType((*FunctionCall)(nil), "function.FunctionCall")
	proto.RegisterType((*FunctionReturn)(nil), "function.FunctionReturn")
	proto.RegisterType((*CreateRequest)(nil), "function.CreateRequest")
	proto.RegisterType((*CreateReply)(nil), "function.CreateReply")
	proto.RegisterType((*ListRequest)(nil), "function.ListRequest")
	proto.RegisterType((*ListReply)(nil), "function.ListReply")
	proto.RegisterType((*DeleteRequest)(nil), "function.DeleteRequest")
	proto.RegisterType((*DeleteReply)(nil), "function.DeleteReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Function service

type FunctionClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error)
}

type functionClient struct {
	cc *grpc.ClientConn
}

func NewFunctionClient(cc *grpc.ClientConn) FunctionClient {
	return &functionClient{cc}
}

func (c *functionClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := grpc.Invoke(ctx, "/function.Function/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/function.Function/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteReply, error) {
	out := new(DeleteReply)
	err := grpc.Invoke(ctx, "/function.Function/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Function service

type FunctionServer interface {
	Create(context.Context, *CreateRequest) (*CreateReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Delete(context.Context, *DeleteRequest) (*DeleteReply, error)
}

func RegisterFunctionServer(s *grpc.Server, srv FunctionServer) {
	s.RegisterService(&_Function_serviceDesc, srv)
}

func _Function_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/function.Function/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Function_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/function.Function/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Function_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/function.Function/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Function_serviceDesc = grpc.ServiceDesc{
	ServiceName: "function.Function",
	HandlerType: (*FunctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Function_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Function_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Function_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/function/function.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/function/function.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x95, 0x4c, 0x89, 0xda, 0xdb, 0xa6, 0xd2, 0x78, 0x86, 0x99, 0x10, 0x21, 0x31, 0xf2,
	0x0a, 0xcd, 0xa2, 0x11, 0xad, 0xd8, 0xc0, 0x06, 0xb5, 0x05, 0x51, 0x89, 0x0d, 0x81, 0x17, 0x70,
	0x53, 0x13, 0x2c, 0x39, 0xb6, 0x71, 0x1c, 0xa4, 0x08, 0xb1, 0x61, 0xcf, 0x8a, 0x47, 0xe3, 0x15,
	0x78, 0x0c, 0x16, 0x28, 0xce, 0x4f, 0x93, 0x22, 0x10, 0x62, 0xe7, 0x73, 0x7b, 0xee, 0xe7, 0x73,
	0x9a, 0x04, 0x9e, 0xa6, 0xcc, 0xbc, 0x2f, 0xf6, 0x8b, 0x44, 0x66, 0x11, 0x51, 0x2a, 0xa1, 0x9c,
	0x6a, 0x62, 0xa4, 0x8e, 0x48, 0xa6, 0x22, 0xa2, 0x58, 0xa4, 0x55, 0x12, 0xbd, 0x2b, 0x44, 0x62,
	0x98, 0x14, 0xdd, 0x61, 0xa1, 0xb4, 0x34, 0x12, 0x8d, 0x5b, 0x1d, 0xde, 0x4f, 0xa5, 0x4c, 0x39,
	0xb5, 0x1b, 0x44, 0x08, 0x69, 0x48, 0x35, 0xce, 0x6b, 0x1f, 0xde, 0x81, 0xff, 0xa2, 0x71, 0x3e,
	0x17, 0x46, 0x97, 0x68, 0x0e, 0x2e, 0x3b, 0x04, 0xce, 0x8d, 0xf3, 0x70, 0x12, 0xbb, 0xec, 0x80,
	0x10, 0x8c, 0x04, 0xc9, 0x68, 0xe0, 0xda, 0x89, 0x3d, 0xa3, 0x4b, 0xb8, 0xc3, 0x32, 0x92, 0xd2,
	0xe0, 0xcc, 0x0e, 0x6b, 0x81, 0xbf, 0x3a, 0x30, 0x6b, 0x59, 0x1b, 0xc2, 0x39, 0xba, 0x02, 0x2f,
	0x21, 0x9c, 0xef, 0xb6, 0x0d, 0xae, 0x51, 0x76, 0x5d, 0xa8, 0xc2, 0x58, 0xe6, 0x2c, 0xae, 0x05,
	0x5a, 0x41, 0x97, 0xd9, 0x72, 0xa7, 0xcb, 0xeb, 0x45, 0x57, 0x6a, 0x90, 0x31, 0xee, 0x8c, 0x28,
	0x84, 0xb1, 0xa6, 0xa6, 0xd0, 0xe2, 0xad, 0x0c, 0x46, 0xf6, 0x92, 0x4e, 0xe3, 0x67, 0x30, 0x6f,
	0xd7, 0x62, 0x3b, 0xfb, 0x63, 0xa0, 0x2b, 0xf0, 0x64, 0x61, 0x8e, 0x89, 0x1a, 0x85, 0xb7, 0xe0,
	0x6f, 0x34, 0x25, 0x86, 0xc6, 0xf4, 0x43, 0x41, 0xf3, 0x61, 0x46, 0xe7, 0x1f, 0x33, 0xe2, 0x35,
	0x4c, 0x5b, 0x8a, 0xe2, 0xe5, 0xff, 0x31, 0x7c, 0x98, 0xbe, 0x62, 0xb9, 0x69, 0x72, 0xe0, 0x35,
	0x4c, 0x6a, 0x59, 0x01, 0x1f, 0xc3, 0xa4, 0xf5, 0xe5, 0x81, 0x73, 0x73, 0xf6, 0x37, 0xe2, 0xd1,
	0x89, 0x1f, 0x80, 0xbf, 0xa5, 0x9c, 0x1e, 0xcb, 0x9d, 0x3c, 0xf9, 0xea, 0xce, 0xd6, 0xa0, 0x78,
	0xb9, 0xfc, 0xe9, 0xc0, 0xb8, 0x85, 0xa1, 0xd7, 0xe0, 0xd5, 0x9d, 0x50, 0xef, 0xaa, 0xc1, 0x7f,
	0x15, 0xde, 0xfd, 0xfd, 0x07, 0xc5, 0x4b, 0x7c, 0xfd, 0xe5, 0xfb, 0x8f, 0x6f, 0xee, 0x39, 0x9e,
	0x45, 0x1f, 0x1f, 0x75, 0x2f, 0xed, 0x13, 0xe7, 0x16, 0xbd, 0x84, 0x51, 0xd5, 0x09, 0xf5, 0xf6,
	0x7a, 0x95, 0xc3, 0x8b, 0xd3, 0x71, 0x05, 0xbb, 0xb4, 0xb0, 0x39, 0x1a, 0xc0, 0xd0, 0x1b, 0xf0,
	0xea, 0xe0, 0xfd, 0x70, 0x83, 0xae, 0xfd, 0x70, 0xbd, 0x8e, 0xf8, 0x9e, 0xe5, 0x5d, 0xdc, 0x9e,
	0xf7, 0x79, 0xd1, 0x27, 0x76, 0xf8, 0xbc, 0xf7, 0xec, 0xf7, 0xb2, 0xfa, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x72, 0xa8, 0x3e, 0x90, 0x96, 0x03, 0x00, 0x00,
}