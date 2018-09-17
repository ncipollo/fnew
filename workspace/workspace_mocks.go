package workspace

import (
	"github.com/stretchr/testify/mock"
	"errors"
)

const BasePath = "/test"

type MockDirectoryCreator struct {
	mock.Mock
}

func (creator MockDirectoryCreator) CreateDirectory(dir string) error {
	args := creator.Called(dir)
	return args.Error(0)
}

func createWorkSpace(creator MockDirectoryCreator) Workspace {
	return New(BasePath, creator)
}

func createDirectoryCreator(shouldError bool) MockDirectoryCreator {
	creator := new(MockDirectoryCreator)
	if shouldError {
		creator.On("CreateDirectory", "/test/manifests").Return(errors.New("I am err"))
	} else {
		creator.On("CreateDirectory", "/test/manifests").Return(nil)
	}
	return *creator
}
