package iwhata_test

import (
	"testing"

	keepertest "github.com/iwhata/testutil/keeper"
	"github.com/iwhata/testutil/nullify"
	"github.com/iwhata/x/iwhata"
	"github.com/iwhata/x/iwhata/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IwhataKeeper(t)
	iwhata.InitGenesis(ctx, *k, genesisState)
	got := iwhata.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
