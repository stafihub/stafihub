package keeper

import (
	"sync"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/app"
	"github.com/stafihub/stafihub/testutil/sample"
	"github.com/stafihub/stafihub/x/sudo/keeper"
	"github.com/stafihub/stafihub/x/sudo/types"
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

var (
	registry = codectypes.NewInterfaceRegistry()
	cdc      = codec.NewProtoCodec(registry)

	db         = tmdb.NewMemDB()
	stateStore = store.NewCommitMultiStore(db)
	encCfg     = app.MakeTestEncodingConfig()

	paramsKeeper  = ParamsKeeper(&encCfg)
	accountKeeper = AccountKeeper(&encCfg, paramsKeeper)
	bankKeeper    = BankKeeper(&encCfg, paramsKeeper, accountKeeper)

	sudoStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	sudoMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	sudoOnce        sync.Once
)

func SudoKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	sudoOnce.Do(func() {
		stateStore.MountStoreWithDB(sudoStoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(sudoMemStoreKey, sdk.StoreTypeMemory, nil)
	})

	require.NoError(t, stateStore.LoadLatestVersion())

	sudoKeeper := keeper.NewKeeper(
		cdc,
		sudoStoreKey,
		sudoMemStoreKey,
	)
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	require.NotNil(t, sample.TestAdminAcc)
	//t.Log("TestAdmin", sample.TestAdmin)
	sudoKeeper.SetAdmin(ctx, sample.TestAdminAcc)
	return sudoKeeper, ctx
}

func ParamsKeeper(encCfg *params.EncodingConfig) *paramskeeper.Keeper {
	keyParams := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	stateStore.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tkeyParams, sdk.StoreTypeIAVL, db)

	k := paramskeeper.NewKeeper(encCfg.Marshaler, encCfg.Amino, keyParams, tkeyParams)
	return &k
}

func AccountKeeper(encCfg *params.EncodingConfig, paramsKeeper *paramskeeper.Keeper) *authkeeper.AccountKeeper {
	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	stateStore.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		types.ModuleName:               {authtypes.Burner, authtypes.Minter},
	}

	k := authkeeper.NewAccountKeeper(
		encCfg.Marshaler, // amino codec
		keyAcc,           // target store
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount, // prototype,
		maccPerms,
	)
	return &k
}

func BankKeeper(encCfg *params.EncodingConfig, paramsKeeper *paramskeeper.Keeper, accountKeeper *authkeeper.AccountKeeper) bankkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)

	blacklistedAddrs := make(map[string]bool)
	return bankkeeper.NewBaseKeeper(
		encCfg.Marshaler,
		storeKey,
		accountKeeper,
		paramsKeeper.Subspace(banktypes.ModuleName),
		blacklistedAddrs,
	)
}
