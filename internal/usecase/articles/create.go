package articles

import (
	"context"
	"fmt"
	"onion/internal/domain/model"
	"onion/internal/domain/repository"
	"onion/internal/domain/service/articles"
	i "onion/internal/usecase/interfaces"
	"time"
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
	// 今日の日付をPrefixにしてタイトルを生成する
	titleWithTimePrefix := fmt.Sprintf("【%s】 %s", time.Now().Format("2006-01-02"), title)
	article, err := a.createArticle.Run(ctx, titleWithTimePrefix, body)
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
