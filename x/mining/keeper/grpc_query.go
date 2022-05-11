package keeper

import (
	"github.com/stafihub/stafihub/x/mining/types"
)

var _ types.QueryServer = Keeper{}
