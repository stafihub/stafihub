package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
)

func (k Keeper) AddPool(ctx sdk.Context, denom string, addr string) error {
	pool, found := k.TryFindPool(ctx, denom, addr, types.PoolPrefix)
	if found {
		return types.ErrPoolAlreadyAdded
	}

	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolPrefix)
	b := k.cdc.MustMarshal(pool)
	store.Set([]byte(denom), b)
	return nil
}

func (k Keeper) SetPool(ctx sdk.Context, pool *types.Pool, pref []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pref)
	b := k.cdc.MustMarshal(pool)
	store.Set([]byte(pool.Denom), b)
}

func (k Keeper) SetEraUnbondLimit(ctx sdk.Context, denom string, limit uint32) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.EraUnbondLimitPrefix)
	eul := &types.EraUnbondLimit{
		Denom: denom,
		Limit: limit,
	}
	b := k.cdc.MustMarshal(eul)
	store.Set([]byte(denom), b)
}

func (k Keeper) EraUnbondLimit(ctx sdk.Context, denom string) (val *types.EraUnbondLimit, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.EraUnbondLimitPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) SetInitBond(ctx sdk.Context, denom, pool string, amount sdk.Int, receiver sdk.AccAddress) error {
	// todo use cacheContext
	_, found := k.TryFindPool(ctx, denom, pool, types.PoolPrefix)
	if !found {
		return types.ErrPoolNotFound
	}

	bpool, found := k.TryFindPool(ctx, denom, pool, types.BondedPoolPrefix)
	if found {
		return types.ErrRepeatInitBond
	}

	rbalance := k.rateKeeper.TokenToRtoken(ctx, denom, amount)
	rcoins := sdk.Coins{
		sdk.NewCoin(denom, rbalance),
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, rcoins); err != nil {
		return err
	}

	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, rcoins)
	if err != nil {
		return err
	}

	if k.rateKeeper.GetRate(ctx, denom) == nil {
		k.rateKeeper.SetRate(ctx, denom, sdk.NewInt(0), sdk.NewInt(0))
	}

	pipe := types.NewBondPipeline(denom, pool)
	k.AddBondedPool(ctx, bpool)
	k.SetBondPipeline(ctx, pipe)
	return nil
}

func (k Keeper) AddBondedPool(ctx sdk.Context, pool *types.Pool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)
	b := k.cdc.MustMarshal(pool)
	store.Set([]byte(pool.Denom), b)
}

func (k Keeper) GetBondedPoolByDenom(ctx sdk.Context, denom string) (val types.Pool, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetChainBondingDuration(ctx sdk.Context, denom string, era uint32) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainBondingDurationPrefix)
	cbd := &types.ChainBondingDuration{
		Denom: denom,
		Era: era,
	}
	b := k.cdc.MustMarshal(cbd)
	store.Set([]byte(denom), b)
}

func (k Keeper) ChainBondingDuration(ctx sdk.Context, denom string) (val *types.ChainBondingDuration, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainBondingDurationPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) SetPoolDetail(ctx sdk.Context, denom string, pool string, subAccounts []string, threshold uint32) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolDetailPrefix)
	cbd := &types.PoolDetail{
		Denom: denom,
		Pool: pool,
		SubAccounts: subAccounts,
		Threshold: threshold,
	}
	b := k.cdc.MustMarshal(cbd)
	store.Set([]byte(denom+pool), b)
}

func (k Keeper) PoolDetail(ctx sdk.Context, denom string, pool string) (val *types.PoolDetail, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolDetailPrefix)

	b := store.Get([]byte(denom+pool))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) SetLeastBond(ctx sdk.Context, denom string, amount sdk.Int) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.LeastBondPrefix)
	lb := &types.LeastBond{
		Denom: denom,
		Amount: amount,
	}
	b := k.cdc.MustMarshal(lb)
	store.Set([]byte(denom), b)
}

func (k Keeper) LeastBond(ctx sdk.Context, denom string) (val *types.LeastBond, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.LeastBondPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) SetCurrentEraSnapShot(ctx sdk.Context, shot types.EraSnapShot) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.CurrentEraSnapShotPrefix)
	b := k.cdc.MustMarshal(&shot)
	store.Set([]byte(shot.Denom), b)
}

func (k Keeper) CurrentEraSnapShots(ctx sdk.Context, denom string) types.EraSnapShot {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.CurrentEraSnapShotPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return types.NewEraSnapShot(denom)
	}

	var val types.EraSnapShot
	k.cdc.MustUnmarshal(b, &val)
	return val
}

