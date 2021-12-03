package keeper

import (
	"github.com/stafiprotocol/stafihub/x/sudo/types"
)

var _ types.QueryServer = Keeper{}
