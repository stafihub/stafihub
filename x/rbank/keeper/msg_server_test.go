package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/x/rbank/keeper"
	"github.com/stafihub/stafihub/x/rbank/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RbankKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
