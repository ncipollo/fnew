package action

import (
    "errors"
    "github.com/ncipollo/fnew/project"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/ncipollo/fnew/transform"
    "github.com/stretchr/testify/assert"
    "testing"
    "github.com/ncipollo/fnew/workspace"
)

const transformActionRepoPath = "path/to/repo"

func TestTransformAction_Perform_FailsOnLoaderError(t *testing.T) {
    action, _ := createTransformActionTestObjects(true, errors.New("loader"), nil)

    err := action.Perform(testmessage.NewTestPrinter())

    assert.Error(t, err)
}

func TestTransformAction_Perform_FailsOnTransformerError(t *testing.T) {
    action, _ := createTransformActionTestObjects(true, nil, errors.New("transformer"))

    err := action.Perform(testmessage.NewTestPrinter())

    assert.Error(t, err)
}

func TestTransformAction_Perform_Success(t *testing.T) {
    action, transformer := createTransformActionTestObjects(true, nil, nil)

    err := action.Perform(testmessage.NewTestPrinter())

    assert.NoError(t, err)
    transformer.AssertExpectations(t)
}

func TestTransformAction_Perform_Success_NoProject(t *testing.T) {
    action, _ := createTransformActionTestObjects(false, nil, errors.New("transformer"))

    err := action.Perform(testmessage.NewTestPrinter())

    assert.NoError(t, err)
}

func createTransformActionTestObjects(
    projectExists bool,
    loaderError error,
    transformerError error) (*TransformAction, *transform.MockTransformer) {

    repoProject := &project.Project{
        Transforms: []transform.Options{
            {
                Type: transform.TypeFileMove,
            },
        },
    }

    variables := transform.NewVariables()
    projectPath := project.PathInRepo(transformActionRepoPath)
    loader := project.NewMockLoader(projectPath, repoProject, loaderError)

    transforms, _ := transform.NewFactory().Transforms(repoProject.Transforms)

    transformer := transform.NewMockTransformer(transforms, variables, transformerError)

    checker := workspace.CreateMockDirectoryChecker(projectExists)

    return NewTransformAction(transformActionRepoPath,
        checker,
        loader,
        transformer,
        variables), transformer
}
