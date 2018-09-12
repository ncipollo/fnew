package main

import (
	"github.com/ncipollo/fnew/workspace"
	"fmt"
)

func main() {
	appWorkspace := workspace.New(workspace.Directory(), workspace.OSDirectoryCreator())
	err := appWorkspace.Setup()
	fmt.Printf("Error: %v", err)
}