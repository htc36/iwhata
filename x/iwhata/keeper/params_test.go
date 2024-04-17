package keeper_test

import (
	"testing"

	testkeeper "github.com/iwhata/testutil/keeper"
	"github.com/iwhata/x/iwhata/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IwhataKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
