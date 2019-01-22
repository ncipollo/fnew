package merger

import (
    "errors"
    "github.com/ncipollo/fnew/config"
    "github.com/ncipollo/fnew/manifest"
    "github.com/ncipollo/fnew/workspace"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "path"
    "testing"
)

const configProject = "config"
const defaultProject = "default"

var configUrl = "http://www.example/config.com"
var defaultUrl = "http://www.example/default.com"

func TestWorkspaceManifestMerger_ErrorLoadingManifestsReturnEmptyManifest(t *testing.T) {
    configLoader := config.MockLoaderWithoutRepoUrl()
    currentWorkspace := workspace.CreateMockWorkSpace(workspace.CreateMockDirectoryChecker(true),
        workspace.CreateMockDirectoryCreator(false))
    manifestLoader := mockLoaderWithErrors()
    merger := NewWorkspaceManifestMerger(configLoader, manifestLoader, currentWorkspace)

    mergedManifest := merger.MergedManifest()

    expectedManifest := &manifest.Manifest{}

    assert.Equal(t, expectedManifest, mergedManifest)
}

func TestWorkspaceManifestMerger_MergedManifest(t *testing.T) {
    configLoader := config.MockLoaderWithoutRepoUrl()
    currentWorkspace := workspace.CreateMockWorkSpace(workspace.CreateMockDirectoryChecker(true),
        workspace.CreateMockDirectoryCreator(false))
    manifestLoader := mockLoaderWithBothManifests(&currentWorkspace)
    merger := NewWorkspaceManifestMerger(configLoader, manifestLoader, currentWorkspace)

    mergedManifest := merger.MergedManifest()

    expectedManifest := &manifest.Manifest{
        configProject:  configUrl,
        defaultProject: defaultUrl,
    }

    assert.Equal(t, expectedManifest, mergedManifest)
}

func mockLoaderWithBothManifests(currentWorkspace *workspace.Workspace) manifest.Loader {
    mockLoader := manifest.MockLoader{}

    mockLoader.On("Load", configManifestPath(currentWorkspace)).Return(mockConfigManifest(), nil)
    mockLoader.On("Load", defaultManifestPath(currentWorkspace)).Return(mockDefaultManifest(), nil)

    return &mockLoader
}

func mockLoaderWithErrors() manifest.Loader {
    var noManifest *manifest.Manifest = nil
    mockLoader := manifest.MockLoader{}
    mockLoader.On("Load", mock.Anything).Return(noManifest, errors.New("whomp whomp"))
    return &mockLoader
}

func configManifestPath(currentWorkspace *workspace.Workspace) string {
    return path.Join(currentWorkspace.ConfigManifestRepoPath(), manifest.FileName)
}

func defaultManifestPath(currentWorkspace *workspace.Workspace) string {
    return path.Join(currentWorkspace.DefaultManifestRepoPath(), manifest.FileName)
}

func mockConfigManifest() *manifest.Manifest {
    return &manifest.Manifest{configProject: configUrl}
}

func mockDefaultManifest() *manifest.Manifest {
    return &manifest.Manifest{defaultProject: defaultUrl}
}
