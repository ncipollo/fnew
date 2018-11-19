package transform

import "github.com/stretchr/testify/mock"

type MockGlobber struct {
    mock.Mock
}

func NewMockGlobber(paths []string, err error) *MockGlobber {
    mockGlobber := MockGlobber{}
    mockGlobber.On("Glob").Return(paths, err)
    return &mockGlobber
}

func (globber *MockGlobber) Glob(pattern string) (matches []string, err error) {
    args := globber.Called()

    return args.Get(0).([]string), args.Error(1)
}
