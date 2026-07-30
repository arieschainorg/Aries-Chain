package ante

import (
	"testing"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/stretchr/testify/suite"

	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/ante"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestEvmUnitAnteTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.AnteIntegrationApp](integration.CreateEvmd, "evm.AnteIntegrationApp")
	suite.Run(t, ante.NewEvmUnitAnteTestSuite(create))
}
