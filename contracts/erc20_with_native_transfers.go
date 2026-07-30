package contracts

import (
	contractutils "github.com/arieschainorg/Aries-Chain/contracts/utils"
	evmtypes "github.com/arieschainorg/Aries-Chain/x/vm/types"
)

func LoadERC20WithNativeTransfers() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("solidity/ERC20WithNativeTransfers.json")
}
