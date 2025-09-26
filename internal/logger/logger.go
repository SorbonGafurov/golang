package logger

import (
	"log/slog"

	"github.com/natefinch/lumberjack"
)

func NewLogger(filename string) *slog.Logger {
	file := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    5,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}
	return slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}
