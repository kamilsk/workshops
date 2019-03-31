// Code generated by protoc-gen-go. DO NOT EDIT.
// source: maintenance.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StatusResponse struct {
	Increment            *StatusResponse_IncrementStatus    `protobuf:"bytes,1,opt,name=increment,proto3" json:"increment,omitempty"`
	Fibonacci            *StatusResponse_FibonacciStatus    `protobuf:"bytes,2,opt,name=fibonacci,proto3" json:"fibonacci,omitempty"`
	Unique               *StatusResponse_UniqueStringStatus `protobuf:"bytes,3,opt,name=unique,proto3" json:"unique,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{0}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetIncrement() *StatusResponse_IncrementStatus {
	if m != nil {
		return m.Increment
	}
	return nil
}

func (m *StatusResponse) GetFibonacci() *StatusResponse_FibonacciStatus {
	if m != nil {
		return m.Fibonacci
	}
	return nil
}

func (m *StatusResponse) GetUnique() *StatusResponse_UniqueStringStatus {
	if m != nil {
		return m.Unique
	}
	return nil
}

type StatusResponse_IncrementStatus struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                int64                `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StatusResponse_IncrementStatus) Reset()         { *m = StatusResponse_IncrementStatus{} }
func (m *StatusResponse_IncrementStatus) String() string { return proto.CompactTextString(m) }
func (*StatusResponse_IncrementStatus) ProtoMessage()    {}
func (*StatusResponse_IncrementStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{0, 0}
}

func (m *StatusResponse_IncrementStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse_IncrementStatus.Unmarshal(m, b)
}
func (m *StatusResponse_IncrementStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse_IncrementStatus.Marshal(b, m, deterministic)
}
func (m *StatusResponse_IncrementStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse_IncrementStatus.Merge(m, src)
}
func (m *StatusResponse_IncrementStatus) XXX_Size() int {
	return xxx_messageInfo_StatusResponse_IncrementStatus.Size(m)
}
func (m *StatusResponse_IncrementStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse_IncrementStatus.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse_IncrementStatus proto.InternalMessageInfo

func (m *StatusResponse_IncrementStatus) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StatusResponse_IncrementStatus) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *StatusResponse_IncrementStatus) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type StatusResponse_FibonacciStatus struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                int64                `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StatusResponse_FibonacciStatus) Reset()         { *m = StatusResponse_FibonacciStatus{} }
func (m *StatusResponse_FibonacciStatus) String() string { return proto.CompactTextString(m) }
func (*StatusResponse_FibonacciStatus) ProtoMessage()    {}
func (*StatusResponse_FibonacciStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{0, 1}
}

func (m *StatusResponse_FibonacciStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse_FibonacciStatus.Unmarshal(m, b)
}
func (m *StatusResponse_FibonacciStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse_FibonacciStatus.Marshal(b, m, deterministic)
}
func (m *StatusResponse_FibonacciStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse_FibonacciStatus.Merge(m, src)
}
func (m *StatusResponse_FibonacciStatus) XXX_Size() int {
	return xxx_messageInfo_StatusResponse_FibonacciStatus.Size(m)
}
func (m *StatusResponse_FibonacciStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse_FibonacciStatus.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse_FibonacciStatus proto.InternalMessageInfo

func (m *StatusResponse_FibonacciStatus) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StatusResponse_FibonacciStatus) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *StatusResponse_FibonacciStatus) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type StatusResponse_UniqueStringStatus struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                string               `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StatusResponse_UniqueStringStatus) Reset()         { *m = StatusResponse_UniqueStringStatus{} }
func (m *StatusResponse_UniqueStringStatus) String() string { return proto.CompactTextString(m) }
func (*StatusResponse_UniqueStringStatus) ProtoMessage()    {}
func (*StatusResponse_UniqueStringStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6053ae89a3b3f561, []int{0, 2}
}

func (m *StatusResponse_UniqueStringStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse_UniqueStringStatus.Unmarshal(m, b)
}
func (m *StatusResponse_UniqueStringStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse_UniqueStringStatus.Marshal(b, m, deterministic)
}
func (m *StatusResponse_UniqueStringStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse_UniqueStringStatus.Merge(m, src)
}
func (m *StatusResponse_UniqueStringStatus) XXX_Size() int {
	return xxx_messageInfo_StatusResponse_UniqueStringStatus.Size(m)
}
func (m *StatusResponse_UniqueStringStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse_UniqueStringStatus.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse_UniqueStringStatus proto.InternalMessageInfo

func (m *StatusResponse_UniqueStringStatus) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StatusResponse_UniqueStringStatus) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *StatusResponse_UniqueStringStatus) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusResponse)(nil), "protobuf.StatusResponse")
	proto.RegisterType((*StatusResponse_IncrementStatus)(nil), "protobuf.StatusResponse.IncrementStatus")
	proto.RegisterType((*StatusResponse_FibonacciStatus)(nil), "protobuf.StatusResponse.FibonacciStatus")
	proto.RegisterType((*StatusResponse_UniqueStringStatus)(nil), "protobuf.StatusResponse.UniqueStringStatus")
}

func init() { proto.RegisterFile("maintenance.proto", fileDescriptor_6053ae89a3b3f561) }

