package config
import (
    "gopkg.in/libgit2/git2go.v22"
    "fmt"
    "strings"
    "os/exec"
    "net/url"
)

func ConfigForCurrentDir() (GitConfig, error) {
    path, e := exec.Command("git", "rev-parse", "--show-toplevel").Output()
    p := strings.TrimRight(string(path), "\n")
    if len(p) < 1 {
        // here is not git directory
        return NewGlobalGitConfig(), e
    }

    return NewLocalGitConfig()
}

type GlobalGitConfig struct {}

func NewGlobalGitConfig() (*GlobalGitConfig) {
    config := new(GlobalGitConfig)
    return config
}

func (this *GlobalGitConfig) Host() (string, error) {
    url, e := this.url()
    host := url.Host
    return host, e
}

func (this *GlobalGitConfig) Scheme() (string, error) {
    url, e := this.url()
    scheme := url.Scheme
    return scheme, e
}

func (this *GlobalGitConfig) Token() (string, error) {
    path, e := git.ConfigFindGlobal()
    global, e := git.OpenOndisk(nil, path)
    token , e := global.LookupString("gitlab.token")
    return token, e
}

func (this *GlobalGitConfig) ApiPath() (string, error) {
    apiPath := fmt.Sprintf("/api/v3")
    return apiPath, nil
}

func (this *GlobalGitConfig) url() (*url.URL, error) {
    path, e := git.ConfigFindGlobal()
    global, e := git.OpenOndisk(nil, path)
    part, e := global.LookupString("gitlab.url")
    url, e := url.Parse(part)
    return url, e
}