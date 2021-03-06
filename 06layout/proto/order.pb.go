// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type OrderReq struct {
	OrderId              int32    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId            int32    `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Addr                 string   `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderReq) Reset()         { *m = OrderReq{} }
func (m *OrderReq) String() string { return proto.CompactTextString(m) }
func (*OrderReq) ProtoMessage()    {}
func (*OrderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_5b6c0f4e002a3c1a, []int{0}
}
func (m *OrderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderReq.Unmarshal(m, b)
}
func (m *OrderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderReq.Marshal(b, m, deterministic)
}
func (dst *OrderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderReq.Merge(dst, src)
}
func (m *OrderReq) XXX_Size() int {
	return xxx_messageInfo_OrderReq.Size(m)
}
func (m *OrderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderReq.DiscardUnknown(m)
}

var xxx_messageInfo_OrderReq proto.InternalMessageInfo

func (m *OrderReq) GetOrderId() int32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *OrderReq) GetProductId() int32 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *OrderReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type OrderResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderResp) Reset()         { *m = OrderResp{} }
func (m *OrderResp) String() string { return proto.CompactTextString(m) }
func (*OrderResp) ProtoMessage()    {}
func (*OrderResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_5b6c0f4e002a3c1a, []int{1}
}
func (m *OrderResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderResp.Unmarshal(m, b)
}
func (m *OrderResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderResp.Marshal(b, m, deterministic)
}
func (dst *OrderResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResp.Merge(dst, src)
}
func (m *OrderResp) XXX_Size() int {
	return xxx_messageInfo_OrderResp.Size(m)
}
func (m *OrderResp) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResp.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResp proto.InternalMessageInfo

func (m *OrderResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderReq)(nil), "proto.OrderReq")
	proto.RegisterType((*OrderResp)(nil), "proto.OrderResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderClient interface {
	PostOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderResp, error)
}

type orderClient struct {
	cc *grpc.ClientConn
}

func NewOrderClient(cc *grpc.ClientConn) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) PostOrder(ctx context.Context, in *OrderReq, opts ...grpc.CallOption) (*OrderResp, error) {
	out := new(OrderResp)
	err := c.cc.Invoke(ctx, "/proto.Order/PostOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
type OrderServer interface {
	PostOrder(context.Context, *OrderReq) (*OrderResp, error)
}

func RegisterOrderServer(s *grpc.Server, srv OrderServer) {
	s.RegisterService(&_Order_serviceDesc, srv)
}

func _Order_PostOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).PostOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Order/PostOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).PostOrder(ctx, req.(*OrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Order_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostOrder",
			Handler:    _Order_PostOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_order_5b6c0f4e002a3c1a) }

var fileDescriptor_order_5b6c0f4e002a3c1a = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x49, 0xb7, 0x76, 0xeb, 0xf3, 0xb0, 0x11, 0x10, 0x46, 0x51, 0x36, 0x0b, 0xc2, 0x10,
	0xd6, 0x80, 0x82, 0x07, 0x8f, 0x82, 0x87, 0x5e, 0x54, 0x7a, 0xf2, 0x26, 0xd9, 0x52, 0x6a, 0x70,
	0xed, 0x8b, 0x49, 0xea, 0xee, 0xde, 0x3c, 0xfb, 0xa7, 0xf9, 0x07, 0x08, 0xe2, 0x1f, 0x22, 0x4d,
	0xa6, 0xd0, 0x53, 0xbe, 0x7c, 0xdf, 0xef, 0xf1, 0xde, 0x07, 0x07, 0xa8, 0x45, 0xa9, 0x33, 0xa5,
	0xd1, 0x22, 0x0d, 0xdd, 0x93, 0x1c, 0x55, 0x88, 0xd5, 0xb6, 0x64, 0x5c, 0x49, 0xc6, 0x9b, 0x06,
	0x2d, 0xb7, 0x12, 0x1b, 0xe3, 0xa1, 0xe4, 0xb2, 0x92, 0xf6, 0xa9, 0x5d, 0x67, 0x1b, 0xac, 0x59,
	0xbd, 0x93, 0xf6, 0x19, 0x77, 0xac, 0xc2, 0x95, 0x0b, 0x57, 0xaf, 0x7c, 0x2b, 0x05, 0xb7, 0xa8,
	0x0d, 0xfb, 0x97, 0x7e, 0x2e, 0x7d, 0x27, 0x30, 0xbe, 0xeb, 0x96, 0x15, 0xe5, 0x0b, 0x3d, 0x81,
	0xb1, 0x5b, 0xfc, 0x28, 0xc5, 0x8c, 0x2c, 0xc8, 0x32, 0xbc, 0x8e, 0xbe, 0xbf, 0xe6, 0xc1, 0x03,
	0x29, 0x46, 0xce, 0xcf, 0x05, 0x9d, 0xc3, 0xa8, 0x35, 0x9e, 0x08, 0x7a, 0x44, 0xd4, 0xd9, 0xb9,
	0xa0, 0xa7, 0x00, 0x4a, 0xa3, 0x68, 0x37, 0xb6, 0x63, 0x06, 0x3d, 0x26, 0xde, 0x27, 0xb9, 0xa0,
	0x14, 0x86, 0x5c, 0x08, 0x3d, 0x1b, 0x2e, 0xc8, 0x32, 0x2e, 0x9c, 0x4e, 0x8f, 0x21, 0xde, 0x9f,
	0x62, 0x14, 0x9d, 0xc2, 0xa0, 0x36, 0x95, 0x3b, 0x23, 0x2e, 0x3a, 0x79, 0x7e, 0x0b, 0xa1, 0x8b,
	0xe9, 0x0d, 0xc4, 0xf7, 0x68, 0xac, 0xff, 0x4c, 0x7c, 0x91, 0xec, 0xaf, 0x44, 0x32, 0xed, 0x1b,
	0x46, 0xa5, 0x87, 0x6f, 0x9f, 0x3f, 0x1f, 0xc1, 0x24, 0x05, 0xa6, 0xd0, 0x58, 0xd7, 0xe4, 0x8a,
	0x9c, 0xad, 0x23, 0xc7, 0x5d, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x85, 0x41, 0xf0, 0x6d,
	0x01, 0x00, 0x00,
}
