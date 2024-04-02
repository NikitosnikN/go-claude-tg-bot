# Go Claude Telegram Bot

A Telegram bot that utilizes the Claude AI language model to provide interactive conversations and assistance to users.

## Introduction

The Go Claude Telegram Bot is a Telegram bot built using the Go programming language and the Claude AI language model.
It enables users to engage in interactive conversations and receive assistance from the powerful Claude AI directly
through their Telegram chats.

Currently only long-polling way of retrieving new messages is supported.

## Features

- Natural language conversation with the Claude AI
- Assistance with various tasks such as writing, analysis, question answering, math, and coding
- Seamless integration with Telegram for easy access and interaction
- Customizable configuration options
- Lightweight and efficient implementation in Go

## Installation

1. Clone the repository:

```shell
git clone https://github.com/NikitosnikN/go-claude-tg-bot.git
```

2. Install the required dependencies:

```shell
cd go-claude-tg-bot
go mod download
```

3. Build the bot:

```shell
go build
```

4. Run the bot:

```shell
go run cmd/bot/main.go run
```

## Configuration

It is possible to configure application via 2 ways:

* via CLI arguments. Run `./bot run --help` or `go run cmd/bot/main.go run --help` for more info.
* via Env variables. See the table with variables below.

| Variable Name           | Description                                                                                  |
|-------------------------|----------------------------------------------------------------------------------------------|
| `APP_TG_BOT_TOKEN`      | Telegram bot token obtained from the BotFather..                                             |
| `APP_ANTHROPIC_API_KEY` | Anthropic API key.                                                                           |
| `APP_ALLOWED_USERNAMES` | (Optional) Allowed usernames to use bot, separated by comma. If not set, anyone can use bot. |
| `APP_PROXY`             | (Optional) Proxy URL. Used for Anthropic API only.                                           |

## Usage

1. Run bot:

```shell
./bot run
```

2. Open your Telegram app and search for your bot using its username.

3. Start a conversation with the bot and send messages to interact with the Claude AI.

4. Use the available commands to perform specific actions or get help:

    - `/start`: Start the conversation with the bot
    - `/help`: Get a list of available commands and their descriptions
    - Send message, send only photo or send photo and message to get completion from Claude AI

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please create an issue or submit
a pull request on the [GitHub repository](https://github.com/NikitosnikN/go-claude-tg-bot).

## License

This project is licensed under the [MIT License](./LICENCE).