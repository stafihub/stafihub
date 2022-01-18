package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

        sudoKeeper types.SudoKeeper
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
    cdc codec.BinaryCodec,
    storeKey,
    memKey sdk.StoreKey,

    sudoKeeper types.SudoKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		sudoKeeper: sudoKeeper,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetLastVoter(ctx sdk.Context, denom, voter string) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.LastVoterPrefix)
	lv := types.LastVoter{
		Denom: denom,
		Voter: voter,
	}
	b := k.cdc.MustMarshal(&lv)
	store.Set([]byte(denom), b)
}

func (k Keeper) LastVoter(ctx sdk.Context, denom string) (val types.LastVoter, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.LastVoterPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
