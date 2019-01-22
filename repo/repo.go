package repo

import (
    "golang.org/x/crypto/ssh"
    "gopkg.in/src-d/go-git.v4"
    "io/ioutil"
    "os"
    "gopkg.in/src-d/go-git.v4/plumbing/protocol/packp/sideband"
    "gopkg.in/src-d/go-git.v4/plumbing/transport"
    go_git_ssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
    "os/user"
    "path/filepath"
)

type Repo interface {
    Clone(localPath string, repoUrl string) (*git.Repository, error)
    Delete(localPath string) error
    Init(localPath string) (*git.Repository, error)
    Open(localPath string) (*git.Repository, error)
    Pull(repository *git.Repository) error
}

type GitRepo struct {
    verbose bool
}

func New(verbose bool) Repo {
    return &GitRepo{verbose: verbose}
}

func (repo *GitRepo) Clone(localPath string, repoUrl string) (*git.Repository, error) {
    auth, err := repo.auth(repoUrl)
    if err != nil {
        return nil, err
    }
    return git.PlainClone(localPath, false, &git.CloneOptions{
        Auth:              auth,
        URL:               repoUrl,
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
        Progress:          repo.progress(),
    })
}

func (repo *GitRepo) Delete(localPath string) error {
    fullPath := filepath.Join(localPath, ".git")
    return os.RemoveAll(fullPath)
}

func (repo *GitRepo) Init(localPath string) (*git.Repository, error) {
    return git.PlainInit(localPath, false)
}

func (GitRepo) Open(localPath string) (*git.Repository, error) {
    return git.PlainOpen(localPath)
}

func (repo *GitRepo) Pull(repository *git.Repository) error {
    remote, err := repository.Remote("origin")
    if err != nil {
        return err
    }

    auth, err := repo.auth(remote.Config().URLs[0])
    if err != nil {
        return err
    }

    tree, err := repository.Worktree()
    if err != nil {
        return err
    }
    return tree.Pull(&git.PullOptions{Auth: auth, RemoteName: "origin"})
}

func (repo *GitRepo) auth(repoUrl string) (transport.AuthMethod, error) {
    ep, err := transport.NewEndpoint(repoUrl)
    if err != nil {
        return nil, err
    }
    switch ep.Protocol {
    case "ssh":
        auth, err := repo.defaultSSHAuth()
        if err != nil {
            auth, err = go_git_ssh.NewSSHAgentAuth(ep.User)
            if err != nil {
                return nil, err
            }
        }
        return auth, nil
    default:
        return nil, nil
    }
}

func (repo *GitRepo) defaultSSHAuth() (transport.AuthMethod, error) {
    usr, err := user.Current()
    if err != nil {
        return nil, err
    }

    sshKeyPath := usr.HomeDir + "/.ssh/id_rsa"

    sshKey, err := ioutil.ReadFile(sshKeyPath)
    if err != nil {
        return nil, err
    }

    signer, err := ssh.ParsePrivateKey([]byte(sshKey))
    if err != nil {
        return nil, err
    }

    auth := &go_git_ssh.PublicKeys{User: "git", Signer: signer}
    return auth, nil
}

func (repo *GitRepo) progress() sideband.Progress {
    if repo.verbose {
        return os.Stdout
    } else {
        return nil
    }
}
