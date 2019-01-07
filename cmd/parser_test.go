package cmd

import (
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestParser_Parse_CreateCommand(t *testing.T) {
    parser := setupParser("project", "/")

    cmd := parser.Parse()
    assert.IsType(t, &CreateCommand{}, cmd)
}

func TestParser_Parse_ListCommand(t *testing.T) {
    parser := setupParser("--list")

    cmd := parser.Parse()
    assert.IsType(t, &ListCommand{}, cmd)
}

func TestParser_Parse_ListCommand_Shorthand(t *testing.T) {
    parser := setupParser("-l")

    cmd := parser.Parse()
    assert.IsType(t, &ListCommand{}, cmd)
}

func TestParser_Parse_UpdateCommand(t *testing.T) {
    parser := setupParser("--update")

    cmd := parser.Parse()
    assert.IsType(t, &UpdateCommand{}, cmd)
}

func TestParser_Parse_UpdateCommand_Shorthand(t *testing.T) {
    parser := setupParser("-u")

    cmd := parser.Parse()
    assert.IsType(t, &UpdateCommand{}, cmd)
}

func setupParser(args ...string) *Parser {
    os.Args = append([]string{"fnew"}, args...)
    return NewParser([]string{})
}
