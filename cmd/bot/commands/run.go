package commands

import (
	"errors"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/config"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/claude_model"
	"github.com/urfave/cli/v2"
	"strings"
)

var RunCommand = &cli.Command{
	Name:  "run",
	Usage: "Run main server and worker",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "tg-bot-token",
			Required: true,
			Usage:    "Telegram bot token obtained from the BotFather.",
			EnvVars:  []string{"APP_TG_BOT_TOKEN"},
		},
		&cli.StringFlag{
			Name:     "anthropic-api-key",
			Required: true,
			Usage:    "Anthropic API key.",
			EnvVars:  []string{"APP_ANTHROPIC_API_KEY"},
		},
		&cli.StringFlag{
			Name:     "db",
			Required: true,
			Usage:    "Database URI. Supported: SQLite",
			EnvVars:  []string{"APP_DB_URI"},
		},
		&cli.StringFlag{
			Name:    "claude-model",
			Aliases: []string{"m"},
			Value:   string(claude_model.ClaudeDefaultModel),
			Usage:   "Claude model name",
			EnvVars: []string{"APP_CLAUDE_MODEL"},
		},
		&cli.StringFlag{
			Name:    "allowed",
			Usage:   "Allowed usernames to use bot, separated by comma. If not set, anyone can use bot",
			EnvVars: []string{"APP_ALLOWED_USERNAMES"},
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
		var allowedUsernames []string

		if cliCtx.String("allowed") != "" {
			allowedUsernames = strings.Split(cliCtx.String("allowed"), ",")
		}

		config := config.Config{
			TgBotToken:       cliCtx.String("tg-bot-token"),
			AnthropicApiKey:  cliCtx.String("anthropic-api-key"),
			ClaudeModel:      cliCtx.String("claude-model"),
			ProxyUrl:         cliCtx.String("proxy"),
			DBUri:            cliCtx.String("db"),
			AllowedUsernames: allowedUsernames,
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
