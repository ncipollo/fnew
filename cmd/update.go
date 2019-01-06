package cmd

import (
    "github.com/ncipollo/fnew/action"
    "github.com/ncipollo/fnew/message"
    "fmt"
)

type UpdateCommand struct {
    setupAction action.Action
    updateAction  action.Action
}

func NewUpdateCommand(setupAction action.Action, updateAction action.Action) *UpdateCommand {
    return &UpdateCommand{setupAction: setupAction, updateAction: updateAction}
}

func (command *UpdateCommand) Run(printer message.Printer) error {
    err := command.setupAction.Perform(printer)
    if err != nil {
        printer.Println(command.setupErrorMessage(err))
        return err
    }

    err = command.updateAction.Perform(printer)
    if err != nil {
        printer.Println(command.updateErrorMessage(err))
        return err
    }

    printer.Println(command.updatedMessage())
    return nil
}


func (UpdateCommand) setupErrorMessage(err error) string {
    return fmt.Sprintf("Setup failed for update command. Error: %v", err)
}

func (UpdateCommand) updateErrorMessage(err error) string {
    return fmt.Sprintf("Unable to update manifests. Error: %v", err)
}

func (UpdateCommand) updatedMessage() string {
    return "Up to date!"
}

