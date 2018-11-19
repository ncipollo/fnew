package transform

import (
    "github.com/stretchr/testify/assert"
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