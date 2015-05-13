package gitlab
import (
	"fmt"
	"strings"
	"github.com/numa08/git-gitlab/config"
)

type Project struct {
	Name string
	Host string
	NameSpace string
	Protocol string
}

func (this *Project) String() string {
	return fmt.Sprintf("%s/%s", this.NameSpace, this.Name)
}

func (this *Project) WebURL(path string) string {
	repoURL := fmt.Sprintf("%s://%s/%s/%s", this.Protocol,this.Host,  this.NameSpace, this.Name)

	if path == "" {
		return repoURL
	} else {
		return fmt.Sprintf("%s/%s", repoURL, path)
	}
}

func NewProject(config config.GitConfig, owner string, name string) (*Project, error) {
	host ,err := config.Host()
	if err != nil {
		return nil, err
	}
	protocol, err := config.Scheme()
	if err != nil {
		return nil, err
	}
	return newProject(owner, name, host, protocol), nil
}


func newProject(nameSpace, name, host, protocol string) *Project {
	if strings.Contains(nameSpace, "/") {
		result := strings.SplitN(nameSpace, "/", 2)
		nameSpace = result[0]
		if name == "" {
			name = result[1]
		}
	} else if strings.Contains(name, "/") {
		result := strings.SplitN(name, "/", 2)
		if nameSpace == "" {
			nameSpace = result[0]
		}
		name = result[1]
	}

	if protocol != "http" && protocol != "https" {
		protocol =  ""
	}

	if protocol == "" {
		protocol = "https"
	}

	return &Project{
		Name:name,
		Host:host,
		NameSpace:nameSpace,
		Protocol:protocol,
	}
}