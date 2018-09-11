package workspace

import (
	"testing"
	"os/user"
	"github.com/stretchr/testify/assert"
	"path"
	"github.com/stretchr/testify/mock"
	"errors"
)

const BASE_PATH = "/test"

func TestDirectory(t *testing.T) {
	currentUser, err := user.Current()
	assert.NoError(t, err)
	expectedPath := path.Join(currentUser.HomeDir, ".fnew")

	assert.Equal(t, expectedPath, Directory())
}

func TestWorkspace_ConfigPath(t *testing.T) {
	testWorkSpace := createWorkSpace(false)

	assert.Equal(t, "/test/config.json", testWorkSpace.ConfigPath())
}

func TestWorkspace_ManifestsPath(t *testing.T) {
	testWorkSpace := createWorkSpace(false)

	assert.Equal(t, "/test/manifests", testWorkSpace.ManifestsPath())
}

type MockDirectoryCreator struct {
	mock.Mock
}

func (creator MockDirectoryCreator) CreateDirectory(dir string) error {
	args := creator.Called(dir)
	return args.Error(0)
}

func createWorkSpace(shouldError bool) Workspace {
	creator := createDirectoryCreator(shouldError)
	return New(BASE_PATH, creator)
}

func createDirectoryCreator(shouldError bool) DirectoryCreator {
	creator := new(MockDirectoryCreator)
	if shouldError {
		creator.On("CreateDirectory").Return(errors.New("I am err"))
	} else {
		creator.On("CreateDirectory").Return(nil)
	}
	return creator
}