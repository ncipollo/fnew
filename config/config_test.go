package config

import (
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "net/url"
    "os"
    "path/filepath"
    "testing"
)

const (
    validJSON = `{
  "repo": "http://www.example.com",
  "manifest": {
    "project1": "http://www.example.com"
  }
}`
    invalidJSON = `{blarg`
    invalidUrl  = `{"repo" : "http]://blarg:foo.[]"}`
)

func TestFromFile(t *testing.T) {
    testDir, err := ioutil.TempDir("", "config")
    assert.NoError(t, err)

    defer os.RemoveAll(testDir)

    configPath := filepath.Join(testDir, "config.json")
    sourceConfig := FullConfig()
    sourceConfig.WriteToFile(configPath, 0777)

    fileConfig, err := FromFile(configPath)
    assert.NoError(t, err)
    assert.Equal(t, FullConfig(), fileConfig)
}

func TestFromString_FailsToParseInvalidJson(t *testing.T) {
    _, err := FromString(invalidJSON)

    assert.Error(t, err)
}

func TestFromString_FailsToParseInvalidUrl(t *testing.T) {
    _, err := FromString(invalidUrl)

    assert.Error(t, err)
}

func TestFromString_ParsesEmptyJSON(t *testing.T) {
    parsedManifest, err := FromString("{}")

    assert.Equal(t, &Config{Manifest: map[string]url.URL{}}, parsedManifest)
    assert.NoError(t, err)
}

func TestFromString_ParsesJSON(t *testing.T) {
    parsedManifest, err := FromString(validJSON)

    assert.Equal(t, FullConfig(), parsedManifest)
    assert.NoError(t, err)
}

func TestString(t *testing.T) {
    jsonString := FullConfig().String()

    assert.Equal(t, validJSON, jsonString)
}
