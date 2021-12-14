package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/stafiprotocol/stafihub/x/ledger/types"
    "github.com/stafiprotocol/stafihub/x/ledger/keeper"
    keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LedgerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
