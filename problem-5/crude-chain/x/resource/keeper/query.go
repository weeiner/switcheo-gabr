package keeper

import (
	"crude-chain/x/resource/types"
)

var _ types.QueryServer = Keeper{}
