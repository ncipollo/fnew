package workspace

import (
    "errors"
    "github.com/stretchr/testify/mock"
    "github.com/ncipollo/fnew/config"
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

func CreateMockWorkSpace(checker MockDirectoryChecker,
    creator MockDirectoryCreator) Workspace {
    configWriter := config.NewMockWriter(config.Config{}, "", nil)
    return New(BasePath, configWriter, checker, creator)
}

func CreateMockWorkSpaceWithConfigWriter(checker MockDirectoryChecker,
    configWriter config.Writer,
    creator MockDirectoryCreator) Workspace {
    return New(BasePath, configWriter, checker, creator)
}

func CreateMockDirectoryChecker(exists bool) MockDirectoryChecker {
    checker := new(MockDirectoryChecker)
    checker.On("DirectoryExists", mock.Anything).Return(exists)
    return *checker
}

func CreateMockDirectoryCreator(shouldError bool) MockDirectoryCreator {
    creator := new(MockDirectoryCreator)
    if shouldError {
        creator.On("CreateDirectory", "/test/manifests").Return(errors.New("I am err"))
    } else {
        creator.On("CreateDirectory", "/test/manifests").Return(nil)
    }
    return *creator
}
