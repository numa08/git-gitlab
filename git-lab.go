package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-lab"
	app.Version = Version
	app.Usage = "command line tool that wrap git in order to extend it with extra features and commands that make working with GitLab easier."
	app.Author = "numa08"
	app.Email = "n511287@gmail.com"
    app.Commands = Commands
    
	app.Run(os.Args)
}


