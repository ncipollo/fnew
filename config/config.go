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
	ManifestRepoUrl string            `json:"repo,omitempty"`
	Manifest        manifest.Manifest `json:"manifest,omitempty"`
}

func FromJSON(data []byte) (Config, error) {
	config := Config{}
	err := json.Unmarshal(data, &config)
	return config, err
}

func FromString(jsonString string) (Config, error) {
	return FromJSON([]byte(jsonString))
}

func (config Config) String() string {
	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (config Config) MarshalJSON() ([]byte, error) {
	rawConfig := rawConfig{
		ManifestRepoUrl: config.ManifestRepoUrl.String(),
		Manifest:        config.Manifest,
	}

	return json.Marshal(rawConfig)
}

func (config *Config) UnmarshalJSON(data []byte) error {
	rawConfig := rawConfig{Manifest: map[string]url.URL{}}
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
