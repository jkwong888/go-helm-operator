// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/billing_setup_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for
// [BillingSetupService.GetBillingSetup][google.ads.googleads.v1.services.BillingSetupService.GetBillingSetup].
type GetBillingSetupRequest struct {
	// The resource name of the billing setup to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBillingSetupRequest) Reset()         { *m = GetBillingSetupRequest{} }
func (m *GetBillingSetupRequest) String() string { return proto.CompactTextString(m) }
func (*GetBillingSetupRequest) ProtoMessage()    {}
func (*GetBillingSetupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_billing_setup_service_6bff980c1ed14b83, []int{0}
}
func (m *GetBillingSetupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBillingSetupRequest.Unmarshal(m, b)
}
func (m *GetBillingSetupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBillingSetupRequest.Marshal(b, m, deterministic)
}
func (dst *GetBillingSetupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBillingSetupRequest.Merge(dst, src)
}
func (m *GetBillingSetupRequest) XXX_Size() int {
	return xxx_messageInfo_GetBillingSetupRequest.Size(m)
}
func (m *GetBillingSetupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBillingSetupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBillingSetupRequest proto.InternalMessageInfo

func (m *GetBillingSetupRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for billing setup mutate operations.
type MutateBillingSetupRequest struct {
	// Id of the customer to apply the billing setup mutate operation to.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform.
	Operation            *BillingSetupOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MutateBillingSetupRequest) Reset()         { *m = MutateBillingSetupRequest{} }
func (m *MutateBillingSetupRequest) String() string { return proto.CompactTextString(m) }
func (*MutateBillingSetupRequest) ProtoMessage()    {}
func (*MutateBillingSetupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_billing_setup_service_6bff980c1ed14b83, []int{1}
}
func (m *MutateBillingSetupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBillingSetupRequest.Unmarshal(m, b)
}
func (m *MutateBillingSetupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBillingSetupRequest.Marshal(b, m, deterministic)
}
func (dst *MutateBillingSetupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBillingSetupRequest.Merge(dst, src)
}
func (m *MutateBillingSetupRequest) XXX_Size() int {
	return xxx_messageInfo_MutateBillingSetupRequest.Size(m)
}
func (m *MutateBillingSetupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBillingSetupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBillingSetupRequest proto.InternalMessageInfo

func (m *MutateBillingSetupRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateBillingSetupRequest) GetOperation() *BillingSetupOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// A single operation on a billing setup, which describes the cancellation of an
// existing billing setup.
type BillingSetupOperation struct {
	// Only one of these operations can be set. "Update" operations are not
	// supported.
	//
	// Types that are valid to be assigned to Operation:
	//	*BillingSetupOperation_Create
	//	*BillingSetupOperation_Remove
	Operation            isBillingSetupOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *BillingSetupOperation) Reset()         { *m = BillingSetupOperation{} }
func (m *BillingSetupOperation) String() string { return proto.CompactTextString(m) }
func (*BillingSetupOperation) ProtoMessage()    {}
func (*BillingSetupOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_billing_setup_service_6bff980c1ed14b83, []int{2}
}
func (m *BillingSetupOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetupOperation.Unmarshal(m, b)
}
func (m *BillingSetupOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetupOperation.Marshal(b, m, deterministic)
}
func (dst *BillingSetupOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetupOperation.Merge(dst, src)
}
func (m *BillingSetupOperation) XXX_Size() int {
	return xxx_messageInfo_BillingSetupOperation.Size(m)
}
func (m *BillingSetupOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetupOperation.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetupOperation proto.InternalMessageInfo

type isBillingSetupOperation_Operation interface {
	isBillingSetupOperation_Operation()
}

type BillingSetupOperation_Create struct {
	Create *resources.BillingSetup `protobuf:"bytes,2,opt,name=create,proto3,oneof"`
}

type BillingSetupOperation_Remove struct {
	Remove string `protobuf:"bytes,1,opt,name=remove,proto3,oneof"`
}

func (*BillingSetupOperation_Create) isBillingSetupOperation_Operation() {}

func (*BillingSetupOperation_Remove) isBillingSetupOperation_Operation() {}

func (m *BillingSetupOperation) GetOperation() isBillingSetupOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *BillingSetupOperation) GetCreate() *resources.BillingSetup {
	if x, ok := m.GetOperation().(*BillingSetupOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *BillingSetupOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*BillingSetupOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BillingSetupOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BillingSetupOperation_OneofMarshaler, _BillingSetupOperation_OneofUnmarshaler, _BillingSetupOperation_OneofSizer, []interface{}{
		(*BillingSetupOperation_Create)(nil),
		(*BillingSetupOperation_Remove)(nil),
	}
}

func _BillingSetupOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BillingSetupOperation)
	// operation
	switch x := m.Operation.(type) {
	case *BillingSetupOperation_Create:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *BillingSetupOperation_Remove:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("BillingSetupOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _BillingSetupOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BillingSetupOperation)
	switch tag {
	case 2: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.BillingSetup)
		err := b.DecodeMessage(msg)
		m.Operation = &BillingSetupOperation_Create{msg}
		return true, err
	case 1: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &BillingSetupOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _BillingSetupOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BillingSetupOperation)
	// operation
	switch x := m.Operation.(type) {
	case *BillingSetupOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BillingSetupOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for a billing setup operation.
type MutateBillingSetupResponse struct {
	// A result that identifies the resource affected by the mutate request.
	Result               *MutateBillingSetupResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MutateBillingSetupResponse) Reset()         { *m = MutateBillingSetupResponse{} }
func (m *MutateBillingSetupResponse) String() string { return proto.CompactTextString(m) }
func (*MutateBillingSetupResponse) ProtoMessage()    {}
func (*MutateBillingSetupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_billing_setup_service_6bff980c1ed14b83, []int{3}
}
func (m *MutateBillingSetupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBillingSetupResponse.Unmarshal(m, b)
}
func (m *MutateBillingSetupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBillingSetupResponse.Marshal(b, m, deterministic)
}
func (dst *MutateBillingSetupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBillingSetupResponse.Merge(dst, src)
}
func (m *MutateBillingSetupResponse) XXX_Size() int {
	return xxx_messageInfo_MutateBillingSetupResponse.Size(m)
}
func (m *MutateBillingSetupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBillingSetupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBillingSetupResponse proto.InternalMessageInfo

func (m *MutateBillingSetupResponse) GetResult() *MutateBillingSetupResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// Result for a single billing setup mutate.
type MutateBillingSetupResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateBillingSetupResult) Reset()         { *m = MutateBillingSetupResult{} }
func (m *MutateBillingSetupResult) String() string { return proto.CompactTextString(m) }
func (*MutateBillingSetupResult) ProtoMessage()    {}
func (*MutateBillingSetupResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_billing_setup_service_6bff980c1ed14b83, []int{4}
}
func (m *MutateBillingSetupResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateBillingSetupResult.Unmarshal(m, b)
}
func (m *MutateBillingSetupResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateBillingSetupResult.Marshal(b, m, deterministic)
}
func (dst *MutateBillingSetupResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateBillingSetupResult.Merge(dst, src)
}
func (m *MutateBillingSetupResult) XXX_Size() int {
	return xxx_messageInfo_MutateBillingSetupResult.Size(m)
}
func (m *MutateBillingSetupResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateBillingSetupResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateBillingSetupResult proto.InternalMessageInfo

func (m *MutateBillingSetupResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBillingSetupRequest)(nil), "google.ads.googleads.v1.services.GetBillingSetupRequest")
	proto.RegisterType((*MutateBillingSetupRequest)(nil), "google.ads.googleads.v1.services.MutateBillingSetupRequest")
	proto.RegisterType((*BillingSetupOperation)(nil), "google.ads.googleads.v1.services.BillingSetupOperation")
	proto.RegisterType((*MutateBillingSetupResponse)(nil), "google.ads.googleads.v1.services.MutateBillingSetupResponse")
	proto.RegisterType((*MutateBillingSetupResult)(nil), "google.ads.googleads.v1.services.MutateBillingSetupResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BillingSetupServiceClient is the client API for BillingSetupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BillingSetupServiceClient interface {
	// Returns a billing setup.
	GetBillingSetup(ctx context.Context, in *GetBillingSetupRequest, opts ...grpc.CallOption) (*resources.BillingSetup, error)
	// Creates a billing setup, or cancels an existing billing setup.
	MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error)
}

type billingSetupServiceClient struct {
	cc *grpc.ClientConn
}

func NewBillingSetupServiceClient(cc *grpc.ClientConn) BillingSetupServiceClient {
	return &billingSetupServiceClient{cc}
}

func (c *billingSetupServiceClient) GetBillingSetup(ctx context.Context, in *GetBillingSetupRequest, opts ...grpc.CallOption) (*resources.BillingSetup, error) {
	out := new(resources.BillingSetup)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.BillingSetupService/GetBillingSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingSetupServiceClient) MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error) {
	out := new(MutateBillingSetupResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.BillingSetupService/MutateBillingSetup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingSetupServiceServer is the server API for BillingSetupService service.
type BillingSetupServiceServer interface {
	// Returns a billing setup.
	GetBillingSetup(context.Context, *GetBillingSetupRequest) (*resources.BillingSetup, error)
	// Creates a billing setup, or cancels an existing billing setup.
	MutateBillingSetup(context.Context, *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error)
}

func RegisterBillingSetupServiceServer(s *grpc.Server, srv BillingSetupServiceServer) {
	s.RegisterService(&_BillingSetupService_serviceDesc, srv)
}

func _BillingSetupService_GetBillingSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillingSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingSetupServiceServer).GetBillingSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.BillingSetupService/GetBillingSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingSetupServiceServer).GetBillingSetup(ctx, req.(*GetBillingSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillingSetupService_MutateBillingSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBillingSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.BillingSetupService/MutateBillingSetup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, req.(*MutateBillingSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BillingSetupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.BillingSetupService",
	HandlerType: (*BillingSetupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBillingSetup",
			Handler:    _BillingSetupService_GetBillingSetup_Handler,
		},
		{
			MethodName: "MutateBillingSetup",
			Handler:    _BillingSetupService_MutateBillingSetup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/billing_setup_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/billing_setup_service.proto", fileDescriptor_billing_setup_service_6bff980c1ed14b83)
}

var fileDescriptor_billing_setup_service_6bff980c1ed14b83 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xe6, 0x52, 0x29, 0x52, 0x2f, 0x20, 0xa4, 0x43, 0xa0, 0x10, 0x21, 0x11, 0x99, 0x0e, 0x55,
	0x86, 0x3b, 0x39, 0x08, 0x05, 0x5d, 0x1b, 0x21, 0x67, 0x49, 0x3b, 0x00, 0x95, 0x2b, 0x32, 0xa0,
	0x48, 0x91, 0x1b, 0x9f, 0x2c, 0x4b, 0xb6, 0xcf, 0xf8, 0x9d, 0xb3, 0x54, 0x5d, 0xd8, 0x98, 0xbb,
	0x33, 0x30, 0xb2, 0xf3, 0x27, 0x18, 0x58, 0xf8, 0x0b, 0x4c, 0xfc, 0x09, 0x90, 0x7d, 0xbe, 0x34,
	0xa1, 0x8e, 0x42, 0xbb, 0x3d, 0x9f, 0xdf, 0xf7, 0x7d, 0xef, 0x7d, 0xef, 0xdd, 0xe1, 0xc3, 0x40,
	0xca, 0x20, 0x12, 0xcc, 0xf3, 0x81, 0xe9, 0xb0, 0x88, 0x16, 0x36, 0x03, 0x91, 0x2d, 0xc2, 0xb9,
	0x00, 0x76, 0x16, 0x46, 0x51, 0x98, 0x04, 0x33, 0x10, 0x2a, 0x4f, 0x67, 0xd5, 0x31, 0x4d, 0x33,
	0xa9, 0x24, 0xe9, 0x6a, 0x08, 0xf5, 0x7c, 0xa0, 0x4b, 0x34, 0x5d, 0xd8, 0xd4, 0xa0, 0x3b, 0x2f,
	0x36, 0xf1, 0x67, 0x02, 0x64, 0x9e, 0x5d, 0x13, 0xd0, 0xc4, 0x9d, 0x27, 0x06, 0x96, 0x86, 0xcc,
	0x4b, 0x12, 0xa9, 0x3c, 0x15, 0xca, 0x04, 0xf4, 0x5f, 0x6b, 0x88, 0x1f, 0x8d, 0x85, 0x1a, 0x69,
	0xdc, 0x69, 0x01, 0x73, 0xc5, 0x87, 0x5c, 0x80, 0x22, 0xcf, 0xf0, 0x3d, 0x43, 0x3c, 0x4b, 0xbc,
	0x58, 0xb4, 0x51, 0x17, 0xed, 0xef, 0xba, 0x77, 0xcd, 0xe1, 0x1b, 0x2f, 0x16, 0xd6, 0x25, 0xc2,
	0x8f, 0x5f, 0xe7, 0xca, 0x53, 0xa2, 0x8e, 0xe2, 0x29, 0x6e, 0xcd, 0x73, 0x50, 0x32, 0x16, 0xd9,
	0x2c, 0xf4, 0x2b, 0x02, 0x6c, 0x8e, 0x8e, 0x7d, 0xf2, 0x0e, 0xef, 0xca, 0x54, 0x64, 0x65, 0x45,
	0xed, 0x46, 0x17, 0xed, 0xb7, 0xfa, 0x03, 0xba, 0xcd, 0x08, 0xba, 0x2a, 0xf5, 0xd6, 0xc0, 0xdd,
	0x2b, 0x26, 0xeb, 0x13, 0xc2, 0x0f, 0x6b, 0x93, 0xc8, 0x31, 0x6e, 0xce, 0x33, 0xe1, 0x29, 0x51,
	0xa9, 0xb1, 0x8d, 0x6a, 0x4b, 0x53, 0xd7, 0xe4, 0x8e, 0xee, 0xb8, 0x15, 0x01, 0x69, 0xe3, 0x66,
	0x26, 0x62, 0xb9, 0xa8, 0x8c, 0x29, 0xfe, 0xe8, 0xef, 0x51, 0x6b, 0xa5, 0x2b, 0x2b, 0xc5, 0x9d,
	0x3a, 0x83, 0x20, 0x95, 0x09, 0x08, 0xe2, 0x16, 0x24, 0x90, 0x47, 0xaa, 0x24, 0x69, 0xf5, 0xf9,
	0xf6, 0xee, 0x6b, 0xd9, 0xf2, 0x48, 0xb9, 0x15, 0x93, 0xf5, 0x0a, 0xb7, 0x37, 0xe5, 0xfc, 0xd7,
	0x50, 0xfb, 0x9f, 0x77, 0xf0, 0x83, 0x55, 0xec, 0xa9, 0x96, 0x26, 0xdf, 0x10, 0xbe, 0xff, 0xcf,
	0xb2, 0x90, 0x97, 0xdb, 0x0b, 0xae, 0xdf, 0xaf, 0xce, 0x4d, 0xad, 0xb7, 0x06, 0x1f, 0x7f, 0xfe,
	0xba, 0x6c, 0xd8, 0x84, 0x15, 0x3b, 0x7f, 0xbe, 0xd6, 0xc6, 0xd0, 0xec, 0x14, 0xb0, 0x9e, 0xb9,
	0x04, 0x25, 0x08, 0x58, 0xef, 0x82, 0xfc, 0x40, 0x98, 0x5c, 0x77, 0x84, 0x1c, 0xdc, 0xce, 0x6b,
	0x5d, 0xfd, 0xe1, 0x2d, 0x07, 0x55, 0x8e, 0xdd, 0x1a, 0x96, 0xad, 0x0c, 0xac, 0x7e, 0xd1, 0xca,
	0x55, 0xed, 0xe7, 0x2b, 0xb7, 0x65, 0xd8, 0xbb, 0x58, 0xef, 0x84, 0xc7, 0x25, 0x1f, 0x47, 0xbd,
	0xd1, 0x1f, 0x84, 0xf7, 0xe6, 0x32, 0xde, 0x5a, 0xc2, 0xa8, 0x5d, 0x33, 0xc6, 0x93, 0xe2, 0xde,
	0x9f, 0xa0, 0xf7, 0x47, 0x15, 0x3a, 0x90, 0x91, 0x97, 0x04, 0x54, 0x66, 0x01, 0x0b, 0x44, 0x52,
	0xbe, 0x0a, 0xe6, 0x79, 0x49, 0x43, 0xd8, 0xfc, 0x9a, 0x1d, 0x98, 0xe0, 0x4b, 0x63, 0x67, 0xec,
	0x38, 0x5f, 0x1b, 0xdd, 0xb1, 0x26, 0x74, 0x7c, 0xa0, 0x3a, 0x2c, 0xa2, 0x89, 0x4d, 0x2b, 0x61,
	0xf8, 0x6e, 0x52, 0xa6, 0x8e, 0x0f, 0xd3, 0x65, 0xca, 0x74, 0x62, 0x4f, 0x4d, 0xca, 0xef, 0xc6,
	0x9e, 0x3e, 0xe7, 0xdc, 0xf1, 0x81, 0xf3, 0x65, 0x12, 0xe7, 0x13, 0x9b, 0x73, 0x93, 0x76, 0xd6,
	0x2c, 0xeb, 0x7c, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x00, 0x54, 0x93, 0x74, 0x05, 0x00,
	0x00,
}
