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

type MockDirectoryChecker struct {
	mock.Mock
}

func (checker MockDirectoryChecker) DirectoryExists(dir string) bool {
	args := checker.Called(dir)
	return args.Bool(0)
}

func createMockWorkSpace(checker MockDirectoryChecker, creator MockDirectoryCreator) Workspace {
	return New(BasePath, checker, creator)
}

func createMockDirectoryChecker(exists bool) MockDirectoryChecker {
	checker := new(MockDirectoryChecker)
	checker.On("DirectoryExists", mock.Anything).Return(exists)
	return *checker
}

func createMockDirectoryCreator(shouldError bool) MockDirectoryCreator {
	creator := new(MockDirectoryCreator)
	if shouldError {
		creator.On("CreateDirectory", "/test/manifests").Return(errors.New("I am err"))
	} else {
		creator.On("CreateDirectory", "/test/manifests").Return(nil)
	}
	return *creator
}