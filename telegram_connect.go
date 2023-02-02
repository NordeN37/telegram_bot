package main

import (
	"context"
	"github.com/NordeN37/telegram_bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog"
	"log"
)

func telegramBotConnect(handler IHandler, logger *zerolog.Logger, cfg *config.Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.MyAwesomeBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	logger.Info().Msgf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logger.Error().Err(err)
	}

	for update := range updates {
		ctx := context.Background()
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			resHandle, err := handler.MakeServiceHandler(ctx, update)
			if err != nil {
				logger.Error().Str("UserName", update.Message.From.UserName).Str("Message", update.Message.Text).Err(err)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Request execution error. Try later.")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				continue
			}

			switch resHandle.(type) {
			case tgbotapi.Chattable:
				bot.Send(resHandle.(tgbotapi.Chattable))
			}
		}
	}
}
