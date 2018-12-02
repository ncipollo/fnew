package action

import (
    "testing"
    "github.com/ncipollo/fnew/project"
    "github.com/ncipollo/fnew/transform"
    "github.com/ncipollo/fnew/message"
    "github.com/stretchr/testify/assert"
    "errors"
)

const transformActionRepoPath = "path/to/repo"

func TestTransformAction_Perform_FailsOnLoaderError(t *testing.T) {
    action, _ := createTransformActionTestObjects(errors.New("loader"), nil)

    err := action.Perform(message.NewTestPrinter())

    assert.Error(t, err)
}

func TestTransformAction_Perform_FailsOnTransformerError(t *testing.T) {
    action, _ := createTransformActionTestObjects(nil, errors.New("transformer"))

    err := action.Perform(message.NewTestPrinter())

    assert.Error(t, err)
}

func TestTransformAction_Perform_Success(t *testing.T) {
    action, transformer := createTransformActionTestObjects(nil, nil)

    err := action.Perform(message.NewTestPrinter())

    assert.NoError(t, err)
    transformer.AssertExpectations(t)
}

func createTransformActionTestObjects(
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

    return NewTransformAction(transformActionRepoPath, loader, transformer, variables), transformer
}
