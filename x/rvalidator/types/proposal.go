package types

import (
	"encoding/hex"
	"fmt"

	"github.com/cometbft/cometbft/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
)

const TypeUpdateRValidatorProposal = "update_r_validator"
const TypeUpdateRValidatorReportProposal = "update_r_validator_report"

var _ sdk.Msg = &UpdateRValidatorProposal{}
var _ sdk.Msg = &UpdateRValidatorReportProposal{}
var _ rvotetypes.Content = &UpdateRValidatorProposal{}
var _ rvotetypes.Content = &UpdateRValidatorReportProposal{}

func init() {
	rvotetypes.RegisterProposalType(TypeUpdateRValidatorProposal)
	rvotetypes.RegisterProposalTypeCodec(&UpdateRValidatorProposal{}, "rvalidator/UpdateRValidator")
	rvotetypes.RegisterProposalType(TypeUpdateRValidatorReportProposal)
	rvotetypes.RegisterProposalTypeCodec(&UpdateRValidatorReportProposal{}, "rvalidator/UpdateRValidatorReport")
}

func NewUpdateRValidatorProposal(creator string, denom string, poolAddress, oldAddress string, newAddress string, cycle *Cycle) *UpdateRValidatorProposal {
	msg := UpdateRValidatorProposal{
		Denom:       denom,
		PoolAddress: poolAddress,
		OldAddress:  oldAddress,
		NewAddress:  newAddress,
		Cycle:       cycle,
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
	if !(msg.Denom == msg.Cycle.Denom && msg.PoolAddress == msg.Cycle.PoolAddress) {
		return fmt.Errorf("denom or pool address not equal")
	}
	return nil
}

func NewUpdateRValidatorReportProposal(creator string, denom string, poolAddress string, cycle *Cycle, status UpdateRValidatorStatus) *UpdateRValidatorReportProposal {
	msg := UpdateRValidatorReportProposal{
		Denom:       denom,
		PoolAddress: poolAddress,
		Cycle:       cycle,
		Status:      status,
	}
	msg.setPropId()

	msg.Creator = creator

	return &msg
}

func (p *UpdateRValidatorReportProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
}

func (p *UpdateRValidatorReportProposal) ProposalRoute() string {
	return ModuleName
}

func (p *UpdateRValidatorReportProposal) ProposalType() string {
	return TypeUpdateRValidatorReportProposal
}

func (p *UpdateRValidatorReportProposal) InFavour() bool {
	return true
}

func (msg *UpdateRValidatorReportProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *UpdateRValidatorReportProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *UpdateRValidatorReportProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !(msg.Denom == msg.Cycle.Denom && msg.PoolAddress == msg.Cycle.PoolAddress) {
		return fmt.Errorf("denom or pool address not equal")
	}
	return nil
}
