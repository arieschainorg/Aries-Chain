package mempool

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/mempool"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestMempoolIntegrationTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.IntegrationNetworkApp](integration.CreateEvmd, "evm.IntegrationNetworkApp")
	suite.Run(t, mempool.NewMempoolIntegrationTestSuite(create))
}
