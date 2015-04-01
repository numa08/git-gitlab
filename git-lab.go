package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-lab"
	app.Version = Version
	app.Usage = ""
	app.Author = "numa08"
	app.Email = "n511287@gmail.com"
    app.Commands = Commands
    
	app.Run(os.Args)
}


