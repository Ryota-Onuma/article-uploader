package error

import (
	"context"

	i "onion/internal/usecase/interfaces"

	"github.com/morikuni/failure/v2"
)

var _ i.ErrorProvider = (*err)(nil)

type err struct{}

func New() *err {
	return &err{}
}

const (
	InternalServerError = "InternalServerError"
	Unauthorized        = "Unauthorized"
	NotFound            = "NotFound"
	Forbidden           = "Forbidden"
)

func (e *err) NewInternalServerError(ctx context.Context, message string) error {
	return failure.New(InternalServerError, failure.Message(message))
}

func (e *err) NewUnauthorizedError(ctx context.Context, message string) error {
	return failure.New(Unauthorized, failure.Message(message))
}

func (e *err) NewNotFoundError(ctx context.Context, message string) error {
	return failure.New(NotFound, failure.Message(message))
}

func (e *err) NewForbiddenError(ctx context.Context, message string) error {
	return failure.New(Forbidden, failure.Message(message))
}

func (e *err) WrapInternalServerError(ctx context.Context, err error) error {
	return failure.Wrap(err, failure.Context{"type": InternalServerError})
}

func (e *err) WrapUnauthorizedError(ctx context.Context, err error) error {
	return failure.Wrap(err, failure.Context{"type": Unauthorized})
}

func (e *err) WrapNotFoundError(ctx context.Context, err error) error {
	return failure.Wrap(err, failure.Context{"type": NotFound})
}

func (e *err) WrapForbiddenError(ctx context.Context, err error) error {
	return failure.Wrap(err, failure.Context{"type": Forbidden})
}
