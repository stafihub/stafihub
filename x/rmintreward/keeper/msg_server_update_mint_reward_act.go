package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rmintreward/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) UpdateMintRewardAct(goCtx context.Context, msg *types.MsgUpdateMintRewardAct) (*types.MsgUpdateMintRewardActResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	k.Keeper.SetMintRewardAct(ctx, msg.Denom, msg.Cycle, msg.Act)
	return &types.MsgUpdateMintRewardActResponse{}, nil
}
