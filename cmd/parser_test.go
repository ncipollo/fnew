package cmd

import (
    "github.com/stretchr/testify/assert"
    "os"
    "testing"
)

func TestParser_Parse(t *testing.T) {
    // TODO: Actually test commands once implemented
    parser := setupParser("project", "/")

    cmd := parser.Parse()
    assert.Nil(t, cmd)
}

func setupParser(args ...string) *Parser {
    os.Args = append([]string{"fnew"}, args...)
    return NewParser([]string{})
}
