// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/staking/v1beta1/lsm_tx.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgRedeemTokensForShares redeems a tokenized share back into a native
// delegation
type MsgRedeemTokensForShares struct {
	DelegatorAddress string     `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	Amount           types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgRedeemTokensForShares) Reset()         { *m = MsgRedeemTokensForShares{} }
func (m *MsgRedeemTokensForShares) String() string { return proto.CompactTextString(m) }
func (*MsgRedeemTokensForShares) ProtoMessage()    {}
func (*MsgRedeemTokensForShares) Descriptor() ([]byte, []int) {
	return fileDescriptor_34c3b474a863e424, []int{0}
}
func (m *MsgRedeemTokensForShares) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRedeemTokensForShares) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRedeemTokensForShares.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRedeemTokensForShares) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRedeemTokensForShares.Merge(m, src)
}
func (m *MsgRedeemTokensForShares) XXX_Size() int {
	return m.Size()
}
func (m *MsgRedeemTokensForShares) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRedeemTokensForShares.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRedeemTokensForShares proto.InternalMessageInfo

// MsgRedeemTokensForSharesResponse defines the Msg/MsgRedeemTokensForShares
// response type.
type MsgRedeemTokensForSharesResponse struct {
	Amount types.Coin `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgRedeemTokensForSharesResponse) Reset()         { *m = MsgRedeemTokensForSharesResponse{} }
func (m *MsgRedeemTokensForSharesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRedeemTokensForSharesResponse) ProtoMessage()    {}
func (*MsgRedeemTokensForSharesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34c3b474a863e424, []int{1}
}
func (m *MsgRedeemTokensForSharesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRedeemTokensForSharesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRedeemTokensForSharesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRedeemTokensForSharesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRedeemTokensForSharesResponse.Merge(m, src)
}
func (m *MsgRedeemTokensForSharesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRedeemTokensForSharesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRedeemTokensForSharesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRedeemTokensForSharesResponse proto.InternalMessageInfo

func (m *MsgRedeemTokensForSharesResponse) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

// MsgTransferTokenizeShareRecord transfer a tokenize share record
type MsgTransferTokenizeShareRecord struct {
	TokenizeShareRecordId uint64 `protobuf:"varint,1,opt,name=tokenize_share_record_id,json=tokenizeShareRecordId,proto3" json:"tokenize_share_record_id,omitempty"`
	Sender                string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	NewOwner              string `protobuf:"bytes,3,opt,name=new_owner,json=newOwner,proto3" json:"new_owner,omitempty"`
}

func (m *MsgTransferTokenizeShareRecord) Reset()         { *m = MsgTransferTokenizeShareRecord{} }
func (m *MsgTransferTokenizeShareRecord) String() string { return proto.CompactTextString(m) }
func (*MsgTransferTokenizeShareRecord) ProtoMessage()    {}
func (*MsgTransferTokenizeShareRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_34c3b474a863e424, []int{2}
}
func (m *MsgTransferTokenizeShareRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferTokenizeShareRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferTokenizeShareRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferTokenizeShareRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferTokenizeShareRecord.Merge(m, src)
}
func (m *MsgTransferTokenizeShareRecord) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferTokenizeShareRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferTokenizeShareRecord.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferTokenizeShareRecord proto.InternalMessageInfo

// MsgTransferTokenizeShareRecordResponse defines the Msg/MsgTransferTokenizeShareRecord response type.
type MsgTransferTokenizeShareRecordResponse struct {
}

func (m *MsgTransferTokenizeShareRecordResponse) Reset() {
	*m = MsgTransferTokenizeShareRecordResponse{}
}
func (m *MsgTransferTokenizeShareRecordResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransferTokenizeShareRecordResponse) ProtoMessage()    {}
func (*MsgTransferTokenizeShareRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34c3b474a863e424, []int{3}
}
func (m *MsgTransferTokenizeShareRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferTokenizeShareRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferTokenizeShareRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferTokenizeShareRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferTokenizeShareRecordResponse.Merge(m, src)
}
func (m *MsgTransferTokenizeShareRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferTokenizeShareRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferTokenizeShareRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferTokenizeShareRecordResponse proto.InternalMessageInfo

// MsgTokenizeShares tokenizes a delegation
type MsgTokenizeShares struct {
	DelegatorAddress    string     `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	ValidatorAddress    string     `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
	Amount              types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	TokenizedShareOwner string     `protobuf:"bytes,4,opt,name=tokenized_share_owner,json=tokenizedShareOwner,proto3" json:"tokenized_share_owner,omitempty"`
}

func (m *MsgTokenizeShares) Reset()         { *m = MsgTokenizeShares{} }
func (m *MsgTokenizeShares) String() string { return proto.CompactTextString(m) }
func (*MsgTokenizeShares) ProtoMessage()    {}
func (*MsgTokenizeShares) Descriptor() ([]byte, []int) {
	return fileDescriptor_34c3b474a863e424, []int{4}
}
func (m *MsgTokenizeShares) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTokenizeShares) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTokenizeShares.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTokenizeShares) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTokenizeShares.Merge(m, src)
}
func (m *MsgTokenizeShares) XXX_Size() int {
	return m.Size()
}
func (m *MsgTokenizeShares) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTokenizeShares.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTokenizeShares proto.InternalMessageInfo

