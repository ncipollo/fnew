package transform

import (
    "bytes"
    "github.com/ncipollo/fnew/message"
    "github.com/stretchr/testify/assert"
    "testing"
)

const inputTransformVariableName = "variable"

func TestInputTransform_Apply_DoesNotSkipIfVariableDoesNotExist(t *testing.T) {
    variables := NewVariables()
    options := Options{OutputVariable: inputTransformVariableName, SkipIfVariableExists: true}
    inputTransform, input, output := createInputTransformTestObjects(options)

    simulateUserInput("input", input)
    inputTransform.Apply(variables)

    output.AssertMessage(t, inputTransform.requestInputMessage(inputTransformVariableName))
}

func TestInputTransform_Apply_ErrorsIfNoVariable(t *testing.T) {
    variables := NewVariables()
    inputTransform, _, _ := createInputTransformTestObjects(Options{})

    assert.Error(t, inputTransform.Apply(variables))
}

func TestInputTransform_Apply_SkipsTransformWhenVariableExists(t *testing.T) {
    variables := NewVariables()
    options := Options{OutputVariable: inputTransformVariableName, SkipIfVariableExists: true}
    inputTransform, _, output := createInputTransformTestObjects(options)

    existingVariable := "existingVariable"
    variables[inputTransformVariableName] = existingVariable
    inputTransform.Apply(variables)

    output.AssertMessage(t, inputTransform.skipInputMessage(inputTransformVariableName))
    assert.Equal(t, existingVariable, variables[inputTransformVariableName])
}

func TestInputTransform_Apply_UserProvidesVariable(t *testing.T) {
    variables := NewVariables()
    options := Options{OutputVariable: inputTransformVariableName}
    inputTransform, input, output := createInputTransformTestObjects(options)

    variables[inputTransformVariableName] = "existingVariable"
    userInput := "input"
    simulateUserInput(userInput, input)
    inputTransform.Apply(variables)

    output.AssertMessage(t, inputTransform.requestInputMessage(inputTransformVariableName))
    assert.Equal(t, userInput, variables[inputTransformVariableName])
}

func createInputTransformTestObjects(options Options) (*InputTransform, *bytes.Buffer, *message.TestPrinter) {
    input := new(bytes.Buffer)
    output := message.NewTestPrinter()
    inputTransform := NewInputTransform(options, input, output)
    return inputTransform, input, output
}

func simulateUserInput(input string, inputBuffer *bytes.Buffer) {
    inputBuffer.WriteString(input)
    inputBuffer.WriteString("\n")
}
