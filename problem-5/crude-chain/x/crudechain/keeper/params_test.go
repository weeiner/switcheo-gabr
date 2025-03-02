package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "crude-chain/testutil/keeper"
	"crude-chain/x/crudechain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CrudechainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
