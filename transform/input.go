package transform

import (
	"io"
	"fmt"
	"bufio"
	"strings"
	"errors"
)

type InputTransform struct {
	Options
	input  io.Reader
	output io.Writer
}

func NewInputTransform(options Options, input io.Reader, output io.Writer) *InputTransform {
	return &InputTransform{Options: options, input: input, output: output}
}

func (transform *InputTransform) Apply(variables Variables) error {
	variableName := transform.OutputVariable
	if len(variableName) == 0 {
		return errors.New("no output variable specified in transform")
	}
	variableValue := variables[variableName]

	if transform.SkipIfVariableExists {
		if len(variableValue) > 0 {
			message := transform.skipInputMessage(variableName)
			fmt.Fprintln(transform.output, message)
			return nil
		}
	}

	message := transform.requestInputMessage(variableName)
	fmt.Fprintln(transform.output, message)

	value, err := transform.readVariable()
	if err == nil {
		variables[variableName] = value
	}
	return err
}

func (InputTransform) requestInputMessage(variableName string) string {
	return fmt.Sprintf("Please enter value for %s", variableName)
}

func (InputTransform) skipInputMessage(variableName string) string {
	return fmt.Sprintf("%s has a value. Skipping...", variableName)
}

func (transform *InputTransform) readVariable() (string, error) {
	reader := bufio.NewReader(transform.input)
	variable, err := reader.ReadString('\n')
	if err == nil{
		return strings.Trim(variable, "\n"), err
	}
	return variable, err
}
