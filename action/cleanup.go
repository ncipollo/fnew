package action

import (
    "github.com/ncipollo/fnew/repo"
    "github.com/ncipollo/fnew/message"
)

type CleanupRepoAction struct {
    localPath string
    repo      repo.Repo
}

func NewCleanupRepoAction(localPath string, repo repo.Repo) *CleanupRepoAction {
    return &CleanupRepoAction{localPath: localPath, repo: repo}
}

func (action *CleanupRepoAction) Perform(output message.Printer) error {
    err := action.repo.Delete(action.localPath)
    if err != nil {
        return err
    }

    _, err = action.repo.Init(action.localPath)
    return err
}
