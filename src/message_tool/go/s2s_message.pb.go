// Code generated by protoc-gen-go.
// source: s2s_message.proto
// DO NOT EDIT!

/*
Package S2SMessage is a generated protocol buffer package.

It is generated from these files:
	s2s_message.proto

It has these top-level messages:
	CS_MsgRoute_Req
	SC_MsgRoute_Rsp
	CS_BaseMessage_Req
	SC_BaseMessage_Rsp
*/
package S2SMessage

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

type MessageID int32

const (
	MessageID_CS_Invalid     MessageID = 0
	MessageID_CS_MsgRoute    MessageID = 1
	MessageID_SC_MsgRoute    MessageID = 2
	MessageID_CS_BaseMessage MessageID = 3
	MessageID_SC_BaseMessage MessageID = 4
)

var MessageID_name = map[int32]string{
	0: "CS_Invalid",
	1: "CS_MsgRoute",
	2: "SC_MsgRoute",
	3: "CS_BaseMessage",
	4: "SC_BaseMessage",
}
var MessageID_value = map[string]int32{
	"CS_Invalid":     0,
	"CS_MsgRoute":    1,
	"SC_MsgRoute":    2,
	"CS_BaseMessage": 3,
	"SC_BaseMessage": 4,
}

func (x MessageID) String() string {
	return proto.EnumName(MessageID_name, int32(x))
}
func (MessageID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EMessageErr int32

const (
	EMessageErr_Invalid EMessageErr = 0
	EMessageErr_Success EMessageErr = 1
	EMessageErr_Fail    EMessageErr = 2
)

var EMessageErr_name = map[int32]string{
	0: "Invalid",
	1: "Success",
	2: "Fail",
}
var EMessageErr_value = map[string]int32{
	"Invalid": 0,
	"Success": 1,
	"Fail":    2,
}

func (x EMessageErr) String() string {
	return proto.EnumName(EMessageErr_name, int32(x))
}
func (EMessageErr) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// route c2s消息路由
type CS_MsgRoute_Req struct {
	Operid int32  `protobuf:"varint,1,opt,name=Operid" json:"Operid,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *CS_MsgRoute_Req) Reset()                    { *m = CS_MsgRoute_Req{} }
func (m *CS_MsgRoute_Req) String() string            { return proto.CompactTextString(m) }
func (*CS_MsgRoute_Req) ProtoMessage()               {}
func (*CS_MsgRoute_Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SC_MsgRoute_Rsp struct {
	Ret EMessageErr `protobuf:"varint,1,opt,name=Ret,enum=S2SMessage.EMessageErr" json:"Ret,omitempty"`
}

func (m *SC_MsgRoute_Rsp) Reset()                    { *m = SC_MsgRoute_Rsp{} }
func (m *SC_MsgRoute_Rsp) String() string            { return proto.CompactTextString(m) }
func (*SC_MsgRoute_Rsp) ProtoMessage()               {}
func (*SC_MsgRoute_Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 基础消息
type CS_BaseMessage_Req struct {
	Srcid int32  `protobuf:"varint,1,opt,name=Srcid" json:"Srcid,omitempty"`
	Dstid int32  `protobuf:"varint,2,opt,name=Dstid" json:"Dstid,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *CS_BaseMessage_Req) Reset()                    { *m = CS_BaseMessage_Req{} }
func (m *CS_BaseMessage_Req) String() string            { return proto.CompactTextString(m) }
func (*CS_BaseMessage_Req) ProtoMessage()               {}
func (*CS_BaseMessage_Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type SC_BaseMessage_Rsp struct {
	Ret EMessageErr `protobuf:"varint,1,opt,name=Ret,enum=S2SMessage.EMessageErr" json:"Ret,omitempty"`
}

func (m *SC_BaseMessage_Rsp) Reset()                    { *m = SC_BaseMessage_Rsp{} }
func (m *SC_BaseMessage_Rsp) String() string            { return proto.CompactTextString(m) }
func (*SC_BaseMessage_Rsp) ProtoMessage()               {}
func (*SC_BaseMessage_Rsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*CS_MsgRoute_Req)(nil), "S2SMessage.CS_MsgRoute_Req")
	proto.RegisterType((*SC_MsgRoute_Rsp)(nil), "S2SMessage.SC_MsgRoute_Rsp")
	proto.RegisterType((*CS_BaseMessage_Req)(nil), "S2SMessage.CS_BaseMessage_Req")
	proto.RegisterType((*SC_BaseMessage_Rsp)(nil), "S2SMessage.SC_BaseMessage_Rsp")
	proto.RegisterEnum("S2SMessage.MessageID", MessageID_name, MessageID_value)
	proto.RegisterEnum("S2SMessage.EMessageErr", EMessageErr_name, EMessageErr_value)
}

func init() { proto.RegisterFile("s2s_message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4f, 0xbb, 0x40,
	0x10, 0xc5, 0xbf, 0xfc, 0x6a, 0xbf, 0x0e, 0x06, 0xd6, 0x89, 0xd1, 0x1e, 0x1b, 0x4e, 0xb5, 0x07,
	0x12, 0xf1, 0xaa, 0x31, 0x11, 0x6a, 0xd2, 0x43, 0x63, 0xb2, 0xeb, 0x9d, 0xac, 0xb0, 0x69, 0x48,
	0xaa, 0x20, 0x43, 0xfd, 0xfb, 0x0d, 0x0b, 0x69, 0x97, 0xa3, 0xb7, 0x79, 0x6f, 0x77, 0xf6, 0xf3,
	0x76, 0x06, 0xae, 0x28, 0xa1, 0xfc, 0x53, 0x11, 0xc9, 0xbd, 0x8a, 0x9b, 0xb6, 0xee, 0x6a, 0x04,
	0x91, 0x88, 0xdd, 0xe0, 0x44, 0x4f, 0x10, 0xa6, 0x22, 0xdf, 0xd1, 0x9e, 0xd7, 0xc7, 0x4e, 0xe5,
	0x5c, 0x7d, 0xe3, 0x0d, 0xcc, 0xde, 0x1a, 0xd5, 0x56, 0xe5, 0xc2, 0x5a, 0x5a, 0x2b, 0x8f, 0x8f,
	0x0a, 0x11, 0xdc, 0x4c, 0x76, 0x72, 0x61, 0x2f, 0xad, 0xd5, 0x25, 0xd7, 0x75, 0xf4, 0x08, 0xa1,
	0x48, 0x8d, 0x76, 0x6a, 0xf0, 0x0e, 0x1c, 0xae, 0x3a, 0xdd, 0x1b, 0x24, 0xb7, 0xf1, 0x99, 0x15,
	0x6f, 0xc6, 0x62, 0xd3, 0xb6, 0xbc, 0xbf, 0x13, 0xbd, 0x03, 0xa6, 0x22, 0x7f, 0x91, 0xa4, 0xc6,
	0x13, 0xcd, 0xbf, 0x06, 0x4f, 0xb4, 0xc5, 0x09, 0x3f, 0x88, 0xde, 0xcd, 0xa8, 0xab, 0x4a, 0x8d,
	0xf7, 0xf8, 0x20, 0x4e, 0x99, 0x1c, 0x23, 0xd3, 0x33, 0xa0, 0x48, 0xa7, 0xaf, 0xfe, 0x29, 0xd6,
	0x5a, 0xc1, 0xc5, 0x68, 0x6d, 0x33, 0x0c, 0x00, 0x52, 0x91, 0x6f, 0xbf, 0x7e, 0xe4, 0xa1, 0x2a,
	0xd9, 0x3f, 0x0c, 0xc1, 0x37, 0x06, 0xc6, 0xac, 0xde, 0x30, 0x46, 0xc0, 0x6c, 0x44, 0x08, 0xa6,
	0xbf, 0x62, 0x4e, 0xef, 0x4d, 0x33, 0x31, 0x77, 0x7d, 0x0f, 0xbe, 0x81, 0x46, 0x1f, 0xe6, 0x67,
	0x8a, 0x0f, 0x73, 0x71, 0x2c, 0x0a, 0x45, 0xc4, 0x2c, 0xfc, 0x0f, 0xee, 0xab, 0xac, 0x0e, 0xcc,
	0xfe, 0x98, 0xe9, 0x05, 0x3e, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x04, 0x14, 0x40, 0x05, 0xd5,
	0x01, 0x00, 0x00,
}
