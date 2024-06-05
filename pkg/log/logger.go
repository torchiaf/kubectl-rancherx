package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

var ctx context.Context

const (
	levelTrace = slog.Level(-8)
	levelFatal = slog.Level(12)
)

var levelNames = map[slog.Leveler]string{
	levelTrace: "TRACE",
	levelFatal: "FATAL",
}

var LogFileName string
var LogLevel int

var logger *slog.Logger

func InitLogger() error {

	ctx = context.Background()

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

	logger = slog.New(slog.NewTextHandler(ioWriter, &slog.HandlerOptions{
		Level: slog.Level(8 - LogLevel), // We want reverse levels
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := levelNames[level]
				if !exists {
					levelLabel = level.String()
				}

				a.Value = slog.StringValue(levelLabel)
			}

			return a
		},
		AddSource: true,
	}))

	slog.SetDefault(logger)

	Info("Logs enabled")

	return nil
}

func Trace(msg string, args ...any) {
	logger.Log(ctx, levelTrace, msg, args...)
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

func Fatal(msg string, args ...any) {
	logger.Log(ctx, levelFatal, msg, args...)
}
