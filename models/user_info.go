package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBotUserInfo struct {
	User User
}

// User represents a Telegram user or bot.
type User struct {
	TelegramId   int64
	FirstName    string
	LastName     string
	UserName     string
	LanguageCode string
}

func (u User) ToModels(user *tgbotapi.User) *User {
	u.TelegramId = user.ID
	u.UserName = user.UserName
	u.LastName = user.LastName
	u.FirstName = user.FirstName
	u.LanguageCode = user.LanguageCode
	return &u
}
