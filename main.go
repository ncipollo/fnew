package main

import (
    "github.com/ncipollo/fnew/cmd"
    "os"
    "github.com/ncipollo/fnew/message"
)

func main() {
    parser := cmd.NewParser(os.Environ())
    command := parser.Parse()
    if command != nil {
        printer := message.NewStandardWriter()
        err := command.Run(printer)
        if err == nil {
            os.Exit(0)
        } else {
            os.Exit(1)
        }
    }
}
