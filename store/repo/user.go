package repo

import (
	"context"
	"github.com/NordeN37/telegram_bot/models"
	"github.com/NordeN37/telegram_bot/utils/errors"
	"gorm.io/gorm"
)

type IUser interface {
	GetByTelegramId(ctx context.Context, telegramId int) (*models.User, error)
	Create(ctx context.Context, data *models.User) error
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) IUser {
	return &user{db: db}
}

type User struct {
	gorm.Model
	TelegramId   int
	FirstName    string
	LastName     string
	UserName     string
	LanguageCode string
}

func (u *User) fromModels(data *models.User) error {
	if data == nil {
		return errors.Ctx().New("data is null")
	}
	u.LastName = data.LastName
	u.UserName = data.UserName
	u.FirstName = data.FirstName
	u.TelegramId = data.ID
	u.LanguageCode = data.LanguageCode
	return nil
}

func (u *User) toModels() *models.User {
	res := &models.User{
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		UserName:     u.UserName,
		ID:           u.TelegramId,
		LanguageCode: u.LanguageCode,
	}
	return res
}

func (u user) GetByTelegramId(ctx context.Context, telegramId int) (*models.User, error) {
	var us User
	if err := u.db.WithContext(ctx).First(&us, "telegram_id = ?", telegramId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Ctx().Just(errors.ErrRecordingNotFound)
		}
		return nil, errors.Ctx().Just(err)
	}
	return us.toModels(), nil
}

func (u user) Create(ctx context.Context, data *models.User) error {
	create := &User{}
	if err := create.fromModels(data); err != nil {
		return errors.Ctx().Just(err)
	}
	if err := u.db.WithContext(ctx).Create(create).Error; err != nil {
		return errors.Ctx().Just(err)
	}
	return nil
}
