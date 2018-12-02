package project

import "github.com/stretchr/testify/mock"

type MockLoader struct {
    mock.Mock
}

func NewMockLoader(filename string, project *Project, err error) *MockLoader {
    loader := MockLoader{}
    loader.On("Load", filename).Return(project, err)
    return &loader
}

func (loader *MockLoader) Load(filename string) (*Project, error) {
    args := loader.Called(filename)
    return args.Get(0).(*Project), args.Error(1)
}


