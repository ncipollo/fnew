package action

import (
    "github.com/ncipollo/fnew/config"
    "github.com/ncipollo/fnew/manifest"
    "github.com/ncipollo/fnew/repo"
    "github.com/ncipollo/fnew/workspace"
    "github.com/ncipollo/fnew/message"
)

type SetupAction struct {
    configLoader config.Loader
    repo         repo.Repo
    workspace    workspace.Workspace
}

func NewSetupAction(configLoader config.Loader,
    repo repo.Repo,
    workspace workspace.Workspace) Action {
    return &SetupAction{configLoader: configLoader, repo: repo, workspace: workspace}
}

func (action *SetupAction) Perform(output message.Printer) error {
    err := action.workspace.Setup()
    if err != nil {
        return err
    }
    userConfig, err := action.configLoader.Load(action.workspace.ConfigPath())
    if err != nil {
        return err
    }
    err = action.fetchConfigManifestIfNeeded(userConfig)
    if err != nil {
        return err
    }

    err = action.fetchDefaultManifestIfNeeded()
    if err != nil {
        return err
    }
    return nil
}

func (action *SetupAction) fetchConfigManifestIfNeeded(userConfig *config.Config) error {
    if userConfig.ManifestRepoUrl == "" {
        return nil
    }
    if !action.workspace.ConfigManifestRepoExists() {
        _, err := action.repo.Clone(action.workspace.ConfigManifestRepoPath(), userConfig.ManifestRepoUrl)
        return err
    }
    return nil
}

func (action *SetupAction) fetchDefaultManifestIfNeeded() error {
    if !action.workspace.DefaultManifestRepoExists() {
        _, err := action.repo.Clone(action.workspace.DefaultManifestRepoPath(), manifest.DefaultRepository)
        return err
    }
    return nil
}
