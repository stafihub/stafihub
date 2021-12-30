package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/rate/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
	}
)

func NewKeeper(
    cdc codec.BinaryCodec,
    storeKey,
    memKey sdk.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// token to rtoken
func (k Keeper) TokenToRtoken(ctx sdk.Context, denom string, balance sdk.Int) sdk.Int {
	er, ok := k.GetExchangeRate(ctx, denom)
	if !ok {
		return balance
	}

	return er.Value.MulInt(balance).TruncateInt()
}

// rtoken to token
func (k Keeper) RtokenToToken(ctx sdk.Context, denom string, rbalance sdk.Int) sdk.Int {
	er, ok := k.GetExchangeRate(ctx, denom)
	if !ok {
		return rbalance
	}

	return sdk.OneDec().MulInt(rbalance).Quo(er.Value).TruncateInt()
}

