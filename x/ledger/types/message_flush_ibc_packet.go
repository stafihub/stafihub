package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFlushIbcPacket = "flush_ibc_packet"

var _ sdk.Msg = &MsgFlushIbcPacket{}

func NewMsgFlushIbcPacket(creator string, portID string, channelID string, sequence uint32) *MsgFlushIbcPacket {
	return &MsgFlushIbcPacket{
		Creator:   creator,
		PortID:    portID,
		ChannelID: channelID,
		Sequence:  sequence,
	}
}

func (msg *MsgFlushIbcPacket) Route() string {
	return RouterKey
}

func (msg *MsgFlushIbcPacket) Type() string {
	return TypeMsgFlushIbcPacket
}

func (msg *MsgFlushIbcPacket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFlushIbcPacket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFlushIbcPacket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
