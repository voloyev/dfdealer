package main

import (
	"fmt"
	"log"
	"os"

	"voloyev/dfdealer/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "dfdealer",
		Usage: "Deal with your dotfiles for U",
		Action: func(*cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "check",
				Aliases: []string{"ch"},
				Usage:   "check existing links",
				Action:  commands.Check,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
