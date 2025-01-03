package articles

import (
	"context"
	"onion/internal/domain/model"
	"onion/internal/domain/repository"
	"onion/internal/domain/service/articles"
	i "onion/internal/usecase/interfaces"
)

var _ i.CreateArticleUsecase = (*CreateArticleUsecaseImpl)(nil)

type CreateArticleUsecaseImpl struct {
	i.BaseUsecase
	createArticle articles.CreateArticleService
	articleRepo   repository.ArticleRepository
}

func NewCreateArticleUsecase(base i.BaseUsecase, createArticle articles.CreateArticleService, articleRepo repository.ArticleRepository) *CreateArticleUsecaseImpl {
	return &CreateArticleUsecaseImpl{
		base,
		createArticle,
		articleRepo,
	}
}

func (a *CreateArticleUsecaseImpl) Run(ctx context.Context, title, body string) (model.Article, error) {
	article, err := a.createArticle.Run(ctx, title, body)
	if err != nil {
		return model.Article{}, a.WrapForbiddenError(ctx, err)
	}

	if err := a.articleRepo.CreateArticle(*article); err != nil {
		return model.Article{}, a.WrapInternalServerError(ctx, err)
	}

	createdArticle, err := a.articleRepo.FetchArticle(article.ID)
	if err != nil {
		return model.Article{}, a.WrapInternalServerError(ctx, err)
	}

	return createdArticle, nil
}
