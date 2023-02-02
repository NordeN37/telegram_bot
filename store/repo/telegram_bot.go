package repo

import "gorm.io/gorm"

type ITelegramBot interface {
}

type TelegramBot struct {
	gorm.Model
}

type telegramBot struct {
	db *gorm.DB
}

func NewTelegramBot(db *gorm.DB) ITelegramBot {
	return &telegramBot{db: db}
}
