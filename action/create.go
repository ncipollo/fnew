package action

import (
	"github.com/ncipollo/fnew/workspace"
	"github.com/ncipollo/fnew/merger"
	"github.com/ncipollo/fnew/repo"
	"io"
)

type CreateAction struct {
	checker     workspace.DirectoryChecker
	localPath   string
	merger      merger.Merger
	projectName string
	repo        repo.Repo
}

func NewCreateAction(checker workspace.DirectoryChecker,
	localPath string,
	merger merger.Merger,
	projectName string,
	repo repo.Repo) *CreateAction {
	return &CreateAction{checker: checker, localPath: localPath, merger: merger, projectName: projectName, repo: repo}
}

func (action CreateAction) Perform(output io.Writer) error {
	return nil
}

