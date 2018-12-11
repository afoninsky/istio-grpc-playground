// Code generated by protoc-gen-go. DO NOT EDIT.
// source: receiver.proto

package internal_receiver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Message struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_receiver_d61c53e73a74d6f0, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_receiver_d61c53e73a74d6f0, []int{1}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "internal.receiver.Message")
	proto.RegisterType((*Empty)(nil), "internal.receiver.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReceiverClient is the client API for Receiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReceiverClient interface {
	// receives messages from any producer
	Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error)
}

type receiverClient struct {
	cc *grpc.ClientConn
}

func NewReceiverClient(cc *grpc.ClientConn) ReceiverClient {
	return &receiverClient{cc}
}

func (c *receiverClient) Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/internal.receiver.Receiver/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceiverServer is the server API for Receiver service.
type ReceiverServer interface {
	// receives messages from any producer
	Publish(context.Context, *Message) (*Empty, error)
}

func RegisterReceiverServer(s *grpc.Server, srv ReceiverServer) {
	s.RegisterService(&_Receiver_serviceDesc, srv)
}

func _Receiver_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.receiver.Receiver/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverServer).Publish(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Receiver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal.receiver.Receiver",
	HandlerType: (*ReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Receiver_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "receiver.proto",
}

func init() { proto.RegisterFile("receiver.proto", fileDescriptor_receiver_d61c53e73a74d6f0) }

var fileDescriptor_receiver_d61c53e73a74d6f0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x4d, 0x4e,
	0xcd, 0x2c, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcc, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x83, 0x49, 0x28, 0xc9, 0x72, 0xb1, 0xfb, 0xa6, 0x16, 0x17, 0x27,
	0xa6, 0xa7, 0x0a, 0x09, 0x71, 0xb1, 0x94, 0xa4, 0x56, 0x94, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x81, 0xd9, 0x4a, 0xec, 0x5c, 0xac, 0xae, 0xb9, 0x05, 0x25, 0x95, 0x46, 0xbe, 0x5c, 0x1c,
	0x41, 0x50, 0x3d, 0x42, 0x8e, 0x5c, 0xec, 0x01, 0xa5, 0x49, 0x39, 0x99, 0xc5, 0x19, 0x42, 0x52,
	0x7a, 0x18, 0x46, 0xea, 0x41, 0xcd, 0x93, 0x92, 0xc0, 0x22, 0x07, 0x36, 0x4c, 0x89, 0x21, 0x89,
	0x0d, 0xec, 0x20, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x99, 0x3a, 0x0f, 0xa2, 0x00,
	0x00, 0x00,
}