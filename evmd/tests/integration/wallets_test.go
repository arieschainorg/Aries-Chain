package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/tests/integration/wallets"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestLedgerTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.IntegrationNetworkApp](CreateEvmd, "evm.IntegrationNetworkApp")
	s := wallets.NewLedgerTestSuite(create)
	suite.Run(t, s)
}
