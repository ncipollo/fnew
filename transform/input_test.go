package transform

import (
	"bytes"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

const inputTransformVariableName = "variable"

func TestInputTransform_Apply_DoesNotSkipIfVariableDoesNotExist(t *testing.T) {
	variables := NewVariables()
	options := Options{OutputVariable: inputTransformVariableName, SkipIfVariableExists: true}
	input, output, inputTransform := createInputTransformTestObjects(options)

	simulateUserInput("input", input)
	inputTransform.Apply(variables)

	assertInputTransformMessage(t, inputTransform.requestInputMessage(inputTransformVariableName), output)
}

func TestInputTransform_Apply_ErrorsIfNoVariable(t *testing.T) {
	variables := NewVariables()
	_, _, inputTransform := createInputTransformTestObjects(Options{})

	assert.Error(t, inputTransform.Apply(variables))
}

func TestInputTransform_Apply_SkipsTransformWhenVariableExists(t *testing.T) {
	variables := NewVariables()
	options := Options{OutputVariable: inputTransformVariableName, SkipIfVariableExists: true}
	_, output, inputTransform := createInputTransformTestObjects(options)

	existingVariable := "existingVariable"
	variables[inputTransformVariableName] = existingVariable
	inputTransform.Apply(variables)

	assertInputTransformMessage(t, inputTransform.skipInputMessage(inputTransformVariableName), output)
	assert.Equal(t, existingVariable, variables[inputTransformVariableName])
}

func TestInputTransform_Apply_UserProvidesVariable(t *testing.T) {
	variables := NewVariables()
	options := Options{OutputVariable: inputTransformVariableName}
	input, output, inputTransform := createInputTransformTestObjects(options)

	variables[inputTransformVariableName] = "existingVariable"
	userInput := "input"
	simulateUserInput(userInput, input)
	inputTransform.Apply(variables)

	assertInputTransformMessage(t, inputTransform.requestInputMessage(inputTransformVariableName), output)
	assert.Equal(t, userInput, variables[inputTransformVariableName])
}

func assertInputTransformMessage(t *testing.T, expectedMessage string, outputBuffer *bytes.Buffer) {
	outputMessage, _ := outputBuffer.ReadString('\n')
	assert.Equal(t, fmt.Sprintln(expectedMessage), outputMessage)
}

func createInputTransformTestObjects(options Options) (*bytes.Buffer, *bytes.Buffer, *InputTransform) {
	input := new(bytes.Buffer)
	output := new(bytes.Buffer)
	inputTransform := NewInputTransform(options, input, output)
	return input, output, inputTransform
}

func simulateUserInput(input string, inputBuffer *bytes.Buffer) {
	inputBuffer.WriteString(input)
	inputBuffer.WriteString("\n")
}
