// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: claim/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8c6eefbad679c0d, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8c6eefbad679c0d, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryIsClaimedRequest struct {
	Round uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	Index uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryIsClaimedRequest) Reset()         { *m = QueryIsClaimedRequest{} }
func (m *QueryIsClaimedRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIsClaimedRequest) ProtoMessage()    {}
func (*QueryIsClaimedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8c6eefbad679c0d, []int{2}
}
func (m *QueryIsClaimedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIsClaimedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIsClaimedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIsClaimedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIsClaimedRequest.Merge(m, src)
}
func (m *QueryIsClaimedRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIsClaimedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIsClaimedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIsClaimedRequest proto.InternalMessageInfo

func (m *QueryIsClaimedRequest) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *QueryIsClaimedRequest) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type QueryIsClaimedResponse struct {
	IsClaimed bool `protobuf:"varint,1,opt,name=isClaimed,proto3" json:"isClaimed,omitempty"`
}

func (m *QueryIsClaimedResponse) Reset()         { *m = QueryIsClaimedResponse{} }
func (m *QueryIsClaimedResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIsClaimedResponse) ProtoMessage()    {}
func (*QueryIsClaimedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8c6eefbad679c0d, []int{3}
}
func (m *QueryIsClaimedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIsClaimedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIsClaimedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIsClaimedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIsClaimedResponse.Merge(m, src)
}
func (m *QueryIsClaimedResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIsClaimedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIsClaimedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIsClaimedResponse proto.InternalMessageInfo

func (m *QueryIsClaimedResponse) GetIsClaimed() bool {
	if m != nil {
		return m.IsClaimed
	}
	return false
}

type QueryClaimSwitchRequest struct {
	Round uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
}

func (m *QueryClaimSwitchRequest) Reset()         { *m = QueryClaimSwitchRequest{} }
func (m *QueryClaimSwitchRequest) String() string { return proto.CompactTextString(m) }
func (*QueryClaimSwitchRequest) ProtoMessage()    {}
func (*QueryClaimSwitchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8c6eefbad679c0d, []int{4}
}
func (m *QueryClaimSwitchRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimSwitchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimSwitchRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimSwitchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimSwitchRequest.Merge(m, src)
}
func (m *QueryClaimSwitchRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimSwitchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimSwitchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimSwitchRequest proto.InternalMessageInfo

func (m *QueryClaimSwitchRequest) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

type QueryClaimSwitchResponse struct {
	IsOpen bool `protobuf:"varint,1,opt,name=isOpen,proto3" json:"isOpen,omitempty"`
}

func (m *QueryClaimSwitchResponse) Reset()         { *m = QueryClaimSwitchResponse{} }
func (m *QueryClaimSwitchResponse) String() string { return proto.CompactTextString(m) }
func (*QueryClaimSwitchResponse) ProtoMessage()    {}
func (*QueryClaimSwitchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8c6eefbad679c0d, []int{5}
}
func (m *QueryClaimSwitchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimSwitchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimSwitchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimSwitchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimSwitchResponse.Merge(m, src)
}
func (m *QueryClaimSwitchResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimSwitchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimSwitchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimSwitchResponse proto.InternalMessageInfo

func (m *QueryClaimSwitchResponse) GetIsOpen() bool {
	if m != nil {
		return m.IsOpen
	}
	return false
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "stafihub.stafihub.claim.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "stafihub.stafihub.claim.QueryParamsResponse")
	proto.RegisterType((*QueryIsClaimedRequest)(nil), "stafihub.stafihub.claim.QueryIsClaimedRequest")
	proto.RegisterType((*QueryIsClaimedResponse)(nil), "stafihub.stafihub.claim.QueryIsClaimedResponse")
	proto.RegisterType((*QueryClaimSwitchRequest)(nil), "stafihub.stafihub.claim.QueryClaimSwitchRequest")
	proto.RegisterType((*QueryClaimSwitchResponse)(nil), "stafihub.stafihub.claim.QueryClaimSwitchResponse")
}

func init() { proto.RegisterFile("claim/query.proto", fileDescriptor_f8c6eefbad679c0d) }

var fileDescriptor_f8c6eefbad679c0d = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3d, 0x6f, 0xd4, 0x40,
	0x10, 0xb5, 0x43, 0x62, 0x91, 0x4d, 0xc5, 0x72, 0x24, 0x27, 0x2b, 0xf2, 0x81, 0x1b, 0x50, 0x00,
	0x2f, 0x77, 0x51, 0x68, 0x10, 0x4d, 0x42, 0x43, 0x05, 0x18, 0x2a, 0x9a, 0x68, 0xed, 0x5b, 0x9c,
	0x95, 0xe2, 0x5d, 0xc7, 0xbb, 0x86, 0x44, 0x51, 0x1a, 0x6a, 0x0a, 0x24, 0x7e, 0x05, 0xff, 0x24,
	0x65, 0x24, 0x1a, 0x2a, 0x84, 0xee, 0xa8, 0xf9, 0x0d, 0xc8, 0xb3, 0x7b, 0x1f, 0xdc, 0xe1, 0xd3,
	0xd1, 0x58, 0x33, 0xe3, 0xf7, 0xe6, 0x3d, 0xbf, 0x91, 0xd1, 0x8d, 0xf4, 0x98, 0xf2, 0x9c, 0x9c,
	0x54, 0xac, 0x3c, 0x8b, 0x8a, 0x52, 0x6a, 0x89, 0xb7, 0x94, 0xa6, 0xef, 0xf8, 0x51, 0x95, 0x44,
	0xe3, 0x02, 0x40, 0x7e, 0x2b, 0x93, 0x99, 0x04, 0x0c, 0xa9, 0x2b, 0x03, 0xf7, 0xb7, 0x33, 0x29,
	0xb3, 0x63, 0x46, 0x68, 0xc1, 0x09, 0x15, 0x42, 0x6a, 0xaa, 0xb9, 0x14, 0xca, 0xbe, 0xdd, 0x49,
	0xa5, 0xca, 0xa5, 0x22, 0x09, 0x55, 0xcc, 0xa8, 0x90, 0xf7, 0xdd, 0x84, 0x69, 0xda, 0x25, 0x05,
	0xcd, 0xb8, 0x00, 0xb0, 0xc5, 0x62, 0xe3, 0xa5, 0xa0, 0x25, 0xcd, 0x2d, 0x3f, 0x6c, 0x21, 0xfc,
	0xaa, 0x66, 0xbd, 0x84, 0x61, 0xcc, 0x4e, 0x2a, 0xa6, 0x74, 0xf8, 0x06, 0xdd, 0xfc, 0x6b, 0xaa,
	0x0a, 0x29, 0x14, 0xc3, 0x4f, 0x91, 0x67, 0xc8, 0x6d, 0xf7, 0xb6, 0x7b, 0x6f, 0xa3, 0xd7, 0x89,
	0x1a, 0x3e, 0x25, 0x32, 0xc4, 0xfd, 0xd5, 0xcb, 0x1f, 0x1d, 0x27, 0xb6, 0xa4, 0xf0, 0x00, 0xdd,
	0x82, 0xad, 0xcf, 0xd5, 0x41, 0x0d, 0x62, 0x7d, 0x2b, 0x87, 0x5b, 0x68, 0xad, 0x94, 0x95, 0xe8,
	0xc3, 0xda, 0xd5, 0xd8, 0x34, 0xf5, 0x94, 0x8b, 0x3e, 0x3b, 0x6d, 0xaf, 0x98, 0x29, 0x34, 0xe1,
	0x63, 0xb4, 0x39, 0xbb, 0xc4, 0xba, 0xdb, 0x46, 0xeb, 0x7c, 0x34, 0x84, 0x4d, 0xd7, 0xe3, 0xc9,
	0x20, 0x24, 0x68, 0x0b, 0x78, 0xd0, 0xbf, 0xfe, 0xc0, 0x75, 0x7a, 0xb4, 0x50, 0x3e, 0xec, 0xa1,
	0xf6, 0x3c, 0xc1, 0x4a, 0x6d, 0x22, 0x8f, 0xab, 0x17, 0x05, 0x13, 0x56, 0xc7, 0x76, 0xbd, 0xdf,
	0xd7, 0xd0, 0x1a, 0x90, 0xf0, 0x27, 0x17, 0x79, 0x26, 0x04, 0x7c, 0xbf, 0x31, 0xa5, 0xf9, 0xe4,
	0xfd, 0x07, 0xcb, 0x81, 0x8d, 0x8f, 0xf0, 0xee, 0xc7, 0x6f, 0xbf, 0xbe, 0xac, 0xdc, 0xc1, 0x1d,
	0x32, 0x02, 0x4f, 0x8a, 0xe9, 0x63, 0xe3, 0xaf, 0x2e, 0x5a, 0x1f, 0x27, 0x86, 0xa3, 0xc5, 0x22,
	0xb3, 0xf7, 0xf1, 0xc9, 0xd2, 0x78, 0xeb, 0xeb, 0x09, 0xf8, 0xda, 0xc3, 0xbb, 0x8d, 0xbe, 0xb8,
	0x3a, 0x4c, 0x0d, 0x89, 0x9c, 0x43, 0xde, 0x17, 0xe4, 0x1c, 0x0e, 0x7c, 0x51, 0x7b, 0xdd, 0x98,
	0x0a, 0x1d, 0x3f, 0x5a, 0xac, 0x3e, 0x7f, 0x50, 0xbf, 0xfb, 0x1f, 0x0c, 0xeb, 0x78, 0x0f, 0x1c,
	0x13, 0xfc, 0xb0, 0xd1, 0x31, 0x3c, 0x0f, 0x15, 0xd0, 0x46, 0x9e, 0xf7, 0x9f, 0x5d, 0x0e, 0x02,
	0xf7, 0x6a, 0x10, 0xb8, 0x3f, 0x07, 0x81, 0xfb, 0x79, 0x18, 0x38, 0x57, 0xc3, 0xc0, 0xf9, 0x3e,
	0x0c, 0x9c, 0xb7, 0x3b, 0x19, 0xd7, 0xa0, 0x2b, 0xf3, 0x7f, 0xac, 0x3c, 0xb5, 0x4b, 0xf5, 0x59,
	0xc1, 0x54, 0xe2, 0xc1, 0xbf, 0xb8, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x82, 0xa0, 0xa1,
	0x2d, 0x04, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of IsClaimed items.
	IsClaimed(ctx context.Context, in *QueryIsClaimedRequest, opts ...grpc.CallOption) (*QueryIsClaimedResponse, error)
	// Queries a list of ClaimSwitch items.
	ClaimSwitch(ctx context.Context, in *QueryClaimSwitchRequest, opts ...grpc.CallOption) (*QueryClaimSwitchResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/stafihub.stafihub.claim.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IsClaimed(ctx context.Context, in *QueryIsClaimedRequest, opts ...grpc.CallOption) (*QueryIsClaimedResponse, error) {
	out := new(QueryIsClaimedResponse)
	err := c.cc.Invoke(ctx, "/stafihub.stafihub.claim.Query/IsClaimed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ClaimSwitch(ctx context.Context, in *QueryClaimSwitchRequest, opts ...grpc.CallOption) (*QueryClaimSwitchResponse, error) {
	out := new(QueryClaimSwitchResponse)
	err := c.cc.Invoke(ctx, "/stafihub.stafihub.claim.Query/ClaimSwitch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of IsClaimed items.
	IsClaimed(context.Context, *QueryIsClaimedRequest) (*QueryIsClaimedResponse, error)
	// Queries a list of ClaimSwitch items.
	ClaimSwitch(context.Context, *QueryClaimSwitchRequest) (*QueryClaimSwitchResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) IsClaimed(ctx context.Context, req *QueryIsClaimedRequest) (*QueryIsClaimedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsClaimed not implemented")
}
func (*UnimplementedQueryServer) ClaimSwitch(ctx context.Context, req *QueryClaimSwitchRequest) (*QueryClaimSwitchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimSwitch not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafihub.stafihub.claim.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IsClaimed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIsClaimedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IsClaimed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafihub.stafihub.claim.Query/IsClaimed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IsClaimed(ctx, req.(*QueryIsClaimedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ClaimSwitch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimSwitchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ClaimSwitch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafihub.stafihub.claim.Query/ClaimSwitch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ClaimSwitch(ctx, req.(*QueryClaimSwitchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stafihub.stafihub.claim.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "IsClaimed",
			Handler:    _Query_IsClaimed_Handler,
		},
		{
			MethodName: "ClaimSwitch",
			Handler:    _Query_ClaimSwitch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "claim/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryIsClaimedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIsClaimedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIsClaimedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if m.Round != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryIsClaimedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIsClaimedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIsClaimedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsClaimed {
		i--
		if m.IsClaimed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryClaimSwitchRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimSwitchRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimSwitchRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Round != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryClaimSwitchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimSwitchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimSwitchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsOpen {
		i--
		if m.IsOpen {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryIsClaimedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovQuery(uint64(m.Round))
	}
	if m.Index != 0 {
		n += 1 + sovQuery(uint64(m.Index))
	}
	return n
}

func (m *QueryIsClaimedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsClaimed {
		n += 2
	}
	return n
}

func (m *QueryClaimSwitchRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovQuery(uint64(m.Round))
	}
	return n
}

func (m *QueryClaimSwitchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsOpen {
		n += 2
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryIsClaimedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryIsClaimedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIsClaimedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryIsClaimedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryIsClaimedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIsClaimedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsClaimed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsClaimed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryClaimSwitchRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryClaimSwitchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimSwitchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryClaimSwitchResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryClaimSwitchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimSwitchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsOpen", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsOpen = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
