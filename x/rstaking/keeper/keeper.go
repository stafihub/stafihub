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
		key := iterator.Key()
		if len(key) <= 1 {
			continue
		}

		valList = append(valList, sdk.ValAddress(key[1:]).String())
	}
	return valList
}

func (k Keeper) ToggleValidatorWhitelistSwitch(ctx sdk.Context) {
	k.SetValidatorWhitelistSwitch(ctx, !k.GetValidatorWhitelistSwitch(ctx))
}

func (k Keeper) SetValidatorWhitelistSwitch(ctx sdk.Context, isOpen bool) {
	store := ctx.KVStore(k.storeKey)
	state := types.SwitchStateClose
	if isOpen {
		state = types.SwitchStateOpen
	}
	store.Set(types.ValidatorWhitelistSwitchKey, state)
}

func (k Keeper) GetValidatorWhitelistSwitch(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ValidatorWhitelistSwitchKey)
	if bts == nil {
		return true
	}
	return bytes.Equal(bts, types.SwitchStateOpen)
}

func (k Keeper) AddDelegatorAddressToWhitelist(ctx sdk.Context, delegatorAddress sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.DelegatorAddressStoreKey(delegatorAddress), []byte{})
}

func (k Keeper) HasDelegatorAddressInWhitelist(ctx sdk.Context, delegatorAddress sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.DelegatorAddressStoreKey(delegatorAddress))
}

func (k Keeper) GetDelegatorAddressWhitelist(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DelegatorAddressStoreKeyPrefix)
	defer iterator.Close()

	valList := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		if len(key) <= 1 {
			continue
		}

		valList = append(valList, sdk.ValAddress(key[1:]).String())
	}
	return valList
}

func (k Keeper) ToggleDelegatorWhitelistSwitch(ctx sdk.Context) {
	k.SetDelegatorWhitelistSwitch(ctx, !k.GetDelegatorWhitelistSwitch(ctx))
}

func (k Keeper) SetDelegatorWhitelistSwitch(ctx sdk.Context, isOpen bool) {
	store := ctx.KVStore(k.storeKey)
	state := types.SwitchStateClose
	if isOpen {
		state = types.SwitchStateOpen
	}
	store.Set(types.DelegatorWhitelistSwitchKey, state)
}

func (k Keeper) GetDelegatorWhitelistSwitch(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.DelegatorWhitelistSwitchKey)
	if bts == nil {
		return true
	}
	return bytes.Equal(bts, types.SwitchStateOpen)
}
