package message

import (
    "fmt"
    "io"
    "os"
)

type Printer interface {
    Println(message string)
}

type standardPrinter struct {
    output io.Writer
}

func (writer *standardPrinter) Println(message string) {
    fmt.Fprintln(writer.output, message)
}

func NewStandardWriter() Printer {
    return &standardPrinter{output: os.Stdout}
}
