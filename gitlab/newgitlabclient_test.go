package gitlab

import (
    "testing"
    "errors"
)

type GitNilConfig  struct{}

func (this *GitNilConfig) Token() (string, error) {
    return "", errors.New("gitlab.token is not found")
}

func (this *GitNilConfig) Host() (string , error) {
    return "", errors.New("gitlab.host is not found")
}

func (this *GitNilConfig) ApiPath() (string, error) {
    return "", errors.New("gitlab.apipath is not found")
}

func TestNewGitlabClient(t *testing.T) {
    config := new(GitNilConfig)
    client , _ := NewGitLabClient(config)
    var expected *GitLabClient = nil
    if client != expected {
        t.Error("client will be nil")
    }
}
