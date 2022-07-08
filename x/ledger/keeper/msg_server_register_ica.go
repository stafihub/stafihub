package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
)

func (k msgServer) RegisterIca(goCtx context.Context, msg *types.MsgRegisterIca) (*types.MsgRegisterIcaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, msg.Owner); err != nil {
		return nil, err
	}

	return &types.MsgRegisterIcaResponse{}, nil
}
