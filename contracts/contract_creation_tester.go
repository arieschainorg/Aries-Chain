package contracts

import (
	contractutils "github.com/arieschainorg/Aries-Chain/contracts/utils"
	evmtypes "github.com/arieschainorg/Aries-Chain/x/vm/types"
)

func LoadContractCreationTester() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("solidity/ContractCreationTester.json")
}
