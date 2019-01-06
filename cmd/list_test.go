package cmd

import (
    "testing"
    "github.com/ncipollo/fnew/testaction"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/stretchr/testify/assert"
    "errors"
)

func TestListCommand_Run_FailsOnListError(t *testing.T) {
    actionErr := errors.New("list")
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, nil)
    listAction := testaction.NewMockAction(printer, actionErr)
    command := NewListCommand(setupAction, listAction)

    err := command.Run(printer)

    assert.Error(t, err)
    setupAction.AssertExpectations(t)
    listAction.AssertExpectations(t)
    printer.AssertMessage(t, command.listErrorMessage(actionErr))
}

func TestListCommand_Run_FailsOnSetupError(t *testing.T) {
    actionErr := errors.New("setup")
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, actionErr)
    listAction := testaction.NewMockAction(printer, nil)
    command := NewListCommand(setupAction, listAction)

    err := command.Run(printer)

    assert.Error(t, err)
    setupAction.AssertExpectations(t)
    printer.AssertMessage(t, command.setupErrorMessage(actionErr))
}

func TestListCommand_Run_Success(t *testing.T) {
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, nil)
    listAction := testaction.NewMockAction(printer, nil)
    command := NewListCommand(setupAction, listAction)

    err := command.Run(printer)

    assert.NoError(t, err)
    setupAction.AssertExpectations(t)
    listAction.AssertExpectations(t)
}
