package main

import (
    "fmt"
    "github.com/ncipollo/fnew/manifest"
    "github.com/ncipollo/fnew/repo"
    "github.com/ncipollo/fnew/transform"
    "github.com/ncipollo/fnew/workspace"
    "os"
)

func main() {
    appWorkspace := workspace.New(workspace.Directory(), workspace.OSDirectoryChecker(), workspace.OSDirectoryCreator())
    err := appWorkspace.Setup()
    fmt.Printf("Workspace Error: %v\n", err)

    manifestRepo := repo.New()
    _, err = manifestRepo.Clone(appWorkspace.ManifestsPath(), manifest.DefaultRepository)

    options := transform.Options{OutputVariable: "$TEST"}
    variables := transform.NewVariables()
    inputTransform := transform.NewInputTransform(options, os.Stdin, os.Stdout)
    inputTransform.Apply(variables)
}
