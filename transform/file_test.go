package transform

import (
    "errors"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

const fileTransformInputPath = "input*"

func TestFileTransform_Apply_GlobErrorFailTransform(t *testing.T) {
    variables := NewVariables()
    paths := []string{"input1", "input2", "input2"}
    transform, output := createFileTransformTestObjects(paths, errors.New("glob"), nil)

    err := transform.Apply(variables)
    output.AssertMessage(t, transform.transformMessage())
    assert.Error(t, err)
}

func TestFileTransform_Apply_ReplacerErrorFailTransform(t *testing.T) {
    variables := NewVariables()
    paths := []string{"input1", "input2", "input2"}
    transform, output := createFileTransformTestObjects(paths, nil, errors.New("replacer"))

    err := transform.Apply(variables)
    output.AssertMessage(t, transform.transformMessage())
    assert.Error(t, err)
}

func TestFileTransform_Apply_Success(t *testing.T) {
    variables := NewVariables()
    paths := []string{"input1", "input2", "input2"}
    transform, output := createFileTransformTestObjects(paths, nil, nil)

    err := transform.Apply(variables)
    output.AssertMessage(t, transform.transformMessage())
    assert.NoError(t, err)
}

func createFileTransformOptions() *Options {
    return &Options{
        InputPath:     fileTransformInputPath,
        StringReplace: StringReplace{},
    }
}

func createFileTransformTestObjects(globberPaths []string,
    globberErr error,
    replacerErr error) (*FileTransform, *testmessage.TestPrinter) {
    options := createFileTransformOptions()
    output := testmessage.NewTestPrinter()
    globber := NewMockGlobber(globberPaths, globberErr)
    replacer := NewMockFileStringReplacer(replacerErr)
    transform := NewFileTransform(*options, output, globber, replacer)
    return transform, output
}

type MockFileStringReplacer struct {
    mock.Mock
}

func NewMockFileStringReplacer(err error) *MockFileStringReplacer {
    mockReplacer := MockFileStringReplacer{}
    mockReplacer.On("Replace").Return(err)
    return &mockReplacer
}

func (replacer *MockFileStringReplacer) Replace(filepath string, stringReplace StringReplace, variables Variables) error {
    args := replacer.Called()
    return args.Error(0)
}
