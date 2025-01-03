package usecase

import i "onion/internal/usecase/interfaces"

var _ i.BaseUsecase = (*BaseUsecaseImpl)(nil)

type BaseUsecaseImpl struct {
	i.ErrorProvider
	i.Logger
}

func NewBaseUsecase(logger i.Logger, err i.ErrorProvider) *BaseUsecaseImpl {
	return &BaseUsecaseImpl{
		ErrorProvider: err,
		Logger:        logger,
	}
}
