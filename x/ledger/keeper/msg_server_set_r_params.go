package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
)

func (k msgServer) SetRParams(goCtx context.Context, msg *types.MsgSetRParams) (*types.MsgSetRParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetRParamsResponse{}, nil
}
