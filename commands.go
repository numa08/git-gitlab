package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandGit_clone,
    commandGit_merge_request,
    commandGit_merge,
    commandGit_show,
    
}


var commandGit_clone = cli.Command{
	Name:  "git-clone",
	Usage: "",
	Description: `
`,
	Action: doGit_clone,
}

var commandGit_merge_request = cli.Command{
	Name:  "git-merge-request",
	Usage: "",
	Description: `
`,
	Action: doGit_merge_request,
}

var commandGit_merge = cli.Command{
	Name:  "git-merge",
	Usage: "",
	Description: `
`,
	Action: doGit_merge,
}

var commandGit_show = cli.Command{
	Name:  "git-show",
	Usage: "",
	Description: `
`,
	Action: doGit_show,
}


func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func doGit_clone(c *cli.Context) {
}

func doGit_merge_request(c *cli.Context) {
}

func doGit_merge(c *cli.Context) {
}

func doGit_show(c *cli.Context) {
}


