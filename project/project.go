package project

import (
	"github.com/ncipollo/fnew/transform"
	"io/ioutil"
	"encoding/json"
)

type Project struct {
	Transforms []transform.Options `json:"transforms,omitempty"`
}

func FromFile(filename string) (*Project, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return FromJSON(data)
}

func FromJSON(data []byte) (*Project, error) {
	config := Project{}
	err := json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func FromString(jsonString string) (*Project, error) {
	return FromJSON([]byte(jsonString))
}