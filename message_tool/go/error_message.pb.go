// Code generated by protoc-gen-go.
// source: error_message.proto
// DO NOT EDIT!

/*
Package ErrMessage is a generated protocol buffer package.

It is generated from these files:
	error_message.proto

It has these top-level messages:
*/
package ErrMessage

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

type ErrMessageId int32

const (
	ErrMessageId_msg_err_invalid       ErrMessageId = 0
	ErrMessageId_msg_err_timeout       ErrMessageId = 1
	ErrMessageId_msg_err_disconnect    ErrMessageId = 2
	ErrMessageId_msg_err_login_faild   ErrMessageId = 3
	ErrMessageId_msg_err_not_find_user ErrMessageId = 4
)

var ErrMessageId_name = map[int32]string{
	0: "msg_err_invalid",
	1: "msg_err_timeout",
	2: "msg_err_disconnect",
	3: "msg_err_login_faild",
	4: "msg_err_not_find_user",
}
var ErrMessageId_value = map[string]int32{
	"msg_err_invalid":       0,
	"msg_err_timeout":       1,
	"msg_err_disconnect":    2,
	"msg_err_login_faild":   3,
	"msg_err_not_find_user": 4,
}

func (x ErrMessageId) String() string {
	return proto.EnumName(ErrMessageId_name, int32(x))
}
func (ErrMessageId) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("ErrMessage.ErrMessageId", ErrMessageId_name, ErrMessageId_value)
}

func init() { proto.RegisterFile("error_message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0xce, 0x41, 0x0a, 0xc2, 0x40,
	0x0c, 0x85, 0x61, 0xab, 0xe2, 0x22, 0x0b, 0x1d, 0x52, 0x54, 0xbc, 0x82, 0x0b, 0x37, 0x9e, 0xa1,
	0x4b, 0xcf, 0x10, 0x6a, 0x27, 0x2d, 0x81, 0x36, 0x91, 0xcc, 0xd4, 0x23, 0x78, 0x6e, 0x41, 0x19,
	0x71, 0xfb, 0xbd, 0xb7, 0xf8, 0xa1, 0x66, 0x77, 0x73, 0x9a, 0x38, 0xa5, 0x76, 0xe0, 0xcb, 0xc3,
	0x2d, 0x1b, 0x42, 0xe3, 0x7e, 0xfb, 0xca, 0xf9, 0x55, 0xc1, 0xb6, 0xf1, 0xdf, 0x83, 0x24, 0x62,
	0x0d, 0xbb, 0x29, 0x0d, 0xc4, 0xee, 0x24, 0xfa, 0x6c, 0x47, 0x89, 0x61, 0xf1, 0x8f, 0x59, 0x26,
	0xb6, 0x39, 0x87, 0x0a, 0x0f, 0x80, 0x05, 0xa3, 0xa4, 0xce, 0x54, 0xb9, 0xcb, 0x61, 0x89, 0x47,
	0xa8, 0x8b, 0x8f, 0x36, 0x88, 0x52, 0xdf, 0xca, 0x18, 0xc3, 0x0a, 0x4f, 0xb0, 0x2f, 0x83, 0x5a,
	0xa6, 0x5e, 0x34, 0xd2, 0x9c, 0xd8, 0xc3, 0xfa, 0xbe, 0xf9, 0xb4, 0x5d, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x4b, 0xa6, 0xcc, 0x93, 0xb2, 0x00, 0x00, 0x00,
}
