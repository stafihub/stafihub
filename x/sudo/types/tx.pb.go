// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stafihub/stafihub/sudo/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgUpdateAdmin struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgUpdateAdmin) Reset()         { *m = MsgUpdateAdmin{} }
func (m *MsgUpdateAdmin) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateAdmin) ProtoMessage()    {}
func (*MsgUpdateAdmin) Descriptor() ([]byte, []int) {
	return fileDescriptor_3330aa0f316e6a27, []int{0}
}
func (m *MsgUpdateAdmin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateAdmin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateAdmin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateAdmin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateAdmin.Merge(m, src)
}
func (m *MsgUpdateAdmin) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateAdmin) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateAdmin.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateAdmin proto.InternalMessageInfo

func (m *MsgUpdateAdmin) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateAdmin) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgUpdateAdminResponse struct {
}

func (m *MsgUpdateAdminResponse) Reset()         { *m = MsgUpdateAdminResponse{} }
func (m *MsgUpdateAdminResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateAdminResponse) ProtoMessage()    {}
func (*MsgUpdateAdminResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3330aa0f316e6a27, []int{1}
}
func (m *MsgUpdateAdminResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateAdminResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateAdminResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateAdminResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateAdminResponse.Merge(m, src)
}
func (m *MsgUpdateAdminResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateAdminResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateAdminResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateAdminResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateAdmin)(nil), "stafihub.stafihub.sudo.MsgUpdateAdmin")
	proto.RegisterType((*MsgUpdateAdminResponse)(nil), "stafihub.stafihub.sudo.MsgUpdateAdminResponse")
}

func init() { proto.RegisterFile("stafihub/stafihub/sudo/tx.proto", fileDescriptor_3330aa0f316e6a27) }

var fileDescriptor_3330aa0f316e6a27 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2e, 0x49, 0x4c,
	0xcb, 0xcc, 0x28, 0x4d, 0xd2, 0x47, 0x30, 0x4a, 0x53, 0xf2, 0xf5, 0x4b, 0x2a, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0x60, 0xe2, 0x7a, 0x08, 0x46, 0x69, 0x4a, 0xbe, 0x94, 0x48, 0x7a,
	0x7e, 0x7a, 0x3e, 0x58, 0x89, 0x3e, 0x88, 0x05, 0x51, 0x2d, 0x25, 0x97, 0x9c, 0x5f, 0x9c, 0x9b,
	0x5f, 0xac, 0x9f, 0x94, 0x98, 0x97, 0xad, 0x5f, 0x66, 0x98, 0x94, 0x5a, 0x92, 0x68, 0x08, 0xe6,
	0x40, 0xe4, 0x95, 0x5c, 0xb8, 0xf8, 0x7c, 0x8b, 0xd3, 0x43, 0x0b, 0x52, 0x12, 0x4b, 0x52, 0x1d,
	0x53, 0x72, 0x33, 0xf3, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x90, 0x4c, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71,
	0xb1, 0x04, 0x13, 0x44, 0x06, 0xca, 0x55, 0x92, 0xe0, 0x12, 0x43, 0x35, 0x25, 0x28, 0xb5, 0xb8,
	0x20, 0x3f, 0xaf, 0x38, 0xd5, 0x28, 0x87, 0x8b, 0xd9, 0xb7, 0x38, 0x5d, 0x28, 0x95, 0x8b, 0x1b,
	0xd9, 0x0e, 0x35, 0x3d, 0xec, 0x9e, 0xd0, 0x43, 0x35, 0x45, 0x4a, 0x8f, 0x38, 0x75, 0x30, 0xdb,
	0x9c, 0x9c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x33, 0x3d, 0xb3,
	0x04, 0xa4, 0x3b, 0x39, 0x3f, 0x57, 0x1f, 0x33, 0x84, 0x2b, 0xa0, 0x61, 0x5c, 0x59, 0x90, 0x5a,
	0x9c, 0xc4, 0x06, 0x0e, 0x19, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x5d, 0x69, 0xd5,
	0x8a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	UpdateAdmin(ctx context.Context, in *MsgUpdateAdmin, opts ...grpc.CallOption) (*MsgUpdateAdminResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateAdmin(ctx context.Context, in *MsgUpdateAdmin, opts ...grpc.CallOption) (*MsgUpdateAdminResponse, error) {
	out := new(MsgUpdateAdminResponse)
	err := c.cc.Invoke(ctx, "/stafihub.stafihub.sudo.Msg/UpdateAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	UpdateAdmin(context.Context, *MsgUpdateAdmin) (*MsgUpdateAdminResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateAdmin(ctx context.Context, req *MsgUpdateAdmin) (*MsgUpdateAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdmin not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafihub.stafihub.sudo.Msg/UpdateAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateAdmin(ctx, req.(*MsgUpdateAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stafihub.stafihub.sudo.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateAdmin",
			Handler:    _Msg_UpdateAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stafihub/stafihub/sudo/tx.proto",
}

func (m *MsgUpdateAdmin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateAdmin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateAdmin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateAdminResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateAdminResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateAdminResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateAdmin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateAdminResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateAdmin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateAdmin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateAdmin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateAdminResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateAdminResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateAdminResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
