package workspace

import (
	"os/user"
	"path"
	"fmt"
	"os"
)

type Workspace struct {
	BasePath         string
	DirectoryCreator DirectoryCreator
}

func New(basePath string, creator DirectoryCreator) Workspace {
	return Workspace{BasePath: basePath, DirectoryCreator: creator}
}

func (workspace Workspace ) Setup() error {
	return workspace.DirectoryCreator.CreateDirectory(workspace.BasePath)
}

func (workspace Workspace ) ConfigPath() string {
	return path.Join(workspace.BasePath, "config.json")
}

func (workspace Workspace ) ManifestsPath() string {
	return path.Join(workspace.BasePath, "manifests")
}

func OSDirectoryCreator() DirectoryCreator {
	return osDirectoryCreator{}
}

func Directory() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(fmt.Sprintf("Unable to get current user! Error: %s", err))
	}
	return path.Join(currentUser.HomeDir, ".fnew")
}

type DirectoryCreator interface {
	CreateDirectory(dir string) error
}

type osDirectoryCreator struct{}

func (osDirectoryCreator) CreateDirectory(dir string) error {
	return os.Mkdir(dir, 777)
}
