package action

import (
    "github.com/ncipollo/fnew/config"
    "github.com/ncipollo/fnew/manifest"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/ncipollo/fnew/testrepo"
    "github.com/ncipollo/fnew/workspace"
    "github.com/stretchr/testify/mock"
    "testing"
    "path/filepath"
)

func TestSetupAction_Perform_ReposDoNotExist(t *testing.T) {
    configLoader := config.MockLoaderWithRepoUrl()
    mockRepo := mockRepo()
    mockWorkSpace := workspaceWithoutRepos()

    setupAction := NewSetupAction(configLoader, mockRepo, mockWorkSpace)
    setupAction.Perform(testmessage.NewTestPrinter())

    mockRepo.AssertCloneCalled(t, mockWorkSpace.DefaultManifestRepoPath(), manifest.DefaultRepository)
    mockRepo.AssertCloneCalled(t, mockWorkSpace.ConfigManifestRepoPath(), config.FullConfig().ManifestRepoUrl.String())
}

func TestSetupAction_Perform_ReposDoNotExist_NoConfigRepoUrl(t *testing.T) {
    configLoader := config.MockLoaderWithoutRepoUrl()
    mockRepo := mockRepo()
    mockWorkSpace := workspaceWithoutRepos()

    setupAction := NewSetupAction(configLoader, mockRepo, mockWorkSpace)
    setupAction.Perform(testmessage.NewTestPrinter())

    mockRepo.AssertCloneCalled(t, mockWorkSpace.DefaultManifestRepoPath(), manifest.DefaultRepository)
    mockRepo.AssertCloneNotCalled(t, mockWorkSpace.ConfigManifestRepoPath(), mock.Anything)
}

func TestSetupAction_Perform_ReposExist(t *testing.T) {
    configLoader := config.MockLoaderWithRepoUrl()
    mockRepo := mockRepo()
    mockWorkSpace := workspaceWithRepos()

    setupAction := NewSetupAction(configLoader, mockRepo, mockWorkSpace)
    setupAction.Perform(testmessage.NewTestPrinter())

    mockRepo.AssertCloneNotCalled(t, mockWorkSpace.DefaultManifestRepoPath(), manifest.DefaultRepository)
    mockRepo.AssertCloneNotCalled(t, mockWorkSpace.ConfigManifestRepoPath(), config.FullConfig().ManifestRepoUrl.String())
}

func mockRepo() *testrepo.MockRepo {
    mockRepo := testrepo.NewMockRepo()
    mockRepo.StubClone(false)
    return mockRepo
}

func workspaceWithRepos() workspace.Workspace {
    checker := workspace.CreateMockDirectoryChecker(true)
    configWriter := config.NewMockWriter(config.Config{},
        filepath.Join(workspace.BasePath, "config.json"),
        nil)
    creator := workspace.CreateMockDirectoryCreator(false)
    return workspace.CreateMockWorkSpaceWithConfigWriter(checker, configWriter, creator)
}

func workspaceWithoutRepos() workspace.Workspace {
    checker := workspace.CreateMockDirectoryChecker(false)
    configWriter := config.NewMockWriter(config.Config{},
        filepath.Join(workspace.BasePath, "config.json"),
        nil)
    creator := workspace.CreateMockDirectoryCreator(false)
    return workspace.CreateMockWorkSpaceWithConfigWriter(checker, configWriter, creator)
}
