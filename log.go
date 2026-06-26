package wisp

import (
	"context"
	"log/slog"
)

func (l *logger) Debug(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelDebug, msg, args...)
}

func (l *logger) Info(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelInfo, msg, args...)
}

func (l *logger) Warn(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelWarn, msg, args...)
}

func (l *logger) Error(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelError, msg, args...)
}

func (l *logger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	argsCopy := make([]any, len(args))
	copy(argsCopy, args)

	for key, value := range l.fields {
		argsCopy = append(argsCopy, slog.Any(key, value))
	}

	requestID := RequestID(ctx)

	if requestID != "" {
		argsCopy = append(argsCopy, slog.String("request_id", requestID))
	}

	contextValues := getContextValues(ctx)

	for key, value := range contextValues {
		argsCopy = append(argsCopy, slog.Any(key, value))
	}

	l.l.Log(ctx, level, msg, argsCopy...)
}
