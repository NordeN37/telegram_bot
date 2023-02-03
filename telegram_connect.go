package main

import (
	"context"
	"github.com/NordeN37/telegram_bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
	"log"
)

func telegramBotConnect(handler IHandler, logger *zerolog.Logger, cfg *config.Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.MyAwesomeBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = cfg.BotDebug
	logger.Info().Msgf("Authorized on account %s", bot.Self.UserName)

	addMyCommands(bot, logger)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		ctx := context.Background()
		if update.Message != nil { // If we got a message
			logger.Info().Str("message", update.Message.Text).Interface("user", update.Message.From).Send()
			resHandle, err := handler.MakeServiceHandler(ctx, update)
			if err != nil {
				logger.Error().Str("UserName", update.Message.From.UserName).Str("Message", update.Message.Text).Err(err).Send()
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

func addMyCommands(bot *tgbotapi.BotAPI, logger *zerolog.Logger) {
	myCommandsConfig := tgbotapi.SetMyCommandsConfig{Commands: []tgbotapi.BotCommand{
		{Command: "/ping", Description: "return pong"},
	}, LanguageCode: "en"}

	resp, err := bot.Request(myCommandsConfig)
	if err != nil {
		logger.Error().Str("request", "SetMyCommandsConfig").Interface("myCommandsConfig", myCommandsConfig).Err(err).Send()
	}
	logger.Info().Str("request", "SetMyCommandsConfig").Interface("myCommandsConfig", myCommandsConfig).Interface("resp", resp).Send()
}
