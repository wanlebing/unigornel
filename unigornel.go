package main

import (
	"os"

	"github.ugent.be/unigornel/build"
	"github.ugent.be/unigornel/env"

	"github.com/urfave/cli"
)

func app() *cli.App {
	app := cli.NewApp()
	app.Name = "unigornel"
	app.Usage = "build unikernels for Go"
	app.HideVersion = true
	app.Commands = []cli.Command{
		env.Env(),
		build.Build(),
		build.CompileGo(),
		build.CompileOS(),
	}
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr
	return app
}

func main() {
	app().Run(os.Args)
}
