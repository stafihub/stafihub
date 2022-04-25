package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewPoolDetail(denom, pool string, subAccounts []string, threshold uint32) PoolDetail {
	return PoolDetail{
		Denom:       denom,
		Pool:        pool,
		SubAccounts: subAccounts,
		Threshold:   threshold,
	}
}

func NewChainEra(denom string) ChainEra {
	return ChainEra{
		Denom: denom,
		Era:   0,
	}
}

func NewPoolUnbond(denom, pool string, era uint32, unbondings []Unbonding) PoolUnbond {
	return PoolUnbond{
		Denom:      denom,
		Pool:       pool,
		Era:        era,
		Unbondings: unbondings,
	}
}

func NewUnbonding(unbonder, recipient string, amount sdk.Int) Unbonding {
	return Unbonding{
		Unbonder:  unbonder,
		Amount:    amount,
		Recipient: recipient,
	}
}

func NewBondPipeline(denom, pool string) BondPipeline {
	return BondPipeline{
		Denom: denom,
		Pool:  pool,
		Chunk: LinkChunk{
			Bond:   SdkNewInt(),
			Unbond: SdkNewInt(),
			Active: SdkNewInt(),
		},
	}
}

func NewBondSnapshot(denom, pool string, era uint32, chunk LinkChunk) BondSnapshot {
	return BondSnapshot{
		Denom:     denom,
		Pool:      pool,
		Era:       era,
		Chunk:     chunk,
		BondState: EraUpdated,
	}
}

func NewEraSnapshot(denom string) EraSnapshot {
	return EraSnapshot{
		Denom:   denom,
		ShotIds: []string{},
	}
}

func NewBondRecord(denom, bonder, pool, txhash string, amount sdk.Int, state LiquidityBondState) BondRecord {
	return BondRecord{
		Denom:  denom,
		Bonder: bonder,
		Pool:   pool,
		Txhash: txhash,
		Amount: amount,
		State:  state,
	}
}

func NewSignature(denom string, era uint32, pool string,
	txType OriginalTxType, propId string) Signature {
	return Signature{
		Denom:  denom,
		Era:    era,
		Pool:   pool,
		TxType: txType,
		PropId: propId,
		Sigs:   []string{},
	}
}

func (bss *BondSnapshot) UpdateState(state PoolBondState) {
	bss.BondState = state
}

func (bss BondSnapshot) Continuable() bool {
	return bss.BondState == TransferSkipped || bss.BondState == TransferReported
}

func SdkNewInt() sdk.Int {
	return sdk.NewInt(0)
}
