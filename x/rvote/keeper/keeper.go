package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stafihub/stafihub/x/rvote/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		sudoKeeper    types.SudoKeeper
		relayerKeeper types.RelayerKeeper
		router        types.Router
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,

	sudoKeeper types.SudoKeeper,
	relayerKeeper types.RelayerKeeper,
	rtr types.Router,
) *Keeper {
	rtr.Seal()

	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		sudoKeeper:    sudoKeeper,
		relayerKeeper: relayerKeeper,
		router:        rtr,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetProposalLife(ctx sdk.Context, proposalLife int64) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: proposalLife})
	store.Set(types.ProposalLifePrefix, b)
}

func (k Keeper) ProposalLife(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.ProposalLifePrefix)
	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshal(b, &intV)

	return intV.GetValue()
}

// Router returns rvoteKeeper's Router
func (keeper Keeper) Router() types.Router {
	return keeper.router
}
