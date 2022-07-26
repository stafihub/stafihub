package keeper

import (
	"github.com/stafihub/stafihub/x/claim/types"
)

var _ types.QueryServer = Keeper{}
