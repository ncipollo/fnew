package transform

import (
    "github.com/stretchr/testify/assert"
    "os"
    "strings"
    "testing"
)

func TestStringReplace_Replace_InputEmpty(t *testing.T) {
    stringReplace := StringReplace{Old: "foo", New: "bar"}
    assert.Equal(t, "", stringReplace.Replace("", NewVariables()))
}

func TestStringReplace_Replace_InputNotEmpty(t *testing.T) {
    stringReplace := StringReplace{Old: "foo", New: "bar"}
    assert.Equal(t, "should be bar", stringReplace.Replace("should be foo", NewVariables()))
}

func TestStringReplace_Replace_NewVariable(t *testing.T) {
    variables := NewVariables()
    variables["VAR"] = "bar"
    stringReplace := StringReplace{Old: "foo", New: "$VAR"}
    assert.Equal(t, "should be bar", stringReplace.Replace("should be foo", variables))
}

func TestStringReplace_Replace_OldVariable(t *testing.T) {
    variables := NewVariables()
    variables["VAR"] = "foo"
    stringReplace := StringReplace{Old: "$VAR", New: "bar"}
    assert.Equal(t, "should be bar", stringReplace.Replace("should be foo", variables))
}

func TestVariables_AddEnv(t *testing.T) {
    variables := NewVariables()
    variables.AddEnv()
    for _, env := range os.Environ() {
        parts := strings.Split(env, "=")
        key := parts[0]
        value := parts[1]
        assert.Equal(t,value, variables[key])
    }
}

func TestVariables_AddProjectName(t *testing.T) {
    projectName := "project"
    variables := NewVariables()
    variables.AddProjectName(projectName)
    assert.Equal(t,projectName, variables["project_name"])
}
