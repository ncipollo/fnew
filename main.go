package main

import (
	"github.com/ncipollo/fnew/workspace"
	"fmt"
)

func main() {
	appWorkspace := workspace.New(workspace.Directory(), workspace.OSDirectoryCreator())
	appWorkspace.Setup()
	fmt.Printf("Created ~/.fnew")
}