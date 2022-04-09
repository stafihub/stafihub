package keeper

import (
	"github.com/stafihub/stafihub/x/rbank/types"
)

var _ types.QueryServer = Keeper{}
