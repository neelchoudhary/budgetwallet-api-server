// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/userfinances/userFinances.proto

package userfinances

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	shared "github.com/neelchoudhary/budgetwallet-api-server/services/shared"
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

type GetFinancialInstitutionsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFinancialInstitutionsRequest) Reset()         { *m = GetFinancialInstitutionsRequest{} }
func (m *GetFinancialInstitutionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetFinancialInstitutionsRequest) ProtoMessage()    {}
func (*GetFinancialInstitutionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec379254392880eb, []int{0}
}

func (m *GetFinancialInstitutionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFinancialInstitutionsRequest.Unmarshal(m, b)
}
func (m *GetFinancialInstitutionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFinancialInstitutionsRequest.Marshal(b, m, deterministic)
}
func (m *GetFinancialInstitutionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFinancialInstitutionsRequest.Merge(m, src)
}
func (m *GetFinancialInstitutionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetFinancialInstitutionsRequest.Size(m)
}
func (m *GetFinancialInstitutionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFinancialInstitutionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFinancialInstitutionsRequest proto.InternalMessageInfo

type GetFinancialInstitutionsResponse struct {
	FinancialInstitutions []*shared.FinancialInstitution `protobuf:"bytes,1,rep,name=financial_institutions,json=financialInstitutions,proto3" json:"financial_institutions,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                       `json:"-"`
	XXX_unrecognized      []byte                         `json:"-"`
	XXX_sizecache         int32                          `json:"-"`
}

func (m *GetFinancialInstitutionsResponse) Reset()         { *m = GetFinancialInstitutionsResponse{} }
func (m *GetFinancialInstitutionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetFinancialInstitutionsResponse) ProtoMessage()    {}
func (*GetFinancialInstitutionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec379254392880eb, []int{1}
}

func (m *GetFinancialInstitutionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFinancialInstitutionsResponse.Unmarshal(m, b)
}
func (m *GetFinancialInstitutionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFinancialInstitutionsResponse.Marshal(b, m, deterministic)
}
func (m *GetFinancialInstitutionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFinancialInstitutionsResponse.Merge(m, src)
}
func (m *GetFinancialInstitutionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetFinancialInstitutionsResponse.Size(m)
}
func (m *GetFinancialInstitutionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFinancialInstitutionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFinancialInstitutionsResponse proto.InternalMessageInfo

func (m *GetFinancialInstitutionsResponse) GetFinancialInstitutions() []*shared.FinancialInstitution {
	if m != nil {
		return m.FinancialInstitutions
	}
	return nil
}

type GetFinancialAccountsRequest struct {
	ItemId               int64    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFinancialAccountsRequest) Reset()         { *m = GetFinancialAccountsRequest{} }
func (m *GetFinancialAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*GetFinancialAccountsRequest) ProtoMessage()    {}
func (*GetFinancialAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec379254392880eb, []int{2}
}

func (m *GetFinancialAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFinancialAccountsRequest.Unmarshal(m, b)
}
func (m *GetFinancialAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFinancialAccountsRequest.Marshal(b, m, deterministic)
}
func (m *GetFinancialAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFinancialAccountsRequest.Merge(m, src)
}
func (m *GetFinancialAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_GetFinancialAccountsRequest.Size(m)
}
func (m *GetFinancialAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFinancialAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFinancialAccountsRequest proto.InternalMessageInfo

func (m *GetFinancialAccountsRequest) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

type GetFinancialAccountsResponse struct {
	FinancialAccounts    []*shared.FinancialAccount `protobuf:"bytes,1,rep,name=financial_accounts,json=financialAccounts,proto3" json:"financial_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetFinancialAccountsResponse) Reset()         { *m = GetFinancialAccountsResponse{} }
func (m *GetFinancialAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*GetFinancialAccountsResponse) ProtoMessage()    {}
func (*GetFinancialAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec379254392880eb, []int{3}
}

func (m *GetFinancialAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFinancialAccountsResponse.Unmarshal(m, b)
}
func (m *GetFinancialAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFinancialAccountsResponse.Marshal(b, m, deterministic)
}
func (m *GetFinancialAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFinancialAccountsResponse.Merge(m, src)
}
func (m *GetFinancialAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_GetFinancialAccountsResponse.Size(m)
}
func (m *GetFinancialAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFinancialAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFinancialAccountsResponse proto.InternalMessageInfo

func (m *GetFinancialAccountsResponse) GetFinancialAccounts() []*shared.FinancialAccount {
	if m != nil {
		return m.FinancialAccounts
	}
	return nil
}

type ToggleFinancialAccountRequest struct {
	ItemId               int64    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Selected             bool     `protobuf:"varint,3,opt,name=selected,proto3" json:"selected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToggleFinancialAccountRequest) Reset()         { *m = ToggleFinancialAccountRequest{} }
func (m *ToggleFinancialAccountRequest) String() string { return proto.CompactTextString(m) }
func (*ToggleFinancialAccountRequest) ProtoMessage()    {}
func (*ToggleFinancialAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec379254392880eb, []int{4}
}

func (m *ToggleFinancialAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToggleFinancialAccountRequest.Unmarshal(m, b)
}
func (m *ToggleFinancialAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToggleFinancialAccountRequest.Marshal(b, m, deterministic)
}
func (m *ToggleFinancialAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToggleFinancialAccountRequest.Merge(m, src)
}
func (m *ToggleFinancialAccountRequest) XXX_Size() int {
	return xxx_messageInfo_ToggleFinancialAccountRequest.Size(m)
}
func (m *ToggleFinancialAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ToggleFinancialAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ToggleFinancialAccountRequest proto.InternalMessageInfo

func (m *ToggleFinancialAccountRequest) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ToggleFinancialAccountRequest) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *ToggleFinancialAccountRequest) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

type ToggleFinancialAccountResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToggleFinancialAccountResponse) Reset()         { *m = ToggleFinancialAccountResponse{} }
func (m *ToggleFinancialAccountResponse) String() string { return proto.CompactTextString(m) }
func (*ToggleFinancialAccountResponse) ProtoMessage()    {}
func (*ToggleFinancialAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec379254392880eb, []int{5}
}

func (m *ToggleFinancialAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToggleFinancialAccountResponse.Unmarshal(m, b)
}
func (m *ToggleFinancialAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToggleFinancialAccountResponse.Marshal(b, m, deterministic)
}
func (m *ToggleFinancialAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToggleFinancialAccountResponse.Merge(m, src)
}
func (m *ToggleFinancialAccountResponse) XXX_Size() int {
	return xxx_messageInfo_ToggleFinancialAccountResponse.Size(m)
}
func (m *ToggleFinancialAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ToggleFinancialAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ToggleFinancialAccountResponse proto.InternalMessageInfo

func (m *ToggleFinancialAccountResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetFinancialTransactionsRequest struct {
	ItemId               int64    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFinancialTransactionsRequest) Reset()         { *m = GetFinancialTransactionsRequest{} }
func (m *GetFinancialTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetFinancialTransactionsRequest) ProtoMessage()    {}
func (*GetFinancialTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec379254392880eb, []int{6}
}

func (m *GetFinancialTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFinancialTransactionsRequest.Unmarshal(m, b)
}
func (m *GetFinancialTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFinancialTransactionsRequest.Marshal(b, m, deterministic)
}
func (m *GetFinancialTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFinancialTransactionsRequest.Merge(m, src)
}
func (m *GetFinancialTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetFinancialTransactionsRequest.Size(m)
}
func (m *GetFinancialTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFinancialTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFinancialTransactionsRequest proto.InternalMessageInfo

func (m *GetFinancialTransactionsRequest) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

type GetFinancialTransactionsResponse struct {
	FinancialTransactions []*shared.FinancialTransaction `protobuf:"bytes,1,rep,name=financial_transactions,json=financialTransactions,proto3" json:"financial_transactions,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                       `json:"-"`
	XXX_unrecognized      []byte                         `json:"-"`
	XXX_sizecache         int32                          `json:"-"`
}

func (m *GetFinancialTransactionsResponse) Reset()         { *m = GetFinancialTransactionsResponse{} }
func (m *GetFinancialTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetFinancialTransactionsResponse) ProtoMessage()    {}
func (*GetFinancialTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec379254392880eb, []int{7}
}

func (m *GetFinancialTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFinancialTransactionsResponse.Unmarshal(m, b)
}
func (m *GetFinancialTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFinancialTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *GetFinancialTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFinancialTransactionsResponse.Merge(m, src)
}
func (m *GetFinancialTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetFinancialTransactionsResponse.Size(m)
}
func (m *GetFinancialTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFinancialTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFinancialTransactionsResponse proto.InternalMessageInfo

func (m *GetFinancialTransactionsResponse) GetFinancialTransactions() []*shared.FinancialTransaction {
	if m != nil {
		return m.FinancialTransactions
	}
	return nil
}

func init() {
	proto.RegisterType((*GetFinancialInstitutionsRequest)(nil), "userfinances.GetFinancialInstitutionsRequest")
	proto.RegisterType((*GetFinancialInstitutionsResponse)(nil), "userfinances.GetFinancialInstitutionsResponse")
	proto.RegisterType((*GetFinancialAccountsRequest)(nil), "userfinances.GetFinancialAccountsRequest")
	proto.RegisterType((*GetFinancialAccountsResponse)(nil), "userfinances.GetFinancialAccountsResponse")
	proto.RegisterType((*ToggleFinancialAccountRequest)(nil), "userfinances.ToggleFinancialAccountRequest")
	proto.RegisterType((*ToggleFinancialAccountResponse)(nil), "userfinances.ToggleFinancialAccountResponse")
	proto.RegisterType((*GetFinancialTransactionsRequest)(nil), "userfinances.GetFinancialTransactionsRequest")
	proto.RegisterType((*GetFinancialTransactionsResponse)(nil), "userfinances.GetFinancialTransactionsResponse")
}

func init() {
	proto.RegisterFile("services/userfinances/userFinances.proto", fileDescriptor_ec379254392880eb)
}

var fileDescriptor_ec379254392880eb = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x57, 0x2a, 0x6d, 0xe5, 0xc1, 0x05, 0x03, 0x23, 0x0a, 0x1d, 0x14, 0x9f, 0xca, 0x8f,
	0xa6, 0xd2, 0x90, 0x38, 0xec, 0x06, 0x87, 0x8d, 0x5e, 0xb3, 0x71, 0xe1, 0x32, 0xb9, 0xce, 0x6b,
	0x6a, 0x29, 0xb3, 0x4b, 0x6c, 0x33, 0x21, 0xfe, 0x70, 0xae, 0xa8, 0x8d, 0x93, 0xb9, 0x99, 0xd3,
	0xf5, 0x54, 0xbd, 0xfa, 0xfb, 0xbe, 0xef, 0xf9, 0xbd, 0x4f, 0x0c, 0x63, 0x8d, 0xe5, 0x6f, 0xc1,
	0x51, 0x4f, 0xad, 0xc6, 0x72, 0x21, 0x24, 0x93, 0x75, 0x70, 0xee, 0x82, 0x64, 0x55, 0x2a, 0xa3,
	0xc8, 0x53, 0x5f, 0x10, 0x0f, 0x9b, 0x3c, 0xbd, 0x64, 0x25, 0x66, 0xee, 0xa7, 0xd2, 0xd2, 0x77,
	0xf0, 0xf6, 0x02, 0x4d, 0x65, 0x20, 0x58, 0x31, 0x93, 0xda, 0x08, 0x63, 0x8d, 0x50, 0x52, 0xa7,
	0xf8, 0xcb, 0xa2, 0x36, 0xf4, 0x16, 0x46, 0xdd, 0x12, 0xbd, 0x52, 0x52, 0x23, 0xb9, 0x84, 0xe3,
	0x45, 0x2d, 0xb8, 0x16, 0x9e, 0x22, 0xea, 0x8d, 0xfa, 0xe3, 0x27, 0xa7, 0xc3, 0xc4, 0x55, 0x0d,
	0xd9, 0xa4, 0x2f, 0x17, 0x21, 0x73, 0xfa, 0x05, 0x5e, 0xfb, 0x85, 0xbf, 0x72, 0xae, 0xac, 0x34,
	0x75, 0x5f, 0xe4, 0x15, 0x1c, 0x09, 0x83, 0x37, 0xd7, 0x22, 0x8b, 0x7a, 0xa3, 0xde, 0xb8, 0x9f,
	0x1e, 0xae, 0xc3, 0x59, 0x46, 0x73, 0x18, 0x86, 0xf3, 0x5c, 0xb3, 0x17, 0x40, 0xee, 0x9a, 0x65,
	0xee, 0xd4, 0x35, 0x1a, 0xdd, 0x6b, 0xd4, 0xa5, 0xa7, 0xcf, 0x16, 0x6d, 0x43, 0xaa, 0xe1, 0xe4,
	0x4a, 0xe5, 0x79, 0x81, 0xf7, 0xc4, 0x0f, 0xb4, 0x48, 0x4e, 0x00, 0x5c, 0xe1, 0xf5, 0xd9, 0xa3,
	0xcd, 0xd9, 0x63, 0xf7, 0xcf, 0x2c, 0x23, 0x31, 0x0c, 0x34, 0x16, 0xc8, 0x0d, 0x66, 0x51, 0x7f,
	0xd4, 0x1b, 0x0f, 0xd2, 0x26, 0xa6, 0x67, 0xf0, 0xa6, 0xab, 0xa8, 0xbb, 0x5f, 0x04, 0x47, 0xda,
	0x72, 0x8e, 0x5a, 0x6f, 0xaa, 0x0e, 0xd2, 0x3a, 0xa4, 0x67, 0xdb, 0xdb, 0xbe, 0x2a, 0x99, 0xd4,
	0x8c, 0xfb, 0xdb, 0xee, 0x9e, 0x6a, 0x0b, 0x83, 0xed, 0xdc, 0x10, 0x06, 0xc6, 0x53, 0x74, 0x62,
	0xe0, 0xd9, 0x78, 0x18, 0xf8, 0xe6, 0xa7, 0xff, 0xfa, 0xf0, 0xfc, 0x87, 0x47, 0xf9, 0x65, 0xc5,
	0x33, 0xf9, 0x0b, 0x51, 0x17, 0x97, 0x64, 0x92, 0xf8, 0xdf, 0x40, 0xf2, 0x00, 0xe2, 0x71, 0xb2,
	0xaf, 0xbc, 0xba, 0x27, 0x3d, 0x20, 0x0a, 0x5e, 0x84, 0x18, 0x23, 0xef, 0xbb, 0x9d, 0x5a, 0xfc,
	0xc6, 0x1f, 0xf6, 0x91, 0x36, 0x05, 0x2d, 0x1c, 0x87, 0xd7, 0x4e, 0x3e, 0x6e, 0xfb, 0xec, 0x24,
	0x32, 0xfe, 0xb4, 0x9f, 0xb8, 0x29, 0xdb, 0x1a, 0xb2, 0xbf, 0x98, 0x5d, 0x43, 0x0e, 0x90, 0xb5,
	0x6b, 0xc8, 0x21, 0x98, 0xe8, 0xc1, 0xb7, 0xef, 0x3f, 0xcf, 0x73, 0x61, 0x96, 0x76, 0x9e, 0x70,
	0x75, 0x33, 0x95, 0x88, 0x05, 0x5f, 0x2a, 0x9b, 0x2d, 0x59, 0xf9, 0x67, 0x3a, 0xb7, 0x59, 0x8e,
	0xe6, 0x96, 0x15, 0x05, 0x9a, 0x09, 0x5b, 0x89, 0xc9, 0xfa, 0x99, 0xc3, 0x72, 0x1a, 0x7c, 0x25,
	0xe7, 0x87, 0x9b, 0xd7, 0xee, 0xf3, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xf3, 0xe8, 0x47,
	0x45, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserFinancesServiceClient is the client API for UserFinancesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserFinancesServiceClient interface {
	GetFinancialInstitutions(ctx context.Context, in *GetFinancialInstitutionsRequest, opts ...grpc.CallOption) (*GetFinancialInstitutionsResponse, error)
	GetFinancialAccounts(ctx context.Context, in *GetFinancialAccountsRequest, opts ...grpc.CallOption) (*GetFinancialAccountsResponse, error)
	ToggleFinancialAccount(ctx context.Context, in *ToggleFinancialAccountRequest, opts ...grpc.CallOption) (*ToggleFinancialAccountResponse, error)
	GetFinancialTransactions(ctx context.Context, in *GetFinancialTransactionsRequest, opts ...grpc.CallOption) (*GetFinancialTransactionsResponse, error)
}

type userFinancesServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserFinancesServiceClient(cc *grpc.ClientConn) UserFinancesServiceClient {
	return &userFinancesServiceClient{cc}
}

func (c *userFinancesServiceClient) GetFinancialInstitutions(ctx context.Context, in *GetFinancialInstitutionsRequest, opts ...grpc.CallOption) (*GetFinancialInstitutionsResponse, error) {
	out := new(GetFinancialInstitutionsResponse)
	err := c.cc.Invoke(ctx, "/userfinances.UserFinancesService/GetFinancialInstitutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFinancesServiceClient) GetFinancialAccounts(ctx context.Context, in *GetFinancialAccountsRequest, opts ...grpc.CallOption) (*GetFinancialAccountsResponse, error) {
	out := new(GetFinancialAccountsResponse)
	err := c.cc.Invoke(ctx, "/userfinances.UserFinancesService/GetFinancialAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFinancesServiceClient) ToggleFinancialAccount(ctx context.Context, in *ToggleFinancialAccountRequest, opts ...grpc.CallOption) (*ToggleFinancialAccountResponse, error) {
	out := new(ToggleFinancialAccountResponse)
	err := c.cc.Invoke(ctx, "/userfinances.UserFinancesService/ToggleFinancialAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFinancesServiceClient) GetFinancialTransactions(ctx context.Context, in *GetFinancialTransactionsRequest, opts ...grpc.CallOption) (*GetFinancialTransactionsResponse, error) {
	out := new(GetFinancialTransactionsResponse)
	err := c.cc.Invoke(ctx, "/userfinances.UserFinancesService/GetFinancialTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserFinancesServiceServer is the server API for UserFinancesService service.
type UserFinancesServiceServer interface {
	GetFinancialInstitutions(context.Context, *GetFinancialInstitutionsRequest) (*GetFinancialInstitutionsResponse, error)
	GetFinancialAccounts(context.Context, *GetFinancialAccountsRequest) (*GetFinancialAccountsResponse, error)
	ToggleFinancialAccount(context.Context, *ToggleFinancialAccountRequest) (*ToggleFinancialAccountResponse, error)
	GetFinancialTransactions(context.Context, *GetFinancialTransactionsRequest) (*GetFinancialTransactionsResponse, error)
}

// UnimplementedUserFinancesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserFinancesServiceServer struct {
}

func (*UnimplementedUserFinancesServiceServer) GetFinancialInstitutions(ctx context.Context, req *GetFinancialInstitutionsRequest) (*GetFinancialInstitutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinancialInstitutions not implemented")
}
func (*UnimplementedUserFinancesServiceServer) GetFinancialAccounts(ctx context.Context, req *GetFinancialAccountsRequest) (*GetFinancialAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinancialAccounts not implemented")
}
func (*UnimplementedUserFinancesServiceServer) ToggleFinancialAccount(ctx context.Context, req *ToggleFinancialAccountRequest) (*ToggleFinancialAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleFinancialAccount not implemented")
}
func (*UnimplementedUserFinancesServiceServer) GetFinancialTransactions(ctx context.Context, req *GetFinancialTransactionsRequest) (*GetFinancialTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFinancialTransactions not implemented")
}

func RegisterUserFinancesServiceServer(s *grpc.Server, srv UserFinancesServiceServer) {
	s.RegisterService(&_UserFinancesService_serviceDesc, srv)
}

func _UserFinancesService_GetFinancialInstitutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFinancialInstitutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFinancesServiceServer).GetFinancialInstitutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userfinances.UserFinancesService/GetFinancialInstitutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFinancesServiceServer).GetFinancialInstitutions(ctx, req.(*GetFinancialInstitutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFinancesService_GetFinancialAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFinancialAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFinancesServiceServer).GetFinancialAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userfinances.UserFinancesService/GetFinancialAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFinancesServiceServer).GetFinancialAccounts(ctx, req.(*GetFinancialAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFinancesService_ToggleFinancialAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleFinancialAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFinancesServiceServer).ToggleFinancialAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userfinances.UserFinancesService/ToggleFinancialAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFinancesServiceServer).ToggleFinancialAccount(ctx, req.(*ToggleFinancialAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFinancesService_GetFinancialTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFinancialTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFinancesServiceServer).GetFinancialTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userfinances.UserFinancesService/GetFinancialTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFinancesServiceServer).GetFinancialTransactions(ctx, req.(*GetFinancialTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserFinancesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userfinances.UserFinancesService",
	HandlerType: (*UserFinancesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFinancialInstitutions",
			Handler:    _UserFinancesService_GetFinancialInstitutions_Handler,
		},
		{
			MethodName: "GetFinancialAccounts",
			Handler:    _UserFinancesService_GetFinancialAccounts_Handler,
		},
		{
			MethodName: "ToggleFinancialAccount",
			Handler:    _UserFinancesService_ToggleFinancialAccount_Handler,
		},
		{
			MethodName: "GetFinancialTransactions",
			Handler:    _UserFinancesService_GetFinancialTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/userfinances/userFinances.proto",
}
