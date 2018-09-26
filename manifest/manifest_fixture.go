package manifest

import (
	"github.com/stretchr/testify/mock"
	"net/url"
	"errors"
)

var url1, _ = url.Parse("http://www.example1.com")
var url2, _ = url.Parse("http://www.example2.com")

type MockLoader struct {
	mock.Mock
}

func NewMockLoader(shouldError bool) Loader  {
	mockLoader := MockLoader{}
	if shouldError {
		mockLoader.On("Load", mock.Anything).Return(nil, errors.New("no manifest for you"))
	} else {
		mockLoader.On("Load", mock.Anything).Return(FullManifest(), nil)
	}
	return &mockLoader
}

func (mockLoader *MockLoader) Load(filename string) (*Manifest, error) {
	args := mockLoader.Called(filename)

	return args.Get(0).(*Manifest), args.Error(1)
}

func FullManifest() *Manifest {
	return &Manifest{
		"project1": *url1,
		"project2": *url2,
	}
}
