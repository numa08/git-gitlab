package main

import (
    "gopkg.in/libgit2/git2go.v22"
    "os/exec"
)

type GitLab struct {
    Repository git.Repository
}

func NewGitLab() (*GitLab, error) {
    path, error := exec.Command("git", "rev-parse", "--show-toplevel").Output()
    repository, error := git.OpenRepository(path)
    gitlab := &GitLab{
        Repository:repository,
    }
//    config, error := repository.Config()
//    url, error := config.LookupString("gitlab.url")
//    token, error := config.LookupString("gitlab.token")
    return gitlab, error
}