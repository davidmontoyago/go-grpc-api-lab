// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tls-auth/secure-service.proto

package api

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

type SecureRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecureRequest) Reset()         { *m = SecureRequest{} }
func (m *SecureRequest) String() string { return proto.CompactTextString(m) }
func (*SecureRequest) ProtoMessage()    {}
func (*SecureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d08ac23261e4757, []int{0}
}

func (m *SecureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecureRequest.Unmarshal(m, b)
}
func (m *SecureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecureRequest.Marshal(b, m, deterministic)
}
func (m *SecureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecureRequest.Merge(m, src)
}
func (m *SecureRequest) XXX_Size() int {
	return xxx_messageInfo_SecureRequest.Size(m)
}
func (m *SecureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SecureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SecureRequest proto.InternalMessageInfo

func (m *SecureRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type SecureResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecureResponse) Reset()         { *m = SecureResponse{} }
func (m *SecureResponse) String() string { return proto.CompactTextString(m) }
func (*SecureResponse) ProtoMessage()    {}
func (*SecureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d08ac23261e4757, []int{1}
}

func (m *SecureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecureResponse.Unmarshal(m, b)
}
func (m *SecureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecureResponse.Marshal(b, m, deterministic)
}
func (m *SecureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecureResponse.Merge(m, src)
}
func (m *SecureResponse) XXX_Size() int {
	return xxx_messageInfo_SecureResponse.Size(m)
}
func (m *SecureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SecureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SecureResponse proto.InternalMessageInfo

func (m *SecureResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*SecureRequest)(nil), "api.SecureRequest")
	proto.RegisterType((*SecureResponse)(nil), "api.SecureResponse")
}

func init() { proto.RegisterFile("tls-auth/secure-service.proto", fileDescriptor_1d08ac23261e4757) }

var fileDescriptor_1d08ac23261e4757 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xc9, 0x29, 0xd6,
	0x4d, 0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0x4e, 0x4d, 0x2e, 0x2d, 0x4a, 0xd5, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52,
	0xe6, 0xe2, 0x0d, 0x06, 0x4b, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1,
	0xa4, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x5a, 0x5c,
	0x7c, 0x30, 0x45, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa5, 0xc9,
	0xc9, 0xa9, 0xc5, 0xc5, 0x60, 0x85, 0x1c, 0x41, 0x30, 0xae, 0x91, 0x07, 0xcc, 0xc0, 0x60, 0x88,
	0x65, 0x42, 0xe6, 0x5c, 0x3c, 0xce, 0x19, 0xa9, 0xc9, 0xd9, 0xbe, 0x95, 0xce, 0x45, 0xa9, 0x29,
	0xc5, 0x42, 0x42, 0x7a, 0x89, 0x05, 0x99, 0x7a, 0x28, 0x96, 0x4a, 0x09, 0xa3, 0x88, 0x41, 0xec,
	0x48, 0x62, 0x03, 0x3b, 0xd3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x48, 0xdd, 0xeb, 0xa6, 0xc7,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecureServiceClient is the client API for SecureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecureServiceClient interface {
	CheckMyCreds(ctx context.Context, in *SecureRequest, opts ...grpc.CallOption) (*SecureResponse, error)
}

type secureServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecureServiceClient(cc *grpc.ClientConn) SecureServiceClient {
	return &secureServiceClient{cc}
}

func (c *secureServiceClient) CheckMyCreds(ctx context.Context, in *SecureRequest, opts ...grpc.CallOption) (*SecureResponse, error) {
	out := new(SecureResponse)
	err := c.cc.Invoke(ctx, "/api.SecureService/CheckMyCreds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecureServiceServer is the server API for SecureService service.
type SecureServiceServer interface {
	CheckMyCreds(context.Context, *SecureRequest) (*SecureResponse, error)
}

// UnimplementedSecureServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecureServiceServer struct {
}

func (*UnimplementedSecureServiceServer) CheckMyCreds(ctx context.Context, req *SecureRequest) (*SecureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckMyCreds not implemented")
}

func RegisterSecureServiceServer(s *grpc.Server, srv SecureServiceServer) {
	s.RegisterService(&_SecureService_serviceDesc, srv)
}

func _SecureService_CheckMyCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureServiceServer).CheckMyCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SecureService/CheckMyCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureServiceServer).CheckMyCreds(ctx, req.(*SecureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecureService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SecureService",
	HandlerType: (*SecureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckMyCreds",
			Handler:    _SecureService_CheckMyCreds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tls-auth/secure-service.proto",
}
