package cmd

import (
    "fmt"
    "github.com/ncipollo/fnew/message"
)

type VersionCommand struct {
    version string
}

func NewVersionCommand(version string) *VersionCommand {
    return &VersionCommand{version:version}
}

func (command *VersionCommand) Run(printer message.Printer) error {
    printer.Println(command.versionMessage())
    return nil
}

func (command *VersionCommand) versionMessage() string {
    return fmt.Sprintf("fnew version: %v", command.version)
}
