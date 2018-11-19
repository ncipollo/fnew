package transform

import (
    "bufio"
    "errors"
    "fmt"
    "github.com/ncipollo/fnew/message"
    "io"
    "strings"
)

type InputTransform struct {
    Options
    input  io.Reader
    output message.Printer
}

func NewInputTransform(options Options, input io.Reader, output message.Printer) *InputTransform {
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
            msg := transform.skipInputMessage(variableName)
            transform.output.Println(msg)
            return nil
        }
    }

    msg := transform.requestInputMessage(variableName)
    transform.output.Println(msg)

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
    if err == nil {
        return strings.Trim(variable, "\n"), err
    }
    return variable, err
}
