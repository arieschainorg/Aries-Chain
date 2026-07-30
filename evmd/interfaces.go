package evmd

import (
	cmn "github.com/arieschainorg/Aries-Chain/precompiles/common"
	evmtypes "github.com/arieschainorg/Aries-Chain/x/vm/types"
)

type BankKeeper interface {
	evmtypes.BankKeeper
	cmn.BankKeeper
}
