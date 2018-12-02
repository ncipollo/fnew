package project

import (
    "github.com/ncipollo/fnew/transform"
    "github.com/stretchr/testify/assert"
    "testing"
)

const (
    allTransformOptions = `{
    "transforms": [
        {
            "arguments" : ["foo","bar"],
            "input_path" : "input/path",
            "input_variable" : "$inputVariable",
            "output_path" : "output/path",
            "output_variable" : "$outputVariable",
            "skip_if_variable_exists" : true,
            "string_replace" : {
                "old" : "old",
                "new" : "new"
            },
        "type": "file_move"
        }
    ] 
}`
    defaultTransformOptions = `{
    "transforms": [
        {
            "type": "file_move"
        }
    ] 
}`
    emptyProject    = `{}`
    emptyTransforms = `{
    "transforms": [] 
}`
    invalidJSON = `{blarg`
)

func TestFromJSON_AllTransformOptions(t *testing.T) {
    parsedProject, err := FromString(allTransformOptions)

    expectedProject := &Project{
        Transforms: []transform.Options{
            {
                Arguments:            []string{"foo", "bar"},
                InputPath:            "input/path",
                InputVariable:        "$inputVariable",
                OutputPath:           "output/path",
                OutputVariable:       "$outputVariable",
                SkipIfVariableExists: true,
                StringReplace: transform.StringReplace{
                    Old: "old",
                    New: "new",
                },
                Type: transform.TypeFileMove,
            },
        },
    }
    assert.Equal(t, expectedProject, parsedProject)
    assert.NoError(t, err)
}

func TestFromJSON_DefaultTransformOptions(t *testing.T) {
    parsedProject, err := FromString(defaultTransformOptions)

    expectedProject := &Project{
        Transforms: []transform.Options{
            {
                Type: transform.TypeFileMove,
            },
        },
    }
    assert.Equal(t, expectedProject, parsedProject)
    assert.NoError(t, err)
}

func TestFromJSON_EmptyProject(t *testing.T) {
    parsedProject, err := FromString(emptyProject)

    expectedProject := &Project{Transforms: nil}
    assert.Equal(t, expectedProject, parsedProject)
    assert.NoError(t, err)
}

func TestFromJSON_EmptyTransforms(t *testing.T) {
    parsedProject, err := FromString(emptyTransforms)

    expectedProject := &Project{Transforms: []transform.Options{}}
    assert.Equal(t, expectedProject, parsedProject)
    assert.NoError(t, err)
}

func TestFromJSON_InvalidJson(t *testing.T) {
    parsedProject, err := FromString(invalidJSON)

    assert.Nil(t, parsedProject)
    assert.Error(t, err)
}

func TestPathInRepo(t *testing.T) {
    assert.Equal(t, "repo/.fnew", PathInRepo("repo"))
}
