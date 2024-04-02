package main

import (
	"github.com/NikitosnikN/go-claude-tg-bot/cmd/bot/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:                 "claude-bot",
		Usage:                "Claude Telegram Bot",
		EnableBashCompletion: true,
		Suggest:              true,
		Commands: []*cli.Command{
			commands.RunCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
