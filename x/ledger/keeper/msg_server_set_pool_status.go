package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetPoolStatus(goCtx context.Context, msg *types.MsgSetPoolStatus) (*types.MsgSetPoolStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Status != types.Active && msg.Status != types.NotActive && msg.Status != types.Reserved {
		return nil, types.ErrPoolStatusUnmatch
	}

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	poolDetail, found := k.Keeper.GetPoolDetail(ctx, msg.Denom, msg.Pool)
	if !found {
		return nil, types.ErrPoolDetailNotFound
	}

	poolDetail.Status = msg.Status

	k.Keeper.SetPoolDetail(ctx, &poolDetail)

	return &types.MsgSetPoolStatusResponse{}, nil
}
