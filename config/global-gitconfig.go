package config
import (
    "gopkg.in/libgit2/git2go.v22"
    "fmt"
    "strings"
    "os/exec"
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
    path, e := git.ConfigFindGlobal()
    global, e := git.OpenOndisk(nil, path)
    host, e := global.LookupString("gitlab.host")
    return host, e
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