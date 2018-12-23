package testmessage

import (
    "bytes"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

type TestPrinter struct {
    output *bytes.Buffer
}

func NewTestPrinter() *TestPrinter {
    return &TestPrinter{output: new(bytes.Buffer)}
}

func (writer *TestPrinter) Println(message string) {
    fmt.Fprintln(writer.output, message)
}

func (writer *TestPrinter) AssertMessage(t *testing.T, expectedMessage string)  {
    outputMessage, _ := writer.output.ReadString('\n')
    assert.Equal(t, fmt.Sprintln(expectedMessage), outputMessage)
}
