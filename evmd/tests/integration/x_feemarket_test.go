package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/tests/integration/x/feemarket"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestFeeMarketKeeperTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.IntegrationNetworkApp](CreateEvmd, "evm.IntegrationNetworkApp")
	s := feemarket.NewTestKeeperTestSuite(create)
	suite.Run(t, s)
}
