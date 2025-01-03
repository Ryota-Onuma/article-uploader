package usecase

import (
	"context"
	"onion/internal/domain/model"
)

type FetchArticlesUsecase interface {
	BaseUsecase
	Run(ctx context.Context) ([]model.Article, error)
}

type CreateArticleUsecase interface {
	BaseUsecase
	Run(ctx context.Context, title, body string) (model.Article, error)
}
