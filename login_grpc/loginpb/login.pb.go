// Code generated by protoc-gen-go. DO NOT EDIT.
// source: loginpb/login.proto

package loginpb // import "./loginpb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_login_b002f8a98383133e, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
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

type Credentials struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_login_b002f8a98383133e, []int{1}
}
func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credentials.Unmarshal(m, b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
}
func (dst *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(dst, src)
}
func (m *Credentials) XXX_Size() int {
	return xxx_messageInfo_Credentials.Size(m)
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ChangeUer struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeUer) Reset()         { *m = ChangeUer{} }
func (m *ChangeUer) String() string { return proto.CompactTextString(m) }
func (*ChangeUer) ProtoMessage()    {}
func (*ChangeUer) Descriptor() ([]byte, []int) {
	return fileDescriptor_login_b002f8a98383133e, []int{2}
}
func (m *ChangeUer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeUer.Unmarshal(m, b)
}
func (m *ChangeUer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeUer.Marshal(b, m, deterministic)
}
func (dst *ChangeUer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeUer.Merge(dst, src)
}
func (m *ChangeUer) XXX_Size() int {
	return xxx_messageInfo_ChangeUer.Size(m)
}
func (m *ChangeUer) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeUer.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeUer proto.InternalMessageInfo

func (m *ChangeUer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChangeUer) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ChangeUer) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ChangeUer) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_login_b002f8a98383133e, []int{3}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "login_grpc.User")
	proto.RegisterType((*Credentials)(nil), "login_grpc.Credentials")
	proto.RegisterType((*ChangeUer)(nil), "login_grpc.ChangeUer")
	proto.RegisterType((*Token)(nil), "login_grpc.Token")
}

func init() { proto.RegisterFile("loginpb/login.proto", fileDescriptor_login_b002f8a98383133e) }

var fileDescriptor_login_b002f8a98383133e = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xc9, 0x4f, 0xcf,
	0xcc, 0x2b, 0x48, 0xd2, 0x07, 0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x60, 0x4e,
	0x7c, 0x7a, 0x51, 0x41, 0xb2, 0x52, 0x08, 0x17, 0x4b, 0x68, 0x71, 0x6a, 0x91, 0x90, 0x14, 0x17,
	0x47, 0x69, 0x71, 0x6a, 0x51, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x9c, 0x2f, 0x24, 0xc2, 0xc5, 0x9a, 0x9a, 0x9b, 0x98, 0x99, 0x23, 0xc1, 0x04, 0x96, 0x80, 0x70,
	0x40, 0x3a, 0x0a, 0x12, 0x8b, 0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24, 0x98, 0x21, 0x3a, 0x60, 0x7c,
	0x25, 0x57, 0x2e, 0x6e, 0xe7, 0xa2, 0xd4, 0x94, 0xd4, 0xbc, 0x92, 0xcc, 0xc4, 0x9c, 0x62, 0xbc,
	0x86, 0x23, 0x1b, 0xc3, 0x84, 0x66, 0x4c, 0x26, 0x17, 0xa7, 0x73, 0x46, 0x62, 0x5e, 0x7a, 0x6a,
	0x68, 0x6a, 0x91, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x0a, 0x58, 0x3b, 0x6b, 0x10, 0x53, 0x66, 0x0a,
	0x8a, 0xa1, 0x4c, 0x78, 0x0c, 0x45, 0x73, 0x1b, 0xc8, 0x37, 0x21, 0xf9, 0xd9, 0xa9, 0x79, 0x12,
	0x2c, 0x10, 0xdf, 0x80, 0x39, 0x4a, 0xb2, 0x50, 0x51, 0x90, 0x74, 0x09, 0x58, 0x1a, 0xe2, 0x50,
	0x08, 0xc7, 0x68, 0x27, 0x23, 0x17, 0x8f, 0x0f, 0x28, 0xd4, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93,
	0x53, 0x85, 0x4c, 0xb8, 0xd8, 0x82, 0x33, 0xd3, 0xf3, 0x42, 0x0b, 0x84, 0x44, 0xf5, 0x10, 0xc1,
	0xa9, 0x07, 0x77, 0xae, 0x94, 0x20, 0xb2, 0x30, 0xc4, 0x0e, 0x06, 0x21, 0x63, 0x2e, 0x56, 0xb0,
	0x29, 0x24, 0x69, 0xb2, 0xe3, 0x12, 0x80, 0xab, 0x08, 0x28, 0xca, 0x4f, 0xcb, 0xcc, 0x49, 0x25,
	0x45, 0xbf, 0x13, 0x77, 0x14, 0xa7, 0x9e, 0x3e, 0x34, 0x1d, 0x24, 0xb1, 0x81, 0x93, 0x80, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x48, 0x38, 0xdb, 0x6a, 0x19, 0x02, 0x00, 0x00,
}
