package keeper

import (
	"github.com/stafihub/stafihub/x/rdex/types"
)

var _ types.QueryServer = Keeper{}
