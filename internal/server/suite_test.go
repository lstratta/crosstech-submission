package server

import (
	"testing"

	"github.com/lstratta/crosstech-submission/config"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	srv *Server
}

func TestSuite_Server(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (ts *TestSuite) SetupSuite() {
	ts.srv = &Server{
		conf: config.New(),
	}
}
