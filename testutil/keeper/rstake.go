package keeper

import (
	"sync"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/rstaking/keeper"
	"github.com/stafihub/stafihub/x/rstaking/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	rstakeStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	rstakeMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	rstakeOnce        sync.Once
)

func RStakingKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	rstakeOnce.Do(func() {
		stateStore.MountStoreWithDB(rstakeStoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(rstakeMemStoreKey, sdk.StoreTypeMemory, nil)
	})
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		rstakeStoreKey,
		rstakeMemStoreKey,
		"RStakingParams",
	)

	sudoKeeper, _ := SudoKeeper(t)
	k := keeper.NewKeeper(
		cdc,
		rstakeStoreKey,
		rstakeMemStoreKey,
		paramsSubspace,

		BankKeeper,
		sudoKeeper,
		authtypes.FeeCollectorName,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
