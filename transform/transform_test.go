package transform

import (
    "bytes"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestStringReplace_Replace_InputEmpty(t *testing.T) {
    stringReplace := StringReplace{Old: "foo", New: "bar"}
    assert.Equal(t, "", stringReplace.Replace(""))
}

func TestStringReplace_Replace_InputNotEmpty(t *testing.T) {
    stringReplace := StringReplace{Old: "foo", New: "bar"}
    assert.Equal(t, "should be bar", stringReplace.Replace("should be foo"))
}

func assertTransformMessage(t *testing.T, expectedMessage string, outputBuffer *bytes.Buffer) {
    outputMessage, _ := outputBuffer.ReadString('\n')
    assert.Equal(t, fmt.Sprintln(expectedMessage), outputMessage)
}
