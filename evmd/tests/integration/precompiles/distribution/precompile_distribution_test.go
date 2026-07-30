package distribution

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/precompiles/distribution"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestDistributionPrecompileTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.DistributionPrecompileApp](integration.CreateEvmd, "evm.DistributionPrecompileApp")
	s := distribution.NewPrecompileTestSuite(create)
	suite.Run(t, s)
}

func TestDistributionPrecompileIntegrationTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.DistributionPrecompileApp](integration.CreateEvmd, "evm.DistributionPrecompileApp")
	distribution.TestPrecompileIntegrationTestSuite(t, create)
}
