package werc20

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/precompiles/werc20"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestWERC20PrecompileUnitTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.WERC20PrecompileApp](integration.CreateEvmd, "evm.WERC20PrecompileApp")
	s := werc20.NewPrecompileUnitTestSuite(create)
	suite.Run(t, s)
}

func TestWERC20PrecompileIntegrationTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.WERC20PrecompileApp](integration.CreateEvmd, "evm.WERC20PrecompileApp")
	werc20.TestPrecompileIntegrationTestSuite(t, create)
}
