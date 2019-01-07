package action

import (
    "github.com/ncipollo/fnew/config"
    "github.com/ncipollo/fnew/manifest"
    "github.com/ncipollo/fnew/merger"
    "github.com/ncipollo/fnew/project"
    "github.com/ncipollo/fnew/repo"
    "github.com/ncipollo/fnew/transform"
    "github.com/ncipollo/fnew/workspace"
)

type Factory struct {
    checker       workspace.DirectoryChecker
    configLoader  config.Loader
    localPath     string
    merger        merger.Merger
    projectLoader project.Loader
    projectName   string
    repo          repo.Repo
    transformer   transform.Transformer
    variables     transform.Variables
    workspace     workspace.Workspace
}

func NewFactory(localPath string, projectName string, verbose bool) *Factory {
    configLoader := config.NewLoader()
    configWriter := config.NewWriter()

    directoryChecker := workspace.OSDirectoryChecker()
    directoryCreator := workspace.OSDirectoryCreator()
    actionWorkspace := workspace.New(workspace.Directory(),
        configWriter,
        directoryChecker,
        directoryCreator)

    manifestLoader := manifest.NewFileLoader()
    manifestMerger := merger.NewWorkspaceManifestMerger(configLoader,
        manifestLoader,
        actionWorkspace)

    actionRepo := repo.New(verbose)
    projectLoader := project.NewLoader()
    transformer := transform.NewTransformer()
    variables := transform.NewVariables()

    return &Factory{
        checker:       directoryChecker,
        configLoader:  configLoader,
        localPath:     localPath,
        merger:        manifestMerger,
        projectName:   projectName,
        projectLoader: projectLoader,
        repo:          actionRepo,
        transformer:   transformer,
        variables:     variables,
        workspace:     actionWorkspace,
    }
}

func (factory *Factory) Create() Action {
    action := NewCreateAction(factory.checker,
        factory.localPath,
        factory.merger,
        factory.projectName,
        factory.repo)

    return action
}

func (factory *Factory) Cleanup() Action {
    action := NewCleanupAction(factory.localPath, factory.repo)
    return action
}

func (factory *Factory) List() Action {
    action := NewListAction(factory.merger)
    return action
}

func (factory *Factory) Setup() Action {
    action := NewSetupAction(factory.configLoader, factory.repo, factory.workspace)
    return action
}

func (factory *Factory) Transform() Action {
    action := NewTransformAction(factory.localPath,
        factory.checker,
        factory.projectLoader,
        factory.transformer,
        factory.variables)
    return action
}

func (factory *Factory) Update() Action {
    action := NewUpdateAction(factory.repo, factory.workspace)
    return action
}
