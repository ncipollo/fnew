package repo

import (
    "gopkg.in/src-d/go-git.v4"
    "os"
    "gopkg.in/src-d/go-git.v4/plumbing/protocol/packp/sideband"
)

type Repo interface {
    Clone(localPath string, repoUrl string) (*git.Repository, error)
    Open(localPath string) (*git.Repository, error)
    Pull(repository *git.Repository) error
}

type GitRepo struct {
    verbose bool
}

func New(verbose bool) Repo {
    return &GitRepo{verbose: verbose}
}

func (repo *GitRepo) Clone(localPath string, repoUrl string) (*git.Repository, error) {
    return git.PlainClone(localPath, false, &git.CloneOptions{
        URL:               repoUrl,
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
        Progress:          repo.progress(),
    })
}

func (GitRepo) Open(localPath string) (*git.Repository, error) {
    return git.PlainOpen(localPath)
}

func (GitRepo) Pull(repository *git.Repository) error {
    tree, err := repository.Worktree()
    if err != nil {
        return err
    }
    return tree.Pull(&git.PullOptions{RemoteName: "origin"})
}

func (repo *GitRepo) progress() sideband.Progress {
    if repo.verbose {
        return os.Stdout
    } else {
        return nil
    }
}