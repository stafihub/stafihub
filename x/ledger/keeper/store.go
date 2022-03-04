package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
)

var (
	defaultStakingRewardCommission = sdk.MustNewDecFromStr("0.1")
	defaultUnbondCommission        = sdk.MustNewDecFromStr("0.002")
	defaultUnbondFee               = sdk.NewCoin("ufis", sdk.ZeroInt())
	defaultEraUnbondLimit          = uint32(200)
)

func (k Keeper) IsBondedPoolExist(ctx sdk.Context, denom string, addr string) bool {
	pool, ok := k.GetBondedPool(ctx, denom)
	if !ok {
		return false
	}

	for _, adr := range pool.Addrs {
		if adr == addr {
			return true
		}
	}

	return false
}

func (k Keeper) GetBondedPool(ctx sdk.Context, denom string) (pool types.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)
	b := store.Get([]byte(denom))
	val := types.Pool{Denom: denom, Addrs: []string{}}
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) RemoveBondedPool(ctx sdk.Context, denom string, addr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)
	pool, ok := k.GetBondedPool(ctx, denom)
	if !ok {
		return
	}

	addrs := make([]string, 0)
	for _, adr := range pool.Addrs {
		if adr != addr {
			addrs = append(addrs, adr)
		}
	}
	pool.Addrs = addrs
	b := k.cdc.MustMarshal(&pool)
	store.Set([]byte(denom), b)
}

func (k Keeper) AddBondedPool(ctx sdk.Context, denom string, addr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)
	pool, _ := k.GetBondedPool(ctx, denom)
	pool.Addrs = append(pool.Addrs, addr)
	b := k.cdc.MustMarshal(&pool)
	store.Set([]byte(denom), b)
}

func (k Keeper) SetBondPipeline(ctx sdk.Context, pipe types.BondPipeline) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondPipelinePrefix)
	b := k.cdc.MustMarshal(&pipe)
	store.Set([]byte(pipe.Denom+pipe.Pool), b)
}

func (k Keeper) GetBondPipeline(ctx sdk.Context, denom string, pool string) (val types.BondPipeline, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondPipelinePrefix)

	b := store.Get([]byte(denom + pool))
	if b == nil {
		return types.NewBondPipeline(denom, pool), false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetEraUnbondLimit(ctx sdk.Context, denom string, limit uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EraUnbondLimitPrefix)
	eul := &types.EraUnbondLimit{
		Denom: denom,
		Limit: limit,
	}
	b := k.cdc.MustMarshal(eul)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetEraUnbondLimit(ctx sdk.Context, denom string) (val types.EraUnbondLimit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EraUnbondLimitPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		val.Denom = denom
		val.Limit = defaultEraUnbondLimit
		return val
	}
	k.cdc.MustUnmarshal(b, &val)
	return val
}

func (k Keeper) SetPoolDetail(ctx sdk.Context, denom string, pool string, subAccounts []string, threshold uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolDetailPrefix)
	cbd := types.NewPoolDetail(denom, pool, subAccounts, threshold)
	b := k.cdc.MustMarshal(&cbd)
	store.Set([]byte(denom+pool), b)
}

func (k Keeper) GetPoolDetail(ctx sdk.Context, denom string, pool string) (val types.PoolDetail, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolDetailPrefix)

	b := store.Get([]byte(denom + pool))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetCurrentEraSnapshot(ctx sdk.Context, shot types.EraSnapshot) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CurrentEraSnapshotPrefix)
	b := k.cdc.MustMarshal(&shot)
	store.Set([]byte(shot.Denom), b)
}

func (k Keeper) CurrentEraSnapshots(ctx sdk.Context, denom string) types.EraSnapshot {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CurrentEraSnapshotPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return types.NewEraSnapshot(denom)
	}

	var val types.EraSnapshot
	k.cdc.MustUnmarshal(b, &val)
	if val.ShotIds == nil {
		return types.NewEraSnapshot(denom)
	}
	return val
}

func (k Keeper) ClearCurrentEraSnapshots(ctx sdk.Context, denom string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CurrentEraSnapshotPrefix)
	shot := types.NewEraSnapshot(denom)
	b := k.cdc.MustMarshal(&shot)
	store.Set([]byte(denom), b)
}

func (k Keeper) SetSnapshot(ctx sdk.Context, shotId string, shot types.BondSnapshot) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SnapshotPrefix)
	b := k.cdc.MustMarshal(&shot)
	store.Set([]byte(shotId), b)
}

func (k Keeper) Snapshot(ctx sdk.Context, shotId string) (val types.BondSnapshot, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SnapshotPrefix)

	b := store.Get([]byte(shotId))
	if b == nil {
		return types.BondSnapshot{}, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetEraSnapshot(ctx sdk.Context, era uint32, shot types.EraSnapshot) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EraSnapshotPrefix)

	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	b := k.cdc.MustMarshal(&shot)
	key := append([]byte(shot.Denom), bera...)
	store.Set(key, b)
}

func (k Keeper) EraSnapshot(ctx sdk.Context, denom string, era uint32) (val types.EraSnapshot) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.EraSnapshotPrefix)
	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom), bera...)
	b := store.Get(key)
	if b == nil {
		return types.NewEraSnapshot(denom)
	}

	k.cdc.MustUnmarshal(b, &val)
	return
}

