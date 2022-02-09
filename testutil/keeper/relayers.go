package keeper

import (
	"sync"
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/keeper"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
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
		stateStore.MountStoreWithDB(relayersStoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(relayersMemStoreKey, sdk.StoreTypeMemory, nil)
	})
	require.NoError(t, stateStore.LoadLatestVersion())

	sudoKeeper, _ := SudoKeeper(t)
	relayersKeeper := keeper.NewKeeper(
		cdc,
		relayersStoreKey,
		relayersMemStoreKey,
		sudoKeeper,
		bankKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return relayersKeeper, ctx
}
