package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/ledger/types"
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

func (k Keeper) SetBondedPool(ctx sdk.Context, pool *types.Pool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)
	b := k.cdc.MustMarshal(pool)
	store.Set([]byte(pool.Denom), b)
}

func (k Keeper) AddBondedPool(ctx sdk.Context, denom string, addr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)
	pool, _ := k.GetBondedPool(ctx, denom)
	pool.Addrs = append(pool.Addrs, addr)
	b := k.cdc.MustMarshal(&pool)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetBondedPoolList(ctx sdk.Context) []*types.Pool {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.BondedPoolPrefix)
	defer iterator.Close()

	list := make([]*types.Pool, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denom := string(key[1:])
		pool, found := k.GetBondedPool(ctx, denom)
		if !found {
			continue
		}
		list = append(list, &pool)
	}
	return list
}

func (k Keeper) SetBondPipeline(ctx sdk.Context, pipe types.BondPipeline) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.BondPipelineStoreKey(pipe.Denom, pipe.Pool), k.cdc.MustMarshal(&pipe))
}

func (k Keeper) GetBondPipeline(ctx sdk.Context, denom string, pool string) (val types.BondPipeline, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.BondPipelineStoreKey(denom, pool))
	if b == nil {
		return types.NewBondPipeline(denom, pool), false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetBondPipelineList(ctx sdk.Context) []*types.BondPipeline {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.BondPipelinePrefix)
	defer iterator.Close()

	list := make([]*types.BondPipeline, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denomLen := int(key[1])
		denom := string(key[2 : 2+denomLen])
		pool := string(key[2+denomLen:])
		pipeline, found := k.GetBondPipeline(ctx, denom, pool)
		if !found {
			continue
		}
		list = append(list, &pipeline)
	}
	return list
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
		val.Limit = types.DefaultEraUnbondLimit
		return val
	}
	k.cdc.MustUnmarshal(b, &val)
	return val
}

func (k Keeper) GetEraUnbondLimitList(ctx sdk.Context) []*types.EraUnbondLimit {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.EraUnbondLimitPrefix)
	defer iterator.Close()

	list := make([]*types.EraUnbondLimit, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denom := string(key[1:])
		unbondLimit := k.GetEraUnbondLimit(ctx, denom)
		list = append(list, &unbondLimit)
	}
	return list
}

func (k Keeper) SetPoolDetail(ctx sdk.Context, denom string, pool string, subAccounts []string, threshold uint32) {
	store := ctx.KVStore(k.storeKey)
	cbd := types.NewPoolDetail(denom, pool, subAccounts, threshold)
	b := k.cdc.MustMarshal(&cbd)
	store.Set(types.PoolDetailStoreKey(denom, pool), b)
}

func (k Keeper) GetPoolDetail(ctx sdk.Context, denom string, pool string) (val types.PoolDetail, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.PoolDetailStoreKey(denom, pool))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetPoolDetailList(ctx sdk.Context) []*types.PoolDetail {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PoolDetailPrefix)
	defer iterator.Close()

	list := make([]*types.PoolDetail, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denomLen := int(key[1])
		denom := string(key[2 : 2+denomLen])
		pool := string(key[2+denomLen:])
		poolDetail, found := k.GetPoolDetail(ctx, denom, pool)
		if !found {
			continue
		}
		list = append(list, &poolDetail)
	}
	return list
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

func (k Keeper) CurrentEraSnapshotList(ctx sdk.Context) []*types.EraSnapshot {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.CurrentEraSnapshotPrefix)
	defer iterator.Close()

	list := make([]*types.EraSnapshot, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denom := string(key[1:])
		snapShots := k.CurrentEraSnapshots(ctx, denom)
		list = append(list, &snapShots)
	}
	return list
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

func (k Keeper) SnapshotList(ctx sdk.Context) []*types.GenesisSnapshot {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SnapshotPrefix)
	defer iterator.Close()

	list := make([]*types.GenesisSnapshot, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		shotId := string(key[1:])
		snapShot, found := k.Snapshot(ctx, shotId)
		if !found {
			continue
		}
		list = append(list, &types.GenesisSnapshot{
			ShotId:   shotId,
			Snapshot: &snapShot,
		})
	}
	return list
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

