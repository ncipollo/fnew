package transform

import "github.com/stretchr/testify/mock"

type MockTransform struct {
    mock.Mock
}

func NewMockTransform(variables Variables, err error) *MockTransform {
    mockTransform := MockTransform{}
    mockTransform.On("Apply", variables).Return(err)
    return &mockTransform
}

func (mockTransform *MockTransform) Apply(variables Variables) error {
    args := mockTransform.Called(variables)
    return args.Error(0)
}
