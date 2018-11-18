package config

import (
    "github.com/stretchr/testify/mock"
    "net/url"
)

type MockLoader struct {
    mock.Mock
}

func (loader *MockLoader) Load(filename string) (*Config, error) {
    args := loader.Called(filename)

    return args.Get(0).(*Config), args.Error(1)
}

func MockLoaderWithRepoUrl() Loader {
    mockLoader := MockLoader{}
    mockLoader.On("Load", mock.Anything).Return(FullConfig(), nil)
    return &mockLoader
}

func MockLoaderWithoutRepoUrl() Loader {
    mockLoader := MockLoader{}
    mockLoader.On("Load", mock.Anything).Return(MinimalConfig(), nil)
    return &mockLoader
}

func FullConfig() *Config {
    repoUrl, _ := url.Parse("http://www.example.com")
    return &Config{
        ManifestRepoUrl: repoUrl,
        Manifest: map[string]url.URL{
            "project1": *repoUrl,
        },
    }
}

func MinimalConfig() *Config {
    return &Config{
        ManifestRepoUrl: nil,
        Manifest:        map[string]url.URL{},
    }
}
