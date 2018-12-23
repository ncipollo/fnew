package action

import (
    "github.com/ncipollo/fnew/manifest"
    "github.com/ncipollo/fnew/merger"
    "github.com/ncipollo/fnew/repo"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/ncipollo/fnew/testrepo"
    "github.com/ncipollo/fnew/workspace"
    "github.com/stretchr/testify/assert"
    "testing"
)

const localPath = "/projects"

var createActionManifestMerger = merger.NewMockMerger(*manifest.FullManifest())

func TestCreateAction_Perform_LocalPathAlreadyExists(t *testing.T) {
    localChecker := localPathChecker(true)
    actionRepo := createActionRepo(false)
    createAction := NewCreateAction(localChecker,
        localPath,
        createActionManifestMerger,
        manifest.MockProject1,
        actionRepo)

    err := createAction.Perform(testmessage.NewTestPrinter())

    assert.Error(t, err)
}

func TestCreateAction_Perform_NoProjectFounds(t *testing.T) {
    localChecker := localPathChecker(false)
    actionRepo := createActionRepo(false)
    createAction := NewCreateAction(localChecker,
        localPath,
        createActionManifestMerger,
        "invalidProject",
        actionRepo)

    err := createAction.Perform(testmessage.NewTestPrinter())

    assert.Error(t, err)
}

func TestCreateAction_Perform_RepoCloneFails(t *testing.T) {
    localChecker := localPathChecker(false)
    actionRepo := createActionRepo(true)
    createAction := NewCreateAction(localChecker,
        localPath,
        createActionManifestMerger,
        manifest.MockProject1,
        actionRepo)

    err := createAction.Perform(testmessage.NewTestPrinter())

    assert.Error(t, err)
}

func TestCreateAction_Perform_Success(t *testing.T) {
    localChecker := localPathChecker(false)
    actionRepo := createActionRepo(false)
    createAction := NewCreateAction(localChecker,
        localPath,
        createActionManifestMerger,
        manifest.MockProject1,
        actionRepo)

    err := createAction.Perform(testmessage.NewTestPrinter())

    assert.NoError(t, err)
}

func localPathChecker(exists bool) workspace.DirectoryChecker {
    return workspace.CreateMockDirectoryChecker(exists)
}

func createActionRepo(shouldError bool) repo.Repo {
    mockRepo := testrepo.NewMockRepo()
    mockRepo.StubClone(shouldError)
    return mockRepo
}
