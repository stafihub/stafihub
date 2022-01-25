package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewPool(denom string) Pool {
	return Pool{
		Denom: denom,
		Addrs: map[string]bool{},
	}
}

func NewChainEra(denom string) ChainEra {
	return ChainEra{
		Denom: denom,
		Era:   0,
	}
}

func NewAccountUnbond(denom, unbonder string, chunks []UserUnlockChunk) AccountUnbond {
	return AccountUnbond{
		Unbonder: unbonder,
		Denom:    denom,
		Chunks:   chunks,
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

func NewBondSnapshot(denom, pool string, era uint32, chunk LinkChunk, voter string) BondSnapshot {
	return BondSnapshot{
		Denom:     denom,
		Pool:      pool,
		Era:       era,
		Chunk:     chunk,
		LastVoter: voter,
		BondState: EraUpdated,
	}
}

func NewEraSnapshot(denom string) EraSnapshot {
	return EraSnapshot{
		Denom:   denom,
		ShotIds: [][]byte{},
	}
}

func NewBondRecord(denom, bonder, pool, blockhash, txhash string, amount sdk.Int) BondRecord {
	return BondRecord{
		Denom:     denom,
		Bonder:    bonder,
		Pool:      pool,
		Blockhash: blockhash,
		Txhash:    txhash,
		Amount:    amount,
	}
}

func NewSignature(denom string, era uint32, pool string,
	txType OriginalTxType, propId []byte) Signature {
	return Signature{
		Denom:   denom,
		Era:     era,
		Pool:    pool,
		TxType:  txType,
		PropId:  propId,
		Sigs:    []string{},
		Signers: map[string]string{},
	}
}

func (bss BondSnapshot) UpdateState(state PoolBondState) {
	// todo need to test if the change was kept
	bss.BondState = state
}

func (bss BondSnapshot) Continuable() bool {
	return bss.BondState == WithdrawSkipped || bss.BondState == TransferReported
}

func SdkNewInt() sdk.Int {
	return sdk.NewInt(0)
}
