package cmd

import (
    "testing"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/ncipollo/fnew/testaction"
    "github.com/stretchr/testify/assert"
    "errors"
)

func TestCreateCommand_Run_FailsOnCleanup(t *testing.T) {
    actionErr := errors.New("cleanup")
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, nil)
    createAction := testaction.NewMockAction(printer, nil)
    transformAction := testaction.NewMockAction(printer, nil)
    cleanupAction := testaction.NewMockAction(printer, actionErr)
    command := NewCreateCommand(setupAction, createAction, transformAction, cleanupAction)

    err := command.Run(printer)

    assert.Error(t, err)
    printer.AssertMessage(t, command.cleanupErrorMessage(actionErr))
}

func TestCreateCommand_Run_FailsOnCreate(t *testing.T) {
    actionErr := errors.New("create")
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, nil)
    createAction := testaction.NewMockAction(printer, actionErr)
    transformAction := testaction.NewMockAction(printer, nil)
    cleanupAction := testaction.NewMockAction(printer, nil)
    command := NewCreateCommand(setupAction, createAction, transformAction, cleanupAction)

    err := command.Run(printer)

    assert.Error(t, err)
    printer.AssertMessage(t, command.createErrorMessage(actionErr))
}

func TestCreateCommand_Run_FailsOnTransform(t *testing.T) {
    actionErr := errors.New("transform")
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, nil)
    createAction := testaction.NewMockAction(printer, nil)
    transformAction := testaction.NewMockAction(printer, actionErr)
    cleanupAction := testaction.NewMockAction(printer, nil)
    command := NewCreateCommand(setupAction, createAction, transformAction, cleanupAction)

    err := command.Run(printer)

    assert.Error(t, err)
    printer.AssertMessage(t, command.transformErrorMessage(actionErr))
}

func TestCreateCommand_Run_FailsOnSetup(t *testing.T) {
    actionErr := errors.New("setup")
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, actionErr)
    createAction := testaction.NewMockAction(printer, nil)
    transformAction := testaction.NewMockAction(printer, nil)
    cleanupAction := testaction.NewMockAction(printer, nil)
    command := NewCreateCommand(setupAction, createAction, transformAction, cleanupAction)

    err := command.Run(printer)

    assert.Error(t, err)
    printer.AssertMessage(t, command.setupErrorMessage(actionErr))
}

func TestCreateCommand_Run_Success(t *testing.T) {
    printer := testmessage.NewTestPrinter()
    setupAction := testaction.NewMockAction(printer, nil)
    createAction := testaction.NewMockAction(printer, nil)
    transformAction := testaction.NewMockAction(printer, nil)
    cleanupAction := testaction.NewMockAction(printer, nil)
    command := NewCreateCommand(setupAction, createAction, transformAction, cleanupAction)

    err := command.Run(printer)

    assert.NoError(t, err)
    setupAction.AssertExpectations(t)
    createAction.AssertExpectations(t)
    transformAction.AssertExpectations(t)
    cleanupAction.AssertExpectations(t)
    printer.AssertMessage(t, command.completionMessage())
}
