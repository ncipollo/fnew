package cmd

import (
    "testing"
    "github.com/ncipollo/fnew/testmessage"
    "github.com/ncipollo/fnew/testaction"
    "github.com/stretchr/testify/assert"
    "errors"
    "github.com/stretchr/testify/mock"
)

const createCommandPath = "test/"

func TestCreateCommand_Run_FailsOnCleanup(t *testing.T) {
    actionErr := errors.New("cleanup")
    printer := testmessage.NewTestPrinter()
    changer := NewMockDirectoryChanger(nil, nil)
    setupAction := testaction.NewMockAction(printer, nil)
    createAction := testaction.NewMockAction(printer, nil)
    transformAction := testaction.NewMockAction(printer, nil)
    cleanupAction := testaction.NewMockAction(printer, actionErr)
    command := NewCreateCommand(createCommandPath,
        changer,
        setupAction,
        createAction,
        transformAction,
        cleanupAction)

    err := command.Run(printer)

    assert.Error(t, err)
    printer.AssertMessage(t, command.cleanupErrorMessage(actionErr))
}

func TestCreateCommand_Run_FailsOnCreate(t *testing.T) {
    actionErr := errors.New("create")
    printer := testmessage.NewTestPrinter()
    changer := NewMockDirectoryChanger(nil, nil)
    setupAction := testaction.NewMockAction(printer, nil)
    createAction := testaction.NewMockAction(printer, actionErr)
    transformAction := testaction.NewMockAction(printer, nil)
    cleanupAction := testaction.NewMockAction(printer, nil)
    command := NewCreateCommand(createCommandPath,
        changer,
        setupAction,
        createAction,
        transformAction,
        cleanupAction)

    err := command.Run(printer)

    assert.Error(t, err)
    printer.AssertMessage(t, command.createErrorMessage(actionErr))
}

func TestCreateCommand_Run_FailsOnMoveIntoProject(t *testing.T) {
    actionErr := errors.New("moveIn")
    printer := testmessage.NewTestPrinter()
    changer := NewMockDirectoryChanger(actionErr, nil)
    setupAction := testaction.NewMockAction(printer, nil)
    createAction := testaction.NewMockAction(printer, nil)
    transformAction := testaction.NewMockAction(printer, nil)
    cleanupAction := testaction.NewMockAction(printer, nil)
    command := NewCreateCommand(createCommandPath,
        changer,
        setupAction,
        createAction,
        transformAction,
        cleanupAction)

    err := command.Run(printer)

    assert.Error(t, err)
    printer.AssertMessage(t, command.moveIntoProjectErrorMessage(actionErr))
}

func TestCreateCommand_Run_FailsOnMoveOutOfProject(t *testing.T) {
    actionErr := errors.New("moveOut")
    printer := testmessage.NewTestPrinter()
    changer := NewMockDirectoryChanger(nil, actionErr)
    setupAction := testaction.NewMockAction(printer, nil)
    createAction := testaction.NewMockAction(printer, nil)
    transformAction := testaction.NewMockAction(printer, nil)
    cleanupAction := testaction.NewMockAction(printer, nil)
    command := NewCreateCommand(createCommandPath,
        changer,
        setupAction,
        createAction,
        transformAction,
        cleanupAction)

    err := command.Run(printer)

    assert.Error(t, err)
    printer.AssertMessage(t, command.moveOutOfProjectErrorMessage(actionErr))
}

func TestCreateCommand_Run_FailsOnTransform(t *testing.T) {
    actionErr := errors.New("transform")
    printer := testmessage.NewTestPrinter()
    changer := NewMockDirectoryChanger(nil, nil)
    setupAction := testaction.NewMockAction(printer, nil)
    createAction := testaction.NewMockAction(printer, nil)
    transformAction := testaction.NewMockAction(printer, actionErr)
    cleanupAction := testaction.NewMockAction(printer, nil)
    command := NewCreateCommand(createCommandPath,
        changer,
        setupAction,
        createAction,
        transformAction,
        cleanupAction)

    err := command.Run(printer)

    assert.Error(t, err)
    printer.AssertMessage(t, command.transformErrorMessage(actionErr))
}

func TestCreateCommand_Run_FailsOnSetup(t *testing.T) {
    actionErr := errors.New("setup")
    printer := testmessage.NewTestPrinter()
    changer := NewMockDirectoryChanger(nil, nil)
    setupAction := testaction.NewMockAction(printer, actionErr)
    createAction := testaction.NewMockAction(printer, nil)
    transformAction := testaction.NewMockAction(printer, nil)
    cleanupAction := testaction.NewMockAction(printer, nil)
    command := NewCreateCommand(createCommandPath,
        changer,
        setupAction,
        createAction,
        transformAction,
        cleanupAction)

    err := command.Run(printer)

    assert.Error(t, err)
    printer.AssertMessage(t, command.setupErrorMessage(actionErr))
}

func TestCreateCommand_Run_Success(t *testing.T) {
    printer := testmessage.NewTestPrinter()
    changer := NewMockDirectoryChanger(nil, nil)
    setupAction := testaction.NewMockAction(printer, nil)
    createAction := testaction.NewMockAction(printer, nil)
    transformAction := testaction.NewMockAction(printer, nil)
    cleanupAction := testaction.NewMockAction(printer, nil)
    command := NewCreateCommand(createCommandPath,
        changer,
        setupAction,
        createAction,
        transformAction,
        cleanupAction)

    err := command.Run(printer)

    assert.NoError(t, err)
    setupAction.AssertExpectations(t)
    createAction.AssertExpectations(t)
    transformAction.AssertExpectations(t)
    cleanupAction.AssertExpectations(t)
    printer.AssertMessage(t, command.completionMessage())
}

type MockDirectoryChanger struct {
    mock.Mock
}

func NewMockDirectoryChanger(inErr error, outErr error) *MockDirectoryChanger {
    changer := MockDirectoryChanger{}
    changer.On("Chdir", createCommandPath).Return(inErr)
    changer.On("Chdir", "..").Return(outErr)

    return &changer
}

func (changer *MockDirectoryChanger) Chdir(dir string) error {
    args := changer.Called(dir)
    return args.Error(0)
}
