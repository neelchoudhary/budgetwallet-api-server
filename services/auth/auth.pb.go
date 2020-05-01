// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/auth/auth.proto

package auth

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fullname             string   `protobuf:"bytes,2,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	CreatedOn            string   `protobuf:"bytes,5,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a414a1b227848998, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

type SignUpUser struct {
	Fullname             string   `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpUser) Reset()         { *m = SignUpUser{} }
func (m *SignUpUser) String() string { return proto.CompactTextString(m) }
func (*SignUpUser) ProtoMessage()    {}
func (*SignUpUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_a414a1b227848998, []int{1}
}

func (m *SignUpUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpUser.Unmarshal(m, b)
}
func (m *SignUpUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpUser.Marshal(b, m, deterministic)
}
func (m *SignUpUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpUser.Merge(m, src)
}
func (m *SignUpUser) XXX_Size() int {
	return xxx_messageInfo_SignUpUser.Size(m)
}
func (m *SignUpUser) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpUser.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpUser proto.InternalMessageInfo

func (m *SignUpUser) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *SignUpUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginUser struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginUser) Reset()         { *m = LoginUser{} }
func (m *LoginUser) String() string { return proto.CompactTextString(m) }
func (*LoginUser) ProtoMessage()    {}
func (*LoginUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_a414a1b227848998, []int{2}
}

func (m *LoginUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginUser.Unmarshal(m, b)
}
func (m *LoginUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginUser.Marshal(b, m, deterministic)
}
func (m *LoginUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginUser.Merge(m, src)
}
func (m *LoginUser) XXX_Size() int {
	return xxx_messageInfo_LoginUser.Size(m)
}
func (m *LoginUser) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginUser.DiscardUnknown(m)
}

var xxx_messageInfo_LoginUser proto.InternalMessageInfo

func (m *LoginUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignupRequest struct {
	SignUpUser           *SignUpUser `protobuf:"bytes,1,opt,name=signUpUser,proto3" json:"signUpUser,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SignupRequest) Reset()         { *m = SignupRequest{} }
func (m *SignupRequest) String() string { return proto.CompactTextString(m) }
func (*SignupRequest) ProtoMessage()    {}
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a414a1b227848998, []int{3}
}

func (m *SignupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupRequest.Unmarshal(m, b)
}
func (m *SignupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupRequest.Marshal(b, m, deterministic)
}
func (m *SignupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupRequest.Merge(m, src)
}
func (m *SignupRequest) XXX_Size() int {
	return xxx_messageInfo_SignupRequest.Size(m)
}
func (m *SignupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignupRequest proto.InternalMessageInfo

func (m *SignupRequest) GetSignUpUser() *SignUpUser {
	if m != nil {
		return m.SignUpUser
	}
	return nil
}

type SignupResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignupResponse) Reset()         { *m = SignupResponse{} }
func (m *SignupResponse) String() string { return proto.CompactTextString(m) }
func (*SignupResponse) ProtoMessage()    {}
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a414a1b227848998, []int{4}
}

