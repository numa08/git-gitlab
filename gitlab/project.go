package gitlab
import (
	"fmt"
	"strings"
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

func newProject(name, host, nameSpace, protocol string) *Project {
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