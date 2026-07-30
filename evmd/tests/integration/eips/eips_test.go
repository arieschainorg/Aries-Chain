package eips_test

import (
	"testing"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/eips"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
	//nolint:revive // dot imports are fine for Ginkgo
	//nolint:revive // dot imports are fine for Ginkgo
)

func TestEIPs(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.IntegrationNetworkApp](integration.CreateEvmd, "evm.IntegrationNetworkApp")
	eips.RunTests(t, create)
}
