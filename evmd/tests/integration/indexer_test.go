package integration

import (
	"testing"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/tests/integration/indexer"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestKVIndexer(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.IntegrationNetworkApp](CreateEvmd, "evm.IntegrationNetworkApp")
	indexer.TestKVIndexer(t, create)
}
