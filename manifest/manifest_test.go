package manifest

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/url"
)

const (
	validJSON = `
{
	"project1" : "http://www.example1.com",
	"project2" : "http://www.example2.com"
}
`
	invalidJSON = `{blarg`
	invaldUrl = `{"project1" : "http]://blarg:foo.[]"}`
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

	assert.Equal(t, expectedManifest(), parsedManifest)
	assert.NoError(t, err)
}

func expectedManifest() Manifest {
	url1, _ := url.Parse("http://www.example1.com")
	url2, _ := url.Parse("http://www.example2.com")
	return Manifest{
		"project1": *url1,
		"project2": *url2,
	}
}
