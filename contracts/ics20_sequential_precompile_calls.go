package contracts

import (
	contractutils "github.com/arieschainorg/Aries-Chain/contracts/utils"
	evmtypes "github.com/arieschainorg/Aries-Chain/x/vm/types"
)

func LoadSequentialICS20Sender() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("solidity/SequentialICS20Sender.json")
}
