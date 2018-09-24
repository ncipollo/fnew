package config

import "github.com/stretchr/testify/mock"

type MockLoader struct {
	mock.Mock
}

func (loader *MockLoader) Load(filename string) (*Config, error) {
	args := loader.Called(filename)

	return args.Get(0).(*Config), args.Error(1)
}
