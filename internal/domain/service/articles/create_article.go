package articles

import (
	"context"
	"onion/internal/domain/model"

	"github.com/google/uuid"
)

type CreateArticleService interface {
	Run(ctx context.Context, title, body string) (*model.Article, error)
}

type CreateArticleServiceImpl struct{}

func NewCreateArticleService() *CreateArticleServiceImpl {
	return &CreateArticleServiceImpl{}
}

func (a *CreateArticleServiceImpl) Run(ctx context.Context, title, body string) (*model.Article, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	id, err := model.NewArticleID(uuid.String())
	if err != nil {
		return nil, err
	}

	articleTitle, err := model.NewArticleTitle(title)
	if err != nil {
		return nil, err
	}

	articleBody, err := model.NewArticleBody(body)
	if err != nil {
		return nil, err
	}

	return model.NewArticle(id, articleTitle, articleBody)
}
