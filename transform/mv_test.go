package transform

import (
    "github.com/stretchr/testify/mock"
    "github.com/ncipollo/fnew/message"
    "testing"
    "github.com/stretchr/testify/assert"
    "errors"
)

func TestFileMoveTransform_Apply_Success(t *testing.T) {
    variables := NewVariables()
    inputPath := "input"
    outputPath := "output"
    options := createFileMoveTransformOptions(inputPath, outputPath)
    transform, output := createFileMoveTransformTestObjects(*options, nil)

    err := transform.Apply(variables)

    output.AssertMessage(t, transform.transformMessage(variables))
    assert.NoError(t, err)
}

func TestFileMoveTransform_Apply_SuccessWithVariables(t *testing.T) {
    variables := NewVariables()
    variables["input"] = "input"
    variables["output"] = "output"
    inputPath := "$input"
    outputPath := "$output"
    options := createFileMoveTransformOptions(inputPath, outputPath)
    transform, output := createFileMoveTransformTestObjects(*options, nil)

    err := transform.Apply(variables)

    output.AssertMessage(t, transform.transformMessage(variables))
    assert.NoError(t, err)
}

func TestFileMoveTransform_Apply_EmptyInputPathFailsTransform(t *testing.T) {
    variables := NewVariables()
    inputPath := ""
    outputPath := "output"
    options := createFileMoveTransformOptions(inputPath, outputPath)
    transform, output := createFileMoveTransformTestObjects(*options, nil)

    err := transform.Apply(variables)

    output.AssertMessage(t, transform.transformMessage(variables))
    assert.Error(t, err)
}

func TestFileMoveTransform_Apply_EmptyInputPathVariableFailsTransform(t *testing.T) {
    variables := NewVariables()
    inputPath := "$input"
    outputPath := "output"
    options := createFileMoveTransformOptions(inputPath, outputPath)
    transform, output := createFileMoveTransformTestObjects(*options, nil)

    err := transform.Apply(variables)

    output.AssertMessage(t, transform.transformMessage(variables))
    assert.Error(t, err)
}

func TestFileMoveTransform_Apply_EmptyOutputPathVariableFailsTransform(t *testing.T) {
    variables := NewVariables()
    inputPath := "input"
    outputPath := "$output"
    options := createFileMoveTransformOptions(inputPath, outputPath)
    transform, output := createFileMoveTransformTestObjects(*options, nil)

    err := transform.Apply(variables)

    output.AssertMessage(t, transform.transformMessage(variables))
    assert.Error(t, err)
}

func TestFileMoveTransform_Apply_EmptyOutputPathFailsTransform(t *testing.T) {
    variables := NewVariables()
    inputPath := "input"
    outputPath := ""
    options := createFileMoveTransformOptions(inputPath, outputPath)
    transform, output := createFileMoveTransformTestObjects(*options, nil)

    err := transform.Apply(variables)

    output.AssertMessage(t, transform.transformMessage(variables))
    assert.Error(t, err)
}

func TestFileMoveTransform_Apply_moverErrorFailsTransform(t *testing.T) {
    variables := NewVariables()
    inputPath := "input"
    outputPath := "output"
    options := createFileMoveTransformOptions(inputPath, outputPath)
    transform, output := createFileMoveTransformTestObjects(*options, errors.New("mover"))

    err := transform.Apply(variables)

    output.AssertMessage(t, transform.transformMessage(variables))
    assert.Error(t, err)
}

func createFileMoveTransformOptions(inputPath string, outputPath string) *Options {
    return &Options{
        InputPath:  inputPath,
        OutputPath: outputPath,
    }
}

func createFileMoveTransformTestObjects(options Options, moverErr error) (*FileMoveTransform, *message.TestPrinter,) {
    output := message.NewTestPrinter()
    mover := NewMockFileMover(moverErr)
    transform := NewFileMoveTransform(options, output, mover)
    return transform, output
}

type MockFileMover struct {
    mock.Mock
}

func NewMockFileMover(err error) *MockFileMover {
    mockMover := MockFileMover{}
    mockMover.On("Move").Return(err)
    return &mockMover
}

func (mover *MockFileMover) Move(oldPath string, newPath string) error {
    args := mover.Called()
    return args.Error(0)
}
