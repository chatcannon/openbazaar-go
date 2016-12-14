// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
	Envelope
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Message_MessageType int32

const (
	Message_PING               Message_MessageType = 0
	Message_MESSAGE            Message_MessageType = 1
	Message_FOLLOW             Message_MessageType = 2
	Message_UNFOLLOW           Message_MessageType = 3
	Message_ORDER              Message_MessageType = 4
	Message_ORDER_REJECT       Message_MessageType = 5
	Message_ORDER_CANCEL       Message_MessageType = 6
	Message_ORDER_CONFIRMATION Message_MessageType = 7
	Message_ORDER_FULFILLMENT  Message_MessageType = 8
	Message_ORDER_COMPLETION   Message_MessageType = 9
	Message_DISPUTE_OPEN       Message_MessageType = 10
	Message_DISPUTE_CLOSE      Message_MessageType = 11
	Message_REFUND             Message_MessageType = 12
	Message_OFFLINE_ACK        Message_MessageType = 13
	Message_OFFLINE_RELAY      Message_MessageType = 14
	Message_ERROR              Message_MessageType = 500
)

var Message_MessageType_name = map[int32]string{
	0:   "PING",
	1:   "MESSAGE",
	2:   "FOLLOW",
	3:   "UNFOLLOW",
	4:   "ORDER",
	5:   "ORDER_REJECT",
	6:   "ORDER_CANCEL",
	7:   "ORDER_CONFIRMATION",
	8:   "ORDER_FULFILLMENT",
	9:   "ORDER_COMPLETION",
	10:  "DISPUTE_OPEN",
	11:  "DISPUTE_CLOSE",
	12:  "REFUND",
	13:  "OFFLINE_ACK",
	14:  "OFFLINE_RELAY",
	500: "ERROR",
}
var Message_MessageType_value = map[string]int32{
	"PING":               0,
	"MESSAGE":            1,
	"FOLLOW":             2,
	"UNFOLLOW":           3,
	"ORDER":              4,
	"ORDER_REJECT":       5,
	"ORDER_CANCEL":       6,
	"ORDER_CONFIRMATION": 7,
	"ORDER_FULFILLMENT":  8,
	"ORDER_COMPLETION":   9,
	"DISPUTE_OPEN":       10,
	"DISPUTE_CLOSE":      11,
	"REFUND":             12,
	"OFFLINE_ACK":        13,
	"OFFLINE_RELAY":      14,
	"ERROR":              500,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type Message struct {
	MessageType Message_MessageType  `protobuf:"varint,1,opt,name=messageType,enum=Message_MessageType" json:"messageType,omitempty"`
	Payload     *google_protobuf.Any `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Message) GetPayload() *google_protobuf.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Envelope struct {
	Message   *Message `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Pubkey    []byte   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Envelope) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Envelope)(nil), "Envelope")
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
}

