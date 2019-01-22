package config

import (
    "encoding/json"
    "github.com/ncipollo/fnew/manifest"
    "io/ioutil"
    "os"
)

type Config struct {
    ManifestRepoUrl string `json:"repo,omitempty"`
    Manifest        manifest.Manifest `json:"manifest,omitempty"`
}

func FromFile(filename string) (*Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return FromJSON(data)
}

func FromJSON(data []byte) (*Config, error) {
    config := Config{Manifest: map[string]string{}}
    err := json.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}

func FromString(jsonString string) (*Config, error) {
    return FromJSON([]byte(jsonString))
}

func (config Config) WriteToFile(filename string, perm os.FileMode) error {
    jsonString := config.String()
    return ioutil.WriteFile(filename, []byte(jsonString), perm)
}

func (config Config) String() string {
    bytes, err := json.MarshalIndent(config, "", "  ")
    if err != nil {
        return ""
    }
    return string(bytes)
}

type Loader interface {
    Load(filename string) (*Config, error)
}

type FileLoader struct{}

func (FileLoader) Load(filename string) (*Config, error) {
    return FromFile(filename)
}

func NewLoader() Loader {
    return &FileLoader{}
}

type Writer interface {
    Write(config Config, filename string) error
}

type FileWriter struct{}

func NewWriter() Writer {
    return &FileWriter{}
}

func (FileWriter) Write(config Config, filename string) error {
    return config.WriteToFile(filename, 0777)
}
