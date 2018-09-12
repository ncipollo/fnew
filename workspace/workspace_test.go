package workspace

import (
	"testing"
	"os/user"
	"github.com/stretchr/testify/assert"
	"path"
	"github.com/stretchr/testify/mock"
	"errors"
)

const BasePath = "/test"

func TestDirectory(t *testing.T) {
	currentUser, err := user.Current()
	assert.NoError(t, err)
	expectedPath := path.Join(currentUser.HomeDir, ".fnew")

	assert.Equal(t, expectedPath, Directory())
}

func TestWorkspace_ConfigPath(t *testing.T) {
	creator := createDirectoryCreator(false)
	testWorkSpace := createWorkSpace(creator)

	assert.Equal(t, "/test/config.json", testWorkSpace.ConfigPath())
}

func TestWorkspace_ManifestsPath(t *testing.T) {
	creator := createDirectoryCreator(false)
	testWorkSpace := createWorkSpace(creator)

	assert.Equal(t, "/test/manifests", testWorkSpace.ManifestsPath())
}

func TestWorkspace_Setup_Error(t *testing.T) {
	creator := createDirectoryCreator(true)
	testWorkSpace := createWorkSpace(creator)

	testWorkSpace.Setup()

	creator.AssertExpectations(t)
}

func TestWorkspace_Setup_NoError(t *testing.T) {
	creator := createDirectoryCreator(false)
	testWorkSpace := createWorkSpace(creator)

	testWorkSpace.Setup()

	creator.AssertExpectations(t)
}

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
		creator.On("CreateDirectory", "/test/config.json").Return(errors.New("I am err"))
	} else {
		creator.On("CreateDirectory", "/test/config.json").Return(nil)
		creator.On("CreateDirectory", "/test/manifests").Return(nil)
	}
	return *creator
}
