package config

type Config struct {
	TgBotToken       string
	AnthropicApiKey  string
	ClaudeModel      string
	ProxyUrl         string
	AllowedUsernames []string
}

func (c *Config) IsFilled() bool {
	return c.TgBotToken != "" && c.AnthropicApiKey != "" && c.ClaudeModel != ""
}
