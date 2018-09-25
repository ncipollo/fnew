package workspace

import (
	"testing"
	"os/user"
	"github.com/stretchr/testify/assert"
	"path"
)

func TestDirectory(t *testing.T) {
	currentUser, err := user.Current()
	assert.NoError(t, err)
	expectedPath := path.Join(currentUser.HomeDir, ".fnew")

	assert.Equal(t, expectedPath, Directory())
}

func TestWorkspace_ConfigPath(t *testing.T) {
	checker := CreateMockDirectoryChecker(true)
	creator := CreateMockDirectoryCreator(false)
	testWorkSpace := CreateMockWorkSpace(checker, creator)

	assert.Equal(t, "/test/config.json", testWorkSpace.ConfigPath())
}

func TestWorkspace_ManifestsPath(t *testing.T) {
	checker := CreateMockDirectoryChecker(true)
	creator := CreateMockDirectoryCreator(false)
	testWorkSpace := CreateMockWorkSpace(checker, creator)

	assert.Equal(t, "/test/manifests", testWorkSpace.ManifestsPath())
}

func TestWorkspace_ConfigManifestRepoPath(t *testing.T) {
	checker := CreateMockDirectoryChecker(true)
	creator := CreateMockDirectoryCreator(false)
	testWorkSpace := CreateMockWorkSpace(checker, creator)

	assert.Equal(t, "/test/manifests/config", testWorkSpace.ConfigManifestRepoPath())
}

func TestWorkspace_ConfigManifestRepoExists(t *testing.T) {
	checker := CreateMockDirectoryChecker(true)
	creator := CreateMockDirectoryCreator(false)
	testWorkSpace := CreateMockWorkSpace(checker, creator)

	assert.True(t, testWorkSpace.ConfigManifestRepoExists())
}

func TestWorkspace_ConfigManifestRepoDoesNotExists(t *testing.T) {
	checker := CreateMockDirectoryChecker(false)
	creator := CreateMockDirectoryCreator(false)
	testWorkSpace := CreateMockWorkSpace(checker, creator)

	assert.False(t, testWorkSpace.ConfigManifestRepoExists())
}

func TestWorkspace_DefaultManifestRepoPath(t *testing.T) {
	checker := CreateMockDirectoryChecker(true)
	creator := CreateMockDirectoryCreator(false)
	testWorkSpace := CreateMockWorkSpace(checker, creator)

	assert.Equal(t, "/test/manifests/default", testWorkSpace.DefaultManifestRepoPath())
}

func TestWorkspace_DefaultManifestRepoExists(t *testing.T) {
	checker := CreateMockDirectoryChecker(true)
	creator := CreateMockDirectoryCreator(false)
	testWorkSpace := CreateMockWorkSpace(checker, creator)

	assert.True(t, testWorkSpace.DefaultManifestRepoExists())
}

func TestWorkspace_DefaultManifestRepoDoesNotExists(t *testing.T) {
	checker := CreateMockDirectoryChecker(false)
	creator := CreateMockDirectoryCreator(false)
	testWorkSpace := CreateMockWorkSpace(checker, creator)

	assert.False(t, testWorkSpace.DefaultManifestRepoExists())
}

func TestWorkspace_Setup_Error(t *testing.T) {
	checker := CreateMockDirectoryChecker(true)
	creator := CreateMockDirectoryCreator(true)
	testWorkSpace := CreateMockWorkSpace(checker, creator)

	testWorkSpace.Setup()

	creator.AssertExpectations(t)
}

func TestWorkspace_Setup_NoError(t *testing.T) {
	checker := CreateMockDirectoryChecker(true)
	creator := CreateMockDirectoryCreator(false)
	testWorkSpace := CreateMockWorkSpace(checker, creator)

	testWorkSpace.Setup()

	creator.AssertExpectations(t)
}
