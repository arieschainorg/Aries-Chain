package common

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/arieschainorg/Aries-Chain/evmd/tests/integration"
	"github.com/arieschainorg/Aries-Chain/tests/integration/precompiles/common"
)

func TestStaticCallTestSuite(t *testing.T) {
	s := common.NewStaticCallTestSuite(integration.CreateEvmd)
	suite.Run(t, s)
}
