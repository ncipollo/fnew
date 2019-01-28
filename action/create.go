package action

import (
    "errors"
    "fmt"
    "github.com/ncipollo/fnew/merger"
    "github.com/ncipollo/fnew/repo"
    "github.com/ncipollo/fnew/workspace"
    "github.com/ncipollo/fnew/message"
)

type CreateAction struct {
    checker     workspace.DirectoryChecker
    localPath   string
    merger      merger.Merger
    projectName string
    repo        repo.Repo
}

func NewCreateAction(checker workspace.DirectoryChecker,
    localPath string,
    merger merger.Merger,
    projectName string,
    repo repo.Repo) *CreateAction {
    return &CreateAction{checker: checker, localPath: localPath, merger: merger, projectName: projectName, repo: repo}
}

func (action *CreateAction) Perform(output message.Printer) error {
    err := action.verifyLocalPath()
    if err != nil {
        return err
    }

    projectUrl, err := action.getProjectUrl()
    if err != nil {
        return err
    }

    output.Println(action.fetchMessage())

    _, err = action.repo.Clone(action.localPath, projectUrl)

    return err
}

func (action *CreateAction) verifyLocalPath() error {
    if action.checker.DirectoryExists(action.localPath) {
        return errors.New("project already exists")
    }
    return nil
}

func (action *CreateAction) getProjectUrl() (string, error) {
    mergedManifest := action.merger.MergedManifest()
    projectUrl := (*mergedManifest)[action.projectName]
    if len(projectUrl) == 0 {
        return "", errors.New("project not found")
    }
    return projectUrl, nil
}


func (action *CreateAction) fetchMessage() string {
    return fmt.Sprintf("Fetching %v...", action.projectName)
}