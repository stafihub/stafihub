package keeper

import (
	"bytes"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/rstaking/types"
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
	store.Set(types.InflationBaseKey, bts)
}

func (k Keeper) GetInflationBase(ctx sdk.Context) (sdk.Int, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.InflationBaseKey)
	if bts == nil {
		return sdk.ZeroInt(), false
	}
	var amount sdk.Int
	err := amount.Unmarshal(bts)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal supply value %v", err))
	}
	return amount, true
}

// impl for mint keeper
func (k Keeper) StakingTokenSupply(ctx sdk.Context) sdk.Int {
	inflationBase, found := k.GetInflationBase(ctx)
	if !found {
		return sdk.ZeroInt()
	}
	return inflationBase
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

func (k Keeper) AddValAddressToWhitelist(ctx sdk.Context, valAddress sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ValAddressStoreKey(valAddress), []byte{})
}

func (k Keeper) HasValAddressInWhitelist(ctx sdk.Context, valAddress sdk.ValAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ValAddressStoreKey(valAddress))
}

func (k Keeper) GetValAddressWhitelist(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValAddressStoreKeyPrefix)
	defer iterator.Close()

	valList := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		valList = append(valList, sdk.ValAddress(iterator.Key()).String())
	}
	return valList
}

func (k Keeper) ToggleWhitelistSwitch(ctx sdk.Context) {
	k.SetWhitelistSwitch(ctx, !k.GetWhitelistSwitch(ctx))
}

func (k Keeper) SetWhitelistSwitch(ctx sdk.Context, isOpen bool) {
	store := ctx.KVStore(k.storeKey)
	state := types.SwitchStateClose
	if isOpen {
		state = types.SwitchStateOpen
	}
	store.Set(types.WhitelistSwitchKey, state)
}

func (k Keeper) GetWhitelistSwitch(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.WhitelistSwitchKey)
	if bts == nil {
		return true
	}
	return bytes.Equal(bts, types.SwitchStateOpen)
}
