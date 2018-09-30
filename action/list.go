package action

import (
	"github.com/ncipollo/fnew/merger"
	"io"
	"fmt"
	"github.com/ncipollo/fnew/manifest"
	"sort"
)

const projectsHeader  = "Available projects:"
const noProjects = "No projects found"

type ListAction struct {
	merger merger.Merger
}

func NewListAction(merger merger.Merger) ListAction {
	return ListAction{merger: merger}
}

func (action ListAction) Perform(output io.Writer) error {
	mergedManifest := action.merger.MergedManifest()
	if len(*mergedManifest) > 0 {
		action.printProjects(*mergedManifest, output)
	} else {
		action.printNoProjects(output)
	}

	return nil
}

func (action ListAction) printProjects(mergedManifest manifest.Manifest ,output io.Writer) {
	keys := action.sortedManifestProjects(mergedManifest)

	fmt.Fprintln(output, projectsHeader)
	for key := range keys {
		fmt.Fprintln(output, key)
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

func (action ListAction) printNoProjects(output io.Writer) {
	fmt.Fprintln(output, noProjects)
}