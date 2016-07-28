// Code generated by protoc-gen-go.
// source: browser.proto
// DO NOT EDIT!

/*
Package transport is a generated protocol buffer package.

It is generated from these files:
	browser.proto

It has these top-level messages:
	BrowserTerminator
*/
package transport

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

// BrowserTerminator is the last message in a DATA stream of HTTP2, or the last chunk of the HTTP1.1 response.
// It's because browsers cannot access the trailing headers of a response, and as such it is crucial to add gRPC
// status information at the end of the data stream.
// The last chunk is prefixed with [0xDEADBEEF][4 bytes length of BrowserTerminator] and the Browser terminator proto.
type BrowserTerminator struct {
	StatusCode int32                               `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	StatusDesc string                              `protobuf:"bytes,2,opt,name=status_desc,json=statusDesc" json:"status_desc,omitempty"`
	Trailer    []*BrowserTerminator_TrailingHeader `protobuf:"bytes,3,rep,name=trailer" json:"trailer,omitempty"`
}

func (m *BrowserTerminator) Reset()                    { *m = BrowserTerminator{} }
func (m *BrowserTerminator) String() string            { return proto.CompactTextString(m) }
func (*BrowserTerminator) ProtoMessage()               {}
func (*BrowserTerminator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BrowserTerminator) GetTrailer() []*BrowserTerminator_TrailingHeader {
	if m != nil {
		return m.Trailer
	}
	return nil
}

type BrowserTerminator_TrailingHeader struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *BrowserTerminator_TrailingHeader) Reset()         { *m = BrowserTerminator_TrailingHeader{} }
func (m *BrowserTerminator_TrailingHeader) String() string { return proto.CompactTextString(m) }
func (*BrowserTerminator_TrailingHeader) ProtoMessage()    {}
func (*BrowserTerminator_TrailingHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func init() {
	proto.RegisterType((*BrowserTerminator)(nil), "grpc.yolo.BrowserTerminator")
	proto.RegisterType((*BrowserTerminator_TrailingHeader)(nil), "grpc.yolo.BrowserTerminator.TrailingHeader")
}

func init() { proto.RegisterFile("browser.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0xb5, 0xca, 0x4e, 0x51, 0x30, 0x78, 0x58, 0xbc, 0x18, 0x7a, 0x5a, 0x10, 0x52,
	0xd0, 0x9b, 0xc7, 0xaa, 0xe0, 0x79, 0xe9, 0xc9, 0x8b, 0x4c, 0x93, 0x21, 0x2c, 0xa4, 0x99, 0x65,
	0x92, 0x2a, 0xbe, 0xa9, 0x8f, 0x23, 0x6d, 0xac, 0x28, 0xde, 0x92, 0xf9, 0xbf, 0xe1, 0xff, 0x18,
	0x38, 0x5b, 0x0b, 0xbf, 0x67, 0x12, 0x3b, 0x0a, 0x17, 0xd6, 0x4d, 0x90, 0xd1, 0xd9, 0x0f, 0x8e,
	0x3c, 0xff, 0x54, 0x70, 0xb1, 0xac, 0xe1, 0x8a, 0x64, 0x33, 0x24, 0x2c, 0x2c, 0xfa, 0x1a, 0x66,
	0xb9, 0x60, 0xd9, 0xe6, 0x57, 0xc7, 0x9e, 0x5a, 0x65, 0x54, 0x37, 0xed, 0xa1, 0x8e, 0x1e, 0xd8,
	0xd3, 0x2f, 0xc0, 0x53, 0x76, 0xed, 0x91, 0x51, 0x5d, 0x73, 0x00, 0x1e, 0x29, 0x3b, 0xfd, 0x04,
	0xa7, 0x45, 0x70, 0x88, 0x24, 0xed, 0xc4, 0x4c, 0xba, 0xd9, 0xed, 0x8d, 0xfd, 0x29, 0xb5, 0xff,
	0x0a, 0xed, 0x6a, 0xc7, 0x0e, 0x29, 0x3c, 0x13, 0x7a, 0x92, 0xfe, 0xb0, 0x7b, 0x75, 0x0f, 0xe7,
	0x7f, 0x23, 0xad, 0xe1, 0x38, 0xe1, 0xa6, 0x3a, 0x35, 0xfd, 0xfe, 0xad, 0x2f, 0x61, 0xfa, 0x86,
	0x71, 0x4b, 0xdf, 0x1e, 0xf5, 0xb3, 0x9c, 0xbf, 0x98, 0xc0, 0x1c, 0x22, 0xd9, 0xc0, 0x11, 0x53,
	0xb0, 0x2c, 0x61, 0xb1, 0x93, 0x58, 0x14, 0xc1, 0x94, 0x47, 0x96, 0xb2, 0x3e, 0xd9, 0x1f, 0xe4,
	0xee, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x81, 0x8e, 0x89, 0x21, 0x01, 0x00, 0x00,
}
