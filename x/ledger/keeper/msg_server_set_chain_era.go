package keeper

import (
	"context"

    "github.com/stafiprotocol/stafihub/x/ledger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) SetChainEra(goCtx context.Context,  msg *types.MsgSetChainEra) (*types.MsgSetChainEraResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgSetChainEraResponse{}, nil
}
