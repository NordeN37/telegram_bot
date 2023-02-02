package telegram_bot

import (
	"github.com/NordeN37/telegram_bot/store"
	"github.com/rs/zerolog"
)

type ITelegramBot interface {
	Ping() string
}

type telegramBot struct {
	log *zerolog.Logger
	db  *store.DBRepo
}

func NewTelegramBot(db *store.DBRepo, log *zerolog.Logger) ITelegramBot {
	return &telegramBot{db: db, log: log}
}

func (t telegramBot) Ping() string {
	return "pong"
}
