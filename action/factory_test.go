package action

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

const factoryLocalPath = "/"
const factoryProjectName = "test"

func TestFactory_Create(t *testing.T) {
    factory := createFactory()

    action := NewCreateAction(factory.checker,
        factory.LocalPath,
        factory.merger,
        factory.ProjectName,
        factory.repo)
    assert.Equal(t, action, factory.Create())
}

func TestFactory_Cleanup(t *testing.T) {
    factory := createFactory()

    action := NewCleanupAction(factory.LocalPath, factory.repo)
    assert.Equal(t, action, factory.Cleanup())
}

func TestFactory_List(t *testing.T) {
    factory := createFactory()

    action := NewListAction(factory.merger)
    assert.Equal(t, action, factory.List())
}

func TestFactory_Setup(t *testing.T) {
    factory := createFactory()

    action := NewSetupAction(factory.configLoader, factory.repo, factory.workspace)
    assert.Equal(t, action, factory.Setup())
}

func TestFactory_Transform(t *testing.T) {
    factory := createFactory()

    action := NewTransformAction(
        factory.LocalPath,
        factory.checker,
        factory.projectLoader,
        factory.transformer,
        factory.variables)
    assert.Equal(t, action, factory.Transform())
}

func TestFactory_Update(t *testing.T) {
    factory := createFactory()

    action := NewUpdateAction(factory.repo, factory.workspace)
    assert.Equal(t, action, factory.Update())
}

func createFactory() *Factory {
    return NewFactory(factoryLocalPath, factoryProjectName, false)
}
