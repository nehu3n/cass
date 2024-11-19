package cli

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func Init() {
	app := &cli.App{
		Name:  "cass",
		Usage: "Advanced commits assistant.",
		Action: func(*cli.Context) error {
			commit, err := initAction()

			if err != nil {
				return err
			}

			confirm, err := ConfirmCommitMessage(commit)

			if err != nil {
				return err
			}

			if confirm {
				println(commit)
				// TODO: push the commit
			} else {
				red := color.New(color.FgRed).SprintFunc()
				println(red("[❌] Commit cancelled"))
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
