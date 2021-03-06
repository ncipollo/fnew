package manifest

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

const (
    validJSON = `{
  "project1": "http://www.example1.com",
  "project2": "http://www.example2.com"
}`
    invalidJSON = `{blarg`
)

func TestDefaultDirectory(t *testing.T) {
    assert.Equal(t, "default", DefaultDirectory)
}

func TestConfigDirectory(t *testing.T) {
    assert.Equal(t, "config", ConfigDirectory)
}

func TestDefaultRepository(t *testing.T) {
    assert.Equal(t, "https://github.com/file-new/fnew-manifest.git", DefaultRepository)
}

func TestFromString_FailsToParseInvalidJson(t *testing.T) {
    _, err := FromString(invalidJSON)

    assert.Error(t, err)
}

func TestFromString_ParsesJSON(t *testing.T) {
    parsedManifest, err := FromString(validJSON)

    assert.Equal(t, FullManifest(), parsedManifest)
    assert.NoError(t, err)
}

func TestMerge(t *testing.T) {
    manifest1 := Manifest{
        "project1": url1,
        "project2": url2,
    }
    manifest2 := Manifest{
        "project2": url1,
        "project3": url2,
    }

    mergedManifest := Manifest{
        "project1": url1,
        "project2": url1,
        "project3": url2,
    }

    assert.Equal(t, mergedManifest, manifest1.Merge(manifest2))
}

func TestString(t *testing.T) {
    jsonString := FullManifest().String()

    assert.Equal(t, validJSON, jsonString)
}