func (k Keeper) ClearCurrentEraSnapShots(ctx sdk.Context, denom string) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.CurrentEraSnapShotPrefix)

	cess := &types.EraSnapShot{
		Denom: denom,
		ShotIds: [][]byte{},
	}
	b := k.cdc.MustMarshal(cess)
	store.Set([]byte(denom), b)
}

func (k Keeper) SetSnapShot(ctx sdk.Context, shotId []byte, shot types.BondSnapshot)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.SnapShotPrefix)
	b := k.cdc.MustMarshal(&shot)
	store.Set(shotId, b)
}

func (k Keeper) SnapShot(ctx sdk.Context, shotId []byte) (val types.BondSnapshot, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.SnapShotPrefix)

	b := store.Get(shotId)
	if b == nil {
		return types.BondSnapshot{}, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetEraSnapShot(ctx sdk.Context, era uint32, shot types.EraSnapShot) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.EraSnapShotPrefix)

	bera:= make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	b := k.cdc.MustMarshal(&shot)
	key := append([]byte(shot.Denom), bera...)
	store.Set(key, b)
}

func (k Keeper) EraSnapShot(ctx sdk.Context, denom string, era uint32) (val types.EraSnapShot) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.EraSnapShotPrefix)
	bera:= make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom), bera...)
	b := store.Get(key)
	if b == nil {
		return types.NewEraSnapShot(denom)
	}

	k.cdc.MustUnmarshal(b, &val)
	return
}

func (k Keeper) SetBondPipeline(ctx sdk.Context, pipe types.BondPipeline) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.BondPipelinePrefix)
	b := k.cdc.MustMarshal(&pipe)
	store.Set([]byte(pipe.Denom), b)
}

func (k Keeper) BondPipeLine(ctx sdk.Context, denom string, pool string) (val types.BondPipeline, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.BondPipelinePrefix)

	b := store.Get([]byte(denom+pool))
	if b == nil {
		return types.NewBondPipeline(denom, pool), false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetChainEra(ctx sdk.Context, denom string, era uint32) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainEraPrefix)
	ce := &types.ChainEra{
		Denom: denom,
		Era: era,
	}

	b := k.cdc.MustMarshal(ce)
	store.Set([]byte(denom), b)
}

func (k Keeper) ChainEra(ctx sdk.Context, denom string) (val types.ChainEra) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainEraPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return types.ChainEra{
			Denom: denom,
			Era: 0,
		}
	}

	k.cdc.MustUnmarshal(b, &val)
	return val
}

func (k Keeper) TryFindPool(ctx sdk.Context, denom, addr string, pref []byte) (val *types.Pool, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), pref)

	b := store.Get([]byte(denom))
	if b == nil {
		return &types.Pool{
			Denom: denom,
			Addrs: map[string]bool{addr: true},
		}, false
	}

	k.cdc.MustUnmarshal(b, val)
	if _, ok := val.Addrs[addr]; ok {
		return val, true
	}

	val.Addrs[addr] = true
	return val, false
}

func (k Keeper) SetCommission(ctx sdk.Context, commission sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	b, _ := commission.Marshal()
	store.Set(types.CommissionPrefix, b)
}

func (k Keeper) Commission(ctx sdk.Context) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.CommissionPrefix)
	if b == nil {
		return sdk.OneDec()
	}

	var val sdk.Dec
	if err := val.Unmarshal(b); err != nil {
		panic(err)
	}

	return val
}

func (k Keeper) SetReceiver(ctx sdk.Context, receiver sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ReceiverPrefix, receiver)
}

func (k Keeper) Receiver(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.CommissionPrefix)
}

func (k Keeper) SetTotalExpectedActive(ctx sdk.Context, denom string, era uint32, active sdk.Int) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.TotalExpectedActivePrefix)

	bera:= make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom), bera...)

	b, _ := active.Marshal()
	store.Set(key, b)
}

func (k Keeper) TotalExpectedActive(ctx sdk.Context, denom string, era uint32) (val sdk.Int) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.TotalExpectedActivePrefix)
	bera:= make([]byte, 4)
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

func (k Keeper) SetPoolUnbond(ctx sdk.Context, denom string, pool string, era uint32, unbondings []types.Unbonding) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolUnbondPrefix)
	pu := types.PoolUnbond{
		Denom: denom,
		Pool: pool,
		Era: era,
		Unbondings: unbondings,
	}
	b := k.cdc.MustMarshal(&pu)
	bera:= make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom+pool), bera...)
	store.Set(key, b)
}

func (k Keeper) PoolUnbond(ctx sdk.Context, denom string, pool string, era uint32) (val types.PoolUnbond, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolUnbondPrefix)
	bera:= make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom+pool), bera...)
	b := store.Get(key)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetPoolByDenom(ctx sdk.Context, denom string) (val *types.Pool, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}






