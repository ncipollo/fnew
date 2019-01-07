package cmd

import (
    "github.com/ncipollo/fnew/action"
    "github.com/ncipollo/fnew/message"
    "fmt"
)

type CreateCommand struct {
    setupAction     action.Action
    createAction    action.Action
    transformAction action.Action
    cleanupAction   action.Action
}

func NewCreateCommand(
    setupAction action.Action,
    createAction action.Action,
    transformAction action.Action,
    cleanupAction action.Action) *CreateCommand {
    return &CreateCommand{
        setupAction:     setupAction,
        createAction:    createAction,
        transformAction: transformAction,
        cleanupAction:   cleanupAction,
    }
}

func (command *CreateCommand) Run(printer message.Printer) error {
    err := command.setupAction.Perform(printer)
    if err != nil {
        printer.Println(command.setupErrorMessage(err))
        return err
    }

    err = command.createAction.Perform(printer)
    if err != nil {
        printer.Println(command.createErrorMessage(err))
        return err
    }

    err = command.transformAction.Perform(printer)
    if err != nil {
        printer.Println(command.transformErrorMessage(err))
        return err
    }

    err = command.cleanupAction.Perform(printer)
    if err != nil {
        printer.Println(command.cleanupErrorMessage(err))
        return err
    }

    printer.Println(command.completionMessage())
    return nil
}

func (CreateCommand) setupErrorMessage(err error) string {
    return fmt.Sprintf("Setup failed. Error: %v", err)
}

func (CreateCommand) createErrorMessage(err error) string {
    return fmt.Sprintf("Project creation failed. Error: %v", err)
}

func (CreateCommand) transformErrorMessage(err error) string {
    return fmt.Sprintf("Transform failed. Error: %v", err)
}

func (CreateCommand) cleanupErrorMessage(err error) string {
    return fmt.Sprintf("Transform failed. Error: %v", err)
}

func (CreateCommand) completionMessage() string {
    return fmt.Sprintf("Project created!")
}
