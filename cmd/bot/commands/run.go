package commands

import (
	"errors"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/config"
	"github.com/urfave/cli/v2"
)

var RunCommand = &cli.Command{
	Name:  "run",
	Usage: "Run main server and worker",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "tg-bot-token",
			Aliases:  []string{"t"},
			Required: true,
			Usage:    "Telegram bot token",
			EnvVars:  []string{"APP_TG_BOT_TOKEN"},
		},
		&cli.StringFlag{
			Name:     "anthropic-api-key",
			Aliases:  []string{"a"},
			Required: true,
			Usage:    "Anthropic API key",
			EnvVars:  []string{"APP_ANTHROPIC_API_KEY"},
		},
		&cli.StringFlag{
			Name:    "proxy",
			Aliases: []string{"p"},
			Value:   "",
			Usage:   "Proxy URL",
			EnvVars: []string{"APP_PROXY"},
		},
	},
	Action: func(cliCtx *cli.Context) error {
		config := config.Config{
			TgBotToken:      cliCtx.String("tg-bot-token"),
			AnthropicApiKey: cliCtx.String("anthropic-api-key"),
			ProxyUrl:        cliCtx.String("proxy"),
		}

		if config.IsFilled() == false {
			return errors.New("config is not filled")
		}

		app := app.NewApp(&config)

		if err := app.Run(); err != nil {
			return err
		}
		return nil
	},
}
