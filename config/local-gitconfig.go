package config
import (
    "gopkg.in/libgit2/git2go.v22"
    "os/exec"
    "strings"
    "fmt"
    "net/url"
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
    url, e := this.url()
    host := url.Host
    return host, e
}

func (this *LocalGitConfig) Scheme() (string, error) {
    url, e := this.url()
    scheme := url.Scheme
    return scheme, e
}

func (this *LocalGitConfig) Token() (string, error) {
    return this.Config.LookupString("gitlab.token")
}

func (this *LocalGitConfig) ApiPath() (string, error) {
    apiPath := fmt.Sprintf("/api/v3")
    return apiPath, nil
}

func (this *LocalGitConfig) url() (*url.URL,error) {
    part, e := this.Config.LookupString("gitlab.url")
    url, e := url.Parse(part)
    return url, e
}