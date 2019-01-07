package cmd

import (
    "github.com/ncipollo/fnew/action"
    "github.com/ncipollo/fnew/message"
    "fmt"
    "os"
)

type CreateCommand struct {
    localPath        string
    directoryChanger DirectoryChanger
    setupAction      action.Action
    createAction     action.Action
    transformAction  action.Action
    cleanupAction    action.Action
}

func NewCreateCommand(
    localPath string,
    directoryChanger DirectoryChanger,
    setupAction action.Action,
    createAction action.Action,
    transformAction action.Action,
    cleanupAction action.Action) *CreateCommand {
    return &CreateCommand{
        localPath:        localPath,
        directoryChanger: directoryChanger,
        setupAction:      setupAction,
        createAction:     createAction,
        transformAction:  transformAction,
        cleanupAction:    cleanupAction,
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

    err = command.moveIntoProject()
    if err != nil {
        printer.Println(command.moveIntoProjectErrorMessage(err))
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

    err = command.moveOutOfProject()
    if err != nil {
        printer.Println(command.moveOutOfProjectErrorMessage(err))
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

func (CreateCommand) moveIntoProjectErrorMessage(err error) string {
    return fmt.Sprintf("Failed to move into project. Error: %v", err)
}

func (CreateCommand) moveOutOfProjectErrorMessage(err error) string {
    return fmt.Sprintf("Failed to move into project. Error: %v", err)
}

func (CreateCommand) completionMessage() string {
    return fmt.Sprintf("Project created!")
}

func (command *CreateCommand) moveIntoProject() error {
    return command.directoryChanger.Chdir(command.localPath)
}

func (command *CreateCommand) moveOutOfProject() error {
    return command.directoryChanger.Chdir("..")
}

type DirectoryChanger interface {
    Chdir(dir string) error
}

type OSDirectoryChanger struct {
}

func (OSDirectoryChanger) Chdir(dir string) error {
    return os.Chdir(dir)
}
