package app

import (
	anthropic "github.com/NikitosnikN/go-anthropic-api"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/adapter/sql_store"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/commands"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/queries"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/config"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/middleware"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/ports/handlers"
	"gopkg.in/telebot.v3"
	"log"
	"time"
)

type Commands struct {
	AddUser            *commands.AddUserHandler
	AddDialog          *commands.AddDialogHandler
	AddMessageToDialog *commands.AddMessagesToDialogHandler
}

type Queries struct {
	GetUserByID       *queries.GetUserHandler
	GetLatestDialog   *queries.GetLatestDialogHandler
	GetDialogMessages *queries.GetDialogMessagesHandler
	GetPhoto          *queries.GetPhoto
}

type App struct {
	db           *sql_store.SQLStore
	config       *config.Config
	bot          *telebot.Bot
	claudeClient *anthropic.Client

	commands *Commands
	queries  *Queries
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
	}
}

func (a *App) build() error {
	// create DB instance
	db, err := sql_store.NewSQLStore(a.config.DBUri)

	if err != nil {
		return err
	}

	err = db.AutomigrateAll()

	if err != nil {
		return err
	}

	a.db = db

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
	bot, err := telebot.NewBot(telebot.Settings{
		Token:       a.config.TgBotToken,
		Poller:      &telebot.LongPoller{Timeout: 30 * time.Second},
		Synchronous: false,
	})

	if err != nil {
		return err
	}
	a.bot = bot

	// build commands
	a.commands = &Commands{
		AddUser:            commands.NewAddUserHandler(),
		AddDialog:          commands.NewAddDialogHandler(),
		AddMessageToDialog: commands.NewAddMessagesToDialogHandler(),
	}

	// build queries
	a.queries = &Queries{
		GetUserByID:       queries.NewGetUserHandler(),
		GetLatestDialog:   queries.NewGetLatestDialogHandler(),
		GetDialogMessages: queries.NewGetDialogMessagesHandler(),
		GetPhoto:          queries.NewGetPhoto(a.bot),
	}

	// build bot middlewares
	//a.bot.Use(middleware.Logger())

	a.bot.Use(middleware.ErrorHandler())
	a.bot.Use(middleware.VerboseLogger())
	if len(a.config.AllowedUsernames) != 0 {
		a.bot.Use(middleware.AllowList(a.config.AllowedUsernames...))
	}
	a.bot.Use(
		middleware.Auth(
			a.db.NewTx,
			a.queries.GetUserByID,
			a.commands.AddUser,
		),
	)

	// build bot handlers
	a.bot.Handle(`/start`, handlers.StartHandler)
	a.bot.Handle(`/help`, handlers.HelpHandler)
	a.bot.Handle(`/echo`, handlers.EchoHandler)
	a.bot.Handle(`/new`, handlers.NewHandler(
		a.db.NewTx,
		a.commands.AddDialog,
	))
	a.bot.Handle(telebot.OnText, handlers.TextMessageHandler(
		a.claudeClient,
		a.db.NewTx,
		a.queries.GetLatestDialog,
		a.queries.GetDialogMessages,
		a.queries.GetPhoto,
		a.commands.AddDialog,
		a.commands.AddMessageToDialog,
	))
	a.bot.Handle(telebot.OnPhoto, handlers.PhotoMessageHandler(
		a.claudeClient,
		a.db.NewTx,
		a.queries.GetLatestDialog,
		a.queries.GetDialogMessages,
		a.queries.GetPhoto,
		a.commands.AddDialog,
		a.commands.AddMessageToDialog,
	))
	return nil
}

func (a *App) Run() error {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("Bot crashed with panic ", r)
		}
	}()

	log.Println("Building app...")

	err := a.build()

	if err != nil {
		return err
	}

	log.Println("Starting bot...")
	a.bot.Start()
	return nil
}
