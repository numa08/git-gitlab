package main

import (
    "github.com/plouc/go-gitlab-client"
)

type GitLabClient struct {
    GitLab *gogitlab.Gitlab
}

func NewGitLabClient(config GitConfig) (*GitLabClient, error) {
    host , e := config.Host()
    path , e := config.ApiPath()
    token , e := config.Token()
    client := GitLabClient{
        GitLab: gogitlab.NewGitlab(host, path, token),
    }
    return &client, e
}