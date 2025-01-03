package usecase

import "context"

type ErrorProvider interface {
	NewInternalServerError(ctx context.Context, message string) error
	NewUnauthorizedError(ctx context.Context, message string) error
	NewNotFoundError(ctx context.Context, message string) error
	NewForbiddenError(ctx context.Context, message string) error
	WrapInternalServerError(ctx context.Context, err error) error
	WrapUnauthorizedError(ctx context.Context, err error) error
	WrapNotFoundError(ctx context.Context, err error) error
	WrapForbiddenError(ctx context.Context, err error) error
}
