package action

import (
	"github.com/ncipollo/fnew/workspace"
	"github.com/ncipollo/fnew/merger"
	"github.com/ncipollo/fnew/manifest"
	"github.com/ncipollo/fnew/repo"
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
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

	err := createAction.Perform(os.Stdout)

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

	err := createAction.Perform(os.Stdout)

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

	err := createAction.Perform(os.Stdout)

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

	err := createAction.Perform(os.Stdout)

	assert.NoError(t, err)
}

func localPathChecker(exists bool) workspace.DirectoryChecker {
	return workspace.CreateMockDirectoryChecker(exists)
}

func createActionRepo(shouldError bool) repo.Repo {
	mockRepo := repo.NewMockRepo()
	mockRepo.StubClone(shouldError)
	return mockRepo
}
