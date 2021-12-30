package keeper

import (
	"github.com/stafiprotocol/stafihub/x/rvote/types"
)

var _ types.QueryServer = Keeper{}
