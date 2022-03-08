package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) MigrateUnbondings(goCtx context.Context, msg *types.MsgMigrateUnbondings) (*types.MsgMigrateUnbondingsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	//should return if exist&&exchangeRate != 1
	rate, found := k.GetExchangeRate(ctx, msg.Denom)
	if found && !rate.Value.Equal(utils.OneDec()) {
		return nil, types.ErrExchangeRateAlreadyExist
	}

	for _, poolUnbonds := range msg.PoolUnbonds {
		err := k.Keeper.CheckAddress(ctx, msg.Denom, poolUnbonds.Pool)
		if err != nil {
			return nil, err
		}
		for _, p := range poolUnbonds.Unbondings {
			err := k.Keeper.CheckAddress(ctx, msg.Denom, p.Recipient)
			if err != nil {
				return nil, err
			}
		}

		// coverable here
		k.SetPoolUnbond(ctx, *poolUnbonds)
	}

	return &types.MsgMigrateUnbondingsResponse{}, nil
}
