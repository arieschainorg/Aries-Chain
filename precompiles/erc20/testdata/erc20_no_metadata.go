package testdata

import (
	contractutils "github.com/arieschainorg/Aries-Chain/contracts/utils"
	evmtypes "github.com/arieschainorg/Aries-Chain/x/vm/types"
)

func LoadERC20NoMetadataContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20NoMetadata.json")
}
