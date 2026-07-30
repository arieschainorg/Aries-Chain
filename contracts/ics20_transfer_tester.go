package contracts

import (
	contractutils "github.com/arieschainorg/Aries-Chain/contracts/utils"
	evmtypes "github.com/arieschainorg/Aries-Chain/x/vm/types"
)

func LoadICS20TransferTester() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("solidity/ICS20TransferTester.json")
}
