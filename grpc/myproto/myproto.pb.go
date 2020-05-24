// Code generated by protoc-gen-go. DO NOT EDIT.
// source: myproto.proto

package myproto

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

//客户端发送给服务端
type HelloReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{0}
}

func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReq.Unmarshal(m, b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return xxx_messageInfo_HelloReq.Size(m)
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//服务端返回给客户端
type HelloRsp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRsp) Reset()         { *m = HelloRsp{} }
func (m *HelloRsp) String() string { return proto.CompactTextString(m) }
func (*HelloRsp) ProtoMessage()    {}
func (*HelloRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{1}
}

func (m *HelloRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRsp.Unmarshal(m, b)
}
func (m *HelloRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRsp.Marshal(b, m, deterministic)
}
func (m *HelloRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRsp.Merge(m, src)
}
func (m *HelloRsp) XXX_Size() int {
	return xxx_messageInfo_HelloRsp.Size(m)
}
func (m *HelloRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRsp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRsp proto.InternalMessageInfo

func (m *HelloRsp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

//客户端发送给服务端
type NameReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameReq) Reset()         { *m = NameReq{} }
func (m *NameReq) String() string { return proto.CompactTextString(m) }
func (*NameReq) ProtoMessage()    {}
func (*NameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{2}
}

func (m *NameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameReq.Unmarshal(m, b)
}
func (m *NameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameReq.Marshal(b, m, deterministic)
}
func (m *NameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameReq.Merge(m, src)
}
func (m *NameReq) XXX_Size() int {
	return xxx_messageInfo_NameReq.Size(m)
}
func (m *NameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NameReq.DiscardUnknown(m)
}

var xxx_messageInfo_NameReq proto.InternalMessageInfo

func (m *NameReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//服务端返回给客户端
type NameRsp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameRsp) Reset()         { *m = NameRsp{} }
func (m *NameRsp) String() string { return proto.CompactTextString(m) }
func (*NameRsp) ProtoMessage()    {}
func (*NameRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{3}
}

func (m *NameRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameRsp.Unmarshal(m, b)
}
func (m *NameRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameRsp.Marshal(b, m, deterministic)
}
func (m *NameRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRsp.Merge(m, src)
}
func (m *NameRsp) XXX_Size() int {
	return xxx_messageInfo_NameRsp.Size(m)
}
func (m *NameRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRsp.DiscardUnknown(m)
}

var xxx_messageInfo_NameRsp proto.InternalMessageInfo

func (m *NameRsp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "myproto.HelloReq")
	proto.RegisterType((*HelloRsp)(nil), "myproto.HelloRsp")
	proto.RegisterType((*NameReq)(nil), "myproto.NameReq")
	proto.RegisterType((*NameRsp)(nil), "myproto.NameRsp")
}

func init() { proto.RegisterFile("myproto.proto", fileDescriptor_04877df457807402) }

var fileDescriptor_04877df457807402 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xad, 0x2c, 0x28,
	0xca, 0x2f, 0xc9, 0xd7, 0x03, 0x93, 0x42, 0xec, 0x50, 0xae, 0x92, 0x1c, 0x17, 0x87, 0x47, 0x6a,
	0x4e, 0x4e, 0x7e, 0x50, 0x6a, 0xa1, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0x24, 0x03, 0x93, 0x2f, 0x2e, 0x10, 0x12, 0xe0, 0x62,
	0xce, 0x2d, 0x4e, 0x87, 0x4a, 0x83, 0x98, 0x4a, 0xb2, 0x5c, 0xec, 0x7e, 0x89, 0xb9, 0xa9, 0xb8,
	0x34, 0x4b, 0x43, 0xa5, 0xb1, 0xe9, 0x35, 0x2a, 0xe2, 0xe2, 0x06, 0x9b, 0x5c, 0x9c, 0x5a, 0x54,
	0x96, 0x5a, 0x24, 0x64, 0xc4, 0xc5, 0x11, 0x9c, 0x58, 0x99, 0x01, 0x12, 0x11, 0x12, 0xd4, 0x83,
	0xb9, 0x16, 0xe6, 0x36, 0x29, 0x74, 0xa1, 0xe2, 0x02, 0x25, 0x06, 0x21, 0x7d, 0x2e, 0xf6, 0xe0,
	0xc4, 0x4a, 0x90, 0x55, 0x42, 0x02, 0x70, 0x79, 0xa8, 0x83, 0xa4, 0xd0, 0x44, 0x40, 0x1a, 0x92,
	0xd8, 0xc0, 0x02, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x4d, 0xab, 0x8c, 0x0e, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloserverClient is the client API for Helloserver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloserverClient interface {
	//一个打招呼的函数
	Sayhello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error)
	//一个说名字的服务
	Sayname(ctx context.Context, in *NameReq, opts ...grpc.CallOption) (*NameRsp, error)
}

type helloserverClient struct {
	cc *grpc.ClientConn
}

func NewHelloserverClient(cc *grpc.ClientConn) HelloserverClient {
	return &helloserverClient{cc}
}

func (c *helloserverClient) Sayhello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRsp, error) {
	out := new(HelloRsp)
	err := c.cc.Invoke(ctx, "/myproto.Helloserver/Sayhello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloserverClient) Sayname(ctx context.Context, in *NameReq, opts ...grpc.CallOption) (*NameRsp, error) {
	out := new(NameRsp)
	err := c.cc.Invoke(ctx, "/myproto.Helloserver/Sayname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloserverServer is the server API for Helloserver service.
type HelloserverServer interface {
	//一个打招呼的函数
	Sayhello(context.Context, *HelloReq) (*HelloRsp, error)
	//一个说名字的服务
	Sayname(context.Context, *NameReq) (*NameRsp, error)
}

// UnimplementedHelloserverServer can be embedded to have forward compatible implementations.
type UnimplementedHelloserverServer struct {
}

func (*UnimplementedHelloserverServer) Sayhello(ctx context.Context, req *HelloReq) (*HelloRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sayhello not implemented")
}
func (*UnimplementedHelloserverServer) Sayname(ctx context.Context, req *NameReq) (*NameRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sayname not implemented")
}

func RegisterHelloserverServer(s *grpc.Server, srv HelloserverServer) {
	s.RegisterService(&_Helloserver_serviceDesc, srv)
}

func _Helloserver_Sayhello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloserverServer).Sayhello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproto.Helloserver/Sayhello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloserverServer).Sayhello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloserver_Sayname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloserverServer).Sayname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproto.Helloserver/Sayname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloserverServer).Sayname(ctx, req.(*NameReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Helloserver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myproto.Helloserver",
	HandlerType: (*HelloserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sayhello",
			Handler:    _Helloserver_Sayhello_Handler,
		},
		{
			MethodName: "Sayname",
			Handler:    _Helloserver_Sayname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myproto.proto",
}
