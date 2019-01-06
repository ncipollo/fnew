package cmd

import (
    "testing"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/ncipollo/fnew/testaction"
    "github.com/stretchr/testify/assert"
    "errors"
)

func TestUpdateCommand_Run_Success(t *testing.T) {
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, nil)
    updateAction := testaction.NewMockAction(printer, nil)
    command := NewUpdateCommand(setupAction, updateAction)

    err := command.Run(printer)

    assert.NoError(t, err)
    setupAction.AssertExpectations(t)
    updateAction.AssertExpectations(t)
    printer.AssertMessage(t, command.updatedMessage())
}

func TestUpdateCommand_Run_FailsOnSetupError(t *testing.T) {
    actionErr := errors.New("setup")
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, actionErr)
    updateAction := testaction.NewMockAction(printer, nil)
    command := NewUpdateCommand(setupAction, updateAction)

    err := command.Run(printer)

    assert.Error(t, err)
    setupAction.AssertExpectations(t)
    printer.AssertMessage(t, command.setupErrorMessage(actionErr))
}

func TestUpdateCommand_Run_FailsOnUpdateError(t *testing.T) {
    actionErr := errors.New("update")
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, nil)
    updateAction := testaction.NewMockAction(printer, actionErr)
    command := NewUpdateCommand(setupAction, updateAction)

    err := command.Run(printer)

    assert.Error(t, err)
    setupAction.AssertExpectations(t)
    updateAction.AssertExpectations(t)
    printer.AssertMessage(t, command.updateErrorMessage(actionErr))
}