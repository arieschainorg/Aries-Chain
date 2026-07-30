package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/tests/integration/x/ibc/callbacks"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestIBCCallback(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.IBCCallbackIntegrationApp](CreateEvmd, "evm.IBCCallbackIntegrationApp")
	suite.Run(t, callbacks.NewKeeperTestSuite(create))
}
