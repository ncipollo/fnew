package manifest

import (
	"net/url"
	"encoding/json"
)

type Manifest map[string]url.URL

func FromJSON(data []byte) (Manifest, error) {
	manifest := Manifest{}
	err := json.Unmarshal(data, &manifest)
	return manifest, err
}

func FromString(jsonString string) (Manifest, error) {
	return FromJSON([]byte(jsonString))
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
