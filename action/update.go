package action

import (
    "github.com/ncipollo/fnew/workspace"
    "github.com/ncipollo/fnew/repo"
    "github.com/ncipollo/fnew/message"
    "gopkg.in/src-d/go-git.v4"
)

type UpdateAction struct {
    repo      repo.Repo
    workspace workspace.Workspace
}

func NewUpdateAction(repo repo.Repo, workspace workspace.Workspace) Action {
    return &UpdateAction{repo: repo, workspace: workspace}
}

func (action *UpdateAction) Perform(output message.Printer) error {
    if action.workspace.DefaultManifestRepoExists() {
        err := action.updateRepoPath(action.workspace.DefaultManifestRepoPath())
        if err != nil && err != git.NoErrAlreadyUpToDate {
            return err
        }
    }

    if action.workspace.ConfigManifestRepoExists() {
        err := action.updateRepoPath(action.workspace.ConfigManifestRepoPath())
        if err != nil && err != git.NoErrAlreadyUpToDate {
            return err
        }
    }

    return nil
}

func (action *UpdateAction) updateRepoPath(filename string) error {
    if action.workspace.DefaultManifestRepoExists() {
        gitRepo, err := action.repo.Open(filename)
        if err != nil {
            return err
        }

        return action.repo.Pull(gitRepo)
    }

    return nil
}
