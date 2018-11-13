// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stellarstation/api/v1/common/common.proto

package common // import "github.com/infostellarinc/go-stellarstation/api/v1/common"

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

// A set of bits - this is not necessarily aligned to bytes (8-bits).
type Bits struct {
	// The number of bits in this message. Only the most-significant `length` bits of `payload` will
	// be used.
	Length uint32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	// The actual bits in this message. Only the most-significant `length` bits of `payload` will
	// be used.
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bits) Reset()         { *m = Bits{} }
func (m *Bits) String() string { return proto.CompactTextString(m) }
func (*Bits) ProtoMessage()    {}
func (*Bits) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_0ecda96a360ec08e, []int{0}
}
func (m *Bits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bits.Unmarshal(m, b)
}
func (m *Bits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bits.Marshal(b, m, deterministic)
}
func (dst *Bits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bits.Merge(dst, src)
}
func (m *Bits) XXX_Size() int {
	return xxx_messageInfo_Bits.Size(m)
}
func (m *Bits) XXX_DiscardUnknown() {
	xxx_messageInfo_Bits.DiscardUnknown(m)
}

var xxx_messageInfo_Bits proto.InternalMessageInfo

func (m *Bits) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Bits) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Bits)(nil), "stellarstation.api.v1.common.Bits")
}

func init() {
	proto.RegisterFile("stellarstation/api/v1/common/common.proto", fileDescriptor_common_0ecda96a360ec08e)
}

var fileDescriptor_common_0ecda96a360ec08e = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x2e, 0x49, 0xcd,
	0xc9, 0x49, 0x2c, 0x2a, 0x2e, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f,
	0x33, 0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x32, 0xa8, 0x4a, 0xf5, 0x12, 0x0b, 0x32, 0xf5, 0xca, 0x0c, 0xf5, 0x20, 0x6a, 0x94, 0x2c,
	0xb8, 0x58, 0x9c, 0x32, 0x4b, 0x8a, 0x85, 0xc4, 0xb8, 0xd8, 0x72, 0x52, 0xf3, 0xd2, 0x4b, 0x32,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0xa0, 0x3c, 0x21, 0x09, 0x2e, 0xf6, 0x82, 0xc4, 0xca,
	0x9c, 0xfc, 0xc4, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x18, 0xd7, 0x29, 0x87, 0x4b,
	0x21, 0x39, 0x3f, 0x57, 0x0f, 0x9f, 0xe9, 0x4e, 0xdc, 0xce, 0x60, 0x3a, 0x00, 0xe4, 0x90, 0x00,
	0xc6, 0x28, 0xcb, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0x90, 0xac, 0x7e, 0x66, 0x5e, 0x5a, 0x3e,
	0x54, 0x6f, 0x66, 0x5e, 0xb2, 0x7e, 0x7a, 0xbe, 0x2e, 0x3e, 0x2f, 0x25, 0xb1, 0x81, 0x3d, 0x63,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x6d, 0xae, 0x91, 0xf9, 0x00, 0x00, 0x00,
}
