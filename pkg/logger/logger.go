package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Gsc23/e-commerce-api/e-commerce-api/pkg/config"
	"github.com/lmittmann/tint"
	slogmulti "github.com/samber/slog-multi"
)

type Logger interface {
	Debug(context.Context, string, ...any)
	Info(context.Context, string, ...any)
	Warn(context.Context, string, ...any)
	Error(context.Context, string, ...any)
}

type LoggerConfig struct {
	Level       slog.Level
	Console     bool
	Color       bool
	JSON        bool
	EnableTrace bool
}

type LoggerFactory struct {
	base    *slog.Logger
	mu      sync.RWMutex
	loggers map[string]*SlogAdapter
}

type SlogAdapter struct {
	log *slog.Logger
}

func newLoggerConfig(cfg config.Config) (*LoggerConfig, error) {
	logger := &LoggerConfig{
		Color:       cfg.LoggerColors(),
		EnableTrace: cfg.LoggerTrace(),
	}

	level := cfg.LoggerLevel()
	if err := logger.handleLogLevel(level); err != nil {
		return nil, err
	}

	return logger, nil
}

func (l *LoggerConfig) handleLogLevel(level string) error {
	switch strings.ToUpper(level) {
	case "DEBUG":
		l.Level = slog.LevelDebug
	case "INFO":
		l.Level = slog.LevelInfo
	case "WARN":
		l.Level = slog.LevelWarn
	case "ERROR":
		l.Level = slog.LevelError
	default:
		return fmt.Errorf("Unknow level name: %s", level)
	}

	return nil
}

func (cfg *LoggerConfig) newHandler(w io.Writer) slog.Handler {
	opts := &slog.HandlerOptions{
		Level:     cfg.Level,
		AddSource: cfg.EnableTrace,
	}

	var handlers []slog.Handler

	if cfg.Color {
		handlers = append(handlers,
			tint.NewHandler(w, &tint.Options{
				Level:      cfg.Level,
				TimeFormat: time.DateTime,
			}),
		)
	} else {
		handlers = append(handlers, slog.NewTextHandler(w, opts))
	}

	return slogmulti.Fanout(handlers...)
}

func NewLoggerFactory(cfg *LoggerConfig) *LoggerFactory {
	handler := cfg.newHandler(os.Stdout)

	return &LoggerFactory{
		base:    slog.New(handler),
		loggers: make(map[string]*SlogAdapter),
	}
}

func (f *LoggerFactory) NewLoggerNamed(name string) Logger {
	f.mu.RLock()
	if l, ok := f.loggers[name]; ok {
		f.mu.RUnlock()
		return l
	}
	f.mu.RUnlock()

	f.mu.Lock()
	defer f.mu.Unlock()

	l := f.base.With("component", name)
	log := &SlogAdapter{log: l}

	f.loggers[name] = log
	return log
}

func (l *SlogAdapter) Info(ctx context.Context, msg string, args ...any) {
	l.log.InfoContext(ctx, msg, args...)
}

func (l *SlogAdapter) Debug(ctx context.Context, msg string, args ...any) {
	l.log.DebugContext(ctx, msg, args...)
}

func (l *SlogAdapter) Warn(ctx context.Context, msg string, args ...any) {
	l.log.WarnContext(ctx, msg, args...)
}

func (l *SlogAdapter) Error(ctx context.Context, msg string, args ...any) {
	l.log.ErrorContext(ctx, msg, args...)
}
