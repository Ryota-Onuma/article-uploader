package router

import "onion/internal/domain/model"

type Article struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func convertArticleDomainModelToArticleDataModel(article model.Article) Article {
	return Article{
		ID:    article.ID.Value(),
		Title: article.Title.Value(),
		Body:  article.Body.Value(),
	}
}
