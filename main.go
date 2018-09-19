package main

import (
	"github.com/ncipollo/fnew/workspace"
	"fmt"
	"github.com/ncipollo/fnew/repo"
	"github.com/ncipollo/fnew/manifest"
)

func main() {
	appWorkspace := workspace.New(workspace.Directory(), workspace.OSDirectoryChecker(), workspace.OSDirectoryCreator())
	err := appWorkspace.Setup()
	fmt.Printf("Workspace Error: %v", err)

	manifestRepo := repo.New()
	_, err = manifestRepo.Clone(appWorkspace.ManifestsPath(), manifest.DefaultRepository)

}
