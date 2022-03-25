package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rmintreward/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddMintRewardAct(goCtx context.Context, msg *types.MsgAddMintRewardAct) (*types.MsgAddMintRewardActResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	willUseCycle := uint64(0)
	latestCycle, found := k.Keeper.GetActLatestCycle(ctx, msg.Denom)
	if found {
		willUseCycle = latestCycle + 1
	}

	k.Keeper.SetMintRewardAct(ctx, msg.Denom, willUseCycle, msg.Act)
	k.Keeper.SetActLatestCycle(ctx, msg.Denom, willUseCycle)

	return &types.MsgAddMintRewardActResponse{}, nil
}
