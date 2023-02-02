package telegram_bot

import (
	"context"
	"github.com/NordeN37/telegram_bot/models"
	"github.com/NordeN37/telegram_bot/store"
	"github.com/rs/zerolog"
)

type IUser interface {
	GetByTelegramId(ctx context.Context, telegramId int) (*models.User, error)
	Create(ctx context.Context, data *models.User) error
}

type user struct {
	log *zerolog.Logger
	db  *store.DBRepo
}

func NewUser(db *store.DBRepo, log *zerolog.Logger) IUser {
	return &user{db: db, log: log}
}

func (u user) GetByTelegramId(ctx context.Context, telegramId int) (*models.User, error) {
	return u.db.User.GetByTelegramId(ctx, telegramId)
}

func (u user) Create(ctx context.Context, data *models.User) error {
	return u.db.User.Create(ctx, data)
}
