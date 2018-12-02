package transform

import "github.com/stretchr/testify/mock"

type MockTransformer struct {
    mock.Mock
}

func NewMockTransformer(transforms []Transform, variables Variables, err error) *MockTransformer {
    transformer := MockTransformer{}
    transformer.On("Apply", transforms, variables).Return(err)
    return &transformer
}

func (transformer *MockTransformer) Apply(transforms []Transform, variables Variables) error {
    args := transformer.Called(transforms, variables)
    return args.Error(0)
}
