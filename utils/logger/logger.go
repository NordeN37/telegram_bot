package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

func New(levelStr string) *zerolog.Logger {
	var level zerolog.Level
	switch levelStr {
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "error":
		level = zerolog.ErrorLevel
	default:
		level = zerolog.DebugLevel
	}
	log := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(level)
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf(" %s:%d ", file, line)
	}
	return &log
}
