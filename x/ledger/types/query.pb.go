// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ledger/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("ledger/query.proto", fileDescriptor_ca959e2684fc7842) }

var fileDescriptor_ca959e2684fc7842 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x49, 0x4d, 0x49,
	0x4f, 0x2d, 0xd2, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0x2d, 0x2e, 0x49, 0x4c, 0xcb, 0x04, 0xb3, 0x93, 0xf3, 0x73, 0xf4, 0xc0, 0xbc, 0x8c, 0xd2, 0x24,
	0x3d, 0x88, 0x52, 0x23, 0x76, 0x2e, 0xd6, 0x40, 0x90, 0x6a, 0x27, 0xaf, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x48, 0xcf, 0x2c, 0x01, 0xe9, 0x48, 0xce, 0xcf, 0xd5,
	0x47, 0x31, 0x4c, 0x1f, 0x66, 0x98, 0x7e, 0x85, 0x3e, 0xd4, 0xe6, 0x92, 0xca, 0x82, 0xd4, 0xe2,
	0x24, 0x36, 0xb0, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x08, 0xc9, 0x8e, 0x90,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stafiprotocol.stafihub.ledger.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "ledger/query.proto",
}