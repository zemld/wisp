package wisp

import (
	"context"
	"maps"
)

func WithField(ctx context.Context, key string, value any) context.Context {
	contextValues := getContextValues(ctx)
	if contextValues == nil {
		contextValues = make(map[string]any)
	}

	contextValuesCopy := maps.Clone(contextValues)

	contextValuesCopy[key] = value

	ctx = context.WithValue(ctx, generalLoggingKey{}, contextValuesCopy)

	return ctx
}

func WithFields(ctx context.Context, fields map[string]any) context.Context {
	contextValues := getContextValues(ctx)
	if contextValues == nil {
		contextValues = make(map[string]any)
	}

	contextValuesCopy := maps.Clone(contextValues)

	maps.Copy(contextValuesCopy, fields)

	ctx = context.WithValue(ctx, generalLoggingKey{}, contextValuesCopy)

	return ctx
}

func getContextValues(ctx context.Context) map[string]any {
	values, ok := ctx.Value(generalLoggingKey{}).(map[string]any)
	if !ok {
		return nil
	}

	return values
}

func (l *logger) WithField(key string, value any) *logger {
	fields := maps.Clone(l.fields)
	fields[key] = value

	return &logger{
		l:      l.l,
		fields: fields,
	}
}
