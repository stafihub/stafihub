package keeper

import (
	"sync"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/rdex/keeper"
	"github.com/stafihub/stafihub/x/rdex/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	rdexStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	rdexMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	rdexOnce        sync.Once
)

func RdexKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	rdexOnce.Do(func() {
		stateStore.MountStoreWithDB(rdexStoreKey, storetypes.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(rdexMemStoreKey, storetypes.StoreTypeMemory, nil)
	})
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		rdexStoreKey,
		rdexMemStoreKey,
		"RdexParams",
	)
	sudoKeeper, _ := SudoKeeper(t)
	k := keeper.NewKeeper(
		cdc,
		rdexStoreKey,
		rdexMemStoreKey,
		paramsSubspace,
		BankKeeper,
		sudoKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
