package log

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

const (
	LevelTrace = slog.Level(-5)
)

var LogFileName string
var LogLevel int

var logger *slog.Logger

func InitLogger() error {

	var ioWriter io.Writer = os.Stdout

	if LogFileName != "" {
		logFile := fmt.Sprintf("%s.log", LogFileName)

		var f *os.File
		var err error

		if f, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
			fmt.Println(err)
			return err
		}

		ioWriter = io.MultiWriter(os.Stdout, f)
	}

	logger = slog.New(slog.NewJSONHandler(ioWriter, &slog.HandlerOptions{
		Level:     slog.Level(8 - LogLevel), // We want reverse levels
		AddSource: true,
	}))

	slog.SetDefault(logger)

	logger.Info("Logs enabled")

	return nil
}

func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}
