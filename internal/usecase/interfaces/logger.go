package usecase

import "context"

type Attribute struct {
	Key   string
	Value any
}

type TraceIDKey struct{}

type Logger interface {
	Error(ctx context.Context, message string, detail string, attributes ...Attribute)
	Warn(ctx context.Context, message string, detail string, attributes ...Attribute)
	Info(ctx context.Context, message string, attributes ...Attribute)
	Debug(ctx context.Context, message string, attributes ...Attribute)
	CreateAttribute(key string, value any) Attribute
}
