package config

type Config struct {
	TgBotToken       string
	AnthropicApiKey  string
	ProxyUrl         string
	AllowedUsernames []string
}

func (c *Config) IsFilled() bool {
	return c.TgBotToken != "" && c.AnthropicApiKey != ""
}
