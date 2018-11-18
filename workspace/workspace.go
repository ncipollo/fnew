package workspace

import (
    "fmt"
    "github.com/ncipollo/fnew/manifest"
    "os"
    "os/user"
    "path"
)

type Workspace struct {
    basePath         string
    directoryChecker DirectoryChecker
    directoryCreator DirectoryCreator
}

func New(basePath string, checker DirectoryChecker, creator DirectoryCreator) Workspace {
    return Workspace{basePath: basePath, directoryChecker: checker, directoryCreator: creator}
}

func (workspace Workspace) Setup() error {
    return workspace.directoryCreator.CreateDirectory(workspace.ManifestsPath())
}

func (workspace Workspace) ConfigPath() string {
    return path.Join(workspace.basePath, "config.json")
}

func (workspace Workspace) ManifestsPath() string {
    return path.Join(workspace.basePath, "manifests")
}

func (workspace Workspace) ConfigManifestRepoPath() string {
    return path.Join(workspace.ManifestsPath(), manifest.ConfigDirectory)
}

func (workspace Workspace) ConfigManifestRepoExists() bool {
    return workspace.directoryChecker.DirectoryExists(workspace.ConfigManifestRepoPath())
}

func (workspace Workspace) DefaultManifestRepoPath() string {
    return path.Join(workspace.ManifestsPath(), manifest.DefaultDirectory)
}

func (workspace Workspace) DefaultManifestRepoExists() bool {
    return workspace.directoryChecker.DirectoryExists(workspace.DefaultManifestRepoPath())
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

func OSDirectoryCreator() DirectoryCreator {
    return osDirectoryCreator{}
}

func (osDirectoryCreator) CreateDirectory(dir string) error {
    return os.MkdirAll(dir, 0777)
}

type DirectoryChecker interface {
    DirectoryExists(dir string) bool
}

type osDirectoryChecker struct{}

func OSDirectoryChecker() DirectoryChecker {
    return osDirectoryChecker{}
}

func (osDirectoryChecker) DirectoryExists(dir string) bool {
    _, err := os.Stat(dir)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false
    }
    return true
}
