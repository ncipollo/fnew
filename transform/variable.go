package transform

import (
    "fmt"
    "io"
)

type VariableTransform struct {
    Options
    output io.Writer
}

func NewVariableTransform(options Options, output io.Writer) *VariableTransform {
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

    fmt.Fprintln(transform.output, transform.transformMessage())
    if skip {
        fmt.Fprintln(transform.output, transform.skipMessage())
        return nil
    }

    variables[outputKey] = stringReplace.Replace(input)

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
