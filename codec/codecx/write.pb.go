// Code generated by protoc-gen-go.
// source: write.proto
// DO NOT EDIT!

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	write.proto

It has these top-level messages:
	RequestHeader
	ResponseHeader
*/
package codecx

import (
	proto "github.com/golang/protobuf/proto"
	"fmt"
	"math"
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

type RequestHeader struct {
	Method string `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Seq    uint64 `protobuf:"varint,2,opt,name=seq" json:"seq,omitempty"`
	Header string `protobuf:"bytes,3,opt,name=header" json:"header,omitempty"`
}

func (m *RequestHeader) Reset()                    { *m = RequestHeader{} }
func (m *RequestHeader) String() string            { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()               {}
func (*RequestHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ResponseHeader struct {
	Method string `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Seq    uint64 `protobuf:"varint,2,opt,name=seq" json:"seq,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	Header string `protobuf:"bytes,4,opt,name=header" json:"header,omitempty"`
}

func (m *ResponseHeader) Reset()                    { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()               {}
func (*ResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*RequestHeader)(nil), "core.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "core.ResponseHeader")
}

func init() { proto.RegisterFile("write.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2f, 0xca, 0x2c,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0xce, 0x2f, 0x4a, 0x55, 0x0a, 0xe4,
	0xe2, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xf1, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x12, 0x12,
	0xe3, 0x62, 0xcb, 0x4d, 0x2d, 0xc9, 0xc8, 0x4f, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
	0xf2, 0x84, 0x04, 0xb8, 0x98, 0x8b, 0x53, 0x0b, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x40,
	0x4c, 0x90, 0xca, 0x0c, 0xb0, 0x1e, 0x09, 0x66, 0x88, 0x4a, 0x08, 0x4f, 0x29, 0x83, 0x8b, 0x2f,
	0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x95, 0x64, 0x33, 0x45, 0xb8, 0x58, 0x53, 0x8b, 0x8a,
	0xf2, 0x61, 0x46, 0x42, 0x38, 0x48, 0x36, 0xb1, 0x20, 0xdb, 0x94, 0xc4, 0x06, 0xf6, 0x89, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x14, 0xb6, 0x6a, 0x9a, 0xd8, 0x00, 0x00, 0x00,
}
