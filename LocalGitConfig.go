package main
import (
    "gopkg.in/libgit2/git2go.v22"
    "os/exec"
    "strings"
    "fmt"
)

type LocalGitConfig struct {
    Config *git.Config
}

func NewLocalGitConfig() (*LocalGitConfig, error) {
    path, e := exec.Command("git", "rev-parse", "--show-toplevel").Output()
    p := strings.TrimRight(string(path), "\n")
    repo, e := git.OpenRepository(p)
    conf, e := repo.Config()
    config := LocalGitConfig{
        Config: conf,
    }
    return &config, e
}

func (this *LocalGitConfig) Host() (string, error) {
    return this.Config.LookupString("gitlab.url")
}

func (this *LocalGitConfig) Token() (string, error) {
    return this.Config.LookupString("gitlab.token")
}

func (this *LocalGitConfig) ApiPath() (string, error) {
    host, e := this.Host()
    apiPath := fmt.Sprintf("%s/api/v3", host)
    return apiPath, e
}