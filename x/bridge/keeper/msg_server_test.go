package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/x/bridge/keeper"
	"github.com/stafihub/stafihub/x/bridge/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BridgeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
