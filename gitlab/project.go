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

func NewProject(config config.GitConfig, remote string) (*Project, error) {
	parts := strings.Split(remote, ":")
	if len(parts) < 2 {
		err := fmt.Errorf("Invalid remote address %s", remote)
		return nil ,err
	}

	name := strings.TrimSuffix(parts[len(parts) - 1], ".git")
	nameSpace := parts[len(parts) - 2]
	host, err := config.Host()
	if err != nil {
		return nil, err
	}
	scheme, err := config.Scheme()
	if err != nil {
		return nil, err
	}
	project := newProject(nameSpace, name, host, scheme)
	return project, nil
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