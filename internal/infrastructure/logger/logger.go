package logger

import (
	"context"
	"fmt"
	"log/slog"
	usecase "onion/internal/usecase/interfaces"
	"os"

	"github.com/m-mizutani/clog"
)

type Logger struct {
	logger *slog.Logger
}

func NewJSONHandler() slog.Handler {
	return slog.NewJSONHandler(os.Stdout, nil)
}

func NewDebugHandler() slog.Handler {
	return clog.New(
		clog.WithColor(true),
		clog.WithLevel(slog.LevelDebug),
		clog.WithTimeFmt("2006-01-02 15:04:05"),
		clog.WithWriter(os.Stdout),
		clog.WithSource(true),
		clog.WithPrinter(clog.PrettyPrinter),
	)
}

func New(handler slog.Handler) *Logger {
	logger := slog.New(handler)
	return &Logger{
		logger: logger,
	}
}

// 以下インターフェイスに準拠
var _ usecase.Logger = (*Logger)(nil)

func (l *Logger) Error(ctx context.Context, message string, detail string, attributes ...usecase.Attribute) {
	attrs := l.createAttributes(ctx, attributes...)
	attrs = append(attrs, slog.String("content", detail))
	l.logger.LogAttrs(
		ctx,
		slog.LevelError,
		fmt.Sprintf("!!! %s !!!", message),
		slog.Group("detail",
			attrs...,
		),
	)
}

func (l *Logger) Warn(ctx context.Context, message string, detail string, attributes ...usecase.Attribute) {
	attrs := l.createAttributes(ctx, attributes...)
	attrs = append(attrs, slog.String("content", detail))
	l.logger.LogAttrs(
		ctx,
		slog.LevelWarn,
		message,
		slog.Group("detail",
			attrs...,
		),
	)
}

func (l *Logger) Info(ctx context.Context, message string, attributes ...usecase.Attribute) {
	attrs := l.createAttributes(ctx, attributes...)

	l.logger.LogAttrs(
		ctx,
		slog.LevelInfo,
		message,
		slog.Group("detail",
			attrs...,
		),
	)
}

func (l *Logger) Debug(ctx context.Context, message string, attributes ...usecase.Attribute) {
	attrs := l.createAttributes(ctx, attributes...)
	l.logger.LogAttrs(
		ctx,
		slog.LevelDebug,
		message,
		slog.Group("detail",
			attrs...,
		),
	)
}

func (l *Logger) createAttributes(ctx context.Context, attributes ...usecase.Attribute) []any {
	attrs := make([]any, 0, len(attributes)+1)
	for _, attr := range attributes {
		attrs = append(attrs, slog.Any(attr.Key, attr.Value))
	}

	traceID := ctx.Value(usecase.TraceIDKey{})
	if traceID != nil && traceID.(string) != "" {
		attrs = append(attrs, slog.String("trace_id", traceID.(string)))
	}

	return attrs
}

func (l *Logger) CreateAttribute(key string, value any) usecase.Attribute {
	return usecase.Attribute{
		Key:   key,
		Value: value,
	}
}
