package merger

import (
    "github.com/ncipollo/fnew/config"
    "github.com/ncipollo/fnew/manifest"
    "github.com/ncipollo/fnew/workspace"
    "path"
)

type Merger interface {
    MergedManifest() *manifest.Manifest
}

type WorkspaceManifestMerger struct {
    configLoader   config.Loader
    manifestLoader manifest.Loader
    workspace      workspace.Workspace
}

func NewWorkspaceManifestMerger(configLoader config.Loader,
    manifestLoader manifest.Loader,
    workspace workspace.Workspace) Merger {
    return &WorkspaceManifestMerger{configLoader: configLoader, manifestLoader: manifestLoader, workspace: workspace}
}

func (merger WorkspaceManifestMerger) MergedManifest() *manifest.Manifest {
    defaultManifest := merger.defaultRepoManifest()
    configManifest := merger.configManifest()
    configRepoManifest := merger.configRepoManifest()

    mergedManifest := defaultManifest.Merge(*configRepoManifest).Merge(*configManifest)
    return &mergedManifest
}

func (merger WorkspaceManifestMerger) defaultRepoManifest() *manifest.Manifest {
    dir := merger.workspace.DefaultManifestRepoPath()
    fileName := path.Join(dir, manifest.FileName)
    loadedManifest, err := merger.manifestLoader.Load(fileName)
    if err != nil {
        return &manifest.Manifest{}
    }
    return loadedManifest
}

func (merger WorkspaceManifestMerger) configManifest() *manifest.Manifest {
    fileName := merger.workspace.ConfigPath()
    userConfig, err := merger.configLoader.Load(fileName)
    if err != nil {
        return &manifest.Manifest{}
    }
    return &userConfig.Manifest
}

func (merger WorkspaceManifestMerger) configRepoManifest() *manifest.Manifest {
    dir := merger.workspace.ConfigManifestRepoPath()
    fileName := path.Join(dir, manifest.FileName)
    loadedManifest, err := merger.manifestLoader.Load(fileName)
    if err != nil {
        return &manifest.Manifest{}
    }
    return loadedManifest
}
