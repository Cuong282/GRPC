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
	return fileDescriptor_login_a0b79c79868b5283, []int{0}
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
	return fileDescriptor_login_a0b79c79868b5283, []int{1}
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
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeUer) Reset()         { *m = ChangeUer{} }
func (m *ChangeUer) String() string { return proto.CompactTextString(m) }
func (*ChangeUer) ProtoMessage()    {}
func (*ChangeUer) Descriptor() ([]byte, []int) {
	return fileDescriptor_login_a0b79c79868b5283, []int{2}
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

func (m *ChangeUer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ChangeUer) GetPassword() string {
	if m != nil {
		return m.Password
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
	return fileDescriptor_login_a0b79c79868b5283, []int{3}
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

func init() { proto.RegisterFile("loginpb/login.proto", fileDescriptor_login_a0b79c79868b5283) }

var fileDescriptor_login_a0b79c79868b5283 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x35, 0xc5, 0x4c, 0x3d, 0xe8, 0xa8, 0x18, 0x0a, 0x82, 0xe4, 0xe4, 0x29, 0x01,
	0x45, 0xf0, 0xe2, 0xc5, 0xe2, 0xcd, 0x53, 0xdb, 0x5c, 0xbc, 0xc8, 0xb6, 0x1d, 0xb7, 0x8b, 0xe9,
	0xee, 0x32, 0x1b, 0xf5, 0x6f, 0xf9, 0x13, 0x65, 0xb7, 0xc6, 0xc6, 0x12, 0x3d, 0xed, 0xbe, 0x19,
	0xbe, 0xc7, 0xcc, 0x1b, 0x38, 0xae, 0x8c, 0x54, 0xda, 0xce, 0x8a, 0xf0, 0xe6, 0x96, 0x4d, 0x6d,
	0x10, 0x82, 0x78, 0x96, 0x6c, 0xe7, 0xd9, 0x14, 0xf6, 0x4a, 0x47, 0x8c, 0x43, 0xd8, 0x7f, 0x73,
	0xc4, 0x5a, 0xac, 0x28, 0x8d, 0x2e, 0xa2, 0xcb, 0x64, 0xfc, 0xa3, 0xf1, 0x04, 0x62, 0x5a, 0x09,
	0x55, 0xa5, 0xbd, 0xd0, 0x58, 0x0b, 0x4f, 0x58, 0xe1, 0xdc, 0x87, 0xe1, 0x45, 0xba, 0xbb, 0x26,
	0x1a, 0x9d, 0x3d, 0xc0, 0x60, 0xc4, 0xb4, 0x20, 0x5d, 0x2b, 0x51, 0xb9, 0x7f, 0xcd, 0xdb, 0x36,
	0xbd, 0x2d, 0x9b, 0x3b, 0x48, 0x46, 0x4b, 0xa1, 0x25, 0x95, 0xc4, 0x9b, 0x29, 0xa2, 0xbf, 0xa6,
	0xd8, 0xc6, 0xcf, 0x21, 0x9e, 0x9a, 0x57, 0xd2, 0x1e, 0xad, 0xfd, 0xa7, 0x41, 0x83, 0xb8, 0xfa,
	0x8c, 0xe0, 0xe0, 0xd1, 0x27, 0x31, 0x21, 0x7e, 0x57, 0x73, 0xc2, 0x02, 0xfa, 0x13, 0x25, 0x75,
	0x69, 0xf1, 0x30, 0xdf, 0x44, 0x94, 0xfb, 0x7c, 0x86, 0x47, 0xed, 0x4a, 0x70, 0xcd, 0x76, 0xf0,
	0x06, 0xe2, 0x60, 0x80, 0x67, 0xed, 0x6e, 0x6b, 0xf3, 0x6e, 0xec, 0x16, 0x60, 0x4c, 0x2f, 0x4c,
	0x6e, 0xa9, 0xb4, 0xc4, 0xd3, 0x5f, 0x6c, 0xb3, 0x6e, 0x27, 0x79, 0x3f, 0x78, 0x4a, 0xf2, 0xe2,
	0xfb, 0xa4, 0xb3, 0x7e, 0xb8, 0xe6, 0xf5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xac, 0xa8,
	0x67, 0xe4, 0x01, 0x00, 0x00,
}
