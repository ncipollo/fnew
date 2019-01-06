package workspace

import (
    "github.com/stretchr/testify/assert"
    "os/user"
    "path"
    "testing"
    "github.com/ncipollo/fnew/config"
    "github.com/stretchr/testify/mock"
    "path/filepath"
    "errors"
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

func TestWorkspace_Setup_ConfigDoesNotExist(t *testing.T) {
    checker := CreateMockDirectoryChecker(false)
    configPath := filepath.Join(BasePath,"config.json")
    configWriter := config.NewMockWriter(config.Config{}, configPath, nil)
    creator := CreateMockDirectoryCreator(false)
    testWorkSpace := CreateMockWorkSpaceWithConfigWriter(checker, configWriter, creator)

    err := testWorkSpace.Setup()

    configWriter.AssertExpectations(t)
    assert.NoError(t, err)
}

func TestWorkspace_Setup_ConfigDoesNotExist_FailsToCreate(t *testing.T) {
    checker := CreateMockDirectoryChecker(false)
    configPath := filepath.Join(BasePath,"config.json")
    configWriter := config.NewMockWriter(config.Config{}, configPath, errors.New("config"))
    creator := CreateMockDirectoryCreator(false)
    testWorkSpace := CreateMockWorkSpaceWithConfigWriter(checker, configWriter, creator)

    err := testWorkSpace.Setup()

    configWriter.AssertExpectations(t)
    assert.Error(t, err)
}

func TestWorkspace_Setup_ConfigExists(t *testing.T) {
    checker := CreateMockDirectoryChecker(true)
    configWriter := config.NewMockWriter(config.Config{}, mock.Anything, nil)
    creator := CreateMockDirectoryCreator(false)
    testWorkSpace := CreateMockWorkSpaceWithConfigWriter(checker, configWriter, creator)

    err := testWorkSpace.Setup()

    configWriter.AssertNotCalled(t, "Writer")
    assert.NoError(t, err)
}

func TestWorkspace_Setup_Error(t *testing.T) {
    checker := CreateMockDirectoryChecker(true)
    creator := CreateMockDirectoryCreator(true)
    testWorkSpace := CreateMockWorkSpace(checker, creator)

    err := testWorkSpace.Setup()

    creator.AssertExpectations(t)
    assert.Error(t, err)
}

func TestWorkspace_Setup_NoError(t *testing.T) {
    checker := CreateMockDirectoryChecker(true)
    creator := CreateMockDirectoryCreator(false)
    testWorkSpace := CreateMockWorkSpace(checker, creator)

    err := testWorkSpace.Setup()

    creator.AssertExpectations(t)
    assert.NoError(t, err)
}
