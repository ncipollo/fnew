package transform

import (
    "fmt"
    "github.com/ncipollo/fnew/message"
)

type VariableTransform struct {
    Options
    output message.Printer
}

func NewVariableTransform(options Options, output message.Printer) *VariableTransform {
    return &VariableTransform{Options: options, output: output}
}

func (transform *VariableTransform) Apply(variables Variables) error {
    inputKey := transform.Options.InputVariable
    input := variables[inputKey]
    outputKey := transform.Options.OutputVariable
    output := variables[outputKey]
    skipIfOutputExists := transform.Options.SkipIfVariableExists
    stringReplace := transform.Options.StringReplace
    skip := len(output) != 0 && skipIfOutputExists

    transform.output.Println(transform.transformMessage())
    if skip {
        transform.output.Println(transform.skipMessage())
        return nil
    }

    variables[outputKey] = stringReplace.Replace(input, variables)

    return nil
}

func (transform *VariableTransform) skipMessage() string {
    outputKey := transform.Options.OutputVariable

    return fmt.Sprintf("Output variable (%s) already exists, skipping", outputKey)
}

func (transform *VariableTransform) transformMessage() string {
    inputKey := transform.Options.InputVariable
    outputKey := transform.Options.OutputVariable

    return fmt.Sprintf("Transforming input variable (%s), saving as (%s)", inputKey, outputKey)
}
