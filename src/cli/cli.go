package cli

import (
	"cass/src/git"
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
			hasChanges, err := git.HasChanges()

			if err != nil {
				return err
			}

			if !hasChanges {
				println(color.New(color.FgRed).Sprint("⚠ You have no changes in the working tree.\n\n"))

				return nil
			}

			hasPendingChanges, err := git.HasPendingChanges()

			if err != nil {
				return err
			}

			if !hasPendingChanges {
				println(color.New(color.FgYellow).Sprintln("⚠ You have no changes in the stage area."))

				confirm, err := ConfirmStageChanges()

				if err != nil {
					return err
				}

				if confirm {
					err := git.StageAllChanges()

					if err != nil {
						return err
					}
				}
			}

			commit, err := initAction()

			if err != nil {
				return err
			}

			confirm, err := ConfirmCommitMessage(commit)

			if err != nil {
				return err
			}

			if confirm {
				err = git.ExecuteCommit(commit)
				if err != nil {
					return err
				}

				println(color.New(color.FgGreen).Println("\n\n✅ Commit successful"))

				// TODO: push commit
			} else {
				color.New(color.FgRed).Println("❌ Commit cancelled")
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
