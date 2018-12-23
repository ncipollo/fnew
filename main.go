package main

import (
    "github.com/ncipollo/fnew/cmd"
    "os"
)

func main() {
    parser := cmd.NewParser(os.Environ())
    parser.Parse()
}
