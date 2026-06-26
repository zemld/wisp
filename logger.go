package wisp

import "log/slog"

var _ Logger = (*logger)(nil)

type logger struct {
	l      *slog.Logger
	fields map[string]any
}

func NewLogger(l *slog.Logger) *logger {
	return &logger{
		l: l,
	}
}
