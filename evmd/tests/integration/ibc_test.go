package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/tests/integration/x/ibc"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestIBCKeeperTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.IBCIntegrationApp](CreateEvmd, "evm.IBCIntegrationApp")
	s := ibc.NewKeeperTestSuite(create)
	suite.Run(t, s)
}