var fileDescriptor_6053ae89a3b3f561 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x69, 0xfa, 0x6f, 0xf8, 0x77, 0x0b, 0x95, 0x2e, 0x22, 0x21, 0x0a, 0x4a, 0x4f, 0x45,
	0x21, 0x81, 0x7a, 0xd2, 0x9b, 0x88, 0x05, 0x0f, 0xbd, 0xa4, 0x7a, 0x13, 0xca, 0x36, 0xd9, 0x96,
	0x81, 0x66, 0x36, 0x26, 0xb3, 0x82, 0x57, 0x5f, 0xc1, 0x57, 0xf2, 0x0d, 0x7c, 0x05, 0x1f, 0x44,
	0x9a, 0xec, 0x5a, 0x4d, 0xa9, 0x1e, 0xc4, 0x53, 0x48, 0xe6, 0x37, 0xdf, 0xf7, 0xcd, 0x64, 0x58,
	0x2f, 0x15, 0x80, 0x24, 0x51, 0x60, 0x2c, 0x83, 0x2c, 0x57, 0xa4, 0xf8, 0xff, 0xf2, 0x31, 0xd3,
	0x73, 0xff, 0x60, 0xa1, 0xd4, 0x62, 0x29, 0x43, 0x91, 0x41, 0x28, 0x10, 0x15, 0x09, 0x02, 0x85,
	0x45, 0xc5, 0xf9, 0xfb, 0xa6, 0x6a, 0xf1, 0x50, 0xa6, 0x19, 0x3d, 0x9a, 0xe2, 0x61, 0xbd, 0x48,
	0x90, 0xca, 0x82, 0x44, 0x9a, 0x55, 0x40, 0xff, 0xe5, 0x1f, 0xeb, 0x4e, 0x48, 0x90, 0x2e, 0x22,
	0x59, 0x64, 0x0a, 0x0b, 0xc9, 0x47, 0xac, 0x0d, 0x18, 0xe7, 0x32, 0x95, 0x48, 0x5e, 0xe3, 0xa8,
	0x31, 0xe8, 0x0c, 0x07, 0x81, 0x15, 0x08, 0xbe, 0xc2, 0xc1, 0xb5, 0x25, 0xcd, 0xf7, 0x75, 0xeb,
	0x4a, 0x67, 0x0e, 0x33, 0x85, 0x22, 0x8e, 0xc1, 0x73, 0x7e, 0xd0, 0x19, 0x59, 0xd2, 0xea, 0x7c,
	0xb4, 0xf2, 0x4b, 0xe6, 0x6a, 0x84, 0x7b, 0x2d, 0xbd, 0x66, 0x29, 0x72, 0xb2, 0x55, 0xe4, 0xb6,
	0xc4, 0x26, 0x94, 0x03, 0x2e, 0x4c, 0xc9, 0xb4, 0xfa, 0x39, 0xdb, 0xa9, 0x45, 0xe5, 0x5d, 0xe6,
	0x40, 0x52, 0x0e, 0xd8, 0x8c, 0x1c, 0x48, 0xf8, 0x2e, 0x6b, 0x3d, 0x88, 0xa5, 0x96, 0x65, 0xd6,
	0x66, 0x54, 0xbd, 0xf0, 0x33, 0xc6, 0x74, 0x96, 0x08, 0x92, 0xc9, 0x54, 0x90, 0x49, 0xe0, 0x07,
	0xd5, 0x5a, 0xd7, 0x41, 0x6e, 0xec, 0x5a, 0xa3, 0xb6, 0xa1, 0x2f, 0x68, 0xe5, 0x59, 0x1b, 0xeb,
	0x93, 0x67, 0xeb, 0x6f, 0x3c, 0x35, 0xe3, 0x9b, 0x5b, 0xf8, 0x7e, 0xd4, 0xf6, 0xef, 0x6d, 0x87,
	0x77, 0xac, 0x33, 0x5e, 0x5f, 0x30, 0x1f, 0x33, 0xd7, 0x38, 0xef, 0x6d, 0xf4, 0x5f, 0xad, 0xce,
	0xd3, 0xf7, 0xb6, 0xfd, 0xc4, 0x7e, 0xef, 0xe9, 0xf5, 0xed, 0xd9, 0xe9, 0xf4, 0xdd, 0x70, 0x0a,
	0x38, 0x57, 0xe7, 0x8d, 0xe3, 0x99, 0x5b, 0xb2, 0xa7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x09,
	0xd9, 0x8e, 0xc4, 0x26, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MaintenanceClient is the client API for Maintenance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaintenanceClient interface {
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
}

type maintenanceClient struct {
	cc *grpc.ClientConn
}

func NewMaintenanceClient(cc *grpc.ClientConn) MaintenanceClient {
	return &maintenanceClient{cc}
}

func (c *maintenanceClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Maintenance/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MaintenanceServer is the server API for Maintenance service.
type MaintenanceServer interface {
	Status(context.Context, *empty.Empty) (*StatusResponse, error)
}

func RegisterMaintenanceServer(s *grpc.Server, srv MaintenanceServer) {
	s.RegisterService(&_Maintenance_serviceDesc, srv)
}

func _Maintenance_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Maintenance/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Maintenance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Maintenance",
	HandlerType: (*MaintenanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Maintenance_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "maintenance.proto",
}
