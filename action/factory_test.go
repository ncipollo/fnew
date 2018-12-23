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
        factory.localPath,
        factory.merger,
        factory.projectName,
        factory.repo)
    assert.Equal(t, action, factory.Create())
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

    action := NewTransformAction(factory.localPath,
        factory.projectLoader,
        factory.transformer,
        factory.variables)
    assert.Equal(t, action, factory.Transform())
}

func createFactory() *Factory {
    return &Factory{localPath: factoryLocalPath, projectName: factoryProjectName}
}
