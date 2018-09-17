package repo

import (
	"gopkg.in/src-d/go-git.v4"
	"os"
)

type Repo interface {
	Clone(localPath string, repoUrl string) (*git.Repository, error)
	Open(localPath string) (*git.Repository, error)
	Pull(repository *git.Repository) error
}

type GitRepo struct {
}

func New() Repo {
	return &GitRepo{}
}

func (GitRepo) Clone(localPath string, repoUrl string) (*git.Repository, error) {
	return git.PlainClone(localPath, false, &git.CloneOptions{
		URL:               repoUrl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
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
