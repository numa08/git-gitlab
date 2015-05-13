package gitlab
import (
	"testing"
)

type mockconfig struct {
	host string
}

func (this *mockconfig) Host() (string, error) {
	return this.host, nil
}

func (this *mockconfig) ApiPath() (string, error) {
	return "api/v3", nil
}

func (this *mockconfig) Token() (string, error) {
	return "token", nil
}

func (this *mockconfig) Scheme() (string, error) {
	return "https", nil
}

func TestProject_NewProjectFromURL(t *testing.T) {
	config := &mockconfig{host: "gitlab.com",}
	project, e := NewProject(config, "numa08", "cookbook-gitlab")

	if e != nil {
		t.Errorf("create project error %s", e)
		return
	}
	if project  == nil {
		t.Error("create nil project")
		return
	}

	if project.Host != "gitlab.com" {
		t.Error("actual project host is " + project.Host + " but expected gitlab.com")
	}
	if project.Name != "cookbook-gitlab" {
		t.Error("actual project name is " + project.Name + " but expected cookbook-gitlab")
	}
	if project.NameSpace != "numa08" {
		t.Error("actual project name space is " + project.NameSpace + " but expected numa08")
	}
	if project.Protocol != "https" {
		t.Error("actual project protocol is " + project.Protocol + " but expected https")
	}
}

func TestWebURL(t *testing.T) {
	config := &mockconfig{host: "gitlab.com"}
	project, _ := NewProject(config, "numa08", "cookbook-gitlab")

	repoURL := project.WebURL("")
	if repoURL != "https://gitlab.com/numa08/cookbook-gitlab" {
		t.Error("actual repository url is " + repoURL + " but expected is https://gitlab.com/numa08/cookbook-gitlab")
	}

	wikiURL := project.WebURL("wiki")
	if wikiURL != "https://gitlab.com/numa08/cookbook-gitlab/wiki" {
		t.Error("actual wiki url is " + wikiURL + " but expected is https://gitlab.com/numa08/cookbook-gitlab/wiki")
	}
}