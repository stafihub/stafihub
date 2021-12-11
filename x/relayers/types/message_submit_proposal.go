package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitProposal{}

func NewMsgSubmitProposal(creator string, denom string, propId []byte, proposalRoute string, proposalType string, params []byte, inFavour bool) *MsgSubmitProposal {
  return &MsgSubmitProposal{
		Creator: creator,
		Denom: denom,
		PropId: propId,
    ProposalRoute: proposalRoute,
    ProposalType: proposalType,
    Params: params,
    InFavour: inFavour,
	}
}

func (msg *MsgSubmitProposal) Route() string {
  return RouterKey
}

func (msg *MsgSubmitProposal) Type() string {
  return "SubmitProposal"
}

func (msg *MsgSubmitProposal) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitProposal) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitProposal) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

func (msg *MsgSubmitProposal) Content() *ProposalContent {
	return &ProposalContent{
		Denom: msg.Denom,
		PropId: msg.PropId,
		ProposalRoute: msg.GetProposalRoute(),
		ProposalType: msg.ProposalType,
		Params: msg.Params,
	}
}



