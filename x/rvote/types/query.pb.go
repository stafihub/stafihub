// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stafihub/stafihub/rvote/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

type QueryGetProposalRequest struct {
	PropId string `protobuf:"bytes,1,opt,name=propId,proto3" json:"propId,omitempty"`
}

func (m *QueryGetProposalRequest) Reset()         { *m = QueryGetProposalRequest{} }
func (m *QueryGetProposalRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetProposalRequest) ProtoMessage()    {}
func (*QueryGetProposalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea308eea6c1a82d2, []int{0}
}
func (m *QueryGetProposalRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetProposalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetProposalRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetProposalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetProposalRequest.Merge(m, src)
}
func (m *QueryGetProposalRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetProposalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetProposalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetProposalRequest proto.InternalMessageInfo

func (m *QueryGetProposalRequest) GetPropId() string {
	if m != nil {
		return m.PropId
	}
	return ""
}

type QueryGetProposalResponse struct {
	Proposal string `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
}

func (m *QueryGetProposalResponse) Reset()         { *m = QueryGetProposalResponse{} }
func (m *QueryGetProposalResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetProposalResponse) ProtoMessage()    {}
func (*QueryGetProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea308eea6c1a82d2, []int{1}
}
func (m *QueryGetProposalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetProposalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetProposalResponse.Merge(m, src)
}
func (m *QueryGetProposalResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetProposalResponse proto.InternalMessageInfo

func (m *QueryGetProposalResponse) GetProposal() string {
	if m != nil {
		return m.Proposal
	}
	return ""
}

type QueryGetProposalLifeRequest struct {
}

func (m *QueryGetProposalLifeRequest) Reset()         { *m = QueryGetProposalLifeRequest{} }
func (m *QueryGetProposalLifeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetProposalLifeRequest) ProtoMessage()    {}
func (*QueryGetProposalLifeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea308eea6c1a82d2, []int{2}
}
func (m *QueryGetProposalLifeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetProposalLifeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetProposalLifeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetProposalLifeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetProposalLifeRequest.Merge(m, src)
}
func (m *QueryGetProposalLifeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetProposalLifeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetProposalLifeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetProposalLifeRequest proto.InternalMessageInfo

type QueryGetProposalLifeResponse struct {
	ProposalLife int64 `protobuf:"varint,1,opt,name=proposalLife,proto3" json:"proposalLife,omitempty"`
}

func (m *QueryGetProposalLifeResponse) Reset()         { *m = QueryGetProposalLifeResponse{} }
func (m *QueryGetProposalLifeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetProposalLifeResponse) ProtoMessage()    {}
func (*QueryGetProposalLifeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea308eea6c1a82d2, []int{3}
}
func (m *QueryGetProposalLifeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetProposalLifeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetProposalLifeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetProposalLifeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetProposalLifeResponse.Merge(m, src)
}
func (m *QueryGetProposalLifeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetProposalLifeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetProposalLifeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetProposalLifeResponse proto.InternalMessageInfo

func (m *QueryGetProposalLifeResponse) GetProposalLife() int64 {
	if m != nil {
		return m.ProposalLife
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryGetProposalRequest)(nil), "stafihub.stafihub.rvote.QueryGetProposalRequest")
	proto.RegisterType((*QueryGetProposalResponse)(nil), "stafihub.stafihub.rvote.QueryGetProposalResponse")
	proto.RegisterType((*QueryGetProposalLifeRequest)(nil), "stafihub.stafihub.rvote.QueryGetProposalLifeRequest")
	proto.RegisterType((*QueryGetProposalLifeResponse)(nil), "stafihub.stafihub.rvote.QueryGetProposalLifeResponse")
}

func init() {
	proto.RegisterFile("stafihub/stafihub/rvote/query.proto", fileDescriptor_ea308eea6c1a82d2)
}

var fileDescriptor_ea308eea6c1a82d2 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x2e, 0x49, 0x4c,
	0xcb, 0xcc, 0x28, 0x4d, 0xd2, 0x87, 0x33, 0x8a, 0xca, 0xf2, 0x4b, 0x52, 0xf5, 0x0b, 0x4b, 0x53,
	0x8b, 0x2a, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0x61, 0x72, 0x7a, 0x70, 0x06, 0x58,
	0x91, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e,
	0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x9b, 0x92, 0x21, 0x97, 0x78, 0x20,
	0xc8, 0x14, 0xf7, 0xd4, 0x92, 0x80, 0xa2, 0xfc, 0x82, 0xfc, 0xe2, 0xc4, 0x9c, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x31, 0x2e, 0xb6, 0x82, 0xa2, 0xfc, 0x02, 0xcf, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0xc9, 0x8c, 0x4b, 0x02, 0x53, 0x4b, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x14, 0x17, 0x47, 0x01, 0x54, 0x0c, 0xaa, 0x0b, 0xce, 0x57, 0x92, 0xe5,
	0x92, 0x46, 0xd7, 0xe7, 0x93, 0x99, 0x96, 0x0a, 0xb5, 0x4e, 0xc9, 0x89, 0x4b, 0x06, 0xbb, 0x34,
	0xd4, 0x68, 0x25, 0x2e, 0x9e, 0x02, 0x24, 0x71, 0xb0, 0xf1, 0xcc, 0x41, 0x28, 0x62, 0x46, 0x2f,
	0x99, 0xb8, 0x58, 0xc1, 0x86, 0x08, 0xad, 0x62, 0xe4, 0xe2, 0x46, 0x32, 0x49, 0xc8, 0x40, 0x0f,
	0x47, 0xf8, 0xe8, 0xe1, 0xf0, 0xbe, 0x94, 0x21, 0x09, 0x3a, 0x20, 0x4e, 0x54, 0x32, 0x6b, 0xba,
	0xfc, 0x64, 0x32, 0x93, 0x81, 0x90, 0x9e, 0x3e, 0xae, 0x18, 0x4b, 0x4f, 0x2d, 0x89, 0x87, 0xb9,
	0x58, 0xbf, 0x1a, 0x12, 0xa0, 0xb5, 0x42, 0x9b, 0x18, 0xb9, 0xf8, 0xd1, 0xbc, 0x2d, 0x64, 0x42,
	0xb4, 0xf5, 0x48, 0x81, 0x28, 0x65, 0x4a, 0xa2, 0x2e, 0xa8, 0xc3, 0x8d, 0xc0, 0x0e, 0xd7, 0x11,
	0xd2, 0x22, 0xca, 0xe1, 0xf1, 0x39, 0x99, 0x69, 0xa9, 0x4e, 0x2e, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x95, 0x9e, 0x59, 0x02, 0xb2, 0x38, 0x39, 0x3f, 0x17, 0x8b,
	0x79, 0x15, 0x50, 0x13, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xc9, 0xd0, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x03, 0x9e, 0x4d, 0x09, 0xe4, 0x02, 0x00, 0x00,
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
	// Queries a list of getProposal items.
	GetProposal(ctx context.Context, in *QueryGetProposalRequest, opts ...grpc.CallOption) (*QueryGetProposalResponse, error)
	// Queries a list of getProposalLife items.
	GetProposalLife(ctx context.Context, in *QueryGetProposalLifeRequest, opts ...grpc.CallOption) (*QueryGetProposalLifeResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetProposal(ctx context.Context, in *QueryGetProposalRequest, opts ...grpc.CallOption) (*QueryGetProposalResponse, error) {
	out := new(QueryGetProposalResponse)
	err := c.cc.Invoke(ctx, "/stafihub.stafihub.rvote.Query/GetProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetProposalLife(ctx context.Context, in *QueryGetProposalLifeRequest, opts ...grpc.CallOption) (*QueryGetProposalLifeResponse, error) {
	out := new(QueryGetProposalLifeResponse)
	err := c.cc.Invoke(ctx, "/stafihub.stafihub.rvote.Query/GetProposalLife", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a list of getProposal items.
	GetProposal(context.Context, *QueryGetProposalRequest) (*QueryGetProposalResponse, error)
	// Queries a list of getProposalLife items.
	GetProposalLife(context.Context, *QueryGetProposalLifeRequest) (*QueryGetProposalLifeResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetProposal(ctx context.Context, req *QueryGetProposalRequest) (*QueryGetProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProposal not implemented")
}
func (*UnimplementedQueryServer) GetProposalLife(ctx context.Context, req *QueryGetProposalLifeRequest) (*QueryGetProposalLifeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProposalLife not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafihub.stafihub.rvote.Query/GetProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetProposal(ctx, req.(*QueryGetProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetProposalLife_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetProposalLifeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetProposalLife(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafihub.stafihub.rvote.Query/GetProposalLife",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetProposalLife(ctx, req.(*QueryGetProposalLifeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stafihub.stafihub.rvote.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProposal",
			Handler:    _Query_GetProposal_Handler,
		},
		{
			MethodName: "GetProposalLife",
			Handler:    _Query_GetProposalLife_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stafihub/stafihub/rvote/query.proto",
}

func (m *QueryGetProposalRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetProposalRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetProposalRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PropId) > 0 {
		i -= len(m.PropId)
		copy(dAtA[i:], m.PropId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PropId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetProposalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetProposalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetProposalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposal) > 0 {
		i -= len(m.Proposal)
		copy(dAtA[i:], m.Proposal)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Proposal)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetProposalLifeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetProposalLifeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetProposalLifeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetProposalLifeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetProposalLifeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetProposalLifeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalLife != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalLife))
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
func (m *QueryGetProposalRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PropId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetProposalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proposal)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetProposalLifeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetProposalLifeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalLife != 0 {
		n += 1 + sovQuery(uint64(m.ProposalLife))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetProposalRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetProposalRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetProposalRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PropId = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetProposalResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetProposalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposal = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetProposalLifeRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetProposalLifeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetProposalLifeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryGetProposalLifeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetProposalLifeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetProposalLifeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalLife", wireType)
			}
			m.ProposalLife = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalLife |= int64(b&0x7F) << shift
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
