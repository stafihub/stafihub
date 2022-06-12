package types

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
	"github.com/tendermint/tendermint/crypto"
)

const TypeUpdateRValidatorProposal = "update_r_validator"

var _ sdk.Msg = &UpdateRValidatorProposal{}
var ExecuteBondProposalType = "ExecuteBondProposal"

func init() {
	rvotetypes.RegisterProposalType(TypeUpdateRValidatorProposal)
	rvotetypes.RegisterProposalTypeCodec(&UpdateRValidatorProposal{}, "rvalidator/UpdateRValidator")
}

func NewUpdateRValidatorProposal(creator string, denom string, oldAddress string, newAddress string, cycle *Cycle) *UpdateRValidatorProposal {
	msg := UpdateRValidatorProposal{
		Denom:      denom,
		OldAddress: oldAddress,
		NewAddress: newAddress,
		Cycle:      cycle,
	}
	msg.setPropId()

	msg.Creator = creator

	return &msg
}

func (p *UpdateRValidatorProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
}

func (p *UpdateRValidatorProposal) ProposalRoute() string {
	return ModuleName
}

func (p *UpdateRValidatorProposal) ProposalType() string {
	return TypeUpdateRValidatorProposal
}

func (p *UpdateRValidatorProposal) InFavour() bool {
	return true
}

func (msg *UpdateRValidatorProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *UpdateRValidatorProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *UpdateRValidatorProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
