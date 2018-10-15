package project

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ncipollo/fnew/transform"
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
		Transforms: []transform.TransformOptions{
			{
				Arguments:      []string{"foo", "bar"},
				InputPath:      "input/path",
				InputVariable:  "$inputVariable",
				OutputPath:     "output/path",
				OutputVariable: "$outputVariable",
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
		Transforms: []transform.TransformOptions{
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

	expectedProject := &Project{Transforms: []transform.TransformOptions{}}
	assert.Equal(t, expectedProject, parsedProject)
	assert.NoError(t, err)
}

func TestFromJSON_InvalidJson(t *testing.T) {
	parsedProject, err := FromString(invalidJSON)

	assert.Nil(t, parsedProject)
	assert.Error(t, err)
}
