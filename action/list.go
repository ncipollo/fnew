package action

import (
    "github.com/ncipollo/fnew/manifest"
    "github.com/ncipollo/fnew/merger"
    "sort"
    "github.com/ncipollo/fnew/message"
)

const projectsHeader = "Available projects:"
const noProjects = "No projects found"

type ListAction struct {
    merger merger.Merger
}

func NewListAction(merger merger.Merger) ListAction {
    return ListAction{merger: merger}
}

func (action ListAction) Perform(output message.Printer) error {
    mergedManifest := action.merger.MergedManifest()
    if len(*mergedManifest) > 0 {
        action.printProjects(*mergedManifest, output)
    } else {
        action.printNoProjects(output)
    }

    return nil
}

func (action ListAction) printProjects(mergedManifest manifest.Manifest, output message.Printer) {
    keys := action.sortedManifestProjects(mergedManifest)

    output.Println(projectsHeader)
    for _, key := range keys {
        output.Println(key)
    }
}

func (action ListAction) sortedManifestProjects(mergedManifest manifest.Manifest) []string {
    keys := make([]string, 0, len(mergedManifest))
    for key := range mergedManifest {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    return keys
}

func (action ListAction) printNoProjects(output message.Printer) {
    output.Println(noProjects)
}
