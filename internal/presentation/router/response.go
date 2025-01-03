package router

import "onion/internal/domain/model"

type Article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

func convertArticleDomainModelToArticleDataModel(article model.Article) Article {
	return Article{
		ID:        article.ID.Value(),
		Title:     article.Title.Value(),
		Body:      article.Body.Value(),
		CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
