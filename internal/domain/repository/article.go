package repository

import "onion/internal/domain/model"

type ArticleRepository interface {
	FetchArticles() ([]model.Article, error)
	CreateArticle(article model.Article) error
	FetchArticle(id model.ArticleID) (model.Article, error)
}
