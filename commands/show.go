package commands
import (
	"github.com/codegangsta/cli"
	"github.com/numa08/git-gitlab/config"
	"github.com/numa08/git-gitlab/gitlab"
	"fmt"
	"github.com/numa08/git-gitlab/utils"
	"os/exec"
)
/**
	$ git lab browse
	> open https://GITLAB_HOST/CURENT_REPO

	$ git lab browse -- issue
	> open https://GITLAB_HOST/CURRENT_REPO/issues

	$ git lab brose gitlab/gitlab
	> open https://GITLAB_HOST/gitlab/gitlab

	$ git lab brose gl
	> open https://GITLAB_HOST/YOUR_LOGIN/gl

	$ git lab brose gl wiki
	> open https://GITLAB_HOST/YOUR_LOGIN/gl/wiki
*/
func Show(c *cli.Context) error {
	var (
		dest	string
		subpage	string
		path	string
		err 	error
		project *gitlab.Project
		bracnh 	*gitlab.Branch
	)

	dest = c.Args().First()

	if dest == "--" {
		dest = ""
	}

	subpage = c.Args().Get(1)

	config, err := config.ConfigForCurrentDir()
	local := gitlab.LocalReppo()

	if dest != "" {
		project, _ := gitlab.NewProject(config, "", dest)
		branch = local.MasterBranch()
	} else if subpage != "" && subpage != "commits" && subpage != "tree" && subpage != "blob" && subpage != "settings" {
		project, err = local.MainProject()
		branch = local.MasterBranch()
	} else {
		branch, err := local.CurrentBranch()
		if err != nil {
			branch = local.MasterBranch()
		}
	}

	if project == nil {
		err := fmt.Errorf(usage of command)
		return err
	}

	if subpage == "commits" {
		path = //create branch commits URL
	} else if subpage == "tree" || subpage == "" {
		if !branch.isDefaultBranch() {
			path = //create tree URL
		}
	} else {
		path = subpage
	}

	pageURL := project.WebURL(path)
	launcher, err := utils.BrowserLauncher()
	if err != nil {
		return err
	}

	run := append(launcher, pageURL)
	exec.Command(run[0], run[1:]...)
	return nil
}