package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) UnsealMigrateInit(goCtx context.Context, msg *types.MsgUnsealMigrateInit) (*types.MsgUnsealMigrateInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	k.Keeper.UnSealMigrateInit(ctx)

	return &types.MsgUnsealMigrateInitResponse{}, nil
}
