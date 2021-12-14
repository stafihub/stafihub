package keeper

import (
	"github.com/stafiprotocol/stafihub/x/ledger/types"
)

var _ types.QueryServer = Keeper{}
