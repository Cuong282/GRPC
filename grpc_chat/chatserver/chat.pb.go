// source: chat.proto

package chatserver // import "./chatserver"

import (
	fmt "fmt"

	math "math"

	"github.com/golang/protobuf/proto"
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

type FromClient struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FromClient) Reset()         { *m = FromClient{} }
func (m *FromClient) String() string { return proto.CompactTextString(m) }
func (*FromClient) ProtoMessage()    {}
func (*FromClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_3d91036648cfa160, []int{0}
}
func (m *FromClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FromClient.Unmarshal(m, b)
}
func (m *FromClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FromClient.Marshal(b, m, deterministic)
}
func (dst *FromClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FromClient.Merge(dst, src)
}
func (m *FromClient) XXX_Size() int {
	return xxx_messageInfo_FromClient.Size(m)
}
func (m *FromClient) XXX_DiscardUnknown() {
	xxx_messageInfo_FromClient.DiscardUnknown(m)
}

var xxx_messageInfo_FromClient proto.InternalMessageInfo

func (m *FromClient) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FromClient) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type FromServer struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FromServer) Reset()         { *m = FromServer{} }
func (m *FromServer) String() string { return proto.CompactTextString(m) }
func (*FromServer) ProtoMessage()    {}
func (*FromServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_3d91036648cfa160, []int{1}
}
func (m *FromServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FromServer.Unmarshal(m, b)
}
func (m *FromServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FromServer.Marshal(b, m, deterministic)
}
func (dst *FromServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FromServer.Merge(dst, src)
}
func (m *FromServer) XXX_Size() int {
	return xxx_messageInfo_FromServer.Size(m)
}
func (m *FromServer) XXX_DiscardUnknown() {
	xxx_messageInfo_FromServer.DiscardUnknown(m)
}

var xxx_messageInfo_FromServer proto.InternalMessageInfo

func (m *FromServer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FromServer) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*FromClient)(nil), "chatserver.FromClient")
	proto.RegisterType((*FromServer)(nil), "chatserver.FromServer")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_chat_3d91036648cfa160) }

var fileDescriptor_chat_3d91036648cfa160 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x02, 0xb3, 0x8b, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0x94,
	0x4c, 0xb8, 0xb8, 0xdc, 0x8a, 0xf2, 0x73, 0x9d, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0x84, 0x84, 0xb8,
	0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58,
	0x52, 0x7e, 0x4a, 0xa5, 0x04, 0x13, 0x44, 0x0c, 0xc4, 0x86, 0xe9, 0x0a, 0x06, 0x9b, 0x41, 0xac,
	0x2e, 0x23, 0x7f, 0x2e, 0x0e, 0x90, 0x8e, 0xcc, 0xe4, 0xd4, 0x62, 0x21, 0x67, 0x2e, 0x6e, 0xe7,
	0x8c, 0xc4, 0x12, 0x28, 0x5f, 0x48, 0x4c, 0x0f, 0xe1, 0x26, 0x3d, 0x84, 0x83, 0xa4, 0x30, 0xc4,
	0x21, 0x56, 0x2a, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x3a, 0xf1, 0x45, 0xf1, 0xe8, 0xe9, 0x23, 0x14,
	0x24, 0xb1, 0x81, 0xfd, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x24, 0x2a, 0xd8, 0xed,
	0x00, 0x00, 0x00,
}
