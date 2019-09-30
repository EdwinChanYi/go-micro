// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro/go-micro/client/proto/client.proto

package go_micro_client

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Endpoint             string   `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d733ae29171347b, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Request) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d733ae29171347b, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Message struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	ContentType          string   `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d733ae29171347b, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Message) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.client.Request")
	proto.RegisterType((*Response)(nil), "go.micro.client.Response")
	proto.RegisterType((*Message)(nil), "go.micro.client.Message")
}

func init() {
	proto.RegisterFile("micro/go-micro/client/proto/client.proto", fileDescriptor_7d733ae29171347b)
}

var fileDescriptor_7d733ae29171347b = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xbb, 0x6d, 0x4c, 0xea, 0x58, 0x10, 0x06, 0x0f, 0x6b, 0x0e, 0x52, 0x73, 0xca, 0xc5,
	0x54, 0xf4, 0x2c, 0x1e, 0x72, 0x16, 0x24, 0x8a, 0x57, 0x49, 0xb6, 0x43, 0x5c, 0x48, 0x77, 0xd7,
	0xec, 0xb6, 0x90, 0x1f, 0xe9, 0x7f, 0x12, 0x36, 0xa9, 0x15, 0x6d, 0x2f, 0xbd, 0xcd, 0x9b, 0x6f,
	0x79, 0x33, 0xfb, 0x06, 0xd2, 0x95, 0x14, 0xad, 0x5e, 0xd4, 0xfa, 0xa6, 0x2f, 0x44, 0x23, 0x49,
	0xb9, 0x85, 0x69, 0xb5, 0xdb, 0x8a, 0xcc, 0x0b, 0x3c, 0xaf, 0x75, 0xe6, 0xdf, 0x64, 0x7d, 0x3b,
	0xd9, 0x40, 0x54, 0xd0, 0xe7, 0x9a, 0xac, 0x43, 0x0e, 0x91, 0xa5, 0x76, 0x23, 0x05, 0x71, 0x36,
	0x67, 0xe9, 0x69, 0xb1, 0x95, 0x18, 0xc3, 0x94, 0xd4, 0xd2, 0x68, 0xa9, 0x1c, 0x1f, 0x7b, 0xf4,
	0xa3, 0xf1, 0x1a, 0x66, 0x42, 0x2b, 0x47, 0xca, 0xbd, 0xbb, 0xce, 0x10, 0x9f, 0x78, 0x7e, 0x36,
	0xf4, 0x5e, 0x3b, 0x43, 0x88, 0x10, 0x54, 0x7a, 0xd9, 0xf1, 0x60, 0xce, 0xd2, 0x59, 0xe1, 0xeb,
	0xe4, 0x0a, 0xa6, 0x05, 0x59, 0xa3, 0x95, 0xdd, 0x71, 0xf6, 0x8b, 0xbf, 0x41, 0xf4, 0x44, 0xd6,
	0x96, 0x35, 0xe1, 0x05, 0x9c, 0x38, 0x6d, 0xa4, 0x18, 0xb6, 0xea, 0xc5, 0xbf, 0xb9, 0xe3, 0xc3,
	0x73, 0x27, 0x3b, 0xdf, 0xbb, 0x2f, 0x06, 0x61, 0xee, 0xbf, 0x8e, 0x0f, 0x10, 0xe4, 0x65, 0xd3,
	0x20, 0xcf, 0xfe, 0x84, 0x92, 0x0d, 0x89, 0xc4, 0x97, 0x7b, 0x48, 0xbf, 0x73, 0x32, 0xc2, 0x1c,
	0xc2, 0x17, 0xd7, 0x52, 0xb9, 0x3a, 0xd2, 0x20, 0x65, 0xb7, 0x0c, 0x1f, 0x21, 0x7a, 0x5e, 0x57,
	0x8d, 0xb4, 0x1f, 0x7b, 0x5c, 0x86, 0x00, 0xe2, 0x83, 0x24, 0x19, 0x55, 0xa1, 0xbf, 0xeb, 0xfd,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x76, 0x1f, 0x51, 0x03, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientClient is the client API for Client service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientClient interface {
	// Call allows a single request to be made
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// Stream is a bidirectional stream
	Stream(ctx context.Context, opts ...grpc.CallOption) (Client_StreamClient, error)
	// Publish publishes a message and returns an empty Message
	Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type clientClient struct {
	cc *grpc.ClientConn
}

func NewClientClient(cc *grpc.ClientConn) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.client.Client/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Client_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Client_serviceDesc.Streams[0], "/go.micro.client.Client/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStreamClient{stream}
	return x, nil
}

type Client_StreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type clientStreamClient struct {
	grpc.ClientStream
}

func (x *clientStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientClient) Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/go.micro.client.Client/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServer is the server API for Client service.
type ClientServer interface {
	// Call allows a single request to be made
	Call(context.Context, *Request) (*Response, error)
	// Stream is a bidirectional stream
	Stream(Client_StreamServer) error
	// Publish publishes a message and returns an empty Message
	Publish(context.Context, *Message) (*Message, error)
}

func RegisterClientServer(s *grpc.Server, srv ClientServer) {
	s.RegisterService(&_Client_serviceDesc, srv)
}

func _Client_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.client.Client/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientServer).Stream(&clientStreamServer{stream})
}

type Client_StreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type clientStreamServer struct {
	grpc.ServerStream
}

func (x *clientStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Client_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.client.Client/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Publish(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Client_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.client.Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Client_Call_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Client_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Client_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "micro/go-micro/client/proto/client.proto",
}
