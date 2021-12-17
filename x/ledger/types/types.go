package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewBondPipeline(denom, pool string) BondPipeline {
	return BondPipeline{
		Denom: denom,
		Pool: pool,
		Chunk: &LinkChunk{
			Bond: SdkNewInt(),
			Unbond: SdkNewInt(),
			Active: SdkNewInt(),
		},
	}
}

func NewBondSnapshot(denom, pool string, era uint32, chunk *LinkChunk, voter string) BondSnapshot {
	return BondSnapshot{
		Denom: denom,
		Pool: pool,
		Era: era,
		Chunk: chunk,
		LastVoter: voter,
		BondState: EraUpdated,
	}
}

func NewEraSnapShot(denom string) CurrentEraSnapShot {
	return CurrentEraSnapShot{
		Denom: denom,
		ShotIds: [][]byte{},
	}
}


func SdkNewInt() *sdk.Int {
	i := sdk.NewInt(0)
	return &i
}