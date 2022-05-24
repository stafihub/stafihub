package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/mining/keeper"
	"github.com/stafihub/stafihub/x/mining/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	miningStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	miningMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
)

func MiningKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	relayersOnce.Do(func() {
		stateStore.MountStoreWithDB(miningStoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(miningMemStoreKey, sdk.StoreTypeMemory, nil)
	})
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		miningStoreKey,
		miningMemStoreKey,
		"MiningParams",
	)

	sudoKeeper, _ := SudoKeeper(t)
	k := keeper.NewKeeper(
		cdc,
		miningStoreKey,
		miningMemStoreKey,
		paramsSubspace,
		sudoKeeper,
		BankKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
