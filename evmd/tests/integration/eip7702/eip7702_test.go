package eip7702

import (
	"testing"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/eip7702"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestEIP7702IntegrationTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.IntegrationNetworkApp](integration.CreateEvmd, "evm.IntegrationNetworkApp")
	eip7702.TestEIP7702IntegrationTestSuite(t, create)
}
