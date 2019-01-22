package action

import (
    "github.com/ncipollo/fnew/manifest"
    "github.com/ncipollo/fnew/merger"
    "github.com/ncipollo/fnew/testmessage"
    "testing"
)

func TestListAction_Perform_EmptyManifest(t *testing.T) {
    action := NewListAction(emptyManifestMerger())
    output := testmessage.NewTestPrinter()

    action.Perform(output)

    output.AssertMessage(t, noProjects)
}

func TestListAction_Perform_FullManifest(t *testing.T) {
    action := NewListAction(fullManifestMerger())
    output := testmessage.NewTestPrinter()

    action.Perform(output)

    output.AssertMessage(t, projectsHeader)
    output.AssertMessage(t, "a")
    output.AssertMessage(t, "b")
    output.AssertMessage(t, "c")
}

func emptyManifestMerger() merger.Merger {
    emptyManifest := manifest.Manifest{}
    return merger.NewMockMerger(emptyManifest)
}

func fullManifestMerger() merger.Merger {
    testUrl := "http://www.example1.com"
    emptyManifest := manifest.Manifest{
        "a": testUrl,
        "b": testUrl,
        "c": testUrl,
    }
    return merger.NewMockMerger(emptyManifest)
}