// MsgTokenizeSharesResponse defines the Msg/MsgTokenizeShares response type.
type MsgTokenizeSharesResponse struct {
	Amount types.Coin `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgTokenizeSharesResponse) Reset()         { *m = MsgTokenizeSharesResponse{} }
func (m *MsgTokenizeSharesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTokenizeSharesResponse) ProtoMessage()    {}
func (*MsgTokenizeSharesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34c3b474a863e424, []int{5}
}
func (m *MsgTokenizeSharesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTokenizeSharesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTokenizeSharesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTokenizeSharesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTokenizeSharesResponse.Merge(m, src)
}
func (m *MsgTokenizeSharesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTokenizeSharesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTokenizeSharesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTokenizeSharesResponse proto.InternalMessageInfo

func (m *MsgTokenizeSharesResponse) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*MsgRedeemTokensForShares)(nil), "cosmos.staking.v1beta1.MsgRedeemTokensForShares")
	proto.RegisterType((*MsgRedeemTokensForSharesResponse)(nil), "cosmos.staking.v1beta1.MsgRedeemTokensForSharesResponse")
	proto.RegisterType((*MsgTransferTokenizeShareRecord)(nil), "cosmos.staking.v1beta1.MsgTransferTokenizeShareRecord")
	proto.RegisterType((*MsgTransferTokenizeShareRecordResponse)(nil), "cosmos.staking.v1beta1.MsgTransferTokenizeShareRecordResponse")
	proto.RegisterType((*MsgTokenizeShares)(nil), "cosmos.staking.v1beta1.MsgTokenizeShares")
	proto.RegisterType((*MsgTokenizeSharesResponse)(nil), "cosmos.staking.v1beta1.MsgTokenizeSharesResponse")
}

func init() {
	// proto.RegisterFile("cosmos/staking/v1beta1/lsm_tx.proto", fileDescriptor_34c3b474a863e424)
}

var fileDescriptor_34c3b474a863e424 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xed, 0x34, 0x8a, 0x9a, 0x63, 0x69, 0x0c, 0x54, 0x6e, 0x41, 0x4e, 0x64, 0x24, 0x14,
	0x09, 0xc9, 0x56, 0xcb, 0x50, 0xa9, 0x1b, 0x41, 0x20, 0x65, 0xa8, 0x90, 0x4c, 0x26, 0x18, 0xac,
	0x73, 0xee, 0xd5, 0xb1, 0x6a, 0xdf, 0x45, 0xf7, 0x2e, 0x4d, 0xcb, 0x27, 0x60, 0x64, 0x64, 0xcc,
	0xc8, 0x47, 0xe9, 0xd8, 0x91, 0xa9, 0x42, 0xc9, 0xc2, 0xcc, 0x27, 0x40, 0x3e, 0xbb, 0x26, 0xd4,
	0x02, 0x29, 0x82, 0xcd, 0xf7, 0xde, 0xcf, 0xff, 0xf7, 0xfe, 0xef, 0xdd, 0x91, 0x27, 0x63, 0x81,
	0x99, 0x40, 0x1f, 0x15, 0x3d, 0x4b, 0x78, 0xec, 0x9f, 0x1f, 0x44, 0xa0, 0xe8, 0x81, 0x9f, 0x62,
	0x16, 0xaa, 0x0b, 0x6f, 0x2a, 0x85, 0x12, 0xd6, 0x6e, 0x01, 0x79, 0x25, 0xe4, 0x95, 0xd0, 0xfe,
	0x83, 0x58, 0xc4, 0x42, 0x23, 0x7e, 0xfe, 0x55, 0xd0, 0xfb, 0x4e, 0x29, 0x19, 0x51, 0x84, 0x4a,
	0x6f, 0x2c, 0x12, 0x5e, 0xe4, 0xdd, 0x2f, 0x26, 0xb1, 0x4f, 0x30, 0x0e, 0x80, 0x01, 0x64, 0x23,
	0x71, 0x06, 0x1c, 0x5f, 0x0b, 0xf9, 0x76, 0x42, 0x25, 0xa0, 0x35, 0x24, 0x1d, 0x06, 0x29, 0xc4,
	0x54, 0x09, 0x19, 0x52, 0xc6, 0x24, 0x20, 0xda, 0x66, 0xcf, 0xec, 0xb7, 0x07, 0x8f, 0x7f, 0xdc,
	0x74, 0xed, 0x4b, 0x9a, 0xa5, 0xc7, 0x6e, 0x0d, 0x71, 0x83, 0x9d, 0x2a, 0xf6, 0xa2, 0x08, 0x59,
	0x47, 0xa4, 0x45, 0x33, 0x31, 0xe3, 0xca, 0x6e, 0xf4, 0xcc, 0xfe, 0xbd, 0xc3, 0x3d, 0xaf, 0xb4,
	0x91, 0x37, 0x76, 0xeb, 0xc1, 0x7b, 0x29, 0x12, 0x3e, 0x68, 0x5e, 0xdd, 0x74, 0x8d, 0xa0, 0xc4,
	0x8f, 0xb7, 0x3f, 0x2e, 0xba, 0xc6, 0xf7, 0x45, 0xd7, 0x70, 0xdf, 0x93, 0xde, 0x9f, 0x3a, 0x0d,
	0x00, 0xa7, 0x82, 0x23, 0xac, 0x95, 0x31, 0x37, 0x2a, 0xe3, 0x7e, 0x36, 0x89, 0x73, 0x82, 0xf1,
	0x48, 0x52, 0x8e, 0xa7, 0x20, 0xb5, 0x7e, 0xf2, 0x01, 0xb4, 0x7c, 0x00, 0x63, 0x21, 0x99, 0x75,
	0x44, 0x6c, 0x55, 0x86, 0x43, 0xcc, 0xe3, 0xa1, 0xd4, 0x89, 0x30, 0x61, 0xba, 0x5a, 0x33, 0x78,
	0xa8, 0xea, 0xbf, 0x0d, 0x99, 0xb5, 0x4b, 0x5a, 0x08, 0x9c, 0x81, 0xd4, 0xde, 0xdb, 0x41, 0x79,
	0xb2, 0x1e, 0x91, 0x36, 0x87, 0x79, 0x28, 0xe6, 0x1c, 0xa4, 0xbd, 0xa5, 0x53, 0xdb, 0x1c, 0xe6,
	0x6f, 0xf2, 0xf3, 0x9a, 0xef, 0x3e, 0x79, 0xfa, 0xf7, 0xce, 0x6e, 0xdd, 0xbb, 0x8b, 0x06, 0xe9,
	0xe4, 0xe8, 0x3a, 0xf2, 0x5f, 0xb7, 0x38, 0x24, 0x9d, 0x73, 0x9a, 0x26, 0xec, 0x37, 0xa9, 0xc6,
	0x5d, 0xa9, 0x1a, 0xe2, 0x06, 0x3b, 0x55, 0xac, 0x7e, 0x21, 0xb6, 0x36, 0xda, 0x94, 0x75, 0x48,
	0xaa, 0x31, 0xb3, 0x72, 0x0f, 0xc5, 0x04, 0x9b, 0x7a, 0x82, 0xf7, 0xab, 0xa4, 0xb6, 0x7f, 0x77,
	0x98, 0x23, 0xb2, 0x57, 0x9b, 0xd0, 0x3f, 0xdf, 0x9e, 0xc1, 0xab, 0xab, 0xa5, 0x63, 0x5e, 0x2f,
	0x1d, 0xf3, 0xdb, 0xd2, 0x31, 0x3f, 0xad, 0x1c, 0xe3, 0x7a, 0xe5, 0x18, 0x5f, 0x57, 0x8e, 0xf1,
	0xee, 0x59, 0x9c, 0xa8, 0xc9, 0x2c, 0xf2, 0xc6, 0x22, 0xcb, 0x9f, 0xf5, 0x69, 0x32, 0x99, 0x45,
	0xbf, 0x3e, 0x2e, 0xfc, 0x14, 0x58, 0x0c, 0xd2, 0x57, 0x97, 0x53, 0xc0, 0xa8, 0xa5, 0xdf, 0xe4,
	0xf3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x2a, 0x65, 0xaa, 0x08, 0x04, 0x00, 0x00,
}

func (m *MsgRedeemTokensForShares) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRedeemTokensForShares) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRedeemTokensForShares) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLsmTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintLsmTx(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRedeemTokensForSharesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRedeemTokensForSharesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRedeemTokensForSharesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLsmTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgTransferTokenizeShareRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferTokenizeShareRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferTokenizeShareRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewOwner) > 0 {
		i -= len(m.NewOwner)
		copy(dAtA[i:], m.NewOwner)
		i = encodeVarintLsmTx(dAtA, i, uint64(len(m.NewOwner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintLsmTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.TokenizeShareRecordId != 0 {
		i = encodeVarintLsmTx(dAtA, i, uint64(m.TokenizeShareRecordId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferTokenizeShareRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferTokenizeShareRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferTokenizeShareRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgTokenizeShares) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTokenizeShares) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTokenizeShares) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenizedShareOwner) > 0 {
		i -= len(m.TokenizedShareOwner)
		copy(dAtA[i:], m.TokenizedShareOwner)
		i = encodeVarintLsmTx(dAtA, i, uint64(len(m.TokenizedShareOwner)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLsmTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintLsmTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintLsmTx(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTokenizeSharesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTokenizeSharesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTokenizeSharesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLsmTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintLsmTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovLsmTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRedeemTokensForShares) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovLsmTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovLsmTx(uint64(l))
	return n
}

func (m *MsgRedeemTokensForSharesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovLsmTx(uint64(l))
	return n
}

func (m *MsgTransferTokenizeShareRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TokenizeShareRecordId != 0 {
		n += 1 + sovLsmTx(uint64(m.TokenizeShareRecordId))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovLsmTx(uint64(l))
	}
	l = len(m.NewOwner)
	if l > 0 {
		n += 1 + l + sovLsmTx(uint64(l))
	}
	return n
}

func (m *MsgTransferTokenizeShareRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgTokenizeShares) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovLsmTx(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovLsmTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovLsmTx(uint64(l))
	l = len(m.TokenizedShareOwner)
	if l > 0 {
		n += 1 + l + sovLsmTx(uint64(l))
	}
	return n
}

func (m *MsgTokenizeSharesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovLsmTx(uint64(l))
	return n
}

func sovLsmTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLsmTx(x uint64) (n int) {
	return sovLsmTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRedeemTokensForShares) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLsmTx
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
			return fmt.Errorf("proto: MsgRedeemTokensForShares: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRedeemTokensForShares: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLsmTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLsmTx
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
func (m *MsgRedeemTokensForSharesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLsmTx
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
			return fmt.Errorf("proto: MsgRedeemTokensForSharesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRedeemTokensForSharesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLsmTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLsmTx
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
func (m *MsgTransferTokenizeShareRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLsmTx
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
			return fmt.Errorf("proto: MsgTransferTokenizeShareRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferTokenizeShareRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenizeShareRecordId", wireType)
			}
			m.TokenizeShareRecordId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenizeShareRecordId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLsmTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLsmTx
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
func (m *MsgTransferTokenizeShareRecordResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLsmTx
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
			return fmt.Errorf("proto: MsgTransferTokenizeShareRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferTokenizeShareRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLsmTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLsmTx
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
func (m *MsgTokenizeShares) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLsmTx
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
			return fmt.Errorf("proto: MsgTokenizeShares: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTokenizeShares: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenizedShareOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenizedShareOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLsmTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLsmTx
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
func (m *MsgTokenizeSharesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLsmTx
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
			return fmt.Errorf("proto: MsgTokenizeSharesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTokenizeSharesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLsmTx
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
				return ErrInvalidLengthLsmTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLsmTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLsmTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLsmTx
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
func skipLsmTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLsmTx
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
					return 0, ErrIntOverflowLsmTx
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
					return 0, ErrIntOverflowLsmTx
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
				return 0, ErrInvalidLengthLsmTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLsmTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLsmTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLsmTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLsmTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLsmTx = fmt.Errorf("proto: unexpected end of group")
)
