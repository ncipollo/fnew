package manifest

import (
    "encoding/json"
    "io/ioutil"
)

const FileName = "manifest.json"
const ConfigDirectory = "config"
const DefaultDirectory = "default"
const DefaultRepository = "https://github.com/file-new/fnew-manifest.git"

type Manifest map[string]string

func FromJSON(data []byte) (*Manifest, error) {
    manifest := Manifest{}
    err := json.Unmarshal(data, &manifest)
    return &manifest, err
}

func FromString(jsonString string) (*Manifest, error) {
    return FromJSON([]byte(jsonString))
}

func (manifest Manifest) Merge(other Manifest) Manifest {
    merged := Manifest{}
    for key, value := range manifest {
        merged[key] = value
    }
    for key, value := range other {
        merged[key] = value
    }
    return merged
}

func (manifest Manifest) String() string {
    bytes, err := json.MarshalIndent(manifest, "", "  ")
    if err != nil {
        return ""
    }
    return string(bytes)
}

type Loader interface {
    Load(filename string) (*Manifest, error)
}

type FileLoader struct{}

func NewFileLoader() Loader {
    return &FileLoader{}
}

func (FileLoader) Load(filename string) (*Manifest, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return FromJSON(data)
}
