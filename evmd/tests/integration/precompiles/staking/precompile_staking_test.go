package staking

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/precompiles/staking"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestStakingPrecompileTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.StakingPrecompileApp](integration.CreateEvmd, "evm.StakingPrecompileApp")
	s := staking.NewPrecompileTestSuite(create)
	suite.Run(t, s)
}

func TestStakingPrecompileIntegrationTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.StakingPrecompileApp](integration.CreateEvmd, "evm.StakingPrecompileApp")
	staking.TestPrecompileIntegrationTestSuite(t, create)
}
