package store

import (
	"github.com/NordeN37/telegram_bot/store/repo"
	"github.com/NordeN37/telegram_bot/store/transaction"
	"github.com/NordeN37/telegram_bot/utils/errors"
	"gorm.io/gorm"
)

// DBRepo - интерфейс работы с базой данных
type DBRepo struct {
	DB          *gorm.DB
	TelegramBot repo.ITelegramBot
	User        repo.IUser
}

// NewDBRepo - конструктор интерфейса работы с базой данных
func NewDBRepo(dbHandler *gorm.DB) *DBRepo {
	return &DBRepo{
		DB:          dbHandler,
		TelegramBot: repo.NewTelegramBot(dbHandler),
		User:        repo.NewUser(dbHandler),
	}
}

// applyAutoMigrations - регистрация авто миграции схемы бд из моделей
func applyAutoMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&repo.User{},
		&repo.TelegramBot{},
	)
}

// WithTransaction - обертка заворачивающая выполнение операция GORM в транзакцию
func (ds *DBRepo) WithTransaction(handler func(tx transaction.ITransaction) error) error {
	tx := ds.DB.Begin()
	if err := handler(tx); err != nil {
		if errRollback := tx.Rollback().Error; errRollback != nil {
			return errors.And(errors.New("ошибка отката изменений транзакции"), errRollback)
		}
		return err
	}
	return tx.Commit().Error
}
