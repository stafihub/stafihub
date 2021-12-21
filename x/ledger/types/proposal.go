package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	relayertypes "github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	SetChainEraProposalType = "SetChainEraProposal"
	BondReportProposalType = "BondReportProposal"
	BondAndReportActiveProposalType = "BondAndReportActiveProposal"
	ActiveReportProposalType = "ActiveReportProposal"
	WithdrawReportProposalType = "WithdrawReportProposal"
	TransferReportProposalType = "TransferReportProposal"
)

//todo add init() to register proposal types

func NewSetChainEraProposal(proposer sdk.AccAddress, denom string, era uint32) *SetChainEraProposal {
	p := &SetChainEraProposal{
		Denom: denom,
		Era: era,
	}

	p.setPropId()
	p.Proposer = proposer.String()
	return p
}

func (p *SetChainEraProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = crypto.Sha256(b)
}

func (p *SetChainEraProposal) ProposalRoute() string {
	return ModuleName
}

func (p *SetChainEraProposal) ProposalType() string {
	return SetChainEraProposalType
}

func (p *SetChainEraProposal) InFavour() bool {
	return true
}

func (p *SetChainEraProposal) ValidateBasic() error {
	err := relayertypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func NewBondReportProposal(proposer sdk.AccAddress, denom string, shotId []byte, action BondAction) *BondReportProposal {
	p := &BondReportProposal{
		Denom: denom,
		ShotId: shotId,
		Action: action,
	}

	p.setPropId()
	p.Proposer = proposer.String()
	return p
}

func (p *BondReportProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = crypto.Sha256(b)
}

func (p *BondReportProposal) ProposalRoute() string {
	return ModuleName
}

func (p *BondReportProposal) ProposalType() string {
	return BondReportProposalType
}

func (p *BondReportProposal) InFavour() bool {
	return true
}

func (p *BondReportProposal) ValidateBasic() error {
	err := relayertypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func NewBondAndReportActiveProposal(proposer sdk.AccAddress, denom string, shotId []byte, action BondAction, staked, unstaked sdk.Int) *BondAndReportActiveProposal {
	p := &BondAndReportActiveProposal{
		Denom: denom,
		ShotId: shotId,
		Action: action,
		Staked: staked,
		Unstaked: unstaked,
	}

	p.setPropId()
	p.Proposer = proposer.String()
	return p
}

func (p *BondAndReportActiveProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = crypto.Sha256(b)
}

func (p *BondAndReportActiveProposal) ProposalRoute() string {
	return ModuleName
}

func (p *BondAndReportActiveProposal) ProposalType() string {
	return BondAndReportActiveProposalType
}

func (p *BondAndReportActiveProposal) InFavour() bool {
	return true
}

func (p *BondAndReportActiveProposal) ValidateBasic() error {
	err := relayertypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func NewActiveReportProposal(proposer sdk.AccAddress, denom string, shotId []byte, staked, unstaked sdk.Int) *ActiveReportProposal {
	p := &ActiveReportProposal{
		Denom: denom,
		ShotId: shotId,
		Staked: staked,
		Unstaked: unstaked,
	}

	p.setPropId()
	p.Proposer = proposer.String()
	return p
}

func (p *ActiveReportProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = crypto.Sha256(b)
}

func (p *ActiveReportProposal) ProposalRoute() string {
	return ModuleName
}

func (p *ActiveReportProposal) ProposalType() string {
	return ActiveReportProposalType
}

func (p *ActiveReportProposal) InFavour() bool {
	return true
}

func (p *ActiveReportProposal) ValidateBasic() error {
	err := relayertypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func NewWithdrawReportProposal(proposer sdk.AccAddress, denom string, shotId []byte) *WithdrawReportProposal {
	p := &WithdrawReportProposal{
		Denom: denom,
		ShotId: shotId,
	}

	p.setPropId()
	p.Proposer = proposer.String()
	return p
}

func (p *WithdrawReportProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = crypto.Sha256(b)
}

func (p *WithdrawReportProposal) ProposalRoute() string {
	return ModuleName
}

func (p *WithdrawReportProposal) ProposalType() string {
	return WithdrawReportProposalType
}

func (p *WithdrawReportProposal) InFavour() bool {
	return true
}

func (p *WithdrawReportProposal) ValidateBasic() error {
	err := relayertypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func NewTransferReportProposal(proposer sdk.AccAddress, denom string, shotId []byte) *TransferReportProposal {
	p := &TransferReportProposal{
		Denom: denom,
		ShotId: shotId,
	}

	p.setPropId()
	p.Proposer = proposer.String()
	return p
}

func (p *TransferReportProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = crypto.Sha256(b)
}

func (p *TransferReportProposal) ProposalRoute() string {
	return ModuleName
}

func (p *TransferReportProposal) ProposalType() string {
	return TransferReportProposalType
}

func (p *TransferReportProposal) InFavour() bool {
	return true
}

func (p *TransferReportProposal) ValidateBasic() error {
	err := relayertypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}








