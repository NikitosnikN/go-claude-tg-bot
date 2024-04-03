package app

import (
	anthropic "github.com/NikitosnikN/go-anthropic-api"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/config"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/handlers"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/middleware"
	"gopkg.in/telebot.v3"
	"log"
	"time"
)

type App struct {
	config       *config.Config
	bot          *telebot.Bot
	claudeClient *anthropic.Client
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
	}
}

func (a *App) build() error {
	// build anthropic client
	client := anthropic.NewClient(a.config.AnthropicApiKey)

	if a.config.ProxyUrl != "" {
		err := client.SetProxy(a.config.ProxyUrl)

		if err != nil {
			return err
		}
	}

	a.claudeClient = client

	// build bot
	botSettings := telebot.Settings{
		Token:       a.config.TgBotToken,
		Poller:      &telebot.LongPoller{Timeout: 30 * time.Second},
		Synchronous: false,
	}

	bot, err := telebot.NewBot(botSettings)

	if err != nil {
		return err
	}
	a.bot = bot

	// add middlewares
	//a.bot.Use(middleware.Logger())

	a.bot.Use(middleware.VerboseLogger())
	if len(a.config.AllowedUsernames) != 0 {
		a.bot.Use(middleware.AllowList(a.config.AllowedUsernames...))
	}
	// build handlers
	a.bot.Handle(`/start`, handlers.StartHandler)
	a.bot.Handle(`/help`, handlers.HelpHandler)
	a.bot.Handle(`/echo`, handlers.EchoHandler)
	a.bot.Handle(telebot.OnText, handlers.TextMessageHandler(a.claudeClient, a.config.ClaudeModel))
	a.bot.Handle(telebot.OnPhoto, handlers.PhotoMessageHandler(a.claudeClient, a.config.ClaudeModel))
	return nil
}

func (a *App) Run() error {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("Bot crashed with panic ", r)
		}
	}()

	log.Println("Building bot...")

	err := a.build()

	if err != nil {
		return err
	}

	log.Println("Starting bot...")
	a.bot.Start()
	return nil
}
