package bech32

import (
	"testing"

	"github.com/stretchr/testify/suite"

	evm "github.com/arieschainorg/Aries-Chain"
	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/precompiles/bech32"
	testapp "github.com/arieschainorg/Aries-Chain/testutil/app"
)

func TestBech32PrecompileTestSuite(t *testing.T) {
	create := testapp.ToEvmAppCreator[evm.Bech32PrecompileApp](integration.CreateEvmd, "evm.Bech32PrecompileApp")
	s := bech32.NewPrecompileTestSuite(create)
	suite.Run(t, s)
}
