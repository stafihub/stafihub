package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVoteStatusProposal{}

func NewMsgVoteStatusProposal(creator sdk.AccAddress, denom string, addresses []string,
	sourceStatus RValidatorStatus, destStatus RValidatorStatus) *MsgVoteStatusProposal {
	return &MsgVoteStatusProposal{
		Creator:      creator.String(),
		Denom:        denom,
		Addresses:    addresses,
		SourceStatus: sourceStatus,
		DestStatus:   destStatus,
	}
}

func (msg *MsgVoteStatusProposal) Route() string {
	return RouterKey
}

func (msg *MsgVoteStatusProposal) Type() string {
	return "VoteStatusProposal"
}

func (msg *MsgVoteStatusProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoteStatusProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteStatusProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
