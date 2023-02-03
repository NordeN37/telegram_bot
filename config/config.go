package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	LogLevel          string `envconfig:"LOG_LEVEL" default:"debug"`
	DBPath            string `envconfig:"DB_PATH" default:"./telegram_bot.db"`
	AutoMigrate       bool   `envconfig:"AUTO_MIGRATE" default:"true"`
	MyAwesomeBotToken string `envconfig:"MY_AWESOME_BOT_TOKEN"`
	TraceSQLCommands  bool   `envconfig:"TRACE_SQL_COMMANDS" default:"false"`
	BotDebug          bool   `envconfig:"BOT_DEBUG" default:"false"`
}

func New() *Config {
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("failed to load envconfig, err: %s", err)
	}
	return &cfg
}
