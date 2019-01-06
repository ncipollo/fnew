package testaction

import (
    "github.com/stretchr/testify/mock"
    "github.com/ncipollo/fnew/message"
)

type MockAction struct {
    mock.Mock
}

func NewMockAction(output message.Printer, err error) *MockAction {
    mockAction := MockAction{}
    mockAction.On("Perform", output).Return(err)
    return &mockAction
}

func (mockAction *MockAction) Perform(output message.Printer) error {
    args := mockAction.Called(output)
    return args.Error(0)
}




