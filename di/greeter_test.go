package di

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GreeterTestSuite struct {
	suite.Suite
}

func (s *GreeterTestSuite) TestGreet() {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jakob")

	s.Equal("Hello, Jakob", buffer.String())
}

func TestGreeterTestSuite(t *testing.T) {
	suite.Run(t, new(GreeterTestSuite))
}