func (k Keeper) EraSnapshotList(ctx sdk.Context) []*types.GenesisEraSnapshot {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.EraSnapshotPrefix)
	defer iterator.Close()

	list := make([]*types.GenesisEraSnapshot, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		eraBts := key[len(key)-4:]
		era := binary.LittleEndian.Uint32(eraBts)

		denom := string(key[1 : len(key)-4])
		snapShot := k.EraSnapshot(ctx, denom, era)

		list = append(list, &types.GenesisEraSnapshot{
			Era:     era,
			Denom:   denom,
			ShotIds: snapShot.ShotIds,
		})
	}
	return list
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

func (k Keeper) GetChainEraList(ctx sdk.Context) []*types.ChainEra {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainEraPrefix)
	defer iterator.Close()

	list := make([]*types.ChainEra, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		denom := string(key[1:])
		chainEra, found := k.GetChainEra(ctx, denom)
		if !found {
			continue
		}
		list = append(list, &chainEra)
	}
	return list
}

func (k Keeper) SetStakingRewardCommission(ctx sdk.Context, denom string, commission utils.Dec) {
	store := ctx.KVStore(k.storeKey)
	b, err := commission.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set(types.StakingRewardCommissionStoreKey(denom), b)
}

func (k Keeper) GetStakingRewardCommission(ctx sdk.Context, denom string) utils.Dec {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.StakingRewardCommissionStoreKey(denom))
	if b == nil {
		return types.DefaultStakingRewardCommission
	}

	var val utils.Dec
	if err := val.Unmarshal(b); err != nil {
		panic(err)
	}

	return val
}

func (k Keeper) GetStakingRewardCommissionList(ctx sdk.Context) []*types.StakingRewardCommission {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.StakingRewardCommissionPrefix)
	defer iterator.Close()

	list := make([]*types.StakingRewardCommission, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		denom := string(key[1:])
		value := k.GetStakingRewardCommission(ctx, denom)

		list = append(list, &types.StakingRewardCommission{
			Denom: denom,
			Value: value,
		})
	}
	return list
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

func (k Keeper) GetRelayFeeReceiverList(ctx sdk.Context) []*types.RelayFeeReceiver {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RelayFeeReceiverPrefix)
	defer iterator.Close()

	list := make([]*types.RelayFeeReceiver, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denom := string(key[1:])
		addr, found := k.GetRelayFeeReceiver(ctx, denom)
		if !found {
			continue
		}

		list = append(list, &types.RelayFeeReceiver{
			Denom:   denom,
			Address: addr.String(),
		})
	}
	return list
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

func (k Keeper) TotalExpectedActiveList(ctx sdk.Context) []*types.TotalExpectedActive {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.TotalExpectedActivePrefix)
	defer iterator.Close()

	list := make([]*types.TotalExpectedActive, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		eraBts := key[len(key)-4:]
		era := binary.LittleEndian.Uint32(eraBts)

		denom := string(key[1 : len(key)-4])
		expectedActive := k.TotalExpectedActive(ctx, denom, era)

		list = append(list, &types.TotalExpectedActive{
			Denom: denom,
			Era:   era,
			Value: expectedActive,
		})
	}
	return list
}

func (k Keeper) SetPoolUnbonding(ctx sdk.Context, denom string, pool string, era, seq uint32, pu *types.Unbonding) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PoolUnbondStoreKey(denom, pool, era, seq), k.cdc.MustMarshal(pu))
}

func (k Keeper) GetPoolUnbonding(ctx sdk.Context, denom string, pool string, era, seq uint32) (*types.Unbonding, bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.PoolUnbondStoreKey(denom, pool, era, seq))
	if b == nil {
		return nil, false
	}

	poolUnbond := types.Unbonding{}
	k.cdc.MustUnmarshal(b, &poolUnbond)
	return &poolUnbond, true
}

