package action

import (
    "errors"
    "github.com/ncipollo/fnew/merger"
    "github.com/ncipollo/fnew/repo"
    "github.com/ncipollo/fnew/workspace"
    "io"
    "net/url"
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

func (action CreateAction) Perform(output io.Writer) error {
    err := action.verifyLocalPath()
    if err != nil {
        return err
    }

    projectUrl, err := action.getProjectUrl()
    if err != nil {
        return err
    }

    _, err = action.repo.Clone(action.localPath, projectUrl.String())

    return err
}

func (action CreateAction) verifyLocalPath() error {
    if action.checker.DirectoryExists(action.localPath) {
        return errors.New("project already exists")
    }
    return nil
}

func (action CreateAction) getProjectUrl() (url.URL, error) {
    mergedManifest := action.merger.MergedManifest()
    projectUrl := (*mergedManifest)[action.projectName]
    if len(projectUrl.String()) == 0 {
        return url.URL{}, errors.New("project not found")
    }
    return projectUrl, nil
}
