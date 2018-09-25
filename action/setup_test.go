package action

import (
	"testing"
	"github.com/ncipollo/fnew/config"
	"github.com/ncipollo/fnew/repo"
	"github.com/ncipollo/fnew/workspace"
	"os"
)

func TestSetupAction_Perform(t *testing.T) {
	configLoader := config.MockLoaderWithRepoUrl()
	mockRepo := repo.NewMockRepo()
	mockRepo.StubClone(false)
	mockRepo.StubOpen(false)
	mockRepo.StubPull(false)
	checker  := workspace.CreateMockDirectoryChecker(true)
	creator := workspace.CreateMockDirectoryCreator(false)
	mockWorkSpace := workspace.CreateMockWorkSpace(checker, creator)

	setupAction := NewSetupAction(configLoader, mockRepo, mockWorkSpace)
	setupAction.Perform(os.Stdout)
}