package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/x/rmintreward/keeper"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RmintrewardKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
