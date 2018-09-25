package repo

import (
	"github.com/stretchr/testify/mock"
	"gopkg.in/src-d/go-git.v4"
	"errors"
	"testing"
)

type MockRepo struct {
	mock.Mock
}

func NewMockRepo() *MockRepo {
	return &MockRepo{}
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

func (repo *MockRepo) AssertCloneCalled(t *testing.T, localPath string, repoUrl string)  {
	repo.AssertCalled(t, "Clone", localPath, repoUrl)
}

func (repo *MockRepo) AssertCloneNotCalled(t *testing.T)  {
	repo.AssertNotCalled(t, "Clone", mock.Anything, mock.Anything)
}

func (repo *MockRepo) StubClone(shouldError bool)  {
	if shouldError {
		repo.On("Clone", mock.Anything, mock.Anything).Return(nil, errors.New("you're the clone"))
	} else {
		repo.On("Clone", mock.Anything, mock.Anything).Return(&git.Repository{}, nil)
	}
}

func (repo *MockRepo) StubOpen(shouldError bool)  {
	if shouldError {
		repo.On("Open", mock.Anything).Return(nil, errors.New("closed"))
	} else {
		repo.On("Open", mock.Anything).Return(git.Repository{}, nil)
	}
}

func (repo *MockRepo) StubPull(shouldError bool)  {
	if shouldError {
		repo.On("Pull", mock.Anything).Return(errors.New("push"))
	} else {
		repo.On("Pull", mock.Anything).Return(nil)
	}
}