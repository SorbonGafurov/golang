package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type AppLogger struct {
	*logrus.Logger
}

func NewLogger() *AppLogger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(logrus.InfoLevel)
	return &AppLogger{log}
}
