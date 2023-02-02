package store

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

func NewLogger(lg *zerolog.Logger, config gormLogger.Config) gormLogger.Interface {
	return &logger{
		lg:     lg,
		Config: config,
	}
}

type logger struct {
	lg *zerolog.Logger
	gormLogger.Config
}

func (l *logger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l logger) Info(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		l.lg.Info().Str("filename", utils.FileWithLineNum()).Msg(fmt.Sprintf(msg, data...))
	}
}

func (l logger) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		l.lg.Warn().Str("filename", utils.FileWithLineNum()).Msg(fmt.Sprintf(msg, data...))
	}
}

func (l logger) Error(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		l.lg.Error().Str("filename", utils.FileWithLineNum()).Msg(fmt.Sprintf(msg, data...))
	}
}

func (l logger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		strElapced := fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
		switch {
		case err != nil && l.LogLevel >= gormLogger.Error:
			var loggerFunc func() *zerolog.Event
			if errors.Is(err, gorm.ErrRecordNotFound) {
				loggerFunc = l.lg.Debug
			} else {
				loggerFunc = l.lg.Error
			}

			sql, rows := fc()
			loggerFunc().
				Err(err).
				Str("rows", rowsToString(rows)).
				Str("sql", sql).
				Str("elapsedTime", strElapced).
				Str("filename", utils.FileWithLineNum()).
				Msg("Trace SQL")
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormLogger.Warn:
			sql, rows := fc()
			l.lg.Warn().Str("rows", rowsToString(rows)).
				Str("sql", sql).
				Str("elapsedTime", strElapced).
				Str("filename", utils.FileWithLineNum()).
				Str("slowThreshold", l.SlowThreshold.String()).Msg("Trace SQL")
		case l.LogLevel >= gormLogger.Info:
			sql, rows := fc()
			l.lg.Warn().Str("rows", rowsToString(rows)).
				Str("sql", sql).
				Str("elapsedTime", strElapced).
				Str("filename", utils.FileWithLineNum()).
				Msg("Trace SQL")
		}
	}
}

func rowsToString(rows int64) string {
	if rows == -1 {
		return "-"
	} else {
		return strconv.FormatInt(rows, 10)
	}
}
