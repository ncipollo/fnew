package action

import (
    "fmt"
    "github.com/ncipollo/fnew/project"
    "github.com/ncipollo/fnew/transform"
    "github.com/ncipollo/fnew/message"
    "github.com/ncipollo/fnew/workspace"
)

type TransformAction struct {
    repoPath    string
    checker     workspace.DirectoryChecker
    loader      project.Loader
    transformer transform.Transformer
    variables   transform.Variables
}

func NewTransformAction(
    repoPath string,
    checker workspace.DirectoryChecker,
    loader project.Loader,
    transformer transform.Transformer,
    variables transform.Variables) *TransformAction {
    return &TransformAction{
        repoPath:    repoPath,
        checker:     checker,
        loader:      loader,
        transformer: transformer,
        variables:   variables,
    }
}

func (action *TransformAction) Perform(output message.Printer) error {

    if !action.checkProjectExists() {
        return nil
    }

    repoProject, err := action.createProject()

    if err != nil {
        return err
    }

    transforms, err := action.createTransforms(*repoProject)

    if err != nil {
        return err
    }

    if len(transforms) > 0 {
        output.Println(action.transformMessage())
    }

    err = action.transformer.Apply(transforms, action.variables)

    return err
}

func (action *TransformAction) checkProjectExists() bool {
    projectPath := project.PathInRepo(action.repoPath)
    return action.checker.DirectoryExists(projectPath)
}

func (action *TransformAction) createProject() (*project.Project, error) {
    projectPath := project.PathInRepo(action.repoPath)
    return action.loader.Load(projectPath)
}

func (action *TransformAction) createTransforms(project project.Project) ([]transform.Transform, error) {
    factory := transform.NewFactory()
    return factory.Transforms(project.Transforms)
}

func (action *TransformAction) transformMessage() string {
    return fmt.Sprintf("Applying Transforms")
}