func (k Keeper) SetChainEra(ctx sdk.Context, denom string, era uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainEraPrefix)
	ce := &types.ChainEra{
		Denom: denom,
		Era:   era,
	}

	b := k.cdc.MustMarshal(ce)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetChainEra(ctx sdk.Context, denom string) (val types.ChainEra, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainEraPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetStakingRewardCommission(ctx sdk.Context, denom string, commission sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	b, err := commission.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set(types.StakingRewardCommissionStoreKey(denom), b)
}

func (k Keeper) GetStakingRewardCommission(ctx sdk.Context, denom string) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.StakingRewardCommissionStoreKey(denom))
	if b == nil {
		return defaultStakingRewardCommission
	}

	var val sdk.Dec
	if err := val.Unmarshal(b); err != nil {
		panic(err)
	}

	return val
}

func (k Keeper) SetProtocolFeeReceiver(ctx sdk.Context, receiver sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ProtocolFeeReceiverPrefix, receiver)
}

func (k Keeper) GetProtocolFeeReceiver(ctx sdk.Context) (sdk.AccAddress, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ProtocolFeeReceiverPrefix)
	if bts == nil {
		return nil, false
	}
	return bts, true
}

func (k Keeper) SetRelayFeeReceiver(ctx sdk.Context, denom string, receiver sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RelayFeeReceiverStorekey(denom), receiver)
}

func (k Keeper) GetRelayFeeReceiver(ctx sdk.Context, denom string) (sdk.AccAddress, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.RelayFeeReceiverStorekey(denom))
	if len(bts) == 0 {
		return nil, false
	}
	return bts, true
}

func (k Keeper) SetTotalExpectedActive(ctx sdk.Context, denom string, era uint32, active sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TotalExpectedActivePrefix)

	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom), bera...)

	b, _ := active.Marshal()
	store.Set(key, b)
}

func (k Keeper) TotalExpectedActive(ctx sdk.Context, denom string, era uint32) (val sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TotalExpectedActivePrefix)
	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom), bera...)
	b := store.Get(key)
	if b == nil {
		return sdk.NewInt(0)
	}

	if err := val.Unmarshal(b); err != nil {
		panic(err)
	}

	return val
}

func (k Keeper) SetPoolUnbond(ctx sdk.Context, pu types.PoolUnbond) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolUnbondPrefix)
	b := k.cdc.MustMarshal(&pu)
	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, pu.Era)
	key := append([]byte(pu.Denom+pu.Pool), bera...)
	store.Set(key, b)
}

func (k Keeper) GetPoolUnbond(ctx sdk.Context, denom string, pool string, era uint32) (val types.PoolUnbond, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolUnbondPrefix)
	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom+pool), bera...)
	b := store.Get(key)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	if val.Unbondings == nil {
		val.Unbondings = []types.Unbonding{}
	}
	return val, true
}

func (k Keeper) SetUnbondRelayFee(ctx sdk.Context, denom string, value sdk.Coin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UnbondFeePrefix)
	uf := types.UnbondRelayFee{
		Denom: denom,
		Value: value,
	}

	b := k.cdc.MustMarshal(&uf)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetUnbondRelayFee(ctx sdk.Context, denom string) (val types.UnbondRelayFee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UnbondFeePrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		val.Denom = denom
		val.Value = defaultUnbondFee
		return
	}

	k.cdc.MustUnmarshal(b, &val)
	return val
}

func (k Keeper) SetUnbondCommission(ctx sdk.Context, denom string, value sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	b, err := value.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set(types.UnbondCommissionStoreKey(denom), b)
}

func (k Keeper) GetUnbondCommission(ctx sdk.Context, denom string) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.UnbondCommissionStoreKey(denom))
	if b == nil {
		return defaultUnbondCommission
	}
	var val sdk.Dec
	if err := val.Unmarshal(b); err != nil {
		panic(err)
	}
	return val
}

func (k Keeper) SetAccountUnbond(ctx sdk.Context, unbond types.AccountUnbond) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountUnbondPrefix)

	b := k.cdc.MustMarshal(&unbond)
	store.Set([]byte(unbond.Denom+unbond.Unbonder), b)
}

func (k Keeper) GetAccountUnbond(ctx sdk.Context, denom, unbonder string) (val types.AccountUnbond, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountUnbondPrefix)
	b := store.Get([]byte(denom + unbonder))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	if val.Chunks == nil {
		val.Chunks = []types.UserUnlockChunk{}
	}
	return val, true
}

func (k Keeper) SetBondRecord(ctx sdk.Context, br types.BondRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondRecordPrefix)
	b := k.cdc.MustMarshal(&br)
	store.Set([]byte(br.Denom+br.Txhash), b)
}

func (k Keeper) GetBondRecord(ctx sdk.Context, denom, txhash string) (val types.BondRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondRecordPrefix)
	b := store.Get([]byte(denom + txhash))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetSignature(ctx sdk.Context, sig types.Signature) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SignaturePrefix)
	b := k.cdc.MustMarshal(&sig)

	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, sig.Era)
	key := append([]byte(sig.Denom+sig.Pool+sig.TxType.String()), bera...)
	key = append(key, sig.PropId...)

	store.Set(key, b)
}

func (k Keeper) GetSignature(ctx sdk.Context, denom string, era uint32, pool string,
	txType types.OriginalTxType, propId string) (types.Signature, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SignaturePrefix)
	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom+pool+txType.String()), bera...)
	key = append(key, propId...)

	val := types.NewSignature(denom, era, pool, txType, propId)
	b := store.Get(key)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetRParams(ctx sdk.Context, rParams types.RParams) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RParamsPrefix)
	b := k.cdc.MustMarshal(&rParams)
	store.Set([]byte(rParams.Denom), b)
}

func (k Keeper) GetRParams(ctx sdk.Context, denom string) (val types.RParams, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RParamsPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
