package manifest

import (
	"net/url"
	"encoding/json"
	"io/ioutil"
	"github.com/ncipollo/fnew/workspace"
	"github.com/ncipollo/fnew/config"
	"path"
)

const FileName = "manifest.json"
const ConfigDirectory = "config"
const DefaultDirectory = "default"
const DefaultRepository = "https://github.com/ncipollo/fnew-manifest.git"

type Manifest map[string]url.URL

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

func (manifest Manifest) MarshalJSON() ([]byte, error) {
	rawManifest := map[string]string{}
	for key, repoUrl := range manifest {
		rawManifest[key] = repoUrl.String()
	}
	return json.Marshal(rawManifest)
}

func (manifest Manifest) UnmarshalJSON(data []byte) error {
	var rawManifest map[string]string
	err := json.Unmarshal(data, &rawManifest)
	if err != nil {
		return err
	}

	for key, value := range rawManifest {
		repoUrl, err := url.Parse(value)
		if err != nil {
			return err
		}
		manifest[key] = *repoUrl
	}

	return nil
}

type Loader interface {
	Load(filename string) (*Manifest, error)
}

type FileLoader struct{}

func (FileLoader) Load(filename string) (*Manifest, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return FromJSON(data)
}

type Merger interface {
	MergedManifest() *Manifest
}

type WorkspaceManifestMerger struct {
	configLoader   config.Loader
	manifestLoader Loader
	workspace      workspace.Workspace
}

func (merger WorkspaceManifestMerger) MergedManifest() *Manifest {
	defaultManifest := merger.defaultRepoManifest()
	configManifest := merger.configManifest()
	configRepoManifest := merger.configRepoManifest()

	mergedManifest := defaultManifest.Merge(*configManifest).Merge(*configRepoManifest)
	return &mergedManifest
}

func (merger WorkspaceManifestMerger) defaultRepoManifest() *Manifest {
	dir := merger.workspace.DefaultManifestRepoPath()
	fileName := path.Join(dir, FileName)
	manifest, err := merger.manifestLoader.Load(fileName)
	if err == nil {
		return &Manifest{}
	}
	return manifest
}

func (merger WorkspaceManifestMerger) configManifest() *Manifest {
	fileName := merger.workspace.ConfigPath()
	userConfig, err := merger.configLoader.Load(fileName)
	if err == nil {
		return &Manifest{}
	}
	return &userConfig.Manifest
}

func (merger WorkspaceManifestMerger) configRepoManifest() *Manifest {
	dir := merger.workspace.ConfigManifestRepoPath()
	fileName := path.Join(dir, FileName)
	manifest, err := merger.manifestLoader.Load(fileName)
	if err == nil {
		return &Manifest{}
	}
	return manifest
}
