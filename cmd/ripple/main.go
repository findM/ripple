package main

import (
	"github.com/codegangsta/cli"
	"os"
	"github.com/bmbstack/ripple/cmd/ripple/scripts"
)

func main() {
	app := cli.NewApp()
	app.Name = "ripple"
	app.Usage = "Command line tool to managing your Ripple application"
	app.Version = "0.0.1"
	app.Author = "wangmingjob"
	app.Email = "wangmingjob@icloud.com"
	app.Commands = scripts.Commands()
	app.Run(os.Args)
}

