package main

import (
	"context"
	"github.com/NordeN37/telegram_bot/bl"
	"github.com/NordeN37/telegram_bot/models"
	"github.com/NordeN37/telegram_bot/utils/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog"
)

type IHandler interface {
	MakeServiceHandler(ctx context.Context, update tgbotapi.Update) (interface{}, error)
}

type handler struct {
	bl     *bl.BL
	logger *zerolog.Logger
}

func NewHandler(bl *bl.BL, logger *zerolog.Logger) IHandler {
	return &handler{bl: bl, logger: logger}
}

func (h handler) MakeServiceHandler(ctx context.Context, update tgbotapi.Update) (interface{}, error) {
	switch update.Message.Text {
	case "/start":
		_, err := h.bl.User.GetByTelegramId(ctx, update.Message.From.ID)
		if err != nil {
			if errors.Is(err, errors.ErrRecordingNotFound) {
				if err = h.bl.User.Create(ctx, (models.User{}).ToModels(update.Message.From)); err != nil {
					return nil, errors.Ctx().Just(err)
				}
			} else {
				return nil, errors.Ctx().Just(err)
			}
		}
	case "/ping":
		res := h.bl.TelegramBot.Ping()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, res)
		msg.ReplyToMessageID = update.Message.MessageID
		return msg, nil
	}
	return nil, nil
}
