package repo

import (
	"github.com/stretchr/testify/mock"
	"gopkg.in/src-d/go-git.v4"
)

type MockRepo struct {
	mock.Mock
}

func (repo *MockRepo) Clone(localPath string, repoUrl string) (*git.Repository, error) {
	args := repo.Called(localPath, repoUrl)
	return args.Get(0).(*git.Repository), args.Error(1)
}

func (repo *MockRepo) Open(localPath string) (*git.Repository, error) {
	args := repo.Called(localPath)
	return args.Get(0).(*git.Repository), args.Error(1)
}

func (repo *MockRepo) Pull(repository *git.Repository) error {
	args := repo.Called(repository)
	return args.Error(0)
}



