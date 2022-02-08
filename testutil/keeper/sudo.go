package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/app"
	"github.com/stafiprotocol/stafihub/x/sudo/keeper"
	"github.com/stafiprotocol/stafihub/x/sudo/types"
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
)

func SudoKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	bankStoreKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	stateStore.MountStoreWithDB(bankStoreKey, sdk.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	encCfg := app.MakeTestEncodingConfig()
	paramsKeeper := ParamsKeeper(&encCfg)
	accountKeeper := AccountKeeper(&encCfg, &paramsKeeper)
	bankKeeper := BankKeeper(&encCfg, &paramsKeeper, &accountKeeper)

	k := SimpleSudoKeeper(storeKey, memStoreKey, codec.NewProtoCodec(registry), bankKeeper)
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return k, ctx
}

func ParamsKeeper(encCfg *params.EncodingConfig) paramskeeper.Keeper {
	keyParams := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	return paramskeeper.NewKeeper(encCfg.Marshaler, encCfg.Amino, keyParams, tkeyParams)
}

func AccountKeeper(encCfg *params.EncodingConfig, paramsKeeper *paramskeeper.Keeper) authkeeper.AccountKeeper {
	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		types.ModuleName:               {authtypes.Burner, authtypes.Minter},
	}

	return authkeeper.NewAccountKeeper(
		encCfg.Marshaler, // amino codec
		keyAcc,           // target store
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount, // prototype,
		maccPerms,
	)
}

func BankKeeper(encCfg *params.EncodingConfig, paramsKeeper *paramskeeper.Keeper, accountKeeper *authkeeper.AccountKeeper) bankkeeper.Keeper {
	keyBank := sdk.NewKVStoreKey(banktypes.StoreKey)
	blacklistedAddrs := make(map[string]bool)
	return bankkeeper.NewBaseKeeper(
		encCfg.Marshaler,
		keyBank,
		accountKeeper,
		paramsKeeper.Subspace(banktypes.ModuleName),
		blacklistedAddrs,
	)
}

func SimpleSudoKeeper(storeKey *sdk.KVStoreKey, memStoreKey *sdk.MemoryStoreKey, cdc *codec.ProtoCodec, bankKeeper bankkeeper.Keeper) *keeper.Keeper {
	return keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		bankKeeper,
	)
}