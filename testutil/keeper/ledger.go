package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/app"
	"github.com/stafiprotocol/stafihub/x/ledger/keeper"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	relayerkeeper "github.com/stafiprotocol/stafihub/x/relayers/keeper"
	sudokeeper "github.com/stafiprotocol/stafihub/x/sudo/keeper"
	sudotypes "github.com/stafiprotocol/stafihub/x/sudo/types"
)

func LedgerKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	encCfg := app.MakeTestEncodingConfig()
	keyParams := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	paramsKeeper := paramskeeper.NewKeeper(encCfg.Marshaler, encCfg.Amino, keyParams, tkeyParams)

	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		types.ModuleName:               {authtypes.Burner, authtypes.Minter},
	}
	accountKeeper := authkeeper.NewAccountKeeper(
		encCfg.Marshaler, // amino codec
		keyAcc,           // target store
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount, // prototype,
		maccPerms,
	)

	keyBank := sdk.NewKVStoreKey(banktypes.StoreKey)
	blacklistedAddrs := make(map[string]bool)
	bankKeeper := bankkeeper.NewBaseKeeper(
		encCfg.Marshaler,
		keyBank,
		accountKeeper,
		paramsKeeper.Subspace(banktypes.ModuleName),
		blacklistedAddrs,
	)

	sudoKey := sdk.NewKVStoreKey(sudotypes.StoreKey)
	sudoMemKey := storetypes.NewMemoryStoreKey(sudotypes.MemStoreKey)
	sudoKeeper := sudokeeper.NewKeeper(
		cdc,
		sudoKey,
		sudoMemKey,
		bankKeeper,
	)

	relayerKeeper := relayerkeeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		sudoKeeper,
		bankKeeper,
	)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		sudoKeeper,
		bankKeeper,
		relayerKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return k, ctx
}
