package keeper

import (
	"sync"
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/relayers/keeper"
	"github.com/stafihub/stafihub/x/relayers/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	relayersStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	relayersMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	relayersOnce        sync.Once
)

func RelayersKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	relayersOnce.Do(func() {
		stateStore.MountStoreWithDB(relayersStoreKey, storetypes.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(relayersMemStoreKey, storetypes.StoreTypeMemory, nil)
	})
	require.NoError(t, stateStore.LoadLatestVersion())

	sudoKeeper, _ := SudoKeeper(t)
	relayersKeeper := keeper.NewKeeper(
		cdc,
		relayersStoreKey,
		relayersMemStoreKey,
		sudoKeeper,
		BankKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return relayersKeeper, ctx
}
