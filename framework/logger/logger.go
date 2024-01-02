package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Logger *logrus.Logger
}

type ILogger interface {
	Info(string)
	Error(string)
	Fatal(string)
}

func NewLogger() ILogger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	if os.Getenv("environment") == "development" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	return &Logger{Logger: logger}
}

func (l *Logger) Info(v string) {
	l.Logger.Info(v)
}

func (l *Logger) Error(v string) {
	l.Logger.Error(v)
}

func (l *Logger) Fatal(v string) {
	l.Logger.Fatal(v)
}
