package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/url"
)

const (
	validJSON = `
{
	"repo": "http://www.example.com",
	"manifest":	
	{
		"project1" : "http://www.example.com"
	}
}
`
	invalidJSON = `{blarg`
	invaldUrl   = `{"repo" : "http]://blarg:foo.[]"}`
)

func TestFromString_FailsToParseInvalidJson(t *testing.T) {
	_, err := FromString(invalidJSON)

	assert.Error(t, err)
}

func TestFromString_FailsToParseInvalidUrl(t *testing.T) {
	_, err := FromString(invaldUrl)

	assert.Error(t, err)
}

func TestFromString_ParsesJSON(t *testing.T) {
	parsedManifest, err := FromString(validJSON)

	assert.Equal(t, expectedConfig(), parsedManifest)
	assert.NoError(t, err)
}

func expectedConfig() Config {
	repoUrl, _ := url.Parse("http://www.example.com")
	return Config{
		ManifestRepoUrl: repoUrl,
		Manifest: map[string]url.URL{
			"project1": *repoUrl,
		},
	}
}
