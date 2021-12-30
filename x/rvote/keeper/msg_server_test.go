package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/stafiprotocol/stafihub/x/rvote/types"
    "github.com/stafiprotocol/stafihub/x/rvote/keeper"
    keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RvoteKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