func (k Keeper) GetPoolUnbondingList(ctx sdk.Context) []*types.GenesisPoolUnbonding {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PoolUnbondPrefix)
	defer iterator.Close()

	list := make([]*types.GenesisPoolUnbonding, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denomLen := int(key[1])
		poolLen := int(key[1+denomLen+1])
		denom := string(key[2 : 2+denomLen])
		pool := string(key[2+denomLen+1 : 2+denomLen+1+poolLen])
		era := binary.LittleEndian.Uint32(key[2+denomLen+1+poolLen : 2+denomLen+1+poolLen+4])
		seq := binary.LittleEndian.Uint32(key[2+denomLen+1+poolLen+4:])

		unbonding, found := k.GetPoolUnbonding(ctx, denom, pool, era, seq)
		if !found {
			continue
		}

		list = append(list, &types.GenesisPoolUnbonding{
			Denom:     denom,
			Era:       era,
			Pool:      pool,
			Sequence:  seq,
			Unbonding: unbonding,
		})
	}
	return list
}

func (k Keeper) GetPoolUnbondNextSequence(ctx sdk.Context, denom string, pool string, era uint32) uint32 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolUnbondNextSequencePrefix)

	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom+pool), bera...)

	seqBts := store.Get(key)
	if seqBts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(seqBts)
	return seq + 1
}

func (k Keeper) SetPoolUnbondSequence(ctx sdk.Context, denom string, pool string, era, seq uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolUnbondNextSequencePrefix)

	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	key := append([]byte(denom+pool), bera...)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, seq)
	store.Set(key, seqBts)
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
		val.Value = types.DefaultUnbondRelayFee
		return
	}

	k.cdc.MustUnmarshal(b, &val)
	return val
}

func (k Keeper) GetUnbondRelayFeeList(ctx sdk.Context) []*types.UnbondRelayFee {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UnbondFeePrefix)
	defer iterator.Close()

	list := make([]*types.UnbondRelayFee, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denom := string(key[1:])
		fee := k.GetUnbondRelayFee(ctx, denom)
		list = append(list, &fee)
	}
	return list
}

func (k Keeper) SetUnbondCommission(ctx sdk.Context, denom string, value utils.Dec) {
	store := ctx.KVStore(k.storeKey)
	b, err := value.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set(types.UnbondCommissionStoreKey(denom), b)
}

func (k Keeper) GetUnbondCommission(ctx sdk.Context, denom string) utils.Dec {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.UnbondCommissionStoreKey(denom))
	if b == nil {
		return types.DefaultUnbondCommission
	}
	var val utils.Dec
	if err := val.Unmarshal(b); err != nil {
		panic(err)
	}
	return val
}

func (k Keeper) GetUnbondCommissionList(ctx sdk.Context) []*types.UnbondCommission {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UnbondCommissionPrefix)
	defer iterator.Close()

	list := make([]*types.UnbondCommission, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denom := string(key[1:])
		commission := k.GetUnbondCommission(ctx, denom)
		list = append(list, &types.UnbondCommission{
			Denom: denom,
			Value: commission,
		})
	}
	return list
}

func (k Keeper) SetBondRecord(ctx sdk.Context, br types.BondRecord) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&br)
	store.Set(types.BondRecordStoreKey(br.Denom, br.Txhash), b)
}

func (k Keeper) GetBondRecord(ctx sdk.Context, denom, txHash string) (val types.BondRecord, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.BondRecordStoreKey(denom, txHash))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetBondRecordList(ctx sdk.Context) []*types.BondRecord {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.BondRecordPrefix)
	defer iterator.Close()

	list := make([]*types.BondRecord, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		denomLen := key[1]
		denom := string(key[2 : 2+denomLen])
		txHash := string(key[denomLen+2:])

		bondRecord, found := k.GetBondRecord(ctx, denom, txHash)
		if !found {
			continue
		}
		list = append(list, &bondRecord)
	}
	return list
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

	val := types.Signature{}
	b := store.Get(key)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetSignatureList(ctx sdk.Context) []*types.Signature {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SignaturePrefix)
	defer iterator.Close()

	list := make([]*types.Signature, 0)
	for ; iterator.Valid(); iterator.Next() {
		val := types.Signature{}
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, &val)
	}
	return list
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

func (k Keeper) GetRParamsList(ctx sdk.Context) []*types.RParams {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RParamsPrefix)
	defer iterator.Close()

	list := make([]*types.RParams, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denom := string(key[1:])
		bondRecord, found := k.GetRParams(ctx, denom)
		if !found {
			continue
		}
		list = append(list, &bondRecord)
	}
	return list
}
