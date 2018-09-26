package merger

import (
	"testing"
	"github.com/ncipollo/fnew/config"
	"github.com/ncipollo/fnew/workspace"
	"path"
	"github.com/ncipollo/fnew/manifest"
	"github.com/magiconair/properties/assert"
)

func TestWorkspaceManifestMerger_MergedManifest(t *testing.T) {
	configLoader := config.MockLoaderWithoutRepoUrl()
	manifestLoader := manifest.NewMockLoader(false)
	currentWorkspace := workspace.CreateMockWorkSpace(workspace.CreateMockDirectoryChecker(true),
		workspace.CreateMockDirectoryCreator(false))
	merger := NewWorkspaceManifestMerger(configLoader, manifestLoader, currentWorkspace)

	mergedManifest := merger.MergedManifest()

	manifestLoader.AssertCalled(t, "Load", path.Join(currentWorkspace.DefaultManifestRepoPath(), manifest.FileName))
	manifestLoader.AssertCalled(t, "Load", path.Join(currentWorkspace.ConfigManifestRepoPath(), manifest.FileName))
	assert.Equal(t, mergedManifest, manifest.FullManifest())
}
