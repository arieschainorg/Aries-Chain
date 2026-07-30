package testutil

import (
	contractutils "github.com/arieschainorg/Aries-Chain/contracts/utils"
	evmtypes "github.com/arieschainorg/Aries-Chain/x/vm/types"
)

func LoadCounterWithCallbacksContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("CounterWithCallbacks.json")
}
