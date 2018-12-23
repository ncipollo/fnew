package transform

import (
    "errors"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "os"
    "os/exec"
    "sort"
    "testing"
)

const scriptTransformInputPath = "script.sh"

func TestScriptTransform_Apply_Success(t *testing.T) {
    args := []string{"arg1", "arg2"}
    variables := NewVariables()
    variables["key1"] = "value1"
    variables["key2"] = "value2"
    environment := []string{"key1=value1", "key2=value2"}
    cmd := expectedCommand(args, environment)
    options := Options{InputPath: scriptTransformInputPath, Arguments: args}
    transform, runner, output := createScriptTransformTestObjects(options, cmd, nil)

    err := transform.Apply(variables)

    runner.AssertExpectations(t)
    output.AssertMessage(t, transform.transformMessage(variables))
    assert.NoError(t, err)
}

func TestScriptTransform_Apply_ScriptErrorFailsTransform(t *testing.T) {
    args := []string{"arg1", "arg2"}
    variables := NewVariables()
    var environment []string
    cmd := expectedCommand(args, environment)
    options := Options{InputPath: scriptTransformInputPath, Arguments: args}
    transform, runner, output := createScriptTransformTestObjects(options, cmd, errors.New("script"))

    err := transform.Apply(variables)

    runner.AssertExpectations(t)
    output.AssertMessage(t, transform.transformMessage(variables))
    assert.Error(t, err)
}

func TestScriptTransform_Apply_EmptyInputPathFailsTransform(t *testing.T) {
    variables := NewVariables()
    cmd := expectedCommand([]string{}, []string{})
    options := Options{}
    transform, _, output := createScriptTransformTestObjects(options, cmd, nil)

    err := transform.Apply(variables)

    output.AssertMessage(t, transform.transformMessage(variables))
    assert.Error(t, err)
}

func TestScriptTransform_Apply_EmptyInputPathVariableFailsTransform(t *testing.T) {
    variables := NewVariables()
    cmd := expectedCommand([]string{}, []string{})
    options := Options{InputPath: "$input"}
    transform, _, output := createScriptTransformTestObjects(options, cmd, nil)

    err := transform.Apply(variables)

    output.AssertMessage(t, transform.transformMessage(variables))
    assert.Error(t, err)
}

func expectedCommand(args []string, environment []string) exec.Cmd {
    argsWithScript := append([]string{scriptTransformInputPath}, args...)
    env := append(os.Environ(), environment...)
    sort.Strings(env)
    cmd := exec.Command("sh", argsWithScript...)
    cmd.Env = env
    return *cmd
}

func createScriptTransformTestObjects(options Options,
    cmd exec.Cmd,
    scriptErr error) (*ScriptTransform, *MockCommandRunner, *testmessage.TestPrinter) {
    output := testmessage.NewTestPrinter()
    runner := NewMockCommandRunner(cmd, scriptErr)
    transform := NewScriptTransform(options, output, runner)
    return transform, runner, output
}

type MockCommandRunner struct {
    mock.Mock
}

func NewMockCommandRunner(cmd exec.Cmd, err error) *MockCommandRunner {
    runner := MockCommandRunner{}
    runner.On("RunCommand", cmd).Return(err)
    return &runner
}

func (runner *MockCommandRunner) RunCommand(cmd exec.Cmd) error {
    args := runner.Called(cmd)
    return args.Error(0)
}
