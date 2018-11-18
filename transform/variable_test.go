package transform

import (
    "bytes"
    "github.com/stretchr/testify/assert"
    "testing"
)

const variableTransformInputName = "input"
const variableTransformOutputName = "output"

func TestVariableTransform_Apply_DoesNotSkipIfVariableDoesNotExist(t *testing.T) {
    oldString := "old"
    newString := "new"
    variables := NewVariables()
    variables[variableTransformInputName] = oldString
    options := createVariableTransformOptions(StringReplace{Old: oldString, New: newString}, true)
    output, transform := createVariableTransformTestObjects(*options)

    transform.Apply(variables)

    assert.Equal(t, newString, variables[variableTransformOutputName])
    assertTransformMessage(t, transform.transformMessage(), output)
}

func TestVariableTransform_Apply_SkipsTransformWhenVariableExists(t *testing.T) {
    oldString := "old"
    newString := "new"
    variables := NewVariables()
    variables[variableTransformInputName] = oldString
    variables[variableTransformOutputName] = oldString
    options := createVariableTransformOptions(StringReplace{Old: oldString, New: newString}, true)
    output, transform := createVariableTransformTestObjects(*options)

    transform.Apply(variables)

    assert.Equal(t, oldString, variables[variableTransformOutputName])
    assertTransformMessage(t, transform.transformMessage(), output)
    assertTransformMessage(t, transform.skipMessage(), output)
}

func TestVariableTransform_Apply_TransformsVariable(t *testing.T) {
    oldString := "old"
    newString := "new"
    variables := NewVariables()
    variables[variableTransformInputName] = oldString
    variables[variableTransformOutputName] = oldString
    options := createVariableTransformOptions(StringReplace{Old: oldString, New: newString}, false)
    output, transform := createVariableTransformTestObjects(*options)

    transform.Apply(variables)

    assert.Equal(t, newString, variables[variableTransformOutputName])
    assertTransformMessage(t, transform.transformMessage(), output)
}

func createVariableTransformOptions(replace StringReplace, skipIfExists bool) *Options {
    return &Options{
        InputVariable:        variableTransformInputName,
        OutputVariable:       variableTransformOutputName,
        SkipIfVariableExists: skipIfExists,
        StringReplace:        replace,
    }
}

func createVariableTransformTestObjects(options Options) (*bytes.Buffer, *VariableTransform) {
    output := new(bytes.Buffer)
    transform := NewVariableTransform(options, output)
    return output, transform
}
