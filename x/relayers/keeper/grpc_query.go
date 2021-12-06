package keeper

import (
	"github.com/stafiprotocol/stafihub/x/relayers/types"
)

var _ types.QueryServer = Keeper{}
