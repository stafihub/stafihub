package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetProposalLife{}

func NewMsgSetProposalLife(creator string, proposalLife int64) *MsgSetProposalLife {
  return &MsgSetProposalLife{
		Creator: creator,
    	ProposalLife: proposalLife,
	}
}

func (msg *MsgSetProposalLife) Route() string {
  return RouterKey
}

func (msg *MsgSetProposalLife) Type() string {
  return "SetProposalLife"
}

func (msg *MsgSetProposalLife) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgSetProposalLife) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgSetProposalLife) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