var fileDescriptor2 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x6e, 0x9b, 0x40,
	0x10, 0xc6, 0x6b, 0x3b, 0x36, 0x78, 0xc0, 0xe9, 0x64, 0x95, 0x56, 0x6e, 0xd5, 0x43, 0xc5, 0xa9,
	0x27, 0x22, 0xa5, 0x52, 0xef, 0x08, 0x0f, 0x11, 0xed, 0xb2, 0x8b, 0x16, 0x50, 0xd5, 0x93, 0x85,
	0x15, 0x8a, 0xaa, 0xba, 0x80, 0xe2, 0xb8, 0x12, 0x4f, 0xd0, 0x07, 0xe8, 0xab, 0xf6, 0x01, 0xca,
	0xbf, 0x55, 0x72, 0x82, 0xf9, 0xbe, 0xdf, 0xec, 0xce, 0x7e, 0x03, 0x9b, 0x5f, 0xc5, 0xe9, 0x94,
	0x97, 0x85, 0xdb, 0x3c, 0xd4, 0x8f, 0xf5, 0xdb, 0x37, 0x65, 0x5d, 0x97, 0xc7, 0xe2, 0x66, 0xa8,
	0x0e, 0xe7, 0xef, 0x37, 0x79, 0xd5, 0x8e, 0x96, 0xf3, 0x67, 0x01, 0x46, 0x34, 0xc2, 0xec, 0x13,
	0x58, 0x53, 0x5f, 0xda, 0x36, 0xc5, 0x76, 0xf6, 0x7e, 0xf6, 0xe1, 0xf2, 0xf6, 0xda, 0x9d, 0x6c,
	0xfd, 0xed, 0x3d, 0xf5, 0x1c, 0x64, 0x2e, 0x18, 0x4d, 0xde, 0x1e, 0xeb, 0xfc, 0x7e, 0x3b, 0xef,
	0x7a, 0xac, 0xae, 0x67, 0xbc, 0xd0, 0xd5, 0x17, 0xba, 0x5e, 0xd5, 0x2a, 0x0d, 0x39, 0x7f, 0xe7,
	0x60, 0x3d, 0x3b, 0x8c, 0x99, 0x70, 0x11, 0x87, 0xe2, 0x0e, 0x5f, 0x30, 0xab, 0x1b, 0x86, 0x92,
	0xc4, 0xbb, 0x23, 0x9c, 0x31, 0x80, 0x55, 0x20, 0x39, 0x97, 0x5f, 0x71, 0xce, 0x6c, 0x30, 0x33,
	0x31, 0x55, 0x0b, 0xb6, 0x86, 0xa5, 0x54, 0x3b, 0x52, 0x78, 0xc1, 0x10, 0xec, 0xe1, 0x77, 0xaf,
	0xe8, 0x33, 0xf9, 0x29, 0x2e, 0x9f, 0x14, 0xdf, 0x13, 0x3e, 0x71, 0x5c, 0xb1, 0xd7, 0xc0, 0x26,
	0x45, 0x8a, 0x20, 0x54, 0x91, 0x97, 0x86, 0x52, 0xa0, 0xc1, 0x5e, 0xc1, 0xd5, 0xa8, 0x07, 0x19,
	0x0f, 0x42, 0xce, 0x23, 0x12, 0x29, 0x9a, 0xec, 0x1a, 0x50, 0xe3, 0x51, 0xcc, 0x69, 0x80, 0xd7,
	0xfd, 0xb1, 0xbb, 0x30, 0x89, 0xb3, 0x94, 0xf6, 0x32, 0x26, 0x81, 0xc0, 0xae, 0x60, 0xa3, 0x15,
	0x9f, 0xcb, 0x84, 0xd0, 0xea, 0x47, 0x56, 0x14, 0x64, 0x62, 0x87, 0x36, 0x7b, 0x09, 0x96, 0x0c,
	0x02, 0x1e, 0x0a, 0xda, 0x7b, 0xfe, 0x17, 0xdc, 0xf4, 0xbc, 0x16, 0x14, 0x71, 0xef, 0x1b, 0x5e,
	0x76, 0xfc, 0x92, 0x94, 0x92, 0x0a, 0xff, 0x2d, 0x9c, 0x7b, 0x30, 0xa9, 0xfa, 0x5d, 0x1c, 0xeb,
	0x2e, 0x11, 0x07, 0x8c, 0x29, 0xe0, 0x61, 0x0b, 0xd6, 0xad, 0xa9, 0xd3, 0x57, 0xda, 0xe8, 0x5e,
	0xb5, 0x6a, 0xce, 0x87, 0x9f, 0x45, 0x3b, 0x84, 0x6e, 0xab, 0xa9, 0x62, 0xef, 0x60, 0x7d, 0xfa,
	0x51, 0x56, 0xf9, 0xe3, 0xf9, 0xa1, 0xd8, 0x2e, 0x06, 0xeb, 0x49, 0x38, 0xac, 0x86, 0x95, 0x7c,
	0xfc, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x77, 0x79, 0xf9, 0x22, 0x02, 0x00, 0x00,
}
