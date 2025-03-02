package keeper

import (
	"crude-chain/x/crudechain/types"
)

var _ types.QueryServer = Keeper{}
