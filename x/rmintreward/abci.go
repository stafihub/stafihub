package rmintreward

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rmintreward/keeper"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.UpdateActLatestCycle(ctx)
}
