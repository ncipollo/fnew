package action

import (
    "github.com/ncipollo/fnew/repo"
    "github.com/ncipollo/fnew/message"
)

type CleanupAction struct {
    localPath string
    repo      repo.Repo
}

func NewCleanupAction(localPath string, repo repo.Repo) *CleanupAction {
    return &CleanupAction{localPath: localPath, repo: repo}
}

func (action *CleanupAction) Perform(output message.Printer) error {
    err := action.repo.Delete(action.localPath)
    if err != nil {
        return err
    }

    _, err = action.repo.Init(action.localPath)
    return err
}
