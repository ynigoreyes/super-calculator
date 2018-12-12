// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/calculator.proto

package calculator

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

type PairOfNumbers struct {
	Number1              float32  `protobuf:"fixed32,1,opt,name=number1,proto3" json:"number1,omitempty"`
	Number2              float32  `protobuf:"fixed32,2,opt,name=number2,proto3" json:"number2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairOfNumbers) Reset()         { *m = PairOfNumbers{} }
func (m *PairOfNumbers) String() string { return proto.CompactTextString(m) }
func (*PairOfNumbers) ProtoMessage()    {}
func (*PairOfNumbers) Descriptor() ([]byte, []int) {
	return fileDescriptor_3454248de0eb79e2, []int{0}
}

func (m *PairOfNumbers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairOfNumbers.Unmarshal(m, b)
}
func (m *PairOfNumbers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairOfNumbers.Marshal(b, m, deterministic)
}
func (m *PairOfNumbers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairOfNumbers.Merge(m, src)
}
func (m *PairOfNumbers) XXX_Size() int {
	return xxx_messageInfo_PairOfNumbers.Size(m)
}
func (m *PairOfNumbers) XXX_DiscardUnknown() {
	xxx_messageInfo_PairOfNumbers.DiscardUnknown(m)
}

var xxx_messageInfo_PairOfNumbers proto.InternalMessageInfo

func (m *PairOfNumbers) GetNumber1() float32 {
	if m != nil {
		return m.Number1
	}
	return 0
}

func (m *PairOfNumbers) GetNumber2() float32 {
	if m != nil {
		return m.Number2
	}
	return 0
}

type Results struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Results) Reset()         { *m = Results{} }
func (m *Results) String() string { return proto.CompactTextString(m) }
func (*Results) ProtoMessage()    {}
func (*Results) Descriptor() ([]byte, []int) {
	return fileDescriptor_3454248de0eb79e2, []int{1}
}

func (m *Results) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Results.Unmarshal(m, b)
}
func (m *Results) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Results.Marshal(b, m, deterministic)
}
func (m *Results) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Results.Merge(m, src)
}
func (m *Results) XXX_Size() int {
	return xxx_messageInfo_Results.Size(m)
}
func (m *Results) XXX_DiscardUnknown() {
	xxx_messageInfo_Results.DiscardUnknown(m)
}

var xxx_messageInfo_Results proto.InternalMessageInfo

func (m *Results) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*PairOfNumbers)(nil), "PairOfNumbers")
	proto.RegisterType((*Results)(nil), "Results")
}

func init() { proto.RegisterFile("proto/calculator.proto", fileDescriptor_3454248de0eb79e2) }

var fileDescriptor_3454248de0eb79e2 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4e, 0xcc, 0x49, 0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x03, 0x0b, 0x28,
	0x39, 0x73, 0xf1, 0x06, 0x24, 0x66, 0x16, 0xf9, 0xa7, 0xf9, 0x95, 0xe6, 0x26, 0xa5, 0x16, 0x15,
	0x0b, 0x49, 0x70, 0xb1, 0xe7, 0x81, 0x99, 0x86, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x4c, 0x41, 0x30,
	0x2e, 0x42, 0xc6, 0x48, 0x82, 0x09, 0x59, 0xc6, 0x48, 0x49, 0x9e, 0x8b, 0x3d, 0x28, 0xb5, 0xb8,
	0x34, 0xa7, 0xa4, 0x58, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x15, 0xaa, 0x19, 0xc2,
	0x31, 0x9a, 0xcb, 0xc8, 0xc5, 0xe5, 0x0c, 0xb7, 0x5a, 0x48, 0x9e, 0x8b, 0xd9, 0x31, 0x25, 0x45,
	0x88, 0x4f, 0x0f, 0xc5, 0x6a, 0x29, 0x0e, 0x3d, 0x98, 0x29, 0x2a, 0x5c, 0x1c, 0xc1, 0xa5, 0x49,
	0x25, 0x45, 0x89, 0xc9, 0x25, 0xf8, 0x55, 0xf9, 0x96, 0xe6, 0x94, 0x64, 0x16, 0xe4, 0x54, 0xe2,
	0x51, 0xa5, 0xc4, 0xc5, 0xe6, 0x92, 0x59, 0x96, 0x99, 0x92, 0x8a, 0x5b, 0x4d, 0x12, 0x1b, 0x38,
	0x30, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x60, 0x4c, 0x66, 0x26, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Add(ctx context.Context, in *PairOfNumbers, opts ...grpc.CallOption) (*Results, error)
	Subtract(ctx context.Context, in *PairOfNumbers, opts ...grpc.CallOption) (*Results, error)
	Multiply(ctx context.Context, in *PairOfNumbers, opts ...grpc.CallOption) (*Results, error)
	Divide(ctx context.Context, in *PairOfNumbers, opts ...grpc.CallOption) (*Results, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *PairOfNumbers, opts ...grpc.CallOption) (*Results, error) {
	out := new(Results)
	err := c.cc.Invoke(ctx, "/Calculator/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Subtract(ctx context.Context, in *PairOfNumbers, opts ...grpc.CallOption) (*Results, error) {
	out := new(Results)
	err := c.cc.Invoke(ctx, "/Calculator/Subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Multiply(ctx context.Context, in *PairOfNumbers, opts ...grpc.CallOption) (*Results, error) {
	out := new(Results)
	err := c.cc.Invoke(ctx, "/Calculator/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Divide(ctx context.Context, in *PairOfNumbers, opts ...grpc.CallOption) (*Results, error) {
	out := new(Results)
	err := c.cc.Invoke(ctx, "/Calculator/Divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Add(context.Context, *PairOfNumbers) (*Results, error)
	Subtract(context.Context, *PairOfNumbers) (*Results, error)
	Multiply(context.Context, *PairOfNumbers) (*Results, error)
	Divide(context.Context, *PairOfNumbers) (*Results, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairOfNumbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calculator/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*PairOfNumbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairOfNumbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calculator/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Subtract(ctx, req.(*PairOfNumbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairOfNumbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calculator/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Multiply(ctx, req.(*PairOfNumbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairOfNumbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calculator/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Divide(ctx, req.(*PairOfNumbers))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _Calculator_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _Calculator_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _Calculator_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/calculator.proto",
}
