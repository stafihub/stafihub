package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

var zeroAddress [20]byte

func (k msgServer) MigrateUnbondings(goCtx context.Context, msg *types.MsgMigrateUnbondings) (*types.MsgMigrateUnbondingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	for _, poolUnbonds := range msg.PoolUnbonds {
		if msg.Denom != poolUnbonds.Denom {
			return nil, types.ErrUnbondingDenomNotMatch
		}
		err := k.Keeper.CheckAddress(ctx, msg.Denom, poolUnbonds.Pool)
		if err != nil {
			return nil, err
		}
		for index, unbonding := range poolUnbonds.Unbondings {
			err := k.Keeper.CheckAddress(ctx, msg.Denom, unbonding.Recipient)
			if err != nil {
				return nil, err
			}
			poolUnbonds.Unbondings[index].Unbonder = sdk.AccAddress(zeroAddress[:]).String()
		}
		if !k.Keeper.IsBondedPoolExist(ctx, msg.Denom, poolUnbonds.Pool) {
			return nil, types.ErrPoolNotBonded
		}

		// coverable here
		k.SetPoolUnbond(ctx, *poolUnbonds)
	}

	return &types.MsgMigrateUnbondingsResponse{}, nil
}
