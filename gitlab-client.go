package main

import (
    "github.com/plouc/go-gitlab-client"
    "strings"
    "gopkg.in/libgit2/git2go.v22"
)

type GitLabClient struct {
    GitLab *gogitlab.Gitlab
}

func NewGitLabClient(config GitConfig) (*GitLabClient, error) {
    host , e := config.Host()
    if e != nil {
        return nil, e
    }
    path , e := config.ApiPath()
    if e != nil {
        return nil, e
    }
    token , e := config.Token()
    if e != nil {
        return nil, e
    }
    client := GitLabClient{
        GitLab: gogitlab.NewGitlab(host, path, token),
    }
    return &client, e
}

func (this *GitLabClient)clone (remote string, path string) (string, error) {
// search remote repository
    projectID := strings.Replace(remote, "/", "%2F", -1)
    project , e := this.GitLab.Project(projectID)
// get remote URL
    remoteURL := project.SshRepoUrl
// clone
    var dest string
    if len(path) < 0 {
        dest = project.Name
    } else {
        dest = path
    }


    repo, e := git.Clone(remoteURL, dest, nil)
    res := repo.Path()
    return res, e
}