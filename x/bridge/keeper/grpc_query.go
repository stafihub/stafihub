package keeper

import (
	"github.com/stafihub/stafihub/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
