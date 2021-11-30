package keeper

import (
	"github.com/stafiprotocol/stafihub/x/rate/types"
)

var _ types.QueryServer = Keeper{}
