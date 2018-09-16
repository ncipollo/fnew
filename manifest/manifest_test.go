package manifest

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/url"
)

const (
	validJSON = `{
  "project1": "http://www.example1.com",
  "project2": "http://www.example2.com"
}`
	invalidJSON = `{blarg`
	invalidUrl  = `{"project1" : "http]://blarg:foo.[]"}`
)

var url1, _ = url.Parse("http://www.example1.com")
var url2, _ = url.Parse("http://www.example2.com")

func TestDefaultRepository(t *testing.T) {
	assert.Equal(t, "https://github.com/ncipollo/fnew-manifest.git", DefaultRepository)
}

func TestFromString_FailsToParseInvalidJson(t *testing.T) {
	_, err := FromString(invalidJSON)

	assert.Error(t, err)
}

func TestFromString_FailsToParseInvalidUrl(t *testing.T) {
	_, err := FromString(invalidUrl)

	assert.Error(t, err)
}

func TestFromString_ParsesJSON(t *testing.T) {
	parsedManifest, err := FromString(validJSON)

	assert.Equal(t, jsonManifest(), parsedManifest)
	assert.NoError(t, err)
}

func TestMerge(t *testing.T) {
	manifest1 := Manifest{
		"project1": *url1,
		"project2": *url2,
	}
	manifest2 := Manifest{
		"project2": *url1,
		"project3": *url2,
	}

	mergedManifest := Manifest{
		"project1": *url1,
		"project2": *url1,
		"project3": *url2,
	}

	assert.Equal(t, mergedManifest, manifest1.Merge(manifest2))
}

func TestString(t *testing.T) {
	jsonString := jsonManifest().String()

	assert.Equal(t, validJSON, jsonString)
}

func jsonManifest() Manifest {
	return Manifest{
		"project1": *url1,
		"project2": *url2,
	}
}
