package bl

import (
	"github.com/NordeN37/telegram_bot/bl/telegram_bot"
	"github.com/NordeN37/telegram_bot/store"
	"github.com/rs/zerolog"
)

type BL struct {
	TelegramBot telegram_bot.ITelegramBot
	User        telegram_bot.IUser
}

func NewBL(r *store.DBRepo, logger *zerolog.Logger) *BL {
	return &BL{
		TelegramBot: telegram_bot.NewTelegramBot(r, logger),
		User:        telegram_bot.NewUser(r, logger),
	}
}
