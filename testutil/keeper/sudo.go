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
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stafihub/stafihub/app"
	"github.com/stafihub/stafihub/testutil/sample"
	miningtypes "github.com/stafihub/stafihub/x/mining/types"
	rdextypes "github.com/stafihub/stafihub/x/rdex/types"
	"github.com/stafihub/stafihub/x/sudo/keeper"
	"github.com/stafihub/stafihub/x/sudo/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

var (
	registry = codectypes.NewInterfaceRegistry()
	cdc      = codec.NewProtoCodec(registry)

	db         = tmdb.NewMemDB()
	stateStore = store.NewCommitMultiStore(db)
	encCfg     = app.MakeTestEncodingConfig()

	ParamsKeeper  = NewParamsKeeper(&encCfg)
	AccountKeeper = NewAccountKeeper(&encCfg, ParamsKeeper)
	BankKeeper    = NewBankKeeper(&encCfg, ParamsKeeper, AccountKeeper)

	sudoStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	sudoMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	sudoOnce        sync.Once
	willMintCoins   = sdk.NewCoins(sdk.NewCoin(sample.TestDenom, sdk.NewInt(500e8)), sdk.NewCoin(sample.TestDenom1, sdk.NewInt(500e8)))
)

func SudoKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	sudoOnce.Do(func() {
		stateStore.MountStoreWithDB(sudoStoreKey, storetypes.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(sudoMemStoreKey, storetypes.StoreTypeMemory, nil)
	})

	require.NoError(t, stateStore.LoadLatestVersion())

	sudoKeeper := keeper.NewKeeper(
		cdc,
		sudoStoreKey,
		sudoMemStoreKey,
	)
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	require.NotNil(t, sample.TestAdminAcc)

	sudoKeeper.SetAdmin(ctx, sample.TestAdminAcc)

	BankKeeper.MintCoins(ctx, types.ModuleName, willMintCoins)
	BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sample.TestAdminAcc, willMintCoins)
	return sudoKeeper, ctx
}

func NewParamsKeeper(encCfg *params.EncodingConfig) *paramskeeper.Keeper {
	keyParams := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	stateStore.MountStoreWithDB(keyParams, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tkeyParams, storetypes.StoreTypeIAVL, db)

	k := paramskeeper.NewKeeper(encCfg.Codec, encCfg.Amino, keyParams, tkeyParams)
	return &k
}

func NewAccountKeeper(encCfg *params.EncodingConfig, paramsKeeper *paramskeeper.Keeper) *authkeeper.AccountKeeper {
	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	stateStore.MountStoreWithDB(keyAcc, storetypes.StoreTypeIAVL, db)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		types.ModuleName:               {authtypes.Burner, authtypes.Minter},
		rdextypes.ModuleName:           {authtypes.Burner, authtypes.Minter},
		miningtypes.ModuleName:         {authtypes.Burner, authtypes.Minter},
	}

	k := authkeeper.NewAccountKeeper(
		encCfg.Codec, // amino codec
		keyAcc,       // target store
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount, // prototype,
		maccPerms,
		sdk.Bech32MainPrefix,
	)
	return &k
}

func NewBankKeeper(encCfg *params.EncodingConfig, paramsKeeper *paramskeeper.Keeper, accountKeeper *authkeeper.AccountKeeper) bankkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	blacklistedAddrs := make(map[string]bool)
	return bankkeeper.NewBaseKeeper(
		encCfg.Codec,
		storeKey,
		accountKeeper,
		paramsKeeper.Subspace(banktypes.ModuleName),
		blacklistedAddrs,
	)
}
