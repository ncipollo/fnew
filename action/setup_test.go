package action

import (
	"testing"
	"github.com/ncipollo/fnew/config"
	"github.com/ncipollo/fnew/repo"
	"os"
	"github.com/ncipollo/fnew/workspace"
	"github.com/ncipollo/fnew/manifest"
)

func TestSetupAction_Perform_ReposDoNotExist(t *testing.T) {
	configLoader := config.MockLoaderWithRepoUrl()
	mockRepo := mockRepo()
	mockWorkSpace := workspaceWithoutRepos()

	setupAction := NewSetupAction(configLoader, mockRepo, mockWorkSpace)
	setupAction.Perform(os.Stdout)

	mockRepo.AssertCloneCalled(t, mockWorkSpace.DefaultManifestRepoPath(), manifest.DefaultRepository)
	mockRepo.AssertCloneCalled(t, mockWorkSpace.ConfigManifestRepoPath(), config.FullConfig().ManifestRepoUrl.String())
}

func TestSetupAction_Perform_ReposExist(t *testing.T) {
	configLoader := config.MockLoaderWithRepoUrl()
	mockRepo := mockRepo()
	mockWorkSpace := workspaceWithRepos()

	setupAction := NewSetupAction(configLoader, mockRepo, mockWorkSpace)
	setupAction.Perform(os.Stdout)

	mockRepo.AssertCloneNotCalled(t)
}

func mockRepo() *repo.MockRepo {
	mockRepo := repo.NewMockRepo()
	mockRepo.StubClone(false)
	return mockRepo
}

func workspaceWithRepos() workspace.Workspace {
	checker := workspace.CreateMockDirectoryChecker(true)
	creator := workspace.CreateMockDirectoryCreator(false)
	return workspace.CreateMockWorkSpace(checker, creator)
}

func workspaceWithoutRepos() workspace.Workspace {
	checker := workspace.CreateMockDirectoryChecker(false)
	creator := workspace.CreateMockDirectoryCreator(false)
	return workspace.CreateMockWorkSpace(checker, creator)
}
