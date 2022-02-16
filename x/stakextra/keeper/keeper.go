package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/stakextra/types"
)

type (
	Keeper struct {
		cdc              codec.BinaryCodec
		storeKey         sdk.StoreKey
		memKey           sdk.StoreKey
		paramstore       paramtypes.Subspace
		feeCollectorName string
		bankKeeper       types.BankKeeper
		sudoKeeper       types.SudoKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	sudoKeeper types.SudoKeeper,
	feeCollectorName string,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:              cdc,
		storeKey:         storeKey,
		memKey:           memKey,
		paramstore:       ps,
		bankKeeper:       bankKeeper,
		sudoKeeper:       sudoKeeper,
		feeCollectorName: feeCollectorName,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetInflationBase(ctx sdk.Context, inflationBase sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bts, err := inflationBase.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value %v", err))
	}
	store.Set(types.InflationBasePrefix, bts)
}

func (k Keeper) GetInflationBase(ctx sdk.Context) sdk.Int {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.InflationBasePrefix)
	if len(bts) == 0 {
		panic(fmt.Errorf("inflationBase not found"))
	}
	var amount sdk.Int
	err := amount.Unmarshal(bts)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal supply value %v", err))
	}

	return amount
}

// impl for mint keeper
func (k Keeper) StakingTokenSupply(ctx sdk.Context) sdk.Int {
	return k.GetInflationBase(ctx)
}

// impl for mint keeper
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	return sdk.ZeroDec()
}

func (k Keeper) GetFeeCollectorName() string {
	return k.feeCollectorName
}

func (k Keeper) GetBankKeeper() types.BankKeeper {
	return k.bankKeeper
}

func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}
