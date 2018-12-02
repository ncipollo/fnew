package action

import (
    "github.com/ncipollo/fnew/project"
    "github.com/ncipollo/fnew/transform"
    "github.com/ncipollo/fnew/message"
)

type TransformAction struct {
    repoPath    string
    loader      project.Loader
    transformer transform.Transformer
    variables   transform.Variables
}

func NewTransformAction(repoPath string,
    loader project.Loader,
    transformer transform.Transformer,
    variables transform.Variables) *TransformAction {
    return &TransformAction{repoPath: repoPath,
        loader:      loader,
        transformer: transformer,
        variables:   variables}
}

func (action *TransformAction) Perform(output message.Printer) error {
    repoProject, err := action.createProject()

    if err != nil {
        return err
    }

    transforms, err := action.createTransforms(*repoProject)

    if err != nil {
        return err
    }

    err = action.transformer.Apply(transforms, action.variables)

    return err
}

func (action *TransformAction) createProject() (*project.Project, error) {
    projectPath := project.PathInRepo(action.repoPath)
    return action.loader.Load(projectPath)
}

func (action *TransformAction) createTransforms(project project.Project) ([]transform.Transform, error) {
    factory := transform.NewFactory()
    return factory.Transforms(project.Transforms)
}
