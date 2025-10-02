package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"wastu/tests"
)

type ExampleTestSuite struct {
	suite.Suite
	tests.TestCase
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func (s *ExampleTestSuite) SetupTest() {
}

func (s *ExampleTestSuite) TearDownTest() {
}

func (s *ExampleTestSuite) TestIndex() {
	s.True(true)
}
