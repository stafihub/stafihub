package types

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	rvotetypes "github.com/stafihub/stafihub/x/rvote/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	SetChainEraProposalType         = "SetChainEraProposal"
	BondReportProposalType          = "BondReportProposal"
	BondAndReportActiveProposalType = "BondAndReportActiveProposal"
	ActiveReportProposalType        = "ActiveReportProposal"
	WithdrawReportProposalType      = "WithdrawReportProposal"
	TransferReportProposalType      = "TransferReportProposal"
	ExecuteBondProposalType         = "ExecuteBondProposal"
)

func init() {
	rvotetypes.RegisterProposalType(SetChainEraProposalType)
	rvotetypes.RegisterProposalTypeCodec(&SetChainEraProposal{}, "ledger/SetChainEraProposal")
	rvotetypes.RegisterProposalType(BondReportProposalType)
	rvotetypes.RegisterProposalTypeCodec(&BondReportProposal{}, "ledger/BondReportProposal")
	rvotetypes.RegisterProposalType(BondAndReportActiveProposalType)
	rvotetypes.RegisterProposalType(ActiveReportProposalType)
	rvotetypes.RegisterProposalTypeCodec(&ActiveReportProposal{}, "ledger/ActiveReportProposal")
	rvotetypes.RegisterProposalType(WithdrawReportProposalType)
	rvotetypes.RegisterProposalType(TransferReportProposalType)
	rvotetypes.RegisterProposalTypeCodec(&TransferReportProposal{}, "ledger/TransferReportProposal")
	rvotetypes.RegisterProposalType(ExecuteBondProposalType)
	rvotetypes.RegisterProposalTypeCodec(&ExecuteBondProposal{}, "ledger/ExecuteBondProposal")
}

func NewSetChainEraProposal(proposer sdk.AccAddress, denom string, era uint32) *SetChainEraProposal {
	p := &SetChainEraProposal{
		Denom: denom,
		Era:   era,
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

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
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
	err := rvotetypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func NewBondReportProposal(proposer sdk.AccAddress, denom string, shotId string, action BondAction) *BondReportProposal {
	p := &BondReportProposal{
		Denom:  denom,
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

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
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
	err := rvotetypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func NewActiveReportProposal(proposer sdk.AccAddress, denom string, shotId string, staked, unstaked sdk.Int) *ActiveReportProposal {
	p := &ActiveReportProposal{
		Denom:    denom,
		ShotId:   shotId,
		Staked:   staked,
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

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
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
	err := rvotetypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func NewTransferReportProposal(proposer sdk.AccAddress, denom string, shotId string) *TransferReportProposal {
	p := &TransferReportProposal{
		Denom:  denom,
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

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
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
	err := rvotetypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func NewExecuteBondProposal(
	proposer sdk.AccAddress, denom string, bonder sdk.AccAddress,
	pool string, txhash string, amount sdk.Int, state LiquidityBondState) *ExecuteBondProposal {
	p := &ExecuteBondProposal{
		Denom:  denom,
		Bonder: bonder.String(),
		Pool:   pool,
		Txhash: txhash,
		Amount: amount,
		State:  state,
	}

	p.setPropId()
	p.Proposer = proposer.String()
	return p
}

func (p *ExecuteBondProposal) setPropId() {
	b, err := p.Marshal()
	if err != nil {
		panic(err)
	}

	p.PropId = hex.EncodeToString(crypto.Sha256(b))
}

func (p *ExecuteBondProposal) ProposalRoute() string {
	return ModuleName
}

func (p *ExecuteBondProposal) ProposalType() string {
	return TransferReportProposalType
}

func (p *ExecuteBondProposal) ValidateBasic() error {
	err := rvotetypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	if p.Bonder == "" {
		return ErrInvalidBonder
	}

	return nil
}
