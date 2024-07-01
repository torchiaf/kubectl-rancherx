package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"

	"github.com/lmittmann/tint"
)

const (
	levelTrace = slog.Level(-8)
	levelDebug = slog.Level(-4)
	levelInfo  = slog.Level(0)
	levelWarn  = slog.Level(4)
	levelError = slog.Level(8)
	levelFatal = slog.Level(12)
)

const (
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	gray    = "\033[37m"
	white   = "\033[97m"
)

type customLevel struct {
	label string
	color string
}

var customLevels = map[slog.Leveler]customLevel{
	levelDebug: {
		label: "DEBUG",
		color: green,
	},
	levelInfo: {
		label: "INFO",
		color: blue,
	},
	levelWarn: {
		label: "WARN",
		color: yellow,
	},
	levelError: {
		label: "ERROR",
		color: red,
	},
	levelTrace: {
		label: "TRACE",
		color: magenta,
	},
	levelFatal: {
		label: "FATAL",
		color: blue,
	},
}

type fileWriter struct {
	w io.Writer
}

func (e fileWriter) Write(p []byte) (int, error) {

	re := regexp.MustCompile(`\033[[0-9;]*m`)

	// Trim colors from output
	newStr := re.ReplaceAllString(string(p), "")

	data := []byte(newStr)

	n, err := e.w.Write(data)

	if err != nil {
		return n, err
	}
	if n != len(data) {
		return n, io.ErrShortWrite
	}

	return len(p), nil
}

type LogConfig struct {
	LogLevel    int
	LogFileName string
}

var logger *slog.Logger

func InitLogger(ctx context.Context, cfg *LogConfig) error {

	var ioWriter io.Writer = os.Stdout

	if cfg.LogFileName != "" {
		logFile := fmt.Sprintf("%s.log", cfg.LogFileName)

		var f *os.File
		var err error

		if f, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err != nil {
			fmt.Println(err)
			return err
		}

		fileWriter := &fileWriter{w: f}

		ioWriter = io.MultiWriter(os.Stdout, fileWriter)
	}

	logger = slog.New(tint.NewHandler(ioWriter, &tint.Options{
		Level: slog.Level(8 - cfg.LogLevel), // We want reverse levels
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := customLevels[level]
				if !exists {
					levelLabel.label = level.String()
				}

				a.Value = slog.StringValue(levelLabel.color + levelLabel.label)
			}

			return a
		},
		AddSource: true,
	}))

	slog.SetDefault(logger)

	Info(ctx, "Logs enabled")

	return nil
}

func Trace(ctx context.Context, msg string, args ...any) {
	logger.Log(ctx, levelTrace, msg, args...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	logger.DebugContext(ctx, msg, args...)
}

func Info(ctx context.Context, msg string, args ...any) {
	logger.InfoContext(ctx, msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	logger.WarnContext(ctx, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	logger.ErrorContext(ctx, msg, args...)
}

func Fatal(ctx context.Context, msg string, args ...any) {
	logger.Log(ctx, levelFatal, msg, args...)
}
