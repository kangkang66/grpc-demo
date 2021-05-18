// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Empty from public import google/protobuf/empty.proto
type Empty = empty.Empty

type Param struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Referer              string   `protobuf:"bytes,2,opt,name=referer,proto3" json:"referer,omitempty"`
	UserAgent            string   `protobuf:"bytes,3,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Param) Reset()         { *m = Param{} }
func (m *Param) String() string { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()    {}
func (*Param) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Param) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Param.Unmarshal(m, b)
}
func (m *Param) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Param.Marshal(b, m, deterministic)
}
func (m *Param) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Param.Merge(m, src)
}
func (m *Param) XXX_Size() int {
	return xxx_messageInfo_Param.Size(m)
}
func (m *Param) XXX_DiscardUnknown() {
	xxx_messageInfo_Param.DiscardUnknown(m)
}

var xxx_messageInfo_Param proto.InternalMessageInfo

func (m *Param) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Param) GetReferer() string {
	if m != nil {
		return m.Referer
	}
	return ""
}

func (m *Param) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

type Response struct {
	Val                  bool     `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVal() bool {
	if m != nil {
		return m.Val
	}
	return false
}

type HttpConfig struct {
	Token                bool     `protobuf:"varint,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpConfig) Reset()         { *m = HttpConfig{} }
func (m *HttpConfig) String() string { return proto.CompactTextString(m) }
func (*HttpConfig) ProtoMessage()    {}
func (*HttpConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *HttpConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpConfig.Unmarshal(m, b)
}
func (m *HttpConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpConfig.Marshal(b, m, deterministic)
}
func (m *HttpConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpConfig.Merge(m, src)
}
func (m *HttpConfig) XXX_Size() int {
	return xxx_messageInfo_HttpConfig.Size(m)
}
func (m *HttpConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpConfig proto.InternalMessageInfo

func (m *HttpConfig) GetToken() bool {
	if m != nil {
		return m.Token
	}
	return false
}

var E_Api = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*HttpConfig)(nil),
	Field:         65801,
	Name:          "mig_one.user.api",
	Tag:           "bytes,65801,opt,name=api",
	Filename:      "user.proto",
}

func init() {
	proto.RegisterType((*Param)(nil), "mig_one.user.Param")
	proto.RegisterType((*Response)(nil), "mig_one.user.Response")
	proto.RegisterType((*HttpConfig)(nil), "mig_one.user.HttpConfig")
	proto.RegisterExtension(E_Api)
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xcd, 0x4b, 0xfb, 0x30,
	0x18, 0xa6, 0xdd, 0xc7, 0xaf, 0x7b, 0x7f, 0x32, 0x24, 0x8a, 0x84, 0xfa, 0x41, 0xe9, 0x69, 0xa7,
	0x14, 0xe6, 0xcd, 0xdb, 0xf4, 0xe2, 0x40, 0xb1, 0x14, 0xbc, 0x78, 0x19, 0x9d, 0x7b, 0x5b, 0xc3,
	0xd6, 0x24, 0x24, 0x99, 0xe0, 0xd5, 0x9b, 0x37, 0xff, 0x2e, 0xff, 0x2a, 0x49, 0xba, 0xa2, 0xf3,
	0x96, 0xe7, 0x83, 0x27, 0xcf, 0xfb, 0x00, 0x6c, 0x0d, 0x6a, 0xa6, 0xb4, 0xb4, 0x92, 0x1c, 0x34,
	0xbc, 0x5e, 0x48, 0x81, 0xcc, 0x71, 0x71, 0x52, 0x4b, 0x59, 0x6f, 0x30, 0xf3, 0xda, 0x72, 0x5b,
	0x65, 0x2b, 0x34, 0xcf, 0x9a, 0x2b, 0x2b, 0x77, 0xfe, 0xf8, 0xf4, 0xaf, 0x03, 0x1b, 0x65, 0xdf,
	0x5a, 0x31, 0xcd, 0x61, 0x90, 0x97, 0xba, 0x6c, 0xc8, 0x18, 0x42, 0xae, 0x68, 0x90, 0x04, 0x93,
	0x51, 0x11, 0x72, 0x45, 0x28, 0xfc, 0xd3, 0x58, 0xa1, 0x46, 0x4d, 0x43, 0x4f, 0x76, 0x90, 0x9c,
	0xb7, 0x6d, 0x16, 0x65, 0x8d, 0xc2, 0xd2, 0x9e, 0x17, 0x47, 0x8e, 0x99, 0x39, 0x22, 0x3d, 0x83,
	0xa8, 0x40, 0xa3, 0xa4, 0x30, 0x48, 0x0e, 0xa1, 0xf7, 0x5a, 0x6e, 0x7c, 0x6a, 0x54, 0xb8, 0x67,
	0x9a, 0x02, 0xdc, 0x5a, 0xab, 0x6e, 0xa4, 0xa8, 0x78, 0x4d, 0x8e, 0x61, 0x60, 0xe5, 0x1a, 0xc5,
	0xce, 0xd1, 0x82, 0xe9, 0x1c, 0xfa, 0x8f, 0x06, 0x35, 0x99, 0x01, 0xcc, 0x4d, 0xc1, 0xcd, 0xda,
	0xa3, 0x23, 0xf6, 0xfb, 0x6e, 0xe6, 0x5b, 0xc7, 0x27, 0xfb, 0x64, 0xf7, 0x71, 0x3a, 0xfc, 0xfa,
	0x4c, 0xc2, 0x28, 0xb8, 0xba, 0x83, 0x5e, 0xa9, 0x38, 0xb9, 0x60, 0xed, 0x06, 0xac, 0xdb, 0x80,
	0xdd, 0xa3, 0x7d, 0x91, 0xab, 0x07, 0x65, 0xb9, 0x14, 0x86, 0x7e, 0xbc, 0xf7, 0x93, 0x60, 0xf2,
	0x7f, 0x4a, 0xf7, 0xe3, 0x7e, 0x9a, 0x16, 0x2e, 0xe6, 0x7a, 0xfc, 0xd4, 0x6d, 0x9f, 0x39, 0x3d,
	0x0f, 0x96, 0x43, 0x1f, 0x78, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x85, 0x34, 0xcc, 0x5c, 0xa0,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	IsRiskUser(ctx context.Context, in *Param, opts ...grpc.CallOption) (*Response, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) IsRiskUser(ctx context.Context, in *Param, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mig_one.user.User/IsRiskUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	IsRiskUser(context.Context, *Param) (*Response, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) IsRiskUser(ctx context.Context, req *Param) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsRiskUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_IsRiskUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Param)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).IsRiskUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mig_one.user.User/IsRiskUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).IsRiskUser(ctx, req.(*Param))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mig_one.user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsRiskUser",
			Handler:    _User_IsRiskUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
