package workspace

import (
	"github.com/stretchr/testify/mock"
)

const BasePath = "/test"

type MockDirectoryCreator struct {
	mock.Mock
}

func (creator MockDirectoryCreator) CreateDirectory(dir string) error {
	args := creator.Called(dir)
	return args.Error(0)
}

type MockDirectoryChecker struct {
	mock.Mock
}

func (checker MockDirectoryChecker) DirectoryExists(dir string) bool {
	args := checker.Called(dir)
	return args.Bool(0)
}


