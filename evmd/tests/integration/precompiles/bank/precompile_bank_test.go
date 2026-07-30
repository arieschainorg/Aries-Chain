package bank

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/precompiles/bank"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestBankPrecompileTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.BankPrecompileApp](integration.CreateEvmd, "evm.BankPrecompileApp")
	s := bank.NewPrecompileTestSuite(create)
	suite.Run(t, s)
}

func TestBankPrecompileIntegrationTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.BankPrecompileApp](integration.CreateEvmd, "evm.BankPrecompileApp")
	bank.TestIntegrationSuite(t, create)
}
