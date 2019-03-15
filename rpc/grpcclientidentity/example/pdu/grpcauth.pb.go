// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcauth.proto

package pdu

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GreetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dfba7be0cf69353, []int{0}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dfba7be0cf69353, []int{1}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "pdu.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "pdu.GreetResponse")
}

func init() { proto.RegisterFile("grpcauth.proto", fileDescriptor_1dfba7be0cf69353) }

var fileDescriptor_1dfba7be0cf69353 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2f, 0x2a, 0x48,
	0x4e, 0x2c, 0x2d, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x48, 0x29, 0x55,
	0x52, 0xe2, 0xe2, 0x71, 0x2f, 0x4a, 0x4d, 0x2d, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x95, 0x14, 0xb9, 0x78, 0xa1, 0x6a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x04, 0xb8, 0x98,
	0x73, 0x8b, 0xd3, 0xa1, 0x6a, 0x40, 0x4c, 0x23, 0x6b, 0x2e, 0x76, 0xb0, 0x92, 0xd4, 0x22, 0x21,
	0x03, 0x2e, 0x56, 0x30, 0x53, 0x48, 0x50, 0xaf, 0x20, 0xa5, 0x54, 0x0f, 0xd9, 0x74, 0x29, 0x21,
	0x64, 0x21, 0x88, 0x61, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xf7, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xa8, 0x53, 0x2f, 0x4c, 0xa1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/pdu.Greeter/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pdu.Greeter/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pdu.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Greeter_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcauth.proto",
}
