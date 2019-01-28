package cmd

import (
    "github.com/ncipollo/fnew/testmessage"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestVersionCommand_Run(t *testing.T) {
    version := "1.0"
    output := testmessage.NewTestPrinter()
    cmd := NewVersionCommand(version)

    err := cmd.Run(output)

    output.AssertMessage(t, cmd.versionMessage())
    assert.NoError(t, err)
}

