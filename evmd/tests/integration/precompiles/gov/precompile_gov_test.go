package gov

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/precompiles/gov"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestGovPrecompileTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.GovPrecompileApp](integration.CreateEvmd, "evm.GovPrecompileApp")
	s := gov.NewPrecompileTestSuite(create)
	suite.Run(t, s)
}

func TestGovPrecompileIntegrationTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.GovPrecompileApp](integration.CreateEvmd, "evm.GovPrecompileApp")
	gov.TestPrecompileIntegrationTestSuite(t, create)
}
