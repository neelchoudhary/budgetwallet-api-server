// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/userfinances/userFinances.proto

package userfinances

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	shared "github.com/neelchoudhary/budgetmanagergrpc/services/shared"
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
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
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

func (m *GetFinancialInstitutionsRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

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
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId               int64    `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
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

func (m *GetFinancialAccountsRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

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
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId               int64    `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	AccountId            int64    `protobuf:"varint,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Selected             bool     `protobuf:"varint,4,opt,name=selected,proto3" json:"selected,omitempty"`
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

func (m *ToggleFinancialAccountRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

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

type GetTransactionsRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId               int64    `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionsRequest) Reset()         { *m = GetTransactionsRequest{} }
func (m *GetTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionsRequest) ProtoMessage()    {}
func (*GetTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec379254392880eb, []int{6}
}

func (m *GetTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionsRequest.Unmarshal(m, b)
}
func (m *GetTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionsRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionsRequest.Merge(m, src)
}
func (m *GetTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionsRequest.Size(m)
}
func (m *GetTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionsRequest proto.InternalMessageInfo

func (m *GetTransactionsRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetTransactionsRequest) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

type GetTransactionsResponse struct {
	Transactions         []*shared.Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetTransactionsResponse) Reset()         { *m = GetTransactionsResponse{} }
func (m *GetTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTransactionsResponse) ProtoMessage()    {}
func (*GetTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec379254392880eb, []int{7}
}

func (m *GetTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionsResponse.Unmarshal(m, b)
}
func (m *GetTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *GetTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionsResponse.Merge(m, src)
}
func (m *GetTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTransactionsResponse.Size(m)
}
func (m *GetTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionsResponse proto.InternalMessageInfo

func (m *GetTransactionsResponse) GetTransactions() []*shared.Transaction {
	if m != nil {
		return m.Transactions
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
	proto.RegisterType((*GetTransactionsRequest)(nil), "userfinances.GetTransactionsRequest")
	proto.RegisterType((*GetTransactionsResponse)(nil), "userfinances.GetTransactionsResponse")
}

func init() {
	proto.RegisterFile("services/userfinances/userFinances.proto", fileDescriptor_ec379254392880eb)
}

var fileDescriptor_ec379254392880eb = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x04, 0xb5, 0x61, 0xa8, 0x84, 0xd8, 0x42, 0x6a, 0x99, 0x14, 0x22, 0x0b, 0xa4, 0xf0,
	0xe5, 0x48, 0xe5, 0x80, 0xd4, 0x13, 0xf4, 0x40, 0x14, 0x2e, 0x48, 0x6e, 0xb9, 0x70, 0x29, 0x9b,
	0xdd, 0x89, 0xb3, 0x52, 0xb2, 0x1b, 0xf6, 0x03, 0x84, 0xf8, 0x0b, 0xfc, 0x44, 0x7e, 0x0c, 0x72,
	0xbc, 0x0e, 0x4e, 0x62, 0x93, 0x88, 0x93, 0x35, 0x9e, 0x37, 0x6f, 0xde, 0xec, 0xbe, 0x59, 0xe8,
	0x1b, 0xd4, 0xdf, 0x04, 0x43, 0x33, 0x70, 0x06, 0xf5, 0x44, 0x48, 0x2a, 0xcb, 0xe0, 0xbd, 0x0f,
	0x92, 0x85, 0x56, 0x56, 0x91, 0xa3, 0x2a, 0x20, 0xea, 0xae, 0xea, 0xcc, 0x94, 0x6a, 0xe4, 0xfe,
	0x53, 0x60, 0xe3, 0x73, 0x78, 0x3c, 0x44, 0x5b, 0x10, 0x08, 0x3a, 0x1b, 0x49, 0x63, 0x85, 0x75,
	0x56, 0x28, 0x69, 0x52, 0xfc, 0xea, 0xd0, 0x58, 0x72, 0x02, 0x87, 0x39, 0xe1, 0xb5, 0xe0, 0x61,
	0xd0, 0x0b, 0xfa, 0xad, 0xf4, 0x20, 0x0f, 0x47, 0x3c, 0xfe, 0x0e, 0xbd, 0xe6, 0x5a, 0xb3, 0x50,
	0xd2, 0x20, 0xb9, 0x84, 0xce, 0xa4, 0x04, 0x5c, 0x8b, 0x0a, 0x22, 0x0c, 0x7a, 0xad, 0xfe, 0x9d,
	0xb3, 0x6e, 0xe2, 0xe5, 0xd4, 0xd1, 0xa4, 0x0f, 0x26, 0x75, 0xe4, 0xf1, 0x47, 0x78, 0x58, 0x6d,
	0xfc, 0x8e, 0x31, 0xe5, 0xa4, 0xdd, 0x29, 0x38, 0x4f, 0x08, 0x8b, 0xf3, 0x3c, 0x71, 0xb3, 0x48,
	0xe4, 0xe1, 0x88, 0xc7, 0x19, 0x74, 0xeb, 0x09, 0xfd, 0x14, 0x43, 0x20, 0x7f, 0xa7, 0xa0, 0x3e,
	0xeb, 0x27, 0x08, 0xb7, 0x26, 0xf0, 0xe5, 0xe9, 0xbd, 0xc9, 0x26, 0x61, 0xfc, 0x2b, 0x80, 0xd3,
	0x2b, 0x95, 0x65, 0x33, 0xdc, 0x42, 0xff, 0xaf, 0x78, 0x72, 0x0a, 0xe0, 0x25, 0xe5, 0xb9, 0xd6,
	0x32, 0x77, 0xdb, 0xff, 0x19, 0x71, 0x12, 0x41, 0xdb, 0xe0, 0x0c, 0x99, 0x45, 0x1e, 0xde, 0xea,
	0x05, 0xfd, 0x76, 0xba, 0x8a, 0xe3, 0x73, 0x78, 0xd4, 0xa4, 0xc6, 0x4f, 0x1e, 0xc2, 0xa1, 0x71,
	0x8c, 0xa1, 0x31, 0x4b, 0x39, 0xed, 0xb4, 0x0c, 0xe3, 0x0f, 0xd0, 0x19, 0xa2, 0xbd, 0xd2, 0x54,
	0x1a, 0xca, 0xf6, 0x32, 0x4c, 0xf3, 0xf9, 0xa7, 0x70, 0xb2, 0xc5, 0xe5, 0x05, 0xbc, 0x81, 0x23,
	0x5b, 0xf9, 0xef, 0x0f, 0xfd, 0xb8, 0x3c, 0xf4, 0x4a, 0x4d, 0xba, 0x06, 0x3c, 0xfb, 0xdd, 0x82,
	0xe3, 0x4f, 0x95, 0xe5, 0xb8, 0x2c, 0xd6, 0x80, 0xfc, 0x84, 0xb0, 0xc9, 0xb5, 0xe4, 0x55, 0x52,
	0x5d, 0x9d, 0x64, 0xc7, 0x66, 0x44, 0xc9, 0xbe, 0xf0, 0x62, 0x96, 0xf8, 0x06, 0x51, 0x70, 0xbf,
	0xce, 0x68, 0xe4, 0x59, 0x33, 0xd3, 0x86, 0xbb, 0xa3, 0xe7, 0xfb, 0x40, 0x57, 0x0d, 0x1d, 0x74,
	0xea, 0x6f, 0x98, 0xbc, 0x58, 0xe7, 0xf9, 0xa7, 0x2b, 0xa3, 0x97, 0xfb, 0x81, 0x57, 0x6d, 0xbf,
	0xc0, 0xdd, 0x8d, 0x0b, 0x25, 0x4f, 0xb6, 0x74, 0xd7, 0x78, 0x27, 0x7a, 0xba, 0x03, 0x55, 0x76,
	0xb8, 0xb8, 0xf8, 0xfc, 0x36, 0x13, 0x76, 0xea, 0xc6, 0x09, 0x53, 0xf3, 0x81, 0x44, 0x9c, 0xb1,
	0xa9, 0x72, 0x7c, 0x4a, 0xf5, 0x8f, 0xc1, 0xd8, 0xf1, 0x0c, 0xed, 0x9c, 0x4a, 0x9a, 0xa1, 0xce,
	0xf4, 0x82, 0x0d, 0x6a, 0xdf, 0xce, 0xf1, 0xc1, 0xf2, 0x0d, 0x7c, 0xfd, 0x27, 0x00, 0x00, 0xff,
	0xff, 0xdc, 0x03, 0x35, 0x23, 0x5b, 0x05, 0x00, 0x00,
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
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
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

func (c *userFinancesServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, "/userfinances.UserFinancesService/GetTransactions", in, out, opts...)
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
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
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
func (*UnimplementedUserFinancesServiceServer) GetTransactions(ctx context.Context, req *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
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

func _UserFinancesService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFinancesServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userfinances.UserFinancesService/GetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFinancesServiceServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
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
			MethodName: "GetTransactions",
			Handler:    _UserFinancesService_GetTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/userfinances/userFinances.proto",
}
