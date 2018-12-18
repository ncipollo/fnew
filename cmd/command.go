package cmd

import "github.com/ncipollo/fnew/message"

type Command interface {
    Run(printer message.Printer)
}