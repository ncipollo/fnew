package transform

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFactory_Transforms_FailsOnInvalidType(t *testing.T) {
    factory := NewFactory()
    optionsList := []Options{
        {Type: TypeFileMove},
        {Type: TypeFileStringReplace},
        {Type: TypeInput},
        {Type: TypeRunScript},
        {Type: TypeVariableStringReplace},
        {Type: "invalid"},
    }

    tranaforms, err := factory.Transforms(optionsList)

    assert.Error(t, err)
    assert.Nil(t, tranaforms)
}

func TestFactory_Transforms_Success(t *testing.T) {
    factory := NewFactory()
    optionsList := []Options{
        {Type: TypeFileMove},
        {Type: TypeFileStringReplace},
        {Type: TypeInput},
        {Type: TypeRunScript},
        {Type: TypeVariableStringReplace},
    }

    tranaforms, err := factory.Transforms(optionsList)

    assert.NoError(t, err)
    assert.IsType(t, &FileMoveTransform{}, tranaforms[0])
    assert.IsType(t, &FileTransform{}, tranaforms[1])
    assert.IsType(t, &InputTransform{}, tranaforms[2])
    assert.IsType(t, &ScriptTransform{}, tranaforms[3])
    assert.IsType(t, &VariableTransform{}, tranaforms[4])
}
