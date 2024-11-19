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
			commit, err := initAction()
			println(commit)

			if err != nil {
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
