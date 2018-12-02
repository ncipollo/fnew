package transform

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "errors"
)

func TestTransformer_Apply_FailsOnTransformError(t *testing.T) {
    transformer := NewTransformer()
    variables := NewVariables()
    transforms := []Transform{
        NewMockTransform(variables, errors.New("transform")),
        NewMockTransform(variables, nil),
    }

    err := transformer.Apply(transforms, variables)

    assert.Error(t, err)
}

func TestTransformer_Apply_NoErrorWhenNoTransforms(t *testing.T) {
    transformer := NewTransformer()
    variables := NewVariables()

    err := transformer.Apply(nil, variables)

    assert.NoError(t, err)
}

func TestTransformer_Apply_NoErrorWhenEmpty(t *testing.T) {
    transformer := NewTransformer()
    variables := NewVariables()
    var transforms []Transform

    err := transformer.Apply(transforms, variables)

    assert.NoError(t, err)
}

func TestTransformer_Apply_Success(t *testing.T) {
    transformer := NewTransformer()
    variables := NewVariables()
    transforms := []Transform{
        NewMockTransform(variables, nil),
        NewMockTransform(variables, nil),
    }

    err := transformer.Apply(transforms, variables)

    assert.NoError(t, err)
    for _, currentTransform := range transforms {
        currentTransform.(*MockTransform).AssertExpectations(t)
    }
}