func (m *SignupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupResponse.Unmarshal(m, b)
}
func (m *SignupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupResponse.Marshal(b, m, deterministic)
}
func (m *SignupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupResponse.Merge(m, src)
}
func (m *SignupResponse) XXX_Size() int {
	return xxx_messageInfo_SignupResponse.Size(m)
}
func (m *SignupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignupResponse proto.InternalMessageInfo

func (m *SignupResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type LoginRequest struct {
	LoginUser            *LoginUser `protobuf:"bytes,1,opt,name=loginUser,proto3" json:"loginUser,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a414a1b227848998, []int{5}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetLoginUser() *LoginUser {
	if m != nil {
		return m.LoginUser
	}
	return nil
}

type LoginResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a414a1b227848998, []int{6}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*SignUpUser)(nil), "auth.SignUpUser")
	proto.RegisterType((*LoginUser)(nil), "auth.LoginUser")
	proto.RegisterType((*SignupRequest)(nil), "auth.SignupRequest")
	proto.RegisterType((*SignupResponse)(nil), "auth.SignupResponse")
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
}

func init() { proto.RegisterFile("services/auth/auth.proto", fileDescriptor_a414a1b227848998) }

var fileDescriptor_a414a1b227848998 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xad, 0xe4, 0x8f, 0x5a, 0xe3, 0xda, 0x2d, 0x6b, 0x1f, 0x84, 0xa1, 0x60, 0x74, 0x32, 0x85,
	0x5a, 0xc5, 0xa5, 0x50, 0x0a, 0xa6, 0x38, 0xe7, 0x40, 0x40, 0xc6, 0x17, 0x5f, 0xc2, 0x5a, 0x9a,
	0x48, 0x22, 0xd2, 0xae, 0xb2, 0x1f, 0xf9, 0x38, 0xe7, 0x8f, 0x07, 0xad, 0x3e, 0x6c, 0x07, 0x42,
	0x0e, 0xb9, 0x08, 0xbd, 0xd9, 0x79, 0xf3, 0xde, 0x1b, 0x06, 0x5c, 0x89, 0xe2, 0x3e, 0x0d, 0x51,
	0xfa, 0x54, 0xab, 0xc4, 0x7c, 0x96, 0x85, 0xe0, 0x8a, 0x93, 0x6e, 0xf9, 0xef, 0x3d, 0x5b, 0xd0,
	0xdd, 0x49, 0x14, 0x64, 0x0c, 0x76, 0x1a, 0xb9, 0xd6, 0xdc, 0x5a, 0x74, 0x02, 0x3b, 0x8d, 0xc8,
	0x0c, 0x06, 0x37, 0x3a, 0xcb, 0x18, 0xcd, 0xd1, 0xb5, 0xe7, 0xd6, 0xc2, 0x09, 0x5a, 0x4c, 0xa6,
	0xd0, 0xc3, 0x9c, 0xa6, 0x99, 0xdb, 0x31, 0x0f, 0x15, 0x28, 0x19, 0x05, 0x95, 0xf2, 0x81, 0x8b,
	0xc8, 0xed, 0x56, 0x8c, 0x06, 0x93, 0xef, 0x00, 0xa1, 0x40, 0xaa, 0x30, 0xba, 0xe6, 0xcc, 0xed,
	0x99, 0x57, 0xa7, 0xae, 0x5c, 0x31, 0x6f, 0x0f, 0xb0, 0x4d, 0x63, 0xb6, 0x2b, 0x8c, 0x95, 0x53,
	0x69, 0xeb, 0xa3, 0xd2, 0xde, 0x1a, 0x9c, 0x4b, 0x1e, 0xa7, 0xcc, 0x8c, 0x6e, 0xe9, 0xd6, 0x5b,
	0x74, 0xfb, 0x15, 0x7d, 0x03, 0xa3, 0xd2, 0x9a, 0x2e, 0x02, 0xbc, 0xd3, 0x28, 0x15, 0xf9, 0x05,
	0x20, 0x5b, 0xaf, 0x66, 0xce, 0x70, 0xf5, 0x6d, 0x69, 0x16, 0x7b, 0xcc, 0x10, 0x9c, 0xf4, 0x78,
	0x3f, 0x60, 0xdc, 0x8c, 0x90, 0x05, 0x67, 0x12, 0x89, 0x0b, 0x9f, 0xa5, 0x0e, 0x43, 0x94, 0xd2,
	0x0c, 0x18, 0x04, 0x0d, 0xf4, 0xd6, 0xf0, 0xc5, 0xb8, 0x6d, 0xd4, 0x7e, 0x82, 0x93, 0x35, 0xee,
	0x6b, 0xb1, 0xaf, 0x95, 0x58, 0x1b, 0x2a, 0x38, 0x76, 0x78, 0xff, 0x61, 0x54, 0xd3, 0xdf, 0x53,
	0x2a, 0x57, 0xa1, 0xf8, 0x2d, 0xb2, 0x3a, 0x71, 0x05, 0x56, 0x8f, 0x30, 0xdc, 0x68, 0x95, 0x6c,
	0xab, 0xab, 0x21, 0x7f, 0xa0, 0x5f, 0x59, 0x27, 0x93, 0x63, 0xc4, 0x76, 0x17, 0xb3, 0xe9, 0x79,
	0xb1, 0xd2, 0xf4, 0x3e, 0x91, 0x15, 0xf4, 0x8c, 0x0d, 0x42, 0x4e, 0xbc, 0x36, 0xa4, 0xc9, 0x59,
	0xad, 0xe1, 0x5c, 0xfc, 0xdb, 0xff, 0x8d, 0x53, 0x95, 0xe8, 0xc3, 0x32, 0xe4, 0xb9, 0xcf, 0x10,
	0xb3, 0x30, 0xe1, 0x3a, 0x4a, 0xa8, 0x78, 0xf2, 0x0f, 0x3a, 0x8a, 0x51, 0xe5, 0x94, 0xd1, 0x18,
	0x45, 0x2c, 0x8a, 0xd0, 0x3f, 0x3b, 0xeb, 0x43, 0xdf, 0x9c, 0xf4, 0xef, 0x97, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x61, 0xda, 0x03, 0x13, 0xee, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Signup(ctx context.Context, req *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _AuthService_Signup_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/auth/auth.proto",
}