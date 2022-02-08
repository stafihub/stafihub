package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/app"
	"github.com/stafiprotocol/stafihub/x/rvote/keeper"
	"github.com/stafiprotocol/stafihub/x/rvote/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ledgermodule "github.com/stafiprotocol/stafihub/x/ledger"
	ledgerkeeper "github.com/stafiprotocol/stafihub/x/ledger/keeper"
	ledgertypes "github.com/stafiprotocol/stafihub/x/ledger/types"
	relayerskeeper "github.com/stafiprotocol/stafihub/x/relayers/keeper"
	relayerstypes "github.com/stafiprotocol/stafihub/x/relayers/types"
	sudokeeper "github.com/stafiprotocol/stafihub/x/sudo/keeper"
	sudotypes "github.com/stafiprotocol/stafihub/x/sudo/types"
)

func RvoteKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	sudoStoreKey := sdk.NewKVStoreKey(sudotypes.StoreKey)
	sudoMemStoreKey := storetypes.NewMemoryStoreKey(sudotypes.MemStoreKey)

	relayersStoreKey := sdk.NewKVStoreKey(relayerstypes.StoreKey)
	relayersMemStoreKey := storetypes.NewMemoryStoreKey(relayerstypes.MemStoreKey)

	ledgerStoreKey := sdk.NewKVStoreKey(ledgertypes.StoreKey)
	ledgerMemStoreKey := storetypes.NewMemoryStoreKey(ledgertypes.MemStoreKey)

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	bankStoreKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	stateStore.MountStoreWithDB(bankStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(sudoStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(sudoMemStoreKey, sdk.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(relayersStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(relayersMemStoreKey, sdk.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(ledgerStoreKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(ledgerMemStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	encCfg := app.MakeTestEncodingConfig()
	paramsKeeper := ParamsKeeper(&encCfg)
	accountKeeper := AccountKeeper(&encCfg, &paramsKeeper)
	bankKeeper := BankKeeper(&encCfg, &paramsKeeper, &accountKeeper)
	sudoKeeper := SimpleSudoKeeper(sudoStoreKey, sudoMemStoreKey, cdc, bankKeeper)
	relayersKeeper := SimpleRelayersKeeper(relayersStoreKey, relayersMemStoreKey, cdc, bankKeeper, sudoKeeper)
	ledgerKeeper := SimpleLedgerKeeper(ledgerStoreKey, ledgerMemStoreKey, cdc, sudoKeeper, bankKeeper, relayersKeeper)
	k := SimpleRvoteKeeper(storeKey, memStoreKey, cdc, sudoKeeper, relayersKeeper, ledgerKeeper)
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return k, ctx
}

func SimpleRvoteKeeper(storeKey *sdk.KVStoreKey, memStoreKey *sdk.MemoryStoreKey, cdc *codec.ProtoCodec, sudoKeeper *sudokeeper.Keeper, relayersKeeper *relayerskeeper.Keeper, ledgerKeeper *ledgerkeeper.Keeper) *keeper.Keeper {
	rvoteRouter := types.NewRouter()
	rvoteRouter.AddRoute(ledgertypes.RouterKey, ledgermodule.NewProposalHandler(*ledgerKeeper))

	return keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,

		sudoKeeper,
		relayersKeeper,
		rvoteRouter,
	)
}


