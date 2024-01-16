// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vault_cred.proto

package vaultcredpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetCredRequest struct {
	CredentialType       string   `protobuf:"bytes,1,opt,name=credentialType,proto3" json:"credentialType,omitempty"`
	CredEntityName       string   `protobuf:"bytes,2,opt,name=credEntityName,proto3" json:"credEntityName,omitempty"`
	CredIdentifier       string   `protobuf:"bytes,3,opt,name=credIdentifier,proto3" json:"credIdentifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCredRequest) Reset()         { *m = GetCredRequest{} }
func (m *GetCredRequest) String() string { return proto.CompactTextString(m) }
func (*GetCredRequest) ProtoMessage()    {}
func (*GetCredRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfee588ca0905ccc, []int{0}
}

func (m *GetCredRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCredRequest.Unmarshal(m, b)
}
func (m *GetCredRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCredRequest.Marshal(b, m, deterministic)
}
func (m *GetCredRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCredRequest.Merge(m, src)
}
func (m *GetCredRequest) XXX_Size() int {
	return xxx_messageInfo_GetCredRequest.Size(m)
}
func (m *GetCredRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCredRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCredRequest proto.InternalMessageInfo

func (m *GetCredRequest) GetCredentialType() string {
	if m != nil {
		return m.CredentialType
	}
	return ""
}

func (m *GetCredRequest) GetCredEntityName() string {
	if m != nil {
		return m.CredEntityName
	}
	return ""
}

func (m *GetCredRequest) GetCredIdentifier() string {
	if m != nil {
		return m.CredIdentifier
	}
	return ""
}

type GetCredResponse struct {
	//service-cred credential, for example: "userName": "iam-root", "password:: "hello"
	//client-cert credential, for example: "clientId": "intelops-user", "ca.crt": "...", "client.crt": "...", "client.key": "..."
	Credential           map[string]string `protobuf:"bytes,1,rep,name=credential,proto3" json:"credential,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetCredResponse) Reset()         { *m = GetCredResponse{} }
func (m *GetCredResponse) String() string { return proto.CompactTextString(m) }
func (*GetCredResponse) ProtoMessage()    {}
func (*GetCredResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfee588ca0905ccc, []int{1}
}

func (m *GetCredResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCredResponse.Unmarshal(m, b)
}
func (m *GetCredResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCredResponse.Marshal(b, m, deterministic)
}
func (m *GetCredResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCredResponse.Merge(m, src)
}
func (m *GetCredResponse) XXX_Size() int {
	return xxx_messageInfo_GetCredResponse.Size(m)
}
func (m *GetCredResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCredResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCredResponse proto.InternalMessageInfo

func (m *GetCredResponse) GetCredential() map[string]string {
	if m != nil {
		return m.Credential
	}
	return nil
}

type PutCredRequest struct {
	CredentialType string `protobuf:"bytes,1,opt,name=credentialType,proto3" json:"credentialType,omitempty"`
	CredEntityName string `protobuf:"bytes,2,opt,name=credEntityName,proto3" json:"credEntityName,omitempty"`
	CredIdentifier string `protobuf:"bytes,3,opt,name=credIdentifier,proto3" json:"credIdentifier,omitempty"`
	//service-cred credential, for example: "userName": "iam-root", "password:: "hello"
	//client-cert credential, for example: "clientId": "intelops-user", "ca.crt": "...", "client.crt": "...", "client.key": "..."
	Credential           map[string]string `protobuf:"bytes,6,rep,name=credential,proto3" json:"credential,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PutCredRequest) Reset()         { *m = PutCredRequest{} }
func (m *PutCredRequest) String() string { return proto.CompactTextString(m) }
func (*PutCredRequest) ProtoMessage()    {}
func (*PutCredRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfee588ca0905ccc, []int{2}
}

func (m *PutCredRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutCredRequest.Unmarshal(m, b)
}
func (m *PutCredRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutCredRequest.Marshal(b, m, deterministic)
}
func (m *PutCredRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutCredRequest.Merge(m, src)
}
func (m *PutCredRequest) XXX_Size() int {
	return xxx_messageInfo_PutCredRequest.Size(m)
}
func (m *PutCredRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutCredRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutCredRequest proto.InternalMessageInfo

func (m *PutCredRequest) GetCredentialType() string {
	if m != nil {
		return m.CredentialType
	}
	return ""
}

func (m *PutCredRequest) GetCredEntityName() string {
	if m != nil {
		return m.CredEntityName
	}
	return ""
}

func (m *PutCredRequest) GetCredIdentifier() string {
	if m != nil {
		return m.CredIdentifier
	}
	return ""
}

func (m *PutCredRequest) GetCredential() map[string]string {
	if m != nil {
		return m.Credential
	}
	return nil
}

type PutCredResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutCredResponse) Reset()         { *m = PutCredResponse{} }
func (m *PutCredResponse) String() string { return proto.CompactTextString(m) }
func (*PutCredResponse) ProtoMessage()    {}
func (*PutCredResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfee588ca0905ccc, []int{3}
}

func (m *PutCredResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutCredResponse.Unmarshal(m, b)
}
func (m *PutCredResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutCredResponse.Marshal(b, m, deterministic)
}
func (m *PutCredResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutCredResponse.Merge(m, src)
}
func (m *PutCredResponse) XXX_Size() int {
	return xxx_messageInfo_PutCredResponse.Size(m)
}
func (m *PutCredResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutCredResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutCredResponse proto.InternalMessageInfo

type DeleteCredRequest struct {
	CredentialType       string   `protobuf:"bytes,1,opt,name=credentialType,proto3" json:"credentialType,omitempty"`
	CredEntityName       string   `protobuf:"bytes,2,opt,name=credEntityName,proto3" json:"credEntityName,omitempty"`
	CredIdentifier       string   `protobuf:"bytes,3,opt,name=credIdentifier,proto3" json:"credIdentifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCredRequest) Reset()         { *m = DeleteCredRequest{} }
func (m *DeleteCredRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCredRequest) ProtoMessage()    {}
func (*DeleteCredRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfee588ca0905ccc, []int{4}
}

func (m *DeleteCredRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCredRequest.Unmarshal(m, b)
}
func (m *DeleteCredRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCredRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCredRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCredRequest.Merge(m, src)
}
func (m *DeleteCredRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCredRequest.Size(m)
}
func (m *DeleteCredRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCredRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCredRequest proto.InternalMessageInfo

func (m *DeleteCredRequest) GetCredentialType() string {
	if m != nil {
		return m.CredentialType
	}
	return ""
}

func (m *DeleteCredRequest) GetCredEntityName() string {
	if m != nil {
		return m.CredEntityName
	}
	return ""
}

func (m *DeleteCredRequest) GetCredIdentifier() string {
	if m != nil {
		return m.CredIdentifier
	}
	return ""
}

type DeleteCredResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCredResponse) Reset()         { *m = DeleteCredResponse{} }
func (m *DeleteCredResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCredResponse) ProtoMessage()    {}
func (*DeleteCredResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfee588ca0905ccc, []int{5}
}

func (m *DeleteCredResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCredResponse.Unmarshal(m, b)
}
func (m *DeleteCredResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCredResponse.Marshal(b, m, deterministic)
}
func (m *DeleteCredResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCredResponse.Merge(m, src)
}
func (m *DeleteCredResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCredResponse.Size(m)
}
func (m *DeleteCredResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCredResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCredResponse proto.InternalMessageInfo

type GetAppRoleTokenRequest struct {
	AppRoleName          string   `protobuf:"bytes,1,opt,name=appRoleName,proto3" json:"appRoleName,omitempty"`
	CredentialPaths      []string `protobuf:"bytes,2,rep,name=credentialPaths,proto3" json:"credentialPaths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppRoleTokenRequest) Reset()         { *m = GetAppRoleTokenRequest{} }
func (m *GetAppRoleTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetAppRoleTokenRequest) ProtoMessage()    {}
func (*GetAppRoleTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfee588ca0905ccc, []int{6}
}

func (m *GetAppRoleTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppRoleTokenRequest.Unmarshal(m, b)
}
func (m *GetAppRoleTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppRoleTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetAppRoleTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppRoleTokenRequest.Merge(m, src)
}
func (m *GetAppRoleTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetAppRoleTokenRequest.Size(m)
}
func (m *GetAppRoleTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppRoleTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppRoleTokenRequest proto.InternalMessageInfo

func (m *GetAppRoleTokenRequest) GetAppRoleName() string {
	if m != nil {
		return m.AppRoleName
	}
	return ""
}

func (m *GetAppRoleTokenRequest) GetCredentialPaths() []string {
	if m != nil {
		return m.CredentialPaths
	}
	return nil
}

type GetAppRoleTokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppRoleTokenResponse) Reset()         { *m = GetAppRoleTokenResponse{} }
func (m *GetAppRoleTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetAppRoleTokenResponse) ProtoMessage()    {}
func (*GetAppRoleTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfee588ca0905ccc, []int{7}
}

func (m *GetAppRoleTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppRoleTokenResponse.Unmarshal(m, b)
}
func (m *GetAppRoleTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppRoleTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetAppRoleTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppRoleTokenResponse.Merge(m, src)
}
func (m *GetAppRoleTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetAppRoleTokenResponse.Size(m)
}
func (m *GetAppRoleTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppRoleTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppRoleTokenResponse proto.InternalMessageInfo

func (m *GetAppRoleTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetCredentialWithAppRoleTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	CredentialPath       string   `protobuf:"bytes,2,opt,name=credentialPath,proto3" json:"credentialPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCredentialWithAppRoleTokenRequest) Reset()         { *m = GetCredentialWithAppRoleTokenRequest{} }
func (m *GetCredentialWithAppRoleTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetCredentialWithAppRoleTokenRequest) ProtoMessage()    {}
func (*GetCredentialWithAppRoleTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfee588ca0905ccc, []int{8}
}

func (m *GetCredentialWithAppRoleTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCredentialWithAppRoleTokenRequest.Unmarshal(m, b)
}
func (m *GetCredentialWithAppRoleTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCredentialWithAppRoleTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetCredentialWithAppRoleTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCredentialWithAppRoleTokenRequest.Merge(m, src)
}
func (m *GetCredentialWithAppRoleTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetCredentialWithAppRoleTokenRequest.Size(m)
}
func (m *GetCredentialWithAppRoleTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCredentialWithAppRoleTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCredentialWithAppRoleTokenRequest proto.InternalMessageInfo

func (m *GetCredentialWithAppRoleTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetCredentialWithAppRoleTokenRequest) GetCredentialPath() string {
	if m != nil {
		return m.CredentialPath
	}
	return ""
}

type GetCredentialWithAppRoleTokenResponse struct {
	Credential           map[string]string `protobuf:"bytes,1,rep,name=credential,proto3" json:"credential,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetCredentialWithAppRoleTokenResponse) Reset()         { *m = GetCredentialWithAppRoleTokenResponse{} }
func (m *GetCredentialWithAppRoleTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetCredentialWithAppRoleTokenResponse) ProtoMessage()    {}
func (*GetCredentialWithAppRoleTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfee588ca0905ccc, []int{9}
}

func (m *GetCredentialWithAppRoleTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCredentialWithAppRoleTokenResponse.Unmarshal(m, b)
}
func (m *GetCredentialWithAppRoleTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCredentialWithAppRoleTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetCredentialWithAppRoleTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCredentialWithAppRoleTokenResponse.Merge(m, src)
}
func (m *GetCredentialWithAppRoleTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetCredentialWithAppRoleTokenResponse.Size(m)
}
func (m *GetCredentialWithAppRoleTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCredentialWithAppRoleTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCredentialWithAppRoleTokenResponse proto.InternalMessageInfo

func (m *GetCredentialWithAppRoleTokenResponse) GetCredential() map[string]string {
	if m != nil {
		return m.Credential
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCredRequest)(nil), "vaultcredpb.GetCredRequest")
	proto.RegisterType((*GetCredResponse)(nil), "vaultcredpb.GetCredResponse")
	proto.RegisterMapType((map[string]string)(nil), "vaultcredpb.GetCredResponse.CredentialEntry")
	proto.RegisterType((*PutCredRequest)(nil), "vaultcredpb.PutCredRequest")
	proto.RegisterMapType((map[string]string)(nil), "vaultcredpb.PutCredRequest.CredentialEntry")
	proto.RegisterType((*PutCredResponse)(nil), "vaultcredpb.PutCredResponse")
	proto.RegisterType((*DeleteCredRequest)(nil), "vaultcredpb.DeleteCredRequest")
	proto.RegisterType((*DeleteCredResponse)(nil), "vaultcredpb.DeleteCredResponse")
	proto.RegisterType((*GetAppRoleTokenRequest)(nil), "vaultcredpb.GetAppRoleTokenRequest")
	proto.RegisterType((*GetAppRoleTokenResponse)(nil), "vaultcredpb.GetAppRoleTokenResponse")
	proto.RegisterType((*GetCredentialWithAppRoleTokenRequest)(nil), "vaultcredpb.GetCredentialWithAppRoleTokenRequest")
	proto.RegisterType((*GetCredentialWithAppRoleTokenResponse)(nil), "vaultcredpb.GetCredentialWithAppRoleTokenResponse")
	proto.RegisterMapType((map[string]string)(nil), "vaultcredpb.GetCredentialWithAppRoleTokenResponse.CredentialEntry")
}

func init() { proto.RegisterFile("vault_cred.proto", fileDescriptor_bfee588ca0905ccc) }

var fileDescriptor_bfee588ca0905ccc = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xde, 0xd9, 0xd0, 0xca, 0x9e, 0x95, 0xdd, 0x76, 0x58, 0x34, 0xac, 0x7f, 0xcb, 0x58, 0x25,
	0xa0, 0xa4, 0xb8, 0xde, 0x88, 0xe0, 0x85, 0xad, 0xb5, 0x88, 0xa2, 0x25, 0x14, 0x05, 0x2f, 0x94,
	0xac, 0x39, 0xd2, 0xd0, 0x98, 0xc4, 0x64, 0x52, 0xc8, 0x1b, 0x14, 0xbc, 0xf5, 0xda, 0x07, 0xf2,
	0x05, 0x7c, 0x1d, 0x99, 0x64, 0x92, 0x49, 0xa6, 0x31, 0xd2, 0x8b, 0x42, 0xef, 0x36, 0xdf, 0x9c,
	0x9f, 0xef, 0x7c, 0xf3, 0x9d, 0x59, 0xd8, 0x38, 0x71, 0xb3, 0x80, 0x7f, 0xfe, 0x92, 0xa0, 0x67,
	0xc7, 0x49, 0xc4, 0x23, 0x3a, 0x2e, 0x10, 0x01, 0xc4, 0x2b, 0x76, 0x4a, 0x60, 0xb2, 0x8f, 0x7c,
	0x37, 0x41, 0xcf, 0xc1, 0xef, 0x19, 0xa6, 0x9c, 0xde, 0x87, 0x89, 0x38, 0xc4, 0x90, 0xfb, 0x6e,
	0x70, 0x98, 0xc7, 0x68, 0x92, 0x05, 0xb1, 0x46, 0x8e, 0x86, 0x56, 0x71, 0x7b, 0x21, 0xf7, 0x79,
	0xfe, 0xd6, 0xfd, 0x86, 0xe6, 0x50, 0xc5, 0x29, 0xb4, 0x8a, 0x7b, 0x55, 0xe4, 0x7e, 0xf5, 0x31,
	0x31, 0x0d, 0x15, 0xa7, 0x50, 0xf6, 0x8b, 0xc0, 0xb4, 0xa6, 0x92, 0xc6, 0x51, 0x98, 0x22, 0x7d,
	0x03, 0xa0, 0xba, 0x9a, 0x64, 0x61, 0x58, 0xe3, 0xe5, 0x43, 0xbb, 0x31, 0x80, 0xad, 0x65, 0xd8,
	0xbb, 0x75, 0xf8, 0x5e, 0xc8, 0x93, 0xdc, 0x69, 0xe4, 0xcf, 0x9f, 0xc1, 0x54, 0x3b, 0xa6, 0x1b,
	0x60, 0x1c, 0x63, 0x2e, 0x27, 0x14, 0x3f, 0xe9, 0x0c, 0xd6, 0x4e, 0xdc, 0x20, 0xab, 0xa6, 0x29,
	0x3f, 0x9e, 0x0e, 0x9f, 0x10, 0xf6, 0x73, 0x08, 0x93, 0x83, 0xec, 0x32, 0x68, 0x45, 0x5f, 0xb7,
	0x74, 0x59, 0x2f, 0x74, 0x79, 0xd0, 0xd2, 0xa5, 0x4d, 0xf4, 0x22, 0x65, 0xd9, 0x84, 0x69, 0xdd,
	0xac, 0xbc, 0x04, 0xf6, 0x83, 0xc0, 0xe6, 0x0b, 0x0c, 0x90, 0xe3, 0x65, 0x30, 0xd6, 0x0c, 0x68,
	0x93, 0x8c, 0xe4, 0xe8, 0xc1, 0xb5, 0x7d, 0xe4, 0xcf, 0xe3, 0xd8, 0x89, 0x02, 0x3c, 0x8c, 0x8e,
	0x31, 0xac, 0x78, 0x2e, 0x60, 0xec, 0x96, 0x70, 0xd1, 0xbc, 0x24, 0xd9, 0x84, 0xa8, 0x05, 0x53,
	0xc5, 0xf9, 0xc0, 0xe5, 0x47, 0xa9, 0x39, 0x5c, 0x18, 0xd6, 0xc8, 0xd1, 0x61, 0xb6, 0x0d, 0xd7,
	0xcf, 0x74, 0x91, 0xde, 0x9e, 0xc1, 0x1a, 0x17, 0x80, 0x6c, 0x50, 0x7e, 0x30, 0x0f, 0xb6, 0xa4,
	0xa5, 0xcb, 0x32, 0x1f, 0x7c, 0x7e, 0xd4, 0x45, 0xb2, 0x33, 0xbb, 0x2d, 0xb1, 0x60, 0xd0, 0x94,
	0x4e, 0xa1, 0xec, 0x37, 0x81, 0x7b, 0xff, 0x69, 0x23, 0x59, 0xae, 0x3a, 0x36, 0x70, 0xa7, 0x6b,
	0x03, 0xfb, 0xeb, 0x5c, 0xa0, 0x01, 0x97, 0x7f, 0x0c, 0x18, 0xbd, 0x17, 0x84, 0x44, 0x11, 0xfa,
	0x12, 0xae, 0x48, 0x46, 0xf4, 0x46, 0xf7, 0x4b, 0x51, 0x08, 0x38, 0xbf, 0xd9, 0xf7, 0x8c, 0xb0,
	0x81, 0xa8, 0x23, 0x6d, 0xad, 0xd5, 0x69, 0x6f, 0x96, 0x56, 0x47, 0xdf, 0x84, 0x01, 0x7d, 0x07,
	0xa0, 0xdc, 0x47, 0x6f, 0xb7, 0xa2, 0xcf, 0xec, 0xc8, 0xfc, 0xce, 0x3f, 0xcf, 0xeb, 0x82, 0x9f,
	0x8a, 0x67, 0xb2, 0x29, 0x32, 0xbd, 0xab, 0xcf, 0xd2, 0xe1, 0x98, 0xf9, 0x56, 0x7f, 0x50, 0x5d,
	0xff, 0x94, 0xc0, 0xad, 0xde, 0x3b, 0xa5, 0x8f, 0xce, 0x73, 0xff, 0x65, 0xf3, 0xe5, 0xf9, 0x2d,
	0xc3, 0x06, 0x3b, 0x93, 0x8f, 0x57, 0xb7, 0x1b, 0x79, 0xab, 0xf5, 0xe2, 0x1f, 0xec, 0xf1, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x04, 0x50, 0xfa, 0xd5, 0x06, 0x00, 0x00,
}
