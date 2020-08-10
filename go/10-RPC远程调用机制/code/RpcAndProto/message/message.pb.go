// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 订单请求参数
type OrderRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	TimeStamp            int64    `protobuf:"varint,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (m *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(m, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderRequest) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

// 订单信息
type OrderInfo struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	OrderName            string   `protobuf:"bytes,2,opt,name=OrderName,proto3" json:"OrderName,omitempty"`
	OrderStatus          string   `protobuf:"bytes,3,opt,name=OrderStatus,proto3" json:"OrderStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderInfo) Reset()         { *m = OrderInfo{} }
func (m *OrderInfo) String() string { return proto.CompactTextString(m) }
func (*OrderInfo) ProtoMessage()    {}
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *OrderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderInfo.Unmarshal(m, b)
}
func (m *OrderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderInfo.Marshal(b, m, deterministic)
}
func (m *OrderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderInfo.Merge(m, src)
}
func (m *OrderInfo) XXX_Size() int {
	return xxx_messageInfo_OrderInfo.Size(m)
}
func (m *OrderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderInfo proto.InternalMessageInfo

func (m *OrderInfo) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderInfo) GetOrderName() string {
	if m != nil {
		return m.OrderName
	}
	return ""
}

func (m *OrderInfo) GetOrderStatus() string {
	if m != nil {
		return m.OrderStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderRequest)(nil), "message.OrderRequest")
	proto.RegisterType((*OrderInfo)(nil), "message.OrderInfo")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xdc, 0xb8,
	0x78, 0xfc, 0x8b, 0x52, 0x52, 0x8b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8,
	0xd8, 0xf3, 0x41, 0x7c, 0xcf, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48,
	0x86, 0x8b, 0xb3, 0x24, 0x33, 0x37, 0x35, 0xb8, 0x24, 0x31, 0xb7, 0x40, 0x82, 0x49, 0x81, 0x51,
	0x83, 0x39, 0x08, 0x21, 0xa0, 0x94, 0xca, 0xc5, 0x09, 0x36, 0xc7, 0x33, 0x2f, 0x2d, 0x1f, 0x64,
	0x88, 0x3f, 0xaa, 0x21, 0xfe, 0x08, 0x43, 0xc0, 0x4c, 0xbf, 0xc4, 0xdc, 0x54, 0xb0, 0x21, 0x9c,
	0x41, 0x08, 0x01, 0x21, 0x05, 0x2e, 0x6e, 0x30, 0x27, 0xb8, 0x24, 0xb1, 0xa4, 0xb4, 0x58, 0x82,
	0x19, 0x2c, 0x8f, 0x2c, 0x94, 0xc4, 0x06, 0x76, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0d,
	0x11, 0xf7, 0x7a, 0xcf, 0x00, 0x00, 0x00,
}
