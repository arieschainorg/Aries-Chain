package testutil

import (
	anteinterfaces "github.com/arieschainorg/Aries-Chain/ante/interfaces"
	"github.com/arieschainorg/Aries-Chain/x/vm/statedb"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewStateDB returns a new StateDB for testing purposes.
func NewStateDB(ctx sdk.Context, evmKeeper anteinterfaces.EVMKeeper) *statedb.StateDB {
	return statedb.New(ctx, evmKeeper, statedb.NewEmptyTxConfig())
}
