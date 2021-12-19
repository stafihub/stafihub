package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewBondPipeline(denom, pool string) BondPipeline {
	return BondPipeline{
		Denom: denom,
		Pool: pool,
		Chunk: LinkChunk{
			Bond: SdkNewInt(),
			Unbond: SdkNewInt(),
			Active: SdkNewInt(),
		},
	}
}

func NewBondSnapshot(denom, pool string, era uint32, chunk LinkChunk, voter string) BondSnapshot {
	return BondSnapshot{
		Denom: denom,
		Pool: pool,
		Era: era,
		Chunk: chunk,
		LastVoter: voter,
		BondState: EraUpdated,
	}
}

func NewEraSnapShot(denom string) EraSnapShot {
	return EraSnapShot{
		Denom: denom,
		ShotIds: [][]byte{},
	}
}

func (bss BondSnapshot) UpdateState(state PoolBondState) {
	bss.BondState = state
}

func (bss BondSnapshot) Continuable() bool {
	return bss.BondState == WithdrawSkipped || bss.BondState == TransferReported
}

func SdkNewInt() sdk.Int {
	return sdk.NewInt(0)
}