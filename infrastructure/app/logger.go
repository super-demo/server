package app

import (
	"log/slog"
	"os"
)

var (
	SLog *slog.Logger
)

func InitLogger() {
	SLog = NewLogger()
}

func NewLogger() *slog.Logger {
	newLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return newLogger
}
