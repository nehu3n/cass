package cli

import (
	"log"
	"os"
	"github.com/urfave/cli/v2"
)
func Init() {
	app := &cli.App{
        Name:  "cass",
        Usage: "Advanced commits assistant.",
        Action: func(*cli.Context) error {
           return initAction()
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}