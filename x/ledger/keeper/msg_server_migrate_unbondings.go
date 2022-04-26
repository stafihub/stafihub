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
	err := k.Keeper.CheckAddress(ctx, msg.Denom, msg.Pool)
	if err != nil {
		return nil, err
	}
	if !k.Keeper.IsBondedPoolExist(ctx, msg.Denom, msg.Pool) {
		return nil, types.ErrPoolNotBonded
	}

	for seq, unbonding := range msg.Unbondings {
		err := k.Keeper.CheckAddress(ctx, msg.Denom, unbonding.Recipient)
		if err != nil {
			return nil, err
		}
		unbonding.Unbonder = sdk.AccAddress(zeroAddress[:]).String()

		// coverable here
		k.SetPoolUnbonding(ctx, msg.Denom, msg.Pool, msg.Era, uint32(seq), unbonding)
	}

	if len(msg.Unbondings) > 0 {
		k.SetPoolUnbondSequence(ctx, msg.Denom, msg.Pool, msg.Era, uint32(len(msg.Unbondings)-1))
	}

	return &types.MsgMigrateUnbondingsResponse{}, nil
}
