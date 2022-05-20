package keeper

import (
	"sync"
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/keeper"
	"github.com/stafihub/stafihub/x/ledger/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	ledgertoreKey     = sdk.NewKVStoreKey(types.StoreKey)
	ledgerMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	ledgerOnce        sync.Once
)

func LedgerKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	ledgerOnce.Do(func() {
		stateStore.MountStoreWithDB(ledgertoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(ledgerMemStoreKey, sdk.StoreTypeMemory, nil)
	})
	require.NoError(t, stateStore.LoadLatestVersion())

	sudoKeeper, _ := SudoKeeper(t)
	relayersKeeper, _ := RelayersKeeper(t)
	rmintRewardKeeper, _ := RmintrewardKeeper(t)
	rBankKeeper, _ := RbankKeeper(t)
	ledgerKeeper := keeper.NewKeeper(
		cdc,
		ledgertoreKey,
		ledgerMemStoreKey,
		sudoKeeper,
		BankKeeper,
		relayersKeeper,
		rmintRewardKeeper,
		rBankKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return ledgerKeeper, ctx
}
