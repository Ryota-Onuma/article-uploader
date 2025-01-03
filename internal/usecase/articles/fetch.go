package articles

import (
	"context"
	"onion/internal/domain/model"
	"onion/internal/domain/repository"
	i "onion/internal/usecase/interfaces"
)

var _ i.FetchArticlesUsecase = (*FetchArticlesUsecaseImpl)(nil)

type FetchArticlesUsecaseImpl struct {
	i.BaseUsecase
	articleRepo repository.ArticleRepository
}

func NewFetchArticlesUsecase(base i.BaseUsecase, articleRepo repository.ArticleRepository) *FetchArticlesUsecaseImpl {
	return &FetchArticlesUsecaseImpl{
		base,
		articleRepo,
	}
}

func (a *FetchArticlesUsecaseImpl) Run(ctx context.Context) ([]model.Article, error) {
	articles, err := a.articleRepo.FetchArticles()
	if err != nil {
		return nil, a.WrapInternalServerError(ctx, err)
	}

	return articles, nil
}
