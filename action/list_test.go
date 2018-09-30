package action

import (
	"github.com/ncipollo/fnew/merger"
	"github.com/ncipollo/fnew/manifest"
	"net/url"
	"testing"
	"bytes"
	"fmt"
	"github.com/magiconair/properties/assert"
)

func TestListAction_Perform_EmptyManifest(t *testing.T) {
	action := NewListAction(emptyManifestMerger())
	output := new(bytes.Buffer)

	action.Perform(output)

	expectedString := fmt.Sprintln(noProjects)
	assert.Equal(t, output.String(), expectedString)
}

func TestListAction_Perform_FullManifest(t *testing.T) {
	action := NewListAction(fullManifestMerger())
	output := new(bytes.Buffer)

	action.Perform(output)

	expectedString := fmt.Sprint(projectsHeader, "\n", "a\n", "b\n", "c\n")
	assert.Equal(t, output.String(), expectedString)
}

func emptyManifestMerger() merger.Merger {
	emptyManifest := manifest.Manifest{}
	return merger.NewMockMerger(emptyManifest)
}

func fullManifestMerger() merger.Merger {
	testUrl, _ := url.Parse("http://www.example1.com")
	emptyManifest := manifest.Manifest{
		"a": *testUrl,
		"b": *testUrl,
		"c": *testUrl,
	}
	return merger.NewMockMerger(emptyManifest)
}
