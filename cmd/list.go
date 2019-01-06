package cmd

import (
    "github.com/ncipollo/fnew/action"
    "github.com/ncipollo/fnew/message"
    "fmt"
)

type ListCommand struct {
    setupAction action.Action
    listAction  action.Action
}

func NewListCommand(setupAction action.Action, listAction action.Action) *ListCommand {
    return &ListCommand{setupAction: setupAction, listAction: listAction}
}

func (command *ListCommand) Run(printer message.Printer) error {
    err := command.setupAction.Perform(printer)
    if err != nil {
        printer.Println(command.setupErrorMessage(err))
        return err
    }

    err = command.listAction.Perform(printer)
    if err != nil {
        printer.Println(command.listErrorMessage(err))
        return err
    }

    return nil
}

func (command *ListCommand) setupErrorMessage(err error) string {
    return fmt.Sprintf("Setup failed for list command. Error: %v", err)
}

func (command *ListCommand) listErrorMessage(err error) string {
    return fmt.Sprintf("Unable to list projects. Error: %v", err)
}
