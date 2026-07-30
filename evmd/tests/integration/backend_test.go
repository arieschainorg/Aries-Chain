package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/tests/integration/rpc/backend"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestBackend(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.IntegrationNetworkApp](CreateEvmd, "evm.IntegrationNetworkApp")
	s := backend.NewTestSuite(create)
	suite.Run(t, s)
}
