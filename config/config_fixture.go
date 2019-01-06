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
    return &Config{}
}

type MockWriter struct {
    mock.Mock
}

func (writer *MockWriter) Write(config Config, filename string) error {
    args := writer.Called(config, filename)
    return args.Error(0)
}

func NewMockWriter(config Config, filename string, err error) *MockWriter {
    mockWriter := MockWriter{}
    mockWriter.On("Write", config, filename).Return(err)
    return &mockWriter
}