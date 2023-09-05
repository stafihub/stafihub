package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetInterchainTxProposalStatus = "set_interchain_tx_proposal_status"

var _ sdk.Msg = &MsgSetInterchainTxProposalStatus{}

func NewMsgSetInterchainTxProposalStatus(creator string, proposalId string, status InterchainTxStatus) *MsgSetInterchainTxProposalStatus {
	return &MsgSetInterchainTxProposalStatus{
		Creator:    creator,
		ProposalId: proposalId,
		Status:     status,
	}
}

func (msg *MsgSetInterchainTxProposalStatus) Route() string {
	return RouterKey
}

func (msg *MsgSetInterchainTxProposalStatus) Type() string {
	return TypeMsgSetInterchainTxProposalStatus
}

func (msg *MsgSetInterchainTxProposalStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetInterchainTxProposalStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetInterchainTxProposalStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
