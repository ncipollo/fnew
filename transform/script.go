package transform

import (
    "os/exec"
    "os"
    "fmt"
    "errors"
    "strings"
    "github.com/ncipollo/fnew/message"
    "sort"
)

type ScriptTransform struct {
    Options
    output message.Printer
    runner CommandRunner
}

func NewScriptTransform(options Options,
    output message.Printer,
    runner CommandRunner) *ScriptTransform {
    return &ScriptTransform{Options: options, output: output, runner: runner}
}

func (transform *ScriptTransform) Apply(variables Variables) error {
    transform.output.Println(transform.transformMessage(variables))

    err := transform.validateInputPath(variables)
    if err != nil {
        return err
    }

    cmd := transform.createCommand(variables)
    return transform.runner.RunCommand(*cmd)
}

func (transform *ScriptTransform) transformMessage(variables Variables) string {
    inputPath := transform.inputPath(variables)

    return fmt.Sprintf("Running script at path: (%s)", inputPath)
}

func (transform *ScriptTransform) validateInputPath(variables Variables) error {
    inputPath := transform.inputPath(variables)
    if len(inputPath) == 0 {
        return errors.New("invalid input path")
    }
    return nil
}

func (transform *ScriptTransform) inputPath(variables Variables) string {
    if strings.HasPrefix(transform.Options.InputPath, "$") {
        return variables[strings.TrimPrefix(transform.Options.InputPath, "$")]
    } else {
        return transform.Options.InputPath
    }
}

func (transform *ScriptTransform) createCommand(variables Variables) *exec.Cmd {
    scriptPath := transform.InputPath
    args := append([]string{scriptPath}, transform.Arguments...)
    cmd := exec.Command("sh", args...)
    cmd.Env = transform.environmentFromVariables(variables)
    return cmd
}

func (ScriptTransform) environmentFromVariables(variables Variables) []string {
    environment := os.Environ()
    for key, value := range variables {
        environment = append(environment, fmt.Sprintf("%s=%s", key, value))
    }
    sort.Strings(environment)
    return environment
}

type CommandRunner interface {
    RunCommand(cmd exec.Cmd) error
}

type osCommandRunner struct{}

func NewCommandRunner() CommandRunner {
    return &osCommandRunner{}
}

func (osCommandRunner) RunCommand(cmd exec.Cmd) error {
    return cmd.Run()
}
