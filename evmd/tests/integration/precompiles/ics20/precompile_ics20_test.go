package ics20

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/precompiles/ics20"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

var ibcAppCreator = testapp.ToIBCAppCreator[evm.ICS20PrecompileApp](integration.SetupEvmd, "evm.ICS20PrecompileApp")

func TestICS20PrecompileTestSuite(t *testing.T) {
	s := ics20.NewPrecompileTestSuite(t, ibcAppCreator)
	suite.Run(t, s)
}

func TestICS20PrecompileIntegrationTestSuite(t *testing.T) {
	ics20.TestPrecompileIntegrationTestSuite(t, ibcAppCreator)
}
