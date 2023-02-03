package store

import (
	"errors"
	"github.com/NordeN37/telegram_bot/config"
	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	gormLogger "gorm.io/gorm/logger"
	"time"

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func OpenMasterDBConnection(logger *zerolog.Logger, settings *config.Config) (db *gorm.DB, err error) {
	if db, err = connectWithGORMRetry(logger, settings); err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(3 * time.Minute)

	if !settings.AutoMigrate {
		return db, nil
	}

	if err := applyAutoMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}

func connectWithGORMRetry(logger *zerolog.Logger, settings *config.Config) (*gorm.DB, error) {
	ticker := time.NewTicker(1 * time.Nanosecond)
	timeout := time.After(15 * time.Minute)
	seconds := 1
	try := 0
	for {
		select {
		case <-ticker.C:
			try++
			ticker.Stop()
			client, err := connectWithGORM(logger, settings.DBPath, settings)
			if err != nil {
				logger.Warn().Err(err).Msgf("не удалось установить соединение с SQLite, попытка № %d", try)

				ticker = time.NewTicker(time.Duration(seconds) * time.Second)
				seconds *= 2
				if seconds > 60 {
					seconds = 60
				}
				continue
			}

			logger.Debug().Msg("соединение с SQLite успешно установлено")
			return client, nil
		case <-timeout:
			return nil, errors.New("SQLite: connection timeout")
		}
	}
}

func connectWithGORM(logger *zerolog.Logger, dbPath string, settings *config.Config) (*gorm.DB, error) {
	logLevel := gormLogger.Warn
	if settings.TraceSQLCommands {
		logLevel = gormLogger.Info
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: NewLogger(logger, gormLogger.Config{
			// временной зазор определения медленных запросов SQL
			SlowThreshold: time.Duration(600) * time.Second,
			LogLevel:      logLevel,
			Colorful:      false,
		}),
		AllowGlobalUpdate: true,
	})
	return db, err
}
