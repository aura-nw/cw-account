package keeper

import (
	"smart-account/x/smartaccount/types"
)

var _ types.QueryServer = Keeper{}
