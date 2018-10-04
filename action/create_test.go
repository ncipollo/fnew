package action

import (
	"github.com/ncipollo/fnew/workspace"
	"github.com/ncipollo/fnew/merger"
	"github.com/ncipollo/fnew/manifest"
	"github.com/ncipollo/fnew/repo"
)

const localPath = "/projects"
var createActionManifest = merger.NewMockMerger(*manifest.FullManifest())


func localPathChecker(exists bool) workspace.DirectoryChecker {
	return workspace.CreateMockDirectoryChecker(exists)
}

func createActionRepo(shouldError bool) repo.Repo {
	mockRepo := repo.NewMockRepo()
	mockRepo.StubClone(shouldError)
	return mockRepo
}