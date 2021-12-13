package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
)

var _ sdk.Msg = &MsgSubmitProposal{}

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
	if msg.Proposer == "" {
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

func (m *MsgSubmitProposal) GetProposer() sdk.AccAddress {
	proposer, _ := sdk.AccAddressFromBech32(m.Proposer)
	return proposer
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