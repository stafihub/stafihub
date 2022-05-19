package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
)

var (
	_ sdk.Msg = &MsgSetProposalLife{}
	_ sdk.Msg = &MsgSubmitProposal{}
)

func NewMsgSetProposalLife(creator sdk.AccAddress, proposalLife int64) *MsgSetProposalLife {
	return &MsgSetProposalLife{
		Creator:      creator.String(),
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
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetProposalLife) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetProposalLife) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	if msg.ProposalLife < 0 {
		return fmt.Errorf("ProposalLife %d should not lt 0", msg.ProposalLife)
	}

	return nil
}

func NewMsgSubmitProposal(proposer sdk.AccAddress, content Content) (*MsgSubmitProposal, error) {
	m := &MsgSubmitProposal{
		Proposer: proposer.String(),
	}
	err := m.SetContent(content)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (msg *MsgSubmitProposal) Route() string {
	return RouterKey
}

func (msg *MsgSubmitProposal) Type() string {
	return "SubmitProposal"
}

func (msg *MsgSubmitProposal) GetSigners() []sdk.AccAddress {
	proposer, _ := sdk.AccAddressFromBech32(msg.Proposer)
	return []sdk.AccAddress{proposer}
}

func (msg *MsgSubmitProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// String implements the Stringer interface
func (m MsgSubmitProposal) String() string {
	out, _ := yaml.Marshal(m)
	return string(out)
}

func (msg *MsgSubmitProposal) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Proposer); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Proposer)
	}

	content := msg.GetContent()
	if content == nil {
		return sdkerrors.Wrap(ErrInvalidProposalContent, "missing content")
	}
	if !IsValidProposalType(content.ProposalType()) {
		return sdkerrors.Wrap(ErrInvalidProposalType, content.ProposalType())
	}
	if err := content.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

func (m *MsgSubmitProposal) SetContent(content Content) error {
	msg, ok := content.(proto.Message)
	if !ok {
		return fmt.Errorf("can't proto marshal %T", msg)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return err
	}
	m.Content = any
	return nil
}

func (m *MsgSubmitProposal) GetContent() Content {
	content, ok := m.Content.GetCachedValue().(Content)
	if !ok {
		return nil
	}
	return content
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgSubmitProposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var content Content
	return unpacker.UnpackAny(m.Content, &content)
}
