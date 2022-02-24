package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
)

func (k Keeper) SetRValidator(ctx sdk.Context, validator types.RValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorPrefix)
	b := k.cdc.MustMarshal(&validator)
	store.Set([]byte(validator.Denom+validator.Address+validator.OperatorAddress), b)
}

func (k Keeper) RValidator(ctx sdk.Context, denom, address, operatorAddress string) (val types.RValidator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorPrefix)
	b := store.Get([]byte(denom + address + operatorAddress))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetRValidatorIndicator(ctx sdk.Context, indicator types.RValidatorIndicator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorIndicatorPrefix)
	b := k.cdc.MustMarshal(&indicator)
	store.Set([]byte(indicator.Denom), b)
}

func (k Keeper) RValidatorIndicator(ctx sdk.Context, denom string) (val types.RValidatorIndicator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorIndicatorPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
