package slashing

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/precompiles/slashing"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestSlashingPrecompileTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.SlashingPrecompileApp](integration.CreateEvmd, "evm.SlashingPrecompileApp")
	s := slashing.NewPrecompileTestSuite(create)
	suite.Run(t, s)
}

func TestStakingPrecompileIntegrationTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.SlashingPrecompileApp](integration.CreateEvmd, "evm.SlashingPrecompileApp")
	slashing.TestPrecompileIntegrationTestSuite(t, create)
}
