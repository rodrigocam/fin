package main

import (
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rodrigocam/fin/pkg/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "fin",
		Usage: "Your personal finance watcher",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Start the watcher server",
				Action: func(cliCtx *cli.Context) error {
					server := server.Default()
					emailPrompt := &survey.Input{
						Message: "Enter your email address:",
					}

					email := ""
					if err := survey.AskOne(emailPrompt, &email); err != nil {
						return err
					}
					prompt := &survey.Password{
						Message: "Enter your password:",
					}

					password := ""
					if err := survey.AskOne(prompt, &password); err != nil {
						return err
					}

					return server.Start(email, password)
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
