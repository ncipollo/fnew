package transform

import (
    "github.com/ncipollo/fnew/testmessage"
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
    options := createVariableTransformOptions(StringReplace{Old: oldString, New: newString},
        true,
        "",
        "")
    transform, output := createVariableTransformTestObjects(*options)

    transform.Apply(variables)

    assert.Equal(t, newString, variables[variableTransformOutputName])
    output.AssertMessage(t, transform.transformMessage())
}

func TestVariableTransform_Apply_SkipsTransformWhenVariableExists(t *testing.T) {
    oldString := "old"
    newString := "new"
    variables := NewVariables()
    variables[variableTransformInputName] = oldString
    variables[variableTransformOutputName] = oldString
    options := createVariableTransformOptions(
        StringReplace{Old: oldString, New: newString},
        true,
        "",
        "")
    transform, output := createVariableTransformTestObjects(*options)

    transform.Apply(variables)

    assert.Equal(t, oldString, variables[variableTransformOutputName])
    output.AssertMessage(t, transform.transformMessage())
    output.AssertMessage(t, transform.skipMessage())
}

func TestVariableTransform_Apply_TransformsVariable(t *testing.T) {
    oldString := "old"
    newString := "new"
    variables := NewVariables()
    variables[variableTransformInputName] = oldString
    variables[variableTransformOutputName] = oldString
    options := createVariableTransformOptions(StringReplace{Old: oldString, New: newString},
        false,
        "",
        "")
    transform, output := createVariableTransformTestObjects(*options)

    transform.Apply(variables)

    assert.Equal(t, newString, variables[variableTransformOutputName])
    output.AssertMessage(t, transform.transformMessage())
}

func TestVariableTransform_Apply_TransformsVariableWithPrefixAndSuffix(t *testing.T) {
    prefix := "prefix"
    suffix := "suffix"
    oldString := "old"
    newString := "new"
    variables := NewVariables()
    variables[variableTransformInputName] = oldString
    variables[variableTransformOutputName] = oldString
    options := createVariableTransformOptions(StringReplace{Old: oldString, New: newString,},
        false,
        prefix,
        suffix)
    transform, output := createVariableTransformTestObjects(*options)

    transform.Apply(variables)

    assert.Equal(t, prefix + newString + suffix, variables[variableTransformOutputName])
    output.AssertMessage(t, transform.transformMessage())
}

func createVariableTransformOptions(replace StringReplace,
    skipIfExists bool,
    prefix string,
    suffix string) *Options {
    return &Options{
        InputVariable:        variableTransformInputName,
        OutputVariable:       variableTransformOutputName,
        SkipIfVariableExists: skipIfExists,
        StringPrefix:         prefix,
        StringReplace:        replace,
        StringSuffix:         suffix,
    }
}

func createVariableTransformTestObjects(options Options) (*VariableTransform, *testmessage.TestPrinter) {
    output := testmessage.NewTestPrinter()
    transform := NewVariableTransform(options, output)
    return transform, output
}
