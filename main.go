package main

import (
	"github.com/NordeN37/telegram_bot/bl"
	"github.com/NordeN37/telegram_bot/config"
	"github.com/NordeN37/telegram_bot/store"
	"github.com/NordeN37/telegram_bot/utils/logger"
)

func main() {
	cfg := config.New()
	log := logger.New(cfg.LogLevel)

	dbConnect, err := store.OpenMasterDBConnection(log, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to postgres")
	}

	repo := store.NewDBRepo(dbConnect)

	serverBl := bl.NewBL(repo, log)

	handler := NewHandler(serverBl, log)
	telegramBotConnect(handler, log, cfg)
}
