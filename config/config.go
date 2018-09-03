package config

import (
	"net/url"
	"github.com/ncipollo/fnew/manifest"
	"encoding/json"
)

type Config struct {
	ManifestRepoUrl *url.URL
	Manifest        manifest.Manifest
}

type rawConfig struct {
	ManifestRepoUrl string            `json:"manifest_repo_url,omitempty"`
	Manifest        manifest.Manifest `json:"manifest,omitempty"`
}

func (config *Config) MarshalJSON() ([]byte, error) {
	rawConfig := rawConfig{
		ManifestRepoUrl: config.ManifestRepoUrl.String(),
		Manifest:        config.Manifest,
	}

	return json.Marshal(rawConfig)
}

func (config *Config) UnmarshalJSON(data []byte) error {
	var rawConfig rawConfig
	err := json.Unmarshal(data, &rawConfig)
	if err != nil {
		return err
	}

	if rawConfig.ManifestRepoUrl != "" {
		config.ManifestRepoUrl, err = url.Parse(rawConfig.ManifestRepoUrl)
		if err != nil {
			return err
		}
	} else {
		config.ManifestRepoUrl = nil
	}
	config.Manifest = rawConfig.Manifest

	return nil
}
